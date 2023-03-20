package main

import "fmt"

func main() {
    var x, y, p int
    fmt.Scan(&x, &y)

    if x >= y {
        p = permutasi(x, y)
    } else {
        p = permutasi(y, x)
    }

    fmt.Println(faktorial(x), faktorial(y), p)
}

func faktorial(n int) int {
    var ans, i int
    ans = 1

    for i = 1; i <= ans; i++ {
        ans *= i
    }

    return ans
}

func permutasi(n int, r int) int {
    var ans, i int
    ans = 1

    for i = n; i >= n-r+1; i-- {
        ans *= i
    }

    return ans;
}
