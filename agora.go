package tspsdk

import (
	"encoding/json"
	"fmt"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
)

type Agora struct {
	Cfg *Config
}

// 更新音视频通话用户列表
func (p *Agora) UpdateAgoraUsers(query *query.UpdateAgoraUsersQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["agora"] = query.Agora
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TspAgoraUsers).SetData(data).HttpRequest()
	if err != nil {
		return &result.Result{
			Status:  1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status:  1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发APP邀请手表进行音视频通话指令
func (p *Agora) SendAppcalldevice(query *query.AgoraAppcalldeviceQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["uid"] = query.Uid
	data["rel_name"] = query.RelName
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TspAgoraAppcalldevice).SetData(data).HttpRequest()
	if err != nil {
		return &result.Result{
			Status:  1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()
	fmt.Println(jsonString)

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status:  1,
			Message: err.Error(),
		}
	}
	return resData
}
