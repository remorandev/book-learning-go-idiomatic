package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	ArraysDeclaration()
	MultiDimensionalArrays()
	Slices()
	Append()
	Make()
	SlicingSlice()
	Maps()
	MapsOk()
	MapsDelete()
	MapsEmpty()
	MapsComparing()
	MapsAsSet()
	SetWithStructs()
	Structs()
	AnonymousStructs()
}

func ArraysDeclaration() {
	var x = [3]int{10, 20, 30}
	var y = []int{1, 5: 4, 6, 10: 100, 15}

	fmt.Printf("ArraysDeclaration: %v \n", x)
	fmt.Printf("ArraysDeclaration: %v \n", y)
}
func MultiDimensionalArrays() {
	var x = [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Printf("MultiDimensionalArrays: %v \n", x)
}

func Slices() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Slices Equal:", slices.Equal(x, y)) // prints true
	fmt.Println("Slices Equal:", slices.Equal(x, z)) // prints false
}

func Append() {
	x := []int{1, 2, 3, 4, 5}
	y := append(x, 6)

	fmt.Printf("Append: %v \n", y)
}

func Make() {
	x := make([]int, 5)
	// Length is 0 and capacity is 10
	y := make([]int, 0, 10)

	fmt.Printf("Make: %v \n", x)
	fmt.Printf("Make: %v Capacity: %d \n", y, cap(y))
}

func SlicingSlice() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("SlicingSlice x:", x)
	fmt.Println("SlicingSlice y:", y)
	fmt.Println("SlicingSlice z:", z)
	fmt.Println("SlicingSlice d:", d)
	fmt.Println("SlicingSlice e:", e)
}

func Maps() {
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams)
}

func MapsReadAndWrite() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
}

func MapsOk() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println("MapsOk", v, ok)

	v, ok = m["world"]
	fmt.Println("MapsOk", v, ok)

	v, ok = m["goodbye"]
	fmt.Println("MapsOk", v, ok)
}

func MapsDelete() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Println("MapsDelete", m)
	delete(m, "world")
	fmt.Println("MapsDelete", m)
}

func MapsEmpty() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println("MapsEmpty:", m, len(m))
	clear(m)
	fmt.Println("MapsEmpty:m", m, len(m))
}

func MapsComparing() {
	a := map[string]int{"A": 0}
	b := map[string]int{"B": 42}
	fmt.Println("MapsComparing:", a, b)
	fmt.Println("MapsComparing:", maps.Equal(a, b))
}

func MapsAsSet() {
	stringSet := map[string]bool{}
	vals := []string{"a", "b", "c", "d", "a", "b"}
	for _, v := range vals {
		stringSet[v] = true
	}
	fmt.Printf("ArrayLen: %d  SetLen: %d\n", len(vals), len(stringSet))
	if stringSet["a"] {
		fmt.Println("a is in the set")
	}
}

func SetWithStructs() {
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = struct{}{}
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}

type person struct {
	name string
	age  int
	pet  string
}

func Structs() {
	people := []person{
		{"Fred", 20, "dog"},
		{"Raul", 30, "cat"},
		{"Waldo", 40, "parrot"},
	}
	fmt.Println("Structs:", people)
}

func AnonymousStructs() {
	people := []struct {
		name string
		age  int
		pet  string
	}{
		{"Fred", 20, "dog"},
		{"Raul", 30, "cat"},
		{"Waldo", 40, "parrot"},
	}
	fmt.Println("AnonymousStructs:", people)
}
