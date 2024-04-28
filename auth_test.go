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
	token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBrZXkiOiIyMDAwIiwiY3JlYXRlZCI6IjIwMjQtMDQtMjggMDk6NTg6MDEiLCJleHBpcmVkIjoiMjAyNC0wNS0yOCAwOTo1ODowMSIsInRyYW5zaWQiOiIyMDAwMjAyNDA0MjgwOTU4MDE5NjEyNDgxNjExOTQifQ.0qVzqJesH4NolatLE-7kOgtqRZFw2rJztsi2RYpM5VY"
	appKey = "2000"
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
