package main

import (
	"github.com/v2pro/plz/test/testify/assert"
	"testing"
)

func Test_UpdateQuality(t *testing.T) {
	t.Run("when a day has passed should update the quality of all the items", func(t *testing.T) {
		var items = []*Item{
			{"+5 Dexterity Vest", 10, 20},
			{"Aged Brie", 2, 0},
			{"Elixir of the Mongoose", 5, 7},
			{"Sulfuras, Hand of Ragnaros", 0, 80},
			{"Sulfuras, Hand of Ragnaros", -1, 80},
			{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
			{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		}
		UpdateQuality(items)

		expectedItem := []*Item{
			{"+5 Dexterity Vest", 9, 19},
			{"Aged Brie", 1, 1},
			{"Elixir of the Mongoose", 4, 6},
			{"Sulfuras, Hand of Ragnaros", 0, 80},
			{"Sulfuras, Hand of Ragnaros", -1, 80},
			{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
			{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
			{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		}
		assert.Equal(t, expectedItem, items)
	})
	t.Run("when five day has passed should update the quality of all the items", func(t *testing.T) {
		var items = []*Item{
			{"+5 Dexterity Vest", 10, 20},
			{"Aged Brie", 2, 0},
			{"Elixir of the Mongoose", 5, 7},
			{"Sulfuras, Hand of Ragnaros", 0, 80},
			{"Sulfuras, Hand of Ragnaros", -1, 80},
			{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
			{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		}
		for i := 0; i < 5; i++ {
			UpdateQuality(items)
		}
		expectedItem := []*Item{
			{"+5 Dexterity Vest", 5, 15},
			{"Aged Brie", -3, 8},
			{"Elixir of the Mongoose", 0, 2},
			{"Sulfuras, Hand of Ragnaros", 0, 80},
			{"Sulfuras, Hand of Ragnaros", -1, 80},
			{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
			{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
			{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		}
		assert.Equal(t, expectedItem, items)
	})
}
