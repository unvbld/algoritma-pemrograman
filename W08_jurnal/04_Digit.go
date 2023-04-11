package main

import "fmt"

func main() {
    var n, num1, num2 int

    fmt.Scan(&n)

    split(n, &num1, &num2)

    fmt.Println(num1, num2)
    fmt.Println(num1 + num2)
}

func len(n int) int {
    var l int

    l = 0

    if n == 0 {
        return 1
    }

    for n > 0 {
        n /= 10
        l++
    }

    return l
}

func pangkat(n int) int {
    var x, i int

    x = 1

    for i = 1; i <= n; i++ {
        x *= 10
    }

    return x
}

func split(num int, num1, num2 *int) {
    var mid int

	mid = pangkat(len(num) / 2)

    *num1 = num / mid
    *num2 = num % mid
}
