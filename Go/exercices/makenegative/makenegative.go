package main

func MakeNegative(x int) int {

	if x <= 0 {
		return x
	}
	return -x
}

// exercício do codewars
// exemplos:
// MakeNegative(1)    // return -1
// MakeNegative(-5)   // return -5
// MakeNegative(0)    // return 0
