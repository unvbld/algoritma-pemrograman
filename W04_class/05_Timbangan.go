package main

import "fmt"

const pi float64 = 3.14

func main() {
    var r, h1, p1, h2, p2 int
    var m1, m2 float64

    fmt.Scan(&r)
    fmt.Scan(&h1, &p1)
    fmt.Scan(&h2, &p2)

    m1 = massa(r, h1, p1)
    m2 = massa(r, h2, p2)

    display(m1, m2)
}

func volume(r, h int) float64 {
    return pi * float64(r * r * h)
}

func massa(r, h, p int) float64 {
    return volume(r, h) * float64(p)
}

func display(m1, m2 float64) {
    if m1 > m2 {
        fmt.Printf("%.3f\n", m1 - m2)
    } else if m2 > m1 {
        fmt.Printf("%.3f\n", m2 - m1)
    } else {
        fmt.Println("BALANCE")
    }
}
