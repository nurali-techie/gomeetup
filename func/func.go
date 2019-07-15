package main

import (
	"fmt"

	"github.com/nurali-techie/gomeetup/func/calc"
	"github.com/nurali-techie/gomeetup/func/closure"
	"github.com/nurali-techie/gomeetup/func/game"
	"github.com/nurali-techie/gomeetup/func/nofunc"
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
