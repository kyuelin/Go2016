package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//hello.SetupHtp()
	//hello.TestGoRoutine()
	//hello.TestChannel()


	stringChan := make(chan string)
	for i:=0; i<5; i++ {
		go makeDough(stringChan)
		go addSource(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond*5000)
	}

	fmt.Println("end of main()")
}

var pNum = 0
var pName = ""
func makeDough(stringChan chan string) {
	pNum++
	pName = "Pizza#" + strconv.Itoa(pNum)

	fmt.Println("Make dough and Send for Sauce ", pName)

	stringChan <- pName
	time.Sleep(time.Millisecond*10)
}

func addSource(stringChan chan string) {
	pizza := <- stringChan

	fmt.Println("Add Sauce and Send ", pizza, " for toppings")

	stringChan <- pizza

	time.Sleep(time.Millisecond*10)
}

func addToppings(stringChan chan string) {
	pizza := <- stringChan

	fmt.Println("Add Toppings ", pizza)

	time.Sleep(time.Millisecond*10)
}
