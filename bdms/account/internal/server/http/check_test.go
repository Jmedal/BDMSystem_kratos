package http

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

type data struct {
	object string
	sign   string
}

func TestMain(m *testing.M) {
	map1 := make(map[string]interface{})
	//map1["courseId"] = []int64{1, 2, 3}
	//map1["time"] = []int64{1580486400, 1582905600}

	//map转为json
	str, err := json.Marshal(map1)
	//base64加密
	encode := base64.StdEncoding.EncodeToString(str)

	randomKey := "higqmysg"

	//生成签名
	d := []byte(encode+randomKey)
	has := md5.Sum(d)
	sign := fmt.Sprintf("%x", has)

	//封装数据
	object := &data{
		object: encode,
		sign:   sign,
	}
	fmt.Println(object)

	//验证sign签名
	if !verifySignature(object, randomKey) {
		fmt.Println("验证失败")
		return
	}
	fmt.Println("验证成功")
	//base64解密
	o, err := base64.StdEncoding.DecodeString(object.object)
	if err != nil {
		fmt.Println(err)
	}
	map2 := make(map[string]interface{})
	err = json.Unmarshal(o, &map2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json to map ", map2)
	for k, v := range map2 {
		fmt.Println(k, v)
	}

}

func verifySignature(object *data, randomKey string) bool {
	d := []byte(object.object + randomKey)
	has := md5.Sum(d)
	sign := fmt.Sprintf("%x", has)
	if sign == object.sign {
		return true
	} else {
		return false
	}
}
