package query

// 获取轨迹
type TracksGetQuery struct {
	StartTime string // 开始时间	格式为2006-01-01 00:00:00（必传）
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00（必传）
	Page      int32  // 页码		范围为1-5000（必传）
	Limit     int32  // 每页条数	范围为1-200（必传）
	Sort      string // 排序方向 DESC/ASC；默认倒序
	ImeiSn    string // 设备Imei号 长度不超过20
	Uuid      string // 设备Uuid		长度不超过40
}
