package gen

// basic &Hiragana chart

type vowlRow string
const (
	noVowl vowlRow = "none"
	a vowlRow = "a"
	i vowlRow = "i"
	u vowlRow = "u"
	e vowlRow = "e"
	o vowlRow = "o"
)

type consRow string
const (
	noCons consRow = "none"
	k consRow = "k"
	g consRow = "g"
	s consRow = "s"
	z consRow = "z"
	t consRow = "t"
	d consRow = "d"
	n consRow = "n"
	h consRow = "h"
	b consRow = "b"
	p consRow = "p"
	m consRow = "m"
	y consRow = "y"
	r consRow = "r"
	w consRow = "w"
	nn consRow = "n"
)

type diacritic string
const (
	noDia diacritic = "none"
	dakuten diacritic = "dakuten"
	handakuten diacritic = "handakuten"
)

// Hiragana represents a &Hiragana symbol
type Hiragana struct {
	glyph string
	consRow
	vowlRow
	diacritic
}

var hiraganaTable = []*Hiragana {
	// vowel row
	&Hiragana {
		glyph: "あ",
		consRow: noCons,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "い",
		consRow: noCons,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "う",
		consRow: noCons,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "え",
		consRow: noCons,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "お",
		consRow: noCons,
		vowlRow: o,
		diacritic: noDia,
	},

	// k row
	&Hiragana {
		glyph: "か",
		consRow: k,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "き",
		consRow: k,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "く",
		consRow: k,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "け",
		consRow: k,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "こ",
		consRow: k,
		vowlRow: o,
		diacritic: noDia,
	},

	// g row
	&Hiragana {
		glyph: "が",
		consRow: k,
		vowlRow: a,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ぎ",
		consRow: k,
		vowlRow: i,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ぐ",
		consRow: k,
		vowlRow: u,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "げ",
		consRow: k,
		vowlRow: e,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ご",
		consRow: k,
		vowlRow: o,
		diacritic: dakuten,
	},
	
	// s row
	&Hiragana {
		glyph: "さ",
		consRow: s,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "し",
		consRow: s,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "す",
		consRow: s,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "せ",
		consRow: s,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "そ",
		consRow: s,
		vowlRow: o,
		diacritic: noDia,
	},
						
	// z row
	&Hiragana {
		glyph: "ざ",
		consRow: z,
		vowlRow: a,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "じ",
		consRow: z,
		vowlRow: i,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ず",
		consRow: z,
		vowlRow: u,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ぜ",
		consRow: z,
		vowlRow: e,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ぞ",
		consRow: z,
		vowlRow: o,
		diacritic: dakuten,
	},
	
	// t row
	&Hiragana {
		glyph: "た",
		consRow: t,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ち",
		consRow: t,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "つ",
		consRow: t,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "て",
		consRow: t,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "と",
		consRow: t,
		vowlRow: o,
		diacritic: noDia,
	},
					
	// d row
	&Hiragana {
		glyph: "だ",
		consRow: d,
		vowlRow: a,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ぢ",
		consRow: d,
		vowlRow: i,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "づ",
		consRow: d,
		vowlRow: u,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "で",
		consRow: d,
		vowlRow: e,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ど",
		consRow: d,
		vowlRow: o,
		diacritic: dakuten,
	},
				
	// n row
	&Hiragana {
		glyph: "な",
		consRow: n,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "に",
		consRow: n,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ぬ",
		consRow: n,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ね",
		consRow: n,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "の",
		consRow: n,
		vowlRow: o,
		diacritic: noDia,
	},

	// h row
	&Hiragana {
		glyph: "は",
		consRow: h,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ひ",
		consRow: h,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ふ",
		consRow: h,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "へ",
		consRow: h,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ほ",
		consRow: h,
		vowlRow: o,
		diacritic: noDia,
	},

	// b row
	&Hiragana {
		glyph: "ば",
		consRow: b,
		vowlRow: a,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "び",
		consRow: b,
		vowlRow: i,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ぶ",
		consRow: b,
		vowlRow: u,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "べ",
		consRow: b,
		vowlRow: e,
		diacritic: dakuten,
	},
	&Hiragana {
		glyph: "ぼ",
		consRow: b,
		vowlRow: o,
		diacritic: dakuten,
	},

	// p row
	&Hiragana {
		glyph: "ぱ",
		consRow: p,
		vowlRow: a,
		diacritic: handakuten,
	},
	&Hiragana {
		glyph: "ぴ",
		consRow: p,
		vowlRow: i,
		diacritic: handakuten,
	},
	&Hiragana {
		glyph: "ぷ",
		consRow: p,
		vowlRow: u,
		diacritic: handakuten,
	},
	&Hiragana {
		glyph: "ぺ",
		consRow: p,
		vowlRow: e,
		diacritic: handakuten,
	},
	&Hiragana {
		glyph: "ぽ",
		consRow: p,
		vowlRow: o,
		diacritic: handakuten,
	},

	// m row
	&Hiragana {
		glyph: "ま",
		consRow: m,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "み",
		consRow: m,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "む",
		consRow: m,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "め",
		consRow: m,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "も",
		consRow: m,
		vowlRow: o,
		diacritic: noDia,
	},

	// y row
	&Hiragana {
		glyph: "や",
		consRow: y,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ゆ",
		consRow: y,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "よ",
		consRow: y,
		vowlRow: o,
		diacritic: noDia,
	},

	// r row
	&Hiragana {
		glyph: "ら",
		consRow: r,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "り",
		consRow: r,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "る",
		consRow: r,
		vowlRow: u,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "れ",
		consRow: r,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ろ",
		consRow: r,
		vowlRow: o,
		diacritic: noDia,
	},
	
	// w row
	&Hiragana {
		glyph: "わ",
		consRow: w,
		vowlRow: a,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ゐ",
		consRow: w,
		vowlRow: i,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "ゑ",
		consRow: w,
		vowlRow: e,
		diacritic: noDia,
	},
	&Hiragana {
		glyph: "を",
		consRow: w,
		vowlRow: o,
		diacritic: noDia,
	},

	// nn row
	&Hiragana {
		glyph: "わ",
		consRow: nn,
		vowlRow: noVowl,
		diacritic: noDia,
	},	

}









