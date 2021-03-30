package tspsdk

const (
	GateWay = "http://192.168.16.191:18001"
)

const (
	AuthURL = GateWay + "/auth/token" // 鉴权（auth）

	/** 体温管理 **/
	TSPTemperatureGetUrl       = GateWay + "/tsp/temperature"        // 获取最新体温数据
	TSPTemperaturesGetUrl      = GateWay + "/tsp/temperatures"       // 获取体温结果列表
	TSPTemperatureUploadGetUrl = GateWay + "/tsp/temperature/upload" // 获取体温测量间隔时间
	TSPTemperatureUploadPutUrl = GateWay + "/tsp/temperature/upload" // 设置体温测量间隔时间
)
