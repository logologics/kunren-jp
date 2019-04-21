package domain

type LexicalItem struct {
	Cat Category
	Class Class
	Forms map[Form]string
}