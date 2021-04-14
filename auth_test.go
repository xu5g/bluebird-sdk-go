package tspsdk

import (
	"fmt"
	"testing"
)

var (
	// 测试时请先获取最新的token，并将此处的token更新
	gateWay = "https://openapi.xu5g.com"
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBrZXkiOiIxMDAwIiwiY3JlYXRlZCI6IjIwMjEtMDQtMTMgMTQ6MjY6MTUiLCJleHBpcmVkIjoiMjAyMS0wNS0xMyAxNDoyNjoxNSIsInRyYW5zaWQiOiIxMDAwMjAyMTA0MTMxNDI2MTU0ODg3MzgyOTMzMjMifQ.ueR8DbGLeAoHC4rmjvl5Dw9HgsP-f-UB3DI3Viqf_hU"
	appKey = "xxxx"
)

// 获取token
func TestAuth_GetToken(t *testing.T) {

	res, err := NewAuth(gateWay, appKey, "xxxxxxxxxxxxxxx").Auth().GetToken()
	if err == nil {
		print(res.Data.Token)
		fmt.Println()
	} else {
		fmt.Println(err)
	}
}
