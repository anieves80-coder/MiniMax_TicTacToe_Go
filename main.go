package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var posStr = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
var pos = make([]int, 9)
var played = []int{}
var player int = 1
var plays int = 0

var winCombo1 = []int{0, 3, 6, 0, 1, 2, 0, 2}
var winCombo2 = []int{1, 4, 7, 3, 4, 5, 4, 4}
var winCombo3 = []int{2, 5, 8, 6, 7, 8, 8, 6}

func main() {
	computerPlays()
	showTable()
	play()
}

func showTable() {
	setSymb()
	clearTerminal()
	fmt.Printf("  %s | %s | %s   \n", posStr[0], posStr[1], posStr[2])
	fmt.Println(" ----------- ")
	fmt.Printf("  %s | %s | %s   \n", posStr[3], posStr[4], posStr[5])
	fmt.Println(" ----------- ")
	fmt.Printf("  %s | %s | %s   \n", posStr[6], posStr[7], posStr[8])
	fmt.Println("")
}

func play() {
	fmt.Println("What position do you want to play?")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter position(1-9): ")
	text, _ := reader.ReadString('\n')
	conv, err := strconv.Atoi(string(text[0]))
	if err != nil || conv == 0 {
		showTable()
		fmt.Println("ERROR: Must enter a number from 1 - 9!")
		play()
	}
	if contains(conv) {
		fmt.Println("ERROR: That position not available... Try again!")
		play()
	}

	pos[conv-1] = -1
	plays++
	played = append(played, conv)
	if verifyWin(pos, 0) == -1 {
		setWinner("You win!!")
	}

	computerPlays()
	if verifyWin(pos, 0) == 1 {
		setWinner("You lose!!")
	}
	if verifyTied() {
		setWinner("Game Tied!!")
	}
	showTable()
	play()
}

// Goes through the array and sets the table to display to either X or O
func setSymb() {
	for i := 0; i < 9; i++ {
		if pos[i] == 1 {
			posStr[i] = "o"
		} else if pos[i] == -1 {
			posStr[i] = "x"
		}
	}
}

// Checks if there is a winner.
func verifyWin(grid []int, depth int) int {
	for i := 0; i < 8; i++ {
		total := grid[winCombo1[i]] + grid[winCombo2[i]] + grid[winCombo3[i]]
		if total == -3 {
			return -1
		}
		if total == 3 {
			return 1
		}
	}
	if depth > 8 {
		return 0
	}
	return 2
}

// Computer calls the minimax to determine the best play by score and plays it.
func computerPlays() {

	bestScoreMx := -2
	bestMove := -4

	for i := 0; i < 9; i++ {
		if pos[i] == 0 {
			pos[i] = 1
			score := minimax(pos, plays+1, false)
			pos[i] = 0
			if score > bestScoreMx {
				bestScoreMx = score
				bestMove = i
			}
		}
	}
	pos[bestMove] = 1
	played = append(played, bestMove)
	plays++
}

// Uses the minimax algorithm with a recursive function and returns
// a score for every posible move.
func minimax(grid []int, depth int, isMax bool) int {

	res := verifyWin(grid, depth)
	if res != 2 {
		return res
	}

	if isMax {
		bestScore := -212
		for i := 0; i < 9; i++ {
			if grid[i] == 0 {
				grid[i] = 1
				score := minimax(grid, depth+1, false)
				grid[i] = 0
				if score > bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	}

	bestScore := 2
	for i := 0; i < 9; i++ {
		if grid[i] == 0 {
			grid[i] = -1
			score := minimax(grid, depth+1, true)
			grid[i] = 0
			if score < bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

// Checks to see if the game is a tide.
func verifyTied() bool {	
	if plays > 8 {
		return true
	}
	return false
}

// Displays the winner
func setWinner(msg string) {
	showTable()
	fmt.Println(msg)
	os.Exit(0)
}

// Clears the terminal ( command for MAC or LINUX )
func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Verifies if the position picked has not already been chosen
func contains(val int) bool {
	for _, d := range played {
		if d == val {
			return true
		}
	}
	return false
}
