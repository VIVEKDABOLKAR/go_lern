package main

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// it is aggregate(high degree) function := take func as parameter and func call it self recursevly
func op(a, b, c int, operation func(int, int) int) int {
	return operation(operation(a, b), c)
}

// currying function := take func as input && give func as output as well
func selfmath(operation func(int, int) int) func(int) int { //selfmath can do operartioj of two parameter in single parameter
	return func(x int) int {
		return operation(x, x)
	}
}

// defer := execult line before return line
func example(s string) {
	defer fmt.Println(s)
	if s == "" {
		fmt.Println("string is empty")
	}
}

// A closure is a function value that references variables from outside its body.
func add_closure() func(int) int {
	sum := 0
	return func(a int) int { //Anonymous function := function which has no name
		sum += a
		return sum
	}

}

func main() {
	fmt.Println("ch_10 advance fnction")

	ans := op(2, 3, 4, add) // higher oreder func

	single_var := selfmath(mul) //currying func

	example("") //defer keyword

	closur_sum := add_closure()
	closur_sum(3)
	closur_sum(5)

	fmt.Println(ans)
	fmt.Println(single_var(5))
	fmt.Println(closur_sum(0))

	fmt.Println("exit")
	time.Sleep(3 * time.Second)
}
