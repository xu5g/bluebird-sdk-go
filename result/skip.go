package result

type SkipEntity struct {
	PrimaryKey        string `json:"primary_key"`
	Id                int64  `json:"id"`
	ImeiSn            string `json:"imei_sn"`
	Uuid              string `json:"uuid"`
	PerSkipNum        int64  `json:"per_skip_num"`        //预置跳绳个数
	ActualSkipNum     int64  `json:"actual_skip_num"`     //实际跳绳个数
	PerSkipTime       int64  `json:"per_skip_time"`       //预置跳绳时间，单位：秒
	ActualSkipTime    int64  `json:"actual_skip_time"`    //实际跳绳时间，单位：秒
	Calorie           int64  `json:"calorie"`             //消耗卡路里
	ActualSkipCreated string `json:"actual_skip_created"` //实际跳绳的起跳时间
	Created           string `json:"created"`             //数据创建时间
}

// 跳绳数据列表
type SkipsResult struct {
	Result
	Data struct {
		Total  int          `json:"total"`
		Result []SkipEntity `json:"result"`
	} `json:"data"`
}
