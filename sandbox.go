package main

import rl "github.com/gen2brain/raylib-go/raylib"

var sandbox = Sandbox{}

type Sandbox struct {
	boards      [3]Board
	tiles       []Tile
	pieces      []Piece
	nextPieceId uint32
}

func NewSandbox() Sandbox {
	var s = Sandbox{}
	setupBoard(&s, 0, BoardStyleHeaven, false)
	setupBoard(&s, 1, BoardStyleEarth, true)
	setupBoard(&s, 2, BoardStyleHell, false)
	return s
}

func setupBoard(sandbox *Sandbox, boardId uint32, style BoardStyle, withPieces bool) {
	var board = sandbox.GetBoard(boardId)
	board.id = boardId
	board.style = style
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			sandbox.NewTile(board.id, Vec2{x, y})
		}
	}
	if !withPieces {
		return
	}
	sandbox.NewPiece(boardId, Vec2{0, 0}, assets.texBlackRook)
	sandbox.NewPiece(boardId, Vec2{1, 0}, assets.texBlackKnight)
	sandbox.NewPiece(boardId, Vec2{2, 0}, assets.texBlackBishop)
	sandbox.NewPiece(boardId, Vec2{3, 0}, assets.texBlackQueen)
	sandbox.NewPiece(boardId, Vec2{4, 0}, assets.texBlackKing)
	sandbox.NewPiece(boardId, Vec2{5, 0}, assets.texBlackBishop)
	sandbox.NewPiece(boardId, Vec2{6, 0}, assets.texBlackKnight)
	sandbox.NewPiece(boardId, Vec2{7, 0}, assets.texBlackRook)
	for x := 0; x < 8; x++ {
		sandbox.NewPiece(boardId, Vec2{x, 1}, assets.texBlackPawn)
	}
	sandbox.NewPiece(boardId, Vec2{0, 7}, assets.texWhiteRook)
	sandbox.NewPiece(boardId, Vec2{1, 7}, assets.texWhiteKnight)
	sandbox.NewPiece(boardId, Vec2{2, 7}, assets.texWhiteBishop)
	sandbox.NewPiece(boardId, Vec2{3, 7}, assets.texWhiteQueen)
	sandbox.NewPiece(boardId, Vec2{4, 7}, assets.texWhiteKing)
	sandbox.NewPiece(boardId, Vec2{5, 7}, assets.texWhiteBishop)
	sandbox.NewPiece(boardId, Vec2{6, 7}, assets.texWhiteKnight)
	sandbox.NewPiece(boardId, Vec2{7, 7}, assets.texWhiteRook)
	for x := 0; x < 8; x++ {
		sandbox.NewPiece(boardId, Vec2{x, 6}, assets.texWhitePawn)
	}
}

func (s *Sandbox) GetBoard(id uint32) *Board {
	return &s.boards[id]
}

func (s *Sandbox) NewTile(board uint32, coord Vec2) *Tile {
	s.tiles = append(s.tiles, Tile{board: board, coord: coord})
	return &s.tiles[len(s.tiles)-1]
}

func (s *Sandbox) GetTile(board uint32, coord Vec2) *Tile {
	for i := 0; i < len(s.tiles); i++ {
		if s.tiles[i].board == board && s.tiles[i].coord == coord {
			return &s.tiles[i]
		}
	}
	return nil
}

func (s *Sandbox) NewPiece(board uint32, coord Vec2, tex rl.Texture2D) *Piece {
	s.nextPieceId++
	s.pieces = append(s.pieces, Piece{
		id:      s.nextPieceId - 1,
		board:   board,
		coord:   coord,
		texture: tex,
	})
	return &s.pieces[len(s.pieces)-1]
}

func (s *Sandbox) GetPiece(id uint32) *Piece {
	for i := 0; i < len(s.pieces); i++ {
		if s.pieces[i].id == id {
			return &s.pieces[i]
		}
	}
	return nil
}

func (s *Sandbox) Render(board uint32) {
	for i := 0; i < len(s.tiles); i++ {
		if s.tiles[i].board == board {
			s.tiles[i].Render(s.boards[board].style)
		}
	}
	for i := 0; i < len(s.pieces); i++ {
		if s.pieces[i].board == board {
			s.pieces[i].Render()
		}
	}
}
