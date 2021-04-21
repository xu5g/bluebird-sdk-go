package result

type ApiLogsGetResult struct {
	Result
	Data struct {
		Total  int           `json:"total"`
		Result []map[string]interface{} `json:"result"`
	} `json:"data,omitempty"` // 返回结果
}
