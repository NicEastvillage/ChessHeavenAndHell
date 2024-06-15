package main

import "math/rand"

const ShopInitiallyUnlockedCount = 3

type ShopEntry struct {
	Id            uint32 `json:"id"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"originalPrice"`
	Description   string `json:"description"`
	Unlocked      bool   `json:"unlocked"`
}

type Shop struct {
	Money         [2]int      `json:"money"`
	Entries       []ShopEntry `json:"entries"`
	HiddenEntries []ShopEntry `json:"hiddenEntries"`
}

func NewShop() Shop {
	var shop = Shop{}
	shop.AddEntry(6, "Stun a non-King piece for 3 rounds.", false)
	shop.AddEntry(12, "Move a non-royal piece to Earth.", false)
	shop.AddEntry(8, "Upgrade a non-royal piece.", false)
	shop.AddEntry(6, "Plant a secret trap.", false)
	shop.AddEntry(5, "Start a fire on an unoccupied tile.", false)
	shop.AddEntry(20, "Kill a non-standard piece.", false)
	shop.AddEntry(4, "Buy a chaos orb.", false)
	shop.AddEntry(8, "Remove ice and fire in a 2x2 area.", false)
	shop.AddEntry(6, "Curse a non-royal piece.", false)
	shop.AddEntry(12, "Spawn new Rook anywhere on your back rank.", false)
	shop.AddEntry(12, "Get two actions next turn.", false)
	shop.AddEntry(6, "Move a piece to anywhere on the same plane.", false)
	shop.AddEntry(6, "Freeze a 2x2 area.", false)
	shop.AddEntry(2, "Forgive a piece for its sins.", false)
	shop.AddEntry(8, "Spawn new Knight anywhere on your back rank.", false)
	shop.AddEntry(10, "Spawn new Bishop anywhere on your back rank.", false)
	rand.Shuffle(len(shop.Entries), func(i, j int) { shop.Entries[i], shop.Entries[j] = shop.Entries[j], shop.Entries[i] })
	for i := 0; i < ShopInitiallyUnlockedCount; i++ {
		shop.Entries[i].Unlocked = true
	}
	return shop
}

func (s *Shop) AddEntry(price int, description string, unlocked bool) *ShopEntry {
	s.Entries = append(s.Entries, ShopEntry{
		Id:            uint32(len(s.Entries)),
		Price:         price,
		OriginalPrice: price,
		Description:   description,
		Unlocked:      unlocked,
	})
	return &s.Entries[len(s.Entries)-1]
}

func (s *Shop) WhiteMoney() *int {
	return &s.Money[0]
}

func (s *Shop) BlackMoney() *int {
	return &s.Money[1]
}

func (s *Shop) GetEntry(id uint32) *ShopEntry {
	for i := 0; i < len(s.Entries); i++ {
		if s.Entries[i].Id == id {
			return &s.Entries[i]
		}
	}
	return nil
}
