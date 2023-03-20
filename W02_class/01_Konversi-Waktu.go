package main

import "fmt"

func main() {
    var h, m, s, total int
    fmt.Scan(&h, &m, &s)

    total = h*3600 + m*60 + s
    fmt.Printf("Hasil konversi = %d detik\n", total)
}
