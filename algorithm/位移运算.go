package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 << 10) // 1024 (1 KiB)
	fmt.Println(1 << 20) // 1048576 (1 MiB)
	fmt.Println(1 << 30) // 1073741824 (1 GiB)
}
