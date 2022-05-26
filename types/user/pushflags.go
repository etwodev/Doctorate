package user

type PushFlags struct {
	HasGifts         int64 `json:"hasGifts"`        
	HasFriendRequest int64 `json:"hasFriendRequest"`
	HasClues         int64 `json:"hasClues"`        
	HasFreeLevelGP   int64 `json:"hasFreeLevelGP"`  
	Status           int64 `json:"status"`          
}