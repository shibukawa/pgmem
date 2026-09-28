package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TParserGet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_TParserGet[0]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= v14 {
		v248 = v2
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	return v248
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v17 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v24 < v23 {
		v248 = v2
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v27 = v23
	v30 = v22
	v31 = v24
	goto L10
L9:
	;
	v248 = v240 & int32(1)
	goto L6
L10:
	;
	if v27 == v31 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+6)))
	v240 = v238
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v44 = int32(0)
	v45 = v30
	goto L12
L14:
	;
	goto L15
L15:
	;
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v35 == v34 {
		v44 = v34
		v45 = v30
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = F_pg_mblen_range(m, v38+v27, v38+v31)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v44 = v41
	v45 = v43
	goto L12
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v58 == int32(0) {
		v78 = v57
		goto L22
	} else {
		goto L23
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = int32(0)
	v57 = v48 + int32(20)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53<<(uint(int32(3))%32))+uint32(_c_F_TParserGet[1])))
	v57 = v56
	goto L18
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v81 != 0 {
		goto L31
	} else {
		goto L32
	}
L23:
	;
	v62 = v57
	goto L24
L24:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)) = uint8(v67)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	v78 = v74
	goto L22
L26:
	;
	if v70 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v78 = v62
	goto L22
L28:
	;
	goto L29
L29:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	v74 = v62 + int32(20)
	if v72 != 0 {
		v62 = v74
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L25
L31:
	;
	m.T0[v81].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+6)))
	if v84&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v90
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v92
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v97
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+6)))
	v100 = v99
	goto L37
L36:
	;
	v100 = v84
	goto L37
L37:
	;
	if v100&int32(2) != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v191 != int32(77) {
		goto L65
	} else {
		goto L66
	}
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
	F_pfree(m, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v109 = v100 & int32(_a_F_TParserGet_0)
	if v109&int32(4) != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v104
	goto L38
L43:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v78
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v116 = F_palloc(m, int32(32))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v109&int32(16) != 0 {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	if v114 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+24)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v116
	goto L38
L48:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v114)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v116)+16)) = v118
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v114)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v116)+8)) = v120
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v114)))
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v122
	goto L47
L49:
	;
	goto L50
L50:
	;
	v124 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v116)+16)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(v116)+8)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v124
	goto L47
L51:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
	F_pfree(m, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v109&int32(64) != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+24)) = v138
	goto L38
L55:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	if v146 == int32(0) {
		goto L38
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v109&int32(32) == int32(0) {
		goto L38
	} else {
		goto L63
	}
L58:
	;
	v150 = v146
	goto L59
L59:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
	F_pfree(m, v150)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L61
	}
L60:
	;
	goto L38
L61:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = v155
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+24))
	if v161 != 0 {
		v150 = v161
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v174)+8)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+12)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v166)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+16)) = v181
	F_pfree(m, v166)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L38
L65:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+20)) = v191
	goto L67
L66:
	;
	goto L67
L67:
	;
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+6)))
	if v196&int32(1) != 0 {
		v240 = v196
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(v196&int32(8) == int32(0))&base.B2i32(v205 <= v204) != 0 {
		v240 = v196
		goto L9
	} else {
		goto L69
	}
L69:
	;
	if v196&int32(10) != 0 {
		v233 = v204
		v235 = v203
		v236 = v205
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v233 <= v236 {
		v27 = v233
		v30 = v235
		v31 = v236
		goto L10
	} else {
		goto L73
	}
L71:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	if v210 == int32(0) {
		v233 = v204
		v235 = v203
		v236 = v205
		goto L70
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v204 + v210
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v216 + v217
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v222 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v221 + v222
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v226 + v222
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v233 = v232
	v235 = v231
	v236 = v230
	goto L70
L73:
	;
	goto L11
}
func F_TeardownHistoricSnapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_TeardownHistoricSnapshot[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_TeardownHistoricSnapshot[1])) = v2
	return
}
func F_TerminateBackgroundWorker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBackgroundWorker[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBackgroundWorker[1]))
	v12 = F_LWLockAcquire(m, v8+int32(_a_F_TerminateBackgroundWorker_0), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = v5 + v6*int32(1488)
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
		if v14 == v18 {
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v17)+17)) = uint8(v20)
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBackgroundWorker[1]))
			F_LWLockRelease(m, v23+int32(_a_F_TerminateBackgroundWorker_0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TerminateBackgroundWorker[2])))
				if v30 == int32(1) {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBackgroundWorker[3]))
					*(*int32)(unsafe.Add(mBase, uint32(v34+int32(28)))) = int32(1)
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBackgroundWorker[4]))
					v43 = F_pgmem_kill(m, v41, int32(10))
					mBase = m.M
				} else {
				}
				return
			}
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBackgroundWorker[1]))
			F_LWLockRelease(m, v45+int32(_a_F_TerminateBackgroundWorker_0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_TerminateBufferIO(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v28 int64
	_ = v28
	var v40 int64
	_ = v40
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v125 int64
	_ = v125
	var v128 int64
	_ = v128
	var v135 int64
	_ = v135
	var v143 int64
	_ = v143
	var v153 int64
	_ = v153
	var v162 int64
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	v7 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = int64(4194304)
	v17 = base.AtomicRmwOr64(m, l0, int32(24), v15)
	if v17&v15 != v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = v17
	goto L4
L2:
	;
	v114 = v17
	goto L3
L3:
	;
	v119 = v114 | int64(4194304)
	if l4 != 0 {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(_a_F_TerminateBufferIO_0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(_a_F_TerminateBufferIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(_a_F_TerminateBufferIO_2)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v40
	if v28&int64(4194304) != v40 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v114 = v103
	goto L3
L6:
	;
	goto L9
L7:
	;
	goto L8
L8:
	;
	v81 = int32(_a_F_TerminateBufferIO_3)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBufferIO[0]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(8))+8))
	if v84 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v13+int32(8))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	v60 = int64(0)
	v63 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v60, v60)
	if v63&int64(4194304) != v60 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v101 = int64(4194304)
	v103 = base.AtomicRmwOr64(m, l0, int32(24), v101)
	if v103&v101 != int64(0) {
		v28 = v103
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TerminateBufferIO[0])) = v99
	goto L15
L17:
	;
	if int32(999) < v82 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v82 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v89 = int32(900)
	if v89 <= v82 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v92 = v89
	goto L23
L22:
	;
	v92 = v82
	goto L23
L23:
	;
	v99 = v92 + int32(100)
	goto L16
L24:
	;
	v99 = v82 - int32(1)
	goto L16
L25:
	;
	goto L5
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
	goto L29
L27:
	;
	v125 = v7
	goto L28
L28:
	;
	if l1 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v125 = int64(-1)
	goto L28
L30:
	;
	v128 = int64(-1283457025)
	goto L32
L31:
	;
	v128 = int64(-201326593)
	goto L32
L32:
	;
	v135 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v119, v125+(v114&v128|l2)&int64(-4194305))
	if v119 != v135 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v143 = v135
	goto L36
L34:
	;
	v162 = v119
	goto L35
L35:
	;
	if l3 != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v153 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v143, (v143&v128|l2)&int64(-4194305)+v125)
	if v143 != v153 {
		v143 = v153
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v162 = v143
	goto L35
L38:
	;
	goto L37
L39:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBufferIO[1]))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerForget(m, v166, base.I64_extend_i32_s(v167+int32(1)), int32(_a_F_TerminateBufferIO_4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_TerminateBufferIO[2]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ConditionVariableBroadcast(m, v175+v176<<(uint(int32(4))%32))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L11
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v182 = int32(0)
	if base.B2i32(l4 == v182)|base.B2i32(v162&int64(536870912) == int64(0)) == v182 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_WakePinCountWaiter(m, l0)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L11
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	m.G0 = v13 + int32(32)
	return
L47:
	;
	goto L46
}
func F_tan(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v97 float64
	_ = v97
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v106 float64
	_ = v106
	var v132 float64
	_ = v132
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 float64
	_ = v160
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v166 int32
	_ = v166
	var v167 float64
	_ = v167
	var v168 float64
	_ = v168
	var v171 float64
	_ = v171
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v214 int32
	_ = v214
	var v218 float64
	_ = v218
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v227 float64
	_ = v227
	var v229 float64
	_ = v229
	var v231 int64
	_ = v231
	var v233 float64
	_ = v233
	var v237 float64
	_ = v237
	var v249 float64
	_ = v249
	var v253 float64
	_ = v253
	var v254 float64
	_ = v254
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v13 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v13) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v13) < base.Ui32(int32(1044381696)) {
			v254 = l0
		} else {
			v18 = float64(0)
			v19 = int32(0)
			v25 = base.I64_reinterpret_f64(l0)
			v29 = base.B2i32(base.Ui64(v25&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)))
			if v29 == v19 {
				v38 = base.B2i32(int64(0) <= v25)
				if int64(0) <= v25 {
					v39 = v18
				} else {
					v39 = base.F64_neg(v18)
				}
				v43 = base.F64_add(base.F64_sub(float64(0.7853981633974483), base.F64_abs(l0)), base.F64_sub(float64(3.061616997868383e-17), v39))
				v44 = float64(0)
				v45 = v38
			} else {
				v43 = l0
				v44 = v18
				v45 = v19
			}
			v46 = base.F64_mul(v43, v43)
			v47 = base.F64_mul(v43, v46)
			v50 = base.F64_mul(v46, v46)
			v89 = base.F64_add(base.F64_mul(v47, float64(0.3333333333333341)), base.F64_add(base.F64_mul(v46, base.F64_add(base.F64_mul(v47, base.F64_add(base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, float64(-1.8558637485527546e-05)), float64(7.817944429395571e-05))), float64(0.0005880412408202641))), float64(0.0035920791075913124))), float64(0.021869488294859542))), float64(0.13333333333320124)), base.F64_mul(v46, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, float64(2.590730518636337e-05)), float64(7.140724913826082e-05))), float64(0.0002464631348184699))), float64(0.0014562094543252903))), float64(0.0088632398235993))), float64(0.05396825397622605))))), v44)), v44))
			v90 = base.F64_add(v43, v89)
			if v29 == int32(0) {
				v97 = base.F64_convert_i32_s(int32(1))
				v102 = base.F64_add(v43, base.F64_sub(v89, base.F64_div(base.F64_mul(v90, v90), base.F64_add(v90, v97))))
				v104 = base.F64_sub(v97, base.F64_add(v102, v102))
				if v45 != 0 {
					v106 = v104
				} else {
					v106 = base.F64_neg(v104)
				}
				v132 = v106
			} else {
				v132 = v90
			}
			v254 = v132
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v13) {
			v254 = base.F64_sub(l0, l0)
		} else {
			v136 = F___rem_pio2(m, l0, v6)
			mBase = m.M
			v137 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
			v138 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			v140 = v136 & int32(1)
			v144 = int32(0)
			v146 = base.I64_reinterpret_f64(v137)
			v150 = base.B2i32(base.Ui64(v146&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)))
			if v150 == v144 {
				v159 = base.B2i32(int64(0) <= v146)
				if int64(0) <= v146 {
					v160 = v138
				} else {
					v160 = base.F64_neg(v138)
				}
				v164 = base.F64_add(base.F64_sub(float64(0.7853981633974483), base.F64_abs(v137)), base.F64_sub(float64(3.061616997868383e-17), v160))
				v165 = float64(0)
				v166 = v159
			} else {
				v164 = v137
				v165 = v138
				v166 = v144
			}
			v167 = base.F64_mul(v164, v164)
			v168 = base.F64_mul(v164, v167)
			v171 = base.F64_mul(v167, v167)
			v210 = base.F64_add(base.F64_mul(v168, float64(0.3333333333333341)), base.F64_add(base.F64_mul(v167, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_add(base.F64_mul(v171, base.F64_add(base.F64_mul(v171, base.F64_add(base.F64_mul(v171, base.F64_add(base.F64_mul(v171, base.F64_add(base.F64_mul(v171, float64(-1.8558637485527546e-05)), float64(7.817944429395571e-05))), float64(0.0005880412408202641))), float64(0.0035920791075913124))), float64(0.021869488294859542))), float64(0.13333333333320124)), base.F64_mul(v167, base.F64_add(base.F64_mul(v171, base.F64_add(base.F64_mul(v171, base.F64_add(base.F64_mul(v171, base.F64_add(base.F64_mul(v171, base.F64_add(base.F64_mul(v171, float64(2.590730518636337e-05)), float64(7.140724913826082e-05))), float64(0.0002464631348184699))), float64(0.0014562094543252903))), float64(0.0088632398235993))), float64(0.05396825397622605))))), v165)), v165))
			v211 = base.F64_add(v164, v210)
			if v150 == int32(0) {
				v214 = int32(1)
				v218 = base.F64_convert_i32_s(v214 - v140<<(uint(v214)%32))
				v223 = base.F64_add(v164, base.F64_sub(v210, base.F64_div(base.F64_mul(v211, v211), base.F64_add(v211, v218))))
				v225 = base.F64_sub(v218, base.F64_add(v223, v223))
				if v166 != 0 {
					v227 = v225
				} else {
					v227 = base.F64_neg(v225)
				}
				v253 = v227
			} else {
				if v140 != 0 {
					v229 = base.F64_div(float64(-1), v211)
					v231 = int64(-4294967296)
					v233 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v229) & v231)
					v237 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v211) & v231)
					v249 = base.F64_add(base.F64_mul(v229, base.F64_add(base.F64_mul(v233, base.F64_sub(v210, base.F64_sub(v237, v164))), base.F64_add(base.F64_mul(v233, v237), float64(1)))), v233)
				} else {
					v249 = v211
				}
				v253 = v249
			}
			v254 = v253
		}
	}
	m.G0 = v6 + int32(16)
	return v254
}
func F_tblspc_identify(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = l0 & int32(240)
	if v5 == int32(16) {
		v8 = int32(_a_F_tblspc_identify_0)
	} else {
		v8 = int32(0)
	}
	if v5 != 0 {
		v10 = v8
	} else {
		v10 = int32(_a_F_tblspc_identify_1)
	}
	return v10
}
func F_tblspc_redo(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+48)))
	v11 = v9 & int32(240)
	if v11 != 0 {
		if v11 == int32(16) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
			v16 = F_EmitProcSignalBarrier(m, int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_WaitForProcSignalBarrier(m, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v22 = F_destroy_tablespace_directories(m, v20, int32(1))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						if v22 != 0 {
							m.G0 = v6 + int32(32)
							return
						} else {
							v25 = int32(0)
							v27 = F_GetConflictingVirtualXIDs(m, v25, v25)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								v29 = int32(1)
								F_ResolveRecoveryConflictWithVirtualXIDs(m, v27, v29, int32(134217773), v29)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
									v36 = F_destroy_tablespace_directories(m, v34, int32(1))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										if v36 != 0 {
											m.G0 = v6 + int32(32)
											return
										} else {
											v40 = F_errstart(m, int32(15), int32(0))
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return
											} else {
												if v40 == int32(0) {
													m.G0 = v6 + int32(32)
													return
												} else {
													F_errcode(m, int32(325))
													mBase = m.M
													v46 = m.ExcPending
													if v46 != 0 {
														return
													} else {
														v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
														*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v47
														F_errmsg(m, int32(_a_F_tblspc_redo_0), v6+int32(16))
														mBase = m.M
														v53 = m.ExcPending
														if v53 != 0 {
															return
														} else {
															F_errhint(m, int32(_a_F_tblspc_redo_1), int32(0))
															mBase = m.M
															v57 = m.ExcPending
															if v57 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_tblspc_redo_2), int32(1590), int32(_a_F_tblspc_redo_3))
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return
																} else {
																	m.G0 = v6 + int32(32)
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
				}
			}
		} else {
			F_errstart_cold(m, int32(24), int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v11
				F_errmsg_internal(m, int32(_a_F_tblspc_redo_4), v6)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_tblspc_redo_2), int32(1594), int32(_a_F_tblspc_redo_3))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
		F_create_tablespace_directories(m, v76+int32(4), v79)
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	}
}
func F_text2ltree(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v15 = F_DirectFunctionCall1Coll(m, int32(_a_F_text2ltree_0), int32(0), base.I64_extend_i32_u(v12))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v12)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v19 != v8 {
						F_pfree(m, v8)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return v15 & int64(4294967295)
						}
					} else {
						return v15 & int64(4294967295)
					}
				}
			}
		}
	}
}
func F_textlename(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(1756), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(base.I32_wrap_i64(v6) <= int32(0)))
	}
}
func F_textnename(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v14 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v44 = F_strlen(m, v13)
	mBase = m.M
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v45 != int32(950) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v20 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v31 = int32(1)
	if v14&v31 != 0 {
		v43 = int32(base.Ui32(v14)>>(uint(v31)%32)) - v31
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v23 = int32(16)
	goto L9
L8:
	;
	v23 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = int32(4)
	goto L12
L11:
	;
	v30 = v23
	goto L12
L12:
	;
	v43 = v30
	goto L3
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v151 != v9 {
		goto L51
	} else {
		goto L52
	}
L15:
	;
	v140 = int32(1)
	if v14&v140 != 0 {
		goto L47
	} else {
		goto L48
	}
L16:
	;
	if v45 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v43 != v44 {
		v150 = int32(1)
		goto L14
	} else {
		goto L25
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_textnename_0), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(_a_F_textnename_1), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_textnename_2), int32(1337), int32(_a_F_textnename_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v70 = int32(1)
	if v14&v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v74 = v70
	goto L28
L27:
	;
	v74 = int32(4)
	goto L28
L28:
	;
	v75 = v9 + v74
	if base.Ui32(int32(4)) <= base.Ui32(v43) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v150 = base.B2i32(v137 != int32(0))
	goto L14
L30:
	;
	v137 = int32(0)
	goto L29
L31:
	;
	v111 = v106
	v112 = v107
	v113 = v108
	goto L41
L32:
	;
	if (v75|v13)&int32(3) != 0 {
		v106 = v75
		v107 = v13
		v108 = v43
		goto L31
	} else {
		goto L35
	}
L33:
	;
	v99 = v75
	v100 = v13
	v101 = v43
	goto L34
L34:
	;
	if v101 == int32(0) {
		goto L30
	} else {
		goto L40
	}
L35:
	;
	v83 = v75
	v84 = v13
	v85 = v43
	goto L36
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v88 != v89 {
		v106 = v83
		v107 = v84
		v108 = v85
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v99 = v94
	v100 = v92
	v101 = v96
	goto L34
L38:
	;
	v91 = int32(4)
	v92 = v84 + v91
	v94 = v83 + v91
	v96 = v85 - v91
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v83 = v94
		v84 = v92
		v85 = v96
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v106 = v99
	v107 = v100
	v108 = v101
	goto L31
L41:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v116 == v117 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v137 = v116 - v117
	goto L29
L43:
	;
	v119 = int32(1)
	v124 = v113 - v119
	if v124 != 0 {
		v111 = v111 + v119
		v112 = v112 + v119
		v113 = v124
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	goto L30
L47:
	;
	v144 = v140
	goto L49
L48:
	;
	v144 = int32(4)
	goto L49
L49:
	;
	v146 = F_varstr_cmp(m, v9+v144, v43, v13, v44, v45)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v150 = base.B2i32(v146 != int32(0))
	goto L14
L51:
	;
	F_pfree(m, v9)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return base.I64_extend_i32_u(v150)
L54:
	;
	goto L53
}
func F_textoverlay(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v12 = F_text_overlay(m, v3, v8, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v12)
			}
		}
	}
}
func F_textregexne(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14379(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_textregexreplace_extended_no_n(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_textregexreplace_extended(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_textregexreplace_noopt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v11 = F_pg_detoast_datum_packed(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v17 = F_replace_text_regexp(m, v3, v8, v11, int32(3), v14, int32(0), int32(1))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v17)
				}
			}
		}
	}
}
func F_textregexsubstr(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = F_RE_compile_and_cache(m, v18, int32(3), v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = int32(1)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				v28 = v26 & v24
				if v28 != 0 {
					v29 = v24
				} else {
					v29 = int32(4)
				}
				if v26 == int32(1) {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
					if v37 == int32(18) {
						v40 = int32(16)
					} else {
						v40 = int32(0)
					}
					if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v47 = int32(4)
					} else {
						v47 = v40
					}
					v58 = v47
				} else {
					v48 = int32(1)
					if v28 != 0 {
						v58 = int32(base.Ui32(v26)>>(uint(v48)%32)) - v48
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v61 = F_palloc_mul(m, int32(4), v58+int32(1))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int64(0)
				} else {
					v63 = F_pg_mb2wchar_with_len(m, v13+v29, v61, v58)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int64(0)
					} else {
						v67 = F_RE_wchar_execute(m, v61, v63, int32(0), int32(2), v10)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v61)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								if v67 == int32(0) {
									v73 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v73)
									v107 = int64(0)
									m.G0 = v10 + int32(16)
									return v107
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, _c_F_textregexsubstr[0]))
									if v79 != 0 {
										v80 = v10 | int32(8)
									} else {
										v80 = v10
									}
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
									if int32(0) <= v81 {
										if v79 != 0 {
											v86 = int32(12)
										} else {
											v86 = int32(4)
										}
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v10|v86)))
										if int32(0) <= v88 {
											v103 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), base.I64_extend_i32_u(v13), base.I64_extend_i32_s(v81+int32(1)), base.I64_extend_i32_s(v88-v81))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int64(0)
											} else {
												v107 = v103
												m.G0 = v10 + int32(16)
												return v107
											}
										} else {
											v92 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v92)
											v107 = int64(0)
											m.G0 = v10 + int32(16)
											return v107
										}
									} else {
										v92 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v92)
										v107 = int64(0)
										m.G0 = v10 + int32(16)
										return v107
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
func F_tfuncFetchRows(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v152 int64
	_ = v152
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
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v280 int32
	_ = v280
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v310 int64
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v481 int64
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v497 int32
	_ = v497
	var v508 int32
	_ = v508
	var v520 int32
	_ = v520
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v549 int32
	_ = v549
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v597 int32
	_ = v597
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v742 int64
	_ = v742
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v771 int32
	_ = v771
	var v772 int64
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v799 int64
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v841 int32
	_ = v841
	var v852 int32
	_ = v852
	var v866 int32
	_ = v866
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v928 int32
	_ = v928
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v1005 int32
	_ = v1005
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1056 int32
	_ = v1056
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1075 int32
	_ = v1075
	var v1086 int32
	_ = v1086
	var v1118 int32
	_ = v1118
	var v1119 int64
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	v3 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(224)
	m.G0 = v34
	v38 = l0 + int32(156)
	v46 = v3
	v47 = v3
	v48 = v3
	v49 = v3
	v50 = v3
	v51 = v3
	v52 = v3
	v53 = v3
	v54 = v3
	v55 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v55 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v1118 = int32(m.ExcTag)
	v1119 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1118 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L7:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v77 = int32(_a_F_tfuncFetchRows_0)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[0]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[0])) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v38
	v90 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[1]))
	v94 = F_tuplestore_begin_heap(m, v90, v90, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v112 = v46
	v113 = v47
	v114 = v48
	v115 = v49
	v116 = v50
	v117 = v51
	v118 = v52
	v119 = v53
	goto L9
L9:
	;
	if v112 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v94
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[0])) = v99
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[2]))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[3]))
	goto L11
