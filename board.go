package main

type Board struct {
	tiles  []Tile
	pieces []Piece
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

func NewStandardBoardWithPieces() Board {
	var board = NewBoard(8, 8)
	board.pieces = make([]Piece, 32)
	board.pieces = append(board.pieces, NewPiece(0, 0, assets.texBlackRook))
	board.pieces = append(board.pieces, NewPiece(1, 0, assets.texBlackKnight))
	board.pieces = append(board.pieces, NewPiece(2, 0, assets.texBlackBishop))
	board.pieces = append(board.pieces, NewPiece(3, 0, assets.texBlackQueen))
	board.pieces = append(board.pieces, NewPiece(4, 0, assets.texBlackKing))
	board.pieces = append(board.pieces, NewPiece(5, 0, assets.texBlackBishop))
	board.pieces = append(board.pieces, NewPiece(6, 0, assets.texBlackKnight))
	board.pieces = append(board.pieces, NewPiece(7, 0, assets.texBlackRook))
	for x := 0; x < 8; x++ {
		board.pieces = append(board.pieces, NewPiece(x, 1, assets.texBlackPawn))
	}
	board.pieces = append(board.pieces, NewPiece(0, 7, assets.texWhiteRook))
	board.pieces = append(board.pieces, NewPiece(1, 7, assets.texWhiteKnight))
	board.pieces = append(board.pieces, NewPiece(2, 7, assets.texWhiteBishop))
	board.pieces = append(board.pieces, NewPiece(3, 7, assets.texWhiteQueen))
	board.pieces = append(board.pieces, NewPiece(4, 7, assets.texWhiteKing))
	board.pieces = append(board.pieces, NewPiece(5, 7, assets.texWhiteBishop))
	board.pieces = append(board.pieces, NewPiece(6, 7, assets.texWhiteKnight))
	board.pieces = append(board.pieces, NewPiece(7, 7, assets.texWhiteRook))
	for x := 0; x < 8; x++ {
		board.pieces = append(board.pieces, NewPiece(x, 6, assets.texWhitePawn))
	}
	return board
}

func (b *Board) Render() {
	for _, tile := range b.tiles {
		tile.Render()
	}
	for _, piece := range b.pieces {
		piece.Render()
	}
}
