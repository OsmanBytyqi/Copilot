package main

// create a game that will pick a random number between 1 and 100
// and ask the user to guess it
//whoever guesses the correct number, the game will tell them they won

//inputs comes from terminal
//outputs goes to terminal


import (
    "fmt"
    "math/rand"
    "os"
)

type Player struct {
    name string
    score int
}


type Game struct {
    player1 Player
    player2 Player
    secretNumber int
}
//create  player1
func (g *Game) createPlayer1() {
    g.player1.name = "Player 1"
    g.player1.score = 0
}


// create  player2
func (g *Game) createPlayer2() {
    g.player2.name = "Player 2"
    g.player2.score = 0
}

//create game
func (g *Game) createGame() {
    g.secretNumber = rand.Intn(100)
}   

//create secretNumber
func (g *Game) createSecretNumber() {
    g.secretNumber = rand.Intn(100)
}


//create  guess
func (g *Game) createGuess() {
    var guess int
    fmt.Println("Guess a number between 1 and 100")
    fmt.Scanln(&guess)
    if guess == g.secretNumber {
        fmt.Println("You guessed the number!")
        g.player1.score++
    } else if guess > g.secretNumber {
        fmt.Println("Your guess is too high")
    } else {
        fmt.Println("Your guess is too low")
    }
}


//create  score
func (g *Game) createScore() {
    fmt.Println("Player 1 score: ", g.player1.score)
    fmt.Println("Player 2 score: ", g.player2.score)
}


//read inputes
func (g *Game) readInput() {
    var input string
    fmt.Scanln(&input)
    if input == "q" {
        os.Exit(0)
    }
}

//create game loop
func (g *Game) createGameLoop() {
    for {
        g.createGuess()
        g.createScore()
        g.readInput()
    }
}

//create main
func main() {
    var g Game
    g.createPlayer1()
    g.createPlayer2()
    g.createGame()
    g.createGameLoop()
    g.createThanksForPlaying()
}
//create thx for playing

func (g *Game) createThanksForPlaying() {
    fmt.Println("Thanks for playing!")
}

