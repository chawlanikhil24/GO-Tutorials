package main

import (
    "fmt"
    "math"
)

type Rectangle struct {
    length, width float64
}

type Circle struct{
    radius float64
}

type Triangle struct{
    base, height float64
}

type Shape interface{
    area() float64 //interface links the functions with their respective structures
    // Interface only allows to declare the functions to be polymorphed
    // Don't Try to Define any functions
}

// The POLYMORPHISM begins here

func (r Rectangle) area() float64 {
    return r.length * r.width
}

func (c Circle) area() float64{
    return math.Pi * math.Pow(c.radius,2)
}

func (t Triangle) area() float64 {
    return 0.5 * t.base * t.height
}

func main() {
    fmt.Println("\n")
    r1 := Rectangle{4.1, 3.3}
    fmt.Println("Rectangle is: ", r1)
    fmt.Println("Rectangle area is: ", r1.area())
    fmt.Println("\n")
    fmt.Println("\n")
    c1 := Circle{10.0}
    fmt.Println("Radius of Cirlce is: ", c1)
    fmt.Println("Area is: ", c1.area())
    fmt.Println("\n")
    fmt.Println("\n")
    t1 := Triangle{44.2,10.5}
    fmt.Println("Triangle is: ", t1)
    fmt.Println("Triangle area is: ", t1.area())
}
