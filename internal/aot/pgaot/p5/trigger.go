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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v172 int64
	_ = v172
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(256)
	m.G0 = v15
	v19 = l3 + l1*int32(416)
	v22 = l2 + l1*int32(28)
	v25 = l1
	v26 = l2
	v30 = v6
	v31 = v6
	v32 = v6
	v35 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	m.G0 = v15 + int32(256)
	return v120
L3:
	;
	goto L2
L4:
	;
	if v35 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v253 = int32(m.ExcTag)
	v254 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v253 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v99 = v25
	v100 = v26
	v101 = v30
	v102 = v31
	v103 = v32
	goto L9
L9:
	;
	if v103 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v31
	v47 = v30 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v47)
	F_fmgr_info(m, v42, v22)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v52 = base.B2i32(l3 == int32(0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v52)
	F_InstrStartNode(m, v19)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v59 = int32(_a_F_ExecCallTriggerFunc_0)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[0])) = l4
	v63 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+238)) = uint16(v63)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+220)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v15)+228)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+236)) = uint8(v63)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v60
	F_pgstat_init_function_usage(m, v15+int32(220), v15+int32(184))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v82 = int32(_a_F_ExecCallTriggerFunc_1)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1])) = v84 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[2]))
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[3]))
	goto L19
L19:
	;
	v93 = v15 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v15 + int32(12)
	goto L22
L20:
	;
	v99 = v91
	v100 = v89
	v101 = v52
	v102 = v60
	v103 = int32(0)
	goto L9
L22:
	;
	goto L20
L23:
	;
	if v116 != 0 {
		goto L3
	} else {
		goto L41
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[2])) = v15 + int32(16)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+220))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v102
	v116 = v101 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v116)
	v120 = m.T0[v111].(func(*base.Module, int32) int32)(m, v15+int32(220))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[3])) = v99
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[2])) = v100
	v227 = int32(_a_F_ExecCallTriggerFunc_1)
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1]))
	v230 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1])) = v229 - v230
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v102
	v237 = v101 & v230
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v237)
	F_pg_re_throw(m)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L40
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[2])) = v100
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[3])) = v99
	v126 = int32(_a_F_ExecCallTriggerFunc_1)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[1])) = v128 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v102
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v116)
	v137 = v15 + int32(184)
	v145 = m.G0
	v147 = v145 - int32(16)
	m.G0 = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if v149 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[0])) = v102
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+236)))
	if v186 == int32(0) {
		goto L23
	} else {
		goto L35
	}
L29:
	;
	F___clock_gettime(m, int32(1), v147)
	mBase = m.M
	v152 = int32(_a_F_ExecCallTriggerFunc_2)
	v153 = *(*int64)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[4]))
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v137)+16))
	v156 = int64(*(*int32)(unsafe.Add(mBase, uint32(v147)+8)))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v137)+24))
	v162 = v156 + v157*int64(1000000000) - v161
	*(*int64)(unsafe.Add(mBase, _c_F_ExecCallTriggerFunc[4])) = v155 + v162
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v137)+8))
	goto L32
L30:
	;
	goto L31
L31:
	;
	m.G0 = v147 + int32(16)
	goto L28
L32:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
	*(*int64)(unsafe.Add(mBase, uint32(v149))) = v167 + int64(1)
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v149)+8)) = v165 + v162
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v149)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v149)+16)) = v172 + (v162 - v153 + v155)
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v102
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v116)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v102
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v116)
	F_errcode(m, int32(16908867))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v15)+220))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v102
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v116)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v205
	F_errmsg(m, int32(_a_F_ExecCallTriggerFunc_3), v15)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v102
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v116)
	F_errfinish(m, int32(_a_F_ExecCallTriggerFunc_4), int32(2389), int32(_a_F_ExecCallTriggerFunc_5))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v102
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v116)
	F_InstrStopNode(m, v19, float64(1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L5
L43:
	;
	v258 = int32(v254)
	m.G0 = v15
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v15+int32(12) == v264 {
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
	if v268 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	v268 = v266
	goto L48
L47:
	;
	v268 = int32(0)
	goto L48
L48:
	;
	goto L45
L49:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v15)+248))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v15)+244))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v15)+240))
	v25 = v272
	v26 = v271
	v30 = v269
	v31 = v270
	v32 = v260
	v35 = v268
	goto L1
L50:
	;
	goto L51
L51:
	;
	F___wasm_longjmp(m, v261, v260)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
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
