package user

type Recruit struct {
	Normal map[string]map[string]Slots `json:"normal"`
}

type Slots struct {
	State         int64         `json:"state"`        
	Tags          []int64       `json:"tags"`         
	SelectTags    []string 		`json:"selectTags"`   	// MISSING DATA | POSSIBLE ERROR
	StartTs       int64         `json:"startTs"`      
	DurationInSEC int64         `json:"durationInSec"`
	MaxFinishTs   int64         `json:"maxFinishTs"`  
	RealFinishTs  int64         `json:"realFinishTs"` 
}
