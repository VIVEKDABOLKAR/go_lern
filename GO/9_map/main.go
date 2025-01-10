package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("ch 9 map(key, element)")

	//map[key]element   --> key must be commpareable type
	//nil is zero value of map
	//map is mutable
	//map is pass by refernce

	//method0-1()
	var age = map[string]int{
		"john": 21,
		"reco": 22,
		"riom": 34,
	}

	//method-2(create empty map)
	name_ph := make(map[string]int)
	name_ph["john"] = 94321
	name_ph["reco"] = 98756
	name_ph["riom"] = 94213

	//element, ok = map[key]
	element, ok := age["rii"]

	//if element is not exist then ok is false, element value is zero value of type
	if ok {
		fmt.Println(element)
	}

	//nested map --> map[key]map[key]element
	nest := make(map[string]map[string]int)
	//to declear-> I) declar first map
	//			   II) assign var useing first and second map
	nest["root"] = map[string]int{}
	nest["root"]["opt"] = 5
	nest["root"]["bin"] = 3

	nest["c:"] = map[string]int{}
	nest["c:"]["download"] = 3

	//we can use struct to alternetive to nested map
	//nest = make(map[struct]int)

	fmt.Println(age)
	fmt.Println(len(name_ph))

	fmt.Println(nest)
	fmt.Println("exit")
	time.Sleep(3 * time.Second)
}
