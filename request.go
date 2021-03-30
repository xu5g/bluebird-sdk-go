package tspsdk

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"strings"
	"time"
)

type Request struct {
	Token  string
	Method string
	Url    string
	AppKey string
	Secret string
}

func (p *Request) SetToken(token string) *Request {
	p.Token = token
	return p
}

func (p *Request) SetMethod(method string) *Request {
	p.Method = strings.ToLower(method)
	return p
}

func (p *Request) SetUrl(url string) *Request {
	p.Url = url
	return p
}

func (p *Request) HttpRequest() (*gjson.Json, error) {
	if p.Method == "get" {
		return p.get(p.Url)
	} else if p.Method == "post" {
		return p.post(p.Url)
	}
	return nil, nil
}

/**
 * @Desc get请求方式
 * @Date 2021/3/12
 */
func (p *Request) get(url string) (*gjson.Json, error) {
	client := g.Client()
	client.SetTimeout(5 * time.Second)
	if p.Token != "" {
		client.SetHeader("token", p.Token)
	}
	res, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	context := res.ReadAllString()
	json := gjson.New(context)
	status := json.GetVar("status").String()
	message := json.GetVar("message").String()
	if status == "0" {
		return json, nil
	} else {
		return nil, errors.New(message)
	}
}

/**
 * @Desc POST
 * @Date 2021/3/12
 */
func (p *Request) post(url string, data ...interface{}) (*gjson.Json, error) {
	client := g.Client()
	client.SetTimeout(5 * time.Second)
	if p.Token != "" {
		client.SetHeader("token", p.Token)
	}
	res, err := client.Post(url, data)
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	if err != nil {
		return nil, err
	}
	context := res.ReadAllString()
	json := gjson.New(context)
	status := json.GetVar("status").String()
	message := json.GetVar("message").String()
	if status == "0" {
		return json, nil
	} else {
		return nil, errors.New(message)
	}
}
