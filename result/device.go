package result

type DeviceEntity struct {
	PrimaryKey        string         `json:"primary_key"`        // 主键
	ImeiSn            string         `json:"imei_sn"`            // 设备唯一编号
	ActiveTime        string         `json:"active_time"`        // 激活时间
	AttenceSn         string         `json:"attence_sn"`         // 2.4G考勤号
	BloodUpload       int64          `json:"blood_upload"`       // 血压上报间隔
	Families          []familyMember `json:"families"`           // 亲情号
	Sos               []familyMember `json:"sos"`                // sos号码
	Dnd               []dndMember    `json:"dnd"`                // 免打扰时间段
	Engine            string         `json:"engine"`             // 协议引擎名称
	IccardSn          string         `json:"iccard_sn"`          // 13.56考勤号
	IsOnline          int64          `json:"is_online"`          // 是否在线
	PartnerId         int64          `json:"partner_id"`         // 企业ID
	AppKey            int64          `json:"appkey"`             // =product_id
	LocateUpload      int64          `json:"locate_upload"`      // 定位上报间隔
	Mobile            string         `json:"mobile"`             // 设备内的手机号
	ModelId           int64          `json:"model_id"`           // 设备型号ID
	HeartUpload       int64          `json:"heart_upload"`       // 心率上报间隔
	TemperatureUpload int64          `json:"temperature_upload"` // 体温上报间隔
	LocateType        int64          `json:"locate_type"`        // 定位模式 1:GPS 2:BTS(基站) 3:WIFI 4:未知
	LocateTime        string         `json:"locate_time"`        // 定位时间
	LocateMode        int64          `json:"locate_mode"`        // 定位模式1：省电模式 2：智能模式 3：性能模式
	LocateDirection   int64          `json:"locate_direction"`   // 定位角度
	LocateAddress     string         `json:"locate_address"`     // 位置信息
	PowerRate         int64          `json:"power_rate"`         // 剩余电量
	PowerStatus       int64          `json:"power_status"`       // 剩余电量状态
	Truename          string         `json:"truename"`           // 终端名称
	Uuid              string         `json:"uuid"`               // 设备uuid
	Lng               float64        `json:"lng"`                // 经度
	Lat               float64        `json:"lat"`                // 维度
	Udtime            udtime         `json:"udtime"`             // 定位时间段
	SleepTime         sleepTime      `json:"sleeptime"`          // 睡眠时间
	Status            int64          `json:"status"`             // 设备状态 1：正常 2：弃用 3：停机
	GsmRate           int64          `json:"gsm_rate"`           // 信号值
	Created           string         `json:"created"`            // 时间
	WhitelistStatus   int64          `json:"whitelist_status"`   // 通话白名单状态 0：暂未使用 1：开启通话白名单 2：关闭通话白名单
	Iccid			  string		 `json:"iccid"`			     // iccid
}

//设置睡眠时间段结构体
type sleepTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type familyMember struct {
	No     string `json:"no"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}
type dndMember struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Dates string `json:"dates"`
}
type udtime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// 获取设备详情
type DeviceGetResult struct {
	Result
	Data DeviceEntity `json:"data"` // 返回结果
}

// 获取设备列表
type DevicesResult struct {
	Result
	Data struct {
		Total  int                      `json:"total"`
		Result []map[string]interface{} `json:"result"`
	} `json:"data"` // 返回结果
}

// 获取设备是否在线
type DeviceIsOnlineResult struct {
	Result
	Data struct {
		IsOnline int `json:"is_online"`
	} `json:"data"` // 返回结果
}

// 获取设备功能清单
type DeviceModulesResult struct {
	Result
	Data struct {
		Id      int64           `json:"id"`
		Alias   string          `json:"alias"`
		Modules []deviceModules `json:"modules"`
	} `json:"data"` // 返回结果
}
type deviceModules struct {
	TrueName string `json:"truename"`
	Remark   string `json:"remark"`
	Icon     string `json:"icon"`
	Alias    string `json:"alias"`
}
