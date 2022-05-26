package user

type Avatar struct {
	AvatarIcon map[string]AvatarIcon `json:"avatar_icon"`
}

type AvatarIcon struct {
	Ts  int64  `json:"ts"` 
	Src string `json:"src"`
}
