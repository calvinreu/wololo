package main

import (
	"container/list"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"./lib/graphic"
)

//Game struct which contains info about everything in the game
type Game struct {
	Graphics  graphic.Graphic
	UnitTypes []UnitType
	Units     list.List
}

//CreateGame creates a game class and should only be used once
func CreateGame() Game {
	var game Game
	var err error
	game.Graphics, err = graphic.New("wololo", 0, 0, 1920, 1080, sdl.WINDOW_FULLSCREEN, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("error:", err)
	}
	game.UnitTypes = nil

	return game
}
