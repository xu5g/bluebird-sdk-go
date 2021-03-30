package result

// 获取token
type AuthResult struct {
	Result
	Data    struct {
		Token string `json:"token"` //权限令牌，调用其它接口时需要提供token，有效期为 1 小时，过期后无法使用
	} `json:"data,omitempty"` // 返回结果
}
