package main

import "fmt"

func main() {
	fmt.Println("begin the count")

//the for loop with defer will let the last println go first
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("this is the end of the count... right?")
}
