package main

import "reflect"

// 要求一：能够实现删除操作就可以。
// 方法 1:利用原来的切片新创建一个切片，新创建的切片不包含给定的元素
func removeSliceElement1(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// 要求二：考虑使用比较高性能的实现。
// 方法 2：在原来的切片上做操作，删除特定元素
func removeSliceElement2(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// 要求三：改造为泛型方法
func removeSliceElement3(slice interface{}, element interface{}) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("RemoveElement: not a slice")
	}

	elemType := reflect.TypeOf(element)
	if s.Type().Elem() != elemType {
		panic("RemoveElement: slice element type and element type mismatch")
	}

	length := s.Len()
	for i := 0; i < length; i++ {
		value := s.Index(i)
		if reflect.DeepEqual(value.Interface(), element) {
			// 将删除元素后的切片重新组装
			reflect.Copy(s.Slice(i, length-1), s.Slice(i+1, length))

			// 返回删除元素后的切片
			newSlice := reflect.MakeSlice(s.Type(), length-1, length-1)
			reflect.Copy(newSlice, s.Slice(0, length-1))
			return newSlice.Interface()
		}
	}

	// 如果切片中不存在要删除的元素，则返回原始切片
	return slice
}

//要求四：支持缩容，并旦设计缩容机制。
