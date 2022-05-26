package user

type Gacha struct {
	Newbee  map[string]int64  				`json:"newbee"` 
	Normal  map[string]Normal 				`json:"normal"` 
	Limit   map[string]LimitedEn			`json:"limit"`  
	Linkage map[string]map[string]Linkage   `json:"linkage"`
}

type LimitedEn struct {
	LeastFree int64 `json:"leastFree"`
}

type Linkage struct {
	Next5      bool   `json:"next5"`     
	Next5Char  string `json:"next5Char"` 
	Must6      bool   `json:"must6"`     
	Must6Char  string `json:"must6Char"` 
	Must6Count int64  `json:"must6Count"`
	Must6Level int64  `json:"must6Level"`
}

type Normal struct {
	Cnt    int64 `json:"cnt"`   
	MaxCnt int64 `json:"maxCnt"`
	Rarity int64 `json:"rarity"`
	Avail  bool  `json:"avail"` 
}
