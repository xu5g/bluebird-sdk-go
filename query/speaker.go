package query

// 音箱-文本转语音播放（tsp.speaker.speech.text，POST /tsp/speaker/speech/text）
// lyds IP网络广播系统 TTS；TSP 侧 queue=true 写死，多次调用自动排队顺播
// 注意：对外接口文档把 loop_times/loop_gap 含义写反了（已确认文档有误），以 TSP 代码为准：
// LoopTimes=次数、LoopDuration=时长秒、LoopGap=间歇秒；LoopTimes 与 LoopDuration 二选一，不可同时传
type SpeakerPlayTextQuery struct {
	ImeiSn       string // 设备号（即广播系统设备编码device_code，播报终端），必填，长度不超过20
	Text         string // 要播放的文字内容，必填，长度1至500个字符
	Vcn          string // 发音人 xiaoyan女声(默认)/xiaofeng男声
	Speed        int    // 语速 30最慢/40较慢/50正常(默认)/70较快
	Volume       int    // 播放音量0-100
	Prompt       bool   // true播放前奏提示叮咚音
	LoopTimes    int    // 循环播放次数(次)，0不循环；与LoopDuration二选一
	LoopDuration int    // 循环播放时长(秒)，0不限制；与LoopTimes二选一
	LoopGap      int    // 循环播放间歇时间(秒)
}