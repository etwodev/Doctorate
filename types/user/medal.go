package user

type Medal struct {
	Medals map[string]MedalValue `json:"medals"`
	Custom Custom                `json:"custom"`
}

type Custom struct {
	CurrentIndex string  `json:"currentIndex"`
	Customs      map[string]Layer `json:"customs"`     
}

type Layer struct {
	Layout []Layout `json:"layout"`
}

type Layout struct {
	ID  string  `json:"id"` 
	Pos []int64 `json:"pos"`
}

type MedalValue struct {
	ID     string    `json:"id"`              
	Val    [][]int64 `json:"val"`             
	Fts    int64     `json:"fts"`             
	RTS    int64     `json:"rts"`             
	Reward *string   `json:"reward,omitempty"`
}
