package main

import "fmt"

func main() {
	var n int
    var s string

    fmt.Scan(&n)
    pola(n, &s)
    fmt.Println(s)
}

func pola(n int, s *string) {
    *s = "\n" + *s

    if n == 1 {
        *s = "*" + *s
    } else {
        for i := 0; i < n; i++ {
            *s = "*" + *s
        }

        pola(n-1, s)
    }
}
