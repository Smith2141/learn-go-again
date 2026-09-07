package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	slice1 := make([]int, 100)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = i + 1
	}

	fmt.Println(slice1)
	fmt.Println(strings.Repeat("-*", 10))

	slice1 = append(slice1[:10], slice1[89:]...)
	fmt.Println(slice1)
	fmt.Println(strings.Repeat("-*", 10))

	sort.Sort(sort.Reverse(sort.IntSlice(slice1)))
	fmt.Println(slice1)
	fmt.Println(strings.Repeat("-*", 10))
}
