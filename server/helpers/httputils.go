package helpers

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithOctet(w http.ResponseWriter, code int, path string) {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(code)
	w.Write(fileBytes)
}

func RespondWithRawJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	response = []byte(response)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func MD5SignWithPrivateKey(location string, data string) (string) {
	hashed := md5.Sum([]byte(data)) 
	tmp, err := ioutil.ReadFile(location)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode([]byte(tmp))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.MD5, hashed[:])
	if err != nil {
		panic(err)
	}
	out := base64.StdEncoding.EncodeToString(sign)

    return out
}