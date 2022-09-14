package query

// 获取最近一次测量的心率数据
type HeartGetQuery struct {
	ImeiSn string // 设备IMEI号  长度不超过20
	Uuid   string // 设备uuid	长度不超过40
}

// 获取心率数据列表
type HeartsGetQuery struct {
	Page      int32  // 页码	范围为1-5000
	Limit     int32  // 每页条数	范围为1-100
	StartTime string // 开始时间	格式为2006-01-01 00:00:00
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00
	ImeiSn    string // 设备号	长度不超过20
	Uuid      string // UUID	长度不超过40
	Sort      string // 排序方向 DESC/ASC；默认倒序
}

// 获取心率数据上报时间间隔
type HeartUploadGetQuery struct {
	ImeiSn string // 设备IMEI号  长度不超过20
}

// 设置心率数据上报时间间隔
type HeartUploadSetQuery struct {
	ImeiSn string // 设备IMEI号  长度不超过20
	Second int    // 体温上报间隔时间（单位秒）
}

// 删除心率
type HeartDeleteQuery struct {
	PrimaryKey string // primary_key的值
	Id      int64  // ID
}
