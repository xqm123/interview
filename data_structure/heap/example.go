package heap

import (
	"fmt"
	"math/rand"
	"time"
)

// FindForKMax 无序数组找到第k大的元素
func FindForKMax() {

	N := int64(1000000)
	k := 5000
	srcSlice := make([]int64, N, N)

	rand.Seed(time.Now().UnixNano())
	for i, _ := range srcSlice {
		srcSlice[i] = rand.Int63n(10*N) - 5*N
	}

	var err error
	minHeap := NewMinHeap(uint64(k))
	if err = minHeap.BuildHeapFromData(srcSlice[:k]); err != nil {
		fmt.Printf("BuildHeapFromData error err: %s\n", err.Error())
		return
	}

	for i := k; i < len(srcSlice); i++ {
		val, _ := minHeap.GetTopValue()
		if srcSlice[i] > val {
			if err = minHeap.ReplaceTopValue(srcSlice[i]); err != nil {
				fmt.Printf("ReplaceTopValue error err: %s\n", err.Error())
				return
			}
		}

	}
	val, _ := minHeap.GetTopValue()
	fmt.Printf("第k大元素：%v\n", val)

}
