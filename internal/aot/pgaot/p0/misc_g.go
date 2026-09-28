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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(33), base.I64_extend_i32_u(l0))
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
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_GetFdwRoutineByRelId_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetFdwRoutineByRelId_1), int32(407), int32(_a_F_GetFdwRoutineByRelId_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30)+4))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_GetFdwRoutineByServerId(m, v32)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v35
				}
			}
		}
	}
}
func F_GetPrivateRefCountEntrySlow(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var __phi141 int32
	_ = __phi141
	var v142 int32
	_ = v142
	var __phi142 int32
	_ = __phi142
	var v143 int32
	_ = v143
	var __phi143 int32
	_ = __phi143
	var v145 int32
	_ = v145
	var __phi145 int32
	_ = __phi145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v221 int32
	_ = v221
	v3 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[0]))
	if v19 != l0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(-1)
	goto L3
L2:
	;
	v21 = v3
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[1]))
	if v23 == l0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = int32(1)
	goto L6
L5:
	;
	v25 = v21
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[2]))
	if v27 == l0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v29 = int32(2)
	goto L9
L8:
	;
	v29 = v25
	goto L9
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[3]))
	if v31 == l0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = int32(3)
	goto L12
L11:
	;
	v33 = v29
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[4]))
	if v35 == l0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = int32(4)
	goto L15
L14:
	;
	v37 = v33
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[5]))
	if v39 == l0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v41 = int32(5)
	goto L18
L17:
	;
	v41 = v37
	goto L18
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[6]))
	if v43 == l0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v45 = int32(6)
	goto L21
L20:
	;
	v45 = v41
	goto L21
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[7]))
	if v47 == l0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v49 = int32(7)
	goto L24
L23:
	;
	v49 = v45
	goto L24
L24:
	;
	if v49 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[8])) = v49
	return v49<<(uint(int32(4))%32) + int32(_a_F_GetPrivateRefCountEntrySlow_0)
L26:
	;
	goto L27
L27:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[9]))
	if v60 == int32(0) {
		v221 = v3
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return v221
L29:
	;
	v63 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[10]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v70 = int32(16)
	v74 = (int32(base.Ui32(l0)>>(uint(v70)%32)) ^ l0) * int32(-2048144789)
	v79 = (int32(base.Ui32(v74)>>(uint(int32(13))%32)) ^ v74) * int32(-1028477387)
	v83 = v69 & (int32(base.Ui32(v79)>>(uint(v70)%32)) ^ v79)
	v86 = v68 + v83<<(uint(int32(4))%32)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)))
	if v87 == v63 {
		v116 = v63
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.B2i32(l1 == v63)|base.B2i32(v116 == int32(0)) != 0 {
		v221 = v116
		goto L28
	} else {
		goto L37
	}
L31:
	;
	v94 = v86
	v96 = v83
	goto L33
L32:
	;
	v116 = v94
	goto L30
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v98 == l0 {
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v116 = int32(0)
	goto L30
L35:
	;
	v102 = (v96 + int32(1)) & v69
	v105 = v68 + v102<<(uint(int32(4))%32)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+4)))
	if v106 != 0 {
		v94 = v105
		v96 = v102
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v116)+8))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v121 - v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v128 = int32(4)
	v132 = v126 & ((v116-v125)>>(uint(v128)%32) + v122)
	v135 = v125 + v132<<(uint(v128)%32)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+4)))
	if v136 != v122 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)) = uint8(v186)
	v188 = int32(_a_F_GetPrivateRefCountEntrySlow_1)
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[9])) = v190 - int32(1)
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L48
	} else {
		goto L49
	}
L39:
	;
	v179 = v116
	goto L38
L40:
	;
	goto L41
L41:
	;
	__phi141 = v135
	__phi142 = v116
	__phi143 = v132
	__phi145 = v126
	v141 = __phi141
	v142 = __phi142
	v143 = __phi143
	v145 = __phi145
	goto L42
L42:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v148 = int32(16)
	v152 = (int32(base.Ui32(v147)>>(uint(v148)%32)) ^ v147) * int32(-2048144789)
	v157 = (int32(base.Ui32(v152)>>(uint(int32(13))%32)) ^ v152) * int32(-1028477387)
	if v143 == (int32(base.Ui32(v157)>>(uint(v148)%32))^v157)&v145 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v179 = v141
	goto L38
L44:
	;
	v179 = v142
	goto L38
L45:
	;
	goto L46
L46:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v141)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v142)+8)) = v163
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	*(*int64)(unsafe.Add(mBase, uint32(v142))) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v169 = int32(1)
	v171 = v168 & (v143 + v169)
	v174 = v167 + v171<<(uint(int32(4))%32)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
	if v175 == v169 {
		__phi141 = v174
		__phi142 = v141
		__phi143 = v171
		__phi145 = v168
		v141 = __phi141
		v142 = __phi142
		v143 = __phi143
		v145 = __phi145
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	return int32(0)
L49:
	;
	v198 = int32(_a_F_GetPrivateRefCountEntrySlow_2)
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[11]))
	v201 = v199 << (uint(int32(4)) % 32)
	*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_GetPrivateRefCountEntrySlow[12]))) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_GetPrivateRefCountEntrySlow[13]))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v199<<(uint(int32(2))%32))+uint32(_c_F_GetPrivateRefCountEntrySlow[0]))) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[11])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntrySlow[8])) = v199
	v221 = v201 + int32(_a_F_GetPrivateRefCountEntrySlow_0)
	goto L28
}
func F_GetPublicationsStr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if l2 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v31 < int32(2) {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v18
	F_appendStringInfo(m, l1, int32(_a_F_GetPublicationsStr_0), v9+int32(16))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v27 = F_quote_literal_cstr(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return
L9:
	;
	goto L4
L10:
	;
	F_appendStringInfoString(m, l1, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v38 = int32(1)
	goto L13
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v38<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L1
L15:
	;
	v59 = v38 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v59 < v60 {
		v38 = v59
		goto L13
	} else {
		goto L23
	}
L16:
	;
	F_appendStringInfoString(m, l1, int32(_a_F_GetPublicationsStr_1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
	F_appendStringInfo(m, l1, int32(_a_F_GetPublicationsStr_2), v9)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L22
	}
L19:
	;
	v50 = F_quote_literal_cstr(m, v46)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoString(m, l1, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	goto L15
L23:
	;
	goto L14
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
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
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
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v16 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = v3
	v29 = v3
	goto L7
L5:
	;
	v260 = v3
	goto L6
L6:
	;
	return v260
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v25<<(uint(int32(2))%32))))
	v38 = int64(0)
	v40 = F_SearchSysCacheList(m, int32(53), int32(1), base.I64_extend_i32_u(v36), v38, v38)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v260 = v243
	goto L6
L9:
	;
	F_ReleaseCatCacheList(m, v40)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L22
	}
L10:
	;
	return int32(0)
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+56))
	if v44 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v87 = int32(0)
	goto L9
L13:
	;
	goto L14
L14:
	;
	v50 = int32(0)
	v55 = v50
	v56 = v44
	v57 = v50
	goto L15
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(-64)+v55<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
	v69 = v67 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+12)))
	if v70 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v87 = v78
	goto L9
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v74 = F_lappend_oid(m, v57, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v77 = v56
	v78 = v57
	goto L19
L19:
	;
	v80 = v55 + int32(1)
	if v80 < v77 {
		v55 = v80
		v56 = v77
		v57 = v78
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v40)+56))
	v77 = v76
	v78 = v74
	goto L19
L21:
	;
	goto L16
L22:
	;
	v96 = v25 + int32(1)
	v97 = int32(0)
	if v87 == v97 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	F_list_free(m, v87)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L10
	} else {
		goto L68
	}
L24:
	;
	v234 = int32(0)
	v243 = v36
	goto L23
L25:
	;
	if v135 != 0 {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	v135 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v103 <= int32(0) {
		v129 = v97
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v135 = v129
	goto L25
L30:
	;
	v106 = int32(0)
	if v106 < v103 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v109 = v103
	goto L33
L32:
	;
	v109 = v106
	goto L33
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v112 = int32(0)
	goto L34
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v110+v112<<(uint(int32(2))%32))))
	v121 = base.B2i32(v120 == l0)
	if v120 == l0 {
		v129 = v121
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v129 = v121
	goto L29
L36:
	;
	v123 = v112 + int32(1)
	if v123 != v109 {
		v112 = v123
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L24
L39:
	;
	goto L40
L40:
	;
	v139 = F_get_rel_namespace(m, v36)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L10
	} else {
		goto L42
	}
L41:
	;
	F_ReleaseCatCacheList(m, v144)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L51
	}
L42:
	;
	v142 = int64(0)
	v144 = F_SearchSysCacheList(m, int32(50), int32(1), base.I64_extend_i32_u(v139), v142, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v144)+56))
	if v146 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v181 = int32(0)
	goto L41
L45:
	;
	goto L46
L46:
	;
	v152 = int32(0)
	v156 = v152
	v157 = v152
	goto L47
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144-int32(-64)+v157<<(uint(int32(2))%32))))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+72))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+22)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169+v170)+4))
	v173 = F_lappend_oid(m, v156, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L49
	}
L48:
	;
	v181 = v173
	goto L41
