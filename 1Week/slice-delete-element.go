package main

// 方法 1:利用原来的切片去掉
func removeSliceElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
