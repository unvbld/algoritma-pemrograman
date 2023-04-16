package main

import "fmt"

const NMAX int = 2022

type student struct {
    name    string
    sid     string
    gpa     float64
}

type tabMhs [NMAX]student

func main() {
    var T tabMhs
    var n int

    T = tabMhs{
        student{"113210689", "Harith", 1.56},
        student{"113212624", "Johnson", 3.19},
        student{"113211834", "Kimmy", 1.32},
        student{"113212925", "Chou", 3.68},
        student{"113210520", "Grock", 1.45},
        student{"113210223", "Lunox", 1.89},
        student{"113212819", "Karrie", 1.05},
        student{"113211273", "Aldous", 2.46},
        student{"113211643", "Franco", 1.6},
        student{"113211992", "Selena", 3.5},
    }

    n = 10

    cetakGPA(T, n)
    sortGPA(&T, n)
    cetakGPA(T, n)
}

func sortGPA(T *tabMhs, n int) {
    var idx, i, j int

    for i = 0; i < n-1; i++ {
        idx = i

        for j = i+1; j < n; j++ {
            if T[idx].gpa > T[j].gpa {
                idx = j
            }
        }

        T[i], T[idx] = T[idx], T[i]
    }
}

func cetakGPA(T tabMhs, n int) {
    var i int
    
    for i = 0; i < n; i++ {
        fmt.Print(T[i].gpa, " ")
    }

    fmt.Println()
}
