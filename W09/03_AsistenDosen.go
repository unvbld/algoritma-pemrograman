package main

import "fmt"

const NMAX int = 2500

type AsDos_T struct {
    kode    string
    nama    string
    nim     int
    mk      string
    jumlah  int
}

type TabAsDos_T [NMAX]AsDos_T

func main() {
    var tabel TabAsDos_T
    var n int
    var mk string

    BacaAsDos(&tabel, &n)
    CetakAsDos(tabel, n, mk)
}

func BacaAsDos(tabel *TabAsDos_T, n *int) {
    var asdos AsDos_T

    fmt.Scan(&asdos.kode, &asdos.nama, &asdos.nim, &asdos.mk, &asdos.jumlah)
    
    for asdos.kode != "XXX" {
        tabel[*n] = asdos
        *n++
        fmt.Scan(&asdos.kode, &asdos.nama, &asdos.nim, &asdos.mk, &asdos.jumlah)
    }
}

func CetakAsDos(tabel TabAsDos_T, n int, mk string) {
    var i int

    for i = 0; i < n; i++ {
        if tabel[i].mk == mk {
            fmt.Println(tabel[i].nama, tabel[i].kode)
        }
    }
}
