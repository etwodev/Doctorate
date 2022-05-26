package user


type StoryReview struct {
	Groups map[string]Storer `json:"groups"`
	Tags   Tags   `json:"tags"`  
}

type Story struct {
	ID  string `json:"id"` 
	Uts int64  `json:"uts"`
	RC  int64  `json:"rc"` 
}

type Storer struct {
	RTS          int64    `json:"rts"`         
	Stories      []Story  `json:"stories"`     
	TrailRewards []string `json:"trailRewards,omitempty"`
}

type Tags struct {
	KnownStoryAcceleration int64 `json:"knownStoryAcceleration"`
}
