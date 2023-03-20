package main

import "fmt"

func main() {
    var c int
    var r, f, k float64

    fmt.Scan(&c)

    r = float64(c) * 0.8
    f = 1.8 * float64(c) + 32
    k = float64(c) + 273.15

    fmt.Printf("%.2fr %.2ff %.2fk\n", r, f, k)
}
