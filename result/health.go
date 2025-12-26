package result

type WaterrecordEntity struct {
	PrimaryKey string `json:"primary_key"`
	Id         string `json:"id"`
	ImeiSn     string `json:"imei_sn"`
	Uuid       string `json:"uuid"`
	Ml         string `json:"ml"`
	Created    string `json:"created"`
}

type WaterrecordsGetResult struct {
	Result
	Data struct {
		Total  int                 `json:"total"`
		Result []WaterrecordEntity `json:"result"`
	} `json:"data"` // 返回结果
}
