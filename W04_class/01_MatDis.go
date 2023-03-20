package main

import "fmt"

func main() {
    var n1, n2, r1, r2 int
    fmt.Scan(&n1, &n2, &r1, &r2)

    fmt.Println(permutation(n1, r1), combination(n1, r1))
    fmt.Println(permutation(n2, r2), combination(n2, r2))
}

func mencariFaktorial(n int, f *int) {
    var i int

    *f = 1

    for i = 1; i <= n; i++ {
        *f = *f * i
    }
}

func permutation(n int, r int) int {
    var a, b int
    mencariFaktorial(n, &a)
    mencariFaktorial(n-r, &b)
    return a / b
}

func combination(n int, r int) int {
    var a, b, c int
    mencariFaktorial(n, &a)
    mencariFaktorial(n-r, &b)
    mencariFaktorial(r, &c)
    return a / (b * c)
}
