package result

// maclist列表
type MacBooksGetResult struct {
	Result
	Data struct {
		Total  int                      `json:"total"`
		Result []map[string]interface{} `json:"result"`
	} `json:"data,omitempty"` // 返回结果
}

type MacbookEntity struct {
	Mac        string  `json:"mac"`
	Name       string  `json:"name"`
	ProvinceId int64   `json:"province_id"`
	CityId     int64   `json:"city_id"`
	DistrictId int64   `json:"district_id"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Status     int64   `json:"status"`
	TotalCount int64   `json:"total_count"`
	Address    string  `json:"address"`
	CreateTime string  `json:"created"`
	Remark     string  `json:"remark"`
}

// maclist详情
type MacBookGetResult struct {
	Result
	Data MacbookEntity `json:"data,omitempty"` // 返回结果
}
