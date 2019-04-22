package domain

import "testing"


func TestFeatureKey(t *testing.T){
	fs := FeatureSet{}
	fs.SetFeature(Polite).SetFeature(Affirmative)

	expKey := "110"
	if fs.Key() != expKey {
		t.Errorf("Expected key %v but got %v", expKey, fs.Key())
	}

	fs = FeatureSet{}
	fs.SetFeature(Past)
	expKey = "1"
	if fs.Key() != expKey {
		t.Errorf("Expected key %v but got %v", expKey, fs.Key())
	}

	fs = FeatureSet{}
	fs.SetFeature(Polite)
	expKey = "10"
	if fs.Key() != expKey {
		t.Errorf("Expected key %v but got %v", expKey, fs.Key())
	}

}


func TestFeatureSet(t *testing.T){
	f := FeatureSet{}

	if f.HasFeature(Past) {
		t.Errorf("FeatureSet has feature Past, but should not have it")
	}
	if f.HasFeature(Polite) {
		t.Errorf("FeatureSet has feature Polite, but should not have it")
	}

	f.SetFeature(Past)
	if !f.HasFeature(Past) {
		t.Errorf("FeatureSet does not have feature Past, but should have it")
	}
	if f.HasFeature(Polite) {
		t.Errorf("FeatureSet has feature Polite, but should not have it")
	}

	f.SetFeature(Polite)
	if !f.HasFeature(Past) {
		t.Errorf("FeatureSet does not have feature Past, but should have it")
	}
	if !f.HasFeature(Polite) {
		t.Errorf("FeatureSet does not have feature Polite, but should have it")
	}

	f.UnsetFeature(Polite).UnsetFeature(Past)
	if f.HasFeature(Past) {
		t.Errorf("FeatureSet has feature Past, but should not have it")
	}
	if f.HasFeature(Polite) {
		t.Errorf("FeatureSet has feature Polite, but should not have it")
	}
}