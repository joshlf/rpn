// Copyright 2013 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type operator func(int) operator

func main() {
	for {
		fmt.Print("> ")
		op := input()
		if !sameFunc(op, zero) {
			fmt.Println("Error: stack bottomed out")
		}
	}
}

func input() operator {
	var s string
	var n int
	for {
		_, err := fmt.Scan(&s)

		if err == io.EOF {
			fmt.Println()
			os.Exit(0)
		}

		_, err = fmt.Sscanf(s, "%d", &n)

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

func number(i int) operator {
	op := input()
	return op(i)
}

func add(top int) operator {
	return func(bottom int) operator {
		return number(bottom + top)
	}
}

func subtract(top int) operator {
	return func(bottom int) operator {
		return number(bottom - top)
	}
}

func multiply(top int) operator {
	return func(bottom int) operator {
		return number(bottom * top)
	}
}

func divide(top int) operator {
	return func(bottom int) operator {
		return number(bottom / top)
	}
}

func or(top int) operator {
	return func(bottom int) operator {
		return number(bottom | top)
	}
}

func and(top int) operator {
	return func(bottom int) operator {
		return number(bottom & top)
	}
}

func negate(i int) operator {
	return number(-i)
}

func not(i int) operator {
	return number(^i)
}

func print(i int) operator {
	fmt.Println(i)
	fmt.Print("> ")
	return number(i)
}

func dup(i int) operator {
	op := number(i)
	return op(i)
}

func swap(top int) operator {
	return func(bottom int) operator {
		op := number(bottom)
		return op(top)
	}
}

func pop(top int) operator {
	return func(bottom int) operator {
		return number(bottom)
	}
}

func zero(i int) operator {
	return zero
}

func sameFunc(i, j operator) bool {
	// Credit to http://stackoverflow.com/a/18067479/836390
	return reflect.ValueOf(i).Pointer() == reflect.ValueOf(j).Pointer()
}
