package query

// 添加maclist
type MacListCreateQuery struct {
	Mac    string // mac地址（必填）
	Name   string // mac名称（必填）
	Signal int64  // 信号值（必填）
	Lat    string // 纬度（必填）
	Lng    string // 经度（必填）
}

// 获取maclist列表
type MacListsGetQuery struct {
	Page      int32  // 页码	范围为1-5000（必传）
	Limit     int32  // 每页条数	范围为1-100（必传）
	StartTime string // 开始时间	格式为2006-01-01 00:00:00（必传）
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00（必传）
	Mac       string // mac地址（必填）
}

// 删除maclist
type MacListDeleteQuery struct {
	MacAddr string // primary_key的值（必传）
	Id      int64  // ID（必传）
}
