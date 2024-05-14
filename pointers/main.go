package main

import "fmt"

type Person struct {
	name string
	age int
}

func main(){
	var a int = 58

	var p *int = &a

	fmt.Println("Address of a:", p)

	fmt.Println("Value of a through pointer p:", *p)

	person := Person{name: "John Doe", age: 21}

	fmt.Println("person's age before change: ", person.age)
	changeAge(&person)
	fmt.Println("person's age after change: ", person.age)
}

func changeAge(p *Person){
	p.age ++
}

