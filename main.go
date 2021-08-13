package main

import (
	"fmt"		// for Printing
	"os"		// for Commands
	"os/exec"	// for Command Execution
	"errors"	// for Errors
	"time"		// for Pausing Execution
)

type Game struct {
	board	[9]	string
	player		string
	turnNumber	int
	err			error
}

var winningCombinations = [8][3] int {
    {1, 2, 3}, {4, 5, 6}, {7, 8, 9},
    {1, 4, 7}, {2, 5, 8}, {3, 6, 9},
    {1, 5, 9}, {3, 5, 7},
}

func clearScreen() {
	command := exec.Command("cmd", "/c", "cls")
	command.Stdout = os.Stdout
	command.Run()

	fmt.Printf("\n      [Tic Tac Toe]\n\n")
}

func (game *Game) handleError() {
	fmt.Printf("Error: %s", game.err)
	time.Sleep(2000 * time.Millisecond)
	game.err = nil
}

func (game *Game) getFirstPlayer() string {
	var firstPlayer string

	clearScreen()
	fmt.Printf("Choose your marker (X or O): ")
	fmt.Scan(&firstPlayer)

	if(firstPlayer != "X" && firstPlayer != "O") {
		firstPlayer = game.getFirstPlayer()
	}

	return firstPlayer
}

func (game *Game) printBoard() {
	clearScreen()
	for index, indexValue := range game.board {
		if indexValue == "" { fmt.Printf(" \t") } else { fmt.Printf("   %s\t", indexValue) }

		if (index+1) != 9 {
			if (index+1)%3 == 0 {
				fmt.Printf("\n  -  -  -  -  -  -  -  -\n\n")
			} else {
				fmt.Printf("|")
			}
		} else {
			fmt.Printf("\n")
		}
	}
}

func (game *Game) getBoxIndex() int {
	var moveInt int

	fmt.Printf("\nTurn for Player %s to play\n", game.player)
	fmt.Printf("Enter box index to mark: ")
	fmt.Scan(&moveInt) // Player input

	if(moveInt > 9 || moveInt < 1) {
		game.err = errors.New("choose an unmarked index between 1 to 9")
	}

	return moveInt
}

func (game *Game) updateGameStatus(index int) {
	if game.board[index - 1] == "" {
		game.board[index - 1] = game.player // Update board
		game.turnNumber += 1				// Update turn to be played
	} else {
		game.err = errors.New("looks like this index is already marked")
	}
}

func (game *Game) checkForWinner() (bool, string) {
	var board = game.board
	var turnNumber = game.turnNumber

	if turnNumber > 4 {
		for row := 0; row < len(winningCombinations); row++ {
			if	board[winningCombinations[row][0] - 1] == board[winningCombinations[row][1] - 1] &&
				board[winningCombinations[row][1] - 1] == board[winningCombinations[row][2] - 1] &&
				board[winningCombinations[row][2] - 1] == board[winningCombinations[row][0] - 1] {
					if	board[winningCombinations[row][0] - 1] == game.player {
						return true, board[winningCombinations[row][0] - 1]
					}
			}
		}
	}

	if turnNumber == 9 { return true, "" } // Game tied

	// Game ongoing (switch player & return status)
	if game.player == "O" { game.player = "X" } else { game.player = "O" }
	return false, ""
}

func askForRematch() string {
	var response string

	fmt.Printf("\nAnother game? (y/n): ")
	fmt.Scan(&response)

	if(response != "y" && response != "n") {
		clearScreen()
		response = askForRematch()
	}

	return response
}


func main() {
	var game Game

	game.player = game.getFirstPlayer()
	gameOver := false

	var winner string
	var response string

	for !gameOver {
		game.printBoard()

		move := game.getBoxIndex()
		if game.err != nil {
			game.handleError()
		} else {
			game.updateGameStatus(move)

			if game.err != nil {
				game.handleError()
			} else {
				gameOver, winner = game.checkForWinner()
			}
		}
	}

	game.printBoard()
	if winner == "" {
		fmt.Println("\n*It's a tie!")
	} else {
		fmt.Printf("\n*Player %s won!!!", winner)
	}

	response = askForRematch()
	if(response == "y") { main() }
}