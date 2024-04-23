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
	Id   uint32       `json:"id"`
	Name string       `json:"name"`
	Tex  rl.Texture2D `json:"tex"`
}

type Obstacle struct {
	Coord Vec2   `json:"coord"`
	Board uint32 `json:"board"`
	Typ   uint32 `json:"type"`
}

func (o *Obstacle) Render(index int, total int) {
	const scale = 0.6
	var typ = sandbox.GetObstacleType(o.Typ)
	var halfsize = Vec2{X: int(float32(typ.Tex.Width) * scale / 2), Y: int(float32(typ.Tex.Height) * scale / 2)}
	var tileCenter = GetBoardOrigo().Add(o.Coord.Scale(TileSize)).Add(ONEONE.Scale(TileSize / 2))
	if total == 1 {
		rl.DrawTextureEx(typ.Tex, tileCenter.Sub(halfsize).ToRlVec(), 0, scale, rl.White)
		return
	}
	const offset = 0.165
	var angleBase = math.Pi / 2
	if total%2 == 0 {
		angleBase = -math.Pi * 3 / 4
	}
	var angleStep = 2 * math.Pi / float64(total)
	var angle = angleBase + angleStep*float64(index)
	var offsetVec = Vec2{X: int(math.Cos(angle) * offset * TileSize), Y: int(math.Sin(angle) * offset * TileSize)}
	var pos = tileCenter.Add(offsetVec).Sub(halfsize)
	rl.DrawTextureEx(typ.Tex, pos.ToRlVec(), 0, scale, rl.White)
}
