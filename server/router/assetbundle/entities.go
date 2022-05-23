package assetbundle

// JSON

type Versions struct {
	IOS				*AssetBundleVersion		`json:"IOS"`
	Android			*AssetBundleVersion		`json:"Android"`
}

type AssetBundleVersion struct {
	ClientVersion 			string  		`json:"clientVersion"`
	ResourceVersion			string 			`json:"resVersion"`
}