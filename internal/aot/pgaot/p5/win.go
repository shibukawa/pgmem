package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WinGetFuncArgInPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+400))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	v20 = v18 + base.I64_extend_i32_s(l1)
	v21 = F_window_gettupleslot(m, l0, v20, v16)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		if v21 == int32(0) {
			if l4 != 0 {
				v27 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v27)
			} else {
			}
			v29 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v29)
			v44 = int32(0)
			m.G0 = v13 + int32(16)
			return v44
		} else {
			if l4 != 0 {
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v32)
			} else {
			}
			if l2 != 0 {
				F_WinSetMarkPosition(m, l0, v20)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v16
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
					v41 = m.T0[v40].(func(*base.Module, int32, int32, int32) int32)(m, v39, v17, l3)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v44 = v41
						m.G0 = v13 + int32(16)
						return v44
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v16
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
				v41 = m.T0[v40].(func(*base.Module, int32, int32, int32) int32)(m, v39, v17, l3)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v44 = v41
					m.G0 = v13 + int32(16)
					return v44
				}
			}
		}
	}
}
func F_WinGetPartitionRowCount(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_spool_tuples(m, v2, int64(-1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)+168))
		return v9
	}
}
