package game

import (
	"fmt"
	"math/rand"
)

/*
This dice game demonstrate how with function one can model solution in simpler way.

Game Rule:
- There are two player
- Each player will get turn to throw dice until one cross 100 score
- For each turn a random strategy will be selected, there are two strategy available
- roll strategy gives same points as on dice
- gamble strategy gives double th points provided it's divisible by 3

Below solution is higher-order function way,
you can model this solution with interface way too, please check comments at the end of this file.
*/

// holds score of both player
type game struct {
	player1, player2 int
}

// two strategy - roll and gamble
type strategy func() (score int)

func roll() int {
	outcome := rand.Intn(6) + 1
	return outcome
}

func gamble() int {
	outcome := rand.Intn(6) + 1
	if outcome%3 == 0 {	// change this to 'outcome%2 == 0' and see result
		return (outcome * 2)
	}
	return 0
}

func randomStrategy() strategy {
	if rand.Intn(2) == 0 {
		return roll
	} else {
		return gamble
	}
}

func play() {
	game := game{}
	for game.player1 < 100 && game.player2 < 100 {
		game.player1 += randomStrategy()()
		game.player2 += randomStrategy()()
	}
	fmt.Println(game)
}

func GameDemo() {
	fmt.Println("GameDemo")
	play()
}

/*
One can model solution using interface as below.

type Strategy interface {
	Excecute() int
}
type Roll struct {
}
type Gamble struct {
}
func (r Roll) Execute() int {
	return 0
}
func (g Gamble) Execute() int {
	return 0
}
*/
