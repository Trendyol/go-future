package future

type Params []any

func GetParam[T any](futureParams Params, index int) (ret T, err error) {
	data, ok := futureParams[index].(T)
	if !ok {
		return ret, ErrParamTypeConversionFailed
	}
	return data, nil
}
