package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	httpMethod := "POST"
	httpUri := "/v1/identify"
	dataType := "audio"
	signatureVersion := "1"
	timestamp := string(time.Now().Unix())
	// Replace "###...###" below with your project's host, access_key and access_secret.
	requrl := "http://identify-us-west-2.acrcloud.com/v1/identify"
	accessKey := "629bcdc80a070d899bb7d9cf86ac87eb"
	accessSecret := "xKrvfcobKOfYiMlGfwp8hQf8BwvvGPV8WR7E5Ad8"
	stringToSign := httpMethod + "\n" +
		httpUri + "\n" +
		accessKey + "\n" +
		dataType + "\n" +
		signatureVersion + "\n" +
		timestamp


	hash := hmac.New(sha1.New, []byte(accessSecret))
	hash.Write([]byte(stringToSign))

	// to lowercase hexits
	hex.EncodeToString(hash.Sum(nil))

	// to base64
	signatureString := b64.StdEncoding.EncodeToString(hash.Sum(nil))

	filePath := "/Users/maxheckel/Sites/vinyl-analytics/api/cmd/listen/test.mp3"
	fileBytes, _ := ioutil.ReadFile(filePath)
	fi, err := os.Stat(filePath);
	if err != nil {
		panic(err)
	}
	// get the size
	size := fi.Size()
	_, err = resty.R().
		SetFileReader("sample", filePath, bytes.NewReader(fileBytes)).
		SetFormData(map[string]string{
			"sample_bytes":      string(size),
			"access_key":        accessKey,
			"data_type":         dataType,
			"signature":         signatureString,
			"signature_version": signatureVersion,
			"timestamp":         timestamp,
		}).
		Post(requrl)
	resp, err := resty.R().
		SetFileReader("sample", "test.mp3", bytes.NewReader(fileBytes)).
		SetFormData(map[string]string{
			//"sample": string(fileBytes),
			"sample_bytes":      string(size),
			"access_key":        accessKey,
			"data_type":         dataType,
			"signature":         signatureString,
			"signature_version": signatureVersion,
			"timestamp":         timestamp,
		}).
		Post(requrl)
	fmt.Println(string(resp.Body()))
}

func hash_hmac_sha1(password, key []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(password)
	return h.Sum(nil)
}
