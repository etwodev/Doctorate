package user

type Welcome struct {
	Crisis Crisis `json:"crisis"`
}

type Crisis struct {
	Current  string         	`json:"current"` 
	Lst      int64          	`json:"lst"`     
	Nst      int64          	`json:"nst"`     
	Map      map[string]Map 	`json:"map"`     
	Shop     CrisisShop          `json:"shop"`    
	Training CrisisTraining     `json:"training"`
	Season   map[string]Season  `json:"season"`  
	Box      []interface{}  	`json:"box"`    // MISSING DATA  
}

type Map struct {
	Rank      int64 `json:"rank"`     
	Confirmed int64 `json:"confirmed"`
}

type CrisisShop struct {
	Coin         int64         `json:"coin"`        
	Info         []interface{} `json:"info"`         // MISSING DATA
	ProgressInfo interface{}   `json:"progressInfo"` // MISSING DATA
}

type CrisisTraining struct {
	CurrentStage []string         		`json:"currentStage"`
	Stage        map[string]CrisisStage `json:"stage"`       
	Nst          int64            		`json:"nst"`         
}

type CrisisStage struct {
	Point int64 `json:"point"`
}

type Season struct {
	Coin 		int64	  `json:"coin"`
	TCoin 		int64	  `json:"tCoin"`
	Permanent	Permanent `json:"permanent"`
}

type Permanent struct {
	Nst		  int64				`json:"nst"`
	Rune	  map[string]int64	`json:"rune"`
	Point	  int64				`json:"point"`
	Challenge Challenge			`json:"challenge"`
	Temporary Temporary			`json:"temporary"`
	SInfo	  SInfo				`json:"sInfo"`
}

type Challenge struct {
	TaskList	map[string]Task		`json:"taskList"`
	TopPoint	int64				`json:"topPoint"`
	PointList	map[string]int64	`json:"pointList"`
}

type Task struct {
	Fts	int64 `json:"fts"`
	Rts int64 `json:"rts"`
}

type Temporary struct {
	Schedule	string
	Nst			int64
	Point		int64
	Challenge	Challenge
}

type SInfo struct {
	AssistCount	int64			`json:"assistCnt"`
	MaxPoint	int64			`json:"maxPnt"`
	Characters	[]interface{}	`json:"chars"`   // MISSING DATA
	History		interface{}		`json:"history"`  // MISSING DATA
}