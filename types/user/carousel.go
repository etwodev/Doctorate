package user

type Carousel struct {
	FurnitureShop FurnitureShop `json:"furnitureShop"` // MISSING DATA
}

type FurnitureShop struct {
	Goods  interface{}  `json:"goods"` 				   // MISSING DATA
	Groups map[string]int64 		`json:"groups"`
}
