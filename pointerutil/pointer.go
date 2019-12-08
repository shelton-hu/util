package pointerutil

// Pointer->Type
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

func GetBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func GetStringMap(v *map[string]interface{}) map[string]interface{} {
	if v == nil {
		return map[string]interface{}{}
	}
	return *v
}

// Type->Interface
func GetStringPointer(v string) *string {
	return &v
}

func GetIntPointer(v int) *int {
	return &v
}

func GetInt32Pointer(v int32) *int32 {
	return &v
}

func GetInt64Pointer(v int64) *int64 {
	return &v
}

func GetUintPointer(v uint) *uint {
	return &v
}

func GetUint32Pointer(v uint32) *uint32 {
	return &v
}

func GetUint64Pointer(v uint64) *uint64 {
	return &v
}

func GetBoolPointer(v bool) *bool {
	return &v
}
