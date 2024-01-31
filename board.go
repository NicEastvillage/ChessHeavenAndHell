package main

type Board struct {
	tiles []Tile
}

func NewBoard(width, height int) Board {
	var tiles = make([]Tile, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tiles[y*width+x] = Tile{coord: Vec2{x: x, y: y}}
		}
	}
	return Board{tiles: tiles}
}

func (b *Board) Render() {
	for _, tile := range b.tiles {
		tile.Render()
	}
}

func (b *Board) BoundingRect() AARect {
	var rect = NewAARectEmpty()
	for _, tile := range b.tiles {
		rect = rect.ExpandedToInclude(tile.coord)
	}
	return rect
}
