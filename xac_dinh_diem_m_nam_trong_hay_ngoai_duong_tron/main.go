package main

import (
    "fmt"
    "math"
)

var (
    xC float64
    yC float64
    xM float64
    yM float64
    R float64
    ketqua float64
    kiemtrakq float64
)

func kiemtra(xC float64, yC float64, xM float64, yM float64, R float64) float64 {
    ketqua = R - math.Sqrt((xM - xC) * (xM * xC) + (yM - yC) * (yM - yC))
    return ketqua
}

func main() {
    fmt.Printf("Nhap vao ban kinh R: ")
    fmt.Scan(&R)

    fmt.Println("\n-----------------------------\n")

    fmt.Printf("Nhap toa do tam C(xC, yC) \n")
    fmt.Printf("xC = ")
    fmt.Scan(&xC)
    fmt.Printf("yC = ")
    fmt.Scan(&yC)

    fmt.Println("\n-----------------------------\n")

    fmt.Printf("Nhap toa do M(xM, yM) \n")
    fmt.Printf("xM = ")
    fmt.Scan(&xM)
    fmt.Printf("yM = ")
    fmt.Scan(&yM)

    fmt.Println("\n-----------------------------\n")

    if (kiemtra(xC, yC, xM, yM, R) > 0) {
       fmt.Printf("Diem M(%g,%g) co ban kinh R = %g nam trong duong tron C(%g,%g)\n", xM, yM, R, xC, yC)
    } else if (kiemtra(xC, yC, xM, yM, R) < 0) {
       fmt.Printf("Diem M(%g,%g) co ban kinh R = %g nam ngoai duong tron C(%g,%g)\n", xM, yM, R, xC, yC)
    } else {
       fmt.Println("M(", xM, ",", yM, ") nam tren duong tron C(", xC, ",", yC, ")")
    }
}
