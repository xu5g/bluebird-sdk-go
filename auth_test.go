package tspsdk

import (
	"fmt"
	"testing"
)

var (
	// 测试时请先获取最新的token，并将此处的token更新
	gateWay = "https://test-openapi.xu5g.com"
	//gateWay = "http://127.0.0.1:18001"
	//token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBrZXkiOiIxMDAwIiwiY3JlYXRlZCI6IjIwMjEtMDQtMTMgMTQ6MjY6MTUiLCJleHBpcmVkIjoiMjAyMS0wNS0xMyAxNDoyNjoxNSIsInRyYW5zaWQiOiIxMDAwMjAyMTA0MTMxNDI2MTU0ODg3MzgyOTMzMjMifQ.ueR8DbGLeAoHC4rmjvl5Dw9HgsP-f-UB3DI3Viqf_hU"
	token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBrZXkiOiIxMDAwIiwiY3JlYXRlZCI6IjIwMjUtMDMtMjYgMDk6MjE6MTciLCJleHBpcmVkIjoiMjAyNS0wNC0yNSAwOToyMToxNyIsInRyYW5zaWQiOiIxMDAwMjAyNTAzMjYwOTIxMTY0Mjk3OTk5NDQ4NDUifQ.3Ai2x-F9qXviUR48ynjQuOoDLThd9HP6gYZaDedL0hI"
	appKey = "1000"
)

// 获取token
func TestAuth_GetToken(t *testing.T) {

	res, err := NewAuth(gateWay, appKey, "wnhMVKn3AoOXaldpDdgrAvo0srRArNCm").Auth().GetToken()
	if err == nil {
		print(res.Data.Token)
		fmt.Println()
	} else {
		fmt.Println(err)
	}
}
