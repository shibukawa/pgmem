package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_GetActiveSnapshot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_GetBulkInsertState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_palloc(m, int32(20))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_GetAccessStrategy(m, int32(2))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v4)+12)) = int64(4294967295)
			*(*int64)(unsafe.Add(mBase, uint32(v4)+4)) = int64(-4294967296)
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v9
			return v4
		}
	}
}
func F_GetFdwRoutineByRelId(m *base.Module, l0 int32) int32 {
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(33), l0)
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
				F_errmsg_internal(m, int32(51600), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491248), int32(364), int32(460545))
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
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+4))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = F_GetFdwRoutineByServerId(m, v31)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v34
				}
			}
		}
	}
}
func F_GetNSItemByRangeTablePosn(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v4 = int32(0)
	if l2 <= v4 {
		v53 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	if v59 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v10 = l2 & int32(7)
	if v10 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if base.Ui32(l2) < base.Ui32(int32(8)) {
		v53 = v25
		goto L1
	} else {
		goto L10
	}
L4:
	;
	v25 = l0
	v28 = l2
	goto L3
L5:
	;
	goto L6
L6:
	;
	v13 = l0
	v16 = l2
	v17 = v4
	goto L7
L7:
	;
	v19 = int32(1)
	v20 = v16 - v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v23 = v17 + v19
	if v23 != v10 {
		v13 = v21
		v16 = v20
		v17 = v23
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v25 = v21
	v28 = v20
	goto L3
L9:
	;
	goto L8
L10:
	;
	v33 = v25
	v36 = v28
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if base.Ui32(v36-int32(9)) < base.Ui32(int32(-2)) {
		v33 = v48
		v36 = v36 - int32(8)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v53 = v48
	goto L1
L13:
	;
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v62 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v65 = int32(0)
	if v65 < v62 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v69 = v62
	goto L19
L18:
	;
	v69 = v65
	goto L19
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v74 = v65
	goto L20
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70+v74<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	if l1 != v81 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	return v80
L22:
	;
	v84 = v74 + int32(1)
	if v69 != v84 {
		v74 = v84
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L14
L26:
	;
	return int32(0)
L27:
	;
	F_errmsg_internal(m, int32(654131), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(490894), int32(536), int32(242013))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetOldestRestartPoint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v8 = F_LWLockAcquire(m, v4+int32(1152), int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[178]))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
		v17 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		F_LWLockRelease(m, v17+int32(1152))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
func F_GetPublicationsStr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if l2 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v28 = int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v29 <= v28 {
		goto L1
	} else {
		goto L14
	}
L5:
	;
	F_appendStringInfoChar(m, l1, int32(34))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v24 = F_quote_literal_cstr(m, v13)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L12
	}
L8:
	;
	return
L9:
	;
	F_appendStringInfoString(m, l1, v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_appendStringInfoChar(m, l1, int32(34))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	F_appendStringInfoString(m, l1, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L4
L14:
	;
	v35 = v28
	goto L15
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v35<<(uint(int32(2))%32))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	F_appendStringInfoString(m, l1, int32(724350))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L17
	}
L16:
	;
	goto L1
L17:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v59 = v35 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v59 < v60 {
		v35 = v59
		goto L15
	} else {
		goto L27
	}
L19:
	;
	v46 = F_quote_literal_cstr(m, v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_appendStringInfoChar(m, l1, int32(34))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	F_appendStringInfoString(m, l1, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	F_appendStringInfoString(m, l1, v42)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	F_appendStringInfoChar(m, l1, int32(34))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	goto L16
}
func F_GetTopMostAncestorInPublication(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	v3 = int32(0)
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v246
L2:
	;
	v22 = v3
	v24 = v3
	goto L7
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v12 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v246 = v3
	goto L1
L6:
	;
	goto L5
L7:
	;
	v26 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v22<<(uint(int32(2))%32))))
	v36 = F_SearchSysCacheList(m, int32(53), int32(1), v33, v26, v26)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v246 = v230
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	if int32(0) < v40 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v48 = int32(0)
	v52 = v26
	goto L14
L12:
	;
	v77 = v26
	goto L13
L13:
	;
	F_ReleaseCatCacheList(m, v36)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L9
	} else {
		goto L18
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(48)+v48<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+22)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61+v62)+4))
	v65 = F_lappend_oid(m, v52, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v77 = v65
	goto L13
L16:
	;
	v68 = v48 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	if v68 < v69 {
		v48 = v68
		v52 = v65
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v85 = v22 + int32(1)
	v86 = int32(0)
	if v77 == v86 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	F_list_free(m, v77)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L63
	}
L20:
	;
	v222 = int32(0)
	v230 = v33
	goto L19
L21:
	;
	if v124 != 0 {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v124 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v92 <= int32(0) {
		v117 = v86
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v124 = v117
	goto L21
L26:
	;
	v95 = int32(0)
	if v95 < v92 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v98 = v92
	goto L29
L28:
	;
	v98 = v95
	goto L29
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v101 = int32(0)
	goto L30
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v99+v101<<(uint(int32(2))%32))))
	v110 = base.B2i32(v109 == l0)
	if v109 == l0 {
		v117 = v110
		goto L25
	} else {
		goto L32
	}
L31:
	;
	v117 = v110
	goto L25
L32:
	;
	v112 = v101 + int32(1)
	if v112 != v98 {
		v101 = v112
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L20
L35:
	;
	goto L36
L36:
	;
	v126 = int32(0)
	v129 = F_get_rel_namespace(m, v33)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v131 = int32(0)
	v133 = F_SearchSysCacheList(m, int32(50), int32(1), v129, v131, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	if int32(0) < v135 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v143 = int32(0)
	v144 = v126
	goto L42
L40:
	;
	v169 = v126
	goto L41
L41:
	;
	F_ReleaseCatCacheList(m, v133)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L46
	}
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v133+int32(48)+v143<<(uint(int32(2))%32))))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+56))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+22)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156+v157)+4))
	v160 = F_lappend_oid(m, v144, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L44
	}
L43:
	;
	v169 = v160
	goto L41
L44:
	;
	v163 = v143 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	if v163 < v164 {
		v143 = v163
		v144 = v160
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v179 = int32(0)
	if v169 == v179 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v217 != 0 {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v217 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v185 <= int32(0) {
		v210 = v179
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v217 = v210
	goto L47
L52:
	;
	v188 = int32(0)
	if v188 < v185 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v191 = v185
	goto L55
L54:
	;
	v191 = v188
	goto L55
L55:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v194 = int32(0)
	goto L56
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v192+v194<<(uint(int32(2))%32))))
	v203 = base.B2i32(v202 == l0)
	if v202 == l0 {
		v210 = v203
		goto L51
	} else {
		goto L58
	}
L57:
	;
	v210 = v203
	goto L51
L58:
	;
	v205 = v194 + int32(1)
	if v205 != v191 {
		v194 = v205
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v218 = v33
	goto L62
L61:
	;
	v218 = v24
	goto L62
L62:
	;
	v222 = v169
	v230 = v218
	goto L19
L63:
	;
	F_list_free(m, v222)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v85 < v235 {
		v22 = v85
		v24 = v230
		goto L7
	} else {
		goto L65
	}
L65:
	;
	goto L8
}
func F_GetTopTransactionId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	v3 = int64(*(*uint32)(unsafe.Add(mBase, _consts[35])))
	if v3 == int64(0) {
		F_AssignTransactionId(m, int32(4371288))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, _consts[35]))
			v13 = v12
			return base.I32_wrap_i64(v13)
		}
	} else {
		v13 = v3
		return base.I32_wrap_i64(v13)
	}
}
func F_GetVirtualXIDsDelayingChkpt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v17 = F_palloc(m, v14<<(uint(int32(3))%32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		v26 = F_LWLockAcquire(m, v22+int32(512), int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if int32(0) < v28 {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[603]))
				v37 = v3
				v38 = v28
				v41 = v3
				for {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)+v41<<(uint(int32(2))%32))))
					v52 = v34 + v49*int32(640)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+120))
					if v53&l1 == int32(0) {
						v69 = v37
						v70 = v38
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
						if v57 == int32(0) {
							v69 = v37
							v70 = v38
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
							v63 = v17 + v37<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v63))) = v60
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v69 = v37 + int32(1)
							v70 = v68
						}
					}
					v74 = v41 + int32(1)
					if v74 < v70 {
						v37 = v69
						v38 = v70
						v41 = v74
						continue
					} else {
						break
					}
					break
				}
				v78 = v69
			} else {
				v78 = v3
			}
			v88 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			F_LWLockRelease(m, v88+int32(512))
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v78
				return v17
			}
		}
	}
}
func F_GlobalVisHorizonKindForRel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	v2 = int32(0)
	if l0 == v2 {
		v50 = v2
		return v50
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+117)))
		if v7 != 0 {
			v50 = v2
			return v50
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
			if v10 == int32(1) {
				v15 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+316))
				v18 = base.B2i32(v16 != int32(2))
				*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v18)
				v20 = v18
			} else {
				v20 = int32(0)
			}
			if v20 != 0 {
				v50 = v2
				return v50
			} else {
				v21 = int32(1)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if base.Ui32(v22) < base.Ui32(int32(12000)) {
					v50 = v21
					return v50
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _consts[8]))
					if v26 < int32(2) {
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v45 != 0 {
							v50 = int32(3)
							return v50
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							if v46 != 0 {
								v50 = int32(3)
								return v50
							} else {
								return int32(2)
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+118)))
						if v30 != int32(112) {
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v45 != 0 {
								v50 = int32(3)
								return v50
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v46 != 0 {
									v50 = int32(3)
									return v50
								} else {
									return int32(2)
								}
							}
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							if base.Ui32(v33) < base.Ui32(int32(12000)) {
								v50 = v21
								return v50
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
								if v36 == int32(0) {
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v45 != 0 {
										v50 = int32(3)
										return v50
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v46 != 0 {
											v50 = int32(3)
											return v50
										} else {
											return int32(2)
										}
									}
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+119)))
									switch v40 - int32(109) {
									case 0, 5:
										v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+104)))
										if v43 != 0 {
											v50 = v21
											return v50
										} else {
											v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v45 != 0 {
												v50 = int32(3)
												return v50
											} else {
												v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v46 != 0 {
													v50 = int32(3)
													return v50
												} else {
													return int32(2)
												}
											}
										}
									default:
										v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v45 != 0 {
											v50 = int32(3)
											return v50
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v46 != 0 {
												v50 = int32(3)
												return v50
											} else {
												return int32(2)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_g_intbig_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == v2 {
		v29 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v29&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v16 == int32(0) {
		v29 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v19 != int32(7) {
		v29 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 != int32(17) {
		v29 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+24)))
	v29 = v25 ^ int32(1)
	goto L2
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = F_get_fn_opclass_options(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v38 = int32(252)
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
	if v40 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	return int32(0)
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = v37
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L52
	}
L13:
	;
	v196 = F_palloc(m, int32(16))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L10
	} else {
		goto L51
	}
L14:
	;
	v43 = F_pg_detoast_datum(m, v39)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	if v146&int32(4) != 0 {
		goto L39
	} else {
		goto L40
	}
L17:
	;
	v46 = v38 + int32(8)
	v47 = F_palloc(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v46 << (uint(int32(2)) % 32)
	v58 = F__emscripten_memset_bulkmem(m, v47+int32(8), base.I32_extend8_s(v49), v38)
	mBase = m.M
	goto L19
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v59 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v60 = F_array_contains_nulls(m, v43)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v64 = v43 + int32(16)
	v65 = F_ArrayGetNItems(m, v62, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	if v60 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	if v65 == int32(0) {
		v193 = v47
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v71 = F_ArrayGetNItems(m, v70, v64)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	if v71 == int32(0) {
		v193 = v47
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v75 = int32(3)
	v76 = v38 << (uint(v75) % 32)
	if v69 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v83 = v69
	goto L31
L30:
	;
	v83 = (v70<<(uint(v75)%32) + int32(23)) & int32(-8)
	goto L31
L31:
	;
	v84 = v43 + v83
	if v71&int32(1) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v88 = base.I32_rem_u_s(v87, v76)
	v91 = v58 + int32(base.Ui32(v88)>>(uint(int32(3))%32))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v93 = int32(1)
	v97 = v92 | v93<<(uint(v88&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v97)
	v103 = v84 + int32(4)
	v106 = v71 - v93
	goto L34
L33:
	;
	v103 = v84
	v106 = v71
	goto L34
L34:
	;
	if v71 == int32(1) {
		v193 = v47
		goto L13
	} else {
		goto L35
	}
L35:
	;
	v109 = v103
	v115 = v106
	goto L36
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v119 = base.I32_rem_u_s(v118, v76)
	v120 = int32(3)
	v122 = v58 + int32(base.Ui32(v119)>>(uint(v120)%32))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v124 = int32(1)
	v125 = int32(7)
	v128 = v123 | v124<<(uint(v119&v125)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v131 = base.I32_rem_u_s(v130, v76)
	v134 = v58 + int32(base.Ui32(v131)>>(uint(v120)%32))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v140 = v135 | v124<<(uint(v131&v125)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v140)
	v145 = v115 - int32(2)
	if v145 != 0 {
		v109 = v109 + int32(8)
		v115 = v145
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v193 = v47
	goto L13
L38:
	;
	goto L37
L39:
	;
	return v10
L40:
	;
	goto L41
L41:
	;
	v150 = int32(0)
	if v38 <= v150 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v182 = F_palloc(m, int32(8))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L10
	} else {
		goto L50
	}
L43:
	;
	v155 = v150
	goto L44
L44:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+(v39+int32(8))))))
	if v165 == int32(255) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	return v10
L46:
	;
	v169 = v155 + int32(1)
	if v38 != v169 {
		v155 = v169
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	goto L42
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v182))) = int64(17179869216)
	v193 = v182
	goto L13
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v193
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+4)) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+8)) = v201
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)))
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+14)) = uint8(v204)
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+12)) = uint16(v203)
	return v196
