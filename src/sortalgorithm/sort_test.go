package sortalgorithm

import (
	  "fmt"
	  "testing"
)

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	  slice := []int{15, 23, 6, 11, 16, 19, 5, 12}
	  BubbleSort(slice)
	  fmt.Println(slice)
}

// 直接插入排序
func TestInsertSort(t *testing.T) {
	  slice := []int{15, 23, 6, 11, 16, 19, 5, 12}
	  InsertSort(slice)
	  fmt.Println(slice)
}

// 快速排序
func TestQuickSort(t *testing.T) {
	  slice := []int{15, 23, 6, 11, 16, 19, 5, 12}
	  QuickSort(slice)
	  fmt.Println(slice)
}

// 归并排序
func TestMerge(t *testing.T) {
	slice := []int{15, 23, 6, 11, 16, 19, 5, 12}
	MergeSort(slice)
	fmt.Println(slice)
}
