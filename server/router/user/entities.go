package user

type AuthData struct {
	DeviceID		string					`xorm:"'DeviceID'"`
	StoreID			string					`xorm:"'StoreID'"`
	TimeNow			int64					`xorm:"'TimeNow'"`
	EntryToken		string					`xorm:"'EntryToken'"`
	AccessToken		string					`xorm:"'AccessToken'"`
	CryptoToken		string					`xorm:"'CryptoToken'"`
	AccessUID		string					`xorm:"'AccessUID'"`
	SnowflakeUID	string					`xorm:"'SnowflakeUID'"`
} 

type AuthyJSON struct {
	Result			*int				  `json:"result,omitempty"`
	UID				int64				  `json:"uid"`
	Token			string			  	  `json:"token"`
	IsNew			*int				  `json:"isNew,omitempty"`
}