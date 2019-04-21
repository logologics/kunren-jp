package gen

import (
	"testing"
	"github.com/go-test/deep"

)

func TestGlyphLookup(t *testing.T){
	h, err := LookupHiraganaByGlyph("せ")
	if err != nil {
		t.Errorf("Could not find せ: %v", err)
	}

	expH := &Hiragana {
		glyph: "せ",
		consRow: s,
		vowlRow: e,
		diacritic: noDia,
	}

	if diff := deep.Equal(expH, h); diff != nil {
		t.Errorf("Incorrect hiragana was retrieved: %v", diff)
	}

}

func TestConsVowelLookup(t *testing.T){
	h, err := LookupHiraganaByConsAndVowel(d, i)
	if err != nil {
		t.Errorf("Could not find hiragana for %v - %v : %v", d, i, err)
	}

	expH := &Hiragana {
		glyph: "ぢ",
		consRow: d,
		vowlRow: i,
		diacritic: dakuten,
	}

	if diff := deep.Equal(expH, h); diff != nil {
		t.Errorf("Incorrect hiragana was retrieved: %v", diff)
	}

}