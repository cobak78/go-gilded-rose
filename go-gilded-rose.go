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
		if items[i].name != agedBrie && items[i].name != backstagePasses {
			if items[i].quality > 0 {
				if items[i].name != sulfuras {
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == backstagePasses {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}
		if items[i].name != sulfuras {
			items[i].sellIn = items[i].sellIn - 1
		}
		if items[i].sellIn < 0 {
			if items[i].name != agedBrie {
				if items[i].name != backstagePasses {
					if items[i].quality > 0 {
						if items[i].name != sulfuras {
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}
}
