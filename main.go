package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xavier2910/tundra"
	"github.com/xavier2910/tundra/commandprocessors"
	"github.com/xavier2910/tundragame/internal/logger"
	"github.com/xavier2910/tundragame/internal/story"
)

func main() {
	fmt.Printf("The Tundra, take 2, version 0.0.0.\n\"bye\" exits\n")
	p := &player{
		input: *bufio.NewReader(os.Stdin),
	}
	err := p.play()
	if err != nil {
		logger.Log(fmt.Errorf("FATAL %s", err))
		log.Fatal(err)
	}
}

type player struct {
	input bufio.Reader
}

func (p *player) play() error {

	defer logger.Close()

	story.MustInitGameData()
	cp := commandprocessors.NewTurnBased(story.GameData)
	story.MustCreateCommands(cp)
	story.MustConnectLocations(cp)
	cp.UpdateContext()

	fmt.Printf("\n# %s\n\n%s\n", story.GameData.PlayerData.CurLoc.Title, story.GameData.PlayerData.CurLoc.Description)

	gameOver := false
	for !gameOver {

		command, err := p.getInput()
		if err != nil {
			return err
		}
		if command == "\n" {
			continue
		}
		if command == "bye\n" {
			return nil
		}
		result, err := cp.Execute(command)

		gameOver = p.display(result)
		if err != nil {
			logger.Log(err)
		}

	}

	return nil
}

func (p *player) getInput() (command string, err error) {
	fmt.Print("> ")

	command, err = p.input.ReadString('\n')
	if err != nil {
		return "", err
	}
	return
}

func (p *player) display(results tundra.CommandResults) (gameEnded bool) {
	switch results.Result {
	case tundra.Ok:
		fmt.Printf("\n%s\n", results.Msg[0])

	case tundra.Expo:
		for _, msg := range results.Msg {
			fmt.Printf("\n%s\n\npress enter to continue . . . ", msg)
			p.pause()
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

func (p *player) pause() {
	p.input.ReadRune()
}
