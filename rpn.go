package main

import (
	"fmt"
	"os"
)

const (
	NONE  = 0
	DUP   = 1
	PRINT = 2
	POP   = 3
)

func main() {
	run(false, 0)
	fmt.Println("Invalid entry: bottom of stack reached")
	os.Exit(1)
}

/*
	dup indicates whether or not there was a duplicate instruction
	if dup is true, n indicates the value to be duplicated

	Returns an operator indicator. This indicator can have values
	NONE, DUP, or PRINT

	NONE:
		The operator entered was a normal arithmetic operator

	DUP:
		The operator entered was the duplicate operator

	PRINT:
		The operator entered was the print operator

	POP:
		The operator entered was the pop operator

	If the indicator is equal to NONE, then one of the two returned
	functions will be the function corresponding to the operator entered.

*/
func run(dup bool, n int) (int, unop, binop) {

	var s string
	var uo unop
	var bo binop

	// Operator indicator
	var ind int

	// If the duplicate operator was entered, then n
	// is already the equal to the value which should
	// be pushed onto the stack.
	if !dup {
		for {
			fmt.Scan(&s)
			_, err := fmt.Sscanf(s, "%d", &n)

			// If it was not a number (ie, an operator)
			if err != nil {
				switch s {
				case "+":
					return NONE, nil, add
				case "-":
					return NONE, nil, subtract
				case "*":
					return NONE, nil, multiply
				case "/":
					return NONE, nil, divide
				case "|":
					return NONE, nil, or
				case "&":
					return NONE, nil, and
				case "c":
					return NONE, negate, nil
				case "~":
					return NONE, not, nil
				case "dup":
					return DUP, nil, nil
				case "print":
					return PRINT, nil, nil
				case "pop":
					return POP, nil, nil
				case "quit":
					os.Exit(0)
				}
				fmt.Println("Unrecognized command")
			} else {
				break
			}
		}
	}

	// Once control reaches this part of the function,
	// n is equal to the value on the top of the stack.

	dup = false

	for {
		ind, uo, bo = run(dup, n)
		dup = false

		switch ind {
		case NONE:
			if uo != nil {
				n = uo(n)
			} else {
				return NONE, bo(n), nil
			}
		case DUP:
			// Simply set dup to true, since on next iteration of loop,
			// run(dup, n) will be called, and n is already the correct value
			dup = true
		case PRINT:
			fmt.Println(n)
		case POP:
			// Effectively letting the previous instance of run perform the call,
			// but way easier than having to pass back sentinal values, etc
			return run(false, 0)
		}
	}
	return NONE, nil, nil
}
