package main

import "fmt"

const NMAX int = 1000

type tabInt [NMAX]int

func main() {
    var tab tabInt
    var n, x, idx int

    tab = tabInt{4,12,17,23,27,29,30,33,37,44,47,60,71,82}
    n = 14

    fmt.Scan(&x)

    idx = binarySearch(tab, n, x)

    if idx != -1 {
        fmt.Printf("angka %d terdapat di dalam array pada urutan ke-%d\n", x, idx)
    } else {
        fmt.Printf("angka %d tidak terdapat di dalam array\n", x)
    }
}

func binarySearch(tab tabInt, n, x int) int {
    var left, right, mid int

    left = 0
    right = n - 1
    mid = (n - 1) / 2

    for mid != left && mid != right {
        if x < tab[mid] {
            right = mid
        } else if x > tab[mid] {
            left = mid
        } else {
            return mid
        }

        mid = (left + right) / 2
    }

    if tab[left] == x {
        return left
    }

    if tab[right] == x {
        return right
    }

    return -1
}
