package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_has_startup_progress_timeout_expired(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v31 int64
	_ = v31
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_has_startup_progress_timeout_expired[0]))
	if v11 != 0 {
		v15 = m.G0
		v16 = int32(16)
		v17 = v15 - v16
		m.G0 = v17
		F_gettimeofday(m, v17)
		mBase = m.M
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+8)))
		m.G0 = v17 + v16
		v31 = *(*int64)(unsafe.Add(mBase, _c_F_has_startup_progress_timeout_expired[1]))
		v38 = v21 + v20*int64(1000000) - int64(946684800000000) - v31
		if v38 <= int64(0) {
			v50 = int32(0)
			v51 = int32(0)
		} else {
			v42 = int64(1000000)
			v43 = base.I64_div_u_s(v38, v42)
			v50 = base.I32_wrap_i64(v43)
			v51 = base.I32_wrap_i64(v38 - v43*v42)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(12)))) = v50
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(8)))) = v51
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v54
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
		*(*int32)(unsafe.Add(mBase, _c_F_has_startup_progress_timeout_expired[0])) = int32(0)
	} else {
	}
	m.G0 = v8 + int32(16)
	return base.B2i32(v11 != int32(0))
}
