package hello

import (
	"fmt"
	"github.com/kyuelin/go2016/string"
	"net/http"
	"time"
	"strconv"
)

func main() {
	var age int = 40
	var favNum float64 = 3.14159

	fmt.Println(age, favNum)

	var numOne = 1.000
	var num99 = .99

	fmt.Println(numOne - num99)

	fmt.Println(string.Reverse("hello, dude!"))
	numlist1 := []int {2,35,6,3,5}
	numlist := make ([]int, 5, 10)
	copy(numlist, numlist1)

	numlist3 := append(numlist, 0, -1, 9)
	fmt.Println(numlist3)


}

func SetupHtp() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func handler1(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello dude!\n")
}

func Count(id int) {
	for i:=0; i<10; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond*1000)
	}
}

func TestGoRoutine() {
	for i:=0; i<10; i++ {
		go Count(i)
	}
	//time.Sleep(time.Millisecond*10*1000)
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

	stringChan <- pizza

	time.Sleep(time.Millisecond*10)
}

func TestChannel() {
	stringChan := make(chan string)
	for i:=0; i<3; i++ {
		go makeDough(stringChan)
		go addSource(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond*5000)
	}
}

