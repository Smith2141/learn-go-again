package main

import "fmt"

func main() {

	var a, b, d int = 1, 2, 4

	var p_nums [4]*int // массив из 4 указателей на значения типа int

	p_nums[0] = &a
	p_nums[1] = &b
	p_nums[3] = &d

	fmt.Println(p_nums) // [0xc000010100 0xc000010108 <nil> 0xc000010110]
}
