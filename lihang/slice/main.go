package main

import "fmt"

// 切片程序实例
func main() {
	forRangeLoop()
}

// 创建切片
func newSlice() {
	var slice []string
	slice = append(slice, "one", "two", "three")
	fmt.Println(slice)
}

// for循环
func forLoop() {
	var slice []string
	slice = append(slice, "one", "two", "three")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[i]: %v\n", slice[i])
	}
}

// forRange循环
func forRangeLoop() {
	var slice []string
	slice = append(slice, "one", "two", "three")
	for _, v := range slice {
		fmt.Printf("v: %v\n", v)
	}
}
