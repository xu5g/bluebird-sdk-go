package query

// 批量设置学生（开架易视频话机 person/create，添加和修改共用）
type PhoneStudentSetQuery struct {
	ImeiSn   string         // 设备Imei号 长度不超过20
	Students []PhoneStudent // 学生信息 单次最多200个
}

type PhoneStudent struct {
	StudentId     int64  `json:"student_id"`     // 学生id（设备侧vipID）必须大于0
	Name          string `json:"name"`           // 学生姓名 长度1至20个字符
	Avatar        string `json:"avatar"`         // 头像地址
	StudentNumber string `json:"student_number"` // 学号 长度1至20个字符
	StudentClass  string `json:"student_class"`  // 班级 长度1至20个字符
	IcCard        string `json:"ic_card"`        // 校徽卡号 长度1至20个字符
}

// 批量删除学生（开架易视频话机 person/delete）
type PhoneStudentDeleteQuery struct {
	ImeiSn     string   // 设备Imei号 长度不超过20
	StudentIds []string // 学生id（设备侧vipID）列表，不能为空且不存在空值
}

// 设置家长关系（开架易视频话机 addParent，增删改一体，删除用 DelFlag=1 条目）
type PhoneParentSetQuery struct {
	ImeiSn  string        // 设备Imei号 长度不超过20
	Parents []PhoneParent // 家长信息 单次最多200个
}

type PhoneParent struct {
	ParentId     int64  `json:"parent_id"`    // 家长id 必须大于0
	StudentId    int64  `json:"student_id"`   // 学生id 必须大于0
	Name         string `json:"name"`         // 家长昵称
	Relationship string `json:"relationship"` // 关系名称 长度1至20个字符
	Avatar       string `json:"avatar"`       // 头像地址，TSP侧下载后转base64下发
	OpenId       string `json:"open_id"`      // 家长openid
	DelFlag      int    `json:"del_flag"`     // 0新增修改 1删除
}

// 家长向学生发送消息（开架易视频话机 sendParentMessage）
type PhoneParentMessageQuery struct {
	ImeiSn    string // 设备Imei号 长度不超过20
	ParentId  int64  // 家长id 必须大于0
	StudentId int64  // 学生id 必须大于0
	Type      int    // 0文字 1语音
	Content   string // 文本内容 type=0时必填，长度不超过200个字符
	AudioUrl  string // 语音mp3的url地址 type=1时必填，后缀必须为mp3
}
