package static

import (
	"github.com/Etwodev/Doctorate/types"
)



const AssetBundle string = "https://ark-us-static-online.yo-star.com/assetbundle/official"
const AssetBundleAsset string = "https://ark-us-static-online.yo-star.com/assetbundle/official/%s/assets/%s/%s"
const AssetBundleVersion string = "https://ark-us-static-online.yo-star.com/assetbundle/official/%s/version"
const AssetBundleHotUpdate string = "https://ark-us-static-online.yo-star.com/assetbundle/official/%s/assets/%s/hot_update_list.json"

const AssetBundleDirectory string = "./static/hotupdate/%s/"
const AssetBundleDirectoryExpired string = "./static/hotupdate/%s/expired/%s/"
const AssetBundleDirectoryExpiredHotUpdate string = "./static/hotupdate/%s/expired/%s/hot_update_list.json"
const AssetBundleDirectoryHotUpdate string = "./static/hotupdate/%s/hot_update_list.json"

const ClientVersion string = "9.0.01"

var AssetBundleVersionHeaders = [...][2]string{{"Connection", "Keep-Alive"},{"User-Agent", "Arknights/31"},{"User-Agent", "CFNetwork/1327.0.4"},{"User-Agent", "Darwin/21.2.0"},{"Accept", "*/*"},{"Accept-Language", "en-GB,en;q=0.9"},{"Accept-Encoding", "gzip, deflate"},{"X-Unity-Version", "2017.4.39f1"},}
var AssetBundleHotUpdateHeaders = [...][2]string{{"Connection", "Keep-Alive"},{"Te", "identity"},{"User-Agent", "BestHTTP"},{"Accept-Language", "en-GB,en;q=0.9"},{"Accept-Encoding", "gzip, identity"},}

const OTP string = "1234567890"

const VerifyEmailSubject string = "Verify your email!"
const VerifyEmailMime string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
const VerifyEmailCode string = "MISSING_CODE_ERROR"
const VerifyEmail string = `<!doctype html><html><head><style>html{width:100%;min-height:100%}body{width:100%;min-height:100%;position:absolute;background-color:rgba(0,0,0,.8)}.container{padding:0 2em 1.5em;color:#fff;top:0;width:100%;min-height:100%;box-sizing:border-box}.container .content{font-size:.8rem;line-height:1.2rem;text-align:justify}</style></head><body><div class="container"><h2 class="head-title">Verification</h2><div class="content">A person has attempted to create or signup to a doctorate server, using this email.<br/>If this was you, please input the code below into the verification box, it will expire in 30 minutes.<br/><br/><b>MISSING_CODE_ERROR</b></div></div></body></html>`

// PRELOAD

var NetworkConfig types.NetworkConfig
var Preannouncement types.Preannouncement
var Settings types.Settings
var Codes []types.Code
var Agreements types.Agreements

var EmailHost string
var EmailIP string
var EmailPassword string
var EmailAddress string

var IOS_VERSION types.Version
var ANDROID_VERSION types.Version

