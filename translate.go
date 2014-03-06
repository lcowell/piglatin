package piglatin

import (
	"strings"
)

const (
	pigLatinSuffix             string = "ay"
	firstLetterExceptions      string = "aeiou"
	firstLetterExceptionSuffix string = "d" + pigLatinSuffix
)

func TranslateMultiple(in ...string) (results []string) {
        for _, word := range in {
                results = append(results, Translate(word))
        }
        return
}

// Translate translates an English word into Pig latin.
func Translate(in string) string {
	first := in[0:1]
	if strings.Contains(firstLetterExceptions, first) {
		return in + firstLetterExceptionSuffix
	} else {
		return in[1:] + first + pigLatinSuffix
	}
}
