package main

import "fmt"

type rectangle struct {
    length, width int
    color string
    property geometry
}

type geometry struct {
    area, perimeter int
}

func main() {
    var persegi rectangle

    isiData(&persegi)
    hitung(&persegi)

    fmt.Println(persegi.property.area, persegi.property.perimeter)
}

func isiData(persegi *rectangle) {
    fmt.Scan(&persegi.length, &persegi.width, &persegi.color)
}

func hitung(persegi *rectangle) {
    var luas, keliling int

    luas = persegi.length * persegi.width
    keliling = 2 * (persegi.length + persegi.width)

    *&persegi.property.area = luas
    *&persegi.property.perimeter = keliling
}
