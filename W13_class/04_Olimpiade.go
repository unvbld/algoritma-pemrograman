package main

import "fmt"

const NMAX int = 100

type peserta struct {
    tim string
    g   int
    s   int
    b   int
}

type olimpiade [NMAX]peserta

func main() {
    var tab olimpiade
    var n int

    isiArray(&tab, &n)
    sorting(&tab, n)
    tampilArray(tab, n)
}

func isiArray(T *olimpiade, n *int) {
    var i int
    var t peserta

    *T = olimpiade{}

    fmt.Scan(n)
    fmt.Scan(&t.tim, &t.g, &t.s, &t.b)

    for i = 0; i < *n; i++ {
        if i < NMAX {
            T[i] = t
        }

        fmt.Scan(&t.tim, &t.g, &t.s, &t.b)
    }

    if *n > NMAX {
        *n = NMAX
    }
}

func tampilArray(T olimpiade, n int) {
    var i int

    for i = 0; i < n; i++ {
        fmt.Println(T[i].tim, T[i].g, T[i].s, T[i].b)
    }
}

func poin1(g, s, b int) int {
    return 4*g + 3*s + b
}

func poin2(p peserta) int {
    return (4 * p.g) + (3 * p.s) + p.b
}

func sorting(T *olimpiade, n int) {
    var i, idx int

    for i = 1; i < n; i++ {
        idx = i

        for idx > 0 && poin2(T[idx]) > poin2(T[idx-1]) {
            T[idx], T[idx-1] = T[idx-1], T[idx]
            idx--
        }
    }
}
