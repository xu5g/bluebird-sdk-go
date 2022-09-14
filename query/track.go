package query

// 获取轨迹
type TracksGetQuery struct {
	StartTime string // 开始时间	格式为2006-01-01 00:00:00
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00
	Page      int32  // 页码		范围为1-5000
	Limit     int32  // 每页条数	范围为1-100
	Sort      string // 排序方向 DESC/ASC；默认倒序
	ImeiSn    string // 设备Imei号 长度不超过20
	Uuid      string // 设备Uuid	  长度不超过40
	IsIgnore  string // 是否查询出被忽略的数据 值：0-不忽略 1：忽略
}

// TracksUpdateQuery 修改轨迹数据
type TracksUpdateQuery struct {
	PrimaryKey string // 主键
	Id         int64  // 结束时间	格式为2006-01-01 00:00:00
	IsIgnore   int64  // 是否忽略，0：未忽略；1：忽略
}

// 删除轨迹数据
type TrackDeleteQuery struct {
	PrimaryKey string // primary_key的值
	Id      int64  // ID
}