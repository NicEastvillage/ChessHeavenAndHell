package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	NameChaosOrb = "Chaos Orb"
	NameCoin     = "Coin"
	NameIce      = "Ice"
	NameFire     = "Fire"
)

type ObstacleType struct {
	id   uint32
	name string
	tex  rl.Texture2D
}

type Obstacle struct {
	coord Vec2
	board uint32
	typ   uint32
}

func (o *Obstacle) Render(index int, total int) {
	const scale = 0.6
	var typ = sandbox.GetObstacleType(o.typ)
	var halfsize = Vec2{x: int(float32(typ.tex.Width) * scale / 2), y: int(float32(typ.tex.Height) * scale / 2)}
	var tileCenter = GetBoardOrigo().Add(o.coord.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
	if total == 1 {
		rl.DrawTextureEx(typ.tex, tileCenter.Sub(halfsize).ToRlVec(), 0, scale, rl.White)
		return
	}
	const offset = 0.165
	var angleBase = math.Pi / 2
	if total%2 == 0 {
		angleBase = -math.Pi * 3 / 4
	}
	var angleStep = 2 * math.Pi / float64(total)
	var angle = angleBase + angleStep*float64(index)
	var offsetVec = Vec2{x: int(math.Cos(angle) * offset * TileSize), y: int(math.Sin(angle) * offset * TileSize)}
	var pos = tileCenter.Add(offsetVec).Sub(halfsize)
	rl.DrawTextureEx(typ.tex, pos.ToRlVec(), 0, scale, rl.White)
}
