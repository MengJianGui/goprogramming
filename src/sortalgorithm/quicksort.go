package sortalgorithm

// QuickSort 快速排序：时间复杂度：O(nlogn)，空间复杂度：O(logn)---每次递归都是使用一个pivot
// 分段找到基准值并放到恰当位置
func QuickSort(slice []int) {
	  length := len(slice)
	  if length <= 1 {
			return
	  }
	  pivot := slice[0]
	  left, right := 0, length-1
	  for left < right {
			for left < right && slice[right] >= pivot {
				  right--
			}
			slice[left] = slice[right]
			for left < right && slice[left] < pivot {
				  left++
			}
			slice[right] = slice[left]
	  }
	  slice[left] = pivot
	  QuickSort(slice[:left])
	  QuickSort(slice[left+1:])
}
