package main

import (
	"encoding/json"
	"fmt"
)

const SerialVersion = 1

type UnknownSerialVersionError struct {
	UnknownVersion int
}

func (e *UnknownSerialVersionError) Error() string {
	return fmt.Sprintf("unknown serial version number %d", e.UnknownVersion)
}

type UnknownTypeMarshallError struct {
	Kind string
	Name string
}

func (e *UnknownTypeMarshallError) Error() string {
	return fmt.Sprintf("unknown %s type %q", e.Kind, e.Name)
}

type UnknownIdMarshallError struct {
	Kind string
	Id   uint32
}

func (e *UnknownIdMarshallError) Error() string {
	return fmt.Sprintf("unknown %s id %d", e.Kind, e.Id)
}

type SerializedSandbox struct {
	Version   int                      `json:"version"`
	Shop      Shop                     `json:"shop"`
	Boards    [3]Board                 `json:"boards"`
	Tiles     []Tile                   `json:"tiles"`
	Pieces    []SerializedPiece        `json:"pieces"`
	Effects   []SerializedStatusEffect `json:"effects"`
	Obstacles []SerializedObstacle     `json:"obstacles"`
}

type SerializedPiece struct {
	Id    uint32     `json:"id"`
	Typ   string     `json:"type"`
	Color PieceColor `json:"color"`
	Board uint32     `json:"board"`
	Coord Vec2       `json:"coord"`
	Scale uint32     `json:"scale"`
}

type SerializedStatusEffect struct {
	Piece uint32 `json:"piece"`
	Typ   string `json:"type"`
}

type SerializedObstacle struct {
	Coord Vec2   `json:"coord"`
	Board uint32 `json:"board"`
	Typ   string `json:"type"`
}

func (s *Sandbox) MarshalJSON() ([]byte, error) {
	var serialized = SerializedSandbox{
		Version:   SerialVersion,
		Shop:      s.Shop,
		Boards:    s.Boards,
		Tiles:     s.Tiles,
		Pieces:    SerializePieces(s),
		Effects:   SerializeStatusEffects(s),
		Obstacles: SerializeObstacles(s),
	}
	return json.Marshal(serialized)
}

func (s *Sandbox) UnmarshalJSON(data []byte) error {
	var serialized = SerializedSandbox{}
	if err := json.Unmarshal(data, &serialized); err != nil {
		return err
	}
	if serialized.Version != SerialVersion {
		return &UnknownSerialVersionError{UnknownVersion: serialized.Version}
	}

	s.Shop = serialized.Shop
	s.Boards = serialized.Boards
	s.Tiles = serialized.Tiles

	err := DeserializePieces(s, serialized.Pieces)
	if err != nil {
		return err
	}
	err = DeserializeStatusEffects(s, serialized.Effects)
	if err != nil {
		return err
	}
	err = DeserializeObstacles(s, serialized.Obstacles)
	if err != nil {
		return err
	}

	return nil
}

func SerializePieces(sb *Sandbox) []SerializedPiece {
	var res = make([]SerializedPiece, len(sb.Pieces))
	for i, piece := range sb.Pieces {
		res[i] = SerializedPiece{
			Id:    piece.Id,
			Typ:   sb.GetPieceType(piece.Typ).Name,
			Color: piece.Color,
			Board: piece.Board,
			Coord: piece.Coord,
			Scale: piece.Scale,
		}
	}
	return res
}

func DeserializePieces(sb *Sandbox, pieces []SerializedPiece) error {
	sb.Pieces = make([]Piece, len(pieces))
	for i, piece := range pieces {
		var typ = sb.GetPieceTypeByName(piece.Typ)
		if typ == nil {
			return &UnknownTypeMarshallError{
				Kind: "piece",
				Name: piece.Typ,
			}
		}
		sb.Pieces[i] = Piece{
			Id:    piece.Id,
			Typ:   typ.Id,
			Color: piece.Color,
			Board: piece.Board,
			Coord: piece.Coord,
			Scale: piece.Scale,
		}
	}
	sb.NextPieceId = 0
	for i := 0; i < len(sb.Pieces); i++ {
		if sb.Pieces[i].Id >= sb.NextPieceId {
			sb.NextPieceId = sb.Pieces[i].Id + 1
		}
	}
	return nil
}

func SerializeStatusEffects(sb *Sandbox) []SerializedStatusEffect {
	var res = make([]SerializedStatusEffect, len(sb.Effects))
	for i, effect := range sb.Effects {
		res[i] = SerializedStatusEffect{
			Piece: effect.Piece,
			Typ:   sb.GetStatusEffectType(effect.Typ).Name,
		}
	}
	return res
}

func DeserializeStatusEffects(sb *Sandbox, effects []SerializedStatusEffect) error {
	sb.Effects = make([]StatusEffect, len(effects))
	for i, effect := range effects {
		if sb.GetPiece(effect.Piece) == nil {
			return &UnknownIdMarshallError{
				Kind: "piece",
				Id:   effect.Piece,
			}
		}
		var typ = sb.GetStatusEffectTypeByName(effect.Typ)
		if typ == nil {
			return &UnknownTypeMarshallError{
				Kind: "status effect",
				Name: effect.Typ,
			}
		}
		sb.Effects[i] = StatusEffect{
			Piece: effect.Piece,
			Typ:   typ.Id,
		}
	}
	return nil
}

func SerializeObstacles(sb *Sandbox) []SerializedObstacle {
	var res = make([]SerializedObstacle, len(sb.Obstacles))
	for i, obstacle := range sb.Obstacles {
		res[i] = SerializedObstacle{
			Coord: obstacle.Coord,
			Board: obstacle.Board,
			Typ:   sb.GetObstacleType(obstacle.Typ).Name,
		}
	}
	return res
}

func DeserializeObstacles(sb *Sandbox, obstacles []SerializedObstacle) error {
	sb.Obstacles = make([]Obstacle, len(obstacles))
	for i, obstacle := range obstacles {
		var typ = sb.GetObstacleTypeByName(obstacle.Typ)
		if typ == nil {
			return &UnknownTypeMarshallError{
				Kind: "obstacle",
				Name: obstacle.Typ,
			}
		}
		sb.Obstacles[i] = Obstacle{
			Coord: obstacle.Coord,
			Board: obstacle.Board,
			Typ:   typ.Id,
		}
	}
	return nil
}
