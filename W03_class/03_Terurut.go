package main

import "fmt"

func main() {
	var x, y int
	
    fmt.Scan(&x, &y)

    for x != y {
        mengurutkan(&x, &y)
        fmt.Println(x, y)
        fmt.Scan(&x, &y)
    }
}

func mengurutkan(a, b *int) {
    var tmp int

    if *a > *b {
        tmp = *a
        *a = *b
        *b = tmp
    }
}
