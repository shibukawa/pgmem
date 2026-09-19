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
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetActiveSnapshot[0]))
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
				F_errmsg_internal(m, int32(_a_F_GetFdwRoutineByRelId_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetFdwRoutineByRelId_1), int32(364), int32(_a_F_GetFdwRoutineByRelId_2))
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
	var v18 int32
	_ = v18
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
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v4 = int32(0)
	if l2 <= v4 {
		v51 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	if v57 == int32(0) {
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
		v51 = v25
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
	v18 = v4
	goto L7
L7:
	;
	v19 = int32(1)
	v20 = v16 - v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v23 = v18 + v19
	if v23 != v10 {
		v13 = v21
		v16 = v20
		v18 = v23
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
	v39 = int32(8)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v39 < v36 {
		v33 = v48
		v36 = v36 - v39
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v51 = v48
	goto L1
L13:
	;
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v63 = int32(0)
	if v63 < v60 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v67 = v60
	goto L19
L18:
	;
	v67 = v63
	goto L19
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v72 = v63
	goto L20
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v68+v72<<(uint(int32(2))%32))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if l1 != v79 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	return v78
L22:
	;
	v82 = v72 + int32(1)
	if v67 != v82 {
		v72 = v82
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
	F_errmsg_internal(m, int32(_a_F_GetNSItemByRangeTablePosn_0), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_GetNSItemByRangeTablePosn_1), int32(536), int32(_a_F_GetNSItemByRangeTablePosn_2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestRestartPoint[0]))
	v8 = F_LWLockAcquire(m, v4+int32(1152), int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestRestartPoint[1]))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestRestartPoint[0]))
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
	F_appendStringInfoString(m, l1, int32(_a_F_GetPublicationsStr_0))
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
func F_GetRecordedFreeSpace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = base.I32_div_u_s(l1, int32(4069))
	v15 = base.I64_extend_i32_u(v12) << (uint(int64(32)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v15
	v19 = F_fsm_readbuf(m, l0, v9, v3)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != 0 {
			if v19 < int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetRecordedFreeSpace[0]))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v19^int32(-1))<<(uint(int32(2))%32))))
				v43 = v35
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_GetRecordedFreeSpace[1]))
				v43 = v37 + v19<<(uint(int32(13))%32) + int32(-8192)
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+(l1-v12*int32(4069)))+uint32(_c_F_GetRecordedFreeSpace[2]))))
			F_ReleaseBuffer(m, v19)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				v53 = v47 << (uint(int32(5)) % 32)
				m.G0 = v9 + int32(16)
				return v53
			}
		} else {
			v53 = v3
			m.G0 = v9 + int32(16)
			return v53
		}
	}
}
func F_GetTopMostAncestorInPublication(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
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
	var v118 int32
	_ = v118
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	v3 = int32(0)
	if l1 == v3 {
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v3
	v26 = v3
	goto L7
L5:
	;
	v242 = v3
	goto L6
L6:
	;
	return v242
L7:
	;
	v28 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v24<<(uint(int32(2))%32))))
	v38 = F_SearchSysCacheList(m, int32(53), int32(1), v35, v28, v28)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v242 = v227
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if int32(0) < v42 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = int32(0)
	v52 = v28
	goto L14
L12:
	;
	v76 = v28
	goto L13
L13:
	;
	F_ReleaseCatCacheList(m, v38)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L9
	} else {
		goto L18
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(48)+v51<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)+4))
	v66 = F_lappend_oid(m, v52, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v76 = v66
	goto L13
L16:
	;
	v69 = v51 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if v69 < v70 {
		v51 = v69
		v52 = v66
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v85 = v24 + int32(1)
	v86 = int32(0)
	if v76 == v86 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	F_list_free(m, v76)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L63
	}
L20:
	;
	v219 = int32(0)
	v227 = v35
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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v92 <= int32(0) {
		v118 = v86
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v124 = v118
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
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v101 = int32(0)
	goto L30
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v99+v101<<(uint(int32(2))%32))))
	v110 = base.B2i32(v109 == l0)
	if v109 == l0 {
		v118 = v110
		goto L25
	} else {
		goto L32
	}
L31:
	;
	v118 = v110
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
	v129 = F_get_rel_namespace(m, v35)
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
	v143 = v126
	v144 = int32(0)
	goto L42
L40:
	;
	v167 = v126
	goto L41
L41:
	;
	F_ReleaseCatCacheList(m, v133)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L46
	}
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v133+int32(48)+v144<<(uint(int32(2))%32))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+56))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+22)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155+v156)+4))
	v159 = F_lappend_oid(m, v143, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L44
	}
L43:
	;
	v167 = v159
	goto L41
L44:
	;
	v162 = v144 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	if v162 < v163 {
		v143 = v159
		v144 = v162
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v177 = int32(0)
	if v167 == v177 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v215 != 0 {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v215 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v183 <= int32(0) {
		v209 = v177
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v215 = v209
	goto L47
L52:
	;
	v186 = int32(0)
	if v186 < v183 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v189 = v183
	goto L55
L54:
	;
	v189 = v186
	goto L55
L55:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v192 = int32(0)
	goto L56
L56:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v190+v192<<(uint(int32(2))%32))))
	v201 = base.B2i32(v200 == l0)
	if v200 == l0 {
		v209 = v201
		goto L51
	} else {
		goto L58
	}
L57:
	;
	v209 = v201
	goto L51
L58:
	;
	v203 = v192 + int32(1)
	if v203 != v189 {
		v192 = v203
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v216 = v35
	goto L62
L61:
	;
	v216 = v26
	goto L62
L62:
	;
	v219 = v167
	v227 = v216
	goto L19
L63:
	;
	F_list_free(m, v219)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v85 < v232 {
		v24 = v85
		v26 = v227
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
	v3 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_GetTopTransactionId[0])))
	if v3 == int64(0) {
		F_AssignTransactionId(m, int32(_a_F_GetTopTransactionId_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, _c_F_GetTopTransactionId[0]))
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_GetVirtualXIDsDelayingChkpt[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v17 = F_palloc(m, v14<<(uint(int32(3))%32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetVirtualXIDsDelayingChkpt[1]))
		v26 = F_LWLockAcquire(m, v22+int32(512), int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if int32(0) < v28 {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetVirtualXIDsDelayingChkpt[2]))
				v37 = v28
				v40 = v3
				v42 = v3
				for {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)+v42<<(uint(int32(2))%32))))
					v52 = v34 + v49*int32(640)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+120))
					if v53&l1 == int32(0) {
						v69 = v37
						v71 = v40
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
						if v57 == int32(0) {
							v69 = v37
							v71 = v40
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
							v63 = v17 + v40<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v63))) = v60
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v69 = v68
							v71 = v40 + int32(1)
						}
					}
					v74 = v42 + int32(1)
					if v74 < v69 {
						v37 = v69
						v40 = v71
						v42 = v74
						continue
					} else {
						break
					}
					break
				}
				v81 = v71
			} else {
				v81 = v3
			}
			v88 = *(*int32)(unsafe.Add(mBase, _c_F_GetVirtualXIDsDelayingChkpt[1]))
			F_LWLockRelease(m, v88+int32(512))
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v81
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
			v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[0])))
			if v10 == int32(1) {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[1]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+316))
				v18 = base.B2i32(v16 != int32(2))
				*(*uint8)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[0])) = uint8(v18)
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
				if base.Ui32(v22) < base.Ui32(int32(_a_F_GlobalVisHorizonKindForRel_0)) {
					v50 = v21
					return v50
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[2]))
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
							if base.Ui32(v33) < base.Ui32(int32(_a_F_GlobalVisHorizonKindForRel_0)) {
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
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
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L54
	}
L13:
	;
	v195 = F_palloc(m, int32(16))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L53
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
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	if v145&int32(4) != 0 {
		goto L41
	} else {
		goto L42
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
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v46 << (uint(int32(2)) % 32)
	v55 = v47 + int32(8)
	if v38 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	base.MemoryFill(m, v55, int32(0), v38)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v58 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v59 = F_array_contains_nulls(m, v43)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v63 = v43 + int32(16)
	v64 = F_ArrayGetNItemsSafe(m, v61, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L27
	}
L25:
	;
	if v59 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	if v64 == int32(0) {
		v190 = v47
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v70 = F_ArrayGetNItemsSafe(m, v69, v63)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v70 == int32(0) {
		v190 = v47
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v74 = int32(3)
	v75 = v38 << (uint(v74) % 32)
	if v68 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v82 = v68
	goto L33
L32:
	;
	v82 = (v69<<(uint(v74)%32) + int32(23)) & int32(-8)
	goto L33
L33:
	;
	v83 = v43 + v82
	if v70&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v87 = base.I32_rem_u_s(v86, v75)
	v90 = v55 + int32(base.Ui32(v87)>>(uint(int32(3))%32))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v92 = int32(1)
	v96 = v91 | v92<<(uint(v87&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v96)
	v102 = v83 + int32(4)
	v105 = v70 - v92
	goto L36
L35:
	;
	v102 = v83
	v105 = v70
	goto L36
L36:
	;
	if v70 == int32(1) {
		v190 = v47
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v108 = v102
	v110 = v105
	goto L38
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v118 = base.I32_rem_u_s(v117, v75)
	v119 = int32(3)
	v121 = v55 + int32(base.Ui32(v118)>>(uint(v119)%32))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v123 = int32(1)
	v124 = int32(7)
	v127 = v122 | v123<<(uint(v118&v124)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v127)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v130 = base.I32_rem_u_s(v129, v75)
	v133 = v55 + int32(base.Ui32(v130)>>(uint(v119)%32))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v139 = v134 | v123<<(uint(v130&v124)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v139)
	v144 = v110 - int32(2)
	if v144 != 0 {
		v108 = v108 + int32(8)
		v110 = v144
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v190 = v47
	goto L13
L40:
	;
	goto L39
L41:
	;
	return v10
L42:
	;
	goto L43
L43:
	;
	v149 = int32(0)
	if v38 <= v149 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v181 = F_palloc(m, int32(8))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L52
	}
L45:
	;
	v154 = v149
	goto L46
L46:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+(v39+int32(8))))))
	if v164 == int32(255) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	return v10
L48:
	;
	v168 = v154 + int32(1)
	if v38 != v168 {
		v154 = v168
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L44
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v181))) = int64(17179869216)
	v190 = v181
	goto L13
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v190
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+4)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v200
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)))
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+14)) = uint8(v203)
	*(*uint16)(unsafe.Add(mBase, uint32(v195)+12)) = uint16(v202)
	return v195
