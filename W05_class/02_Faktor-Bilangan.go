package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    faktor(n, 1)
}

func faktor(n int, i int) {
    if i == n {
        fmt.Println(n)
    } else {
        if n%i == 0 {
            fmt.Printf("%d ", i)
        }
        faktor(n, i+1)
    }
}
