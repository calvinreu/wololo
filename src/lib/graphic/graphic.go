//Package graphic is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

//Graphic contains the information required to render a window with diffrent Sprites
type Graphic struct {
	Sprites  []Sprite
	Renderer *sdl.Renderer
	window   *sdl.Window
}

//RunOutput will render FPS frames every second until running is false
func (graphic Graphic) RunOutput(FPS uint32, running *bool) {
	var timeStamp, frameTime uint32
	frameTime = 1000 / FPS

	for *running {
		timeStamp = sdl.GetTicks()
		graphic.Render()
		if sdl.GetTicks()-timeStamp < frameTime {
			sdl.Delay(frameTime - (sdl.GetTicks() - timeStamp))
		}
	}
}

//Render renders the information from the graphic object to the screen
func (graphic Graphic) Render() {
	graphic.Renderer.Clear()

	for _, i := range graphic.Sprites {
		for j := i.instances.Front(); j != nil; j = j.Next() {
			if instance, ok := j.Value.(Instance); ok {
				graphic.Renderer.CopyEx(i.texture, &i.srcRect, &instance.destRect, instance.angle, instance.center, sdl.FLIP_HORIZONTAL)
			} else {
				fmt.Println("list of sprite does not contain Instances")
			}
		}
	}
	graphic.Renderer.Present()
}

//New returns a Graphic object with initialized renderer and window note that Sprites have to be added manual
func New(title string, x, y, width, heigh int32, WindowFlags, RendererFlags uint32) (Graphic, error) {
	var graphic Graphic
	var err = sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER)
	if err != nil {
		return graphic, err
	}

	graphic.window, err = sdl.CreateWindow(title, x, y, width, heigh, WindowFlags)
	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO)
		return graphic, err
	}

	graphic.Renderer, err = sdl.CreateRenderer(graphic.window, -1, RendererFlags)

	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO)
		return graphic, err
	}

	return graphic, nil
}

//AddSprite adds another sprite which can be used be creating a instance of it see Sprite.NewInstance
func (graphic *Graphic) AddSprite(imgPath string, srcRect sdl.Rect) (int, error) {
	var err error
	var sprite Sprite
	sprite, err = NewSprite(graphic.Renderer, imgPath, srcRect)
	if err != nil {
		fmt.Print(err)
		return -1, err
	}
	graphic.Sprites = append(graphic.Sprites, sprite)

	return len(graphic.Sprites) - 1, err
}
