package speak

import (
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
)

// Sentence speaks word
func Word(word string) {
	speech := htgotts.Speech{Folder: audioFolder, Language: voices.English, Handler: &handlers.Native{}}
	speech.Speak(word)
}
