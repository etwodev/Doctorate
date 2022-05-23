package app

type Code struct {
	Number			string		`json:"codestr"`
	Message			string		`json:"codemessage"`
}

type Codes struct {
	Result				int 			`json:"result"`
	Data				[]Code			`json:"data"`
}
