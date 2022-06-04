package account

type OTPVerify struct {
	Email		string	`xorm:"pk varchar(32) not null 'Email'"`
	OTP			string	`xorm:"varchar(6) not null 'OTP'"`
	CurrentTime int64	`xorm:"bigint not null 'CurrentTime'"`
}

type GeneralAccount struct {
	Email			string `xorm:"pk varchar(32) not null 'Email'"`
	MasterUID		string `xorm:"varchar(32) not null 'MasterUID'"`
	MasterToken		string `xorm:"varchar(128) not null 'MasterToken'"`
}

type AuthSubmit struct {
	Result			int64  `json:"result"`
	YostarUID		string `json:"yostar_uid"`
	YostarToken		string `json:"yostar_token"`
	YostarAccount	string `json:"yostar_account"`
}