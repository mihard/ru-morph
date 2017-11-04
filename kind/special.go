package kind

import (
	"fmt"
	"regexp"
	"strings"
)

type Special struct {
	word string
	gen  string
}

var numRx *regexp.Regexp

func BuildIfSpecial(input []string) (*Special, bool) {

	lcFirst := strings.ToLower(input[0])

	if numRx == nil {
		numRx = regexp.MustCompile("^\\d+$")
	}

	if lcFirst == "на" && len(input) == 2 {
		ending := input[1][len(input[1])-2:]

		if ending == "е" || ending == "у" {
			return &Special{word: fmt.Sprintf("-%s-%s", input[0], input[1])}, true
		}
	}

	if lcFirst == "cан" || lcFirst == "санкт" || lcFirst == "лос" {
		return &Special{word: input[0]}, false
	}

	if numRx.MatchString(lcFirst) {
		return &Special{word: input[0]}, false
	}

	return nil, false
}

func (s *Special) Gen() string {
	return s.gen
}

func (s *Special) Word() string {
	return s.word
}

func (a *Special) Locative(_ string) (string, bool) {
	return a.word, false
}
