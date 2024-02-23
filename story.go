package main

import "github.com/xavier2910/tundra"

var theworld *tundra.World

func mustInitWorld() {
	theworld = tundra.NewWorld(
		tundra.NewPlayer(),
		[]*tundra.Location{},
	)
}
