package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/veandco/go-sdl2/sdl"

	"./lib/graphic"
)

//Unit Is a single unit in the game
type Unit struct {
	parent          *UnitType
	Animation       list.Element
	AttackAnimation list.Element
}

//UnitType contains the information which all units of one type have in common
type UnitType struct {
	Speed              uint32
	AnimationSRC       uint32
	AttackSpeed        uint32
	AttackAnimationSRC uint32
}

//UnitTypeRaw contains the same data as a unit object in the json file
type UnitTypeRaw struct {
	Speed                  float64
	Image                  string
	AttackSpeed            float64
	AttackAnimation        string
	SrcRectImage           sdl.Rect
	SrcRectAttackAnimation sdl.Rect
}

//LoadUnits creates a UnitType class with the information in resources/units.json under the unitName
func LoadUnits(graphic graphic.Graphic, unitJSON string) []UnitType {

	data, err := ioutil.ReadFile(unitJSON)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	var unitTypesRaw map[string]UnitTypeRaw
	var unitTypes []UnitType

	err = json.Unmarshal(data, &unitTypesRaw)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	for _, i := range unitTypesRaw {
		var unitType UnitType

		unitType.Speed = uint32(i.Speed)
		unitType.AttackSpeed = uint32(i.AttackSpeed)
		unitType.AnimationSRC, err = graphic.AddSprite(i.Image, i.SrcRectImage)
		if err != nil {
			fmt.Println("error:", err)
			return nil
		}
		unitType.AttackAnimationSRC, err = graphic.AddSprite(i.AttackAnimation, i.SrcRectAttackAnimation)
		if err != nil {
			fmt.Println("error:", err)
			return nil
		}
	}

	return unitTypes
}
