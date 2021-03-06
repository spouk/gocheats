package utils

import (
	"math/rand"
	"errors"
	"time"
)

var (
	letters        = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	asciiLowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	asciiLetters   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	asciiUppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits         = []rune("0123456789")
	hexdigits      = []rune("0123456789abcdefABCDEF")
)

const (
	LasciiLetters   = 1 << iota
	LasciiUppercase
	Ldigits
	Lhexdigits
	LasciiLowercase
	Lletters
)

type Randomizer struct{}

func NewRandomize() *Randomizer {
	rand.Seed(time.Now().UTC().UnixNano())
	return new(Randomizer)
}
func (r *Randomizer) Hexdigits() string {
	return string(hexdigits)
}
func (r *Randomizer) Digits() string {
	return string(digits)
}
func (r *Randomizer) AsciiUppercase() string {
	return string(asciiUppercase)
}
func (r *Randomizer) Letters() string {
	return string(letters)
}
func (r *Randomizer) AsciiLowercase() string {
	return string(asciiLowercase)
}
func (r *Randomizer) AsciiLetters() string {
	return string(asciiLetters)
}
func (r *Randomizer) RandomString(count int) string {
	var result = make([]rune, count)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
func (r *Randomizer) RandomStringChoice(count int, flag int) string {
	var result = make([]rune, count)

	for i := range result {
		switch flag {
		case LasciiLetters:
			result[i] = asciiLetters[rand.Intn(len(asciiLetters))]
		case LasciiLowercase:
			result[i] = asciiLowercase[rand.Intn(len(asciiLowercase))]
		case LasciiUppercase:
			result[i] = asciiUppercase[rand.Intn(len(asciiUppercase))]
		case Ldigits:
			result[i] = digits[rand.Intn(len(digits))]
		case Lhexdigits:
			result[i] = hexdigits[rand.Intn(len(hexdigits))]
		case Lletters:
			result[i] = letters[rand.Intn(len(letters))]
		default:
			panic(errors.New("wrong type flag"))
		}
	}
	return string(result)
}
func (r *Randomizer) RandomStringSliceChoice(countLen, lengthEach int, flag int) []string {
	var result []string
	for i := 0; i < countLen; i++ {
		var element = make([]rune, lengthEach)
		for x := range element {
			switch flag {
			case LasciiLetters:
				element[x] = asciiLetters[rand.Intn(len(asciiLetters))]
			case LasciiLowercase:
				element[x] = asciiLowercase[rand.Intn(len(asciiLowercase))]
			case LasciiUppercase:
				element[x] = asciiUppercase[rand.Intn(len(asciiUppercase))]
			case Ldigits:
				element[x] = digits[rand.Intn(len(digits))]
			case Lhexdigits:
				element[x] = hexdigits[rand.Intn(len(hexdigits))]
			case Lletters:
				element[x] = letters[rand.Intn(len(letters))]
			default:
				panic(errors.New("wrong type flag"))
			}
		}
		result = append(result, string(element))
	}
	return result
}

func (r *Randomizer) RandomStringSlice(countLen, lengthEach int) []string {
	var result []string
	for i := 0; i < countLen; i++ {
		var element = make([]rune, lengthEach)
		for x := range element {
			element[x] = letters[rand.Intn(len(letters))]
		}
		result = append(result, string(element))
	}
	return result
}
