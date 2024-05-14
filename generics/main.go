package main

import "fmt"

func Filter[T any](slice []T, criteria func(T) bool) []T{
	var result []T
	for _, v := range slice {
		if criteria(v){
			result = append(result, v)
		}
	}
	return result
}

func main(){
	ints := []int{1, 2, 3, 4, 5}
	even := Filter(ints, func(n int) bool {return n%2 == 0})
	fmt.Println(even)

	strings := []string{"apple", "banana", "cherry"}
	withA := Filter(strings, func(c string) bool {return c[0]=='a'})
	fmt.Println(withA)
}