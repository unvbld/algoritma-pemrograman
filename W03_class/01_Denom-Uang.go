package main

import "fmt"

func main() {
    var uang, k10, k2, k1 int
    fmt.Scan(&uang)

    pecahUang(uang, &k10, &k2, &k1)

    fmt.Println(k10, "lembar uang 10000")
    fmt.Println(k2, "lembar uang 2000")
    fmt.Println(k1, "lembar uang 1000")
}

func pecahUang(uang int, k10, k2, k1 *int) {
    *k10 = uang / 10000
    *k2 = (uang % 10000) / 2000
    *k1 = ((uang % 10000) % 2000) / 1000
}
