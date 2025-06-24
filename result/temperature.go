package result

type TemperatureEntity struct {
	AppKey      int64  `json:"appkey"`
	CreateTime  string `json:"created"`
	Id          string `json:"id"`
	ImeiSn      string `json:"imei_sn"`
	PrimaryKey  string `json:"primary_key"`
	Temperature string `json:"temperature"`
	UserId      string `json:"user_id"`
	Uuid        string `json:"uuid"`
}

// 体温最近测量数据
type TemperatureResult struct {
	Result
	Data TemperatureEntity `json:"data"` // 返回结果
}

// 体温测量数据列表
type TemperaturesResult struct {
	Result
	Data struct {
		Total  int                      `json:"total"`
		Result []map[string]interface{} `json:"result"`
	} `json:"data"` // 返回结果
}

// 体温测量间隔时间
type TemperatureUploadResult struct {
	Result
	Data float64 `json:"data"` // 返回结果
}
