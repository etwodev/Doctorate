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

	stat "github.com/Etwodev/Doctorate/static"
	"github.com/Etwodev/Doctorate/types"
	"github.com/rs/zerolog/log"
)

func Updater(platform string) (error) {
	err := Handler(platform)
	if err != nil {
		return fmt.Errorf("Updater: failed update: %w", err)
	}
	return nil
}


func Handler(platform string) (error) {
	var ver types.Version
	var old types.Payload
	var new types.Payload

	log.Info().Str("Platform", platform).Msg("Creating directory")
	err := os.MkdirAll("./static/hotupdate/" + platform, 0755)
	if err != nil {
		return fmt.Errorf("Handler: failed to create directory: %w", err)
	}
	
	response, err := GetURLData(fmt.Sprintf(stat.AssetBundleVersion, platform), stat.AssetBundleVersionHeaders[:])
	if err != nil {
		return fmt.Errorf("Handler: failed get response: %w", err)
	}

	err = json.Unmarshal(response, &ver)
	if err != nil {
		return fmt.Errorf("Handler: failed get unmarshalling: %w", err)
	}

	info, err := os.Stat(fmt.Sprintf(stat.AssetBundleDirectoryHotUpdate, platform))
	if err != nil {
		if os.IsNotExist(err) {
			log.Info().Str("Platform", platform).Str("Name", fmt.Sprintf(stat.AssetBundleDirectoryHotUpdate, platform)).Msg("Downloading list file")
			err := Creator(stat.AssetBundleHotUpdate, stat.AssetBundleDirectoryHotUpdate, platform, ver.Resource)
			if err != nil {
				return fmt.Errorf("Handler: failed creating data: %w", err)
			}
			
			err = Marshaller(fmt.Sprintf(stat.AssetBundleDirectoryHotUpdate, platform), &new)
			if err != nil {
				return fmt.Errorf("Handler: failed unmarshalling data: %w", err)
			}

			err = General(new, platform, ver.Resource)
			if err != nil {
				return fmt.Errorf("Handler: failed update: %w", err)
			}
			return nil
		} else {
			return fmt.Errorf("Handler: failed checking data: %w", err)
		}
	}
	
	err = Marshaller(fmt.Sprintf(stat.AssetBundleDirectoryHotUpdate, platform), &old)
	if err != nil {
		return fmt.Errorf("Handler: failed unmarshalling data: %w", err)
	}

	if old.VersionID == ver.Resource {
		err := General(old, platform, ver.Resource)
		if err != nil {
			return fmt.Errorf("Handler: failed update: %w", err)
		}
		return nil
	}

	log.Info().Str("Platform", platform).Str("Name", info.Name()).Int64("Size", info.Size()).Msg("Creating expired directory")
	err = os.MkdirAll(fmt.Sprintf(stat.AssetBundleDirectoryExpired, platform, old.VersionID), 0755)
	if err != nil {
		return fmt.Errorf("Handler: failed to create expired directory: %w", err)
	}

	err = Mover(fmt.Sprintf(stat.AssetBundleDirectoryHotUpdate, platform), fmt.Sprintf(stat.AssetBundleDirectoryExpiredHotUpdate, platform, old.VersionID))
	if err != nil {
		return fmt.Errorf("Handler: failed copy: %w", err)
	}
	
	err = Creator(stat.AssetBundleHotUpdate, stat.AssetBundleDirectoryHotUpdate, platform, ver.Resource)
	if err != nil {
		return fmt.Errorf("Handler: failed creating data: %w", err)
	}
	
	err = Marshaller(fmt.Sprintf(stat.AssetBundleDirectoryHotUpdate, platform), &new)
	if err != nil {
		return fmt.Errorf("Handler: failed unmarshalling data: %w", err)
	}

	l, err := Checker(old, new)
	if err != nil {
		return fmt.Errorf("Handler: failed checking data: %w", err)
	}

	for _, x := range l {
		_, err = os.Stat(fmt.Sprintf(stat.AssetBundleDirectory, platform) + x)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return fmt.Errorf("Handler: failed checking data: %w", err)
			}
		} else {
			err = Mover(fmt.Sprintf(stat.AssetBundleDirectory, platform) + x, fmt.Sprintf(stat.AssetBundleDirectoryExpired, platform, old.VersionID) + x)
			if err != nil {
				return fmt.Errorf("Handler: failed moving file: %w", err)
			}
		}
	}

	err = General(new, platform, ver.Resource)
	if err != nil {
		return fmt.Errorf("Handler: failed downloading file: %w", err)
	}
	return nil
}

func Mover(old string, new string) (error) {
	err := os.Rename(old, new)
	if err != nil {
		return fmt.Errorf("Mover: failed renaming file path: %w", err)
	}
	return nil
}

