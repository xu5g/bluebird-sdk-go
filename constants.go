package tspsdk

const (
	AuthPath = "/auth/token" // 鉴权（auth）

	/** 体温管理 **/
	TSPTemperatureGetPath       = "/tsp/temperature"        // 获取最新体温数据
	TSPTemperaturesGetPath      = "/tsp/temperatures"       // 获取体温结果列表
	TSPTemperatureUploadGetPath = "/tsp/temperature/upload" // 获取体温测量间隔时间
	TSPTemperatureUploadPutPath = "/tsp/temperature/upload" // 设置体温测量间隔时间

	/** 心率管理 **/
	TSPHeartGetPath       = "/tsp/heart"        // 获取最新心率数据
	TSPHeartsGetPath      = "/tsp/hearts"       // 获取心率结果列表
	TSPHeartUploadGetPath = "/tsp/heart/upload" // 获取心率测量间隔时间
	TSPHeartUploadPutPath = "/tsp/heart/upload" // 设置心率测量间隔时间
)
