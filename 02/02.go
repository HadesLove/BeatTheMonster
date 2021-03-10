package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//var i int = 10
	var i = 10
	/*var (
		j int = 0
		k int = 1
	)*/
	var (
		j = 0
		k = 1
	)
	fmt.Println(i)
	fmt.Println(j, k)

	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is", f32, ", f64 is", f64)

	var bf bool = false
	var bt bool = true
	fmt.Println("bf is", bf, ", bt is", bt)

	var s1 string = "Hello"
	var s2 string = "Golang"
	fmt.Println("s1 is", s1, ", s2 is", s2)
	fmt.Println("s1+s2", s1+s2)

	s1 += s2
	fmt.Println("new s1 is", s1)

	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)

	/*i := 10
	bf := false
	s1 := "Hello"*/

	pi := &i
	fmt.Println(*pi)

	i = 20
	fmt.Println("i 的新值是", i)

	const name = "Golang"

	/*const(
		one = 1
		two = 2
		three = 3
		four = 4
	)*/

	const (
		one = iota + 1
		two
		three
		four
	)

	fmt.Println(one, two, three, four)

	i2s := strconv.Itoa(i)
	s2i, _ := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i)

	fmt.Println(strings.HasPrefix(s1, "H"))
	fmt.Println(strings.Index(s1, "o"))
	fmt.Println(strings.ToUpper(s1))
}