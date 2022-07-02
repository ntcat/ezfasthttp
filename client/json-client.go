package client

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

func (f *FastClient) RequestGet() error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(f.url)
	req.Header.SetMethod(fasthttp.MethodGet)
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

func (f *FastClient) RequestPostJson(jsonBody string) error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetBodyString(jsonBody)
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
