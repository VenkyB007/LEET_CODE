// leet - 5

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s []string) bool {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	str := strings.Split(s, "")
	max := 1
	result := ""
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		first := string(str[i])
		var checkLetters []string
		checkLetters = append(checkLetters, first)
		for j := i + 1; j < len(str); j++ {
			next := string(str[j])
			checkLetters = append(checkLetters, next)
			if isPalindrome(checkLetters) {
				if len(checkLetters) > max {
					max = len(checkLetters)
					result = ""
					result = strings.Join(checkLetters, "")
				}
			}
		}
	}
	if result == "" {
		if len(s) != 0 {
			return string(s[0])
		}
	}
	return result
}

func main() {
	res := longestPalindrome("CCC")
	fmt.Println(res)
}
