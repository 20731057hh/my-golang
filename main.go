package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("=====================")
	fmt.Println("Hello World")
	fmt.Println("1 + 1 =",1+1)
	var x int = 1
	fmt.Printf("%v\n", x+1)
	fmt.Println("=====================")
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, t, 1)
	fmt.Printf("%T %v %t\n", f, f, 1)
	fmt.Println("=====================")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("=====================13")
	y := 1
	fmt.Printf("%T %v\n",y,y)
	yy := float64(y)
	fmt.Printf("%T %v\n",yy,yy)
	var s string = "14"
	fmt.Printf("%T %v\n",s,s)
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n",i,i)
	h := "Hello World"
	fmt.Println(h)
	fmt.Println(h[0])
	fmt.Println(string(h[0]))
	fmt.Println("=====================14")
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)
	var b[]int = []int{100,200}
	b = append(b,300)
	fmt.Println(b)
	n := []int{1,2,3,4,5,6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:5])
	fmt.Println(n[:2])
	fmt.Println(n[:])
	//pythonのように−1はできない
	//fmt.Println(n[-1])
	var board = [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
	}
	fmt.Println(board)
	fmt.Println("=====================16")
	n = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	c := make([]int, 5)
	//c := make([]int, 0, 5)
	for i := 0; i < 5; i++{
		c = append(c,i)
		fmt.Println(c)
	}
	fmt.Println(c)
	fmt.Println("=====================17")
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["new"] = 500
	fmt.Println(m)
	fmt.Println(m["nothing"])
	v, ok := m["nothing"]
	fmt.Println(v, ok)
	m2 := make(map[string]int)
	m2["pc"] = 50000
	fmt.Println(m2)
	var s2 []int
	if s2 == nil{
		fmt.Println("Nil")
	}
	fmt.Println("=====================18")
	b2 := []byte{72, 73}
	fmt.Println(b2)
	fmt.Println(string(b2))
	c2 := []byte("HI")
	fmt.Println(c2)
	fmt.Println(string(c2))
	fmt.Println("=====================19")
	r1, r2 := add(1,2)
	fmt.Println(r1, r2)
	r3 := cal(10,2)
	fmt.Println(r3)
	f2 := func(x int){
		fmt.Println("inner func",x)
	}
	fmt.Println("func 0")
	f2(1)
	func(x int){
		fmt.Println("inner func",x)
	}(2)
	fmt.Println("=====================20")
	counter := incremantGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	cc := circleArea(3.14)
	fmt.Println(cc(2))
	fmt.Println(cc(3))
	cc1 := circleArea(3.15)
	fmt.Println(cc1(2))
	fmt.Println(cc1(3))
	fmt.Println("=====================")

}

func circleArea(pi float64) func(radius float64) float64{
	return func(radius float64) float64{
		return pi * radius * radius
	}
}

func incremantGenerator() (func() int){
	x := 0
	return func() int{
		x++
		return x
	}

}

func add(x, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x -y
}

func cal(price, item int) (result int){
	result = price * item
	return result
}

func convert(price int) float64{
	return float64(price)
}