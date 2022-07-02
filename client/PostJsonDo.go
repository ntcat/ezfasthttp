package client

func (f *FastClient) PostJsonDo(
	jsonBody string,
	bodyHandle BodyHandle,
	dataMapHandle DataMapHandle) (result any, err error) {
	if err = f.RequestPostJson(jsonBody); err != nil {
		return nil, err
	}
	if bodyHandle != nil {
		if err = f.BodyHandle(bodyHandle); err != nil {
			return nil, err
		}
	}
	if err = f.PraseJsonHandle(PraseJsonCommonHandle); err != nil {
		return nil, err
	}

	if dataMapHandle != nil {
		if result, err = f.DataMapHandle(dataMapHandle); err != nil {
			return nil, err
		}
	} else {
		result = f.dataMap
	}

	return result, nil

}

func (f *FastClient) PostJsonArrayDo(
	jsonBody string,
	bodyHandle BodyHandle,
	dataMapArrayHandle DataMapArrayHandle) (result []any, err error) {
	if err = f.RequestPostJson(jsonBody); err != nil {
		return nil, err
	}
	if bodyHandle != nil {
		if err = f.BodyHandle(bodyHandle); err != nil {
			return nil, err
		}
	}
	if err = f.PraseJsonArrayHandle(PraseJsonArrayCommonHandle); err != nil {
		return nil, err
	}

	if dataMapArrayHandle != nil {
		if result, err = f.DataMapArrayHandle(dataMapArrayHandle); err != nil {
			return nil, err
		}
	} else {
		result = f.dataMapArray
	}

	return result, nil

}
