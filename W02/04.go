package main

import "fmt"

func main() {
    var tahun, total_awal, d1, d34, diskon, total_akhir int

    fmt.Scan(&tahun, &total_awal)

    d1 = tahun / 1000
    d34 = tahun % 100
    diskon = d1 * d34
    total_akhir = total_awal * (100 - diskon) / 100

    fmt.Printf("Besar diskon: %d%%\n", diskon)
    fmt.Printf("Jumlah yang dibayar: %d\n", total_akhir)
}
