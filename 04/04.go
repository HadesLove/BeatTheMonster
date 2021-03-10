package main

import (
	"errors"
	"fmt"
)

/*func funcName(params) result {
	body
}*/

/*func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	return a + b, nil
}*/

func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}
// func Println(a ...interface{}) (n int, err error)

func sumAll(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}

	return sum
}

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	result, err := sum(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(sumAll(1, 2))
	fmt.Println(sumAll(1, 2, 3))
	fmt.Println(sumAll(1, 2, 3, 4))

	sum2 := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum2(1, 3))

	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	age := Age(25)
	age.String()
	(&age).Modify()
	age.String()
}

type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}

func (age *Age) Modify() {
	*age = Age(30)
}