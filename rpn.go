// Copyright 2013 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type operator func(int) operator

var operators map[string]operator

func init() {
	// Do this in init() to avoid
	// initialization loop
	operators = map[string]operator{
		"+":     add,
		"-":     subtract,
		"*":     multiply,
		"/":     divide,
		"|":     or,
		"&":     and,
		"c":     negate,
		"~":     not,
		"dup":   dup,
		"print": print,
		"pop":   pop,
		"swap":  swap,
		"zero":  zero,
	}
}

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

		n, err = strconv.Atoi(s)

		if err != nil {
			if s == "quit" {
				os.Exit(0)
			}
			op, ok := operators[s]
			if ok {
				return op
			} else {
				fmt.Printf("Unrecognized command: %s\n", s)
				fmt.Print("> ")
			}
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
