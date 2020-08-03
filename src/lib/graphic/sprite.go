// this package is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/img"
import "container/list"

//Instance position angle and the center of an instance of a sprite
type Instance struct
{
	destRect sdl.Rect
	angle float64
	center* sdl.Point
}

//Sprite contains the texture a list of instances and a srcRect
type Sprite struct
{
	texture* sdl.Texture
	instances* list.List
	srcRect  sdl.Rect
}

//NewSprite creates a sprite based on a renderer, the image path and a src rectangle
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

//NewInstance adds a instance to the sprite and initializes the width and height of the dest rectangle with the src rectangle
func (sprite *Sprite) NewInstance() *list.Element {
	var instance Instance

	instance.destRect.W = sprite.srcRect.W
	instance.destRect.H = sprite.srcRect.H
	instance.angle = 0

	return sprite.instances.PushBack(instance)
}