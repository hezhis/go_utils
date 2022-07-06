package bitflag

// IsSetBit16 是否设置指定位
func IsSetBit16(value uint16, bit int) bool {
	if bit > 31 {
		return false
	}
	return (value & (1 << bit)) > 0
}

// SetBit16 设置位
func SetBit16(value uint16, bit int) uint16 {
	if bit > 31 {
		return value
	}
	return value | (1 << bit)
}

// ClearBit16 清空位
func ClearBit16(value uint16, bit int) uint16 {
	if bit > 31 {
		return value
	}
	return value & ^(1 << bit)
}
