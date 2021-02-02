package main

import (
	"fmt"
	mySearch "interview/algorithm/search"
	"time"
)

func main() {
	//slice := []int{1, 10, -2, 21, 2, 30, -29, 83, 2, -10, -3, 67, 188, 23, -1, -1, 230, 121, -140, 52, 51, 43, 67, -1, 234, -187, 23}
	//slice := []string{"abs1", "13870669286", "13845761209", "", "13209181203", "1392309", "13870659302"}

	n := 20000001
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i - 100000000
	}

	var t1 int64
	var t2 int64

	//二分查找
	//copy(a, slice)
	t1 = time.Now().UnixNano() / 1e6
	index := mySearch.BinarySearchV1(slice, 100000000)
	t2 = time.Now().UnixNano() / 1e6
	//fmt.Println(a)
	fmt.Printf("BinarySearchV1 search result: %d  spendtime: %vms\n", index, t2-t1)

	//遍历查找
	a := make([]int, n)
	copy(a, slice)
	t1 = time.Now().UnixNano() / 1e6
	index = mySearch.Sample(a, 100000000)
	t2 = time.Now().UnixNano() / 1e6
	//fmt.Println(a)
	fmt.Printf("Sample search result: %d  spendtime: %vms\n", index, t2-t1)
}
