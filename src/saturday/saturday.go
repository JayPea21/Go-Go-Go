package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday!")
	case today + 1:
		fmt.Println("Saturday is Tomorrow!")
	case today + 2:
		fmt.Println("It's Saturday In two days.")
  case today + 3:
    fmt.Println("It's Saturday In three days.")
  case today + 4:
    fmt.Println("It's Saturday In four days.")
  case today + 5:
    fmt.Println("It's Saturday In five days.")
  case today + 6:
    fmt.Println("It's Saturday In six days.")
	default:
		fmt.Println("Too far away.")
	}
}
