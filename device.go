package tspsdk

type Device struct {
	Cfg *Config
}

// 创建设备
func (p *Device) Create() {
}

// 判断设备是否在线
func (p *Device) IsOnline() (bool,error){
	return false, nil
}

// 设备详情
func (p *Device) GetOne(imeiSn string) (bool,error){

	return false, nil
}
