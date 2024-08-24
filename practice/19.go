package main

import (
	"context"
	"fmt"
	"letGo/utils"
	"sync"
	"time"
)

// 资源的创建函数
func createStringResource() string {
	return "Hello, World!"
}

// 资源的销毁函数
func destroyStringResource(resource string) {
	fmt.Printf("Destroying resource: %s\n", resource)
}

func main() {
	// 创建一个容量为3的LimitedResource实例，处理字符串资源
	lr := utils.NewLimitedResource(3, createStringResource, destroyStringResource)

	// 定义资源借用的超时时间
	timeout := 2 * time.Second

	// 模拟并发借用资源
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			resource, err := lr.Borrow(ctx, timeout)
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
