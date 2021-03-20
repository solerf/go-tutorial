package main

// functions for formatting text
import "fmt"

// from https://pkg.go.dev
import "rsc.io/quote"


func main() {
	// upper case means can be accessed outside its package
    fmt.Println(quote.Go())
}