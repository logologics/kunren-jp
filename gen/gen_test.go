package gen

import (
	"testing"
	dom "github.com/logologics/kunren-jp/domain"

)

func TestGenTaberuPast(t *testing.T){
	li := &dom.LexicalItem{
		Cat: dom.V,
		Class: dom.ICHIDAN,
		Forms: map[dom.Form]string{
			dom.DICT: "たべる",
			dom.TFORM: "たべて",
			dom.KANJI: "食べる",
		},
	}

	fs := &dom.FeatureSet{
		Features: []dom.Feature{
			dom.PLAIN, dom.PAST,
		},
	}
	
	infl, err := gen(li, fs)
	if err!= nil {
		t.Errorf("Form generation failed with error: %v", err)
	}

	expHiragana := "smthg"
	if infl.Hiragana != expHiragana {
		t.Errorf("Expected %v but got %v", expHiragana, infl.Hiragana)
	}
}