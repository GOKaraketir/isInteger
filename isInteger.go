package isInteger

func IsInteger(val interface{}) bool {
	switch val.(type) {
	case uint, int, uint8, uint16, uint32, uint64, int8, int16, int32, int64:
		return true
	default:
		return false
	}
}
