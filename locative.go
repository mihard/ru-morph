package rumorph

import (
	"github.com/mihard/ru-morph/kind"
)

func Locative(input string) (locative string, err error) {

	words, err := MakePhrase(input)

	if err != nil {
		return input, err
	}

	_words := words.GetWordsOfOriginal()

	for l, w := range _words {
		adj := kind.BuildAdjective(w)
		if adj != nil {
			words.AddWord(adj)
			continue
		}

		sp, phraseFinished := kind.BuildIfSpecial(_words[l:])
		if sp != nil {
			words.AddWord(sp)
			if phraseFinished {
				break
			} else {
				continue
			}
		}

		n := kind.BuildNoun(w)

		if n == nil {
			continue
		}

		words.AddWord(n)
	}

	return words.Locative(), nil
}
