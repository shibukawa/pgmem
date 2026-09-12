package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LockBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 < int32(0) {
		m.G0 = v6 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		v14 = v11 + l0<<(uint(int32(6))%32)
		switch l1 {
		case 0:
			F_LWLockRelease(m, v14-int32(16))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		case 1:
			v18 = F_LWLockAcquire(m, v14-int32(16), int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		case 2:
			v23 = F_LWLockAcquire(m, v14-int32(16), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
				F_errmsg_internal(m, int32(508226), v6)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errfinish(m, int32(518165), int32(5617), int32(237597))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_allocNewBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v14 = int32(3)
	v15 = l1 & v14
	v17 = base.B2i32(v15 == v14)
	v20 = l1<<(uint(int32(1))%32)&int32(8) | v17<<(uint(int32(2))%32)
	v21 = F_spgGetCache(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v36 = F_SpGistNewBuffer(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	return v36
L5:
	;
	if v57&int32(3) != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v38 = int32(0)
	v39 = base.B2i32(v38 <= v36)
	if v39 == v38 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v36^int32(-1))<<(uint(int32(2))%32))))
	v57 = v49
	goto L5
L8:
	;
	goto L9
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v57 = v51 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L10:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
	v100 = v57 + v99
	v101 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)) = uint16(v101)
	*(*uint16)(unsafe.Add(mBase, uint32(v100))) = uint16(v20)
	if v15 == v14 {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+10)) = int32(1572864)
	v90 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+18)) = uint16(v90)
	v96 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)) = uint16(v96)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+14)) = uint16(v96)
	goto L10
L12:
	;
	v84 = F___memset(m, v57, int32(0), int32(8192))
	mBase = m.M
	goto L11
L13:
	;
	goto L12
L20:
	;
	goto L4
L21:
	;
	if v36 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v124 = base.I32_rem_u_s(v122, int32(3))
	if v15 == v124 {
		goto L20
	} else {
		goto L26
	}
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v36^int32(-1))<<(uint(int32(6))%32))+16))
	v122 = v113
	goto L22
L24:
	;
	goto L25
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115+v36<<(uint(int32(6))%32)+int32(-64))+16))
	v122 = v121
	goto L22
L26:
	;
	if v20 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v128 = v124 | int32(4)
	goto L29
L28:
	;
	v128 = v124
	goto L29
L29:
	;
	v131 = v21 - int32(-64) + v128<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v122
	if v39 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+14)))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+12)))
	v153 = v151 - v152
	v154 = int32(0)
	if v154 < v153 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v36^int32(-1))<<(uint(int32(2))%32))))
	v150 = v142
	goto L30
L32:
	;
	goto L33
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v150 = v144 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v157
	F_UnlockReleaseBuffer(m, v36)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L38
	}
L35:
	;
	v157 = v153
	goto L37
L36:
	;
	v157 = v154
	goto L37
L37:
	;
	goto L34
L38:
	;
	goto L3
}
