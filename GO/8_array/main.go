package main

import (
	"fmt"
	"time"
)

//slice => dynamicly sized, flexible view of array (subarray)  ----> slice is pointer of array
//var_name := arr_name[] || arr_name[low : high]  (low is inclusion,  high is exclusion )
//can decler explicitly
//change in slice is reflect to array ,  slice is refernce of array (slice == &array)

//make func ---> make slice func of given length and capacity
//make([]type, length, capacity) ---> len() -> return no of element,, cap() -> return max no of element

//veridic--> vector -take any number of argument

func sum(num ...int) int {
	ans := 0
	for i := 0; i < len(num); i++ {
		ans += num[i]
	}
	return ans
}

func main() {
	fmt.Println("ch8 Array")

	//arr declartion
	var arr1 [4]int        //array
	arr2 := []int{1, 2, 3} //slice
	var append_slice []int //slice

	fmt.Println(len(append_slice))
	//slice
	mySlice := arr2[0:2]
	s := sum(1, 2, 3, 3, 2, 3, 3)

	//sperad ---> able to pass slice when intend to pass varidic
	//cannot pass arr1 because [3]
	s += sum(arr2...)

	//append slice --> append(slice []type,element ...type )

	append_slice = append(arr2, arr2...) //do not stor append result indifferent var slice
	arr2 = append(arr2, arr2...)

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(mySlice)
	fmt.Println(s)
	fmt.Println(append_slice)

	time.Sleep(3 * time.Second)
	fmt.Println("exit")
}
