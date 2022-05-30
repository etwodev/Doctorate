package user

type Rooms struct {
	Control     map[string]ControlRoom           		`json:"CONTROL"`    
	Elevator    map[string]interface{}  				`json:"ELEVATOR"`   // THIS IS NOT MISSING DATA
	Power       map[string]PowerRoom            	    `json:"POWER"`      
	Manufacture map[string]ManufactureRoom       		`json:"MANUFACTURE"`
	Trading     map[string]TradingRoom           		`json:"TRADING"`    
	Corridor    map[string]interface{}          		`json:"CORRIDOR"`   // THIS IS NOT MISSING DATA
	Dormitory   map[string]DormitoryRoom         		`json:"DORMITORY"`  
	Workshop    map[string]WorkshopRoom          		`json:"WORKSHOP"`   
	Meeting     map[string]MeetingRoom           		`json:"MEETING"`    
	Hire        map[string]HireRoom        				`json:"HIRE"`       
	Training    map[string]TrainingRoom          		`json:"TRAINING"`   
}

type ControlRoom struct {
	Buff		   ControlBuff	`json:"buff"`
	ApCost         int64 		`json:"apCost"`        
	LastUpdateTime int64 		`json:"lastUpdateTime"`
}

type ControlBuff struct {
	Global      ControlBuffGlobal           `json:"global"`     
	Manufacture ControlBuffManufacture      `json:"manufacture"`
	Trading     ControlBuffManufacture      `json:"trading"`    
	ApCost      map[string]int64 			`json:"apCost"`     
	Meeting     ControlBuffMeeting          `json:"meeting"`    
	Point       interface{}      			`json:"point"`      // MISSING DATA
	Hire        ControlBuffHire             `json:"hire"`       
}

type ControlBuffGlobal struct {
	ApCost  int64 		`json:"apCost"`  
	RoomCnt interface{} `json:"roomCnt"` // MISSING DATA
}

type ControlBuffHire struct {
	SPUp ControlSPU `json:"spUp"`
}

type ControlBuffSPUp struct {
	Base int64 `json:"base"`
	Up   int64 `json:"up"`  
}

type ControlBuffManufacture struct {
	Speed     int64 	  `json:"speed"`    
	SSpeed    int64 	  `json:"sSpeed"`   
	RoomSpeed interface{} `json:"roomSpeed"`	// MISSING DATA
}

type ControlBuffMeeting struct {
	Clue int64 `json:"clue"`
}

type ControlSPU struct {
	Base int64 `json:"base"`
	Up   int64 `json:"up"`  
}


type PowerRoom struct {
	Buff PowerBuff `json:"buff"`
}

type PowerBuff struct {
	LaborSpeed  float64     			`json:"laborSpeed"` 
	ApCost      map[string]interface{}  `json:"apCost"`     // MISSING DATA
	Global      map[string]interface{}	`json:"global"`     	// MISSING DATA
	Manufacture map[string]interface{}	`json:"manufacture"`	// MISSING DATA
}


type ManufactureRoom struct {
	Buff              ManufactureBuff `json:"buff"`             
	State             int64       	  `json:"state"`            
	FormulaID         string      	  `json:"formulaId"`        
	RemainSolutionCnt int64       	  `json:"remainSolutionCnt"`
	OutputSolutionCnt int64       	  `json:"outputSolutionCnt"`
	LastUpdateTime    int64      	  `json:"lastUpdateTime"`   
	SaveTime          int64      	  `json:"saveTime"`         
	TailTime          int64     	  `json:"tailTime"`         
	ApCost            int64     	  `json:"apCost"`           
	CompleteWorkTime  int64       	  `json:"completeWorkTime"` 
	Capacity          int64       	  `json:"capacity"`         
	Display           Display     	  `json:"display"`          
	ProcessPoint      float64     	  `json:"processPoint"`
}


type ManufactureBuff struct {
	ApCost   map[string]map[string]int64 	`json:"apCost"`  
	Speed    float64      					`json:"speed"`   
	SSpeed   int64        					`json:"sSpeed"`  
	Capacity int64        					`json:"capacity"`
	MaxSpeed int64        					`json:"maxSpeed"`
	TSpeed   interface{}        			`json:"tSpeed"`  // MISSING DATA
	CSpeed   float64        				`json:"cSpeed"`  
	CapFrom  map[string]int64      			`json:"capFrom"` 
	Point    interface{}         			`json:"point"`   // MISSING DATA
	Flag     interface{}         			`json:"flag"`    // MISSING DATA
}

