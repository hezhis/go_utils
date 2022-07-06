package bitflag

// IsSetBit32 是否设置指定位
func IsSetBit32(value uint32, bit int) bool {
	if bit > 31 {
		return false
	}
	return (value & (1 << bit)) > 0
}

// SetBit32 设置位
func SetBit32(value uint32, bit int) uint32 {
	if bit > 31 {
		return value
	}
	return value | (1 << bit)
}

// ClearBit32 清空位
func ClearBit32(value uint32, bit int) uint32 {
	if bit > 31 {
		return value
	}
	return value & ^(1 << bit)
}
