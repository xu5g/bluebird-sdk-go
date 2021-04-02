package query

// 获取设备详情
type DeviceGetQuery struct {
	ImeiSn    string // 设备Imei号	长度不超过20（PS: ImeiSn和AttenceSn不能同时为空）
	AttenceSn string // 设备考勤号	长度不超过10
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
	ImeiSn string	// 设备Imei号 长度不超过20（必传）
}

// 获取设备是否在线
type DeviceIsOnlineQuery struct {
	ImeiSn string // 设备号 长度不超过20（必传）
}

// 透传报文
type DeviceMessageQuery struct {
	ImeiSn  string // 设备号	设备号 长度不超过20（必传）
	Message string // 报文内容（必传）
}

type DeviceModulesQuery struct {
	ImeiSn string //not null
}

type DeviceBindQuery struct {
	ImeiSn   string
	TrueName string
	Mobile   string
	Uuid     string
}

type DeviceUnBindQuery struct {
	ImeiSn string
}

type DeviceFindQuery struct {
	ImeiSn string
}

// 下发定位上报间隔指令
type DeviceLocateUploadQuery struct {
	ImeiSn string // 必传 设备号
	Second int    // 必传 定位上报间隔，单位：秒
}

// 下发定位时间段指令
type DeviceUdtimeQuery struct {
	ImeiSn string // 必传
	Start  string // 必传 开始时间，例如：06:00
	End    string // 必传	结束时间，例如：22:00
}

// 下发设置亲情号码指令
type DeviceFamilyQuery struct {
	ImeiSn   string   // 必传
	Families []family // 必传
}
type family struct {
	Relation string // 亲人名称
	Mobile   string // 亲人手机号
}

type DeviceLocateModeQuery struct {
	ImeiSn     string
	LocateMode string
}

type DeviceHostQuery struct {
	ImeiSn string
	Host   string
	Post   string
}

// 下发关机指令
type DevicePowerOffQuery struct {
	ImeiSn string
	Host   string
	Post   string
}

// 下发重启指令
type DeviceRestartQuery struct {
	ImeiSn string // 必传
}
