package main

//
//import (
//	"context"
//	"fmt"
//	"sync"
//	"time"
//)
//
//// LimitedResource 限制资源实例数量的结构
//type LimitedResource struct {
//	capacity     int        // 资源实例的最大数量
//	idle         int        // 当前可用资源的数量
//	resourceCond *sync.Cond // 条件变量，用于同步
//}
//
//// NewLimitedResource 创建一个新的LimitedResource实例
//func NewLimitedResource(capacity int) *LimitedResource {
//	return &LimitedResource{
//		capacity:     capacity,
//		idle:         capacity,
//		resourceCond: sync.NewCond(&sync.Mutex{}),
//	}
//}
//
//// Borrow 借用资源，如果资源不可用则等待，超时则返回错误
//func (lr *LimitedResource) Borrow(timeout time.Duration) (interface{}, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	defer cancel()
//
//	lr.resourceCond.L.Lock()
//	defer lr.resourceCond.L.Unlock()
//
//	for lr.idle == 0 {
//		if err := lr.resourceCond.WaitWithContext(ctx); err != nil {
//			return nil, fmt.Errorf("resource acquisition timed out after %s", timeout)
//		}
//	}
//
//	// 减少可用资源计数
//	lr.idle--
//
//	// 模拟资源创建
//	resource := fmt.Sprintf("Resource#%d", lr.capacity-lr.idle)
//
//	return resource, nil
//}
//
//// Return 归还资源
//func (lr *LimitedResource) Return(resource interface{}) {
//	lr.resourceCond.L.Lock()         // 锁定
//	defer lr.resourceCond.L.Unlock() // 解锁
//
//	// 增加可用资源计数
//	lr.idle++
//	// 通知等待的goroutine
//	lr.resourceCond.Signal()
//}
//
//// Destroy 销毁资源池
//func (lr *LimitedResource) Destroy() error {
//	lr.resourceCond.L.Lock()
//	defer lr.resourceCond.L.Unlock()
//
//	// 确保所有资源都被归还
//	if lr.idle != lr.capacity {
//		return fmt.Errorf("cannot destroy resource pool with borrowed resources")
//	}
//
//	// 销毁资源池的逻辑
//	// 这里可以加入实际的销毁逻辑，例如关闭资源连接等
//
//	return nil
//}
//
//func main() {
//	// 创建一个容量为3的LimitedResource实例
//	lr := NewLimitedResource(3)
//
//	// 定义资源借用的超时时间
//	timeout := 2 * time.Second
//
//	// 模拟并发借用资源
//	var wg sync.WaitGroup
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(id int) {
//			defer wg.Done()
//			resource, err := lr.Borrow(timeout)
//			if err != nil {
//				fmt.Printf("Attempt %d failed: %s\n", id, err)
//				return
//			}
//			fmt.Printf("Attempt %d acquired resource: %v\n", id, resource)
//			// 模拟资源使用一段时间后归还
//			time.Sleep(1 * time.Second)
//			lr.Return(resource)
//			fmt.Printf("Attempt %d returned resource: %v\n", id, resource)
//		}(i)
//	}
//
//	// 等待所有借用goroutine完成
//	wg.Wait()
//
//	// 销毁资源池
//	if err := lr.Destroy(); err != nil {
//		fmt.Println("Error destroying resource pool:", err)
//	} else {
//		fmt.Println("Resource pool destroyed.")
//	}
//}
