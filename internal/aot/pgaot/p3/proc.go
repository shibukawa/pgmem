package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecProcNodeFirst(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	F_check_stack_depth(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v8 == int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v12 = v11
		} else {
			v12 = int32(633)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
		v14 = m.T0[v12].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_ExecProcNodeInstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_InstrStartNode(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v9 = m.T0[v8].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v9 == int32(0) {
				F_InstrStopNode(m, v11, float64(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					return v9
				}
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
				if v20&int32(2) != 0 {
					v23 = float64(0)
				} else {
					v23 = float64(1)
				}
				F_InstrStopNode(m, v11, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					return v9
				}
			}
		}
	}
}
