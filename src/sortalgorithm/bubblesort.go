package sortalgorithm

/**
冒泡排序Go实现
时间复杂度：O(n^2)，空间复杂度：O(1)
*/
// BubbleSort 冒泡排序算法：Go语言数组是值传递，而切片就是引用传递
// 冒泡排序是相邻两个数进行比较，每次找出最大的数放到最右边。
func BubbleSort(slice []int) {
	  length := len(slice)
	  for i := 0; i < length; i++ {
			for j := 0; j < length - i -1; j++ {
				  if slice[j] > slice[j+1] {
						slice[j], slice[j+1] = slice[j+1], slice[j]
				  }
			}
	  }
}
