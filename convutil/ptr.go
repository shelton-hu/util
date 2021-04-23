package convutil

func GetBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func GetString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func GetInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func GetInt8(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

func GetInt16(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

func GetInt32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func GetInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func GetUint(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

func GetUint8(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

func GetUint16(v *uint16) uint16 {
	if v == nil {
		return 0
	}
	return *v
}

func GetUint32(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

func GetUint64(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

func GetFloat32(v *float32) float32 {
	if v == nil {
		return 0.0
	}
	return *v
}

func GetFloat64(v *float64) float64 {
	if v == nil {
		return 0.0
	}
	return *v
}

func GetBoolPtr(v bool) *bool {
	return &v
}

func GetStringPtr(v string) *string {
	return &v
}

func GetIntPtr(v int) *int {
	return &v
}

func GetInt8Ptr(v int8) *int8 {
	return &v
}

func GetInt16Ptr(v int16) *int16 {
	return &v
}

func GetInt32Ptr(v int32) *int32 {
	return &v
}

func GetInt64Ptr(v int64) *int64 {
	return &v
}

func GetUintPtr(v uint) *uint {
	return &v
}

func GetUint8Ptr(v uint8) *uint8 {
	return &v
}

func GetUint16Ptr(v uint16) *uint16 {
	return &v
}

func GetUint32Ptr(v uint32) *uint32 {
	return &v
}

func GetUint64Ptr(v uint64) *uint64 {
	return &v
}

func GetFloat32Ptr(v float32) *float32 {
	return &v
}

func GetFloat64Ptr(v float64) *float64 {
	return &v
}
