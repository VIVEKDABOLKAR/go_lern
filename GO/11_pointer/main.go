package main

//pointer := store address of the variable
//ptr = &var --str add of var
//*ptr --point to the value of the var
//var ptr *int -->zero value of pointer is nil  dereferencing
//ptr has no arthmetic operation --> !ptr++ like c (no illegel access to memory location)

import (
	"fmt"
	"time"
)

type circle struct {
	r     int
	x, y  int
	color int
}

func (c *circle) setcolor() int {
	c.color = 222
	return 0
}

func main() {
	fmt.Println("ch11 Pointer")
	variable := 5
	ptr := &variable //store address of the variable it is call pointer

	c := circle{
		r:     1,
		x:     2,
		y:     3,
		color: 268,
	}

	//c is not pointer in below function
	//but method still get access to the pointer method (it casted to the pointer)
	c.setcolor()

	fmt.Println(c.color) //
	fmt.Println(*ptr)    //*gain access to the value of the variable
	fmt.Println("exit")
	time.Sleep(time.Second * 3)
}
