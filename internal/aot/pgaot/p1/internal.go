package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_internal_size(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v53 int32
	_ = v53
	if l1 <= int32(0) {
		v43 = int64(0)
	} else {
		v9 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+4)))
		v10 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0))))
		v13 = v9 - v10 + int64(1)
		if base.Ui32(l1) < base.Ui32(int32(3)) {
			v43 = v13
		} else {
			v19 = v13
			v20 = int32(2)
			for {
				v25 = l0 + v20<<(uint(int32(2))%32)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(4))))
				if v26 != v29 {
					v31 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+4)))
					v37 = v19 + v31 - base.I64_extend_i32_s(v26) + int64(1)
				} else {
					v37 = v19
				}
				v39 = v20 + int32(2)
				if v39 < l1 {
					v19 = v37
					v20 = v39
					continue
				} else {
					break
				}
				break
			}
			v43 = v37
		}
	}
	if base.Ui64(v43-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		v53 = int32(-1)
	} else {
		v53 = base.I32_wrap_i64(v43)
	}
	return v53
}
