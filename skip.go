package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Skip struct {
	Cfg *Config
}

// 获取跳绳数据列表
func (p *Skip) GetSkips(query *query.SkipsGetQuery) *result.SkipsResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)
	params.Set("appkey", strconv.Itoa(int(query.AppKey)))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TspSkipsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.SkipsResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.SkipsResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.SkipsResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 设置手表和蓝牙跳绳定时开启自动建立连接
func (p *Skip) SetAcdateskip(query *query.SetAcdateSkipRequest) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["skipdate"] = query.Skipdate
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TspSkipAcdateskipPath).SetData(data).HttpRequest()

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

// 设置手表和蓝牙跳绳自动连接以及跳绳数据的倒计时上报时间
func (p *Skip) SetAcuskip(query *query.SetAcuSkipRequest) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["countdown"] = query.Countdown
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TspSkipAcuskipPath).SetData(data).HttpRequest()

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
