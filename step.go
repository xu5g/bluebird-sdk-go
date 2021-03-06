package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Step struct {
	Cfg *Config
}

// 获取计步列表
func (p *Step) GetSteps(query *query.StepsGetQuery) *result.StepsResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPStepsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.StepsResult{
			Result:result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.StepsResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.StepsResult{
			Result:result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}
	return resData
}