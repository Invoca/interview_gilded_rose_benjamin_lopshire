package main

type Item struct {
	Name   string `json:"name"`
	SellBy int32  `json:"sellBy"`
	Price  int32  `json:"price"`
}

type Inventory struct {
	Items []*Item `json:"items"`
}

func (inventory *Inventory) UpdatePrice() {
	for _, item := range inventory.Items {
		if item.Name != "Fine Art" && item.Name != "Concert Tickets" {
			if item.Price > 0 {
				if item.Name != "Gold Coins" {
					item.Price = item.Price - 1
				}
			}
		} else {
			if item.Price < 50 {
				item.Price = item.Price + 1
				if item.Name == "Concert Tickets" {
					if item.SellBy < 11 {
						if item.Price < 50 {
							item.Price = item.Price + 1
						}
					}
					if item.SellBy < 6 {
						if item.Price < 50 {
							item.Price = item.Price + 1
						}
					}
				}
			}
		}

		if item.Name != "Gold Coins" {
			item.SellBy = item.SellBy - 1
		}

		if item.SellBy < 0 {
			if item.Name != "Fine Art" {
				if item.Name != "Concert Tickets" {
					if item.Price > 0 {
						if item.Name != "Gold Coins" {
							item.Price = item.Price - 1
						}
					}
				} else {
					item.Price = item.Price - item.Price
				}
			} else {
				if item.Price < 50 {
					item.Price = item.Price + 1
				}
			}
		}
	}
}

func main() {
}
