package query

// 获取api日志列表
type ApiLogsGetQuery struct {
	Page      int32  // 页码	范围为1-5000
	Limit     int32  // 每页条数	范围为1-100
	StartTime string // 开始时间	格式为2006-01-01 00:00:00
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00
	Sort      string // 排序方向 DESC/ASC；默认倒序
	TransId   string // 事件id
	LogType   string // 日志类型
	AppKey    string // 应用id
	ApiAlias  string // 接口别名
	ReqMethod string // 请求方式
}

// 删除api日志
type ApiLogDeleteQuery struct {
	PrimaryKey string // primary_key的值
	Id      int64  // ID
}