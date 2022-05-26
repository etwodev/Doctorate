package user

type Social struct {
	AssistCharList  []AssistCharList `json:"assistCharList"` 
	YesterdayReward YesterdayReward  `json:"yesterdayReward"`
	YCrisisSs       string           `json:"yCrisisSs"`      
	MedalBoard      MedalBoard       `json:"medalBoard"`     
}

type AssistCharList struct {
	CharInstID int64 `json:"charInstId"`
	SkillIndex int64 `json:"skillIndex"`
}

type MedalBoard struct {
	Type              string	`json:"type"`             
	Custom            string	`json:"custom"`           
	Template          string	`json:"template"`
	TemplateMedalList []string	`json:"templateMedalList"`
}

type YesterdayReward struct {
	CanReceive    int64 `json:"canReceive"`   
	AssistAmount  int64 `json:"assistAmount"` 
	ComfortAmount int64 `json:"comfortAmount"`
	First         int64 `json:"first"`        
}