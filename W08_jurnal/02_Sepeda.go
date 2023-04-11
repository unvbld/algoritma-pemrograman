package main

import "fmt"

func main() {
    var jam, menit int
    var member bool
    var sewa float64

    fmt.Scan(&jam, &menit, &member)
    hitungSewa(jam, menit, member, &sewa)

    fmt.Println(sewa)
}

func durasi(jam, menit int) int {
    if jam == 0 && menit != 0 {
        return 1
    }

    if menit >= 10 {
        return jam + 1
    }

    return jam
}

func potongan(durasi, tarif int) float64 {
    if durasi > 3 {
        return 0.9
    }

    return 1
}

func tarif(member bool) int {
    if member {
        return 3500
    }

    return 5000
}

func hitungSewa(jam, menit int, member bool, biaya *float64) {
    var d, t int
    var p float64

    d = durasi(jam, menit)
    t = tarif(member)
    p = potongan(d, t)

    *biaya = float64(d * t) * p
}
