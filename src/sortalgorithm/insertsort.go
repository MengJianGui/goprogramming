package sortalgorithm

// InsertSort 直接插入排序：时间复杂度：O(n^2)，空间复杂度：O(1)
// 直接插入排序的基本思想是：将数组中的所有元素依次跟前面已经排好的元素相比较，
//如果选择的元素比已排序的元素小，则交换，直到全部元素都比较过为止。
// 直接插入排序前面的都是有序的。
func InsertSort(slice []int) {
	  length := len(slice)
	  for i := 1; i < length; i++ {
			temp := slice[i]
			for j := i - 1; j >= 0; j-- {
				  if temp < slice[j] {
						slice[j+1] = slice[j]
						slice[j] = temp
				  } else if j >= 1 && slice[j] > slice[j-1] {
						break
				  }
			}
	  }

}
