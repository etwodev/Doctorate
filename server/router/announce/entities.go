package announce


type Announcers struct {
	Announcements	    []Announcement `json:"announceList"`
	Extra		   		*Xtrannounce
	Focus  				string      `json:"focusAnnounceId"`
}

type Announcement struct {
	AnnoucementID		string	 	`json:"announceId"`
	Day					int 		`json:"day"`
	Group				string		`json:"group"`
	IsWeb				bool		`json:"isWebUrl"`
	Month				int			`json:"month"`
	Title				string		`json:"title"`
	URL					string		`json:"webUrl"`
}

type Xtrannounce struct {
	Enable				bool		`json:"enable"`
	Name				string		`json:"name"`
}

type Preannouncement struct {
	Activated			bool	 	`json:"actived"`
	PreannouncementID	string 		`json:"preAnnounceId"`
	Type				int			`json:"preAnnounceType"`
	URL					string		`json:"preAnnounceUrl"`
}

type Version struct {
	Client 		        string		`json:"clientVersion"`
	Asset				string		`json:"resVersion"`
}