L49:
	;
	v176 = v157 + int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v144)+56))
	if v176 < v177 {
		v156 = v173
		v157 = v176
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v192 = int32(0)
	if v181 == v192 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v230 != 0 {
		goto L65
	} else {
		goto L66
	}
L53:
	;
	v230 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v198 <= int32(0) {
		v224 = v192
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v230 = v224
	goto L52
L57:
	;
	v201 = int32(0)
	if v201 < v198 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v204 = v198
	goto L60
L59:
	;
	v204 = v201
	goto L60
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v207 = int32(0)
	goto L61
L61:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v205+v207<<(uint(int32(2))%32))))
	v216 = base.B2i32(v215 == l0)
	if v215 == l0 {
		v224 = v216
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v224 = v216
	goto L56
L63:
	;
	v218 = v207 + int32(1)
	if v218 != v204 {
		v207 = v218
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v231 = v36
	goto L67
L66:
	;
	v231 = v29
	goto L67
L67:
	;
	v234 = v181
	v243 = v231
	goto L23
L68:
	;
	F_list_free(m, v234)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v96 < v248 {
		v25 = v96
		v29 = v243
		goto L7
	} else {
		goto L70
	}
L70:
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
func F_GetUserMappingExtended(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = base.I64_extend_i32_u(l1)
	v15 = F_SearchSysCache2(m, int32(84), base.I64_extend_i32_u(l0), v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			v49 = v15
			v51 = F_palloc(m, int32(16))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54)))
				*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
				v64 = F_SysCacheGetAttr(m, int32(84), v49, int32(4), v10+int32(15))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v66 != 0 {
						v70 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v70
						F_ReleaseCatCache(m, v49)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					} else {
						v68 = F_untransformRelOptions(m, v64)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v70 = v68
							*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v70
							F_ReleaseCatCache(m, v49)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v19 = F_SearchSysCache2(m, int32(84), int64(0), v14)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if v19 != 0 {
					v49 = v19
					v51 = F_palloc(m, int32(16))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54)))
						*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
						v64 = F_SysCacheGetAttr(m, int32(84), v49, int32(4), v10+int32(15))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v66 != 0 {
								v70 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v70
								F_ReleaseCatCache(m, v49)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								v68 = F_untransformRelOptions(m, v64)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = v68
									*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v70
									F_ReleaseCatCache(m, v49)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						}
					}
				} else {
					v22 = F_GetForeignServerExtended(m, l1, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v26 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							if v26 == int32(0) {
								m.G0 = v10 + int32(16)
								return
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									if l0 != 0 {
										v34 = F_GetUserNameFromId(m, l0, int32(0))
										mBase = m.M
										v35 = m.ExcPending
										if v35 != 0 {
											return
										} else {
											v37 = v34
											v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v38
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v37
											F_errmsg(m, int32(_a_F_GetUserMappingExtended_0), v10)
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_GetUserMappingExtended_1), int32(268), int32(_a_F_GetUserMappingExtended_2))
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									} else {
										v37 = int32(_a_F_GetUserMappingExtended_3)
										v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v38
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v37
										F_errmsg(m, int32(_a_F_GetUserMappingExtended_0), v10)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_GetUserMappingExtended_1), int32(268), int32(_a_F_GetUserMappingExtended_2))
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return
											} else {
												m.G0 = v10 + int32(16)
												return
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
func F_GetVirtualXIDsDelayingChkpt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v3 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_GetVirtualXIDsDelayingChkpt[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = F_palloc_mul(m, int32(8), v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_GetVirtualXIDsDelayingChkpt[1]))
		v25 = F_LWLockAcquire(m, v21+int32(512), int32(1))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if int32(0) < v27 {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_GetVirtualXIDsDelayingChkpt[2]))
				v36 = v27
				v39 = v3
				v41 = v3
				for {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(36)+v41<<(uint(int32(2))%32))))
					v51 = v33 + v48*int32(768)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+336))
					if v52&l1 == int32(0) {
						v68 = v36
						v70 = v39
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+44))
						if v56 == int32(0) {
							v68 = v36
							v70 = v39
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
							v62 = v16 + v39<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v59
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
							v68 = v67
							v70 = v39 + int32(1)
						}
					}
					v73 = v41 + int32(1)
					if v73 < v68 {
						v36 = v68
						v39 = v70
						v41 = v73
						continue
					} else {
						break
					}
					break
				}
				v80 = v70
			} else {
				v80 = v3
			}
			v87 = *(*int32)(unsafe.Add(mBase, _c_F_GetVirtualXIDsDelayingChkpt[1]))
			F_LWLockRelease(m, v87+int32(512))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80
				return v16
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	v2 = int32(0)
	if l0 == v2 {
		v63 = v2
		return v63
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+117)))
		if v7 != 0 {
			v63 = v2
			return v63
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[0])))
			if v10 == int32(1) {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[1]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+308))
				v18 = base.B2i32(v16 != int32(2))
				*(*uint8)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[0])) = uint8(v18)
				v20 = v18
			} else {
				v20 = int32(0)
			}
			if v20 != 0 {
				v63 = v2
				return v63
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if base.Ui32(v21) < base.Ui32(int32(_a_F_GlobalVisHorizonKindForRel_0)) {
					return int32(1)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[2]))
					if v27 <= int32(1) {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GlobalVisHorizonKindForRel[3])))
						if v31&int32(1) == int32(0) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v58 != 0 {
								v63 = int32(3)
								return v63
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v59 != 0 {
									v63 = int32(3)
									return v63
								} else {
									return int32(2)
								}
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+118)))
							if v37 != int32(112) {
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v58 != 0 {
									v63 = int32(3)
									return v63
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v59 != 0 {
										v63 = int32(3)
										return v63
									} else {
										return int32(2)
									}
								}
							} else {
								if v27 <= int32(0) {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v42 != 0 {
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v58 != 0 {
											v63 = int32(3)
											return v63
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v59 != 0 {
												v63 = int32(3)
												return v63
											} else {
												return int32(2)
											}
										}
									} else {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v43 != 0 {
											v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v58 != 0 {
												v63 = int32(3)
												return v63
											} else {
												v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v59 != 0 {
													v63 = int32(3)
													return v63
												} else {
													return int32(2)
												}
											}
										} else {
											v44 = int32(1)
											v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if base.Ui32(v45) < base.Ui32(int32(_a_F_GlobalVisHorizonKindForRel_0)) {
												v63 = v44
												return v63
											} else {
												v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
												if v48 == int32(0) {
													v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v58 != 0 {
														v63 = int32(3)
														return v63
													} else {
														v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v59 != 0 {
															v63 = int32(3)
															return v63
														} else {
															return int32(2)
														}
													}
												} else {
													v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+119)))
													switch v52 - int32(109) {
													case 0, 5:
														v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+112)))
														if v55 != 0 {
															v63 = v44
															return v63
														} else {
															v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															if v58 != 0 {
																v63 = int32(3)
																return v63
															} else {
																v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
																if v59 != 0 {
																	v63 = int32(3)
																	return v63
																} else {
																	return int32(2)
																}
															}
														}
													default:
														v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v58 != 0 {
															v63 = int32(3)
															return v63
														} else {
															v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v59 != 0 {
																v63 = int32(3)
																return v63
															} else {
																return int32(2)
															}
														}
													}
												}
											}
										}
									}
								} else {
									v44 = int32(1)
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if base.Ui32(v45) < base.Ui32(int32(_a_F_GlobalVisHorizonKindForRel_0)) {
										v63 = v44
										return v63
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										if v48 == int32(0) {
											v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v58 != 0 {
												v63 = int32(3)
												return v63
											} else {
												v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v59 != 0 {
													v63 = int32(3)
													return v63
												} else {
													return int32(2)
												}
											}
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+119)))
											switch v52 - int32(109) {
											case 0, 5:
												v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+112)))
												if v55 != 0 {
													v63 = v44
													return v63
												} else {
													v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v58 != 0 {
														v63 = int32(3)
														return v63
													} else {
														v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v59 != 0 {
															v63 = int32(3)
															return v63
														} else {
															return int32(2)
														}
													}
												}
											default:
												v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v58 != 0 {
													v63 = int32(3)
													return v63
												} else {
													v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v59 != 0 {
														v63 = int32(3)
														return v63
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
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+118)))
						if v37 != int32(112) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v58 != 0 {
								v63 = int32(3)
								return v63
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v59 != 0 {
									v63 = int32(3)
									return v63
								} else {
									return int32(2)
								}
							}
						} else {
							if v27 <= int32(0) {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v42 != 0 {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v58 != 0 {
										v63 = int32(3)
										return v63
									} else {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v59 != 0 {
											v63 = int32(3)
											return v63
										} else {
											return int32(2)
										}
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v43 != 0 {
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v58 != 0 {
											v63 = int32(3)
											return v63
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v59 != 0 {
												v63 = int32(3)
												return v63
											} else {
												return int32(2)
											}
										}
									} else {
										v44 = int32(1)
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if base.Ui32(v45) < base.Ui32(int32(_a_F_GlobalVisHorizonKindForRel_0)) {
											v63 = v44
											return v63
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											if v48 == int32(0) {
												v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v58 != 0 {
													v63 = int32(3)
													return v63
												} else {
													v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v59 != 0 {
														v63 = int32(3)
														return v63
													} else {
														return int32(2)
													}
												}
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+119)))
												switch v52 - int32(109) {
												case 0, 5:
													v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+112)))
													if v55 != 0 {
														v63 = v44
														return v63
													} else {
														v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v58 != 0 {
															v63 = int32(3)
															return v63
														} else {
															v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v59 != 0 {
																v63 = int32(3)
																return v63
															} else {
																return int32(2)
															}
														}
													}
												default:
													v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v58 != 0 {
														v63 = int32(3)
														return v63
													} else {
														v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v59 != 0 {
															v63 = int32(3)
															return v63
														} else {
															return int32(2)
														}
													}
												}
											}
										}
									}
								}
							} else {
								v44 = int32(1)
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if base.Ui32(v45) < base.Ui32(int32(_a_F_GlobalVisHorizonKindForRel_0)) {
									v63 = v44
									return v63
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if v48 == int32(0) {
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v58 != 0 {
											v63 = int32(3)
											return v63
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v59 != 0 {
												v63 = int32(3)
												return v63
											} else {
												return int32(2)
											}
										}
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+119)))
										switch v52 - int32(109) {
										case 0, 5:
											v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+112)))
											if v55 != 0 {
												v63 = v44
												return v63
											} else {
												v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v58 != 0 {
													v63 = int32(3)
													return v63
												} else {
													v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v59 != 0 {
														v63 = int32(3)
														return v63
													} else {
														return int32(2)
													}
												}
											}
										default:
											v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v58 != 0 {
												v63 = int32(3)
												return v63
											} else {
												v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v59 != 0 {
													v63 = int32(3)
													return v63
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
}
func F_g_intbig_compress(m *base.Module, l0 int32) int64 {
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
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+32)))
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
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+18)))
	if v40 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	return int64(0)
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = v37
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L54
	}
