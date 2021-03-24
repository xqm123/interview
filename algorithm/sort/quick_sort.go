package sort

//快速排序
//从小到大
func QuickSort(slice []int) {
	quick_sort_c(slice, 0, len(slice)-1)
}

func quick_sort_c(slice []int, left, right int) {
	if right <= left {
		return
	}
	q := quick_sort_partition(slice, left, right)
	quick_sort_c(slice, left, q-1)
	quick_sort_c(slice, q+1, right)
}

//quick_sort_partition 分区函数 这是提升效率的核心
func quick_sort_partition(slice []int, left, right int) int {

	i := left
	point := slice[right] //选择最后一个为分区点

	//原地分区巧妙实现  类似于选择排序的已排序区和未排序区
	for j := left; j < right; j++ {
		if slice[j] < point {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}

	slice[i], slice[right] = slice[right], slice[i]

	return i
}