type Display struct {
	Base int64 `json:"base"`
	Buff int64 `json:"buff"`
}

type TradingRoom struct {
	Buff             TradingBuff   `json:"buff"`            
	State            int64         `json:"state"`           
	LastUpdateTime   int64         `json:"lastUpdateTime"`  
	Strategy         string        `json:"strategy"`        
	StockLimit       int64         `json:"stockLimit"`      
	ApCost           int64         `json:"apCost"`          
	Stock            []interface{} `json:"stock"`           // MISSING DATA
	Next             TradingNext          `json:"next"`            
	CompleteWorkTime int64         `json:"completeWorkTime"`
	Display          Display       `json:"display"`         
}

type TradingBuff struct {
	Speed     float64       `json:"speed"`    
	Limit     int64         `json:"limit"`    
	ApCost    TradingAPCost  `json:"apCost"`   
	Rate      interface{}     `json:"rate"`     // MISSING DATA
	Tgw       []interface{} `json:"tgw"`      	// MISSING DATA
	Point     interface{}     `json:"point"`    // MISSING DATA
	ManuLines interface{}     `json:"manuLines"`	// MISSING DATA
	OrderBuff []interface{} `json:"orderBuff"`		// MISSING DATA
}

type TradingAPCost struct {
	All    int64     			 `json:"all"`   
	Single interface{} 			 `json:"single"`	// MISSING DATA
	Self   map[string]int64      `json:"self"`  
}

type TradingNext struct {
	Order        int64   `json:"order"`       
	ProcessPoint float64 `json:"processPoint"`
	MaxPoint     int64   `json:"maxPoint"`    
	Speed        float64 `json:"speed"`       
}


type DormitoryRoom struct {
	Buff        DormitoryBuff        `json:"buff"`       
	Comfort     int64       		 `json:"comfort"`    
	DiySolution Solution 			 `json:"diySolution"`
}

type DormitoryBuff struct {
	ApCost DormitoryAPCost `json:"apCost"`
	Point  interface{}     `json:"point"` 	// MISSING DATA
}

type DormitoryAPCost struct {
	All    int64  `json:"all"`   
	Single DormitorySingle `json:"single"`
	Self   interface{}  `json:"self"`  	// MISSING DATA
}

type DormitorySingle struct {
	Target interface{} `json:"target"`
	Value  int64       `json:"value"` 
}

type WorkshopRoom struct {
	Buff      WorkshopBuff      `json:"buff"`     
	Statistic WorkshopStatistic `json:"statistic"`
}

type WorkshopBuff struct {
	Rate      WorkshopRate          `json:"rate"`     
	Frate     []interface{} `json:"frate"`    
	Cost      WorkshopCost          `json:"cost"`     
	CostRe    WorkshopCostRe        `json:"costRe"`   
	Recovery  WorkshopRecovery      `json:"recovery"` 
	GoldFree  interface{}      `json:"goldFree"` 
	CostForce WorkshopCostForce     `json:"costForce"`
	FFix      map[string]map[string]interface{}          `json:"fFix"`     // MISSING DATA
}

type WorkshopCost struct {
	Type      string `json:"type"`     
	Limit     int64  `json:"limit"`    
	Reduction int64  `json:"reduction"`
}

type WorkshopCostForce struct {
	Type string `json:"type"`
	Cost int64  `json:"cost"`
}

type WorkshopCostRe struct {
	Type   string `json:"type"`  
	From   int64  `json:"from"`  
	Change int64  `json:"change"`
}


type WorkshopRate struct {
	All       int64 `json:"all"`       
	WEvolve   int64 `json:"W_EVOLVE"`  
	WBuilding int64 `json:"W_BUILDING"`
	WSkill    int64 `json:"W_SKILL"`   
	WAsc      int64 `json:"W_ASC"`     
}

type WorkshopRecovery struct {
	Type    string `json:"type"`   
	Pace    int64  `json:"pace"`   
	Recover int64  `json:"recover"`
}

type WorkshopStatistic struct {
	NoAddition int64 `json:"noAddition"`
}

