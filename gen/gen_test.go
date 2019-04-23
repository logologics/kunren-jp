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

func checkGen(t *testing.T, li *dom.LexicalItem, fs *dom.FeatureSet, expForm string){
	infl, err := Gen(li, fs)
	if err!= nil {
		t.Errorf("Form generation failed with error: %v", err)
	}

	if infl == nil {
		t.Errorf("No form was generated")
	}

	if infl.Form != expForm {
		t.Errorf("Expected %v but got %v", expForm, infl.Form)
	}
}

var taberu = &dom.LexicalItem{
	Cat: dom.V,
	Class: dom.ICHIDAN,
	Forms: map[dom.Form]string{
		dom.DICT: "たべる",
		dom.KANJI: "食べる",
	},
}

func TestGenTaberuPresentAffirmative(t *testing.T){
	fs := &dom.FeatureSet{}
	checkGen(t, taberu, fs, "たべる")	
}

func TestGenTaberuPresentPoliteAffirmative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Polite)
	checkGen(t, taberu, fs, "たべます")	
}

func TestGenTaberuPresentNegative(t *testing.T){	
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Negative)
	checkGen(t, taberu, fs, "たべない")
}

func TestGenTaberuPresentPoliteNegative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Polite).SetFeature(dom.Negative)
	checkGen(t, taberu, fs, "たべません")	
}

func TestGenTaberuPastAffirmative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Past)
	checkGen(t, taberu, fs, "たべた")	
}

func TestGenTaberuPastNegative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Past).SetFeature(dom.Negative)
	checkGen(t, taberu, fs, "たべなかった")	
}

var kau = &dom.LexicalItem{
	Cat: dom.V,
	Class: dom.GODAN,
	Forms: map[dom.Form]string{
		dom.DICT: "かう",
		dom.KANJI: "買う",
	},
}

func TestGenKauPresentAffirmative(t *testing.T){
	fs := &dom.FeatureSet{}
	checkGen(t, kau, fs, "かう")	
}

func TestGenKauPresentPoliteAffirmative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Polite)
	checkGen(t, kau, fs, "かいます")	
}

func TestKauPresentNegative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Negative)
	checkGen(t, kau, fs, "かわない")
}

func TestGenKauPresentPoliteNegative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Polite).SetFeature(dom.Negative)
	checkGen(t, kau, fs, "かいません")	
}

func TestGenKauPastAffirmative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Past)
	checkGen(t, kau, fs, "かった")	
}

func TestGenKauPastNegative(t *testing.T){
	fs := &dom.FeatureSet{}
	fs.SetFeature(dom.Past).SetFeature(dom.Negative)
	checkGen(t, kau, fs, "かわなかった")	
}
