package tspsdk

type Result struct {
	Status  int    `json:"status"`            // 必选,返回码
	Message string `json:"message,omitempty"` // 可选，返回消息
}

// 获取token
type AuthResult struct {
	Result
	Data    struct {
		Token string `json:"token"` //权限令牌，调用其它接口时需要提供token，有效期为 1 小时，过期后无法使用
	} `json:"data,omitempty"` // 返回结果
}


/*====================== 体温 ==========================*/
type TemperatureEntity struct {
	PrimaryKey  string `json:"primary_key"`
	Id          int64  `json:"id"`
	ImeiSn      string `json:"imei_sn"`
	Uuid        string `json:"uuid"`
	ProductId   int64  `json:"product_id"`
	Temperature string `json:"temperature"`
	CreateTime  string `json:"created"`
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