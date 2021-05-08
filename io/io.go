package io

import (
	"fmt"
	"strconv"
)

func GetOption(text string, options []string) int {
	var sel string

	fmt.Println("\n" + text)
	for i, descr := range options {
		fmt.Printf("%d: %s\n", i+1, descr)

	}

	for {
		fmt.Printf("\nSelect option (1-%d): ", len(options))
		fmt.Scanln(&sel)
		index := getOptIndex(sel, len(options))

		if index != -1 {
			return index
		}
		fmt.Printf("Invalid input. ")
	}
}

func GetConfirmation(text string) bool {
	var answer string
	fmt.Printf("\n%s (y/n): ", text)
	for {
		fmt.Scanln(&answer)
		if answer == "y" {
			return true
		}
		if answer == "n" {
			return false
		}
		fmt.Print("Invalid input. Try again (y/n): ")
	}

}

func getOptIndex(sel string, count int) int {

	index, err := strconv.Atoi(sel)
	if err == nil && index > 0 && index <= count {
		return index
	}
	return -1
}

func GetInt(text string) int {
	var input string
	fmt.Printf("%s: ", text)
	for {
		fmt.Scanln(&input)
		index, err := strconv.Atoi(input)
		if err == nil {
			return index
		}
		fmt.Print("Invalid input. Please enter a number: ")
	}

}

func GetString(text string) string {
	var input string
	fmt.Printf("%s: ", text)
	fmt.Scanln(&input)
	return input
}
