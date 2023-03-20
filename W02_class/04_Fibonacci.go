package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
    var ans, s1, s2, i, tmp int

    ans = 0
    s1 = 0
    s2 = 1

    for i = 0; i < n; i++ {
        ans += s1
        tmp = s1 + s2
        s1 = s2
        s2 = tmp
    }

    return ans
}
