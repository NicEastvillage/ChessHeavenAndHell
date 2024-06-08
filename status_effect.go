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
	RenderStyleOverlay StatusEffectRenderStyle = iota
	RenderStyleBottom
	RenderStyleStun
)

type StatusEffectType struct {
	Id    uint32
	Name  string
	Style StatusEffectRenderStyle
	Tex   rl.Texture2D
}

type StatusEffect struct {
	Piece uint32
	Typ   uint32
}

func (t *StatusEffectType) RenderOnTop(coord Vec2, scale float32) {
	var tilePos = GetBoardOrigo().Add(coord.Scale(TileSize))
	var bottomLeft = tilePos.Add(DOWN.Scale(TileSize))
	var texScale = float32(TileSize) / float32(t.Tex.Width)
	var corner = bottomLeft.Add(UP.Scale(int(texScale * scale * float32(t.Tex.Height))))
	rl.DrawTextureEx(t.Tex, corner.ToRlVec(), 0, texScale*scale, rl.White)
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
	var texScale = float32(TileSize*0.4) / float32(max(t.Tex.Width, t.Tex.Height))
	var corner = effectCenter.Sub(Vec2{int(scale * texScale * float32(t.Tex.Width) / 2), int(scale * texScale * float32(t.Tex.Height) / 2)})
	rl.DrawTextureEx(t.Tex, corner.ToRlVec(), 0, scale*texScale, rl.White)
}

func (t *StatusEffectType) RenderAsStun(coord Vec2, index int, scale float32) {
	const startY = 0.5
	const stepY = -0.11
	var tilePos = GetBoardOrigo().Add(coord.Scale(TileSize))
	var bottomCenter = tilePos.Add(Vec2{int(scale * TileSize / 2), int((startY + stepY*float32(index)) * scale * TileSize)})
	var corner = bottomCenter.Sub(Vec2{int(scale * float32(t.Tex.Width) / 2), int(scale * float32(t.Tex.Height))})
	rl.DrawTextureEx(t.Tex, corner.ToRlVec(), 0, scale, rl.White)
}