L11:
	;
	v106 = v34 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v34 + int32(28)
	goto L14
L12:
	;
	v112 = int32(0)
	v113 = l0 + int32(176)
	v114 = v76
	v115 = v104
	v116 = v102
	v117 = v78
	v118 = l0 + int32(180)
	v119 = v38
	goto L9
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[2])) = v34 + int32(32)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	m.T0[v126].(func(*base.Module, int32, int32))(m, l0, v129)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[3])) = v115
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[2])) = v116
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v1064 != 0 {
		goto L138
	} else {
		goto L139
	}
L18:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v152 = m.T0[v141].(func(*base.Module, int32, int32, int32) int64)(m, v140, l1, v34+int32(189))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+189)))
	if v154 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+80))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+64))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	m.T0[v161].(func(*base.Module, int32, int64))(m, l0, v152)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	v1005 = v54
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[3])) = v115
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[2])) = v116
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v1029 != 0 {
		goto L133
	} else {
		goto L134
	}
L23:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v178 = int32(0)
	goto L25
L24:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v416 == int32(0) {
		v589 = v415
		v597 = v54
		goto L58
	} else {
		goto L59
	}
L25:
	;
	v206 = int32(0)
	if v173 == v206 {
		v216 = v206
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v400 = F_text_to_cstring(m, base.I32_wrap_i64(v241))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L6
	} else {
		goto L56
	}
L27:
	;
	if v172 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v210 <= v178 {
		v216 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v216 = v212 + v178<<(uint(int32(2))%32)
	goto L27
L30:
	;
	goto L26
L31:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v224+v178<<(uint(int32(2))%32))))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v310 = m.T0[v299].(func(*base.Module, int32, int32, int32) int64)(m, v298, l1, v34+int32(190))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L43
	}
L32:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v226 == int32(0) {
		goto L24
	} else {
		goto L36
	}
L33:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if base.B2i32(v216 == int32(0))|base.B2i32(v221 <= v178) != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	if v224 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v241 = m.T0[v230].(func(*base.Module, int32, int32, int32) int64)(m, v229, l1, v34+int32(190))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+190)))
	if v243 != int32(1) {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errcode(m, int32(67108994))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errmsg(m, int32(_a_F_tfuncFetchRows_1), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errfinish(m, int32(_a_F_tfuncFetchRows_2), int32(390), int32(_a_F_tfuncFetchRows_3))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L3
L43:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+190)))
	if v312 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L6
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v372 = F_text_to_cstring(m, base.I32_wrap_i64(v310))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L51
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errcode(m, int32(67108994))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errmsg(m, int32(_a_F_tfuncFetchRows_4), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errfinish(m, int32(_a_F_tfuncFetchRows_2), int32(371), int32(_a_F_tfuncFetchRows_3))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	goto L3
L51:
	;
	if v297 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v376 = v374
	goto L54
L53:
	;
	v376 = int32(0)
	goto L54
L54:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	m.T0[v377].(func(*base.Module, int32, int32, int32))(m, l0, v376, v372)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v178 = v178 + int32(1)
	goto L25
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	m.T0[v390].(func(*base.Module, int32, int32))(m, l0, v400)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	goto L24
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = int64(1)
	v619 = int32(_a_F_tfuncFetchRows_0)
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[0]))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v589)+20))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v589)+16))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)+80))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)+64))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[0])) = v630
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v641 = m.T0[v632].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L6
	} else {
		goto L82
	}
L59:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v419 <= int32(0) {
		v589 = v415
		v597 = v54
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v415)+12))
	v427 = int32(0)
	v435 = v54
	v436 = v419
	goto L61
L61:
	;
	if v427 != v159 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v589 = v585
	v597 = v578
	goto L58
L63:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	v462 = v422 + v456<<(uint(int32(3))%32) + v427*int32(100)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v465+v427<<(uint(int32(2))%32))))
	if v469 != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v578 = v435
	v579 = v436
	goto L65
L65:
	;
	v583 = v427 + int32(1)
	if v583 < v579 {
		v427 = v583
		v435 = v578
		v436 = v579
		goto L61
	} else {
		goto L81
	}
L66:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v563
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	m.T0[v566].(func(*base.Module, int32, int32, int32))(m, l0, v565, v427)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L6
	} else {
		goto L80
	}
L67:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v481 = m.T0[v470].(func(*base.Module, int32, int32, int32) int64)(m, v469, l1, v34+int32(190))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L6
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v563 = v435
	v565 = v462 + int32(32)
	goto L66
L70:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+190)))
	if v483 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L6
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v559 = F_text_to_cstring(m, base.I32_wrap_i64(v481))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L6
	} else {
		goto L79
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errcode(m, int32(67108994))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errmsg(m, int32(_a_F_tfuncFetchRows_5), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v462 + int32(32)
	v535 = F_errdetail(m, int32(_a_F_tfuncFetchRows_6), v34+int32(16))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errfinish(m, int32(_a_F_tfuncFetchRows_2), int32(419), int32(_a_F_tfuncFetchRows_3))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	goto L3
L79:
	;
	v563 = v559
	v565 = v559
	goto L66
L80:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v578 = v563
	v579 = v577
	goto L65
L81:
	;
	goto L62
L82:
	;
	if v641 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	goto L86
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[0])) = v620
	v1005 = v597
	goto L22
L86:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v678 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L85
L88:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)+12))
	v681 = v679
	goto L90
L89:
	;
	v681 = int32(0)
	goto L90
L90:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[4]))
	if v683 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_ProcessInterrupts(m)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+8))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	m.T0[v696].(func(*base.Module, int32))(m, v694)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L6
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	v707 = int32(0)
	if v707 < v627 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v713 = v707
	v722 = v681
	goto L99
L97:
	;
	goto L98
L98:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_tuplestore_putvalues(m, v928, v626, v622, v621)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L6
	} else {
		goto L129
	}
L99:
	;
	if v713 == v625 {
		goto L105
	} else {
		goto L106
	}
L100:
	;
	goto L98
L101:
	;
	v895 = v713 + int32(1)
	if v895 != v627 {
		v713 = v895
		v722 = v893
		goto L99
	} else {
		goto L128
	}
L102:
	;
	v881 = v722 + int32(4)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)+12))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	if base.Ui32(v881) < base.Ui32(v884+v885<<(uint(int32(2))%32)) {
		goto L125
	} else {
		goto L126
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L6
	} else {
		goto L121
	}
L104:
	;
	if v722 != 0 {
		goto L102
	} else {
		goto L120
	}
L105:
	;
	v742 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v742 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v622+v625<<(uint(int32(3))%32)))) = base.I64_extend32_s(v742)
	v748 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v625+v621))) = uint8(v748)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	v756 = v626 + v750<<(uint(int32(3))%32) + v713*int32(100)
	v758 = v756 + int32(28)
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)+76))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v758)+68))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v628)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v771 = v34 + int32(191)
	v772 = m.T0[v761].(func(*base.Module, int32, int32, int32, int32, int32) int64)(m, l0, v713, v760, v759, v771)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	v776 = v622 + v713<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v776))) = v772
	v779 = int32(0)
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+191)))
	if base.B2i32(v722 == v779)|base.B2i32(v781 != int32(1)) == v779 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v713+v621))) = uint8(v821)
	goto L104
L110:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v816 = F_bms_is_member(m, v713, v807)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L6
	} else {
		goto L118
	}
L111:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	if v787 == int32(0) {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v781 != 0 {
		goto L110
	} else {
		goto L117
	}
L114:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v787)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v799 = m.T0[v790].(func(*base.Module, int32, int32, int32) int64)(m, v787, l1, v771)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v776))) = v799
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+191)))
	if v802 != 0 {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v821 = int32(0)
	goto L109
L117:
	;
	v821 = int32(0)
	goto L109
L118:
	;
	if v816 != 0 {
		goto L103
	} else {
		goto L119
	}
L119:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+191)))
	v821 = v818
	goto L109
L120:
	;
	v893 = int32(0)
	goto L101
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errcode(m, int32(67108994))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v756 + int32(32)
	F_errmsg(m, int32(_a_F_tfuncFetchRows_7), v34)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_errfinish(m, int32(_a_F_tfuncFetchRows_2), int32(509), int32(_a_F_tfuncFetchRows_8))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	goto L3
L125:
	;
	v890 = v881
	goto L127
L126:
	;
	v890 = int32(0)
	goto L127
L127:
	;
	v893 = v890
	goto L101
L128:
	;
	goto L100
L129:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_MemoryContextReset(m, v939)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v628)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	v959 = m.T0[v950].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	if v959 != 0 {
		goto L86
	} else {
		goto L132
	}
L132:
	;
	goto L87
L133:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v114)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	m.T0[v1030].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L6
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tfuncFetchRows[0])) = v117
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_MemoryContextReset(m, v1046)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L6
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	goto L135
L137:
	;
	m.G0 = v34 + int32(224)
	return
L138:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v114)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	m.T0[v1065].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L6
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+196)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v34)+200)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v34)+204)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+212)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v34)+216)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v34)+220)) = v119
	F_pg_re_throw(m)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L6
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	goto L5
L143:
	;
	v1123 = int32(v1119)
	m.G0 = v34
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+4))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1123)))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1126)))
	if v34+int32(28) == v1129 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	m.ExcPending = 1
	goto L152
L145:
	;
	if v1133 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+4))
	v1133 = v1131
	goto L148
L147:
	;
	v1133 = int32(0)
	goto L148
L148:
	;
	goto L145
L149:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v34)+220))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v34)+216))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v34)+212))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v34)+204))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v34)+200))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v34)+196))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v34)+192))
	v46 = v1125
	v47 = v1138
	v48 = v1135
	v49 = v1140
	v50 = v1139
	v51 = v1136
	v52 = v1137
	v53 = v1134
	v54 = v1141
	v55 = v1133
	goto L1
L150:
	;
	goto L151
L151:
	;
	F___wasm_longjmp(m, v1126, v1125)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	return
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tidge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return base.I64_extend_i32_u(base.B2i32(int32(0) <= v27))
}
func F_tidlt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return base.I64_extend_i32_u(int32(base.Ui32(v27) >> (uint(int32(31)) % 32)))
}
func F_timetztypmodin(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14385(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_tliOfPointInHistory(m *base.Module, l0 int64, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v3 = int32(0)
	if l1 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	return v54
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v11 = int32(0)
	if v11 < v8 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v14 = v8
	goto L7
L6:
	;
	v14 = v11
	goto L7
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v19 = v3
	goto L8
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15+v19<<(uint(int32(2))%32))))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v24)+8))
	if base.Ui64(v25) <= base.Ui64(l0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)+16))
	if base.Ui64(l0) <= base.Ui64(v27-int64(1)) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v32 = v19 + int32(1)
	if v32 != v14 {
		v19 = v32
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	goto L9
L15:
	;
	return int32(0)
L16:
	;
	F_errmsg_internal(m, int32(_a_F_tliOfPointInHistory_0), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_tliOfPointInHistory_1), int32(562), int32(_a_F_tliOfPointInHistory_2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tlist_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v9 <= int32(0) {
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
	v16 = v3
	goto L8
L7:
	;
	return v22
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v16<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = F_equal(m, l0, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	return int32(0)
L11:
	;
	if v24 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v29 = v16 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v29 < v30 {
		v16 = v29
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
}
func F_towlower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v119 int32
	_ = v119
	v2 = int32(0)
	if base.Ui32(int32(_a_F_towlower_0)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v119
L2:
	;
	v119 = l0
	goto L1
L3:
	;
	v12 = int32(255)
	v13 = l0 & v12
	v14 = int32(3)
	v15 = base.I32_div_u_s(v13, v14)
	v21 = int32(2)
	v23 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_c_F_towlower[0])))
	v24 = int32(8)
	v25 = int32(base.Ui32(l0) >> (uint(v24) % 32))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_towlower[1]))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v26*int32(86))+uint32(_c_F_towlower[1]))))
	v35 = base.I32_rem_u_s(int32(base.Ui32(v23*v30)>>(uint(int32(11))%32)), int32(6))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_towlower[2]))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(v21)%32)+v38<<(uint(v21)%32))+uint32(_c_F_towlower[3])))
	v44 = v42 >> (uint(v24) % 32)
	v46 = v42 & v12
	if base.Ui32(v46) <= base.Ui32(int32(1)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v119 = v44&(int32(0)-(v2^v46)) + l0
	goto L1
L5:
	;
	goto L6
L6:
	;
	v55 = v44 & int32(255)
	if v55 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v62 = int32(base.Ui32(v44) >> (uint(int32(8)) % 32))
	v63 = v55
	goto L8
L8:
	;
	v69 = int32(1)
	v70 = int32(base.Ui32(v63) >> (uint(v69) % 32))
	v71 = v70 + v62
	v73 = v71 << (uint(v69) % 32)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_towlower[4]))))
	if v74 == v13 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_towlower[5]))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_c_F_towlower[3])))
	v83 = v81 & int32(255)
	if base.Ui32(v83) <= base.Ui32(int32(1)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v97 = base.B2i32(base.Ui32(v13) < base.Ui32(v74))
	if base.Ui32(v13) < base.Ui32(v74) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v119 = (int32(0)-(v2^v83))&(v81>>(uint(int32(8))%32)) + l0
	goto L1
L14:
	;
	goto L15
L15:
	;
	goto L17
L17:
	;
	goto L18
L18:
	;
	v119 = int32(1) + l0
	goto L1
L19:
	;
	v98 = v62
	goto L21
L20:
	;
	v98 = v71
	goto L21
L21:
	;
	if base.Ui32(v13) < base.Ui32(v74) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v100 = v70
	goto L24
L23:
	;
	v100 = v63 - v70
	goto L24
L24:
	;
	if v100 != 0 {
		v62 = v98
		v63 = v100
		goto L8
	} else {
		goto L25
	}
L25:
	;
	goto L9
}
func F_trackitem_compare_frequencies_desc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	return v5 - v7
}
func F_transfer_first_span(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v4 = l3
	v11 = l1 + int32(16)
	v14 = v11 + l2<<(uint(int32(2))%32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(0)
	v19 = base.AtomicRmwOr32(m, v16, int32(_a_F_transfer_first_span_0), v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+1468))
	if v20 != v22 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return base.B2i32(v15 != int32(0))
L4:
	;
	v27 = F_LWLockAcquire(m, v21+int32(1476), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v41 = int32(base.Ui32(v15) >> (uint(int32(27)) % 32))
	v44 = l0 + v41*int32(20)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	return int32(0)
L8:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v33+int32(1476))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v49 = v45
	goto L13
L12:
	;
	v46 = F_get_segment_by_index(m, l0, v41)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v50 = v49 + v15&int32(134217727)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v51
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v49 = v48
	goto L13
L15:
	;
	v53 = int32(0)
	v56 = base.AtomicRmwOr32(m, v53, int32(_a_F_transfer_first_span_0), v53)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+1468))
	if v57 != v59 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v93 = v11 + v4<<(uint(int32(2))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v15
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	if v97 != 0 {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v64 = F_LWLockAcquire(m, v58+int32(1476), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v74 = int32(base.Ui32(v51) >> (uint(int32(27)) % 32))
	v77 = l0 + v74*int32(20)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v78 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v68+int32(1476))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v82 = v78
	goto L26
L25:
	;
	v79 = F_get_segment_by_index(m, l0, v74)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82+v51&int32(134217727))+4)) = int32(0)
	goto L17
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v82 = v81
	goto L26
