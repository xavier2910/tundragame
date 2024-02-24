package story

import "github.com/xavier2910/tundra"

var GameData *tundra.World

func MustInitGameData() {
	GameData = tundra.NewWorld(
		tundra.NewPlayer(),
		[]*tundra.Location{},
	)
}
