package main

import rl "github.com/gen2brain/raylib-go/raylib"

const FontChars = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~æøåÆØÅ"

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
	texAngel              rl.Texture2D
	texImp                rl.Texture2D

	texObstacleChaosOrb rl.Texture2D
	texObstacleCoin     rl.Texture2D
	texObstacleFire     rl.Texture2D
	texObstacleIce      rl.Texture2D

	texEffectBlood       rl.Texture2D
	texEffectMedal       rl.Texture2D
	texEffectCurse       rl.Texture2D
	texEffectForcedMove  rl.Texture2D
	texEffectPaid2ndMove rl.Texture2D
	texEffectPortalGun   rl.Texture2D
	texEffectStonks      rl.Texture2D
	texEffectStun        rl.Texture2D
	texEffectWizardHat   rl.Texture2D
	texPieceScale        rl.Texture2D

	fontComicSansMs    rl.Font
	fontComicSansMsBig rl.Font

	sfxTileAddRemove      rl.Sound
	sfxMove               rl.Sound
	sfxCapture            rl.Sound
	sfxPromote            rl.Sound
	sfxCreatePiece        rl.Sound
	sfxDeletePiece        rl.Sound
	sfxChangeColor        rl.Sound
	sfxPieceSizeChange    rl.Sound
	sfxStatusEffectAdd    rl.Sound
	sfxStatusEffectRemove rl.Sound
	sfxObstacleAdd        rl.Sound
	sfxObstacleRemove     rl.Sound
	sfxEarnMoney          rl.Sound
	sfxQuickbuy           rl.Sound
	sfxUndoQuickbuy       rl.Sound
	sfxShopUnlock         rl.Sound
	sfxShopShuffle        rl.Sound
	sfxShopPriceChange    rl.Sound
}

