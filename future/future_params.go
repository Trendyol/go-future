package future

type Params []any

func GetParam[T any](futureParams Params, index int) (ret T, err error) {
	data, ok := futureParams[index].(T)
	if !ok {
		return ret, ErrParamTypeConversionFailed
	}
	return data, nil
}

func (p Params) GetIntParam(index int) int {
	return p[index].(int)
}

func (p Params) GetInt32Param(index int) int32 {
	return p[index].(int32)
}

func (p Params) GetInt64Param(index int) int64 {
	return p[index].(int64)
}

func (p Params) GetUIntParam(index int) uint {
	return p[index].(uint)
}

func (p Params) GetUInt8Param(index int) uint8 {
	return p[index].(uint8)
}

func (p Params) GetUInt32Param(index int) uint32 {
	return p[index].(uint32)
}

func (p Params) GetUInt64Param(index int) uint64 {
	return p[index].(uint64)
}

func (p Params) GetFloat32Param(index int) float32 {
	return p[index].(float32)
}

func (p Params) GetFloat64Param(index int) float64 {
	return p[index].(float64)
}

func (p Params) GetStringParam(index int) string {
	return p[index].(string)
}

func (p Params) GetBoolParam(index int) bool {
	return p[index].(bool)
}

func (p Params) GetBytesParam(index int) []byte {
	return p[index].([]byte)
}
