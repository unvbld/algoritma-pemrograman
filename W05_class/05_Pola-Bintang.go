package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
    bintang(n, 1)
}

func bintang(n int, i int) {
    if i <= n {
        for j := 1; j <= i; j++ {
            fmt.Printf("*")
        }
        fmt.Println()
        bintang(n, i+1)
    }
}