func (am *AssetManager) LoadAll() {
	am.texWhitePawn = rl.LoadTexture("assets/pieces/white_pawn.png")
	am.texWhiteKnight = rl.LoadTexture("assets/pieces/white_knight.png")
	am.texWhiteBishop = rl.LoadTexture("assets/pieces/white_bishop.png")
	am.texWhiteRook = rl.LoadTexture("assets/pieces/white_rook.png")
	am.texWhiteQueen = rl.LoadTexture("assets/pieces/white_queen.png")
	am.texWhiteKing = rl.LoadTexture("assets/pieces/white_king.png")
	am.texWhiteBomber = rl.LoadTexture("assets/pieces/white_bomber.png")
	am.texWhiteLeopard = rl.LoadTexture("assets/pieces/white_leopard.png")
	am.texWhiteChecker = rl.LoadTexture("assets/pieces/white_checker.png")
	am.texWhiteMountedArcher = rl.LoadTexture("assets/pieces/white_mounted_archer.png")
	am.texWhiteWizard = rl.LoadTexture("assets/pieces/white_wizard.png")
	am.texWhiteArchbishop = rl.LoadTexture("assets/pieces/white_archbishop.png")
	am.texWhiteFortress = rl.LoadTexture("assets/pieces/white_fortress.png")
	am.texWhiteScout = rl.LoadTexture("assets/pieces/white_scout.png")
	am.texWhiteWarlock = rl.LoadTexture("assets/pieces/white_warlock.png")
	am.texBlackPawn = rl.LoadTexture("assets/pieces/black_pawn.png")
	am.texBlackKnight = rl.LoadTexture("assets/pieces/black_knight.png")
	am.texBlackBishop = rl.LoadTexture("assets/pieces/black_bishop.png")
	am.texBlackRook = rl.LoadTexture("assets/pieces/black_rook.png")
	am.texBlackQueen = rl.LoadTexture("assets/pieces/black_queen.png")
	am.texBlackKing = rl.LoadTexture("assets/pieces/black_king.png")
	am.texBlackBomber = rl.LoadTexture("assets/pieces/black_bomber.png")
	am.texBlackLeopard = rl.LoadTexture("assets/pieces/black_leopard.png")
	am.texBlackChecker = rl.LoadTexture("assets/pieces/black_checker.png")
	am.texBlackMountedArcher = rl.LoadTexture("assets/pieces/black_mounted_archer.png")
	am.texBlackWizard = rl.LoadTexture("assets/pieces/black_wizard.png")
	am.texBlackArchbishop = rl.LoadTexture("assets/pieces/black_archbishop.png")
	am.texBlackFortress = rl.LoadTexture("assets/pieces/black_fortress.png")
	am.texBlackScout = rl.LoadTexture("assets/pieces/black_scout.png")
	am.texBlackWarlock = rl.LoadTexture("assets/pieces/black_warlock.png")
	am.texAngel = rl.LoadTexture("assets/pieces/angel.png")
	am.texImp = rl.LoadTexture("assets/pieces/imp.png")

	am.texObstacleChaosOrb = rl.LoadTexture("assets/obstacles/chaos_orb.png")
	am.texObstacleCoin = rl.LoadTexture("assets/obstacles/coin.png")
	am.texObstacleFire = rl.LoadTexture("assets/obstacles/fire.png")
	am.texObstacleIce = rl.LoadTexture("assets/obstacles/ice.png")

	am.texEffectBlood = rl.LoadTexture("assets/effects/blood.png")
	am.texEffectMedal = rl.LoadTexture("assets/effects/medal.png")
	am.texEffectCurse = rl.LoadTexture("assets/effects/curse.png")
	am.texEffectForcedMove = rl.LoadTexture("assets/effects/forced_move.png")
	am.texEffectPaid2ndMove = rl.LoadTexture("assets/effects/paid_2nd_move.png")
	am.texEffectPortalGun = rl.LoadTexture("assets/effects/portal_gun.png")
	am.texEffectStonks = rl.LoadTexture("assets/effects/stonks.png")
	am.texEffectStun = rl.LoadTexture("assets/effects/stun.png")
	am.texEffectWizardHat = rl.LoadTexture("assets/effects/wizard_hat.png")
	am.texPieceScale = rl.LoadTexture("assets/effects/piece_scale.png")

	am.fontComicSansMs = rl.LoadFontEx("assets/comic.ttf", 20, []rune(FontChars))
	am.fontComicSansMsBig = rl.LoadFontEx("assets/comic.ttf", 28, []rune(FontChars))

	am.sfxTileAddRemove = rl.LoadSound("assets/sfx/tile_add_remove.ogg")
	am.sfxMove = rl.LoadSound("assets/sfx/move.mp3")
	am.sfxCapture = rl.LoadSound("assets/sfx/capture.mp3")
	am.sfxPromote = rl.LoadSound("assets/sfx/promote.mp3")
	am.sfxCreatePiece = rl.LoadSound("assets/sfx/create_piece.ogg")
	am.sfxDeletePiece = rl.LoadSound("assets/sfx/delete_piece.ogg")
	am.sfxChangeColor = rl.LoadSound("assets/sfx/change_color.ogg")
	am.sfxPieceSizeChange = rl.LoadSound("assets/sfx/piece_size_change.wav")
	am.sfxStatusEffectAdd = rl.LoadSound("assets/sfx/status_effect_add.ogg")
	am.sfxStatusEffectRemove = rl.LoadSound("assets/sfx/status_effect_remove.ogg")
	am.sfxObstacleAdd = rl.LoadSound("assets/sfx/obstacle_add.wav")
	am.sfxObstacleRemove = rl.LoadSound("assets/sfx/obstacle_remove.wav")
	am.sfxEarnMoney = rl.LoadSound("assets/sfx/earn_money.ogg")
	am.sfxQuickbuy = rl.LoadSound("assets/sfx/quickbuy.ogg")
	am.sfxUndoQuickbuy = rl.LoadSound("assets/sfx/undo_quickbuy.ogg")
	am.sfxShopUnlock = rl.LoadSound("assets/sfx/shop_unlock.ogg")
	am.sfxShopShuffle = rl.LoadSound("assets/sfx/shop_shuffle.ogg")
	am.sfxShopPriceChange = rl.LoadSound("assets/sfx/shop_price_change.ogg")
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
	rl.UnloadTexture(am.texAngel)
	rl.UnloadTexture(am.texImp)

	rl.UnloadTexture(am.texObstacleChaosOrb)
	rl.UnloadTexture(am.texObstacleCoin)
	rl.UnloadTexture(am.texObstacleFire)
	rl.UnloadTexture(am.texObstacleIce)

	rl.UnloadTexture(am.texEffectBlood)
	rl.UnloadTexture(am.texEffectMedal)
	rl.UnloadTexture(am.texEffectCurse)
	rl.UnloadTexture(am.texEffectForcedMove)
	rl.UnloadTexture(am.texEffectPaid2ndMove)
	rl.UnloadTexture(am.texEffectPortalGun)
	rl.UnloadTexture(am.texEffectStonks)
	rl.UnloadTexture(am.texEffectStun)
	rl.UnloadTexture(am.texEffectWizardHat)
	rl.UnloadTexture(am.texPieceScale)

	rl.UnloadFont(am.fontComicSansMs)
	rl.UnloadFont(am.fontComicSansMsBig)

	rl.UnloadSound(am.sfxTileAddRemove)
	rl.UnloadSound(am.sfxMove)
	rl.UnloadSound(am.sfxCapture)
	rl.UnloadSound(am.sfxPromote)
	rl.UnloadSound(am.sfxCreatePiece)
	rl.UnloadSound(am.sfxDeletePiece)
	rl.UnloadSound(am.sfxChangeColor)
	rl.UnloadSound(am.sfxPieceSizeChange)
	rl.UnloadSound(am.sfxStatusEffectAdd)
	rl.UnloadSound(am.sfxStatusEffectRemove)
	rl.UnloadSound(am.sfxObstacleAdd)
	rl.UnloadSound(am.sfxObstacleRemove)
	rl.UnloadSound(am.sfxEarnMoney)
	rl.UnloadSound(am.sfxQuickbuy)
	rl.UnloadSound(am.sfxUndoQuickbuy)
	rl.UnloadSound(am.sfxShopUnlock)
	rl.UnloadSound(am.sfxShopShuffle)
	rl.UnloadSound(am.sfxShopPriceChange)
}
