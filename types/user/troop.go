package user

type Troop struct {
	CurCharInstID int64                				   `json:"curCharInstId"`
	CurSquadCount int64                				   `json:"curSquadCount"`
	Squads        map[string]Squad     				   `json:"squads"`    
	Chars         map[string]Char      				   `json:"chars"`  
	CharGroup     map[string]CharGroup 				   `json:"charGroup"`
	CharMission   map[string]map[string]int64          `json:"charMission"`
	Addon         map[string]Addon                	   `json:"addon,omitempty"`
}

type Addon struct {
	Stage	map[string]Mem			`json:"stage,omitempty"`
	Story	map[string]StorySet		`json:"story,omitempty"`
} 

type Mem struct {
	StartTimes    int64 `json:"startTimes"`   
	CompleteTimes int64 `json:"completeTimes"`
	State         int64 `json:"state"`        
	Fts           int64 `json:"fts"`          
	RTS           int64 `json:"rts"`          
	StartTime     int64 `json:"startTime"`    
}

type StorySet struct {
	Fts int64 `json:"fts"`
	RTS int64 `json:"rts"`
}

type CharGroup struct {
	FavorPoint int64 `json:"favorPoint"`
}

type Char struct {
	InstID            int64            `json:"instId"`            
	CharID            string           `json:"charId"`            
	FavorPoint        int64            `json:"favorPoint"`        
	PotentialRank     int64            `json:"potentialRank"`     
	MainSkillLvl      int64            `json:"mainSkillLvl"`      
	Skin              string           `json:"skin"`              
	Level             int64            `json:"level"`             
	Exp               int64            `json:"exp"`               
	EvolvePhase       int64            `json:"evolvePhase"`       
	DefaultSkillIndex int64            `json:"defaultSkillIndex"` 
	GainTime          int64            `json:"gainTime"`          
	Skills            []Skill          `json:"skills"`            
	CurrentEquip      interface{}      `json:"currentEquip"` // MISSING DATA
	Equip             map[string]Equip `json:"equip"`             
	VoiceLAN          string           `json:"voiceLan"`          
	StarMark          *int64           `json:"starMark,omitempty"`
}

type Equip struct {
	Hide   int64 `json:"hide"`  
	Locked int64 `json:"locked"`
	Level  int64 `json:"level"` 
}

type Skill struct {
	SkillID             string `json:"skillId"`            
	Unlock              int64  `json:"unlock"`             
	State               int64  `json:"state"`              
	SpecializeLevel     int64  `json:"specializeLevel"`    
	CompleteUpgradeTime int64  `json:"completeUpgradeTime"`
}

type Squad struct {
	SquadID string  `json:"squadId"`
	Name    string  `json:"name"`   
	Slots   []*Slot `json:"slots"`  
}

type Slot struct {
	CharInstID int64 `json:"charInstId"`
	SkillIndex int64 `json:"skillIndex"`
}