L54:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_g_intbig_compress_0), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_g_intbig_compress_1), int32(159), int32(_a_F_g_intbig_compress_2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
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
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
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
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
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
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
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
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
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
	var v379 int32
	_ = v379
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
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
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
		v426 = v45
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
	v456 = m.ExcPending
	if v456 != 0 {
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
	return v426 & int32(1)
L16:
	;
	if v18&int32(_a_F_g_intbig_consistent_0) == int32(20) {
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
		v414 = v59
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v426 = v59
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
	switch v18&int32(_a_F_g_intbig_consistent_0) - int32(3) {
	case 0:
		goto L32
	default:
		v402 = v66
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
		v426 = v402
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
		v402 = v267
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
	v75 = F_ArrayGetNItemsSafe(m, v72, v13+int32(16))
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
		v402 = v66
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
	v98 = v75
	v102 = v89 + v13
	goto L41
L41:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v108 = base.I32_rem_u_s(v107, v44<<(uint(int32(3))%32))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+int32(8)+int32(base.Ui32(v108)>>(uint(int32(3))%32))))))
	v115 = int32(base.Ui32(v112) >> (uint(v108&int32(7)) % 32))
	if v115&int32(1) != 0 {
		v402 = v115
		goto L28
	} else {
		goto L43
	}
L42:
	;
	v402 = v115
	goto L28
L43:
	;
	v121 = v98 - int32(1)
	if v121 != 0 {
		v98 = v121
		v102 = v102 + int32(4)
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
	v131 = F_ArrayGetNItemsSafe(m, v128, v13+int32(16))
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
		v251 = int32(1)
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
	v177 = v169
	v178 = v172
	goto L59
L59:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v186 = base.I32_rem_u_s(v185, v149)
	v187 = int32(3)
	v189 = v144 + int32(base.Ui32(v186)>>(uint(v187)%32))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v191 = int32(1)
	v192 = int32(7)
	v195 = v190 | v191<<(uint(v186&v192)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	v198 = base.I32_rem_u_s(v197, v149)
	v201 = v144 + int32(base.Ui32(v198)>>(uint(v187)%32))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v207 = v202 | v191<<(uint(v198&v192)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v207)
	v212 = v178 - int32(2)
	if v212 != 0 {
		v177 = v177 + int32(8)
		v178 = v212
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
	v231 = int32(0)
	goto L64
L64:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+(v226+int32(8))))))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v144))))
	v244 = base.B2i32(v241 == v243)
	if v241 != v243 {
		v251 = v244
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v251 = v244
	goto L62
L66:
	;
	v247 = v231 + int32(1)
	if v247 != v44 {
		v231 = v247
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v402 = v251
	goto L28
L69:
	;
	v402 = v262
	goto L28
L70:
	;
	v402 = v265
	goto L28
L71:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v279 = F_ArrayGetNItemsSafe(m, v276, v13+int32(16))
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
		v402 = int32(1)
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
	v325 = v317
	v326 = v320
	goto L83
L83:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v334 = base.I32_rem_u_s(v333, v297)
	v335 = int32(3)
	v337 = v292 + int32(base.Ui32(v334)>>(uint(v335)%32))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v339 = int32(1)
	v340 = int32(7)
	v343 = v338 | v339<<(uint(v334&v340)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v337))) = uint8(v343)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v346 = base.I32_rem_u_s(v345, v297)
	v349 = v292 + int32(base.Ui32(v346)>>(uint(v335)%32))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	v355 = v350 | v339<<(uint(v346&v340)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v349))) = uint8(v355)
	v360 = v326 - int32(2)
	if v360 != 0 {
		v325 = v325 + int32(8)
		v326 = v360
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
	v379 = int32(0)
	goto L87
L87:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379+(v374+int32(8))))))
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379+v292))))
	v394 = v389 & (v391 ^ int32(255))
	v396 = base.B2i32(v394 == int32(0))
	if v394 != 0 {
		v402 = v396
		goto L28
	} else {
		goto L89
	}
L88:
	;
	v402 = v396
	goto L28
L89:
	;
	v398 = v379 + int32(1)
	if v398 != v44 {
		v379 = v398
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v414 = v402
	goto L17
L92:
	;
	v426 = v414
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
	F_errmsg(m, int32(_a_F_g_intbig_consistent_1), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_g_intbig_consistent_2), int32(492), int32(_a_F_g_intbig_consistent_3))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
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
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_g_intbig_consistent_1), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_g_intbig_consistent_2), int32(82), int32(_a_F_g_intbig_consistent_4))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
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
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_g_intbig_options_0), int32(_a_F_g_intbig_options_1), int32(252), int32(1), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
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
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
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
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v416 int64
	_ = v416
	var v418 int64
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
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
		v56 = v21
		v60 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v200 = base.B2i32(v198 < int32(0))
	if v198 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L2:
	;
	v198 = v188
	goto L1
L3:
	;
	if base.B2i32(v23 <= int32(0))|base.B2i32(v24 <= v56) != 0 {
		v92 = v24
		v94 = v4
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v37 = v21
	v41 = v4
	goto L5
L5:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v41<<(uint(int32(1))%32)))))
	if v47 != 0 {
		v188 = int32(1)
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v56 = v51
	v60 = v49
	goto L3
L7:
	;
	v48 = int32(1)
	v49 = v41 + v48
	v51 = v37 - v48
	if v51 <= v24 {
		v56 = v51
		v60 = v49
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v49 < v20 {
		v37 = v51
		v41 = v49
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	if v56 != v92 {
		v134 = v60
		v135 = v94
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v73 = v24
	v75 = v4
	goto L12
L12:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v75<<(uint(int32(1))%32)))))
	if v80 != 0 {
		v188 = int32(-1)
		goto L2
	} else {
		goto L14
	}
L13:
	;
	v92 = v84
	v94 = v82
	goto L10
L14:
	;
	v81 = int32(1)
	v82 = v75 + v81
	v84 = v73 - v81
	if v84 <= v56 {
		v92 = v84
		v94 = v82
		goto L10
	} else {
		goto L15
	}
L15:
	;
	if v82 < v23 {
		v73 = v84
		v75 = v82
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	if v20 < v134 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v103 = v60
	v104 = v94
	goto L19
L19:
	;
	if base.B2i32(v20 <= v103)|base.B2i32(v23 <= v104) != 0 {
		v134 = v103
		v135 = v104
		goto L17
	} else {
		goto L21
	}
L20:
	;
	if base.I32_extend16_s(v120) < base.I32_extend16_s(v118) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v109 = int32(1)
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v103<<(uint(v109)%32)))))
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104<<(uint(v109)%32)+v22))))
	if v118 == v120 {
		v103 = v103 + v109
		v104 = v104 + v109
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v127 = int32(1)
	goto L25
L24:
	;
	v127 = int32(-1)
	goto L25
L25:
	;
	v198 = v127
	goto L1
L26:
	;
	v138 = v134
	goto L28
L27:
	;
	v138 = v20
	goto L28
L28:
	;
	v145 = v134
	goto L29
L29:
	;
	if v138 == v145 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v188 = v171
	goto L2
L31:
	;
	if v23 < v135 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v171 = int32(1)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v145<<(uint(v171)%32)))))
	if v177 == int32(0) {
		v145 = v145 + v171
		goto L29
	} else {
		goto L43
	}
L34:
	;
	v150 = v135
	goto L36
L35:
	;
	v150 = v23
	goto L36
L36:
	;
	v158 = v135
	goto L37
L37:
	;
	if v150 == v158 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v188 = int32(-1)
	goto L2
L39:
	;
	v198 = int32(0)
	goto L1
L40:
	;
	goto L41
L41:
	;
	v162 = int32(1)
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158<<(uint(v162)%32)+v22))))
	if v167 == int32(0) {
		v158 = v158 + v162
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	goto L30
L44:
	;
	v201 = l1
	goto L46
L45:
	;
	v201 = l0
	goto L46
L46:
	;
	if v198 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v202 = v23
	goto L49
L48:
	;
	v202 = v20
	goto L49
L49:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v204 < v203 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v206 = v203
	goto L52
L51:
	;
	v206 = v204
	goto L52
L52:
	;
	if v198 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	m.G0 = v17 + int32(80)
	return
L54:
	;
	v208 = v20
	goto L56
L55:
	;
	v208 = v23
	goto L56
L56:
	;
	if v208 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v209 = v198
	goto L59
L58:
	;
	v209 = int32(0)
	goto L59
L59:
	;
	if v209 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v216 = F_palloc(m, v202<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v198 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L63:
	;
	return
L64:
	;
	v218 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v216))) = uint16(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v220 <= v218 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v232 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v224 = v220 << (uint(int32(1)) % 32)
	if v224 == int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v201)+20))
	base.MemoryCopy(m, v216+int32(2), v229, v224)
	goto L65
L68:
	;
	F_pfree(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L63
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v201)))
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v201)+8))
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v201)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v236
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v235
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v216 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	goto L53
L71:
	;
	goto L70
L72:
	;
	v248 = l0
	goto L74
L73:
	;
	v248 = l1
	goto L74
L74:
	;
	v249 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v249
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v249
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v249
	v259 = F_palloc(m, v202<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L63
	} else {
		goto L75
	}
L75:
	;
	v261 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v259))) = uint16(v261)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v263 <= v261 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v201)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v276
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v201)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v259
	v281 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v259 + v281
	v288 = F_palloc(m, v275<<(uint(int32(1))%32)+v281)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L63
	} else {
		goto L79
	}
L77:
	;
	v267 = v263 << (uint(int32(1)) % 32)
	if v267 == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v201)+20))
	base.MemoryCopy(m, v259+int32(2), v272, v267)
	goto L76
L79:
	;
	v290 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v288))) = uint16(v290)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if v292 <= v290 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v304 != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v296 = v292 << (uint(int32(1)) % 32)
	if v296 == int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v248)+20))
	base.MemoryCopy(m, v288+int32(2), v301, v296)
	goto L80
L83:
	;
	F_pfree(m, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L63
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v248)))
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v248)+8))
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v248)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v309
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v308
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v307
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v288
	v314 = v288
	v315 = v259
	goto L87
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v314 + int32(2)
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_gcd_var[0]))
	if v332 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	F_pfree(m, v315)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L63
	} else {
		goto L116
	}
L89:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L63
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v335 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v335
	v342 = v17 + int32(32)
	v344 = v17 + int32(56)
	v345 = int32(0)
	F_div_var(m, v342, l2, v344, v345, v345, int32(1))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L63
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_mul_var(m, l2, v344, v344, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L63
	} else {
		goto L94
	}
L94:
	;
	F_sub_var(m, v342, v344, v17+int32(8))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L63
	} else {
		goto L95
	}
L95:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v357 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_pfree(m, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L63
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v360 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v366 = F_palloc(m, v361<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L63
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	goto L88
L103:
	;
	v368 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v366))) = uint16(v368)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v370 <= v368 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	F_pfree(m, v315)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L63
	} else {
		goto L107
	}
L105:
	;
	v374 = v370 << (uint(int32(1)) % 32)
	if v374 == int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	base.MemoryCopy(m, v366+int32(2), v379, v374)
	goto L104
L107:
	;
	v384 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v384
	v386 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v366
	v389 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v366 + v389
	v393 = v360 << (uint(int32(1)) % 32)
	v396 = F_palloc(m, v393+v389)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L63
	} else {
		goto L108
	}
