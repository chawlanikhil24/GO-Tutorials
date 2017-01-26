package main

import "fmt"

func main() {
    for i:=1; i <= 25; i += 1{
        //LOGICAL OPERATORS
        if (i <= 10) && (i >=5){ //conditional AND
            fmt.Println("IMPLEMENTING conditional AND")
            fmt.Println(i)
        }else if (i > 20) || (i < 10){ //conditional OR
            fmt.Println("IMPLEMENTING conditional OR")
            fmt.Println(i)
        }else if (!false){
            fmt.Println("IMPLEMENTING conditional OR")
            fmt.Println("This is conditional NOT")
        }
    }
    fmt.Println("--------------------------------------------------\n")
    for i:=1; i <= 11; i += 1{
        //COMPARISON OPERATORS
        if i==10 {
            fmt.Println("used comparison operator:","EQUAL TO\n",i )
            } else if i!=10{
            fmt.Println("used comparison operator:","NOT EQUAL TO\n",i )
            if i <= 10{
                fmt.Println("used comparison operator:","LESS THAN EQUAL TO\n",i)
            } else if i > 8{
                fmt.Println("used comparison operator:","LESS THAN EQUAL TO\n",i)
            }
        }
    }

    fmt.Println("--------------------------------------------------\n")
    a := 10
    b := 11
    fmt.Println("ADD:",a+b)
    fmt.Println("SUB:",a-b)
    fmt.Println("MUL:",a*b)
    fmt.Println("DIVISION:",a/b)
    fmt.Println("MODULUS:",a%b)

    // PERFORMING BITWISE OPERATIONS:
    fmt.Println("BITWISE AND:",a&b)
    fmt.Println("BITWISE |:",a|b)
    fmt.Println("BITWISE XOR:",a^b)
    fmt.Println("BITWISE CLEAR(AND_NOT):",a&^b)
    fmt.Println("MODULUS:",a%b)

    // PERFORMING BIT SHIFTS
    // left shift, If you don't know, then shiting the number by "n", multiplies the  number by 2^n
    var n uint64 = 3 //To perform bit-shift "n" has to an unsigned it , else "Error"
    var z int64 = 10
    r := z << n //that means I am multiplying by 2^r = 8
    fmt.Println("LEFT BIT SHIFT RESULT:", r)

    // right shift, It is the converse of the left shit, that means it, divides the number by 2^n
    r = r >> n
    fmt.Println("RIGHT BIT SHIFT RESULT:",r)

    //WELL I HOPE, I MADE THE IMPLEMENTATION OF THESE OPERATORS SIMPLE. :)
}