L28:
	;
	v98 = int32(0)
	v101 = base.AtomicRmwOr32(m, v98, int32(_a_F_transfer_first_span_0), v98)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+1468))
	if v102 != v104 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v50)+30)) = uint16(v4)
	goto L3
L31:
	;
	v109 = F_LWLockAcquire(m, v103+int32(1476), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v119 = int32(base.Ui32(v97) >> (uint(int32(27)) % 32))
	v122 = l0 + v119*int32(20)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	if v123 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v113+int32(1476))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v127 = v123
	goto L39
L38:
	;
	v124 = F_get_segment_by_index(m, l0, v119)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127+v97&int32(134217727))+4)) = v15
	goto L30
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v127 = v126
	goto L39
}
func F_transformDistinctClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	v5 = int32(0)
	if l2 == v5 {
		v70 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v73 == int32(0) {
		v111 = v70
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v11 <= int32(0) {
		v70 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v5
	v19 = v5
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v18<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = F_get_sortgroupclause_tle(m, v26, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L14
	}
L6:
	;
	return int32(0)
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+26)))
	if v32 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = F_copyObjectImpl(m, v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L5
L11:
	;
	v37 = F_lappend(m, v19, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v40 = v18 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v40 < v41 {
		v18 = v40
		v19 = v37
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v70 = v37
	goto L1
L14:
	;
	F_errcode(m, int32(_a_F_transformDistinctClause_0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if l3 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v52 = int32(_a_F_transformDistinctClause_1)
	goto L18
L17:
	;
	v52 = int32(_a_F_transformDistinctClause_2)
	goto L18
L18:
	;
	F_errmsg(m, v52, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v57 = F_exprLocation(m, v56)
	mBase = m.M
	F_parser_errposition(m, l0, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_transformDistinctClause_3), int32(3021), int32(_a_F_transformDistinctClause_4))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	if v111 != 0 {
		goto L32
	} else {
		goto L33
	}
L23:
	;
	v76 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v77 <= v76 {
		v111 = v70
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v84 = v76
	v85 = v70
	goto L25
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v84<<(uint(int32(2))%32))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+26)))
	if v93 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v111 = v101
	goto L22
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v98 = F_exprLocation(m, v97)
	mBase = m.M
	v99 = F_addTargetToGroupList(m, l0, v92, v85, v96, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	v101 = v85
	goto L29
L29:
	;
	v103 = v84 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v103 < v104 {
		v84 = v103
		v85 = v101
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v101 = v99
	goto L29
L31:
	;
	goto L26
L32:
	;
	return v111
L33:
	;
	goto L34
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	if l3 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v124 = int32(_a_F_transformDistinctClause_5)
	goto L39
L38:
	;
	v124 = int32(_a_F_transformDistinctClause_6)
	goto L39
L39:
	;
	F_errmsg(m, v124, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_transformDistinctClause_3), int32(3052), int32(_a_F_transformDistinctClause_4))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformReturningClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v18 = v16
	goto L3
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	if l2 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L22
	} else {
		goto L72
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L22
	} else {
		goto L67
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L22
	} else {
		goto L64
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L22
	} else {
		goto L59
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L22
	} else {
		goto L54
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v19 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	m.G0 = v13 - int32(-64)
	return
L12:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v74 != 0 {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v32 = int32(0)
	goto L15
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v32<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	switch v40 {
	case 0:
		goto L19
	case 1:
		goto L18
	default:
		goto L6
	}
L16:
	;
	goto L12
L17:
	;
	v48 = int32(0)
	v51 = F_refnameNamespaceItem(m, l0, v48, v47, int32(-1), v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	if v44 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v41 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v42
	v47 = v42
	goto L17
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v45
	v47 = v45
	goto L17
L22:
	;
	return
L23:
	;
	if v51 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v57 = int32(2)
	goto L27
L26:
	;
	v57 = int32(1)
	goto L27
L27:
	;
	F_addNSItemForReturning(m, l0, v53, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v61 = v32 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v61 < v62 {
		v32 = v61
		goto L15
	} else {
		goto L29
	}
L29:
	;
	goto L16
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	if v87 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v75 = int32(0)
	v79 = F_refnameNamespaceItem(m, l0, v75, int32(_a_F_transformReturningClause_0), int32(-1), v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L22
	} else {
		goto L32
	}
L32:
	;
	if v79 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v81 = int32(_a_F_transformReturningClause_0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v81
	F_addNSItemForReturning(m, l0, v81, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v104 = F_transformTargetList(m, l0, v103, l3)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L22
	} else {
		goto L40
	}
L36:
	;
	v88 = int32(0)
	v92 = F_refnameNamespaceItem(m, l0, v88, int32(_a_F_transformReturningClause_1), int32(-1), v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	if v92 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v94 = int32(_a_F_transformReturningClause_1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v94
	F_addNSItemForReturning(m, l0, v94, int32(2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v104
	if v104 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_markTargetListOrigins(m, l0, v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+81)))
	if v111 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	F_resolveTargetListUnknowns(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L22
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v118 = int32(0)
	if base.B2i32(v117 == v118)|base.B2i32(v18 <= v118) != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v128
	goto L11
L48:
	;
	v128 = int32(0)
	goto L50
L49:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v18 < v125 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L47
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v18
	goto L53
L52:
	;
	goto L53
L53:
	;
	v128 = v117
	goto L50
L54:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L22
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_transformReturningClause_2)
	F_errmsg(m, int32(_a_F_transformReturningClause_3), v11+int32(-32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L22
	} else {
		goto L56
	}
L56:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	F_parser_errposition(m, l0, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L22
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_transformReturningClause_4), int32(2742), int32(_a_F_transformReturningClause_5))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L22
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(_a_F_transformReturningClause_6)
	F_errmsg(m, int32(_a_F_transformReturningClause_3), v11+int32(-16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L22
	} else {
		goto L61
	}
L61:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	F_parser_errposition(m, l0, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_transformReturningClause_4), int32(2752), int32(_a_F_transformReturningClause_5))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L22
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v192
	F_errmsg_internal(m, int32(_a_F_transformReturningClause_7), v13)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_transformReturningClause_4), int32(2757), int32(_a_F_transformReturningClause_5))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L22
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v209
	F_errmsg(m, int32(_a_F_transformReturningClause_8), v11+int32(-48))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L22
	} else {
		goto L69
	}
L69:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	F_parser_errposition(m, l0, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L22
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_transformReturningClause_4), int32(2765), int32(_a_F_transformReturningClause_5))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L22
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L22
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_transformReturningClause_9), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L22
	} else {
		goto L74
	}
L74:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v238 = F_exprLocation(m, v237)
	mBase = m.M
	F_parser_errposition(m, l0, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L22
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_transformReturningClause_4), int32(2813), int32(_a_F_transformReturningClause_5))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L22
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_traverse_lacons(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v46 int32
	_ = v46
	F_check_stack_depth(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+l1<<(uint(int32(2))%32))))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	if v15 != int32(_a_F_traverse_lacons_0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = v14
	v24 = v15
	goto L6
L4:
	;
	goto L5
L5:
	;
	return
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.I32_extend16_s(v24) < v25 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)))
	if v46 != int32(_a_F_traverse_lacons_0) {
		v23 = v23 + int32(8)
		v24 = v46
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28 + int32(1)
	if l4 <= v28 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	F_traverse_lacons(m, l0, v40, l2, l3, l4)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v35 = l3 + v28<<(uint(int32(3))%32)
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23))))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v38
	goto L8
L13:
	;
	goto L8
L14:
	;
	goto L7
}
func F_trigram_qsort_signed(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
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
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v885 int32
	_ = v885
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v991 int32
	_ = v991
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1021 int32
	_ = v1021
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1087 int32
	_ = v1087
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v973 = l0
		v974 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	if base.Ui32(v974) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L316
	}
L3:
	;
	v19 = l0
	v20 = l1
	goto L4
L4:
	;
	v36 = v19 + int32(3)
	v38 = v20
	goto L6
L5:
	;
	v973 = v19
	v974 = v970
	goto L2
L6:
	;
	v54 = v38 * int32(3)
	if base.Ui32(v54) < base.Ui32(int32(4)) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v57 = v19 + v54
	v61 = v36
	goto L9
L9:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61-int32(3)))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v76 != v77 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v106 = v19 + int32(base.Ui32(v38)>>(uint(int32(1))%32))*int32(3)
	if v38 != int32(7) {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	goto L10
L12:
	;
	v98 = v61 + int32(3)
	if base.Ui32(v98) < base.Ui32(v57) {
		v61 = v98
		goto L9
	} else {
		goto L22
	}
L13:
	;
	if base.I32_extend8_s(v76) < base.I32_extend8_s(v77) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61-int32(2)))))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v84 != v85 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L11
L17:
	;
	if base.I32_extend8_s(v85) <= base.I32_extend8_s(v84) {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v92 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61-int32(1)))))
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+2)))
	if v93 < v92 {
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L12
L21:
	;
	goto L12
L22:
	;
	goto L1
L23:
	;
	v110 = v57 - int32(3)
	if base.Ui32(v38) < base.Ui32(int32(41)) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v508 = v106
	goto L25
L25:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v512)
	*(*uint8)(unsafe.Add(mBase, uint32(v508))) = uint8(v511)
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)) = uint8(v516)
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)) = uint8(v515)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)) = uint8(v520)
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+2)) = uint8(v519)
	v524 = v57 - int32(3)
	v526 = v36
	v528 = v36
	v530 = v524
	v534 = v524
	goto L242
L26:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	v417 = base.I32_extend8_s(v416)
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	v419 = base.I32_extend8_s(v418)
	if v416 != v418 {
		goto L193
	} else {
		goto L194
	}
L27:
	;
	v409 = v106
	v410 = v19
	v411 = v110
	goto L26
L28:
	;
	goto L29
L29:
	;
	v113 = int32(3)
	v114 = int32(base.Ui32(v38) >> (uint(v113) % 32))
	v116 = v114 * v113
	v117 = v19 + v116
	v119 = v114 * int32(6)
	v120 = v19 + v119
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v126 = base.I32_extend8_s(v125)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v128 = base.I32_extend8_s(v127)
	if v125 != v127 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v215 = v106 - v116
	v216 = v116 + v106
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v222 = base.I32_extend8_s(v221)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v224 = base.I32_extend8_s(v223)
	if v221 != v223 {
		goto L87
	} else {
		goto L88
	}
L31:
	;
	v214 = v204
	goto L30
L32:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v175 = base.I32_extend8_s(v174)
	if v174 != v126&int32(255) {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v143 = base.I32_extend8_s(v142)
	if v142 != v126&int32(255) {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	if v128 < v126 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v131 != v132 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L32
L38:
	;
	if base.I32_extend8_s(v132) <= base.I32_extend8_s(v131) {
		goto L32
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+2)))
	v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v117)+2)))
	if v138 <= v137 {
		goto L32
	} else {
		goto L42
	}
L41:
	;
	goto L33
L42:
	;
	goto L33
L43:
	;
	if v143 != v128 {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	if v143 <= v126 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	if v148 != v149 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v204 = v117
	goto L31
L48:
	;
	if base.I32_extend8_s(v149) <= base.I32_extend8_s(v148) {
		goto L43
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v117)+2)))
	v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v120)+2)))
	if v154 < v155 {
		v204 = v117
		goto L31
	} else {
		goto L52
	}
L51:
	;
	v204 = v117
	goto L31
L52:
	;
	goto L43
L53:
	;
	v214 = v19
	goto L30
L54:
	;
	if v143 <= v128 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	if v161 != v162 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v204 = v120
	goto L31
L58:
	;
	if base.I32_extend8_s(v162) <= base.I32_extend8_s(v161) {
		goto L53
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v167 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+2)))
	v168 = int32(*(*int8)(unsafe.Add(mBase, uint32(v120)+2)))
	if v167 < v168 {
		v204 = v120
		goto L31
	} else {
		goto L62
	}
L61:
	;
	v204 = v120
	goto L31
L62:
	;
	goto L53
L63:
	;
	if v175 != v128 {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	if v126 < v175 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	if v180 != v181 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v204 = v117
	goto L31
L68:
	;
	if base.I32_extend8_s(v180) < base.I32_extend8_s(v181) {
		goto L63
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v186 = int32(*(*int8)(unsafe.Add(mBase, uint32(v117)+2)))
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v120)+2)))
	if v187 < v186 {
		v204 = v117
		goto L31
	} else {
		goto L72
	}
L71:
	;
	v204 = v117
	goto L31
L72:
	;
	goto L63
L73:
	;
	v204 = v120
	goto L31
L74:
	;
	if v175 <= v128 {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	if v193 != v194 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v204 = v19
	goto L31
L78:
	;
	if base.I32_extend8_s(v194) <= base.I32_extend8_s(v193) {
		goto L73
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v199 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+2)))
	v200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v120)+2)))
	if v199 < v200 {
		v204 = v19
		goto L31
	} else {
		goto L82
	}
L81:
	;
	v204 = v19
	goto L31
L82:
	;
	goto L73
L83:
	;
	v311 = v110 - v119
	v312 = v110 - v116
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312))))
	v318 = base.I32_extend8_s(v317)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	v320 = base.I32_extend8_s(v319)
	if v317 != v319 {
		goto L140
	} else {
		goto L141
	}
L84:
	;
	v310 = v300
	goto L83
L85:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v271 = base.I32_extend8_s(v270)
	if v270 != v222&int32(255) {
		goto L117
	} else {
		goto L118
	}
L86:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v239 = base.I32_extend8_s(v238)
	if v238 != v222&int32(255) {
		goto L97
	} else {
		goto L98
	}
L87:
	;
	if v224 < v222 {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v227 != v228 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L85
L91:
	;
	if base.I32_extend8_s(v228) <= base.I32_extend8_s(v227) {
		goto L85
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(v215)+2)))
	v234 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+2)))
	if v234 <= v233 {
		goto L85
	} else {
		goto L95
	}
L94:
	;
	goto L86
L95:
	;
	goto L86
L96:
	;
	if v239 != v224 {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	if v239 <= v222 {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v244 != v245 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v300 = v106
	goto L84
L101:
	;
	if base.I32_extend8_s(v245) <= base.I32_extend8_s(v244) {
		goto L96
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v250 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+2)))
	v251 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+2)))
	if v250 < v251 {
		v300 = v106
		goto L84
	} else {
		goto L105
	}
L104:
	;
	v300 = v106
	goto L84
L105:
	;
	goto L96
L106:
	;
	v310 = v215
	goto L83
L107:
	;
	if v239 <= v224 {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v257 != v258 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v300 = v216
	goto L84
L111:
	;
	if base.I32_extend8_s(v258) <= base.I32_extend8_s(v257) {
		goto L106
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v263 = int32(*(*int8)(unsafe.Add(mBase, uint32(v215)+2)))
	v264 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+2)))
	if v263 < v264 {
		v300 = v216
		goto L84
	} else {
		goto L115
	}
L114:
	;
	v300 = v216
	goto L84
L115:
	;
	goto L106
L116:
	;
	if v271 != v224 {
		goto L127
	} else {
		goto L128
	}
L117:
	;
	if v222 < v271 {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v276 != v277 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v300 = v106
	goto L84
L121:
	;
	if base.I32_extend8_s(v276) < base.I32_extend8_s(v277) {
		goto L116
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v282 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+2)))
	v283 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+2)))
	if v283 < v282 {
		v300 = v106
		goto L84
	} else {
		goto L125
	}
L124:
	;
	v300 = v106
	goto L84
L125:
	;
	goto L116
L126:
	;
	v300 = v216
	goto L84
L127:
	;
	if v271 <= v224 {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v289 != v290 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v300 = v215
	goto L84
L131:
	;
	if base.I32_extend8_s(v290) <= base.I32_extend8_s(v289) {
		goto L126
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v295 = int32(*(*int8)(unsafe.Add(mBase, uint32(v215)+2)))
	v296 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+2)))
	if v295 < v296 {
		v300 = v215
		goto L84
	} else {
		goto L135
	}
L134:
	;
	v300 = v215
	goto L84
L135:
	;
	goto L126
L136:
	;
	v409 = v310
	v410 = v214
	v411 = v406
	goto L26
L137:
	;
	v406 = v396
	goto L136
L138:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v367 = base.I32_extend8_s(v366)
	if v366 != v318&int32(255) {
		goto L170
	} else {
		goto L171
	}
L139:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v335 = base.I32_extend8_s(v334)
	if v334 != v318&int32(255) {
		goto L150
	} else {
		goto L151
	}
L140:
	;
	if v320 < v318 {
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	if v323 != v324 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L138
L144:
	;
	if base.I32_extend8_s(v324) <= base.I32_extend8_s(v323) {
		goto L138
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v329 = int32(*(*int8)(unsafe.Add(mBase, uint32(v311)+2)))
	v330 = int32(*(*int8)(unsafe.Add(mBase, uint32(v312)+2)))
	if v330 <= v329 {
		goto L138
	} else {
		goto L148
	}
L147:
	;
	goto L139
L148:
	;
	goto L139
L149:
	;
	if v335 != v320 {
		goto L160
	} else {
		goto L161
	}
L150:
	;
	if v335 <= v318 {
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v340 != v341 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v396 = v312
	goto L137
L154:
	;
	if base.I32_extend8_s(v341) <= base.I32_extend8_s(v340) {
		goto L149
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v312)+2)))
	v347 = int32(*(*int8)(unsafe.Add(mBase, uint32(v110)+2)))
	if v346 < v347 {
		v396 = v312
		goto L137
	} else {
		goto L158
	}
