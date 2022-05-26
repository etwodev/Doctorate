package user

type Building struct {
	Status             Status                       `json:"status"`   
	Chars              map[string]CharValue         `json:"chars"`  
	RoomSlots          map[string]RoomSlot          `json:"roomSlots"`
	Rooms              Rooms					    `json:"rooms"`  
	Furniture          map[string]Furniture         `json:"furniture"`
	DiyPresetSolutions map[string]DiyPresetSolution `json:"diyPresetSolutions"` 
	Assist             []float64                      `json:"assist"`
}

type Status struct {
	Labor Labor `json:"labor"`
}

type Labor struct {
	BuffSpeed      float64 `json:"buffSpeed"`     
	ProcessPoint   float64 `json:"processPoint"`  
	Value          float64 `json:"value"`         
	LastUpdateTime float64 `json:"lastUpdateTime"`
	MaxValue       float64 `json:"maxValue"`      
}

type CharValue struct {
	CharID        string `json:"charId"`       
	LastApAddTime float64  `json:"lastApAddTime"`
	Ap            float64  `json:"ap"`           
	RoomSlotID    string `json:"roomSlotId"`   
	Index         float64  `json:"index"`        
	ChangeScale   float64  `json:"changeScale"`  
	Bubble        Bubble `json:"bubble"`       
	WorkTime      float64  `json:"workTime"`     
}

type Bubble struct {
	Normal Assist `json:"normal"`
	Assist Assist `json:"assist"`
}

type Assist struct {
	Add float64 `json:"add"`
	Ts  float64 `json:"ts"` 
}

type DiyPresetSolution struct {
	Name      string   `json:"name"`     
	Solution  Solution `json:"solution"` 
	Thumbnail string   `json:"thumbnail"`
}

type Solution struct {
	WallPaper string   `json:"wallPaper"`
	Floor     string   `json:"floor"`    
	Carpet    []Carpet `json:"carpet"`   
	Other     []Carpet `json:"other"`    
}

type Carpet struct {
	ID         string     `json:"id"`        
	Coordinate Coordinate `json:"coordinate"`
}

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Furniture struct {
	Count float64 `json:"count"`
	InUse float64 `json:"inUse"`
}

type RoomSlot struct {
	Level                 float64   `json:"level"`                
	State                 float64   `json:"state"`                
	RoomID                string  `json:"roomId"`               
	CharInstIDS           []float64 `json:"charInstIds"`          
	CompleteConstructTime float64   `json:"completeConstructTime"`
}