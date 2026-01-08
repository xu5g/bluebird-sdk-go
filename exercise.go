package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/v2/query"
	"github.com/xu5g/bluebird-sdk-go/v2/result"
	"github.com/xu5g/bluebird-sdk-go/v2/util"
	"net/url"
	"strconv"
)

type Exercise struct {
	Cfg *Config
}

// 获取运动数据列表
func (p *Exercise) GetExercises(query *query.ExercisesGetQuery) *result.ExercisesResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)
	params.Set("type", strconv.Itoa(int(query.Type)))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPExercisesPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.ExercisesResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.MustToJsonString()
	var resData = new(result.ExercisesResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.ExercisesResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}
