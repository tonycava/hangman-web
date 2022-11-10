package hangman_classic

import (
	"RootHangManWeb/state"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func hangmanRandomWord() string { //Word creation
	var selectedLine int
	rand.Seed(time.Now().UnixNano())
	switch state.Difficulty {
	case "words1.txt":
		selectedLine = rand.Intn(36-1) + 1
	case "words2.txt":
		selectedLine = rand.Intn(22-1) + 1
	case "words3.txt":
		selectedLine = rand.Intn(23-1) + 1
	}

	fileIO, err := os.OpenFile("hangman-classic/word-file/"+state.Difficulty, os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(any(err))
	}
	lines := strings.Split(string(rawBytes), "\n")
	if selectedLine > len(lines) || selectedLine < 0 {
		selectedLine = 0
	}
	return lines[selectedLine]
}

func randomLetters(Word string) []string {
	var InitialLetter []string
	if len(Word) < 5 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < (len(Word)/3)-1; i++ {
		char := string(Word[rand.Intn(len(Word)-1)])
		InitialLetter = append(InitialLetter, char)
	}
	return InitialLetter
}

func revealLetters(revealLetter []string) {
	state.DisplayAtScreen = []string{}
	for _, char := range state.Word {
		if contains(revealLetter, string(char)) {
			state.DisplayAtScreen = append(state.DisplayAtScreen, string(char))
			continue
		}
		state.DisplayAtScreen = append(state.DisplayAtScreen, "_")
	}
}
