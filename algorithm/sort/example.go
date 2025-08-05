package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func Example(n int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if n <= 0 {
		return
	}

	slice1 := make([]int, n)
	for i := 0; i < n; i++ {
		slice1[i] = rand.Intn(100*n) - (100*n)/3
	}
	// copy 切片
	slice2 := make([]int, n)
	slice3 := make([]int, n)
	slice4 := make([]int, n)
	slice5 := make([]int, n)
	copy(slice2, slice1)
	copy(slice3, slice1)
	copy(slice4, slice1)
	copy(slice5, slice1)

	fmt.Printf("before slice len: %d\n", n)
	// fmt.Printf("before slice: %v\n", utils.JsonString(slice1))

	t1 := time.Now().UnixNano() / 1e6
	QuickSort(slice1)
	t2 := time.Now().UnixNano() / 1e6
	fmt.Printf("QuickSort spend time: %d ms\n", t2-t1)
	// fmt.Printf("QuickSort after slice: %v\n", utils.JsonString(slice1))

	t1 = time.Now().UnixNano() / 1e6
	MergeSort(slice2)
	t2 = time.Now().UnixNano() / 1e6
	fmt.Printf("MergeSort spend time: %d ms\n", t2-t1)
	// fmt.Printf("MergeSort after slice: %v\n", utils.JsonString(slice2))

	//t1 = time.Now().UnixNano() / 1e6
	//InsertionSort(slice3)
	//t2 = time.Now().UnixNano() / 1e6
	//fmt.Printf("InsertionSort spend time: %d ms\n", t2-t1)
	//// fmt.Printf("InsertionSort after slice: %v\n", utils.JsonString(slice3))
	//
	//t1 = time.Now().UnixNano() / 1e6
	//SelectionSort(slice4)
	//t2 = time.Now().UnixNano() / 1e6
	//fmt.Printf("SelectionSort spend time: %d ms\n", t2-t1)
	//// fmt.Printf("SelectionSort after slice: %v\n", utils.JsonString(slice4))
}
