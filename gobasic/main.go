package main

import (
	"fmt"
	"gobasic/customer"
	"gobasic/user"
	"unicode/utf8"
)

// global variable cannot use short declare
var g int = 10

func main() {

	// declare variable in golang have to used
	// if not used it cannot compile
	var x int = 3
	var y = 4

	// short declare variable z
	// golang know the
	z := 10

	//if declare but not used
	test := 20
	_ = test

	// note that golang has set the default value at the beginning
	// call it zero value
	var booltype bool     //false
	var inttype int       // 0
	var stringtype string // ""

	fmt.Printf("Declare variable\n")
	fmt.Printf("*****************\n")
	//print the solution
	fmt.Printf("x=%d\ny=%d\n", x, y)
	fmt.Printf("z=%d\n", z)
	fmt.Printf("booltype = %v\n", booltype)
	fmt.Printf("inttype = %v\n", inttype)
	fmt.Printf("stringtype = %v\n", stringtype)

	fmt.Printf("if else condition\n")
	fmt.Printf("*****************\n")
	// if else condition
	point := 50
	if point >= 50 && point <= 100 {
		fmt.Printf("point is more than 50 <= 100\n")
	} else if point >= 20 {
		fmt.Printf("point is more than 20\n")
	}

	fmt.Printf("*****************\n")
	fmt.Printf("Array\n")
	fmt.Printf("*****************\n")
	// array datatype
	var xarray = [3]int{}
	// or delcare like this
	yarray := [3]int{1, 2, 3}
	fmt.Printf("xarray = %v\n", xarray) // return [0,0,0]
	fmt.Printf("yarray = %v\n", yarray) // return [1,2,3]

	// assign array
	xarray[0] = 10
	fmt.Printf("xarray = %v\n", xarray) // return [10,0,0]

	// 2D array
	x2darray := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{1, 2, 3},
	}

	x2darray[0][1] = 10
	fmt.Printf("2d array = %v\n", x2darray)

	// not default size of array in golang
	zarray := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("z array = %v\n", zarray)

	//list (slice) in golang
	xlist := []int{1, 2, 3}
	//append 4 to xlist
	xlist = append(xlist, 4)
	xlist = append(xlist, 4)
	fmt.Printf("xlist = %v\n", xlist)
	//length of list or array
	fmt.Printf("length of xlist is %v\n", len(xlist))

	// length of string
	name := "กจ"
	fmt.Printf("length of name = %d\n", utf8.RuneCountInString(name))

	// slide slice and array
	//			0	1	2	3	4	5	6	7	8
	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	b := a[1:]
	_ = a
	c := a[1:6]
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)

	fmt.Printf("*****************\n")
	fmt.Printf("Mapping\n")
	fmt.Printf("*****************\n")

	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["en"] = "United State"
	println(countries["th"])

	country, ok := countries["jp"]
	if !ok {
		println("no value")
	} else {
		println(country)
	}

	fmt.Printf("*****************\n")
	fmt.Printf("for loop\n")
	fmt.Printf("*****************\n")

	value := []int{10, 20, 30, 40, 50, 60}
	for i := 0; i < len(value); i++ {
		fmt.Println(value[i])
	}

	// while loop
	j := 0
	for j < len(value) {
		println(value[j])
		j++
	}

	// foreach
	for index, v := range value { //if not use index we can use _
		println(index, v)
	}

	fmt.Printf("*****************\n")
	fmt.Printf("function\n")
	fmt.Printf("*****************\n")

	fmt.Printf("sum function %v\n", sum(2, 4))

	result, hello := summultiple(3, 4)
	fmt.Printf("summultiple %v %v\n", result, hello)

	// spcieal function way
	xfunc := func(a, b int) int {
		return a + b
	}

	sumresult := xfunc(10, 20)
	fmt.Printf("sumresult = %v\n", sumresult)

	cal(sum) // function that has function as a arguement

	cal(func(a, b int) int {
		return a - b
	})

	fmt.Printf("sum slice = %v\n", sumslice(1, 2, 3, 4))

	// package
	fmt.Printf("package customer.Sum = %v\n", customer.Sum(2, 3))


	// struct
	xstruct := user.Person{}
	xstruct.SetName("Bond")
	fmt.Printf("struct user.person %v\n",xstruct.GetName())

}

// we can write func sum(a,b int) instate of a int, b int
func sum(a int, b int) int {
	return a + b
}

func summultiple(a, b int) (int, string) {
	return a + b, "hello"
}

func sumslice(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

// function that have func as a arguement
func cal(f func(int, int) int) {
	result := f(50, 20)
	fmt.Printf("cal function = %v\n", result)
}
