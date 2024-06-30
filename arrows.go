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
	Board uint32
}

func (a Arrow) Render() {
	if a.Begin == a.End {
		a.renderAsCircle()
	} else {
		a.renderAsArrow()
	}
}

func (a Arrow) renderAsArrow() {
	const headSize = TileSize / 3
	const endOffset = headSize
	const lineHeadOffset = 13
	var color = ArrowColors[a.Color]
	var origo = GetBoardOrigo()
	var beginPixCent = origo.Add(a.Begin.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
	var endPixCent = origo.Add(a.End.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
	var rldir = rl.Vector2Normalize(endPixCent.Sub(beginPixCent).ToRlVec())
	var angle = float32(math.Atan2(float64(a.End.Y-a.Begin.Y), float64(a.End.X-a.Begin.X)))
	rl.DrawLineEx(beginPixCent.ToRlVec(), rl.Vector2Subtract(endPixCent.ToRlVec(), rl.Vector2Scale(rldir, endOffset+lineHeadOffset)), 10, color)
	rl.DrawPoly(rl.Vector2Subtract(endPixCent.ToRlVec(), rl.Vector2Scale(rldir, endOffset)), 3, headSize, angle*rl.Rad2deg, color)
}

func (a Arrow) renderAsCircle() {
	var color = ArrowColors[a.Color]
	var origo = GetBoardOrigo()
	var coord = origo.Add(a.Begin.Scale(TileSize))
	rl.DrawTexture(assets.texVfxArrowCircle, int32(coord.X), int32(coord.Y), color)
}

type ArrowDrawer struct {
	HasCurrent bool
	Current    Arrow
	Arrows     []Arrow
}

func (ad *ArrowDrawer) Update(board uint32) {
	if !rl.IsKeyDown(rl.KeyLeftControl) && !rl.IsKeyDown(rl.KeyRightControl) {
		ad.HasCurrent = false
		return
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		ad.Clear()
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
			Board: board,
		}
	} else if rl.IsMouseButtonDown(rl.MouseButtonRight) && ad.HasCurrent {
		ad.Current.End = coord
		ad.Current.Color = color
		ad.Current.Board = board
	} else if rl.IsMouseButtonReleased(rl.MouseButtonRight) && ad.HasCurrent {
		ad.HasCurrent = false
		if i := slices.Index(ad.Arrows, ad.Current); i != -1 {
			if ad.Current.Begin == ad.Current.End {
				ad.RemoveStartingAtWithColor(ad.Current.Begin, board, ad.Current.Color)
			} else {
				ad.Arrows[i] = ad.Arrows[len(ad.Arrows)-1]
				ad.Arrows = ad.Arrows[:len(ad.Arrows)-1]
			}
		} else {
			ad.Arrows = append(ad.Arrows, ad.Current)
		}
	}
}

func (ad *ArrowDrawer) RemoveStartingAtWithColor(pos Vec2, board uint32, color uint8) {
	for i := len(ad.Arrows) - 1; i >= 0; i-- {
		if ad.Arrows[i].Begin == pos && ad.Arrows[i].Board == board && ad.Arrows[i].Color == color {
			ad.Arrows[i] = ad.Arrows[len(ad.Arrows)-1]
			ad.Arrows = ad.Arrows[:len(ad.Arrows)-1]
		}
	}
}

func (ad *ArrowDrawer) Clear() {
	ad.HasCurrent = false
	ad.Arrows = []Arrow{}
}

func (ad *ArrowDrawer) Render(board uint32) {
	if ad.HasCurrent && ad.Current.Board == board {
		ad.Current.Render()
	}
	for _, arrow := range ad.Arrows {
		if arrow.Board == board {
			arrow.Render()
		}
	}
}
