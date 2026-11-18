package query

// kjy-设置学生信息（tsp.kjyphone.student，PUT /tsp/kjyphone/student）
// 开架易视频话机 person/create，批量/单个 添加与修改共用；删除学生见 PhoneStudentDeleteQuery
type PhoneStudentSetQuery struct {
	ImeiSn   string         // 设备Imei号 长度不超过20
	Students []PhoneStudent // 学生信息 单次最多200个（扁平数组；文档示例里的 [{students:[...],imei_sn:...}] 包裹形式与其参数表矛盾，TSP 按参数表解析，实测以扁平为准）
}

type PhoneStudent struct {
	StudentId     int64  `json:"student_id"`     // 学生id（设备侧vipID）必须大于0
	Name          string `json:"name"`           // 学生姓名 必填，长度1至20个字符
	Avatar        string `json:"avatar"`         // 学生面部图片链接（非必填）
	StudentNumber string `json:"student_number"` // 学号 可为空，不为空时校验长度1至20
	StudentClass  string `json:"student_class"`  // 班级 可为空，不为空时校验长度1至20
	IcCard        string `json:"ic_card"`        // 校徽卡号：文档标「可为空」，但 TSP 校验为 required|length:1,20，实际必填，空值会被整批驳回
}

// kjy-删除学生（tsp.kjyphone.student，DELETE /tsp/kjyphone/student）
// 开架易视频话机 person/delete
type PhoneStudentDeleteQuery struct {
	ImeiSn     string   // 设备Imei号 长度不超过20
	StudentIds []string // 学生id（设备侧vipID），不能为空且不含空值；文档写作逗号分隔字符串「1,2」，此处按 TSP 参数结构的 []string 传数组，TSP 侧 gstr.Join 后下发的报文与文档形态一致
}

// kjy-设置家长关系（tsp.kjyphone.parent，PUT /tsp/kjyphone/parent）
// 开架易视频话机 addParent，批量/单个 增删改一体，删除用 DelFlag=1 条目
type PhoneParentSetQuery struct {
	ImeiSn  string        // 设备Imei号 长度不超过20
	Parents []PhoneParent // 家长信息 单次最多200个
}

type PhoneParent struct {
	ParentId     int64  `json:"parent_id"`    // 家长id 必须大于0
	StudentId    int64  `json:"student_id"`   // 学生id 必须大于0
	Name         string `json:"name"`         // 家长名称（TSP 未校验，随文档字段保留）
	Relationship string `json:"relationship"` // 关系名称 必填，长度1至20个字符（DelFlag=1 时 TSP 仍校验，删除条目也要带上）
	Avatar       string `json:"avatar"`       // 家长头像：TSP 下载后转base64，必填（DelFlag=1 时同样校验）
	OpenId       string `json:"open_id"`      // 家长openid
	DelFlag      int    `json:"del_flag"`     // 0不删除 1删除
}

// kjy-家长向学生发送消息（tsp.kjyphone.parent.message，POST /tsp/kjyphone/parent/message）
// 开架易视频话机 sendParentMessage
type PhoneParentMessageQuery struct {
	ImeiSn    string // 设备Imei号 长度不超过20
	ParentId  int64  // 家长id 必须大于0
	StudentId int64  // 学生id 必须大于0
	Type      int    // 0文本消息 1语音消息
	Content   string // 文本内容 Type=0时必填，长度不超过200个字符
	AudioUrl  string // mp3格式的音频文件链接 Type=1时必填，且后缀必须为mp3
}
