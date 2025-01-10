package main

import "fmt"

type person struct {
	name string
	age  int
}

// f func(func(int, int) int, int) int  ????
func math(operation func(int, int) int, x, y int) int {
	return operation(x, y)

}
func add(i1, i2 int) int {
	i1 = i1 + i2
	return i1
}
func minus(i1, i2 int) int {
	return i1 - i2
}

//multiple return value
func getName() (string, int) {
	return "name", 0
}

func name(operaton func(int, int) int) (string, string) {
	return "hello", "world"
}

func main() {
	sendedSofer := 438
	const sendsToAdd = 25
	sendedSofer = add(sendedSofer, sendsToAdd)

	//print line
	fmt.Println("you've sent", sendedSofer, "message")

	p := person{}
	p.name, p.age = getName()

	fmt.Println(p.name, p.age)
	fmt.Println(math(add, 5, 3))
	fmt.Println(name(add))

}
