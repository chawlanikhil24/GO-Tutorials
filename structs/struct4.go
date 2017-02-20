package main

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

func (object Student) showDetails() {
	fmt.Println("The details entered are:")
	fmt.Println("Name", object.Name)
	fmt.Println("Class", object.Class)
	fmt.Println("Age", object.Age)
	fmt.Println("RollNo", object.RollNo)
	fmt.Println("Father's Name", object.Father_Name)
	fmt.Println("Mother's Name", object.Mother_Name)
	fmt.Println("Locality", object.Locality)
	fmt.Println("Pincode", object.Pincode)
}

func main() {
	var Object1 Student
	Object2 := Student{}
	fmt.Println("Object1: ", Object1)
	fmt.Println("")
	fmt.Println("Object2: ", Object2)
	fmt.Println("")
	Object2.Name = "abcd"
	Object2.Class = "7th"
	Object2.Pincode = 110022
	Object2.RollNo = 34
	Object2.Age = 21
	Object2.Father_Name = "XTZ"
	Object2.Mother_Name = "pqrs"
	Object2.Locality = "Westwood"
	fmt.Println("Object2: ", Object2)
	fmt.Println("")
	Object3 := Student{
		Name:        "abcd",
		Class:       "7th",
		Pincode:     110022,
		RollNo:      34,
		Age:         21,
		Father_Name: "XYZ",
		Mother_Name: "pqrs",
		Locality:    "Westwood",
	}
	fmt.Println("Object3:", Object3)
	fmt.Println("")
	Object4 := Student{"abcd", "7th", 21, 34, "XYZ", "pqrs", "Westwood", 110022}
	fmt.Println("Object4:", Object4)
	fmt.Println("")
	Object5 := Student{"Westwood", "7th", 110022, 34, "XYZ", "pqrs", "abcd", 21}
	fmt.Println("")
	fmt.Println("Object5.Age:", Object5.Age)
	fmt.Println("")
	Object5.showDetails()
	Object6 := new(Student)
	Object6.Name = "abcd"
	Object6.Class = "7th"
	Object6.Pincode = 110022
	Object6.RollNo = 34
	Object6.Age = 21
	Object6.Father_Name = "XTZ"
	Object6.Mother_Name = "pqrs"
	Object6.Locality = "Westwood"
	fmt.Println("Object6", Object6)
}
