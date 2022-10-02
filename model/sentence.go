package model

import "gorm.io/gorm"

// Sentence
type Sentence struct {
	gorm.Model
	Text      string
	ExplainID uint
}
