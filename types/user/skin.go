package user

type Skin struct {
	CharacterSkins map[string]int64 `json:"characterSkins"`
	SkinTs         map[string]int64 `json:"skinTs"`        
}
