package result

type BloodEntity struct {
	PrimaryKey string `json:"primary_key"`
	Id         int64  `json:"id"`
	ImeiSn     string `json:"imei_sn"`
	Uuid       string `json:"uuid"`
	ProductId  int64  `json:"product_id"`
	Blood      string `json:"temperature"`
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
