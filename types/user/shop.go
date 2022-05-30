package user

type Shop struct {
	Ls     Ls     `json:"LS"`    
	Hs     Hs     `json:"HS"`    
	Es     Es     `json:"ES"`    
	Cash   Cash   `json:"CASH"`  
	Gp     Gp     `json:"GP"`    
	Furni  Furni  `json:"FURNI"` 
	Social ShopSocial `json:"SOCIAL"`
	Epgs   Cash   `json:"EPGS"`  
	Lmtgs  Cash   `json:"LMTGS"` 
}

type Cash struct {
	Info []ShopInfo `json:"info"`
}

type ShopInfo struct {
	ID    string `json:"id"`   
	Count int64  `json:"count"`
}

type Es struct {
	CurShopID string `json:"curShopId"`
	Info      []ShopInfo `json:"info"`     
	LastClick int64  `json:"lastClick"`
}

type Furni struct {
	Info      []ShopInfo           `json:"info"`     
	GroupInfo map[string]int64 `json:"groupInfo"`
}

type Gp struct {
	OneTime Cash    `json:"oneTime"`
	Level   Cash    `json:"level"`  
	Weekly  Monthly `json:"weekly"` 
	Monthly Monthly `json:"monthly"`
}

type Monthly struct {
	CurGroupID string	`json:"curGroupId"`
	Info       []ShopInfo	`json:"info"`      
}

type Hs struct {
	CurShopID    string                  `json:"curShopId"`   
	Info         []ShopInfo         		 `json:"info"`        
	ProgressInfo map[string]ProgressInfo `json:"progressInfo"`
}

type ProgressInfo struct {
	Count int64 `json:"count"`
	Order int64 `json:"order"`
}

type Ls struct {
	CurShopID  string `json:"curShopId"` 
	CurGroupID string `json:"curGroupId"`
	Info       []ShopInfo `json:"info"`      
}

type ShopSocial struct {
	CurShopID    string        	 `json:"curShopId"`   
	Info         []ShopInfo 		 `json:"info"`        
	CharPurchase map[string]int  `json:"charPurchase"`
}

