package main

import "fmt"

func main() {
    var n, m int
    fmt.Scan(&n, &m)
    fmt.Println(jumlahBus(n, m), "bus")
}

func jumlahBus(penumpang int, kapasitas int) int {
    var n int
    n = 1

    for (kapasitas * n < penumpang) {
        n++
    }

    return n
}
