package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommandCounterIncrement(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[0])))
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L60
	} else {
		goto L70
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[1])) = v20
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L60
	} else {
		goto L66
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L60
	} else {
		goto L62
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[2]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v12 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	return
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+76)))
	if v13 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[3]))
	if int32(0) <= v15 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v18 = int32(_a_F_CommandCounterIncrement_0)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[1]))
	v22 = v20 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[1])) = v22
	if v22 == int32(-1) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[0])) = uint8(v27)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[4])))
	if v30 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[5]))
	if v43 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[6]))
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v22
	goto L15
L14:
	;
	goto L15
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[7]))
	if v37 == int32(0) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v22
	goto L11
L17:
	;
	F_CommandEndInvalidationMessages(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L60
	} else {
		goto L61
	}
L18:
	;
	if int32(0) < v43 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[8]))
	if v130 != 0 {
		goto L39
	} else {
		goto L40
	}
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[9]))
	v48 = v47
	v50 = int32(0)
	goto L24
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[5])) = int32(0)
	goto L20
L24:
	;
	v56 = v50 << (uint(int32(3)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_CommandCounterIncrement[10])))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_CommandCounterIncrement[11])))
	v59 = int32(0)
	if v59 < v48 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L23
L26:
	;
	v110 = v50 + int32(1)
	if v110 != v43 {
		v48 = v102
		v50 = v110
		goto L24
	} else {
		goto L38
	}
L27:
	;
	v63 = v59
	goto L31
L28:
	;
	goto L29
L29:
	;
	v89 = v48 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_c_F_CommandCounterIncrement[12]))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_c_F_CommandCounterIncrement[13]))) = v58
	v96 = int32(_a_F_CommandCounterIncrement_1)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[9]))
	v100 = v98 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[9])) = v100
	v102 = v100
	goto L26
L30:
	;
	if int32(64) <= v48 {
		goto L1
	} else {
		goto L37
	}
L31:
	;
	v70 = v63 << (uint(int32(3)) % 32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_CommandCounterIncrement[13])))
	if v71 != v58 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_CommandCounterIncrement[12]))) = v57
	v102 = v48
	goto L26
L33:
	;
	v74 = v63 + int32(1)
	if v48 != v74 {
		v63 = v74
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L30
L37:
	;
	goto L29
L38:
	;
	goto L25
L39:
	;
	if int32(0) < v130 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	goto L17
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[14]))
	v136 = v134
	v138 = int32(0)
	goto L45
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[8])) = int32(0)
	goto L41
L45:
	;
	v144 = v138 << (uint(int32(3)) % 32)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+uint32(_c_F_CommandCounterIncrement[15])))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v144)+uint32(_c_F_CommandCounterIncrement[16])))
	v147 = int32(0)
	if v147 < v136 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L44
L47:
	;
	v198 = v138 + int32(1)
	if v198 != v130 {
		v136 = v190
		v138 = v198
		goto L45
	} else {
		goto L59
	}
L48:
	;
	v151 = v147
	goto L52
L49:
	;
	goto L50
L50:
	;
	v177 = v136 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_CommandCounterIncrement[17]))) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_CommandCounterIncrement[18]))) = v146
	v184 = int32(_a_F_CommandCounterIncrement_2)
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[14]))
	v188 = v186 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CommandCounterIncrement[14])) = v188
	v190 = v188
	goto L47
L51:
	;
	if int32(64) <= v136 {
		goto L1
	} else {
		goto L58
	}
L52:
	;
	v158 = v151 << (uint(int32(3)) % 32)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+uint32(_c_F_CommandCounterIncrement[18])))
	if v159 != v146 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+uint32(_c_F_CommandCounterIncrement[17]))) = v145
	v190 = v136
	goto L47
