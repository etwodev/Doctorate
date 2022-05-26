package user

type TShop struct {
	Coin         int64                   `json:"coin"`        
	Info         []Info                  `json:"info"`        
	ProgressInfo map[string]CharProgress `json:"progressInfo"`
}

type Info struct {
	ID    string `json:"id"`   
	Count int64  `json:"count"`
}

type CharProgress struct {
	Count int64 `json:"count"`
	Order int64 `json:"order"`
}
