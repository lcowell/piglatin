package piglatin

import (
        "io"
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

type Writer struct {
        w io.Writer
}

func NewWriter(w io.Writer) *Writer {

        return &Writer{w: w}
}

func (w *Writer) Write(p []byte) (n int, err error) {
        words := strings.Split(string(p), " ")
        wordIndex := len(words) - 1

        for n, word := range words {
                w.w.Write([]byte(Translate(word)))
                if n < wordIndex {
                        w.w.Write([]byte(" "))
                }
        }
        return
}