L108:
	;
	v398 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v396))) = uint16(v398)
	if base.B2i32(v393 == v398)|base.B2i32(v360 <= v398) == v398 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	base.MemoryCopy(m, v396+int32(2), v409, v393)
	goto L111
L110:
	;
	goto L111
L111:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v411 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_pfree(m, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L63
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v414
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v416
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v418
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v396
	v314 = v396
	v315 = v366
	goto L87
L115:
	;
	goto L114
L116:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v426 == int32(0) {
		goto L53
	} else {
		goto L117
	}
L117:
	;
	F_pfree(m, v426)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L63
	} else {
		goto L118
	}
L118:
	;
	goto L53
}
func F_gen_random_uuid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v3 = F_palloc(m, int32(16))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = m.Env.Pgmem_random_bytes(m, v3, int32(16))
		mBase = m.M
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_gen_random_uuid_0), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_gen_random_uuid_1), int32(541), int32(_a_F_gen_random_uuid_2))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
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
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
			v31 = v27&int32(15) | int32(64)
			*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)) = uint8(v31)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+8)))
			v37 = v33&int32(63) | int32(128)
			*(*uint8)(unsafe.Add(mBase, uint32(v3)+8)) = uint8(v37)
			return v3
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
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
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
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
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
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
	v228 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L3
	} else {
		goto L69
	}
L2:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v192 != 0 {
		goto L62
	} else {
		goto L63
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
	v181 = int32(0)
	v191 = v9
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
	if v156 == int32(0) {
		v181 = v153
		v191 = v158
		goto L2
	} else {
		goto L50
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v38 <= v32 {
		v153 = v27
		v156 = v35
		v158 = v37
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_errorConflictingDefElem(m, v44, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L49
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v32<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v46 = int32(_a_F_generateSerialExtraStmts_0)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generateSerialExtraStmts[0])))
	if base.B2i32(v49 == int32(0))|base.B2i32(v49 != v52) != 0 {
		v70 = v49
		v71 = v52
		goto L16
	} else {
		goto L17
	}
L12:
	;
	goto L10
L13:
	;
	if v144 != 0 {
		v27 = v144
		v32 = v145 + int32(1)
		v35 = v146
		v37 = v147
		goto L9
	} else {
		goto L48
	}
L14:
	;
	v142 = F_list_delete_nth_cell(m, v27, v32)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L47
	}
L15:
	;
	if v70-v71 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	goto L15
L17:
	;
	v55 = v45
	v56 = v46
	goto L18
L18:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v70 = v60
		v71 = v59
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v70 = v60
	v71 = v59
	goto L16
L20:
	;
	v63 = int32(1)
	if v60 == v59 {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v35 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v80 = int32(_a_F_generateSerialExtraStmts_1)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generateSerialExtraStmts[1])))
	if base.B2i32(v83 == int32(0))|base.B2i32(v83 != v86) != 0 {
		v104 = v83
		v105 = v86
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v138 = v44
	v139 = v37
	goto L14
L26:
	;
	goto L27
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_errorConflictingDefElem(m, v44, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	if v37 != 0 {
		goto L12
	} else {
		goto L46
	}
L30:
	;
	if v104-v105 == int32(0) {
		goto L29
	} else {
		goto L37
	}
L31:
	;
	goto L30
L32:
	;
	v89 = v45
	v90 = v80
	goto L33
L33:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v94 == int32(0) {
		v104 = v94
		v105 = v93
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v104 = v94
	v105 = v93
	goto L31
L35:
	;
	v97 = int32(1)
	if v94 == v93 {
		v89 = v89 + v97
		v90 = v90 + v97
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v109 = int32(_a_F_generateSerialExtraStmts_2)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generateSerialExtraStmts[2])))
	if base.B2i32(v112 == int32(0))|base.B2i32(v112 != v115) != 0 {
		v133 = v112
		v134 = v115
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v133-v134 == int32(0) {
		goto L29
	} else {
		goto L45
	}
L39:
	;
	goto L38
L40:
	;
	v118 = v45
	v119 = v109
	goto L41
L41:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	if v123 == int32(0) {
		v133 = v123
		v134 = v122
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v133 = v123
	v134 = v122
	goto L39
L43:
	;
	v126 = int32(1)
	if v123 == v122 {
		v118 = v118 + v126
		v119 = v119 + v126
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v144 = v27
	v145 = v32
	v146 = v35
	v147 = v37
	goto L13
L46:
	;
	v138 = v35
	v139 = v44
	goto L14
L47:
	;
	v144 = v142
	v145 = v32 - int32(1)
	v146 = v138
	v147 = v139
	goto L13
L48:
	;
	v153 = v144
	v156 = v146
	v158 = v147
	goto L8
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v162 = F_makeRangeVarFromNameList(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	if v164 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v167 != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v176 = v164
	goto L54
L54:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v214 = v153
	v221 = v176
	v224 = v158
	v225 = v177
	goto L1
L55:
	;
	v174 = F_get_namespace_name(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L60
	}
L56:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+48))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+68))
	v173 = v169
	goto L55
L57:
	;
	goto L58
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v171 = F_RangeVarGetCreationNamespace(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v173 = v171
	goto L55
L60:
	;
	v176 = v174
	goto L54
L61:
	;
	v202 = F_get_namespace_name(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L67
	}
L62:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+48))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+68))
	v201 = v194
	goto L61
L63:
	;
	goto L64
L64:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v196 = F_RangeVarGetCreationNamespace(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_RangeVarAdjustRelationPersistence(m, v198, v196)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v201 = v196
	goto L61
L67:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v209 = F_ChooseRelationName(m, v205, v206, int32(_a_F_generateSerialExtraStmts_3), v201, int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v214 = v181
	v221 = v202
	v224 = v191
	v225 = v209
	goto L1
L69:
	;
	if v228 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v230
	F_errmsg_internal(m, int32(_a_F_generateSerialExtraStmts_4), v17+int32(16))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L3
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v250 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	F_errfinish(m, int32(_a_F_generateSerialExtraStmts_5), int32(486), int32(_a_F_generateSerialExtraStmts_6))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	if v224 != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+48))
	v257 = v251 + int32(118)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v257 = v254 + int32(17)
	goto L75
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L3
	} else {
		goto L126
	}
L80:
	;
	if v258&int32(255) == int32(116) {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	v294 = v258
	goto L82
L82:
	;
	v296 = F_palloc0(m, int32(20))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L3
	} else {
		goto L94
	}
L83:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	v266 = int32(_a_F_generateSerialExtraStmts_1)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generateSerialExtraStmts[1])))
	if base.B2i32(v269 == int32(0))|base.B2i32(v269 != v272) != 0 {
		v290 = v269
		v291 = v272
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v290-v291 != 0 {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	goto L84
L86:
	;
	v275 = v265
	v276 = v266
	goto L87
L87:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	if v280 == int32(0) {
		v290 = v280
		v291 = v279
		goto L85
	} else {
		goto L89
	}
L88:
	;
	v290 = v280
	v291 = v279
	goto L85
L89:
	;
	v283 = int32(1)
	if v280 == v279 {
		v275 = v275 + v283
		v276 = v276 + v283
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v293 = int32(117)
	goto L93
L92:
	;
	v293 = int32(112)
	goto L93
L93:
	;
	v294 = v293
	goto L82
L94:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v296)+16)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v296))) = int32(189)
	v302 = F_makeRangeVar(m, v221, v225, int32(-1))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+4)) = v302
	*(*uint8)(unsafe.Add(mBase, uint32(v302)+17)) = uint8(v294)
	*(*int32)(unsafe.Add(mBase, uint32(v296)+8)) = v214
	if l2 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v308 = F_makeTypeNameFromOid(m, l2)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v317 != 0 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v311 = F_makeDefElem(m, int32(_a_F_generateSerialExtraStmts_7), v308, int32(-1))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v296)+8))
	v314 = F_lcons(m, v311, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+8)) = v314
	goto L98
L102:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+48))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+80))
	v321 = v319
	goto L104
L103:
	;
	v321 = int32(0)
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+12)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v324 = F_lappend(m, v323, v296)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v324
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v327
	v330 = F_palloc0(m, int32(16))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = int32(190)
	v335 = F_makeRangeVar(m, v221, v225, int32(-1))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+4)) = v335
	v338 = F_makeString(m, v221)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v338
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v343 = F_makeString(m, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v343
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v347 = F_makeString(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v347
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v351
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v353
	v362 = F_list_make3_impl(m, v17+int32(12), v17+int32(8), v17+int32(4))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	v365 = F_makeDefElem(m, int32(_a_F_generateSerialExtraStmts_8), v362, int32(-1))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v365
	v370 = F_list_make1_impl(m, int32(1), v17)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v330)+12)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+8)) = v370
	if l5 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if l6 != 0 {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v375 = F_lappend(m, v374, v330)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L3
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v379 = F_lappend(m, v378, v330)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L3
	} else {
		goto L119
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v375
	goto L114
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v379
	goto L114
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v221
	goto L122
L121:
	;
	goto L122
L122:
	;
	if l7 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v225
	goto L125
L124:
	;
	goto L125
L125:
	;
	m.G0 = v17 + int32(48)
	return
L126:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(_a_F_generateSerialExtraStmts_9), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	F_parser_errposition(m, v398, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_generateSerialExtraStmts_5), int32(505), int32(_a_F_generateSerialExtraStmts_6))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
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
	var v64 int32
	_ = v64
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
	F_pg_qsort(m, v27, v20, int32(3), int32(_a_F_generate_trgm_0))
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
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_generate_trgm[0]))
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
	v78 = v64 + int32(1)
	goto L5
L9:
	;
	v66 = v34 + int32(1)
	if v66 != v20 {
		v34 = v66
		v35 = v64
		goto L7
	} else {
		goto L13
	}
L10:
	;
	if v49 == int32(0) {
		v64 = v35
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v54 = v35 + int32(1)
	if v54 == v34 {
		v64 = v34
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v58 = v27 + v54*int32(3)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)) = uint8(v59)
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
	*(*uint16)(unsafe.Add(mBase, uint32(v58))) = uint16(v61)
	v64 = v54
	goto L9
L13:
	;
	goto L8
}
func F_getNextFlagFromString(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
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
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v15 == int32(0) {
		v179 = l2
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L13
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L13
	} else {
		goto L69
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L65
	}
L4:
	;
	v221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v221)
	m.G0 = v12 + int32(128)
	return
L5:
	;
	v205 = F_pg_mblen_cstr(m, v28)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L13
	} else {
		goto L61
	}
L6:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v186 != int32(1) {
		v214 = v179
		goto L4
	} else {
		goto L56
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if base.Ui32(int32(2)) <= base.Ui32(v18) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getNextFlagFromString[0])) = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v52 = F_strtox_2(m, v47, v12+int32(124), int32(10), int64(2147483648))
	mBase = m.M
	v53 = base.I32_wrap_i64(v52)
	goto L22
