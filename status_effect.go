package main

import rl "github.com/gen2brain/raylib-go/raylib"

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

type StatusEffectType struct {
	id   uint32
	name string
	tex  rl.Texture2D
}

type StatusEffect struct {
	piece uint32
	typ   uint32
}

func (e *StatusEffect) Render(coord Vec2, index int, total int, scale uint32) {
	const margin = 0.2
	var typ = sandbox.GetStatusEffectType(e.typ)
	var tilePos = GetWorldOrigo().Add(coord.Scale(TileSize))
	var effectRightMost = tilePos.Add(ONEONE.Scale(int(float32(scale) * TileSize * (1 - margin))))
	var step = ZEROZERO
	if total == 1 {
		effectRightMost = tilePos.Add(Vec2{int(scale * TileSize / 2), int(float32(scale) * TileSize * (1 - margin))})
	} else {
		step = LEFT.Scale(int(float32(scale) * (1 - 2*margin) * TileSize / float32(total-1)))
	}
	var effectCenter = effectRightMost.Add(step.Scale(index))
	var texScale = float32(TileSize*0.4) / float32(max(typ.tex.Width, typ.tex.Height))
	var corner = effectCenter.Sub(Vec2{int(texScale * float32(typ.tex.Width) / 2), int(texScale * float32(typ.tex.Height) / 2)})
	rl.DrawTextureEx(typ.tex, corner.ToRlVec(), 0, float32(scale)*texScale, rl.White)
}
