package main

import "fmt"

func main() {
	// testLoop()
	// doWhileLoop()
	// whileLoop()
	// normIfElse()
	// processBeforeIf()
}
func testLoop() {
	for i := 1; i < 9; i++ {
		println(i)
	}
}
func doWhileLoop() {
	i := 1
	for {
		if i == 2 {
			println(i)
			break

		}
		i++
	}
}
func whileLoop() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}
}
func normIfElse() {
	x := 10
	if x < 10 {
		println("x มีค่าน้อยกว่า 10")
	} else {
		println("x มีค่ามากว่า หรือ เท่ากับ 10")
	}
}

func processBeforeIf() {
	i := 2
	j := 3
	k := 0
	k = i + j
	if k == 5 {
		println("k มีค่าเท่ากับ 5")
	}
}
