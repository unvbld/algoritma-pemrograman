package main

import "fmt"

const N int = 100

type himpunan [N]int

func main() {
    var A, B himpunan
    var n1, n2 int

    A = himpunan{11, 28, 33, 64, 95, 16}
    B = himpunan{11, 28, 33, 64, 95, 16, 100, 28, 33, 64, 95, 16}
    n1 = 6
    n2 = 12

    fmt.Println(valid(A, n1))
    fmt.Println(valid(B, n2))
}

func valid(himpunan himpunan, n int) bool {
    var idx, i, x int

    idx = 0

    for idx < n {
        x = himpunan[idx]

        for i = idx+1; i < n; i++ {
            if himpunan[i] == x {
                return false
            }
        }

        idx++
    }

    return true
}
