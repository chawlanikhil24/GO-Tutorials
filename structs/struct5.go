`package main

import "fmt"

type Student struct {
Name        string
Class       string
Age         int
RollNo      int
Father_Name string
Mother_Name string
Locality    string
Pincode     int
}

func main() {
fmt.Println("")
Object7 := &Student{"abcd", "7th", 21, 34, "XYZ", "pqrs", "Westwood", 110022}
fmt.Println("")
fmt.Println("Object7: ", Object7)
Object8 := new(Student)
fmt.Println("")
fmt.Println("Object8: ", Object8)
Object8.Name = "abcd"
Object8.Class = "7th"
Object8.Pincode = 110022
Object8.RollNo = 34
Object8.Age = 21
Object8.Father_Name = "XTZ"
Object8.Mother_Name = "pqrs"
Object8.Locality = "Westwood"
fmt.Println("")
fmt.Println("Object8: ", Object8)
fmt.Println("")
}
`
