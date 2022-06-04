package user

type Dungeon struct {
	Stages   map[string]Stage    `json:"stages"`
	CowLevel map[string]CowLevel `json:"cowLevel"`
}

type CowLevel struct {
	ID   string        `json:"id"`  
	Type string        `json:"type"`
	Val  []interface{} `json:"val"` // MISSING DATA
	Fts  int64         `json:"fts"` 
	RTS  int64         `json:"rts"` 
}

type Stage struct {
	StageID         string `json:"stageId"`        
	CompleteTimes   int64  `json:"completeTimes"`  
	StartTimes      int64  `json:"startTimes"`     
	PracticeTimes   int64  `json:"practiceTimes"`  
	State           int64  `json:"state"`          
	HasBattleReplay int64  `json:"hasBattleReplay"`
	NoCostCnt       int64  `json:"noCostCnt"`      
}