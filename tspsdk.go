package tspsdk

import "github.com/xu5g/bluebird-sdk-go/util"

type Config struct {
	HttpClient *util.Request
}

type Tsp struct {
	Cfg *Config
}

// 获取其它接口数据时在header里面携带token
func NewClient(gateWay, token string) *Tsp {
	return &Tsp{
		Cfg: &Config{
			HttpClient: &util.Request{
				Token: token,
				GateWay: gateWay,
			},
		},
	}
}

// 获取token（只用来获取token）
func NewAuth(gateWay, appKey, secret string) *Tsp {
	return &Tsp{
		Cfg: &Config{
			&util.Request{
				AppKey: appKey,
				Secret: secret,
				GateWay: gateWay,
			},
		},
	}
}

// 鉴权
func (p *Tsp) Auth() *Auth {
	return &Auth{
		Cfg: p.Cfg,
	}
}


// 体温
func (p *Tsp) Temperature() *Temperature {
	return &Temperature{
		Cfg: p.Cfg,
	}
}

// 心率
func (p *Tsp) Heart() *Heart {
	return &Heart{
		Cfg: p.Cfg,
	}
}

// 血压
func (p *Tsp) Blood() *Blood {
	return &Blood{
		Cfg: p.Cfg,
	}
}

// 计步
func (p *Tsp) Step() *Step {
	return &Step{
		Cfg: p.Cfg,
	}
}

// 睡眠
func (p *Tsp) Sleep() *Sleep {
	return &Sleep{
		Cfg: p.Cfg,
	}
}

