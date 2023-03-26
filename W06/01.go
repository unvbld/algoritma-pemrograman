package main

func main() {

}

func abc(x int) string {
    if x == 0 {
        return ""
    }

    if x%2 == 0 {
        return abc(x / 2) + "0"
    }

    return abc(x / 2) + "1"
}
