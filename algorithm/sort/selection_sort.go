package sort

//选择排序

//选择排序(Selection-sort)是一种简单直观的排序算法。它的工作原理：首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。

/**
 * 选择排序
 * Avg O(n^2)
 * Bad O(n^2)
 * Best O(n^2)
 */

//算法分析: 表现最稳定的排序算法之一，因为无论什么数据进去都是O(n2)的时间复杂度，所以用到它的时候，数据规模越小越好。唯一的好处可能就是不占用额外的内存空间了吧。理论上讲，选择排序可能也是平时排序一般人想到的最多的排序方法了吧。

//从小到大
func SelectionSort(slice []int) {
	len := len(slice)
	var maxIndex int
	for i := 0; i < len-1; i++ {
		maxIndex = i
		for j := i + 1; j < len; j++ {
			if slice[j] < slice[maxIndex] {
				maxIndex = j
			}
		}
		slice[i], slice[maxIndex] = slice[maxIndex], slice[i]
	}

	return
}
