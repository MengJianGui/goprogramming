package sortalgorithm

/**
归并排序：时间复杂度：O(n*logn)，空间复杂度：O(n)
归并排序算法是将两个（或两个以上）有序表合并成一个新的有序表，
即把待排序序列分为若干个子序列，每个子序列是有序的。
然后再把有序子序列合并为整体有序序列。
*/
func MergeSort(slice []int) {
	slice = mergeSort(slice)
}

func mergeSort(slice []int) []int {
	length := len(slice)
	if len(slice) <= 1 {
		return slice
	}
	mid := length / 2
	left := mergeSort(slice[:mid])		// 递归的每个时间复杂度为O(logn)
	right := mergeSort(slice[mid:])
	return Merge(left, right)		// 时间复杂度为O(n)
}

// merge() 合并两个有序数组成为一个新的有序数组(包含重复数字)
func Merge(left, right []int) []int {
	lenL, lenR := len(left), len(right)
	var result []int

	l, r := 0, 0
	for l < lenL && r < lenR {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	for l < lenL && r == lenR {
		result = append(result, left[l])
		l++
	}
	for r < lenR && l == lenL {
		result = append(result, right[r])
		r++
	}

	return result
}
