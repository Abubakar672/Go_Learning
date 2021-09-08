/*

GO lEARNING
Authors: Ali Wisam, AbuBakar Mujahid

*/


/*whenever we want to start the code of go we have to import the main package like package main    
import fmt used for formatting the code in the golang
and your codes end on the main function

command: go fmt

if in terminal you will go to the folder and enter the command of the go fmt so it will format all the code according to the go lang 

& 
command: go fmt ./...

if you will enter the go fmt ./... this command in a repo so it will format all the code files in that repo into the golang format
*/

//Simple program to print Hello Wordl in Go

// package main

// import "fmt"

// func main() {
//   fmt.Println("Hello World")
// }


/*
command: go run <filename>
As for the case if you want to run your .go file through terminal and want to see output so simply in terminal go to the specific folder and enter the command go run <filename>

command: go build

if you enter the command go build in that specific folder it will build the exe file of the program

command: ./..go build

*/


/*

Control Flows in Go lang

There are three type of control flow in Go 
Control flow is the flow control in which the function call are been executed means it is the flow that how the computer will read the code 

1)Sequential 
2)Loop Iterative
3)Conditional 

*/


//First Program in Go to print data 

// package main 

// import "fmt" 

// func main(){

//   fmt.Println("Hello ABubkar")
// }


//Go Programming codes here 


/*
**************************************************************
*                         CODE 1                             *
**************************************************************
*/


// package main

// import "fmt"

// func main(){
//   fmt.Println("a")
//   foo()
//   even()
// }


/* := short decleration operator as it allows to write code and declare available 
*/

/*   

x:=42    here we are declaring the variabel x with a value of 42
fmt.Println(x)
x=2   now here we are asssinging x with a value of 2 as x is already declared up there 
fmt.Println(x)

*/

// func foo(){
//   fmt.Println("HI ")
// }

// func even(){
// for i:=0 ;i<100;i++{
//   if i%2==0{
//   fmt.Println(i)
// }}

// }



/*
**************************************************************
*                         CODE 2                             *
**************************************************************
*/

// package main

// import "fmt"

// func main(){
//   x:= 45
//   fmt.Println(x)
//   x=120
//   fmt.Println(x)
// }

/*
**************************************************************
*                         CODE 3                             *
**************************************************************

difference between the var keyword and the short hand operator =: is this that the var have global scope and the short hand operator := can only be used inside the function and in order to print out the 
*/

//creating our own datatype 
/*
package main
import "fmt"
var a string
type boy int 
var b boy
type student string
var male student

func main(){
a= "asd"
b=12
male= "AbuBakar"

fmt.Println(a);
fmt.Println(b);
fmt.Println(male);
 }
*/


/*
**************************************************************
*                         CODE 4                             *
**************************************************************

conversion of datatypes in golang 
*/

// package main 
// import "fmt"

// var a int 
// type boy int
// var b boy

// func main (){
// a=2
// b=256

// //As now if we do as a=b so it will give error as the a has an int data type 
// //while b is of boy datatype so here comes the conversion shown as followed 

// fmt.Println("Before conversion")
// fmt.Println(a)
// fmt.Println(b)

// a=int(b)

// fmt.Println("After conversion")
// fmt.Println(a)
// fmt.Println(b)

// }

/*
**************************************************************
*                         CODE 5                             *
**************************************************************

NINJA LEVEL 1
*/

//Exercise 1

// package main
// import "fmt"

// func main(){
 /*1*/
// x := 123
// y:= "James Bond"
// z:= "True"

 /* 2 */
// //Single Printing
// fmt.Println(x,y,z)

// //Multiple Printing
// fmt.Println(x)
// fmt.Println(y)
// fmt.Println(z)
// }

//Exercise 2

// package main 
// import "fmt"

// var x int;

// var y string;

// vaar z bool;

// func main(){

// fmt.Println(x)
// fmt.Println(y)
// fmt.Println(z)
// }

//Exercise 3

// package main
// import "fmt"

// var x int = 42
// var y string = "James Bond"
// var z bool = true

// func main(){

// s := fmt.Sprintf("%v\t","%v\t","%v\t", x,y,z)
// fmt.Println(s)
// }






