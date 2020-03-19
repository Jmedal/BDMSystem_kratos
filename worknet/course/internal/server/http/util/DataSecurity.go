package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
)


//解析校验加密数据
func Read(data []byte, randomKey string) (decodeData []byte, err error) {

	//提取数据
	encodeMap := make(map[string]string)
	if err = json.Unmarshal(data, &encodeMap); err != nil {
		return
	}

	//校验签名
	if !verifySignature(encodeMap, randomKey) {
		decodeData = nil
		return
	}

	//base64解密
	decodeData, err = base64.StdEncoding.DecodeString(encodeMap["object"])
	return
}

//校验签名
func verifySignature(data map[string]string, randomKey string) bool {
	d := []byte(data["object"] + randomKey)
	has := md5.Sum(d)
	sign := fmt.Sprintf("%x", has)

	if data["sign"] == sign {
		return true
	} else {
		return false
	}

}
