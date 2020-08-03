package graphic

import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/img"
import "container/list"


type Instance struct
{
	destRect sdl.Rect
	angle float64
	center* sdl.Point
}

type Sprite struct
{
	texture* sdl.Texture
	instances* list.List
	srcRect  sdl.Rect
}

func NewSprite(renderer *sdl.Renderer, imgPath string, srcRect sdl.Rect) (Sprite, error) {
	var sprite Sprite
	var err error
	sprite.texture, err = img.LoadTexture(renderer, imgPath)

	if err != nil {
		return sprite, err
	}

	sprite.srcRect = srcRect
	sprite.instances = list.New()

	return sprite, err
}

func (sprite *Sprite) NewInstance() *list.Element {
	var instance Instance

	instance.destRect.W = sprite.srcRect.W
	instance.destRect.H = sprite.srcRect.H
	instance.angle = 0

	return sprite.instances.PushBack(instance)
}