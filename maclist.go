package tspsdk

import (
	"encoding/json"
	"github.com/gogf/gf/container/gmap"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type MacList struct {
	Cfg *Config
}

// 添加maclist
func (p *MacList) MacListCreate(query *query.MacListCreateQuery) (*result.Result, error) {
	var data = gmap.New(true)
	data.Set("name", query.Name)
	data.Set("lng", query.Lng)
	data.Set("lat", query.Lat)
	data.Set("mac", query.Mac)
	data.Set("signal", query.Signal)

	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacListCreatePath).SetData(data).HttpRequest()
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

// 获取mac列表
func (p *MacList) GetMacLists(query *query.MacListsGetQuery) (*result.MacListsGetResult, error) {
	params := url.Values{}
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("mac", query.Mac)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacListsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var result = new(result.MacListsGetResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 删除maclist
func (p *MacList) DeleteMacList(query *query.MacListDeleteQuery) (*result.Result, error) {
	var data = gmap.New(true)
	data.Set("id", query.Id)
	data.Set("macaddr", query.PrimaryKey)

	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacListDeletePath).SetData(data).HttpRequest()
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
