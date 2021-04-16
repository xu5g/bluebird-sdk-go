package tspsdk

import (
	"encoding/json"
	"fmt"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Device struct {
	Cfg *Config
}

// 设备详情
func (p *Device) GetDevice(query *query.DeviceGetQuery) *result.DeviceGetResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("attence_sn", query.AttenceSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.DeviceGetResult{
			Result:result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.DeviceGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.DeviceGetResult{
			Result:result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 更新设备信息
func (p *Device) DeviceUpdate(query *query.DeviceUpdateQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["truename"] = query.Truename

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceUpdatePath).SetData(data).HttpRequest()
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}

	jsonString := res.Export()
	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 获取设备列表
func (p *Device) GetDevices(query *query.DevicesGetQuery) *result.DevicesResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("attence_sn", query.AttenceSn)
	params.Set("mobile", query.Mobile)
	params.Set("is_online", query.IsOnline)
	params.Set("uuid", query.Uuid)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("model_id", strconv.FormatInt(query.ModelId, 10))
	params.Set("appkey", strconv.FormatInt(query.AppKey, 10))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDevicesGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.DevicesResult{
			Result:result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.DevicesResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.DevicesResult{
			Result:result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 下发定位指令到终端
func (p *Device) SendLocate(query *query.DeviceLocateQuery) *result.Result {

	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceLocatePath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()
	fmt.Println(jsonString)

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 判断设备是否在线
func (p *Device) GetDeviceIsOnline(query *query.DeviceIsOnlineQuery) *result.DeviceIsOnlineResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceOnlinePath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.DeviceIsOnlineResult{
			Result: result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.DeviceIsOnlineResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.DeviceIsOnlineResult{
			Result: result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 透传报文
func (p *Device) SendMessage(query *query.DeviceMessageQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["message"] = query.Message
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceMessagePath).SetData(data).HttpRequest()

	fmt.Println(res, err)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 获取设备功能清单
func (p *Device) GetDeviceModules(query *query.DeviceModulesQuery) *result.DeviceModulesResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceModulesPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.DeviceModulesResult{
			Result: result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}

	jsonString := res.Export()
	var resData = new(result.DeviceModulesResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.DeviceModulesResult{
			Result: result.Result{
				Status: 1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 设备绑定
func (p *Device) BindDevice(query *query.DeviceBindQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["mobile"] = query.Mobile
	data["uuid"] = query.Uuid
	data["truename"] = query.TrueName
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceBindPath).SetData(data).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 设备解绑
func (p *Device) UnBindDevice(query *query.DeviceUnBindQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceUnBindPath).SetData(data).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发设备寻找指令
func (p *Device) SendFindDevice(query *query.DeviceFindQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceFindPath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发定位上报间隔指令
func (p *Device) SendLocateUpload(query *query.DeviceLocateUploadQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["second"] = query.Second
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceLocateUploadPath).SetData(data).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发定位时间段指令
func (p *Device) SendUdtime(query *query.DeviceUdtimeQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["start"] = query.Start
	data["end"] = query.End
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceUdtimePath).SetData(data).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发设置亲情号码指令
func (p *Device) SendFamily(query *query.DeviceFamilyQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["families"] = query.Families
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceFamilyPath).SetData(data).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发设置定位模式指令
func (p *Device) SendLocateMode(query *query.DeviceLocateModeQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["locate_mode"] = query.LocateMode
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceLocateModePath).SetData(data).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发设置终端host指令
func (p *Device) SendHost(query *query.DeviceHostQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["host"] = query.Host
	data["port"] = query.Port
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceHostPath).SetData(data).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发关机指令
func (p *Device) SendPowerOff(query *query.DevicePowerOffQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDevicePowerOffPath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发重启指令
func (p *Device) SendRestart(query *query.DeviceRestartQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceRestartPath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发聆听指令
func (p *Device) SendMonitor(query *query.DeviceMonitorQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("mobile", query.Mobile)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceMonitorPath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}

// 下发设置免打扰时间段指令
func (p *Device) SendDnd(query *query.DeviceDndQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("dnd", query.Dnd)

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceDndPath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	jsonString := res.Export()

	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status: 1,
			Message: err.Error(),
		}
	}
	return resData
}
