package helpers

import (
	"bytes"
	"compress/gzip"
	"strings"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

type Pack struct {
	Name		string		`json:"name"`
	Hash		string		`json:"hash"`
	MD5			string		`json:"md5"`
	TotalSize	int			`json:"totalSize"`
	ABSize		int			`json:"abSize"`
	Type		string		`json:"type"`
	CID			int			`json:"cid"`
}

type AssetBundleData struct {
	FullPack	*Pack		`json:"fullPack"`
	Version		string		`json:"versionId"`
	ABInfo		[]Pack		`json:"abInfos"`
	CountRes	int			`json:"countOfTypedRes"`
	PackInfo	[]Pack		`json:"packInfos"`
}

type AssetBundleVersion struct {
	ClientVersion 			string  		`json:"clientVersion"`
	ResourceVersion			string 			`json:"resVersion"`
}

type AssetBundleConfig struct {
	// Version GET header keys for parsing
	VersionKeys		[]string
	// Version GET header values for parsing
	VersionValues	[]string
	// Version URL e.g. 'https://hostname/assetbundle/{device}/version'
	VersionURL		string
	// Asset GET header keys for parsing
	AssetKeys		[]string
	// Asset GET header values for parsing
	AssetValues		[]string
	// Asset URL e.g. 'https://hostname/assetbundle/{device}/{version}/assets/'
	AssetURL		string
	// Asset Directory e.g. './public/assets/{device}/{version}/'
	AssetDirectory	string
	// Asset List Path e.g. './public/assets/{device}/{version}/hot_update_list.json'
	AssetListPath	string
	// Asset List URL e.g. 'https://hostname/assetbundle/{device}/{version}/assets/host_update_list.json'
	AssetListURL	string
}

type HotUpdateManager struct {
	Data		*AssetBundleData
	Versions	*AssetBundleVersion
	Config		*AssetBundleConfig
}

func HotUpdater(url string, platform string) {
	pkg := AssetBundleData{}
	ver := AssetBundleVersion{}

	v_key := []string{"Connection", "User-Agent", "User-Agent", "User-Agent", "Accept", "Accept-Language", "Accept-Encoding", "X-Unity-Version"}
	v_val := []string{"Keep-Alive", "Arknights/31", "CFNetwork/1327.0.4", "Darwin/21.2.0", "*/*", "en-GB,en;q=0.9", "gzip, deflate", "2017.4.39f1"}
	v_url := fmt.Sprintf("%s/%s/version", url, platform)

	res := getData(v_url, v_key, v_val)
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(resp, &ver)
	if err != nil {
		panic(err)
	}

	a_key := []string{"Connection", "Te", "User-Agent", "Accept-Language", "Accept-Encoding"}
	a_val := []string{"Keep-Alive", "identity", "BestHTTP", "en-GB,en;q=0.9", "gzip, identity"}
	a_url := fmt.Sprintf("%s/%s/assets/%s/", url, platform, ver.ResourceVersion)
	a_dir := fmt.Sprintf("./public/%s/%s/", platform, ver.ResourceVersion)
	a_pat := a_dir + "hot_update_list.json"
	a_lit := a_url + "hot_update_list.json"

	conf := AssetBundleConfig{
		VersionKeys: v_key,
		VersionValues: v_val,
		VersionURL: v_url,
		AssetKeys: a_key,
		AssetValues: a_val,
		AssetURL: a_url,
		AssetDirectory: a_dir,
		AssetListPath: a_pat,
		AssetListURL: a_lit,
	}

	m := HotUpdateManager{
		Data: &pkg,
		Versions: &ver,
		Config: &conf,
	}

	if !m.genCheck() {
		return
	}
	
	genDir(m.Config.AssetDirectory)
	
	reader := parseData(getData(m.Config.AssetListURL, m.Config.AssetKeys, m.Config.AssetValues))
	genAltFile(m.Config.AssetListPath, reader)

	reader = parseData(getData(m.Config.AssetListURL, m.Config.AssetKeys, m.Config.AssetValues))

	err = json.NewDecoder(reader).Decode(&m.Data)
    if err != nil && err != io.EOF {
        panic(err)
    }


	for _, pack := range pkg.PackInfo {
		log.Debug().Msgf("Downloading: %s | Size: %s", pack.Name, pack.TotalSize)
		respon := getData(m.Config.AssetURL + pack.Name + ".dat", m.Config.AssetKeys, m.Config.AssetValues)
		genResponseFile(m.Config.AssetDirectory + pack.Name + ".dat", respon)
	}
	for _, pack := range pkg.ABInfo {
		log.Debug().Msgf("Downloading: %s | Size: %s", pack.Name, pack.TotalSize)
		fp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(pack.Name, "/", "_"), ".ab", ".dat"), "#", "__"), ".mp4", ".dat")
		respons := getData(m.Config.AssetURL + fp, m.Config.AssetKeys, m.Config.AssetValues)
		genResponseFile(m.Config.AssetDirectory + fp, respons)
	}
}

