package main
import "fmt"

func main() {
    var i float64
    var m int
    var p bool

    fmt.Scan(&i, &m, &p)

    if cumlaude(i, m, p) {
        fmt.Println("Cumlaude")
    } else if sangatMemuaskan(i, m, p) {
        fmt.Println("Sangat memuaskan")
    } else if memuaskan(i, m, p) {
        fmt.Println("Memuaskan")
    }
}

func cumlaude(ipk float64, masaStudi int, publikasi bool) bool {
    return ipk >= 3.51 && ipk <= 4.0 && masaStudi <= 8 && publikasi
}

func sangatMemuaskan(ipk float64, masaStudi int, publikasi bool) bool {
    return ipk >= 2.76 && ipk <= 3.5 && masaStudi <= 14
}

func memuaskan(ipk float64, masaStudi int, publikasi bool) bool {
    return ipk >= 2.0 && ipk <= 2.75 && masaStudi <= 14
}