L52:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(151167), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(487301), int32(159), int32(125898))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_g_intbig_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(0)
	if v20 == v21 {
		v37 = v21
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v37&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v24 == int32(0) {
		v37 = v21
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v27 != int32(7) {
		v37 = v21
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v30 != int32(17) {
		v37 = v21
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
	v37 = v33 ^ int32(1)
	goto L4
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = F_get_fn_opclass_options(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v44 = int32(252)
	goto L11
L11:
	;
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v45)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
	if v49&int32(4) != 0 {
		v425 = v45
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v44 = v43
	goto L11
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L97
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L93
	}
L15:
	;
	return v425 & int32(1)
L16:
	;
	if v18&int32(65535) == int32(20) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_pfree(m, v13)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L92
	}
L18:
	;
	v59 = F_signconsistent(m, v13, v48+int32(8), v44, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v63 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v13 != v61 {
		v413 = v59
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v425 = v59
	goto L15
L23:
	;
	v64 = F_array_contains_nulls(m, v13)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v66 = int32(0)
	switch v18&int32(65535) - int32(3) {
	case 0:
		goto L32
	default:
		v401 = v66
		goto L28
	case 3:
		goto L31
	case 4, 10:
		goto L30
	case 5, 11:
		goto L29
	}
L26:
	;
	if v64 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v13 == v410 {
		v425 = v401
		goto L15
	} else {
		goto L91
	}
L29:
	;
	v267 = int32(1)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268)+16)))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v269)+12)))
	if v271&v267 == int32(0) {
		v401 = v267
		goto L28
	} else {
		goto L71
	}
L30:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v265 = F__intbig_contains(m, v264, v13, v44)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L70
	}
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+16)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v123)+12)))
	if v125&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L32:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v75 = F_ArrayGetNItems(m, v72, v13+int32(16))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v77 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v75 == int32(0) {
		v401 = v66
		goto L28
	} else {
		goto L40
	}
L35:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v89 = (v80<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v87 = F_array_contains_nulls(m, v13)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v87 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v89 = v77
	goto L34
L40:
	;
	v100 = v75
	v103 = v13 + v89
	goto L41
L41:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v108 = base.I32_rem_u_s(v107, v44<<(uint(int32(3))%32))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+int32(8)+int32(base.Ui32(v108)>>(uint(int32(3))%32))))))
	v115 = int32(base.Ui32(v112) >> (uint(v108&int32(7)) % 32))
	if v115&int32(1) != 0 {
		v401 = v115
		goto L28
	} else {
		goto L43
	}
L42:
	;
	v401 = v115
	goto L28
L43:
	;
	v121 = v100 - int32(1)
	if v121 != 0 {
		v100 = v121
		v103 = v103 + int32(4)
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v131 = F_ArrayGetNItems(m, v128, v13+int32(16))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v262 = F__intbig_contains(m, v261, v13, v44)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L69
	}
L48:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v133 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v143 = (v136<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L51
L50:
	;
	v143 = v133
	goto L51
L51:
	;
	v144 = F_palloc0(m, v44)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v131 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v44 <= int32(0) {
		v250 = int32(1)
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v149 = v44 << (uint(int32(3)) % 32)
	v150 = v143 + v13
	if v131&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v154 = base.I32_rem_u_s(v153, v149)
	v157 = v144 + int32(base.Ui32(v154)>>(uint(int32(3))%32))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v159 = int32(1)
	v163 = v158 | v159<<(uint(v154&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v163)
	v169 = v150 + int32(4)
	v172 = v131 - v159
	goto L57
L56:
	;
	v169 = v150
	v172 = v131
	goto L57
L57:
	;
	if v131 == int32(1) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v176 = v169
	v182 = v172
	goto L59
L59:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v186 = base.I32_rem_u_s(v185, v149)
	v187 = int32(3)
	v189 = v144 + int32(base.Ui32(v186)>>(uint(v187)%32))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v191 = int32(1)
	v192 = int32(7)
	v195 = v190 | v191<<(uint(v186&v192)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v198 = base.I32_rem_u_s(v197, v149)
	v201 = v144 + int32(base.Ui32(v198)>>(uint(v187)%32))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v207 = v202 | v191<<(uint(v198&v192)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v207)
	v212 = v182 - int32(2)
	if v212 != 0 {
		v176 = v176 + int32(8)
		v182 = v212
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L53
L61:
	;
	goto L60
L62:
	;
	F_pfree(m, v144)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L68
	}
L63:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v233 = int32(0)
	goto L64
L64:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233+(v226+int32(8))))))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233+v144))))
	v244 = base.B2i32(v241 == v243)
	if v243 != v241 {
		v250 = v244
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v250 = v244
	goto L62
L66:
	;
	v247 = v233 + int32(1)
	if v247 != v44 {
		v233 = v247
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v401 = v250
	goto L28
L69:
	;
	v401 = v262
	goto L28
L70:
	;
	v401 = v265
	goto L28
L71:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v279 = F_ArrayGetNItems(m, v276, v13+int32(16))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v281 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v291 = (v284<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L75
L74:
	;
	v291 = v281
	goto L75
L75:
	;
	v292 = F_palloc0(m, v44)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v279 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if v44 <= int32(0) {
		v401 = int32(1)
		goto L28
	} else {
		goto L86
	}
L78:
	;
	v297 = v44 << (uint(int32(3)) % 32)
	v298 = v291 + v13
	if v279&int32(1) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	v302 = base.I32_rem_u_s(v301, v297)
	v305 = v292 + int32(base.Ui32(v302)>>(uint(int32(3))%32))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	v307 = int32(1)
	v311 = v306 | v307<<(uint(v302&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v305))) = uint8(v311)
	v317 = v298 + int32(4)
	v320 = v279 - v307
	goto L81
L80:
	;
	v317 = v298
	v320 = v279
	goto L81
L81:
	;
	if v279 == int32(1) {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v324 = v317
	v330 = v320
	goto L83
L83:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v334 = base.I32_rem_u_s(v333, v297)
	v335 = int32(3)
	v337 = v292 + int32(base.Ui32(v334)>>(uint(v335)%32))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v339 = int32(1)
	v340 = int32(7)
	v343 = v338 | v339<<(uint(v334&v340)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v337))) = uint8(v343)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v346 = base.I32_rem_u_s(v345, v297)
	v349 = v292 + int32(base.Ui32(v346)>>(uint(v335)%32))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	v355 = v350 | v339<<(uint(v346&v340)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v349))) = uint8(v355)
	v360 = v330 - int32(2)
	if v360 != 0 {
		v324 = v324 + int32(8)
		v330 = v360
		goto L83
	} else {
		goto L85
	}
L84:
	;
	goto L77
L85:
	;
	goto L84
L86:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v381 = int32(0)
	goto L87
L87:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381+(v374+int32(8))))))
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381+v292))))
	v394 = v389 & (v391 ^ int32(255))
	v396 = base.B2i32(v394 == int32(0))
	if v394 != 0 {
		v401 = v396
		goto L28
	} else {
		goto L89
	}
L88:
	;
	v401 = v396
	goto L28
L89:
	;
	v398 = v381 + int32(1)
	if v398 != v44 {
		v381 = v398
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v413 = v401
	goto L17
L92:
	;
	v425 = v413
	goto L15
L93:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(151167), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(487301), int32(492), int32(91835))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(151167), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(487301), int32(82), int32(236383))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_g_intbig_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(278953), int32(157724), int32(252), int32(1), int32(2024))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_gcd_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v304 int64
	_ = v304
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v396 int64
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v24 < v21)&base.B2i32(v4 < v20) == v4 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v198 = base.B2i32(v196 < int32(0))
	if v196 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L2:
	;
	v196 = v186
	goto L1
L3:
	;
	if v24 <= v56 {
		v91 = v24
		v93 = v4
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v56 = v21
	v60 = v4
	goto L3
L5:
	;
	goto L6
L6:
	;
	v37 = v21
	v41 = v4
	goto L7
L7:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v41<<(uint(int32(1))%32)))))
	if v47 != 0 {
		v186 = int32(1)
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v56 = v51
	v60 = v49
	goto L3
L9:
	;
	v48 = int32(1)
	v49 = v41 + v48
	v51 = v37 - v48
	if v51 <= v24 {
		v56 = v51
		v60 = v49
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v49 < v20 {
		v37 = v51
		v41 = v49
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	if v56 != v91 {
		v132 = v60
		v133 = v93
		goto L20
	} else {
		goto L21
	}
L13:
	;
	if v23 <= int32(0) {
		v91 = v24
		v93 = v4
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v72 = v24
	v74 = v4
	goto L15
L15:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v74<<(uint(int32(1))%32)))))
	if v79 != 0 {
		v186 = int32(-1)
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v91 = v83
	v93 = v81
	goto L12
L17:
	;
	v80 = int32(1)
	v81 = v74 + v80
	v83 = v72 - v80
	if v83 <= v56 {
		v91 = v83
		v93 = v81
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if v81 < v23 {
		v72 = v83
		v74 = v81
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	if v20 < v132 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v102 = v60
	v103 = v93
	goto L22
L22:
	;
	if v20 <= v102 {
		v132 = v102
		v133 = v103
		goto L20
	} else {
		goto L24
	}
L23:
	;
	if base.I32_extend16_s(v118) < base.I32_extend16_s(v116) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	if v23 <= v103 {
		v132 = v102
		v133 = v103
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v107 = int32(1)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v102<<(uint(v107)%32)))))
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103<<(uint(v107)%32)+v22))))
	if v116 == v118 {
		v102 = v102 + v107
		v103 = v103 + v107
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v125 = int32(1)
	goto L29
L28:
	;
	v125 = int32(-1)
	goto L29
L29:
	;
	v196 = v125
	goto L1
L30:
	;
	v136 = v132
	goto L32
L31:
	;
	v136 = v20
	goto L32
L32:
	;
	v143 = v132
	goto L33
L33:
	;
	if v136 == v143 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v186 = v169
	goto L2
L35:
	;
	if v23 < v133 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v169 = int32(1)
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v143<<(uint(v169)%32)))))
	if v175 == int32(0) {
		v143 = v143 + v169
		goto L33
	} else {
		goto L47
	}
L38:
	;
	v148 = v133
	goto L40
L39:
	;
	v148 = v23
	goto L40
L40:
	;
	v156 = v133
	goto L41
L41:
	;
	if v148 == v156 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v186 = int32(-1)
	goto L2
L43:
	;
	v196 = int32(0)
	goto L1
L44:
	;
	goto L45
L45:
	;
	v160 = int32(1)
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v156<<(uint(v160)%32)))))
	if v165 == int32(0) {
		v156 = v156 + v160
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	goto L34
L48:
	;
	v199 = l1
	goto L50
L49:
	;
	v199 = l0
	goto L50
L50:
	;
	if v196 < int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v200 = v23
	goto L53
L52:
	;
	v200 = v20
	goto L53
L53:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v202 < v201 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v204 = v201
	goto L56
L55:
	;
	v204 = v202
	goto L56
L56:
	;
	if v196 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	m.G0 = v17 + int32(80)
	return
L58:
	;
	v206 = v20
	goto L60
L59:
	;
	v206 = v23
	goto L60
L60:
	;
	if v206 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v207 = v196
	goto L63
L62:
	;
	v207 = int32(0)
	goto L63
L63:
	;
	if v207 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v214 = F_palloc(m, v200<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if v196 < int32(0) {
		goto L80
	} else {
		goto L81
	}
L67:
	;
	return
L68:
	;
	v216 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v214))) = uint16(v216)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v216 < v218 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v225 = v218 << (uint(int32(1)) % 32)
	if v225 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	goto L71
L71:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v228 != 0 {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	goto L71
L73:
	;
	v226 = F__emscripten_memcpy_bulkmem(m, v214+int32(2), v223, v225)
	mBase = m.M
	goto L75
L74:
	;
	goto L75
L75:
	;
	goto L72
L76:
	;
	F_pfree(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L67
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v199)+8))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v199)))
	v234 = l2 + int32(16)
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v234))) = v235
	v238 = l2 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v238))) = v231
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v232
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v214 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v214
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = int32(0)
	goto L57
L79:
	;
	goto L78
L80:
	;
	v248 = l0
	goto L82
L81:
	;
	v248 = l1
	goto L82
L82:
	;
	v249 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v249
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v249
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v249
	v259 = F_palloc(m, v200<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L67
	} else {
		goto L83
	}
L83:
	;
	v261 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v259))) = uint16(v261)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v261 < v263 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v270 = v263 << (uint(int32(1)) % 32)
	if v270 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	goto L86
L86:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v199)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v274
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v199)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v259
	v279 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v259 + v279
	v286 = F_palloc(m, v273<<(uint(int32(1))%32)+v279)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L67
	} else {
		goto L91
	}
L87:
	;
	goto L86
L88:
	;
	v271 = F__emscripten_memcpy_bulkmem(m, v259+int32(2), v268, v270)
	mBase = m.M
	goto L90
L89:
	;
	goto L90
L90:
	;
	goto L87
L91:
	;
	v288 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v286))) = uint16(v288)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if v288 < v290 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v248)+20))
	v297 = v290 << (uint(int32(1)) % 32)
	if v297 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	goto L94
L94:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v300 != 0 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	goto L94
L96:
	;
	v298 = F__emscripten_memcpy_bulkmem(m, v286+int32(2), v295, v297)
	mBase = m.M
	goto L98
L97:
	;
	goto L98
