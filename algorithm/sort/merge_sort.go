package sort

//归并排序
//从小到大

func MergeSort(slice []int) {
	merge_sort_c(slice, 0, len(slice)-1)
}
func merge_sort_c(slice []int, left, right int) {
	if right <= left {
		return
	}
	mid := left + (right-left)>>1
	merge_sort_c(slice, left, mid)
	merge_sort_c(slice, mid+1, right)
	re := merge_sort_merge(slice[left:mid+1], slice[mid+1:right+1])
	for index, v := range re {
		slice[left+index] = v
	}
}

func merge_sort_merge(leftSlice, rightSlice []int) []int {
	leftNum := len(leftSlice)
	rightNum := len(rightSlice)
	leftCurrent := 0
	rightCurrent := 0

	result := make([]int, 0)
	for leftCurrent < leftNum && rightCurrent < rightNum {
		if leftSlice[leftCurrent] <= rightSlice[rightCurrent] {
			result = append(result, leftSlice[leftCurrent])
			leftCurrent++
		} else {
			result = append(result, rightSlice[rightCurrent])
			rightCurrent++
		}
	}
	for leftCurrent < leftNum {
		result = append(result, leftSlice[leftCurrent])
		leftCurrent++
	}
	for rightCurrent < rightNum {
		result = append(result, rightSlice[rightCurrent])
		rightCurrent++
	}
	return result
}
