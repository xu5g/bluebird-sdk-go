package result

// 设置配对蓝牙
type SetBluetoothLinkResult struct {
	Result
	Data interface{} `json:"data"` // 返回结果
}
