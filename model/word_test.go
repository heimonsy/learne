package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWord(t *testing.T) {
	db, callback := testGetTempDB(t)
	// defer callback()
	_ = callback

	err := db.AutoMigrate(&Word{}, &Explain{}, &Sentence{})
	require.NoError(t, err)

	err = db.Create(&Word{Word: "learn", Pos: "verb", Explains: []Explain{
		{
			ZHS: "学习", EN: "to get knowledge or skill in a new subject or activity", Sentences: []Sentence{
				{Text: "They learn Russian at school."},
				{Text: "I've learned a lot about computers since I started work here."},
			},
		},
	}}).Error
	require.NoError(t, err)

	var words []Word
	err = db.Model(&Word{}).Preload("Explains").Preload("Explains.Sentences").Find(&words, "words.word = ?", "learn").Error
	require.NoError(t, err)
	require.Len(t, words, 1)
	require.Equal(t, "learn", words[0].Word)
	require.Len(t, words[0].Explains, 1)
	require.Equal(t, "学习", words[0].Explains[0].ZHS)
	require.Len(t, words[0].Explains[0].Sentences, 2)
	require.Equal(t, "They learn Russian at school.", words[0].Explains[0].Sentences[0].Text)
	require.Equal(t, "I've learned a lot about computers since I started work here.", words[0].Explains[0].Sentences[1].Text)
}
