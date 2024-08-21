package main

import "fmt"

const helloWorld = "Hello, World!"

func main() {
	NumericTypeConversions()
	IntegerTypeConversions()
	VarDefinition()
	ConstVar()
	VariableNamesNeverUse()
}

func NumericTypeConversions() {
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println("NumericTypeConversions:", sum1, sum2)
}

func IntegerTypeConversions() {
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	fmt.Println("IntegerTypeConversions:", sum3, sum4)
}

func VarDefinition() {
	var x int = 10
	var y = 20
	var a, b int = 10, 20
	var c, d = 10, "hello"
	e := 10
	f, g := 10, "hello"

	fmt.Println("VarDefinition:", x, y, a, b, c, d, e, f, g)
}

func ConstVar() {
	const a = 20

	fmt.Println("ConstVar:", a, helloWorld)
}

func VariableNamesNeverUse() {
	_0 := 0_0
	_𝟙 := 20
	π := 3
	ａ := "hello"              // Unicode U+FF41
	__ := "double underscore" // two underscores
	fmt.Println(_0)
	fmt.Println(_𝟙)
	fmt.Println(π)
	fmt.Println(ａ)
	fmt.Println(__)
}
