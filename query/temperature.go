package query

// 获取最新体温数据
type TemperatureGetQuery struct {
	ImeiSn string // 设备IMEI号  长度不超过20（ImeiSn和Uuid不能同时为空）
	Uuid   string // 设备uuid	长度不超过40
}

type TemperaturesQuery struct {
	StartTime string // 开始时间	格式为2006-01-01 00:00:00（必传）
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00（必传）
	Page      int32  // 页码	范围为1-5000（必传）
	Limit     int32  // 每页条数	范围为1-200（必传）
	ImeiSn    string // 设备号	长度不超过20
	Uuid      string // UUID	长度不超过40
	Sort      string // 排序方向 DESC/ASC；默认倒序
}

type TemperatureUploadQuery struct {
	ImeiSn string // 设备IMEI号  长度不超过20
}

// 设置体温测量间隔时间
type TemperatureUpload struct {
	ImeiSn string // 设备IMEI号  长度不超过20
	Second int    // 体温上报间隔时间（单位秒）
}
