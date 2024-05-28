package story

import (
	"fmt"
	"math/rand"

	"github.com/xavier2910/tundra"
	"github.com/xavier2910/tundra/commands"
)

var GameData *tundra.World
var (
	//////////////////////////////STATE FOR ICY PLANET//////////////////////////
	lampOn    = false
	hatchOpen = false
)

func MustInitGameData() {

	///////////////////////////////ICY PLANET///////////////////////////////////////////////
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
		Description: "The forest surrounding you is very dark.",
		Objects:     map[string]*tundra.Object{},
		Commands:    map[string]tundra.Command{},
	}
	//cave depths
	ladder := tundra.NewObject(tundra.WithDescription("The ladder is made of steel loops directly embedded in the stone"))
	hatch := tundra.NewObject(tundra.WithDescription("The door is round, with one of those submarine screwy things for a handle."))
	tjoint := &tundra.Location{
		Title:       "Vertical Tunnel",
		Description: "You are hanging on to a metal ladder in a dark, vertical stone tunnel. A faint grey light shows the top of the tunnel, but down is only darkness.",
		Objects: map[string]*tundra.Object{
			"ladder": ladder,
			"door":   hatch,
		},
		Commands: map[string]tundra.Command{},
	}
	button := tundra.NewObject(tundra.WithDescription("The button is very shiny, and bears the sigil: '->'"))
	telroom := &tundra.Location{
		Title:       "Teleporter Room",
		Description: "You are standing at the bottom of a steel ladder in a dark room illuminated by a brilliant circle of glowing blue.",
		Objects: map[string]*tundra.Object{
			"button": button,
		},
		Commands: map[string]tundra.Command{},
	}
	weprack := tundra.NewObject()
	laser := tundra.NewObject(tundra.WithDescription("The laser is in a pistol form. It has all pertaining input devices."))
	cavestor := &tundra.Location{
		Title:       "Storage Room",
		Description: "You are standing in a small storage room. A row of large tanks, sort of like huge propane tanks, lines one wall. On the other is a weapons rack.",
		Objects: map[string]*tundra.Object{
			"rack": weprack,
			"door": hatch,
		},
		Commands: map[string]tundra.Command{},
	}

	////////////////////////////PC-b1///////////////////////
	hangar := /*todo*/ &tundra.Location{}

	GameData = tundra.NewWorld(
		tundra.NewPlayer(
			tundra.WithStartingLocation(start),
			tundra.WithAdditionalContext(map[string]*tundra.Object{}),
			tundra.WithStartingInventory(map[string]*tundra.Object{}),
		),
		[]*tundra.Location{
			start, //0
			incave,
			nearforest,
			overcave,
			westnearforest,
			house, //5
			inhouse,
			inforest,
			tjoint,
			telroom,
			cavestor, //10

			hangar, //dummy
		},
	)

	////////////////////////////// COMMANDS FOR ICY PLANET  ////////////////////////////////////////
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
	penny.AddCommand("take", commands.Take("penny", penny, GameData))
	penny.AddCommand("drop", commands.Drop("penny", penny, GameData))
	chair.AddCommand("examine", commands.Examine(chair))
	chair.AddCommand("take", commands.Take("chair", chair, GameData))
	chair.AddCommand("drop", commands.Drop("chair", chair, GameData))
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
	lamp.AddCommand("take", commands.Take("lamp", lamp, GameData))
	lamp.AddCommand("drop", commands.Drop("lamp", lamp, GameData))
	lamp.AddCommand("examine", func(o []*tundra.Object) (tundra.CommandResults, error) {
		var msg string
		if lampOn {
			msg = fmt.Sprintf("%s The lamp is on.", lamp.Description)
		} else {
			msg = fmt.Sprintf("%s The lamp is off.", lamp.Description)
		}
		return tundra.CommandResults{
			Result: 0,
			Msg:    []string{msg},
		}, nil
	})
	ladder.AddCommand("examine", commands.Examine(ladder))
	hatch.AddCommand("examine", func(o []*tundra.Object) (tundra.CommandResults, error) {
		var msg string
		if hatchOpen {
			msg = fmt.Sprintf("%s The hatch is open.", hatch.Description)
		} else {
			msg = fmt.Sprintf("%s The hatch is closed.", hatch.Description)
		}
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{msg},
		}, nil
	})
	button.AddCommand("examine", commands.Examine(button))
	weprack.AddObject("laser", laser)
	laser.AddCommand("examine", commands.Examine(laser))
	laser.AddCommand("take", func(o []*tundra.Object) (tundra.CommandResults, error) {
		//attempt a standard take. if that fails, try to do a from weprack take
		//no promises for a hot take
		//or for decent puns...
		res, err := commands.Take("laser", laser, GameData)(o)
		if err != nil {
			if weprack.GetObject("laser") == nil {
				return tundra.CommandResults{
					Result: tundra.Ok,
					Msg:    []string{"Oops, it appears the laser is inaccessable. You may have already grabbed it, but the command processor hasn't figured out it's gone...."},
				}, fmt.Errorf("laser on weprack is nil")
			}
			weprack.RemoveObject("laser")
			GameData.PlayerData.AddObject("laser", laser)
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg:    []string{"laser taken."},
			}, nil
		} else {
			return res, err
		}
	})
	laser.AddCommand("drop", commands.Drop("laser", laser, GameData))

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
	GameData.Places[8].GetObject("door").AddCommand("open", func(o []*tundra.Object) (tundra.CommandResults, error) {
		if hatchOpen {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg:    []string{"The hatch is already open. Therefore, you cannot open it. (If you ever do manage to open an already open door, shoot me an email so I can fix the game.)"},
			}, nil
		}
		GameData.Places[8].SetConnection(tundra.In, GameData.Places[10], GameData.PlayerData, cp)
		GameData.Places[10].SetConnection(tundra.Out, GameData.Places[8], GameData.PlayerData, cp)
		cp.UpdateContext()
		hatchOpen = true
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{"You twist the steering-wheel-esque handle and open the heavy, but smoothly greased hatch."},
		}, nil
	})
	GameData.Places[8].GetObject("door").AddCommand("close", func(o []*tundra.Object) (tundra.CommandResults, error) {
		if !hatchOpen {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg:    []string{"The hatch is already closed. Therefore, you cannot close it. (If you ever do manage to close an already closed door, shoot me an email so I can fix the game.)"},
			}, nil
		}
		GameData.Places[8].RemoveConnection(tundra.In)
		GameData.Places[10].RemoveConnection(tundra.Out)
		cp.UpdateContext()
		hatchOpen = false
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{"The hatch swings heavily closed. You screw the handle to secure it."},
		}, nil
	})
	weprack := GameData.Places[10].GetObject("rack")
	weprack.AddCommand("examine", func(o []*tundra.Object) (tundra.CommandResults, error) {
		var msg string
		if weprack.GetObject("laser") != nil {
			msg = fmt.Sprintf("%s There is a solitary laser pistol hanging on the rack.", weprack.Description)
			cp.InjectContext("laser", weprack.GetObject("laser"))
			cp.InjectContext("pistol", weprack.GetObject("laser"))
		} else {
			msg = fmt.Sprintf("%s The rack is empty.", weprack.Description)
		}
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{msg},
		}, nil
	})
	GameData.Places[9].GetObject("button").AddCommand("push", func(o []*tundra.Object) (tundra.CommandResults, error) {
		GameData.PlayerData.CurLoc = GameData.Places[11]
		cp.UpdateContext()
		return tundra.CommandResults{
			Result: tundra.Expo,
			Msg: []string{
				"You press the button, and immediately you are violently sucked forward in a brilliant flash of light.",
				"You hear a slightly fake-sounding voice say:\n\t\"We thank you for bravely volunteering to save our humble planet and, perhaps, indeed, quite likely, all of humanity. A force of an alien faction is rapidly approaching this planet in order to destroy it. If you cannot save us, we will be forced to warp space to hide the earth from the universe. But that would mean the interruption of trade and the starvation of many. Therefore, we are sending you into one of their central facilities to destroy their fleet or stop them by any other way. The fate of humanity is in your hands. Good luck.",
				"You are sucked forward in another blinding flash of light.",
			},
		}, nil
	})
}

func MustConnectLocations(cp tundra.CommandProcessor) {
	/////////////////////////ICY PLANET////////////////////////
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
		// inforest
		GameData.Places[7].Objects[string(direction)] = conn
	}
	//to nearforest
	hookUpinforestTo(tundra.North, GameData.Places[2])
	//to westnearforest
	hookUpinforestTo(tundra.East, GameData.Places[4])
	//to overcave
	hookUpinforestTo(tundra.West, GameData.Places[3])
	//to house
	hookUpinforestTo(tundra.South, GameData.Places[5])
	//incave to tjoint
	GameData.Places[1].SetConnection(tundra.Down, GameData.Places[8], GameData.PlayerData, cp)
	GameData.Places[1].SetConnection(tundra.In, GameData.Places[8], GameData.PlayerData, cp)
	//tjoint to incave
	GameData.Places[8].SetConnection(tundra.Up, GameData.Places[1], GameData.PlayerData, cp)
	//tjoint to telroom
	GameData.Places[8].SetConnection(tundra.Down, GameData.Places[9], GameData.PlayerData, cp)
	//telroom to tjoint
	GameData.Places[9].SetConnection(tundra.Up, GameData.Places[8], GameData.PlayerData, cp)

}
