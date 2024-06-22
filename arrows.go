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
	const headSize = TileSize / 3
	const endOffset = headSize
	const lineHeadOffset = 13
	var origo = GetBoardOrigo()
	var beginPixCent = origo.Add(a.Begin.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
	var endPixCent = origo.Add(a.End.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
	var rldir = rl.Vector2Normalize(endPixCent.Sub(beginPixCent).ToRlVec())
	var angle = float32(math.Atan2(float64(a.End.Y-a.Begin.Y), float64(a.End.X-a.Begin.X)))
	rl.DrawLineEx(beginPixCent.ToRlVec(), rl.Vector2Subtract(endPixCent.ToRlVec(), rl.Vector2Scale(rldir, endOffset+lineHeadOffset)), 10, rl.NewColor(0, 0, 255, 130))
	rl.DrawPoly(rl.Vector2Subtract(endPixCent.ToRlVec(), rl.Vector2Scale(rldir, endOffset)), 3, headSize, angle*rl.Rad2deg, rl.NewColor(0, 0, 255, 130))
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

}

func (ad *ArrowDrawer) Render() {
	if ad.HasCurrent || true {
		ad.Current.Render()
	}
	for _, arrow := range ad.Arrows {
		arrow.Render()
	}
}
