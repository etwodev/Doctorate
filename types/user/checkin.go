package user

type CheckIn struct {
	CanCheckIn         int64   `json:"canCheckIn"`        
	CheckInGroupID     string  `json:"checkInGroupId"`    
	CheckInRewardIndex int64   `json:"checkInRewardIndex"`
	CheckInHistory     []int64 `json:"checkInHistory"`    
}