L9:
	;
	if v18 == int32(2) {
		v41 = l2
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = F_pg_mblen_cstr(m, v14)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L1
L13:
	;
	return
L14:
	;
	if v24 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	base.MemoryCopy(m, l2, v14, v24)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = v27 + v24
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28
	v30 = l2 + v24
	if v23 != int32(1) {
		v214 = v30
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v33 == int32(0) {
		v179 = v30
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if base.Ui32(v36) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	if v36 != int32(2) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v41 = v30
	goto L8
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v54 == v55 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_getNextFlagFromString[0]))
	if v58 == int32(68) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(int32(_a_F_getNextFlagFromString_0)) <= base.Ui32(v53) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v53
	v67 = F_pg_sprintf(m, v41, int32(_a_F_getNextFlagFromString_1), v12+int32(112))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v54
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v70 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v179 = v41 + v67
	goto L6
L28:
	;
	v78 = int32(0)
	v79 = v70
	v80 = v54
	goto L29
L29:
	;
	if base.Ui32((v79-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L27
L31:
	;
	if v78 != 0 {
		goto L27
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if base.Ui32(v79-int32(9)) < base.Ui32(int32(5)) {
		v159 = v78
		goto L39
	} else {
		goto L40
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v96
	F_errmsg(m, int32(_a_F_getNextFlagFromString_2), v12-int32(-64))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_getNextFlagFromString_3), int32(402), int32(_a_F_getNextFlagFromString_4))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v161 = F_pg_mblen_cstr(m, v80)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L54
	}
L40:
	;
	v113 = v79 - int32(32)
	if v113 == int32(0) {
		v159 = v78
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if v113 == int32(12) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v78 == int32(0) {
		v159 = int32(1)
		goto L39
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L13
	} else {
		goto L50
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v128
	F_errmsg(m, int32(_a_F_getNextFlagFromString_2), v12+int32(96))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_getNextFlagFromString_3), int32(411), int32(_a_F_getNextFlagFromString_4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v147
	F_errmsg(m, int32(_a_F_getNextFlagFromString_5), v12+int32(80))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_getNextFlagFromString_3), int32(419), int32(_a_F_getNextFlagFromString_4))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v164 = v161 + v163
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v164
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v166 != 0 {
		v78 = v159
		v79 = v166
		v80 = v164
		goto L29
	} else {
		goto L55
	}
L55:
	;
	goto L30
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
	F_errmsg(m, int32(_a_F_getNextFlagFromString_6), v12)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_getNextFlagFromString_3), int32(439), int32(_a_F_getNextFlagFromString_4))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	if v205 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryCopy(m, v30, v28, v205)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v208 + v205
	v214 = v205 + v30
	goto L4
L65:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v233
	F_errmsg(m, int32(_a_F_getNextFlagFromString_2), v12+int32(32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_getNextFlagFromString_3), int32(384), int32(_a_F_getNextFlagFromString_4))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L13
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
	F_errcode(m, int32(22))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v252
	F_errmsg(m, int32(_a_F_getNextFlagFromString_7), v12+int32(48))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_getNextFlagFromString_3), int32(389), int32(_a_F_getNextFlagFromString_4))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v271
	F_errmsg_internal(m, int32(_a_F_getNextFlagFromString_8), v12+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_getNextFlagFromString_3), int32(428), int32(_a_F_getNextFlagFromString_4))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	v3 = int32(0)
	if l0 == v3 {
		v65 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v65
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(0) {
		v65 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v11 <= int32(0) {
		v65 = v3
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
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v32 == int32(0))|base.B2i32(v32 != v35) != 0 {
		v53 = v32
		v54 = v35
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v65 = int32(0)
	goto L1
L10:
	;
	if v53-v54 == int32(0) {
		v65 = v28
		goto L1
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v38 = v29
	v39 = l1
	goto L13
L13:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v43
		v54 = v42
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v53 = v43
	v54 = v42
	goto L11
L15:
	;
	v46 = int32(1)
	if v43 == v42 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v59 = v20 + int32(1)
	if v59 != v17 {
		v20 = v59
		goto L8
	} else {
		goto L18
	}
L18:
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
	var v45 int32
	_ = v45
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v45
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 <= int32(0) {
		v45 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v45 = int32(0)
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
		v45 = v23
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
		v45 = v23
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
	var v28 int32
	_ = v28
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
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
	v93 = m.ExcPending
	if v93 != 0 {
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
		v88 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v28 = v3
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
	v88 = v79
	goto L4
L10:
	;
	v81 = v30 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v81 < v82 {
		v28 = v79
		v30 = v81
		goto L8
	} else {
		goto L26
	}
L11:
	;
	v79 = v28
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
		v88 = v48
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
	v88 = v48
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
	v72 = base.B2i32(v68 != v70) | v28
	if v68 == v70 {
		v79 = v72
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
	v88 = v72
	goto L4
L26:
	;
	goto L9
L27:
	;
	return v88 & int32(1)
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
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
	F_appendStringInfoString(m, v14, int32(_a_F_get_const_expr_0))
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
	F_appendStringInfo(m, v14, int32(_a_F_get_const_expr_1), v12+int32(16))
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
	F_appendStringInfo(m, v36, int32(_a_F_get_const_expr_2), v12)
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
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L71
	}
L18:
	;
	v221 = int32(0)
	goto L17
L19:
	;
	F_appendStringInfoString(m, v14, v57)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L70
	}
L20:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if base.Ui32((v117-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L47
	} else {
		goto L48
	}
L21:
	;
	F_appendStringInfoChar(m, v14, int32(39))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
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
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v72 != int32(116) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v62 != int32(45) {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v57
	F_appendStringInfo(m, v14, int32(_a_F_get_const_expr_3), v12-int32(-64))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v221 = int32(1)
	goto L17
L27:
	;
	F_appendStringInfoString(m, v14, int32(_a_F_get_const_expr_4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L31
	}
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v75 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_appendStringInfoString(m, v14, int32(_a_F_get_const_expr_5))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	goto L18
L31:
	;
	goto L18
L32:
	;
	goto L21
L33:
	;
	v93 = v57
	goto L34
L34:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v97 = base.I32_extend8_s(v96)
	if v96 != int32(39) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	F_appendStringInfoChar(m, v14, int32(39))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L46
	}
L36:
	;
	goto L35
L37:
	;
	F_appendStringInfoChar(m, v14, v97)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L45
	}
L38:
	;
	if v96 == int32(0) {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_appendStringInfoChar(m, v14, v97)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L44
	}
L41:
	;
	if v97 != int32(92) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_const_expr[0])))
	if v105&int32(1) != 0 {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	goto L37
L45:
	;
	v93 = v93 + int32(1)
	goto L34
L46:
	;
	goto L18
L47:
	;
	v124 = int32(_a_F_get_const_expr_6)
	v128 = m.G0
	v130 = v128 - int32(32)
	m.G0 = v130
	v132 = int32(*(*int8)(unsafe.Add(mBase, _c_F_get_const_expr[1])))
	if v132 != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v57
	F_appendStringInfo(m, v14, int32(_a_F_get_const_expr_3), v12+int32(80))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L69
	}
L50:
	;
	v191 = F_strlen(m, v57)
	mBase = m.M
	if v183-v57 != v191 {
		goto L19
	} else {
		goto L68
	}
L51:
	;
	m.G0 = v130 + int32(32)
	goto L50
L52:
	;
	F___memset(m, v130, int32(0), int32(32))
	mBase = m.M
	v138 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_const_expr[1])))
	if v138 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_const_expr[2])))
	if v133 != 0 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v134 = F___strchrnul(m, v57, v132)
	mBase = m.M
	v183 = v134
	goto L51
L56:
	;
	goto L55
L57:
	;
	v140 = v124
	v141 = v138
	goto L60
L58:
	;
	goto L59
L59:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v162 == int32(0) {
		v183 = v57
		goto L51
	} else {
		goto L63
	}
L60:
	;
	v148 = v130 + int32(base.Ui32(v141)>>(uint(int32(3))%32))&int32(28)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v150 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v149 | v150<<(uint(v141)%32)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)))
	if v154 != 0 {
		v140 = v140 + v150
		v141 = v154
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	goto L61
L63:
	;
	v166 = v57
	v167 = v162
	goto L64
L64:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v130+int32(base.Ui32(v167)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v175)>>(uint(v167)%32))&int32(1) != 0 {
		v183 = v166
		goto L51
	} else {
		goto L66
	}
L65:
	;
	v183 = v181
	goto L51
L66:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	v181 = v166 + int32(1)
	if v179 != 0 {
		v166 = v181
		v167 = v179
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	goto L49
L69:
	;
	v221 = int32(1)
	goto L17
L70:
	;
	goto L18
L71:
	;
	if l2 < int32(0) {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v226 - int32(16) {
	case 0:
		goto L77
	case 1, 2, 3, 4, 5, 6:
		goto L74
	case 7:
		v238 = v221
		goto L75
	default:
		goto L78
	}
L73:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v252 == int32(0) {
		goto L1
	} else {
		goto L85
	}
L74:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v243 = F_format_type_with_typemod(m, v226, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L83
	}
L75:
	;
	if l2 != 0 {
		goto L74
	} else {
		goto L81
	}
L76:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v238 = v221 | base.B2i32(int32(0) <= v234)
	goto L75
L77:
	;
	v238 = int32(0)
	goto L75
L78:
	;
	if v226 == int32(1700) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	if v226 != int32(705) {
		goto L74
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if v238 == int32(0) {
		goto L73
	} else {
		goto L82
	}
L82:
	;
	goto L74
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v243
	F_appendStringInfo(m, v14, int32(_a_F_get_const_expr_1), v12+int32(48))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	goto L73
L85:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v257 = F_get_typcollation(m, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v257 == v259 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v261 = F_generate_collation_name(m, v259)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v261
	F_appendStringInfo(m, v255, int32(_a_F_get_const_expr_2), v12+int32(32))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v11 = int32(2)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	switch v12 - int32(4) {
	case 0:
		v67 = int32(3)
		m.G0 = v9 + int32(112)
		return v67
	default:
		if l2|base.B2i32(v12 != int32(10)) == int32(0) {
			v67 = int32(4)
			m.G0 = v9 + int32(112)
			return v67
		} else {
			if l2 != 0 {
				v26 = F___fstatat(m, int32(-100), l0, v9+int32(16), int32(0))
				mBase = m.M
				v32 = v26
			} else {
				v31 = F___fstatat(m, int32(-100), l0, v9+int32(16), int32(256))
				mBase = m.M
				v32 = v31
			}
			if v32 < int32(0) {
				v35 = int32(0)
				v37 = F_errstart(m, l3, v35)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					if v37 == int32(0) {
						v67 = v35
						m.G0 = v9 + int32(112)
						return v67
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg(m, int32(_a_F_get_dirent_type_0), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_dirent_type_1), int32(592), int32(_a_F_get_dirent_type_2))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v67 = v35
									m.G0 = v9 + int32(112)
									return v67
								}
							}
						}
					}
				}
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				v56 = v54 & int32(_a_F_get_dirent_type_3)
				if v56 != int32(_a_F_get_dirent_type_4) {
					if v56 == int32(_a_F_get_dirent_type_5) {
						v67 = v11
					} else {
						if v56 != int32(_a_F_get_dirent_type_6) {
							v67 = int32(1)
						} else {
							v67 = int32(3)
						}
					}
				} else {
					v67 = int32(4)
				}
				m.G0 = v9 + int32(112)
				return v67
			}
		}
	case 4:
		v67 = v11
		m.G0 = v9 + int32(112)
		return v67
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
	var v36 int32
	_ = v36
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
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v161 int32
	_ = v161
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
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
	v176 = int32(0)
	goto L3
