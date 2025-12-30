package result

type WaterrecordEntity struct {
	PrimaryKey string `json:"primary_key"`
	Id         string `json:"id"`
	ImeiSn     string `json:"imei_sn"`
	Uuid       string `json:"uuid"`
	Ml         int    `json:"ml"`
	WaterType  int    `json:"water_type"`
	Created    string `json:"created"`
}

type WaterrecordsGetResult struct {
	Result
	Data struct {
		Total  int                 `json:"total"`
		Result []WaterrecordEntity `json:"result"`
	} `json:"data"` // 返回结果
}
