package main

import (
	"fmt"
	"log"
	"os"
)

func sequenceGenerator() func() int {
	i := 0
	return func() int {
		i+= 1
		return i
	}
}

func readfile(fileName string){
	file, err := os.Open(fileName)
	if err != nil {
	log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	fmt.Println(file)

}

func main(){
	nextNumber := sequenceGenerator()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	readfile("hello.txt")
}