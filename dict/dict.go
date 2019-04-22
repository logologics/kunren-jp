package dict

import (
	"fmt"
	e "github.com/pkg/errors"
	d "github.com/logologics/kunren-jp/domain"
)

// Lexicon represents dictionary data 
type Lexicon struct {
	Items map[string]*d.LexicalItem
}

// a little test dictionary to get started
var lex = Lexicon{
	Items : map[string]*d.LexicalItem{
		"たべる" : &d.LexicalItem{
			Cat: d.V,
			Class: d.ICHIDAN,
			Forms: map[d.Form]string{
				d.DICT: "たべる",
				d.TFORM: "たべて",
				d.KANJI: "食べる",
			},
		},
		"かう" : &d.LexicalItem{
			Cat: d.V,
			Class: d.ICHIDAN,
			Forms: map[d.Form]string{
				d.DICT: "かう",
				d.TFORM: "かって",
				d.KANJI: "買う",
			},
		},
	},
}

// GetLexicalItem returns a lexical item 
func GetLexicalItem(dictForm string) (*d.LexicalItem, error) {
	l, found := lex.Items[dictForm]
	if !found {
		e.New(fmt.Sprintf("Dictionary form could not be found: %v", dictForm))
	}

	return l, nil
}