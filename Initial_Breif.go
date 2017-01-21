package main //we Just Imported the main package
import (
    "fmt"
    "strconv"
    ) // we are import the "format" from the main package

func main()  { //Just like C,C++,JAVA, Python and other languages, here also we begins with main function
    fmt.Println("Hi Buddy! I am Your GoBuddy... :)")
    /*    The game begins If we know all the available Data Types
          There are multiple ways of declaring data types
          If you are habitual of using ; after end of a line, You can use, Go has No problems with that. (AND I FORGOT, THIS IS A MULTILINE COMMMENT)*/
    // ret := justChecking()
    // fmt.Println(ret)
    // learnDataTypes()
    // playWithFORLOOP()
    // anonymousFuntionExample()
    // funcWithParameter(456) //456 is the function parameter
    // ret2 := funcWith3Parameter(10,11,12)
    // fmt.Println("The result of funcWith3Parameter() is: ",ret2)
    // val1,val2 := funtionReturning2values()
    // fmt.Println("VAL1: ",val1,",VAL2: ",val2)
    // practiceArray()
    // practiceSlices()
    // use_defer()
    // how_to_RECOVER()
    // panicAndRecover()
    // learnPointers()
    // ---Uncomment this part to learn the Implementation of structures and their associated functions--------------
            // rec1 := Rectangle{10,20,10,20}
            // //We just created an object named rec of Rectangle structure. Through this object we can access the structure data items
            // fmt.Println("Side1 is: ",rec1.side1)
            // fmt.Println("Side2 is: ",rec1.side2)
            // fmt.Println("Side3 is: ",rec1.side3)
            // fmt.Println("Side4 is: ",rec1.side4)
            // fmt.Println("The area of the Rectangle is: ",rec1.Area() )
    // ------------------------------------------------------------------------------------------------------------
}

type Rectangle struct{
    side1 int
    side2 int
    side3 int
    side4 int
} //We just defined a structure named Rectangle with 4 sides

// If you understand the above sturucture, Now lets create an "Rectangle structure" associated funtion
func (r Rectangle) Area() int{
    return r.side1 * r.side2
}

func learnPointers(){
    /* Pointers has always been tough to understand and implement for me. But now when I developed a certain method to understand them */
    valX := 123
    fmt.Println("Initial Value of X: ",valX)
    changeValUsingReference(&valX) //Implementation of pointers need to send address of variable not the variable
    fmt.Println("Function manipulated Value of X: ",valX)
    // Another Experiment to do it
    addresX := &valX
    changeValUsingReference(addresX)
    fmt.Println("Function manipulated Value of X: ",valX)
}

func changeValUsingReference(x *int){   //The parameter must be a pointer to accept address
    fmt.Println("Location of X in memory: ",x)
    *x = *x * 2     // "  *= content of  " always consider * as the mentioned value, you will never jumble again
    // I just changed the content of memory address, which was earlier 0  and now 213
}

func panicAndRecover(){
    // This is GOLANG style of TRY AND CATCH
        defer func(){
            fmt.Println(recover())
        }()
        panic("I m PANICCCCC!!, Sending Instructions to recover function")
}

func how_to_RECOVER(){
    fmt.Println(perfromDivision1(100,0))
    fmt.Println(perfromDivision1(100,1))
}

func perfromDivision1(n1 int, n2 int)string{
    defer func(){
        fmt.Println(recover())
        fmt.Println("I am the saviour of this program, I didn't let the program stop :)")
    }()
    res := n1/n2
    fmt.Println(res)
    r := "The result of the Division is: "+strconv.Itoa(res) // I just converted Int64 to Sting, strconv is imported for this purpose
    return r
    /*After the Execution of this function, you must have observed, that "Runtime error" must have occured, but with the use of "recover"
    go routine, the flow of program didn't stop and continued its operation. So recover can be used wherever there is a
    chance of occurence of Error*/
}

func use_defer()  {
    /* The defer statement in  go helps to run a function at the last when all other funtions in this block has executed.
     Consider and Example to understand this, when we get some books issued from library,which are of our interest.
     Now, we performed the tasks, of 1) Reading 2) understanding 3) Creating Notes for future.
     After all the tasks executed, the final task is to return the book(This is a clean Up task).
     This is why defer is used, to execute the "CLEAN UP TASK".*/
     defer playWithFORLOOP()
     fmt.Println("-----------------------------------------------------------------")
     learnDataTypes()
     fmt.Println("-----------------------------------------------------------------")
     ret := justChecking()
     fmt.Println(ret)
     fmt.Println("-----------------------------------------------------------------")
    // playWithFORLOOP will execute at last
    // learnDataTypes will execute first since its the first in queue Now
    // justChecking will execute second
}

func practiceSlices()bool{
    // Slices can be considered as subpart of the Array
    exampleArray := []int{1,2,3,4,5,6}
    // The above used array can be called a dynamic array
    slice1 := exampleArray[2:4]
    slice2 := exampleArray[2:]
    fmt.Println(slice1)
    fmt.Println(slice2)
    slice3 := make([]int,5,10)
    for i,value := range slice3{
        fmt.Println(value,i)
    }
    return true
}

