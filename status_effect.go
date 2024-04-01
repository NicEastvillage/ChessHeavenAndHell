package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	NameExperience  = "Experience"
	NameBloody      = "Bloody"
	NameCurse       = "Curse"
	NameForcedMove  = "Forced Move"
	NamePaid2ndMove = "Paid 2nd Move"
	NamePortalGun   = "Portal Gun"
	NameStonks      = "Stonks"
	NameStun        = "Stun"
	NameWizardHat   = "Wizard Hat"
)

type StatusEffectRenderStyle = uint32

const (
	RenderStyleBottom StatusEffectRenderStyle = iota
	RenderStyleStun
)

type StatusEffectType struct {
	id    uint32
	name  string
	style StatusEffectRenderStyle
	tex   rl.Texture2D
}

type StatusEffect struct {
	piece uint32
	typ   uint32
}

func (t *StatusEffectType) RenderAtBottom(coord Vec2, index int, total int, scale float32) {
	const margin = 0.2
	var tilePos = GetBoardOrigo().Add(coord.Scale(TileSize))
	var effectRightMost = tilePos.Add(ONEONE.Scale(int(scale * TileSize * (1 - margin))))
	var step = ZEROZERO
	if total == 1 {
		effectRightMost = tilePos.Add(Vec2{int(scale * TileSize / 2), int(scale * TileSize * (1 - margin))})
	} else {
		step = LEFT.Scale(int(scale * (1 - 2*margin) * TileSize / float32(total-1)))
	}
	var effectCenter = effectRightMost.Add(step.Scale(index))
	var texScale = float32(TileSize*0.4) / float32(max(t.tex.Width, t.tex.Height))
	var corner = effectCenter.Sub(Vec2{int(scale * texScale * float32(t.tex.Width) / 2), int(scale * texScale * float32(t.tex.Height) / 2)})
	rl.DrawTextureEx(t.tex, corner.ToRlVec(), 0, scale*texScale, rl.White)
}

func (t *StatusEffectType) RenderAbove(coord Vec2, index int, scale float32) {
	const startY = 0.2
	const stepY = -0.1
	var tilePos = GetBoardOrigo().Add(coord.Scale(TileSize))
	var center = tilePos.Add(Vec2{int(scale * TileSize / 2), int((startY + stepY*float32(index)) * scale * TileSize)})
	var corner = center.Sub(Vec2{int(scale * float32(t.tex.Width) / 2), int(scale * float32(t.tex.Height) / 2)})
	rl.DrawTextureEx(t.tex, corner.ToRlVec(), 0, scale, rl.White)
}
