package main

import "fmt"

func main() {
    var x1, x2, x3 int
    fmt.Scan(&x1, &x2, &x3)

    fmt.Println(f(g(h(x1))))
    fmt.Println(g(h(f(x2))))
    fmt.Println(h(f(g(x3))))
}

func f(x int) int {
    return x * x
}

func g(x int) int {
    return x - 2
}

func h(x int) int {
    return x + 1
}
