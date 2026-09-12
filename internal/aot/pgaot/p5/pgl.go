package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_atexit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1338]))
	if v6 <= int32(31) {
		*(*int32)(unsafe.Add(mBase, _consts[1338])) = v6 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(2))%32))+uint32(_consts[1339]))) = l0
		v19 = int32(0)
	} else {
		v19 = int32(-1)
	}
	return v19
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
		v7 = *(*int32)(unsafe.Add(mBase, _consts[265]))
		v8 = F_freopen(m, l0, l1, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1340])) = v8
			return v8
		}
	case 1:
		v16 = *(*int32)(unsafe.Add(mBase, _consts[266]))
		v17 = F_freopen(m, l0, l1, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1341])) = v17
			return v17
		}
	case 2:
		v22 = *(*int32)(unsafe.Add(mBase, _consts[467]))
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	return v2
}
func F_pgl_getPGliteExitStatus(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	return v2
}
func F_pgl_set_rw_cbs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1342])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[1343])) = l0
	return
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
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v4 = m.G0
	v6 = v4 - int32(144)
	m.G0 = v6
	v8 = int32(4520272)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(1)
	v21 = F__emscripten_memset_bulkmem(m, v6+int32(20), base.I32_extend8_s(int32(0)), int32(120))
	mBase = m.M
	v22 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)) = uint16(v22)
	v31 = int32(14357)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+14)) = uint16(v31)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+140)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(16777343)
	v54 = F_pq_init(m, v6+int32(8))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[9])) = v9
		*(*int32)(unsafe.Add(mBase, _consts[381])) = v54
		*(*int32)(unsafe.Add(mBase, _consts[282])) = int32(2)
		v64 = int32(1)
		*(*int32)(unsafe.Add(mBase, _consts[19])) = v64
		*(*uint8)(unsafe.Add(mBase, _consts[811])) = uint8(v64)
		*(*uint8)(unsafe.Add(mBase, _consts[184])) = uint8(v64)
		v73 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[812])) = uint8(v73)
		v75 = F_load_hba(m)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return
		} else {
			if v75 == int32(0) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, _consts[813]))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v84
					F_errmsg(m, int32(197621), v6)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						F_errfinish(m, int32(495279), int32(260), int32(351351))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
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
