package main

import "fmt"

const NMAX = 37

type tHimpunan struct {
    anggota [NMAX]int
    panjang int
}

func main() {
    var himpunan1, himpunan2 tHimpunan

    fmt.Print("Anggota himpunan 1: ")
    bacaMasukan(&himpunan1)

    fmt.Print("Anggota himpunan 2: ")
    bacaMasukan(&himpunan2)

    fmt.Println("Himpunan 1 = Himpunan 2?", sama(himpunan1, himpunan2))
}

func bacaMasukan(set *tHimpunan) {
    var x int

    set.panjang = 0

    fmt.Scan(&x)

    for !ada(*set, x) {
        set.anggota[set.panjang] = x
        set.panjang++

        fmt.Scan(&x)
    }
}

func ada(set tHimpunan, x int) bool {
    var i int

    for i = 0; i < set.panjang; i++ {
        if x == set.anggota[i] {
            return true
        }
    }

    return false
}

func urut(set *tHimpunan) {
    var i, idx int

    for i = 1; i < set.panjang; i++ {
        idx = i

        for idx > 0 && set.anggota[idx] < set.anggota[idx-1] {
            set.anggota[idx], set.anggota[idx-1] = set.anggota[idx-1], set.anggota[idx]
            idx--
        }
    }
}

func sama(set1, set2 tHimpunan) bool {
    var i int

    if set1.panjang != set2.panjang {
        return false
    }

    for i = 0; i < set1.panjang; i++ {
        if set1.anggota[i] != set2.anggota[i] {
            return false
        }
    }

    return true
}
