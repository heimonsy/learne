package model

import (
	"gorm.io/gorm"
)

// Word
type Word struct {
	gorm.Model
	Word        string `gorm:"index:idx_word"`
	Pos         string
	USPronounce string

	Explains []Explain
}
