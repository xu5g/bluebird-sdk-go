package query

// 获取设备详情
type DeviceGetQuery struct {
	ImeiSn    string // 设备Imei号	长度不超过20
	AttenceSn string // 设备考勤号	长度不超过10
}

// 创建设备
type DeviceCreateQuery struct {
	ImeiSn       string // 设备Imei号 长度不超过20
	ModelId      int64  // 设备型号ID
	AppKey       int64  // 应用ID
	AttenceSn    string // 2.4G考勤号
	IccardSn     string // 13.56考勤号
	PartnerId    int64
	Engine       string
	LocateUpload int64
}

// 更新设备信息
type DeviceUpdateQuery struct {
	ImeiSn    string // 设备Imei号 长度不超过20
	Truename  string // 设备名称 长度不超过50个字符
	ModelId   int64  // 设备型号ID
	Mobile    string // 手机号 长度11
	AppKey    int64  // 应用ID
	AttenceSn string // 2.4G考勤号
	IccardSn  string // 13.56考勤号
	PartnerId int64  // 企业ID
}

// 获取设备列表
type DevicesGetQuery struct {
	Page      int32  // 页码 范围为1-5000
	Limit     int32  // 每页条数 范围为1-200
	ImeiSn    string // 设备Imei号 长度不超过20
	Uuid      string // 设备Uuid 长度不超过40
	Mobile    string // 手机号 长度11
	AttenceSn string // 设备考勤号 长度不超过10
	IsOnline  string // 是否在线 1在线
	ModelId   int64  // 设备型号ID
	AppKey    int64  // 应用ID
	PartnerId int64  // 合作方ID
}

// 下发定位指令到终端
type DeviceLocateQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
}

// 获取设备是否在线
type DeviceIsOnlineQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
}

// 透传报文
type DeviceMessageQuery struct {
	ImeiSn  string // 设备Imei号 长度不超过20
	Message string // 报文内容
}

// 获取设备功能清单
type DeviceModulesQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
}

// 绑定设备
type DeviceBindQuery struct {
	ImeiSn   string // 设备Imei号 长度不超过20
	TrueName string // 设备名称
	Mobile   string // 设备内SIM卡的手机号
	Uuid     string // 设备UUID
}

// 解绑设备
type DeviceUnBindQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
}

// 下发寻找设备指令
type DeviceFindQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
}

// 下发定位上报间隔指令
type DeviceLocateUploadQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
	Second int    // 必传 定位上报间隔，单位：秒
}

// 下发定位时间段指令
type DeviceUdtimeQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20必传
	Start  string // 开始时间，格式：06:00
	End    string // 结束时间，格式：22:00
}

// 下发设置亲情号码指令
type DeviceFamilyQuery struct {
	ImeiSn   string   // 设备Imei号 长度不超过20
	Families []Family // 亲情号号码信息
	Status   int      // 状态 0： 设置亲情号与sos 1：设置亲情号 2：设置sos
}
type Family struct {
	Relation string `json:"relation"` // 亲人名称
	Mobile   string `json:"mobile"`   // 亲人手机号
}

// 下发设置定位模式指令
type DeviceLocateModeQuery struct {
	ImeiSn       string // 设备Imei号 长度不超过20
	LocateMode   string // 定位模式 1：省电模式 2：智能模式 3：性能模式
	LocateUpload int64  // 定位上报
}

// 下发设置终端host指令
type DeviceHostQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
	Host   string // 主机地址
	Port   string // 端口号
}

// 下发关机指令
type DevicePowerOffQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
}

// 下发重启指令
type DeviceRestartQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
}

// 下发聆听指令
type DeviceMonitorQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
	Mobile string // 手机号
}

// 下发设置免打扰时间段指令
type DeviceDndQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
	Dnd    string // 免打扰时间段 格式：8:00-11:30|123456;14:00-17:30|12345 12345表示周一到周五生效，1234567表示每天都生效
}

// 变更设备状态
type DeviceStatusQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
	Status string // 状态：1：正常 2：弃用 3：停机
}

// 删除设备
type DeviceDeleteQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
}

// 设置睡眠时间段
type SendSleepTimeQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20
	Start  string // 开始时间，格式：22:00
	End    string // 结束时间，格式：07:00
}

// 下发传输微聊音频文件到设备的指令
type DeviceWechatQuery struct {
	ImeiSn         string // 设备Imei号 长度不超过20
	WechatAudioUrl string // amr音频文件的地址
}

// DeviceWhitelistStatus 变更通话白名单状态
type DeviceWhitelistStatus struct {
	ImeiSn          string // 设备Imei号 长度不超过20
	WhitelistStatus int    // 1：开启通话白名单，只有亲情号码才可来电  2：关闭通话白名单，即不阻止陌生来电
}

// DeviceBatchLocateMode 批量设置设备的定位模式指令
type DeviceBatchLocateMode struct {
	LocateMode interface{}
}
