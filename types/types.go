package types

type Version struct {
	Client 			string  		`json:"clientVersion"`
	Resource		string 			`json:"resVersion"`
}

type Preannouncement struct {
	Actived         bool   `json:"actived"`        
	PreAnnounceID   int64  `json:"preAnnounceId"`  
	PreAnnounceType int64  `json:"preAnnounceType"`
	PreAnnounceURL  string `json:"preAnnounceUrl"` 
}

type NetworkConfig struct {
	Sign		string 		`json:"sign"`
	Content		string		`json:"content"`
}

type NetworkContent struct {
	ConfigVersion 	string
	FunctionVersion string					
	Configs			map[string]Networker
}

type Networker struct {
	Override		bool     `json:"override"`
	Network			Networks `json:"network"` 
}

type Networks struct {
	Gs     string `json:"gs"`    
	As     string `json:"as"`    
	U8     string `json:"u8"`    
	Hu     string `json:"hu"`    
	Hv     string `json:"hv"`    
	RC     string `json:"rc"`    
	An     string `json:"an"`    
	Prean  string `json:"prean"` 
	Sl     string `json:"sl"`    
	Of     string `json:"of"`    
	PkgAd  string `json:"pkgAd"` 
	PkgIOS string `json:"pkgIOS"`
	Secure bool   `json:"secure"`
}

type Payload struct {
	FullPack        FullPack `json:"fullPack"`       
	VersionID       string   `json:"versionId"`      
	ABInfo          []Pack   `json:"abInfos"`        
	Resources		int64    `json:"countOfTypedRes"`
	PackInfo	    []Pack   `json:"packInfos"`      
}

type Pack struct {
	Name      string `json:"name"`          
	Hash      string `json:"hash"`          
	MD5       string `json:"md5"`           
	TotalSize int64  `json:"totalSize"`     
	ABSize    int64  `json:"abSize"`        
	CID       int64  `json:"cid"`           
	PID       *PID   `json:"pid,omitempty"` 
	Type      *Type  `json:"type,omitempty"`
}

type FullPack struct {
	TotalSize int64  `json:"totalSize"`
	ABSize    int64  `json:"abSize"`   
	Type      string `json:"type"`     
	CID       int64  `json:"cid"`      
}

type PID string
const (
	LpackChar PID = "lpack_char"
	LpackFurn PID = "lpack_furn"
	LpackInit PID = "lpack_init"
	LpackLart PID = "lpack_lart"
	LpackLcom PID = "lpack_lcom"
	LpackLmesh PID = "lpack_lmesh"
	LpackMisc PID = "lpack_misc"
	LpackMusic PID = "lpack_music"
	LpackScene PID = "lpack_scene"
	LpackV023 PID = "lpack_v023"
	LpackVoice PID = "lpack_voice"
	LpackVoicn PID = "lpack_voicn"
)

type Type string
const (
	DynIllust Type = "dyn_illust"
	Video Type = "video"
	Voice Type = "voice"
)

type Settings struct {
	Settings	Setting   `json:"settings"`
}

type Setting struct {
	AppAndroidkey          string              `json:"APP_ANDROIDKEY"`         
	AdjustEnabled          int64               `json:"ADJUST_ENABLED"`         
	AdjustIsdebug          int64               `json:"ADJUST_ISDEBUG"`         
	AdjustAppid            string              `json:"ADJUST_APPID"`           
	AdjustChargeeventtoken string              `json:"ADJUST_CHARGEEVENTTOKEN"`
	AppGl                  string              `json:"app_gl"`                 
	UserAgreement          UserAgreement       `json:"USER_AGREEMENT"`         
	AppAdjustInitScript    AppAdjustInitScript `json:"app_adjust_init_script"` 
	AppFire                int64               `json:"app_fire"`               
	AppDebug               int64               `json:"app_debug"`              
	UserDestroyDays        int64               `json:"UserDestroyDays"`        
	AihelpMode             int64               `json:"AIHELP_MODE"`            
	AdjustEventtokens      map[string]string   `json:"ADJUST_EVENTTOKENS"`     
	TwitterKey             string              `json:"TWITTER_KEY"`            
	TwitterSecret          string              `json:"TWITTER_SECRET"`         
	FacebookAppid          string              `json:"FACEBOOK_APPID"`         
	FacebookClienttoken    string              `json:"FACEBOOK_CLIENTTOKEN"`   
	AppBirthsetEnabled     int64               `json:"APP_BIRTHSET_ENABLED"`   
	RemoteConfig           []RemoteConfig      `json:"REMOTE_CONFIG"`          
}

type AppAdjustInitScript struct {
	Android Android `json:"android"`
	Ios     Android `json:"ios"`    
}

type Android struct {
	SecretID string `json:"secretId"`
	Info1    string `json:"info1"`   
	Info2    string `json:"info2"`   
	Info3    string `json:"info3"`   
	Info4    string `json:"info4"`   
}

type RemoteConfig struct {
	ConfigKey string `json:"config_key"`
	EventName string `json:"event_name"`
}

type UserAgreement struct {
	Latest Latest `json:"LATEST"`
}

type Latest struct {
	Version string `json:"version"`
}

type Code struct {
	Number			string		`json:"codestr"`
	Message			string		`json:"codemessage"`
}

type Codes struct {
	Result				int 			`json:"result"`
	Data				[]Code			`json:"data"`
}

type Agreements struct {
	Version			string			`json:"version"`
	Data			[]string		`json:"data"`
}