package main

import "fmt"

func main() {
    var x, tPemasukan, tPengeluaran, tUang float64

    fmt.Scan(&x)
    tUang = x

    membeliKain(x, &tUang, &tPengeluaran)
    membeliBenangJahit(x, &tUang, &tPengeluaran)
    membuatPolaBaju(x, &tUang, &tPengeluaran)
    membuatPolaBaju(x, &tUang, &tPengeluaran)
    menjahitBaju(x, &tUang, &tPengeluaran)
    menjahitBaju(x, &tUang, &tPengeluaran)
    menjahitBaju(x, &tUang, &tPengeluaran)
    menjahitBaju(x, &tUang, &tPengeluaran)
    mengemasBaju(x, &tUang, &tPengeluaran)
    mengemasBaju(x, &tUang, &tPengeluaran)
    mendistribusikan(x, &tUang, &tPemasukan, &tPengeluaran)

    fmt.Printf("%.0f %.0f %.0f\n", tPengeluaran, tPemasukan, tUang)
}

func membeliKain(uangAwal float64, tUang, tPengeluaran *float64) {
    *tPengeluaran += uangAwal / 4
    *tUang -= uangAwal / 4
}

func membeliBenangJahit(uangAwal float64, tUang, tPengeluaran *float64) {
    *tPengeluaran += uangAwal / 20
    *tUang -= uangAwal / 20
}

func membuatPolaBaju(uangAwal float64, tUang, tPengeluaran *float64) {
    *tPengeluaran += uangAwal / 200
    *tUang -= uangAwal / 200
}

func menjahitBaju(uangAwal float64, tUang, tPengeluaran *float64) {
    *tPengeluaran += uangAwal / 200
    *tUang -= uangAwal / 200
}

func mengemasBaju(uangAwal float64, tUang, tPengeluaran *float64) {
    *tPengeluaran += 3 * uangAwal / 200
    *tUang -= 3 * uangAwal / 200
}

func mendistribusikan(uangAwal float64, tUang, tPemasukan, tPengeluaran *float64) {
    *tPemasukan += uangAwal / 2
    *tPengeluaran += 3 * uangAwal / 200
    *tUang += 97 * uangAwal / 200
}
