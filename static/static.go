package static

const OTP string = "1234567890"

const IOS string = "IOS"
const Android string = "Android"

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