package main

import "fmt"

const N int = 100

type himpunan [N]int

func main() {
    var A, B, C himpunan
    var n1, n2, n3 int

    A = himpunan{11, 28, 33, 64, 95, 16, 100, 15}
    B = himpunan{3, 11, 7, 28, 33, 6}
    n1 = 8
    n2 = 6

    cariIrisan(A, B, &C, n1, n2, &n3)

    fmt.Println(C[:n3])
}

func cariIrisan(A, B himpunan, C *himpunan, n1, n2 int, n3 *int) {
    var i, j int

    *C = himpunan{}
    *n3 = 0

    for i = 0; i < n1; i++ {
        for j = 0; j < n2; j++ {
            if A[i] == B[j] {
                *&C[*n3] = A[i]
                *n3++
            }
        }
    }
}
