package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"net/url"
	"strconv"
)

type Temperature struct {
	Cfg *Config
}

// 获取最新体温数据
func (p *Temperature) GetTemperature(query *query.TemperatureRecentQuery) (*result.TemperatureResult, error) {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetToken(p.Cfg.HttpClient.Token).SetUrl(TSPTemperatureGetUrl + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var temperatureEntity = new(result.TemperatureResult)
	err = json.Unmarshal([]byte(jsonString), temperatureEntity)
	if err != nil {
		return nil, err
	}
	return temperatureEntity, nil
}

// 获取体温列表
func (p *Temperature) GetTemperatures(query *query.TemperaturesQuery) (*result.TemperaturesResult, error) {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("product_id", strconv.FormatInt(query.ProductId, 10))
	params.Set("sort", query.Sort)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetToken(p.Cfg.HttpClient.Token).SetUrl(TSPTemperaturesGetUrl + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var temperatureEntity = new(result.TemperaturesResult)
	err = json.Unmarshal([]byte(jsonString), temperatureEntity)
	if err != nil {
		return nil, err
	}
	return temperatureEntity, nil
}

// 获取体温测量间隔时间
func (p *Temperature) GetTemperatureUpload(query *query.TemperatureUploadQuery) (*result.TemperatureUploadResult, error) {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetToken(p.Cfg.HttpClient.Token).SetUrl(TSPTemperatureUploadGetUrl + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var temperatureUploadResult = new(result.TemperatureUploadResult)
	err = json.Unmarshal([]byte(jsonString), temperatureUploadResult)
	if err != nil {
		return nil, err
	}
	return temperatureUploadResult, nil
}

// 设置体温测量间隔时间
func (p *Temperature) UpdateTemperatureUpload(param *query.TemperatureUpload) (*result.Result, error) {

	var data = make(map[string]interface{})
	data["imei_sn"] = param.ImeiSn
	data["second"] = param.Second

	res, err := p.Cfg.HttpClient.SetMethod("put").SetToken(p.Cfg.HttpClient.Token).SetUrl(TSPTemperatureUploadPutUrl).SetData(data).HttpRequest()
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
