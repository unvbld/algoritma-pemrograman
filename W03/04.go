package main

import "fmt"

func main() {
    var Uang, TotalTabungan, maxUang, minUang, Jumlah int
    var rata float64

    fmt.Scan(&Uang)
    Jumlah = 0
    TotalTabungan = 0
    maxUang = Uang
    minUang = Uang

    for Uang > 0 {
        TotalTabungan += Uang

        if Uang > maxUang {
            maxUang = Uang
        }

        if Uang < minUang {
            minUang = Uang
        }

        Jumlah = Jumlah + 1
        fmt.Scan(&Uang)
    }

    rata = float64(TotalTabungan) / float64(Jumlah)

    fmt.Println("Jumlah =", TotalTabungan)
    fmt.Println("Rata - rata =", rata)
    fmt.Println("Uang tertinggi =", maxUang)
    fmt.Println("Uang terendah =", minUang)
}
