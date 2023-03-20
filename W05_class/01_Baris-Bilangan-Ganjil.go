package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
    ganjil(n)
}

func ganjil(R int) { 
    if R >= 1 {
        if R%2 == 1 {
            ganjil(R-2)
            fmt.Printf("%d ", R)
        } else {
            ganjil(R-1)
        }
    }
}
