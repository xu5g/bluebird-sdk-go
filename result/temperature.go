package result


type TemperatureEntity struct {
	PrimaryKey  string `json:"primary_key"`
	Id          int64  `json:"id"`
	ImeiSn      string `json:"imei_sn"`
	Uuid        string `json:"uuid"`
	AppKey      int64  `json:"appkey"`
	Temperature string `json:"temperature"`
	CreateTime  string `json:"create_time"`
}

// 体温最近测量数据
type TemperatureResult struct {
	Result
	Data TemperatureEntity  `json:"data,omitempty"` // 返回结果
}

// 体温测量数据列表
type TemperaturesResult struct {
	Result
	Data struct{
		Total int	`json:"total"`
		Result []TemperatureEntity `json:"result"`
	}  `json:"data,omitempty"` // 返回结果
}

// 体温测量间隔时间
type TemperatureUploadResult struct {
	Result
	Data float64  `json:"data,omitempty"` // 返回结果
}