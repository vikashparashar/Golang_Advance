package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Lets Take First Number As Arguments From CLI
	a := os.Args[1]    // first number as args
	x , err := strconv.Atoi(a)   // lets convert 1st string into int 
	if err!=nil{
		log.Fatal("failed to convert agrs from string into and int to perform calculation's.")
	}



	// Lets Take Second Number As Arguments From CLI
	b := os.Args[2]    // second number as args
	y,err:=strconv.Atoi(b) // lets convert 2nd string into int 
	if err!=nil{
		log.Fatal("failed to convert agrs from string into and int to perform calculation's.")
	}


// Lets Take Choice of User To Perform Operation On Rest Of Arguments Arguments From CLI
	c := os.Args[3]    // second number as args
	// z,err:=strconv.Atoi(b) // lets convert 2nd string into int
	// if err!=nil{
	// 	log.Fatal("failed to convert agrs from string into and int to perform calculation's.")
	// }

	switch c {
	case "+":
		Addition(x,y)
		
	case "-":
		Subtraction(x,y)
		
	case "*":
		Multiplication(x,y)
		
	case "/":
		Division(x,y)
		
	case "%":
		Remainder(x,y)
		
	}
}


func Addition(a,b int){
	fmt.Println("-----------------------------------------")
	fmt.Printf("\t       %d + %d = %d         \n",a,b,a+b)
	fmt.Println("-----------------------------------------")
}
func Subtraction(a,b int){
	fmt.Println("-----------------------------------------")
	fmt.Printf("\t      %d - %d = %d           \n",a,b,a-b)
	fmt.Println("-----------------------------------------")
}

func Multiplication(a,b int){
	fmt.Println("-----------------------------------------")
	fmt.Printf("\t       %d * %d = %d          \n",a,b,a*b)
	fmt.Println("-----------------------------------------")
}

func Division(a,b int){
	fmt.Println("-----------------------------------------")
	fmt.Printf("\t       %d / %d = %d          \n",a,b,a/b)
	fmt.Println("-----------------------------------------")
}

func Remainder(a,b int){
	fmt.Println("-----------------------------------------")
	fmt.Printf("  Remainder Of  %d / %d = %d          \n",a,b,a%b)
	fmt.Println("-----------------------------------------")
}