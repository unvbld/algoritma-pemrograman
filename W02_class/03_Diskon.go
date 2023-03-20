package main

import "fmt"

func main() {
    var n int
    var b bool
    fmt.Scan(&n, &b)

    fmt.Println(float64(n) * float64(1 - diskon(n, b)))
}

func diskon(belanja int, member bool) float64 {
    if belanja > 100000 {
        if member {
            return 0.1
        }

        return 0.05
    }

    return 0
}
