package rumorph

import (
	"errors"
	"fmt"
	"github.com/mihard/ru-morph/kind"
	"regexp"
	"strings"
)

type Phrase struct {
	original  string
	origWords []string
	words     []kind.IWord
	gen       string
}

func MakePhrase(original string) (*Phrase, error) {
	words := regexp.MustCompile("[—\\s-]+").Split(original, 10)

	p := &Phrase{original: original, origWords: words, words: []kind.IWord{}}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Phrase) AddWord(w kind.IWord) {
	p.words = append(p.words, w)
}

func (p *Phrase) GetWordsOfOriginal() []string {
	return p.origWords
}

func (p *Phrase) Locative() string {
	locative := p.original

	fg := ""
	pLen := len(p.words)
	for i, w := range p.words {
		l, changed := w.Locative(fg)

		if changed {
			fg = w.Gen()
		}

		locative = strings.Replace(locative, w.Word(), l, 10)

		if pLen > 1 && i == pLen-1 && w.Gen() == kind.GEN_MIDDLE && !changed {
			return p.locativeForUnchangableMiddleGenNouns()
		}
	}

	return locative
}

func (p *Phrase) locativeForUnchangableMiddleGenNouns() string {
	locative := p.original

	fg := kind.ENFORCED_ORIGINAL
	for _, w := range p.words {
		l, _ := w.Locative(fg)

		locative = strings.Replace(locative, w.Word(), l, 10)
	}

	return locative
}

func (p *Phrase) validate() error {
	vrx := regexp.MustCompile("^[0-9А-ЯЁа-яё]+$")

	for _, w := range p.origWords {
		if !vrx.MatchString(w) {
			return errors.New(fmt.Sprintf("The word '%s' contains one or more not-Cyrillic characters", w))
		}
	}

	return nil
}
