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
var lex = []*d.LexicalItem{
	&d.LexicalItem{
		Cat: d.V,
		Class: d.ICHIDAN,
		Forms: map[d.Form]string{
			d.DICT: "たべる",
			d.KANJI: "食べる",
		},
	},
	&d.LexicalItem{
		Cat: d.V,
		Class: d.GODAN,
		Forms: map[d.Form]string{
			d.DICT: "かう",
			d.KANJI: "買う",
		},
	},
}

var indexedLex *Lexicon

func getLex() *Lexicon {
	if indexedLex == nil {
		items := make(map[string]*d.LexicalItem)
		indexedLex = &Lexicon{Items: items}
		for _, l := range lex {
			indexedLex.Items[l.Forms[d.DICT]] = l
		}
	}

	return indexedLex
}


// GetLexicalItem returns a lexical item 
func GetLexicalItem(dictForm string) (*d.LexicalItem, error) {
	l, found := getLex().Items[dictForm]
	if !found {
		e.New(fmt.Sprintf("Dictionary form could not be found: %v", dictForm))
	}

	return l, nil
}

func GetLexicalItemByIndex(i int) (*d.LexicalItem, error) {
	if i >= DictSize() {
		return nil, e.New(fmt.Sprintf("Index is out of bounds: %v", i))
	}

	return lex[i], nil
}


func DictSize() int {
	return len(lex)
}