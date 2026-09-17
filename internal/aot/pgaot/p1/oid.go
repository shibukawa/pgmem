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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v10 = v5 + int32(-56)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_OidFunctionCall1Coll[0]))
	F_fmgr_info_cxt_security(m, l0, v10, v12, int32(0))
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
		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v10
		v27 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+54)) = uint16(v27)
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v32 = m.T0[v31].(func(*base.Module, int32) int32)(m, v5+int32(-28))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+52)))
			if v34 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v41
					F_errmsg_internal(m, int32(_a_F_OidFunctionCall1Coll_0), v7)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_OidFunctionCall1Coll_1), int32(1143), int32(_a_F_OidFunctionCall1Coll_2))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
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
				return v32
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
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L26
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L23
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	switch v13 {
	case 0:
		v70 = int32(0)
		goto L6
	case 1:
		goto L7
	default:
		goto L2
	}
L4:
	;
	goto L5
L5:
	;
	goto L1
L6:
	;
	m.G0 = v8 + int32(16)
	return v70
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != int32(40) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v20 = F_palloc(m, l0<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if int32(0) < l0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v28 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v58 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v58 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L14:
	;
	v33 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v33 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v36 == int32(41) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v45 = F_strtox_2(m, v33, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v28<<(uint(int32(2))%32)))) = base.I32_wrap_i64(v45)
	v49 = v28 + int32(1)
	if v49 != l0 {
		v28 = v49
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v61 != int32(1) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v64 != int32(41) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v70 = v20
	goto L6
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v81
	F_errmsg_internal(m, int32(_a_F_readOidCols_0), v8)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_readOidCols_1), int32(693), int32(_a_F_readOidCols_2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F_errmsg_internal(m, int32(_a_F_readOidCols_3), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_readOidCols_1), int32(693), int32(_a_F_readOidCols_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
