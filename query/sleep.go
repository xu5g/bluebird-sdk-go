package query

// 获取睡眠数据
type SleepsGetQuery struct {
	StartTime string // 开始时间	格式为2006-01-01 00:00:00
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00
	Page      int32  // 页码	范围为1-5000
	Limit     int32  // 每页条数	范围为1-100
	ImeiSn    string // 设备号	长度不超过20
	Uuid      string // UUID	长度不超过40
	Sort      string // 排序方向 DESC/ASC；默认倒序
}

// 获取最近一次测量的睡眠数据
type SleepGetQuery struct {
	ImeiSn    string // 设备号	长度不超过20
	Uuid      string // UUID	长度不超过40
}

// 删除睡眠数据
type SleepDeleteQuery struct {
	PrimaryKey string // primary_key的值
	Id      int64  // ID
}