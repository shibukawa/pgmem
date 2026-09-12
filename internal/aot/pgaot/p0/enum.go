package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enum_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v4^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_enum_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_SearchSysCache1(m, int32(23), v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
					F_errmsg(m, int32(55951), v6)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(475437), int32(167), int32(64203))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
			v37 = F_pstrdup(m, v32+v33+int32(12))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v10)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v37
				}
			}
		}
	}
}
func F_enum_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_SearchSysCache1(m, int32(23), v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v10 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v36 = v32 + v33 + int32(12)
	F_pq_begintypsend(m, v6+int32(16))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
	F_errmsg(m, int32(55951), v6)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(475437), int32(233), int32(407392))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	if v36&int32(3) == int32(0) {
		v66 = v36
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_pq_sendtext(m, v6+int32(16), v36, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L28
	}
L12:
	;
	v99 = v91 - v36
	goto L11
L13:
	;
	v70 = v66
	goto L22
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v50 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v99 = int32(0)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v55 = v36
	goto L18
L18:
	;
	v59 = v55 + int32(1)
	if v59&int32(3) == int32(0) {
		v66 = v59
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v91 = v59
	goto L12
L20:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v64 != 0 {
		v55 = v59
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v79 = int32(-2139062144)
	if (int32(16843008)-v76|v76)&v79 == v79 {
		v70 = v70 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v85 = v70
	goto L25
L24:
	;
	goto L23
L25:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v89 != 0 {
		v85 = v85 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v91 = v85
	goto L12
L27:
	;
	goto L26
L28:
	;
	F_ReleaseCatCache(m, v10)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v105 = v6 + int32(16)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v108 << (uint(int32(2)) % 32)
	goto L30
L30:
	;
	m.G0 = v6 + int32(32)
	return v107
}