L13:
	;
	v197 = F_palloc(m, int32(24))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
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
		v192 = v47
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
		v192 = v47
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
		v192 = v47
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
	v192 = v47
	goto L13
L40:
	;
	goto L39
L41:
	;
	return base.I64_extend_i32_u(v10)
L42:
	;
	goto L43
L43:
	;
	v150 = int32(0)
	if v38 <= v150 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v183 = F_palloc(m, int32(8))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L10
	} else {
		goto L52
	}
L45:
	;
	v155 = v150
	goto L46
L46:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+(v39+int32(8))))))
	if v165 == int32(255) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	return base.I64_extend_i32_u(v10)
L48:
	;
	v169 = v155 + int32(1)
	if v38 != v169 {
		v155 = v169
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
	*(*int64)(unsafe.Add(mBase, uint32(v183))) = int64(17179869216)
	v192 = v183
	goto L13
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v197))) = base.I64_extend_i32_u(v192)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+12)) = v203
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+18)) = uint8(v206)
	*(*uint16)(unsafe.Add(mBase, uint32(v197)+16)) = uint16(v205)
	return base.I64_extend_i32_u(v197)
L54:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_g_intbig_compress_0), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_g_intbig_compress_1), int32(157), int32(_a_F_g_intbig_compress_2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
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
func F_g_intbig_consistent(m *base.Module, l0 int32) int64 {
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
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
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return int64(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
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
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+32)))
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
		v422 = v45
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
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L97
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L93
	}
L15:
	;
	return base.I64_extend_i32_u(v422) & int64(1)
L16:
	;
	if v18 == int32(20) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_pfree(m, v13)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L92
	}
L18:
	;
	v57 = F_signconsistent(m, v13, v48+int32(8), v44, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v61 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v13 != v59 {
		v410 = v57
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v422 = v57
	goto L15
L23:
	;
	v62 = F_array_contains_nulls(m, v13)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v64 = int32(0)
	switch v18 - int32(3) {
	case 0:
		goto L32
	default:
		v398 = v64
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
	if v62 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v13 == v406 {
		v422 = v398
		goto L15
	} else {
		goto L91
	}
L29:
	;
	v263 = int32(1)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+16)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264+v265)+12)))
	if v267&v263 == int32(0) {
		v398 = v263
		goto L28
	} else {
		goto L71
	}
L30:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v261 = F__intbig_contains(m, v260, v13, v44)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L70
	}
L31:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v119)+12)))
	if v121&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L32:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v71 = F_ArrayGetNItemsSafe(m, v68, v13+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v73 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v71 == int32(0) {
		v398 = v64
		goto L28
	} else {
		goto L40
	}
L35:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v85 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v83 = F_array_contains_nulls(m, v13)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v83 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v85 = v73
	goto L34
L40:
	;
	v94 = v71
	v98 = v85 + v13
	goto L41
L41:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v104 = base.I32_rem_u_s(v103, v44<<(uint(int32(3))%32))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(8)+int32(base.Ui32(v104)>>(uint(int32(3))%32))))))
	v111 = int32(base.Ui32(v108) >> (uint(v104&int32(7)) % 32))
	if v111&int32(1) != 0 {
		v398 = v111
		goto L28
	} else {
		goto L43
	}
L42:
	;
	v398 = v111
	goto L28
L43:
	;
	v117 = v94 - int32(1)
	if v117 != 0 {
		v94 = v117
		v98 = v98 + int32(4)
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v127 = F_ArrayGetNItemsSafe(m, v124, v13+int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v258 = F__intbig_contains(m, v257, v13, v44)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L69
	}
L48:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v129 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v139 = (v132<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L51
L50:
	;
	v139 = v129
	goto L51
L51:
	;
	v140 = F_palloc0(m, v44)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v127 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v44 <= int32(0) {
		v247 = int32(1)
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v145 = v44 << (uint(int32(3)) % 32)
	v146 = v139 + v13
	if v127&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v150 = base.I32_rem_u_s(v149, v145)
	v153 = v140 + int32(base.Ui32(v150)>>(uint(int32(3))%32))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v155 = int32(1)
	v159 = v154 | v155<<(uint(v150&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v159)
	v165 = v146 + int32(4)
	v168 = v127 - v155
	goto L57
L56:
	;
	v165 = v146
	v168 = v127
	goto L57
L57:
	;
	if v127 == int32(1) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v173 = v165
	v174 = v168
	goto L59
L59:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v182 = base.I32_rem_u_s(v181, v145)
	v183 = int32(3)
	v185 = v140 + int32(base.Ui32(v182)>>(uint(v183)%32))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	v187 = int32(1)
	v188 = int32(7)
	v191 = v186 | v187<<(uint(v182&v188)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v194 = base.I32_rem_u_s(v193, v145)
	v197 = v140 + int32(base.Ui32(v194)>>(uint(v183)%32))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v203 = v198 | v187<<(uint(v194&v188)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v203)
	v208 = v174 - int32(2)
	if v208 != 0 {
		v173 = v173 + int32(8)
		v174 = v208
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
	F_pfree(m, v140)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L68
	}
L63:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v227 = int32(0)
	goto L64
L64:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227+(v222+int32(8))))))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227+v140))))
	v240 = base.B2i32(v237 == v239)
	if v237 != v239 {
		v247 = v240
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v247 = v240
	goto L62
L66:
	;
	v243 = v227 + int32(1)
	if v243 != v44 {
		v227 = v243
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v398 = v247
	goto L28
L69:
	;
	v398 = v258
	goto L28
L70:
	;
	v398 = v261
	goto L28
L71:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v275 = F_ArrayGetNItemsSafe(m, v272, v13+int32(16))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v277 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v287 = (v280<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L75
L74:
	;
	v287 = v277
	goto L75
L75:
	;
	v288 = F_palloc0(m, v44)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v275 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if v44 <= int32(0) {
		v398 = int32(1)
		goto L28
	} else {
		goto L86
	}
L78:
	;
	v293 = v44 << (uint(int32(3)) % 32)
	v294 = v287 + v13
	if v275&int32(1) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v298 = base.I32_rem_u_s(v297, v293)
	v301 = v288 + int32(base.Ui32(v298)>>(uint(int32(3))%32))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v303 = int32(1)
	v307 = v302 | v303<<(uint(v298&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v307)
	v313 = v294 + int32(4)
	v316 = v275 - v303
	goto L81
L80:
	;
	v313 = v294
	v316 = v275
	goto L81
L81:
	;
	if v275 == int32(1) {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v321 = v313
	v322 = v316
	goto L83
L83:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v330 = base.I32_rem_u_s(v329, v293)
	v331 = int32(3)
	v333 = v288 + int32(base.Ui32(v330)>>(uint(v331)%32))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	v335 = int32(1)
	v336 = int32(7)
	v339 = v334 | v335<<(uint(v330&v336)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v333))) = uint8(v339)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v342 = base.I32_rem_u_s(v341, v293)
	v345 = v288 + int32(base.Ui32(v342)>>(uint(v331)%32))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	v351 = v346 | v335<<(uint(v342&v336)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v351)
	v356 = v322 - int32(2)
	if v356 != 0 {
		v321 = v321 + int32(8)
		v322 = v356
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
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v375 = int32(0)
	goto L87
L87:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375+(v370+int32(8))))))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375+v288))))
	v390 = v385 & (v387 ^ int32(255))
	v392 = base.B2i32(v390 == int32(0))
	if v390 != 0 {
		v398 = v392
		goto L28
	} else {
		goto L89
	}
L88:
	;
	v398 = v392
	goto L28