L3:
	;
	F_list_sort(m, v176, int32(1059))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L20
	} else {
		goto L41
	}
L4:
	;
	v36 = v6
	goto L7
L5:
	;
	goto L6
L6:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v176 = v161
	goto L3
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v36<<(uint(int32(2))%32))))
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
	v145 = v36 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v145 < v146 {
		v36 = v145
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
	F_errmsg_internal(m, int32(_a_F_get_policies_for_relation_0), v16)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_policies_for_relation_1), int32(590), int32(_a_F_get_policies_for_relation_2))
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
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+12)))
	if v125 != 0 {
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
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	if v86 <= v85 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v94 = v85
	goto L31
L31:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v81+v94<<(uint(int32(2))%32))))
	v106 = F_has_privs_of_role(m, l2, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L20
	} else {
		goto L33
	}
L32:
	;
	goto L9
L33:
	;
	if v106 != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v109 = v94 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	if v109 < v110 {
		v94 = v109
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v126 = l3
	goto L38
L37:
	;
	v126 = l4
	goto L38
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = F_lappend(m, v127, v46)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v128
	goto L9
L40:
	;
	goto L8
L41:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_get_policies_for_relation[0]))
	if v181 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_get_policies_for_relation[1]))
	if v302 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L43:
	;
	v184 = m.T0[v181].(func(*base.Module, int32, int32) int32)(m, l1, l0)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	F_list_sort(m, v184, int32(1059))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	if v184 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v191 <= int32(0) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v202 = int32(0)
	goto L48
L48:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208+v202<<(uint(int32(2))%32))))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	if v214 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L42
L50:
	;
	v285 = v202 + int32(1)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v285 < v286 {
		v202 = v285
		goto L48
	} else {
		goto L63
	}
L51:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v268 = F_lappend(m, v267, v212)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L20
	} else {
		goto L62
	}
L52:
	;
	v222 = v214
	goto L54
L53:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v222 = (v215<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L54
L54:
	;
	v223 = v222 + v213
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v224 == int32(0) {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v227 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	if v228 <= v227 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v236 = v227
	goto L57
L57:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v223+v236<<(uint(int32(2))%32))))
	v248 = F_has_privs_of_role(m, l2, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L20
	} else {
		goto L59
	}
L58:
	;
	goto L50
L59:
	;
	if v248 != 0 {
		goto L51
	} else {
		goto L60
	}
L60:
	;
	v251 = v236 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	if v251 < v252 {
		v236 = v251
		goto L57
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v268
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
	v305 = m.T0[v302].(func(*base.Module, int32, int32) int32)(m, l1, l0)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	if v305 == int32(0) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v309 <= int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v320 = int32(0)
	goto L69
L69:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v305)+12))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326+v320<<(uint(int32(2))%32))))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	if v332 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	goto L64
L71:
	;
	v403 = v320 + int32(1)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v403 < v404 {
		v320 = v403
		goto L69
	} else {
		goto L84
	}
L72:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v386 = F_lappend(m, v385, v330)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L20
	} else {
		goto L83
	}
L73:
	;
	v340 = v332
	goto L75
L74:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	v340 = (v333<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L75
L75:
	;
	v341 = v340 + v331
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	if v342 == int32(0) {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v345 = int32(0)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v331)+16))
	if v346 <= v345 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v354 = v345
	goto L78
L78:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v341+v354<<(uint(int32(2))%32))))
	v366 = F_has_privs_of_role(m, l2, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L20
	} else {
		goto L80
	}
L79:
	;
	goto L71
L80:
	;
	if v366 != 0 {
		goto L72
	} else {
		goto L81
	}
L81:
	;
	v369 = v354 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v331)+16))
	if v369 < v370 {
		v354 = v369
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v386
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
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v10 == int32(0) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
		if v9&int32(1) != 0 {
			if v13&int32(1) == int32(0) {
				return float64(0)
			} else {
				v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
				if v83 != 0 {
					v84 = float64(0)
				} else {
					v84 = float64(1)
				}
				v86 = v84
				return v86
			}
		} else {
			v16 = float64(0.5)
			if v13&int32(1) != 0 {
				v86 = v16
				return v86
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				if v19 == int32(0) {
					v86 = v16
					return v86
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
						if base.F64_le(v31, float64(0))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v31)&int64(9223372036854775807))) != 0 {
							v86 = v16
							return v86
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v43 = F_FunctionCall2Coll(m, v23, v40, v41, v42)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return float64(0)
							} else {
								v45 = *(*float64)(unsafe.Add(mBase, uint32(v43)))
								v46 = base.F64_div(v45, v31)
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v46)&int64(9223372036854775807)) {
									v86 = v16
									return v86
								} else {
									v52 = float64(0)
									if base.F64_gt(v46, v52) != 0 {
										v55 = v46
									} else {
										v55 = v52
									}
									v56 = float64(1)
									if base.F64_lt(v55, v56) != 0 {
										v59 = v55
									} else {
										v59 = v56
									}
									return v59
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
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			if v65 != int32(1) {
				return float64(1)
			} else {
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
				if v72 != 0 {
					v73 = float64(0)
				} else {
					v73 = float64(1)
				}
				return v73
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
								if v28 == int32(_a_F_get_promoted_array_type_0) {
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
						if v28 == int32(_a_F_get_promoted_array_type_0) {
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	v52 = int32(_a_F_get_reloptions_0)
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
	F_appendStringInfoString(m, l0, int32(_a_F_get_reloptions_1))
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
	F_appendStringInfo(m, l0, int32(_a_F_get_reloptions_2), v10)
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
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L43
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
		goto L34
	} else {
		goto L35
	}
L31:
	;
	F_appendStringInfoChar(m, l0, int32(39))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L42
	}
L32:
	;
	goto L31
L33:
	;
	F_appendStringInfoChar(m, l0, v78)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L41
	}
L34:
	;
	if v77 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_appendStringInfoChar(m, l0, v78)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L40
	}
L37:
	;
	if v78 != int32(92) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_reloptions[0])))
	if v86&int32(1) != 0 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	goto L33
L41:
	;
	v71 = v71 + int32(1)
	goto L30
L42:
	;
	goto L23
L43:
	;
	v108 = v28 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v108 < v109 {
		v28 = v108
		goto L7
	} else {
		goto L44
	}
L44:
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
	F_make_relative_path(m, l0, int32(_a_F_get_share_path_0), int32(_a_F_get_share_path_1))
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
	if v26 == l1&int32(_a_F_get_tle_by_resno_0) {
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	v6 = m.G0
	v8 = v6 - int32(1056)
	m.G0 = v8
	v10 = int32(_a_F_get_tsearch_config_filename_0)
	v14 = m.G0
	v16 = v14 - int32(32)
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v17
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_tsearch_config_filename[0])))
	if v25 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v94 = F_strlen(m, l0)
	mBase = m.M
	if v93 != v94 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v93 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_tsearch_config_filename[1])))
	if v29 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = l0
	goto L8
L6:
	;
	goto L7
L7:
	;
	v43 = v10
	v44 = v25
	goto L11
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v39 == v25 {
		v33 = v33 + int32(1)
		goto L8
	} else {
		goto L10
	}
L9:
	;
	v93 = v33 - l0
	goto L1
L10:
	;
	goto L9
