package gen

import 	( 
	"unicode/utf8"
	dom "github.com/logologics/kunren-jp/domain"
)

func trimFinalRune(s string) string {
	_, l := utf8.DecodeLastRuneInString(s)
	return s[:len(s)-l]
}

func splitFinalRune(s string) (string, string) {
   _, l := utf8.DecodeLastRuneInString(s)
   p := len(s)-l
   return s[:p], s[p:]
}

func infl(form string, lex *dom.LexicalItem, fs *dom.FeatureSet) *dom.InflectedForm{
	return &dom.InflectedForm {
		LexicalItem: lex,
		FeatureSet: fs,
		Form: form,
	}
}

func stem(lex *dom.LexicalItem) (string, error) {
	if (lex.Class == dom.ICHIDAN){
		return trimFinalRune(lex.Forms[dom.DICT]), nil
	} else {
		pref, suff := splitFinalRune(lex.Forms[dom.DICT])
		h, err := HiraganaChangeVowel(suff, i)
		if err != nil {
			return "", err
		}
		return pref + h.glyph, nil
	}
} 
