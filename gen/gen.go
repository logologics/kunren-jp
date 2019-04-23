package gen

import 	( 
	"fmt"
	dom "github.com/logologics/kunren-jp/domain"
	"github.com/pkg/errors"
)

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
	return infl(form, lex, fs ), nil
}

func genPlainPresentNegative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	
	var form string 

	if (lex.Class == dom.ICHIDAN){
		form = trimFinalRune(lex.Forms[dom.DICT]) + sNai
	} else {
		pref, suff := splitFinalRune(lex.Forms[dom.DICT])
		suffHira, err := LookupHiraganaByGlyph(suff)
		if err != nil {
			return nil, err
		}

		// exception う
		cRow := suffHira.consRow
		if (suff == "う"){
			cRow = w
		}

		h, err  := LookupHiraganaByConsAndVowel(cRow, a)
		form = pref + h.glyph  + sNai
	}

	return infl(form, lex, fs ), nil
}

func genPlainPastAffirmative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	
	var form string 

	if (lex.Class == dom.ICHIDAN){
		form = trimFinalRune(lex.Forms[dom.DICT]) + sTa
	} else {
		pref, suff := splitFinalRune(lex.Forms[dom.DICT])
		ch, found := taMap[suff]
		if !found {
			return nil, errors.New(fmt.Sprintf("Illegal final glyph om dictionary form: %v", suff))
		}
		form = pref + ch
	}

	return infl(form, lex, fs ), nil
}

func genPlainPasttNegative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	nai, err := genPlainPresentNegative(lex, fs)
	if err != nil {
		return nil, err
	}
	
	form := trimFinalRune(nai.Form) + sKatta
	return infl(form, lex, fs ), nil
}

func genPolitePresentAffirmative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	
	stem, err := stem(lex)
	if err != nil {
		return nil, err
	} 
	form := stem + sMasu
	
	return infl(form, lex, fs ), nil
}

func genPolitePresentNegative(lex *dom.LexicalItem, fs *dom.FeatureSet) (*dom.InflectedForm, error){
	
	stem, err := stem(lex)
	if err != nil {
		return nil, err
	} 
	form := stem + sMasen

	return infl(form, lex, fs ), nil
}
