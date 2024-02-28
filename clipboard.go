package main

type Clipboard struct {
	isEmpty bool
	typ     uint32
	color   PieceColor
	scale   uint32
	effects []uint32
}

func NewClipboard() Clipboard {
	return Clipboard{
		isEmpty: true,
	}
}

func (c *Clipboard) StorePiece(typ uint32, color PieceColor, scale uint32, effects []uint32) {
	c.isEmpty = false
	c.typ = typ
	c.color = color
	c.scale = scale
	c.effects = make([]uint32, len(effects))
	copy(c.effects, effects)
}
