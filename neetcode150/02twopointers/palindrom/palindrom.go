package main

import (
	"strings"
	"unicode"
)

func main() {
	println(isPalindrome("Was it a car or a cat I saw?"))
	println(isPalindrome("tab a cat"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	lp := 0
	rp := len(s) - 1

	for lp <= rp {
		if !isAlphaNumeric(rune(s[lp])) {
			lp++
			continue
		}

		if !isAlphaNumeric(rune(s[rp])) {
			rp--
			continue
		}

		if s[lp] != s[rp] {
			return false
		}

		lp++
		rp--
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
