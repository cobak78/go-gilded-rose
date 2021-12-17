package main

type Item struct {
	name            string
	sellIn, quality int
}

const agedBrie = "Aged Brie"
const backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
const sulfuras = "Sulfuras, Hand of Ragnaros"

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		updateQualityItem(items[i])
	}
}

func updateQualityItem(item *Item) {
	if item.name == backstagePasses {
		if item.sellIn < 11 {
			incrementQuality(item, 2)
		}
		if item.sellIn < 6 {
			incrementQuality(item, 1)
		}
	} else {
		decreaseQuality(item)
	}
	if item.name == agedBrie || item.name == sulfuras {
		incrementQuality(item, 1)
	} else {
		decreaseQuality(item)
	}

	if item.name != sulfuras {
		item.sellIn--
	}
	checkSellInDateHasPassed(item)
}

func incrementQuality(item *Item, incrementValue int) {
	if item.quality < 50 {
		item.quality += incrementValue
	}
}

func checkSellInDateHasPassed(item *Item) {
	if item.sellIn < 0 {
		if item.name == agedBrie {
			if item.quality < 50 {
				item.quality++
			}
			return
		}
		if item.name == backstagePasses {
			item.quality--
			return
		}
		if item.quality > 0 {
			if item.name != sulfuras {
				item.quality--
			}
		}
	}
}

func decreaseQuality(item *Item) {
	if item.quality > 0 {
		item.quality--
	}
}
