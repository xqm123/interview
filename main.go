package main

import (
	"fmt"
	mysort "interview/algorithm/sort"
	"math/rand"
	"time"
)

func main() {
	//slice := []int{1, 10, -2, 21, 2, 30, -29, 83, 2, -10, -3, 67, 188, 23, -1, -1, 230, 121, -140, 52, 51, 43, 67, -1, 234, -187, 23}
	//slice := []string{"abs1", "13870669286", "13845761209", "", "13209181203", "1392309", "13870659302"}

	rand.Seed(time.Now().UnixNano())
	n := 100000
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Int()
	}

	var t1 int64
	var t2 int64
	a := make([]int, n)

	//快速排序
	copy(a, slice)
	t1 = time.Now().UnixNano() / 1e6
	mysort.QuickSort(a)
	t2 = time.Now().UnixNano() / 1e6
	//fmt.Println(a)
	fmt.Printf("QuickSort spendtime: %vms\n", t2-t1)

	//归并排序
	copy(a, slice)
	t1 = time.Now().UnixNano() / 1e6
	mysort.MergeSort(a)
	t2 = time.Now().UnixNano() / 1e6
	//fmt.Println(a)
	fmt.Printf("MergeSort spendtime: %vms\n", t2-t1)

	//冒泡排序
	copy(a, slice)
	t1 = time.Now().UnixNano() / 1e6
	mysort.BubbleSort(a)
	t2 = time.Now().UnixNano() / 1e6
	//fmt.Println(a)
	fmt.Printf("BubbleSort spendtime: %vms\n", t2-t1)

	//插入排序
	copy(a, slice)
	t1 = time.Now().UnixNano() / 1e6
	mysort.InsertionSort(a)
	t2 = time.Now().UnixNano() / 1e6
	//fmt.Println(a)
	fmt.Printf("InsertionSort spendtime: %vms\n", t2-t1)

	//选择排序
	copy(a, slice)
	t1 = time.Now().UnixNano() / 1e6
	mysort.SelectionSort(a)
	t2 = time.Now().UnixNano() / 1e6
	//fmt.Println(a)
	fmt.Printf("SelectionSort spendtime: %vms\n", t2-t1)
}
