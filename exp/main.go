package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Dog  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John Smith",
		Dog:  "Rick",
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Roy Wei"
	data.Dog = "Ricky"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
