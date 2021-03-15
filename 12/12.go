package main

import "fmt"

func main()  {
	name := "Hello Golang"
	nameP := &name

	fmt.Println("name变量的值为:", name)
	fmt.Println("nameP变量的内存地址为:", nameP)

	//var intP *int
	//intP = &name

	//intP := new(int)

	nameV := *nameP
	fmt.Println("nameP指针指向的值为:", nameV)

	*nameP = "Hello Golang"
	fmt.Println("nameP指针指向的值为:", *nameP)
	fmt.Println("name变量的值为:", name)

	var intP *int = new(int)
	*intP = 10

	fmt.Println(*intP)

	age := 18
	modifyAge(&age)
	fmt.Println("age的值为:", age)
}

func modifyAge(age *int) {
	*age = 20
}