package main

import (
	"fmt"
	"reflect" // для сравнения срезов
	"sort"
)

func main() {

	slice1 := []int{1, 2, 3, 4}
	slice2 := []string{"Tom", "Bob", "Sam"}
	slice3 := []int{1, 2, 3}
	slice4 := []int{1, 2, 4, 3}
	sort.Ints(slice4)

	fmt.Println("slice1 == slice2:", reflect.DeepEqual(slice1, slice2)) // false
	fmt.Println("slice1 == slice3:", reflect.DeepEqual(slice1, slice3)) // false
	fmt.Println("slice1 == slice4:", reflect.DeepEqual(slice1, slice4)) // true
}
