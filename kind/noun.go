package kind

import (
	"regexp"

	"github.com/mihard/ru-morph/phonetic"
)

type Noun struct {
	word string
	gen  string
}

func BuildNoun(input string) *Noun {

	wLen := len(input)

	if wLen < 4 {
		return nil
	}

	e1 := input[len(input)-2:]
	e2 := input[len(input)-4 : len(input)-2]

	g := GEN_MALE

	switch e1 {
	case "я":
		g = GEN_FEMALE
	case "а":
		g = GEN_FEMALE
	case "о":
		g = GEN_MIDDLE
	case "е":
		g = GEN_MIDDLE
	case "э":
		g = GEN_MIDDLE
	case "и":
		g = GEN_MIDDLE
	case "ь":
		if phonetic.IsVoiceless(e2) {
			g = GEN_FEMALE
		}
	}

	return &Noun{gen: g, word: input}
}

func (a *Noun) Gen() string {
	return a.gen
}

func (a *Noun) Word() string {
	return a.word
}

func (a *Noun) Locative(fg string) (string, bool) {
	e1 := a.word[len(a.word)-2:]
	e2 := a.word[len(a.word)-4 : len(a.word)-2]

	if fg != "" && fg != a.gen && fg != GEN_MALE {
		return a.word, false
	}

	if a.gen == GEN_MIDDLE {
		switch e1 {
		case "о":
			if e2 == "н" || phonetic.IsVowel(e2) {
				return a.word, false
			}
			return regexp.MustCompile("о$").ReplaceAllString(a.word, "е"), true
		}

		return a.word, false
	} else if a.gen == GEN_FEMALE {
		if e1 == "я" {
			if e2 == "е" {
				return regexp.MustCompile(".$").ReplaceAllString(a.word, "е"), true
			} else {
				return regexp.MustCompile(".$").ReplaceAllString(a.word, "и"), true
			}
		}

		if e1 == "ь" {
			return regexp.MustCompile(".$").ReplaceAllString(a.word, "и"), true
		}

		if phonetic.IsVowel(e2) {
			return a.word, false
		}

		return regexp.MustCompile(".$").ReplaceAllString(a.word, "е"), true
	} else {
		if e1 == "ь" {
			return regexp.MustCompile(".$").ReplaceAllString(a.word, "е"), true
		}

		return a.word + "е", true
	}
}
