package main

import (
	"fmt"
	"unicode/utf8"
)

func ParseSlices(s *[]string) []string{
	temp := make(map[string]struct{}, len(*s))
	for i := 0; i < len(*s); {
		if _, ok := temp[(*s)[i]]; ok {
			*s = append((*s)[:i], (*s)[i+1:]...)
			continue
		}
		temp[(*s)[i]] = struct{}{}
		i++
	}
	return *s
}

func main()  {
	//array := [5]string{"a", "b", "c", "d", "e"}
	array := [...]string{"a", "b", "c", "d", "e"}

	array1 := [5]string{1:"b", 3:"d"}
	array2 := [...]string{1:"b", 3:"d"}
	fmt.Println("array1 len is: ", len(array1))
	fmt.Println("array2 len is: ", len(array2))

	fmt.Println(array[2])

	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引: %d, 对应值: %s\n", i, array[i])
	}

	for i, v := range array {
		fmt.Printf("数组索引: %d, 对应值: %s\n", i, v)
	}

	slice := array[2:5]
	fmt.Println(slice)

	slice[1] = "f"
	fmt.Println(array)

	/*slice1 := make([]string, 4)
	slice1 := make([]string, 4, 8)*/

	slice1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice1), cap(slice1))

	/*slice2 := append(slice1, "f")
	slice2 := append(slice1, "f", "g")
	slice2 := append(slice1, slice...)*/

	nameAgeMap := make(map[string]int)
	//nameAgeMap := map[string]int{"Golang": 5}

	nameAgeMap["Golang"] = 20

	age, ok := nameAgeMap["Golang1"]
	if ok {
		fmt.Println(age)
	}

	delete(nameAgeMap, "Golang")

	/*s := []string{"cc", "a", "c", "cc", "b", "a", "c"}
	fmt.Println(ParseSlices(&s))*/

	nameAgeMap["飞雪无情"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22

	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ", Value is", v)
	}

	fmt.Println(len(nameAgeMap))

	s := "Hello飞雪无情"
	bs := []byte(s)

	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[15])
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Println(i, r)
	}
}
