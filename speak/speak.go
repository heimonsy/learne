package speak

import (
	"fmt"
	"os"
	"path/filepath"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
)

var audioFolder string

func init() {
	audioFolder = filepath.Join(os.TempDir(), "learne_speak")
	err := os.Mkdir(filepath.Join("", ""), 0666)
	if err != nil {
		fmt.Println("make temporary directory for speak error:", err)
		os.Exit(1)
	}
}

// Sentence speaks sentence
func Sentence(sentence string) {
	speech := htgotts.Speech{Folder: audioFolder, Language: voices.English, Handler: &handlers.Native{}}
	speech.Speak(sentence)
}
