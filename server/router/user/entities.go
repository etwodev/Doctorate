package user

type DeviceAccount struct {
	MasterUID		string `xorm:"pk varchar(32) not null 'MasterUID'"`
	DeviceUID		string `xorm:"pk varchar(32) not null 'DeviceUID'"`
	DeviceToken		string `xorm:"pk varchar(32) not null 'DeviceToken'"`
	Channel			string `xorm:"pk varchar(32) not null 'Channel'"`
}