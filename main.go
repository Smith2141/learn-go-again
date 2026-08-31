package main

import "fmt"

func main() {

	for j := 1; j <= 100; j++ {
		var fizz, buzz, sum int

		// fmt.Print(j, ":")
		if j%3 == 0 {
			fizz = 1
		}
		if j%5 == 0 {
			buzz = 2
		}
		sum = fizz + buzz
		// fmt.Println("f b s", fizz, buzz, sum)

		switch sum {
		case 3:
			fmt.Println("FizzBuzz")
		case 2:
			fmt.Println("Buzz")
		case 1:
			fmt.Println("Fizz")
		default:
			fmt.Println(j)
		}
	}
}
