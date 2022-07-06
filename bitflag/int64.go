package bitflag

// IsSetBit64 是否设置指定位
func IsSetBit64(value uint64, bit int) bool {
	if bit > 63 {
		return false
	}
	return (value & (1 << bit)) > 0
}

// SetBit64 设置位
func SetBit64(value uint64, bit int) uint64 {
	if bit > 63 {
		return value
	}
	return value | (1 << bit)
}

// ClearBit64 清空位
func ClearBit64(value uint64, bit int) uint64 {
	if bit > 63 {
		return value
	}
	return value & ^(1 << bit)
}
