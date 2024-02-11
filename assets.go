package main

import rl "github.com/gen2brain/raylib-go/raylib"

var assets = AssetManager{}

type AssetManager struct {
	texWhitePawn          rl.Texture2D
	texWhiteKnight        rl.Texture2D
	texWhiteBishop        rl.Texture2D
	texWhiteRook          rl.Texture2D
	texWhiteQueen         rl.Texture2D
	texWhiteKing          rl.Texture2D
	texWhiteBomber        rl.Texture2D
	texWhiteLeopard       rl.Texture2D
	texWhiteChecker       rl.Texture2D
	texWhiteMountedArcher rl.Texture2D
	texWhiteWizard        rl.Texture2D
	texWhiteArchbishop    rl.Texture2D
	texWhiteFortress      rl.Texture2D
	texWhiteScout         rl.Texture2D
	texWhiteWarlock       rl.Texture2D
	texBlackPawn          rl.Texture2D
	texBlackKnight        rl.Texture2D
	texBlackBishop        rl.Texture2D
	texBlackRook          rl.Texture2D
	texBlackQueen         rl.Texture2D
	texBlackKing          rl.Texture2D
	texBlackBomber        rl.Texture2D
	texBlackLeopard       rl.Texture2D
	texBlackChecker       rl.Texture2D
	texBlackMountedArcher rl.Texture2D
	texBlackWizard        rl.Texture2D
	texBlackArchbishop    rl.Texture2D
	texBlackFortress      rl.Texture2D
	texBlackScout         rl.Texture2D
	texBlackWarlock       rl.Texture2D

	texObstacleChaosOrb rl.Texture2D
	texObstacleCoin     rl.Texture2D
	texObstacleFire     rl.Texture2D
	texObstacleIce      rl.Texture2D

	texEffectBlood rl.Texture2D
	texEffectMedal rl.Texture2D
	texEffectCurse rl.Texture2D

	texPieceScale rl.Texture2D
}

func (am *AssetManager) LoadAll() {
	am.texWhitePawn = rl.LoadTexture("Assets/pieces/white_pawn.png")
	am.texWhiteKnight = rl.LoadTexture("Assets/pieces/white_knight.png")
	am.texWhiteBishop = rl.LoadTexture("Assets/pieces/white_bishop.png")
	am.texWhiteRook = rl.LoadTexture("Assets/pieces/white_rook.png")
	am.texWhiteQueen = rl.LoadTexture("Assets/pieces/white_queen.png")
	am.texWhiteKing = rl.LoadTexture("Assets/pieces/white_king.png")
	am.texWhiteBomber = rl.LoadTexture("Assets/pieces/white_bomber.png")
	am.texWhiteLeopard = rl.LoadTexture("Assets/pieces/white_leopard.png")
	am.texWhiteChecker = rl.LoadTexture("Assets/pieces/white_checker.png")
	am.texWhiteMountedArcher = rl.LoadTexture("Assets/pieces/white_mounted_archer.png")
	am.texWhiteWizard = rl.LoadTexture("Assets/pieces/white_wizard.png")
	am.texWhiteArchbishop = rl.LoadTexture("Assets/pieces/white_archbishop.png")
	am.texWhiteFortress = rl.LoadTexture("Assets/pieces/white_fortress.png")
	am.texWhiteScout = rl.LoadTexture("Assets/pieces/white_scout.png")
	am.texWhiteWarlock = rl.LoadTexture("Assets/pieces/white_warlock.png")
	am.texBlackPawn = rl.LoadTexture("Assets/pieces/black_pawn.png")
	am.texBlackKnight = rl.LoadTexture("Assets/pieces/black_knight.png")
	am.texBlackBishop = rl.LoadTexture("Assets/pieces/black_bishop.png")
	am.texBlackRook = rl.LoadTexture("Assets/pieces/black_rook.png")
	am.texBlackQueen = rl.LoadTexture("Assets/pieces/black_queen.png")
	am.texBlackKing = rl.LoadTexture("Assets/pieces/black_king.png")
	am.texBlackBomber = rl.LoadTexture("Assets/pieces/black_bomber.png")
	am.texBlackLeopard = rl.LoadTexture("Assets/pieces/black_leopard.png")
	am.texBlackChecker = rl.LoadTexture("Assets/pieces/black_checker.png")
	am.texBlackMountedArcher = rl.LoadTexture("Assets/pieces/black_mounted_archer.png")
	am.texBlackWizard = rl.LoadTexture("Assets/pieces/black_wizard.png")
	am.texBlackArchbishop = rl.LoadTexture("Assets/pieces/black_archbishop.png")
	am.texBlackFortress = rl.LoadTexture("Assets/pieces/black_fortress.png")
	am.texBlackScout = rl.LoadTexture("Assets/pieces/black_scout.png")
	am.texBlackWarlock = rl.LoadTexture("Assets/pieces/black_warlock.png")

	am.texObstacleChaosOrb = rl.LoadTexture("Assets/obstacle_chaos_orb.png")
	am.texObstacleCoin = rl.LoadTexture("Assets/obstacle_coin.png")
	am.texObstacleFire = rl.LoadTexture("Assets/obstacle_fire.png")
	am.texObstacleIce = rl.LoadTexture("Assets/obstacle_ice.png")

	am.texEffectBlood = rl.LoadTexture("Assets/se_blood.png")
	am.texEffectMedal = rl.LoadTexture("Assets/se_medal.png")
	am.texEffectCurse = rl.LoadTexture("Assets/se_curse.png")

	am.texPieceScale = rl.LoadTexture("Assets/piece_scale.png")
}

