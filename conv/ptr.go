package conv

// Val 解引用任意类型指针，nil 时返回类型零值。
// 推荐用法，可替代 String/Int/Bool 等一系列特化函数。
func Val[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// Ptr 返回任意类型值的指针。
// 推荐用法，可替代 StringPtr/IntPtr/BoolPtr 等一系列特化函数。
func Ptr[T any](v T) *T {
	return &v
}

func String(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func StringPtr(v string) *string {
	return &v
}

func Bool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func BoolPtr(v bool) *bool {
	return &v
}

func Int(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}
func IntPtr(v int) *int {
	return &v
}

func Int16(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

func Int16Ptr(v int16) *int16 {
	return &v
}

func Int32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func Int32Ptr(v int32) *int32 {
	return &v
}

func Int64(e *int64) int64 {
	if e == nil {
		return 0
	}
	return *e
}

func Int64Ptr(e int64) *int64 {
	return &e
}

func Int64PtrToInt(v *int64) int {
	if v == nil {
		return 0
	}
	return int(*v)
}

func Int64PtrToIntPtr(v *int64) *int {
	if v == nil {
		return nil
	}
	return IntPtr(int(*v))
}

func IntPtrToInt64Ptr(v *int) *int64 {
	if v == nil {
		return nil
	}
	return Int64Ptr(int64(*v))
}
