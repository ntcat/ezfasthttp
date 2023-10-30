package client

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

func (f *FastClient) RequestGet() error {
	f.req.SetRequestURI(f.url)
	f.req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := fasthttp.DoTimeout(f.req, resp, f.timeout)
	fasthttp.ReleaseRequest(f.req)
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

func (f *FastClient) RequestPostJson(jsonBody string) error {
	f.req.Header.SetMethod(fasthttp.MethodPost)
	f.req.Header.SetContentType("application/json")
	f.req.SetBodyString(jsonBody)
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

func PraseJsonCommonHandle(body BodyType) (dataMap DataMapType, err error) {
	var r ResJson
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}

	if r.Code != 0 {
		err = fmt.Errorf("r.Code=%d,r.Msg=%s", r.Code, r.Msg)
	} else {
		if r.Data != nil {
			dataMap = (r.Data).(DataMapType)
		}
	}
	return
}

func PraseJsonArrayCommonHandle(body BodyType) (dataMap []any, err error) {
	var r ResJsonArray
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}

	if r.Code != 0 {
		err = fmt.Errorf("r.Code=%d,r.Msg=%s", r.Code, r.Msg)
	} else {
		if r.Data != nil {
			for _, dm := range r.Data {
				dataMap = append(dataMap, dm.(DataMapType))
			}
		}
	}
	return
}

// type PraseJsonFn func(body BodyType) (dataMap DataMapType, err error)

// type PraseJsonFn1 interface {
// 	PraseJsonCommonHandle(body BodyType) (dataMap DataMapType, err error)
// 	PraseJsonArrayCommonHandle(body BodyType) (dataMap []any, err error)
// 	praseJsonCustomHandle(body BodyType) (dataMap []any, err error)
// }
