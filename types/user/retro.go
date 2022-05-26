package user

type Retro struct {
	Coin       int64            			`json:"coin"`      
	Supplement int64            			`json:"supplement"`
	Block      map[string]Block 			`json:"block"`     
	Lst        int64            			`json:"lst"`       
	Nst        int64            			`json:"nst"`       
	Trail      map[string]map[string]int64  `json:"trail"`     
}

type Block struct {
	Locked int64 `json:"locked"`
	Open   int64 `json:"open"`  
}