L98:
	;
	goto L95
L99:
	;
	F_pfree(m, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L67
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v248)+8))
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v248)))
	v306 = l2 + int32(16)
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v248)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v307
	v310 = l2 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v310))) = v303
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v286
	v314 = v286
	v319 = v259
	goto L103
L102:
	;
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v314 + int32(2)
	v332 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v332 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	F_pfree(m, v319)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L67
	} else {
		goto L140
	}
L105:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L67
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v336 = v17 + int32(72)
	v337 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v336))) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v17-int32(-64)))) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v337
	v349 = int32(0)
	F_div_var(m, v17+int32(32), l2, v17+int32(56), v349, v349, int32(1))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L67
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	v355 = v17 + int32(56)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_mul_var(m, l2, v355, v355, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L67
	} else {
		goto L110
	}
L110:
	;
	F_sub_var(m, v17+int32(32), v17+int32(56), v17+int32(8))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L67
	} else {
		goto L111
	}
L111:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v369 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_pfree(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L67
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v372 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	goto L114
L116:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v378 = F_palloc(m, v373<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L67
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	goto L104
L119:
	;
	v380 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v378))) = uint16(v380)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v380 < v382 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v389 = v382 << (uint(int32(1)) % 32)
	if v389 != 0 {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L122
L122:
	;
	F_pfree(m, v319)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L67
	} else {
		goto L127
	}
L123:
	;
	goto L122
L124:
	;
	v390 = F__emscripten_memcpy_bulkmem(m, v378+int32(2), v387, v389)
	mBase = m.M
	goto L126
L125:
	;
	goto L126
L126:
	;
	goto L123
L127:
	;
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v310)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v394
	v396 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v378
	v399 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v378 + v399
	v403 = v372 << (uint(int32(1)) % 32)
	v406 = F_palloc(m, v403+v399)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L67
	} else {
		goto L128
	}
L128:
	;
	v408 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v406))) = uint16(v408)
	if v408 < v372 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v403 != 0 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	goto L131
L131:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v417 != 0 {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	goto L131
L133:
	;
	v415 = F__emscripten_memcpy_bulkmem(m, v406+int32(2), v414, v403)
	mBase = m.M
	goto L135
L134:
	;
	goto L135
L135:
	;
	goto L132
L136:
	;
	F_pfree(m, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L67
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v420
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v422
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v310))) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v406
	v314 = v406
	v319 = v378
	goto L103
L139:
	;
	goto L138
L140:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v432 == int32(0) {
		goto L57
	} else {
		goto L141
	}
L141:
	;
	F_pfree(m, v432)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L67
	} else {
		goto L142
	}
L142:
	;
	goto L57
}
func F_gen_random_uuid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v4 = F_palloc(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = m.Env.Pgmem_random_bytes(m, v4, int32(16))
		mBase = m.M
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(156676), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494210), int32(541), int32(428133))
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
			}
		} else {
			v29 = v4 + int32(6)
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
			v34 = v30&int32(15) | int32(64)
			*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v34)
			v37 = v4 + int32(8)
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
			v42 = v38&int32(63) | int32(128)
			*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v42)
			return v4
		}
	}
}
func F_generateSerialExtraStmts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	v5 = l4
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = F_list_copy(m, l3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v225 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L3
	} else {
		goto L72
	}
L2:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v189 != 0 {
		goto L65
	} else {
		goto L66
	}
L3:
	;
	return
L4:
	;
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v178 = int32(0)
	v188 = v9
	goto L2
L6:
	;
	goto L7
L7:
	;
	v27 = v19
	v32 = v9
	v35 = v9
	v37 = v9
	goto L9
L8:
	;
	if v153 == int32(0) {
		v178 = v150
		v188 = v155
		goto L2
	} else {
		goto L53
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v38 <= v32 {
		v150 = v27
		v153 = v35
		v155 = v37
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_errorConflictingDefElem(m, v44, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L52
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v32<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v46 = int32(376207)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[243])))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v50 == int32(0) {
		v69 = v49
		v70 = v50
		goto L16
	} else {
		goto L17
	}
L12:
	;
	goto L10
L13:
	;
	if v141 != 0 {
		v27 = v141
		v32 = v142 + int32(1)
		v35 = v143
		v37 = v144
		goto L9
	} else {
		goto L51
	}
L14:
	;
	v139 = F_list_delete_nth_cell(m, v27, v32)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L3
	} else {
		goto L50
	}
L15:
	;
	if v70-v69 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	goto L15
L17:
	;
	if v49 != v50 {
		v69 = v49
		v70 = v50
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v54 = v45
	v55 = v46
	goto L19
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v59 == int32(0) {
		v69 = v58
		v70 = v59
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v69 = v58
	v70 = v59
	goto L16
L21:
	;
	v62 = int32(1)
	if v58 == v59 {
		v54 = v54 + v62
		v55 = v55 + v62
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v35 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v79 = int32(455057)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, _consts[244])))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v83 == int32(0) {
		v102 = v82
		v103 = v83
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v135 = v44
	v136 = v37
	goto L14
L27:
	;
	goto L28
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_errorConflictingDefElem(m, v44, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
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
	if v37 != 0 {
		goto L12
	} else {
		goto L49
	}
L31:
	;
	if v103-v102 == int32(0) {
		goto L30
	} else {
		goto L39
	}
L32:
	;
	goto L31
L33:
	;
	if v82 != v83 {
		v102 = v82
		v103 = v83
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v87 = v45
	v88 = v79
	goto L35
L35:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v92 == int32(0) {
		v102 = v91
		v103 = v92
		goto L32
	} else {
		goto L37
	}
L36:
	;
	v102 = v91
	v103 = v92
	goto L32
L37:
	;
	v95 = int32(1)
	if v91 == v92 {
		v87 = v87 + v95
		v88 = v88 + v95
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v107 = int32(454977)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, _consts[245])))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v111 == int32(0) {
		v130 = v110
		v131 = v111
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v131-v130 == int32(0) {
		goto L30
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	if v110 != v111 {
		v130 = v110
		v131 = v111
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v115 = v45
	v116 = v107
	goto L44
L44:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v120 == int32(0) {
		v130 = v119
		v131 = v120
		goto L41
	} else {
		goto L46
	}
L45:
	;
	v130 = v119
	v131 = v120
	goto L41
L46:
	;
	v123 = int32(1)
	if v119 == v120 {
		v115 = v115 + v123
		v116 = v116 + v123
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v141 = v27
	v142 = v32
	v143 = v35
	v144 = v37
	goto L13
L49:
	;
	v135 = v35
	v136 = v44
	goto L14
L50:
	;
	v141 = v139
	v142 = v32 - int32(1)
	v143 = v135
	v144 = v136
	goto L13
L51:
	;
	v150 = v141
	v153 = v143
	v155 = v144
	goto L8
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v159 = F_makeRangeVarFromNameList(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	if v161 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v164 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v173 = v161
	goto L57
L57:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v211 = v150
	v218 = v173
	v221 = v155
	v222 = v174
	goto L1
L58:
	;
	v171 = F_get_namespace_name(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L63
	}
L59:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+68))
	v170 = v166
	goto L58
L60:
	;
	goto L61
L61:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v168 = F_RangeVarGetCreationNamespace(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	v170 = v168
	goto L58
L63:
	;
	v173 = v171
	goto L57
L64:
	;
	v199 = F_get_namespace_name(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L70
	}
L65:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+48))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+68))
	v198 = v191
	goto L64
L66:
	;
	goto L67
L67:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v193 = F_RangeVarGetCreationNamespace(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_RangeVarAdjustRelationPersistence(m, v195, v193)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v198 = v193
	goto L64
L70:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v206 = F_ChooseRelationName(m, v202, v203, int32(228702), v198, int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v211 = v178
	v218 = v199
	v221 = v188
	v222 = v206
	goto L1
L72:
	;
	if v225 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v227
	F_errmsg_internal(m, int32(670150), v17+int32(16))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v247 != 0 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	F_errfinish(m, int32(494160), int32(486), int32(122177))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v221 != 0 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+48))
	v254 = v248 + int32(118)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v254 = v251 + int32(17)
	goto L78
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L3
	} else {
		goto L130
	}
L83:
	;
	if v255&int32(255) == int32(116) {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	v290 = v255
	goto L85
L85:
	;
	v292 = F_palloc0(m, int32(20))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L3
	} else {
		goto L98
	}
L86:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	v263 = int32(455057)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, _consts[244])))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v267 == int32(0) {
		v286 = v266
		v287 = v267
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v287-v286 != 0 {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	goto L87
L89:
	;
	if v266 != v267 {
		v286 = v266
		v287 = v267
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v271 = v262
	v272 = v263
	goto L91
L91:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	if v276 == int32(0) {
		v286 = v275
		v287 = v276
		goto L88
	} else {
		goto L93
	}
L92:
	;
	v286 = v275
	v287 = v276
	goto L88
L93:
	;
	v279 = int32(1)
	if v275 == v276 {
		v271 = v271 + v279
		v272 = v272 + v279
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v289 = int32(117)
	goto L97
L96:
	;
	v289 = int32(112)
	goto L97
L97:
	;
	v290 = v289
	goto L85
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v292)+16)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = int32(189)
	v298 = F_makeRangeVar(m, v218, v222, int32(-1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+4)) = v298
	*(*uint8)(unsafe.Add(mBase, uint32(v298)+17)) = uint8(v290)
	*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v211
	if l2 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v304 = F_makeTypeNameFromOid(m, l2)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L3
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v313 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v307 = F_makeDefElem(m, int32(172829), v304, int32(-1))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v292)+8))
	v310 = F_lcons(m, v307, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v310
	goto L102
L106:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+48))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+80))
	v317 = v315
	goto L108
L107:
	;
	v317 = int32(0)
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+12)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v320 = F_lappend(m, v319, v292)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v320
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v323
	v326 = F_palloc0(m, int32(16))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v326))) = int32(190)
	v331 = F_makeRangeVar(m, v218, v222, int32(-1))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v326)+4)) = v331
	v334 = F_makeString(m, v218)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v334
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v339 = F_makeString(m, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v339
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v343 = F_makeString(m, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v343
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v347
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v349
	v358 = F_list_make3_impl(m, v17+int32(12), v17+int32(8), v17+int32(4))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v361 = F_makeDefElem(m, int32(23848), v358, int32(-1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v361
	v366 = F_list_make1_impl(m, int32(1), v17)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v326)+12)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v326)+8)) = v366
	if l5 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if l6 != 0 {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v371 = F_lappend(m, v370, v326)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L3
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v375 = F_lappend(m, v374, v326)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L3
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v371
	goto L118
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v375
	goto L118
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v218
	goto L126
L125:
	;
	goto L126
L126:
	;
	if l7 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v222
	goto L129
L128:
	;
	goto L129
L129:
	;
	m.G0 = v17 + int32(48)
	return
L130:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(411495), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	F_parser_errposition(m, v394, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(494160), int32(505), int32(122177))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_generate_trgm(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	F_generate_trgm_only(m, v11+int32(4), l0, l1, v3)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v22)
	if int32(2) <= v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = v21 + int32(5)
	F_pg_qsort(m, v27, v20, int32(3), int32(6598))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v78 = v20
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v78*int32(12) + int32(20)
	m.G0 = v11 + int32(16)
	return v21
L6:
	;
	v34 = int32(1)
	v35 = v3
	goto L7
L7:
	;
	v41 = int32(3)
	v43 = v27 + v34*v41
	v48 = *(*int32)(unsafe.Add(mBase, _consts[1097]))
	v49 = m.T0[v48].(func(*base.Module, int32, int32) int32)(m, v43, v27+v35*v41)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v78 = v63 + int32(1)
	goto L5
L9:
	;
	v66 = v34 + int32(1)
	if v66 != v20 {
		v34 = v66
		v35 = v63
		goto L7
	} else {
		goto L13
	}
L10:
	;
	if v49 == int32(0) {
		v63 = v35
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v54 = v35 + int32(1)
	if v34 == v54 {
		v63 = v34
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v58 = v27 + v54*int32(3)
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
	*(*uint16)(unsafe.Add(mBase, uint32(v58))) = uint16(v59)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)) = uint8(v61)
	v63 = v54
	goto L9
L13:
	;
	goto L8
}
func F_getNextFlagFromString(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v16 = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v17 == v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = int32(2)
	goto L3
L2:
	;
	v20 = v16
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v24 = l2
	v27 = v21
	v28 = v20
	goto L7
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L20
	} else {
		goto L69
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L20
	} else {
		goto L65
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L20
	} else {
		goto L61
	}
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v221)
	m.G0 = v13 + int32(128)
	return
L9:
	;
	goto L8
