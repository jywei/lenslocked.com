package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Dog   string
	Int   int
	Float float64
	Slice []string
	Map   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:  "John Smith",
		Dog:   "Rick",
		Int:   123,
		Float: 3.14,
		Slice: []string{"a", "b", "c"},
		Map: map[string]string{
			"k1": "v1",
			"k2": "v2",
		},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
