package utils

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// LimitedResource 通用的限制资源实例数量的结构
type LimitedResource[T any] struct {
	capacity     int              // 资源实例的最大数量
	idle         int              // 当前可用资源的数量
	resourceCond *sync.Cond       // 条件变量，用于同步
	createFunc   func() T         // 资源创建函数
	destroyFunc  func(resource T) // 资源销毁函数
}

// NewLimitedResource 创建一个新的LimitedResource实例
func NewLimitedResource[T any](capacity int, createFunc func() T, destroyFunc func(resource T)) *LimitedResource[T] {
	return &LimitedResource[T]{
		capacity:     capacity,
		idle:         capacity,
		resourceCond: sync.NewCond(&sync.Mutex{}),
		createFunc:   createFunc,
		destroyFunc:  destroyFunc,
	}
}

// Borrow 借用资源，如果资源不可用则等待，超时则返回错误
func (lr *LimitedResource[T]) Borrow(ctx context.Context, timeout time.Duration) (T, error) {
	var zero T // 声明一个零值，用于在超时时返回
	timeoutChan := time.After(timeout)

	lr.resourceCond.L.Lock()
	defer lr.resourceCond.L.Unlock()

	for lr.idle == 0 {
		select {
		case <-timeoutChan:
			return zero, fmt.Errorf("resource acquisition timed out after %s", timeout)
		default:
			lr.resourceCond.Wait()
		}
	}

	// 减少可用资源计数
	lr.idle--

	// 创建资源
	resource := lr.createFunc()

	// 启动一个 goroutine 来监控资源使用时间
	go func(res T) {
		select {
		case <-time.After(timeout): // 超时后自动归还资源
			fmt.Printf("Resource %v timed out and will be returned and destroyed.\n", res)
			lr.Return(res)
			lr.DestroyResource(res)
		case <-ctx.Done(): // 手动取消
			return
		}
	}(resource)

	return resource, nil
}

// Return 归还资源
func (lr *LimitedResource[T]) Return(resource T) {
	lr.resourceCond.L.Lock()
	defer lr.resourceCond.L.Unlock()

	// 增加可用资源计数
	lr.idle++
	// 通知等待的goroutine
	lr.resourceCond.Signal()
}

// Destroy 销毁资源池
func (lr *LimitedResource[T]) Destroy() error {
	lr.resourceCond.L.Lock()
	defer lr.resourceCond.L.Unlock()

	// 确保所有资源都被归还
	if lr.idle != lr.capacity {
		return fmt.Errorf("cannot destroy resource pool with borrowed resources")
	}

	fmt.Println("All resources are returned. Destroying resource pool...")
	return nil
}

// DestroyResource 销毁单个资源
func (lr *LimitedResource[T]) DestroyResource(resource T) {
	// 实际的资源销毁逻辑由外部提供的 destroyFunc 实现
	lr.destroyFunc(resource)
}
