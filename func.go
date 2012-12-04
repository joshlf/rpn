package main

/*
	Binary operators are curried.
*/

type unop func(int) int
type binop func(int) unop

func add(j int) unop {
	return func(i int) int {
		return i + j
	}
}

func subtract(j int) unop {
	return func(i int) int {
		return i - j
	}
}

func multiply(j int) unop {
	return func(i int) int {
		return i * j
	}
}

func divide(j int) unop {
	return func(i int) int {
		return i / j
	}
}

func or(j int) unop {
	return func(i int) int {
		return i | j
	}
}

func and(j int) unop {
	return func(i int) int {
		return i & j
	}
}

func negate(i int) int {
	return -1 * i
}

func not(i int) int {
	return ^i
}

func swap(i int) unop {
	return func(j int) int {
		return i
	}
}