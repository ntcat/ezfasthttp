package client

func (f *FastClient) PostJsonDo(jsonBody string, bodyHandle func(BodyType) (body BodyType, err error),
	praseJsonCustomHandle func(BodyType) (DataMapType, error),
	dataMapHandle func(DataMapType) ([]any, error)) (result []any, err error) {
	if err = f.RequestPostJson(jsonBody); err != nil {
		return nil, err
	}
	if bodyHandle != nil {
		if err = f.BodyHandle(bodyHandle); err != nil {
			return nil, err
		}
	}
	praseJsonFn := PraseJsonCommonHandle
	if praseJsonCustomHandle != nil {
		praseJsonFn = praseJsonCustomHandle
	}
	if err = f.PraseJsonHandle(praseJsonFn); err != nil {
		return nil, err
	}
	if dataMapHandle != nil {
		if result, err = f.DataMapHandle(dataMapHandle); err != nil {
			return nil, err
		}
	}

	return result, nil

}
