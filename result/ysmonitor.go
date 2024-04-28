package result

type YsmonitorTokenGetResult struct {
	Result
	Data map[string]interface{} `json:"data"` // 返回结果
}
