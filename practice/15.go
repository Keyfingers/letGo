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
	done := make(chan bool, 1) // 用于超时控制

	lr.resourceCond.L.Lock() // 锁定
	for lr.idle == 0 {
		// 如果没有可用资源，等待
		lr.resourceCond.Wait()
	}
	// 减少可用资源计数
	lr.idle--
	lr.resourceCond.L.Unlock() // 解锁

	// 模拟资源创建
	resource := fmt.Sprintf("Resource#%d", lr.capacity-lr.idle)

	// 超时控制
	go func() {
		time.Sleep(timeout)
		done <- true
	}()

	// 检查是否超时
	if <-done {
		// 超时，归还资源
		lr.Return(resource)
		return nil, fmt.Errorf("resource acquisition timed out after %s", timeout)
	}

	return resource, nil
}

// Return 归还资源
func (lr *LimitedResource) Return(resource interface{}) {
	lr.resourceCond.L.Lock()         // 锁定
	defer lr.resourceCond.L.Unlock() // 解锁

	// 销毁资源
	// 资源实际的销毁逻辑应该在这里实现

	// 增加可用资源计数
	lr.idle++
	// 通知等待的goroutine
	lr.resourceCond.Signal()
}

// Destroy 销毁资源池
func (lr *LimitedResource) Destroy() {
	lr.resourceCond.L.Lock()
	defer lr.resourceCond.L.Unlock()

	// 确保所有资源都被归还
	if lr.idle != lr.capacity {
		panic("cannot destroy resource pool with borrowed resources")
	}

	// 销毁资源池的逻辑
	// ...
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

	// 等待所有借用goroutine完成
	wg.Wait()

	// 销毁资源池
	lr.Destroy()
	fmt.Println("Resource pool destroyed.")
}
