package piscine

func IterativeFactorial(c int) int {
	var a int
	if c == 0 || c == 1 {
		a = 1
	} else if c > 0 {
		for i := 1; i <= 20; i++ {
			a = a * i
		}

	} else {
		a = 0
	}
	return a
}
