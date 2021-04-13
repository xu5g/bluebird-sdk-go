package query

// 获取最近一次测量的血压数据
type BloodGetQuery struct {
	ImeiSn string // 设备IMEI号  长度不超过20（ImeiSn和Uuid不能同时为空）
	Uuid   string // 设备uuid	长度不超过40
}

// 获取血压数据
type BloodsGetQuery struct {
	StartTime string // 开始时间	格式为2006-01-01 00:00:00（必传）
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00（必传）
	Page      int32  // 页码	范围为1-5000（必传）
	Limit     int32  // 每页条数	范围为1-100（必传）
	ImeiSn    string // 设备号	长度不超过20
	Uuid      string // UUID	长度不超过40
	Sort      string // 排序方向 DESC/ASC；默认倒序
}
