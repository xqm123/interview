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
	quick_sort_c(slice, left, q)
	quick_sort_c(slice, q+1, right)
}

//quick_sort_partition 分区函数 这是提升效率的核心
func quick_sort_partition(slice []int, left, right int) int {

	var re int
	var q int
	var mins []int
	var maxs []int
	for i := right; i >= left; i-- {
		pivot := slice[i] //简单点 这里分区点从后往前取 直到合格为止
		mins = make([]int, 0)
		maxs = make([]int, 0)
		for j := left; j <= right; j++ {
			if slice[j] <= pivot {
				mins = append(mins, slice[j])
			} else {
				maxs = append(maxs, slice[j])
			}
		}

		q = left + len(mins)
		re = q - 1
		if re == left || re == right {
			//这是不合格的分区点
			continue
		}

		//走到这说明这是合格的
		break

	}

	if re == left || re == right {
		//走到这都没有合格的 只能说明任意取一个都可以
		re = left + (right-left)>>1
	}

	for k, v := range mins {
		slice[left+k] = v
	}
	for k, v := range maxs {
		slice[q+k] = v
	}
	return re
}
