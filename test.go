package main

import "fmt"

// Exported - accessible from other files in the 'main' package
func PrintHello() {
    fmt.Println("Hello from file1!")
}

// not exported - only usable within file1.go
func helperFunction() { 
    fmt.Println("This is a helper function.")
}