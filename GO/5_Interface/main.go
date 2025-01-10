package main

import (
	"fmt"
)

type message interface {
	getmessage()
	sendmessage(msg string) string
}

// multiple interface
type offic_message interface {
	message
	to_all(msg string) string
}

// implicity(automaticly) impletment interface
type user1msg struct {
	msg, sent string
}

func (u user1msg) getmessage() {
	fmt.Println("hii, user1")
}

func (u user1msg) sendmessage(s string) string {
	return "message " + s + "send successfully"
}

type user2msg struct {
	msg, sent string
}

func (u user2msg) getmessage() {
	fmt.Println("hii, user2")
}

func (u user2msg) sendmessage(s string) string {
	return "message " + s + "send successfully"
}

// return which user , using type assertion
func whichUser(m message) string {
	u1, ok := m.(user1msg)
	if ok {
		u1.getmessage()
		return "user1"
	}

	u2, ok := m.(user2msg)
	if ok {
		u2.getmessage()
		return "user2"
	} else {
		return "unknow_user"
	}
}

// make type asseration using type switch
func whichuserS(m message) string {
	switch m.(type) {
	case user1msg:
		return "user_1"

	case user2msg:
		return "user_2"

	default:
		return "unknow_user"
	}

}

// you can give m as user1msg or user2msg
func test(m message) {
	m.getmessage()
}

func main() {
	//can not declear struct(non-primitive) like var_name var_type
	var u1 user1msg
	var u2 user2msg

	fmt.Println("ch5 Interface")

	//pass user1msg to message(interface)
	test(user1msg{
		msg:  "",
		sent: "hello",
	})

	//type assertion => child_verabile, boolen_var := interface_instance.(child_type)
	usr := whichUser(u1)
	fmt.Println(usr)

	fmt.Println(whichuserS(u2))

}
