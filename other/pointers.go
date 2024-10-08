package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	X, Y int
}

type Coords struct {
	Lat, Long float64
}

func main() {
	pointerInt()
	vertex := createVertex()
	changeVertex(vertex)

	p := &vertex
	p.X = 10
	fmt.Println(vertex)

	assignMultipleVertex()

	arrayInit()
	arraySlice()
	namesSlice()
	sliceLiterals()
	sliceDefaults()
	sliceAndCapacity()
	nilSlice()
	sliceWithMake()
	ticTacToe()
	extendArrayCapacity(7)

	r := []int{1, 2, 3}
	r = appendInt(r, 5, 6, 7)
	fmt.Println(r)
	r = appendInt(r, 9, 10, 11)
	fmt.Println(r)

	appendToSlice()
	forRangeSimple()
	forRangeOmitValues()
	mapSimple()
	mapMultiple()
	mapMutation()
	fmt.Println(wordCount("a b c a a c b b c a b c c c a a b c a"))
}

func pointerInt() {
	i, j := 42, 65

	// Make p as pointer to variable i
	p := &i
	fmt.Println("p =", p)
	fmt.Println("pointer p =", *p)
	// Change value of variable i through pointer p
	*p = 21
	fmt.Println("i = ", i)

	// Re-assign pointer p from variable j to variable i
	p = &j
	// Change value of variable j through pointer p
	*p = *p / 5
	fmt.Println("j =", j)
}

func createVertex() Vertex {
	v := Vertex{1, 2}
	fmt.Println("Vertex =", v)
	return v
}

func changeVertex(v Vertex) {
	v.X = 5
	fmt.Println("Vertex.X =", v.X)
}

func assignMultipleVertex() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{Y: 5}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)

	fmt.Println(v1, v2, v3, p)
}

func arrayInit() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	longArray := [5]int{5, 4, 3, 2, 1}
	fmt.Println(longArray)
}

func arraySlice() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var subArr []int = arr[3:5] // arr[3..4]
	fmt.Println(subArr)
}

func namesSlice() {
	names := [4]string{
		"Alex",
		"Tina",
		"John",
		"Fred",
	}
	fmt.Println(names)

	a := names[0:2] // names[0..1]
	b := names[1:3] // names [1..2]
	fmt.Println(a, b)

	b[0] = "Emma"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
		s string
	}{
		{2, true, "qqq"},
		{3, false, "www"},
		{5, true, "eee"},
		{7, true, "aaa"},
		{11, false, "sss"},
		{13, true, "ddd"},
	}
	fmt.Println(s)
}

func sliceDefaults() {
	a := []int{1, 2, 3, 4, 5, 6}

	a = a[:]
	fmt.Println(a)

	a = a[1:3]
	fmt.Println(a)

	a = a[:2]
	fmt.Println(a)

	a = a[1:]
	fmt.Println(a)
}

// Length - actual length of an array.
// Capacity - possible max length of an array.
// Slice array from the end: length decrease, capacity same.
// Slice array from the start: length decrease, capacity decrease.
func sliceAndCapacity() {
	s := []int{0, 1, 2, 3, 4, 5}

	printSlice(s)
	printSlice(s[:0])
	printSlice(s[:4])
	printSlice(s[3:6])
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSlice() {
	var a []int
	printSlice(a)

	if a == nil {
		fmt.Println("slice is empty")
	}
}

func sliceWithMake() {
	a := make([]int, 5)
	printSlice(a)

	// Second argument - length
	// Third argument - capacity
	b := make([]int, 5, 7)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := b[2:cap(b)]
	printSlice(d)

	e := b[len(b):cap(b)]
	printSlice(e)
}

func ticTacToe() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		row := strings.Join(board[i], " ")
		fmt.Printf("%s\n", row)
	}
}

func extendArrayCapacity(newCapacity int) {
	a := []int{1, 2, 3}
	t := make([]int, len(a), newCapacity)
	copy(t, a)
	a = t
	fmt.Println(a, len(a), cap(a))
}

func appendInt(slice []int, data ...int) []int {
	// [1, 2, 3]   					[5, 6, 7]
	// [1, 2, 3, x, x, x]   [5, 6, 7]

	sliceLength := len(slice)
	newLength := sliceLength + len(data)

	if newLength > cap(slice) {
		newSlice := make([]int, newLength)
		copy(newSlice, slice)
		// [1, 2, 3, x, x, x]
		slice = newSlice
	}

	// [1, 2, 3, x, x, x]
	slice = slice[0:newLength]
	// [x, x, x] << [5, 6, 7]
	copy(slice[sliceLength:newLength], data)
	return slice
}

func appendToSlice() {
	a := make([]int, 1)
	a = append(a, 1, 2, 3)
	printSlice(a)

	b := []string{"aaa", "bbb", "ccc"}
	c := []string{"ddd", "eee"}
	b = append(b, c...)
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	var s []int
	s = append(s, 0, 1, 2)
	printSlice(s)
}

func forRangeSimple() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d == %d\n", i, v)
	}
}

func forRangeOmitValues() {
	pow := make([]int, 7)

	for i := range pow {
		pow[i] = int(math.Pow(2, float64(i)))
	}

	for _, value := range pow {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func mapSimple() {
	var m map[string]Coords
	m = make(map[string]Coords)
	m["Africa"] = Coords{40.68433, -74.39967}

	fmt.Println(m)
	fmt.Println(m["Africa"])
}

func mapMultiple() {
	var countries = map[string]Coords{
		"Italy": {40.68433, -74.39967},
		"Spain": {37.42202, -122.08408},
	}

	fmt.Println(countries)
	for name, coords := range countries {
		fmt.Printf("%v: %v\n", name, coords)
	}

	for name := range countries {
		fmt.Println(countries[name])
	}
}

func mapMutation() {
	answers := make(map[string]int)
	name := "First"

	answers[name] = 15
	fmt.Printf("%v - %v\n", name, answers[name])

	answers[name] = 17
	fmt.Printf("%v - %v\n", name, answers[name])

	delete(answers, name)
	fmt.Printf("%v - %v\n", name, answers[name])

	val, ok := answers[name]
	fmt.Println("The value:", val, "present?", ok)
}

func wordCount(str string) map[string]int {
	wordsMap := make(map[string]int)
	words := strings.Fields(str)

	for _, word := range words {
		_, ok := wordsMap[word]
		if ok {
			wordsMap[word] += 1
		} else {
			wordsMap[word] = 1
		}
	}

	return wordsMap
}
