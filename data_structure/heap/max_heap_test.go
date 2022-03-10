package heap

import (
    "math/rand"
    "testing"
    "time"
)

func getData(n int) []int64 {
    rand.Seed(time.Now().UnixNano())
    var data []int64
    a := 2*n
    for i := 0; i < n; i++ {
        elem := rand.Intn(a) - a/2
        data = append(data, int64(elem))
    }
    return data
}

func BenchmarkAddElement(b *testing.B) {
    capn := uint64(1000000000)
    heapIns := NewMaxHeap(capn)
    preN := 10000000
    if err := heapIns.BuildHeapFromData(getData(preN)); err != nil {
        b.Error(err)
        return
    }
    rand.Seed(time.Now().UnixNano())
    n := 3
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        if b.N > n {
            break
        }
        elem := rand.Intn(n) - n/2
        err := heapIns.AddElement(int64(elem))
        if err != nil {
            b.Error(err)
        }
    }
}

