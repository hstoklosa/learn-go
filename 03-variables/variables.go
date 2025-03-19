package main

import "fmt"

/*
 * The var statement declares a list of variables;
 * the type is last, unless it is initialised.
 */
var c, python, java bool
var p, q = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var c2, python2, java2 = true, false, "no!"
	fmt.Println(p, q, c2, python2, java2)

	/*
 	* := short assignment statement can be used in place of a var declaration with implicit type.
 	* Outside a function, every statement begins with a keyword (:= construct is not available).
 	*/
	var r, s = 1, 2
	q := 3
	c3, python3, java3 := true, false, "no!"

	fmt.Println(r, s, q, c3, python3, java3)
}
