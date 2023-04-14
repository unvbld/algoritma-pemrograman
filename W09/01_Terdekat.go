package main

import (
    "fmt"
    "math"
)

type Point struct {
    x, y int
}

func main() {
    var A, B, C, D Point
    var d1, d2 float64

    fmt.Scan(&A.x, &A.y, &B.x, &B.y)
    fmt.Scan(&C.x, &C.y, &D.x, &D.y)

    d1 = jarak(A, B)
    d2 = jarak(C, D)

    if d1 < d2 {
        fmt.Printf("Titik terdekat adalah titik A(%d,%d) ", A.x, A.y)
        fmt.Printf("dengan B(%d,%d) dengan jarak %.2f.\n", B.x, B.y, d1)
    } else {
        fmt.Printf("Titik terdekat adalah titik C(%d,%d) ", C.x, C.y)
        fmt.Printf("dengan D(%d,%d) dengan jarak %.2f.\n", D.x, D.y, d2)
    }
}

func jarak(p1, p2 Point) float64 {
    var a, b float64

    a = float64((p1.x - p2.x) * (p1.x - p2.x))
    b = float64((p1.y - p2.y) * (p1.y - p2.y))

    return math.Sqrt(a + b)
}
