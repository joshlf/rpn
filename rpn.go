// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		for {
			fmt.Print("> ")
			run(in)
			fmt.Println("Bottom of stack reached")
		}
	}
}

/*
	Operators take integer arguments and return operators (they are curried).
	For example, "add" takes its first argument, and returns an operator 
	which will take the second argument, add them, and return them.

	This introduces a caveat - numbers are themselves operators. Specifically,
	numbers are closures, and their integer values are local variables captured
	inside their closure environments. When called, these functions call the
	"run" function, which continues the stack. Once this function returns an
	operator, the number applies the operator to itself (that is, to its integer
	value), and returns the resulting operator.

	Opgens are functions which take no argument and return an operator. They are
	generated as closures. Their purpose is to ferry around operators until those
	operators are needed, and which point the opgen itself will be called, and it
	will return the operator.

	The "run" function comprises the main loop of the program. Each iteration
	corresponds to another value or operator on the stack (though, since values
	are operators, "run" can't tell the difference). During each iteration of
	"run", input is taken from the user by calling an input function (more on
	these in the next paragraph). The value that the input function returns
	is an opgen. Once this value has been obtained by calling the input function,
	it is itself run in order to obtain an operator, which is returned to lower
	layers of the stack. If you think about what happens to items on the stack
	when an arithmetic operator is pushed, it will make sense why this value should
	be returned to previous layers.

	Input functions are the way that different layers of the stack communicate.
	A certain layer of the stack will recursively call "run", passing an input
	function. "run" will now call that input function. The input function can do
	a number of things. If it is the standard input function, it will simply wait
	for user input, and return it (be it number or operator). However, certain
	circumstances require higher layers in the stack to have their value set by
	lower layers in the stack. For example, the duplicate operator pops the value
	off the top of the stack, and then pushes two copies of it. Only one value can
	be stored per layer of the stack, so one recursive call to "run" must be made
	to store the second pushed value. In order to tell "run" to store this value
	instead of asking for user input, we pass it a specialized input function which,
	when run, will simply return the value we want to push. What's nice about this
	approach is that "run" doesn't have to know anything about the duplicate operator,
	since it can't tell the difference between values taken from user input and values
	which were intentionally passed up the stack by lower layers.

*/
type operator func(int) operator
type opgen func() operator
type input func() opgen

func run(in input) operator {

	var op opgen
	op = in()

	return op()
}

func in() opgen {
	var s string
	var n int
	for {
		fmt.Scan(&s)
		_, err := fmt.Sscanf(s, "%d", &n)

		if err != nil {
			switch s {
			case "+":
				return add
			case "-":
				return subtract
			case "*":
				return multiply
			case "/":
				return divide
			case "|":
				return or
			case "&":
				return and
			case "c":
				return negate
			case "~":
				return not
			case "dup":
				return dup
			case "print":
				return print
			case "pop":
				return pop
			case "swap":
				return swap
			case "zero":
				return zero
			case "quit":
				os.Exit(0)
			}
			fmt.Printf("Unrecognized command: %s\n", s)
			fmt.Print("> ")
		} else {
			break
		}
	}
	return number(n)
}

func in_gen(n int) func() opgen {
	return func() opgen {
		return number(n)
	}
}

func number(n int) opgen {
	return func() operator {
		op := run(in)
		return op(n)
	}
}

func add() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(in_gen(i + j))
		}
	}
}

func subtract() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(in_gen(j - i))
		}
	}
}

func multiply() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(in_gen(i * j))
		}
	}
}

func divide() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(in_gen(j / i))
		}
	}
}

func or() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(in_gen(i | j))
		}
	}
}

func and() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(in_gen(i & j))
		}
	}
}

func negate() operator {
	return func(i int) operator {
		return run(in_gen(-i))
	}
}

func not() operator {
	return func(i int) operator {
		return run(in_gen(^i))
	}
}

func print() operator {
	return func(i int) operator {
		fmt.Println(i)
		fmt.Print("> ")
		return run(in_gen(i))
	}
}

func dup() operator {
	return func(i int) operator {

		input := func() opgen {

			return func() operator {
				op := run(in_gen(i))
				return op(i)
			}
		}

		return run(input)
	}
}

func swap() operator {
	return func(i int) operator {
		return func(j int) operator {

			input := func() opgen {
				return func() operator {
					op := run(in_gen(j))
					return op(i)
				}
			}

			return run(input)
		}
	}
}

func pop() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(in_gen(j))
		}
	}
}

func zero() operator {
	return zero_helper
}

func zero_helper(i int) operator {
	return zero_helper
}