func (am *AssetManager) UnloadAll() {
	rl.UnloadTexture(am.texWhitePawn)
	rl.UnloadTexture(am.texWhiteKnight)
	rl.UnloadTexture(am.texWhiteBishop)
	rl.UnloadTexture(am.texWhiteRook)
	rl.UnloadTexture(am.texWhiteQueen)
	rl.UnloadTexture(am.texWhiteKing)
	rl.UnloadTexture(am.texWhiteBomber)
	rl.UnloadTexture(am.texWhiteLeopard)
	rl.UnloadTexture(am.texWhiteChecker)
	rl.UnloadTexture(am.texWhiteMountedArcher)
	rl.UnloadTexture(am.texWhiteWizard)
	rl.UnloadTexture(am.texWhiteArchbishop)
	rl.UnloadTexture(am.texWhiteFortress)
	rl.UnloadTexture(am.texWhiteScout)
	rl.UnloadTexture(am.texWhiteWarlock)
	rl.UnloadTexture(am.texBlackPawn)
	rl.UnloadTexture(am.texBlackKnight)
	rl.UnloadTexture(am.texBlackBishop)
	rl.UnloadTexture(am.texBlackRook)
	rl.UnloadTexture(am.texBlackQueen)
	rl.UnloadTexture(am.texBlackKing)
	rl.UnloadTexture(am.texBlackBomber)
	rl.UnloadTexture(am.texBlackLeopard)
	rl.UnloadTexture(am.texBlackChecker)
	rl.UnloadTexture(am.texBlackMountedArcher)
	rl.UnloadTexture(am.texBlackWizard)
	rl.UnloadTexture(am.texBlackArchbishop)
	rl.UnloadTexture(am.texBlackFortress)
	rl.UnloadTexture(am.texBlackScout)
	rl.UnloadTexture(am.texBlackWarlock)

	rl.UnloadTexture(am.texObstacleChaosOrb)
	rl.UnloadTexture(am.texObstacleCoin)
	rl.UnloadTexture(am.texObstacleFire)
	rl.UnloadTexture(am.texObstacleIce)

	rl.UnloadTexture(am.texEffectBlood)
	rl.UnloadTexture(am.texEffectMedal)
	rl.UnloadTexture(am.texEffectCurse)

	rl.UnloadTexture(am.texPieceScale)
}
