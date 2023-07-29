package main

// 方法 1：
func removeSliceElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
