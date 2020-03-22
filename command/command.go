package command

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// GetChecksum - calculates BBB checksum based on Method, Param & Secret (salt)
func GetChecksum(method string, param string, salt string) string {
	private := []byte(method + param + salt)
	ciphertext := sha1.Sum(private)

	return hex.EncodeToString(ciphertext[:])
}

// HttpGet - makes HTTP GET request
func HttpGet(url string) string {
	response, err := http.Get(url)

	if nil != err {
		log.Println("HTTP GET ERROR: " + err.Error())
		return "ERROR"
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if nil != err {
		log.Println("HTTP GET ERROR: " + err.Error())
		return "ERROR"
	}

	return string(body)
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}
