package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"sort"
)

var sandbox = Sandbox{
	Shop: NewShop(),
}

type Sandbox struct {
	Shop          Shop               `json:"shop"`
	Boards        [3]Board           `json:"boards"`
	Tiles         []Tile             `json:"tiles"`
	PieceTypes    []PieceType        `json:"pieceTypes"`
	Pieces        []Piece            `json:"pieces"`
	NextPieceId   uint32             `json:"nextPieceId"`
	EffectTypes   []StatusEffectType `json:"effectTypes"`
	Effects       []StatusEffect     `json:"effects"`
	ObstacleTypes []ObstacleType     `json:"obstacleTypes"`
	Obstacles     []Obstacle         `json:"obstacles"`
}

func (s *Sandbox) GetBoard(id uint32) *Board {
	return &s.Boards[id]
}

func IsOffBoard(coord Vec2) bool {
	return coord.X < 0 || coord.X >= 8 || coord.Y < 0 || coord.Y >= 8
}

func (s *Sandbox) FindUnoccupiedOffBoardCoord() Vec2 {
	for x := 9; x < 12; x++ {
		for y := 0; y < 8; y++ {
			if s.GetPieceAtVisual(Vec2{x, y}, OffBoard) == nil {
				return Vec2{x, y}
			}
		}
	}
	for x := 2; x < 4; x++ {
		for y := 0; y < 8; y++ {
			if s.GetPieceAtVisual(Vec2{-x, y}, OffBoard) == nil {
				return Vec2{-x, y}
			}
		}
	}
	return Vec2{X: 9, Y: -1}
}

func (s *Sandbox) NewTile(board uint32, coord Vec2) *Tile {
	s.Tiles = append(s.Tiles, Tile{Board: board, Coord: coord})
	return &s.Tiles[len(s.Tiles)-1]
}

func (s *Sandbox) GetTile(board uint32, coord Vec2) *Tile {
	for i := 0; i < len(s.Tiles); i++ {
		if s.Tiles[i].Board == board && s.Tiles[i].Coord == coord {
			return &s.Tiles[i]
		}
	}
	return nil
}

func (s *Sandbox) RemoveTile(board uint32, coord Vec2) bool {
	if len(s.Tiles) == 0 {
		return false
	}
	// The slice is unordered, so we insert the last Tile where there removed Tile was and shorten the slice
	var last = s.Tiles[len(s.Tiles)-1]
	for i := 0; i < len(s.Tiles); i++ {
		if s.Tiles[i].Board == board && s.Tiles[i].Coord == coord {
			s.Tiles[i] = last
			s.Tiles = s.Tiles[:len(s.Tiles)-1]
			return true
		}
	}
	return false
}

func (s *Sandbox) RegisterPieceType(name string, texWhite rl.Texture2D, texBlack rl.Texture2D) *PieceType {
	// We assume Piece types are never unregistered
	s.PieceTypes = append(s.PieceTypes, PieceType{
		Id:       uint32(len(s.PieceTypes)),
		Name:     name,
		TexWhite: texWhite,
		TexBlack: texBlack,
	})
	return &s.PieceTypes[len(s.PieceTypes)-1]
}

func (s *Sandbox) GetPieceType(id uint32) *PieceType {
	return &s.PieceTypes[id]
}

func (s *Sandbox) GetPieceTypeByName(name string) *PieceType {
	for i := 0; i < len(s.PieceTypes); i++ {
		if s.PieceTypes[i].Name == name {
			return &s.PieceTypes[i]
		}
	}
	return nil
}

func (s *Sandbox) NewPiece(typ uint32, color PieceColor, board uint32, coord Vec2) *Piece {
	s.NextPieceId++
	s.Pieces = append(s.Pieces, Piece{
		Id:    s.NextPieceId - 1,
		Typ:   typ,
		Color: color,
		Board: board,
		Coord: coord,
		Scale: 1,
	})
	return &s.Pieces[len(s.Pieces)-1]
}

func (s *Sandbox) NewPieceFromName(typ string, color PieceColor, board uint32, coord Vec2) *Piece {
	return s.NewPiece(s.GetPieceTypeByName(typ).Id, color, board, coord)
}

