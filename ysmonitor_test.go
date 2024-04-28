package tspsdk

import (
	"fmt"
	"testing"
)

// 轨迹数据列表
func TestYsmonitor_GetToken(t *testing.T) {
	res := NewClient(gateWay, appKey, token).YsMonitor().TSPYsmonitorAccesstokenPath()
	fmt.Println(res.Data)
}
