package http

import (
	"bytes"
	"encoding/json"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/render"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	xtime "github.com/bilibili/kratos/pkg/time"
	"io/ioutil"
	"net/http"
	pb "role/api"
	"role/internal/server/http/util"
	"time"
)

//初始化token注册客户端
func initTokenClient() (grpcClient pb.TokenClient, err error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 10),
		Timeout:           xtime.Duration(time.Millisecond * 250),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
	}

	return pb.NewClient(grpccfg)
}

//检查http请求签名认证
func dataSecurityAction() bm.HandlerFunc {
	return func(c *bm.Context) {
		req := c.Request
		log.Info(req.URL.String())
		token := req.Header.Get("Authorization")
		if token == "" {
			renderErrMsg(c, ecode.Unauthorized.Code(), "API token required")
			return
		}

		// TODO: get userId,randomKey from some code
		resp, err := tokenClient.Verify(c, &pb.VerifyTokenReq{AccessToken: token})
		if err != nil {
			renderErrMsg(c, ecode.ServerErr.Code(), "Token server error")
			return
		}
		if resp.UserId < 0 {
			renderErrMsg(c, ecode.AccessDenied.Code(), "Invalid API token")
			return
		}

		//取出请求数据
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("获取Body数据发生错误，错误信息:", err)
		}

		//解析加密数据
		if data != nil && len(data) > 0 {
			decodeData, err := util.Read(data, resp.RandomKey)
			if err != nil || decodeData == nil {
				log.Error("解析加密数据时发生错误，错误信息:", err)
				renderErrMsg(c, ecode.NothingFound.Code(), "Invalid encodeData")
				return
			}
			//添加userId
			body, err := setBodyParam(decodeData, "userId", resp.UserId)
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
