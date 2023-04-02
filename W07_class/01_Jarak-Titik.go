package main

import (
    "fmt"
    "math"
)

type titik struct {
    x, y int
}

func main() {
    var p1, p2 titik
    var j float64

    fmt.Scan(&p1.x, &p2.x, &p1.y, &p2.y)
    j = jarak(p1, p2)
    fmt.Printf("%.2f\n", j)
}

func jarak(P1, P2 titik) float64 {
    var x_komponen, y_komponen float64

    x_komponen = float64(P1.x - P2.x)
    y_komponen = float64(P1.y - P2.y)

    return akar(x_komponen*x_komponen + y_komponen*y_komponen)
}

func akar(x float64) float64 {
    return math.Sqrt(x)
}
