package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_lsn_mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(288)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(v10) < base.Ui64(v12) {
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v12 - v10
		v20 = F_pg_snprintf(m, v7+int32(32), int32(256), int32(39653), v7)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v35 = int32(0)
			v40 = F_DirectFunctionCall3Coll(m, int32(408), v35, v7+int32(32), v35, int32(-1))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(288)
				return v40
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v10 - v12
		v32 = F_pg_snprintf(m, v7+int32(32), int32(256), int32(40068), v7+int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v35 = int32(0)
			v40 = F_DirectFunctionCall3Coll(m, int32(408), v35, v7+int32(32), v35, int32(-1))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(288)
				return v40
			}
		}
	}
}
func F_pg_lsn_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui64(v5) < base.Ui64(v7) {
		v9 = v5
	} else {
		v9 = v7
	}
	v10 = F_Int64GetDatum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