L89:
	;
	v394 = v375 + int32(1)
	if v394 != v44 {
		v375 = v394
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v410 = v398
	goto L17
L92:
	;
	v422 = v410
	goto L15
L93:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_g_intbig_consistent_0), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_g_intbig_consistent_1), int32(491), int32(_a_F_g_intbig_consistent_2))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
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
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_g_intbig_consistent_0), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_g_intbig_consistent_1), int32(80), int32(_a_F_g_intbig_consistent_3))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
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
func F_g_intbig_options(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_g_intbig_options_0), int32(_a_F_g_intbig_options_1), int32(252), int32(1), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		return int64(0)
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
func F_gen_random_uuid(m *base.Module, l0 int32) int64 {
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
		return int64(0)
	} else {
		v8 = m.Env.Pgmem_random_bytes(m, v3, int32(16))
		mBase = m.M
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_gen_random_uuid_0), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_gen_random_uuid_1), int32(554), int32(_a_F_gen_random_uuid_2))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
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
			return base.I64_extend_i32_u(v3)
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
	F_errfinish(m, int32(_a_F_generateSerialExtraStmts_5), int32(492), int32(_a_F_generateSerialExtraStmts_6))
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
	F_errfinish(m, int32(_a_F_generateSerialExtraStmts_5), int32(511), int32(_a_F_generateSerialExtraStmts_6))
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	F_generate_trgm_only(m, v11+int32(4), l0, l1, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v22 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v22)
		if int32(2) <= v20 {
			v27 = v21 + int32(5)
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_generate_trgm[0]))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+272)))
			if v30 != 0 {
				F_trigram_qsort_signed(m, v27, v20)
				mBase = m.M
			} else {
				F_trigram_qsort_unsigned(m, v27, v20)
				mBase = m.M
			}
			v35 = int32(0)
			v36 = int32(1)
			for {
				v43 = int32(3)
				v45 = v27 + v36*v43
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
				v49 = v27 + v35*v43
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
				if v46 != v50 {
					v59 = v35 + int32(1)
					if v59 == v36 {
						v68 = v36
					} else {
						v63 = v27 + v59*int32(3)
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
						*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)) = uint8(v64)
						v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
						*(*uint16)(unsafe.Add(mBase, uint32(v63))) = uint16(v66)
						v68 = v59
					}
				} else {
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
					if v52 != v53 {
						v59 = v35 + int32(1)
						if v59 == v36 {
							v68 = v36
						} else {
							v63 = v27 + v59*int32(3)
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
							*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)) = uint8(v64)
							v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
							*(*uint16)(unsafe.Add(mBase, uint32(v63))) = uint16(v66)
							v68 = v59
						}
					} else {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
						if v55 == v56 {
							v68 = v35
						} else {
							v59 = v35 + int32(1)
							if v59 == v36 {
								v68 = v36
							} else {
								v63 = v27 + v59*int32(3)
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
								*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)) = uint8(v64)
								v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
								*(*uint16)(unsafe.Add(mBase, uint32(v63))) = uint16(v66)
								v68 = v59
							}
						}
					}
				}
				v71 = v36 + int32(1)
				if v71 != v20 {
					v35 = v68
					v36 = v71
					continue
				} else {
					break
				}
				break
			}
			v83 = v68 + int32(1)
		} else {
			v83 = v20
		}
		*(*int32)(unsafe.Add(mBase, uint32(v21))) = v83*int32(12) + int32(20)
		m.G0 = v11 + int32(16)
		return v21
	}
}
func F_getCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(1056)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v11 == v3 {
		v168 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(1056)
	return v168
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v15 == int32(0) {
		v168 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v3
	goto L4
L4:
	;
	v27 = v9 + int32(16)
	F_getNextFlagFromString(m, l0, v9+int32(12), v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v168 = v161
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v32 == int32(2) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+1044)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+1052)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+1048)) = v147
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v157 = F_bsearch(m, v9+int32(1044), v153, v154, int32(12), int32(1266))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L44
	}
L9:
	;
	v35 = F_parseNumericAffixFlag(m, v27)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v39 = F_strlen(m, v9+int32(16))
	mBase = m.M
	v41 = v39 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v41) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v145 = v35
	goto L8
L13:
	;
	v67 = v9 + int32(16)
	if (v67^v63)&int32(3) != 0 {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	v44 = F_palloc0(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v49 = (v39 + int32(8)) & int32(4088)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v49) <= base.Ui32(v50) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v63 = v44
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v57 - v49
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v58 + v49
	v63 = v58
	goto L13
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v57 = v50
	v58 = v52
	goto L18
L20:
	;
	goto L21
L21:
	;
	v53 = int32(_a_F_getCompoundAffixFlagValue_0)
	v55 = F_palloc0(m, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v57 = v53
	v58 = v55
	goto L18
L23:
	;
	v145 = v63
	goto L8
L24:
	;
	goto L23
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v121)
	if v121&int32(255) == int32(0) {
		goto L24
	} else {
		goto L40
	}
L26:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v120 = v67
	v121 = v73
	v122 = v63
	goto L25
L27:
	;
	goto L28
L28:
	;
	if v67&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v77 = v67
	v79 = v63
	goto L32
L30:
	;
	v91 = v67
	v93 = v63
	goto L31
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v98 = int32(-2139062144)
	if (int32(16843008)-v95|v95)&v98 != v98 {
		v120 = v91
		v121 = v95
		v122 = v93
		goto L25
	} else {
		goto L36
	}
L32:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v80)
	if v80 == int32(0) {
		goto L24
	} else {
		goto L34
	}
L33:
	;
	v91 = v87
	v93 = v85
	goto L31
L34:
	;
	v84 = int32(1)
	v85 = v79 + v84
	v87 = v77 + v84
	if v87&int32(3) != 0 {
		v77 = v87
		v79 = v85
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v103 = v91
	v104 = v95
	v105 = v93
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v104
	v107 = int32(4)
	v108 = v105 + v107
	v110 = v103 + v107
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v115 = int32(-2139062144)
	if (int32(16843008)-v112|v112)&v115 == v115 {
		v103 = v110
		v104 = v112
		v105 = v108
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v120 = v110
	v121 = v112
	v122 = v108
	goto L25
L39:
	;
	goto L38
L40:
	;
	v129 = v120
	v131 = v122
	goto L41
L41:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)) = uint8(v132)
	v134 = int32(1)
	if v132 != 0 {
		v129 = v129 + v134
		v131 = v131 + v134
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L24
L43:
	;
	goto L42
L44:
	;
	if v157 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v161 = v159 | v22
	goto L47
L46:
	;
	v161 = v22
	goto L47
L47:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v163 != 0 {
		v22 = v161
		goto L4
	} else {
		goto L48
	}
L48:
	;
	goto L5
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
	var v7 int32
	_ = v7
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	v7 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
			v13 = v11 + v12
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
			if l1 != 0 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
			} else {
			}
			F_ReleaseCatCache(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v28 = F_OidFunctionCall0Coll(m, v14)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v33 = base.I32_wrap_i64(v28)
						return v33
					}
				} else {
					return int32(0)
				}
			}
		} else {
			v21 = int32(0)
			if l1 == v21 {
				v33 = v21
				return v33
			} else {
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
				return v24
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v6 = F_SearchSysCache2(m, int32(7), base.I64_extend_i32_u(l0), base.I64_extend_i32_s(l1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14+v15)+68))
			F_ReleaseCatCache(m, v6)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v17
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
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
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
	v15 = int64(0)
	v17 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(l0), v15, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_ReleaseCatCacheList(m, v17)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L27
	}
L5:
	;
	return int32(0)
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	if v21 <= int32(0) {
		v93 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v30 = v3
	v32 = v3
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17-int32(-64)+v32<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v41 = v39 + v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v42 != int32(405) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v93 = v84
	goto L4
L10:
	;
	v86 = v32 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	if v86 < v87 {
		v30 = v84
		v32 = v86
		goto L8
	} else {
		goto L26
	}
L11:
	;
	v84 = v30
	goto L10
L12:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+16)))
	if v45 != int32(1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if v48 == v49 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = int32(1)
	if l1 == int32(0) {
		v93 = v51
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
	v93 = v51
	goto L4
L18:
	;
	v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+4)))
	v59 = base.I64_extend_i32_u(v48)
	v61 = F_SearchSysCache4(m, int32(4), v58, v59, v59, int64(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v72
	v74 = int32(0)
	v76 = base.B2i32(v72 != v74) | v30
	if v72 == v74 {
		v84 = v76
		goto L10
	} else {
		goto L25
	}
L20:
	;
	if v61 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = int32(0)
	goto L19
L22:
	;
	goto L23
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v67)+20))
	F_ReleaseCatCache(m, v61)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v72 = v69
	goto L19
L25:
	;
	v93 = v76
	goto L4
L26:
	;
	goto L9
