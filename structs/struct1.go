package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name       string
	age        int
	gender     string
	city       string
	profession string
	experience int
	pincode    int
}

func (obj *Person) details() {
	fmt.Println("")
	fmt.Println("Hi!.. My name is: ", obj.name)
	fmt.Println("My age is: ", obj.age)
	fmt.Println("My gender is: ", obj.gender)
	fmt.Println("City I reside in: ", obj.city)
	fmt.Println("My profession is: ", obj.profession)
	if strings.ToLower(obj.profession) == "student" {
		if obj.experience != 0 {
			fmt.Println("Bitch Please! being a student your experience is not admissible, Wrap up with your course first.")
		}
	} else {
		fmt.Println("My experience in mu profession is: ", obj.experience)
	}
	fmt.Println("Postal Address is: ", obj.pincode)
}

func main() {
	person1 := Person{"Nikhil Chawla", 22, "MALE", "Delhi", "Student", 0, 110093}
	person2 := Person{"Saurabh", 21, "MALE", "Delhi", "Student", 1, 110092}
	person3 := Person{"KAnika Murarka", 21, "Female", "Bengaluru", "Intern", 11, 323322}
	fmt.Println(person1, person2, person3)
	person1.details()
	person2.details()
	person3.details()
	person4 := Person{}
	person4.details()
}
