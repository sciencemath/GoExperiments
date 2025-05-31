package main

import (
	"fmt"
)

func main() {
	// fmt.Println("The answer to everything is 42")

	var name string = "Mathias"
	fmt.Printf("This is my name %s\n", name)

	age := 28
	fmt.Printf("This is my age %d\n", age)

	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("its the weekend")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("this is i", i)
	}

	counter := 0
	for counter < 3 {
		fmt.Println("this is loop is like a while", counter)
		counter++
	}

	iterations := 0
	for {
		if iterations > 3 {
			break
		}
		iterations++
	}

	// arrays cannot change, holds the same types, reassigning is ok
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("this is out array %v\n", numbers)
	fmt.Println("this is the last value", numbers[len(numbers)-1])

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("this is out matrix: %v\n", matrix)

	// allNumbers := numbers[:]
	// firstThree := numbers[0:3]

	fruits := []string{"apple", "banana", "kiwi"}
	fmt.Printf("these are my fruits %v\n", fruits)

	fruits = append(fruits, "orange")
	fmt.Printf("these are my fruits with orange %v\n", fruits)

	fruits = append(fruits, "mango", "pineapple")
	fmt.Printf("these are my fruits with more fruit %v\n", fruits)

	moreFruits := []string{"blueberries", "grapefruit"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("these are my fruits with even more fruit %v\n", fruits)

	for index, value := range numbers {
		fmt.Printf("index %d and value %d\n", index, value)
	}

	var city string
	city = "St Pete"
	fmt.Printf("this is my city %s\n", city)

	var country, continent string = "USA", "North America"
	fmt.Printf("this is my country %s and this is my continent %s\n", country, continent)

	var (
		isEmployed bool   = true
		salary     int    = 50000
		position   string = "developer"
	)

	fmt.Printf("isEmployed: %t this is my salary: %d and this is my position: %s\n", isEmployed, salary, position)

	// Zero Values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("int: %d float %f string %s and bool %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	fmt.Printf("Monday: %d - Tuesday %d Wednesday %d\n", Monday, Tuesday, Wednesday)

	const typedAge int = 25
	const untypedAge = 25

	fmt.Println(typedAge == untypedAge)

	const (
		Jan = iota + 1 // 1
		Feb            // 2
		Mar            // 3
		Apr            // 4
	)

	fmt.Printf("jan - %d feb - %d mar - %d apr - %d\n", Jan, Feb, Mar, Apr)

	result := add(3, 4)
	fmt.Println("this is the result", result)

	sum, product := calculateSumAndProduct(10, 10)
	fmt.Printf("this is sum: %d and this is product: %d\n", sum, product)

	capitalCities := map[string]string{
		"USA":   "Washington D.C.",
		"India": "New Delhi",
		"UK":    "London",
	}

	fmt.Println(capitalCities["USA"])

	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Println("this is the capital", capital)
	} else {
		fmt.Println("Does not exists")
	}

	delete(capitalCities, "UK")
	fmt.Printf("this is after deleted map: %v", capitalCities)
}

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
