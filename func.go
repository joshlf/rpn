// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

/*
	Binary operators are curried.
*/

type unop func(int) (int, bool, int)
type binop func(int) unop

func add(j int) unop {
	return func(i int) (int, bool, int) {
		return i + j, false, 0
	}
}

func subtract(j int) unop {
	return func(i int) (int, bool, int) {
		return i - j, false, 0
	}
}

func multiply(j int) unop {
	return func(i int) (int, bool, int) {
		return i * j, false, 0
	}
}

func divide(j int) unop {
	return func(i int) (int, bool, int) {
		return i / j, false, 0
	}
}

func or(j int) unop {
	return func(i int) (int, bool, int) {
		return i | j, false, 0
	}
}

func and(j int) unop {
	return func(i int) (int, bool, int) {
		return i & j, false, 0
	}
}

func negate(i int) (int, bool, int) {
	return -1 * i, false, 0
}




func not(i int) (int, bool, int) {
	return ^i, false, 0
}

/*
	Non-arithmetic operators
*/

func swap(i int) unop {
	return func(j int) (int, bool, int) {
		return i, true, j
	}
}

func dup(i int) (int, bool, int) {
	return i, true, i
}

func popBottom(i int) (int, bool, int) {
	return i, false, 0
}

func pop(i int) unop {
	return popBottom
	// return func(j int) (int, bool, int) {
	// 	return j, false, 0
	// }
}

func prnt(i int) (int, bool, int) {
	fmt.Println(i)
	fmt.Print("> ")
	return i, false, 0
}

