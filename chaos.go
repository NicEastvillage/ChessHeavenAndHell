package main

import (
	"fmt"
	"math/rand"
)

type RngStuff struct {
	chaosEntries   []string
	chaosShown     [3]string
	piece          string
	plane          string
	tile           string
	unoccupiedTile string // No piece on it, but obstacle allowed
	emptyTile      string
}

func NewRngStuff() RngStuff {
	var res = RngStuff{
		chaosShown: [3]string{"", "", ""},
	}
	res.AddChaosOption("Shuffle the shop ordering and gain 6 coins.")
	res.AddChaosOption("Give both players 30 coins at the start of your next round.")
	res.AddChaosOption("Freeze all tiles with Pawns on them.")
	res.AddChaosOption("Summon two Imps on Earth on random empty tiles.")
	res.AddChaosOption("Pick a file in a plane. Curse all non-royal pieces in the chosen file.")
	res.AddChaosOption("Replace all Bishops with Rooks and vice versa.")
	res.AddChaosOption("Reset both players' coin count to 0.")
	res.AddChaosOption("Spread all ice to adjacent tiles.")
	res.AddChaosOption("Summon the Devil (Giant Imp with double move) in the middle of Hell.")
	res.AddChaosOption("Summon an Angel on an empty tile in Heaven.")
	res.AddChaosOption("Destroy the outermost ranks OR files of Heaven.")
	res.AddChaosOption("Move each Pawn a step backwards, if the destination tile is unoccupied.")
	res.AddChaosOption("Stun 4 random pieces for 3 rounds.")
	res.AddChaosOption("Destroy the outermost ranks OR files of Hell.")
	res.AddChaosOption("Duplicate a non-King piece to an adjacent empty tile. Then Curse both.")
	res.AddChaosOption("Summon a Rook on a back rank for each player. Owner chooses where. Opponent starts.")
	res.AddChaosOption("Spawn a Chaos Orb in each plane on a random empty tile.")
	res.AddChaosOption("Make a piece giant.")
	res.AddChaosOption("Swap a piece in Hell with one in Heaven.")
	res.AddChaosOption("Curse a random piece of each player.")
	res.AddChaosOption("Upgrade all Knights.")
	res.AddChaosOption("Replace all ice with fire OR all fire with ice.")
	res.AddChaosOption("Spread all fire to adjacent tiles.")
	res.AddChaosOption("Pick two Pawns. They move forward at the end of their owner's turn if the tile is unoccupied.")
	res.AddChaosOption("Fill empty tiles on Earth with coins.")
	res.AddChaosOption("Start 3 fires on random empty tiles in a plane of your choosing.")
	res.AddChaosOption("Freeze one unoccupied tile and start a fire on a different unoccupied tile.")
	res.AddChaosOption("Downgrade an upgraded piece of your choosing. Give experience to every piece around it.")
	res.AddChaosOption("You have two actions on your next turn.")
	res.AddChaosOption("Remove four options from the shop. Reduce the price of remaining options by 1.")
	res.AddChaosOption("Destroy all traps. You may plant a trap for each destroyed.")
	res.AddChaosOption("Move all Bishops to their starting position on Earth. Then give them experience.")
	res.AddChaosOption("Upgrade all Kings.")
	res.AddChaosOption("Swap your King and a Pawn on the same plane.")
	res.AddChaosOption("Pick a plane. Downgrade all pieces on that plane.")
	res.AddChaosOption("Replace all Pawns with Suicide Bombers.")
	res.AddChaosOption("Pick a Pawn of yours. While it is a Pawn make as many moves and captures with it as you desire/can.")
	res.AddChaosOption("Pick a plane. Spawn white Pawns on empty rank 2 tiles and black Pawns on empty rank 7 tiles.")
	res.AddChaosOption("Gain 8 coins. Then buy something from the shop. Then remove that option from the shop.")
	res.AddChaosOption("Pick a non-royal piece. Move it to a random position on a random plane.")
	res.AddChaosOption("From now on, piece only need 1 experience to be upgraded.")
	res.AddChaosOption("Gain an infinite number of coins, but you can only buy odd-cost options from the shop.")
	res.AddChaosOption("Something random. Reroll the Chaos options and pick the top one.")
	return res
}

