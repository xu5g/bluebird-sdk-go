package query

type TemperatureRecentQuery struct {
	ImeiSn string
	Uuid   string
}

type TemperaturesQuery struct {
	StartTime string // 开始时间
	EndTime   string // 结束时间
	ImeiSn    string // 设备号
	Uuid      string // UUID
	Page      int32  // 页码
	Limit     int32  // 每页条数
	ProductId int64  // 按照ProductId查询
	Sort      string // 排序方向 DESC/ASC；默认倒序
}

type TemperatureUploadQuery struct {
	ImeiSn string // 设备号
}

// 设置体温测量间隔时间
type TemperatureUpload struct {
	ImeiSn string // 设备号
	Second int    // 体温上报间隔时间（单位秒）
}