type MeetingRoom struct {
	Buff             MeetingBuff          `json:"buff"`            
	State            int64         `json:"state"`           
	Speed            int64         `json:"speed"`           
	ProcessPoint     int64         `json:"processPoint"`    
	OwnStock         []MeetingDailyReward `json:"ownStock"`        
	ReceiveStock     []interface{} `json:"receiveStock"`    // MISSING DATA
	Board            MeetingBoard         `json:"board"`           
	SocialReward     MeetingSocialReward  `json:"socialReward"`    
	DailyReward      MeetingDailyReward   `json:"dailyReward"`     
	ExpiredReward    int64         `json:"expiredReward"`   
	Received         int64         `json:"received"`        
	InfoShare        MeetingInfoShare     `json:"infoShare"`       
	LastUpdateTime   int64         `json:"lastUpdateTime"`  
	MFC              interface{}           `json:"mfc"`             // MISSING DATA
	CompleteWorkTime int64         `json:"completeWorkTime"`
}

type MeetingBoard struct {
	Rhodes     string `json:"RHODES"`    
	Glasgow    string `json:"GLASGOW"`   
	Kjerag     string `json:"KJERAG"`    
	Blacksteel string `json:"BLACKSTEEL"`
	Ursus      string `json:"URSUS"`     
}

type MeetingBuff struct {
	Speed  float64 `json:"speed"` 
	Weight MeetingWeight  `json:"weight"`
	Flag   interface{}     `json:"flag"`  // MISSING DATA
}

type MeetingWeight struct {
	Rhine      int64   `json:"RHINE"`     
	Penguin    int64   `json:"PENGUIN"`   
	Blacksteel int64   `json:"BLACKSTEEL"`
	Ursus      int64   `json:"URSUS"`     
	Glasgow    int64   `json:"GLASGOW"`   
	Kjerag     float64 `json:"KJERAG"`    
	Rhodes     int64   `json:"RHODES"`    
}

type MeetingDailyReward struct {
	ID      string `json:"id"`     
	Type    string `json:"type"`   
	Number  int64  `json:"number"` 
	Uid     string `json:"uid"`    
	Name    string `json:"name"`   
	NickNum string `json:"nickNum"`
	Chars   []MeetingChar `json:"chars"`  
	InUse   int64  `json:"inUse"`  
}

type MeetingChar struct {
	CharID      string `json:"charId"`     
	Level       int64  `json:"level"`      
	Skin        string `json:"skin"`       
	EvolvePhase int64  `json:"evolvePhase"`
}

type MeetingInfoShare struct {
	Ts     int64 `json:"ts"`    
	Reward int64 `json:"reward"`
}

type MeetingSocialReward struct {
	Daily  int64 `json:"daily"` 
	Search int64 `json:"search"`
}

type HireRoom struct {
	Buff             HireBuff  `json:"buff"`            
	State            int64 `json:"state"`           
	RefreshCount     int64 `json:"refreshCount"`    
	LastUpdateTime   int64 `json:"lastUpdateTime"`  
	ProcessPoint     int64 `json:"processPoint"`    
	Speed            int64 `json:"speed"`           
	CompleteWorkTime int64 `json:"completeWorkTime"`
}

type HireBuff struct {
	Speed   float64 `json:"speed"`  
	Meeting HireMeeting `json:"meeting"`
	Stack   HireStack   `json:"stack"`  
	Point   interface{}   `json:"point"`  // MISSING DATA
	ApCost  map[string]interface{}  `json:"apCost"` // MISSING DATA
}

type HireMeeting struct {
	SpeedUp int64 `json:"speedUp"`
}

type HireStack struct {
	Char       []HireChar `json:"char"`      
	ClueWeight interface{}  `json:"clueWeight"` // MISSING DATA
}

type HireChar struct {
	Refresh int64 `json:"refresh"`
}

type TrainingRoom struct {
	Buff           TrainingBuff    `json:"buff"`          
	State          int64   	`json:"state"`         
	LastUpdateTime int64   `json:"lastUpdateTime"`
	Trainee        TrainingTrainee `json:"trainee"`       
	Trainer        TrainingTrainer `json:"trainer"`       
}

type TrainingBuff struct {
	Speed  float64 		`json:"speed"` 
	LVEx   interface{}      `json:"lvEx"`   // MISSING DATA
	LVCost interface{}      `json:"lvCost"` // MISSING DATA
}

type TrainingTrainee struct {
	CharInstID   int64       `json:"charInstId"`  
	State        int64       `json:"state"`       
	TargetSkill  int64       `json:"targetSkill"` 
	ProcessPoint int64       `json:"processPoint"`
	Speed        float64     `json:"speed"`       
	CharTemplate interface{} `json:"charTemplate"`	// MISSING DATA
}

type TrainingTrainer struct {
	CharInstID int64 `json:"charInstId"`
	State      int64 `json:"state"`     
}