L10:
	;
	v201 = F_pg_mblen_cstr(m, v27)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L20
	} else {
		goto L55
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if base.Ui32(v33) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v172 = v24
	goto L13
L13:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v180 != int32(1) {
		v213 = v172
		goto L9
	} else {
		goto L49
	}
L14:
	;
	if v33 != int32(2) {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v46 = F_strtox_2(m, v41, v13+int32(124), int32(10), int64(2147483648))
	mBase = m.M
	v47 = base.I32_wrap_i64(v46)
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v48 == v49 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v52 == int32(68) {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(int32(65537)) <= base.Ui32(v47) {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v47
	v61 = F_pg_sprintf(m, v24, int32(460864), v13+int32(112))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v48
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v64 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v172 = v24 + v61
	goto L13
L23:
	;
	v71 = v64
	v72 = v48
	v74 = int32(0)
	goto L24
L24:
	;
	if base.Ui32((v71-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L22
L26:
	;
	if v74 != 0 {
		goto L22
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v103 = v71 & int32(255)
	if base.Ui32(v103-int32(9)) < base.Ui32(int32(5)) {
		v152 = v74
		goto L34
	} else {
		goto L35
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v90
	F_errmsg(m, int32(691342), v13-int32(-64))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(491922), int32(402), int32(327142))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v153 = F_pg_mblen_cstr(m, v72)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L20
	} else {
		goto L47
	}
L35:
	;
	switch v103 - int32(32) {
	case 0:
		v152 = v74
		goto L34
	default:
		goto L36
	case 12:
		goto L37
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L43
	}
L37:
	;
	if v74 == int32(0) {
		v152 = int32(1)
		goto L34
	} else {
		goto L38
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v120
	F_errmsg(m, int32(691342), v13+int32(96))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(491922), int32(411), int32(327142))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v139
	F_errmsg(m, int32(691305), v13+int32(80))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(491922), int32(419), int32(327142))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v156 = v153 + v155
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v156
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v158 != 0 {
		v71 = v158
		v72 = v156
		v74 = v152
		goto L24
	} else {
		goto L48
	}
L48:
	;
	goto L25
L49:
	;
	if v28 <= int32(0) {
		v213 = v172
		goto L9
	} else {
		goto L50
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v21
	F_errmsg(m, int32(342213), v13)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(491922), int32(439), int32(327142))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	if v201 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v206 = v201 + v205
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v206
	v208 = v204 + v201
	v210 = v28 - int32(1)
	if v210 != 0 {
		v24 = v208
		v27 = v206
		v28 = v210
		goto L7
	} else {
		goto L60
	}
L57:
	;
	v203 = F__emscripten_memcpy_bulkmem(m, v24, v27, v201)
	mBase = m.M
	v204 = v203
	goto L59
L58:
	;
	v204 = v24
	goto L59
L59:
	;
	goto L56
L60:
	;
	v213 = v208
	goto L9
L61:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v233
	F_errmsg(m, int32(691342), v13+int32(32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(491922), int32(384), int32(327142))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v252
	F_errmsg(m, int32(398022), v13+int32(48))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L20
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(491922), int32(389), int32(327142))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v268
	F_errmsg_internal(m, int32(481576), v13+int32(16))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(491922), int32(428), int32(327142))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_getSubscriptingRoutines(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v6 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
			v12 = v10 + v11
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
			if l1 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
			} else {
			}
			F_ReleaseCatCache(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v13 != 0 {
					v27 = F_OidFunctionCall0Coll(m, v13)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v31 = v27
						return v31
					}
				} else {
					return int32(0)
				}
			}
		} else {
			v20 = int32(0)
			if l1 == v20 {
				v31 = v20
				return v31
			} else {
				v23 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
				return v23
			}
		}
	}
}
func F_get_ENR(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v3 = int32(0)
	if l0 == v3 {
		v63 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v63
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(0) {
		v63 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v11 <= int32(0) {
		v63 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = int32(0)
	if v14 < v11 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v17 = v11
	goto L7
L6:
	;
	v17 = v14
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v20 = int32(0)
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v33 == int32(0) {
		v52 = v32
		v53 = v33
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v63 = int32(0)
	goto L1
L10:
	;
	if v53-v52 == int32(0) {
		v63 = v28
		goto L1
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v32 != v33 {
		v52 = v32
		v53 = v33
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = v29
	v38 = l1
	goto L14
L14:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v41
		v53 = v42
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v52 = v41
	v53 = v42
	goto L11
L16:
	;
	v45 = int32(1)
	if v41 == v42 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v58 = v20 + int32(1)
	if v58 != v17 {
		v20 = v58
		goto L8
	} else {
		goto L19
	}
L19:
	;
	goto L9
}
func F_get_atttype(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache2(m, int32(7), l0, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+68))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_cheapest_parallel_safe_total_inner(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v44
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 <= int32(0) {
		v44 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v44 = int32(0)
	goto L1
L5:
	;
	v9 = int32(0)
	if v9 < v6 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v12 = v6
	goto L8
L7:
	;
	v12 = v9
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = int32(0)
	goto L9
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15<<(uint(int32(2))%32))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+21)))
	if v24 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v27 == int32(0) {
		v44 = v23
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v35 = v15 + int32(1)
	if v35 != v12 {
		v15 = v35
		goto L9
	} else {
		goto L16
	}
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 == int32(0) {
		v44 = v23
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	goto L10
}
func F_get_compatible_hash_operators(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	v3 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v13 = int32(0)
	v15 = F_SearchSysCacheList(m, int32(3), int32(1), l0, v13, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_ReleaseCatCacheList(m, v15)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L27
	}
L5:
	;
	return int32(0)
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v19 <= int32(0) {
		v86 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v27 = v3
	v30 = v3
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(48)+v30<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	v38 = v36 + v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	if v39 != int32(405) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v86 = v78
	goto L4
L10:
	;
	v80 = v30 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v80 < v81 {
		v27 = v78
		v30 = v80
		goto L8
	} else {
		goto L26
	}
L11:
	;
	v78 = v27
	goto L10
L12:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	if v42 != int32(1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v45 == v46 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = int32(1)
	if l1 == int32(0) {
		v86 = v48
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
	v86 = v48
	goto L4
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v57 = F_SearchSysCache4(m, int32(4), v55, v45, v45, int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v68
	v70 = int32(0)
	v72 = base.B2i32(v68 != v70) | v27
	if v68 == v70 {
		v78 = v72
		goto L10
	} else {
		goto L25
	}
L20:
	;
	if v57 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v68 = int32(0)
	goto L19
L22:
	;
	goto L23
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)+20))
	F_ReleaseCatCache(m, v57)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v68 = v65
	goto L19
L25:
	;
	v86 = v72
	goto L4
L26:
	;
	goto L9
L27:
	;
	return v86 & int32(1)
}
func F_get_const_expr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v15 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(96)
	return
L2:
	;
	F_appendStringInfoString(m, v14, int32(527204))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_getTypeOutputInfo(m, v48, v12+int32(92), v12+int32(91))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L15
	}
L5:
	;
	return
L6:
	;
	if l2 < int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = F_format_type_with_typemod(m, v23, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v25
	F_appendStringInfo(m, v14, int32(175247), v12+int32(16))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v33 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = F_get_typcollation(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v38 == v40 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v42 = F_generate_collation_name(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v42
	F_appendStringInfo(m, v36, int32(196527), v12)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = F_OidOutputFunctionCall(m, v55, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v59 - int32(16) {
	case 0:
		goto L23
	case 1, 2, 3, 4, 5, 6:
		goto L21
	case 7:
		goto L24
	default:
		goto L22
	}
L17:
	;
	F_pfree(m, v57)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L91
	}
L18:
	;
	F_appendStringInfoString(m, v14, v57)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L90
	}
L19:
	;
	v266 = int32(1)
	goto L17
L20:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if base.Ui32((v115-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L48
	} else {
		goto L49
	}
L21:
	;
	F_appendStringInfoChar(m, v14, int32(39))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L33
	}
L22:
	;
	if v59 == int32(1700) {
		goto L20
	} else {
		goto L32
	}
L23:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v71 != int32(116) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v62 != int32(45) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v57
	F_appendStringInfo(m, v14, int32(665535), v12-int32(-64))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	F_appendStringInfoString(m, v14, int32(357375))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L31
	}
L28:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v74 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_appendStringInfoString(m, v14, int32(340544))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v266 = int32(0)
	goto L17
L31:
	;
	v266 = int32(0)
	goto L17
L32:
	;
	goto L21
L33:
	;
	v94 = v57
	goto L34
L34:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v98 = base.I32_extend8_s(v97)
	if v97 != int32(39) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_appendStringInfoChar(m, v14, v98)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L47
	}
L37:
	;
	if v97 != int32(92) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	F_appendStringInfoChar(m, v14, v98)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L46
	}
L40:
	;
	if v97 != 0 {
		goto L36
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[848])))
	if v108 != 0 {
		goto L36
	} else {
		goto L45
	}
L43:
	;
	F_appendStringInfoChar(m, v14, int32(39))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v266 = int32(0)
	goto L17
L45:
	;
	goto L39
L46:
	;
	goto L36
L47:
	;
	v94 = v94 + int32(1)
	goto L34
L48:
	;
	v122 = int32(637241)
	v126 = m.G0
	v128 = v126 - int32(32)
	m.G0 = v128
	v130 = int32(*(*int8)(unsafe.Add(mBase, _consts[849])))
	if v130 != 0 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v57
	F_appendStringInfo(m, v14, int32(665535), v12+int32(80))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L89
	}
L51:
	;
	if v57&int32(3) == int32(0) {
		v212 = v57
		goto L73
	} else {
		goto L74
	}
L52:
	;
	m.G0 = v128 + int32(32)
	goto L51
L53:
	;
	v135 = F___memset(m, v128, int32(0), int32(32))
	mBase = m.M
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _consts[849])))
	if v136 != 0 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, _consts[850])))
	if v131 != 0 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v132 = F___strchrnul(m, v57, v130)
	mBase = m.M
	v183 = v132
	goto L52
L57:
	;
	goto L56
L58:
	;
	v138 = v122
	v139 = v136
	goto L61
L59:
	;
	goto L60
L60:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v160 == int32(0) {
		v183 = v57
		goto L52
	} else {
		goto L64
	}
L61:
	;
	v146 = v128 + int32(base.Ui32(v139)>>(uint(int32(3))%32))&int32(28)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v147 | v148<<(uint(v139)%32)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v152 != 0 {
		v138 = v138 + v148
		v139 = v152
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	goto L62
L64:
	;
	v164 = v57
	v165 = v160
	goto L65
L65:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(base.Ui32(v165)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v173)>>(uint(v165)%32))&int32(1) != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v183 = v179
	goto L52
L67:
	;
	v183 = v164
	goto L52
L68:
	;
	goto L69
L69:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	v179 = v164 + int32(1)
	if v177 != 0 {
		v164 = v179
		v165 = v177
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	if v183-v57 != v245 {
		goto L18
	} else {
		goto L88
	}
L72:
	;
	v245 = v237 - v57
	goto L71
L73:
	;
	v216 = v212
	goto L82
L74:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v196 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v245 = int32(0)
	goto L71
L76:
	;
	goto L77
L77:
	;
	v201 = v57
	goto L78
L78:
	;
	v205 = v201 + int32(1)
	if v205&int32(3) == int32(0) {
		v212 = v205
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v237 = v205
	goto L72
L80:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v210 != 0 {
		v201 = v205
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v225 = int32(-2139062144)
	if (int32(16843008)-v222|v222)&v225 == v225 {
		v216 = v216 + int32(4)
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v231 = v216
	goto L85
L84:
	;
	goto L83
L85:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v235 != 0 {
		v231 = v231 + int32(1)
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v237 = v231
	goto L72
L87:
	;
	goto L86
L88:
	;
	goto L50
L89:
	;
	goto L19
L90:
	;
	v266 = int32(0)
	goto L17
L91:
	;
	if l2 < int32(0) {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v271 - int32(16) {
	case 0:
		goto L97
	case 1, 2, 3, 4, 5, 6:
		goto L94
	case 7:
		v283 = v266
		goto L95
	default:
		goto L98
	}
L93:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v297 == int32(0) {
		goto L1
	} else {
		goto L105
	}
L94:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v288 = F_format_type_with_typemod(m, v271, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L103
	}
L95:
	;
	if l2 != 0 {
		goto L94
	} else {
		goto L101
	}
L96:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = v266 | base.B2i32(int32(0) <= v279)
	goto L95
L97:
	;
	v283 = int32(0)
	goto L95
L98:
	;
	if v271 == int32(1700) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	if v271 != int32(705) {
		goto L94
	} else {
		goto L100
	}
L100:
	;
	goto L97
L101:
	;
	if v283 == int32(0) {
		goto L93
	} else {
		goto L102
	}
L102:
	;
	goto L94
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v288
	F_appendStringInfo(m, v14, int32(175247), v12+int32(48))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	goto L93
L105:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v302 = F_get_typcollation(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v302 == v304 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v306 = F_generate_collation_name(m, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v306
	F_appendStringInfo(m, v300, int32(196527), v12+int32(32))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	goto L1
}
func F_get_dirent_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	switch v12 - int32(4) {
	case 0:
		v64 = int32(3)
		m.G0 = v9 + int32(112)
		return v64
	default:
		if l2 == int32(0) {
			v28 = F___fstatat(m, int32(-100), l0, v9+int32(16), int32(256))
			mBase = m.M
			v29 = v28
		} else {
			v23 = F___fstatat(m, int32(-100), l0, v9+int32(16), int32(0))
			mBase = m.M
			v29 = v23
		}
		if v29 < int32(0) {
			v32 = int32(0)
			v34 = F_errstart(m, l3, v32)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				if v34 == int32(0) {
					v64 = v32
					m.G0 = v9 + int32(112)
					return v64
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(294607), v9)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488559), int32(592), int32(362286))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v64 = v32
								m.G0 = v9 + int32(112)
								return v64
							}
						}
					}
				}
			}
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			v53 = v51 & int32(61440)
			if v53 != int32(40960) {
				if v53 == int32(32768) {
					v64 = int32(2)
				} else {
					if v53 != int32(16384) {
						v64 = int32(1)
					} else {
						v64 = int32(3)
					}
				}
			} else {
				v64 = int32(4)
			}
			m.G0 = v9 + int32(112)
			return v64
		}
	case 4:
		v64 = int32(2)
		m.G0 = v9 + int32(112)
		return v64
	case 6:
		if l2 != 0 {
			v23 = F___fstatat(m, int32(-100), l0, v9+int32(16), int32(0))
			mBase = m.M
			v29 = v23
			if v29 < int32(0) {
				v32 = int32(0)
				v34 = F_errstart(m, l3, v32)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v34 == int32(0) {
						v64 = v32
						m.G0 = v9 + int32(112)
						return v64
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg(m, int32(294607), v9)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(488559), int32(592), int32(362286))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v64 = v32
									m.G0 = v9 + int32(112)
									return v64
								}
							}
						}
					}
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				v53 = v51 & int32(61440)
				if v53 != int32(40960) {
					if v53 == int32(32768) {
						v64 = int32(2)
					} else {
						if v53 != int32(16384) {
							v64 = int32(1)
						} else {
							v64 = int32(3)
						}
					}
				} else {
					v64 = int32(4)
				}
				m.G0 = v9 + int32(112)
				return v64
			}
		} else {
			v64 = int32(4)
			m.G0 = v9 + int32(112)
			return v64
		}
	}
}
func F_get_policies_for_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v163 int32
	_ = v163
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v6
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if int32(0) < v24 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v177 = v6
	goto L3
