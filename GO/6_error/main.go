package main

//error pkg ==> error.new()
import (
	"errors"
	"fmt"
	"time"
)

//error interface
/*
type error interface{
	Error() string
}*/

// implement struct as error interface
/*type Errormsg struct {
	name string
}
func (e1 Errormsg) Error() string {
	return "error occured"
}*/

//easily can use error.new to define error insted of using error
//var errmsg error = error.new(errror_string)

type person struct {
	id       int
	msg      int
	bill     float64
	distance float64
}

func (p1 *person) sendmessage(p *person, s int) (float64, error) {
	if p.distance == -1 {
		return 0.00, fmt.Errorf("recever person is down")
	}

	if s == 000 {
		return 0.00, errors.New("error is send empty message using 111")
	}
	p.msg = s
	current_bill := (p1.distance + p.distance) * 0.5
	p1.bill += current_bill
	fmt.Printf("msg send successfully and bill : %f \n", p1.bill)
	return current_bill, nil
}

func main() {
	fmt.Println("ch_6 error")

	p1 := person{
		id:       000,
		msg:      -1,
		bill:     0.00,
		distance: 20,
	}

	p2 := person{001, -1, 0.00, 20}
	p3 := person{002, -1, 0.00, 20}

	//p1--->p2, p2---->p3
	bill, err := p1.sendmessage(&p2, 000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("your bill %f", bill)
	}
	p2.sendmessage(&p3, 002)
	p3.sendmessage(&p1, 301)

	fmt.Printf("%d \n %d \n %d", p1.msg, p2.msg, p3.msg)

	time.Sleep(8 * time.Second)
	fmt.Println("exit")
}
