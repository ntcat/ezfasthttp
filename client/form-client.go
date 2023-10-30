package client

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func (f *FastClient) RequestPostForm(formBody string) error {
	f.req.Header.SetMethod(fasthttp.MethodPost)
	f.req.Header.SetContentType("multipart/form-data")
	f.req.SetBodyString(formBody)
	f.req.SetRequestURI(f.url)
	err := fasthttp.DoTimeout(f.req, f.resp, f.timeout)
	fasthttp.ReleaseRequest(f.req)
	status := f.resp.StatusCode()
	f.body = f.resp.Body()
	if err != nil {
		return err
	}
	if status != fasthttp.StatusOK {
		return fmt.Errorf("request failed,status:%d", status)
	}
	fasthttp.ReleaseResponse(f.resp)

	return nil
}
