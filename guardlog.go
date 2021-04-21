package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type GuardLog struct {
	Cfg *Config
}

// 获取api日志列表
func (p *GuardLog) GetGuardLogs(query *query.GuardLogsGetQuery) *result.GuardLogsGetResult {
	params := url.Values{}
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)
	params.Set("event_id", query.EventId)
	params.Set("status", strconv.Itoa(query.Status))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPGuardLogsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.GuardLogsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.GuardLogsGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.GuardLogsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}
