package main

import (
	student ".."
	"fmt"
)

func main() {
	nbits := student.ActiveBits(15)
	fmt.Println(nbits)
	nbits = student.ActiveBits(17)
	fmt.Println(nbits)
	nbits = student.ActiveBits(4)
	fmt.Println(nbits)
	nbits = student.ActiveBits(11)
	fmt.Println(nbits)
	nbits = student.ActiveBits(9)
	fmt.Println(nbits)
	nbits = student.ActiveBits(12)
	fmt.Println(nbits)
	nbits = student.ActiveBits(2)
	fmt.Println(nbits)

}

// 4
// 2
// 1
// 3
// 2
// 2
// 1