L157:
	;
	v396 = v312
	goto L137
L158:
	;
	goto L149
L159:
	;
	v406 = v311
	goto L136
L160:
	;
	if v335 <= v320 {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v353 != v354 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v396 = v110
	goto L137
L164:
	;
	if base.I32_extend8_s(v354) <= base.I32_extend8_s(v353) {
		goto L159
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v359 = int32(*(*int8)(unsafe.Add(mBase, uint32(v311)+2)))
	v360 = int32(*(*int8)(unsafe.Add(mBase, uint32(v110)+2)))
	if v359 < v360 {
		v396 = v110
		goto L137
	} else {
		goto L168
	}
L167:
	;
	v396 = v110
	goto L137
L168:
	;
	goto L159
L169:
	;
	if v367 != v320 {
		goto L180
	} else {
		goto L181
	}
L170:
	;
	if v318 < v367 {
		goto L169
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v372 != v373 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v396 = v312
	goto L137
L174:
	;
	if base.I32_extend8_s(v372) < base.I32_extend8_s(v373) {
		goto L169
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v378 = int32(*(*int8)(unsafe.Add(mBase, uint32(v312)+2)))
	v379 = int32(*(*int8)(unsafe.Add(mBase, uint32(v110)+2)))
	if v379 < v378 {
		v396 = v312
		goto L137
	} else {
		goto L178
	}
L177:
	;
	v396 = v312
	goto L137
L178:
	;
	goto L169
L179:
	;
	v396 = v110
	goto L137
L180:
	;
	if v367 <= v320 {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v385 != v386 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v396 = v311
	goto L137
L184:
	;
	if base.I32_extend8_s(v386) <= base.I32_extend8_s(v385) {
		goto L179
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v391 = int32(*(*int8)(unsafe.Add(mBase, uint32(v311)+2)))
	v392 = int32(*(*int8)(unsafe.Add(mBase, uint32(v110)+2)))
	if v391 < v392 {
		v396 = v311
		goto L137
	} else {
		goto L188
	}
L187:
	;
	v396 = v311
	goto L137
L188:
	;
	goto L179
L189:
	;
	v508 = v505
	goto L25
L190:
	;
	v505 = v495
	goto L189
L191:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	v466 = base.I32_extend8_s(v465)
	if v465 != v417&int32(255) {
		goto L223
	} else {
		goto L224
	}
L192:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	v434 = base.I32_extend8_s(v433)
	if v433 != v417&int32(255) {
		goto L203
	} else {
		goto L204
	}
L193:
	;
	if v419 < v417 {
		goto L192
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+1)))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+1)))
	if v422 != v423 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L191
L197:
	;
	if base.I32_extend8_s(v423) <= base.I32_extend8_s(v422) {
		goto L191
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v428 = int32(*(*int8)(unsafe.Add(mBase, uint32(v410)+2)))
	v429 = int32(*(*int8)(unsafe.Add(mBase, uint32(v409)+2)))
	if v429 <= v428 {
		goto L191
	} else {
		goto L201
	}
L200:
	;
	goto L192
L201:
	;
	goto L192
L202:
	;
	if v434 != v419 {
		goto L213
	} else {
		goto L214
	}
L203:
	;
	if v434 <= v417 {
		goto L202
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+1)))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+1)))
	if v439 != v440 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v495 = v409
	goto L190
L207:
	;
	if base.I32_extend8_s(v440) <= base.I32_extend8_s(v439) {
		goto L202
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v445 = int32(*(*int8)(unsafe.Add(mBase, uint32(v409)+2)))
	v446 = int32(*(*int8)(unsafe.Add(mBase, uint32(v411)+2)))
	if v445 < v446 {
		v495 = v409
		goto L190
	} else {
		goto L211
	}
L210:
	;
	v495 = v409
	goto L190
L211:
	;
	goto L202
L212:
	;
	v505 = v410
	goto L189
L213:
	;
	if v434 <= v419 {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+1)))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+1)))
	if v452 != v453 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v495 = v411
	goto L190
L217:
	;
	if base.I32_extend8_s(v453) <= base.I32_extend8_s(v452) {
		goto L212
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v458 = int32(*(*int8)(unsafe.Add(mBase, uint32(v410)+2)))
	v459 = int32(*(*int8)(unsafe.Add(mBase, uint32(v411)+2)))
	if v458 < v459 {
		v495 = v411
		goto L190
	} else {
		goto L221
	}
L220:
	;
	v495 = v411
	goto L190
L221:
	;
	goto L212
L222:
	;
	if v466 != v419 {
		goto L233
	} else {
		goto L234
	}
L223:
	;
	if v417 < v466 {
		goto L222
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+1)))
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+1)))
	if v471 != v472 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v495 = v409
	goto L190
L227:
	;
	if base.I32_extend8_s(v471) < base.I32_extend8_s(v472) {
		goto L222
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v477 = int32(*(*int8)(unsafe.Add(mBase, uint32(v409)+2)))
	v478 = int32(*(*int8)(unsafe.Add(mBase, uint32(v411)+2)))
	if v478 < v477 {
		v495 = v409
		goto L190
	} else {
		goto L231
	}
L230:
	;
	v495 = v409
	goto L190
L231:
	;
	goto L222
L232:
	;
	v495 = v411
	goto L190
L233:
	;
	if v466 <= v419 {
		goto L232
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+1)))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+1)))
	if v484 != v485 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v495 = v410
	goto L190
L237:
	;
	if base.I32_extend8_s(v485) <= base.I32_extend8_s(v484) {
		goto L232
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v490 = int32(*(*int8)(unsafe.Add(mBase, uint32(v410)+2)))
	v491 = int32(*(*int8)(unsafe.Add(mBase, uint32(v411)+2)))
	if v490 < v491 {
		v495 = v410
		goto L190
	} else {
		goto L241
	}
L240:
	;
	v495 = v410
	goto L190
L241:
	;
	goto L232
L242:
	;
	if base.Ui32(v530) < base.Ui32(v528) {
		v596 = v526
		v598 = v528
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v966 = int32(3)
	v967 = base.I32_div_u_s(v809, v966)
	F_trigram_qsort_signed(m, v57-v809, v967)
	mBase = m.M
	v970 = base.I32_div_u_s(v682, v966)
	if base.Ui32(int32(21)) <= base.Ui32(v682) {
		v38 = v970
		goto L6
	} else {
		goto L315
	}
L244:
	;
	if base.Ui32(v598) <= base.Ui32(v530) {
		goto L264
	} else {
		goto L265
	}
L245:
	;
	v543 = v526
	v545 = v528
	goto L246
L246:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	v559 = base.I32_extend8_s(v558)
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v558 != v560 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	v596 = v589
	v598 = v593
	goto L244
L248:
	;
	v593 = v545 + int32(3)
	if base.Ui32(v593) <= base.Ui32(v530) {
		v543 = v589
		v545 = v593
		goto L246
	} else {
		goto L261
	}
L249:
	;
	if v559 < base.I32_extend8_s(v560) {
		v589 = v543
		goto L248
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+1)))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v564 != v565 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v596 = v543
	v598 = v545
	goto L244
L253:
	;
	if base.I32_extend8_s(v564) < base.I32_extend8_s(v565) {
		v589 = v543
		goto L248
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+2)))
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	if v570 == v571 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v596 = v543
	v598 = v545
	goto L244
L257:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	*(*uint8)(unsafe.Add(mBase, uint32(v543))) = uint8(v559)
	*(*uint8)(unsafe.Add(mBase, uint32(v545))) = uint8(v573)
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+1)))
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v543)+1)) = uint8(v577)
	*(*uint8)(unsafe.Add(mBase, uint32(v545)+1)) = uint8(v576)
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+2)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v543)+2)) = uint8(v581)
	*(*uint8)(unsafe.Add(mBase, uint32(v545)+2)) = uint8(v580)
	v589 = v543 + int32(3)
	goto L248
L258:
	;
	goto L259
L259:
	;
	if base.I32_extend8_s(v571) <= base.I32_extend8_s(v570) {
		v596 = v543
		v598 = v545
		goto L244
	} else {
		goto L260
	}
L260:
	;
	v589 = v543
	goto L248
L261:
	;
	goto L247
L262:
	;
	goto L243
L263:
	;
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	*(*uint8)(unsafe.Add(mBase, uint32(v598))) = uint8(v629)
	*(*uint8)(unsafe.Add(mBase, uint32(v617))) = uint8(v950)
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+1)))
	v954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v598)+1)) = uint8(v954)
	*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)) = uint8(v953)
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+2)))
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v598)+2)) = uint8(v958)
	*(*uint8)(unsafe.Add(mBase, uint32(v617)+2)) = uint8(v957)
	v961 = int32(3)
	v526 = v596
	v528 = v598 + v961
	v530 = v617 - v961
	v534 = v621
	goto L242
L264:
	;
	v617 = v530
	v621 = v534
	goto L267
L265:
	;
	v670 = v530
	v674 = v534
	goto L266
L266:
	;
	v681 = v596 - v19
	v682 = v598 - v596
	if v681 < v682 {
		goto L284
	} else {
		goto L285
	}
L267:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617))))
	v629 = base.I32_extend8_s(v628)
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v628 != v630 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v670 = v663
	v674 = v661
	goto L266
L269:
	;
	v663 = v617 - int32(3)
	if base.Ui32(v598) <= base.Ui32(v663) {
		v617 = v663
		v621 = v661
		goto L267
	} else {
		goto L282
	}
L270:
	;
	if base.I32_extend8_s(v630) <= v629 {
		v661 = v621
		goto L269
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v634 != v635 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	goto L263
L274:
	;
	if base.I32_extend8_s(v634) < base.I32_extend8_s(v635) {
		goto L263
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+2)))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	if v640 == v641 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v661 = v621
	goto L269
L278:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	*(*uint8)(unsafe.Add(mBase, uint32(v617))) = uint8(v643)
	*(*uint8)(unsafe.Add(mBase, uint32(v621))) = uint8(v629)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)))
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)) = uint8(v647)
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+1)) = uint8(v646)
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+2)))
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v617)+2)) = uint8(v651)
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+2)) = uint8(v650)
	v661 = v621 - int32(3)
	goto L269
L279:
	;
	goto L280
L280:
	;
	if base.I32_extend8_s(v640) < base.I32_extend8_s(v641) {
		goto L263
	} else {
		goto L281
	}
L281:
	;
	v661 = v621
	goto L269
L282:
	;
	goto L268
L283:
	;
	v809 = v674 - v670
	v812 = v57 - v674 - int32(3)
	if base.Ui32(v809) < base.Ui32(v812) {
		goto L299
	} else {
		goto L300
	}
L284:
	;
	v684 = v681
	goto L286
L285:
	;
	v684 = v682
	goto L286
L286:
	;
	if v684 == int32(0) {
		goto L283
	} else {
		goto L287
	}
L287:
	;
	v687 = v598 - v684
	v689 = v684 & int32(3)
	v690 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v684) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v698 = int32(0)
	v701 = v690
	goto L291
L289:
	;
	v754 = v690
	goto L290
L290:
	;
	v768 = v690
	v770 = v754
	goto L295
L291:
	;
	v713 = v19 + v701
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	v715 = v701 + v687
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715))))
	*(*uint8)(unsafe.Add(mBase, uint32(v713))) = uint8(v716)
	*(*uint8)(unsafe.Add(mBase, uint32(v715))) = uint8(v714)
	v720 = v701 | int32(1)
	v721 = v19 + v720
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721))))
	v723 = v687 + v720
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723))))
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v724)
	*(*uint8)(unsafe.Add(mBase, uint32(v723))) = uint8(v722)
	v728 = v701 | int32(2)
	v729 = v19 + v728
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729))))
	v731 = v687 + v728
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731))))
	*(*uint8)(unsafe.Add(mBase, uint32(v729))) = uint8(v732)
	*(*uint8)(unsafe.Add(mBase, uint32(v731))) = uint8(v730)
	v736 = v701 | int32(3)
	v737 = v19 + v736
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737))))
	v739 = v687 + v736
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739))))
	*(*uint8)(unsafe.Add(mBase, uint32(v737))) = uint8(v740)
	*(*uint8)(unsafe.Add(mBase, uint32(v739))) = uint8(v738)
	v743 = int32(4)
	v744 = v701 + v743
	v746 = v698 + v743
	if v746 != v684&int32(-4) {
		v698 = v746
		v701 = v744
		goto L291
	} else {
		goto L293
	}
L292:
	;
	if v689 == int32(0) {
		goto L283
	} else {
		goto L294
	}
L293:
	;
	goto L292
L294:
	;
	v754 = v744
	goto L290
L295:
	;
	v782 = v19 + v770
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v782))))
	v784 = v770 + v687
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784))))
	*(*uint8)(unsafe.Add(mBase, uint32(v782))) = uint8(v785)
	*(*uint8)(unsafe.Add(mBase, uint32(v784))) = uint8(v783)
	v788 = int32(1)
	v791 = v768 + v788
	if v791 != v689 {
		v768 = v791
		v770 = v770 + v788
		goto L295
	} else {
		goto L297
	}
L296:
	;
	goto L283
L297:
	;
	goto L296
L298:
	;
	if base.Ui32(v809) < base.Ui32(v682) {
		goto L262
	} else {
		goto L313
	}
L299:
	;
	v814 = v809
	goto L301
L300:
	;
	v814 = v812
	goto L301
L301:
	;
	if v814 == int32(0) {
		goto L298
	} else {
		goto L302
	}
L302:
	;
	v817 = v57 - v814
	v819 = v814 & int32(3)
	v820 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v814) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v829 = int32(0)
	v832 = v820
	goto L306
L304:
	;
	v885 = v820
	goto L305
L305:
	;
	v897 = v820
	v901 = v885
	goto L310
L306:
	;
	v843 = v598 + v832
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v843))))
	v845 = v832 + v817
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	*(*uint8)(unsafe.Add(mBase, uint32(v843))) = uint8(v846)
	*(*uint8)(unsafe.Add(mBase, uint32(v845))) = uint8(v844)
	v850 = v832 | int32(1)
	v851 = v598 + v850
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851))))
	v853 = v850 + v817
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853))))
	*(*uint8)(unsafe.Add(mBase, uint32(v851))) = uint8(v854)
	*(*uint8)(unsafe.Add(mBase, uint32(v853))) = uint8(v852)
	v858 = v832 | int32(2)
	v859 = v598 + v858
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859))))
	v861 = v858 + v817
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	*(*uint8)(unsafe.Add(mBase, uint32(v859))) = uint8(v862)
	*(*uint8)(unsafe.Add(mBase, uint32(v861))) = uint8(v860)
	v866 = v832 | int32(3)
	v867 = v598 + v866
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867))))
	v869 = v866 + v817
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869))))
	*(*uint8)(unsafe.Add(mBase, uint32(v867))) = uint8(v870)
	*(*uint8)(unsafe.Add(mBase, uint32(v869))) = uint8(v868)
	v873 = int32(4)
	v874 = v832 + v873
	v876 = v829 + v873
	if v876 != v814&int32(-4) {
		v829 = v876
		v832 = v874
		goto L306
	} else {
		goto L308
	}
L307:
	;
	if v819 == int32(0) {
		goto L298
	} else {
		goto L309
	}
L308:
	;
	goto L307
L309:
	;
	v885 = v874
	goto L305
L310:
	;
	v912 = v598 + v901
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
	v914 = v901 + v817
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	*(*uint8)(unsafe.Add(mBase, uint32(v912))) = uint8(v915)
	*(*uint8)(unsafe.Add(mBase, uint32(v914))) = uint8(v913)
	v918 = int32(1)
	v921 = v897 + v918
	if v921 != v819 {
		v897 = v921
		v901 = v901 + v918
		goto L310
	} else {
		goto L312
	}
L311:
	;
	goto L298
L312:
	;
	goto L311
L313:
	;
	v940 = int32(3)
	v941 = base.I32_div_u_s(v682, v940)
	F_trigram_qsort_signed(m, v19, v941)
	mBase = m.M
	v944 = base.I32_div_u_s(v809, v940)
	v945 = v57 - v809
	if base.Ui32(int32(21)) <= base.Ui32(v809) {
		v19 = v945
		v20 = v944
		goto L4
	} else {
		goto L314
	}
L314:
	;
	v973 = v945
	v974 = v944
	goto L2
L315:
	;
	goto L7
L316:
	;
	v991 = int32(3)
	v1007 = v973 + v991
	goto L317
L317:
	;
	if base.Ui32(v1007) <= base.Ui32(v973) {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	goto L1
L319:
	;
	v1087 = v1007 + int32(3)
	if base.Ui32(v1087) < base.Ui32(v973+v974*v991) {
		v1007 = v1087
		goto L317
	} else {
		goto L334
	}
L320:
	;
	v1013 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1007))))
	v1021 = v1007
	goto L321
L321:
	;
	v1033 = v1021 - int32(3)
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	v1035 = base.I32_extend8_s(v1034)
	if v1034 != v1013&int32(255) {
		goto L324
	} else {
		goto L325
	}
L322:
	;
	goto L319
L323:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1021))) = uint8(v1035)
	*(*uint8)(unsafe.Add(mBase, uint32(v1033))) = uint8(v1013)
	*(*uint8)(unsafe.Add(mBase, uint32(v1021-int32(2)))) = uint8(v1056)
	*(*uint8)(unsafe.Add(mBase, uint32(v1021)+1)) = uint8(v1055)
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021)+2)))
	v1065 = v1021 - int32(1)
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1021)+2)) = uint8(v1066)
	*(*uint8)(unsafe.Add(mBase, uint32(v1065))) = uint8(v1063)
	if base.Ui32(v973) < base.Ui32(v1033) {
		v1021 = v1033
		goto L321
	} else {
		goto L333
	}
L324:
	;
	if v1035 < v1013 {
		goto L319
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021-int32(2)))))
	v1045 = base.I32_extend8_s(v1044)
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021)+1)))
	if v1044 != v1046 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021-int32(2)))))
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021)+1)))
	v1055 = v1040
	v1056 = v1041
	goto L323
L328:
	;
	v1048 = base.I32_extend8_s(v1046)
	if v1048 <= v1045 {
		v1055 = v1045
		v1056 = v1048
		goto L323
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1052 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1021-int32(1)))))
	v1053 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1021)+2)))
	if v1052 <= v1053 {
		goto L319
	} else {
		goto L332
	}
L331:
	;
	goto L319
