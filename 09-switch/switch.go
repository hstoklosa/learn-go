package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
 * Go's switch only runs the selected case, hence 'break' statement is not needed.
 * Also the cases need not be constants, and values involved need not be integers.
 */

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	/* Switch cases evaluate cases from top to bottom, stopping when a case succeeds. */
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In 2 days.")
	default:
		fmt.Println("Too far away.")
	}

	/* Switch without condition */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}