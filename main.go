package main

import (
	"fmt"
	"learnpackage/simpleInterest" // module name/ package name
	"log"
)

var a, b = 2,-3
func init() {  
    println("Main package initialized")
    if a < 0 {
        log.Fatal("Principal is less than zero")
    }
    if b < 0 {
        log.Fatal("Rate of interest is less than zero")
    }
}

func main() {
	var _ = simpleInterest.Kara(a,b)
	fmt.Println("sasha simple")
	// fmt.Println(simpleInterest.Kara(a,b))
}