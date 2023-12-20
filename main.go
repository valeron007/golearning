package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:c"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Bogart", "Bergman"}},
	{Title: "Cool hand luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullit", Year: 1968, Color: true, Actors: []string{"Steve McQuuen", "Bisset"}},
}

func main() {
	var dilbert Employee
	dilbert.Salary = -5000
	position := &dilbert.Position
	*position = "Senior" + *position
	fmt.Printf("position=%s\n", string(*position))
	var employee *Employee = &dilbert
	employee.Position += " (активный участник команды)"
	fmt.Printf("Change position=%s\n", string(*position))
	fmt.Println(employee.Position)

	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("Сбой маршалинга json: %s", err)
	}
	fmt.Printf("%s", data)
}
