package main

import "fmt"

func main() {
    var n int
    var avrg float64

    fmt.Scan(&n)
    average(n, 1, &avrg)
    fmt.Println(avrg)
}

func average(bil, i int, hasil *float64) {
    if bil == 0 {
        if i > 1 {
            *hasil = *hasil / float64(i-1)
        } else {
            *hasil = 0
        }
    } else {
        if i == 1 {
            *hasil = 0
        }

        *hasil = *hasil + float64(bil % 10)
        average(bil / 10, i+1, hasil)
    }
}
