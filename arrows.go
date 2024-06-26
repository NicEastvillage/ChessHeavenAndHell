package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"slices"
)

var ArrowColors = []rl.Color{
	{20, 80, 255, 130},
	{255, 40, 20, 130},
	{0, 125, 0, 130},
	{255, 210, 0, 130},
}

type Arrow struct {
	Begin Vec2
	End   Vec2
	Color uint8
}

func (a Arrow) Render() {
	var color = ArrowColors[a.Color]
	if a.Begin == a.End {
		var origo = GetBoardOrigo()
		var coord = origo.Add(a.Begin.Scale(TileSize))
		rl.DrawTexture(assets.texVfxArrowCircle, int32(coord.X), int32(coord.Y), color)
		return
	}

	const headSize = TileSize / 3
	const endOffset = headSize
	const lineHeadOffset = 13
	var origo = GetBoardOrigo()
	var beginPixCent = origo.Add(a.Begin.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
	var endPixCent = origo.Add(a.End.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
	var rldir = rl.Vector2Normalize(endPixCent.Sub(beginPixCent).ToRlVec())
	var angle = float32(math.Atan2(float64(a.End.Y-a.Begin.Y), float64(a.End.X-a.Begin.X)))
	rl.DrawLineEx(beginPixCent.ToRlVec(), rl.Vector2Subtract(endPixCent.ToRlVec(), rl.Vector2Scale(rldir, endOffset+lineHeadOffset)), 10, color)
	rl.DrawPoly(rl.Vector2Subtract(endPixCent.ToRlVec(), rl.Vector2Scale(rldir, endOffset)), 3, headSize, angle*rl.Rad2deg, color)
}

type ArrowDrawer struct {
	HasCurrent bool
	Current    Arrow
	Arrows     []Arrow
}

func (ad *ArrowDrawer) Update() {
	if !rl.IsKeyDown(rl.KeyLeftControl) && !rl.IsKeyDown(rl.KeyRightControl) {
		ad.HasCurrent = false
		return
	}
	var coord = GetHoveredCoord()
	var color = uint8(0)
	if rl.IsKeyDown(rl.KeyRightShift) || rl.IsKeyDown(rl.KeyLeftShift) {
		color = 1
	}
	if rl.IsKeyDown(rl.KeyRightAlt) || rl.IsKeyDown(rl.KeyLeftAlt) {
		color = 2
	}
	if (rl.IsKeyDown(rl.KeyRightShift) || rl.IsKeyDown(rl.KeyLeftShift)) && (rl.IsKeyDown(rl.KeyRightAlt) || rl.IsKeyDown(rl.KeyLeftAlt)) {
		color = 3
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		ad.HasCurrent = true
		ad.Current = Arrow{
			Begin: coord,
			End:   coord,
			Color: color,
		}
	} else if rl.IsMouseButtonDown(rl.MouseButtonRight) && ad.HasCurrent {
		ad.Current.End = coord
		ad.Current.Color = color
	} else if rl.IsMouseButtonReleased(rl.MouseButtonRight) {
		ad.HasCurrent = false
		if slices.Index(ad.Arrows, ad.Current) == -1 {
			ad.Arrows = append(ad.Arrows, ad.Current)
		} else if ad.Current.Begin == ad.Current.End {
			for i := len(ad.Arrows) - 1; i >= 0; i-- {
				if ad.Arrows[i].Begin == ad.Current.Begin {
					ad.Arrows[i] = ad.Arrows[len(ad.Arrows)-1]
					ad.Arrows = ad.Arrows[:len(ad.Arrows)-1]
				}
			}
		}
	}
}

func (ad *ArrowDrawer) Render() {
	if ad.HasCurrent {
		ad.Current.Render()
	}
	for _, arrow := range ad.Arrows {
		arrow.Render()
	}
}
