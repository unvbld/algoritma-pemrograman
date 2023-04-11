package main

import "fmt"

func main() {
    var m, jumlah, menu, total int
    var bersisa bool

    fmt.Scan(&m)

    for m > 0 {
        fmt.Scan(&menu, &jumlah, &bersisa)

        hitungTarif(menu, bersisa, jumlah, &total)
        fmt.Println(total)

        m--
    }
}

func tarif(menu int) int {
    if menu <= 3 {
        return 10000
    }

    if menu <= 50 {
        return menu * 2500
    }

    return 100000
}

func hitungTarif(menu int, bersisa bool, n int, total *int) {
    if bersisa {
        *total = tarif(menu) * n
    } else {
		*total = tarif(menu)
	}
}
