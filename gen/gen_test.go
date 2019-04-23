package gen

import (
	"testing"
	dom "github.com/logologics/kunren-jp/domain"

)


func TestGeneratorRegistry(t *testing.T){
	
	fs := &dom.FeatureSet{}

	f, found := genReg[fs.Key()]

	if f == nil  {
		t.Errorf("Plain affirmative generated was not registered")
	}

	if !found  {
		t.Errorf("Plain affirmative generated was not registered")
	}
}


func TestGenTaberuPresentAffirmative(t *testing.T){
	li := &dom.LexicalItem{
		Cat: dom.V,
		Class: dom.ICHIDAN,
		Forms: map[dom.Form]string{
			dom.DICT: "たべる",
			dom.TFORM: "たべて",
			dom.KANJI: "食べる",
		},
	}

	fs := &dom.FeatureSet{}
	
	infl, err := Gen(li, fs)
	if err!= nil {
		t.Errorf("Form generation failed with error: %v", err)
	}

	if infl == nil {
		t.Errorf("No form was generated")
	}

	expForm := "たべる"
	if infl.Form != expForm {
		t.Errorf("Expected %v but got %v", expForm, infl.Form)
	}
}

func TestGenTaberuPresentNegative(t *testing.T){
	li := &dom.LexicalItem{
		Cat: dom.V,
		Class: dom.ICHIDAN,
		Forms: map[dom.Form]string{
			dom.DICT: "たべる",
			dom.TFORM: "たべて",
			dom.KANJI: "食べる",
		},
	}

	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Negative)
	
	infl, err := Gen(li, fs)
	if err!= nil {
		t.Errorf("Form generation failed with error: %v", err)
	}

	if infl == nil {
		t.Errorf("No form was generated")
	}

	expForm := "たべない"
	if infl.Form != expForm {
		t.Errorf("Expected %v but got %v", expForm, infl.Form)
	}
}