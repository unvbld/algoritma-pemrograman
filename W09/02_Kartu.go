package main

import "fmt"

func main() {
    var arr [52]int
    var x, n, i int

    n = 0
    fmt.Scan(&x)

    for x != 0 {
        arr[n] = x
        n++
        fmt.Scan(&x)
    }

    for i = n-1; i > -1; i-- {
        fmt.Print(arr[i], " ")
    }

    fmt.Println()
}
