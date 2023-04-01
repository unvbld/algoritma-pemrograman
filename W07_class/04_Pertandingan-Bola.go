package main

import "fmt"

const N int = 256

type tabGol int

func main() {
    var t1, t2 [N]tabGol
    var n1, n2 int

    inputData(&t1, &n1)
    inputData(&t2, &n2)

    fmt.Println(rataan(t1, n1), rataan(t2, n2))
}

func inputData(t *[N]tabGol, n *int) {
    var x int

    fmt.Scan(&x)

    for x >= 0 {
        *&t[*n] = tabGol(x)
        *n++
        fmt.Scan(&x)
    }
}

func rataan(t [N]tabGol, n int) float64 {
    var i, total int

    for i = 0; i < n; i++ {
        total += int(t[i])
    }

    return float64(total) / float64(n)
}
