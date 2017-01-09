/*
Find the largest Palindrome within the given text. Asks user to enter text to be checked.

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/asaskevich/govalidator"
)

func main() {

	text := GetValidText()
	fmt.Println("This palindrome was found: " + FindPalindrome(text, 0))
}

// Ask the user to enter text.
// Validate text to make sure it only contains alphabetic characters.
func GetValidText() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter text to check for the largest palindrome. ")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		panic(`Failed to read input.`)
	}

	if text == "\n" {
		log.Fatal(`No text entered.`)
		panic(`Console input is required, but none entered.`)
	}

	if !govalidator.IsAlpha(strings.TrimSuffix(text, "\n")) {
		log.Fatal(`None alphabetic text eneted.`)
		panic(`None alphabetic text eneted.`)
	}

	return text
}

// Recursively check for palindromes
func FindPalindrome(text string, offset int) string {

	a := []byte(strings.ToLower(text))

	if offset > len(text)-2 {
		return ""
	}

	for i := 0; i <= offset; i++ {
		sub := string(a[i : len(a)-offset+i])
		// fmt.Println("Checking " + sub)

		if sub == Reverse(sub) {
			return sub
		}
	}

	return FindPalindrome(text, offset+1)
}

// Returns reversed text
func Reverse(text string) string {
	a := []byte(text)

	// Reverse a
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return string(a[:])
}
