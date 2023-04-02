package main

import (
    "fmt"
    "math/rand"
)

const N int = 256

func main() {
    var arr1, arr2 [N]int
    var n int

    arr1 = [N]int{10,20,30,20,10}
    n = 5
    fill(n, &arr2)

    fmt.Println("arr1 =", arr1[:n])
    fmt.Println("arr2 =", arr2[:n])

    reverse(n, &arr1)
    reverse(n, &arr2)

    fmt.Println("arr1 reversed =", arr1[:n])
    fmt.Println("arr2 reversed =", arr2[:n])

    fmt.Println("arr1 is palindrom =", isPalindrom(n, arr1))
    fmt.Println("arr2 is palindrom =", isPalindrom(n, arr2))
}

func fill(n int, arr *[N]int) {
    for i := 0; i < n; i++ {
        arr[i] = rand.Int()
    }
}

func reverse(n int, arr *[N]int) {
    var mid, i int

    if n%2 == 0 {
        mid = n / 2 - 1
    } else {
        mid = n / 2
    }

    for i = 0; i < mid; i++ {
        arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
    }
}

func isPalindrom(n int, arr [N]int) bool {
    var mid, i int

    if n%2 == 0 {
        mid = n / 2 - 1
    } else {
        mid = n / 2
    }

    for i = 0; i < mid; i++ {
        if arr[i] != arr[n-i-1] {
            return false
        }
    }

    return true
}
