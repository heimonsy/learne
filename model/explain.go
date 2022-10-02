package model

import "gorm.io/gorm"

// Explain
type Explain struct {
	gorm.Model
	EN     string
	ZHS    string
	WordID uint

	Sentences []Sentence
}