// AddPiece adds a Piece with full details. It is up to the called to ensure that another Piece with
// the same Id does not exist.
func (s *Sandbox) AddPiece(piece Piece) *Piece {
	s.Pieces = append(s.Pieces, piece)
	return &s.Pieces[len(s.Pieces)-1]
}

func (s *Sandbox) GetPiece(id uint32) *Piece {
	for i := 0; i < len(s.Pieces); i++ {
		if s.Pieces[i].Id == id {
			return &s.Pieces[i]
		}
	}
	return nil
}

func (s *Sandbox) RemovePiece(id uint32) bool {
	if len(s.Pieces) == 0 {
		return false
	}

	s.RemoveEffectsFromPiece(id)

	// The slice is unordered, so we insert the last Piece where there removed Piece was and shorten the slice
	for i := 0; i < len(s.Pieces); i++ {
		if s.Pieces[i].Id == id {
			s.Pieces[i] = s.Pieces[len(s.Pieces)-1]
			s.Pieces = s.Pieces[:len(s.Pieces)-1]
			return true
		}
	}
	return false
}

func (s *Sandbox) RemoveEffectsFromPiece(pieceId uint32) {
	var removedEffects = 0
	for i := len(s.Effects) - 1; i >= 0; i-- {
		if s.Effects[i].Piece == pieceId {
			removedEffects++
			s.Effects[i] = s.Effects[len(s.Effects)-removedEffects]
		}
	}
	s.Effects = s.Effects[:len(s.Effects)-removedEffects]
}

func (s *Sandbox) GetPieceAt(coord Vec2, board uint32) *Piece {
	for i := 0; i < len(s.Pieces); i++ {
		if s.Pieces[i].Board == board && s.Pieces[i].Coord == coord {
			return &s.Pieces[i]
		}
	}
	return nil
}

func (s *Sandbox) GetPieceAtVisual(coord Vec2, board uint32) *Piece {
	for i := 0; i < len(s.Pieces); i++ {
		for x := 0; x < int(s.Pieces[i].Scale); x++ {
			for y := 0; y < int(s.Pieces[i].Scale); y++ {
				if (s.Pieces[i].Board == board || s.Pieces[i].Board == OffBoard) && s.Pieces[i].Coord.Add(Vec2{x, y}) == coord {
					return &s.Pieces[i]
				}
			}
		}
	}
	return nil
}

func (s *Sandbox) RegisterEffectType(name string, style StatusEffectRenderStyle, tex rl.Texture2D) *StatusEffectType {
	// We assume effect types are never unregistered
	s.EffectTypes = append(s.EffectTypes, StatusEffectType{
		Id:    uint32(len(s.EffectTypes)),
		Name:  name,
		Style: style,
		Tex:   tex,
	})
	return &s.EffectTypes[len(s.EffectTypes)-1]
}

func (s *Sandbox) GetStatusEffectType(id uint32) *StatusEffectType {
	return &s.EffectTypes[id]
}

func (s *Sandbox) GetStatusEffectTypeByName(name string) *StatusEffectType {
	for i := 0; i < len(s.EffectTypes); i++ {
		if s.EffectTypes[i].Name == name {
			return &s.EffectTypes[i]
		}
	}
	return nil
}

func (s *Sandbox) NewStatusEffect(piece uint32, typ uint32) *StatusEffect {
	s.Effects = append(s.Effects, StatusEffect{
		Piece: piece,
		Typ:   typ,
	})
	return &s.Effects[len(s.Effects)-1]
}

func (s *Sandbox) RemoveStatusEffect(piece uint32, typ uint32) {
	for i, effect := range s.Effects {
		if effect.Typ == typ && effect.Piece == piece {
			s.Effects = append(s.Effects[:i], s.Effects[i+1:]...)
			return
		}
	}
}

func (s *Sandbox) GetStatusEffectCount(pieceId uint32, statusType uint32) int {
	var count = 0
	for _, effect := range sandbox.Effects {
		if effect.Typ == statusType && effect.Piece == pieceId {
			count++
		}
	}
	return count
}

func (s *Sandbox) GetStatusEffectsOnPiece(pieceId uint32) []uint32 {
	var effects = make([]uint32, 0)
	for _, effect := range s.Effects {
		if effect.Piece == pieceId {
			effects = append(effects, effect.Typ)
		}
	}
	return effects
}