func General(payload types.Payload, platform string, version string) (error) {
	for _, x := range payload.PackInfo {
		err := Downloader(x.Name + ".dat", version, platform, x)
		if err != nil {
			return fmt.Errorf("General: failed downloading file: %w", err)
		}
		continue
	}
	for _, x := range payload.ABInfo {
		err := Downloader(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(x.Name, "/", "_"), ".ab", ".dat"), "#", "__"), ".mp4", ".dat"), version, platform, x)
		if err != nil {
			return fmt.Errorf("General: failed downloading file: %w", err)
		}
		continue
	}
	return nil
}

func Downloader(file string, version string, platform string, x types.Pack) (error) {
	info, err := os.Stat(fmt.Sprintf(stat.AssetBundleDirectory, platform) + file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info().Str("Platform", platform).Str("Name", x.Name).Int64("Size", x.TotalSize).Msg("File doesn't exist, downloading asset")
			response, err := GetURLData(fmt.Sprintf(stat.AssetBundleAsset, platform, version, file), stat.AssetBundleHotUpdateHeaders[:])
			if err != nil {
				return fmt.Errorf("Downloader: failed downloading file: %w", err)
			}
			err = GenerateFile(fmt.Sprintf(stat.AssetBundleDirectory, platform) + file, response)
			if err != nil {
				return fmt.Errorf("Downloader: failed generating file: %w", err)
			}
		} else {
			return fmt.Errorf("Downloader: failed checking file: %w", err)
		}
	}  else {
		if info.Size() == x.TotalSize {
			return nil
		} else {
			log.Info().Str("Platform", platform).Str("Name", x.Name).Int64("Size", x.TotalSize).Msg("File doesn't exist, downloading asset")
			response, err := GetURLData(fmt.Sprintf(stat.AssetBundleAsset, platform, version, file), stat.AssetBundleHotUpdateHeaders[:])
			if err != nil {
				return fmt.Errorf("Downloader: failed downloading file: %w", err)
			}
			err = GenerateFile(fmt.Sprintf(stat.AssetBundleDirectory, platform) + file, response)
			if err != nil {
				return fmt.Errorf("Downloader: failed generating file: %w", err)
			}
		}
	}
	return nil
}

func GenerateFile(path string, data []byte) error {
	out, err := os.Create(path)
	if err != nil  {
	  return fmt.Errorf("GenerateFile: failed creating file: %w", err)
	}
	defer out.Close()
  
	_, err = out.Write(data)
	if err != nil {
		return fmt.Errorf("GenerateFile: failed writing file: %w", err)
	}
	return nil
}

func Checker(old types.Payload, new types.Payload) ([]string, error) {
	var l []string
	for _, x := range old.PackInfo {
		if !ListMatch(x, new.PackInfo) {
			log.Info().Str("Name", x.Name).Int64("Size", x.TotalSize).Msg("File exists but doesn't match size")
			l = append(l, x.Name + ".dat")
		} else {
			log.Info().Str("Name", x.Name).Int64("Size", x.TotalSize).Msg("File exists and matches size")
		}
		continue
	}
	for _, x := range old.ABInfo {
		if !ListMatch(x, new.ABInfo) {
			log.Info().Str("Name", x.Name).Int64("Size", x.TotalSize).Msg("File exists but doesn't match size")
			l = append(l, strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(x.Name, "/", "_"), ".ab", ".dat"), "#", "__"), ".mp4", ".dat"))
		} else {
			log.Info().Str("Name", x.Name).Int64("Size", x.TotalSize).Msg("File exists and matches size")
		}
		continue
	}
	return l, nil
}

func ListMatch(a types.Pack, l []types.Pack) bool {
    for _, b := range l {
        if (b.Name == a.Name && b.TotalSize == a.TotalSize) {
            return true
        }
    }
    return false
}


func Creator(url string, path string, platform string, version string) (error) {
	response, err := GetURLData(fmt.Sprintf(url, platform, version), stat.AssetBundleHotUpdateHeaders[:])
	if err != nil {
		return fmt.Errorf("Creator: failed to get response: %w", err)
	}

	buffer := bytes.NewBuffer(response)
	reader, err := gzip.NewReader(buffer)
	if err != nil {
		return fmt.Errorf("Creator: failed to read gzip: %w", err)
	}

	err = Generator(fmt.Sprintf(path, platform), reader)
	if err != nil {
		return fmt.Errorf("Creator: failed to generate gzip: %w", err)
	}
	return nil
}


func Marshaller(path string, payload interface{}) (error) {
	bin, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Marshaller: failed reading filee: %w", err)
	}

	err = json.Unmarshal(bin, &payload)
	if err != nil {
		return fmt.Errorf("Marshaller: failed unmarshalling: %w", err)
	}
	return nil
}


func Generator(path string, reader *gzip.Reader) error {
	out, err := os.Create(path)
	if err != nil  {
	  return fmt.Errorf("Generator: failed creating file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, reader)
	if err != nil  {
		return fmt.Errorf("Generator: failed copying file: %w", err)
	}
	return nil
}