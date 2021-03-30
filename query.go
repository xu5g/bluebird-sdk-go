package tspsdk

/*================== 体温 ====================*/
type TemperatureRecentQuery struct {
	ImeiSn string
	Uuid   string
}

type TemperaturesQuery struct {
	StartTime string // 开始时间
	EndTime   string // 结束时间
	ImeiSn    string // 设备号
	Uuid      string // UUID
	Page      int32  // 页码
	Limit     int32  // 每页条数
	ProductId int64  // 按照ProductId查询
	Sort      string // 排序方向 DESC/ASC；默认倒序
}

type TemperatureUploadQuery struct {
	ImeiSn string // 设备号
}