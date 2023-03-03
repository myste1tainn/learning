package main

import (
	"fmt"
	"strconv"
)

func ruleOfFunctional() {
	// #1 Any function must taking input
	// #2 Any function must returning output
	// #3 Any function must be immutable

	// Pros
	// #1 Deterministic
	// #2 Testable
	// #3 Readability, conciseness
	// #4 Reusability

	// Cons
	// #1 Sucks at performance
}

// functional
func add(a, b int) int {
	return a + b
}

// not functional
func greetName(name string) {
	fmt.Println("Hello, " + name)
}

// not functional
func greetIce() {
	fmt.Println("Hello, Ice")
}

// not functional
func greetIceText() string {
	return "Hello, Ice"
}

var name string = "john"

// not functional
func greetCurrentText(n string) string {
	name = n
	return "Hello, " + name
}

func example() {
	var arrayOfNumbers []int = []int{1, 2, 3}
	var arrayOfGreets []string
	for _, v := range arrayOfNumbers {
		// filteration
		if v%2 == 0 {
			continue
		}
		// transformation
		str := strconv.Itoa(v)
		// transformation
		str = "Hello, " + str
		// resulting
		fmt.Println(str)
		arrayOfGreets = append(arrayOfGreets, str)
	}

	// var list = arrayOfNumbers
	//     .filter { it -> it % 2 != 0 }
	//     .map { it -> strconv.Itota(v) }
	//     .map { it -> "Hello, " + str }

	// list.forEach { it -> fmt.Println(it) }

	// arrayOfNumbers.takeOddNumber.toString.addPrefix("Hello, ")

	var oddNumber = func(n int) bool {
		return n%2 != 0
	}

	var toString = func(n int) string {
		return strconv.Itoa(n)
	}

	var createGreet = func(name string) string {
		return "Hello, " + name
	}

	// var greetTexts = arrayOfNumbers.filter(oddNumber).map(toString).map(createGreet)

	processArrayOfNumbers(
		arrayOfNumbers,
		func(n int) bool { return n%2 != 0 },
		func(n int) string { return strconv.Itoa(n) },
		func(name string) string { return "Hello, " },
	)
	processArrayOfNumbers(
		arrayOfNumbers,
		func(n int) bool { return n%2 != 0 },
		func(n int) string { return strconv.Itoa(n) },
		func(name string) string { return "Goodbye, " },
	)
	processArrayOfNumbers(
		arrayOfNumbers,
		oddNumber,
		toString,
		createGreet,
	)

	processArrays[string, int](
		[]string{"0", "1", "10"},
		func(n string) bool { return len(n) == 1 },
		func(n string) int {
			number, _ := strconv.Atoi(n)
			return number
		},
		func(name int) int { return name * 20 },
	)
}

func processArrayOfNumbers(n []int, filter func(n int) bool, transformOne func(n int) string, transformTwo func(name string) string) []string {
	var result []string
	for _, v := range n {
		// filteration
		ok := filter(v)
		if !ok {
			continue
		}

		// transformation
		str := transformOne(v)
		// transformation
		str = transformTwo(str)
		// resulting
		fmt.Println(str)
		result = append(result, str)
	}
	return result
}

func processArrays[Input comparable, Output any](n []Input, filter func(n Input) bool, transformOne func(n Input) Output, transformTwo func(name Output) Output) []Output {
	var result []Output
	for _, v := range n {
		// filteration
		ok := filter(v)
		if !ok {
			continue
		}

		// transformation
		str := transformOne(v)
		// transformation
		str = transformTwo(str)
		// resulting
		fmt.Println(str)
		result = append(result, str)
	}

	return result
}
