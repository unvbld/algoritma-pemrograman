package main

import "fmt"

const NMAX int = 100

type tabInt [NMAX]int

func main() {
    var T tabInt
    var n int

    isiData(&T, &n)

    fmt.Println(T[:n])
    fmt.Println(median(T, n))
}

func isiData(A *tabInt, n *int) {
    var i, x int

    fmt.Scan(n)

    for i = 0; i < *n; i++ {
        fmt.Scan(&x)

        if i < NMAX {
            A[i] = x
        }
    }

    if *n > NMAX {
        *n = NMAX
    }
}

func insertionSort(A *tabInt, n int) {
    var i, idx int

    for i = 1; i < n; i++ {
        idx = i

        for idx > 0 && A[idx] < A[idx-1] {
            A[idx], A[idx-1] = A[idx-1], A[idx]
            idx--
        }
    }
}

func median(A tabInt, n int) float64 {
    var T tabInt

    T = A

    insertionSort(&T, n)

    if n % 2 == 0 {
        return float64(T[n/2 - 1] + T[n/2]) / 2
    }

    return float64(T[n/2])
}
