package client

type BodyType = []byte
type DataMapType = map[string]interface{}
type BodyHandle func(BodyType) (body BodyType, err error)
type DataMapHandle func(any) ([]any, error)
type DataMapArrayHandle func([]any) ([]any, error)
