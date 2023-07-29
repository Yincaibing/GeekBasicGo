package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	removeSliceElement(slice1, 2)
	fmt.Println(slice1)
}
