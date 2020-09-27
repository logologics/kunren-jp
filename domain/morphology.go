package domain

import (
	"math/big"
)

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
	IRREG Class = "irregular"
)

type Form string
const (
	DICT Form = "dictinary form"
	KANJI Form = "kanji stem"
)

type Feature int
const (
	// boolean features
	Past Feature = iota
	Polite
	Negative
	// non boolean features
	Indicative
	Volitional
	Imperative
	Presumptive
	Progressive
	BaConditional
	TaraConditional
	Potential
	Causative
)

type FeatureSet struct {
	bitSet big.Int
}

func (fs *FeatureSet) SetFeature(f Feature) *FeatureSet {
	fs.bitSet.SetBit(&fs.bitSet,int(f), 1)
	return fs
}

func (fs *FeatureSet) UnsetFeature(f Feature) *FeatureSet {
	fs.bitSet.SetBit(&fs.bitSet,int(f),  0)
	return fs
}
func (fs *FeatureSet) HasFeature(f Feature) bool {
	bit := fs.bitSet.Bit(int(f))
	return bit == 1
}
func (fs *FeatureSet) Key() string {
	return fs.bitSet.Text(2)
}

func MakeFeatureSet(features []Feature) *FeatureSet{
	fs := &FeatureSet{}
	for _, f := range features {
		fs.SetFeature(f)
	}
	return fs
}


type InflectedForm struct {
	*LexicalItem
	*FeatureSet
	Form string
}
