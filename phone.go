package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/v2/query"
	"github.com/xu5g/bluebird-sdk-go/v2/result"
	"github.com/xu5g/bluebird-sdk-go/v2/util"
)

type Phone struct {
	Cfg *Config
}

// 批量设置学生（添加和修改共用，删除学生请用 DeletePhoneStudents）
func (p *Phone) SetPhoneStudents(query *query.PhoneStudentSetQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["students"] = query.Students
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TspPhoneStudent).SetData(data).HttpRequest()
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

// 批量删除学生
func (p *Phone) DeletePhoneStudents(query *query.PhoneStudentDeleteQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["student_ids"] = query.StudentIds
	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TspPhoneStudent).SetData(data).HttpRequest()
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

// 设置家长关系（增删改一体，删除用 del_flag=1 条目）
func (p *Phone) SetPhoneParents(query *query.PhoneParentSetQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["parents"] = query.Parents
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TspPhoneParent).SetData(data).HttpRequest()
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

// 家长向学生发送消息（type=0 文字传 content，type=1 语音传 audio_url）
func (p *Phone) SendPhoneMessage(query *query.PhoneParentMessageQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["parent_id"] = query.ParentId
	data["student_id"] = query.StudentId
	data["type"] = query.Type
	data["content"] = query.Content
	data["audio_url"] = query.AudioUrl
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TspPhoneParentMessage).SetData(data).HttpRequest()
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
