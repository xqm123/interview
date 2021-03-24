package main

import (
	"fmt"
	"interview/algorithm/sort"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	n := 1000000

	slice := make([]int, n)
	slice2 := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(n)-n/2
		slice2[i] = slice[i]
	}

	//fmt.Println(slice)

	t1 := time.Now().UnixNano()/1e6
	sort.QuickSort(slice)
	t2 := time.Now().UnixNano()/1e6
	//fmt.Println(slice)
	fmt.Printf("QuickSort spend time: %d \n", t2-t1)

	t1 = time.Now().UnixNano()/1e6
	sort.MergeSort(slice)
	t2 = time.Now().UnixNano()/1e6
	//fmt.Println(slice)
	fmt.Printf("MergeSort spend time: %d \n", t2-t1)

}
