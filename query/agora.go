package query

// 更新音视频通话用户列表
type UpdateAgoraUsersQuery struct {
	ImeiSn string       // 设备Imei号 长度不超过20
	Agora  []AgoraUsers // 视频通话列表信息
}

type AgoraUsers struct {
	Uid     int64  `json:"uid"`      // 业务层用户唯一ID
	RelName string `json:"rel_name"` // 关系名称
	RelIcon string `json:"rel_icon"` // https头像地址
	Mobile  string `json:"mobile"`   // 手机号
	Status  int    `json:"status"`
}

// 下发APP邀请手表进行音视频通话指令
type AgoraAppcalldeviceQuery struct {
	ImeiSn  string
	Uid     int64
	RelName string
}

// 下发APP挂断视频通话
type AgoraApphangupdeviceQuery struct {
	ImeiSn string
}
