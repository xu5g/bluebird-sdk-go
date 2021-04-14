package tspsdk

import (
	"fmt"
	"github.com/xu5g/bluebird-sdk-go/query"
	"testing"
)

// 设备详情
func TestDevice_GetDevice(t *testing.T) {
	query := &query.DeviceGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}
	res, err := NewClient(gateWay, appKey, token).Device().GetDevice(query)
	fmt.Println(res, err)
}

// 设备列表
func TestDevice_GetDevices(t *testing.T) {
	query := &query.DevicesGetQuery{
		Page: 1,
		Limit: 10,
	}
	res, err := NewClient(gateWay, appKey, token).Device().GetDevices(query)
	fmt.Println(res, err)
}

// 下发定位指令
func TestDevice_SendLocate(t *testing.T) {
	query := &query.DeviceLocateQuery{
		ImeiSn: "864002050010256",
	}
	res, err := NewClient(gateWay, appKey, token).Device().SendLocate(query)
	fmt.Println(res, err)
}

// 获取设备是否在线
func TestDevice_GetDeviceIsOnline(t *testing.T) {
	query := &query.DeviceIsOnlineQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}

	res, err := NewClient(gateWay, appKey, token).Device().GetDeviceIsOnline(query)
	fmt.Println(res, err)
}

// 透传报文
func TestDevice_SendMessage(t *testing.T) {
	query := &query.DeviceMessageQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Message: "ABABAB",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendMessage(query)
	fmt.Println(res, err)
}

// 获取设备功能清单
func TestDevice_GetDeviceModules(t *testing.T) {
	query := &query.DeviceModulesQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}

	res, err := NewClient(gateWay, appKey, token).Device().GetDeviceModules(query)
	fmt.Println(res, err)
}

// 绑定设备
func TestDevice_BindDevice(t *testing.T) {
	query := &query.DeviceBindQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		TrueName: "设备名称",
		Mobile: "11111111111",
		Uuid: "1000-0-0000-xxxxxxxxxxxxxxx",
	}

	res, err := NewClient(gateWay, appKey, token).Device().BindDevice(query)
	fmt.Println(res, err)
}

// 解绑设备
func TestDevice_UnBindDevice(t *testing.T) {
	query := &query.DeviceUnBindQuery{
		ImeiSn: "68946459036",
	}

	res, err := NewClient(gateWay, appKey, token).Device().UnBindDevice(query)
	fmt.Println(res, err)
}

// 下发寻找设备指令
func TestDevice_SendFindDevice(t *testing.T) {
	query := &query.DeviceFindQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendFindDevice(query)
	fmt.Println(res, err)
}

// 下发定位上报间隔指令
func TestDevice_SendLocateUpload(t *testing.T) {
	query := &query.DeviceLocateUploadQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Second: 400,
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendLocateUpload(query)
	fmt.Println(res, err)
}

// 下发定位时间段指令
func TestDevice_SendUdtime(t *testing.T) {
	query := &query.DeviceUdtimeQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Start: "06:00",
		End: "18:00",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendUdtime(query)
	fmt.Println(res, err)
}

// 下发设置亲情号码指令
func TestDevice_SendFamily(t *testing.T) {
	query := &query.DeviceFamilyQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Families: []query.Family{
			{
				Mobile: "xxxxxxxxx",
				Relation: "主人",
			},
		},
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendFamily(query)
	fmt.Println(res, err)
}

// 下发设置定位模式指令
func TestDevice_SendLocateMode(t *testing.T) {
	query := &query.DeviceLocateModeQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		LocateMode: "1",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendLocateMode(query)
	fmt.Println(res, err)
}

// 下发设置终端host指令
func TestDevice_SendHost(t *testing.T) {
	query := &query.DeviceHostQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Host: "test.xxx.xu5g.com",
		Port: "2232",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendHost(query)
	fmt.Println(res, err)
}

// 下发关机指令
func TestDevice_SendPowerOff(t *testing.T) {
	query := &query.DevicePowerOffQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendPowerOff(query)
	fmt.Println(res, err)
}

// 下发重启指令
func TestDevice_SendRestart(t *testing.T) {
	query := &query.DeviceRestartQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendRestart(query)
	fmt.Println(res, err)
}

// 下发聆听指令
func TestDevice_SendMonitor(t *testing.T) {
	query := &query.DeviceMonitorQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Mobile: "XXXXXXXXXXX",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendMonitor(query)
	fmt.Println(res, err)
}

// 下发设置免打扰时间段指令
func TestDevice_SendDnd(t *testing.T) {
	query := &query.DeviceDndQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Dnd: "8:00-11:30|123456;14:00-17:30|12345",
	}

	res, err := NewClient(gateWay, appKey, token).Device().SendDnd(query)
	fmt.Println(res, err)
}

// 更新设备信息
func TestDevice_DeviceUpdate(t *testing.T) {
	query := &query.DeviceUpdateQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Truename: "小明",
	}

	res, err := NewClient(gateWay, appKey, token).Device().DeviceUpdate(query)
	fmt.Println(res, err)
}

