package story

import (
	"fmt"
	"math/rand"

	"github.com/xavier2910/tundra"
	"github.com/xavier2910/tundra/commands"
)

var GameData *tundra.World

func MustInitGameData() {

	start := &tundra.Location{
		Title:       "Near Cave",
		Description: "You are looking into the mouth of a dark cave in the side of a tall, icy cliff to the immediate east. Surrounding you is a tundra so open you can even see your hand in front of your face.",
		Objects:     map[string]*tundra.Object{},
		Commands:    map[string]tundra.Command{},
	}

	penny := tundra.NewObject(tundra.WithDescription("The penny is very shiny indeed, but it has no face."))
	chair := tundra.NewObject(tundra.WithDescription("This chair appears to be made out of gold."))

	incave := &tundra.Location{
		Title:       "In Cave",
		Description: "You are standing in a very dark cave. You can barely make out a sloping upward passage to the west. In the middle of the room is a round hole with a ladder in it.",
		Objects: map[string]*tundra.Object{
			"penny": penny,
			"chair": chair,
		},
		Commands: map[string]tundra.Command{},
	}

	GameData = tundra.NewWorld(
		tundra.NewPlayer(
			tundra.WithStartingLocation(start),
			tundra.WithAdditionalContext(map[string]*tundra.Object{}),
			tundra.WithStartingInventory(map[string]*tundra.Object{}),
		),
		[]*tundra.Location{
			start,
			incave,
		},
	)

	penny.AddCommand("examine", commands.Examine(penny))
	penny.AddCommand("flip", func(o []*tundra.Object) (tundra.CommandResults, error) {
		flip := rand.Intn(2) != 0
		var msg string
		if flip {
			msg = "heads"
		} else {
			msg = "tails"
		}
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg: []string{
				fmt.Sprintf("The only way you can tell which way the coin lands by it's tails side. The coin is %s.", msg),
			},
		}, nil

	})

	penny.AddCommand("take", take("penny", penny))
	penny.AddCommand("drop", drop("penny", penny))
	chair.AddCommand("examine", commands.Examine(chair))
	chair.AddCommand("take", take("chair", chair))
	chair.AddCommand("drop", drop("chair", chair))
}

func MustCreateCommands(cp tundra.CommandProcessor) {
	GameData.PlayerData.AddCommand("inventory", func(o []*tundra.Object) (tundra.CommandResults, error) {
		if len(GameData.PlayerData.Inventory) == 0 {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg:    []string{"You aren't holding anything."},
			}, nil
		}

		var list string
		for name, obj := range GameData.PlayerData.Inventory {
			list = fmt.Sprintf("%s\n%s", list, name)
			cp.InjectContext(name, obj)
		}
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{fmt.Sprintf("You are holding:%s", list)},
		}, nil
	})
}

func MustConnectLocations(cp tundra.CommandProcessor) {
	GameData.Places[0].SetConnection(tundra.East, GameData.Places[1], GameData.PlayerData, cp)
	GameData.Places[1].SetConnection(tundra.West, GameData.Places[0], GameData.PlayerData, cp)
}

func take(name string, obj *tundra.Object) tundra.Command {
	return func(o []*tundra.Object) (tundra.CommandResults, error) {
		if GameData.PlayerData.Inventory[name] != nil {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg: []string{
					fmt.Sprintf("You already have the %s.", name),
				},
			}, nil
		}
		if GameData.PlayerData.CurLoc.GetObject(name) == nil {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg: []string{
					fmt.Sprintf("You already have the %s, or the %s has ceased to exist.", name, name),
				},
			}, fmt.Errorf("object \"%s\": %#v is nil at location %#v", name, obj, GameData.PlayerData.CurLoc)
		}
		GameData.PlayerData.CurLoc.RemoveObject(name)
		GameData.PlayerData.AddObject(name, obj)
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{fmt.Sprintf("%s taken.", name)},
		}, nil
	}
}

func drop(name string, obj *tundra.Object) tundra.Command {
	return func(o []*tundra.Object) (tundra.CommandResults, error) {
		if GameData.PlayerData.Inventory[name] == nil {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg: []string{
					fmt.Sprintf("You don't have a %s.", name),
				},
			}, nil
		}
		if GameData.PlayerData.CurLoc.GetObject(name) != nil {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg: []string{
					fmt.Sprintf("There is already a %s here.", name),
				},
			}, fmt.Errorf("duplicate object \"%s\" at location %#v", name, GameData.PlayerData.CurLoc)
		}
		GameData.PlayerData.CurLoc.AddObject(name, obj)
		GameData.PlayerData.RemoveObject(name)
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{fmt.Sprintf("%s dropped.", name)},
		}, nil
	}
}
