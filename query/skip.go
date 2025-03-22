package query

// 获取跳绳数据
type SkipsGetQuery struct {
	ImeiSn    string
	Uuid      string
	StartTime string
	EndTime   string
	Page      int32
	Limit     int32
	Sort      string
	AppKey    int64
}

type SetAcdateSkipRequest struct {
	ImeiSn   string
	Skipdate []Skip
}

type Skip struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Dates string `json:"dates"`
}

type SetAcuSkipRequest struct {
	ImeiSn    string
	Countdown int
}
