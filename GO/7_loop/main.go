package main

import (
	"fmt"
	"time"
)

// skip for prie and end if i == 100
func fizzbuzz() {
	i := 0
	fmt.Println("fireebuzz")
	for {
		i++
		if i == 100000 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)

	}
}
func main() {
	fmt.Println("ch_7 loops")
	cnt := 3
	for i := 0; i < 5; i++ { //can omitt any clause
		fmt.Println(i)
	}

	//no while loop so,
	for cnt < 4 {
		fmt.Println(cnt)
		cnt++
	}

	arr := []int{1, 2, 3}
	//for INDEX, ELEMENT := range Array_name
	for i, num := range arr {
		fmt.Println(i + num)
	}

	fizzbuzz()
	time.Sleep(3 * time.Second)
	fmt.Println("exit")
}
