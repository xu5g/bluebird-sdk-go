package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/v2/query"
	"github.com/xu5g/bluebird-sdk-go/v2/result"
	"github.com/xu5g/bluebird-sdk-go/v2/util"
	"net/url"
	"strconv"
)

type Health struct {
	Cfg *Config
}

// 获取喝水记录列表
func (p *Health) GetWaters(query *query.WaterrecordsGetQuery) *result.WaterrecordsGetResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)
	params.Set("partner_id", strconv.Itoa(int(query.PartnerId)))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPWaterrecordsPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.WaterrecordsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.MustToJsonString()
	var resData = new(result.WaterrecordsGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.WaterrecordsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// DeviceRemind 设置健康提醒
func (p *Health) SetHealthremind(query *query.HealthremindSetQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["healthremind_type"] = query.HealthremindType
	data["is_open"] = query.IsOpen
	data["interval"] = query.Interval
	data["start_time"] = query.StartTime
	data["end_time"] = query.EndTime
	data["is_noondnd"] = query.IsNoondnd
	data["dnd_startTime"] = query.DndStartTime
	data["dnd_endTime"] = query.DndEndTime
	data["week"] = query.Week
	data["target"] = query.Target
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPHealthremindPath).SetData(data).HttpRequest()

	if err != nil {
		return &result.Result{
			Status:  1,
			Message: err.Error(),
		}
	}
	jsonString := res.MustToJsonString()

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
