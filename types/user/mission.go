package user

type Mission struct {
	Missions       Missions         `json:"missions"`      
	MissionRewards MissionRewards   `json:"missionRewards"`
	MissionGroups  map[string]int64 `json:"missionGroups"` 
}

type MissionRewards struct {
	DailyPoint  int64   `json:"dailyPoint"` 
	WeeklyPoint int64   `json:"weeklyPoint"`
	Rewards     Rewards `json:"rewards"`    
}

type Rewards struct {
	Daily  map[string]int64 `json:"DAILY"` 
	Weekly map[string]int64 `json:"WEEKLY"`
}

type Missions struct {
	Openserver map[string]Activity `json:"OPENSERVER"`
	Daily      map[string]Activity `json:"DAILY"`     
	Weekly     map[string]Activity `json:"WEEKLY"`    
	Guide      map[string]Activity `json:"GUIDE"`     
	Main       map[string]Activity `json:"MAIN"`      
	Activity   map[string]Activity `json:"ACTIVITY"`  
	Sub        map[string]Activity `json:"SUB"`       
}

type Activity struct {
	State    int64      `json:"state"`   
	Progress []Progress `json:"progress"`
}

type Progress struct {
	Target int64 `json:"target"`
	Value  int64 `json:"value"` 
}