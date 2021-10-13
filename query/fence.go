package query

// FencesGetQuery 获取围栏列表
type FencesGetQuery struct {
	Page      int    // 页码 范围为1-5000
	Limit     int    // 每页条数 范围为1-100
	Uuid      string // 设备UUID 长度不超过40
	FenceType int64  // 围栏类型 1：应用围栏；2：设备围栏
	ShapeType int64  // 围栏形状类型 1：点状围栏（圆形围栏）；2：多边形围栏
}

// FenceCreateQuery 添加围栏
type FenceCreateQuery struct {
	Truename    string // 围栏名称 长度不超过32
	FenceType   int64  // 围栏类型 1：应用围栏；2：设备围栏
	ShapeType   int64  // 围栏形状类型 1：点状围栏（圆形围栏）；2：多边形围栏
	CollideType int64  // 告警类型 1：进出告警；2：靠近告警；3：进出+靠近告警
	Points      interface{}
	PointRadius int64
	NearRadius  int64  // 靠近围栏预警半径（单位：米）
	ValidStart  string // 有效开始时间 PS: 08:00
	ValidEnd    string // 有效结束时间 PS: 18:00
	ValidWeek   string // 有效周天 周一至周日分别为1-7，逗号分割（1,2,3,4,5,6,7）
	Uuid        string // uuid
}

// FenceDeleteQuery 删除围栏
type FenceDeleteQuery struct {
	Id int64 // 主键ID
}