func practiceArray()bool {
    //Below this line suppose to be the decalaration of an array
    // var <array_name><[size of array]> <data-type> of elements of array
    var exampleArray[5] int
    // naive assignment method which we all are familiar with
     exampleArray[0] = 1
     exampleArray[1] = 12
     exampleArray[2] = 123
     exampleArray[3] = 234
     exampleArray[4] = 345
     fmt.Println("--------------------------------------------------")
     fmt.Println(exampleArray)
     exampleArray2 := [5]float32{1,2,3,4,5}
     fmt.Println("--------------------------------------------------")
     fmt.Println(exampleArray2)
    // How to traverse in an array
    fmt.Println("--------------------------------------------------")
     for i,value := range exampleArray2{
         fmt.Println(value,i) //You must have observed that there were indexes as well,
     }
     fmt.Println("--------------------------------------------------")
    //  Use this traversal if you don't want the index to be printed
    for _,value := range exampleArray2{
        fmt.Println(value) //You must have observed that there were indexes as well,
    }
    return true
}

func learnDataTypes()bool  {
    var int_EX1 uint8 = 10  //First way to declare data type, UNSIGNED 8-BIT INT
    const int_EX2 uint16 = 10 //CONSTANT UNSIGNED 16-BIT INT
    var int_EX3 uint32 = 10 //UNSIGNED 32-BIT INT
    var int_EX4 uint64 = 11 //UNSIGNED 64-BIT INT
    var int_EX5 int8 = 10   //SIGNED 8-BIT INT
    var int_EX6 int16 = 10  //SIGNED 16-BIT INT
    var int_EX7 int32 = 10  //SIGNED 32-BIT INT
    var int_EX8 int64 = 10  //SIGNED 64-BIT INT
    var int_EX9 int = 10   //value is 32-BIT on 32-BIT system, mine 64
    var int_EX10 uint = 10 //value is 32-BIT on 32-BIT system, mine 64
    var int_EX11 uintptr = 10 //value is 32-BIT on 32-BIT system, mine 64
    var EX1 byte = 11 // alias for uint8
    var EX2 rune = 89 //alias for int32
    int_EX12 := 11 //DEFAULT data type  SIGNED 64-BIT INT
    var float_EX1 float32 = 10
    var float_EX2 = 1.22    //The we are used to declare variables #DEFAULT
    var bool_EX bool = true //BOOL Example
    var string_EX = "Hi.. This is being Done by Nikhil Chawla" //Simple string data-type
    float_EX3 := 26666666666666666666666666666666.33 //Here we are moving from right to left,variable data type is decided by the data on the right
    fmt.Println("Following are the data types")
    fmt.Printf("%T : %d\n",int_EX1,int_EX1)
    fmt.Printf("%T : %d\n",int_EX2,int_EX2)
    fmt.Printf("%T : %d\n",int_EX3,int_EX3)
    fmt.Printf("%T : %d\n",int_EX4,int_EX4)
    fmt.Printf("%T : %d\n",int_EX5,int_EX5)
    fmt.Printf("%T : %d\n",int_EX6,int_EX6)
    fmt.Printf("%T : %d\n",int_EX7,int_EX7)
    fmt.Printf("%T : %d\n",int_EX8,int_EX8)
    fmt.Printf("%T : %d\n",int_EX9,int_EX9)
    fmt.Printf("%T : %d\n",int_EX10,int_EX10)
    fmt.Printf("%T : %d\n",int_EX11,int_EX11)
    fmt.Printf("%T : %d\n",EX1,EX1)
    fmt.Printf("%T : %d\n",EX2,EX2)
    fmt.Printf("%T : %d\n",int_EX12,int_EX12)
    fmt.Printf("%T : %f\n",float_EX1,float_EX1)
    fmt.Printf("%T : %f\n",float_EX2,float_EX2)
    fmt.Printf("%T : %f\n",float_EX3,float_EX3)
    fmt.Printf("%T : %v\n",string_EX,string_EX)
    fmt.Printf("%T : %v\n",bool_EX,bool_EX)
    return true
}

//I just proved that I was compiled not Interpreted .. ;)
func justChecking()string {
    fmt.Println("YO .. man !.. I am A userdefined function ,, ready to GO")
    return "The function is working"
}

func playWithFORLOOP()bool{
    //2 WAYS TO EXECUTE
    //1. TYPICAL WAY
    for i := 1; i <= 10; i += 1 {
        if (i % 2 == 0) || (i % 3 == 0){ //LOGICAL OR
            fmt.Println("Condition 1")
            fmt.Println(i)
        } else if (i % 1 == 0) && (i % 2 == 0){ // LOGICAL AND
            fmt.Println("Condition 2")
            fmt.Println(i)
        } else {
            fmt.Println("Condition DEFAULT")
            fmt.Println("No If else is satisfied")
        }

    }
    // 2. LIKE WHILE LOOP in C,C++,JAVA,PYTHON
    i := 1
    for i <=10{
        fmt.Println(i)
        i += 1
    }
    return true
}

func anonymousFuntionExample()bool  {
    //IF YOU've ever done javascript, You must be familiar about the Anonymous function
    num := 5
    anonyEx := func ()int  {
            num = num*2
            return num
    }
    fmt.Println("OUTPUT OF THE ANONYMOUS FUNCTION: ",anonyEx)
    return true
}

func funcWithParameter(para int)bool {
    fmt.Println("The parameter passed to the function is: ",para)
    return true
}

func funcWith3Parameter(para1 int,para2 int,para3 int,)int  {
    fmt.Println("The 3 parameters are: ",para1,para2,para3)
    // NOW I AM THINKING TO RETURN THE PRODUCT OF THESE NUMBERS
    return para1*para2*para3
}

func funtionReturning2values()(string,int)  {
    // DON'T FORGET TO CHANGE AND INCREASE THE RETURN TYPES
    ret1 := "My AGe is"
    ret2 := 22
    return  ret1, ret2
}
