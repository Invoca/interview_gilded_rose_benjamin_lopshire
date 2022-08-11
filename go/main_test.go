package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInventoryUpdatePrices(t *testing.T) {
	tests := []struct {
		desc           string
		inputItem      Item
		expectedSellBy int32
		expectedPrice  int32
	}{
		{
			desc:           "reduces Price and SellBy for normal items",
			inputItem:      Item{Name: "Normal Item", SellBy: 10, Price: 20},
			expectedSellBy: 9,
			expectedPrice:  19,
		},
		{
			desc:           "reduces Price twice as fast for normal items past SellBy",
			inputItem:      Item{Name: "Normal Item", SellBy: -1, Price: 20},
			expectedSellBy: -2,
			expectedPrice:  18,
		},
		{
			desc:           "does not allow Price to go negative",
			inputItem:      Item{Name: "Normal Item", SellBy: 10, Price: 0},
			expectedSellBy: 9,
			expectedPrice:  0,
		},
		{
			desc:           "increases Price for Fine Art",
			inputItem:      Item{Name: "Fine Art", SellBy: 10, Price: 20},
			expectedSellBy: 9,
			expectedPrice:  21,
		},
		{
			desc:           "does not allow Price of appreciating items to exceed 50 for Fine Art",
			inputItem:      Item{Name: "Fine Art", SellBy: 10, Price: 50},
			expectedSellBy: 9,
			expectedPrice:  50,
		},
		{
			desc:           "does not allow gold coin Price to exceed 80",
			inputItem:      Item{Name: "Gold Coins", SellBy: 10, Price: 80},
			expectedSellBy: 10,
			expectedPrice:  80,
		},
		{
			desc:           "does not allow Price of appreciating items to exceed 50 for Concert Tickets",
			inputItem:      Item{Name: "Concert Tickets", SellBy: 10, Price: 50},
			expectedSellBy: 9,
			expectedPrice:  50,
		},
		{
			desc:           "increases Price for Concert Tickets by 1 when more than 10 days before SellBy",
			inputItem:      Item{Name: "Concert Tickets", SellBy: 12, Price: 20},
			expectedSellBy: 11,
			expectedPrice:  21,
		},
		{
			desc:           "increases Price for Concert Tickets by 2 when between 6 and 10 days before SellBy",
			inputItem:      Item{Name: "Concert Tickets", SellBy: 7, Price: 20},
			expectedSellBy: 6,
			expectedPrice:  22,
		},
		{
			desc:           "increases Price for Concert Tickets by 3 when less than 6 days before SellBy",
			inputItem:      Item{Name: "Concert Tickets", SellBy: 5, Price: 20},
			expectedSellBy: 4,
			expectedPrice:  23,
		},
		{
			desc:           "reduces Price to 0 when SellBy for Concert Tickets is zero",
			inputItem:      Item{Name: "Concert Tickets", SellBy: 0, Price: 20},
			expectedSellBy: -1,
			expectedPrice:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			inventory := &Inventory{Items: []*Item{&test.inputItem}}
			inventory.UpdatePrice()
			assert.Equal(t, test.expectedPrice, inventory.Items[0].Price)
			assert.Equal(t, test.expectedSellBy, inventory.Items[0].SellBy)
		})
	}
}
