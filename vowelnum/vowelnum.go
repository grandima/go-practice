package vowelnum

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var encoderMap = []string{
	"a", "1",
	"e", "2",
	"i", "3",
	"o", "4",
	"u", "5",
	"y", "6",
}

var decoderMap = []string{
	"1", "a",
	"2", "e",
	"3", "i",
	"4", "o",
	"5", "u",
	"6", "y",
}

func convertText(mapping []string, text string) string {
	r := strings.NewReplacer(mapping...)
	lowerCase := strings.ToLower(text)
	return r.Replace(lowerCase)
}

func VowelConvert(decode bool) {
	var result string

	fmt.Println("Enter vowel text:")
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	if s.Scan() {
		text := s.Text()
		if decode {
			result = convertText(decoderMap, text)
		} else {
			result = convertText(encoderMap, text)
		}
		fmt.Println(result)
	}
	if err := s.Err(); err != nil {
		log.Printf("reading input: %s", err)
	}
}