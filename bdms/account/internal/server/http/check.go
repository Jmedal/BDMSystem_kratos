package http

import (
	"account/internal/server/http/util"
	"bytes"
	"encoding/json"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/render"
	"io/ioutil"
	"net/http"
)

//检查http请求签名认证
func dataSecurityAction() bm.HandlerFunc {
	return func(c *bm.Context) {
		req := c.Request
		log.Info(req.URL.String())

		//取出请求数据
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("获取Body数据发生错误，错误信息:", err)
		}

		//登录接口的数据处理
		if url := req.URL.String(); url == "/service.v1.Account/Login" {
			if data != nil && len(data) > 0 {
				decodeData, err := util.Read(data, "")
				if err != nil || decodeData == nil {
					log.Error("解析加密数据时发生错误，错误信息:", err)
					renderErrMsg(c, ecode.NothingFound.Code(), "Invalid encodeData")
					return
				}
				//把读过的字节流重新放到body
				req.Body = ioutil.NopCloser(bytes.NewBuffer(decodeData))
			}
			c.Next()
			return
		}

		token := req.Header.Get("Authorization")
		if token == "" {
			renderErrMsg(c, ecode.Unauthorized.Code(), "API token required")
			return
		}

		// TODO: get userId,randomKey from some code
		userId, randomKey, err := svd.VerifyToken(c, token)
		if err != nil {
			renderErrMsg(c, ecode.ServerErr.Code(), "Token server error")
			return
		}
		if userId < 0 {
			renderErrMsg(c, ecode.AccessDenied.Code(), "Invalid API token")
			return
		}

		//解析加密数据
		if data != nil && len(data) > 0 {
			decodeData, err := util.Read(data, randomKey)
			if err != nil || decodeData == nil {
				log.Error("解析加密数据时发生错误，错误信息:", err)
				renderErrMsg(c, ecode.NothingFound.Code(), "Invalid encodeData")
				return
			}
			//添加userId
			body, err := setBodyParam(decodeData, "userId", userId)
			if err != nil {
				log.Error("Body添加数据时发生错误，错误信息:", err)
			}
			//把读过的字节流重新放到body
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
	}
}

//向body中添加数据
func setBodyParam(old []byte, key string, value interface{}) (new []byte, err error) {
	dataMap := make(map[string]interface{})
	err = json.Unmarshal(old, &dataMap)
	if err != nil {
		log.Error("数据转Map时发生错误，错误信息:", err)
	}
	dataMap[key] = value
	new, err = json.Marshal(dataMap)
	if err != nil {
		log.Error("Map转数据时发生错误，错误信息:", err)
	}
	return
}

//设置Response内容
func renderErrMsg(c *bm.Context, code int, msg string) {
	data := map[string]interface{}{
		"code":    code,
		"message": msg,
	}
	c.Render(http.StatusOK, render.MapJSON(data))
	c.Abort()
}
