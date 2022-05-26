package user

type DexNav struct {
	Character map[string]Character 			`json:"character"`
	Formula   Formula             			`json:"formula"`  
	Enemy     Enemy                			`json:"enemy"`    
	TeamV2    map[string]map[string]int64	`json:"teamV2"`   
}

type Character struct {
	CharInstID int64 `json:"charInstId"`
	Count      int64 `json:"count"`     
}

type Enemy struct {
	Enemies map[string]int64    `json:"enemies"`
	Stage   map[string][]string `json:"stage"`  
}

type Formula struct {
	Shop        interface{}      `json:"shop"`     // MISSING DATA
	Manufacture map[string]int64 `json:"manufacture"`
	Workshop    map[string]int64 `json:"workshop"`   
}