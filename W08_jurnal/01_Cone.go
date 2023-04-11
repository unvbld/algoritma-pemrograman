package main

import (
    "fmt"
    "math"
)

func main() {
    var R, r, t, t1 int
    var luasP, luasSK, luasSB float64

    fmt.Scan(&R, &r, &t, &t1)

    hitungLuasSelimut(R, t, &luasSB)
    hitungLuasSelimut(r, t1, &luasSK)

    luasP = luasALas(R) + luasALas(r) + luasSB - luasSK

    fmt.Println(luasSB)
    fmt.Println(luasSK)
    fmt.Println(luasP)
}

func luasALas(r int) float64 {
    return 3.14 * float64(r*r)
}

func garisPelukis(r, t int) float64 {
    return math.Sqrt(float64(r*r + t*t))
}

func hitungLuasSelimut(r, t int, luas *float64) {
    var s float64
    s = garisPelukis(r, t)
    *luas = 3.14 * float64(r) * s
}
