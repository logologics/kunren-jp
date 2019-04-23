package gen

import 	(
	"fmt"
	"github.com/pkg/errors"
)

// create some usefull hiragana indexes

// HiraganaLookup allows lookup of Hiragana by glyph
var hiraganaLookup map[string]*Hiragana

// HiraganaLookup allows lookup of Hiragana by consonant and vowl row
var hiraganaConsVowel map[consRow]map[vowlRow]*Hiragana

func init(){
	mapGlyphLookup()
	mapConsVowlLookup()
}

func mapGlyphLookup() {
	hiraganaLookup = make(map[string]*Hiragana, 72)
	for _, h := range hiraganaTable{
		hiraganaLookup[h.glyph] = h
	}
}

func mapConsVowlLookup() {
	hiraganaConsVowel = make (map[consRow]map[vowlRow]*Hiragana)
	for _, h := range hiraganaTable{
		if h.consRow != nn {
			if (hiraganaConsVowel[h.consRow] == nil){
				hiraganaConsVowel[h.consRow] = make(map[vowlRow]*Hiragana)
			}

			hiraganaConsVowel[h.consRow][h.vowlRow]=h
		}
	}
}

func LookupHiraganaByGlyph(glyph string) (*Hiragana, error){
	h, err := hiraganaLookup[glyph]
	if !err {
		return nil, errors.New(fmt.Sprintf("Cannot find glyph %v", glyph))
	} 	
	return h, nil
}

func LookupHiraganaByConsAndVowel(c consRow, v vowlRow) (*Hiragana, error){
	h, err := hiraganaConsVowel[c][v]
	if !err {
		return nil, errors.New(fmt.Sprintf("Cannot find glyph %v:%v", c, v))
	} 	
	return h, nil
}

func HiraganaChangeVowel(glyph string, v vowlRow) (*Hiragana, error) {
	h, err := LookupHiraganaByGlyph(glyph)
	if err != nil {
		return nil, err
	}
	return LookupHiraganaByConsAndVowel(h.consRow, v)
}