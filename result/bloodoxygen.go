package result

type BloodOxygenEntity struct {
	PrimaryKey      string `json:"primary_key"`
	Id              string `json:"id"`
	PartnerId       int64  `json:"partner_id"`
	AppKey          int64  `json:"appkey"`
	ImeiSn          string `json:"imei_sn"`
	Uuid            string `json:"uuid"`
	BloodOxygenRate int64  `json:"blood_oxygen_rate"`
	Created         string `json:"created"`
}

type BloodOxygenGetResult struct {
	Result
	Data BloodOxygenEntity `json:"data"` // 返回结果
}

type BloodOxygensGetResult struct {
	Result
	Data struct {
		Total  int                 `json:"total"`
		Result []BloodOxygenEntity `json:"result"`
	} `json:"data"` // 返回结果
}
