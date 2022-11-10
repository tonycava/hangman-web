package hangman_server

import (
	"RootHangManWeb/hangman-classic"
	"RootHangManWeb/state"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func HangManServer() {
	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)
	server.HandleFunc("/hangman", Hangman)
	server.HandleFunc("/startgame", StartGame)
	server.HandleFunc("/endgame", EndGame)
	server.HandleFunc("/favicon.ico", FavIcon)
	server.HandleFunc("/credit", Credit)
	server.HandleFunc("/difficulty", Difficulty)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/startgame")
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/templates/Home.gohtml"))

	if state.GameState != "Playing..." {
		http.Redirect(w, r, "/endgame", 303)
		return
	}

	type HomePageProps struct {
		Output       int
		OutputSearch string
	}

	_ = tmpl.Execute(w, HomePageProps{Output: state.Try, OutputSearch: strings.Join(state.DisplayAtScreen, "")})
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		guess := r.FormValue("name")
		hangman_classic.Main(guess)
	}
	http.Redirect(w, r, "/", 303)
}

func StartGame(w http.ResponseWriter, r *http.Request) { //run the logic for the start of the game
	tmpl := template.Must(template.ParseFiles("public/templates/StartGame.gohtml"))
	_ = tmpl.Execute(w, nil)
	hangman_classic.Start()
}

func EndGame(w http.ResponseWriter, r *http.Request) { //load the end page with the corresponding variable
	tmpl := template.Must(template.ParseFiles("public/templates/EndGame.gohtml"))

	type EndGameProps struct {
		OutputWin  string
		OutputTry  int
		OutputText string
	}

	_ = tmpl.Execute(w, EndGameProps{OutputWin: state.GameState, OutputTry: state.Try, OutputText: state.Word})
}

func FavIcon(w http.ResponseWriter, r *http.Request) { //get icon for the site
	http.ServeFile(w, r, "favicon/favicon.ico")
}

func Credit(w http.ResponseWriter, r *http.Request) { //load the credit file for the credit page
	tmpl := template.Must(template.ParseFiles("public/templates/Credit.gohtml"))
	_ = tmpl.Execute(w, nil)
}

func Difficulty(w http.ResponseWriter, r *http.Request) { //select the difficulty level ( different word file )
	if r.Method == "POST" {
		state.Difficulty = r.FormValue("difficulty")
	}

	var difficulty string
	switch state.Difficulty {
	case "words1.txt":
		difficulty = "Easy"
	case "words2.txt":
		difficulty = "Medium"
	case "words3.txt":
		difficulty = "Hard"
	default:
		difficulty = "Easy"
	}

	type DifficultyProps struct {
		Difficulty string
	}

	tmpl := template.Must(template.ParseFiles("public/templates/Difficulty.gohtml"))
	_ = tmpl.Execute(w, DifficultyProps{Difficulty: difficulty})
}
