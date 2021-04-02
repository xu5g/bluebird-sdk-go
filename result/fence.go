package result

type FenceEntity struct {
	Id         int64  `json:"id,omitempty"`
	Created    string `json:"created,omitempty"`     // 测量时间
	ImeiSn     string `json:"imei_sn,omitempty"`     // 设备号
	Uuid       string `json:"uuid,omitempty"`        // uuid
	PrimaryKey string `json:"primary_key,omitempty"` // 主键
}

// 获取围栏列表
type FencesResult struct {
	Result
	Data struct {
		Total  int           `json:"total"`
		Result []FenceEntity `json:"result"`
	} `json:"data,omitempty"` // 返回结果
}

// 添加围栏
type FenceCreateResult struct {
	Result
	Data int64 `json:"data,omitempty"` // 围栏ID
}
