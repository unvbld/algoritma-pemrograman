package main

import "fmt"

const NMAX int = 1000

type arrayBunga [NMAX]string

func main() {
    var T arrayBunga
    var n int

    isiArray(&T, &n)
    urut(&T, n)
    tampilArray(T, n)
}

func panjang(bunga string) int {
    var i, n int
    var b1, b2 bool

    i = 0
    n = 0

    for bunga[i] != '.' {
        b1 = 'A' <= bunga[i] && bunga[i] <= 'Z'
        b2 = 'a' <= bunga[i] && bunga[i] <= 'z'

        if b1 || b2 {
            n++
        }

        i++
    }

    return n
}

func urut(tabBunga *arrayBunga, n int) {
    var i, idx int

    for i = 1; i < n; i++ {
        idx = i

        for idx > 0 && panjang(tabBunga[idx]) < panjang(tabBunga[idx-1]) {
            tabBunga[idx], tabBunga[idx-1] = tabBunga[idx-1], tabBunga[idx]
            idx--
        }
    }
}

func isiArray(tabBunga *arrayBunga, n *int) {
    var i, n1 int
    var s string

    *tabBunga = arrayBunga{}

    fmt.Scan(&n1)
    fmt.Scan(&s)
    
    for i = 0; i < n1; i++ {
        if i < NMAX {
            tabBunga[i] = s
            *n++
        }

        fmt.Scan(&s)
    }
}

func tampilArray(tabBunga arrayBunga, n int) {
    var i int

    for i = 0; i < n; i++ {
        fmt.Println(tabBunga[i])
    }
}
