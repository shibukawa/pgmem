package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecCallTriggerFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v174 int64
	_ = v174
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int64
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(256)
	m.G0 = v16
	v20 = l3 + l1*int32(368)
	v23 = l2 + l1*int32(28)
	v26 = l1
	v27 = l2
	v31 = v6
	v32 = v6
	v33 = v6
	v36 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	m.G0 = v16 + int32(256)
	return base.I32_wrap_i64(v122)
L3:
	;
	goto L2
L4:
	;
	if v36 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v255 = int32(m.ExcTag)
	v256 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v255 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v101 = v26
	v102 = v27
	v103 = v31
	v104 = v32
	v105 = v33
	goto L9
L9:
	;
	if v105 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v32
	v49 = v31 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v49)
	F_fmgr_info(m, v44, v23)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v54 = base.B2i32(l3 == int32(0))
	if l3 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v32
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v54)
	F_InstrStart(m, v20)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v61 = int32(_a_F_ExecCallTriggerFunc_0)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[0])) = l4
	v65 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+234)) = uint16(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+220)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v16)+216)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v16)+224)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+232)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v27
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v62
	F_pgstat_init_function_usage(m, v16+int32(216), v16+int32(184))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v84 = int32(_a_F_ExecCallTriggerFunc_1)
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1])) = v86 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[2]))
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[3]))
	goto L19
L19:
	;
	v95 = v16 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v16 + int32(12)
	goto L22
L20:
	;
	v101 = v93
	v102 = v91
	v103 = v54
	v104 = v62
	v105 = int32(0)
	goto L9
L22:
	;
	goto L20
L23:
	;
	if v118 != 0 {
		goto L3
	} else {
		goto L41
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[2])) = v16 + int32(16)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v16)+216))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v104
	v118 = v103 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v118)
	v122 = m.T0[v113].(func(*base.Module, int32) int64)(m, v16+int32(216))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[3])) = v101
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[2])) = v102
	v229 = int32(_a_F_ExecCallTriggerFunc_1)
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1]))
	v232 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1])) = v231 - v232
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v104
	v239 = v103 & v232
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v239)
	F_pg_re_throw(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L6
	} else {
		goto L40
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[2])) = v102
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[3])) = v101
	v128 = int32(_a_F_ExecCallTriggerFunc_1)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1])) = v130 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v118)
	v139 = v16 + int32(184)
	v147 = m.G0
	v149 = v147 - int32(16)
	m.G0 = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v151 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[0])) = v104
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+232)))
	if v188 == int32(0) {
		goto L23
	} else {
		goto L35
	}
L29:
	;
	F___clock_gettime(m, int32(1), v149)
	mBase = m.M
	v154 = int32(_a_F_ExecCallTriggerFunc_2)
	v155 = *(*int64)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[4]))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v139)+16))
	v158 = int64(*(*int32)(unsafe.Add(mBase, uint32(v149)+8)))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v139)+24))
	v164 = v158 + v159*int64(1000000000) - v163
	*(*int64)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[4])) = v157 + v164
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v139)+8))
	goto L32
L30:
	;
	goto L31
L31:
	;
	m.G0 = v149 + int32(16)
	goto L28
L32:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
	*(*int64)(unsafe.Add(mBase, uint32(v151))) = v169 + int64(1)
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v151)+8)) = v167 + v164
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v151)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+16)) = v174 + (v164 - v155 + v157)
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v118)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v118)
	F_errcode(m, int32(16908867))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v16)+216))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v118)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v207
	F_errmsg(m, int32(_a_F_ExecCallTriggerFunc_3), v16)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v118)
	F_errfinish(m, int32(_a_F_ExecCallTriggerFunc_4), int32(2404), int32(_a_F_ExecCallTriggerFunc_5))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)) = uint8(v118)
	F_InstrStopTrigger(m, v20)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L5
L43:
	;
	v260 = int32(v256)
	m.G0 = v16
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	if v16+int32(12) == v266 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	m.ExcPending = 1
	goto L52
L45:
	;
	if v270 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v270 = v268
	goto L48
L47:
	;
	v270 = int32(0)
	goto L48
L48:
	;
	goto L45
L49:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v16)+248))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v16)+244))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v16)+240))
	v26 = v274
	v27 = v273
	v31 = v271
	v32 = v272
	v33 = v262
	v36 = v270
	goto L1
L50:
	;
	goto L51
L51:
	;
	F___wasm_longjmp(m, v263, v262)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	return int32(0)
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