L3:
	;
	F_list_sort(m, v177, int32(1059))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L20
	} else {
		goto L41
	}
L4:
	;
	v37 = v6
	goto L7
L5:
	;
	goto L6
L6:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v177 = v163
	goto L3
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v37<<(uint(int32(2))%32))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v47 == int32(42) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v147 = v37 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v147 < v148 {
		v37 = v147
		goto L7
	} else {
		goto L40
	}
L10:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v72 != 0 {
		goto L26
	} else {
		goto L27
	}
L11:
	;
	switch l1 - int32(1) {
	case 0:
		goto L16
	case 1:
		goto L14
	case 2:
		goto L15
	case 3:
		goto L12
	case 4:
		goto L9
	default:
		goto L13
	}
L12:
	;
	if v47 != int32(100) {
		goto L9
	} else {
		goto L24
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if v47 == int32(119) {
		goto L10
	} else {
		goto L19
	}
L15:
	;
	if v47 == int32(97) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	if v47 == int32(114) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	goto L9
L19:
	;
	goto L9
L20:
	;
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg_internal(m, int32(471388), v16)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(486917), int32(590), int32(260434))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	goto L10
L25:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+12)))
	if v127 != 0 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v80 = v72
	goto L28
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v80 = (v73<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L28
L28:
	;
	v81 = v80 + v71
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v82 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v85 = int32(0)
	v87 = v71 + int32(16)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 <= v85 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v96 = v85
	goto L31
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v81+v96<<(uint(int32(2))%32))))
	v108 = F_has_privs_of_role(m, l2, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L20
	} else {
		goto L33
	}
L32:
	;
	goto L9
L33:
	;
	if v108 != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v111 = v96 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v111 < v112 {
		v96 = v111
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v128 = l3
	goto L38
L37:
	;
	v128 = l4
	goto L38
L38:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v130 = F_lappend(m, v129, v46)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v130
	goto L9
L40:
	;
	goto L8
L41:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[553]))
	if v182 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[554]))
	if v305 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L43:
	;
	v185 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l1, l0)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	F_list_sort(m, v185, int32(1059))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	if v185 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v192 <= int32(0) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v204 = int32(0)
	goto L48
L48:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v204<<(uint(int32(2))%32))))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+8))
	if v215 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L42
L50:
	;
	v288 = v204 + int32(1)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v288 < v289 {
		v204 = v288
		goto L48
	} else {
		goto L63
	}
L51:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v271 = F_lappend(m, v270, v213)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L20
	} else {
		goto L62
	}
L52:
	;
	v223 = v215
	goto L54
L53:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v223 = (v216<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L54
L54:
	;
	v224 = v223 + v214
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	if v225 == int32(0) {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v228 = int32(0)
	v230 = v214 + int32(16)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v231 <= v228 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v239 = v228
	goto L57
L57:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v224+v239<<(uint(int32(2))%32))))
	v251 = F_has_privs_of_role(m, l2, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L59
	}
L58:
	;
	goto L50
L59:
	;
	if v251 != 0 {
		goto L51
	} else {
		goto L60
	}
L60:
	;
	v254 = v239 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v254 < v255 {
		v239 = v254
		goto L57
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v271
	goto L50
L63:
	;
	goto L49
L64:
	;
	m.G0 = v16 + int32(16)
	return
L65:
	;
	v308 = m.T0[v305].(func(*base.Module, int32, int32) int32)(m, l1, l0)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	if v308 == int32(0) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	if v312 <= int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v324 = int32(0)
	goto L69
L69:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329+v324<<(uint(int32(2))%32))))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	if v335 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	goto L64
L71:
	;
	v408 = v324 + int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	if v408 < v409 {
		v324 = v408
		goto L69
	} else {
		goto L84
	}
L72:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v391 = F_lappend(m, v390, v333)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L20
	} else {
		goto L83
	}
L73:
	;
	v343 = v335
	goto L75
L74:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	v343 = (v336<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L75
L75:
	;
	v344 = v343 + v334
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v345 == int32(0) {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v348 = int32(0)
	v350 = v334 + int32(16)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	if v351 <= v348 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v359 = v348
	goto L78
L78:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v344+v359<<(uint(int32(2))%32))))
	v371 = F_has_privs_of_role(m, l2, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L20
	} else {
		goto L80
	}
L79:
	;
	goto L71
L80:
	;
	if v371 != 0 {
		goto L72
	} else {
		goto L81
	}
L81:
	;
	v374 = v359 + int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	if v374 < v375 {
		v359 = v374
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v391
	goto L71
L84:
	;
	goto L70
}
func F_get_position(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 float64
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v51 float64
	_ = v51
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v58 float64
	_ = v58
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v10 == int32(0) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
		if v9&int32(1) != 0 {
			if v13&int32(1) == int32(0) {
				return float64(0)
			} else {
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
				if v82 != 0 {
					v83 = float64(0)
				} else {
					v83 = float64(1)
				}
				v85 = v83
				return v85
			}
		} else {
			v16 = float64(0.5)
			if v13&int32(1) != 0 {
				v85 = v16
				return v85
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				if v19 == int32(0) {
					v85 = v16
					return v85
				} else {
					v23 = l0 + int32(268)
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v27 = F_FunctionCall2Coll(m, v23, v24, v25, v26)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return float64(0)
					} else {
						v31 = *(*float64)(unsafe.Add(mBase, uint32(v27)))
						if base.F64_le(v31, float64(0)) != 0 {
							v85 = v16
							return v85
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v31)&int64(9223372036854775807)) {
								v85 = v16
								return v85
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v42 = F_FunctionCall2Coll(m, v23, v39, v40, v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return float64(0)
								} else {
									v44 = *(*float64)(unsafe.Add(mBase, uint32(v42)))
									v45 = base.F64_div(v44, v31)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v45)&int64(9223372036854775807)) {
										v85 = v16
										return v85
									} else {
										v51 = float64(0)
										if base.F64_gt(v45, v51) != 0 {
											v54 = v45
										} else {
											v54 = v51
										}
										v55 = float64(1)
										if base.F64_lt(v54, v55) != 0 {
											v58 = v54
										} else {
											v58 = v55
										}
										return v58
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		if v9&int32(1) != 0 {
			return float64(0.5)
		} else {
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			if v64 != int32(1) {
				return float64(1)
			} else {
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
				if v71 != 0 {
					v72 = float64(0)
				} else {
					v72 = float64(1)
				}
				return v72
			}
		}
	}
}
func F_get_promoted_array_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10)+96))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v12 != 0 {
					v36 = v12
					return v36
				} else {
					v17 = F_SearchSysCache1(m, int32(82), l0)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						if v17 == int32(0) {
							return int32(0)
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
							v25 = v23 + v24
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
							if v26 != 0 {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
								if v28 == int32(6179) {
									v31 = l0
								} else {
									v31 = int32(0)
								}
								v33 = v31
							} else {
								v33 = int32(0)
							}
							F_ReleaseCatCache(m, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = v33
								return v36
							}
						}
					}
				}
			}
		} else {
			v17 = F_SearchSysCache1(m, int32(82), l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
					v25 = v23 + v24
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
					if v26 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
						if v28 == int32(6179) {
							v31 = l0
						} else {
							v31 = int32(0)
						}
						v33 = v31
					} else {
						v33 = int32(0)
					}
					F_ReleaseCatCache(m, v17)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = v33
						return v36
					}
				}
			}
		}
	}
}
func F_get_reloptions(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_deconstruct_array_builtin(m, v12, int32(25), v10+int32(12), int32(0), v10+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if int32(0) < v22 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v10 + int32(16)
	return
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v28<<(uint(int32(2))%32))))
	v37 = F_text_to_cstring(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v39 = int32(61)
	v40 = F___strchrnul(m, v37, v39)
	mBase = m.M
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v42 == v39 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v46 = v40
	goto L13
L12:
	;
	v46 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v47)
	v52 = v46 + int32(1)
	goto L16
L15:
	;
	v52 = int32(735586)
	goto L16
L16:
	;
	if v28 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(724350))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v56 = F_quote_identifier(m, v37)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v56
	F_appendStringInfo(m, l0, int32(540211), v10)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v62 = F_quote_identifier(m, v52)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	F_pfree(m, v37)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L44
	}
L24:
	;
	if v62 == v52 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_appendStringInfoString(m, l0, v52)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_appendStringInfoChar(m, l0, int32(39))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L23
L29:
	;
	v71 = v52
	goto L30
L30:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v78 = base.I32_extend8_s(v77)
	if v77 != int32(39) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	F_appendStringInfoChar(m, l0, v78)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L43
	}
L33:
	;
	if v77 != int32(92) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	F_appendStringInfoChar(m, l0, v78)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L42
	}
L36:
	;
	if v77 != 0 {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _consts[848])))
	if v87 != 0 {
		goto L32
	} else {
		goto L41
	}
L39:
	;
	F_appendStringInfoChar(m, l0, int32(39))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L23
L41:
	;
	goto L35
L42:
	;
	goto L32
L43:
	;
	v71 = v71 + int32(1)
	goto L30
L44:
	;
	v104 = v28 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v104 < v105 {
		v28 = v104
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L8
}
func F_get_rolespec_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = F_get_rolespec_tuple(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
		v12 = F_pstrdup(m, v7+v8+int32(4))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_ReleaseCatCache(m, v3)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_get_share_path(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_make_relative_path(m, l0, int32(297155), int32(4472000))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_get_steps_using_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l6 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l3
		v21 = F_list_make1_impl(m, int32(1), v12+int32(16))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l4
			v30 = F_list_make1_impl(m, int32(472), v12+int32(12))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v33 = F_palloc0(m, int32(24))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(377)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v37 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v30
					*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v21
					if l2 != 0 {
						v45 = int32(0)
					} else {
						v45 = l1
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v33)+8)) = uint16(v45)
					*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v37
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v49 = F_lappend(m, v48, v33)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v33
						v57 = F_list_make1_impl(m, int32(1), v12+int32(8))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							v68 = v57
							m.G0 = v12 + int32(32)
							return v68
						}
					}
				}
			}
		}
	} else {
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
		v60 = int32(0)
		v62 = F_get_steps_using_prefix_recurse(m, l0, l1, l2, l3, l4, l5, l6, v59, v60, v60)
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return int32(0)
		} else {
			v68 = v62
			m.G0 = v12 + int32(32)
			return v68
		}
	}
}
func F_get_tle_by_resno(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v40
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 <= int32(0) {
		v40 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v40 = int32(0)
	goto L1
L5:
	;
	v9 = int32(0)
	if v9 < v6 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v12 = v6
	goto L8
L7:
	;
	v12 = v9
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = int32(0)
	goto L9
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13+v17<<(uint(int32(2))%32))))
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)))
	if v26 == l1&int32(65535) {
		v40 = v25
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v29 = v17 + int32(1)
	if v29 != v12 {
		v17 = v29
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_get_tsearch_config_filename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	v5 = m.G0
	v7 = v5 - int32(1056)
	m.G0 = v7
	v9 = int32(501853)
	v13 = m.G0
	v15 = v13 - int32(32)
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v16
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[680])))
	if v24 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l0&int32(3) == int32(0) {
		v116 = l0
		goto L24
	} else {
		goto L25
	}
L2:
	;
	v92 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[681])))
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = l0
	goto L8
L6:
	;
	goto L7
L7:
	;
	v42 = v9
	v43 = v24
	goto L11
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v38 == v24 {
		v32 = v32 + int32(1)
		goto L8
	} else {
		goto L10
	}
L9:
	;
	v92 = v32 - l0
	goto L1
L10:
	;
	goto L9
