package main

//func main() {
// arguments:=os.Args
//	fmt.Println(RecursivePower(2, 5))
//}

func RecursivePower(nb int, power int) int {
	if power > 1 {
		return nb * RecursivePower(nb, power-1)
	} else if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	} else {
		return 0
	}

}
