package result

type StepEntity struct {
	PrimaryKey string `json:"primary_key"`
	Id         int64  `json:"id"`
	ImeiSn     string `json:"imei_sn"`
	Uuid       string `json:"uuid"`
	StepRate   int    `json:"step_rate"`
	CreateTime string `json:"create_time"`
}

// 计步数据列表
type StepsResult struct {
	Result
	Data struct {
		Total  int          `json:"total"`
		Result []StepEntity `json:"result"`
	} `json:"data,omitempty"` // 返回结果
}
