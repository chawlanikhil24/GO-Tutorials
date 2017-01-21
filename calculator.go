package main

import (
    "fmt"
    "math"
)

func main()  {
    var num1,num2 int
    fmt.Scanf("%d",&num1)
    fmt.Scanf("%d",&num2)
    fmt.Println("Num1:",num1,"Num2:",num2)
    fmt.Print("\n")
    fmt.Println("Sum of Given Numbers: ",Add(num1,num2))
    fmt.Println("Subtraction of Given Numbers: ",Sub(num1,num2))
    fmt.Println("Product of Given Numbers: ",Mul(num1,num2))
    fmt.Println("Division of Given Numbers: ",Div(num1,num2))
    fmt.Println("Squares of the given Numbers: ",math.Pow(float64(num1),2),math.Pow(float64(num2),2))
    fmt.Println("Square root of the Given Numbers: ",math.Sqrt(float64(num1)),math.Sqrt(float64(num2)))
    fmt.Println("Using Exponential Calculating num1 ^ num2:",math.Pow(float64(num1),float64(num2)))
    fmt.Println("Log of the Num1 and Num2: ",math.Log(float64(num1)),math.Log(float64(num2)))
}

func Add(a,b int)int{
    return a+b
}

func Sub(a,b int)int{
    if a > b{
        return a-b
    }else{
        return b-a
    }
}

func Mul(a,b int)int{
    return a*b
}

func Div(a,b int)float64{
    defer func(){
        recover()
    }()
    sol := float64(a/b)
    return sol
}
