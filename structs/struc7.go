package main

import "fmt"

type Sofa struct{}
type Table struct{}
type Chimney struct{}
type Knife struct{}

type Home interface {
	purpose() string
}

func (S Sofa) purpose() string {
	return "I am Sofa, I am one of the finest and relaxing object on which you can sit "
}

func (T Table) purpose() string {
	return "I am a table, used to keep plates, trays and other occassional objects"
}

func (C Chimney) purpose() string {
	return "I am a chimney, used to exhaust heat produced in kitchen while cooking"
}

func (K Knife) purpose() string {
	return "I am a KNIFE, used to cut any object"
}

func main() {
	sofa := Sofa{}
	table := Table{}
	chimney := Chimney{}
	knife := Knife{}
	fmt.Println(sofa.purpose())
	fmt.Println(table.purpose())
	fmt.Println(chimney.purpose())
	fmt.Println(knife.purpose())
}
