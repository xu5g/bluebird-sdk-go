package result

type BloodEntity struct {
	PrimaryKey string `json:"primary_key"`
	Id         int64  `json:"id"`
	ImeiSn     string `json:"imei_sn"`
	Uuid       string `json:"uuid"`
	AppKey     int64  `json:"appkey"`
	BloodRate  string `json:"blood_rate"`
	MaxRate    int64  `json:"max_rate"`
	MinRate    int64  `json:"min_rate"`
	CreateTime string `json:"created"`
}


type BloodGetResult struct {
	Result
	Data BloodEntity `json:"data,omitempty"` // 返回结果
}

type BloodsGetResult struct {
	Result
	Data struct {
		Total  int           `json:"total"`
		Result []BloodEntity `json:"result"`
	} `json:"data,omitempty"` // 返回结果
}

// 血压测量间隔时间
type BloodUploadResult struct {
	Result
	Data float64  `json:"data,omitempty"` // 返回结果
}
