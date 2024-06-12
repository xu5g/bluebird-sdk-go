package query

// 获取跳绳数据
type SkipsGetQuery struct {
	ImeiSn    string
	Uuid      string
	StartTime string
	EndTime   string
	Page      int32
	Limit     int32
	Sort      string
	AppKey    int64
}
