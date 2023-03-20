package main

import "fmt"

func main() {
    var w, h, skala int
    var s string
    fmt.Scan(&w, &h)
    fmt.Scan(&s, &skala)

    if s == "+" {
        zoomIn(skala, &w, &h)
    } else if s == "-" {
        zoomOut(skala, &w, &h)
    }

    fmt.Println(w, h)
}

func zoomIn(skala int, p, l *int) {
    *p = *p * skala
    *l = *l * skala
}

func zoomOut(skala int, p, l *int) {
    *p = *p / skala
    *l = *l / skala
}
