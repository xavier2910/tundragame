package main

import (
	"fmt"

	"github.com/xavier2910/tundra"
)

func main() {
	fmt.Printf("The Tundra, take 2, version 0.0.0.\nNothing implemented yet\n")
	fmt.Printf("%#v\n", tundra.NewWorld(
		tundra.NewPlayer(),
		[]*tundra.Location{},
	))
}
