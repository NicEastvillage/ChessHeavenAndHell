package main

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var sandbox = Sandbox{}

type Sandbox struct {
	boards        [3]Board
	tiles         []Tile
	pieceTypes    []PieceType
	pieces        []Piece
	nextPieceId   uint32
	effectTypes   []StatusEffectType
	effects       []StatusEffect
	obstacleTypes []ObstacleType
	obstacles     []Obstacle
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

func (s *Sandbox) RemoveTile(board uint32, coord Vec2) bool {
	if len(s.tiles) == 0 {
		return false
	}
	// The slice is unordered, so we insert the last Tile where there removed Tile was and shorten the slice
	var last = s.tiles[len(s.tiles)-1]
	for i := 0; i < len(s.tiles); i++ {
		if s.tiles[i].board == board && s.tiles[i].coord == coord {
			s.tiles[i] = last
			s.tiles = s.tiles[:len(s.tiles)-1]
			return true
		}
	}
	return false
}

func (s *Sandbox) RegisterPieceType(name string, texWhite rl.Texture2D, texBlack rl.Texture2D) *PieceType {
	// We assume piece types are never unregistered
	s.pieceTypes = append(s.pieceTypes, PieceType{
		id:       uint32(len(s.pieceTypes)),
		name:     name,
		texWhite: texWhite,
		texBlack: texBlack,
	})
	return &s.pieceTypes[len(s.pieceTypes)-1]
}

func (s *Sandbox) GetPieceType(id uint32) *PieceType {
	return &s.pieceTypes[id]
}

func (s *Sandbox) GetPieceTypeByName(name string) *PieceType {
	for i := 0; i < len(s.pieceTypes); i++ {
		if s.pieceTypes[i].name == name {
			return &s.pieceTypes[i]
		}
	}
	return nil
}

func (s *Sandbox) NewPiece(typ uint32, color PieceColor, board uint32, coord Vec2) *Piece {
	s.nextPieceId++
	s.pieces = append(s.pieces, Piece{
		id:    s.nextPieceId - 1,
		typ:   typ,
		color: color,
		board: board,
		coord: coord,
	})
	return &s.pieces[len(s.pieces)-1]
}

func (s *Sandbox) NewPieceFromName(typ string, color PieceColor, board uint32, coord Vec2) *Piece {
	return s.NewPiece(s.GetPieceTypeByName(typ).id, color, board, coord)
}

func (s *Sandbox) GetPiece(id uint32) *Piece {
	for i := 0; i < len(s.pieces); i++ {
		if s.pieces[i].id == id {
			return &s.pieces[i]
		}
	}
	return nil
}

func (s *Sandbox) RemovePiece(id uint32) bool {
	if len(s.pieces) == 0 {
		return false
	}

	s.RemoveEffectsFromPiece(id)

	// The slice is unordered, so we insert the last Piece where there removed Piece was and shorten the slice
	for i := 0; i < len(s.pieces); i++ {
		if s.pieces[i].id == id {
			s.pieces[i] = s.pieces[len(s.pieces)-1]
			s.pieces = s.pieces[:len(s.pieces)-1]
			return true
		}
	}
	return false
}

func (s *Sandbox) RemoveEffectsFromPiece(pieceId uint32) {
	var removedEffects = 0
	for i := len(s.effects) - 1; i >= 0; i-- {
		if s.effects[i].piece == pieceId {
			removedEffects++
			s.effects[i] = s.effects[len(s.effects)-removedEffects]
		}
	}
	s.effects = s.effects[:len(s.effects)-removedEffects]
}

func (s *Sandbox) GetPieceAt(coord Vec2) *Piece {
	for i := 0; i < len(s.pieces); i++ {
		if s.pieces[i].coord == coord {
			return &s.pieces[i]
		}
	}
	return nil
}

func (s *Sandbox) RegisterEffectType(name string, tex rl.Texture2D) *StatusEffectType {
	// We assume effect types are never unregistered
	s.effectTypes = append(s.effectTypes, StatusEffectType{
		id:   uint32(len(s.effectTypes)),
		name: name,
		tex:  tex,
	})
	return &s.effectTypes[len(s.effectTypes)-1]
}

func (s *Sandbox) GetStatusEffectType(id uint32) *StatusEffectType {
	return &s.effectTypes[id]
}

func (s *Sandbox) GetStatusEffectTypeByName(name string) *StatusEffectType {
	for i := 0; i < len(s.effectTypes); i++ {
		if s.effectTypes[i].name == name {
			return &s.effectTypes[i]
		}
	}
	return nil
}

func (s *Sandbox) NewStatusEffect(piece uint32, typ uint32) *StatusEffect {
	s.effects = append(s.effects, StatusEffect{
		piece: piece,
		typ:   typ,
	})
	return &s.effects[len(s.effects)-1]
}

func (s *Sandbox) RemoveStatusEffect(piece uint32, typ uint32) {
	for i, effect := range s.effects {
		if effect.typ == typ && effect.piece == piece {
			s.effects = append(s.effects[:i], s.effects[i+1:]...)
			return
		}
	}
}

func (s *Sandbox) GetStatusEffectCount(pieceId uint32, statusType uint32) int {
	var count = 0
	for _, effect := range sandbox.effects {
		if effect.typ == statusType && effect.piece == pieceId {
			count++
		}
	}
	return count
}

func (s *Sandbox) RegisterObstacleType(name string, tex rl.Texture2D) *ObstacleType {
	// We assume obstacle types are never unregistered
	s.obstacleTypes = append(s.obstacleTypes, ObstacleType{
		id:   uint32(len(s.obstacleTypes)),
		name: name,
		tex:  tex,
	})
	return &s.obstacleTypes[len(s.obstacleTypes)-1]
}

func (s *Sandbox) GetObstacleType(id uint32) *ObstacleType {
	return &s.obstacleTypes[id]
}

func (s *Sandbox) GetObstacleTypeByName(name string) *ObstacleType {
	for i := 0; i < len(s.obstacleTypes); i++ {
		if s.obstacleTypes[i].name == name {
			return &s.obstacleTypes[i]
		}
	}
	return nil
}

func (s *Sandbox) NewObstacle(coord Vec2, board uint32, typ uint32) *Obstacle {
	s.obstacles = append(s.obstacles, Obstacle{
		coord: coord,
		board: board,
		typ:   typ,
	})
	return &s.obstacles[len(s.obstacles)-1]
}

func (s *Sandbox) GetObstacleCount(coord Vec2, typ uint32) int {
	var count = 0
	for i := 0; i < len(s.obstacles); i++ {
		if s.obstacles[i].typ == typ && s.obstacles[i].coord == coord {
			count++
		}
	}
	return count
}

func (s *Sandbox) RemoveObstacle(coord Vec2, typ uint32) bool {
	for i := 0; i < len(s.obstacles); i++ {
		if s.obstacles[i].typ == typ && s.obstacles[i].coord == coord {
			s.obstacles[i] = s.obstacles[len(s.obstacles)-1]
			s.obstacles = s.obstacles[:len(s.obstacles)-1]
			return true
		}
	}
	return false
}

func (s *Sandbox) Render(board uint32, selection *Selection) {
	for i := 0; i < len(s.tiles); i++ {
		if s.tiles[i].board == board {
			s.tiles[i].Render(s.boards[board].style)
		}
	}
	var obstacleHasBeenRenderedFlag = make([]bool, len(s.obstacles))
	for i := 0; i < len(s.obstacles); i++ {
		if !obstacleHasBeenRenderedFlag[i] && s.obstacles[i].board == board {
			var obstaclesOnThisCoord = make([]*Obstacle, 0)
			for j := i; j < len(s.obstacles); j++ {
				if !obstacleHasBeenRenderedFlag[j] && s.obstacles[i].coord == s.obstacles[j].coord && s.obstacles[i].board == board {
					obstacleHasBeenRenderedFlag[j] = true
					obstaclesOnThisCoord = append(obstaclesOnThisCoord, &s.obstacles[j])
				}
			}
			for j := 0; j < len(obstaclesOnThisCoord); j++ {
				obstaclesOnThisCoord[j].Render(j, len(obstaclesOnThisCoord))
			}
		}
	}
	for i := 0; i < len(s.pieces); i++ {
		if s.pieces[i].board == board {
			s.pieces[i].Render(selection)
		}
	}
	for i := 0; i < len(s.pieces); i++ {
		if s.pieces[i].board == board {
			var effectsToRender = make([]*StatusEffect, 0)
			for j := 0; j < len(s.effects); j++ {
				if s.effects[j].piece == s.pieces[i].id {
					effectsToRender = append(effectsToRender, &s.effects[j])
				}
			}
			sort.Slice(effectsToRender, func(a, b int) bool {
				return effectsToRender[a].typ > effectsToRender[b].typ
			})
			for j := 0; j < len(effectsToRender); j++ {
				effectsToRender[j].Render(s.pieces[i].coord, j, len(effectsToRender))
			}
		}
	}
	if coord, ok := selection.GetSelectedCoord(); ok {
		var pos = GetWorldOrigo().Add(coord.Scale(TileSize))
		rl.DrawRectangleLines(int32(pos.x)+4, int32(pos.y)+4, TileSize-8, TileSize-8, rl.Red)
	}
}
