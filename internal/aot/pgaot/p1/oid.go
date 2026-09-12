package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OidFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v12 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	F_fmgr_info_cxt_security(m, l0, v5+int32(-56), v12, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+60)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+52)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v5 + int32(-56)
		v29 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+54)) = uint16(v29)
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v34 = m.T0[v33].(func(*base.Module, int32) int32)(m, v5+int32(-28))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+52)))
			if v36 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v43
					F_errmsg_internal(m, int32(523043), v7)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(486266), int32(1143), int32(298999))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v7 - int32(-64)
				return v34
			}
		}
	}
}
func F_oid_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v4) < base.Ui32(v3) {
		v8 = v3 - v4
	} else {
		v8 = v4 - v3
	}
	return v8
}
func F_readOidCols(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L10
	} else {
		goto L33
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L30
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L27
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	switch v13 {
	case 0:
		v70 = int32(0)
		goto L7
	case 1:
		goto L8
	default:
		goto L3
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L24
	}
L7:
	;
	m.G0 = v8 + int32(16)
	return v70
L8:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != int32(40) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v20 = F_palloc(m, l0<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if int32(0) < l0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v28 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v58 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v58 == int32(0) {
		goto L2
	} else {
		goto L21
	}
L15:
	;
	v33 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v33 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v36 == int32(41) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v45 = F_strtox_2(m, v33, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v28<<(uint(int32(2))%32)))) = base.I32_wrap_i64(v45)
	v49 = v28 + int32(1)
	if v49 != l0 {
		v28 = v49
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v61 != int32(1) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v64 != int32(41) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v70 = v20
	goto L7
L24:
	;
	F_errmsg_internal(m, int32(24646), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(486040), int32(693), int32(149244))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v94
	F_errmsg_internal(m, int32(664451), v8)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(486040), int32(693), int32(149244))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	F_errmsg_internal(m, int32(24646), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(486040), int32(693), int32(149244))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errmsg_internal(m, int32(24646), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(486040), int32(693), int32(149244))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
