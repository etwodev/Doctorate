package user

type UserStatus struct {
	NickName                     string           `json:"nickName"`                    
	NickNumber                   string           `json:"nickNumber"`                  
	Level                        int64            `json:"level"`                       
	Exp                          int64            `json:"exp"`                         
	SocialPoint                  int64            `json:"socialPoint"`                 
	GachaTicket                  int64            `json:"gachaTicket"`                 
	TenGachaTicket               int64            `json:"tenGachaTicket"`              
	InstantFinishTicket          int64            `json:"instantFinishTicket"`         
	HggShard                     int64            `json:"hggShard"`                    
	LggShard                     int64            `json:"lggShard"`                    
	RecruitLicense               int64            `json:"recruitLicense"`              
	Progress                     int64            `json:"progress"`                    
	BuyApRemainTimes             int64            `json:"buyApRemainTimes"`            
	ApLimitUpFlag                int64            `json:"apLimitUpFlag"`               
	UID                          string           `json:"uid"`                         
	Flags                        map[string]int64 `json:"flags"`                   
	Ap                           int64            `json:"ap"`                          
	MaxAp                        int64            `json:"maxAp"`                       
	PayDiamond                   int64            `json:"payDiamond"`                  
	FreeDiamond                  int64            `json:"freeDiamond"`                 
	DiamondShard                 int64            `json:"diamondShard"`                
	Gold                         int64            `json:"gold"`                        
	PracticeTicket               int64            `json:"practiceTicket"`              
	LastRefreshTs                int64            `json:"lastRefreshTs"`               
	LastApAddTime                int64            `json:"lastApAddTime"`               
	MainStageProgress            string           `json:"mainStageProgress"`           
	RegisterTs                   int64            `json:"registerTs"`                  
	LastOnlineTs                 int64            `json:"lastOnlineTs"`                
	ServerName                   string           `json:"serverName"`                  
	AvatarID                     string           `json:"avatarId"`                    
	Resume                       string           `json:"resume"`                      
	FriendNumLimit               int64            `json:"friendNumLimit"`              
	MonthlySubscriptionStartTime int64            `json:"monthlySubscriptionStartTime"`
	MonthlySubscriptionEndTime   int64            `json:"monthlySubscriptionEndTime"`  
	Secretary                    string           `json:"secretary"`                   
	SecretarySkinID              string           `json:"secretarySkinId"`             
	TipMonthlyCardExpireTs       int64            `json:"tipMonthlyCardExpireTs"`      
	Avatar                       UserAvatar       `json:"avatar"`                      
	GlobalVoiceLAN               string           `json:"globalVoiceLan"`              
}

type UserAvatar struct {
	Type string `json:"type"`
	ID   string `json:"id"`  
}
