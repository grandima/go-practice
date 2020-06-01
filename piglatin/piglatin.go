package piglatin

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

)

var vowels = map[rune]bool {
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'y': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
	'Y': true,
}

const (
	suffix  string = "ay"
	vowSuffix string = "y" + suffix
)

func Encode(s string) string {
	if len(s) < 1 {
		return s
	}

	first := []rune(s)[0]

	if !unicode.IsLetter(first) {
		return s
	}

	if _, ok := vowels[first]; ok {
		return encodeVowel(s)
	}

	return encodeConsonant(s)
}

func encodeVowel(s string) string {
	return s + vowSuffix
}

func encodeConsonant(s string) string {
	prefix := ""

	for _, r := range s {
		if _, ok := vowels[r]; ok {
			break
		}
		prefix += string(r)
	}

	return strings.TrimLeft(s, prefix) + prefix + suffix
}

func translateToPig(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		wordWithoutSeparators := removeSeparators(w)
		w = strings.Replace(w, wordWithoutSeparators, Encode(wordWithoutSeparators), -1)
		words[i] = w
	}
	return strings.Join(words, " ")
}

const separators = ",;.-"

func removeSeparators(w string) string {
	if len(w) < 1 {
		return w
	}
	for _, c := range separators {
		w = strings.ReplaceAll(w, string(c), "")
	}
	return w
}

func PigLatin() {
	var result string

	fmt.Println("Enter text to translate to pig:")
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	if s.Scan() {
		text := s.Text()
			result = translateToPig(text)
		fmt.Println(result)
	}
	if err := s.Err(); err != nil {
		log.Printf("reading input: %s", err)
	}
}

