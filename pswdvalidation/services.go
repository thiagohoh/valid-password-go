package pswdvalidation

import "unicode"

type unicodeFn func(rune) bool

// unicodeChecker apply a unicode function to given string and check it's total.
func unicodeChecker(p string, v int, fn unicodeFn) bool {
	total := 0
	runes := []rune(p)

	for _, ch := range runes {
		if fn(ch) {
			total++
		}
	}
	return total >= v
}

// MinSize returns true if p has a minimun length of v.
func MinSize(p string, v int) bool {
	return len(p) >= v
}

// MinUpperCase returns true if p has at least v upper case characters.
func MinUpperCase(p string, v int) bool {
	return unicodeChecker(p, v, unicode.IsUpper)
}

// MinLowerCase returns true if p has at least v lower case characters.
func MinLowerCase(p string, v int) bool {
	return unicodeChecker(p, v, unicode.IsLower)
}

// MinDigit returns true if p has at least v digits.
func MinDigit(p string, v int) bool {
	return unicodeChecker(p, v, unicode.IsDigit)
}

// MinSpecialChars true if p has at least v special characters.
func MinSpecialChars(p string, v int) bool {
	list := []rune("!@#$%^&*()-+\\/{}[]")

	total := 0
	runes := []rune(p)
	for _, ch := range runes {
		for _, s := range list {
			if ch == s {
				total++
			}
		}
	}
	return total >= v
}

// HasRepeated returns true if p has consecutive repeated char
func HasRepeated(p string) bool {
	length := len(p)

	for j := 0; j < length; j++ {
		if j+1 >= length {
			return false
		}
		if p[j] == p[j+1] {
			return true
		}
	}
	return false
}
