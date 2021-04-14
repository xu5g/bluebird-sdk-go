package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Blood struct {
	Cfg *Config
}

// 获取最新血压数据
func (p *Blood) GetBlood(query *query.BloodGetQuery) (*result.BloodGetResult, error) {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var result = new(result.BloodGetResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取血压列表
func (p *Blood) GetBloods(query *query.BloodsGetQuery) (*result.BloodsGetResult, error) {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var result = new(result.BloodsGetResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取血压测量间隔时间
func (p *Blood) GetBloodUpload(query *query.BloodUploadGetQuery) (*result.TemperatureUploadResult, error) {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodUploadGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var result = new(result.TemperatureUploadResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 设置血压测量间隔时间
func (p *Blood) UpdateBloodUpload(param *query.BloodUploadSetQuery) (*result.Result, error) {

	var data = make(map[string]interface{})
	data["imei_sn"] = param.ImeiSn
	data["second"] = param.Second

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodUploadSetPath).SetData(data).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()

	var result = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
