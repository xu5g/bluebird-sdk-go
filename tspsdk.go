package tspsdk

type Config struct {
	HttpClient *Request
}

type Tsp struct {
	Cfg *Config
}

// 获取其它接口数据时在header里面携带token
func NewClient(token string) *Tsp {
	return &Tsp{
		Cfg: &Config{
			HttpClient: &Request{
				Token: token,
			},
		},
	}
}

// 获取token（只用来获取token）
func NewAuth(appKey, secret string) *Tsp {
	return &Tsp{
		Cfg: &Config{
			&Request{
				AppKey: appKey,
				Secret: secret,
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

