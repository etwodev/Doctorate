package types

type Version struct {
	Client 			string  		`json:"clientVersion"`
	Resource		string 			`json:"resVersion"`
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