L11:
	;
	v50 = v15 + int32(base.Ui32(v43)>>(uint(int32(3))%32))&int32(28)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51 | v52<<(uint(v43)%32)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v56 != 0 {
		v42 = v42 + v52
		v43 = v56
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v59 == int32(0) {
		v84 = l0
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v92 = v84 - l0
	goto L1
L15:
	;
	v63 = l0
	v64 = v59
	goto L16
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(base.Ui32(v64)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v72)>>(uint(v64)%32))&int32(1) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v84 = v80
	goto L14
L18:
	;
	v84 = v63
	goto L14
L19:
	;
	goto L20
L20:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v80 = v63 + int32(1)
	if v78 != 0 {
		v63 = v80
		v64 = v78
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	if v92 != v149 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v149 = v141 - l0
	goto L22
L24:
	;
	v120 = v116
	goto L33
L25:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v100 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v149 = int32(0)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v105 = l0
	goto L29
L29:
	;
	v109 = v105 + int32(1)
	if v109&int32(3) == int32(0) {
		v116 = v109
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v141 = v109
	goto L23
L31:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v114 != 0 {
		v105 = v109
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v129 = int32(-2139062144)
	if (int32(16843008)-v126|v126)&v129 == v129 {
		v120 = v120 + int32(4)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v135 = v120
	goto L36
L35:
	;
	goto L34
L36:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v139 != 0 {
		v135 = v135 + int32(1)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v141 = v135
	goto L23
L38:
	;
	goto L37
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	F_get_share_path(m, v7+int32(32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L42
	} else {
		goto L47
	}
L42:
	;
	return int32(0)
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	F_errmsg(m, int32(693230), v7+int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(488548), int32(53), int32(373899))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v176 = F_palloc(m, int32(1024))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L42
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(32)
	v185 = F_pg_snprintf(m, v176, int32(1024), int32(175531), v7)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L42
	} else {
		goto L49
	}
L49:
	;
	m.G0 = v7 + int32(1056)
	return v176
}
func F_gettoken_query_standard(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v224 int32
	_ = v224
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v22 = l0 + int32(8)
	goto L1
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v35 - int32(1) {
	case 0, 2:
		goto L9
	case 1:
		goto L8
	default:
		goto L4
	}
L3:
	;
	v379 = F_pg_mblen_cstr(m, v373)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L14
	} else {
		goto L88
	}
L4:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v373 = v366
	goto L3
L5:
	;
	m.G0 = v15 + int32(32)
	return v358
L6:
	;
	v358 = int32(3)
	goto L5
L7:
	;
	v330 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38 + v330
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v330)
	goto L6
L8:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v164 != int32(124) {
		goto L46
	} else {
		goto L47
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v39-int32(9)) < base.Ui32(int32(5)) {
		v373 = v38
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v44 = int32(1)
	switch v39 - int32(32) {
	case 0:
		v373 = v38
		goto L3
	case 1:
		goto L7
	default:
		goto L11
	case 8:
		goto L12
	case 26:
		v358 = v44
		goto L5
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v38
	goto L13
L12:
	;
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38 + v47
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v52 + v47
	v358 = int32(4)
	goto L5
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v60 = int32(0)
	v62 = F_gettoken_tsvector(m, v59, l3, l2, v60, v60, v22)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v62 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v67 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v67)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v67)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 != int32(58) {
		v125 = v66
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v135 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L19:
	;
	v131 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v125
	v358 = v131
	goto L5
L20:
	;
	v75 = v66 + int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v76 == int32(0) {
		v125 = v75
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v85 = v75
	goto L22
L22:
	;
	v91 = F_pg_mblen_cstr(m, v85)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L14
	} else {
		goto L24
	}
L23:
	;
	v125 = v117
	goto L19
L24:
	;
	if v91 != int32(1) {
		v125 = v85
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	switch v95 - int32(42) {
	case 0:
		goto L27
	default:
		v125 = v85
		goto L19
	case 23, 55:
		goto L31
	case 24, 56:
		goto L30
	case 25, 57:
		goto L29
	case 26, 58:
		goto L28
	}
L26:
	;
	v117 = v85 + int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v118 != 0 {
		v85 = v117
		goto L22
	} else {
		goto L32
	}
L27:
	;
	v114 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v114)
	goto L26
L28:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	v112 = v110 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v112)
	goto L26
L29:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	v108 = v106 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v108)
	goto L26
L30:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	v104 = v102 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v104)
	goto L26
L31:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	v100 = v98 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v100)
	goto L26
L32:
	;
	goto L23
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v142 == int32(3) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v138 != int32(447) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+4)))
	if v141 != 0 {
		v358 = v44
		goto L5
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v358 = int32(0)
	goto L5
L38:
	;
	goto L39
L39:
	;
	v146 = F_errsave_start(m, v135)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	if v146 == int32(0) {
		v358 = v44
		goto L5
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v153
	F_errmsg(m, int32(703146), v15)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, v135, int32(486956), int32(345), int32(417252))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	v358 = v44
	goto L5
L45:
	;
	if v164 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	if v164 != int32(38) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v176 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163 + v176
	v181 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v181)
	v358 = v181
	goto L5
L49:
	;
	v169 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163 + v169
	v174 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v174)
	goto L6
L50:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v296 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L51:
	;
	v194 = v163
	v195 = v164
	v196 = int32(0)
	v199 = int32(1)
	goto L52
L52:
	;
	switch v196 - int32(1) {
	case 0:
		goto L59
	case 1:
		v211 = v194
		v212 = v195
		goto L58
	case 2:
		goto L55
	default:
		goto L57
	}
L53:
	;
	goto L50
L54:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v283 != 0 {
		v194 = v279
		v195 = v283
		v196 = v281
		v199 = v282
		goto L52
	} else {
		goto L76
	}
L55:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v199)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	v277 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v277)
	goto L6
L56:
	;
	if base.Ui32(int32(9)) < base.Ui32((v195-int32(48))&int32(255)) {
		goto L50
	} else {
		goto L64
	}
L57:
	;
	if v195&int32(255) != int32(60) {
		goto L50
	} else {
		goto L63
	}
L58:
	;
	if v212&int32(255) != int32(62) {
		goto L50
	} else {
		goto L62
	}
L59:
	;
	if v195&int32(255) != int32(45) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v207 = v194 + int32(1)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v208 == int32(0) {
		goto L50
	} else {
		goto L61
	}
L61:
	;
	v211 = v207
	v212 = v208
	goto L58
L62:
	;
	v279 = v211 + int32(1)
	v281 = int32(3)
	v282 = v199
	goto L54
L63:
	;
	v224 = int32(1)
	v279 = v194 + v224
	v281 = v224
	v282 = v199
	goto L54
L64:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(0)
	v240 = F_strtox_2(m, v194, v15+int32(28), int32(10), int64(2147483648))
	mBase = m.M
	v241 = base.I32_wrap_i64(v240)
	goto L65
L65:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v194 == v242 {
		goto L50
	} else {
		goto L66
	}
L66:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v245 != int32(68) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if base.Ui32(v241) < base.Ui32(int32(16385)) {
		v279 = v242
		v281 = int32(2)
		v282 = v241
		goto L54
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v254 = F_errsave_start(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L14
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	if v254 == int32(0) {
		goto L50
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L14
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(16384)
	F_errmsg(m, int32(339814), v15+int32(16))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L14
	} else {
		goto L74
	}
L74:
	;
	F_errsave_finish(m, v253, int32(486956), int32(211), int32(207320))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L14
	} else {
		goto L75
	}
L75:
	;
	goto L50
L76:
	;
	goto L53
L77:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if base.Ui32(v307-int32(9)) < base.Ui32(int32(5)) {
		v373 = v306
		goto L3
	} else {
		goto L81
	}
L78:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	if v299 != int32(447) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+4)))
	if v302 == int32(0) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v358 = int32(1)
	goto L5
L81:
	;
	v312 = int32(1)
	switch v307 - int32(32) {
	case 0:
		v373 = v306
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8:
		v358 = v312
		goto L5
	case 9:
		goto L83
	default:
		goto L82
	}
L82:
	;
	if v307 != 0 {
		v358 = v312
		goto L5
	} else {
		goto L87
	}
L83:
	;
	v315 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v306 + v315
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v320 = v318 - v315
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v320
	if v320 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v326 = v315
	goto L86
L85:
	;
	v326 = int32(5)
	goto L86
L86:
	;
	v358 = v326
	goto L5
L87:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v358 = base.B2i32(v327 != int32(0))
	goto L5
L88:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v379 + v381
	goto L1
}
func F_gimme_tree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v4 = int32(0)
	if l2 <= v4 {
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v17 = v4
	v19 = v4
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v24 = int32(2)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1+v17<<(uint(v24)%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23+v27<<(uint(v24)%32)-int32(4))))
	v35 = F_palloc(m, int32(8))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v48 = int32(0)
	if v43 == v48 {
		v96 = v48
		goto L10
	} else {
		goto L11
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v33
	v43 = F_merge_clump(m, l0, v19, v35, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v46 = v17 + int32(1)
	if v46 != l2 {
		v17 = v46
		v19 = v43
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	return v96
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v51 < int32(2) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v84 != int32(1) {
		v96 = v48
		goto L10
	} else {
		goto L23
	}
L13:
	;
	v84 = v51
	v85 = v43
	goto L12
L14:
	;
	goto L15
L15:
	;
	v54 = int32(0)
	v59 = v54
	v60 = v54
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
	v70 = F_merge_clump(m, l0, v60, v68, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	if v70 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v73 = v59 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v73 < v74 {
		v59 = v73
		v60 = v70
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(0)
L21:
	;
	goto L22
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v84 = v80
	v85 = v70
	goto L12
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v96 = v93
	goto L10
}
func F_ginarrayextract_2args(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v13 <= int32(2) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(120551), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494319), int32(71), int32(154021))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v32 = F_pg_detoast_datum_copy(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
			F_get_typlenbyvalalign(m, v36, v11+int32(14), v11+int32(13), v11+int32(12))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+14)))
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)))
				v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+12)))
				F_deconstruct_array(m, v32, v46, v47, v48, v11+int32(8), v11+int32(4), v11)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v35))) = v55
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v34))) = v57
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					m.G0 = v11 + int32(16)
					return v59
				}
			}
		}
	}
}
func F_ginarraytriconsistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = int32(2)
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	switch v17 - int32(1) {
	case 0:
		goto L6
	case 1:
		goto L7
	case 2:
		v109 = v16
		goto L2
	case 3:
		goto L8
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return v109 & int32(255)
L3:
	;
	v109 = int32(0)
	goto L2
L4:
	;
	v83 = v20
	goto L33
L5:
	;
	v61 = v23
	v62 = int32(1)
	goto L25
L6:
	;
	v28 = int32(0)
	if v14 <= v28 {
		goto L3
	} else {
		goto L11
	}
L7:
	;
	v23 = int32(0)
	if v23 < v14 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v20 = int32(0)
	if v14 <= v20 {
		v109 = v16
		goto L2
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	v109 = int32(1)
	goto L2
L11:
	;
	v31 = v28
	v32 = int32(0)
	goto L12
L12:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v13))))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v109 = v55
	goto L2
L14:
	;
	v43 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v15))))
	if v45 == v43 {
		v109 = v43
		goto L2
	} else {
		goto L17
	}
L15:
	;
	v55 = v32
	goto L16
L16:
	;
	v59 = v31 + int32(1)
	if v59 != v14 {
		v31 = v59
		v32 = v55
		goto L12
	} else {
		goto L24
	}
L17:
	;
	if v32&int32(255) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = v32
	goto L20
L19:
	;
	v51 = int32(2)
	goto L20
L20:
	;
	if v45 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v54 = v51
	goto L23
L22:
	;
	v54 = v32
	goto L23
L23:
	;
	v55 = v54
	goto L16
L24:
	;
	goto L13
L25:
	;
	v69 = int32(0)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v15))))
	if v71 == v69 {
		v109 = v69
		goto L2
	} else {
		goto L27
	}
L26:
	;
	v109 = v79
	goto L2
L27:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v13))))
	if v75 != 0 {
		v109 = v69
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v76 = int32(2)
	if v71 == v76 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v79 = v76
	goto L31
L30:
	;
	v79 = v62
	goto L31
L31:
	;
	v81 = v61 + int32(1)
	if v81 != v14 {
		v61 = v81
		v62 = v79
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v15))))
	if v92 == int32(0) {
		goto L3
	} else {
		goto L35
	}
L34:
	;
	v109 = v16
	goto L2
L35:
	;
	v96 = v83 + int32(1)
	if v14 != v96 {
		v83 = v96
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	return int32(0)
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
	F_errmsg_internal(m, int32(476835), v11)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(494319), int32(300), int32(90807))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ginbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v4 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(5708))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[15]))) = int64(0)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v19 = F_AllocSetContextCreateInternal(m, v14, int32(59463), int32(0), int32(8192), int32(8388608))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
				v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v28 = F_AllocSetContextCreateInternal(m, v23, int32(59703), int32(0), int32(8192), int32(8388608))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[16]))) = v28
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
					F_initGinState(m, v9+int32(4), v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v9
						return v4
					}
				}
			}
		}
	}
}
func F_ginhandler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v53 int32
	_ = v53
	var v63 int64
	_ = v63
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+10)) = v7
		v9 = int32(7)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+8)) = uint16(v9)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(1970324836975030)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(61)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(62)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(63)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(64)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = int32(65)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(66)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+80)) = int32(67)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+76)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(68)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(69)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(70)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(71)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(72)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(73)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(74)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v7
		v53 = int32(5)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+29)) = uint8(v53)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+25)) = int32(65537)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+17)) = int64(281479271678209)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+13)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+128)) = v7
		v63 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v63
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v63
		return v3
	}
}
func F_ginqueryarrayextract(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_copy(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
		F_get_typlenbyvalalign(m, v24, v13+int32(30), v13+int32(29), v13+int32(28))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+30)))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+29)))
			v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+28)))
			F_deconstruct_array(m, v16, v34, v35, v36, v13+int32(24), v13+int32(20), v13+int32(16))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v45
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v47
				switch v21 - int32(1) {
				case 0:
					v72 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
					m.G0 = v13 + int32(32)
					return v74
				case 1:
					v72 = base.B2i32(v45 <= int32(0)) << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
					m.G0 = v13 + int32(32)
					return v74
				case 2:
					v72 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
					m.G0 = v13 + int32(32)
					return v74
				case 3:
					v72 = base.B2i32(v45 <= int32(0))
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
					m.G0 = v13 + int32(32)
					return v74
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v21
						F_errmsg_internal(m, int32(476981), v13)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494319), int32(131), int32(110637))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
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
	}
}
func F_gistchoose(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 float32
	_ = v202
	var v204 float32
	_ = v204
	var v207 float32
	_ = v207
	var v213 float32
	_ = v213
	var v218 float32
	_ = v218
	var v220 float32
	_ = v220
	var v222 float32
	_ = v222
	var v227 int32
	_ = v227
	var v228 float32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	v23 = m.G0
	v25 = v23 - int32(720)
	m.G0 = v25
	F_gistDeCompressAtt(m, l3, l0, l2, v25+int32(48), v25+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+576)) = int32(-1082130432)
	v37 = int32(1)
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
	if base.Ui32(v38) < base.Ui32(int32(25)) {
		v374 = v37
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v25 + int32(720)
	return v374 & int32(65535)
L4:
	;
	v46 = int32(base.Ui32(v38+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v46 == int32(0) {
		v374 = v37
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v52 = l3 + int32(8084)
	v65 = int32(-1)
	v67 = int32(1)
	v69 = v37
	goto L6
L6:
	;
	v82 = v67 & int32(65535)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+10)))
	if v84 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v374 = v324
	goto L3
L8:
	;
	if v271 != base.I32_extend16_s(v269) {
		v323 = v268
		v324 = v272
		goto L52
	} else {
		goto L53
	}
L9:
	;
	v268 = v65
	v269 = v84
	v271 = int32(0)
	v272 = v69
	v273 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v82<<(uint(int32(2))%32)+(l1+int32(24))-int32(4))))
	v106 = v65
	v109 = int32(0)
	v110 = v69
	v111 = int32(1)
	goto L12
