package main

import (
	"math/rand"
)

type ShopEntry struct {
	id            uint32
	price         int
	originalPrice int
	description   string
}

type Shop struct {
	money         [2]int
	entries       []ShopEntry
	unlockedCount int
}

func NewShop() Shop {
	var shop = Shop{
		unlockedCount: 3,
	}
	shop.AddEntry(6, "Stun a non-King piece for 3 rounds.")
	shop.AddEntry(12, "Move a non-royal piece to Earth.")
	shop.AddEntry(8, "Upgrade a non-royal piece.")
	shop.AddEntry(6, "Plant a secret trap.")
	shop.AddEntry(5, "Start a fire on an unoccupied tile.")
	shop.AddEntry(20, "Kill a non-standard piece.")
	shop.AddEntry(4, "Buy a chaos orb.")
	shop.AddEntry(8, "Remove ice and fire in a 2x2 area.")
	shop.AddEntry(6, "Curse a non-royal piece.")
	shop.AddEntry(12, "Spawn new Rook anywhere on your back rank.")
	shop.AddEntry(12, "Get two actions next turn.")
	shop.AddEntry(6, "Move a piece t anywhere on the same plane.")
	shop.AddEntry(6, "Freeze a 2x2 area.")
	shop.AddEntry(2, "Forgive a piece for its sins.")
	shop.AddEntry(8, "Spawn new Knight anywhere on your back rank.")
	shop.AddEntry(10, "Spawn new Bishop anywhere on your back rank.")
	shop.Shuffle()
	return shop
}

func (s *Shop) AddEntry(price int, description string) *ShopEntry {
	s.entries = append(s.entries, ShopEntry{
		id:            uint32(len(s.entries)),
		price:         price,
		originalPrice: price,
		description:   description,
	})
	return &s.entries[len(s.entries)-1]
}

func (s *Shop) WhiteMoney() *int {
	return &s.money[0]
}

func (s *Shop) BlackMoney() *int {
	return &s.money[1]
}

func (s *Shop) Shuffle() {
	rand.Shuffle(len(s.entries), func(i, j int) { s.entries[i], s.entries[j] = s.entries[j], s.entries[i] })
}
