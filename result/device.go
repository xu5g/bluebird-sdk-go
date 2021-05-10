package result

type DeviceEntity struct {
	PrimaryKey        string         `json:"primary_key,omitempty"`        // 主键
	ImeiSn            string         `json:"imei_sn,omitempty"`            // 设备唯一编号
	ActiveTime        string         `json:"active_time,omitempty"`        // 激活时间
	AttenceSn         string         `json:"attence_sn,omitempty"`         // 2.4G考勤号
	BloodUpload       int64          `json:"blood_upload,omitempty"`       // 血压上报间隔
	Families          []familyMember `json:"families,omitempty"`           // 亲情号
	Sos               []familyMember `json:"sos,omitempty"`                // sos号码
	Dnd               []dndMember    `json:"dnd,omitempty"`                // 免打扰时间段
	Engine            string         `json:"engine,omitempty"`             // 协议引擎名称
	IccardSn          string         `json:"iccard_sn,omitempty"`          // 13.56考勤号
	IsOnline          int64          `json:"is_online,omitempty"`          // 是否在线
	AppKey            int64          `json:"appkey,omitempty"`             // =product_id
	LocateUpload      int64          `json:"locate_upload,omitempty"`      // 定位上报间隔
	Mobile            string         `json:"mobile,omitempty"`             // 设备内的手机号
	ModelId           int64          `json:"model_id,omitempty"`           // 设备型号ID
	HeartUpload       int64          `json:"heart_upload,omitempty"`       // 心率上报间隔
	TemperatureUpload int64          `json:"temperature_upload,omitempty"` // 体温上报间隔
	LocateType        int64          `json:"locate_type,omitempty"`        // 定位模式 1:GPS 2:BTS(基站) 3:WIFI 4:未知
	LocateTime        string         `json:"locate_time,omitempty"`        // 定位时间
	LocateMode        int64          `json:"locate_mode,omitempty"`        // 定位模式1：省电模式 2：智能模式 3：性能模式
	LocateDirection   int64          `json:"locate_direction,omitempty"`   // 定位角度
	LocateAddress     string         `json:"locate_address,omitempty"`     // 位置信息
	PowerRate         int64          `json:"power_rate,omitempty"`         // 剩余电量
	PowerStatus       int64          `json:"power_status,omitempty"`       // 剩余电量状态
	Truename          string         `json:"truename,omitempty"`           // 终端名称
	Uuid              string         `json:"uuid,omitempty"`               // 设备uuid
	Lng               float64        `json:"lng,omitempty"`                // 经度
	Lat               float64        `json:"lat,omitempty"`                // 维度
	Udtime            udtime         `json:"udtime,omitempty"`             // 定位时间段
	SleepTime		  sleepTime		 `json:"sleeptime,omitempty"`		   // 睡眠时间
	Status            int64          `json:"status,omitempty"`             // 设备状态 1：正常 2：弃用 3：停机
	GsmRate           int64          `json:"gsm_rate,omitempty"`           // 信号值

}

//设置睡眠时间段结构体
type sleepTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type familyMember struct {
	No     string `json:"no,omitempty"`
	Name   string `json:"name,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}
type dndMember struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
	Dates string `json:"dates,omitempty"`
}
type udtime struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

// 获取设备详情
type DeviceGetResult struct {
	Result
	Data DeviceEntity `json:"data,omitempty"` // 返回结果
}

// 获取设备列表
type DevicesResult struct {
	Result
	Data struct {
		Total  int                      `json:"total,omitempty"`
		Result []map[string]interface{} `json:"result,omitempty"`
	} `json:"data,omitempty"` // 返回结果
}

// 获取设备是否在线
type DeviceIsOnlineResult struct {
	Result
	Data struct {
		IsOnline int `json:"is_online,omitempty"`
	} `json:"data,omitempty"` // 返回结果
}

// 获取设备功能清单
type DeviceModulesResult struct {
	Result
	Data struct {
		Id      int64           `json:"id,omitempty"`
		Alias   string          `json:"alias,omitempty"`
		Modules []deviceModules `json:"modules,omitempty"`
	} `json:"data,omitempty"` // 返回结果
}
type deviceModules struct {
	TrueName string `json:"truename,omitempty"`
	Remark   string `json:"remark,omitempty"`
	Icon     string `json:"icon,omitempty"`
	Alias    string `json:"alias,omitempty"`
}
