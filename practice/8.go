package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	number, letter := make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				print(i)
				i++
				print(i)
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
				if i >= 'Z' {
					wg.Done()
					return
				}
				print(string(i))
				i++
				print(string(i))
				i++
				number <- true
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}
