package main

import ("fmt")

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

func (c Cat) Speak() string{
	return "Meow! My name is " + c.Name
}

func introduceSpeaker(s Speaker){
	fmt.Println(s.Speak())
}

func main(){
	dog := Dog{Name: "Tom"}
	cat := Cat{Name: "Bil"}
	introduceSpeaker(dog)
	introduceSpeaker(cat)
}
