package kind

import "regexp"

type Adjective struct {
	word string
	gen  string
}

func BuildAdjective(input string) *Adjective {

	wLen := len(input)

	if wLen < 4 {
		return nil
	}

	ending := input[wLen-4:]

	switch ending {
	case "ая":
		return &Adjective{gen: GEN_FEMALE, word: input}
	case "яя":
		return &Adjective{gen: GEN_FEMALE, word: input}
	case "ий":
		return &Adjective{gen: GEN_MALE, word: input}
	case "ый":
		return &Adjective{gen: GEN_MALE, word: input}
	case "ое":
		return &Adjective{gen: GEN_MIDDLE, word: input}
	case "ее":
		return &Adjective{gen: GEN_MIDDLE, word: input}
	}

	return nil
}

func (a *Adjective) Gen() string {
	return a.gen
}

func (a *Adjective) Word() string {
	return a.word
}

func (a *Adjective) Locative(_ string) (string, bool) {
	wl := len(a.word)
	ending := a.word[wl-4:]

	switch ending {
	case "ая":
		return regexp.MustCompile("ая$").ReplaceAllString(a.word, "ой"), true
	case "яя":
		return regexp.MustCompile("яя$").ReplaceAllString(a.word, "ей"), true
	case "ий":
		e3 := a.word[wl-6 : wl-4]
		if e3 == "н" {
			return regexp.MustCompile("ий$").ReplaceAllString(a.word, "ем"), true
		}

		return regexp.MustCompile("ий$").ReplaceAllString(a.word, "ом"), true
	case "ый":
		return regexp.MustCompile("ый$").ReplaceAllString(a.word, "ом"), true
	case "ое":
		return regexp.MustCompile("ое$").ReplaceAllString(a.word, "ом"), true
	case "ее":
		return regexp.MustCompile("ee$").ReplaceAllString(a.word, "eм"), true
	}

	return a.word, false
}
