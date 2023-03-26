package main

import "fmt"

func main() {
    var n, m int

    fmt.Print("Kapasitas rak dan jumlah buku saat ini? ")
    fmt.Scan(&n, &m)

    fmt.Printf("Sisa rak kosong %d buku\n", n-m)
    beliBuku(n, m)
}

func beliBuku(n, m int) {
    if n-m == 1 {
        fmt.Println("Beli 1 buku, rak buku penuh")
    } else {
        fmt.Println("Beli 1 buku, sisa", n-m-1)
        beliBuku(n, m+1)
    }
}
