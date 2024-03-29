package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	board uint32
	coord Vec2
}

func (t *Tile) Render(style BoardStyle) {
	var color = ColorAt(t.coord, style)
	var pos = GetBoardOrigo().Add(t.coord.Scale(TileSize))
	rl.DrawRectangle(int32(pos.x), int32(pos.y), TileSize, TileSize, color)
}

func ColorAt(coord Vec2, style BoardStyle) rl.Color {
	var light = coord.x%2 == coord.y%2
	var color = rl.NewColor(227, 193, 111, 255)
	if style == BoardStyleEarth && !light {
		color = rl.NewColor(184, 139, 74, 255)
	} else if style == BoardStyleHeaven {
		if light {
			color = rl.NewColor(255, 239, 178, 255)
		} else {
			color = rl.NewColor(234, 210, 28, 255)
		}
	} else if style == BoardStyleHell {
		if light {
			color = rl.NewColor(255, 127, 107, 255)
		} else {
			color = rl.NewColor(193, 96, 77, 255)
		}
	}
	return color
}
