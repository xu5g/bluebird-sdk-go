package query

// 获取兔盯云日志列表
type TdcloudLogsGetQuery struct {
	TransId   string //流水序号
	EventId   string //事件ID
	PartnerId int64  //企业ID
	Appkey    int64  `v:"Appkey @integer#appkey为整数"` // 应用KEY
	ImeiSn    string //设备号
	PushType  int64  //推送类型1：HTTP；2：MQ
	StartTime string `v:"StartTime @required|date-format:Y-m-d H:i:s#时间不能为空|时间格式不正确"` // 查询记录开始时间
	EndTime   string `v:"EndTime @date-format:Y-m-d H:i:s#时间格式不正确"`                   // 查询记录结束时间
	Sort      string `v:"Sort @in:asc,desc#排序方式不正确"`                                  // 排序方向 DESC/ASC；默认倒序
	Page      int32  `v:"Page @min:1|max:5000#页码最小值为1|页码最大值为5000"`                    // 当前页码
	Limit     int32  `v:"Limit @min:1|max:100#每页最小值为1|每页最大值为100"`                     // 每页多少条
}
