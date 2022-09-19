// Date : 19/09/2022 00:55 IST
// Simple Testing And Table Testing In Golang ,  BenchMarking In Golang

// you must have to test the all functions that you created in your program
// refactor you code and try to do run multiple test for your function
package main

import (
	"fmt"

	"github.com/gic-vikash/golang-advance/testing/add"
)

func main(){
	Result := add.Addition(5,123)
	fmt.Println(Result)
}

