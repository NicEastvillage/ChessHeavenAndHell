package main

import rl "github.com/gen2brain/raylib-go/raylib"

type StatusEffectType struct {
	id  uint32
	tex rl.Texture2D
}

type StatusEffect struct {
	piece uint32
	typ   uint32
}

func (e *StatusEffect) Render(coord Vec2, index int, total int) {
	const margin = 0.2
	var typ = sandbox.GetStatusEffectType(e.typ)
	var tilePos = GetWorldOrigo().Add(coord.Scale(TileSize))
	var effectRightMost = tilePos.Add(ONEONE.Scale(TileSize * (1 - margin)))
	var step = ZEROZERO
	if total == 1 {
		effectRightMost = tilePos.Add(Vec2{TileSize / 2, TileSize * (1 - margin)})
	} else {
		step = LEFT.Scale((1 - 2*margin) * TileSize / (total - 1))
	}
	var effectCenter = effectRightMost.Add(step.Scale(index))
	var corner = effectCenter.Sub(Vec2{int(typ.tex.Width / 2), int(typ.tex.Height / 2)})
	rl.DrawTexture(typ.tex, int32(corner.x), int32(corner.y), rl.White)
}
