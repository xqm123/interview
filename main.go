package main

import (
	"fmt"
	mysort "interview/algorithm/sort"
)

func main() {
	//slice := []int{1, 10, -2, 21, 2, 30, -29, 83, 2, -10, -3, 67, 188, 23, -1, -1, 230, 121, -140, 52, 51, 43, 67, -1, 234, -187, 23}
	slice := []string{"abs1", "13870669286", "13845761209", "13209181203", "1392309", "13870659302"}
	mysort.RadixSort(slice)
	fmt.Printf("%#v", slice)
}
