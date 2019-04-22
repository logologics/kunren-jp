package gen

import 	( 
	"fmt"
	dom "github.com/logologics/kunren-jp/domain"
	"github.com/pkg/errors"
)


var genReg map[string]func(*dom.LexicalItem, *dom.FeatureSet)(*dom.InflectedForm, error)

// create register of generation functions
func init(){
	genReg = make(map[string]func(*dom.LexicalItem, *dom.FeatureSet)(*dom.InflectedForm, error))
	
	plain := &dom.FeatureSet{}
	plain.SetFeature(dom.Affirmative)
	genReg[plain.Key()] = genPlainPresentAffirmative

	plainNeg := &dom.FeatureSet{}
	genReg[plainNeg.Key()] = genPlainPresentNegative

	plainPast := &dom.FeatureSet{}
	plainPast.SetFeature(dom.Past).SetFeature(dom.Affirmative)
	genReg[plainPast.Key()] = genPlainPastAffirmative

	plainPastNeg := &dom.FeatureSet{}
	plainPastNeg.SetFeature(dom.Past)
	genReg[plainPastNeg.Key()] = genPlainPasttNegative

}

// Gen generates a inflected form for a lexical item given the passed feature set
func Gen(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	if lex == nil {
		return nil, errors.New(fmt.Sprintf("A lexical item is required"))
	}

	if fs == nil {
		return nil, errors.New(fmt.Sprintf("A feature set is required"))
	}

	f, found := genReg[fs.Key()]
	if !found {
		return nil, errors.New(fmt.Sprintf("No generator exists for feature set %v", fs))
	}

	return f(lex, fs)
} 


func genPlainPresentAffirmative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	form := lex.Forms[dom.DICT]
	infl := &dom.InflectedForm {
		LexicalItem: lex,
		FeatureSet: fs,
		Form: form,
	}
	
	return infl, nil
}

func genPlainPresentNegative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	return nil, nil

}

func genPlainPastAffirmative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	return nil, nil

}

func genPlainPasttNegative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	return nil, nil
}