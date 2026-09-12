package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpfamilyIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_SearchSysCache1(m, int32(42), l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v159
L2:
	;
	return int32(0)
L3:
	;
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v22)
	v159 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg_internal(m, int32(37633), v13)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(480234), int32(2283), int32(62006))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v42 = v38 + v39
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
	if v43 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L44
	}
L15:
	;
	v46 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	if v48 == v46 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v87 == int32(0) {
		v147 = v46
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v87 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v55 <= int32(0) {
		v80 = v46
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v87 = v80
	goto L18
L23:
	;
	v58 = int32(0)
	if v58 < v55 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v61 = v55
	goto L26
L25:
	;
	v61 = v58
	goto L26
L26:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v64 = int32(0)
	goto L27
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62+v64<<(uint(int32(2))%32))))
	v73 = base.B2i32(v72 == v43)
	if v72 == v43 {
		v80 = v73
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v80 = v73
	goto L22
L29:
	;
	v75 = v64 + int32(1)
	if v75 != v61 {
		v64 = v75
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	if v95 == int32(0) {
		v138 = v3
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v147 = base.B2i32(l0 == v138)
	goto L14
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v98 <= int32(0) {
		v138 = v3
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v107 = int32(0)
	v109 = v104
	v113 = v98
	goto L36
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v107<<(uint(int32(2))%32))))
	if v109 != v120 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v138 = int32(0)
	goto L33
L38:
	;
	v124 = F_GetSysCacheOid(m, int32(41), v91, v42+int32(8), v120, int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	v129 = v109
	v130 = v113
	goto L40
L40:
	;
	v132 = v107 + int32(1)
	if v132 < v130 {
		v107 = v132
		v109 = v129
		v113 = v130
		goto L36
	} else {
		goto L43
	}
L41:
	;
	if v124 != 0 {
		v138 = v124
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v128 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v129 = v128
	v130 = v126
	goto L40
L43:
	;
	goto L37
L44:
	;
	v159 = v147
	goto L1
}
func F_get_opfamily_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(42), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(37909), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(479899), int32(1404), int32(362791))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v33 = F_pstrdup(m, v28+v29+int32(8))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v9)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v33
				}
			}
		}
	}
}
func F_opfamily_can_sort_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	v3 = int32(0)
	v13 = F_SearchSysCacheList(m, int32(13), int32(1), int32(403), v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if int32(0) < v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = v3
	goto L6
L4:
	;
	v52 = v3
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v13)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(48)+v24<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+80))
	if v36 != l0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v52 = int32(0)
	goto L5
L8:
	;
	v47 = v24 + int32(1)
	if v47 != v17 {
		v24 = v47
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+84))
	if v38 != l1 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	F_ReleaseCatCacheList(m, v13)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	return base.B2i32(v40 != int32(0))
L12:
	;
	goto L7
L13:
	;
	return v52
}