L27:
	;
	return v93 & int32(1)
}
func F_get_const_expr(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v14 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(96)
	return
L2:
	;
	F_appendStringInfoString(m, v13, int32(_a_F_get_const_expr_0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_getTypeOutputInfo(m, v47, v11+int32(92), v11+int32(91))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_format_type_with_typemod(m, v22, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v24
	F_appendStringInfo(m, v13, int32(_a_F_get_const_expr_1), v11+int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v32 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = F_get_typcollation(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v37 == v39 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v41 = F_generate_collation_name(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v41
	F_appendStringInfo(m, v35, int32(_a_F_get_const_expr_2), v11)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v56 = F_OidOutputFunctionCall(m, v54, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v58 - int32(16) {
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
	F_pfree(m, v56)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L68
	}
L18:
	;
	F_appendStringInfoString(m, v13, v56)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L67
	}
L19:
	;
	v204 = int32(1)
	goto L17
L20:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if base.Ui32((v110-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L44
	} else {
		goto L45
	}
L21:
	;
	F_appendStringInfoChar(m, v13, int32(39))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L33
	}
L22:
	;
	if v58 == int32(1700) {
		goto L20
	} else {
		goto L32
	}
L23:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v70 != int32(116) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v61 != int32(45) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v56
	F_appendStringInfo(m, v13, int32(_a_F_get_const_expr_3), v11-int32(-64))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	F_appendStringInfoString(m, v13, int32(_a_F_get_const_expr_4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L31
	}
L28:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v73 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_appendStringInfoString(m, v13, int32(_a_F_get_const_expr_5))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v204 = int32(0)
	goto L17
L31:
	;
	v204 = int32(0)
	goto L17
L32:
	;
	goto L21
L33:
	;
	v93 = v56
	goto L34
L34:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v95 != int32(39) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_appendStringInfoChar(m, v13, base.I32_extend8_s(v95))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L43
	}
L37:
	;
	if v95 != 0 {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_appendStringInfoChar(m, v13, int32(39))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	F_appendStringInfoChar(m, v13, int32(39))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v204 = int32(0)
	goto L17
L42:
	;
	goto L36
L43:
	;
	v93 = v93 + int32(1)
	goto L34
L44:
	;
	v117 = int32(_a_F_get_const_expr_6)
	v121 = m.G0
	v123 = v121 - int32(32)
	m.G0 = v123
	v125 = int32(*(*int8)(unsafe.Add(mBase, _c_F_get_const_expr[0])))
	if v125 != 0 {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v56
	F_appendStringInfo(m, v13, int32(_a_F_get_const_expr_3), v11+int32(80))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L66
	}
L47:
	;
	v184 = F_strlen(m, v56)
	mBase = m.M
	if v176-v56 != v184 {
		goto L18
	} else {
		goto L65
	}
L48:
	;
	m.G0 = v123 + int32(32)
	goto L47
L49:
	;
	F___memset(m, v123, int32(0), int32(32))
	mBase = m.M
	v131 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_const_expr[0])))
	if v131 != 0 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_const_expr[1])))
	if v126 != 0 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v127 = F___strchrnul(m, v56, v125)
	mBase = m.M
	v176 = v127
	goto L48
L53:
	;
	goto L52
L54:
	;
	v133 = v117
	v134 = v131
	goto L57
L55:
	;
	goto L56
L56:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v155 == int32(0) {
		v176 = v56
		goto L48
	} else {
		goto L60
	}
L57:
	;
	v141 = v123 + int32(base.Ui32(v134)>>(uint(int32(3))%32))&int32(28)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v143 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v142 | v143<<(uint(v134)%32)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	if v147 != 0 {
		v133 = v133 + v143
		v134 = v147
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	goto L58
L60:
	;
	v159 = v56
	v160 = v155
	goto L61
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v123+int32(base.Ui32(v160)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v168)>>(uint(v160)%32))&int32(1) != 0 {
		v176 = v159
		goto L48
	} else {
		goto L63
	}
L62:
	;
	v176 = v174
	goto L48
L63:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	v174 = v159 + int32(1)
	if v172 != 0 {
		v159 = v174
		v160 = v172
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L46
L66:
	;
	goto L19
L67:
	;
	v204 = int32(0)
	goto L17
L68:
	;
	if l2 < int32(0) {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v209 - int32(16) {
	case 0:
		goto L74
	case 1, 2, 3, 4, 5, 6:
		goto L71
	case 7:
		v221 = v204
		goto L72
	default:
		goto L75
	}
L70:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v235 == int32(0) {
		goto L1
	} else {
		goto L82
	}
L71:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v226 = F_format_type_with_typemod(m, v209, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L80
	}
L72:
	;
	if l2 != 0 {
		goto L71
	} else {
		goto L78
	}
L73:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v221 = v204 | base.B2i32(int32(0) <= v217)
	goto L72
L74:
	;
	v221 = int32(0)
	goto L72
L75:
	;
	if v209 == int32(1700) {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	if v209 != int32(705) {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	if v221 == int32(0) {
		goto L70
	} else {
		goto L79
	}
L79:
	;
	goto L71
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v226
	F_appendStringInfo(m, v13, int32(_a_F_get_const_expr_1), v11+int32(48))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	goto L70
L82:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v240 = F_get_typcollation(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v240 == v242 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v244 = F_generate_collation_name(m, v242)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v244
	F_appendStringInfo(m, v238, int32(_a_F_get_const_expr_2), v11+int32(32))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
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
	F_list_sort(m, v176, int32(1137))
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
	F_errfinish(m, int32(_a_F_get_policies_for_relation_1), int32(605), int32(_a_F_get_policies_for_relation_2))
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
	F_list_sort(m, v184, int32(1137))
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v36 float64
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
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
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v11 == int32(0) {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
		if v10&int32(1) != 0 {
			if v14&int32(1) == int32(0) {
				return float64(0)
			} else {
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
				if v82 != 0 {
					v83 = float64(0)
				} else {
					v83 = float64(1)
				}
				v85 = v83
				return v85
			}
		} else {
			v17 = float64(0.5)
			if v14&int32(1) != 0 {
				v85 = v17
				return v85
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				if v20 == int32(0) {
					v85 = v17
					return v85
				} else {
					v24 = l0 + int32(268)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
					v26 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
					v27 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					v28 = F_FunctionCall2Coll(m, v24, v25, v26, v27)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return float64(0)
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v28&int64(9223372036854775807)) {
							v85 = v17
							return v85
						} else {
							v36 = base.F64_reinterpret_i64(v28)
							if base.F64_le(v36, float64(0)) != 0 {
								v85 = v17
								return v85
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
								v41 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
								v42 = F_FunctionCall2Coll(m, v24, v39, v40, v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return float64(0)
								} else {
									v45 = base.F64_div(base.F64_reinterpret_i64(v42), v36)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v45)&int64(9223372036854775807)) {
										v85 = v17
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
		if v10&int32(1) != 0 {
			return float64(0.5)
		} else {
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			if v64 != int32(1) {
				return float64(1)
			} else {
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
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
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = base.I64_extend_i32_u(l0)
	v7 = F_SearchSysCache1(m, int32(82), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11+v12)+96))
			F_ReleaseCatCache(m, v7)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v38 = v14
					return v38
				} else {
					v19 = F_SearchSysCache1(m, int32(82), v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						if v19 == int32(0) {
							return int32(0)
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
							v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
							v27 = v25 + v26
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
							if v28 != 0 {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
								if v30 == int32(_a_F_get_promoted_array_type_0) {
									v33 = l0
								} else {
									v33 = int32(0)
								}
								v35 = v33
							} else {
								v35 = int32(0)
							}
							F_ReleaseCatCache(m, v19)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = v35
								return v38
							}
						}
					}
				}
			}
		} else {
			v19 = F_SearchSysCache1(m, int32(82), v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
					v27 = v25 + v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
					if v28 != 0 {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
						if v30 == int32(_a_F_get_promoted_array_type_0) {
							v33 = l0
						} else {
							v33 = int32(0)
						}
						v35 = v33
					} else {
						v35 = int32(0)
					}
					F_ReleaseCatCache(m, v19)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = v35
						return v38
					}
				}
			}
		}
	}
}
func F_get_reloptions(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_pg_detoast_datum(m, base.I32_wrap_i64(l1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_deconstruct_array_builtin(m, v13, int32(25), v10+int32(12), int32(0), v10+int32(8))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if int32(0) < v23 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = int32(0)
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v30<<(uint(int32(3))%32))))
	v38 = F_text_to_cstring(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v40 = int32(61)
	v41 = F___strchrnul(m, v38, v40)
	mBase = m.M
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v43 == v40 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v47 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v47 = v41
	goto L13
L12:
	;
	v47 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v48)
	v53 = v47 + int32(1)
	goto L16
L15:
	;
	v53 = int32(_a_F_get_reloptions_0)
	goto L16
L16:
	;
	if v30 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_get_reloptions_1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v57 = F_quote_identifier(m, v38)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v57
	F_appendStringInfo(m, l0, int32(_a_F_get_reloptions_2), v10)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v63 = F_quote_identifier(m, v53)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	F_pfree(m, v38)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L40
	}
L24:
	;
	if v63 == v53 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_appendStringInfoString(m, l0, v53)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L23
L29:
	;
	v73 = v53
	goto L30
L30:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v78 != int32(39) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	F_appendStringInfoChar(m, l0, base.I32_extend8_s(v78))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L39
	}
L33:
	;
	if v78 != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_appendStringInfoChar(m, l0, int32(39))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	F_appendStringInfoChar(m, l0, int32(39))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L23
L38:
	;
	goto L32
L39:
	;
	v73 = v73 + int32(1)
	goto L30
L40:
	;
	v102 = v30 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v102 < v103 {
		v30 = v102
		goto L7
	} else {
		goto L41
	}
L41:
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
			v30 = F_list_make1_impl(m, int32(480), v12+int32(12))
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
					*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(381)
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
	if v138 != int32(453) {
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
	if v298 != int32(453) {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_gimme_tree[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v10 < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v10<<(uint(int32(2))%32))))
	v18 = v17
	goto L3
L2:
	;
	v18 = v4
	goto L3
L3:
	;
	if l2 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v27 = int32(0)
	v28 = v4
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v34 = int32(2)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v27<<(uint(v34)%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33+v37<<(uint(v34)%32)-int32(4))))
	v45 = F_palloc(m, int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v58 = int32(0)
	if v53 == v58 {
		v106 = v58
		goto L13
	} else {
		goto L14
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v43
	v53 = F_merge_clump(m, l0, v28, v45, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v56 = v27 + int32(1)
	if v56 != l2 {
		v27 = v56
		v28 = v53
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	return v106
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v61 < int32(2) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v94 != int32(1) {
		v106 = v58
		goto L13
	} else {
		goto L26
	}
L16:
	;
	v92 = v53
	v94 = v61
	goto L15
L17:
	;
	goto L18
L18:
	;
	v64 = int32(0)
	v67 = v64
	v69 = v64
	goto L19
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v69<<(uint(int32(2))%32))))
	v80 = F_merge_clump(m, l0, v67, v78, int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	if v80 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v83 = v69 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v83 < v84 {
		v67 = v80
		v69 = v83
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	return int32(0)
L24:
	;
	goto L25
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v92 = v80
	v94 = v90
	goto L15
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v106 = v103
	goto L13
}
func F_ginarrayextract_2args(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v10 <= int32(2) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_ginarrayextract_2args_0), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_ginarrayextract_2args_1), int32(71), int32(_a_F_ginarrayextract_2args_2))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v29 = F_pg_detoast_datum_copy(m, v28)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
			F_get_typlenbyvalalign(m, v33, v8+int32(14), v8+int32(13), v8+int32(12))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+14)))
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)))
				v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8)+12)))
				F_deconstruct_array(m, v29, v43, v44, v45, v8+int32(8), v8+int32(4), v8)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					*(*int32)(unsafe.Add(mBase, uint32(v32))) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v54
					v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)))
					m.G0 = v8 + int32(16)
					return v56
				}
			}
		}
	}
}
func F_ginarraytriconsistent(m *base.Module, l0 int32) int64 {
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
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = int32(2)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = base.I32_wrap_i64(v17) & int32(_a_F_ginarraytriconsistent_0)
	switch v20 - int32(1) {
	case 0:
		goto L6
	case 1:
		goto L7
	case 2:
		v111 = v16
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
	v130 = m.ExcPending
	if v130 != 0 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return base.I64_extend_i32_u(v111) & int64(255)
L3:
	;
	v111 = int32(0)
	goto L2
L4:
	;
	v86 = v23
	goto L33
L5:
	;
	v64 = v26
	v66 = int32(1)
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
	v26 = int32(0)
	if v26 < v14 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v23 = int32(0)
	if v14 <= v23 {
		v111 = v16
		goto L2
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	v111 = int32(1)
	goto L2
L11:
	;
	v34 = int32(0)
	v36 = int32(0)
	goto L12
L12:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v13))))
	if v43 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v111 = v59
	goto L2
