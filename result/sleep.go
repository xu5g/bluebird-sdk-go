package result

type SleepEntity struct {
	PrimaryKey string `json:"primary_key"`
	Id         int64  `json:"id"`
	ImeiSn     string `json:"imei_sn"`
	Uuid       string `json:"uuid"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	CreateTime string `json:"created"`
	SleepType  int    `json:"sleep_type"`
	SleepTime  int    `json:"sleep_time"`
}

// 睡眠数据列表
type SleepsResult struct {
	Result
	Data struct {
		Total  int           `json:"total"`
		Result []SleepEntity `json:"result"`
	} `json:"data"`
}

//
type SleepResult struct {
	Result
	Data SleepEntity `json:"data"`
}
