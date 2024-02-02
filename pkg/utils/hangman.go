package utils

import (
	"math/rand"
	"unicode"
)

// PickRandomWord get random word for Hangman game
func PickRandomWord() string {
	//TODO add more words
	words := []string{"apple", "banana", "orange", "computer", "programming", "blockchain"}
	return words[rand.Intn(len(words))]
}

// IsLetter check if user guess is a letter
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// HangmanAscii ascii hangman pictures
var HangmanAscii = [7]string{
	"\n+-----+\n" +
		"|/      |\n" +
		"|		  \n" +
		"|		  \n" +
		"|		  \n" +
		"|		  \n" +
		"|		  \n" +
		"+=====+\n",

	"\n+-----+ \n" +
		"|/     | \n" +
		"|	  0\n" +
		"|		  \n" +
		"|		  \n" +
		"|		  \n" +
		"|		  \n" +
		"+=====+\n",

	"\n+-----+ \n" +
		"|/     | \n" +
		"|	  0\n" +
		"|	   | \n" +
		"|	   | \n" +
		"|		  \n" +
		"|		  \n" +
		"+=====+\n",

	"\n+-----+ \n" +
		"|/     | \n" +
		"|	  0\n" +
		"|	   |--\n" +
		"|	   | \n" +
		"|		  \n" +
		"|		  \n" +
		"+=====+\n",

	"\n+-----+ \n" +
		"|/     | \n" +
		"|	  0\n" +
		"|   --|--\n" +
		"|	   | \n" +
		"|		  \n" +
		"|		  \n" +
		"+=====+\n",
	"\n+-----+ \n" +
		"|/     | \n" +
		"|	  0\n" +
		"|   --|--\n" +
		"|	  | \n" +
		"|    /  \n" +
		"|		  \n" +
		"+=====+\n",
	"\n+-----+ \n" +
		"|/     | \n" +
		"|	  0\n" +
		"|   --|--\n" +
		"|	  | \n" +
		"|    / \\\n" +
		"|		  \n" +
		"+=====+\n",
}
