package main

import (
	"fmt"
	"sync"
	"time"
)

// LimitedResource 限制资源实例数量的结构
type LimitedResource struct {
	capacity     int        // 资源实例的最大数量
	idle         int        // 当前可用资源的数量
	resourceCond *sync.Cond // 条件变量，用于同步
}

// NewLimitedResource 创建一个新的LimitedResource实例
func NewLimitedResource(capacity int) *LimitedResource {
	return &LimitedResource{
		capacity:     capacity,
		idle:         capacity,
		resourceCond: sync.NewCond(&sync.Mutex{}),
	}
}

// Borrow 借用资源，如果资源不可用则等待，超时则返回错误
func (lr *LimitedResource) Borrow(timeout time.Duration) (interface{}, error) {
	timeoutChan := time.After(timeout)
	lr.resourceCond.L.Lock()
	defer lr.resourceCond.L.Unlock()

	for lr.idle == 0 {
		select {
		case <-timeoutChan:
			return nil, fmt.Errorf("resource acquisition timed out after %s", timeout)
		default:
			lr.resourceCond.Wait()
		}
	}

	// 减少可用资源计数
	lr.idle--

	// 模拟资源创建
	resource := fmt.Sprintf("Resource#%d", lr.capacity-lr.idle)

	// 启动一个 goroutine 来监控资源使用时间
	go func(res interface{}) {
		select {
		case <-time.After(timeout): // 超时后自动归还资源
			fmt.Printf("Resource %v timed out and will be returned and destroyed.\n", res)
			lr.Return(res)
			lr.DestroyResource(res)
		}
	}(resource)

	return resource, nil
}

// Return 归还资源
func (lr *LimitedResource) Return(resource interface{}) {
	lr.resourceCond.L.Lock()
	defer lr.resourceCond.L.Unlock()

	// 增加可用资源计数
	lr.idle++
	// 通知等待的goroutine
	lr.resourceCond.Signal()
}

// Destroy 销毁资源池
func (lr *LimitedResource) Destroy() error {
	lr.resourceCond.L.Lock()
	defer lr.resourceCond.L.Unlock()

	// 确保所有资源都被归还
	if lr.idle != lr.capacity {
		return fmt.Errorf("cannot destroy resource pool with borrowed resources")
	}

	// 销毁资源池的逻辑
	fmt.Println("All resources are returned. Destroying resource pool...")

	return nil
}

// DestroyResource 销毁单个资源
func (lr *LimitedResource) DestroyResource(resource interface{}) {
	// 实际的资源销毁逻辑应该在这里实现
	fmt.Printf("Resource %v is destroyed.\n", resource)
}

func main() {
	// 创建一个容量为3的LimitedResource实例
	lr := NewLimitedResource(3)

	// 定义资源借用的超时时间
	timeout := 2 * time.Second

	// 模拟并发借用资源
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			resource, err := lr.Borrow(timeout)
			if err != nil {
				fmt.Printf("Attempt %d failed: %s\n", id, err)
				return
			}
			fmt.Printf("Attempt %d acquired resource: %v\n", id, resource)
			// 模拟资源使用一段时间后归还
			time.Sleep(1 * time.Second)
			lr.Return(resource)
			fmt.Printf("Attempt %d returned resource: %v\n", id, resource)
		}(i)
	}

	// 等待所有借用资源的 goroutine 完成
	wg.Wait()

	// 销毁资源池
	if err := lr.Destroy(); err != nil {
		fmt.Println("Failed to destroy resource pool:", err)
	} else {
		fmt.Println("Resource pool destroyed.")
	}
}
