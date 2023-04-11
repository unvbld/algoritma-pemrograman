package main

import "fmt"

const N int = 100

type mahasiswa struct {
    nama        string
    nim         string
    eprt        float64
    semester    int
    ipk         float64
}

type tabMahasiswa [N]mahasiswa

func main() {
    var T tabMahasiswa
    var n int

    isiData(&T, &n)

    fmt.Println(maxEPRT(T, n))
    fmt.Println(minIPK(T, n))
    fmt.Println(rataRataSemester(T, n))
}

func isiData(T *tabMahasiswa, n *int) {
    var t mahasiswa

    *n = 0

    fmt.Scan(&t.nama, &t.nim, &t.eprt, &t.semester, &t.ipk)

    for t.nim != "none" {
        if *n < N {
            *&T[*n] = t
            *n++
        }

        fmt.Scan(&t.nama, &t.nim, &t.eprt, &t.semester, &t.ipk)
    }
}

func maxEPRT(T tabMahasiswa, n int) float64 {
    var max float64
    var i int

    max = T[0].eprt

    for i = 1; i < n; i++ {
        if max < T[i].eprt {
            max = T[i].eprt
        }
    }

    return max
}

func minIPK(T tabMahasiswa, n int) float64 {
    var min float64
    var i int

    min = T[0].ipk

    for i = 1; i < n; i++ {
        if min > T[i].ipk {
            min = T[i].ipk
        }
    }

    return min
}

func rataRataSemester(T tabMahasiswa, n int) float64 {
    var total, i int

    total = 0

    for i = 0; i < n; i++ {
        total = total + T[i].semester
    }

    return float64(total) / float64(n)
}
