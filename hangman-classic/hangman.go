package hangman_classic

import (
	"RootHangManWeb/state"
	"fmt"
	"strings"
)

func Start() {
	state.Try = 10
	state.GameState = "Playing..."
	state.AlreadyTry = []string{}
	state.Word = hangmanRandomWord()
	// Only on Windows but not need on docker apparently
	// state.Word = state.Word[:len(state.Word)-1]
	fmt.Println(state.Word)
	state.RevealLetter = randomLetters(state.Word)
	revealLetters(state.RevealLetter)
}

func Main(guess string) {
	guessCheck(guess)
}

func guessCheck(guess string) {
	if guess == state.Word {
		state.GameState = "WINNER"
		return
	}

	if len(guess) >= 2 && guess != state.Word {
		state.Try -= 2
		if state.Try <= 0 {
			state.GameState = "LOOSER"
			state.Try = 0
			return
		}
		return
	}
	if contains(state.AlreadyTry, guess) {
		return
	}
	state.AlreadyTry = append(state.AlreadyTry, guess)

	if !strings.Contains(state.Word, guess) {
		state.Try--
		if state.Try <= 0 {
			state.GameState = "LOOSER"
			state.Try = 0
			return
		}
	}

	state.RevealLetter = append(state.RevealLetter, guess)

	revealLetters(state.RevealLetter)

	if strings.Join(state.DisplayAtScreen, "") == state.Word {
		state.GameState = "WINNER"
		return
	}
}
