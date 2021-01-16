package sort

//基数排序
//只支持英文或数字(ascii<=127)字符串切片
//按位按ascii从小到大排序
func RadixSort(slice []string) {
	count := len(slice)
	maxDiag := 0
	for _, v := range slice {
		strlen := len(v)
		if strlen > maxDiag {
			maxDiag = strlen
		}
	}

	cp := make([]string, count)
	for i := maxDiag; i > 0; i-- {
		buckets := make([][]int, 129)
		bucket_r := make([]int, 0) //字符串在相应位缺位的桶
		copy(cp, slice)
		for k, v := range cp {
			if len(v) >= i {
				index := int(v[i-1])
				if len(buckets[index]) == 0 {
					buckets[index] = []int{k}
				} else {
					buckets[index] = append(buckets[index], k)
				}
			} else {
				bucket_r = append(bucket_r, k)
			}
		}
		inx := 0
		//先遍历缺位桶  这个是最小的
		for _, buck := range bucket_r {
			slice[inx] = cp[buck]
			inx++
		}
		for _, bucket := range buckets {
			for _, buck := range bucket {
				slice[inx] = cp[buck]
				inx++
			}
		}
	}
	return
}