L14:
	;
	v46 = int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v15))))
	if v48 == v46 {
		v111 = v46
		goto L2
	} else {
		goto L17
	}
L15:
	;
	v59 = v36
	goto L16
L16:
	;
	v62 = v34 + int32(1)
	if v62 != v14 {
		v34 = v62
		v36 = v59
		goto L12
	} else {
		goto L24
	}
L17:
	;
	if v36&int32(255) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = v36
	goto L20
L19:
	;
	v54 = int32(2)
	goto L20
L20:
	;
	if v48 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v57 = v54
	goto L23
L22:
	;
	v57 = v36
	goto L23
L23:
	;
	v59 = v57
	goto L16
L24:
	;
	goto L13
L25:
	;
	v72 = int32(0)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v15))))
	if v74 == v72 {
		v111 = v72
		goto L2
	} else {
		goto L27
	}
L26:
	;
	v111 = v82
	goto L2
L27:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v13))))
	if v78 != 0 {
		v111 = v72
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v79 = int32(2)
	if v74 == v79 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v82 = v79
	goto L31
L30:
	;
	v82 = v66
	goto L31
L31:
	;
	v84 = v64 + int32(1)
	if v84 != v14 {
		v64 = v84
		v66 = v82
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v15))))
	if v95 == int32(0) {
		goto L3
	} else {
		goto L35
	}
L34:
	;
	v111 = v16
	goto L2
L35:
	;
	v99 = v86 + int32(1)
	if v14 != v99 {
		v86 = v99
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	return int64(0)
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v20
	F_errmsg_internal(m, int32(_a_F_ginarraytriconsistent_1), v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_ginarraytriconsistent_2), int32(306), int32(_a_F_ginarraytriconsistent_3))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
func F_ginhandler(m *base.Module, l0 int32) int64 {
	return int64(829312)
}
func F_ginqueryarrayextract(m *base.Module, l0 int32) int64 {
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
	var v16 int64
	_ = v16
	var v17 int64
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
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_copy(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		F_get_typlenbyvalalign(m, v20, v9+int32(30), v9+int32(29), v9+int32(28))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+30)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
			v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+28)))
			F_deconstruct_array(m, v12, v30, v31, v32, v9+int32(24), v9+int32(20), v9+int32(16))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v43
				v48 = base.I32_wrap_i64(v17) & int32(_a_F_ginqueryarrayextract_0)
				switch v48 - int32(1) {
				case 0:
					v72 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)))) = v72
					v75 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
					m.G0 = v9 + int32(32)
					return v75
				case 1:
					v68 = int32(0)
					if v41 <= v68 {
						v71 = int32(2)
					} else {
						v71 = v68
					}
					v72 = v71
					*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)))) = v72
					v75 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
					m.G0 = v9 + int32(32)
					return v75
				case 2:
					v72 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)))) = v72
					v75 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
					m.G0 = v9 + int32(32)
					return v75
				case 3:
					v72 = base.B2i32(v41 <= int32(0))
					*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)))) = v72
					v75 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
					m.G0 = v9 + int32(32)
					return v75
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48
						F_errmsg_internal(m, int32(_a_F_ginqueryarrayextract_1), v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_ginqueryarrayextract_2), int32(132), int32(_a_F_ginqueryarrayextract_3))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v65 int64
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v211 float32
	_ = v211
	var v213 float32
	_ = v213
	var v216 float32
	_ = v216
	var v222 float32
	_ = v222
	var v226 float32
	_ = v226
	var v228 float32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 float32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	v26 = m.G0
	v28 = v26 - int32(992)
	m.G0 = v28
	F_gistDeCompressAtt(m, l3, l0, l2, v28+int32(48), v28+int32(16))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+848)) = int32(-1082130432)
	v40 = int32(1)
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
	if base.Ui32(v41) < base.Ui32(int32(25)) {
		v333 = v40
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v28 + int32(992)
	return v333 & int32(_a_F_gistchoose_0)
L4:
	;
	v49 = int32(base.Ui32(v41+int32(_a_F_gistchoose_1))>>(uint(int32(2))%32)) & int32(_a_F_gistchoose_0)
	if v49 == int32(0) {
		v333 = v40
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v55 = l3 + int32(_a_F_gistchoose_2)
	v65 = base.I64_extend_i32_u(v28 + int32(824))
	v74 = int32(-1)
	v75 = int32(1)
	v77 = v40
	goto L6
L6:
	;
	v94 = v75 & int32(_a_F_gistchoose_0)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v95)+10)))
	if v96 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v333 = v309
	goto L3
L8:
	;
	if base.B2i32(v275&int32(_a_F_gistchoose_0) == v94)|base.B2i32(v274 != base.I32_extend16_s(v271)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L9:
	;
	v271 = v96
	v272 = v74
	v274 = int32(0)
	v275 = v77
	v276 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v94<<(uint(int32(2))%32))))
	v116 = v74
	v118 = int32(0)
	v119 = v77
	v120 = int32(1)
	goto L12
L12:
	;
	v136 = v118 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v140 = F_index_getattr_2(m, l1+v104&int32(_a_F_gistchoose_3), v136, v137, v28+int32(15))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263)+10)))
	v271 = v264
	v272 = v116
	v274 = v118
	v275 = v119
	v276 = int32(0)
	goto L8
L14:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)))
	if v142 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(16)+v118))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+988)) = int32(0)
	v194 = l3 + int32(3604) + v118*int32(28)
	if (v189|v142)&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+842)) = uint8(v183)
	goto L15
L17:
	;
	v145 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+842)) = uint8(v145)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+840)) = uint16(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+836)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+832)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v28)+824)) = v140
	v153 = l3 + int32(2708) + v118*int32(28)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v154 == v145 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+840)) = uint16(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+836)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+832)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v28)+824)) = int64(0)
	v183 = int32(0)
	goto L16
L20:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v55+v118<<(uint(int32(2))%32))))
	v161 = F_FunctionCall1Coll(m, v153, v160, v65)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v163 = base.I32_wrap_i64(v161)
	if v163 == v28+int32(824) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v163)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+824)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+832)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+836)) = v171
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+840)) = uint16(v173)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+18)))
	v183 = v175
	goto L16
L23:
	;
	v230 = v28 + int32(848)
	v233 = v230 + v118<<(uint(int32(2))%32)
	v234 = *(*float32)(unsafe.Add(mBase, uint32(v233)))
	if base.F32_lt(v234, float32(0))|base.F32_lt(v228, v234) != 0 {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	if v189&v142 != 0 {
		goto L36
	} else {
		goto L37
	}
L25:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+10)))
	if v198 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v55+v118<<(uint(int32(2))%32))))
	v209 = F_FunctionCall3Coll(m, v194, v202, v65, base.I64_extend_i32_u(v28+int32(48)+v118*int32(24)), base.I64_extend_i32_u(v28+int32(988)))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v211 = float32(0)
	v213 = *(*float32)(unsafe.Add(mBase, uint32(v28)+988))
	if base.F32_lt(v213, v211) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v216 = v211
	goto L32
L31:
	;
	v216 = v213
	goto L32
L32:
	;
	if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v213)&int32(2147483647)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v222 = v211
	goto L35
L34:
	;
	v222 = v216
	goto L35
L35:
	;
	v228 = v222
	goto L23
L36:
	;
	v226 = float32(0)
	goto L38
L37:
	;
	v226 = math.Float32frombits(uint32(0x7f800000))
	goto L38
L38:
	;
	v228 = v226
	goto L23
L39:
	;
	goto L13
L40:
	;
	v261 = base.B2i32(base.F32_gt(v228, float32(0)) == int32(0)) & v120
	if v136 < v254 {
		v116 = v255
		v118 = v136
		v119 = v256
		v120 = v261
		goto L12
	} else {
		goto L48
	}
