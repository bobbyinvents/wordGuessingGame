package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/sethvargo/go-diceware/diceware"
)

func main() {
	mainMenu()
}

func getRandomWord() string {
	list, err := diceware.Generate(1)
	if err != nil {
		log.Fatal(err)
	}
	return list[0]
}

func mainMenu() {
	fmt.Println(strings.ToUpper("Word Guessing Game"))
	fmt.Printf("Instructions: \tGuess letters in a hidden word!\n\n")

	prompt := "Commands:\n\t1: Start a new game.\n\t2: Exit.\n"
	fmt.Println(prompt)
	fmt.Print("Enter a command: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input string
	input = scanner.Text()

	commands := map[string]func(){
		"1": gameMode,
		"2": exitGame,
	}
	out, found := commands[input]
	for !found {
		fmt.Printf("\n%s\nERROR: Invalid command.\nEnter a command: ", prompt)
		scanner.Scan()
		input = scanner.Text()
		out, found = commands[input]
	}
	out()
}

func gameMode() {
	randomWord := getRandomWord()
	hidden := make([]rune, len(randomWord))
	for i := range randomWord {
		hidden[i] += rune("_"[0])
	}
	fmt.Printf("%v\n", string(hidden))

	var guessCount int = 10
	for !gameOver(hidden, guessCount) {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Number of guesses left: %d. Guess a letter: ", guessCount)
		scanner.Scan()
		guess := scanner.Text()
		if len(guess) == 0 || len(guess) > 1 || !unicode.IsLetter(rune(guess[0])) {
			fmt.Println("ERROR: Input needs to be a single letter.")
			continue
		}
		input := rune(scanner.Text()[0])
		guessInWord := false
		for i, v := range randomWord {
			if v == input {
				hidden[i] = v
				guessInWord = true
			}
		}

		if !guessInWord {
			guessCount -= 1
		}
		fmt.Printf("%v\n", string(hidden))
	}
	fmt.Println("Game over!")
	if wordGuessed(hidden) {
		fmt.Println("🎉 YOU WIN! 🎉")
	} else {
		fmt.Printf("😿 YOU LOSE! 😿\nThe word was %q\n", randomWord)
	}
	fmt.Println()
	mainMenu()
}

func gameOver(word []rune, guessCount int) bool {
	if guessCount <= 0 {
		return true
	}
	return wordGuessed(word)
}

func wordGuessed(word []rune) bool {
	for _, v := range word {
		if v == rune("_"[0]) {
			return false
		}
	}
	return true
}

func exitGame() {
	fmt.Println("Exiting game...")
	os.Exit(0)
}
