package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)

    if segitiga(a, b, c) {
        fmt.Println("segitiga")
    } else {
        fmt.Println("bukan segitiga")
    }
}

func segitiga(a, b, c int) bool {
    var tmp int
    
    if (a > c) {
        tmp = a
        a = c
        c = tmp
    }
    if (a > b) {
        tmp = a
        a = b
        b = tmp
    }
    if (b > c) {
        tmp = b
        b = c
        c = tmp
    }

    return a + b > c
}
