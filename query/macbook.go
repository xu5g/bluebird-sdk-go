package query

// 获取macbook列表
type MacBooksGetQuery struct {
	Page       int32  // 页码	范围为1-5000
	Limit      int32  // 每页条数	范围为1-100
	StartTime  string // 开始时间	格式为2006-01-01
	EndTime    string // 结束时间	格式为2006-01-01
	Sort       string // 排序方式 DESC/ASC；默认倒序
	Name       string // 按照MAC名称查询
	Mac        string // 按照Mac查询
	Status     string // 按照mac地址状态查询
	TotalCount int64  // 定位数量
	CityId     int64  // 按照城市id查询
}

// 获取macbook详情
type MacBookGetQuery struct {
	Mac string // mac地址
}

// 修改macbook
type MacbookUpdateQuery struct {
	Mac        string  // mac地址
	Name       string  // mac名称
	CityId     int64   // 市ID
	ProvinceId int64   // 省ID
	DistrictId int64   // 区ID
	Lat        float64 // 纬度
	Lng        float64 // 经度
	Status     int64   // mac状态
	TotalCount int64   // maclist条数
	Street     string  // 街道
	Address    string  // 地址
	Remark     string  // 备注
}

// 重绘macbook
type MacbookDrawQuery struct {
	Mac  string // mac地址
	Name string // mac名称
}

// 删除macbook
type MacbookDeleteQuery struct {
	MakAddr string // 主键
	Mac     string // mac名称
}
