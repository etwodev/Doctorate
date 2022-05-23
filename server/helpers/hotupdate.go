package helpers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

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
	// Base platform string
	BasePlatform	string
	// Base URL string
	BaseURL			string
	// Version GET headers
	VersionHeaders	[][2]string
	// Version URL e.g. 'https://hostname/assetbundle/{device}/version'
	VersionURL		string
	// Asset GET headers
	AssetHeaders	[][2]string
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


func HotUpdater(url string, platform string) error {
	ver := AssetBundleVersion{}
	cfg := genConf(url, platform)
	
	response, err := GetURLData(cfg.VersionURL, cfg.VersionHeaders)
	if err != nil {
		return fmt.Errorf("HotUpdater: failed get response: %w", err)
	}

	err = json.Unmarshal(response, &ver)
	if err != nil {
		return fmt.Errorf("HotUpdater: failed get unmarshalling: %w", err)
	}

	cfg.AssetURL = fmt.Sprintf("%s/%s/assets/%s/", cfg.BaseURL, cfg.BasePlatform, ver.ResourceVersion)
	cfg.AssetDirectory = fmt.Sprintf("./static/hotupdate/%s/%s/", cfg.BasePlatform, ver.ResourceVersion)
	cfg.AssetListPath = cfg.AssetDirectory + "hot_update_list.json"
	cfg.AssetListURL = cfg.AssetURL + "hot_update_list.json"

	m := genManager(&ver, cfg)
	logic, err := m.genCheck()
	if err != nil {
		return fmt.Errorf("HotUpdater: failed to get logic: %w", err)
	}

	if !logic {
		return nil
	} else {
		err := os.MkdirAll(cfg.AssetDirectory, 0755)
		if err != nil {
			return fmt.Errorf("HotUpdater: failed to create directory: %w", err)
		}
		response, err := GetURLData(cfg.AssetListURL, cfg.AssetHeaders)
		if err != nil {
			return fmt.Errorf("HotUpdater: failed to get response: %w", err)
		}

		buffer := bytes.NewBuffer(response)
		reader, err := gzip.NewReader(buffer)
		if err != nil {
			return fmt.Errorf("HotUpdater: failed to read gzip: %w", err)
		}

		err = genGZIP(m.Config.AssetListPath, reader)
		if err != nil {
			return fmt.Errorf("HotUpdater: failed to generate gzip: %w", err)
		}
		
		logic, err := m.genCheck()
		if err != nil {
			return fmt.Errorf("HotUpdater: failed to get logic on second iteration: %w", err)
		}

		if !logic {
			return nil
		} else {
			return fmt.Errorf("HotUpdater: second iteration failed: %w", err)
		}
	}
}

func (m *HotUpdateManager) genCheck() (bool, error) {
	pkg := AssetBundleData{}
	bin, err := ioutil.ReadFile(m.Config.AssetListPath)
	if err != nil {
		return true, nil
	}

	err = json.Unmarshal(bin, &pkg)
	if err != nil {
		return true, fmt.Errorf("genCheck: failed unmarshalling: %w", err)
	}
	
	for _, pack := range pkg.PackInfo {	
		file := pack.Name + ".dat"
		url := m.Config.AssetURL + file
		path := m.Config.AssetDirectory + file
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			log.Info().Str("Name", pack.Name).Int("Size", pack.TotalSize).Msg("Downloading asset")
			response, err := GetURLData(url, m.Config.AssetHeaders)
			if err != nil {
				return true, fmt.Errorf("genCheck: failed downloading data: %w", err)
			}
			err = genFile(path, response)
			if err != nil {
				return true, fmt.Errorf("genCheck: failed generating data: %w", err)
			}
		}
	}

	for _, pack := range pkg.ABInfo {
		fp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(pack.Name, "/", "_"), ".ab", ".dat"), "#", "__"), ".mp4", ".dat")
		m.genDown(fp, &pack)
	}
	return false, nil
}

func (m *HotUpdateManager) genDown(file string, pack *Pack) (error) {
	url := m.Config.AssetURL + file
	path := m.Config.AssetDirectory + file
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Info().Str("Name", pack.Name).Int("Size", pack.TotalSize).Msg("Downloading asset")
		response, err := GetURLData(url, m.Config.AssetHeaders)
		if err != nil {
			return fmt.Errorf("genCheck: failed downloading data: %w", err)
		}		
		genFile(path, response)
	} else {
		if pack.TotalSize == int(info.Size()) {
			return nil
		} else {
			log.Info().Str("Name", pack.Name).Int("Size", pack.TotalSize).Msg("Downloading asset")
			response, err := GetURLData(url, m.Config.AssetHeaders)
			if err != nil {
				return fmt.Errorf("genCheck: failed downloading data: %w", err)
			}		
			genFile(path, response)
		}
	}
	return nil
}

func genGZIP(path string, reader *gzip.Reader) error {
	out, err := os.Create(path)
	if err != nil  {
	  return fmt.Errorf("genFile: failed creating file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, reader)
	if err != nil  {
	  panic(err)
	}
	return nil
}

func genFile(path string, data []byte) error {
	out, err := os.Create(path)
	if err != nil  {
	  return fmt.Errorf("genFile: failed creating file: %w", err)
	}
	defer out.Close()
  
	_, err = out.Write(data)
	if err != nil {
		return fmt.Errorf("genFile: failed writing file: %w", err)
	}
	return nil
}

func genConf(url string, platform string) *AssetBundleConfig {
	return &AssetBundleConfig{
		BasePlatform: platform,
		BaseURL: url,
		VersionHeaders: [][2]string{
			{"Connection", "Keep-Alive"},
			{"User-Agent", "Arknights/31"},
			{"User-Agent", "CFNetwork/1327.0.4"},
			{"User-Agent", "Darwin/21.2.0"},
			{"Accept", "*/*"},
			{"Accept-Language", "en-GB,en;q=0.9"},
			{"Accept-Encoding", "gzip, deflate"},
			{"X-Unity-Version", "2017.4.39f1"},
		},
		VersionURL: fmt.Sprintf("%s/%s/version", url, platform),
		AssetHeaders: [][2]string{
			{"Connection", "Keep-Alive"},
			{"Te", "identity"},
			{"User-Agent", "BestHTTP"},
			{"Accept-Language", "en-GB,en;q=0.9"},
			{"Accept-Encoding", "gzip, identity"},
		},
	}
}

func genManager(ver *AssetBundleVersion, cfg *AssetBundleConfig) *HotUpdateManager {
	return &HotUpdateManager{
		Versions: ver,
		Config: cfg,
	}
}