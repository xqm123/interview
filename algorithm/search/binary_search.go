package search

//BinarySearchV1 二分查找  --只要找到等于给定值的元素就返回index
// slice 需从小到大排序
func BinarySearchV1(slice []int, search int) int {

	low := 0
	high := len(slice) - 1
	var mid int
	for high >= low {
		mid = low + (high-low)>>1
		if slice[mid] == search {
			return mid
		} else if slice[mid] > search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

//BinarySearchV2 二分查找  --查找第一个等于给定值的元素
// slice 需从小到大排序
func BinarySearchV2(slice []int, search int) int {

	low := 0
	high := len(slice) - 1
	var mid int
	for high >= low {
		mid = low + (high-low)>>1
		if slice[mid] == search {
			if mid == 0 || slice[mid-1] != search {
				return mid
			}
			high = mid - 1
		} else if slice[mid] > search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

//BinarySearchV3 二分查找  --查找最后一个等于给定值的元素
// slice 需从小到大排序
func BinarySearchV3(slice []int, search int) int {

	n := len(slice) - 1
	low := 0
	high := n
	var mid int
	for high >= low {
		mid = low + (high-low)>>1
		if slice[mid] == search {
			if mid == n || slice[mid+1] != search {
				return mid
			}
			low = mid + 1
		} else if slice[mid] > search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

//BinarySearchV4 二分查找  --查找最后一个小于等于给定值的元素
// slice 需从小到大排序
func BinarySearchV4(slice []int, search int) int {

	n := len(slice) - 1
	low := 0
	high := n
	var mid int
	for high >= low {
		mid = low + (high-low)>>1
		if slice[mid] <= search {
			if mid == n || slice[mid+1] > search {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

//BinarySearchV5 二分查找  --查找第一个大于等于给定值的元素
// slice 需从小到大排序
func BinarySearchV5(slice []int, search int) int {

	n := len(slice) - 1
	low := 0
	high := n
	var mid int
	for high >= low {
		mid = low + (high-low)>>1
		if slice[mid] >= search {
			if mid == 0 || slice[mid-1] < search {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func Sample(slice []int, search int) int {
	for i, v := range slice {
		if v == search {
			return i
		}
	}
	return -1
}
