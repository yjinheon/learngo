package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name} // 사람 초기화
	p.age = 42
	return &p // return a pointer to the new person
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{
		name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{
		name: "Ann", age: 40})

	fmt.Println(newPerson("jon"))

	s := person{name: "Sean", age: 50}

	fmt.Println(s.name)

	sp := &s
	fmt.Println(s.age)

	sp.age = 51
	fmt.Println(sp.age)

}
