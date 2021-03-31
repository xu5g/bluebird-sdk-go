package result


type HeartEntity struct {
	PrimaryKey  string `json:"primary_key"`
	Id          int64  `json:"id"`
	ImeiSn      string `json:"imei_sn"`
	Uuid        string `json:"uuid"`
	ProductId   int64  `json:"product_id"`
	Heart string `json:"temperature"`
	CreateTime  string `json:"created"`
}

// 体温最近测量数据
type HeartResult struct {
	Result
	Data HeartEntity  `json:"data,omitempty"` // 返回结果
}

// 体温测量数据列表
type HeartsResult struct {
	Result
	Data struct{
		Total int	`json:"total"`
		Result []HeartEntity `json:"result"`
	}  `json:"data,omitempty"` // 返回结果
}

// 体温测量间隔时间
type HeartUploadResult struct {
	Result
	Data float64  `json:"data,omitempty"` // 返回结果
}