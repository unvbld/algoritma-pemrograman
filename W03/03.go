package main

import "fmt"

func main() {
    var g string
    var jr, jl, gaji, total_gaji int
    fmt.Scan(&g, &jr, &jl)

    total_gaji = 0

    for g != "Z" {
        if g == "A" {
            gaji = (jr * 75000) + (jl * 90000)
            total_gaji += gaji
            fmt.Println(gaji)
        } else if g == "B" {
            gaji = (jr * 125000) + (jl * 70000)
            total_gaji += gaji
            fmt.Println(gaji)
        } else {
            fmt.Println("Golongan tidak dikenali")
        }

        fmt.Scan(&g, &jr, &jl)
    }

    fmt.Println("Biaya yang dikeluarkan perusahaan untuk gaji karyawan:", total_gaji)
}
