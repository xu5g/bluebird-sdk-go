package result

type TrackEntity struct {
	ImeiSn       string `json:"imei_sn"`
	Lat          string `json:"lat"`
	Lng          string `json:"lng"`
	IsLrd        int64  `json:"is_lrd"`
	Height       int64  `json:"height"`
	ProvinceId   int64  `json:"province_id"`
	CityId       int64  `json:"city_id"`
	DistrictId   int64  `json:"district_id"`
	Speed        int64  `json:"speed"`
	Address      string `json:"address"`
	Direction    int64  `json:"direction"`
	Radius       int64  `json:"radius"`
	LocateType   int64  `json:"locate_type"`
	LocateTime   string `json:"locate_time"`
	IsIgnore     int64  `json:"is_ignore"`
	IgnoreReason string `json:"ignore_reason"`
	Message      string `json:"message"`
}

// 获取轨迹
type TracksGetResult struct {
	Result
	Data struct {
		Total  int           `json:"total"`
		Result []TrackEntity `json:"result"`
	} `json:"data,omitempty"` // 返回结果
}
