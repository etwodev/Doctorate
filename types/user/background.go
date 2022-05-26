package user

type Background struct {
	Selected string 			`json:"selected"`
	Bgs      map[string]States    `json:"bgs"`     
}

type States struct {
	Unlock	int64						`json:"unlock,omitempty"`
	Alt		map[string]Alt	`json:"conditions,omitempty"`
}

type Alt struct {
	V int64 `json:"v"`
	T int64 `json:"t"`
}