L41:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v233))) = v228
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241)+10)))
	if v118 < v242-int32(1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if base.F32_ne(v228, v234) != 0 {
		goto L39
	} else {
		goto L47
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136<<(uint(int32(2))%32)+v230))) = int32(-1082130432)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v254 = v242
	v255 = int32(-1)
	v256 = v75
	goto L40
L47:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252)+10)))
	v254 = v253
	v255 = v116
	v256 = v119
	goto L40
L48:
	;
	v271 = v254
	v272 = v255
	v274 = v136
	v275 = v256
	v276 = v261
	goto L8
L49:
	;
	if v272 == int32(-1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v308 = v272
	v309 = v275
	goto L51
L51:
	;
	if v276 != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v302 = Fn14349(m, int64(63))
	mBase = m.M
	goto L55
L53:
	;
	v303 = v272
	goto L54
L54:
	;
	if v303 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v303 = v302
	goto L54
L56:
	;
	v304 = v275
	goto L58
L57:
	;
	v304 = v75
	goto L58
L58:
	;
	if v303 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v307 = int32(1)
	goto L61
L60:
	;
	v307 = int32(-1)
	goto L61
L61:
	;
	v308 = v307
	v309 = v304
	goto L51
L62:
	;
	if v308 == int32(-1) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v318 = v308
	goto L64
L64:
	;
	v320 = v75 + int32(1)
	if base.Ui32(v320&int32(_a_F_gistchoose_0)) <= base.Ui32(v49) {
		v74 = v318
		v75 = v320
		v77 = v309
		goto L6
	} else {
		goto L70
	}
L65:
	;
	v313 = Fn14349(m, int64(63))
	mBase = m.M
	goto L68
L66:
	;
	v314 = v308
	goto L67
L67:
	;
	if v314 == int32(1) {
		v333 = v309
		goto L3
	} else {
		goto L69
	}
L68:
	;
	v314 = v313
	goto L67
L69:
	;
	v318 = int32(0)
	goto L64
L70:
	;
	goto L7
}
func F_gistcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 float64
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v70 float64
	_ = v70
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v21 = v17 + int32(8)
	base.MemoryFill(m, v21, int32(0), int32(72))
	F_genericcostestimate(m, l0, l1, l2, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
		if v27 < int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			if base.Ui32(int32(2)) <= base.Ui32(v31) {
				v35 = F_log(m, base.F64_convert_i32_u(v31))
				mBase = m.M
				v39 = base.I32_trunc_sat_f64_s(base.F64_div(v35, float64(4.605170185988092)))
			} else {
				v39 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v39
			v42 = v39
		} else {
			v42 = v27
		}
		v44 = *(*float64)(unsafe.Add(mBase, _c_F_gistcostestimate[0]))
		v45 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
		v46 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
		if base.F64_gt(v46, float64(1)) == int32(0) {
			v51 = *(*float64)(unsafe.Add(mBase, uint32(v17)+64))
			v52 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
			v61 = v45
			v62 = v51
			v64 = v52
		} else {
			v53 = F_log(m, v46)
			mBase = m.M
			v55 = base.F64_mul(base.F64_ceil(v53), v44)
			v57 = *(*float64)(unsafe.Add(mBase, uint32(v17)+64))
			v59 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
			v61 = base.F64_add(v45, v55)
			v62 = v57
			v64 = base.F64_add(base.F64_mul(v57, v55), v59)
		}
		v70 = base.F64_mul(v44, base.F64_mul(base.F64_convert_i32_s(v42+int32(1)), float64(50)))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v61, v70)
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(v62, v70), v64)
		v76 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v76
		v78 = *(*float64)(unsafe.Add(mBase, uint32(v17)+32))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v78
		v80 = *(*float64)(unsafe.Add(mBase, uint32(v17)+40))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v80
		m.G0 = v17 + int32(80)
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
	var v27 int32
	_ = v27
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
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(v14^int32(-1))*int32(56))+16))
		v33 = v24
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_gistinserttuple[1]))
		v27 = int32(56)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+v14*v27-v27)+16))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = F_repalloc_mul(m, l0, int32(4), v7+l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = l3 << (uint(int32(2)) % 32)
		if v14 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			base.MemoryCopy(m, v9+v15<<(uint(int32(2))%32), l2, v14)
		} else {
		}
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20 + l3
		return v9
	}
}
func F_greek_UTF_8_create_env(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_SN_new_env(m, int32(32))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			v7 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v3)+28)) = uint8(v7)
		} else {
		}
		return v3
	}
}
func F_gseg_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v56 float32
	_ = v56
	var v57 float32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v69 float32
	_ = v69
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = base.I32_wrap_i64(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v21 = (v17 - int32(1)) & int32(_a_F_gseg_picksplit_0)
	v24 = F_palloc(m, v21*int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v17&int32(_a_F_gseg_picksplit_0) != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = int32(1)
	if base.Ui32(v21) <= base.Ui32(v32) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_pg_qsort(m, v24, v21, int32(12), int32(_a_F_gseg_picksplit_1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	v35 = v32
	goto L8
L7:
	;
	v35 = v21
	goto L8
L8:
	;
	v39 = int32(1)
	goto L9
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(8)+v39*int32(24))))
	v56 = *(*float32)(unsafe.Add(mBase, uint32(v55)))
	v57 = *(*float32)(unsafe.Add(mBase, uint32(v55)+4))
	v58 = int32(12)
	v60 = v24 + v39*v58
	*(*int32)(unsafe.Add(mBase, uint32(v60-int32(4)))) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v60-int32(8)))) = uint16(v39)
	v69 = float32(0.5)
	*(*float32)(unsafe.Add(mBase, uint32(v60-v58))) = base.F32_add(base.F32_mul(v56, v69), base.F32_mul(v57, v69))
	if v39 != v35 {
		v39 = v39 + int32(1)
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
	v95 = int32(1)
	v97 = v21 << (uint(v95) % 32)
	v98 = F_palloc(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v98
	v101 = F_palloc(m, v97)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v101
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v110 = F_palloc(m, int32(12))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+8)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v112)))
	*(*int64)(unsafe.Add(mBase, uint32(v110))) = v115
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v108))) = uint16(v117)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v120 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v119 + v120
	v124 = int32(base.Ui32(v21) >> (uint(v120) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v21) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v128 = v95
	v130 = v108
	v132 = v110
	goto L19
L17:
	;
	v166 = v110
	goto L18
L18:
	;
	v177 = F_palloc(m, int32(12))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v145 = v24 + v128*int32(12)
	v146 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v145)+8)))
	v147 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_picksplit_2), int32(0), base.I64_extend_i32_u(v132), v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v166 = v157
	goto L18
L21:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)) = uint16(v149)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v152 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v151 + v152
	v157 = base.I32_wrap_i64(v147)
	v159 = v128 + v152
	if v159 != v124 {
		v128 = v159
		v130 = v130 + int32(2)
		v132 = v157
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v181 = v24 + v124*int32(12)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+8)) = v183
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(v177))) = v185
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v101))) = uint16(v187)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v190 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v189 + v190
	v194 = v124 + v190
	if base.Ui32(v194) < base.Ui32(v21) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v196 = v101
	v197 = v177
	v199 = v194
	goto L27
L25:
	;
	v231 = v177
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = base.I64_extend_i32_u(v231)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = base.I64_extend_i32_u(v166)
	return v14 & int64(4294967295)
L27:
	;
	v214 = v24 + v199*int32(12)
	v215 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v214)+8)))
	v216 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_picksplit_2), int32(0), base.I64_extend_i32_u(v197), v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v231 = v226
	goto L26
