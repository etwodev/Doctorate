package user

type Actor struct {
	Token       int64         `json:"token"`      
	Agenda      int64         `json:"agenda"`     
	Flag        Flag          `json:"flag"`       
	LastRefresh int64         `json:"lastRefresh"`
	Mission     ActorMission  `json:"mission"`    
	FavorList   []interface{} `json:"favorList"`  // MISSING DATA
}

type Flag struct {
	Agenda  bool `json:"agenda"` 
	Mission bool `json:"mission"`
}

type ActorMission struct {
	Random    int64            `json:"random"`   
	Condition Condition        `json:"condition"`
	Pool      []Pool           `json:"pool"`     
	Board     []ActorBoard     `json:"board"`    
	Complete  map[string]int64 `json:"complete"` 
}

type ActorBoard struct {
	Mission  Pool     `json:"mission"` 
	Progress ActorProgress `json:"progress"`
}

type Pool struct {
	MissionID        string `json:"missionId"`       
	OrgID            string `json:"orgId"`           
	PrincipalID      string `json:"principalId"`     
	PrincipalDescIdx int64  `json:"principalDescIdx"`
	RewardGroupID    string `json:"rewardGroupId"`   
}

type ActorProgress struct {
	Value  int64 `json:"value"` 
	Target int64 `json:"target"`
}

type Condition struct {
	OrgID  string      `json:"orgId"` 
	Reward interface{} `json:"reward"`  // MISSING DATA
}