L11:
	;
	v51 = v16 + int32(base.Ui32(v44)>>(uint(int32(3))%32))&int32(28)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v52 | v53<<(uint(v44)%32)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v57 != 0 {
		v43 = v43 + v53
		v44 = v57
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v60 == int32(0) {
		v83 = l0
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v93 = v83 - l0
	goto L1
L15:
	;
	v64 = l0
	v65 = v60
	goto L16
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(base.Ui32(v65)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v73)>>(uint(v65)%32))&int32(1) == int32(0) {
		v83 = v64
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v83 = v81
	goto L14
L18:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v81 = v64 + int32(1)
	if v79 != 0 {
		v64 = v81
		v65 = v79
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v117 = v8 + int32(32)
	F_get_share_path(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L23
	} else {
		goto L28
	}
L23:
	;
	return int32(0)
L24:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	F_errmsg(m, int32(_a_F_get_tsearch_config_filename_1), v8+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_get_tsearch_config_filename_2), int32(53), int32(_a_F_get_tsearch_config_filename_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v121 = F_palloc(m, int32(1024))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v117
	v128 = F_pg_snprintf(m, v121, int32(1024), int32(_a_F_get_tsearch_config_filename_4), v8)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	m.G0 = v8 + int32(1056)
	return v121
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
	var v36 int32
	_ = v36
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
	var v86 int32
	_ = v86
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
	var v116 int32
	_ = v116
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
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v374 int32
	_ = v374
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v22 = l0 + int32(8)
	goto L3
L1:
	;
	m.G0 = v15 + int32(32)
	return v374
L2:
	;
	v374 = int32(3)
	goto L1
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v36 - int32(1) {
	case 0, 2:
		goto L8
	case 1:
		goto L7
	default:
		v335 = v35
		goto L6
	}
L4:
	;
	v346 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35 + v346
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v346)
	goto L2
L5:
	;
	goto L4
L6:
	;
	v341 = F_pg_mblen_cstr(m, v335)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L12
	} else {
		goto L88
	}
L7:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v163 != int32(124) {
		goto L46
	} else {
		goto L47
	}
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if base.Ui32(v39-int32(9)) < base.Ui32(int32(5)) {
		v335 = v35
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v44 = int32(1)
	switch v39 - int32(32) {
	case 0:
		v335 = v35
		goto L6
	case 1:
		goto L5
	default:
		goto L10
	case 8:
		goto L11
	case 26:
		v374 = v44
		goto L1
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v35
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v60 = int32(0)
	v62 = F_gettoken_tsvector(m, v59, l3, l2, v60, v60, v22)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35 + v47
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v52 + v47
	v374 = int32(4)
	goto L1
L12:
	;
	return int32(0)
L13:
	;
	if v62 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v67 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v67)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v67)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 != int32(58) {
		v125 = v66
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v135 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L17:
	;
	v131 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v125
	v374 = v131
	goto L1
L18:
	;
	v75 = v66 + int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v76 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v125 = v75
	goto L17
L20:
	;
	goto L21
L21:
	;
	v86 = v75
	goto L22
L22:
	;
	v91 = F_pg_mblen_cstr(m, v86)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L24
	}
L23:
	;
	v125 = v118
	goto L17
L24:
	;
	if v91 != int32(1) {
		v125 = v86
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	switch v95 - int32(42) {
	case 0:
		goto L27
	default:
		v125 = v86
		goto L17
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
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v118 = v86 + int32(1)
	if v116 != 0 {
		v86 = v118
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
		v374 = v44
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v374 = int32(0)
	goto L1
L38:
	;
	goto L39
L39:
	;
	v146 = F_errsave_start(m, v135)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	if v146 == int32(0) {
		v374 = v44
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v153
	F_errmsg(m, int32(_a_F_gettoken_query_standard_0), v15)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, v135, int32(_a_F_gettoken_query_standard_1), int32(345), int32(_a_F_gettoken_query_standard_2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	v374 = v44
	goto L1
L45:
	;
	if v163 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	if v163 != int32(38) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v175 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35 + v175
	v180 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v180)
	v374 = v180
	goto L1
L49:
	;
	v168 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35 + v168
	v173 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v173)
	goto L2
L50:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v295 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L51:
	;
	v193 = v35
	v194 = v163
	v195 = int32(0)
	v198 = int32(1)
	goto L52
L52:
	;
	switch v195 - int32(1) {
	case 0:
		goto L59
	case 1:
		v210 = v193
		v211 = v194
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
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v282 != 0 {
		v193 = v278
		v194 = v282
		v195 = v280
		v198 = v281
		goto L52
	} else {
		goto L76
	}
L55:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v198)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	v276 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v276)
	goto L2
L56:
	;
	if base.Ui32(int32(9)) < base.Ui32((v194-int32(48))&int32(255)) {
		goto L50
	} else {
		goto L64
	}
L57:
	;
	if v194&int32(255) != int32(60) {
		goto L50
	} else {
		goto L63
	}
L58:
	;
	if v211&int32(255) != int32(62) {
		goto L50
	} else {
		goto L62
	}
L59:
	;
	if v194&int32(255) != int32(45) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	if v205 == int32(0) {
		goto L50
	} else {
		goto L61
	}
L61:
	;
	v210 = v193 + int32(1)
	v211 = v205
	goto L58
L62:
	;
	v278 = v210 + int32(1)
	v280 = int32(3)
	v281 = v198
	goto L54
L63:
	;
	v223 = int32(1)
	v278 = v193 + v223
	v280 = v223
	v281 = v198
	goto L54
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gettoken_query_standard[0])) = int32(0)
	v239 = F_strtox_2(m, v193, v15+int32(28), int32(10), int64(2147483648))
	mBase = m.M
	v240 = base.I32_wrap_i64(v239)
	goto L65
L65:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v193 == v241 {
		goto L50
	} else {
		goto L66
	}
L66:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_gettoken_query_standard[0]))
	if v244 != int32(68) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if base.Ui32(v240) < base.Ui32(int32(_a_F_gettoken_query_standard_3)) {
		v278 = v241
		v280 = int32(2)
		v281 = v240
		goto L54
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v253 = F_errsave_start(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L12
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	if v253 == int32(0) {
		goto L50
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_gettoken_query_standard_4)
	F_errmsg(m, int32(_a_F_gettoken_query_standard_5), v15+int32(16))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	F_errsave_finish(m, v252, int32(_a_F_gettoken_query_standard_1), int32(211), int32(_a_F_gettoken_query_standard_6))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
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
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if base.Ui32(v306-int32(9)) < base.Ui32(int32(5)) {
		v335 = v305
		goto L6
	} else {
		goto L81
	}
L78:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v298 != int32(447) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+4)))
	if v301 == int32(0) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v374 = int32(1)
	goto L1
L81:
	;
	v311 = int32(1)
	switch v306 - int32(32) {
	case 0:
		v335 = v305
		goto L6
	case 1, 2, 3, 4, 5, 6, 7, 8:
		v374 = v311
		goto L1
	case 9:
		goto L83
	default:
		goto L82
	}
L82:
	;
	if v306 != 0 {
		v374 = v311
		goto L1
	} else {
		goto L87
	}
L83:
	;
	v314 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v305 + v314
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v319 = v317 - v314
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v319
	if v319 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v325 = v314
	goto L86
L85:
	;
	v325 = int32(5)
	goto L86
L86:
	;
	v374 = v325
	goto L1
L87:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v374 = base.B2i32(v326 != int32(0))
	goto L1
L88:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v341 + v343
	goto L3
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
	var v18 int32
	_ = v18
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
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
	v18 = v4
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
	v43 = F_merge_clump(m, l0, v18, v35, int32(0))
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
		v18 = v43
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
	v82 = v43
	v84 = v51
	goto L12
L14:
	;
	goto L15
L15:
	;
	v54 = int32(0)
	v57 = v54
	v59 = v54
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
	v70 = F_merge_clump(m, l0, v57, v68, int32(1))
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
		v57 = v70
		v59 = v73
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
	v82 = v70
	v84 = v80
	goto L12
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v96 = v93
	goto L10
}
func F_ginarrayextract_2args(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v9 <= int32(2) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_ginarrayextract_2args_0), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_ginarrayextract_2args_1), int32(71), int32(_a_F_ginarrayextract_2args_2))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v28 = F_pg_detoast_datum_copy(m, v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
			F_get_typlenbyvalalign(m, v32, v7+int32(14), v7+int32(13), v7+int32(12))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+14)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)))
				v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7)+12)))
				F_deconstruct_array(m, v28, v42, v43, v44, v7+int32(8), v7+int32(4), v7)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v53
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					m.G0 = v7 + int32(16)
					return v55
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v108 int32
	_ = v108
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
		v108 = v16
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
	return v108 & int32(255)
L3:
	;
	v108 = int32(0)
	goto L2
L4:
	;
	v83 = v20
	goto L33
L5:
	;
	v61 = v23
	v63 = int32(1)
	goto L25
L6:
	;
	if v14 <= int32(0) {
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
		v108 = v16
		goto L2
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	v108 = int32(1)
	goto L2
L11:
	;
	v31 = int32(0)
	v33 = int32(0)
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
	v108 = v56
	goto L2
L14:
	;
	v43 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v15))))
	if v45 == v43 {
		v108 = v43
		goto L2
	} else {
		goto L17
	}
L15:
	;
	v56 = v33
	goto L16
L16:
	;
	v59 = v31 + int32(1)
	if v59 != v14 {
		v31 = v59
		v33 = v56
		goto L12
	} else {
		goto L24
	}
L17:
	;
	if v33&int32(255) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = v33
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
	v54 = v33
	goto L23
L23:
	;
	v56 = v54
	goto L16
L24:
	;
	goto L13
L25:
	;
	v69 = int32(0)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v15))))
	if v71 == v69 {
		v108 = v69
		goto L2
	} else {
		goto L27
	}
L26:
	;
	v108 = v79
	goto L2
L27:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v13))))
	if v75 != 0 {
		v108 = v69
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
	v79 = v63
	goto L31
L31:
	;
	v81 = v61 + int32(1)
	if v81 != v14 {
		v61 = v81
		v63 = v79
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
	v108 = v16
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
	F_errmsg_internal(m, int32(_a_F_ginarraytriconsistent_0), v11)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_ginarraytriconsistent_1), int32(300), int32(_a_F_ginarraytriconsistent_2))
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
		v9 = F_palloc(m, int32(_a_F_ginbeginscan_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_ginbeginscan[0]))) = int64(0)
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_ginbeginscan[1]))
			v19 = F_AllocSetContextCreateInternal(m, v14, int32(_a_F_ginbeginscan_1), int32(0), int32(_a_F_ginbeginscan_2), int32(_a_F_ginbeginscan_3))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_ginbeginscan[1]))
				v28 = F_AllocSetContextCreateInternal(m, v23, int32(_a_F_ginbeginscan_4), int32(0), int32(_a_F_ginbeginscan_2), int32(_a_F_ginbeginscan_3))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_ginbeginscan[2]))) = v28
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
	var v55 int32
	_ = v55
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
		*(*int32)(unsafe.Add(mBase, uint32(v3)+13)) = v7
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
		v55 = int32(5)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+29)) = uint8(v55)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+25)) = int32(_a_F_ginhandler_0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+17)) = int64(281479271678209)
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_copy(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		F_get_typlenbyvalalign(m, v20, v9+int32(30), v9+int32(29), v9+int32(28))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+30)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
			v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+28)))
			F_deconstruct_array(m, v12, v30, v31, v32, v9+int32(24), v9+int32(20), v9+int32(16))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v43
				switch v17 - int32(1) {
				case 0:
					v69 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v69
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					m.G0 = v9 + int32(32)
					return v71
				case 1:
					v65 = int32(0)
					if v41 <= v65 {
						v68 = int32(2)
					} else {
						v68 = v65
					}
					v69 = v68
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v69
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					m.G0 = v9 + int32(32)
					return v71
				case 2:
					v69 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v69
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					m.G0 = v9 + int32(32)
					return v71
				case 3:
					v69 = base.B2i32(v41 <= int32(0))
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v69
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					m.G0 = v9 + int32(32)
					return v71
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(_a_F_ginqueryarrayextract_0), v9)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ginqueryarrayextract_1), int32(131), int32(_a_F_ginqueryarrayextract_2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 float32
	_ = v200
	var v202 float32
	_ = v202
	var v205 float32
	_ = v205
	var v211 float32
	_ = v211
	var v215 float32
	_ = v215
	var v217 float32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 float32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
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
		v319 = v37
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v25 + int32(720)
	return v319 & int32(_a_F_gistchoose_0)
L4:
	;
	v46 = int32(base.Ui32(v38+int32(_a_F_gistchoose_1))>>(uint(int32(2))%32)) & int32(_a_F_gistchoose_0)
	if v46 == int32(0) {
		v319 = v37
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v52 = l3 + int32(_a_F_gistchoose_2)
	v65 = int32(-1)
	v66 = int32(1)
	v68 = v37
	goto L6
L6:
	;
	v82 = v66 & int32(_a_F_gistchoose_0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+10)))
	if v84 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v319 = v295
	goto L3
L8:
	;
	if base.B2i32(v264&int32(_a_F_gistchoose_0) == v82)|base.B2i32(v263 != base.I32_extend16_s(v260)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L9:
	;
	v260 = v84
	v261 = v65
	v263 = int32(0)
	v264 = v68
	v265 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v82<<(uint(int32(2))%32))))
	v104 = v65
	v106 = int32(0)
	v107 = v68
	v108 = int32(1)
	goto L12
L12:
	;
	v121 = v106 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v125 = F_index_getattr_2(m, l1+v92&int32(_a_F_gistchoose_3), v121, v122, v25+int32(15))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+10)))
	v260 = v253
	v261 = v104
	v263 = v106
	v264 = v107
	v265 = int32(0)
	goto L8
L14:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+15)))
	if v127 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(16)+v106))))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+716)) = int32(0)
	v180 = l3 + int32(3604) + v106*int32(28)
	if (v175|v127)&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+574)) = uint8(v168)
	goto L15
L17:
	;
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+574)) = uint8(v130)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+572)) = uint16(v66)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+568)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v25)+564)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v25)+560)) = v125
	v138 = l3 + int32(2708) + v106*int32(28)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v139 == v130 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+572)) = uint16(v66)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+568)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v25)+564)) = l0
	v163 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+560)) = v163
	v168 = v163
	goto L16