L12:
	;
	v123 = v109 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v127 = F_index_getattr_2(m, l1+v94&int32(32767), v123, v124, v25+int32(15))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+10)))
	v268 = v106
	v269 = v260
	v271 = v109
	v272 = v110
	v273 = int32(0)
	goto L8
L14:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+15)))
	if v129 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(16)+v109))))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+716)) = int32(0)
	v182 = l3 + int32(3604) + v109*int32(28)
	if (v177|v129)&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+574)) = uint8(v171)
	goto L15
L17:
	;
	v132 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+574)) = uint8(v132)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+572)) = uint16(v67)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+568)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v25)+564)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v25)+560)) = v127
	v140 = l3 + int32(2708) + v109*int32(28)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v141 == v132 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+572)) = uint16(v67)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+568)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v25)+564)) = l0
	v167 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+560)) = v167
	v171 = v167
	goto L16
L20:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v52+v109<<(uint(int32(2))%32))))
	v150 = F_FunctionCall1Coll(m, v140, v147, v25+int32(560))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v150 == v25+int32(560) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+560)) = v155
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+564)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+568)) = v159
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+572)) = uint16(v161)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+14)))
	v171 = v163
	goto L16
L23:
	;
	v227 = v25 + int32(576) + v109<<(uint(int32(2))%32)
	v228 = *(*float32)(unsafe.Add(mBase, uint32(v227)))
	if base.F32_lt(v228, float32(0))|base.F32_lt(v222, v228) != 0 {
		goto L44
	} else {
		goto L45
	}
L24:
	;
	if v177&int32(1) != 0 {
		goto L36
	} else {
		goto L37
	}
L25:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+10)))
	if v186 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v52+v109<<(uint(int32(2))%32))))
	v200 = F_FunctionCall3Coll(m, v182, v190, v25+int32(560), v25+int32(48)+v109<<(uint(int32(4))%32), v25+int32(716))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v202 = float32(0)
	v204 = *(*float32)(unsafe.Add(mBase, uint32(v25)+716))
	if base.F32_lt(v204, v202) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v207 = v202
	goto L32
L31:
	;
	v207 = v204
	goto L32
L32:
	;
	if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v204)&int32(2147483647)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v213 = v202
	goto L35
L34:
	;
	v213 = v207
	goto L35
L35:
	;
	v222 = v213
	goto L23
L36:
	;
	v218 = float32(0)
	goto L38
L37:
	;
	v218 = math.Float32frombits(uint32(0x7f800000))
	goto L38
L38:
	;
	if v129 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v220 = v218
	goto L41
L40:
	;
	v220 = math.Float32frombits(uint32(0x7f800000))
	goto L41
L41:
	;
	v222 = v220
	goto L23
L42:
	;
	goto L13
L43:
	;
	v257 = base.B2i32(base.F32_gt(v222, float32(0)) == int32(0)) & v111
	if v123 < v251 {
		v106 = v250
		v109 = v123
		v110 = v252
		v111 = v257
		goto L12
	} else {
		goto L51
	}
L44:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v227))) = v222
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235)+10)))
	if v109 < v236-int32(1) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	if base.F32_ne(v222, v228) != 0 {
		goto L42
	} else {
		goto L50
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(576)+v123<<(uint(int32(2))%32)))) = int32(-1082130432)
	goto L49
L48:
	;
	goto L49
L49:
	;
	v250 = int32(-1)
	v251 = v236
	v252 = v67
	goto L43
L50:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+10)))
	v250 = v106
	v251 = v249
	v252 = v110
	goto L43
L51:
	;
	v268 = v250
	v269 = v251
	v271 = v123
	v272 = v252
	v273 = v257
	goto L8
L52:
	;
	if v273 != 0 {
		goto L65
	} else {
		goto L66
	}
L53:
	;
	if v272&int32(65535) == v82 {
		v323 = v268
		v324 = v272
		goto L52
	} else {
		goto L54
	}
L54:
	;
	if v268 == int32(-1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v293 = int32(4559656)
	v294 = int32(4559648)
	v295 = *(*int64)(unsafe.Add(mBase, _consts[23]))
	v297 = *(*int64)(unsafe.Add(mBase, _consts[24]))
	v298 = v295 ^ v297
	*(*int64)(unsafe.Add(mBase, _consts[24])) = base.I64_rotl(v298, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[23])) = v298<<(uint(int64(16))%64) ^ base.I64_rotl(v295, int64(24)) ^ v298
	goto L58
L56:
	;
	v319 = v268
	goto L57
L57:
	;
	if v319 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v319 = base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v295*int64(5), int64(7))*int64(9)) >> (uint(int64(63)) % 64)))
	goto L57
L59:
	;
	v320 = v272
	goto L61
L60:
	;
	v320 = v67
	goto L61
L61:
	;
	if v319 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v322 = v319
	goto L64
L63:
	;
	v322 = int32(-1)
	goto L64
L64:
	;
	v323 = v322
	v324 = v320
	goto L52
L65:
	;
	if v323 == int32(-1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v358 = v323
	goto L67
L67:
	;
	v360 = v67 + int32(1)
	if base.Ui32(v360&int32(65535)) <= base.Ui32(v46) {
		v65 = v358
		v67 = v360
		v69 = v324
		goto L6
	} else {
		goto L73
	}
L68:
	;
	v329 = int32(4559656)
	v330 = int32(4559648)
	v331 = *(*int64)(unsafe.Add(mBase, _consts[23]))
	v333 = *(*int64)(unsafe.Add(mBase, _consts[24]))
	v334 = v331 ^ v333
	*(*int64)(unsafe.Add(mBase, _consts[24])) = base.I64_rotl(v334, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[23])) = v334<<(uint(int64(16))%64) ^ base.I64_rotl(v331, int64(24)) ^ v334
	goto L71
L69:
	;
	v355 = v323
	goto L70
L70:
	;
	if v355 == int32(1) {
		v374 = v324
		goto L3
	} else {
		goto L72
	}
L71:
	;
	v355 = base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v331*int64(5), int64(7))*int64(9)) >> (uint(int64(63)) % 64)))
	goto L70
L72:
	;
	v358 = v355
	goto L67
L73:
	;
	goto L7
}
func F_gistcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v86 float64
	_ = v86
	var v92 float64
	_ = v92
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v19
	F_genericcostestimate(m, l0, l1, l2, v16)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
		if v37 < int32(0) {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			if base.Ui32(v41) < base.Ui32(int32(2)) {
				v54 = int32(0)
			} else {
				v45 = F_log(m, base.F64_convert_i32_u(v41))
				mBase = m.M
				v47 = base.F64_div(v45, float64(4.605170185988092))
				if base.F64_lt(base.F64_abs(v47), float64(2.147483648e+09)) != 0 {
					v51 = base.I32_trunc_f64_s(v47)
					v54 = v51
				} else {
					v54 = int32(-2147483648)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v54
			v57 = v54
		} else {
			v57 = v37
		}
		v60 = *(*float64)(unsafe.Add(mBase, _consts[384]))
		v61 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
		v62 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
		if base.F64_gt(v62, float64(1)) == int32(0) {
			v67 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
			v68 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
			v77 = v61
			v79 = v67
			v80 = v68
		} else {
			v69 = F_log(m, v62)
			mBase = m.M
			v71 = base.F64_mul(base.F64_ceil(v69), v60)
			v73 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
			v75 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
			v77 = base.F64_add(v61, v71)
			v79 = v73
			v80 = base.F64_add(base.F64_mul(v73, v71), v75)
		}
		v86 = base.F64_mul(v60, base.F64_mul(base.F64_convert_i32_s(v57+int32(1)), float64(50)))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v77, v86)
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(v79, v86), v80)
		v92 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v92
		v94 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v94
		v96 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v96
		m.G0 = v16 - int32(-64)
		return
	}
}
func F_gistfillbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l3 != 0 {
		v22 = l3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < l2 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v13) < base.Ui32(int32(25)) {
		v22 = int32(1)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(base.Ui32(v13+int32(262120))>>(uint(int32(2))%32)) + int32(1)
	goto L1
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L14
	}
L5:
	;
	v30 = v22
	v31 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v10 + int32(16)
	return
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v31<<(uint(int32(2))%32))))
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+6)))
	v40 = v38 & int32(8191)
	v44 = F_PageAddItemExtended(m, l0, v37, v40, v30&int32(65535), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	return
L11:
	;
	if v44 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v48 = int32(1)
	v51 = v31 + v48
	if v51 != l2 {
		v30 = v30 + v48
		v31 = v51
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
	F_errmsg_internal(m, int32(158197), v10)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(491946), int32(50), int32(223387))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gistgetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int64
	_ = v96
	var v103 int64
	_ = v103
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v7
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	if v15 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v103 = v7
	goto L3
L3:
	;
	m.G0 = v10 + int32(48)
	return v103
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)))
	if v22 != int32(1) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v31 = v19
	goto L7
L7:
	;
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v32 + int64(1)
	goto L4
L8:
	;
	F_pgstat_assoc_relation(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int64(0)
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+272))
	v31 = v30
	goto L7
L11:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	*(*int64)(unsafe.Add(mBase, uint32(v37))) = v38 + int64(1)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[20]))) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v42
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[21])))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_MemoryContextReset(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v51
	F_gistScanPage(m, l0, v10+int32(8), v51, l1, v10+int32(40))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
	v103 = v96
	goto L3
L20:
	;
	v66 = v60
	goto L21
L21:
	;
	v71 = F_pairingheap_remove_first(m, v66)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L23
	}
L22:
	;
	goto L19
L23:
	;
	if v71 == int32(0) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v76 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_gistScanPage(m, l0, v71, v71+int32(32), l1, v10+int32(40))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	F_pfree(m, v71)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	if v88 != 0 {
		v66 = v87
		goto L21
	} else {
		goto L31
	}
