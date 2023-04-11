package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)
    display(x, y)
}

func isFaktor(num, x int) bool {
    return num % x == 0
}

func jumlahFaktor(num int, total*int) {
    var i int

    *total = 0

    for i = 1; i < num; i++ {
        if isFaktor(num, i) {
            *total += i
        }
    }
}

func perfect(num int) bool {
    var total int

    jumlahFaktor(num, &total)

    return num == total
}

func display(x, y int) {
    var i int

    fmt.Print("Barisan Perfect Number: ")

    for i = x; i <= y; i++ {
        if perfect(i) {
            fmt.Print(i, " ")
        }
    }

    fmt.Println()
}
