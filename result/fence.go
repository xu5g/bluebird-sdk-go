package result

type FenceEntity struct {
	Id         int64  `json:"id"`
	Created    string `json:"created"`     // 测量时间
	ImeiSn     string `json:"imei_sn"`     // 设备号
	Uuid       string `json:"uuid"`        // uuid
	PrimaryKey string `json:"primary_key"` // 主键
	TrueName   string `json:"truename"`
}

// 获取围栏列表
type FencesGetResult struct {
	Result
	Data struct {
		Total  int                      `json:"total"`
		Result []map[string]interface{} `json:"result"`
	} `json:"data"` // 返回结果
}

// 添加围栏
type FenceCreateResult struct {
	Result
	Data int64 `json:"data"` // 围栏ID
}
