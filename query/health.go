package query

// 获取喝水记录数据
type WaterrecordsGetQuery struct {
	Page      int32  // 页码	范围为1-5000
	Limit     int32  // 每页条数	范围为1-200
	StartTime string // 开始时间	格式为2006-01-01 00:00:00
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00
	ImeiSn    string // 设备号	长度不超过20
	Uuid      string // UUID	长度不超过40
	Sort      string // 排序方向 DESC/ASC；默认倒序
	PartnerId int32
}

// 设置健康提醒
type HealthremindSetQuery struct {
	ImeiSn           string
	HealthremindType int64
	IsOpen           int64
	Interval         int64
	StartTime        string
	EndTime          string
	IsNoondnd        int64
	DndStartTime     string
	DndEndTime       string
	Week             int64
	Target           int64
}
