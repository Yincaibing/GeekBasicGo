package main

import "fmt"

func main() {

	//验证方法 1
	slice1 := []int{1, 2, 3, 4, 5}
	removeSliceElement1(slice1, 0)
	fmt.Printf("验证方法 1：" + "删除数字 1")
	fmt.Println(slice1)

	//验证方法 2
	slice2 := []int{1, 2, 3, 4, 5, 6}
	removeSliceElement2(slice2, 1)
	fmt.Printf("验证方法 2：" + "删除数字 2")
	fmt.Println(slice2)

	//验证方法 3
	//删除切片中的数字2
	s := []int{1, 2, 3, 4, 5}
	element := 3
	newSlice := removeSliceElement3(s, element)
	fmt.Printf("验证方法 3：" + "删除数字 3")
	fmt.Println(newSlice) // 输出 [1 2 4 5]

	// 删除切片中的字符串"world"
	s2 := []string{"hello", "world", "welcome", "to", "go"}
	element2 := "world"
	newSlice2 := removeSliceElement3(s2, element2)
	fmt.Printf("验证方法 3：" + "删除字符串 world")
	fmt.Println(newSlice2) // 输出 [hello welcome to go]

}
