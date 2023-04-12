package main

import "fmt"

const NMAX int = 30

type intArray struct {
    tabInt [NMAX]int
    N int
}

func main() {
    var array1, array2 intArray

    inputArray(&array1)
    inputArray(&array2)
    sortArray(&array1)
    sortArray(&array2)

    fmt.Println(isSimilar(array1, array2))
}

func inputArray(T *intArray) {
    var x int

    fmt.Scan(&x)

    for x != 0 {
        if T.N < NMAX {
            T.tabInt[T.N] = x
            T.N++
        }

        fmt.Scan(&x)
    }
}

func sortArray(T *intArray) {
    var idx int

    for i := 0; i < T.N-1; i++ {
        idx = i

        for j := i+1; j < T.N; j++ {
            if T.tabInt[idx] > T.tabInt[j] {
                idx = j
            }
        }

        T.tabInt[i], T.tabInt[idx] = T.tabInt[idx], T.tabInt[i]
    }
}

func isSimilar(T1, T2 intArray) bool {
    var i int

    if T1.N != T2.N {
        return false
    }

    for i = 0; i < T1.N; i++ {
        if T1.tabInt[i] != T2.tabInt[i] {
            return false
        }
    }

    return true
}
