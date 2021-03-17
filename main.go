package main

import (
	"fmt"
	"interview/data_structure/heap"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	var o int
	n := 100
	maxHeap := heap.NewMaxHeap(uint64(n))

	slice := make([]int64, n)
	for i := 0; i < n; i++ {
		slice[i] = int64(rand.Intn(1000) - 500)
		// if o = maxHeap.AddElement(int64(rand.Intn(100) - 50)); o != 1 {
		// 	fmt.Printf("AddElement err return: %d\n", o)
		// 	os.Exit(2)
		// }
	}

	fmt.Println(slice)

	if o = maxHeap.BuildHeapFromData(slice); o != 1 {
		fmt.Printf("BuildHeapFromData err return: %d\n", o)
		os.Exit(2)
	}

	fmt.Println(maxHeap.GetData())
	fmt.Println(maxHeap.GetCap())
	fmt.Println(maxHeap.GetLen())
	if o = maxHeap.Sort(); o != 1 {
		fmt.Printf("Sort err return: %d\n", o)
		os.Exit(2)
	}

	fmt.Println(maxHeap.GetData())
	fmt.Println(maxHeap.GetCap())
	fmt.Println(maxHeap.GetLen())

}
