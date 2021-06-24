package query

// 添加maclist
type MacListCreateQuery struct {
	Mac    string // mac地址
	Name   string // mac名称
	Signal int64  // 信号值
	Lat    string // 纬度
	Lng    string // 经度
	Remark string // 备注
	Status int64  // 状态
}

// 获取maclist列表
type MacListsGetQuery struct {
	Page      int32  // 页码	范围为1-5000
	Limit     int32  // 每页条数	范围为1-100
	StartTime string // 开始时间	格式为2006-01-01 00:00:00
	EndTime   string // 结束时间	格式为2006-01-01 00:00:00
	Mac       string // mac地址
}

// 删除maclist
type MacListDeleteQuery struct {
	MacAddr string // primary_key的值
	Id      int64  // ID
	Mac     string // mac地址
}
