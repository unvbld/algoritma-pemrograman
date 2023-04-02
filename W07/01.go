package main

import "fmt"

func main() {
    var mean float64
    var min, max, n int

    hitung(&mean, &min, &max, &n)

    if n != 0 {
        fmt.Println(mean, max, min, n)
    } else {
        fmt.Println("- - - -")
    }
}

func inputBilangan(bil *int) {
    fmt.Scan(bil)

    for *bil < 0 {
        fmt.Scan(bil)
    }
}

func stop(bil int) bool {
    return bil == 0
}

func hitung(mean *float64, min, max, n *int) {
    var total, x int

    *mean = 0
    *min = 0
    *max = 0
    *n = 0

    inputBilangan(&x)

    for !stop(x) {
        if *min > x || *n == 0 {
            *min = x
        }

        if *max < x {
            *max = x
        }

        total += x
        *n++
        inputBilangan(&x)
    }

    if *n != 0 {
        *mean = float64(total) / float64(*n)
    }
}