L20:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v52+v106<<(uint(int32(2))%32))))
	v147 = v25 + int32(560)
	v148 = F_FunctionCall1Coll(m, v138, v145, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v148 == v147 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+560)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+564)) = v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+568)) = v155
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+572)) = uint16(v157)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+14)))
	v168 = v159
	goto L16
L23:
	;
	v219 = v25 + int32(576)
	v222 = v219 + v106<<(uint(int32(2))%32)
	v223 = *(*float32)(unsafe.Add(mBase, uint32(v222)))
	if base.F32_lt(v223, float32(0))|base.F32_lt(v217, v223) != 0 {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	if v175&v127 != 0 {
		goto L36
	} else {
		goto L37
	}
L25:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+10)))
	if v184 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v52+v106<<(uint(int32(2))%32))))
	v198 = F_FunctionCall3Coll(m, v180, v188, v25+int32(560), v25+int32(48)+v106<<(uint(int32(4))%32), v25+int32(716))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v200 = float32(0)
	v202 = *(*float32)(unsafe.Add(mBase, uint32(v25)+716))
	if base.F32_lt(v202, v200) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v205 = v200
	goto L32
L31:
	;
	v205 = v202
	goto L32
L32:
	;
	if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v202)&int32(2147483647)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v211 = v200
	goto L35
L34:
	;
	v211 = v205
	goto L35
L35:
	;
	v217 = v211
	goto L23
L36:
	;
	v215 = float32(0)
	goto L38
L37:
	;
	v215 = math.Float32frombits(uint32(0x7f800000))
	goto L38
L38:
	;
	v217 = v215
	goto L23
L39:
	;
	goto L13
L40:
	;
	v250 = base.B2i32(base.F32_gt(v217, float32(0)) == int32(0)) & v108
	if v121 < v243 {
		v104 = v244
		v106 = v121
		v107 = v245
		v108 = v250
		goto L12
	} else {
		goto L48
	}
L41:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v222))) = v217
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v231 = int32(*(*int16)(unsafe.Add(mBase, uint32(v230)+10)))
	if v106 < v231-int32(1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if base.F32_ne(v217, v223) != 0 {
		goto L39
	} else {
		goto L47
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121<<(uint(int32(2))%32)+v219))) = int32(-1082130432)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v243 = v231
	v244 = int32(-1)
	v245 = v66
	goto L40
L47:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241)+10)))
	v243 = v242
	v244 = v104
	v245 = v107
	goto L40
L48:
	;
	v260 = v243
	v261 = v244
	v263 = v121
	v264 = v245
	v265 = v250
	goto L8
L49:
	;
	if v261 == int32(-1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v294 = v261
	v295 = v264
	goto L51
L51:
	;
	if v265 != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v288 = Fn13966(m, int64(63))
	mBase = m.M
	goto L55
L53:
	;
	v289 = v261
	goto L54
L54:
	;
	if v289 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v289 = v288
	goto L54
L56:
	;
	v290 = v264
	goto L58
L57:
	;
	v290 = v66
	goto L58
L58:
	;
	if v289 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v293 = int32(1)
	goto L61
L60:
	;
	v293 = int32(-1)
	goto L61
L61:
	;
	v294 = v293
	v295 = v290
	goto L51
L62:
	;
	if v294 == int32(-1) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v304 = v294
	goto L64
L64:
	;
	v306 = v66 + int32(1)
	if base.Ui32(v306&int32(_a_F_gistchoose_0)) <= base.Ui32(v46) {
		v65 = v304
		v66 = v306
		v68 = v295
		goto L6
	} else {
		goto L70
	}
L65:
	;
	v299 = Fn13966(m, int64(63))
	mBase = m.M
	goto L68
L66:
	;
	v300 = v294
	goto L67
L67:
	;
	if v300 == int32(1) {
		v319 = v295
		goto L3
	} else {
		goto L69
	}
L68:
	;
	v300 = v299
	goto L67
L69:
	;
	v304 = int32(0)
	goto L64
L70:
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v80 float64
	_ = v80
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
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
			if base.Ui32(int32(2)) <= base.Ui32(v41) {
				v45 = F_log(m, base.F64_convert_i32_u(v41))
				mBase = m.M
				v49 = base.I32_trunc_sat_f64_s(base.F64_div(v45, float64(4.605170185988092)))
			} else {
				v49 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v49
			v52 = v49
		} else {
			v52 = v37
		}
		v54 = *(*float64)(unsafe.Add(mBase, _c_F_gistcostestimate[0]))
		v55 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
		v56 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
		if base.F64_gt(v56, float64(1)) == int32(0) {
			v61 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
			v62 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
			v71 = v55
			v72 = v61
			v74 = v62
		} else {
			v63 = F_log(m, v56)
			mBase = m.M
			v65 = base.F64_mul(base.F64_ceil(v63), v54)
			v67 = *(*float64)(unsafe.Add(mBase, uint32(v16)+56))
			v69 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
			v71 = base.F64_add(v55, v65)
			v72 = v67
			v74 = base.F64_add(base.F64_mul(v67, v65), v69)
		}
		v80 = base.F64_mul(v54, base.F64_mul(base.F64_convert_i32_s(v52+int32(1)), float64(50)))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v71, v80)
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(v72, v80), v74)
		v86 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v86
		v88 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v88
		v90 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v90
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
	v22 = int32(base.Ui32(v13+int32(_a_F_gistfillbuffer_0))>>(uint(int32(2))%32)) + int32(1)
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
	v40 = v38 & int32(_a_F_gistfillbuffer_1)
	v44 = F_PageAddItemExtended(m, l0, v37, v40, v30&int32(_a_F_gistfillbuffer_2), int32(0))
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
	F_errmsg_internal(m, int32(_a_F_gistfillbuffer_3), v10)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_gistfillbuffer_4), int32(50), int32(_a_F_gistfillbuffer_5))
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_gistgetbitmap[0]))) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v42
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_gistgetbitmap[1])))
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
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_gistgetbitmap[2]))
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
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_gistinserttuple[0]))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(v14^int32(-1))<<(uint(int32(6))%32))+16))
		v33 = v24
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_gistinserttuple[1]))
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
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_repalloc(m, l0, (v6+l3)<<(uint(int32(2))%32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = l3 << (uint(int32(2)) % 32)
		if v15 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			base.MemoryCopy(m, v10+v16<<(uint(int32(2))%32), l2, v15)
		} else {
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21 + l3
		return v10
	}
}
func F_gseg_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 float32
	_ = v51
	var v52 float32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 float32
	_ = v64
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v18 = (v14 - int32(1)) & int32(_a_F_gseg_picksplit_0)
	v21 = F_palloc(m, v18*int32(12))
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
	if v14&int32(_a_F_gseg_picksplit_0) != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v29 = int32(1)
	if base.Ui32(v18) <= base.Ui32(v29) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_pg_qsort(m, v21, v18, int32(12), int32(_a_F_gseg_picksplit_1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	v32 = v29
	goto L8
L7:
	;
	v32 = v18
	goto L8
L8:
	;
	v36 = int32(1)
	goto L9
L9:
	;
	v47 = int32(4)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(4)+v36<<(uint(v47)%32))))
	v51 = *(*float32)(unsafe.Add(mBase, uint32(v50)))
	v52 = *(*float32)(unsafe.Add(mBase, uint32(v50)+4))
	v53 = int32(12)
	v55 = v21 + v36*v53
	*(*int32)(unsafe.Add(mBase, uint32(v55-v47))) = v50
	*(*uint16)(unsafe.Add(mBase, uint32(v55-int32(8)))) = uint16(v36)
	v64 = float32(0.5)
	*(*float32)(unsafe.Add(mBase, uint32(v55-v53))) = base.F32_add(base.F32_mul(v51, v64), base.F32_mul(v52, v64))
	if v36 != v32 {
		v36 = v36 + int32(1)
		goto L9
	} else {
		goto L11
	}
L10:
	;
	goto L5
L11:
	;
	goto L10
L12:
	;
	v88 = int32(1)
	v90 = v18 << (uint(v88) % 32)
	v91 = F_palloc(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v91
	v94 = F_palloc(m, v90)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v94
	v97 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v97
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v103 = F_palloc(m, int32(12))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	*(*int64)(unsafe.Add(mBase, uint32(v103))) = v108
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v101))) = uint16(v110)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v113 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v112 + v113
	v117 = int32(base.Ui32(v18) >> (uint(v113) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v18) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v122 = v88
	v123 = v101
	v127 = v103
	goto L19
L17:
	;
	v157 = v103
	goto L18
L18:
	;
	v162 = F_palloc(m, int32(12))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v135 = v21 + v122*int32(12)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v137 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_picksplit_2), int32(0), v127, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v157 = v137
	goto L18
L21:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v123)+2)) = uint16(v139)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v142 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v141 + v142
	v148 = v122 + v142
	if v148 != v117 {
		v122 = v148
		v123 = v123 + int32(2)
		v127 = v137
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v166 = v21 + v117*int32(12)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v167)))
	*(*int64)(unsafe.Add(mBase, uint32(v162))) = v170
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v94))) = uint16(v172)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v175 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v174 + v175
	v179 = v117 + v175
	if base.Ui32(v179) < base.Ui32(v18) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v181 = v94
	v183 = v162
	v184 = v179
	goto L27
L25:
	;
	v213 = v162
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v157
	return v12
L27:
	;
	v196 = v21 + v184*int32(12)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	v198 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_picksplit_2), int32(0), v183, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v213 = v198
	goto L26