L29:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+2)) = uint16(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v221 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v220 + v221
	v226 = base.I32_wrap_i64(v216)
	v228 = v199 + v221
	if v228 != v21 {
		v196 = v196 + int32(2)
		v197 = v226
		v199 = v228
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
func F_gtsquery_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v130 int32
	_ = v130
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
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
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v258 int32
	_ = v258
	var v265 int64
	_ = v265
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v293 int32
	_ = v293
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v326 int64
	_ = v326
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v371 int32
	_ = v371
	var v373 int64
	_ = v373
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v439 int32
	_ = v439
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v452 int32
	_ = v452
	var v453 int64
	_ = v453
	var v475 int32
	_ = v475
	var v498 int32
	_ = v498
	var v500 int64
	_ = v500
	var v503 int64
	_ = v503
	var v507 int64
	_ = v507
	var v516 int32
	_ = v516
	var v529 int32
	_ = v529
	var v552 int32
	_ = v552
	var v554 int64
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v588 int64
	_ = v588
	var v589 int64
	_ = v589
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v607 int32
	_ = v607
	var v612 int64
	_ = v612
	var v613 int64
	_ = v613
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	v8 = int32(0)
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = base.I32_wrap_i64(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v29 = (v25 + int32(_a_F_gtsquery_picksplit_0)) & int32(_a_F_gtsquery_picksplit_1)
	v33 = v29<<(uint(int32(1))%32) + int32(4)
	v34 = F_palloc(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = v34
		v39 = F_palloc(m, v33)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int64(0)
		} else {
			v41 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v41
			*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v39
			*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v41
			if base.Ui32(int32(2)) <= base.Ui32(v29) {
				v49 = v24 + int32(8)
				v60 = int32(1)
				v61 = int32(-1)
				v69 = v8
				v70 = v8
				for {
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v49+v60*int32(24))))
					v78 = v60 + int32(1)
					v79 = v78
					v88 = v61
					v89 = v78
					v96 = v69
					v97 = v70
					for {
						v103 = *(*int64)(unsafe.Add(mBase, uint32(v49+v79*int32(24))))
						v104 = v76 ^ v103
						v107 = int32(0)
						v108 = int64(0)
						for {
							v130 = int32(1)
							v153 = base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(v108)%64)))&v130 + v107 + base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(v108|int64(1))%64)))&v130 + base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(v108|int64(2))%64)))&v130 + base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(v108|int64(3))%64)))&v130
							v155 = v108 + int64(4)
							if v155 != int64(64) {
								v107 = v153
								v108 = v155
								continue
							} else {
								break
							}
							break
						}
						v158 = base.B2i32(v88 < v153)
						if v88 < v153 {
							v159 = v89
						} else {
							v159 = v96
						}
						if v88 < v153 {
							v160 = v60
						} else {
							v160 = v97
						}
						if v88 < v153 {
							v161 = v153
						} else {
							v161 = v88
						}
						v163 = v89 + int32(1)
						v165 = v163 & int32(_a_F_gtsquery_picksplit_1)
						if base.Ui32(v165) <= base.Ui32(v29) {
							v79 = v165
							v88 = v161
							v89 = v163
							v96 = v159
							v97 = v160
							continue
						} else {
							break
						}
						break
					}
					if v78 != v29 {
						v60 = v78
						v61 = v161
						v69 = v159
						v70 = v160
						continue
					} else {
						break
					}
					break
				}
				v185 = v159
				v186 = v160
			} else {
				v185 = v8
				v186 = v8
			}
			v192 = v24 + int32(8)
			v194 = int32(_a_F_gtsquery_picksplit_1)
			v196 = int32(0)
			v202 = base.B2i32(v186&v194 == v196) | base.B2i32(v185&v194 == v196)
			if v202 != 0 {
				v203 = int32(2)
			} else {
				v203 = v185
			}
			v208 = v192 + v203&int32(_a_F_gtsquery_picksplit_1)*int32(24)
			v209 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
			if v202 != 0 {
				v211 = int32(1)
			} else {
				v211 = v186
			}
			v212 = int32(_a_F_gtsquery_picksplit_1)
			v216 = v192 + v211&v212*int32(24)
			v217 = *(*int64)(unsafe.Add(mBase, uint32(v216)))
			v220 = v25 + v212
			v222 = v220 & v212
			v223 = F_palloc_mul(m, int32(8), v222)
			mBase = m.M
			v224 = m.ExcPending
			if v224 != 0 {
				return int64(0)
			} else {
				if v25&int32(_a_F_gtsquery_picksplit_1) == int32(1) {
					F_pg_qsort(m, v223, v222, int32(8), int32(1724))
					mBase = m.M
					v232 = m.ExcPending
					if v232 != 0 {
						return int64(0)
					} else {
						v612 = v209
						v613 = v217
						v620 = v39
						v622 = v34
						v630 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v622))) = uint16(v630)
						*(*uint16)(unsafe.Add(mBase, uint32(v620))) = uint16(v630)
						*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v612
						*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v613
						return v22 & int64(4294967295)
					}
				} else {
					v233 = int32(1)
					v235 = v233
					v244 = v233
					for {
						v258 = v223 + v235<<(uint(int32(3))%32)
						*(*uint16)(unsafe.Add(mBase, uint32(v258-int32(8)))) = uint16(v244)
						v265 = *(*int64)(unsafe.Add(mBase, uint32(v192+v235*int32(24))))
						v266 = *(*int64)(unsafe.Add(mBase, uint32(v216)))
						v267 = v265 ^ v266
						v270 = int32(0)
						v271 = int64(0)
						for {
							v293 = int32(1)
							v316 = base.I32_wrap_i64(int64(base.Ui64(v267)>>(uint(v271)%64)))&v293 + v270 + base.I32_wrap_i64(int64(base.Ui64(v267)>>(uint(v271|int64(1))%64)))&v293 + base.I32_wrap_i64(int64(base.Ui64(v267)>>(uint(v271|int64(2))%64)))&v293 + base.I32_wrap_i64(int64(base.Ui64(v267)>>(uint(v271|int64(3))%64)))&v293
							v318 = v271 + int64(4)
							if v318 != int64(64) {
								v270 = v316
								v271 = v318
								continue
							} else {
								break
							}
							break
						}
						v321 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
						v322 = v321 ^ v265
						v326 = int64(0)
						v335 = int32(0)
						for {
							v348 = int32(1)
							v371 = base.I32_wrap_i64(int64(base.Ui64(v322)>>(uint(v326)%64)))&v348 + v335 + base.I32_wrap_i64(int64(base.Ui64(v322)>>(uint(v326|int64(1))%64)))&v348 + base.I32_wrap_i64(int64(base.Ui64(v322)>>(uint(v326|int64(2))%64)))&v348 + base.I32_wrap_i64(int64(base.Ui64(v322)>>(uint(v326|int64(3))%64)))&v348
							v373 = v326 + int64(4)
							if v373 != int64(64) {
								v326 = v373
								v335 = v371
								continue
							} else {
								break
							}
							break
						}
						v378 = v316 - v371
						v380 = v378 >> (uint(int32(31)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(v258-int32(4)))) = v378 ^ v380 - v380
						v385 = v244 + int32(1)
						v386 = int32(_a_F_gtsquery_picksplit_1)
						v387 = v385 & v386
						if base.Ui32(v387) <= base.Ui32(v220&v386) {
							v235 = v387
							v244 = v385
							continue
						} else {
							break
						}
						break
					}
					F_pg_qsort(m, v223, v222, int32(8), int32(1724))
					mBase = m.M
					v394 = m.ExcPending
					if v394 != 0 {
						return int64(0)
					} else {
						v395 = int32(1)
						if base.Ui32(v222) <= base.Ui32(v395) {
							v398 = v395
						} else {
							v398 = v222
						}
						v405 = v209
						v406 = v217
						v411 = int32(0)
						v413 = v39
						v415 = v34
						for {
							v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223+v411<<(uint(int32(3))%32)))))
							if v211&int32(_a_F_gtsquery_picksplit_1) == v426 {
								*(*uint16)(unsafe.Add(mBase, uint32(v415))) = uint16(v211)
								v429 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v429 + int32(1)
								v588 = v405
								v589 = v406
								v596 = v413
								v598 = v415 + int32(2)
							} else {
								if v203&int32(_a_F_gtsquery_picksplit_1) == v426 {
									*(*uint16)(unsafe.Add(mBase, uint32(v413))) = uint16(v203)
									v439 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v439 + int32(1)
									v588 = v405
									v589 = v406
									v596 = v413 + int32(2)
									v598 = v415
								} else {
									v448 = *(*int64)(unsafe.Add(mBase, uint32(v192+v426*int32(24))))
									v449 = v448 ^ v406
									v452 = int32(0)
									v453 = int64(0)
									for {
										v475 = int32(1)
										v498 = base.I32_wrap_i64(int64(base.Ui64(v449)>>(uint(v453)%64)))&v475 + v452 + base.I32_wrap_i64(int64(base.Ui64(v449)>>(uint(v453|int64(1))%64)))&v475 + base.I32_wrap_i64(int64(base.Ui64(v449)>>(uint(v453|int64(2))%64)))&v475 + base.I32_wrap_i64(int64(base.Ui64(v449)>>(uint(v453|int64(3))%64)))&v475
										v500 = v453 + int64(4)
										if v500 != int64(64) {
											v452 = v498
											v453 = v500
											continue
										} else {
											break
										}
										break
									}
									v503 = v405 ^ v448
									v507 = int64(0)
									v516 = int32(0)
									for {
										v529 = int32(1)
										v552 = base.I32_wrap_i64(int64(base.Ui64(v503)>>(uint(v507)%64)))&v529 + v516 + base.I32_wrap_i64(int64(base.Ui64(v503)>>(uint(v507|int64(1))%64)))&v529 + base.I32_wrap_i64(int64(base.Ui64(v503)>>(uint(v507|int64(2))%64)))&v529 + base.I32_wrap_i64(int64(base.Ui64(v503)>>(uint(v507|int64(3))%64)))&v529
										v554 = v507 + int64(4)
										if v554 != int64(64) {
											v507 = v554
											v516 = v552
											continue
										} else {
											break
										}
										break
									}
									v559 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
									v560 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
									v561 = v559 - v560
									if base.F64_lt(base.F64_convert_i32_s(v498), base.F64_add(base.F64_convert_i32_s(v552), base.F64_mul(base.F64_convert_i32_s(v561*v561*v561), float64(-0.05)))) != 0 {
										*(*uint16)(unsafe.Add(mBase, uint32(v415))) = uint16(v426)
										v570 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v570 + int32(1)
										v588 = v405
										v589 = v406 | v448
										v596 = v413
										v598 = v415 + int32(2)
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v413))) = uint16(v426)
										v578 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v578 + int32(1)
										v588 = v405 | v448
										v589 = v406
										v596 = v413 + int32(2)
										v598 = v415
									}
								}
							}
							v607 = v411 + int32(1)
							if v607 != v398 {
								v405 = v588
								v406 = v589
								v411 = v607
								v413 = v596
								v415 = v598
								continue
							} else {
								break
							}
							break
						}
						v612 = v588
						v613 = v589
						v620 = v596
						v622 = v598
						v630 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v622))) = uint16(v630)
						*(*uint16)(unsafe.Add(mBase, uint32(v620))) = uint16(v630)
						*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v612
						*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v613
						return v22 & int64(4294967295)
					}
				}
			}
		}
	}
}
