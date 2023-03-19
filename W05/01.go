package main

import "fmt"

func main() {
    var tgl1, bln1, thn1 int
    var tgl2, bln2, thn2 int

    inputTglPinjam(&tgl1, &bln1, &thn1)

    for valid(tgl1, bln1, thn1) {
        hitungTanggalKembali(tgl1, bln1, thn1, &tgl2, &bln2, &thn2)
        fmt.Println(tgl2, bln2, thn2)
        inputTglPinjam(&tgl1, &bln1, &thn1)
    }

    fmt.Println("input tidak valid")
}

func inputTglPinjam(tanggal, bulan, tahun *int) {
    fmt.Scan(tanggal, bulan, tahun)
}

func valid(tanggal, bulan, tahun int) bool {
    var jumlahHari int
    getJumlahHari(bulan, tahun, &jumlahHari)
    return tanggal >= 1 && tanggal <= jumlahHari && bulan >= 1 && bulan <= 12
}

func kabisat(tahun int) bool {
    return (tahun % 4 == 0 && tahun % 100 != 0) || (tahun % 400 == 0)
}

func getJumlahHari(bulan, tahun int, jmlHari *int) {
    if bulan == 2 {
        if kabisat(tahun) {
            *jmlHari = 29
        } else {
            *jmlHari = 28
        }
    } else {
        if bulan <= 7 {
            *jmlHari = 30 + (bulan % 2)
        } else {
            *jmlHari = 31 - (bulan % 2)
        }
    }
}

func hitungTanggalKembali(tanggal1, bulan1, tahun1 int, tanggal2, bulan2, tahun2 *int) {
    var jumlahHari int

    getJumlahHari(bulan1, tahun1, &jumlahHari)

    if tanggal1 + 3 <= jumlahHari {
        *tanggal2 = tanggal1 + 3
        *bulan2 = bulan1
        *tahun2 = tahun1
    } else {
        *tanggal2 = tanggal1 + 3 - jumlahHari
        *bulan2 = (bulan1 % 12) + 1
        *tahun2 = tahun1 + (bulan1 / 12)
    }
}
