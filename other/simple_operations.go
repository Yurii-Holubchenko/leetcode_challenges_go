package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Automatically assigned default value - false
var python, java bool
var ruby = "yes!"

func main() {
	fmt.Println("Random number", rand.Intn(10))
	fmt.Printf("Sqrt number %g \n", math.Sqrt(16))
	fmt.Println("Pi", math.Pi)

	sum := func(a, b int) int { return a + b }(3, 4)
	fmt.Println("Sum by anonymized func", sum)

	var a [1]int
	a[0] = 1
	fmt.Println("Tiny array", a)

	fmt.Println("Sum by named func", add(3, 10))

	str1, str2 := swap("Hello", "World")
	fmt.Println("Swapped strings", str1, str2)

	fmt.Println(split(17))

	var iterator int
	var inter int = 2
	fmt.Println(iterator, inter, python, java, ruby)

	x, y := 2, 4
	f := math.Sqrt(float64(x*x + y*y))
	u := uint(f)
	fmt.Println(x, y, f, u)

	v := 2.13
	fmt.Printf("v is of type %T\n", v)

	const World = "Bla"
	fmt.Println(World)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Returns both x and y as a results
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