func (s *Sandbox) RegisterObstacleType(name string, tex rl.Texture2D) *ObstacleType {
	// We assume obstacle types are never unregistered
	s.ObstacleTypes = append(s.ObstacleTypes, ObstacleType{
		Id:   uint32(len(s.ObstacleTypes)),
		Name: name,
		Tex:  tex,
	})
	return &s.ObstacleTypes[len(s.ObstacleTypes)-1]
}

func (s *Sandbox) GetObstacleType(id uint32) *ObstacleType {
	return &s.ObstacleTypes[id]
}

func (s *Sandbox) GetObstacleTypeByName(name string) *ObstacleType {
	for i := 0; i < len(s.ObstacleTypes); i++ {
		if s.ObstacleTypes[i].Name == name {
			return &s.ObstacleTypes[i]
		}
	}
	return nil
}

func (s *Sandbox) NewObstacle(coord Vec2, board uint32, typ uint32) *Obstacle {
	s.Obstacles = append(s.Obstacles, Obstacle{
		Coord: coord,
		Board: board,
		Typ:   typ,
	})
	return &s.Obstacles[len(s.Obstacles)-1]
}

func (s *Sandbox) GetObstaclesAt(coord Vec2, board uint32) []uint32 {
	var obstacles = make([]uint32, 0)
	for _, obstacle := range s.Obstacles {
		if obstacle.Coord == coord && obstacle.Board == board {
			obstacles = append(obstacles, obstacle.Typ)
		}
	}
	return obstacles
}

func (s *Sandbox) GetObstacleCount(coord Vec2, board uint32, typ uint32) int {
	var count = 0
	for i := 0; i < len(s.Obstacles); i++ {
		if s.Obstacles[i].Typ == typ && s.Obstacles[i].Board == board && s.Obstacles[i].Coord == coord {
			count++
		}
	}
	return count
}

func (s *Sandbox) RemoveObstacle(coord Vec2, board uint32, typ uint32) bool {
	for i := 0; i < len(s.Obstacles); i++ {
		if s.Obstacles[i].Typ == typ && s.Obstacles[i].Board == board && s.Obstacles[i].Coord == coord {
			s.Obstacles[i] = s.Obstacles[len(s.Obstacles)-1]
			s.Obstacles = s.Obstacles[:len(s.Obstacles)-1]
			return true
		}
	}
	return false
}

