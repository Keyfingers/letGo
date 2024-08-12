package main

import (
	"fmt"
	"sync"
)

func main() {
	number, letter := make(chan bool), make(chan bool)
	var wg sync.WaitGroup

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true

			}
		}
	}(&wg)
	number <- true
	wg.Wait()

}

//
//func find(target int) int  {
//	var arr []int
//	go func() {
//		for i := 0; i < 100; i++ {
//			arr = append(arr, arr[i]*2)
//		}
//	}()
//
//	var size = 2
//
//
//
//}
