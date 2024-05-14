package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Email string `json:"email"`
}

func inspectVariable(variable interface{}){
	t:= reflect.TypeOf(variable)
	v:= reflect.ValueOf(variable)
	fmt.Println("Type: ", t)
	fmt.Println("Value: ", v)
}


func main(){
	person := Person{Name: "Tahmid", Age: 26, Email: "tahmidakter56@gmail.com"}

	// inspectVariable(person)

	//Marshaling
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Error marshaling to Json: %s", err)
	}

	fmt.Println(string(jsonData))

	//Unmarshaling
	var p2 Person

	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		log.Fatalf("Error unmarshaling Json: %s", err)
	}

	fmt.Println(p2)
}