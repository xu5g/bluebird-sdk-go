package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Fence struct {
	Cfg *Config
}


// 获取围栏列表
func (p *Fence) GetFences(query *query.FencesGetQuery) (*result.FencesGetResult, error) {
	params := url.Values{}
	params.Set("uuid", query.Uuid)
	params.Set("page", strconv.Itoa(query.Page))
	params.Set("limit", strconv.Itoa(query.Limit))
	params.Set("fence_type", strconv.Itoa(int(query.FenceType)))
	params.Set("shape_type", strconv.Itoa(int(query.ShapeType)))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPFencesGetPath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var result = new(result.FencesGetResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 创建围栏
func (p *Fence) CreateFence(query *query.FenceCreateQuery) (*result.FenceCreateResult, error) {
	var data = make(map[string]interface{})
	data["truename"] = query.Truename
	data["fence_type"] = query.FenceType
	data["shape_type"] = query.ShapeType
	data["collide_type"] = query.CollideType
	data["geo"] = query.Geo
	data["near_radius"] = query.NearRadius
	data["valid_start"] = query.ValidStart
	data["valid_end"] = query.ValidEnd
	data["valid_week"] = query.ValidWeek
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPFenceCreatePath).SetData(data).HttpRequest()

	if err != nil {
		return nil, err
	}
	jsonString := res.Export()

	var result = new(result.FenceCreateResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 删除围栏
func (p *Fence) DeleteFence(query *query.FenceDeleteQuery) (*result.Result, error) {
	var data = make(map[string]interface{})
	data["id"] = query.Id
	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TspFenceDeletePath).SetData(data).HttpRequest()

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
