package hello

import (
	"fmt"
	"math"
	"strings"
	"sort"
	"os"
	"io/ioutil"
	"log"
	"strconv"
)

func math2() {


	randInt := 6
	randFloat := 6.7
	randString := "6"
	randString2 := "6.2"
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)
	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)

	file, error := os.Create("s.txt")

	if error != nil {
		log.Fatal(error)
	}
	file.WriteString("dudududu")
	file.Close()

	stream, error := ioutil.ReadFile("s.txt")
	if error != nil {
		log.Fatal(error)
	}
	readstr := string(stream)
	fmt.Println(readstr)

	n3 := 3
	//closure
	doubleNum := func() int {
		n3 *= 2
		return n3
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
	fmt.Println(n3)
	changeValue(&n3)
	fmt.Println(n3)

	defer printOne()
	printTwo()


	fmt.Println(divNum(2,0))


	yPtr := new(int)

	changeValue(yPtr)
	fmt.Println("y=",*yPtr)

	rect1 := Rectangle{20,50}
	fmt.Println(getArea(rect1))
	cir := Circle{3}
	fmt.Println(getArea(cir))

	fmt.Println("=======================")

	ss := "1,3,4,5,5,6,"
	fmt.Println(strings.Split(ss, ","))

	listl := []string{"2", "c","a", "1"}
	sort.Strings(listl)
	fmt.Println(strings.Join(listl, "|"))

}

type Shape interface {
	area() float64
}

type Shape2 interface {
	area() float64
}


type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) area() float64 {
	return rect.height*rect.width
}

func (cir Circle ) area() float64 {
	return math.Pi * math.Pow(cir.radius, 2)
}

func getArea(s Shape) float64{
	return s.area()
}

func changeValue(x *int) {
	*x =2

}

func divNum(n int, dn int) int {
	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC")

	res := n/dn
	return res
}

func printOne() {fmt.Println("from printOne()", 1)}
func printTwo() {fmt.Println("from printTwo()", 2)}


func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func subtractThem(args ...int) int {
	fv := 0
	for _, val := range args {
		fv -= val
	}
	return fv
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _,  val := range numbers {
		sum+=val
	}
	return sum
}

func nextValue(n int) (int,int) {
	return n+1, n+2
}

