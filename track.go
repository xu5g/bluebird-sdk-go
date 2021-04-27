package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Track struct {
	Cfg *Config
}

// 获取轨迹列表
func (p *Track) GetTracks(query *query.TracksGetQuery) *result.TracksGetResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)
	params.Set("is_ignore", query.IsIgnore)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTracksGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.TracksGetResult{
			Result:result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.TracksGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.TracksGetResult{
			Result:result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}
	return resData
}