L332:
	;
	v1055 = v1045
	v1056 = v1045
	goto L323
L333:
	;
	goto L322
L334:
	;
	goto L318
}
func F_trigram_qsort_signed_med3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v9 = base.I32_extend8_s(v8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v11 = base.I32_extend8_s(v10)
	if v8 != v10 {
		if v11 < v9 {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
			v26 = base.I32_extend8_s(v25)
			if v25 != v9&int32(255) {
				if v26 <= v9 {
					if v26 != v11 {
						if v26 <= v11 {
							return l0
						} else {
							v88 = l2
							return v88
						}
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
						if v44 != v45 {
							if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
								return l0
							} else {
								v88 = l2
								return v88
							}
						} else {
							v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
							v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
							if v50 < v51 {
								v88 = l2
								return v88
							} else {
								return l0
							}
						}
					}
				} else {
					v88 = l1
					return v88
				}
			} else {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
				if v31 != v32 {
					if base.I32_extend8_s(v32) <= base.I32_extend8_s(v31) {
						if v26 != v11 {
							if v26 <= v11 {
								return l0
							} else {
								v88 = l2
								return v88
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v44 != v45 {
								if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
									return l0
								} else {
									v88 = l2
									return v88
								}
							} else {
								v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
								v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
								if v50 < v51 {
									v88 = l2
									return v88
								} else {
									return l0
								}
							}
						}
					} else {
						v88 = l1
						return v88
					}
				} else {
					v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
					v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
					if v37 < v38 {
						v88 = l1
						return v88
					} else {
						if v26 != v11 {
							if v26 <= v11 {
								return l0
							} else {
								v88 = l2
								return v88
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v44 != v45 {
								if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
									return l0
								} else {
									v88 = l2
									return v88
								}
							} else {
								v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
								v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
								if v50 < v51 {
									v88 = l2
									return v88
								} else {
									return l0
								}
							}
						}
					}
				}
			}
		} else {
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
			v59 = base.I32_extend8_s(v58)
			if v58 != v9&int32(255) {
				if v9 < v59 {
					if v59 != v11 {
						if v59 <= v11 {
							v88 = l2
						} else {
							v88 = l0
						}
					} else {
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
						if v77 != v78 {
							if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
								v88 = l2
							} else {
								v88 = l0
							}
						} else {
							v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
							v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
							if v83 < v84 {
								v88 = l0
							} else {
								v88 = l2
							}
						}
					}
				} else {
					v88 = l1
				}
			} else {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
				if v64 != v65 {
					if base.I32_extend8_s(v64) < base.I32_extend8_s(v65) {
						if v59 != v11 {
							if v59 <= v11 {
								v88 = l2
							} else {
								v88 = l0
							}
						} else {
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v77 != v78 {
								if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
									v88 = l2
								} else {
									v88 = l0
								}
							} else {
								v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
								v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
								if v83 < v84 {
									v88 = l0
								} else {
									v88 = l2
								}
							}
						}
					} else {
						v88 = l1
					}
				} else {
					v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
					v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
					if v71 < v70 {
						v88 = l1
					} else {
						if v59 != v11 {
							if v59 <= v11 {
								v88 = l2
							} else {
								v88 = l0
							}
						} else {
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v77 != v78 {
								if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
									v88 = l2
								} else {
									v88 = l0
								}
							} else {
								v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
								v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
								if v83 < v84 {
									v88 = l0
								} else {
									v88 = l2
								}
							}
						}
					}
				}
			}
			return v88
		}
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v14 != v15 {
			if base.I32_extend8_s(v15) <= base.I32_extend8_s(v14) {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				v59 = base.I32_extend8_s(v58)
				if v58 != v9&int32(255) {
					if v9 < v59 {
						if v59 != v11 {
							if v59 <= v11 {
								v88 = l2
							} else {
								v88 = l0
							}
						} else {
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v77 != v78 {
								if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
									v88 = l2
								} else {
									v88 = l0
								}
							} else {
								v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
								v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
								if v83 < v84 {
									v88 = l0
								} else {
									v88 = l2
								}
							}
						}
					} else {
						v88 = l1
					}
				} else {
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v64 != v65 {
						if base.I32_extend8_s(v64) < base.I32_extend8_s(v65) {
							if v59 != v11 {
								if v59 <= v11 {
									v88 = l2
								} else {
									v88 = l0
								}
							} else {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v77 != v78 {
									if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
										v88 = l2
									} else {
										v88 = l0
									}
								} else {
									v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
									v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
									if v83 < v84 {
										v88 = l0
									} else {
										v88 = l2
									}
								}
							}
						} else {
							v88 = l1
						}
					} else {
						v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
						v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
						if v71 < v70 {
							v88 = l1
						} else {
							if v59 != v11 {
								if v59 <= v11 {
									v88 = l2
								} else {
									v88 = l0
								}
							} else {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v77 != v78 {
									if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
										v88 = l2
									} else {
										v88 = l0
									}
								} else {
									v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
									v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
									if v83 < v84 {
										v88 = l0
									} else {
										v88 = l2
									}
								}
							}
						}
					}
				}
				return v88
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				v26 = base.I32_extend8_s(v25)
				if v25 != v9&int32(255) {
					if v26 <= v9 {
						if v26 != v11 {
							if v26 <= v11 {
								return l0
							} else {
								v88 = l2
								return v88
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v44 != v45 {
								if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
									return l0
								} else {
									v88 = l2
									return v88
								}
							} else {
								v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
								v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
								if v50 < v51 {
									v88 = l2
									return v88
								} else {
									return l0
								}
							}
						}
					} else {
						v88 = l1
						return v88
					}
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v31 != v32 {
						if base.I32_extend8_s(v32) <= base.I32_extend8_s(v31) {
							if v26 != v11 {
								if v26 <= v11 {
									return l0
								} else {
									v88 = l2
									return v88
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v44 != v45 {
									if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
										return l0
									} else {
										v88 = l2
										return v88
									}
								} else {
									v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
									v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
									if v50 < v51 {
										v88 = l2
										return v88
									} else {
										return l0
									}
								}
							}
						} else {
							v88 = l1
							return v88
						}
					} else {
						v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
						v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
						if v37 < v38 {
							v88 = l1
							return v88
						} else {
							if v26 != v11 {
								if v26 <= v11 {
									return l0
								} else {
									v88 = l2
									return v88
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v44 != v45 {
									if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
										return l0
									} else {
										v88 = l2
										return v88
									}
								} else {
									v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
									v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
									if v50 < v51 {
										v88 = l2
										return v88
									} else {
										return l0
									}
								}
							}
						}
					}
				}
			}
		} else {
			v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
			v21 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v21 <= v20 {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				v59 = base.I32_extend8_s(v58)
				if v58 != v9&int32(255) {
					if v9 < v59 {
						if v59 != v11 {
							if v59 <= v11 {
								v88 = l2
							} else {
								v88 = l0
							}
						} else {
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v77 != v78 {
								if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
									v88 = l2
								} else {
									v88 = l0
								}
							} else {
								v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
								v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
								if v83 < v84 {
									v88 = l0
								} else {
									v88 = l2
								}
							}
						}
					} else {
						v88 = l1
					}
				} else {
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v64 != v65 {
						if base.I32_extend8_s(v64) < base.I32_extend8_s(v65) {
							if v59 != v11 {
								if v59 <= v11 {
									v88 = l2
								} else {
									v88 = l0
								}
							} else {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v77 != v78 {
									if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
										v88 = l2
									} else {
										v88 = l0
									}
								} else {
									v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
									v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
									if v83 < v84 {
										v88 = l0
									} else {
										v88 = l2
									}
								}
							}
						} else {
							v88 = l1
						}
					} else {
						v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
						v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
						if v71 < v70 {
							v88 = l1
						} else {
							if v59 != v11 {
								if v59 <= v11 {
									v88 = l2
								} else {
									v88 = l0
								}
							} else {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v77 != v78 {
									if base.I32_extend8_s(v78) <= base.I32_extend8_s(v77) {
										v88 = l2
									} else {
										v88 = l0
									}
								} else {
									v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
									v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
									if v83 < v84 {
										v88 = l0
									} else {
										v88 = l2
									}
								}
							}
						}
					}
				}
				return v88
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				v26 = base.I32_extend8_s(v25)
				if v25 != v9&int32(255) {
					if v26 <= v9 {
						if v26 != v11 {
							if v26 <= v11 {
								return l0
							} else {
								v88 = l2
								return v88
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v44 != v45 {
								if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
									return l0
								} else {
									v88 = l2
									return v88
								}
							} else {
								v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
								v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
								if v50 < v51 {
									v88 = l2
									return v88
								} else {
									return l0
								}
							}
						}
					} else {
						v88 = l1
						return v88
					}
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v31 != v32 {
						if base.I32_extend8_s(v32) <= base.I32_extend8_s(v31) {
							if v26 != v11 {
								if v26 <= v11 {
									return l0
								} else {
									v88 = l2
									return v88
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v44 != v45 {
									if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
										return l0
									} else {
										v88 = l2
										return v88
									}
								} else {
									v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
									v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
									if v50 < v51 {
										v88 = l2
										return v88
									} else {
										return l0
									}
								}
							}
						} else {
							v88 = l1
							return v88
						}
					} else {
						v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
						v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
						if v37 < v38 {
							v88 = l1
							return v88
						} else {
							if v26 != v11 {
								if v26 <= v11 {
									return l0
								} else {
									v88 = l2
									return v88
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v44 != v45 {
									if base.I32_extend8_s(v45) <= base.I32_extend8_s(v44) {
										return l0
									} else {
										v88 = l2
										return v88
									}
								} else {
									v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
									v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+2)))
									if v50 < v51 {
										v88 = l2
										return v88
									} else {
										return l0
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
func F_trivial_subqueryscan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	v9 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	switch v10 - v9 {
	case 0:
		v99 = v9
		goto L1
	case 1:
		goto L3
	default:
		goto L2
	}
L1:
	;
	return v99
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return int32(0)
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = v22
	goto L9
L8:
	;
	v23 = int32(0)
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v28 = v26
	goto L12
L11:
	;
	v28 = int32(0)
	goto L12
L12:
	;
	if v28 != v23 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v35 = int32(0)
	v40 = int32(1)
	goto L16
L16:
	;
	v42 = int32(0)
	if v21 == v42 {
		v52 = v42
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v91 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v91
	v99 = v91
	goto L1
L18:
	;
	if v25 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v46 <= v35 {
		v52 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v52 = v48 + v35<<(uint(int32(2))%32)
	goto L18
L21:
	;
	goto L17
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if base.B2i32(v52 == int32(0))|base.B2i32(v57 <= v35) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v60 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+26)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60+v35<<(uint(int32(2))%32))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+26)))
	if v64 != v69 {
		v99 = v20
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v71 == int32(0) {
		v99 = v20
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	switch v74 - int32(6) {
	case 0:
		goto L29
	case 1:
		goto L28
	default:
		v99 = v20
		goto L1
	}
L27:
	;
	v86 = int32(1)
	v35 = v35 + v86
	v40 = v40 + v86
	goto L16
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v80 = F_equal(m, v71, v79)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+8)))
	if v40 == v77 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v99 = v20
	goto L1
L31:
	;
	return int32(0)
L32:
	;
	if v80 == int32(0) {
		v99 = v20
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L27
}
func F_truncate_useless_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v643 int32
	_ = v643
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v707 int32
	_ = v707
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v756 int32
	_ = v756
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	v4 = int32(0)
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v16 = v14
	goto L3
L2:
	;
	v16 = int32(0)
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if l2 == v17 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v71 == v16 {
		v797 = l2
		goto L29
	} else {
		goto L30
	}
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v71 = v64
	goto L4
L6:
	;
	if v17 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v19 = int32(0)
	if base.B2i32(l2 == v19)|base.B2i32(v17 == v19) != 0 {
		v71 = v4
		goto L4
	} else {
		goto L10
	}
L9:
	;
	v71 = v4
	goto L4
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v25 = int32(0)
	if v25 < v24 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v28 = v24
	goto L13
L12:
	;
	v28 = v25
	goto L13
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v33 = v4
	goto L14
L14:
	;
	if v33 < v29 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v49 = v45 + v33<<(uint(int32(2))%32)
	goto L18
L17:
	;
	v49 = int32(0)
	goto L18
L18:
	;
	if v33 == v28 {
		v71 = v28
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v49 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v71 = v33
	goto L4
L21:
	;
	goto L22
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v53 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v71 = v33
	goto L4
L24:
	;
	goto L25
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53+v33<<(uint(int32(2))%32))))
	if v56 != v60 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v71 = v33
	goto L4
L27:
	;
	v33 = v33 + int32(1)
	goto L14
L29:
	;
	return v797
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if l2 == v79 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v134 == v16 {
		v797 = l2
		goto L29
	} else {
		goto L56
	}
L32:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v134 = v129
	goto L31
L33:
	;
	if v79 != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v82 = int32(0)
	if base.B2i32(l2 == v82)|base.B2i32(v79 == v82) != 0 {
		v134 = v82
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v134 = int32(0)
	goto L31
L37:
	;
	v88 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v88 < v89 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v93 = v89
	goto L40
L39:
	;
	v93 = v88
	goto L40
L40:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v98 = v88
	goto L41
L41:
	;
	if v98 < v94 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v114 = v110 + v98<<(uint(int32(2))%32)
	goto L45
L44:
	;
	v114 = int32(0)
	goto L45
L45:
	;
	if v98 == v93 {
		v134 = v93
		goto L31
	} else {
		goto L46
	}
L46:
	;
	if v114 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v134 = v98
	goto L31
L48:
	;
	goto L49
L49:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v118 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v134 = v98
	goto L31
L51:
	;
	goto L52
L52:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v118+v98<<(uint(int32(2))%32))))
	if v121 != v125 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v134 = v98
	goto L31
L54:
	;
	v98 = v98 + int32(1)
	goto L41
L56:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if l2 == v144 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	if v202 == v16 {
		v797 = l2
		goto L29
	} else {
		goto L82
	}
L58:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v202 = v194
	goto L57
L59:
	;
	if v144 != 0 {
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v147 = int32(0)
	if base.B2i32(l2 == v147)|base.B2i32(v144 == v147) != 0 {
		v202 = v147
		goto L57
	} else {
		goto L63
	}
L62:
	;
	v202 = int32(0)
	goto L57
L63:
	;
	v153 = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v153 < v154 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v158 = v154
	goto L66
L65:
	;
	v158 = v153
	goto L66
L66:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v163 = v153
	goto L67
L67:
	;
	if v163 < v159 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v179 = v175 + v163<<(uint(int32(2))%32)
	goto L71
L70:
	;
	v179 = int32(0)
	goto L71
L71:
	;
	if v163 == v158 {
		v202 = v158
		goto L57
	} else {
		goto L72
	}
L72:
	;
	if v179 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v202 = v163
	goto L57
L74:
	;
	goto L75
L75:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v183 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v202 = v163
	goto L57
L77:
	;
	goto L78
L78:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183+v163<<(uint(int32(2))%32))))
	if v186 != v190 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v202 = v163
	goto L57
L80:
	;
	v163 = v163 + int32(1)
	goto L67
L82:
	;
	v209 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if base.B2i32(v210 == v209)|base.B2i32(l2 == v209) != 0 {
		v285 = v209
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v285 == v16 {
		v797 = l2
		goto L29
	} else {
		goto L103
	}
L84:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v216 <= int32(0) {
		v285 = v209
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v222 = v209
	goto L86
L86:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232+v222<<(uint(int32(2))%32))))
	v237 = int32(0)
	if v210 == v237 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v285 = v279
	goto L83
L88:
	;
	if v275 == int32(0) {
		v285 = v222
		goto L83
	} else {
		goto L101
	}
L89:
	;
	v275 = int32(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v243 <= int32(0) {
		v269 = v237
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v275 = v269
	goto L88
L93:
	;
	v246 = int32(0)
	if v246 < v243 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v249 = v243
	goto L96
L95:
	;
	v249 = v246
	goto L96
L96:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v252 = int32(0)
	goto L97
L97:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v250+v252<<(uint(int32(2))%32))))
	v261 = base.B2i32(v260 == v236)
	if v260 == v236 {
		v269 = v261
		goto L92
	} else {
		goto L99
	}
L98:
	;
	v269 = v261
	goto L92
L99:
	;
	v263 = v252 + int32(1)
	if v263 != v249 {
		v252 = v263
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v279 = v222 + int32(1)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v279 < v280 {
		v222 = v279
		goto L86
	} else {
		goto L102
	}
L102:
	;
	goto L87
L103:
	;
	v296 = int32(0)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if base.B2i32(v297 == v296)|base.B2i32(l2 == v296) != 0 {
		v374 = v296
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if v374 == v16 {
		v797 = l2
		goto L29
	} else {
		goto L124
	}
L105:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v303 <= int32(0) {
		v374 = v296
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v311 = v296
	goto L107
L107:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319+v311<<(uint(int32(2))%32))))
	v324 = int32(0)
	if v297 == v324 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v374 = v366
	goto L104
L109:
	;
	if v362 == int32(0) {
		v374 = v311
		goto L104
	} else {
		goto L122
	}
L110:
	;
	v362 = int32(0)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if v330 <= int32(0) {
		v356 = v324
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v362 = v356
	goto L109
L114:
	;
	v333 = int32(0)
	if v333 < v330 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v336 = v330
	goto L117
L116:
	;
	v336 = v333
	goto L117
L117:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	v339 = int32(0)
	goto L118
L118:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v337+v339<<(uint(int32(2))%32))))
	v348 = base.B2i32(v347 == v323)
	if v347 == v323 {
		v356 = v348
		goto L113
	} else {
		goto L120
	}
L119:
	;
	v356 = v348
	goto L113
L120:
	;
	v350 = v339 + int32(1)
	if v350 != v336 {
		v339 = v350
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v366 = v311 + int32(1)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v366 < v367 {
		v311 = v366
		goto L107
	} else {
		goto L123
	}
L123:
	;
	goto L108
L124:
	;
	if v134 < v71 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v384 = v71
	goto L127
L126:
	;
	v384 = v134
	goto L127
L127:
	;
	if v202 < v384 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v386 = v384
	goto L130
L129:
	;
	v386 = v202
	goto L130
L130:
	;
	if v285 < v386 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v388 = v386
	goto L133
L132:
	;
	v388 = v285
	goto L133
L133:
	;
	if v374 < v388 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v390 = v388
	goto L136
L135:
	;
	v390 = v374
	goto L136
L136:
	;
	if l2 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v780 < v390 {
		goto L227
	} else {
		goto L228
	}
L138:
	;
	v780 = int32(0)
	goto L137
L139:
	;
	goto L140
L140:
	;
	v394 = int32(0)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v395 <= v394 {
		v780 = v394
		goto L137
	} else {
		goto L141
	}
L141:
	;
	v403 = v394
	goto L142
L142:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411+v403<<(uint(int32(2))%32))))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v416 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v780 = v772
	goto L137
