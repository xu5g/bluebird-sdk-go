package result

type TrackEntity struct {
	PrimaryKey string `json:"primary_key"`
	Id         int64  `json:"id"`
	ImeiSn     string `json:"imei_sn"`
	Uuid       string `json:"uuid"`
	Address    string `json:"address"`
	isIgnore   int    `json:"is_ignore"`
	Lat        string `json:"lat"`
	Lng        string `json:"lng"`
}

// 轨迹数据列表
type TracksResult struct {
	Result
	Data struct {
		Total  int          `json:"total"`
		Result []StepEntity `json:"result"`
	} `json:"data,omitempty"` // 返回结果
}
