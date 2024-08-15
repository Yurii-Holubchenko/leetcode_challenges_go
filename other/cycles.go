package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Until/While cycle: ")
	num := 0
	for num < 5 {
		fmt.Printf("%d ", num)
		num += 1
	}
	fmt.Println()

	fmt.Print("For cycle: ")
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	systemCheck(runtime.GOOS)
	saturdayCheck(time.Now().Weekday())
	timeOfDay(time.Now())
	deferTry()
	deferMultiple()
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func systemCheck(osVal string) {
	fmt.Print("Go runs on ")

	switch os := osVal; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func saturdayCheck(today time.Weekday) {
	fmt.Print("When is Saturday? ")

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func timeOfDay(t time.Time) {
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferTry() {
	defer fmt.Println("hello!")

	fmt.Print("Say ")
}

func deferMultiple() {
	fmt.Print("counting ")

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Print("done: ")
}
