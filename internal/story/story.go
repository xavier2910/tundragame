package story

import (
	"fmt"
	"math/rand"

	"github.com/xavier2910/tundra"
	"github.com/xavier2910/tundra/commands"
)

var GameData *tundra.World
var (
	lampOn = false
)

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
	nearforest := &tundra.Location{
		Title:       "North of Forest",
		Description: "You are standing at the north edge of a dark conifer forest, which bends northward to the west and southward to the east.",
		Objects:     map[string]*tundra.Object{},
		Commands:    map[string]tundra.Command{},
	}
	overcave := &tundra.Location{
		Title:       "Top of Cliff",
		Description: "You are standing atop a tall icy cliff facing west. A conifer forest lies to your east, stretching away to the north and south.",
		Objects:     map[string]*tundra.Object{},
		Commands:    map[string]tundra.Command{},
	}
	westnearforest := &tundra.Location{
		Title:       "East of Forest",
		Description: "You are standing under the eaves of the west edge of a conifer forest. The forest edge continues off to the north and to the southeast. To the east you can see a cliff.",
		Objects:     map[string]*tundra.Object{},
		Commands:    map[string]tundra.Command{},
	}
	house := &tundra.Location{
		Title:       "Forest House",
		Description: "You are standing at the southern edge of a forest, which stretches away to the east and west. There is a small empty-looking house nestled under the eaves of the dark conifer forest.",
		Objects:     map[string]*tundra.Object{},
		Commands:    map[string]tundra.Command{},
	}
	fireplace := tundra.NewObject(tundra.WithDescription("The fireplace is bare and black."))
	lamp := tundra.NewObject(tundra.WithDescription("This lamp appears to still have plenty of oil in it."))
	inhouse := &tundra.Location{
		Title:       "Inside House",
		Description: "You are standing inside a small house.",
		Objects: map[string]*tundra.Object{
			"fireplace": fireplace,
			"lamp":      lamp,
		},
		Commands: map[string]tundra.Command{},
	}
	inforest := &tundra.Location{
		Title:       "In Forest",
		Description: "The forest surrounding you is very dark",
		Objects:     map[string]*tundra.Object{},
		Commands:    map[string]tundra.Command{},
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
			nearforest,
			overcave,
			westnearforest,
			house,
			inhouse,
			inforest,
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
	fireplace.AddCommand("examine", commands.Examine(fireplace))
	lamp.AddCommand("light", func(o []*tundra.Object) (tundra.CommandResults, error) {
		if lampOn {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg:    []string{"The lamp is already on."},
			}, nil
		}
		lampOn = true
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{"You turn the lamp on. It has enough oil to last for quite a while."},
		}, nil
	})
	lamp.AddCommand("snuff", func(o []*tundra.Object) (tundra.CommandResults, error) {
		if !lampOn {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg:    []string{"The lamp is already off."},
			}, nil
		}
		lampOn = false
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{"You turn the lamp off. It still has enough oil to last for quite a while."},
		}, nil
	})
	lamp.AddCommand("take", take("lamp", lamp))
	lamp.AddCommand("examine", func(o []*tundra.Object) (tundra.CommandResults, error) {
		var msg string
		if lampOn {
			msg = fmt.Sprintf("%s. The lamp is on.", lamp.Description)
		} else {
			msg = fmt.Sprintf("%s. The lamp is off.", lamp.Description)
		}
		return tundra.CommandResults{
			Result: 0,
			Msg:    []string{msg},
		}, nil
	})
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
	//start to incave
	GameData.Places[0].SetConnection(tundra.East, GameData.Places[1], GameData.PlayerData, cp)
	GameData.Places[0].SetConnection(tundra.In, GameData.Places[1], GameData.PlayerData, cp)
	//incave to start
	GameData.Places[1].SetConnection(tundra.West, GameData.Places[0], GameData.PlayerData, cp)
	GameData.Places[1].SetConnection(tundra.Out, GameData.Places[0], GameData.PlayerData, cp)
	//nearforest to start
	GameData.Places[2].SetConnection(tundra.North, GameData.Places[0], GameData.PlayerData, cp)
	//start to nearforest
	GameData.Places[0].SetConnection(tundra.South, GameData.Places[2], GameData.PlayerData, cp)
	//overcave to nearforest
	GameData.Places[3].SetConnection(tundra.Southwest, GameData.Places[2], GameData.PlayerData, cp)
	GameData.Places[3].SetConnection(tundra.South, GameData.Places[2], GameData.PlayerData, cp)
	//nearforest to overcave
	GameData.Places[2].SetConnection(tundra.Northeast, GameData.Places[3], GameData.PlayerData, cp)
	GameData.Places[2].SetConnection(tundra.East, GameData.Places[3], GameData.PlayerData, cp)
	//start to westnearforest
	GameData.Places[0].SetConnection(tundra.West, GameData.Places[4], GameData.PlayerData, cp)
	//wesstnearforest to start
	GameData.Places[4].SetConnection(tundra.East, GameData.Places[0], GameData.PlayerData, cp)
	//westnearforest to nearforest
	GameData.Places[4].SetConnection(tundra.Southeast, GameData.Places[2], GameData.PlayerData, cp)
	GameData.Places[4].SetConnection(tundra.South, GameData.Places[2], GameData.PlayerData, cp)
	//start to house
	GameData.Places[0].SetConnection(tundra.North, GameData.Places[5], GameData.PlayerData, cp)
	//house to start
	GameData.Places[5].SetConnection(tundra.South, GameData.Places[0], GameData.PlayerData, cp)
	//house to overcave
	GameData.Places[5].SetConnection(tundra.Southwest, GameData.Places[3], GameData.PlayerData, cp)
	GameData.Places[5].SetConnection(tundra.West, GameData.Places[3], GameData.PlayerData, cp)
	//house to westnearforest
	GameData.Places[5].SetConnection(tundra.Southeast, GameData.Places[4], GameData.PlayerData, cp)
	GameData.Places[5].SetConnection(tundra.East, GameData.Places[4], GameData.PlayerData, cp)
	//westnearforest to house
	GameData.Places[4].SetConnection(tundra.North, GameData.Places[5], GameData.PlayerData, cp)
	GameData.Places[4].SetConnection(tundra.Northeast, GameData.Places[5], GameData.PlayerData, cp)
	//house to inhouse
	GameData.Places[5].SetConnection(tundra.In, GameData.Places[6], GameData.PlayerData, cp)
	//inhouse to house
	GameData.Places[6].SetConnection(tundra.Out, GameData.Places[5], GameData.PlayerData, cp)
	//house to inforest
	GameData.Places[5].SetConnection(tundra.North, GameData.Places[7], GameData.PlayerData, cp)
	GameData.Places[5].SetConnection(tundra.Northwest, GameData.Places[7], GameData.PlayerData, cp)
	GameData.Places[5].SetConnection(tundra.Northeast, GameData.Places[7], GameData.PlayerData, cp)
	//nearforest to inforest
	GameData.Places[2].SetConnection(tundra.South, GameData.Places[7], GameData.PlayerData, cp)
	GameData.Places[2].SetConnection(tundra.Southwest, GameData.Places[7], GameData.PlayerData, cp)
	GameData.Places[2].SetConnection(tundra.Southeast, GameData.Places[7], GameData.PlayerData, cp)
	//westnearforest to inforest
	GameData.Places[4].SetConnection(tundra.West, GameData.Places[7], GameData.PlayerData, cp)
	GameData.Places[4].SetConnection(tundra.Southwest, GameData.Places[7], GameData.PlayerData, cp)
	GameData.Places[4].SetConnection(tundra.Northwest, GameData.Places[7], GameData.PlayerData, cp)
	//overcave to inforest
	GameData.Places[3].SetConnection(tundra.East, GameData.Places[7], GameData.PlayerData, cp)
	GameData.Places[3].SetConnection(tundra.Northeast, GameData.Places[7], GameData.PlayerData, cp)
	GameData.Places[3].SetConnection(tundra.Southeast, GameData.Places[7], GameData.PlayerData, cp)

	// now we allow the player to go various places from in the forest if they have the lamp and it's on
	// but if they don't have the lamp or it's off, they get eaten by a wild beast.
	hookUpinforestTo := func(direction tundra.Direction, other *tundra.Location) {
		conn := tundra.NewObject()
		conn.AddCommand("go", func(o []*tundra.Object) (tundra.CommandResults, error) {
			if lampOn && GameData.PlayerData.Inventory["lamp"] != nil {
				GameData.PlayerData.CurLoc = other
				cp.UpdateContext()
				return tundra.CommandResults{
					Result: tundra.Ok,
					Msg:    []string{"# " + other.Title + "\n\n" + other.Description},
				}, nil
			} else {
				return tundra.CommandResults{
					Result: tundra.Death,
					Msg:    []string{"You blunder around in the dark woods and are eaten by a wild beast."},
				}, nil
			}
		})
		GameData.Places[7]. /*inforest*/ Objects[string(direction)] = conn
	}
	//to nearforest
	hookUpinforestTo(tundra.North, GameData.Places[2])
	//to westnearforest
	hookUpinforestTo(tundra.East, GameData.Places[4])
	//to overcave
	hookUpinforestTo(tundra.West, GameData.Places[3])
	//to house
	hookUpinforestTo(tundra.South, GameData.Places[5])
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
