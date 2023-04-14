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

func poin(g, s, b int) int {
    return 4*g + 3*s + b
}

func sorting(T *olimpiade, n int) {
    var i, idx, poin1, poin2 int

    for i = 1; i < n; i++ {
        idx = i
        poin1 = poin(T[idx].g, T[idx].s, T[idx].b)
        poin2 = poin(T[idx-1].g, T[idx-1].s, T[idx-1].b)

        for idx > 0 && poin1 > poin2 {
            T[idx], T[idx-1] = T[idx-1], T[idx]

            idx--

            if idx > 0 {
                poin1 = poin(T[idx].g, T[idx].s, T[idx].b)
                poin2 = poin(T[idx-1].g, T[idx-1].s, T[idx-1].b)
            }
        }
    }
}
