package main

import "fmt"

/**
swap 的写法
*/
func main() {
	i, j := 10, 11
	j, i = i, j
	fmt.Printf("hello i=%d j=%d", i, j)
	i2 := swap(i, j)
	fmt.Print(i2)
}

func swap(i int, j int) int {
	fmt.Print(i + j)
	return i + j
}
