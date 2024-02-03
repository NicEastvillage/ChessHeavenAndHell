package main

import rl "github.com/gen2brain/raylib-go/raylib"

var assets = AssetManager{}

type AssetManager struct {
	texWhitePawn   rl.Texture2D
	texWhiteKnight rl.Texture2D
	texWhiteBishop rl.Texture2D
	texWhiteRook   rl.Texture2D
	texWhiteQueen  rl.Texture2D
	texWhiteKing   rl.Texture2D
	texBlackPawn   rl.Texture2D
	texBlackKnight rl.Texture2D
	texBlackBishop rl.Texture2D
	texBlackRook   rl.Texture2D
	texBlackQueen  rl.Texture2D
	texBlackKing   rl.Texture2D

	texObstacleChaosOrb rl.Texture2D
	texObstacleCoin     rl.Texture2D
	texObstacleFire     rl.Texture2D
	texObstacleIce      rl.Texture2D

	texEffectBlood rl.Texture2D
	texEffectMedal rl.Texture2D
	texEffectCurse rl.Texture2D
}

func (am *AssetManager) LoadAll() {
	am.texWhitePawn = rl.LoadTexture("Assets/white_pawn.png")
	am.texWhiteKnight = rl.LoadTexture("Assets/white_knight.png")
	am.texWhiteBishop = rl.LoadTexture("Assets/white_bishop.png")
	am.texWhiteRook = rl.LoadTexture("Assets/white_rook.png")
	am.texWhiteQueen = rl.LoadTexture("Assets/white_queen.png")
	am.texWhiteKing = rl.LoadTexture("Assets/white_king.png")
	am.texBlackPawn = rl.LoadTexture("Assets/black_pawn.png")
	am.texBlackKnight = rl.LoadTexture("Assets/black_knight.png")
	am.texBlackBishop = rl.LoadTexture("Assets/black_bishop.png")
	am.texBlackRook = rl.LoadTexture("Assets/black_rook.png")
	am.texBlackQueen = rl.LoadTexture("Assets/black_queen.png")
	am.texBlackKing = rl.LoadTexture("Assets/black_king.png")

	am.texObstacleChaosOrb = rl.LoadTexture("Assets/obstacle_chaos_orb.png")
	am.texObstacleCoin = rl.LoadTexture("Assets/obstacle_coin.png")
	am.texObstacleFire = rl.LoadTexture("Assets/obstacle_fire.png")
	am.texObstacleIce = rl.LoadTexture("Assets/obstacle_ice.png")

	am.texEffectBlood = rl.LoadTexture("Assets/se_blood.png")
	am.texEffectMedal = rl.LoadTexture("Assets/se_medal.png")
	am.texEffectCurse = rl.LoadTexture("Assets/se_curse.png")
}

func (am *AssetManager) UnloadAll() {
	rl.UnloadTexture(am.texWhitePawn)
	rl.UnloadTexture(am.texWhiteKnight)
	rl.UnloadTexture(am.texWhiteBishop)
	rl.UnloadTexture(am.texWhiteRook)
	rl.UnloadTexture(am.texWhiteQueen)
	rl.UnloadTexture(am.texWhiteKing)
	rl.UnloadTexture(am.texBlackPawn)
	rl.UnloadTexture(am.texBlackKnight)
	rl.UnloadTexture(am.texBlackBishop)
	rl.UnloadTexture(am.texBlackRook)
	rl.UnloadTexture(am.texBlackQueen)
	rl.UnloadTexture(am.texBlackKing)

	rl.UnloadTexture(am.texObstacleChaosOrb)
	rl.UnloadTexture(am.texObstacleCoin)
	rl.UnloadTexture(am.texObstacleFire)
	rl.UnloadTexture(am.texObstacleIce)

	rl.UnloadTexture(am.texEffectBlood)
	rl.UnloadTexture(am.texEffectMedal)
	rl.UnloadTexture(am.texEffectCurse)
}
