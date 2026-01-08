package result

// 运动数据列表
type ExercisesResult struct {
	Result
	Data struct {
		Total  int                      `json:"total"`
		Result []map[string]interface{} `json:"result"`
	} `json:"data"` // 返回结果
}
