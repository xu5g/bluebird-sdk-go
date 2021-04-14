package query

// 获取设备详情
type DeviceGetQuery struct {
	ImeiSn    string // 设备Imei号	长度不超过20（必传）
	AttenceSn string // 设备考勤号	长度不超过10
}

// 更新设备信息
type DeviceUpdateQuery struct {
	ImeiSn   string // 设备Imei号 长度不超过20（必传）
	Truename string // 设备名称 长度不超过50个字符（必传）
}

// 获取设备列表
type DevicesGetQuery struct {
	Page      int32  // 页码 范围为1-5000（必传）
	Limit     int32  // 每页条数 范围为1-200（必传）
	ImeiSn    string // 设备Imei号 长度不超过20
	Uuid      string // 设备Uuid 长度不超过40
	Mobile    string // 手机号 长度11
	AttenceSn string // 设备考勤号 长度不超过10
	IsOnline  string // 是否在线 1在线
	ModelId   int64  // 设备型号ID
}

// 下发定位指令到终端
type DeviceLocateQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
}

// 获取设备是否在线
type DeviceIsOnlineQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
}

// 透传报文
type DeviceMessageQuery struct {
	ImeiSn  string // 设备Imei号 长度不超过20（必传）
	Message string // 报文内容（必传）
}

// 获取设备功能清单
type DeviceModulesQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
}

// 绑定设备
type DeviceBindQuery struct {
	ImeiSn   string // 设备Imei号 长度不超过20（必传）
	TrueName string // 设备名称（必传）
	Mobile   string // 设备内SIM卡的手机号（必传）
	Uuid     string // 设备UUID（必传）
}

// 解绑设备
type DeviceUnBindQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
}

// 下发寻找设备指令
type DeviceFindQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
}

// 下发定位上报间隔指令
type DeviceLocateUploadQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
	Second int    // 必传 定位上报间隔，单位：秒（必传）
}

// 下发定位时间段指令
type DeviceUdtimeQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20必传（必传）
	Start  string // 开始时间，格式：06:00（必传）
	End    string // 结束时间，格式：22:00（必传）
}

// 下发设置亲情号码指令
type DeviceFamilyQuery struct {
	ImeiSn   string   // 设备Imei号 长度不超过20（必传）
	Families []Family // 亲情号号码信息（必传）
}
type Family struct {
	Relation string `json:"relation"` // 亲人名称
	Mobile   string `json:"mobile"`   // 亲人手机号
}

// 下发设置定位模式指令
type DeviceLocateModeQuery struct {
	ImeiSn     string // 设备Imei号 长度不超过20（必传）
	LocateMode string // 定位模式 1：省电模式 2：智能模式 3：性能模式 （笔必传）
}

// 下发设置终端host指令
type DeviceHostQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
	Host   string // 主机地址（必传）
	Port   string // 端口号（必传）
}

// 下发关机指令
type DevicePowerOffQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
}

// 下发重启指令
type DeviceRestartQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
}

// 下发聆听指令
type DeviceMonitorQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
	Mobile string // 手机号（必传）
}

// 下发设置免打扰时间段指令
type DeviceDndQuery struct {
	ImeiSn string // 设备Imei号 长度不超过20（必传）
	Dnd    string // 免打扰时间段 格式：8:00-11:30|123456;14:00-17:30|12345 12345表示周一到周五生效，1234567表示每天都生效（必传）
}
