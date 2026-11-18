package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/v2/query"
	"github.com/xu5g/bluebird-sdk-go/v2/result"
	"github.com/xu5g/bluebird-sdk-go/v2/util"
)

type Speaker struct {
	Cfg *Config
}

// 音箱-文本转语音播放（lyds TTS；TSP 侧队列模式写死，多次调用自动排队顺播）
func (p *Speaker) PlayText(query *query.SpeakerPlayTextQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["text"] = query.Text
	data["vcn"] = query.Vcn
	data["speed"] = query.Speed
	data["volume"] = query.Volume
	data["prompt"] = query.Prompt
	data["loop_times"] = query.LoopTimes
	data["loop_duration"] = query.LoopDuration
	data["loop_gap"] = query.LoopGap
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TspSpeakerSpeechText).SetData(data).HttpRequest()
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