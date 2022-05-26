package user

type CampaignsV2 struct {
	CampaignCurrentFee int64               `json:"campaignCurrentFee"`
	CampaignTotalFee   int64               `json:"campaignTotalFee"`  
	LastRefreshTs      int64               `json:"lastRefreshTs"`     
	Instances          map[string]Instance `json:"instances"`         
	Open               Open                `json:"open"`              
	Missions           map[string]int64    `json:"missions"`          
}

type Instance struct {
	MaxKills     int64   `json:"maxKills"`    
	RewardStatus []int64 `json:"rewardStatus"`
}

type Open struct {
	Permanent []string `json:"permanent"`
	Rotate    string   `json:"rotate"`   
	RGroup    string   `json:"rGroup"`   
	Training  []string `json:"training"` 
	TGroup    string   `json:"tGroup"`   
	TAllOpen  string   `json:"tAllOpen"` 
}
