package main

import "fmt"

/*
 * := short assignment statement can be used in place of a var declaration with implicit type.
 * Outside a function, every statement begins with a keyword (:= construct is not available).
 */

func main9() {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}