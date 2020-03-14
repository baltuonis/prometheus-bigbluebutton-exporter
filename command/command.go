package command

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

/*******************************************************************************
* 根据请求的接口, 参数以及公钥生成密文
* 参数: method, 请求的接口
*	   param, 请求携带的参数
*      salt, 服务器提供的公钥
* 返回: 加密后的checksum密文
*******************************************************************************/
func GetChecksum(method string, param string, salt string) string {
	private := []byte(method + param + salt)
	ciphertext := sha1.Sum(private)

	return hex.EncodeToString(ciphertext[:])
}

/*******************************************************************************
* 执行HTTP GET请求, 返回请求结果
* 参数: url, 携带参数的请求地址
* 返回: 请求结果, 如果返回ERROR说明请求过程中出错, 详细信息可以查看log
*******************************************************************************/
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

/*******************************************************************************
* 将Struct转换为Map格式
* type demo struct {              key    value
*     id string        ----\      id     001
*     name string      ----/      name   名字
* }
* 参数: obj, 需要转换的结构体实例
* 返回: Map类型的结果
*******************************************************************************/
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}
