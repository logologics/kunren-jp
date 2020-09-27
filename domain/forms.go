package domain

//TODO DELETE
type FormType int
const (
	PresentIndicativePlainAffirmative FormType = iota
	PresentIndicativePlainNegative
	PresentIndicativePoliteAffirmative
	PresentIndicativePoliteNegative

	PastIndicativePlainAffirmative
	PastIndicativePlainNegative
	PastIndicativePoliteAffirmative
	PastIndicativePoliteNegative
	
	VolitionalPlainAffirmative
	VolitionalPlainNegative
	VolitionalPoliteAffirmative
	VolitionalPoliteNegative

	ImperativePlainAffirmative
	ImperativePlainNegative
	ImperativePoliteAffirmative
	ImperativePoliteNegative

	PastPresumptivePlainAffirmative
	PastPresumptivePlainNegative
	PastPresumptivePoliteAffirmative
	PastPresumptivePoliteMefative
	
	PresentProgressivePlainAffirmative
	PresentProgressivePlainNegative
	PresentProgressivePoliteAffirmative
	PresentProgressivePoliteNegative
	
	PastProgressivePlainAffirmative
	PastProgressivePlainNegative
	PastProgressivePoliteAffirmative
	PastProgressivePoliteNegative

	BaConditionalPlainAffirmative
	BaConditionalPlainNegative
	
	TaraConditionalPlainAffirmative
	TaraConditionalPlainNegative
	TaraConditionalPoliteAffirmative
	TaraConditionalPoliteNegative

	PotentialPlainAffirmative
	PotentialPlainNegative
	PotentialPoliteAffirmative
	PotentialPoliteNegative

	CaustativePlainAffirmative
	CaustativePlainNegative
	CaustativePoliteAffirmative
	CaustativePoliteNegative

)

// TODO DELETE
func FeatureSets() []*FeatureSet {
	fs := make([]*FeatureSet, 42)
	fs[PresentIndicativePlainAffirmative] = MakeFeatureSet([]Feature{ Indicative  })
	fs[PresentIndicativePlainNegative] = MakeFeatureSet([]Feature{ Indicative  })
	fs[PresentIndicativePoliteAffirmative] = MakeFeatureSet([]Feature{ Indicative  })
	fs[PresentIndicativePoliteNegative] = MakeFeatureSet([]Feature{ Indicative  })



	
	
	

	return fs
}

