package main

import "fmt"

const N int = 100

type tabBunga [N]string

func main() {
    var n int
    var T tabBunga

    fmt.Scan(&n)
    isiData(&T, n)

    fmt.Println(T)
    renameBunga(&T, n, "lily", "tulip")
    fmt.Println(T)
    deleteBunga(&T, n, "tulip")
    fmt.Println(T)
}

func isiData(T *tabBunga, n int) {
    var i int
    var s string

    for i = 0; i < n; i++ {
        fmt.Scan(&s)

        if i < N {
            *&T[i] = s
        }
    }
}

func indexBunga(T tabBunga, n int, nama string) int {
    var idx int

    idx = 0

    for idx < n {
        if T[idx] == nama {
            return idx
        }

        idx++
    }

    return -1
}

func renameBunga(T *tabBunga, n int, bungaLama, bungaBaru string) {
    var idx int

    idx = indexBunga(*T, n, bungaLama)

    if idx != -1 {
        *&T[idx] = bungaBaru
    }
}

func deleteBunga(T *tabBunga, n int, bunga string) {
    var idx int

    idx = indexBunga(*T, n, bunga)

    if idx != -1 {
        for idx < n - 1 {
            *&T[idx] = *&T[idx+1]
            idx++
        }

        *&T[n] = ""
    }
}
