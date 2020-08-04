// this package is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import "github.com/veandco/go-sdl2/sdl"
import "fmt"

//Graphic contains the information required to render a window with diffrent Sprites
type Graphic struct
{
	Sprites []Sprite
	renderer* sdl.Renderer
	window* sdl.Window
}

//Render 
func (graphic Graphic) Render()  {
	graphic.renderer.Clear()

	for _, i := range graphic.Sprites {
		for j := i.instances.Front(); j != nil; j = j.Next() {
			if instance, ok := j.Value.(Instance); ok {
				graphic.renderer.CopyEx(i.texture, &i.srcRect, &instance.destRect, instance.angle, instance.center, sdl.FLIP_HORIZONTAL)
			} else {
				fmt.Println("list of sprite does not contain Instances")
			}
		}
	}
	graphic.renderer.Present()
}

//New returns a Graphic object with initialized renderer and window note that Sprites have to be added manual
func New(title string, x, y, width, heigh int32, WindowFlags, RendererFlags uint32) (Graphic, error) {
	var graphic Graphic
	var err = sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return graphic, err
	}

	graphic.window, err = sdl.CreateWindow(title, x, y, width, heigh, WindowFlags)
	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO)
		return graphic, err
	}

	graphic.renderer, err = sdl.CreateRenderer(graphic.window, -1, RendererFlags)

	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO)
		return graphic, err
	}

	return graphic, nil
}

//Add Sprite adds another sprite which can be used be creating a instance of it see Sprite.NewInstance
func (graphic* Graphic) AddSprite(imgPath string, srcRect sdl.Rect) (int, error) {
	var err error
	var sprite Sprite
	sprite, err = NewSprite(graphic.renderer, imgPath, srcRect)
	if err != nil {
		fmt.Print(err)
		return -1, err
	}
	graphic.Sprites = append(graphic.Sprites, sprite)

	return len(graphic.Sprites)-1, err
}