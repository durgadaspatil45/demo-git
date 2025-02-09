// 1417. Reformat The String
// Easy
// Topics
// Companies
// Hint
// You are given an alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).

// You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.

// Return the reformatted string or return an empty string if it is impossible to reformat the string.

// Example 1:

// Input: s = "a0b1c2"
// Output: "0a1b2c"
// Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
// Example 2:

// Input: s = "leetcode"
// Output: ""
// Explanation: "leetcode" has only characters so we cannot separate them by digits.

package main

import (
	"fmt"
)

func main() {
	s := "1a2b3c4dr"
	s1 := "asdf123"
	fmt.Println("Lenght", len(s), len(s1), len(s)/2, len(s1)/2)
	output := reformatString(s)
	if len(output) == 0 {
		fmt.Println("Not Valid string to make foramt")
	} else {
		fmt.Printf("Output of formatted string is - %s", output)
	}

}

// this function refacters the string
func reformatString(s string) string {
	number := 0
	character := 0
	runeArrChar := []rune{}
	runeArrDigit := []rune{}
	finalString := []rune{}
	for i, v := range s {
		if v >= 48 && v <= 57 {
			runeArrDigit = append(runeArrDigit, v)
			number++
		} else if v >= 97 && v <= 122 {
			runeArrChar = append(runeArrChar, v)
			character++
		}
		fmt.Println(i, v)
	}
	halfLen := len(s) / 2
	fmt.Printf("Number %v, Chaer %v, half %v\n", number, character, halfLen)
	if halfLen == character && halfLen == number || halfLen <= (character+1) && halfLen <= (halfLen+1) || halfLen == character && halfLen <= (halfLen+1) || halfLen <= (character+1) && halfLen == number {
		if character > number {
			for i, v := range runeArrChar {
				finalString = append(finalString, v)
				finalString = append(finalString, runeArrDigit[i])
			}
		} else {
			for i, v := range runeArrDigit {
				finalString = append(finalString, v)
				finalString = append(finalString, runeArrChar[i])
			}
		}

	} else {
		return ""
	}
	return string(finalString)
}
