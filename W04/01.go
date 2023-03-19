package main

import (
	"fmt"
	"math"
)

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Println(z(a, b), z(b, a))
}

func z(x int, y int) float64 {
    return 1.5 * math.Sqrt(2 * float64(x))
}
