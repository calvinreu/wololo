package main

import (
	"container/list"
	"fmt"

	graphic ".."
	"github.com/veandco/go-sdl2/sdl"
)

var sprite1 int
var sprite2 int

var instances1 [10]list.Element
var instances2 [20]list.Element

func main() {
	example, err := graphic.New("example", 0, 0, 1920, 1080, sdl.WINDOW_FULLSCREEN, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Print(err)
		return
	}

	sprite1, err = example.AddSprite("./img/wololo.png", sdl.Rect{0, .0, 400, 200})
	sprite2, err = example.AddSprite("./img/blabla.png", sdl.Rect{0, .0, 200, 400})

	for i := 0; i < len(instances1); i++ {
		instances1[i].Value = example.Sprites[sprite1].NewInstance((float64)(i*36), &sdl.Point{(int32)(i*192 + 100), (int32)(i*108 + 100)})
	}

	for i := 0; i < len(instances1); i++ {
		instances2[i].Value = example.Sprites[sprite2].NewInstance((float64)(i*36), &sdl.Point{(int32)(i * 192), (int32)(i*108 + 100)})
	}
	running := true
	go example.RunOutput(10, &running)
	sdl.Delay(1000)
	running = false

	fmt.Print("the result should look very chaotic \n")

	sdl.Quit()
}
