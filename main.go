package main

import (
	"fmt"
	"runtime"
)

func main() {
	slice := make([]int, 3, 5)
	fmt.Println(runtime.Version())
	fmt.Println("111")
	for i, v := range slice {
		fmt.Printf("%p\n", &v)
		fmt.Println(i)
	}
}
