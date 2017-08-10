package encoding

import (
	"math"
	"strings"
)

// Base62Alphabet ... Letters for use in the short URLS.
// Each number, lowercase, and uppercase letter is a distinct character.
var Base62Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Base ... Default base for encoding
var Base = len(Base62Alphabet)

// HashToID ... Take a hash and convert back to a numeric ID
func HashToID(hash string) (result int) {
	length := len(hash)
	for i, r := range hash {
		index := strings.Index(Base62Alphabet, string(r))
		power := math.Pow(float64(Base), float64(length-(i+1)))
		num := int(float64(index) * power)
		result = result + num
	}
	return
}

// IDToHash ... Take a numeric id and convert to a user friendly string
// for use.
func IDToHash(id int) string {
	digits := FindDigitsForInt(id)
	return MapDigitsToAlphabet(digits)
}

func MapDigitsToAlphabet(digits []int) string {
	shortenedURLHash := ""

	for i := range digits {
		// Get the alphabet indice from our converted digits
		indice := digits[i]

		// Get a one letter range to easily get a string and add to the hash
		shortenedURLHash = shortenedURLHash + Base62Alphabet[indice:indice+1]
	}
	return shortenedURLHash
}

// FindDigitsForInt ... Obtain the individual digits that will be used
// to find the replacement letters in our base62 alphabet.
func FindDigitsForInt(dividend int) []int {
	if dividend == 0 {
		return []int{0}
	}

	digits := make([]int, 0)

	for dividend > 0 {
		remainder := dividend % Base
		digits = append(digits, remainder)
		dividend = dividend / Base
	}

	return reverse(digits)
}

func reverse(digits []int) []int {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}
