package helpers

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	stat "github.com/Etwodev/Doctorate/static"
	"github.com/bwmarrin/snowflake"
)

func GenerateOTP(length int) (string, error) {
    buffer := make([]byte, length)
    _, err := rand.Read(buffer)
    if err != nil {
        return "", fmt.Errorf("GenerateOTP: failed reading buffer: %w", err)
    }

    l := len(stat.OTP)
    for i := 0; i < length; i++ {
        buffer[i] = stat.OTP[int(buffer[i])%l]
    }

    return string(buffer), nil
}

func GenerateSnowflake(n int64) (snowflake.ID, error) {
	node, err := snowflake.NewNode(n)
	if err != nil {
		return 0, fmt.Errorf("GenerateSnowflake: failed generating snowflake: %w", err)
	}
	return node.Generate(), nil
}

func Serialization(path string) (string, error) {
	bin, err := ioutil.ReadFile(path)
	if err != nil {
		return "",  fmt.Errorf("Serialization: failed reading file: %w", err)
	}
	return string(bin), nil
}

func GenerateSecureToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}

func HashWithMD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func SignatureWithMD5(str string, path string) (string, error) {
	hash := md5.Sum([]byte(str))
	bin, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("SignatureWithMD5: failed reading file: %w", err)
	}

	block, _ := pem.Decode([]byte(bin))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("SignatureWithMD5: failed parsing key: %w", err)
	}

	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.MD5, hash[:])
	if err != nil {
		return "", fmt.Errorf("SignatureWithMD5: failed signing data: %w", err)
	}

	return base64.StdEncoding.EncodeToString(sign), nil
}

func GetURLData(url string, headers [][2]string) ([]byte, error) {
	c := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("GetURLData: failed creating request: %w", err)
	}

	for _, header := range headers {
		request.Header.Add(header[0], header[1])
	}

	response, err := c.Do(request)
	if err != nil {
		return nil, fmt.Errorf("GetURLData: failed sending request: %w", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("GetURLData: failed reading body: %w", err)
	}

	return body, nil
}

func OpenFile(path string) ([]byte, error) {
	dub, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("OpenJson: failed opening file: %w", err)
	}

	bin, err := ioutil.ReadAll(dub)
	if err != nil {
		return nil, fmt.Errorf("OpenJson: failed reading file: %w", err)
	}
	return bin, nil
}