L54:
	;
	v162 = v151 + int32(1)
	if v136 != v162 {
		v151 = v162
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L51
L58:
	;
	goto L50
L59:
	;
	goto L46
L60:
	;
	return
L61:
	;
	goto L6
L62:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_CommandCounterIncrement_3), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L60
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_CommandCounterIncrement_4), int32(1118), int32(_a_F_CommandCounterIncrement_5))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L60
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L60
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_CommandCounterIncrement_6), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L60
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_CommandCounterIncrement_4), int32(1126), int32(_a_F_CommandCounterIncrement_5))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L60
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
	F_errmsg_internal(m, int32(_a_F_CommandCounterIncrement_7), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L60
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_CommandCounterIncrement_8), int32(403), int32(_a_F_CommandCounterIncrement_9))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L60
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CommandEndInvalidationMessages(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_CommandEndInvalidationMessages[0]))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v6 < v7 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v9 = v6
	goto L7
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v25 < v26 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_CommandEndInvalidationMessages[1]))
	F_LocalExecuteInvalidationMessage(m, v13+v9<<(uint(int32(4))%32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	v20 = v9 + int32(1)
	if v20 != v7 {
		v9 = v20
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v28 = v25
	goto L15
L13:
	;
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_CommandEndInvalidationMessages[2]))
	if int32(2) <= v45 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_CommandEndInvalidationMessages[3]))
	F_LocalExecuteInvalidationMessage(m, v32+v28<<(uint(int32(4))%32))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v39 = v28 + int32(1)
	if v39 != v26 {
		v28 = v39
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_LogLogicalInvalidations(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_CommandEndInvalidationMessages[0]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+32)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v55
	goto L3
L22:
	;
	goto L21
}
func F_EndCommand(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = int32(2)
	if base.Ui32(l1-v10) <= base.Ui32(v10) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = v14 << (uint(int32(3)) % 32)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_EndCommand[0]))))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_EndCommand[1])))
			base.MemoryCopy(m, v8, v18, v17)
		} else {
		}
		v20 = v17 + v8
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_EndCommand[2]))))
		if v21&int32(1) != 0 {
			if v14 == int32(158) {
				v26 = int32(_a_F_EndCommand_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v20))) = uint16(v26)
				v30 = v20 + int32(2)
			} else {
				v30 = v20
			}
			v31 = int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v31)
			v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			v35 = v30 + int32(1)
			v36 = int32(0)
			if v33 == int64(0) {
				v45 = int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v45)
				v216 = int32(1)
			} else {
				v52 = int32(1233)
				v57 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(v33))^int32(63))*v52+v52) >> (uint(int32(12)) % 32))
				v60 = *(*int64)(unsafe.Add(mBase, uint32(v57<<(uint(int32(3))%32))+uint32(_c_F_EndCommand[3])))
				v62 = v57 + base.B2i32(base.Ui64(v60) <= base.Ui64(v33))
				if base.Ui64(int64(100000000)) <= base.Ui64(v33) {
					v66 = v33
					v69 = v36
					for {
						v75 = v35 + v62 - v69
						v76 = int32(8)
						v79 = base.I64_div_u_s(v66, int64(100000000))
						v83 = base.I32_wrap_i64(v66 + v79*int64(4194967296))
						v85 = base.I32_div_u_s(v83, int32(_a_F_EndCommand_1))
						v86 = int32(1)
						v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85<<(uint(v86)%32))+uint32(_c_F_EndCommand[4]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v75-v76))) = uint16(v88)
						v92 = int32(_a_F_EndCommand_2)
						v93 = base.I32_div_u_s(v83, v92)
						v94 = int32(100)
						v95 = base.I32_rem_u_s(v93, v94)
						v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95<<(uint(v86)%32))+uint32(_c_F_EndCommand[4]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v75-int32(6)))) = uint16(v98)
						v104 = v83 - v93*v92
						v105 = int32(_a_F_EndCommand_3)
						v108 = base.I32_div_u_s(v104&v105, v94)
						v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108<<(uint(v86)%32))+uint32(_c_F_EndCommand[4]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v75-int32(4)))) = uint16(v111)
						v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v104-v108*v94)&v105<<(uint(v86)%32))+uint32(_c_F_EndCommand[4]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v75-int32(2)))) = uint16(v122)
						v125 = v69 + v76
						if base.Ui64(int64(9999999999999999)) < base.Ui64(v66) {
							v66 = v79
							v69 = v125
							continue
						} else {
							break
						}
						break
					}
					v128 = v79
					v131 = v125
				} else {
					v128 = v33
					v131 = v36
				}
				v137 = base.I32_wrap_i64(v128)
				if base.Ui64(int64(10000)) <= base.Ui64(v128) {
					v141 = v35 + v62 - v131
					v142 = int32(4)
					v145 = base.I32_div_u_s(v137, int32(_a_F_EndCommand_2))
					v148 = v137 + v145*int32(-10000)
					v149 = int32(100)
					v150 = base.I32_div_u_s(v148, v149)
					v151 = int32(1)
					v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150<<(uint(v151)%32))+uint32(_c_F_EndCommand[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v141-v142))) = uint16(v153)
					v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v148-v150*v149)<<(uint(v151)%32))+uint32(_c_F_EndCommand[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v141-int32(2)))) = uint16(v162)
					v166 = v145
					v167 = v131 | v142
				} else {
					v166 = v137
					v167 = v131
				}
				if base.Ui32(int32(100)) <= base.Ui32(v166) {
					v175 = int32(2)
					v177 = int32(_a_F_EndCommand_3)
					v179 = int32(100)
					v180 = base.I32_div_u_s(v166&v177, v179)
					v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v166-v180*v179)&v177<<(uint(int32(1))%32))+uint32(_c_F_EndCommand[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v35+v62-v167-v175))) = uint16(v188)
					v192 = v180
					v193 = v167 + v175
				} else {
					v192 = v166
					v193 = v167
				}
				if base.Ui32(int32(10)) <= base.Ui32(v192) {
					v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192<<(uint(int32(1))%32))+uint32(_c_F_EndCommand[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v35+v62-v193-int32(2)))) = uint16(v202)
					v216 = v62
				} else {
					v205 = v192 | int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v205)
					v216 = v62
				}
			}
			v219 = v216 + v35
		} else {
			v219 = v20
		}
		v220 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v220)
		v227 = *(*int32)(unsafe.Add(mBase, _c_F_EndCommand[5]))
		v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
		v229 = m.T0[v228].(func(*base.Module, int32, int32, int32) int32)(m, int32(67), v8, v219-v8+int32(1))
		mBase = m.M
		v230 = m.ExcPending
		if v230 != 0 {
			return
		} else {
			m.G0 = v8 - int32(-64)
			return
		}
	} else {
		m.G0 = v8 - int32(-64)
		return
	}
}
