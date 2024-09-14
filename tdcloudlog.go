package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type TdcloudLog struct {
	Cfg *Config
}

// 获取api日志列表
func (p *TdcloudLog) GetTdcloudLogs(query *query.TdcloudLogsGetQuery) *result.TdcloudLogsGetResult {
	params := url.Values{}
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)
	params.Set("event_id", query.EventId)
	params.Set("trans_id", query.TransId)
	params.Set("imei_sn", query.ImeiSn)
	params.Set("appkey", strconv.Itoa(int(query.Appkey)))
	params.Set("partner_id", strconv.Itoa(int(query.PartnerId)))
	params.Set("push_type", strconv.Itoa(int(query.PushType)))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTdcloudLogsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.TdcloudLogsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.TdcloudLogsGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.TdcloudLogsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}
