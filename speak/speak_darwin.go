package speak

import "os/exec"

// Word speak a word
func Word(word string) error {
	return exec.Command("say", "-v", "Alex", "error").Run()
}
