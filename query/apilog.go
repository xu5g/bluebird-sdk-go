package query

// 获取api日志列表
type ApiLogsGetQuery struct {
	Page      int32  // 页码	范围为1-5000（必传）
	Limit     int32  // 每页条数	范围为1-100（必传）
	StartTime string // 开始时间	格式为2006-01-01 00:00:00（必传）
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00（必传）
	Sort      string // 排序方向 DESC/ASC；默认倒序
	TransId   string // 事件id
	LogType   string // 日志类型
	AppKey    string // 应用id
}
