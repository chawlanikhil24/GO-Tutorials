package main

import "fmt"

//THIS FUNCTION ACCEPTS STRING ARGUMENT IN PARAMETER AND RETURNS A STRING ARRAY
func createList(arg string) []string {
	var retObject = []string{arg}
	return retObject
}

//THIS FUNCTION ACCEPTS STRING ARGUMENTS IN PARAMETERS AND RETURNS A MAP with KEY=STRINGS, VAL=STRINGS
func mapCreate(name, rollNo, height, weight, averageMarks, extraCurricularGrade string) map[string]string {
	var mapRet = map[string]string{
		"name":              name,
		"roll_no":           rollNo,
		"weight":            weight,
		"height":            height,
		"avg_marks":         averageMarks,
		"extra_curr_grades": extraCurricularGrade,
	}
	return mapRet
}

//THIS FUNCTION ACCEPTS STRING ARGUMENTS IN PARAMETERS AND RETURNS A MAP with KEY=STRINGS, VAL=STRING ARRAY
func mapCreateList(name, rollNo, height, weight, averageMarks, extraCurricularGrade string) map[string][]string {
	var mapRet = map[string][]string{
		"name":              createList(name),
		"roll_no":           createList(rollNo),
		"weight":            createList(weight),
		"height":            createList(height),
		"avg_marks":         createList(averageMarks),
		"extra_curr_grades": createList(extraCurricularGrade),
	}
	return mapRet
}

func main() {
	kanika := mapCreate("kanika", "4", "5'4", "52", "85", "B")
	saurabh := mapCreate("saurabh", "5", "5'9", "50", "76", "C")
	deepti := mapCreate("deepti", "6", "5'6", "59", "83", "B")
	//CREATED 3 STUDENT RECORDS AS MAP
	var studentEX1 = []map[string]string{kanika, saurabh, deepti}
	//STORED THESE MAPS IN AN ARRAY
	for _, value := range studentEX1 {
		fmt.Println(" ")
		fmt.Println(value)
		fmt.Println("KEY", "VALUE")
		for key, val := range value {
			fmt.Println(key, val)
		}
	}

	nikhil := mapCreateList("nikhil", "1", "5'11", "78", "80", "A")
	sparsh := mapCreateList("sparsh", "2", "5'9", "60", "79", "B")
	shivangi := mapCreateList("shivangi", "3", "5'10", "60", "81", "A")
	//CREATED 3 STUDENT RECORDS AS MAP with values= ARRAY
	var studentEX2 = []map[string][]string{nikhil, sparsh, shivangi}
	//STORED THESE MAPS IN AN ARRAY
	for _, value := range studentEX2 {
		fmt.Println(" ")
		fmt.Println(value)
		fmt.Println("KEY", "VALUE")
		for key, value := range nikhil {
			fmt.Println(key, value)
		}
	}
}

//AT THE END WE IMPLEMENTED A JSON LIKE STRUCTURE IN GO
