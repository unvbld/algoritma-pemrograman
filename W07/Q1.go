package main

import "fmt"

func inputBilangan(bil *int) {
	for {
		fmt.Scan(bil)
		if *bil >= 0 {
			break
		}
	}
}

func stop(bil int) bool {
	if bil == 0 {
		return true
	} else {
		return false
	}
}

func hitung(mean, min, max *float64, n *int) {
	var bil int
	var sum float64

	*min = 9999999.0
	*max = -9999999.0

	for {
		inputBilangan(&bil)
		if stop(bil) {
			break
		}

		sum += float64(bil)
		*n++

		if float64(bil) < *min {
			*min = float64(bil)
		}

		if float64(bil) > *max {
			*max = float64(bil)
		}
	}

	if *n == 0 {
		*mean = 0.0
	} else {
		*mean = sum / float64(*n)
	}
}

func main() {
	var mean, min, max float64
	var n int

	hitung(&mean, &min, &max, &n)

	if n == 0 {
		fmt.Println("- - - -")
	} else {
		fmt.Printf("%.2f %.0f %.0f %d\n", mean, min, max, n)
	}
}
