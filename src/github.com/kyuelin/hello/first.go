package main

import "fmt"

//comments

func main() {
	fmt.Println("Hello World")
	var age int = 40
	var favNum float64=1.61234
	randNum := "hello" 
	fmt.Println(age, favNum, randNum)

	fmt.Println("6+4=", 6+4)
	const pi float64 = 31.4159
	var myName string = "Kenneth Lincoln"
	fmt.Println(len(myName))

	j := 0
	for j <10 {
		//fmt.Println(j)
		j++
	}
	
	yourAge:=18
	switch yourAge {
		case 14: fmt.Println("can drive")
		default: fmt.Println("can't drive")
	}

	var favNums[4] float64
	fmt.Println(favNums)

	for _, value :=range favNums {
		fmt.Println(value)
	}

	numSlice := []int {5,4,4,5,6}
	var numSlice1 []int = numSlice[1:3]
	fmt.Println(numSlice)
	fmt.Println(numSlice1)
	numSlice1[1]=99
	fmt.Println(numSlice)
	fmt.Println(numSlice1)
	
	numSlice1 = append(numSlice1,0,-1,1,2,34,5)
	fmt.Println(numSlice1)
	fmt.Println("----")
	numSlice1[0]=999
	fmt.Println(numSlice1)
	fmt.Println(numSlice)
}