package main

import (
    "fmt"
    "math"
)

// Start variable
var (
    dientich float64
    pi float64 = 3.141593
    v float64
)
// End varable

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func thetich(dientich float64) float64 {
    v = ((4 * pi) / 3 ) * math.Pow((math.Sqrt(dientich / (4 * pi))), 3)
    return v
}

func main() {
    fmt.Printf("Nhap vao dien tich mat cau: ")
    fmt.Scan(&dientich)
    fmt.Println("The Tich Cua Mat Cau La:", toFixed(thetich(dientich), 6) )
}
