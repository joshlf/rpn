// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

type opgen func() operator
type operator func(int) operator

func input() opgen {
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

func input_gen(n int) func() opgen {
	return func() opgen {
		return number(n)
	}
}

func number(n int) opgen {
	return func() operator {
		op := run(input)
		return op(n)
	}
}

func add() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(input_gen(i + j))
		}
	}
}

func subtract() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(input_gen(j - i))
		}
	}
}

func multiply() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(input_gen(i * j))
		}
	}
}

func divide() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(input_gen(i / j))
		}
	}
}

func or() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(input_gen(i | j))
		}
	}
}

func and() operator {
	return func(i int) operator {
		return func(j int) operator {
			return run(input_gen(i & j))
		}
	}
}

func negate() operator {
	return func(i int) operator {
		return run(input_gen(-i))
	}
}

func not() operator {
	return func(i int) operator {
		return run(input_gen(^i))
	}
}

func print() operator {
	return func(i int) operator {
		fmt.Println(i)
		fmt.Print("> ")
		return run(input_gen(i))
	}
}

func dup() operator {
	return func(i int) operator {

		input := func() opgen {

			return func() operator {
				op := run(input_gen(i))
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
					op := run(input_gen(j))
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
			return run(input_gen(j))
		}
	}
}

func zero() operator {
	return zero_helper
}

func zero_helper(i int) operator {
	return zero_helper
}
