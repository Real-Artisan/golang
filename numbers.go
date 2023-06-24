package main

import "fmt"

func main() {
	var numbers [3]int

	// Get user input for three numbers
	for i := 0; i < 3; i++ {
		fmt.Scan(&numbers[i])
		if numbers[i] < 0 || numbers[i] > 10 {
			fmt.Println("Invalid input. Please enter a number between 0 and 10.")
			i-- // Retry input
		}
	}

	// Convert numbers to English texts
	englishTexts := make([]string, 3)
	for i, num := range numbers {
		englishTexts[i] = convertToText(num)
	}

	// Output the English texts
	for _, text := range englishTexts {
		fmt.Println(text)
	}
}

// Function to convert a number to its corresponding English text
func convertToText(number int) string {
	// Array of English text for numbers 0 to 10
	numberTexts := []string{
		"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	}

	return numberTexts[number]
}
