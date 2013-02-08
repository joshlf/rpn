// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	for {
		for {
			fmt.Print("> ")
			run(input)
			fmt.Println("Bottom of stack reached")
		}
	}
}

func run(input func() opgen) operator {

	var op opgen
	op = input()

	return op()
}