func (r *RngStuff) AddChaosOption(description string) {
	r.chaosEntries = append(r.chaosEntries, description)
}

func (r *RngStuff) RerollChaosShown() {
	// First
	if len(r.chaosEntries) == 0 {
		r.chaosShown = [3]string{"", "", ""}
		return
	}
	var first = rand.Intn(len(r.chaosEntries))
	r.chaosShown[0] = r.chaosEntries[first]

	// Second
	if len(r.chaosEntries) == 1 {
		r.chaosShown[1] = ""
		r.chaosShown[2] = ""
		return
	}
	var second = rand.Intn(len(r.chaosEntries))
	for first == second {
		second = rand.Intn(len(r.chaosEntries))
	}
	r.chaosShown[1] = r.chaosEntries[second]

	// Third
	if len(r.chaosEntries) == 2 {
		r.chaosShown[2] = ""
		return
	}
	var third = rand.Intn(len(r.chaosEntries))
	for first == third || second == third {
		third = rand.Intn(len(r.chaosEntries))
	}
	r.chaosShown[2] = r.chaosEntries[third]
}

func (r *RngStuff) RerollPiece(sandbox *Sandbox) {
	if len(sandbox.Pieces) == 0 {
		r.piece = ""
		return
	}
	var piece = sandbox.Pieces[rand.Intn(len(sandbox.Pieces))]
	var file = piece.Coord.X
	var rank = piece.Coord.Y
	var fileLetter = 'a' + (file+('z'-'a'+1))%('z'-'a'+1)
	r.piece = fmt.Sprintf("%s, %c%d %s", sandbox.GetPieceType(piece.Typ).Name, fileLetter, rank+1, NameOfBoard(piece.Board))
}

func (r *RngStuff) RerollPlane() {
	var n = rand.Intn(3)
	r.plane = NameOfBoard(uint32(n))
}

func (r *RngStuff) RerollTile() {
	var file = rand.Intn(8)
	var rank = rand.Intn(8)
	r.tile = fmt.Sprintf("%c%d", 'a'+file, rank+1)
}

func (r *RngStuff) RerollUnoccupiedTile(sandbox *Sandbox) {
	// No Piece on it, but obstacle allowed
	// There is a possibility that such a tile does not exist
	var tiles = make([]Tile, 0)
	for _, tile := range sandbox.Tiles {
		if sandbox.GetPieceAt(tile.Coord, tile.Board) == nil {
			tiles = append(tiles, tile)
		}
	}
	if len(tiles) == 0 {
		r.unoccupiedTile = ""
		return
	}
	var tile = tiles[rand.Intn(len(tiles))]
	var file = tile.Coord.X
	var rank = tile.Coord.Y
	var fileLetter = 'a' + (file+('z'-'a'+1))%('z'-'a'+1)
	r.unoccupiedTile = fmt.Sprintf("%c%d %s", fileLetter, rank+1, NameOfBoard(tile.Board))
}

func (r *RngStuff) RerollEmptyTile(sandbox *Sandbox) {
	// There is a possibility that such a tile does not exist
	var tiles = make([]Tile, 0)
	for _, tile := range sandbox.Tiles {
		if sandbox.GetPieceAt(tile.Coord, tile.Board) == nil && len(sandbox.GetObstaclesAt(tile.Coord, tile.Board)) == 0 {
			tiles = append(tiles, tile)
		}
	}
	if len(tiles) == 0 {
		r.emptyTile = ""
		return
	}
	var tile = tiles[rand.Intn(len(tiles))]
	var file = tile.Coord.X
	var rank = tile.Coord.Y
	var fileLetter = 'a' + (file+('z'-'a'+1))%('z'-'a'+1)
	r.emptyTile = fmt.Sprintf("%c%d %s", fileLetter, rank+1, NameOfBoard(tile.Board))
}
