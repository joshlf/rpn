// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*
	TODO: Get pop to work either as a unary operator
	or have the main loop inspect for its return value
*/

import (
	"fmt"
	"os"
)

const (
	NONE = iota
	DUP
	PRINT
	POP
	SWAP
	ZERO
)

func main() {
	for {
		fmt.Print("> ")
		uOp, bOp := run(false, 0)
		// ind := ZERO
		// 		for ind == ZERO {
		// 			fmt.Print("> ")
		// 			ind, _, _ = run(false, 0)
		// 		}
		fmt.Println("Invalid entry: bottom of stack reached")
	}
}

/*
	push indicates whether to push the passed value onto the stack or
	to wait for user input.

	Returns an operator indicator. This indicator can have values
	DUP, or PRINT

	NONE:
		The operator entered was a normal arithmetic operator

	DUP:
		The operator entered was the duplicate operator

	PRINT:
		The operator entered was the print operator

	POP:
		The operator entered was the pop operator

	SWAP:
		The operator entered was the swap operator

	ZERO:
		The operator entered was the zero operator

	If the indicator is equal to then one of the two returned
	functions will be the function corresponding to the operator entered.

*/
func run(push bool, n int) (unop, binop) {

	var s string
	var uOp unop
	var bOp binop

	// Operator indicator
	// var ind int

	// If the duplicate operator was entered, then n
	// is already the equal to the value which should
	// be pushed onto the stack.
	if !push {
		for {
			fmt.Scan(&s)
			_, err := fmt.Sscanf(s, "%d", &n)

			// If it was not a number (ie, an operator)
			if err != nil {
				switch s {
				case "+":
					return nil, add
				case "-":
					return nil, subtract
				case "*":
					return nil, multiply
				case "/":
					return nil, divide
				case "|":
					return nil, or
				case "&":
					return nil, and
				case "c":
					return negate, nil
				case "~":
					return not, nil
				case "dup":
					return dup, nil
				case "print":
					// fmt.Println("returning print")
					return prnt, nil
				case "pop":
					return nil, pop
				case "swap":
					return nil, swap
				// case "zero":
				// return ZERO, nil, nil
				case "quit":
					os.Exit(0)
				}
				fmt.Printf("Unrecognized command: %s\n", s)
				fmt.Print("> ")
			} else {
				break
			}
		}
	}

	// Once control reaches this part of the function,
	// n is equal to the value on the top of the stack.

	push = false
	newN := n
	
	for {
		// fmt.Printf("Recursive call for n = %d\n", n)
		uOp, bOp = run(push, newN)
		if uOp != nil {
			// var newN int
			n, push, newN = uOp(n)
			// uOp, bOp = run(push, newN)
		} else {
			return bOp(n), nil
		}
		// 
		// TOP:
		// 
		// 	push = false
		// 
		// 	switch ind {
		// 	case NONE:
		// 		if uOp != nil {
		// 			n = uOp(n)
		// 		} else {
		// 			return bOp(n), nil
		// 		}
		// 	case DUP:
		// 		// Simply set push to true, since on next iteration of loop,
		// 		// run(push, n) will be called, and n is already the correct value
		// 		push = true
		// 	case PRINT:
		// 		fmt.Println(n)
		// 		fmt.Print("> ")
		// 	case POP:
		// 		// Effectively letting the previous instance of run perform the call,
		// 		// but way easier than having to pass back sentinal values, etc
		// 		return run(false, 0)
		// 	case SWAP:
		// 		if uOp == nil {
		// 			// bOp will return a function which, when given an argument,
		// 			// will discard the argument and simply return this n
		// 			return SWAP, bOp(n), nil
		// 		} else {
		// 			// uOp will discard its argument and return the argument which
		// 			// was passed in the previous call
		// 			m := uOp(0)
		// 			ind, uOp, bOp = run(true, n)
		// 			n = m
		// 
		// 			// Necessary to avoid double-calling run
		// 			goto TOP
		// 		}
		// 	case ZERO:
		// 		return ZERO, nil, nil
		// 	}
	}

	// Control should never reach this
	return nil, nil
}