L144:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	if v482 == int32(1) {
		goto L159
	} else {
		goto L160
	}
L145:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v415)+12))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v441)+12))
	if v466 != v467 {
		v780 = v403
		goto L137
	} else {
		goto L157
	}
L146:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v415)+12))
	if v463 == int32(1) {
		goto L144
	} else {
		goto L156
	}
L147:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v419 <= int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	v428 = int32(0)
	goto L149
L149:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v423+v428<<(uint(int32(2))%32))))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	if v442 == v422 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L146
L151:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	if v444 == v445 {
		goto L145
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v448 = v428 + int32(1)
	if v419 != v448 {
		v428 = v448
		goto L149
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	goto L150
L156:
	;
	v780 = v403
	goto L137
L157:
	;
	goto L144
L158:
	;
	v772 = v403 + int32(1)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v772 < v773 {
		v403 = v772
		goto L142
	} else {
		goto L226
	}
L159:
	;
	v485 = int32(0)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+40)))
	if v487 != 0 {
		v643 = v485
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v665 == int32(0) {
		v780 = v403
		goto L137
	} else {
		goto L206
	}
L162:
	;
	if v643 != 0 {
		goto L158
	} else {
		goto L205
	}
L163:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v486)+16))
	if v488 == int32(0) {
		v643 = v485
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v491 < int32(2) {
		v643 = v485
		goto L162
	} else {
		goto L165
	}
L165:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v486)+36))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v497 = v495 - int32(2)
	if base.Ui32(v497) <= base.Ui32(int32(3)) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v497<<(uint(int32(2))%32))+uint32(_c_F_truncate_useless_pathkeys[0])))
	v504 = v502
	goto L168
L167:
	;
	v504 = int32(8)
	goto L168
L168:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v504+l1)))
	v507 = int32(0)
	if v494 == v507 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v560 != 0 {
		v643 = v485
		goto L162
	} else {
		goto L183
	}
L170:
	;
	v560 = int32(1)
	goto L169
L171:
	;
	goto L172
L172:
	;
	if v506 == int32(0) {
		v553 = v507
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v560 = v553
	goto L169
L174:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	if v517 < v516 {
		v553 = v507
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v519 = int32(1)
	if v516 <= v519 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v522 = v519
	goto L178
L177:
	;
	v522 = v516
	goto L178
L178:
	;
	v523 = int32(8)
	v528 = int32(0)
	goto L179
L179:
	;
	v535 = v528 << (uint(int32(2)) % 32)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v494+v523+v535)))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v506+v523+v535)))
	v542 = v537 & (v539 ^ int32(-1))
	v544 = base.B2i32(v542 == int32(0))
	if v542 != 0 {
		v553 = v544
		goto L173
	} else {
		goto L181
	}
L180:
	;
	v553 = v544
	goto L173
L181:
	;
	v546 = v528 + int32(1)
	if v546 != v522 {
		v528 = v546
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v486)+16))
	if v561 == int32(0) {
		v643 = v485
		goto L162
	} else {
		goto L184
	}
L184:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v564 <= int32(0) {
		v643 = v485
		goto L162
	} else {
		goto L185
	}
L185:
	;
	v571 = v485
	goto L186
L186:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v580+v571<<(uint(int32(2))%32))))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)+8))
	v586 = int32(0)
	if base.B2i32(v585 == v586)|base.B2i32(v506 == v586) != 0 {
		v631 = v586
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v643 = v631 ^ int32(1)
	goto L162
L188:
	;
	if v631 != 0 {
		goto L201
	} else {
		goto L202
	}
L189:
	;
	goto L188
L190:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	if v596 < v597 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v599 = v596
	goto L193
L192:
	;
	v599 = v597
	goto L193
L193:
	;
	if v599 <= int32(1) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v602 = int32(1)
	goto L196
L195:
	;
	v602 = v599
	goto L196
L196:
	;
	v603 = int32(8)
	v608 = int32(0)
	goto L197
L197:
	;
	v615 = v608 << (uint(int32(2)) % 32)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v506+v603+v615)))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v585+v603+v615)))
	v620 = v617 & v619
	v622 = base.B2i32(v620 != int32(0))
	if v620 != 0 {
		v631 = v622
		goto L189
	} else {
		goto L199
	}
L198:
	;
	v631 = v622
	goto L189
L199:
	;
	v624 = v608 + int32(1)
	if v624 != v602 {
		v608 = v624
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	v633 = v571 + int32(1)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v633 < v634 {
		v571 = v633
		goto L186
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	goto L187
L204:
	;
	goto L203
L205:
	;
	goto L161
L206:
	;
	v668 = int32(0)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v665)+4))
	if v669 <= v668 {
		v780 = v403
		goto L137
	} else {
		goto L207
	}
L207:
	;
	v675 = v669
	v676 = v668
	goto L208
L208:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v665)+12))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v685+v676<<(uint(int32(2))%32))))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)+96))
	if v690 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v780 = v403
	goto L137
L210:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v689)+100))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v691)+56))
	if v692 != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v745 = v675
	goto L212
L212:
	;
	v756 = v676 + int32(1)
	if v756 < v745 {
		v675 = v745
		v676 = v756
		goto L208
	} else {
		goto L225
	}
L213:
	;
	v696 = v692
	goto L216
L214:
	;
	v714 = v691
	goto L215
L215:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v689)+104))
	v725 = v721
	goto L219
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v689)+100)) = v696
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v696)+56))
	if v707 != 0 {
		v696 = v707
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v714 = v696
	goto L215
L218:
	;
	goto L217
L219:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v725)+56))
	if v735 != 0 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	if base.B2i32(v714 == v737)|base.B2i32(v725 == v737) != 0 {
		goto L158
	} else {
		goto L224
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v689)+104)) = v735
	v725 = v735
	goto L219
L222:
	;
	goto L223
L223:
	;
	goto L220
L224:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v665)+4))
	v745 = v741
	goto L212
L225:
	;
	goto L209
L226:
	;
	goto L143
L227:
	;
	v789 = v390
	goto L229
L228:
	;
	v789 = v780
	goto L229
L229:
	;
	if v789 == v16 {
		v797 = l2
		goto L29
	} else {
		goto L230
	}
L230:
	;
	v791 = F_list_copy_head(m, l2, v789)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	return int32(0)
L232:
	;
	v797 = v791
	goto L29
}
func F_tsearch_readline_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v10 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
			F_errcontext_msg(m, int32(_a_F_tsearch_readline_callback_0), v8+int32(16))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
			F_errcontext_msg(m, int32(_a_F_tsearch_readline_callback_1), v8)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	}
}
func F_tsearch_readline_end(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v4 != v3 {
			F_pfree(m, v3)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_pfree(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v14 = F_FreeFile(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, _c_F_tsearch_readline_end[0])) = v17
						return
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pfree(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v14 = F_FreeFile(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, _c_F_tsearch_readline_end[0])) = v17
					return
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_pfree(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = F_FreeFile(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, _c_F_tsearch_readline_end[0])) = v17
				return
			}
		}
	}
}
func F_tsm_system_time_handler(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14392(m, l0, int32(0), int32(_a_F_tsm_system_time_handler_0), int32(_a_F_tsm_system_time_handler_1), int32(_a_F_tsm_system_time_handler_2), int32(_a_F_tsm_system_time_handler_3), int32(_a_F_tsm_system_time_handler_4), int32(701))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_tsqueryrecv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v366 int32
	_ = v366
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = F_pq_getmsgint(m, v24, int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L10
	} else {
		goto L91
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L10
	} else {
		goto L88
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L10
	} else {
		goto L85
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L10
	} else {
		goto L82
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L10
	} else {
		goto L79
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L10
	} else {
		goto L76
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L10
	} else {
		goto L73
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L10
	} else {
		goto L70
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L10
	} else {
		goto L67
	}
L10:
	;
	return int64(0)
L11:
	;
	if base.Ui32(v26) < base.Ui32(int32(89478486)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v34 = F_palloc(m, v26<<(uint(int32(2))%32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L10
	} else {
		goto L64
	}
L15:
	;
	v39 = v26*int32(12) + int32(8)
	v40 = F_palloc0(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v26
	if v26 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v49 = v40 + int32(8)
	v56 = v2
	v58 = v2
	goto L20
L18:
	;
	v271 = v2
	goto L19
L19:
	;
	v281 = v271 + v39
	v282 = F_repalloc(m, v40, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L10
	} else {
		goto L48
	}
L20:
	;
	v67 = F_pq_getmsgint(m, v24, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L22
	}
L21:
	;
	v271 = v247
	goto L19
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v67)
	switch v67&int32(255) - int32(1) {
	case 0:
		goto L25
	case 1:
		goto L24
	default:
		goto L2
	}
L23:
	;
	v260 = v58 + int32(1)
	if v260 != v26 {
		v49 = v49 + int32(12)
		v56 = v247
		v58 = v260
		goto L20
	} else {
		goto L47
	}
L24:
	;
	v214 = F_pq_getmsgint(m, v24, int32(1))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L41
	}
L25:
	;
	v75 = F_pq_getmsgint(m, v24, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v78 = F_pq_getmsgint(m, v24, int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v80 = F_pq_getmsgstring(m, v24)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v82 = F_strlen(m, v80)
	mBase = m.M
	if v75&int32(240) != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v82 == int32(0) {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(int32(2048)) <= base.Ui32(v82) {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if int32(_a_F_tsqueryrecv_0) < v56 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v91 = int32(-1)
	if v82 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)) = uint8(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v175 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v82 | v56<<(uint(int32(12))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)) = uint8(base.B2i32(v78&int32(255) != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v34+v58<<(uint(int32(2))%32)))) = v80
	v247 = v82 + v56 + int32(1)
	goto L23
L34:
	;
	v99 = v80
	v100 = v91
	v106 = int32(0)
	goto L37
L35:
	;
	v145 = v80
	v146 = v91
	goto L36
L36:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v170 = *(*int32)(unsafe.Add(mBase, uint32((v164^int32(base.Ui32(v146)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_c_F_tsqueryrecv[0])))
	v175 = v170 ^ v146<<(uint(int32(8))%32)
	goto L33
L37:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v120 = int32(24)
	v123 = int32(2)
	v125 = *(*int32)(unsafe.Add(mBase, uint32((v119^int32(base.Ui32(v100)>>(uint(v120)%32)))<<(uint(v123)%32))+uint32(_c_F_tsqueryrecv[0])))
	v126 = int32(8)
	v128 = v125 ^ v100<<(uint(v126)%32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32((v118^int32(base.Ui32(v128)>>(uint(v120)%32)))<<(uint(v123)%32))+uint32(_c_F_tsqueryrecv[0])))
	v137 = v134 ^ v128<<(uint(v126)%32)
	v139 = v99 + v123
	v141 = v106 + v123
	if v141 != v82&int32(2046) {
		v99 = v139
		v100 = v137
		v106 = v141
		goto L37
	} else {
		goto L39
	}
L38:
	;
	if v82&int32(1) == int32(0) {
		v175 = v137
		goto L33
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v145 = v139
	v146 = v137
	goto L36
L41:
	;
	v217 = v214 << (uint(int32(24)) % 32)
	v219 = base.B2i32(v217 != int32(67108864))
	if v219&base.B2i32(base.B2i32(v214&int32(253) == int32(1))|base.B2i32(v217 == int32(33554432)) == int32(0)) != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if v58 == v26-int32(1) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)) = uint8(v214)
	if v217 != int32(67108864) {
		v247 = v56
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v233 = F_pq_getmsgint(m, v24, int32(2))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	if base.Ui32(int32(_a_F_tsqueryrecv_1)) <= base.Ui32(v233) {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+2)) = uint16(v233)
	v247 = v56
	goto L23
L47:
	;
	goto L21
L48:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(0)
	v288 = v282 + int32(8)
	F_findoprnd_recurse(m, v288, v22+int32(44), v26, v22+int32(43))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v295 != v26 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v26 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v301 = v288
	v302 = int32(0)
	v308 = v288 + v284*int32(12)
	goto L54
L52:
	;
	goto L53
L53:
	;
	F_pfree(m, v34)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L10
	} else {
		goto L63
	}
L54:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v320 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	v327 = v323&int32(4095) + int32(1)
	if v327 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v340 = v308
	goto L58
L58:
	;
	v344 = v302 + int32(1)
	if v344 != v26 {
		v301 = v301 + int32(12)
		v302 = v344
		v308 = v340
		goto L54
	} else {
		goto L62
	}
L59:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v34+v302<<(uint(int32(2))%32))))
	base.MemoryCopy(m, v308, v331, v327)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	v340 = v308 + v333&int32(4095) + int32(1)
	goto L58
L62:
	;
	goto L55
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v281 << (uint(int32(2)) % 32)
	m.G0 = v22 + int32(48)
	return base.I64_extend_i32_u(v282)
L64:
	;
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_2), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1241), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_5), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1274), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_6), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1277), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L10
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
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_7), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1280), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_8), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1283), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v217 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_9), v22+int32(32))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1312), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_10), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1314), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v233
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_11), v22+int32(16))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1323), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v489 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v489
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_12), v22)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(1328), int32(_a_F_tsqueryrecv_4))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errmsg_internal(m, int32(_a_F_tsqueryrecv_13), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_tsqueryrecv_3), int32(793), int32(_a_F_tsqueryrecv_14))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsquerysend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = v12 + int32(16)
	F_pq_begintypsend(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	F_enlargeStringInfo(m, v16, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v30 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v26))) = base.I32_rotr(v21, int32(24))&v30 | base.I32_rotr(v21&v30, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v25 + int32(4)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if int32(0) < v41 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v46 = v14 + int32(8)
	v49 = v46
	v53 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v170 != v14 {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	F_enlargeStringInfo(m, v12+int32(16), int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v63))) = uint8(v56)
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v62 + v66
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	switch v69 - v66 {
	case 0:
		goto L11
	case 1:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v158 = v53 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v158 < v159 {
		v49 = v49 + int32(12)
		v53 = v158
		goto L7
	} else {
		goto L23
	}
L11:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v120 = v12 + int32(16)
	F_enlargeStringInfo(m, v120, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L20
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v74 = v12 + int32(16)
	F_enlargeStringInfo(m, v74, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v78+v79))) = uint8(v72)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v78 + int32(1)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v85 != int32(4) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+2)))
	F_enlargeStringInfo(m, v74, int32(2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v95 = int32(8)
	v99 = v88<<(uint(v95)%32) | int32(base.Ui32(v88)>>(uint(v95)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v92+v93))) = uint16(v99)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v92 + int32(2)
	goto L10
L17:
	;
	v108 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v108
	F_errmsg_internal(m, int32(_a_F_tsquerysend_0), v12)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_tsquerysend_1), int32(1215), int32(_a_F_tsquerysend_2))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v124+v125))) = uint8(v118)
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v124 + v128
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
	F_enlargeStringInfo(m, v120, v128)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v136))) = uint8(v131)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v135 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v143 = int32(12)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	F_pq_sendstring(m, v120, v46+v142*v143+int32(base.Ui32(v146)>>(uint(v143)%32)))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L10
L23:
	;
	goto L8
L24:
	;
	F_pfree(m, v14)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v175 = v12 + int32(16)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v178 << (uint(int32(2)) % 32)
	goto L28
L27:
	;
	goto L26
L28:
	;
	m.G0 = v12 + int32(32)
	return base.I64_extend_i32_u(v177)
}
func F_tsquerytree(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 == v2 {
		v18 = F_palloc(m, int32(4))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(16)
			v105 = v18
			m.G0 = v11 + int32(32)
			return base.I64_extend_i32_u(v105)
		}
	} else {
		v24 = int32(8)
		v26 = m.G0
		v28 = v26 - int32(16)
		m.G0 = v28
		v31 = v13 + v24
		v32 = F_maketree(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int64(0)
		} else {
			v34 = F_clean_NOT_intree(m, v32)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = int64(16)
				v38 = int32(0)
				if v34 == v38 {
					v60 = v2
					v61 = v38
					*(*int32)(unsafe.Add(mBase, uint32(v11+v24))) = v60
					m.G0 = v28 + int32(16)
					if v61 == int32(0) {
						v69 = F_cstring_to_text(m, int32(_a_F_tsquerytree_0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							v100 = v69
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v13 == v101 {
								v105 = v100
								m.G0 = v11 + int32(32)
								return base.I64_extend_i32_u(v105)
							} else {
								F_pfree(m, v13)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int64(0)
								} else {
									v105 = v100
									m.G0 = v11 + int32(32)
									return base.I64_extend_i32_u(v105)
								}
							}
						}
					} else {
						v71 = int32(32)
						*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v61
						v76 = F_palloc_mul(m, int32(1), v71)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v76
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v76
							v80 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v80)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v83 = int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31 + v82*v83
							F_infix_1(m, v11+v83, int32(-1), v80)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
								v96 = F_cstring_to_text_with_len(m, v93, v94-v93)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									F_pfree(m, v61)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int64(0)
									} else {
										v100 = v96
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v13 == v101 {
											v105 = v100
											m.G0 = v11 + int32(32)
											return base.I64_extend_i32_u(v105)
										} else {
											F_pfree(m, v13)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int64(0)
											} else {
												v105 = v100
												m.G0 = v11 + int32(32)
												return base.I64_extend_i32_u(v105)
											}
										}
									}
								}
							}
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
					v44 = int32(1)
					if base.Ui32(v44) < base.Ui32((v43-v44)&int32(255)) {
						v60 = v2
						v61 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v11+v24))) = v60
						m.G0 = v28 + int32(16)
						if v61 == int32(0) {
							v69 = F_cstring_to_text(m, int32(_a_F_tsquerytree_0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								v100 = v69
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v13 == v101 {
									v105 = v100
									m.G0 = v11 + int32(32)
									return base.I64_extend_i32_u(v105)
								} else {
									F_pfree(m, v13)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int64(0)
									} else {
										v105 = v100
										m.G0 = v11 + int32(32)
										return base.I64_extend_i32_u(v105)
									}
								}
							}
						} else {
							v71 = int32(32)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v61
							v76 = F_palloc_mul(m, int32(1), v71)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v76
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v76
								v80 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v80)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v83 = int32(12)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31 + v82*v83
								F_infix_1(m, v11+v83, int32(-1), v80)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
									v96 = F_cstring_to_text_with_len(m, v93, v94-v93)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										F_pfree(m, v61)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int64(0)
										} else {
											v100 = v96
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											if v13 == v101 {
												v105 = v100
												m.G0 = v11 + int32(32)
												return base.I64_extend_i32_u(v105)
											} else {
												F_pfree(m, v13)
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return int64(0)
												} else {
													v105 = v100
													m.G0 = v11 + int32(32)
													return base.I64_extend_i32_u(v105)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v51 = F_palloc(m, int32(192))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v51
							F_plainnode(m, v28+int32(4), v34)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
								v60 = v58
								v61 = v59
								*(*int32)(unsafe.Add(mBase, uint32(v11+v24))) = v60
								m.G0 = v28 + int32(16)
								if v61 == int32(0) {
									v69 = F_cstring_to_text(m, int32(_a_F_tsquerytree_0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int64(0)
									} else {
										v100 = v69
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v13 == v101 {
											v105 = v100
											m.G0 = v11 + int32(32)
											return base.I64_extend_i32_u(v105)
										} else {
											F_pfree(m, v13)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int64(0)
											} else {
												v105 = v100
												m.G0 = v11 + int32(32)
												return base.I64_extend_i32_u(v105)
											}
										}
									}
								} else {
									v71 = int32(32)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v71
									*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v61
									v76 = F_palloc_mul(m, int32(1), v71)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v76
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v76
										v80 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v80)
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v83 = int32(12)
										*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31 + v82*v83
										F_infix_1(m, v11+v83, int32(-1), v80)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int64(0)
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
											v96 = F_cstring_to_text_with_len(m, v93, v94-v93)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int64(0)
											} else {
												F_pfree(m, v61)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int64(0)
												} else {
													v100 = v96
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													if v13 == v101 {
														v105 = v100
														m.G0 = v11 + int32(32)
														return base.I64_extend_i32_u(v105)
													} else {
														F_pfree(m, v13)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int64(0)
														} else {
															v105 = v100
															m.G0 = v11 + int32(32)
															return base.I64_extend_i32_u(v105)
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
	}
}
func F_tupledesc_match(m *base.Module, l0 int32, l1 int32) {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v12 == v13 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L11
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L25
	}
L3:
	;
	if int32(0) < v12 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L20
	}
