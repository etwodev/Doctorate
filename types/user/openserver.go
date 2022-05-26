package user

type OpenServer struct {
	CheckIn    CheckInOS    `json:"checkIn"`   
	ChainLogin ChainLogin `json:"chainLogin"`
}

type ChainLogin struct {
	IsAvailable bool    `json:"isAvailable"`
	NowIndex    int64   `json:"nowIndex"`   
	History     []int64 `json:"history"`    
}

type CheckInOS struct {
	IsAvailable bool    `json:"isAvailable"`
	History     []int64 `json:"history"`    
}
