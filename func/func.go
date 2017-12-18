package main

import (
	"fmt"
	"github.com/gomeetup/func/game"
	"github.com/gomeetup/func/calc"
	"github.com/gomeetup/func/closure"
	"github.com/gomeetup/func/nofunc"
)

func main() {
	fmt.Println("FunctionDemo")
	fmt.Println()

	calc.CalcDemo()
	fmt.Println()

	closure.ClosureDemo()
	fmt.Println()

	nofunc.NoFuncDemo()
	fmt.Println()

	game.GameDemo()
}

