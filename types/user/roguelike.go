package user

type Roguelike struct {
	Current interface{} `json:"current"` // MISSING DATA
	Stable  Stable  `json:"stable"` 
}

type Stable struct {
	OutBuff interface{}                 `json:"outBuff"` // MISSING DATA
	Relic   map[string]map[string]int64 `json:"relic"`
	Stages  interface{}                 `json:"stages"`  // MISSING DATA
	Ending  interface{}                 `json:"ending"`  // MISSING DATA
	Mode    interface{}                 `json:"mode"`    // MISSING DATA
	Stats   interface{}                 `json:"stats"`   // MISSING DATA
}

type Mode struct {
	Easy      map[string]int64 `json:"easy"`       // MISSING DATA
	Normal    map[string]int64 `json:"normal"`     // MISSING DATA
	Difficult map[string]int64 `json:"difficult"`  // MISSING DATA
}
