package tspsdk

import (
	"fmt"
	"github.com/xu5g/bluebird-sdk-go/query"
	"testing"
)

// 轨迹数据列表
func TestTrack_GetTracks(t *testing.T) {
	query := &query.TracksGetQuery{
		StartTime: "2021-02-01 00:00:00",
		EndTime: "2021-04-30 23:00:00",
		Page: 1,
		Limit: 10,
		ImeiSn: "863659041488510",
	}

	res, err := NewClient(gateWay, appKey, token).Track().GetTracks(query)
	fmt.Println(res, err)
}
