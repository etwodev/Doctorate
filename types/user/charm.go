package user

type Charm struct {
	Charms map[string]int64  `json:"charms"`
	Squad  []string 		 `json:"squad"` 
}