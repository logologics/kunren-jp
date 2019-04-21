package domain


type Language string
const (
	JP Language = "Japanese"
)


type Category string
const (
	V Category = "verb"
	N Category = "noun"
	A Category = "adjective"
	NA Category = "na-adjective" 
)

type Class string
const (
	ICHIDAN Class = "ichidan"
	GODAN Class = "godan"
)

type Feature string
const (
	PLAIN Feature = "plain"
	PAST Feature = "past"
	POLITE Feature = "polite"
	AFFIRMATIVE Feature = "affirmative"
)

type Form string
const (
	DICT Form = "dictinary form"
	TFORM Form = "t-form"
	KANJI Form = "kanji stem"

)

type FeatureSet struct {
	ID string
	Features []Feature
}

type InflectedForm struct {
	*LexicalItem
	*FeatureSet
	Hiragana string
}
