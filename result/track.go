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

// 获取轨迹
type TracksGetResult struct {
	Result
	Data struct {
		Total  int          `json:"total"`
		Result []TrackEntity `json:"result"`
	} `json:"data,omitempty"` // 返回结果
}
