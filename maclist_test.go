package tspsdk

import (
	"fmt"
	"github.com/xu5g/bluebird-sdk-go/query"
	"testing"
)

// 添加maclist
func TestMacList_MacListCreate(t *testing.T) {
	params := &query.MacListCreateQuery{
		Name:   "xxxxxxxxxxxxxxxx",
		Mac:    "xxxxxxxxxxxxxxxx",
		Lat:    "35.097426",
		Lng:    "112.60424",
		Signal: 11,
	}

	res, err := NewClient(gateWay, appKey, token).MacList().MacListCreate(params)
	fmt.Println(err, res)
}

// 获取maclist列表
func TestMacList_GetMacLists(t *testing.T) {
	params := &query.MacListsGetQuery{
		Mac:       "xxxxxxxxxxxxxxxx",
		StartTime: "2021-04-12 00:00:00",
		EndTime:   "2021-04-22 00:00:00",
		Page: 1,
		Limit: 10,
	}

	res, err := NewClient(gateWay, appKey, token).MacList().GetMacLists(params)
	fmt.Println(err, res)
}

// 删除maclist
func TestMacList_(t *testing.T) {
	params := &query.MacListDeleteQuery{
		MacAddr: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Id: 0,
	}

	res, err := NewClient(gateWay, appKey, token).MacList().DeleteMacList(params)
	fmt.Println(err, res)
}
