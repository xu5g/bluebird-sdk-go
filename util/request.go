package util

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"strings"
	"time"
)

type Request struct {
	GateWay string
	Token   string
	AppKey  string
	Secret  string
	Method  string
	Url     string
	Data    []interface{}
}

func (p *Request) SetMethod(method string) *Request {
	p.Method = strings.ToLower(method)
	return p
}

func (p *Request) SetUrl(url string) *Request {
	p.Url = url
	return p
}
func (p *Request) SetData(data ...interface{}) *Request {
	p.Data = data
	return p
}

func (p *Request) HttpRequest() (*gjson.Json, error) {
	if p.Method == "get" {
		return p.get()
	} else if p.Method == "post" {
		return p.post()
	} else if p.Method == "put" {
		return p.put()
	}
	return nil, nil
}

/**
 * @Desc GET
 */
func (p *Request) get() (*gjson.Json, error) {
	client := g.Client()
	client.SetTimeout(5 * time.Second)
	if p.Token != "" {
		client.SetHeader("token", p.Token)
	}
	res, err := client.Get(p.Url)

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
 */
func (p *Request) post() (*gjson.Json, error) {
	client := g.Client()
	client.SetTimeout(5 * time.Second)
	if p.Token != "" {
		client.SetHeader("token", p.Token)
	}
	res, err := client.Post(p.Url, p.Data)
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

/**
 * @Desc PUT
 */
func (p *Request) put() (*gjson.Json, error) {
	client := g.Client()
	client.SetTimeout(5 * time.Second)
	if p.Token != "" {
		client.SetHeader("token", p.Token)
	}
	res, err := client.Put(p.Url, p.Data)
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