L29:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v181)+2)) = uint16(v200)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v203 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v202 + v203
	v209 = v184 + v203
	if v209 != v18 {
		v181 = v181 + int32(2)
		v183 = v198
		v184 = v209
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
func F_gtsquery_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v127 int32
	_ = v127
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v265 int64
	_ = v265
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v290 int32
	_ = v290
	var v313 int32
	_ = v313
	var v315 int64
	_ = v315
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v324 int64
	_ = v324
	var v332 int32
	_ = v332
	var v345 int32
	_ = v345
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v445 int64
	_ = v445
	var v446 int64
	_ = v446
	var v449 int32
	_ = v449
	var v450 int64
	_ = v450
	var v471 int32
	_ = v471
	var v494 int32
	_ = v494
	var v496 int64
	_ = v496
	var v499 int64
	_ = v499
	var v503 int64
	_ = v503
	var v511 int32
	_ = v511
	var v524 int32
	_ = v524
	var v547 int32
	_ = v547
	var v549 int64
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v583 int64
	_ = v583
	var v584 int64
	_ = v584
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v606 int64
	_ = v606
	var v607 int64
	_ = v607
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	v7 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v27 = (v23 + int32(_a_F_gtsquery_picksplit_0)) & int32(_a_F_gtsquery_picksplit_1)
	v31 = v27<<(uint(int32(1))%32) + int32(4)
	v32 = F_palloc(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v21))) = v32
		v37 = F_palloc(m, v31)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v39
			*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v37
			*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v39
			if base.Ui32(int32(2)) <= base.Ui32(v27) {
				v47 = v22 + int32(4)
				v57 = int32(1)
				v58 = int32(-1)
				v66 = v7
				v67 = v7
				for {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v47+v57<<(uint(int32(4))%32))))
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)))
					v76 = v57 + int32(1)
					v77 = v76
					v85 = v58
					v86 = v76
					v93 = v66
					v94 = v67
					for {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v47+v77<<(uint(int32(4))%32))))
						v101 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
						v102 = v74 ^ v101
						v105 = int32(0)
						v106 = int64(0)
						for {
							v127 = int32(1)
							v150 = base.I32_wrap_i64(int64(base.Ui64(v102)>>(uint(v106)%64)))&v127 + v105 + base.I32_wrap_i64(int64(base.Ui64(v102)>>(uint(v106|int64(1))%64)))&v127 + base.I32_wrap_i64(int64(base.Ui64(v102)>>(uint(v106|int64(2))%64)))&v127 + base.I32_wrap_i64(int64(base.Ui64(v102)>>(uint(v106|int64(3))%64)))&v127
							v152 = v106 + int64(4)
							if v152 != int64(64) {
								v105 = v150
								v106 = v152
								continue
							} else {
								break
							}
							break
						}
						v155 = base.B2i32(v85 < v150)
						if v85 < v150 {
							v156 = v86
						} else {
							v156 = v93
						}
						if v85 < v150 {
							v157 = v57
						} else {
							v157 = v94
						}
						if v85 < v150 {
							v158 = v150
						} else {
							v158 = v85
						}
						v160 = v86 + int32(1)
						v162 = v160 & int32(_a_F_gtsquery_picksplit_1)
						if base.Ui32(v162) <= base.Ui32(v27) {
							v77 = v162
							v85 = v158
							v86 = v160
							v93 = v156
							v94 = v157
							continue
						} else {
							break
						}
						break
					}
					if v76 != v27 {
						v57 = v76
						v58 = v158
						v66 = v156
						v67 = v157
						continue
					} else {
						break
					}
					break
				}
				v181 = v156
				v182 = v157
			} else {
				v181 = v7
				v182 = v7
			}
			v186 = v22 + int32(4)
			v188 = int32(_a_F_gtsquery_picksplit_1)
			v190 = int32(0)
			v196 = base.B2i32(v182&v188 == v190) | base.B2i32(v181&v188 == v190)
			if v196 != 0 {
				v197 = int32(2)
			} else {
				v197 = v181
			}
			v202 = v186 + v197&int32(_a_F_gtsquery_picksplit_1)<<(uint(int32(4))%32)
			v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
			v204 = *(*int64)(unsafe.Add(mBase, uint32(v203)))
			if v196 != 0 {
				v206 = int32(1)
			} else {
				v206 = v182
			}
			v207 = int32(_a_F_gtsquery_picksplit_1)
			v211 = v186 + v206&v207<<(uint(int32(4))%32)
			v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
			v213 = *(*int64)(unsafe.Add(mBase, uint32(v212)))
			v215 = v23 + v207
			v217 = v215 & v207
			v220 = F_palloc(m, v217<<(uint(int32(3))%32))
			mBase = m.M
			v221 = m.ExcPending
			if v221 != 0 {
				return int32(0)
			} else {
				if v23&int32(_a_F_gtsquery_picksplit_1) == int32(1) {
					F_pg_qsort(m, v220, v217, int32(8), int32(1511))
					mBase = m.M
					v229 = m.ExcPending
					if v229 != 0 {
						return int32(0)
					} else {
						v606 = v204
						v607 = v213
						v613 = v37
						v615 = v32
						v623 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v615))) = uint16(v623)
						*(*uint16)(unsafe.Add(mBase, uint32(v613))) = uint16(v623)
						v627 = F_Int64GetDatum(m, v607)
						mBase = m.M
						v628 = m.ExcPending
						if v628 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v627
							v630 = F_Int64GetDatum(m, v606)
							mBase = m.M
							v631 = m.ExcPending
							if v631 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v630
								return v21
							}
						}
					}
				} else {
					v230 = int32(1)
					v232 = v230
					v240 = v230
					for {
						v254 = v220 + v232<<(uint(int32(3))%32)
						*(*uint16)(unsafe.Add(mBase, uint32(v254-int32(8)))) = uint16(v240)
						v261 = *(*int32)(unsafe.Add(mBase, uint32(v186+v232<<(uint(int32(4))%32))))
						v262 = *(*int64)(unsafe.Add(mBase, uint32(v261)))
						v263 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
						v264 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
						v265 = v262 ^ v264
						v268 = int32(0)
						v269 = int64(0)
						for {
							v290 = int32(1)
							v313 = base.I32_wrap_i64(int64(base.Ui64(v265)>>(uint(v269)%64)))&v290 + v268 + base.I32_wrap_i64(int64(base.Ui64(v265)>>(uint(v269|int64(1))%64)))&v290 + base.I32_wrap_i64(int64(base.Ui64(v265)>>(uint(v269|int64(2))%64)))&v290 + base.I32_wrap_i64(int64(base.Ui64(v265)>>(uint(v269|int64(3))%64)))&v290
							v315 = v269 + int64(4)
							if v315 != int64(64) {
								v268 = v313
								v269 = v315
								continue
							} else {
								break
							}
							break
						}
						v318 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
						v319 = *(*int64)(unsafe.Add(mBase, uint32(v318)))
						v320 = v319 ^ v262
						v324 = int64(0)
						v332 = int32(0)
						for {
							v345 = int32(1)
							v368 = base.I32_wrap_i64(int64(base.Ui64(v320)>>(uint(v324)%64)))&v345 + v332 + base.I32_wrap_i64(int64(base.Ui64(v320)>>(uint(v324|int64(1))%64)))&v345 + base.I32_wrap_i64(int64(base.Ui64(v320)>>(uint(v324|int64(2))%64)))&v345 + base.I32_wrap_i64(int64(base.Ui64(v320)>>(uint(v324|int64(3))%64)))&v345
							v370 = v324 + int64(4)
							if v370 != int64(64) {
								v324 = v370
								v332 = v368
								continue
							} else {
								break
							}
							break
						}
						v375 = v313 - v368
						v377 = v375 >> (uint(int32(31)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(v254-int32(4)))) = v375 ^ v377 - v377
						v382 = v240 + int32(1)
						v383 = int32(_a_F_gtsquery_picksplit_1)
						v384 = v382 & v383
						if base.Ui32(v384) <= base.Ui32(v215&v383) {
							v232 = v384
							v240 = v382
							continue
						} else {
							break
						}
						break
					}
					F_pg_qsort(m, v220, v217, int32(8), int32(1511))
					mBase = m.M
					v391 = m.ExcPending
					if v391 != 0 {
						return int32(0)
					} else {
						v392 = int32(1)
						if base.Ui32(v217) <= base.Ui32(v392) {
							v395 = v392
						} else {
							v395 = v217
						}
						v402 = v204
						v403 = v213
						v407 = int32(0)
						v409 = v37
						v411 = v32
						for {
							v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220+v407<<(uint(int32(3))%32)))))
							if v206&int32(_a_F_gtsquery_picksplit_1) == v422 {
								*(*uint16)(unsafe.Add(mBase, uint32(v411))) = uint16(v206)
								v425 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v425 + int32(1)
								v583 = v402
								v584 = v403
								v590 = v409
								v592 = v411 + int32(2)
							} else {
								if v197&int32(_a_F_gtsquery_picksplit_1) == v422 {
									*(*uint16)(unsafe.Add(mBase, uint32(v409))) = uint16(v197)
									v435 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v435 + int32(1)
									v583 = v402
									v584 = v403
									v590 = v409 + int32(2)
									v592 = v411
								} else {
									v444 = *(*int32)(unsafe.Add(mBase, uint32(v186+v422<<(uint(int32(4))%32))))
									v445 = *(*int64)(unsafe.Add(mBase, uint32(v444)))
									v446 = v445 ^ v403
									v449 = int32(0)
									v450 = int64(0)
									for {
										v471 = int32(1)
										v494 = base.I32_wrap_i64(int64(base.Ui64(v446)>>(uint(v450)%64)))&v471 + v449 + base.I32_wrap_i64(int64(base.Ui64(v446)>>(uint(v450|int64(1))%64)))&v471 + base.I32_wrap_i64(int64(base.Ui64(v446)>>(uint(v450|int64(2))%64)))&v471 + base.I32_wrap_i64(int64(base.Ui64(v446)>>(uint(v450|int64(3))%64)))&v471
										v496 = v450 + int64(4)
										if v496 != int64(64) {
											v449 = v494
											v450 = v496
											continue
										} else {
											break
										}
										break
									}
									v499 = v402 ^ v445
									v503 = int64(0)
									v511 = int32(0)
									for {
										v524 = int32(1)
										v547 = base.I32_wrap_i64(int64(base.Ui64(v499)>>(uint(v503)%64)))&v524 + v511 + base.I32_wrap_i64(int64(base.Ui64(v499)>>(uint(v503|int64(1))%64)))&v524 + base.I32_wrap_i64(int64(base.Ui64(v499)>>(uint(v503|int64(2))%64)))&v524 + base.I32_wrap_i64(int64(base.Ui64(v499)>>(uint(v503|int64(3))%64)))&v524
										v549 = v503 + int64(4)
										if v549 != int64(64) {
											v503 = v549
											v511 = v547
											continue
										} else {
											break
										}
										break
									}
									v554 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
									v555 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
									v556 = v554 - v555
									if base.F64_lt(base.F64_convert_i32_s(v494), base.F64_add(base.F64_convert_i32_s(v547), base.F64_mul(base.F64_convert_i32_s(v556*v556*v556), float64(-0.05)))) != 0 {
										*(*uint16)(unsafe.Add(mBase, uint32(v411))) = uint16(v422)
										v565 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v565 + int32(1)
										v583 = v402
										v584 = v403 | v445
										v590 = v409
										v592 = v411 + int32(2)
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v409))) = uint16(v422)
										v573 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v573 + int32(1)
										v583 = v402 | v445
										v584 = v403
										v590 = v409 + int32(2)
										v592 = v411
									}
								}
							}
							v601 = v407 + int32(1)
							if v601 != v395 {
								v402 = v583
								v403 = v584
								v407 = v601
								v409 = v590
								v411 = v592
								continue
							} else {
								break
							}
							break
						}
						v606 = v583
						v607 = v584
						v613 = v590
						v615 = v592
						v623 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v615))) = uint16(v623)
						*(*uint16)(unsafe.Add(mBase, uint32(v613))) = uint16(v623)
						v627 = F_Int64GetDatum(m, v607)
						mBase = m.M
						v628 = m.ExcPending
						if v628 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v627
							v630 = F_Int64GetDatum(m, v606)
							mBase = m.M
							v631 = m.ExcPending
							if v631 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v630
								return v21
							}
						}
					}
				}
			}
		}
	}
}