L31:
	;
	goto L22
}
func F_gistinserttuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l3
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 < v6 {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(v14^int32(-1))<<(uint(int32(6))%32))+16))
		v33 = v24
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _consts[10]))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+v14<<(uint(int32(6))%32)+int32(-64))+16))
		v33 = v32
	}
	F_CheckForSerializableConflictIn(m, v12, v6, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v43 = int32(1)
		v44 = int32(0)
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		v51 = F_gistplacetopage(m, v38, v39, l2, v40, v9+int32(8), v43, l4, v44, v44, v9+int32(12), v43, v49, v50)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if v53 != 0 {
				F_gistfinishsplit(m, l0, l1, l2, v53, int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v51
				}
			} else {
				m.G0 = v9 + int32(16)
				return v51
			}
		}
	}
}
func F_gistjoinvector(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = F_repalloc(m, l0, (v5+l3)<<(uint(int32(2))%32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v14 = int32(2)
	v16 = v9 + v13<<(uint(v14)%32)
	v18 = l3 << (uint(v14) % 32)
	if v16 == l2 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v163 + l3
	return v9
L4:
	;
	goto L3
L5:
	;
	v22 = v16 + v18
	if base.Ui32(l2-v22) <= base.Ui32(int32(0)-v18<<(uint(int32(1))%32)) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = F___memcpy(m, v16, l2, v18)
	mBase = m.M
	goto L3
L7:
	;
	goto L8
L8:
	;
	v32 = (v16 ^ l2) & int32(3)
	if base.Ui32(v16) < base.Ui32(l2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v134 == int32(0) {
		goto L4
	} else {
		goto L45
	}
L10:
	;
	if base.Ui32(v112) <= base.Ui32(int32(3)) {
		v133 = v111
		v134 = v112
		v135 = v113
		goto L9
	} else {
		goto L41
	}
L11:
	;
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v32 != 0 {
		v94 = v18
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v133 = l2
	v134 = v18
	v135 = v16
	goto L9
L15:
	;
	goto L16
L16:
	;
	if v16&int32(3) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v111 = l2
	v112 = v18
	v113 = v16
	goto L10
L18:
	;
	goto L19
L19:
	;
	v39 = l2
	v40 = v18
	v41 = v16
	goto L20
L20:
	;
	if v40 == int32(0) {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	v111 = v48
	v112 = v50
	v113 = v52
	goto L10
L22:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v45)
	v47 = int32(1)
	v48 = v39 + v47
	v50 = v40 - v47
	v52 = v41 + v47
	if v52&int32(3) != 0 {
		v39 = v48
		v40 = v50
		v41 = v52
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v94 == int32(0) {
		goto L4
	} else {
		goto L37
	}
L25:
	;
	if v22&int32(3) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v59 = v18
	goto L29
L27:
	;
	v74 = v18
	goto L28
L28:
	;
	if base.Ui32(v74) <= base.Ui32(int32(3)) {
		v94 = v74
		goto L24
	} else {
		goto L33
	}
L29:
	;
	if v59 == int32(0) {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	v74 = v65
	goto L28
L31:
	;
	v65 = v59 - int32(1)
	v66 = v16 + v65
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v65))))
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v68)
	if v66&int32(3) != 0 {
		v59 = v65
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v81 = v74
	goto L34
L34:
	;
	v85 = v81 - int32(4)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2+v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v16+v85))) = v88
	if base.Ui32(int32(3)) < base.Ui32(v85) {
		v81 = v85
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v94 = v85
	goto L24
L36:
	;
	goto L35
L37:
	;
	v101 = v94
	goto L38
L38:
	;
	v105 = v101 - int32(1)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v105))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16+v105))) = uint8(v108)
	if v105 != 0 {
		v101 = v105
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L4
L40:
	;
	goto L39
L41:
	;
	v118 = v111
	v119 = v112
	v120 = v113
	goto L42
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v122
	v124 = int32(4)
	v125 = v118 + v124
	v127 = v120 + v124
	v129 = v119 - v124
	if base.Ui32(int32(3)) < base.Ui32(v129) {
		v118 = v125
		v119 = v129
		v120 = v127
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v133 = v125
	v134 = v129
	v135 = v127
	goto L9
L44:
	;
	goto L43
L45:
	;
	v140 = v133
	v141 = v134
	v142 = v135
	goto L46
L46:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v144)
	v146 = int32(1)
	v151 = v141 - v146
	if v151 != 0 {
		v140 = v140 + v146
		v141 = v151
		v142 = v142 + v146
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L4
L48:
	;
	goto L47
}
func F_gtsquery_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v131 int32
	_ = v131
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int64
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v269 int32
	_ = v269
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v297 int32
	_ = v297
	var v320 int32
	_ = v320
	var v322 int64
	_ = v322
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v331 int64
	_ = v331
	var v337 int32
	_ = v337
	var v353 int32
	_ = v353
	var v376 int32
	_ = v376
	var v378 int64
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v454 int64
	_ = v454
	var v455 int64
	_ = v455
	var v458 int32
	_ = v458
	var v459 int64
	_ = v459
	var v481 int32
	_ = v481
	var v504 int32
	_ = v504
	var v506 int64
	_ = v506
	var v509 int64
	_ = v509
	var v513 int64
	_ = v513
	var v519 int32
	_ = v519
	var v535 int32
	_ = v535
	var v558 int32
	_ = v558
	var v560 int64
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v595 int64
	_ = v595
	var v596 int64
	_ = v596
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v613 int32
	_ = v613
	var v619 int64
	_ = v619
	var v620 int64
	_ = v620
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	v7 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v28 = (v24 + int32(65534)) & int32(65535)
	v32 = v28<<(uint(int32(1))%32) + int32(4)
	v33 = F_palloc(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v33
		v38 = F_palloc(m, v32)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v40
			if base.Ui32(int32(2)) <= base.Ui32(v28) {
				v48 = v23 + int32(4)
				v59 = v7
				v60 = int32(-1)
				v61 = v7
				v63 = int32(1)
				for {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v48+v63<<(uint(int32(4))%32))))
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
					v78 = v63 + int32(1)
					v79 = v78
					v86 = v78
					v87 = v59
					v88 = v60
					v89 = v61
					for {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v48+v79<<(uint(int32(4))%32))))
						v104 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
						v105 = v76 ^ v104
						v108 = int32(0)
						v109 = int64(0)
						for {
							v131 = int32(1)
							v154 = base.I32_wrap_i64(int64(base.Ui64(v105)>>(uint(v109)%64)))&v131 + v108 + base.I32_wrap_i64(int64(base.Ui64(v105)>>(uint(v109|int64(1))%64)))&v131 + base.I32_wrap_i64(int64(base.Ui64(v105)>>(uint(v109|int64(2))%64)))&v131 + base.I32_wrap_i64(int64(base.Ui64(v105)>>(uint(v109|int64(3))%64)))&v131
							v156 = v109 + int64(4)
							if v156 != int64(64) {
								v108 = v154
								v109 = v156
								continue
							} else {
								break
							}
							break
						}
						v159 = base.B2i32(v88 < v154)
						if v88 < v154 {
							v160 = v86
						} else {
							v160 = v89
						}
						if v88 < v154 {
							v161 = v63
						} else {
							v161 = v87
						}
						if v88 < v154 {
							v162 = v154
						} else {
							v162 = v88
						}
						v164 = v86 + int32(1)
						v166 = v164 & int32(65535)
						if base.Ui32(v166) <= base.Ui32(v28) {
							v79 = v166
							v86 = v164
							v87 = v161
							v88 = v162
							v89 = v160
							continue
						} else {
							break
						}
						break
					}
					if v78 != v28 {
						v59 = v161
						v60 = v162
						v61 = v160
						v63 = v78
						continue
					} else {
						break
					}
					break
				}
				v177 = v161
				v179 = v160
			} else {
				v177 = v7
				v179 = v7
			}
			v191 = v23 + int32(4)
			v193 = int32(65535)
			v195 = int32(0)
			v201 = base.B2i32(v177&v193 == v195) | base.B2i32(v179&v193 == v195)
			if v201 != 0 {
				v202 = int32(2)
			} else {
				v202 = v179
			}
			v207 = v191 + v202&int32(65535)<<(uint(int32(4))%32)
			v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
			v209 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
			if v201 != 0 {
				v211 = int32(1)
			} else {
				v211 = v177
			}
			v212 = int32(65535)
			v216 = v191 + v211&v212<<(uint(int32(4))%32)
			v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
			v218 = *(*int64)(unsafe.Add(mBase, uint32(v217)))
			v220 = v24 + v212
			v222 = v220 & v212
			v225 = F_palloc(m, v222<<(uint(int32(3))%32))
			mBase = m.M
			v226 = m.ExcPending
			if v226 != 0 {
				return int32(0)
			} else {
				if v24&int32(65535) == int32(1) {
					F_pg_qsort(m, v225, v222, int32(8), int32(1527))
					mBase = m.M
					v234 = m.ExcPending
					if v234 != 0 {
						return int32(0)
					} else {
						v619 = v209
						v620 = v218
						v626 = v38
						v629 = v33
						v636 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v629))) = uint16(v636)
						*(*uint16)(unsafe.Add(mBase, uint32(v626))) = uint16(v636)
						v640 = F_Int64GetDatum(m, v620)
						mBase = m.M
						v641 = m.ExcPending
						if v641 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v640
							v643 = F_Int64GetDatum(m, v619)
							mBase = m.M
							v644 = m.ExcPending
							if v644 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v643
								return v22
							}
						}
					}
				} else {
					v235 = int32(1)
					v237 = v235
					v246 = v235
					for {
						v260 = v225 + v237<<(uint(int32(3))%32)
						*(*uint16)(unsafe.Add(mBase, uint32(v260-int32(8)))) = uint16(v246)
						v267 = *(*int32)(unsafe.Add(mBase, uint32(v191+v237<<(uint(int32(4))%32))))
						v268 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
						v269 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
						v270 = *(*int64)(unsafe.Add(mBase, uint32(v269)))
						v271 = v268 ^ v270
						v274 = int32(0)
						v275 = int64(0)
						for {
							v297 = int32(1)
							v320 = base.I32_wrap_i64(int64(base.Ui64(v271)>>(uint(v275)%64)))&v297 + v274 + base.I32_wrap_i64(int64(base.Ui64(v271)>>(uint(v275|int64(1))%64)))&v297 + base.I32_wrap_i64(int64(base.Ui64(v271)>>(uint(v275|int64(2))%64)))&v297 + base.I32_wrap_i64(int64(base.Ui64(v271)>>(uint(v275|int64(3))%64)))&v297
							v322 = v275 + int64(4)
							if v322 != int64(64) {
								v274 = v320
								v275 = v322
								continue
							} else {
								break
							}
							break
						}
						v325 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
						v326 = *(*int64)(unsafe.Add(mBase, uint32(v325)))
						v327 = v326 ^ v268
						v331 = int64(0)
						v337 = int32(0)
						for {
							v353 = int32(1)
							v376 = base.I32_wrap_i64(int64(base.Ui64(v327)>>(uint(v331)%64)))&v353 + v337 + base.I32_wrap_i64(int64(base.Ui64(v327)>>(uint(v331|int64(1))%64)))&v353 + base.I32_wrap_i64(int64(base.Ui64(v327)>>(uint(v331|int64(2))%64)))&v353 + base.I32_wrap_i64(int64(base.Ui64(v327)>>(uint(v331|int64(3))%64)))&v353
							v378 = v331 + int64(4)
							if v378 != int64(64) {
								v331 = v378
								v337 = v376
								continue
							} else {
								break
							}
							break
						}
						v383 = v320 - v376
						v385 = v383 >> (uint(int32(31)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(v260-int32(4)))) = v383 ^ v385 - v385
						v390 = v246 + int32(1)
						v391 = int32(65535)
						v392 = v390 & v391
						if base.Ui32(v392) <= base.Ui32(v220&v391) {
							v237 = v392
							v246 = v390
							continue
						} else {
							break
						}
						break
					}
					F_pg_qsort(m, v225, v222, int32(8), int32(1527))
					mBase = m.M
					v399 = m.ExcPending
					if v399 != 0 {
						return int32(0)
					} else {
						v400 = int32(1)
						if base.Ui32(v222) <= base.Ui32(v400) {
							v403 = v400
						} else {
							v403 = v222
						}
						v411 = v209
						v412 = v218
						v416 = int32(0)
						v418 = v38
						v421 = v33
						for {
							v431 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225+v416<<(uint(int32(3))%32)))))
							if v211&int32(65535) == v431 {
								*(*uint16)(unsafe.Add(mBase, uint32(v421))) = uint16(v211)
								v434 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v434 + int32(1)
								v595 = v411
								v596 = v412
								v602 = v418
								v605 = v421 + int32(2)
							} else {
								if v202&int32(65535) == v431 {
									*(*uint16)(unsafe.Add(mBase, uint32(v418))) = uint16(v202)
									v444 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v444 + int32(1)
									v595 = v411
									v596 = v412
									v602 = v418 + int32(2)
									v605 = v421
								} else {
									v453 = *(*int32)(unsafe.Add(mBase, uint32(v191+v431<<(uint(int32(4))%32))))
									v454 = *(*int64)(unsafe.Add(mBase, uint32(v453)))
									v455 = v454 ^ v412
									v458 = int32(0)
									v459 = int64(0)
									for {
										v481 = int32(1)
										v504 = base.I32_wrap_i64(int64(base.Ui64(v455)>>(uint(v459)%64)))&v481 + v458 + base.I32_wrap_i64(int64(base.Ui64(v455)>>(uint(v459|int64(1))%64)))&v481 + base.I32_wrap_i64(int64(base.Ui64(v455)>>(uint(v459|int64(2))%64)))&v481 + base.I32_wrap_i64(int64(base.Ui64(v455)>>(uint(v459|int64(3))%64)))&v481
										v506 = v459 + int64(4)
										if v506 != int64(64) {
											v458 = v504
											v459 = v506
											continue
										} else {
											break
										}
										break
									}
									v509 = v454 ^ v411
									v513 = int64(0)
									v519 = int32(0)
									for {
										v535 = int32(1)
										v558 = base.I32_wrap_i64(int64(base.Ui64(v509)>>(uint(v513)%64)))&v535 + v519 + base.I32_wrap_i64(int64(base.Ui64(v509)>>(uint(v513|int64(1))%64)))&v535 + base.I32_wrap_i64(int64(base.Ui64(v509)>>(uint(v513|int64(2))%64)))&v535 + base.I32_wrap_i64(int64(base.Ui64(v509)>>(uint(v513|int64(3))%64)))&v535
										v560 = v513 + int64(4)
										if v560 != int64(64) {
											v513 = v560
											v519 = v558
											continue
										} else {
											break
										}
										break
									}
									v565 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
									v566 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
									v567 = v565 - v566
									if base.F64_lt(base.F64_convert_i32_s(v504), base.F64_add(base.F64_convert_i32_s(v558), base.F64_mul(base.F64_convert_i32_s(v567*v567*v567), float64(-0.05)))) != 0 {
										*(*uint16)(unsafe.Add(mBase, uint32(v421))) = uint16(v431)
										v576 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v576 + int32(1)
										v595 = v411
										v596 = v454 | v412
										v602 = v418
										v605 = v421 + int32(2)
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v418))) = uint16(v431)
										v584 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v584 + int32(1)
										v595 = v454 | v411
										v596 = v412
										v602 = v418 + int32(2)
										v605 = v421
									}
								}
							}
							v613 = v416 + int32(1)
							if v613 != v403 {
								v411 = v595
								v412 = v596
								v416 = v613
								v418 = v602
								v421 = v605
								continue
							} else {
								break
							}
							break
						}
						v619 = v595
						v620 = v596
						v626 = v602
						v629 = v605
						v636 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v629))) = uint16(v636)
						*(*uint16)(unsafe.Add(mBase, uint32(v626))) = uint16(v636)
						v640 = F_Int64GetDatum(m, v620)
						mBase = m.M
						v641 = m.ExcPending
						if v641 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v640
							v643 = F_Int64GetDatum(m, v619)
							mBase = m.M
							v644 = m.ExcPending
							if v644 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v643
								return v22
							}
						}
					}
				}
			}
		}
	}
}
