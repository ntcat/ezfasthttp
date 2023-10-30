package client

import (
	"github.com/valyala/fasthttp"
	"time"
)

type FastClient struct {
	url          string
	req          *fasthttp.Request
	resp         *fasthttp.Response
	timeout      time.Duration
	body         BodyType
	dataMap      any
	dataMapArray []any
}

type ResJson struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data any    `json:"data"`
}

type ResJsonArray struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []any  `json:"data"`
}

func NewFastClient(url string) (f *FastClient) {
	f = &FastClient{}
	f.SetUrl(url)
	f.timeout = time.Second * 30
	f.req = fasthttp.AcquireRequest()
	f.resp = fasthttp.AcquireResponse()
	return
}

func (f *FastClient) GetUrl() string {
	return f.url
}
func (f *FastClient) SetUrl(url string) {
	f.url = url
}
func (f *FastClient) SetHeadReq(key, value string) {
	f.req.Header.Set(key, value)
}
func (f *FastClient) SetHeadResp(key, value string) {
	f.resp.Header.Set(key, value)
}
func (f *FastClient) SetUserAgent(value string) {
	f.req.Header.SetUserAgent(value)
}
func (f *FastClient) SetCookie(key, value string) {
	f.req.Header.SetCookie(key, value)
}
func (f *FastClient) GetBody() BodyType {
	return f.body
}
func (f *FastClient) SetBody(body BodyType) {
	f.body = body
}
func (f *FastClient) GetTimeout() time.Duration {
	return f.timeout
}
func (f *FastClient) SetTimeout(timeout time.Duration) {
	f.timeout = timeout
}
func (f *FastClient) BodyHandle(callback func(BodyType) (BodyType, error)) error {
	out, err := callback(f.body)
	f.body = out
	return err
}
func (f *FastClient) PraseJsonHandle(callback func(BodyType) (DataMapType, error)) error {
	out, err := callback(f.body)
	f.dataMap = out
	return err
}

func (f *FastClient) DataMapHandle(callback DataMapHandle) ([]any, error) {
	return callback(f.dataMap)
}

func (f *FastClient) PraseJsonArrayHandle(callback func(BodyType) ([]any, error)) error {
	out, err := callback(f.body)
	f.dataMapArray = out
	return err
}

func (f *FastClient) DataMapArrayHandle(callback DataMapArrayHandle) ([]any, error) {
	return callback(f.dataMapArray)
}
