package api

import (
	"github.com/logologics/kunren-jp/dict"
	g "github.com/logologics/kunren-jp/gen"
	d "github.com/logologics/kunren-jp/domain"
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomInflectedForm() (*d.InflectedForm, error) {
	li, _ := RandomDictionaryItem()
	//TODO handle err
	fs := RandomFeatureSet(li)
	return GeneratedForm(li,fs)
}

func RandomDictionaryItem() (*d.LexicalItem, error) {
	
	i := rand.Intn(dict.DictSize())
	return dict.GetLexicalItemByIndex(i)
}

func RandomFeatureSet(li *d.LexicalItem) *d.FeatureSet {
	
	fs := &d.FeatureSet{}
	if (rand.Intn(1) == 1){
		fs.SetFeature(d.Past)
	}
	if (rand.Intn(1) == 1){
		fs.SetFeature(d.Polite)
	}
	if (rand.Intn(1) == 1){
		fs.SetFeature(d.Negative)
	}
	
	offset := int(d.Indicative)
	nonBooleanFeatureRand := rand.Intn(int(d.Causative)-offset)
	nonBooleanFeature := d.Feature(nonBooleanFeatureRand + offset)
	fs.SetFeature(nonBooleanFeature)
	
	return &d.FeatureSet{}
}

func GeneratedForm(li *d.LexicalItem, fs *d.FeatureSet) (*d.InflectedForm, error) {
	return g.Gen(li,fs)
}