func (m *HotUpdateManager) genCheck() bool {
	file, err := ioutil.ReadFile(m.Config.AssetListPath)
	if err != nil {
		return true
	}

	pkg := AssetBundleData{}
	err = json.Unmarshal([]byte(file), &pkg)
	if err != nil {
		panic(err)
	}
	
	for _, pack := range pkg.PackInfo {
		// These packs DO NOT have checksums, so to save time, we skip verification and stick to file checks.
		
		file := pack.Name + ".dat"
		url := m.Config.AssetURL + file
		path := m.Config.AssetDirectory + file

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			log.Debug().Msgf("Downloading: %s | Size: %s", pack.Name, pack.TotalSize)
			respons := getData(url, m.Config.AssetKeys, m.Config.AssetValues)
			genResponseFile(path, respons)
		}
	}

	for _, pack := range pkg.ABInfo {
		fp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(pack.Name, "/", "_"), ".ab", ".dat"), "#", "__"), ".mp4", ".dat")
		m.genPredown(fp, &pack)
	}
	return false
}

func (m *HotUpdateManager) genPredown(file string, pack *Pack) {
	url := m.Config.AssetURL + file
	path := m.Config.AssetDirectory + file
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Debug().Msgf("Downloading: %s | Size: %s", pack.Name, pack.TotalSize)
		respons := getData(url, m.Config.AssetKeys, m.Config.AssetValues)
		genResponseFile(path, respons)
	} else {
		if pack.TotalSize == int(info.Size()) {
			log.Debug().Msgf("%s Already downloaded!", pack.Name)
			return
		} else {
			log.Debug().Msgf("Downloading: %s | Size: %s", pack.Name, pack.TotalSize)
			respons := getData(url, m.Config.AssetKeys, m.Config.AssetValues)
			genResponseFile(path, respons)
		}
		
	}
}


func genDir(path string) {
	err := os.MkdirAll(path, 0755)
    if err != nil {
        panic(err)
    }
}

func getData(url string, keys []string, values []string) *http.Response {
	c := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	for i, key := range keys {
		req.Header.Add(key, values[i])
	}

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	return res
}

func parseData(res *http.Response) *gzip.Reader {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(body)
    reader, err := gzip.NewReader(buf)
    if err != nil {
        panic(err)
    }
	
    return reader

}

func genAltFile(path string, reader *gzip.Reader) {

	out, err := os.Create(path)
	if err != nil  {
	  panic(err)
	}

	_, err = io.Copy(out, reader)
	if err != nil  {
	  panic(err)
	}
}

func genResponseFile(path string, res *http.Response) {
	out, err := os.Create(path)
	if err != nil  {
	  panic(err)
	}
	defer out.Close()
  
	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}
  
	_, err = io.Copy(out, res.Body)
	if err != nil  {
	  panic(err)
	}
}