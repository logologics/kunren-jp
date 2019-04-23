package gen

import 	( 
	dom "github.com/logologics/kunren-jp/domain"
)

var taMap = map[string]string{
	"う":"った", "る":"った", "つ":"った",
	"む":"んだ", "ぬ":"んだ", "ぶ":"んだ",
	"く":"いた",
	"ぐ":"いだ",
	"す":"した",
}


var genReg map[string]func(*dom.LexicalItem, *dom.FeatureSet)(*dom.InflectedForm, error)

// create register of generation functions
func init(){
	genReg = make(map[string]func(*dom.LexicalItem, *dom.FeatureSet)(*dom.InflectedForm, error))
	
	plain := &dom.FeatureSet{}
	genReg[plain.Key()] = genPlainPresentAffirmative

	plainNeg := &dom.FeatureSet{}
	plainNeg.SetFeature(dom.Negative)
	genReg[plainNeg.Key()] = genPlainPresentNegative

	plainPast := &dom.FeatureSet{}
	plainPast.SetFeature(dom.Past)
	genReg[plainPast.Key()] = genPlainPastAffirmative

	plainPastNeg := &dom.FeatureSet{}
	plainPastNeg.SetFeature(dom.Past).SetFeature(dom.Negative)
	genReg[plainPastNeg.Key()] = genPlainPasttNegative

	polite := &dom.FeatureSet{}
	polite.SetFeature(dom.Polite)
	genReg[polite.Key()] = genPolitePresentAffirmative

	politeNeg := &dom.FeatureSet{}
	politeNeg.SetFeature(dom.Polite).SetFeature(dom.Negative)
	genReg[politeNeg.Key()] = genPolitePresentNegative

}