L6:
	;
	v20 = v12
	v22 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v10 + int32(48)
	return
L9:
	;
	v25 = v22 * int32(100)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v27 = int32(3)
	v30 = v25 + (l1 + v26<<(uint(v27)%32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	v35 = l0 + v20<<(uint(v27)%32) + v25
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+96))
	v37 = F_IsBinaryCoercible(m, v31, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = int32(28)
	v42 = v30 + v41
	v44 = v35 + v41
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+91)))
	if v45 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v57 = v22 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v57 < v58 {
		v20 = v58
		v22 = v57
		goto L9
	} else {
		goto L19
	}
L16:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+72)))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+72)))
	if v48 != v49 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+83)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+83)))
	if v51 != v52 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	goto L10
L20:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_tupledesc_match_0), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v81
	F_errdetail_plural(m, int32(_a_F_tupledesc_match_1), int32(_a_F_tupledesc_match_2), v81, v10+int32(32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_tupledesc_match_3), int32(957), int32(_a_F_tupledesc_match_4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_tupledesc_match_0), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v108 = F_format_type_be(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v44)+68))
	v111 = F_format_type_be(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v22 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v108
	v121 = F_errdetail(m, int32(_a_F_tupledesc_match_5), v10+int32(16))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_tupledesc_match_3), int32(973), int32(_a_F_tupledesc_match_4))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_tupledesc_match_0), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22 + int32(1)
	v143 = F_errdetail(m, int32(_a_F_tupledesc_match_6), v10)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_tupledesc_match_3), int32(981), int32(_a_F_tupledesc_match_4))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tzparse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v478 int32
	_ = v478
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v742 int32
	_ = v742
	var v766 int64
	_ = v766
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v789 int64
	_ = v789
	var v792 int64
	_ = v792
	var v796 int64
	_ = v796
	var v797 int64
	_ = v797
	var v810 int32
	_ = v810
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v900 int64
	_ = v900
	var v905 int32
	_ = v905
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v930 int32
	_ = v930
	var v947 int32
	_ = v947
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int64
	_ = v976
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1011 int64
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1069 int32
	_ = v1069
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1504 int64
	_ = v1504
	var v1507 int64
	_ = v1507
	var v1511 int64
	_ = v1511
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1529 int64
	_ = v1529
	var v1532 int64
	_ = v1532
	var v1536 int64
	_ = v1536
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1555 int64
	_ = v1555
	var v1558 int64
	_ = v1558
	var v1562 int64
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1601 int64
	_ = v1601
	var v1603 int64
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1619 int32
	_ = v1619
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1667 int32
	_ = v1667
	v4 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(48)
	m.G0 = v31
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v31 + int32(48)
	return v1667
L2:
	;
	v370 = v348 + int32(1)
	if base.Ui32(int32(512)) < base.Ui32(v370) {
		v1667 = v345
		goto L1
	} else {
		goto L61
	}
L3:
	;
	v33 = F_strlen(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = int32(0)
	v341 = l0
	v345 = v4
	v348 = v33
	v349 = l0 + v33
	goto L2
L4:
	;
	goto L5
L5:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v183 = int32(0)
	if v157&int32(255) == v183 {
		v1667 = v183
		goto L1
	} else {
		goto L30
	}
L7:
	;
	v115 = l0 + int32(1)
	v118 = v115
	goto L22
L8:
	;
	v155 = l0
	v157 = v87
	v158 = v88
	v162 = v88 - l0
	goto L6
L9:
	;
	v87 = v37
	v88 = l0
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v37 == int32(60) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v44 = v37
	v46 = l0
	goto L13
L13:
	;
	if base.Ui32(int32(252)) < base.Ui32((v44-int32(46))&int32(255)) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v87 = int32(0)
	v88 = v83
	goto L8
L15:
	;
	v87 = v44
	v88 = v46
	goto L8
L16:
	;
	goto L17
L17:
	;
	if base.Ui32(int32(-11)) < base.Ui32(base.I32_extend8_s(v44)-int32(58)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v87 = v44
	v88 = v46
	goto L8
L19:
	;
	goto L20
L20:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	v83 = v46 + int32(1)
	if v81 != 0 {
		v44 = v81
		v46 = v83
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v144 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	v155 = v115
	v157 = v154
	v158 = v118 + int32(1)
	v162 = v118 - v115
	goto L6
L24:
	;
	v1667 = v4
	goto L1
L25:
	;
	goto L26
L26:
	;
	if v144 != int32(62) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v118 = v118 + int32(1)
	goto L22
L28:
	;
	goto L29
L29:
	;
	goto L23
L30:
	;
	v189 = v31 + int32(44)
	v190 = int32(0)
	v196 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	switch v197 - int32(43) {
	case 0:
		goto L33
	default:
		v205 = v158
		v206 = v196
		goto L32
	case 2:
		goto L34
	}
L31:
	;
	if v338 == int32(0) {
		v1667 = v183
		goto L1
	} else {
		goto L60
	}
L32:
	;
	v207 = int32(*(*int8)(unsafe.Add(mBase, uint32(v205))))
	if base.Ui32(int32(9)) < base.Ui32(v207-int32(48)) {
		v338 = v190
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v205 = v158 + int32(1)
	v206 = v196
	goto L32
L34:
	;
	v205 = v158 + int32(1)
	v206 = int32(0)
	goto L32
L35:
	;
	goto L31
L36:
	;
	v212 = v205
	v214 = v190
	v215 = v207
	goto L37
L37:
	;
	v225 = base.I32_extend8_s(v215) + v214*int32(10) - int32(48)
	if int32(167) < v225 {
		v338 = v190
		goto L35
	} else {
		goto L39
	}
L38:
	;
	if v225 < int32(0) {
		v338 = v190
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v229 = v212 + int32(1)
	v230 = int32(*(*int8)(unsafe.Add(mBase, uint32(v212)+1)))
	if base.Ui32(v230-int32(48)) < base.Ui32(int32(10)) {
		v212 = v229
		v214 = v225
		v215 = v230
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v238 = v225 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v238
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v240 != int32(58) {
		v318 = v229
		v323 = v238
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v206 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	v243 = int32(*(*int8)(unsafe.Add(mBase, uint32(v212)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v243-int32(48)) {
		v338 = v190
		goto L35
	} else {
		goto L44
	}
L44:
	;
	v251 = v212 + int32(2)
	v253 = int32(0)
	v254 = v243
	goto L45
L45:
	;
	v264 = base.I32_extend8_s(v254) + v253*int32(10) - int32(48)
	if int32(59) < v264 {
		v338 = v190
		goto L35
	} else {
		goto L47
	}
L46:
	;
	if v264 < int32(0) {
		v338 = v190
		goto L35
	} else {
		goto L49
	}
L47:
	;
	v268 = v251 + int32(1)
	v269 = int32(*(*int8)(unsafe.Add(mBase, uint32(v251)+1)))
	if base.Ui32(v269-int32(48)) < base.Ui32(int32(10)) {
		v251 = v268
		v253 = v264
		v254 = v269
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v278 = v264*int32(60) + v238
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v278
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v280 != int32(58) {
		v318 = v268
		v323 = v278
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v283 = int32(*(*int8)(unsafe.Add(mBase, uint32(v251)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v283-int32(48)) {
		v338 = v190
		goto L35
	} else {
		goto L51
	}
L51:
	;
	v291 = v251 + int32(2)
	v293 = v283
	v294 = int32(0)
	goto L52
L52:
	;
	v304 = base.I32_extend8_s(v293) + v294*int32(10) - int32(48)
	if int32(60) < v304 {
		v338 = v190
		goto L35
	} else {
		goto L54
	}
L53:
	;
	if v304 < int32(0) {
		v338 = v190
		goto L35
	} else {
		goto L56
	}
L54:
	;
	v307 = int32(*(*int8)(unsafe.Add(mBase, uint32(v291)+1)))
	v309 = v291 + int32(1)
	if base.Ui32(v307-int32(48)) < base.Ui32(int32(10)) {
		v291 = v309
		v293 = v307
		v294 = v304
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v316 = v304 + v278
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v316
	v318 = v309
	v323 = v316
	goto L42
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = int32(0) - v323
	goto L59
L58:
	;
	goto L59
L59:
	;
	v338 = v318
	goto L35
L60:
	;
	v341 = v155
	v345 = v183
	v348 = v162
	v349 = v338
	goto L2
L61:
	;
	v373 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v373)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v373
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	if v377 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v517 = int32(0)
	if v516 == v517 {
		v1667 = v517
		goto L1
	} else {
		goto L89
	}
L63:
	;
	v449 = v349 + int32(1)
	v452 = v449
	goto L81
L64:
	;
	if v377 == int32(60) {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(4294967296)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v429 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[0]))) = v429
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[1]))) = uint16(v429)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[2]))) = v429
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[3]))) = uint8(v429)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[4]))) = v429 - v428
	v442 = l1 + int32(_a_F_tzparse_0)
	if v348 != 0 {
		goto L78
	} else {
		goto L79
	}
L67:
	;
	v382 = v377
	v384 = v349
	goto L68
L68:
	;
	if base.Ui32(int32(252)) < base.Ui32((v382-int32(46))&int32(255)) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v491 = v423
	v496 = v349
	v516 = v423 - v349
	goto L62
L70:
	;
	goto L69
L71:
	;
	v423 = v384
	goto L70
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(-11)) < base.Ui32(base.I32_extend8_s(v382)-int32(58)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v423 = v384
	goto L70
L75:
	;
	goto L76
L76:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+1)))
	v421 = v384 + int32(1)
	if v419 != 0 {
		v382 = v419
		v384 = v421
		goto L68
	} else {
		goto L77
	}
L77:
	;
	v423 = v421
	goto L70
L78:
	;
	base.MemoryCopy(m, v442, v341, v348)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v445 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v442+v348))) = uint8(v445)
	v1667 = int32(1)
	goto L1
L81:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	if v478 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v491 = v452 + int32(1)
	v496 = v449
	v516 = v452 - v449
	goto L62
L83:
	;
	v1667 = v345
	goto L1
L84:
	;
	goto L85
L85:
	;
	if v478 != int32(62) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v452 = v452 + int32(1)
	goto L81
L87:
	;
	goto L88
L88:
	;
	goto L82
L89:
	;
	v522 = v348 + v516 + int32(2)
	if base.Ui32(int32(512)) < base.Ui32(v522) {
		v1667 = v517
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491))))
	switch v525 - int32(44) {
	case 0, 15:
		goto L92
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L93
	default:
		goto L94
	}
L91:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685))))
	if v687 != 0 {
		goto L135
	} else {
		goto L136
	}
L92:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v681 - int32(3600)
	v685 = v491
	goto L91
L93:
	;
	v531 = v31 + int32(40)
	v532 = int32(0)
	v538 = int32(1)
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491))))
	switch v539 - int32(43) {
	case 0:
		goto L98
	default:
		v547 = v491
		v548 = v538
		goto L97
	case 2:
		goto L99
	}
L94:
	;
	if v525 == int32(0) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	if v680 != 0 {
		v685 = v680
		goto L91
	} else {
		goto L125
	}
L97:
	;
	v549 = int32(*(*int8)(unsafe.Add(mBase, uint32(v547))))
	if base.Ui32(int32(9)) < base.Ui32(v549-int32(48)) {
		v680 = v532
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v547 = v491 + int32(1)
	v548 = v538
	goto L97
L99:
	;
	v547 = v491 + int32(1)
	v548 = int32(0)
	goto L97
L100:
	;
	goto L96
L101:
	;
	v554 = v547
	v556 = v532
	v557 = v549
	goto L102
L102:
	;
	v567 = base.I32_extend8_s(v557) + v556*int32(10) - int32(48)
	if int32(167) < v567 {
		v680 = v532
		goto L100
	} else {
		goto L104
	}
L103:
	;
	if v567 < int32(0) {
		v680 = v532
		goto L100
	} else {
		goto L106
	}
L104:
	;
	v571 = v554 + int32(1)
	v572 = int32(*(*int8)(unsafe.Add(mBase, uint32(v554)+1)))
	if base.Ui32(v572-int32(48)) < base.Ui32(int32(10)) {
		v554 = v571
		v556 = v567
		v557 = v572
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v580 = v567 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = v580
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	if v582 != int32(58) {
		v660 = v571
		v665 = v580
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if v548 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L108:
	;
	v585 = int32(*(*int8)(unsafe.Add(mBase, uint32(v554)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v585-int32(48)) {
		v680 = v532
		goto L100
	} else {
		goto L109
	}
L109:
	;
	v593 = v554 + int32(2)
	v595 = int32(0)
	v596 = v585
	goto L110
L110:
	;
	v606 = base.I32_extend8_s(v596) + v595*int32(10) - int32(48)
	if int32(59) < v606 {
		v680 = v532
		goto L100
	} else {
		goto L112
	}
L111:
	;
	if v606 < int32(0) {
		v680 = v532
		goto L100
	} else {
		goto L114
	}
L112:
	;
	v610 = v593 + int32(1)
	v611 = int32(*(*int8)(unsafe.Add(mBase, uint32(v593)+1)))
	if base.Ui32(v611-int32(48)) < base.Ui32(int32(10)) {
		v593 = v610
		v595 = v606
		v596 = v611
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v620 = v606*int32(60) + v580
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = v620
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	if v622 != int32(58) {
		v660 = v610
		v665 = v620
		goto L107
	} else {
		goto L115
	}
L115:
	;
	v625 = int32(*(*int8)(unsafe.Add(mBase, uint32(v593)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v625-int32(48)) {
		v680 = v532
		goto L100
	} else {
		goto L116
	}
L116:
	;
	v633 = v593 + int32(2)
	v635 = v625
	v636 = int32(0)
	goto L117
L117:
	;
	v646 = base.I32_extend8_s(v635) + v636*int32(10) - int32(48)
	if int32(60) < v646 {
		v680 = v532
		goto L100
	} else {
		goto L119
	}
L118:
	;
	if v646 < int32(0) {
		v680 = v532
		goto L100
	} else {
		goto L121
	}
L119:
	;
	v649 = int32(*(*int8)(unsafe.Add(mBase, uint32(v633)+1)))
	v651 = v633 + int32(1)
	if base.Ui32(v649-int32(48)) < base.Ui32(int32(10)) {
		v633 = v651
		v635 = v649
		v636 = v646
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v658 = v646 + v620
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = v658
	v660 = v651
	v665 = v658
	goto L107
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = int32(0) - v665
	goto L124
L123:
	;
	goto L124
L124:
	;
	v680 = v660
	goto L100
L125:
	;
	v1667 = v517
	goto L1
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v522
	v1651 = l1 + int32(_a_F_tzparse_0)
	if v348 != 0 {
		goto L288
	} else {
		goto L289
	}
L127:
	;
	if v1607-v974 < int32(401) {
		goto L126
	} else {
		goto L287
	}
L128:
	;
	v1601 = *(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[5])))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[2]))) = v1601
	v1603 = *(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[6])))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[4]))) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	goto L126
L129:
	;
	v979 = l1 + int32(_a_F_tzparse_1)
	v981 = l1 + int32(24)
	v988 = v974
	v989 = v975
	v991 = v4
	v996 = v742 + int32(400)
	v1011 = v976
	goto L173
L130:
	;
	if v742 < int32(2147483248) {
		v974 = v770
		v975 = v785
		v976 = v766
		goto L129
	} else {
		goto L172
	}
L131:
	;
	v947 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[1]))) = uint16(v947)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[2]))) = v947
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[3]))) = uint8(v947)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[4]))) = v947 - v930
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[7]))) = uint16(v947)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[5]))) = v370
	v960 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[8]))) = uint8(v960)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[0]))) = v947
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[6]))) = v947 - v956
	goto L126
L132:
	;
	v810 = v517
	goto L158
L133:
	;
	v700 = F_getrule(m, v688+int32(1), v31+int32(20))
	mBase = m.M
	if v700 == int32(0) {
		v1667 = v517
		goto L1
	} else {
		goto L140
	}
L134:
	;
	if v689 != 0 {
		v1667 = v517
		goto L1
	} else {
		goto L138
	}
L135:
	;
	v688 = v685
	goto L137
L136:
	;
	v688 = int32(_a_F_tzparse_2)
	goto L137
L137:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688))))
	switch v689 - int32(44) {
	case 0, 15:
		goto L133
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		v1667 = v517
		goto L1
	default:
		goto L134
	}
L138:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v692 {
		goto L132
	} else {
		goto L139
	}
L139:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v930 = v695
	goto L131
