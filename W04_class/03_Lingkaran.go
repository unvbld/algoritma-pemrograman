package main

import (
	"fmt"
	"math"
)

func main() {
    var x_1, y_1, r_1, x_2, y_2, r_2 int
    var x_koor, y_koor int
    var b1, b2 bool

    fmt.Scan(&x_1, &y_1, &r_1)
    fmt.Scan(&x_2, &y_2, &r_2)
    fmt.Scan(&x_koor, &y_koor)

    b1 = didalam(x_1, y_1, r_1, x_koor, y_koor)
    b2 = didalam(x_2, y_2, r_2, x_koor, y_koor)

    if b1 && b2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if b1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if b2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}

func jarak(a, b, c, d int) float64 {
    return akar(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func didalam(cx, cy, r, x, y int) bool {
    return jarak(cx, cy, x, y) < float64(r)
}

func akar(x float64) float64 {
    return math.Sqrt(x)
}
