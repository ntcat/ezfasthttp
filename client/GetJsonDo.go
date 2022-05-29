package client

func (f *FastClient) GetJsonDo(bodyHandle func([]byte) (body []byte, err error),
	praseJsonCustomHandle func([]byte) (map[string]interface{}, error),
	dataMapHandle func(map[string]interface{}) ([]interface{}, error)) (result []interface{}, err error) {
	if err = f.RequestGet(); err != nil {
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
