package main

import "fmt"

func main() {
	var x int
	fmt.Scanln(&x)

	var sum, result float64
	angkaHitung := hitungAngka_1303220150(x)
	ratarata_1303220150(x, angkaHitung, &sum)

	if angkaHitung > 0 {
		result = sum / float64(angkaHitung)
	}
	fmt.Printf("%.2f\n", result)
}

func hitungAngka_1303220150(x int) int {
	if x < 10 {
		return 1
	}
	return 1 + hitungAngka_1303220150(x/10)
}

func ratarata_1303220150(bil int, i int, hasil *float64) {
	if bil < 10 {
		*hasil += float64(bil)
		return
	}

	digit := bil % 10
	*hasil += float64(digit)
	ratarata_1303220150(bil/10, i-1, hasil)
}
