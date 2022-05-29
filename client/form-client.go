package client

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func (f *FastClient) RequestPostForm(formBody string) error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("multipart/form-data")
	req.SetBodyString(formBody)
	req.SetRequestURI(f.url)
	resp := fasthttp.AcquireResponse()
	err := fasthttp.DoTimeout(req, resp, f.timeout)
	fasthttp.ReleaseRequest(req)
	status := resp.StatusCode()
	f.body = resp.Body()
	if err != nil {
		return err
	}
	if status != fasthttp.StatusOK {
		return fmt.Errorf("request failed,status:%d", status)
	}
	fasthttp.ReleaseResponse(resp)

	return nil
}
