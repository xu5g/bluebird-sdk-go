package result


type BloodEntity struct {
	PrimaryKey  string `json:"primary_key"`
	Id          int64  `json:"id"`
	ImeiSn      string `json:"imei_sn"`
	Uuid        string `json:"uuid"`
	ProductId   int64  `json:"product_id"`
	Blood string `json:"temperature"`
	CreateTime  string `json:"created"`
}

// 血压最近测量数据
type BloodResult struct {
	Result
	Data BloodEntity  `json:"data,omitempty"` // 返回结果
}

// 血压测量数据列表
type BloodsResult struct {
	Result
	Data struct{
		Total int	`json:"total"`
		Result []BloodEntity `json:"result"`
	}  `json:"data,omitempty"` // 返回结果
}

// 血压测量间隔时间
type BloodUploadResult struct {
	Result
	Data float64  `json:"data,omitempty"` // 返回结果
}