L140:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v703 != int32(44) {
		v1667 = v517
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v708 = F_getrule(m, v700+int32(1), v31)
	mBase = m.M
	if v708 == int32(0) {
		v1667 = v517
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	if v711 != 0 {
		v1667 = v517
		goto L1
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(2)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v715 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[1]))) = uint16(v715)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[2]))) = v715
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[3]))) = uint8(v715)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[4]))) = v715 - v714
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[0]))) = v715
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[7]))) = uint16(v715)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[5]))) = v370
	v730 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[8]))) = uint8(v730)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_tzparse[6]))) = v715 - v724
	v742 = int32(1970)
	v766 = int64(0)
	goto L144
L144:
	;
	v770 = v742 - int32(1)
	if v770&int32(3) != 0 {
		v780 = int32(0)
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v974 = int32(1770)
	v975 = int32(0)
	v976 = v797
	goto L129
L146:
	;
	v797 = v766 + v796
	if base.Ui32(int32(1771)) < base.Ui32(v742) {
		v742 = v770
		v766 = v797
		goto L144
	} else {
		goto L155
	}
L147:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v780<<(uint(int32(2))%32))+uint32(_c_F_tzparse[9])))
	v785 = v783 * int32(-86400)
	if v785 < int32(0) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v775 = base.I32_rem_u_s(v770, int32(100))
	if v775 != 0 {
		v780 = int32(1)
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v777 = base.I32_rem_u_s(v770, int32(400))
	v780 = base.B2i32(v777 == int32(0))
	goto L147
L150:
	;
	v789 = base.I64_extend_i32_s(v785)
	if int64(-9223372036854775807-1)-v789 <= v766 {
		v796 = v789
		goto L146
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v792 = base.I64_extend_i32_u(v785)
	if v792^int64(9223372036854775807) < v766 {
		goto L130
	} else {
		goto L154
	}
L153:
	;
	goto L130
L154:
	;
	v796 = v792
	goto L146
L155:
	;
	goto L145
L156:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v863 = int32(0)
	v865 = v850
	goto L162
L157:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	v850 = int32(0) - v847
	goto L156
L158:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(_a_F_tzparse_1)+v810))))
	v838 = l1 + int32(_a_F_tzparse_3) + v835<<(uint(int32(4))%32)
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838)+4)))
	if v839 == int32(0) {
		goto L157
	} else {
		goto L160
	}
L159:
	;
	v850 = int32(0)
	goto L156
L160:
	;
	v843 = v810 + int32(1)
	if v843 != v692 {
		v810 = v843
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	v887 = l1 + int32(_a_F_tzparse_1) + v863
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	v891 = l1 + int32(_a_F_tzparse_3) + v888<<(uint(int32(4))%32)
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v887))) = uint8(v892)
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891)+13)))
	if v894 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v930 = v857
	goto L131
L164:
	;
	v917 = v863 + int32(1)
	if v917 < v692 {
		v863 = v917
		v865 = v915
		goto L162
	} else {
		goto L171
	}
L165:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v891)))
	v915 = int32(0) - v912
	goto L164
L166:
	;
	v899 = l1 + int32(24) + v863<<(uint(int32(3))%32)
	v900 = *(*int64)(unsafe.Add(mBase, uint32(v899)))
	*(*int64)(unsafe.Add(mBase, uint32(v899))) = v900 + base.I64_extend_i32_s(v857-v865)
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891)+4)))
	if v905 == int32(0) {
		goto L165
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	if v892&int32(1) != 0 {
		v915 = v865
		goto L164
	} else {
		goto L170
	}
L169:
	;
	v915 = v865
	goto L164
L170:
	;
	goto L165
L171:
	;
	goto L163
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	goto L128
L173:
	;
	v1014 = v31 + int32(20)
	v1015 = int32(0)
	if v988&int32(3) != 0 {
		v1032 = v1015
		goto L177
	} else {
		goto L178
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1545
	if v1545 != 0 {
		v1607 = v1568
		goto L127
	} else {
		goto L286
	}
L175:
	;
	v1551 = v989 + v1483
	if v1551 < int32(0) {
		goto L280
	} else {
		goto L281
	}
L176:
	;
	v1240 = int32(0)
	if v988&int32(3) != 0 {
		v1257 = v1240
		goto L215
	} else {
		goto L216
	}
L177:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	switch v1033 {
	case 0:
		goto L183
	case 1:
		goto L182
	case 2:
		goto L181
	default:
		v1231 = v1015
		goto L180
	}
L178:
	;
	v1027 = base.I32_rem_s(v988, int32(100))
	if v1027 != 0 {
		v1032 = int32(1)
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v1029 = base.I32_rem_s(v988, int32(400))
	v1032 = base.B2i32(v1029 == int32(0))
	goto L177
L180:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+16))
	v1239 = v1237 + (v714 + v1231)
	goto L176
L181:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+12))
	v1050 = v988 - base.B2i32(v1047 < int32(3))
	v1052 = base.I32_div_s(v1050, int32(400))
	v1054 = base.I32_rem_s(v1050, int32(100))
	v1057 = base.I32_div_s(v1050, int32(-100))
	v1058 = int32(1)
	v1063 = base.I32_div_s(base.I32_extend8_s(v1054), int32(4))
	v1069 = base.I32_rem_s(v1047+int32(9), int32(12))
	v1076 = base.I32_div_s(base.I32_extend16_s(v1069*int32(26)+int32(24)), int32(10))
	v1081 = int32(7)
	v1082 = base.I32_rem_s(v1052+v1054+v1057<<(uint(v1058)%32)+base.I32_extend8_s(v1063)+base.I32_extend16_s(v1076+v1058), v1081)
	if v1082 < int32(0) {
		goto L190
	} else {
		goto L191
	}
L182:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	v1231 = v1043 * int32(_a_F_tzparse_4)
	goto L180
L183:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+4))
	v1035 = int32(_a_F_tzparse_4)
	v1036 = v1034 * v1035
	v1038 = v1036 - v1035
	if int32(59) < v1034 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1041 = v1036
	goto L186
L185:
	;
	v1041 = v1038
	goto L186
L186:
	;
	if v1032 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1042 = v1041
	goto L189
L188:
	;
	v1042 = v1038
	goto L189
L189:
	;
	v1231 = v1042
	goto L180
L190:
	;
	v1087 = v1082 + v1081
	goto L192
L191:
	;
	v1087 = v1082
	goto L192
L192:
	;
	v1088 = v1046 - v1087
	if v1088 < int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1093 = v1088 + int32(7)
	goto L195
L194:
	;
	v1093 = v1088
	goto L195
L195:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+8))
	if v1094 < int32(2) {
		v1127 = v1093
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1138 = v1127 * int32(_a_F_tzparse_4)
	v1140 = v1047 - int32(1)
	if v1140 <= int32(0) {
		v1231 = v1138
		goto L180
	} else {
		goto L202
	}
L197:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1032*int32(48)+v1047<<(uint(int32(2))%32))+uint32(_c_F_tzparse[10])))
	v1105 = int32(7)
	v1111 = v1093
	v1115 = int32(1)
	goto L198
L198:
	;
	v1122 = v1111 + int32(7)
	if v1104 <= v1122 {
		v1127 = v1111
		goto L196
	} else {
		goto L200
	}
L199:
	;
	v1127 = v1093 + v1094*v1105 - v1105
	goto L196
L200:
	;
	v1125 = v1115 + int32(1)
	if v1125 != v1094 {
		v1111 = v1122
		v1115 = v1125
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	v1143 = int32(3)
	v1144 = v1140 & v1143
	v1148 = v1032*int32(48) + int32(_a_F_tzparse_5)
	if base.Ui32(v1047-int32(2)) < base.Ui32(v1143) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v1205 = v1195
	v1209 = v1199
	v1212 = int32(0)
	goto L211
L204:
	;
	v1195 = int32(0)
	v1199 = v1138
	goto L203
L205:
	;
	goto L206
L206:
	;
	v1157 = int32(0)
	v1159 = v1157
	v1163 = v1138
	v1164 = v1157
	goto L207
L207:
	;
	v1171 = v1148 + v1159<<(uint(int32(2))%32)
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+12))
	v1173 = int32(_a_F_tzparse_4)
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1171)))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+4))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+8))
	v1187 = v1172*v1173 + (v1175*v1173 + v1163 + v1179*v1173 + v1183*v1173)
	v1188 = int32(4)
	v1189 = v1159 + v1188
	v1191 = v1164 + v1188
	if v1191 != v1140&int32(2147483644) {
		v1159 = v1189
		v1163 = v1187
		v1164 = v1191
		goto L207
	} else {
		goto L209
	}
L208:
	;
	if v1144 == int32(0) {
		v1231 = v1187
		goto L180
	} else {
		goto L210
	}
L209:
	;
	goto L208
L210:
	;
	v1195 = v1189
	v1199 = v1187
	goto L203
L211:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1148+v1205<<(uint(int32(2))%32))))
	v1221 = v1218*int32(_a_F_tzparse_4) + v1209
	v1222 = int32(1)
	v1225 = v1212 + v1222
	if v1225 != v1144 {
		v1205 = v1205 + v1222
		v1209 = v1221
		v1212 = v1225
		goto L211
	} else {
		goto L213
	}
L212:
	;
	v1231 = v1221
	goto L180
L213:
	;
	goto L212
L214:
	;
	v1465 = base.B2i32(v1464 < v1239)
	v1466 = int32(0)
	if v988&int32(3) != 0 {
		v1478 = v1466
		goto L252
	} else {
		goto L253
	}
L215:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	switch v1258 {
	case 0:
		goto L221
	case 1:
		goto L220
	case 2:
		goto L219
	default:
		v1456 = v1240
		goto L218
	}
L216:
	;
	v1252 = base.I32_rem_s(v988, int32(100))
	if v1252 != 0 {
		v1257 = int32(1)
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v1254 = base.I32_rem_s(v988, int32(400))
	v1257 = base.B2i32(v1254 == int32(0))
	goto L215
L218:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v1464 = v1462 + (v724 + v1456)
	goto L214
L219:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v1275 = v988 - base.B2i32(v1272 < int32(3))
	v1277 = base.I32_div_s(v1275, int32(400))
	v1279 = base.I32_rem_s(v1275, int32(100))
	v1282 = base.I32_div_s(v1275, int32(-100))
	v1283 = int32(1)
	v1288 = base.I32_div_s(base.I32_extend8_s(v1279), int32(4))
	v1294 = base.I32_rem_s(v1272+int32(9), int32(12))
	v1301 = base.I32_div_s(base.I32_extend16_s(v1294*int32(26)+int32(24)), int32(10))
	v1306 = int32(7)
	v1307 = base.I32_rem_s(v1277+v1279+v1282<<(uint(v1283)%32)+base.I32_extend8_s(v1288)+base.I32_extend16_s(v1301+v1283), v1306)
	if v1307 < int32(0) {
		goto L228
	} else {
		goto L229
	}
L220:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1456 = v1268 * int32(_a_F_tzparse_4)
	goto L218
L221:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1260 = int32(_a_F_tzparse_4)
	v1261 = v1259 * v1260
	v1263 = v1261 - v1260
	if int32(59) < v1259 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1266 = v1261
	goto L224
L223:
	;
	v1266 = v1263
	goto L224
L224:
	;
	if v1257 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1267 = v1266
	goto L227
L226:
	;
	v1267 = v1263
	goto L227
L227:
	;
	v1456 = v1267
	goto L218
L228:
	;
	v1312 = v1307 + v1306
	goto L230
L229:
	;
	v1312 = v1307
	goto L230
L230:
	;
	v1313 = v1271 - v1312
	if v1313 < int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1318 = v1313 + int32(7)
	goto L233
L232:
	;
	v1318 = v1313
	goto L233
L233:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v1319 < int32(2) {
		v1352 = v1318
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1363 = v1352 * int32(_a_F_tzparse_4)
	v1365 = v1272 - int32(1)
	if v1365 <= int32(0) {
		v1456 = v1363
		goto L218
	} else {
		goto L240
	}
L235:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1257*int32(48)+v1272<<(uint(int32(2))%32))+uint32(_c_F_tzparse[10])))
	v1330 = int32(7)
	v1336 = v1318
	v1340 = int32(1)
	goto L236
L236:
	;
	v1347 = v1336 + int32(7)
	if v1329 <= v1347 {
		v1352 = v1336
		goto L234
	} else {
		goto L238
	}
L237:
	;
	v1352 = v1318 + v1319*v1330 - v1330
	goto L234
L238:
	;
	v1350 = v1340 + int32(1)
	if v1350 != v1319 {
		v1336 = v1347
		v1340 = v1350
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v1368 = int32(3)
	v1369 = v1365 & v1368
	v1373 = v1257*int32(48) + int32(_a_F_tzparse_5)
	if base.Ui32(v1272-int32(2)) < base.Ui32(v1368) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1430 = v1420
	v1434 = v1424
	v1437 = int32(0)
	goto L249
L242:
	;
	v1420 = int32(0)
	v1424 = v1363
	goto L241
L243:
	;
	goto L244
L244:
	;
	v1382 = int32(0)
	v1384 = v1382
	v1388 = v1363
	v1389 = v1382
	goto L245
L245:
	;
	v1396 = v1373 + v1384<<(uint(int32(2))%32)
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+12))
	v1398 = int32(_a_F_tzparse_4)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1396)))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+4))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+8))
	v1412 = v1397*v1398 + (v1400*v1398 + v1388 + v1404*v1398 + v1408*v1398)
	v1413 = int32(4)
	v1414 = v1384 + v1413
	v1416 = v1389 + v1413
	if v1416 != v1365&int32(2147483644) {
		v1384 = v1414
		v1388 = v1412
		v1389 = v1416
		goto L245
	} else {
		goto L247
	}
L246:
	;
	if v1369 == int32(0) {
		v1456 = v1412
		goto L218
	} else {
		goto L248
	}
L247:
	;
	goto L246
L248:
	;
	v1420 = v1414
	v1424 = v1412
	goto L241
L249:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1373+v1430<<(uint(int32(2))%32))))
	v1446 = v1443*int32(_a_F_tzparse_4) + v1434
	v1447 = int32(1)
	v1450 = v1437 + v1447
	if v1450 != v1369 {
		v1430 = v1430 + v1447
		v1434 = v1446
		v1437 = v1450
		goto L249
	} else {
		goto L251
	}
L250:
	;
	v1456 = v1446
	goto L218
L251:
	;
	goto L250
L252:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1478<<(uint(int32(2))%32))+uint32(_c_F_tzparse[9])))
	v1483 = v1481 * int32(_a_F_tzparse_4)
	if base.B2i32(v1465 == v1466)&(base.B2i32(v1483+(v714-v724) <= v1464-v1239)|base.B2i32(v1464 <= v1239)) != 0 {
		v1545 = v991
		v1547 = v996
		goto L175
	} else {
		goto L255
	}
L253:
	;
	v1473 = base.I32_rem_u_s(v988, int32(100))
	if v1473 != 0 {
		v1478 = int32(1)
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v1475 = base.I32_rem_u_s(v988, int32(400))
	v1478 = base.B2i32(v1475 == int32(0))
	goto L252
L255:
	;
	if int32(1999) <= v991 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v991
	v1607 = v988
	goto L127
L257:
	;
	goto L258
L258:
	;
	if v1239 < v1464 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1494 = v1464
	goto L261
L260:
	;
	v1494 = v1239
	goto L261
L261:
	;
	v1497 = v981 + v991<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1497))) = v1011
	if v1464 < v1239 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v1523 = v981 + v1519<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1523))) = v1011
	v1525 = v989 + v1494
	if v1525 < int32(0) {
		goto L273
	} else {
		goto L274
	}
L263:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1497))) = v1011 + v1511
	*(*uint8)(unsafe.Add(mBase, uint32(v991+v979))) = uint8(base.B2i32(v1239 <= v1464))
	v1519 = v991 + int32(1)
	goto L262
L264:
	;
	v1499 = v1464
	goto L266
L265:
	;
	v1499 = v1239
	goto L266
L266:
	;
	v1500 = v1499 + v989
	if v1500 < int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1504 = base.I64_extend_i32_s(v1500)
	if int64(-9223372036854775807-1)-v1504 <= v1011 {
		v1511 = v1504
		goto L263
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1507 = base.I64_extend_i32_u(v1500)
	if v1507^int64(9223372036854775807) < v1011 {
		v1519 = v991
		goto L262
	} else {
		goto L271
	}
L270:
	;
	v1519 = v991
	goto L262
L271:
	;
	v1511 = v1507
	goto L263
L272:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1523))) = v1011 + v1536
	*(*uint8)(unsafe.Add(mBase, uint32(v1519+v979))) = uint8(v1465)
	v1545 = v1519 + int32(1)
	v1547 = v988 + int32(401)
	goto L175
L273:
	;
	v1529 = base.I64_extend_i32_s(v1525)
	if int64(-9223372036854775807-1)-v1529 <= v1011 {
		v1536 = v1529
		goto L272
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1532 = base.I64_extend_i32_u(v1525)
	if v1532^int64(9223372036854775807) < v1011 {
		v1545 = v1519
		v1547 = v996
		goto L175
	} else {
		goto L277
	}
L276:
	;
	v1545 = v1519
	v1547 = v996
	goto L175
L277:
	;
	v1536 = v1532
	goto L272
L278:
	;
	goto L174
L279:
	;
	v1566 = v988 + int32(1)
	if v1566 < v1547 {
		v988 = v1566
		v989 = int32(0)
		v991 = v1545
		v996 = v1547
		v1011 = v1011 + v1562
		goto L173
	} else {
		goto L285
	}
L280:
	;
	v1555 = base.I64_extend_i32_s(v1551)
	if int64(-9223372036854775807-1)-v1555 <= v1011 {
		v1562 = v1555
		goto L279
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1558 = base.I64_extend_i32_u(v1551)
	if v1558^int64(9223372036854775807) < v1011 {
		v1568 = v988
		goto L278
	} else {
		goto L284
	}
L283:
	;
	v1568 = v988
	goto L278
L284:
	;
	v1562 = v1558
	goto L279
L285:
	;
	v1568 = v1566
	goto L278
L286:
	;
	goto L128
L287:
	;
	v1619 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v1619)
	goto L126
L288:
	;
	base.MemoryCopy(m, v1651, v341, v348)
	goto L290
L289:
	;
	goto L290
L290:
	;
	v1653 = v1651 + v348
	v1654 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1653))) = uint8(v1654)
	v1656 = int32(1)
	v1658 = v1653 + v1656
	if v516 != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	base.MemoryCopy(m, v1658, v496, v516)
	goto L293
L292:
	;
	goto L293
L293:
	;
	v1661 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1658+v516))) = uint8(v1661)
	v1667 = v1656
	goto L1
}
