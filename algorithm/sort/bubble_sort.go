package sort

//冒泡排序

//冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。

/**
 * 冒泡排序
 * Avg O(n^2)
 * Bad O(n^2)
 * Best O(n)
 */

//从小到大
func BubbleSort(slice []int) {
	len := len(slice)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if slice[j+1] < slice[j] {
				tmp := slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = tmp
			}
		}
	}

	return
}

//eg:
// func main() {
// 	slice := []int{1, 10, -2, 2, 2, 3, 2, 8, 2, -10, -3, 67, 188, 23, -1, -1, 230, 12, 14, 52, 51, 43, 67, -1, 234, -187, 23}

// 	fmt.Println(BubbleSort(slice))
// }
