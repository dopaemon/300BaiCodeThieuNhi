package main

import (
    "fmt"
    "math"
)

var (
    ab float64
    a float64
    b float64
    c float64
    d float64
)

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func khoangcachab(xa float64, xb float64, ya float64, yb float64) float64 {
     ab = math.Sqrt(math.Pow((xb - xa), 2) + math.Pow((yb - ya), 2))
     return ab
}

func main() {
     fmt.Printf("Nhap vao Xa: ")
     fmt.Scan(&a)
     fmt.Printf("Nhap vao Ya: ")
     fmt.Scan(&c)
     fmt.Printf("Nhap vao Xb: ")
     fmt.Scan(&b)
     fmt.Printf("Nhap vao Yb: ")
     fmt.Scan(&d)
     fmt.Println("|AB| = Sqrt(Pow((", a, "-", b, ") + (", c, "-", d,"))")
     fmt.Println("|AB| = ", toFixed(khoangcachab(a, b, c, d), 4))
}
