package main

import "fmt"

type person struct {
	name string
	age	 uint
	address
}

type address struct {
	province string
	city	 string
}

type Stringer interface {
	String() string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func NewPerson(name string) *person {
	return &person{name: name}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWrite interface {
	Reader
	Writer
}

func main() {
	//var p person
	//p := person{"Golang", 5}
	//p := person{age: 7, name: "Golang"}
	p := person{
		age: 5,
		name: "Golang",
		address:address{
			province: "广东",
			city: "广州",
		},
	}
	fmt.Println(p.province, p.city)
	printString(p)
	printString(p.address)

	p1 := NewPerson("小宋")
	var s fmt.Stringer
	s = p1
	p2 := s.(*person)
	fmt.Println(p2)

	a, ok := s.(address)
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("s 不是一个 address")
	}
	fmt.Println(a)

}