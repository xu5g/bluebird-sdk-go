package tspsdk

import (
	"fmt"
	"github.com/xu5g/bluebird-sdk-go/query"
	"testing"
)

// 最新血压数据
func TestBlood_GetBlood(t *testing.T) {
	query := &query.BloodGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}
	res, err := NewClient(gateWay, appKey, token).Blood().GetBlood(query)
	fmt.Println(res, err)
}

// 血压数据列表
func TestBlood_GetBloods(t *testing.T) {
	query := &query.BloodsGetQuery{
		Page: 1,
		Limit: 10,
	}
	res, err := NewClient(gateWay, appKey, token).Blood().GetBloods(query)
	fmt.Println(res, err)
}

// 获取血压上报间隔
func TestBlood_GetBloodUpload(t *testing.T) {
	query := &query.BloodUploadGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}
	res, err := NewClient(gateWay, appKey, token).Blood().GetBloodUpload(query)
	fmt.Println(res, err)
}

// 设置血压上报间隔
func TestBlood_UpdateBloodUpload(t *testing.T) {
	query := &query.BloodUploadSetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Second: 500,
	}
	res, err := NewClient(gateWay, appKey, token).Blood().UpdateBloodUpload(query)
	fmt.Println(res, err)
}