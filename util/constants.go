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
	TSPTemperatureDeletePath    = "/tsp/temperature/delete"     // 删除体温数据

	/** 心率管理 **/
	TSPHeartGetPath       = "/tsp/heart/get"        // 获取最新心率数据
	TSPHeartsGetPath      = "/tsp/hearts/get"       // 获取心率结果列表
	TSPHeartUploadGetPath = "/tsp/heart/upload/get" // 获取心率测量间隔时间
	TSPHeartUploadSetPath = "/tsp/heart/upload/set" // 设置心率测量间隔时间
	TSPHeartDeletePath    = "/tsp/heart/delete"     // 删除心率数据

	/** 血压管理 **/
	TSPBloodGetPath       = "/tsp/blood/get"        // 获取最新血压数据
	TSPBloodsGetPath      = "/tsp/bloods/get"       // 获取血压结果列表
	TSPBloodUploadGetPath = "/tsp/blood/upload/get" // 获取血压测量间隔时间
	TSPBloodUploadSetPath = "/tsp/blood/upload/set" // 设置血压测量间隔时间
	TSPBloodDeletePath    = "/tsp/blood/delete"     // 删除血压数据

	/** 血氧管理 **/
	TSPBloodOxygenGetPath    = "/tsp/bloodoxygen/get"    // 获取最新血氧数据
	TSPBloodOxygensGetPath   = "/tsp/bloodoxygens/get"   // 获取血氧结果列表
	TSPBloodOxygenDeletePath = "/tsp/bloodoxygen/delete" // 删除血氧数据

	/** 计步管理 **/
	TSPStepsGetPath = "/tsp/steps/get" // 获取计步数据列表

	/** 睡眠管理 **/
	TSPSleepsGetPath   = "/tsp/sleeps/get"   // 获取睡眠数据列表
	TSPSleepGetPath    = "/tsp/sleep/get"    // 获取最近一次测量的睡眠数据
	TSPSleepDeletePath = "/tsp/sleep/delete" // 删除睡眠数据

	/** 轨迹管理 **/
	TSPTracksGetPath    = "/tsp/tracks/get"   // 获取轨迹
	TSPTracksUpdatePath = "/tsp/track/update" // 修改轨迹数据
	TSPTrackDeletePath  = "/tsp/track/delete" // 删除轨迹数据

	/** 设备管理 **/
	TSPDeviceGetPath                   = "/tsp/device/get"               // 获取设备详情
	TSPDeviceUpdatePath                = "/tsp/device/update"            // 更新设备信息
	TSPDeviceCreatePath                = "/tsp/device/create"            // 创建设备信息
	TSPDevicesGetPath                  = "/tsp/devices/get"              // 获取设备列表
	TSPDevicesBatchGetPath             = "/tsp/devices/batch/get"        // 批量获取设备列表
	TSPDeviceLocatePath                = "/tsp/device/locate"            // 下发定位指令到终端
	TSPDeviceOnlinePath                = "/tsp/device/online"            // 获取设备是否在线
	TSPDeviceMessagePath               = "/tsp/device/message"           // 向设备透传报文
	TSPDeviceModulesPath               = "/tsp/device/modules"           // 获取设备功能清单
	TSPDeviceBindPath                  = "/tsp/device/bind"              // 设备绑定
	TSPDeviceUnBindPath                = "/tsp/device/unbind"            // 设备解绑
	TSPDeviceFindPath                  = "/tsp/device/find"              // 下发寻找设备指令
	TSPDeviceLocateUploadPath          = "/tsp/device/locate/upload"     // 下发定位上报间隔指令
	TSPDeviceUdtimePath                = "/tsp/device/udtime"            // 下发定位时间段指令
	TSPDeviceFamilyPath                = "/tsp/device/family"            // 下发设置亲情号码指令
	TSPDeviceLocateModePath            = "/tsp/device/locate/mode"       // 下发设置定位模式指令
	TSPDeviceHostPath                  = "/tsp/device/host"              // 下发设置终端host指令
	TSPDevicePowerOffPath              = "/tsp/device/poweroff"          // 下发关机指令
	TSPDeviceRestartPath               = "/tsp/device/restart"           // 下发重启指令
	TSPDeviceMonitorPath               = "/tsp/device/monitor"           // 下发聆听指令
	TSPDeviceDndPath                   = "/tsp/device/dnd"               // 下发设置免打扰指令
	TSPDeviceStatusPath                = "/tsp/device/status"            // 更改设备状态
	TSPDeviceDeletePath                = "/tsp/device/delete"            // 删除设备
	TSPDeviceSleepTimePath             = "/tsp/device/sleeptime"         // 下发睡眠时间段指令
	TSPDeviceDeviceWechatPath          = "/tsp/device/wchat"             // 下发传输微聊音频文件到设备的指令
	TSPDeviceDeviceWhitelistStatusPath = "/tsp/device/whitelist/status"  // 变更通话白名单状态
	TSPDeviceBatchLocateMode           = "/tsp/device/batch/locate/mode" // 批量设置设备的定位模式指令
	TSPDeviceRemind                    = "/tsp/device/remind"            // 设置设备闹钟
	TSPDeviceCronshutdown              = "/tsp/device/cronshutdown"      // 设置设备定时开关机
	TspDeviceWordsmessage              = "/tsp/device/wordsmessage"      // 发送文字消息到设备
	TspDeviceWordbook                  = "/tsp/device/wordbook"          // 下发单词
	TspDevicePoetrybook                = "/tsp/device/poetrybook"        // 下发古诗词

	/** 围栏管理 **/
	TSPFencesGetPath   = "/tsp/fences/get"   // 获取围栏列表
	TSPFenceCreatePath = "/tsp/fence/create" // 添加围栏
	TspFenceDeletePath = "/tsp/fence/delete" // 删除围栏

	/** 定位纠偏管理 **/
	TspMacListCreatePath = "/tsp/maclist/create" // 添加maclist
	TspMacListsGetPath   = "/tsp/maclists/get"   // 获取maclist列表
	TspMacListDeletePath = "/tsp/maclist/delete" // 删除maclist
	TspMacBooksGetPath   = "/tsp/macbooks/get"   // 获取macbook列表
	TspMacBookGetPath    = "/tsp/macbook/get"    // 获取macbook详情
	TspMacBookUpdatePath = "/tsp/macbook/update" // 更新macbook
	TspMacBookDrawPath   = "/tsp/macbook/draw"   // 重绘macbook
	TspMacBookDeletePath = "/tsp/macbook/delete" // 删除macbook

	/** 考勤管理 **/
	TSPAttencesGetPath   = "/tsp/attences/get"   // 获取考勤数据列表
	TSPAttenceDeletePath = "/tsp/attence/delete" // 删除考勤数据

	/** 报文管理 **/
	TSPMessagesGetPath    = "/tsp/messages/get"   // 获取报文数据列表
	TSPMessagesDeletePath = "/tsp/message/delete" // 删除报文数据

	/** 日志管理 **/
	TSPApiLogsGetPath     = "/tsp/apilogs/get"     // 获取api日志列表
	TSPApiLogDeletePath   = "/tsp/apilog/delete"   // 删除api日志
	TSPCoreLogsGetPath    = "/tsp/corelogs/get"    // 获取core日志列表
	TSPCoreLogDeletePath  = "/tsp/corelog/delete"  // 删除core日志
	TSPGuardLogsGetPath   = "/tsp/guardlogs/get"   // 获取guard日志列表
	TSPGuardLogDeletePath = "/tsp/guardlog/delete" // 删除guard日志

	/** 获取萤石token **/
	TSPYsmonitorAccesstokenGetPath = "/tsp/ysmonitor/accesstoken" // 获取萤石token

)
