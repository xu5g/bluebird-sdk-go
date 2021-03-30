package tspsdk

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"net/url"
)

type Auth struct {
	Cfg *Config
}

// 获取token
func (p *Auth) GetToken() (*AuthResult, error) {
	appKey := p.Cfg.HttpClient.AppKey
	tranSid := fmt.Sprintf("%s%s%s", appKey, gtime.Now().Format("YmdHis"), grand.Letters(12))
	sign := gsha1.Encrypt(fmt.Sprintf("%s%s%s%s", p.Cfg.HttpClient.Secret, appKey, tranSid, p.Cfg.HttpClient.Secret))

	params := url.Values{}
	params.Set("transid", tranSid)
	params.Set("appkey", appKey)
	params.Set("sign", sign)

	res, err := p.Cfg.HttpClient.SetUrl(AuthURL + "?" + params.Encode()).SetMethod("get").HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var authEntity = new(AuthResult)
	err = json.Unmarshal([]byte(jsonString), authEntity)
	if err != nil {
		return nil, err
	}

	return authEntity, err
}



