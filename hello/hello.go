package hello

import (
	"fmt"
	"github.com/kyuelin/go2016/string"
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