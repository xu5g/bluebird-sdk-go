package util

const SdkVersion = "1.0.0"

const (
	// 身份认证
	TSPAuthPath = "/tsp/auth/token" // 鉴权（auth）

	/** 体温管理 **/
	TSPTemperatureGetPath       = "/tsp/temperature/get"        // 获取最新体温数据
	TSPTemperaturesGetPath      = "/tsp/temperatures/get"       // 获取体温结果列表
	TSPTemperatureUploadGetPath = "/tsp/temperature/upload/get" // 获取体温测量间隔时间
	TSPTemperatureUploadSetPath = "/tsp/temperature/upload/set" // 设置体温测量间隔时间

	/** 心率管理 **/
	TSPHeartGetPath       = "/tsp/heart/get"        // 获取最新心率数据
	TSPHeartsGetPath      = "/tsp/hearts/get"       // 获取心率结果列表
	TSPHeartUploadGetPath = "/tsp/heart/upload/get" // 获取心率测量间隔时间
	TSPHeartUploadSetPath = "/tsp/heart/upload/set" // 设置心率测量间隔时间

	/** 血压管理 **/
	TSPBloodGetPath  = "/tsp/blood/get"  // 获取最新血压数据
	TSPBloodsGetPath = "/tsp/bloods/get" // 获取血压结果列表

	/** 计步管理 **/
	TSPStepsGetPath = "/tsp/steps/get" // 获取计步数据列表

	/** 睡眠管理 **/
	TSPSleepsGetPath = "/tsp/sleeps/get" // 获取睡眠数据列表

	/** 轨迹管理 **/
	TSPTracksGetPath = "/tsp/tracks/get" // 获取轨迹

	/** 设备管理 **/
	TSPDeviceGetPath          = "/tsp/device/get"           // 获取设备详情
	TSPDevicesGetPath         = "/tsp/devices/get"          // 获取设备列表
	TSPDeviceLocatePath       = "/tsp/device/locate"        // 下发定位指令到终端
	TSPDeviceOnlinePath       = "/tsp/device/online"        // 获取设备是否在线
	TSPDeviceMessagePath      = "/tsp/device/message"       // 向设备透传报文
	TSPDeviceModulesPath      = "/tsp/device/modules"       // 获取设备功能清单
	TSPDeviceBindPath         = "/tsp/device/bind"          // 设备绑定
	TSPDeviceUnBindPath       = "/tsp/device/unbind"        // 设备解绑
	TSPDeviceFindPath         = "/tsp/device/find"          // 下发寻找设备指令
	TSPDeviceLocateUploadPath = "/tsp/device/locate/upload" // 下发定位上报间隔指令
	TSPDeviceUdtimePath       = "/tsp/device/udtime"        // 下发定位时间段指令
	TSPDeviceFamilyPath       = "/tsp/device/family"        // 下发设置亲情号码指令
	TSPDeviceLocateModePath   = "/tsp/device/locatemode"    // 下发设置定位模式指令
	TSPDeviceHostPath         = "/tsp/device/host"          // 下发设置终端host指令
	TSPDevicePowerOffPath     = "/tsp/device/poweroff"      // 下发关机指令
	TSPDeviceRestartPath      = "/tsp/device/restart"       // 下发重启指令

	/** 围栏管理 **/
	TSPFencesGetPath   = "/tsp/fences/get"   // 获取围栏列表
	TSPFenceCreatePath = "/tsp/fence/create" // 添加围栏
	TspFenceDeletePath = "/tsp/fence/delete" // 删除围栏

)
