package main

import (
	"fmt"
	"log"

	"github.com/xavier2910/tundra"
	"github.com/xavier2910/tundra/commandprocessors"
	"github.com/xavier2910/tundragame/internal/story"
)

func main() {
	fmt.Printf("The Tundra, take 2, version 0.0.0.\nNothing implemented yet\n")
	story.MustInitGameData()
	//mustOpenLogFile()
	err := play()
	if err != nil {
		log.Fatal(err)
	}
}

func play() error {

	cp := commandprocessors.NewTurnBased(story.GameData)
	fmt.Printf("\n%s\n%s\n", story.GameData.PlayerData.CurLoc.Title, story.GameData.PlayerData.CurLoc.Description)

	gameOver := false
	for !gameOver {

		command := getInput()
		result, _ := cp.Execute(command)

		gameOver = display(result)

	}

	return nil
}

func getInput() (command string) {
	fmt.Print("> ")

	fmt.Scanln(&command)

	return command
}

func display(results tundra.CommandResults) (gameEnded bool) {
	switch results.Result {
	case tundra.Ok:
		fmt.Printf("\n%s\n", results.Msg[0])

	case tundra.Expo:
		for _, msg := range results.Msg {
			fmt.Printf("\n%s\n\npress enter to continue . . . ", msg)
			getInput()
		}
	case tundra.Death:
		fmt.Printf("\n%s\n\nYou have died. The end.\n", results.Msg[0])
		gameEnded = true

	case tundra.Win:
		fmt.Printf("\n%s\n\nYou Win! The end.\n", results.Msg[0])
		gameEnded = true
	}

	return
}
