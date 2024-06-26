package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Arrow struct {
	Begin Vec2
	End   Vec2
}

func (a Arrow) Render() {
	var color = rl.NewColor(0, 0, 255, 130)
	if a.Begin == a.End {
		var origo = GetBoardOrigo()
		var beginPixCent = origo.Add(a.Begin.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
		rl.DrawCircleLines(int32(beginPixCent.X), int32(beginPixCent.Y), TileSize/2, color)
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

func NewArrowDrawer() ArrowDrawer {
	return ArrowDrawer{
		Current: Arrow{
			Begin: Vec2{1, 2},
			End:   Vec2{4, 3},
		},
	}
}

func (ad *ArrowDrawer) Update() {
	if !rl.IsKeyDown(rl.KeyLeftControl) && !rl.IsKeyDown(rl.KeyRightControl) {
		ad.HasCurrent = false
		return
	}
	var coord = GetHoveredCoord()
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		ad.HasCurrent = true
		ad.Current = Arrow{
			Begin: coord,
			End:   coord,
		}
	} else if rl.IsMouseButtonDown(rl.MouseButtonRight) && ad.HasCurrent {
		ad.Current.End = coord
	} else if rl.IsMouseButtonReleased(rl.MouseButtonRight) {
		ad.HasCurrent = false
		ad.Arrows = append(ad.Arrows, ad.Current)
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