func (s *Sandbox) Render(board uint32, preview bool, selection *Selection) {
	var origo = GetBoardOrigo()
	for i := 0; i < len(s.Tiles); i++ {
		if s.Tiles[i].Board == board {
			s.Tiles[i].Render(s.Boards[board].Style)
		}
	}
	if !preview {
		var rankFileTextColor = ColorAt(Vec2{1, 0}, s.Boards[board].Style)
		for x := 0; x < 8; x++ {
			rl.DrawTextEx(assets.fontComicSansMs, string(rune('a'+x)), rl.NewVector2(float32(origo.X+x*TileSize+7), float32(origo.Y+4+TileSize*8)), UiFontSize, 1, rankFileTextColor)
		}
		for y := 0; y < 8; y++ {
			rl.DrawTextEx(assets.fontComicSansMs, string(rune('1'+y)), rl.NewVector2(float32(origo.X-7-10), float32(origo.Y-7-20+TileSize*8-y*TileSize)), UiFontSize, 1, rankFileTextColor)
		}
	} else {
		const fontSizeGiant = 60
		var rankFileTextColor = ColorAt(Vec2{1, 0}, s.Boards[board].Style)
		var offsetX = TileSize/2 - fontSizeGiant/3
		var offsetY = TileSize/2 - fontSizeGiant/2
		for x := 0; x < 8; x++ {
			rl.DrawTextEx(assets.fontComicSansMs, string(rune('a'+x)), rl.NewVector2(float32(origo.X+x*TileSize+offsetX), float32(origo.Y+offsetY+TileSize*8)), fontSizeGiant, 1, rankFileTextColor)
		}
		for y := 0; y < 8; y++ {
			rl.DrawTextEx(assets.fontComicSansMs, string(rune('1'+y)), rl.NewVector2(float32(origo.X-TileSize+offsetX), float32(origo.Y+offsetY+TileSize*7-y*TileSize)), fontSizeGiant, 1, rankFileTextColor)
		}
	}
	var obstacleHasBeenRenderedFlag = make([]bool, len(s.Obstacles))
	for i := 0; i < len(s.Obstacles); i++ {
		if !obstacleHasBeenRenderedFlag[i] && s.Obstacles[i].Board == board {
			var obstaclesOnThisCoord = make([]*Obstacle, 0)
			for j := i; j < len(s.Obstacles); j++ {
				if !obstacleHasBeenRenderedFlag[j] && s.Obstacles[i].Coord == s.Obstacles[j].Coord && (s.Obstacles[j].Board == board || (s.Obstacles[j].Board == OffBoard && !preview)) {
					obstacleHasBeenRenderedFlag[j] = true
					obstaclesOnThisCoord = append(obstaclesOnThisCoord, &s.Obstacles[j])
				}
			}
			for j := 0; j < len(obstaclesOnThisCoord); j++ {
				obstaclesOnThisCoord[j].Render(j, len(obstaclesOnThisCoord))
			}
		}
	}
	for i := 0; i < len(s.Pieces); i++ {
		if s.Pieces[i].Board == board || (s.Pieces[i].Board == OffBoard && !preview) {
			s.Pieces[i].Render()
			if selection.IsPieceSelected(s.Pieces[i].Id) {
				var pos = GetBoardOrigo().Add(s.Pieces[i].Coord.Scale(TileSize))
				var rect = rl.NewRectangle(float32(pos.X+4), float32(pos.Y+4), float32(TileSize*s.Pieces[i].Scale-8), float32(TileSize*s.Pieces[i].Scale-8))
				var thickness = 1
				if preview {
					thickness = 4
				}
				rl.DrawRectangleLinesEx(rect, float32(thickness), rl.Blue)
			}
		}
	}
	for i := 0; i < len(s.Pieces); i++ {
		if s.Pieces[i].Board == board || (s.Pieces[i].Board == OffBoard && !preview) {
			s.RenderStatusEffectsOfPiece(&s.Pieces[i])
		}
	}
	if coord, ok := selection.GetSelectedCoord(); ok {
		var pos = GetBoardOrigo().Add(coord.Scale(TileSize))
		var rect = rl.NewRectangle(float32(pos.X+4), float32(pos.Y+4), TileSize-8, TileSize-8)
		var thickness = 1
		if preview {
			thickness = 4
		}
		rl.DrawRectangleLinesEx(rect, float32(thickness), rl.Red)
	}

	var selectedId, hasSelection = selection.GetSelectedPieceId()
	var selectedPiece = s.GetPiece(selectedId)
	if hasSelection && selectedPiece.Board != board && !(selectedPiece.Board == OffBoard && !preview) {
		selectedPiece.RenderCrossPlaneIndicator()
	}
}

func (s *Sandbox) RenderStatusEffectsOfPiece(piece *Piece) {
	var effectsToRenderAtBottom = make([]*StatusEffect, 0)
	var stuns = 0
	for j := 0; j < len(s.Effects); j++ {
		if s.Effects[j].Piece == piece.Id {
			var typ = s.GetStatusEffectType(s.Effects[j].Typ)
			switch typ.Style {
			case RenderStyleBottom:
				effectsToRenderAtBottom = append(effectsToRenderAtBottom, &s.Effects[j])
			case RenderStyleStun:
				typ.RenderAbove(piece.Coord, stuns, float32(piece.Scale))
				stuns++
			}
		}
	}
	sort.Slice(effectsToRenderAtBottom, func(a, b int) bool {
		return effectsToRenderAtBottom[a].Typ > effectsToRenderAtBottom[b].Typ
	})
	for j := 0; j < len(effectsToRenderAtBottom); j++ {
		var typ = s.GetStatusEffectType(effectsToRenderAtBottom[j].Typ)
		typ.RenderAtBottom(piece.Coord, j, len(effectsToRenderAtBottom), float32(piece.Scale))
	}
}
