package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_atexit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_atexit[0]))
	if v3 <= int32(31) {
		*(*int32)(unsafe.Add(mBase, _c_F_pgl_atexit[0])) = v3 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_pgl_atexit[1]))) = int32(1103)
	} else {
	}
	return
}
func F_pgl_freopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	switch l2 {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[0]))
		v8 = F_freopen(m, l0, l1, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[1])) = v8
			return v8
		}
	case 1:
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[2]))
		v17 = F_freopen(m, l0, l1, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[3])) = v17
			return v17
		}
	case 2:
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[4]))
		v23 = F_freopen(m, l0, l1, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = v23
			return v25
		}
	default:
		v25 = int32(0)
		return v25
	}
}
func F_pgl_getMyProcPort(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_getMyProcPort[0]))
	return v2
}
func F_pgl_startPGlite(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v4 = m.G0
	v6 = v4 - int32(144)
	m.G0 = v6
	v8 = int32(_a_F_pgl_startPGlite_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_startPGlite[0]))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_startPGlite[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_startPGlite[0])) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(1)
	base.MemoryFill(m, v6+int32(20), int32(0), int32(120))
	v21 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)) = uint16(v21)
	v30 = int32(_a_F_pgl_startPGlite_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+14)) = uint16(v30)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+140)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = base.I32_rotr(int32(2130706433), int32(24))&int32(16711935) | base.I32_rotr(int32(1), int32(8))
	v47 = F_pq_init(m, v6+int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pgl_startPGlite[0])) = v9
		*(*int32)(unsafe.Add(mBase, _c_F_pgl_startPGlite[2])) = v47
		*(*int32)(unsafe.Add(mBase, _c_F_pgl_startPGlite[3])) = int32(2)
		v57 = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_pgl_startPGlite[4])) = v57
		*(*uint8)(unsafe.Add(mBase, _c_F_pgl_startPGlite[5])) = uint8(v57)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgl_startPGlite[6])) = uint8(v57)
		v66 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgl_startPGlite[7])) = uint8(v66)
		v68 = F_load_hba(m)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return
		} else {
			if v68 == int32(0) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_startPGlite[8]))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v77
					F_errmsg(m, int32(_a_F_pgl_startPGlite_2), v6)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pgl_startPGlite_3), int32(260), int32(_a_F_pgl_startPGlite_4))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v6 + int32(144)
				return
			}
		}
	}
}
