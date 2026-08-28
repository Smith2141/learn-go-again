package main

import "fmt"

func main() {
	// определите переменные ver, id, pi
	var (
		ver = 0.001
		id int
		pi = 3.1415
	)

	var verV string = fmt.Sprintf("v%g", ver)

	fmt.Println("ver =", verV, "id =", id, "pi =", pi)
}
