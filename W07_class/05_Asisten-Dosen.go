package main

import "fmt"

const N int = 2022

type tabNasabah struct {
    kode, nasabah, bank, rekening string
}

func main() {
    var T [N]tabNasabah
    var x string
    var n, i int

    // banyak data nasabah yang tersimpan
    n = 0

    for i = 0; i < 10; i++ {
        addNasabah(&T, &n)
    }

    fmt.Scan(&x)
    cetak(T, n, x)
}

func addNasabah(T *[N]tabNasabah, n *int) {
    var t tabNasabah

    fmt.Scan(&t.kode, &t.nasabah, &t.bank, &t.rekening)

    if *n < N {
        *&T[*n] = t
        *n++
    } else {
        fmt.Println("data penuh")
    }
}

func cetak(T [N]tabNasabah, n int, x string) {
    var i int

    for i = 0; i < n; i++ {
        if T[i].bank == x {
            fmt.Printf("Kode: %s, ", T[i].kode)
            fmt.Printf("Nasabah: %s, ", T[i].nasabah)
            fmt.Printf("Bank: %s, ", T[i].bank)
            fmt.Printf("Rek: %s\n", T[i].rekening)
        }
    }
}
