package main

import (
	"fmt"
	"sort"
	"strings"
)

const dim int = 100

func main() {
	slice1 := make([]int, dim)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = i + 1
	}

	fmt.Println(slice1)
	fmt.Println(strings.Repeat("-*", 10))

	slice1 = append(slice1[:10], slice1[dim-10:]...)
	fmt.Println(slice1)
	fmt.Println(strings.Repeat("-*", 10))

	sort.Sort(sort.Reverse(sort.IntSlice(slice1)))
	fmt.Println(slice1)
	fmt.Println(strings.Repeat("-*", 10))
}
