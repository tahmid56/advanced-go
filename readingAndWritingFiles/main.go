package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	
	"os"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
}

func saveBooks(filename string, books []Book) error{
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, book := range books {
		jsonData, err := json.Marshal(book)
		if err != nil {
			return err
		}
		_, err = writer.WriteString(string(jsonData) + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func loadBooks(fileName string) ([]Book, error) {
	var books []Book
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var book Book
		if err := json.Unmarshal([]byte(scanner.Text()), &book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, scanner.Err()
}

func main(){
	books:= []Book{
		{"The Go Programming Language", "Alan A. A. Donovan", 380},
		{"Go in Action", "William Kennedy", 300},
	}

	filename := "books.json"
	if err := saveBooks(filename, books); err != nil {
		fmt.Println("Error saving books:", err)
		return
	}

	loadedBooks, err := loadBooks(filename)
	if err != nil {
		fmt.Println("Error loading books:", err)
		return
	}
	fmt.Println("Loaded Books:", loadedBooks)
}