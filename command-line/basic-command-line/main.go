// Date : 19/09/2022 00:55 IST
// Basics-Command-Line_Args & Flangs In Golang

package main

import (
	"fmt"
	"os"
)

func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}