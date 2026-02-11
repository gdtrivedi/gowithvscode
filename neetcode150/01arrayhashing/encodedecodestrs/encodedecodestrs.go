package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := &Solution{}

	encodedstr := s.Encode([]string{"hello", "world", "how", "are", "you"})
	fmt.Println("Encoded String:", encodedstr)

	decodedstrs := s.Decode(encodedstr)
	fmt.Println("Decoded Strings:", decodedstrs)
}

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encodedstr := ""
	for _, s := range strs {
		encodedstr = encodedstr + strconv.Itoa(len(s)) + "#" + s
	}
	return encodedstr
}

func (s *Solution) Decode(encoded string) []string {
	encodedstr := encoded
	result := make([]string, 0)
	for len(encodedstr) > 0 {
		idx := strings.Index(encodedstr, "#")
		n := encodedstr[0:idx]
		ni, _ := strconv.Atoi(n)
		end := idx + 1 + ni

		s1 := encodedstr[idx+1 : end]

		result = append(result, s1)

		if len(encodedstr) > end {
			encodedstr = encodedstr[end:]
		} else {
			encodedstr = ""
		}
	}
	return result
}
