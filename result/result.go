package result

type Result struct {
	Status  int    `json:"status"`            // 必选,返回码
	Message string `json:"message"` // 可选，返回消息
}
