package go_utils

// High32 64位高32位值
func High32(value uint64) uint32 {
	return uint32(value >> 32 & 0xFFFFFFFF)
}

// Low32 64位低32位值
func Low32(value uint64) uint32 {
	return uint32(value & 0xFFFFFFFF)
}

// Make64 32+32组装64
func Make64(low, high uint32) uint64 {
	return uint64(low) | uint64(high)<<32
}

// High16 32位高16位值
func High16(value uint32) uint16 {
	return uint16(value >> 16 & 0xFFFF)
}

// Low16 32位低16位值
func Low16(value uint32) uint16 {
	return uint16(value & 0xFFFF)
}

// Make32 16+16组装32
func Make32(low, high uint16) uint32 {
	return uint32(low) | uint32(high)<<16
}
