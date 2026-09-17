package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ScanKeyEntryInitialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	v3 = l2
	v4 = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l4
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v16 = l0 + int32(16)
	if l6 != 0 {
		F_fmgr_info(m, l6, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(0)
		v21 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v16))) = v21
		return
	}
}
func F_SerialPagePrecedesLogically(m *base.Module, l0 int64, l1 int64) int32 {
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = int32(10)
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(v4)%32) | v9
	v12 = base.I32_wrap_i64(l1) << (uint(v4) % 32)
	v15 = F_TransactionIdPrecedes(m, v10, v12|v9)
	if v15 != 0 {
		v17 = F_TransactionIdPrecedes(m, v10, v12+int32(1027))
		v19 = v17
	} else {
		v19 = int32(0)
	}
	return v19
}
func F_ShmemAllocNoError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	if v7 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[0]))
		F_s_lock(m, v11, int32(_a_F_ShmemAllocNoError_0), int32(208), int32(_a_F_ShmemAllocNoError_1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[1]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
			v26 = (l0+int32(127))&int32(-128) + v25
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
			if base.Ui32(v26) <= base.Ui32(v27) {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v26
				v33 = v30 + v25
			} else {
				v33 = int32(0)
			}
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(0)
			return v33
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[1]))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
		v26 = (l0+int32(127))&int32(-128) + v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
		if base.Ui32(v26) <= base.Ui32(v27) {
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v26
			v33 = v30 + v25
		} else {
			v33 = int32(0)
		}
		v35 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(0)
		return v33
	}
}
func F_SignalHandlerForConfigReload(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForConfigReload[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForConfigReload[1]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_SplitIdentifierString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v11 = l0
	goto L1
L1:
	;
	v21 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
	goto L3
L2:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v31 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if base.B2i32(v21 == int32(32))|base.B2i32(base.Ui32((v21-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v11 = v11 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	return int32(1)
L6:
	;
	v34 = v31
	v37 = v11
	goto L7
L7:
	;
	v43 = v34 & int32(255)
	if v43 != int32(34) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v308 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v240))) = uint8(v308)
	v310 = F_strlen(m, v241)
	mBase = m.M
	F_truncate_identifier(m, v241, v310, v308)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L27
	} else {
		goto L83
	}
L10:
	;
	return int32(0)
L11:
	;
	v244 = v236
	goto L71
L12:
	;
	if v43 == int32(0) {
		v79 = v37
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v202 = v37 + int32(1)
	v203 = int32(34)
	v204 = F___strchrnul(m, v202, v203)
	mBase = m.M
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v206 == v203 {
		goto L56
	} else {
		goto L57
	}
L15:
	;
	if v37 == v79 {
		goto L10
	} else {
		goto L26
	}
L16:
	;
	v49 = l1 & int32(255)
	if v43 == v49 {
		v79 = v37
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v51 = v34
	v56 = v37
	goto L18
L18:
	;
	v59 = base.I32_extend8_s(v51)
	goto L20
L19:
	;
	v79 = v70
	goto L15
L20:
	;
	if base.B2i32(v59 == int32(32))|base.B2i32(base.Ui32((v59-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v79 = v56
	goto L15
L22:
	;
	goto L23
L23:
	;
	v70 = v56 + int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v71 == int32(0) {
		v79 = v70
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v71 != v49 {
		v51 = v71
		v56 = v70
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v84 = v79 - v37
	v86 = F_downcase_truncate_identifier(m, v37, v84, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	if (v86^v37)&int32(3) != 0 {
		v160 = v86
		v161 = v84
		v162 = v37
		goto L33
	} else {
		goto L34
	}
L29:
	;
	F_pfree(m, v86)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L27
	} else {
		goto L54
	}
L30:
	;
	F___memset(m, v195, int32(0), v194)
	mBase = m.M
	goto L29
L31:
	;
	v194 = int32(0)
	v195 = v189
	goto L30
L32:
	;
	v172 = v167
	v173 = v168
	v174 = v169
	goto L50
L33:
	;
	if v161 == int32(0) {
		v189 = v162
		goto L31
	} else {
		goto L49
	}
L34:
	;
	v95 = int32(0)
	if base.B2i32(v86&int32(3) == v95)|base.B2i32(v84 == v95) != 0 {
		v126 = v86
		v127 = v84
		v128 = v37
		v129 = base.B2i32(v84 != v95)
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v129 == int32(0) {
		v189 = v128
		goto L31
	} else {
		goto L42
	}
L36:
	;
	v105 = v86
	v106 = v84
	v107 = v37
	goto L37
L37:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v109)
	if v109 == int32(0) {
		v194 = v106
		v195 = v107
		goto L30
	} else {
		goto L39
	}
L38:
	;
	v126 = v120
	v127 = v116
	v128 = v114
	v129 = v118
	goto L35
L39:
	;
	v113 = int32(1)
	v114 = v107 + v113
	v116 = v106 - v113
	v117 = int32(0)
	v118 = base.B2i32(v116 != v117)
	v120 = v105 + v113
	if v120&int32(3) == v117 {
		v126 = v120
		v127 = v116
		v128 = v114
		v129 = v118
		goto L35
	} else {
		goto L40
	}
L40:
	;
	if v116 != 0 {
		v105 = v120
		v106 = v116
		v107 = v114
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v132 == int32(0) {
		v194 = v127
		v195 = v128
		goto L30
	} else {
		goto L43
	}
L43:
	;
	if base.Ui32(v127) < base.Ui32(int32(4)) {
		v160 = v126
		v161 = v127
		v162 = v128
		goto L33
	} else {
		goto L44
	}
L44:
	;
	v138 = v126
	v139 = v127
	v140 = v128
	goto L45
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v146 = int32(-2139062144)
	if (int32(16843008)-v143|v143)&v146 != v146 {
		v167 = v138
		v168 = v139
		v169 = v140
		goto L32
	} else {
		goto L47
	}
L46:
	;
	v160 = v154
	v161 = v156
	v162 = v152
	goto L33
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v143
	v151 = int32(4)
	v152 = v140 + v151
	v154 = v138 + v151
	v156 = v139 - v151
	if base.Ui32(int32(3)) < base.Ui32(v156) {
		v138 = v154
		v139 = v156
		v140 = v152
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v167 = v160
	v168 = v161
	v169 = v162
	goto L32
L50:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v176)
	if v176 == int32(0) {
		v194 = v173
		v195 = v174
		goto L30
	} else {
		goto L52
	}
L51:
	;
	v189 = v181
	goto L31
L52:
	;
	v180 = int32(1)
	v181 = v174 + v180
	v185 = v173 - v180
	if v185 != 0 {
		v172 = v172 + v180
		v173 = v185
		v174 = v181
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v236 = v79
	v240 = v79
	v241 = v37
	goto L11
L55:
	;
	if v210 == int32(0) {
		goto L10
	} else {
		goto L59
	}
L56:
	;
	v210 = v204
	goto L58
L57:
	;
	v210 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v217 = v210
	goto L60
L60:
	;
	v222 = v217 + int32(1)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v223 != int32(34) {
		v236 = v222
		v240 = v217
		v241 = v202
		goto L11
	} else {
		goto L62
	}
L61:
	;
	goto L10
L62:
	;
	v226 = F_strlen(m, v217)
	mBase = m.M
	if v226 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	base.MemoryCopy(m, v217, v222, v226)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v228 = int32(34)
	v229 = F___strchrnul(m, v222, v228)
	mBase = m.M
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v231 == v228 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v235 != 0 {
		v217 = v235
		goto L60
	} else {
		goto L70
	}
L67:
	;
	v235 = v229
	goto L69
L68:
	;
	v235 = int32(0)
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L61
L71:
	;
	v254 = int32(*(*int8)(unsafe.Add(mBase, uint32(v244))))
	goto L73
L72:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v266 = l1 & int32(255)
	if v264 == v266 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	if base.B2i32(v254 == int32(32))|base.B2i32(base.Ui32((v254-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v244 = v244 + int32(1)
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v271 = v244
	goto L78
L76:
	;
	goto L77
L77:
	;
	if v264 == int32(0) {
		v303 = v244
		goto L9
	} else {
		goto L82
	}
L78:
	;
	v277 = v271 + int32(1)
	v278 = int32(*(*int8)(unsafe.Add(mBase, uint32(v271)+1)))
	goto L80
L80:
	;
	if base.B2i32(v278 == int32(32))|base.B2i32(base.Ui32((v278-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v271 = v277
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v303 = v277
	goto L9
L82:
	;
	goto L10
L83:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v315 = F_lappend(m, v314, v241)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L27
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v315
	if v264 != v266 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	v34 = v319
	v37 = v303
	goto L7
}
func F_StatementTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_StatementTimeoutHandler[0]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StatementTimeoutHandler[1])))
	if v9 != 0 {
		v10 = int32(15)
	} else {
		v10 = int32(2)
	}
	v11 = F_kill(m, int32(0)-v4, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_StatementTimeoutHandler[0]))
		v15 = F_kill(m, v14, v10)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64 {
	var v7 float64
	_ = v7
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	v7 = base.F64_mul(l0, l0)
	v22 = base.F64_add(base.F64_mul(base.F64_mul(v7, base.F64_mul(v7, v7)), base.F64_add(base.F64_mul(v7, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))
	v23 = base.F64_mul(l0, v7)
	if l2 == int32(0) {
		return base.F64_add(base.F64_mul(v23, base.F64_add(base.F64_mul(v7, v22), float64(-0.16666666666666632))), l0)
	} else {
		return base.F64_sub(l0, base.F64_add(base.F64_sub(base.F64_mul(v7, base.F64_sub(base.F64_mul(l1, float64(0.5)), base.F64_mul(v23, v22))), l1), base.F64_mul(v23, float64(0.16666666666666632))))
	}
}
func F___stdio_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l2 - base.B2i32(v13 != v4)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v18
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v27 = m.Wasi_snapshot_preview1.Fd_read(m, v21, v10+int32(16), int32(2), v10+int32(12))
	mBase = m.M
	if v27 == v4 {
		v34 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___stdio_read[0])) = v27
		v34 = int32(-1)
	}
	if v34 != 0 {
		v43 = int32(32)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v43 | v44
		v65 = v4
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		if int32(0) < v36 {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			if base.Ui32(v36) <= base.Ui32(v47) {
				v65 = v36
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v49 + (v36 - v47)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v54 != 0 {
					v55 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 + v55
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1+l2-v55))) = uint8(v61)
				} else {
				}
				v65 = l2
			}
		} else {
			if v36 != 0 {
				v41 = int32(32)
			} else {
				v41 = int32(16)
			}
			v43 = v41
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v43 | v44
			v65 = v4
		}
	}
	m.G0 = v10 + int32(32)
	return v65
}
func F___stdio_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v5 = F___lseek(m, v4, l1, l2)
	mBase = m.M
	return v5
}
func F_s_lock_stuck(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		if l2 != 0 {
			v16 = l2
		} else {
			v16 = int32(_a_F_s_lock_stuck_0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
		F_errmsg_internal(m, int32(_a_F_s_lock_stuck_1), v7)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_errfinish(m, int32(_a_F_s_lock_stuck_2), int32(90), int32(_a_F_s_lock_stuck_3))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_save_ps_display_args(m *base.Module, l0 int32, l1 int32) int32 {
	return l1
}
func F_sbrk(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_sbrk[0]))
	v12 = base.I64_extend_i32_u(v5) + (base.I64_extend_i32_u(l0)+int64(7))&int64(8589934584)
	if base.Ui64(v12) <= base.Ui64(int64(4294967295)) {
		v15 = base.I32_wrap_i64(v12)
		if base.Ui32(v15) <= base.Ui32(base.MemorySize(m)<<(uint(int32(16))%32)) {
			*(*int32)(unsafe.Add(mBase, _c_F_sbrk[0])) = v15
			return v5
		} else {
			v20 = m.Env.Emscripten_resize_heap(m, v15)
			mBase = m.M
			if v20 != 0 {
				*(*int32)(unsafe.Add(mBase, _c_F_sbrk[0])) = v15
				return v5
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_sbrk[1])) = int32(48)
				return int32(-1)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_sbrk[1])) = int32(48)
		return int32(-1)
	}
}
func F_scalarltjoinsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.3333333333333333))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_scalarltsel(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_scalarineqsel_wrapper(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_searchstoplist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 == v3 {
		v28 = v3
		m.G0 = v7 + int32(16)
		return v28
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v14 <= int32(0) {
			v28 = v3
			m.G0 = v7 + int32(16)
			return v28
		} else {
			v21 = F_bsearch(m, v7+int32(12), v11, v14, int32(4), int32(1170))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v28 = base.B2i32(v21 != int32(0))
				m.G0 = v7 + int32(16)
				return v28
			}
		}
	}
}
func F_send(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = int32(0)
	v6 = F_sendto(m, l0, l1, l2, v4, v4)
	return v6
}
func F_send_feedback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
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
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v352 int64
	_ = v352
	var v355 int64
	_ = v355
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v361 int64
	_ = v361
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int64
	_ = v404
	var v405 int64
	_ = v405
	var v410 int64
	_ = v410
	var v414 int64
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int64
	_ = v436
	var v441 int64
	_ = v441
	var v446 int64
	_ = v446
	v3 = l2
	v7 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[0]))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[1]))
	if base.Ui64(v21) < base.Ui64(l0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v23 = l0
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	v27 = int32(_a_F_send_feedback_0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[2]))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+280)) = v29
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[3])) = v29
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[2]))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+272)) = v35
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[4])) = v35
	goto L11
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[5]))
	v47 = int32(0)
	if base.B2i32(v46 == v47)|base.B2i32(v46 == int32(_a_F_send_feedback_1)) == v47 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[3]))
	goto L9
L13:
	;
	v115 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[6]))
	if base.Ui64(v115) < base.Ui64(v110) {
		goto L33
	} else {
		goto L34
	}
L14:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v46)+8))
	if base.Ui64(v54) <= base.Ui64(v44) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v94 = v46
	v97 = v7
	goto L16
L16:
	;
	if v94 != int32(_a_F_send_feedback_1) {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[5]))
	v94 = v90
	v97 = v65
	goto L16
L18:
	;
	v59 = v46
	goto L21
L19:
	;
	v83 = v7
	goto L20
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[7]))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
	v110 = v83
	v113 = v88
	goto L13
L21:
	;
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v59)+16))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v69
	F_pfree(m, v59)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v83 = v65
	goto L20
L23:
	;
	return
L24:
	;
	if v67 == int32(_a_F_send_feedback_1) {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	if base.Ui64(v75) <= base.Ui64(v44) {
		v59 = v67
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v102 = v97
	goto L29
L28:
	;
	v102 = v23
	goto L29
L29:
	;
	if v94 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v103 = v102
	goto L32
L31:
	;
	v103 = v23
	goto L32
L32:
	;
	v110 = v103
	v113 = v103
	goto L13
L33:
	;
	v117 = v110
	goto L35
L34:
	;
	v117 = v115
	goto L35
L35:
	;
	v119 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[8]))
	if base.Ui64(v119) < base.Ui64(v113) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v121 = v113
	goto L38
L37:
	;
	v121 = v119
	goto L38
L38:
	;
	v125 = m.G0
	v126 = int32(16)
	v127 = v125 - v126
	m.G0 = v127
	F_gettimeofday(m, v127)
	mBase = m.M
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
	v131 = int64(*(*int32)(unsafe.Add(mBase, uint32(v127)+8)))
	m.G0 = v127 + v126
	v139 = v131 + v130*int64(1000000) - int64(946684800000000)
	goto L39
L39:
	;
	if l1 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[9])) = v139
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	if v162 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v141 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[8]))
	if v121 != v141 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v144 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[6]))
	if v117 != v144 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v147 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[9]))
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[0]))
	goto L44
L44:
	;
	if base.B2i32(base.I64_extend_i32_s(v149*int32(1000))*int64(1000) <= v139-v147) == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	F_enlargeStringInfo(m, v186, int32(1))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L23
	} else {
		goto L52
	}
L47:
	;
	v165 = int32(_a_F_send_feedback_2)
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[11]))
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_send_feedback[11])) = v169
	v171 = F_makeStringInfo(m)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L23
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v178)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+12)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v178
	goto L51
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_send_feedback[11])) = v166
	*(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10])) = v171
	v186 = v171
	goto L46
L51:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	v186 = v185
	goto L46
L52:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v194 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v192))) = uint8(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v191 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v200, int32(8))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v207 = int64(56)
	v209 = int64(65280)
	v211 = int64(40)
	v214 = int64(16711680)
	v216 = int64(24)
	v218 = int64(4278190080)
	v220 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v204+v205))) = v23<<(uint(v207)%64) | v23&v209<<(uint(v211)%64) | (v23&v214<<(uint(v216)%64) | v23&v218<<(uint(v220)%64)) | (int64(base.Ui64(v23)>>(uint(v220)%64))&v218 | int64(base.Ui64(v23)>>(uint(v216)%64))&v214 | (int64(base.Ui64(v23)>>(uint(v211)%64))&v209 | int64(base.Ui64(v23)>>(uint(v207)%64))))
	v243 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v204 + v243
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v247, v243)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v254 = int64(56)
	v256 = int64(65280)
	v258 = int64(40)
	v261 = int64(16711680)
	v263 = int64(24)
	v265 = int64(4278190080)
	v267 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v251+v252))) = v117<<(uint(v254)%64) | v117&v256<<(uint(v258)%64) | (v117&v261<<(uint(v263)%64) | v117&v265<<(uint(v267)%64)) | (int64(base.Ui64(v117)>>(uint(v267)%64))&v265 | int64(base.Ui64(v117)>>(uint(v263)%64))&v261 | (int64(base.Ui64(v117)>>(uint(v258)%64))&v256 | int64(base.Ui64(v117)>>(uint(v254)%64))))
	v290 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v247)+4)) = v251 + v290
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v294, v290)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v301 = int64(56)
	v303 = int64(65280)
	v305 = int64(40)
	v308 = int64(16711680)
	v310 = int64(24)
	v312 = int64(4278190080)
	v314 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v298+v299))) = v121<<(uint(v301)%64) | v121&v303<<(uint(v305)%64) | (v121&v308<<(uint(v310)%64) | v121&v312<<(uint(v314)%64)) | (int64(base.Ui64(v121)>>(uint(v314)%64))&v312 | int64(base.Ui64(v121)>>(uint(v310)%64))&v308 | (int64(base.Ui64(v121)>>(uint(v305)%64))&v303 | int64(base.Ui64(v121)>>(uint(v301)%64))))
	v337 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+4)) = v298 + v337
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v341, v337)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v348 = int64(56)
	v350 = int64(65280)
	v352 = int64(40)
	v355 = int64(16711680)
	v357 = int64(24)
	v359 = int64(4278190080)
	v361 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v345+v346))) = v139<<(uint(v348)%64) | v139&v350<<(uint(v352)%64) | (v139&v355<<(uint(v357)%64) | v139&v359<<(uint(v361)%64)) | (int64(base.Ui64(v139)>>(uint(v361)%64))&v359 | int64(base.Ui64(v139)>>(uint(v357)%64))&v355 | (int64(base.Ui64(v139)>>(uint(v352)%64))&v350 | int64(base.Ui64(v139)>>(uint(v348)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v341)+4)) = v345 + int32(8)
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	F_enlargeStringInfo(m, v388, int32(1))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	*(*uint8)(unsafe.Add(mBase, uint32(v392+v393))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v388)+4)) = v392 + int32(1)
	v401 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	if v401 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+24)) = uint32(v117)
	v404 = int64(32)
	v405 = int64(base.Ui64(v117) >> (uint(v404) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+20)) = uint32(v405)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+16)) = uint32(v121)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	v410 = int64(base.Ui64(v121) >> (uint(v404) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+12)) = uint32(v410)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+8)) = uint32(v23)
	v414 = int64(base.Ui64(v23) >> (uint(v404) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v414)
	F_errmsg_internal(m, int32(_a_F_send_feedback_3), v12)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L23
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[13]))
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[10]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_send_feedback[14]))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+44))
	m.T0[v432].(func(*base.Module, int32, int32, int32))(m, v425, v428, v429)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L23
	} else {
		goto L64
	}
L62:
	;
	F_errfinish(m, int32(_a_F_send_feedback_4), int32(3910), int32(_a_F_send_feedback_5))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v436 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[1]))
	if base.Ui64(v436) < base.Ui64(v23) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[1])) = v23
	goto L67
L66:
	;
	goto L67
L67:
	;
	v441 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[8]))
	if base.Ui64(v441) < base.Ui64(v121) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[8])) = v121
	goto L70
L69:
	;
	goto L70
L70:
	;
	v446 = *(*int64)(unsafe.Add(mBase, _c_F_send_feedback[6]))
	if base.Ui64(v117) <= base.Ui64(v446) {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_send_feedback[6])) = v117
	goto L1
}
func F_setCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v13 == int32(2) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L15
	} else {
		goto L47
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L15
	} else {
		goto L43
	}
L3:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v139
	m.G0 = v11 + int32(32)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_setCompoundAffixFlagValue[0])) = int32(0)
	v23 = F_strtox_2(m, l2, v11+int32(28), int32(10), int64(2147483648))
	mBase = m.M
	v24 = base.I32_wrap_i64(v23)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v34 = F_strlen(m, l2)
	mBase = m.M
	v36 = v34 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v36) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if l2 == v25 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_setCompoundAffixFlagValue[0]))
	if v28 == int32(68) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if base.Ui32(int32(_a_F_setCompoundAffixFlagValue_0)) <= base.Ui32(v24) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
	goto L3
L11:
	;
	if (l2^v58)&int32(3) != 0 {
		goto L25
	} else {
		goto L26
	}
L12:
	;
	v39 = F_palloc0(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v44 = (v34 + int32(8)) & int32(4088)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v44) <= base.Ui32(v45) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	return
L16:
	;
	v58 = v39
	goto L11
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v52 - v44
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v53 + v44
	v58 = v53
	goto L11
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v52 = v45
	v53 = v47
	goto L17
L19:
	;
	goto L20
L20:
	;
	v48 = int32(_a_F_setCompoundAffixFlagValue_1)
	v50 = F_palloc0(m, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v52 = v48
	v53 = v50
	goto L17
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v58
	goto L3
L23:
	;
	goto L22
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v114)
	if v114&int32(255) == int32(0) {
		goto L23
	} else {
		goto L39
	}
L25:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v113 = l2
	v114 = v66
	v115 = v58
	goto L24
L26:
	;
	goto L27
L27:
	;
	if l2&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v70 = l2
	v72 = v58
	goto L31
L29:
	;
	v84 = l2
	v86 = v58
	goto L30
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v91 = int32(-2139062144)
	if (int32(16843008)-v88|v88)&v91 != v91 {
		v113 = v84
		v114 = v88
		v115 = v86
		goto L24
	} else {
		goto L35
	}
L31:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v73)
	if v73 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L32:
	;
	v84 = v80
	v86 = v78
	goto L30
L33:
	;
	v77 = int32(1)
	v78 = v72 + v77
	v80 = v70 + v77
	if v80&int32(3) != 0 {
		v70 = v80
		v72 = v78
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v96 = v84
	v97 = v88
	v98 = v86
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v97
	v100 = int32(4)
	v101 = v98 + v100
	v103 = v96 + v100
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v108 = int32(-2139062144)
	if (int32(16843008)-v105|v105)&v108 == v108 {
		v96 = v103
		v97 = v105
		v98 = v101
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v113 = v103
	v114 = v105
	v115 = v101
	goto L24
L38:
	;
	goto L37
L39:
	;
	v122 = v113
	v124 = v115
	goto L40
L40:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)) = uint8(v125)
	v127 = int32(1)
	if v125 != 0 {
		v122 = v122 + v127
		v124 = v124 + v127
		goto L40
	} else {
		goto L42
	}
L41:
	;
	goto L23
L42:
	;
	goto L41
L43:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
	F_errmsg(m, int32(_a_F_setCompoundAffixFlagValue_2), v11)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_setCompoundAffixFlagValue_3), int32(1041), int32(_a_F_setCompoundAffixFlagValue_4))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L15
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
	F_errcode(m, int32(22))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l2
	F_errmsg(m, int32(_a_F_setCompoundAffixFlagValue_5), v11+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_setCompoundAffixFlagValue_3), int32(1045), int32(_a_F_setCompoundAffixFlagValue_4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_set_deparse_for_query(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
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
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v954 int32
	_ = v954
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	v4 = int32(0)
	base.MemoryFill(m, l0, v4, int32(80))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v22
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v26
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
	F_set_rtable_names(m, l0, l2, v4)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v46 = v4
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v46
	if v46 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v69 != 0 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v58 = v56
	goto L7
L6:
	;
	v58 = int32(0)
	goto L7
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v59 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v62 = v60
	goto L10
L9:
	;
	v62 = int32(0)
	goto L10
L10:
	;
	if v58 < v62 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v65 = F_palloc0(m, int32(52))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L4
L14:
	;
	v67 = F_lappend(m, v46, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v46 = v67
	goto L3
L16:
	;
	v70 = F_has_dangerous_join_using(m, l0, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v79 = v46
	v80 = v59
	goto L18
L18:
	;
	v93 = v4
	goto L21
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v70)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	F_set_using_names(m, l0, v73, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = v78
	v80 = v77
	goto L18
L21:
	;
	v99 = int32(0)
	if v80 == v99 {
		v109 = v99
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v79 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v103 <= v93 {
		v109 = int32(0)
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v109 = v105 + v93<<(uint(int32(2))%32)
	goto L23
L26:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if int32(0) < v699 {
		goto L156
	} else {
		goto L157
	}
L27:
	;
	return
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if base.B2i32(v109 == int32(0))|base.B2i32(v114 <= v93) != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v117 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v120 = int32(2)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117+v93<<(uint(v120)%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	if v125 == v120 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123)+32))
	v131 = int32(2)
	v134 = int32(4)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v129+v130<<(uint(v131)%32)-v134)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v123)+28))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v129+v137<<(uint(v131)%32)-v134)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	if v145 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	F_set_relation_column_names(m, l0, v124, v123)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L155
	}
L34:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v148 = v146
	goto L36
L35:
	;
	v148 = int32(0)
	goto L36
L36:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v149 < v148 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v151 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	F_build_colinfo_names_hash(m, v123)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L46
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v164
	goto L39
L41:
	;
	v156 = F_palloc0(m, v148<<(uint(int32(2))%32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v158 = int32(2)
	v162 = F_repalloc0(m, v151, v149<<(uint(v158)%32), v148<<(uint(v158)%32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v164 = v156
	goto L40
L45:
	;
	v164 = v162
	goto L40
L46:
	;
	v171 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v123)+44))
	if v173 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v175 = v174
	goto L49
L48:
	;
	v175 = v171
	goto L49
L49:
	;
	if v175 < v148 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v179 = v175
	v184 = v171
	goto L53
L51:
	;
	v314 = v173
	v317 = v171
	goto L52
L52:
	;
	v328 = int32(0)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v314 != 0 {
		goto L91
	} else {
		goto L92
	}
L53:
	;
	v196 = v179 << (uint(int32(2)) % 32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196+v197)))
	if int32(0) < v199 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v123)+44))
	v314 = v309
	v317 = v305
	goto L52
L55:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v229 = v228 + v196
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if v230 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v227 = v202 + v199<<(uint(int32(2))%32) - int32(4)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208+v196)))
	if int32(0) < v210 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v227 = v213 + v210<<(uint(int32(2))%32) - int32(4)
	goto L55
L60:
	;
	goto L61
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221+v196)))
	v227 = v223 + int32(4)
	goto L55
L62:
	;
	v307 = v179 + int32(1)
	if v307 != v148 {
		v179 = v307
		v184 = v305
		goto L53
	} else {
		goto L90
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = int32(0)
	v305 = v184
	goto L62
L64:
	;
	goto L65
L65:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v235 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v230
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v239 == int32(0) {
		v305 = v184
		goto L62
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v246 != 0 {
		v270 = v246
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v244 = F_hash_search(m, v239, v230, int32(1), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v305 = v184
	goto L62
L71:
	;
	v271 = int32(1)
	if v184&v271 != 0 {
		v305 = v271
		goto L62
	} else {
		goto L82
	}
L72:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	if v247 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v250 = v248
	goto L75
L74:
	;
	v250 = int32(0)
	goto L75
L75:
	;
	if v179 < v250 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v252+v196)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v256 = v255
	goto L78
L77:
	;
	v256 = v230
	goto L78
L78:
	;
	v257 = F_make_colname_unique(m, v256, l0, v123)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v259+v196))) = v257
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v262 == int32(0) {
		v270 = v257
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v267 = F_hash_search(m, v262, v257, int32(1), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v270 = v257
	goto L71
L82:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if base.B2i32(v276 == int32(0))|base.B2i32(v276 != v279) != 0 {
		v297 = v276
		v298 = v279
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v305 = base.B2i32(v297-v298 != int32(0))
	goto L62
L84:
	;
	goto L83
L85:
	;
	v282 = v270
	v283 = v230
	goto L86
L86:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	if v287 == int32(0) {
		v297 = v287
		v298 = v286
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v297 = v287
	v298 = v286
	goto L84
L88:
	;
	v290 = int32(1)
	if v287 == v286 {
		v282 = v282 + v290
		v283 = v283 + v290
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	goto L54
L91:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	v334 = v332
	goto L93
L92:
	;
	v334 = int32(0)
	goto L93
L93:
	;
	v335 = v329 + v330 - v334
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = v335
	v339 = F_palloc0(m, v335<<(uint(int32(2))%32))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v339
	v342 = F_palloc0(m, v335)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+16)) = v342
	if v148 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v428 <= int32(0) {
		goto L113
	} else {
		goto L114
	}
L97:
	;
	v347 = int32(0)
	v411 = v347
	v421 = v347
	v426 = v328
	goto L96
L98:
	;
	goto L99
L99:
	;
	v349 = int32(0)
	v352 = v349
	v362 = v349
	v367 = v328
	goto L100
L100:
	;
	v370 = v352 << (uint(int32(2)) % 32)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v370+v371)))
	if v373 == int32(0) {
		v411 = v352
		v421 = v362
		v426 = v367
		goto L96
	} else {
		goto L102
	}
L101:
	;
	v411 = v148
	v421 = v398
	v426 = v406
	goto L96
L102:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376+v370)))
	if v378 == int32(0) {
		v411 = v352
		v421 = v362
		v426 = v367
		goto L96
	} else {
		goto L103
	}
L103:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v383+v370)))
	*(*int32)(unsafe.Add(mBase, uint32(v381+v370))) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v389 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v387+v352))) = uint8(v389)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391+v370)))
	if v389 < v393 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v396 = F_bms_add_member(m, v362, v393)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	v398 = v362
	goto L106
L106:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399+v370)))
	if int32(0) < v401 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v398 = v396
	goto L106
L108:
	;
	v404 = F_bms_add_member(m, v367, v401)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	v406 = v367
	goto L110
L110:
	;
	v408 = v352 + int32(1)
	if v408 != v148 {
		v352 = v408
		v362 = v398
		v367 = v406
		goto L100
	} else {
		goto L112
	}
L111:
	;
	v406 = v404
	goto L110
L112:
	;
	goto L101
L113:
	;
	v682 = v411
	v683 = v411
	v688 = v317
	goto L26
L114:
	;
	goto L115
L115:
	;
	v431 = int32(0)
	v434 = v411
	v435 = v411
	v437 = v431
	v439 = v431
	v440 = v317
	goto L116
L116:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451+v439))))
	if v453 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v682 = v655
	v683 = v656
	v688 = v661
	goto L26
L118:
	;
	v673 = v439 + int32(1)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v673 < v674 {
		v434 = v655
		v435 = v656
		v437 = v658
		v439 = v673
		v440 = v661
		goto L116
	} else {
		goto L154
	}
L119:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648+v439))))
	*(*uint8)(unsafe.Add(mBase, uint32(v646+v434))) = uint8(v650)
	v655 = v434 + int32(1)
	v656 = v630
	v658 = v632
	v661 = v635
	goto L118
L120:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v456 <= v437 {
		v488 = v437
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563+v439<<(uint(int32(2))%32))))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v568 != 0 {
		goto L138
	} else {
		goto L139
	}
L123:
	;
	v503 = v488 + int32(1)
	v504 = F_bms_is_member(m, v503, v421)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L129
	}
L124:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v463 = v437
	goto L125
L125:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v458+v463<<(uint(int32(2))%32))))
	if v480 != 0 {
		v488 = v463
		goto L123
	} else {
		goto L127
	}
L126:
	;
	v488 = v456
	goto L123
L127:
	;
	v482 = v463 + int32(1)
	if v482 != v456 {
		v463 = v482
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	if v504 != 0 {
		v655 = v434
		v656 = v435
		v658 = v503
		v661 = v440
		goto L118
	} else {
		goto L130
	}
L130:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v507 <= v435 {
		v536 = v435
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v553 = int32(2)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v506+v536<<(uint(v553)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v552+v434<<(uint(v553)%32)))) = v559
	v630 = v536 + int32(1)
	v632 = v503
	v635 = v440
	goto L119
L132:
	;
	v511 = v435
	goto L133
L133:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v506+v511<<(uint(int32(2))%32))))
	if v530 != 0 {
		v536 = v511
		goto L131
	} else {
		goto L135
	}
L134:
	;
	v536 = v507
	goto L131
L135:
	;
	v532 = v511 + int32(1)
	if v532 != v507 {
		v511 = v532
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v618 != 0 {
		goto L150
	} else {
		goto L151
	}
L138:
	;
	v569 = F_make_colname_unique(m, v567, l0, v123)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v610+v434<<(uint(int32(2))%32)))) = v567
	v617 = v440
	goto L137
L141:
	;
	v572 = v434 << (uint(int32(2)) % 32)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v572+v573))) = v569
	v576 = int32(1)
	if v440&v576 != 0 {
		v617 = v576
		goto L137
	} else {
		goto L142
	}
L142:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v579+v572)))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581))))
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	if base.B2i32(v584 == int32(0))|base.B2i32(v584 != v587) != 0 {
		v605 = v584
		v606 = v587
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v617 = base.B2i32(v605-v606 != int32(0))
	goto L137
L144:
	;
	goto L143
L145:
	;
	v590 = v581
	v591 = v567
	goto L146
L146:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+1)))
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+1)))
	if v595 == int32(0) {
		v605 = v595
		v606 = v594
		goto L144
	} else {
		goto L148
	}
L147:
	;
	v605 = v595
	v606 = v594
	goto L144
L148:
	;
	v598 = int32(1)
	if v595 == v594 {
		v590 = v590 + v598
		v591 = v591 + v598
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v619+v434<<(uint(int32(2))%32))))
	v626 = F_hash_search(m, v618, v623, int32(1), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v630 = v435
	v632 = v437
	v635 = v617
	goto L119
L153:
	;
	goto L152
L154:
	;
	goto L117
L155:
	;
	v93 = v93 + int32(1)
	goto L21
L156:
	;
	v702 = int32(0)
	v705 = v682
	v706 = v683
	v708 = v702
	v711 = v688
	v712 = v702
	goto L159
L157:
	;
	v954 = v688
	goto L158
L158:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v965 != 0 {
		goto L198
	} else {
		goto L199
	}
L159:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722+v712))))
	if v724 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v954 = v932
	goto L158
L161:
	;
	v944 = v712 + int32(1)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if v944 < v945 {
		v705 = v926
		v706 = v927
		v708 = v929
		v711 = v932
		v712 = v944
		goto L159
	} else {
		goto L197
	}
L162:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919+v712))))
	*(*uint8)(unsafe.Add(mBase, uint32(v917+v705))) = uint8(v921)
	v926 = v705 + int32(1)
	v927 = v901
	v929 = v903
	v932 = v906
	goto L161
L163:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v727 <= v708 {
		v759 = v708
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L165
L165:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v834+v712<<(uint(int32(2))%32))))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v839 != 0 {
		goto L181
	} else {
		goto L182
	}
L166:
	;
	v774 = v759 + int32(1)
	v775 = F_bms_is_member(m, v774, v426)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L172
	}
L167:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v734 = v708
	goto L168
L168:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v729+v734<<(uint(int32(2))%32))))
	if v751 != 0 {
		v759 = v734
		goto L166
	} else {
		goto L170
	}
L169:
	;
	v759 = v727
	goto L166
L170:
	;
	v753 = v734 + int32(1)
	if v753 != v727 {
		v734 = v753
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	if v775 != 0 {
		v926 = v705
		v927 = v706
		v929 = v774
		v932 = v711
		goto L161
	} else {
		goto L173
	}
L173:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v778 <= v706 {
		v807 = v706
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v824 = int32(2)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v777+v807<<(uint(v824)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v823+v705<<(uint(v824)%32)))) = v830
	v901 = v807 + int32(1)
	v903 = v774
	v906 = v711
	goto L162
L175:
	;
	v782 = v706
	goto L176
L176:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v777+v782<<(uint(int32(2))%32))))
	if v801 != 0 {
		v807 = v782
		goto L174
	} else {
		goto L178
	}
L177:
	;
	v807 = v778
	goto L174
L178:
	;
	v803 = v782 + int32(1)
	if v803 != v778 {
		v782 = v803
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v889 != 0 {
		goto L193
	} else {
		goto L194
	}
L181:
	;
	v840 = F_make_colname_unique(m, v838, l0, v123)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v881+v705<<(uint(int32(2))%32)))) = v838
	v888 = v711
	goto L180
L184:
	;
	v843 = v705 << (uint(int32(2)) % 32)
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v843+v844))) = v840
	v847 = int32(1)
	if v711&v847 != 0 {
		v888 = v847
		goto L180
	} else {
		goto L185
	}
L185:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v850+v843)))
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852))))
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838))))
	if base.B2i32(v855 == int32(0))|base.B2i32(v855 != v858) != 0 {
		v876 = v855
		v877 = v858
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v888 = base.B2i32(v876-v877 != int32(0))
	goto L180
L187:
	;
	goto L186
L188:
	;
	v861 = v852
	v862 = v838
	goto L189
L189:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+1)))
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+1)))
	if v866 == int32(0) {
		v876 = v866
		v877 = v865
		goto L187
	} else {
		goto L191
	}
L190:
	;
	v876 = v866
	v877 = v865
	goto L187
L191:
	;
	v869 = int32(1)
	if v866 == v865 {
		v861 = v861 + v869
		v862 = v862 + v869
		goto L189
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v890+v705<<(uint(int32(2))%32))))
	v897 = F_hash_search(m, v889, v894, int32(1), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v901 = v706
	v903 = v708
	v906 = v888
	goto L162
L196:
	;
	goto L195
L197:
	;
	goto L160
L198:
	;
	F_hash_destroy(m, v965)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v973 = base.B2i32(v970 != int32(0)) & v954
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+20)) = uint8(v973)
	v93 = v93 + int32(1)
	goto L21
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+48)) = int32(0)
	goto L200
}
func F_set_pathtarget_cost_width(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int64
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v90 int32
	_ = v90
	v3 = int32(0)
	v8 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v8
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == v3 {
		v90 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v90
	m.G0 = v13 + int32(32)
	return l1
L2:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v24 <= v23 {
		v90 = v23
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = v13 + int32(16)
	v34 = v3
	v36 = v8
	goto L4
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
	v44 = F_get_expr_width(m, l0, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v75 = int64(1073741823)
	if v75 <= v70 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	return int32(0)
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v49 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
	v53 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v53
	v59 = F_cost_qual_eval_walker(m, v43, v13+int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v70 = v36 + base.I64_extend_i32_s(v44)
	v72 = v34 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v72 < v73 {
		v34 = v72
		v36 = v70
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v61 = *(*float64)(unsafe.Add(mBase, uint32(v13)+24))
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
	v63 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v62, v63)
	v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = base.F64_add(v61, v66)
	goto L10
L12:
	;
	goto L5
L13:
	;
	v78 = v75
	goto L15
L14:
	;
	v78 = v70
	goto L15
L15:
	;
	v90 = base.I32_wrap_i64(v78)
	goto L1
}
func F_set_rtable_names(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v299 int32
	_ = v299
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = int64(292057776192)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_set_rtable_names[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v30 = F_hash_create(m, int32(_a_F_set_rtable_names_0), v26, v15+int32(32), int32(1048))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v15 + int32(80)
	return
L4:
	;
	return
L5:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v117 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v34 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = v4
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v43<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v102 = v43 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v102 < v103 {
		v43 = v102
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v58 <= v57 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v66 = v57
	goto L14
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v66<<(uint(int32(2))%32))))
	if v77 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L11
L16:
	;
	v81 = F_hash_search(m, v30, v77, int32(1), v15+int32(31))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v86 = v66 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v86 < v87 {
		v66 = v86
		goto L14
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+64)) = int32(0)
	goto L18
L20:
	;
	goto L15
L21:
	;
	goto L10
L22:
	;
	F_hash_destroy(m, v30)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L72
	}
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v120 <= int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v133 = int32(1)
	v134 = v4
	goto L25
L25:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v134<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_set_rtable_names[1]))
	if v142 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L22
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l2 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L29
L31:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v277 = F_lappend(m, v276, v265)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L70
	}
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v150 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v147 = F_bms_is_member(m, v133, l2)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v147 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v265 = int32(0)
	goto L31
L36:
	;
	if v160 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v160 = v151
	goto L36
L38:
	;
	goto L39
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	switch v153 {
	case 0:
		goto L41
	default:
		goto L40
	case 2:
		v265 = int32(0)
		goto L31
	}
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v140)+8))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v160 = v158
	goto L36
L41:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	v155 = F_get_rel_name(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v160 = v155
	goto L36
L43:
	;
	v265 = int32(0)
	goto L31
L44:
	;
	goto L45
L45:
	;
	v167 = F_hash_search(m, v30, v160, int32(1), v15+int32(31))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v169 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v172 = F_strlen(m, v160)
	mBase = m.M
	v175 = F_palloc(m, v172+int32(16))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	v250 = v160
	v261 = v167
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+64)) = int32(0)
	v265 = v250
	goto L31
L50:
	;
	v182 = v172
	goto L51
L51:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v167)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+64)) = v189 + int32(1)
	if v182 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v250 = v175
	v261 = v246
	goto L49
L53:
	;
	base.MemoryCopy(m, v175, v160, v182)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v167)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v194
	v200 = F_pg_sprintf(m, v175+v182, int32(_a_F_set_rtable_names_1), v15+int32(16))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v202 = F_strlen(m, v175)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v202) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v210 = v182
	goto L60
L58:
	;
	v236 = v182
	goto L59
L59:
	;
	v246 = F_hash_search(m, v30, v175, int32(1), v15+int32(31))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L68
	}
L60:
	;
	v219 = F_pg_mbcliplen(m, v160, v210, v210-int32(1))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v236 = v219
	goto L59
L62:
	;
	if v219 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	base.MemoryCopy(m, v175, v160, v219)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v167)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v222
	v226 = F_pg_sprintf(m, v175+v219, int32(_a_F_set_rtable_names_1), v15)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v228 = F_strlen(m, v175)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v228) {
		v210 = v219
		goto L60
	} else {
		goto L67
	}
L67:
	;
	goto L61
L68:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v248 != 0 {
		v182 = v236
		goto L51
	} else {
		goto L69
	}
L69:
	;
	goto L52
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v277
	v280 = int32(1)
	v283 = v134 + v280
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v283 < v284 {
		v133 = v133 + v280
		v134 = v283
		goto L25
	} else {
		goto L71
	}
L71:
	;
	goto L26
L72:
	;
	goto L3
}
func F_set_using_names(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
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
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
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
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v24 = l1
	v25 = l2
	goto L1
L1:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v41 != int32(64) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L13
	} else {
		goto L137
	}
L3:
	;
	goto L2
L4:
	;
	switch v41 - int32(63) {
	case 0:
		goto L7
	default:
		goto L3
	case 2:
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v106 = int32(4)
	v107 = v103<<(uint(int32(2))%32) - v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v109)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113+v107)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	switch v119 - int32(63) {
	case 0:
		v140 = v106
		goto L16
	case 1:
		goto L17
	default:
		goto L18
	}
L7:
	;
	m.G0 = v21 + int32(48)
	return
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v46 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v57 = int32(0)
	goto L11
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v57<<(uint(int32(2))%32))))
	F_set_using_names(m, l0, v75, v25)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L7
L13:
	;
	return
L14:
	;
	v79 = v57 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v79 < v80 {
		v57 = v79
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140+v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	switch v145 - int32(63) {
	case 0:
		v166 = v106
		goto L22
	case 1:
		goto L23
	default:
		goto L24
	}
L17:
	;
	v140 = int32(36)
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v127
	F_errmsg_internal(m, int32(_a_F_set_using_names_0), v21+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_set_using_names_1), int32(_a_F_set_using_names_2), int32(_a_F_set_using_names_3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
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
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144+v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+32)) = v168
	v170 = int32(0)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v115)+52))
	if v172 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v166 = int32(36)
	goto L22
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v153
	F_errmsg_internal(m, int32(_a_F_set_using_names_0), v21+int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_set_using_names_1), int32(_a_F_set_using_names_4), int32(_a_F_set_using_names_3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
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
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v176 = v173 << (uint(int32(2)) % 32)
	goto L30
L29:
	;
	v176 = v170
	goto L30
L30:
	;
	v177 = F_palloc0(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v177
	v180 = F_palloc0(m, v176)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+40)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v115)+56))
	if v183 == int32(0) {
		v223 = v170
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v115)+60))
	if v237 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v186 <= int32(0) {
		v223 = v170
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v193 = v170
	goto L36
L36:
	;
	v208 = v193 << (uint(int32(2)) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211+v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v208+v209))) = v213
	v216 = v193 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v216 < v217 {
		v193 = v216
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v223 = v216
	goto L33
L38:
	;
	goto L37
L39:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	v301 = int32(2)
	v304 = int32(4)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v299+v300<<(uint(v301)%32)-v304)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v299+v307<<(uint(v301)%32)-v304)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v111)+40))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v316 != 0 {
		goto L48
	} else {
		goto L49
	}
L40:
	;
	v240 = int32(0)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v241 <= v240 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v247 = v240
	v248 = v223
	goto L42
L42:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v111)+40))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v115)+48))
	v264 = base.B2i32(v263 <= v247)
	if v263 <= v247 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L39
L44:
	;
	v265 = v248
	goto L46
L45:
	;
	v265 = v247
	goto L46
L46:
	;
	v266 = int32(2)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269+v247<<(uint(v266)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v262+v265<<(uint(v266)%32)))) = v273
	v277 = v247 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v277 < v278 {
		v247 = v277
		v248 = v248 + v264
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v437 == int32(0) {
		v609 = v25
		goto L78
	} else {
		goto L79
	}
L49:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v317 <= int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v325 = int32(0)
	goto L51
L51:
	;
	v340 = v325 << (uint(int32(2)) % 32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v340+v341)))
	if v343 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L48
L53:
	;
	v416 = v325 + int32(1)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v416 < v417 {
		v325 = v416
		goto L51
	} else {
		goto L77
	}
L54:
	;
	v346 = v340 + v315
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	if int32(0) < v347 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	if v351 < v347 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v379 = v340 + v314
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	if v380 <= int32(0) {
		goto L53
	} else {
		goto L67
	}
L58:
	;
	if v350 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v369 = v350
	v370 = v347
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370<<(uint(int32(2))%32)+v369-int32(4)))) = v343
	goto L57
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = v365
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v369 = v365
	v370 = v368
	goto L60
L62:
	;
	v357 = F_palloc0(m, v347<<(uint(int32(2))%32))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L13
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v359 = int32(2)
	v363 = F_repalloc0(m, v350, v351<<(uint(v359)%32), v347<<(uint(v359)%32))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L13
	} else {
		goto L66
	}
L65:
	;
	v365 = v357
	goto L61
L66:
	;
	v365 = v363
	goto L61
L67:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v384 < v380 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v383 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v402 = v383
	v403 = v380
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403<<(uint(int32(2))%32)+v402-int32(4)))) = v343
	goto L53
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v398
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v402 = v398
	v403 = v401
	goto L70
L72:
	;
	v390 = F_palloc0(m, v380<<(uint(int32(2))%32))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L13
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v392 = int32(2)
	v396 = F_repalloc0(m, v383, v384<<(uint(v392)%32), v380<<(uint(v392)%32))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L13
	} else {
		goto L76
	}
L75:
	;
	v398 = v390
	goto L71
L76:
	;
	v398 = v396
	goto L71
L77:
	;
	goto L52
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+24)) = v609
	*(*int32)(unsafe.Add(mBase, uint32(v306)+24)) = v609
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	F_set_using_names(m, l0, v627, v609)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L13
	} else {
		goto L136
	}
L79:
	;
	v440 = F_list_copy(m, v25)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v442 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v445 = v443
	goto L83
L82:
	;
	v445 = int32(0)
	goto L83
L83:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v446 < v445 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v448 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v467 = v442
	goto L86
L86:
	;
	if v467 == int32(0) {
		v609 = v440
		goto L78
	} else {
		goto L93
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v461
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v467 = v464
	goto L86
L88:
	;
	v453 = F_palloc0(m, v445<<(uint(int32(2))%32))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v455 = int32(2)
	v459 = F_repalloc0(m, v448, v446<<(uint(v455)%32), v445<<(uint(v455)%32))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L13
	} else {
		goto L92
	}
L91:
	;
	v461 = v453
	goto L87
L92:
	;
	v461 = v459
	goto L87
L93:
	;
	v470 = int32(0)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	if v471 <= v470 {
		v609 = v440
		goto L78
	} else {
		goto L94
	}
L94:
	;
	v476 = v440
	v479 = v470
	goto L95
L95:
	;
	v493 = v479 << (uint(int32(2)) % 32)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v493+v494)))
	if v496 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v609 = v535
	goto L78
L97:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v467)+12))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v499+v493)))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)+4))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v503 == int32(0) {
		v516 = v502
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v530 = v496
	goto L99
L99:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	v532 = F_lappend(m, v531, v530)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L13
	} else {
		goto L109
	}
L100:
	;
	v517 = F_make_colname_unique(m, v516, l0, v111)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L13
	} else {
		goto L104
	}
L101:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v503)+8))
	if v506 == int32(0) {
		v516 = v502
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	if v509 <= v479 {
		v516 = v502
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v511+v493)))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	v516 = v514
	goto L100
L104:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v519 == int32(1) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v523 = F_lappend(m, v522, v517)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L13
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v526+v493))) = v517
	v530 = v517
	goto L99
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v523
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = v532
	v535 = F_lappend(m, v476, v530)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	v537 = v493 + v315
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	if int32(0) < v538 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	if v542 < v538 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	v570 = v493 + v314
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	if int32(0) < v571 {
		goto L123
	} else {
		goto L124
	}
L114:
	;
	if v541 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v560 = v541
	v561 = v538
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561<<(uint(int32(2))%32)+v560-int32(4)))) = v530
	goto L113
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v538
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = v556
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	v560 = v556
	v561 = v559
	goto L116
L118:
	;
	v548 = F_palloc0(m, v538<<(uint(int32(2))%32))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L13
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v550 = int32(2)
	v554 = F_repalloc0(m, v541, v542<<(uint(v550)%32), v538<<(uint(v550)%32))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L13
	} else {
		goto L122
	}
L121:
	;
	v556 = v548
	goto L117
L122:
	;
	v556 = v554
	goto L117
L123:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v575 < v571 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	v604 = v479 + int32(1)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	if v604 < v605 {
		v476 = v535
		v479 = v604
		goto L95
	} else {
		goto L135
	}
L126:
	;
	if v574 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v593 = v574
	v594 = v571
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v594<<(uint(int32(2))%32)+v593-int32(4)))) = v530
	goto L125
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v589
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	v593 = v589
	v594 = v592
	goto L128
L130:
	;
	v581 = F_palloc0(m, v571<<(uint(int32(2))%32))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L13
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v583 = int32(2)
	v587 = F_repalloc0(m, v574, v575<<(uint(v583)%32), v571<<(uint(v583)%32))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L13
	} else {
		goto L134
	}
L133:
	;
	v589 = v581
	goto L129
L134:
	;
	v589 = v587
	goto L129
L135:
	;
	goto L96
L136:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v24 = v630
	v25 = v609
	goto L1
L137:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v635
	F_errmsg_internal(m, int32(_a_F_set_using_names_5), v21)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L13
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_set_using_names_1), int32(_a_F_set_using_names_6), int32(_a_F_set_using_names_7))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setop_load_group(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v4)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, v8, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L9
	}
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v9&int32(2) == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	m.T0[v16].(func(*base.Module, int32))(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	goto L1
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_ExecStoreMinimalTuple(m, v24, v26, int32(1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v35 = int64(1)
	goto L11
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ExecReScan(m, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = m.T0[v40].(func(*base.Module, int32) int32)(m, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v41
	if v41 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v46&int32(2) != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = F_setop_compare_slots(m, v49, v41, l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if v50 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = v52 + int64(1)
	goto L11
}
func F_setval_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	F_do_setval(m, v3, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_Int64GetDatum(m, v5)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_shdepDropDependency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v60 int32
	_ = v60
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	v12 = m.G0
	v14 = v12 - int32(192)
	m.G0 = v14
	v18 = int32(1)
	if l1 <= int32(3591) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v90 = int32(3)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_shdepDropDependency[0]))
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	goto L1
L3:
	;
	v89 = int32(0)
	goto L2
L4:
	;
	if base.B2i32(base.Ui32(l1-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l1-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v89 = v18
		goto L2
	} else {
		goto L25
	}
L5:
	;
	if l1 <= int32(2670) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if l1 <= int32(_a_F_shdepDropDependency_0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	switch l1 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v89 = v18
		goto L2
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L3
	default:
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v30 = l1 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v30))|base.B2i32(int32(1)<<(uint(v30)%32)&int32(226492515) == int32(0)) != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(2396)) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v89 = v18
	goto L2
L13:
	;
	v89 = v18
	goto L2
L14:
	;
	if base.Ui32(l1-int32(3592)) < base.Ui32(int32(2)) {
		v89 = v18
		goto L2
	} else {
		goto L23
	}
L15:
	;
	v43 = l1 - int32(_a_F_shdepDropDependency_1)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v43))|base.B2i32(int32(1)<<(uint(v43)%32)&int32(963) == int32(0)) != 0 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	switch l1 - int32(_a_F_shdepDropDependency_2) {
	case 0, 1, 2, 3, 4, 59, 60:
		v89 = v18
		goto L2
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L3
	default:
		goto L19
	}
L18:
	;
	v89 = v18
	goto L2
L19:
	;
	if base.Ui32(l1-int32(_a_F_shdepDropDependency_3)) < base.Ui32(int32(3)) {
		v89 = v18
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v60 = l1 - int32(_a_F_shdepDropDependency_4)
	if base.Ui32(int32(15)) < base.Ui32(v60) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if int32(1)<<(uint(v60)%32)&int32(_a_F_shdepDropDependency_5) != 0 {
		v89 = v18
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L3
L23:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(4060)) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v89 = v18
	goto L2
L25:
	;
	goto L3
L26:
	;
	v97 = int32(0)
	goto L28
L27:
	;
	v97 = v96
	goto L28
L28:
	;
	F_ScanKeyInit(m, v14, int32(1), v90, int32(184), v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return
L30:
	;
	F_ScanKeyInit(m, v14+int32(48), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v109 = int32(3)
	F_ScanKeyInit(m, v14+int32(96), v109, v109, int32(184), l2)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if l4 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v116 = int32(4)
	F_ScanKeyInit(m, v14+int32(144), v116, int32(3), int32(65), l3)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L36
	}
L34:
	;
	v124 = v90
	goto L35
L35:
	;
	v128 = F_systable_beginscan(m, l0, int32(1232), int32(1), int32(0), v124, v14)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L29
	} else {
		goto L37
	}
L36:
	;
	v124 = v116
	goto L35
L37:
	;
	v130 = F_systable_getnext(m, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	if v130 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v133 = v130
	goto L42
L40:
	;
	goto L41
L41:
	;
	F_systable_endscan(m, v128)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L29
	} else {
		goto L60
	}
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+22)))
	v145 = v143 + v144
	if l5 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	v156 = F_systable_getnext(m, v128)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L29
	} else {
		goto L58
	}
L45:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
	if v146 != l5 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l6 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+20))
	if v148 != l6 {
		goto L44
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if l7 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v145)+24)))
	if l7 != v150 {
		goto L44
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_simple_heap_delete(m, l0, v133+int32(4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L29
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	goto L44
L58:
	;
	if v156 != 0 {
		v133 = v156
		goto L42
	} else {
		goto L59
	}
L59:
	;
	goto L43
L60:
	;
	m.G0 = v14 + int32(192)
	return
}
func F_shell_archive_init(m *base.Module) int32 {
	return int32(_a_F_shell_archive_init_0)
}
func F_shell_archive_shutdown(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v4 = F_errstart(m, int32(14), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if v4 != 0 {
			F_errmsg_internal(m, int32(_a_F_shell_archive_shutdown_0), int32(0))
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_shell_archive_shutdown_1), int32(141), int32(_a_F_shell_archive_shutdown_2))
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	}
}
func F_shift_jis_2004_to_euc_jis_2004(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
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
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v280 int32
	_ = v280
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v15, v16, v17, int32(41), int32(5))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v17 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L109
	}
L4:
	;
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v269)
	return v260 - v14
L5:
	;
	v260 = v14
	v262 = v13
	goto L4
L6:
	;
	goto L7
L7:
	;
	v26 = v17
	v28 = v14
	v30 = v13
	goto L8
L8:
	;
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	if int32(0) <= v37 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v260 = v249
	v262 = v251
	goto L4
L10:
	;
	if int32(0) < v255 {
		v26 = v255
		v28 = v249
		v30 = v251
		goto L8
	} else {
		goto L108
	}
L11:
	;
	if v37 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v50 = F_pg_encoding_verifymbchar(m, int32(41), v28, v26)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	if v12 != 0 {
		v260 = v28
		v262 = v30
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v37)
	v43 = int32(1)
	v249 = v28 + v43
	v251 = v30 + v43
	v255 = v26 - v43
	goto L10
L17:
	;
	goto L3
L18:
	;
	if base.Ui32(v26) < base.Ui32(v50) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v12 != 0 {
		v260 = v28
		v262 = v30
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.B2i32(v50 != int32(1))|base.B2i32(base.Ui32(int32(62)) < base.Ui32((v37+int32(95))&int32(255))) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L3
L23:
	;
	v249 = v28 + v50
	v251 = v243
	v255 = v26 - v50
	goto L10
L24:
	;
	v243 = v236 + int32(2)
	goto L23
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)) = uint8(v37)
	v65 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v65)
	v236 = v30
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v50 != int32(2) {
		v243 = v30
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v70 = v37 & int32(255)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v72 = base.I32_extend8_s(v71)
	if base.B2i32(v37 == int32(-128))|base.B2i32(base.Ui32(int32(-97)) < base.Ui32(v37)) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v228 = int32(96)
	v229 = v226 - v228
	*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)) = uint8(v229)
	v232 = v225 - v228
	*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v232)
	v236 = v227
	goto L24
L30:
	;
	v225 = v70<<(uint(int32(1))%32) + v218 - int32(256)
	v226 = v219
	v227 = v30
	goto L29
L31:
	;
	v81 = v71 + int32(-64)
	if base.Ui32(v81) <= base.Ui32(int32(62)) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v37&int32(-16) == int32(-32) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v218 = int32(-1)
	v219 = v71 - int32(63)
	goto L30
L35:
	;
	goto L36
L36:
	;
	if v72 < int32(-97) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v12 != 0 {
		v260 = v28
		v262 = v30
		goto L4
	} else {
		goto L43
	}
L38:
	;
	v95 = v81
	v96 = int32(-1)
	goto L40
L39:
	;
	if int32(-4) < v72 {
		goto L37
	} else {
		goto L41
	}
L40:
	;
	if int32(0) <= v95 {
		v218 = v96
		v219 = v95
		goto L30
	} else {
		goto L42
	}
L41:
	;
	v95 = v71 - int32(158)
	v96 = int32(0)
	goto L40
L42:
	;
	goto L37
L43:
	;
	goto L3
L44:
	;
	v225 = v70<<(uint(int32(1))%32) + v211 - int32(384)
	v226 = v212
	v227 = v30
	goto L29
L45:
	;
	v106 = v71 + int32(-64)
	if base.Ui32(v106) <= base.Ui32(int32(62)) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v37&int32(-4) == int32(-16) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v211 = int32(-1)
	v212 = v71 - int32(63)
	goto L44
L49:
	;
	goto L50
L50:
	;
	if v72 < int32(-97) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v12 != 0 {
		v260 = v28
		v262 = v30
		goto L4
	} else {
		goto L57
	}
L52:
	;
	v120 = v106
	v121 = int32(-1)
	goto L54
L53:
	;
	if int32(-4) < v72 {
		goto L51
	} else {
		goto L55
	}
L54:
	;
	if int32(0) <= v120 {
		v211 = v121
		v212 = v120
		goto L44
	} else {
		goto L56
	}
L55:
	;
	v120 = v71 - int32(158)
	v121 = int32(0)
	goto L54
L56:
	;
	goto L51
L57:
	;
	goto L3
L58:
	;
	v207 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v207)
	v225 = v206
	v226 = v205
	v227 = v30 + int32(1)
	goto L29
L59:
	;
	switch v70 - int32(240) {
	case 0:
		goto L92
	case 1:
		goto L95
	case 2:
		goto L94
	default:
		goto L93
	}
L60:
	;
	v131 = v71 + int32(-64)
	if base.Ui32(v131) <= base.Ui32(int32(62)) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if base.Ui32((v37+int32(12))&int32(255)) <= base.Ui32(int32(8)) {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	v188 = int32(1)
	v189 = v71 - int32(63)
	goto L59
L64:
	;
	goto L65
L65:
	;
	if v72 < int32(-97) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v12 != 0 {
		v260 = v28
		v262 = v30
		goto L4
	} else {
		goto L72
	}
L67:
	;
	v145 = v131
	v146 = int32(1)
	goto L69
L68:
	;
	if int32(-4) < v72 {
		goto L66
	} else {
		goto L70
	}
L69:
	;
	if int32(0) <= v145 {
		v188 = v146
		v189 = v145
		goto L59
	} else {
		goto L71
	}
L70:
	;
	v145 = v71 - int32(158)
	v146 = int32(0)
	goto L69
L71:
	;
	goto L66
L72:
	;
	goto L3
L73:
	;
	if v37 == int32(-12) {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	v158 = v71 + int32(-64)
	if base.Ui32(v158) <= base.Ui32(int32(62)) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v12 != 0 {
		v260 = v28
		v262 = v30
		goto L4
	} else {
		goto L87
	}
L77:
	;
	v178 = int32(1)
	v179 = v71 - int32(63)
	goto L73
L78:
	;
	goto L79
L79:
	;
	if v72 < int32(-97) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v12 != 0 {
		v260 = v28
		v262 = v30
		goto L4
	} else {
		goto L86
	}
L81:
	;
	v172 = v158
	v173 = int32(1)
	goto L83
L82:
	;
	if int32(-4) < v72 {
		goto L80
	} else {
		goto L84
	}
L83:
	;
	if int32(0) <= v172 {
		v178 = v173
		v179 = v172
		goto L73
	} else {
		goto L85
	}
L84:
	;
	v172 = v71 - int32(158)
	v173 = int32(0)
	goto L83
L85:
	;
	goto L80
L86:
	;
	goto L3
L87:
	;
	goto L3
L88:
	;
	if v178 != 0 {
		v205 = v179
		v206 = int32(15)
		goto L58
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v205 = v179
	v206 = v70<<(uint(int32(1))%32) - v178 - int32(410)
	goto L58
L91:
	;
	goto L90
L92:
	;
	if v188 != 0 {
		goto L105
	} else {
		goto L106
	}
L93:
	;
	if v188 != 0 {
		goto L102
	} else {
		goto L103
	}
L94:
	;
	if v188 != 0 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	if v188 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v194 = int32(3)
	goto L98
L97:
	;
	v194 = int32(4)
	goto L98
L98:
	;
	v205 = v189
	v206 = v194
	goto L58
L99:
	;
	v197 = int32(5)
	goto L101
L100:
	;
	v197 = int32(12)
	goto L101
L101:
	;
	v205 = v189
	v206 = v197
	goto L58
L102:
	;
	v200 = int32(13)
	goto L104
L103:
	;
	v200 = int32(14)
	goto L104
L104:
	;
	v205 = v189
	v206 = v200
	goto L58
L105:
	;
	v203 = int32(1)
	goto L107
L106:
	;
	v203 = int32(8)
	goto L107
L107:
	;
	v205 = v189
	v206 = v203
	goto L58
L108:
	;
	goto L9
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_shortest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
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
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v555 int32
	_ = v555
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v638 int32
	_ = v638
	v8 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if l5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v22 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v638
L8:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+8)))
	if v134&int32(2) != 0 {
		goto L47
	} else {
		goto L48
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = v25 + v22<<(uint(int32(3))%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 < int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = int32(0)
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+66)))
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+64)))
	if base.I32_extend16_s(v36) <= v35 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if base.B2i32(l5 == v32)|base.B2i32(v125 == int32(0)) != 0 {
		v638 = v125
		goto L7
	} else {
		goto L46
	}
L12:
	;
	v39 = l2
	goto L14
L13:
	;
	v39 = v32
	goto L14
L14:
	;
	if l2 == l3 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = v39
	goto L17
L16:
	;
	v42 = int32(0)
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 == v43 {
		v125 = v42
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v45 = v43 - v29
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v51 = int32(1)
	v53 = base.I32_div_u_s((l3-l2)>>(uint(int32(2))%32)-v51, v45)
	v55 = v53 + v51
	if base.Ui32(v36) < base.Ui32(v55) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v58 = v36
	goto L21
L21:
	;
	v63 = base.I32_div_u_s((l4-l2)>>(uint(int32(2))%32), v45)
	if base.Ui32(v63) < base.Ui32(v35) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v57 = v55
	goto L24
L23:
	;
	v57 = v36
	goto L24
L24:
	;
	v58 = v57
	goto L21
L25:
	;
	v65 = v63
	goto L27
L26:
	;
	v65 = v35
	goto L27
L27:
	;
	if v35 == int32(256) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v68 = v63
	goto L30
L29:
	;
	v68 = v65
	goto L30
L30:
	;
	if base.Ui32(v68) < base.Ui32(v58) {
		v638 = v8
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v58 == int32(0) {
		v125 = l2
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v72 = int32(2)
	v85 = int32(0)
	v87 = l2
	goto L34
L33:
	;
	if base.Ui32(v58) <= base.Ui32(v106) {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	if v85 == v68 {
		v105 = v87
		v106 = v68
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v105 = v100
	v106 = v58
	goto L33
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+420))
	v96 = m.T0[v95].(func(*base.Module, int32, int32, int32) int32)(m, v46+v29<<(uint(v72)%32), v87, v45)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(0)
L38:
	;
	if v96 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v105 = v87
	v106 = v85
	goto L33
L40:
	;
	goto L41
L41:
	;
	v100 = v45<<(uint(v72)%32) + v87
	v102 = v85 + int32(1)
	if v102 != v58 {
		v85 = v102
		v87 = v100
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L35
L43:
	;
	v109 = v105
	goto L45
L44:
	;
	v109 = int32(0)
	goto L45
L45:
	;
	v125 = v109
	goto L11
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l2
	return v125
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133)+44))
	v142 = (l3 - l2) >> (uint(int32(2)) % 32)
	if base.B2i32(v137 != int32(256))&base.B2i32(base.Ui32(v137) < base.Ui32(v142)) != 0 {
		v638 = v8
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v159 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v159 < v167 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	if (l4-l2)>>(uint(int32(2))%32) < v145 {
		v638 = v8
		goto L7
	} else {
		goto L51
	}
L51:
	;
	if base.Ui32(v142) < base.Ui32(v145) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v154 = l2 + v145<<(uint(int32(2))%32)
	goto L54
L53:
	;
	v154 = l3
	goto L54
L54:
	;
	if l5 == int32(0) {
		v638 = v154
		goto L7
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l2
	return v154
L56:
	;
	if v380 == int32(0) {
		v638 = v8
		goto L7
	} else {
		goto L93
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
	v380 = v359
	goto L56
L58:
	;
	v334 = int32(0)
	goto L90
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)))
	if v171&int32(1) != 0 {
		v326 = v170
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v175 = F_getvacant(m, l0, l1, l2, l2)
	mBase = m.M
	if v175 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v380 = int32(0)
	goto L56
L64:
	;
	goto L65
L65:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v179 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v183 = int32(0)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v222 = v215 + int32(base.Ui32(v217)>>(uint(int32(3))%32))&int32(536870908)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v223 | v224<<(uint(v217)%32)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v229 == v224 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	*(*int32)(unsafe.Add(mBase, uint32(v194+v183<<(uint(int32(2))%32)))) = int32(0)
	v201 = v183 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v201 < v202 {
		v183 = v201
		goto L69
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v308
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v319 <= int32(0) {
		v359 = v175
		goto L57
	} else {
		goto L89
	}
L73:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v308 = v232
	goto L72
L74:
	;
	goto L75
L75:
	;
	if v229 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v308 = int32(0)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v237 = v229 & int32(3)
	v238 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v229) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v244 = v238
	v247 = v238
	v253 = v159
	goto L82
L80:
	;
	v273 = v238
	v276 = v238
	goto L81
L81:
	;
	v284 = v273
	v287 = v276
	v294 = v159
	goto L86
L82:
	;
	v257 = v228 + v244<<(uint(int32(2))%32)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v265 = v258 ^ (v259 ^ (v260 ^ (v261 ^ v247)))
	v266 = int32(4)
	v267 = v244 + v266
	v269 = v253 + v266
	if v269 != v229&int32(2147483644) {
		v244 = v267
		v247 = v265
		v253 = v269
		goto L82
	} else {
		goto L84
	}
L83:
	;
	if v237 == int32(0) {
		v308 = v265
		goto L72
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	v273 = v267
	v276 = v265
	goto L81
L86:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v228+v284<<(uint(int32(2))%32))))
	v299 = v298 ^ v287
	v300 = int32(1)
	v303 = v294 + v300
	if v303 != v237 {
		v284 = v284 + v300
		v287 = v299
		v294 = v303
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v308 = v299
	goto L72
L88:
	;
	goto L87
L89:
	;
	v326 = v175
	goto L58
L90:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v345+v334<<(uint(int32(5))%32))+20)) = int32(0)
	v352 = v334 + int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v352 < v353 {
		v334 = v352
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v359 = v326
	goto L57
L92:
	;
	goto L91
L93:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v383 == l2 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v409 = F_miss(m, l0, l1, v380, base.I32_extend16_s(v407), l2, l2)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L37
	} else {
		goto L101
	}
L95:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v389 = int32(1)
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385+(v386^int32(-1))&v389<<(uint(v389)%32))+20)))
	v407 = v394
	goto L94
L96:
	;
	goto L97
L97:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l2-int32(4))))
	if base.Ui32(v397) <= base.Ui32(int32(2047)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400+v397<<(uint(int32(1))%32)))))
	v407 = v404
	goto L94
L99:
	;
	goto L100
L100:
	;
	v405 = F_pg_reg_getcolor(m, v17, v397)
	mBase = m.M
	v407 = v405
	goto L94
L101:
	;
	if v409 == int32(0) {
		v638 = v8
		goto L7
	} else {
		goto L102
	}
L102:
	;
	if l4 != v16 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v416 = int32(4)
	goto L105
L104:
	;
	v416 = int32(0)
	goto L105
L105:
	;
	if l3 != v16 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v421 = int32(4)
	goto L108
L107:
	;
	v421 = int32(0)
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+20)) = l2
	v432 = v409
	v433 = l2
	goto L109
L109:
	;
	if base.Ui32(l4+v416) <= base.Ui32(v433) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	if l5 != 0 {
		goto L125
	} else {
		goto L126
	}
L111:
	;
	goto L110
L112:
	;
	v475 = v432
	v477 = v433
	goto L111
L113:
	;
	goto L114
L114:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	if base.Ui32(v440) <= base.Ui32(int32(2047)) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v432)+24))
	v451 = base.I32_extend16_s(v449)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v450+v451<<(uint(int32(2))%32))))
	if v455 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v443+v440<<(uint(int32(1))%32)))))
	v449 = v447
	goto L115
L117:
	;
	goto L118
L118:
	;
	v448 = F_pg_reg_getcolor(m, v17, v440)
	mBase = m.M
	v449 = v448
	goto L115
L119:
	;
	v460 = F_miss(m, l0, l1, v432, v451, v433+int32(4), l2)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L37
	} else {
		goto L122
	}
L120:
	;
	v464 = v455
	goto L121
L121:
	;
	v466 = v433 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v464)+20)) = v466
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+8)))
	if base.B2i32(v468&int32(2) == int32(0))|base.B2i32(base.Ui32(v466) < base.Ui32(l3+v421)) != 0 {
		v432 = v464
		v433 = v466
		goto L109
	} else {
		goto L124
	}
L122:
	;
	if v460 == int32(0) {
		v638 = v8
		goto L7
	} else {
		goto L123
	}
L123:
	;
	v464 = v460
	goto L121
L124:
	;
	v475 = v464
	v477 = v466
	goto L111
L125:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v479 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v475)+8))
	v578 = v576 & int32(2)
	v579 = int32(0)
	if base.B2i32(v578 == v579)|base.B2i32(base.Ui32(v477) <= base.Ui32(l3)) == v579 {
		goto L158
	} else {
		goto L159
	}
L128:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v483 = v482
	goto L130
L129:
	;
	v483 = v479
	goto L130
L130:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v484 <= int32(0) {
		v555 = v483
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v555
	goto L127
L132:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v484&int32(1) != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+8)))
	if v490&int32(8) != 0 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v502 = v487
	v503 = v483
	v505 = v484
	goto L135
L135:
	;
	if v484 == int32(1) {
		v555 = v503
		goto L131
	} else {
		goto L142
	}
L136:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v487)+20))
	if base.Ui32(v483) < base.Ui32(v493) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v496 = v483
	goto L138
L138:
	;
	v502 = v487 + int32(32)
	v503 = v496
	v505 = v484 - int32(1)
	goto L135
L139:
	;
	v495 = v493
	goto L141
L140:
	;
	v495 = v483
	goto L141
L141:
	;
	v496 = v495
	goto L138
L142:
	;
	v516 = v502
	v518 = v503
	v519 = v505
	goto L143
L143:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+8)))
	if v523&int32(8) != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v555 = v537
	goto L131
L145:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v516)+20))
	if base.Ui32(v518) < base.Ui32(v526) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v529 = v518
	goto L147
L147:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+40)))
	if v531&int32(8) != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v528 = v526
	goto L150
L149:
	;
	v528 = v518
	goto L150
L150:
	;
	v529 = v528
	goto L147
L151:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v516)+52))
	if base.Ui32(v529) < base.Ui32(v534) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v537 = v529
	goto L153
L153:
	;
	v541 = int32(2)
	if v541 < v519 {
		v516 = v516 - int32(-64)
		v518 = v537
		v519 = v519 - v541
		goto L143
	} else {
		goto L157
	}
L154:
	;
	v536 = v534
	goto L156
L155:
	;
	v536 = v529
	goto L156
L156:
	;
	v537 = v536
	goto L153
L157:
	;
	goto L144
L158:
	;
	return v477 - int32(4)
L159:
	;
	goto L160
L160:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v477 != v588)|base.B2i32(v588 != l4) != 0 {
		v616 = v578
		goto L162
	} else {
		goto L163
	}
L161:
	;
	if l6 == int32(0) {
		v638 = v8
		goto L7
	} else {
		goto L170
	}
L162:
	;
	if v616 != 0 {
		goto L167
	} else {
		goto L168
	}
L163:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v599 = int32(*(*int16)(unsafe.Add(mBase, uint32(v592+(v593^int32(-1))&int32(2))+24)))
	v600 = F_miss(m, l0, l1, v475, v599, v477, l2)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L37
	} else {
		goto L164
	}
L164:
	;
	if v600 == int32(0) {
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v600)+8))
	v606 = v604 & int32(2)
	if v606|base.B2i32(l6 == int32(0)) != 0 {
		v616 = v606
		goto L162
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(1)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v600)+8))
	v616 = v612 & int32(2)
	goto L162
L167:
	;
	v618 = v477
	goto L169
L168:
	;
	v618 = int32(0)
	goto L169
L169:
	;
	return v618
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(1)
	v638 = v8
	goto L7
}
func F_show_incremental_sort_group_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
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
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v149 int64
	_ = v149
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v169 int32
	_ = v169
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int64
	_ = v258
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v14&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L5
L2:
	;
	v29 = int32(0)
	v30 = v14
	goto L3
L3:
	;
	if v30&int32(2) != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v26 = F_lappend(m, int32(0), v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_show_incremental_sort_group_info[0]))
	goto L7
L7:
	;
	goto L4
L8:
	;
	return
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v29 = v26
	v30 = v28
	goto L3
L10:
	;
	goto L14
L11:
	;
	v44 = v29
	v45 = v30
	goto L12
L12:
	;
	if v45&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v41 = F_lappend(m, v29, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L17
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_show_incremental_sort_group_info[1]))
	goto L16
L16:
	;
	goto L13
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v44 = v41
	v45 = v43
	goto L12
L18:
	;
	goto L22
L19:
	;
	v59 = v44
	v60 = v45
	goto L20
L20:
	;
	if v60&int32(8) != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v56 = F_lappend(m, v44, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L25
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_show_incremental_sort_group_info[2]))
	goto L24
L24:
	;
	goto L21
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v59 = v56
	v60 = v58
	goto L20
L26:
	;
	goto L30
L27:
	;
	v73 = v59
	goto L28
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v74 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v71 = F_lappend(m, v59, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L33
	}
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_show_incremental_sort_group_info[3]))
	goto L32
L32:
	;
	goto L29
L33:
	;
	v73 = v71
	goto L28
L34:
	;
	m.G0 = v12 + int32(160)
	return
L35:
	;
	if l2 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v193 = v12 + int32(144)
	F_initStringInfo(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L75
	}
L38:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	F_appendStringInfoSpaces(m, v77, v78<<(uint(int32(1))%32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l1
	F_appendStringInfo(m, v83, int32(_a_F_show_incremental_sort_group_info_0), v12-int32(-64))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	if v73 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) < v149 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_appendStringInfoString(m, v94, int32(_a_F_show_incremental_sort_group_info_1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if int32(1) < v101 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L43
L48:
	;
	v104 = int32(_a_F_show_incremental_sort_group_info_2)
	goto L50
L49:
	;
	v104 = int32(_a_F_show_incremental_sort_group_info_1)
	goto L50
L50:
	;
	F_appendStringInfoString(m, v98, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v107 <= int32(0) {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	v117 = int32(0)
	goto L53
L53:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v117<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v120, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L55
	}
L54:
	;
	goto L43
L55:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v117 < v128-int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_appendStringInfoString(m, v132, int32(_a_F_show_incremental_sort_group_info_3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v137 = v117 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v137 < v138 {
		v117 = v137
		goto L53
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L54
L61:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v155 = int32(_a_F_show_incremental_sort_group_info_4)
	goto L65
L62:
	;
	goto L63
L63:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v173 <= int64(0) {
		goto L34
	} else {
		goto L69
	}
L64:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v155
	v162 = base.I64_div_s(v153, v152)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v155
	F_appendStringInfo(m, v158, int32(_a_F_show_incremental_sort_group_info_5), v12+int32(32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L68
	}
L65:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L63
L69:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v180 = int32(_a_F_show_incremental_sort_group_info_6)
	goto L72
L70:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v180
	v186 = base.I64_div_s(v177, v176)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v180
	F_appendStringInfo(m, v182, int32(_a_F_show_incremental_sort_group_info_5), v12)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	goto L34
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = l1
	F_appendStringInfo(m, v193, int32(_a_F_show_incremental_sort_group_info_7), v12+int32(112))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
	F_ExplainOpenGroup(m, int32(_a_F_show_incremental_sort_group_info_8), v203, int32(1), l3)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v209 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_9), int32(0), v209, l3)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_ExplainPropertyList(m, int32(_a_F_show_incremental_sort_group_info_10), v73, l3)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) < v215 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v219 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	goto L84
L81:
	;
	goto L82
L82:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if int64(0) < v258 {
		goto L93
	} else {
		goto L94
	}
L83:
	;
	v225 = v12 + int32(128)
	F_initStringInfo(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L87
	}
L84:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = int32(_a_F_show_incremental_sort_group_info_4)
	F_appendStringInfo(m, v225, int32(_a_F_show_incremental_sort_group_info_11), v12+int32(96))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	F_ExplainOpenGroup(m, int32(_a_F_show_incremental_sort_group_info_12), v235, int32(1), l3)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v241 = base.I64_div_s(v219, v218)
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_13), int32(_a_F_show_incremental_sort_group_info_14), v241, l3)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v246 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_15), int32(_a_F_show_incremental_sort_group_info_14), v246, l3)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	F_ExplainCloseGroup(m, int32(_a_F_show_incremental_sort_group_info_12), int32(1), l3)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	goto L82
L93:
	;
	v261 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v262 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	goto L98
L94:
	;
	goto L95
L95:
	;
	F_ExplainCloseGroup(m, int32(_a_F_show_incremental_sort_group_info_8), int32(1), l3)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L8
	} else {
		goto L106
	}
L96:
	;
	v268 = v12 + int32(128)
	F_initStringInfo(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(_a_F_show_incremental_sort_group_info_6)
	F_appendStringInfo(m, v268, int32(_a_F_show_incremental_sort_group_info_11), v12+int32(80))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	F_ExplainOpenGroup(m, int32(_a_F_show_incremental_sort_group_info_12), v278, int32(1), l3)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	v284 = base.I64_div_s(v262, v261)
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_13), int32(_a_F_show_incremental_sort_group_info_14), v284, l3)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExplainPropertyInteger(m, int32(_a_F_show_incremental_sort_group_info_15), int32(_a_F_show_incremental_sort_group_info_14), v289, l3)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	F_ExplainCloseGroup(m, int32(_a_F_show_incremental_sort_group_info_12), int32(1), l3)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	goto L95
L106:
	;
	goto L34
}
func F_sigemptyset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
}
func F_slice_to(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v7 < int32(0) {
		if l1 == int32(0) {
			return int32(0)
		} else {
			F_pfree(m, l1-int32(8))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v10 < v7 {
			if l1 == int32(0) {
				return int32(0)
			} else {
				F_pfree(m, l1-int32(8))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v12 < v10 {
				if l1 == int32(0) {
					return int32(0)
				} else {
					F_pfree(m, l1-int32(8))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v14 == int32(0) {
					if l1 == int32(0) {
						return int32(0)
					} else {
						F_pfree(m, l1-int32(8))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v14-int32(4))))
					if v12 <= v19 {
						v36 = v10 - v7
						v38 = l1 - int32(8)
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						if v39 < v36 {
							v43 = F_repalloc(m, v38, v36+int32(29))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								if v43 == int32(0) {
									F_pfree(m, v38)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										return int32(0)
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v43))) = v36 + int32(20)
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v54 = v43 + int32(8)
									v55 = v50
									v56 = v51
									if v36 != 0 {
										base.MemoryCopy(m, v54, v55+v56, v36)
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v54-int32(4)))) = v36
									return v54
								}
							}
						} else {
							v54 = l1
							v55 = v7
							v56 = v14
							if v36 != 0 {
								base.MemoryCopy(m, v54, v55+v56, v36)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v54-int32(4)))) = v36
							return v54
						}
					} else {
						if l1 == int32(0) {
							return int32(0)
						} else {
							F_pfree(m, l1-int32(8))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						}
					}
				}
			}
		}
	}
}
func F_slotsync_worker_onexit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[0]))
	if v4 != 0 {
		F_ReplicationSlotRelease(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_ReplicationSlotCleanup(m, int32(0))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(1)
				if v12 != 0 {
					v16 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
					F_s_lock(m, v16+int32(16), int32(_a_F_slotsync_worker_onexit_0), int32(1231), int32(_a_F_slotsync_worker_onexit_1))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
						v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])))
						if v29 != 0 {
							v30 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v30)
							*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])) = uint8(v30)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
						return
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])))
					if v29 != 0 {
						v30 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v30)
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])) = uint8(v30)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
					return
				}
			}
		}
	} else {
		F_ReplicationSlotCleanup(m, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(1)
			if v12 != 0 {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
				F_s_lock(m, v16+int32(16), int32(_a_F_slotsync_worker_onexit_0), int32(1231), int32(_a_F_slotsync_worker_onexit_1))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])))
					if v29 != 0 {
						v30 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v30)
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])) = uint8(v30)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
					return
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
				v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])))
				if v29 != 0 {
					v30 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v30)
					*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_worker_onexit[2])) = uint8(v30)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
				return
			}
		}
	}
}
func F_smgrdestroyall(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v1 = int32(0)
	v5 = int32(_a_F_smgrdestroyall_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0])) = v7 + int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[1]))
	if base.B2i32(v12 == v1)|base.B2i32(v12 == int32(_a_F_smgrdestroyall_1)) == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L15
	}
L2:
	;
	v20 = v12
	goto L5
L3:
	;
	goto L4
L4:
	;
	v70 = int32(_a_F_smgrdestroyall_0)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0])) = v72 - int32(1)
	return
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v25 = int32(_a_F_smgrdestroyall_0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0])) = v27 + int32(1)
	v32 = v20 - int32(76)
	F_mdclose(m, v32, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	return
L8:
	;
	F_mdclose(m, v32, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_mdclose(m, v32, int32(2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_mdclose(m, v32, int32(3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v48
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[2]))
	v54 = F_hash_search(m, v51, v32, int32(2), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v54 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v58 = int32(_a_F_smgrdestroyall_0)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdestroyall[0])) = v60 - int32(1)
	if v24 != int32(_a_F_smgrdestroyall_1) {
		v20 = v24
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L6
L15:
	;
	F_errmsg_internal(m, int32(_a_F_smgrdestroyall_2), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_smgrdestroyall_3), int32(339), int32(_a_F_smgrdestroyall_4))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_smgrnblocks_cached(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrnblocks_cached[0])))
	if v4 == int32(1) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20))
		if v10 != int32(-1) {
			v15 = v10
		} else {
			v15 = int32(-1)
		}
	} else {
		v15 = int32(-1)
	}
	return v15
}
func F_smgrprefetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v11 = int32(_a_F_smgrprefetch_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[0])) = v13 + int32(1)
	if base.Ui64(int64(4294967295)) < base.Ui64(base.I64_extend_i32_s(l3)+base.I64_extend_i32_u(l2)) {
		v114 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v118 = int32(_a_F_smgrprefetch_0)
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[0])) = v120 - int32(1)
	return v114
L2:
	;
	if l3 <= int32(0) {
		v114 = int32(1)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = l2
	v28 = l3
	goto L4
L4:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrprefetch[1])))
	if v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v114 = v46
	goto L1
L6:
	;
	v40 = int32(2)
	goto L8
L7:
	;
	v40 = int32(1)
	goto L8
L8:
	;
	v41 = F__mdfd_getseg(m, l0, l1, v27, int32(0), v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v45 = int32(0)
	v46 = base.B2i32(v41 != v45)
	if v41 == v45 {
		v114 = v46
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v50 = v27 & int32(_a_F_smgrprefetch_1)
	v55 = int32(_a_F_smgrprefetch_2) - v50
	if base.Ui32(v28) < base.Ui32(v55) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v57 = v28
	goto L14
L13:
	;
	v57 = v55
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v62 = F_FileAccess(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if int32(0) <= v62 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	goto L19
L17:
	;
	goto L18
L18:
	;
	v105 = v28 - v57
	if int32(0) < v105 {
		v27 = v27 + v57
		v28 = v105
		goto L4
	} else {
		goto L22
	}
L19:
	;
	v78 = int32(_a_F_smgrprefetch_3)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(167772180)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[3]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v61*int32(48))))
	v87 = F_posix_fadvise(m, v85, base.I64_extend_i32_u(v50<<(uint(int32(13))%32)), base.I64_extend_i32_u(v57<<(uint(int32(13))%32)), int32(3))
	mBase = m.M
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_smgrprefetch[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(0)
	if v87 == int32(27) {
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	goto L20
L22:
	;
	goto L5
}
func F_sort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(1)
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v18 != int32(2) {
			v66 = int32(0)
			v67 = v17
			v69 = v2
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v70 != 0 {
				v71 = F_array_contains_nulls(m, v13)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if v71 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_sort_0), int32(0))
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_sort_1), int32(206), int32(_a_F_sort_2))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
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
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v75 = v13 + int32(16)
						v76 = F_ArrayGetNItemsSafe(m, v73, v75)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							if int32(2) <= v76 {
								v80 = int32(1)
								if v67 != 0 {
									v119 = v80
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v121 = F_ArrayGetNItemsSafe(m, v120, v75)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
										if v124 != 0 {
											v132 = v124
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v132+v13, v121, v10+int32(15))
										mBase = m.M
										m.G0 = v10 + int32(16)
										return v13
									}
								} else {
									switch v69 - int32(3) {
									case 0:
										v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
										if v83|int32(32) != int32(97) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
											if v88|int32(32) != int32(115) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
												if v93|int32(32) == int32(99) {
													v119 = v80
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v121 = F_ArrayGetNItemsSafe(m, v120, v75)
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														if v124 != 0 {
															v132 = v124
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														F_isort(m, v132+v13, v121, v10+int32(15))
														mBase = m.M
														m.G0 = v10 + int32(16)
														return v13
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
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
									case 1:
										v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
										if v98|int32(32) != int32(100) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
											v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
											if v103|int32(32) != int32(101) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
												if v108|int32(32) != int32(115) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
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
													v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
													if v114|int32(32) != int32(99) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_sort_3), int32(0))
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
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
														v119 = int32(0)
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v121 = F_ArrayGetNItemsSafe(m, v120, v75)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
															if v124 != 0 {
																v132 = v124
															} else {
																v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															F_isort(m, v132+v13, v121, v10+int32(15))
															mBase = m.M
															m.G0 = v10 + int32(16)
															return v13
														}
													}
												}
											}
										}
									default:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
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
							} else {
								m.G0 = v10 + int32(16)
								return v13
							}
						}
					}
				}
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				v75 = v13 + int32(16)
				v76 = F_ArrayGetNItemsSafe(m, v73, v75)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					if int32(2) <= v76 {
						v80 = int32(1)
						if v67 != 0 {
							v119 = v80
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v121 = F_ArrayGetNItemsSafe(m, v120, v75)
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								if v124 != 0 {
									v132 = v124
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v132+v13, v121, v10+int32(15))
								mBase = m.M
								m.G0 = v10 + int32(16)
								return v13
							}
						} else {
							switch v69 - int32(3) {
							case 0:
								v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
								if v83|int32(32) != int32(97) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_sort_3), int32(0))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
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
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
									if v88|int32(32) != int32(115) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
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
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
										if v93|int32(32) == int32(99) {
											v119 = v80
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v121 = F_ArrayGetNItemsSafe(m, v120, v75)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
												v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												if v124 != 0 {
													v132 = v124
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												F_isort(m, v132+v13, v121, v10+int32(15))
												mBase = m.M
												m.G0 = v10 + int32(16)
												return v13
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
							case 1:
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
								if v98|int32(32) != int32(100) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_sort_3), int32(0))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
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
									v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
									if v103|int32(32) != int32(101) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
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
										v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
										if v108|int32(32) != int32(115) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
											if v114|int32(32) != int32(99) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
												v119 = int32(0)
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v121 = F_ArrayGetNItemsSafe(m, v120, v75)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													if v124 != 0 {
														v132 = v124
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v132+v13, v121, v10+int32(15))
													mBase = m.M
													m.G0 = v10 + int32(16)
													return v13
												}
											}
										}
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_sort_3), int32(0))
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
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
					} else {
						m.G0 = v10 + int32(16)
						return v13
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v23 = F_pg_detoast_datum_packed(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v66 = int32(0)
					v67 = v17
					v69 = v2
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
					if v28 == int32(1) {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
						if v34 == int32(18) {
							v37 = int32(16)
						} else {
							v37 = int32(0)
						}
						if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v44 = int32(4)
						} else {
							v44 = v37
						}
						v57 = v17
						v58 = v44
					} else {
						if v28&int32(1) != 0 {
							v47 = int32(1)
							v57 = v28
							v58 = int32(base.Ui32(v28)>>(uint(v47)%32)) - v47
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
							v57 = v51
							v58 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v59 = int32(1)
					if v57&v59 != 0 {
						v63 = v59
					} else {
						v63 = int32(4)
					}
					v66 = v23 + v63
					v67 = int32(0)
					v69 = v58
				}
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v70 != 0 {
					v71 = F_array_contains_nulls(m, v13)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						if v71 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_sort_0), int32(0))
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_sort_1), int32(206), int32(_a_F_sort_2))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
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
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v75 = v13 + int32(16)
							v76 = F_ArrayGetNItemsSafe(m, v73, v75)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								if int32(2) <= v76 {
									v80 = int32(1)
									if v67 != 0 {
										v119 = v80
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v121 = F_ArrayGetNItemsSafe(m, v120, v75)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
											if v124 != 0 {
												v132 = v124
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											F_isort(m, v132+v13, v121, v10+int32(15))
											mBase = m.M
											m.G0 = v10 + int32(16)
											return v13
										}
									} else {
										switch v69 - int32(3) {
										case 0:
											v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
											if v83|int32(32) != int32(97) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
												v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
												if v88|int32(32) != int32(115) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
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
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
													if v93|int32(32) == int32(99) {
														v119 = v80
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v121 = F_ArrayGetNItemsSafe(m, v120, v75)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
															if v124 != 0 {
																v132 = v124
															} else {
																v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															F_isort(m, v132+v13, v121, v10+int32(15))
															mBase = m.M
															m.G0 = v10 + int32(16)
															return v13
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_sort_3), int32(0))
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
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
										case 1:
											v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
											if v98|int32(32) != int32(100) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
												v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
												if v103|int32(32) != int32(101) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
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
													v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
													if v108|int32(32) != int32(115) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_sort_3), int32(0))
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
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
														v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
														if v114|int32(32) != int32(99) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50856066))
																mBase = m.M
																v167 = m.ExcPending
																if v167 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_sort_3), int32(0))
																	mBase = m.M
																	v171 = m.ExcPending
																	if v171 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
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
															v119 = int32(0)
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v121 = F_ArrayGetNItemsSafe(m, v120, v75)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int32(0)
															} else {
																*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
																v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
																if v124 != 0 {
																	v132 = v124
																} else {
																	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																	v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
																}
																F_isort(m, v132+v13, v121, v10+int32(15))
																mBase = m.M
																m.G0 = v10 + int32(16)
																return v13
															}
														}
													}
												}
											}
										default:
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
								} else {
									m.G0 = v10 + int32(16)
									return v13
								}
							}
						}
					}
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v75 = v13 + int32(16)
					v76 = F_ArrayGetNItemsSafe(m, v73, v75)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						if int32(2) <= v76 {
							v80 = int32(1)
							if v67 != 0 {
								v119 = v80
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v121 = F_ArrayGetNItemsSafe(m, v120, v75)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									if v124 != 0 {
										v132 = v124
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									F_isort(m, v132+v13, v121, v10+int32(15))
									mBase = m.M
									m.G0 = v10 + int32(16)
									return v13
								}
							} else {
								switch v69 - int32(3) {
								case 0:
									v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
									if v83|int32(32) != int32(97) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
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
										v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
										if v88|int32(32) != int32(115) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
											if v93|int32(32) == int32(99) {
												v119 = v80
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v121 = F_ArrayGetNItemsSafe(m, v120, v75)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													if v124 != 0 {
														v132 = v124
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v132+v13, v121, v10+int32(15))
													mBase = m.M
													m.G0 = v10 + int32(16)
													return v13
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
								case 1:
									v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
									if v98|int32(32) != int32(100) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_sort_3), int32(0))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
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
										v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
										if v103|int32(32) != int32(101) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_sort_3), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
											if v108|int32(32) != int32(115) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_sort_3), int32(0))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
												v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
												if v114|int32(32) != int32(99) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_sort_3), int32(0))
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
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
													v119 = int32(0)
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v121 = F_ArrayGetNItemsSafe(m, v120, v75)
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v119)
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														if v124 != 0 {
															v132 = v124
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v132 = (v125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														F_isort(m, v132+v13, v121, v10+int32(15))
														mBase = m.M
														m.G0 = v10 + int32(16)
														return v13
													}
												}
											}
										}
									}
								default:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_sort_3), int32(0))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_sort_1), int32(224), int32(_a_F_sort_2))
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
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
						} else {
							m.G0 = v10 + int32(16)
							return v13
						}
					}
				}
			}
		}
	}
}
func F_sort_pending_writebacks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v720 int64
	_ = v720
	var v722 int64
	_ = v722
	var v724 int32
	_ = v724
	var v726 int64
	_ = v726
	var v728 int64
	_ = v728
	var v730 int32
	_ = v730
	var v732 int64
	_ = v732
	var v734 int64
	_ = v734
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v787 int64
	_ = v787
	var v789 int64
	_ = v789
	var v791 int32
	_ = v791
	var v793 int64
	_ = v793
	var v795 int64
	_ = v795
	var v797 int32
	_ = v797
	var v799 int64
	_ = v799
	var v801 int64
	_ = v801
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int64
	_ = v860
	var v862 int64
	_ = v862
	var v864 int32
	_ = v864
	var v866 int64
	_ = v866
	var v868 int64
	_ = v868
	var v870 int32
	_ = v870
	var v872 int64
	_ = v872
	var v874 int64
	_ = v874
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v914 int32
	_ = v914
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int64
	_ = v927
	var v929 int64
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int64
	_ = v934
	var v936 int64
	_ = v936
	var v938 int32
	_ = v938
	var v940 int64
	_ = v940
	var v942 int64
	_ = v942
	var v945 int32
	_ = v945
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v980 int32
	_ = v980
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int64
	_ = v992
	var v994 int64
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int64
	_ = v999
	var v1001 int64
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int64
	_ = v1005
	var v1007 int64
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1044 int32
	_ = v1044
	var v1046 int64
	_ = v1046
	var v1048 int64
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int64
	_ = v1052
	var v1054 int64
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int64
	_ = v1058
	var v1060 int64
	_ = v1060
	var v1062 int32
	_ = v1062
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = l0
	v19 = l1
	goto L1
L1:
	;
	v32 = v18 + int32(20)
	v34 = v19
	goto L3
L2:
	;
	m.G0 = v16 + int32(32)
	return
L3:
	;
	v48 = v18 + v34*int32(20)
	if base.Ui32(v34) <= base.Ui32(int32(6)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	goto L4
L6:
	;
	if base.Ui32(v34) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v148 = v32
	goto L28
L9:
	;
	v63 = v32
	goto L10
L10:
	;
	if base.Ui32(v63) <= base.Ui32(v18) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	v144 = v63 + int32(20)
	if base.Ui32(v144) < base.Ui32(v48) {
		v63 = v144
		goto L10
	} else {
		goto L27
	}
L13:
	;
	v70 = v63
	goto L14
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(12))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	if base.Ui32(v82) < base.Ui32(v83) {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v86 = v70 - int32(20)
	if base.Ui32(v83) < base.Ui32(v82) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v86)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+8)) = v119
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v123
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v127
	if base.Ui32(v18) < base.Ui32(v86) {
		v70 = v86
		goto L14
	} else {
		goto L26
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(16))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if base.Ui32(v90) < base.Ui32(v91) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v91) < base.Ui32(v90) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if base.Ui32(v94) < base.Ui32(v95) {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(v95) < base.Ui32(v94) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(8))))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if v100 < v101 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	if v101 < v100 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(4))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	if base.Ui32(v106) <= base.Ui32(v107) {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	goto L17
L26:
	;
	goto L15
L27:
	;
	goto L11
L28:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	if base.Ui32(v161) < base.Ui32(v162) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v199 = v18 + int32(base.Ui32(v34)>>(uint(int32(1))%32))*int32(20)
	if v34 != int32(7) {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	goto L29
L31:
	;
	v191 = v148 + int32(20)
	if base.Ui32(v191) < base.Ui32(v48) {
		v148 = v191
		goto L28
	} else {
		goto L41
	}
L32:
	;
	if base.Ui32(v162) < base.Ui32(v161) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if base.Ui32(v167) < base.Ui32(v168) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32(v168) < base.Ui32(v167) {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(20))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if base.Ui32(v173) < base.Ui32(v174) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v174) < base.Ui32(v173) {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(8))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	if v179 < v180 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v180 < v179 {
		goto L30
	} else {
		goto L39
	}
L39:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(4))))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	if base.Ui32(v186) < base.Ui32(v185) {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	goto L31
L41:
	;
	goto L5
L42:
	;
	v203 = v48 - int32(20)
	if base.Ui32(v34) < base.Ui32(int32(41)) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v713 = v199
	goto L44
L44:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v718
	v720 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v720
	v722 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v713)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v724
	v726 = *(*int64)(unsafe.Add(mBase, uint32(v713)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v726
	v728 = *(*int64)(unsafe.Add(mBase, uint32(v713)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v728
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v713)+16)) = v730
	v732 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v713)+8)) = v732
	v734 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v713))) = v734
	v737 = v48 - int32(20)
	v740 = v737
	v741 = v32
	v743 = v32
	v745 = v737
	goto L277
L45:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v586)))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v585)))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v585)+8))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v586)+8))
	if base.Ui32(v604) < base.Ui32(v605) {
		goto L225
	} else {
		goto L226
	}
L46:
	;
	v585 = v18
	v586 = v199
	v587 = v203
	goto L45
L47:
	;
	goto L48
L48:
	;
	v207 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	v209 = v207 * int32(20)
	v210 = v18 + v209
	v213 = v18 + v207*int32(40)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	if base.Ui32(v227) < base.Ui32(v228) {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v336 = v207 * int32(-20)
	v337 = v199 + v336
	v338 = v199 + v209
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v337)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
	if base.Ui32(v352) < base.Ui32(v353) {
		goto L111
	} else {
		goto L112
	}
L50:
	;
	v334 = v18
	goto L49
L51:
	;
	v334 = v213
	goto L49
L52:
	;
	v334 = v310
	goto L49
L53:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	if base.Ui32(v228) < base.Ui32(v280) {
		goto L85
	} else {
		goto L86
	}
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	if base.Ui32(v228) < base.Ui32(v244) {
		v310 = v210
		goto L52
	} else {
		goto L64
	}
L55:
	;
	if base.Ui32(v228) < base.Ui32(v227) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	if base.Ui32(v225) < base.Ui32(v223) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(v223) < base.Ui32(v225) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	if base.Ui32(v226) < base.Ui32(v224) {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v224) < base.Ui32(v226) {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	if v235 < v236 {
		goto L54
	} else {
		goto L61
	}
L61:
	;
	if v236 < v235 {
		goto L53
	} else {
		goto L62
	}
L62:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	if base.Ui32(v240) <= base.Ui32(v239) {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	goto L54
L64:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if base.Ui32(v244) < base.Ui32(v228) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if base.Ui32(v227) < base.Ui32(v244) {
		goto L51
	} else {
		goto L74
	}
L66:
	;
	if base.Ui32(v223) < base.Ui32(v246) {
		v310 = v210
		goto L52
	} else {
		goto L67
	}
L67:
	;
	if base.Ui32(v246) < base.Ui32(v223) {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	if base.Ui32(v224) < base.Ui32(v247) {
		v310 = v210
		goto L52
	} else {
		goto L69
	}
L69:
	;
	if base.Ui32(v247) < base.Ui32(v224) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	if v253 < v254 {
		v310 = v210
		goto L52
	} else {
		goto L71
	}
L71:
	;
	if v254 < v253 {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	if base.Ui32(v257) < base.Ui32(v258) {
		v310 = v210
		goto L52
	} else {
		goto L73
	}
L73:
	;
	goto L65
L74:
	;
	if base.Ui32(v244) < base.Ui32(v227) {
		goto L50
	} else {
		goto L75
	}
L75:
	;
	if base.Ui32(v225) < base.Ui32(v246) {
		goto L51
	} else {
		goto L76
	}
L76:
	;
	if base.Ui32(v246) < base.Ui32(v225) {
		goto L50
	} else {
		goto L77
	}
L77:
	;
	if base.Ui32(v226) < base.Ui32(v247) {
		goto L51
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(v247) < base.Ui32(v226) {
		goto L50
	} else {
		goto L79
	}
L79:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	if v268 < v269 {
		goto L51
	} else {
		goto L80
	}
L80:
	;
	if v269 < v268 {
		goto L50
	} else {
		goto L81
	}
L81:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	if base.Ui32(v272) < base.Ui32(v273) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v275 = v213
	goto L84
L83:
	;
	v275 = v18
	goto L84
L84:
	;
	v334 = v275
	goto L49
L85:
	;
	if base.Ui32(v227) < base.Ui32(v280) {
		goto L50
	} else {
		goto L95
	}
L86:
	;
	if base.Ui32(v280) < base.Ui32(v228) {
		v310 = v210
		goto L52
	} else {
		goto L87
	}
L87:
	;
	if base.Ui32(v223) < base.Ui32(v278) {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(v278) < base.Ui32(v223) {
		v310 = v210
		goto L52
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(v224) < base.Ui32(v279) {
		goto L85
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(v279) < base.Ui32(v224) {
		v310 = v210
		goto L52
	} else {
		goto L91
	}
L91:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	if v287 < v288 {
		goto L85
	} else {
		goto L92
	}
L92:
	;
	if v288 < v287 {
		v310 = v210
		goto L52
	} else {
		goto L93
	}
L93:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	if base.Ui32(v292) < base.Ui32(v291) {
		v310 = v210
		goto L52
	} else {
		goto L94
	}
L94:
	;
	goto L85
L95:
	;
	if base.Ui32(v280) < base.Ui32(v227) {
		goto L51
	} else {
		goto L96
	}
L96:
	;
	if base.Ui32(v225) < base.Ui32(v278) {
		goto L50
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(v278) < base.Ui32(v225) {
		goto L51
	} else {
		goto L98
	}
L98:
	;
	if base.Ui32(v226) < base.Ui32(v279) {
		goto L50
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(v279) < base.Ui32(v226) {
		goto L51
	} else {
		goto L100
	}
L100:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	if v302 < v303 {
		goto L50
	} else {
		goto L101
	}
L101:
	;
	if v303 < v302 {
		goto L51
	} else {
		goto L102
	}
L102:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	if base.Ui32(v306) < base.Ui32(v307) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v309 = v18
	goto L105
L104:
	;
	v309 = v213
	goto L105
L105:
	;
	v310 = v309
	goto L52
L106:
	;
	v462 = v203 + v207*int32(-40)
	v463 = v203 + v336
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v463)))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v462)+4))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v462)+8))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	if base.Ui32(v477) < base.Ui32(v478) {
		goto L168
	} else {
		goto L169
	}
L107:
	;
	v459 = v337
	goto L106
L108:
	;
	v459 = v338
	goto L106
L109:
	;
	v459 = v435
	goto L106
L110:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	if base.Ui32(v353) < base.Ui32(v405) {
		goto L142
	} else {
		goto L143
	}
L111:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	if base.Ui32(v353) < base.Ui32(v369) {
		v435 = v199
		goto L109
	} else {
		goto L121
	}
L112:
	;
	if base.Ui32(v353) < base.Ui32(v352) {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	if base.Ui32(v350) < base.Ui32(v348) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	if base.Ui32(v348) < base.Ui32(v350) {
		goto L110
	} else {
		goto L115
	}
L115:
	;
	if base.Ui32(v351) < base.Ui32(v349) {
		goto L111
	} else {
		goto L116
	}
L116:
	;
	if base.Ui32(v349) < base.Ui32(v351) {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	if v360 < v361 {
		goto L111
	} else {
		goto L118
	}
L118:
	;
	if v361 < v360 {
		goto L110
	} else {
		goto L119
	}
L119:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	if base.Ui32(v365) <= base.Ui32(v364) {
		goto L110
	} else {
		goto L120
	}
L120:
	;
	goto L111
L121:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	if base.Ui32(v369) < base.Ui32(v353) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	if base.Ui32(v352) < base.Ui32(v369) {
		goto L108
	} else {
		goto L131
	}
L123:
	;
	if base.Ui32(v348) < base.Ui32(v371) {
		v435 = v199
		goto L109
	} else {
		goto L124
	}
L124:
	;
	if base.Ui32(v371) < base.Ui32(v348) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	if base.Ui32(v349) < base.Ui32(v372) {
		v435 = v199
		goto L109
	} else {
		goto L126
	}
L126:
	;
	if base.Ui32(v372) < base.Ui32(v349) {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	if v378 < v379 {
		v435 = v199
		goto L109
	} else {
		goto L128
	}
L128:
	;
	if v379 < v378 {
		goto L122
	} else {
		goto L129
	}
L129:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v338)+16))
	if base.Ui32(v382) < base.Ui32(v383) {
		v435 = v199
		goto L109
	} else {
		goto L130
	}
L130:
	;
	goto L122
L131:
	;
	if base.Ui32(v369) < base.Ui32(v352) {
		goto L107
	} else {
		goto L132
	}
L132:
	;
	if base.Ui32(v350) < base.Ui32(v371) {
		goto L108
	} else {
		goto L133
	}
L133:
	;
	if base.Ui32(v371) < base.Ui32(v350) {
		goto L107
	} else {
		goto L134
	}
L134:
	;
	if base.Ui32(v351) < base.Ui32(v372) {
		goto L108
	} else {
		goto L135
	}
L135:
	;
	if base.Ui32(v372) < base.Ui32(v351) {
		goto L107
	} else {
		goto L136
	}
L136:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	if v393 < v394 {
		goto L108
	} else {
		goto L137
	}
L137:
	;
	if v394 < v393 {
		goto L107
	} else {
		goto L138
	}
L138:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v338)+16))
	if base.Ui32(v397) < base.Ui32(v398) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v400 = v338
	goto L141
L140:
	;
	v400 = v337
	goto L141
L141:
	;
	v459 = v400
	goto L106
L142:
	;
	if base.Ui32(v352) < base.Ui32(v405) {
		goto L107
	} else {
		goto L152
	}
L143:
	;
	if base.Ui32(v405) < base.Ui32(v353) {
		v435 = v199
		goto L109
	} else {
		goto L144
	}
L144:
	;
	if base.Ui32(v348) < base.Ui32(v403) {
		goto L142
	} else {
		goto L145
	}
L145:
	;
	if base.Ui32(v403) < base.Ui32(v348) {
		v435 = v199
		goto L109
	} else {
		goto L146
	}
L146:
	;
	if base.Ui32(v349) < base.Ui32(v404) {
		goto L142
	} else {
		goto L147
	}
L147:
	;
	if base.Ui32(v404) < base.Ui32(v349) {
		v435 = v199
		goto L109
	} else {
		goto L148
	}
L148:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	if v412 < v413 {
		goto L142
	} else {
		goto L149
	}
L149:
	;
	if v413 < v412 {
		v435 = v199
		goto L109
	} else {
		goto L150
	}
L150:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v338)+16))
	if base.Ui32(v417) < base.Ui32(v416) {
		v435 = v199
		goto L109
	} else {
		goto L151
	}
L151:
	;
	goto L142
L152:
	;
	if base.Ui32(v405) < base.Ui32(v352) {
		goto L108
	} else {
		goto L153
	}
L153:
	;
	if base.Ui32(v350) < base.Ui32(v403) {
		goto L107
	} else {
		goto L154
	}
L154:
	;
	if base.Ui32(v403) < base.Ui32(v350) {
		goto L108
	} else {
		goto L155
	}
L155:
	;
	if base.Ui32(v351) < base.Ui32(v404) {
		goto L107
	} else {
		goto L156
	}
L156:
	;
	if base.Ui32(v404) < base.Ui32(v351) {
		goto L108
	} else {
		goto L157
	}
L157:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	if v427 < v428 {
		goto L107
	} else {
		goto L158
	}
L158:
	;
	if v428 < v427 {
		goto L108
	} else {
		goto L159
	}
L159:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v338)+16))
	if base.Ui32(v431) < base.Ui32(v432) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v434 = v337
	goto L162
L161:
	;
	v434 = v338
	goto L162
L162:
	;
	v435 = v434
	goto L109
L163:
	;
	v585 = v334
	v586 = v459
	v587 = v584
	goto L45
L164:
	;
	v584 = v462
	goto L163
L165:
	;
	v584 = v203
	goto L163
L166:
	;
	v584 = v560
	goto L163
L167:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	if base.Ui32(v478) < base.Ui32(v530) {
		goto L199
	} else {
		goto L200
	}
L168:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	if base.Ui32(v478) < base.Ui32(v494) {
		v560 = v463
		goto L166
	} else {
		goto L178
	}
L169:
	;
	if base.Ui32(v478) < base.Ui32(v477) {
		goto L167
	} else {
		goto L170
	}
L170:
	;
	if base.Ui32(v475) < base.Ui32(v473) {
		goto L168
	} else {
		goto L171
	}
L171:
	;
	if base.Ui32(v473) < base.Ui32(v475) {
		goto L167
	} else {
		goto L172
	}
L172:
	;
	if base.Ui32(v476) < base.Ui32(v474) {
		goto L168
	} else {
		goto L173
	}
L173:
	;
	if base.Ui32(v474) < base.Ui32(v476) {
		goto L167
	} else {
		goto L174
	}
L174:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	if v485 < v486 {
		goto L168
	} else {
		goto L175
	}
L175:
	;
	if v486 < v485 {
		goto L167
	} else {
		goto L176
	}
L176:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v462)+16))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v463)+16))
	if base.Ui32(v490) <= base.Ui32(v489) {
		goto L167
	} else {
		goto L177
	}
L177:
	;
	goto L168
L178:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if base.Ui32(v494) < base.Ui32(v478) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	if base.Ui32(v477) < base.Ui32(v494) {
		goto L165
	} else {
		goto L188
	}
L180:
	;
	if base.Ui32(v473) < base.Ui32(v496) {
		v560 = v463
		goto L166
	} else {
		goto L181
	}
L181:
	;
	if base.Ui32(v496) < base.Ui32(v473) {
		goto L179
	} else {
		goto L182
	}
L182:
	;
	if base.Ui32(v474) < base.Ui32(v497) {
		v560 = v463
		goto L166
	} else {
		goto L183
	}
L183:
	;
	if base.Ui32(v497) < base.Ui32(v474) {
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v503 < v504 {
		v560 = v463
		goto L166
	} else {
		goto L185
	}
L185:
	;
	if v504 < v503 {
		goto L179
	} else {
		goto L186
	}
L186:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v463)+16))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	if base.Ui32(v507) < base.Ui32(v508) {
		v560 = v463
		goto L166
	} else {
		goto L187
	}
L187:
	;
	goto L179
L188:
	;
	if base.Ui32(v494) < base.Ui32(v477) {
		goto L164
	} else {
		goto L189
	}
L189:
	;
	if base.Ui32(v475) < base.Ui32(v496) {
		goto L165
	} else {
		goto L190
	}
L190:
	;
	if base.Ui32(v496) < base.Ui32(v475) {
		goto L164
	} else {
		goto L191
	}
L191:
	;
	if base.Ui32(v476) < base.Ui32(v497) {
		goto L165
	} else {
		goto L192
	}
L192:
	;
	if base.Ui32(v497) < base.Ui32(v476) {
		goto L164
	} else {
		goto L193
	}
L193:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v518 < v519 {
		goto L165
	} else {
		goto L194
	}
L194:
	;
	if v519 < v518 {
		goto L164
	} else {
		goto L195
	}
L195:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v462)+16))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	if base.Ui32(v522) < base.Ui32(v523) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v525 = v203
	goto L198
L197:
	;
	v525 = v462
	goto L198
L198:
	;
	v584 = v525
	goto L163
L199:
	;
	if base.Ui32(v477) < base.Ui32(v530) {
		goto L164
	} else {
		goto L209
	}
L200:
	;
	if base.Ui32(v530) < base.Ui32(v478) {
		v560 = v463
		goto L166
	} else {
		goto L201
	}
L201:
	;
	if base.Ui32(v473) < base.Ui32(v528) {
		goto L199
	} else {
		goto L202
	}
L202:
	;
	if base.Ui32(v528) < base.Ui32(v473) {
		v560 = v463
		goto L166
	} else {
		goto L203
	}
L203:
	;
	if base.Ui32(v474) < base.Ui32(v529) {
		goto L199
	} else {
		goto L204
	}
L204:
	;
	if base.Ui32(v529) < base.Ui32(v474) {
		v560 = v463
		goto L166
	} else {
		goto L205
	}
L205:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v537 < v538 {
		goto L199
	} else {
		goto L206
	}
L206:
	;
	if v538 < v537 {
		v560 = v463
		goto L166
	} else {
		goto L207
	}
L207:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v463)+16))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	if base.Ui32(v542) < base.Ui32(v541) {
		v560 = v463
		goto L166
	} else {
		goto L208
	}
L208:
	;
	goto L199
L209:
	;
	if base.Ui32(v530) < base.Ui32(v477) {
		goto L165
	} else {
		goto L210
	}
L210:
	;
	if base.Ui32(v475) < base.Ui32(v528) {
		goto L164
	} else {
		goto L211
	}
L211:
	;
	if base.Ui32(v528) < base.Ui32(v475) {
		goto L165
	} else {
		goto L212
	}
L212:
	;
	if base.Ui32(v476) < base.Ui32(v529) {
		goto L164
	} else {
		goto L213
	}
L213:
	;
	if base.Ui32(v529) < base.Ui32(v476) {
		goto L165
	} else {
		goto L214
	}
L214:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v552 < v553 {
		goto L164
	} else {
		goto L215
	}
L215:
	;
	if v553 < v552 {
		goto L165
	} else {
		goto L216
	}
L216:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v462)+16))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	if base.Ui32(v556) < base.Ui32(v557) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v559 = v462
	goto L219
L218:
	;
	v559 = v203
	goto L219
L219:
	;
	v560 = v559
	goto L166
L220:
	;
	v713 = v711
	goto L44
L221:
	;
	v711 = v585
	goto L220
L222:
	;
	v711 = v587
	goto L220
L223:
	;
	v711 = v687
	goto L220
L224:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	if base.Ui32(v605) < base.Ui32(v657) {
		goto L256
	} else {
		goto L257
	}
L225:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	if base.Ui32(v605) < base.Ui32(v621) {
		v687 = v586
		goto L223
	} else {
		goto L235
	}
L226:
	;
	if base.Ui32(v605) < base.Ui32(v604) {
		goto L224
	} else {
		goto L227
	}
L227:
	;
	if base.Ui32(v602) < base.Ui32(v600) {
		goto L225
	} else {
		goto L228
	}
L228:
	;
	if base.Ui32(v600) < base.Ui32(v602) {
		goto L224
	} else {
		goto L229
	}
L229:
	;
	if base.Ui32(v603) < base.Ui32(v601) {
		goto L225
	} else {
		goto L230
	}
L230:
	;
	if base.Ui32(v601) < base.Ui32(v603) {
		goto L224
	} else {
		goto L231
	}
L231:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v585)+12))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	if v612 < v613 {
		goto L225
	} else {
		goto L232
	}
L232:
	;
	if v613 < v612 {
		goto L224
	} else {
		goto L233
	}
L233:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v585)+16))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v586)+16))
	if base.Ui32(v617) <= base.Ui32(v616) {
		goto L224
	} else {
		goto L234
	}
L234:
	;
	goto L225
L235:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	if base.Ui32(v621) < base.Ui32(v605) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	if base.Ui32(v604) < base.Ui32(v621) {
		goto L222
	} else {
		goto L245
	}
L237:
	;
	if base.Ui32(v600) < base.Ui32(v623) {
		v687 = v586
		goto L223
	} else {
		goto L238
	}
L238:
	;
	if base.Ui32(v623) < base.Ui32(v600) {
		goto L236
	} else {
		goto L239
	}
L239:
	;
	if base.Ui32(v601) < base.Ui32(v624) {
		v687 = v586
		goto L223
	} else {
		goto L240
	}
L240:
	;
	if base.Ui32(v624) < base.Ui32(v601) {
		goto L236
	} else {
		goto L241
	}
L241:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	if v630 < v631 {
		v687 = v586
		goto L223
	} else {
		goto L242
	}
L242:
	;
	if v631 < v630 {
		goto L236
	} else {
		goto L243
	}
L243:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v586)+16))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v587)+16))
	if base.Ui32(v634) < base.Ui32(v635) {
		v687 = v586
		goto L223
	} else {
		goto L244
	}
L244:
	;
	goto L236
L245:
	;
	if base.Ui32(v621) < base.Ui32(v604) {
		goto L221
	} else {
		goto L246
	}
L246:
	;
	if base.Ui32(v602) < base.Ui32(v623) {
		goto L222
	} else {
		goto L247
	}
L247:
	;
	if base.Ui32(v623) < base.Ui32(v602) {
		goto L221
	} else {
		goto L248
	}
L248:
	;
	if base.Ui32(v603) < base.Ui32(v624) {
		goto L222
	} else {
		goto L249
	}
L249:
	;
	if base.Ui32(v624) < base.Ui32(v603) {
		goto L221
	} else {
		goto L250
	}
L250:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v585)+12))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	if v645 < v646 {
		goto L222
	} else {
		goto L251
	}
L251:
	;
	if v646 < v645 {
		goto L221
	} else {
		goto L252
	}
L252:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v585)+16))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v587)+16))
	if base.Ui32(v649) < base.Ui32(v650) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v652 = v587
	goto L255
L254:
	;
	v652 = v585
	goto L255
L255:
	;
	v711 = v652
	goto L220
L256:
	;
	if base.Ui32(v604) < base.Ui32(v657) {
		goto L221
	} else {
		goto L266
	}
L257:
	;
	if base.Ui32(v657) < base.Ui32(v605) {
		v687 = v586
		goto L223
	} else {
		goto L258
	}
L258:
	;
	if base.Ui32(v600) < base.Ui32(v655) {
		goto L256
	} else {
		goto L259
	}
L259:
	;
	if base.Ui32(v655) < base.Ui32(v600) {
		v687 = v586
		goto L223
	} else {
		goto L260
	}
L260:
	;
	if base.Ui32(v601) < base.Ui32(v656) {
		goto L256
	} else {
		goto L261
	}
L261:
	;
	if base.Ui32(v656) < base.Ui32(v601) {
		v687 = v586
		goto L223
	} else {
		goto L262
	}
L262:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	if v664 < v665 {
		goto L256
	} else {
		goto L263
	}
L263:
	;
	if v665 < v664 {
		v687 = v586
		goto L223
	} else {
		goto L264
	}
L264:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v586)+16))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v587)+16))
	if base.Ui32(v669) < base.Ui32(v668) {
		v687 = v586
		goto L223
	} else {
		goto L265
	}
L265:
	;
	goto L256
L266:
	;
	if base.Ui32(v657) < base.Ui32(v604) {
		goto L222
	} else {
		goto L267
	}
L267:
	;
	if base.Ui32(v602) < base.Ui32(v655) {
		goto L221
	} else {
		goto L268
	}
L268:
	;
	if base.Ui32(v655) < base.Ui32(v602) {
		goto L222
	} else {
		goto L269
	}
L269:
	;
	if base.Ui32(v603) < base.Ui32(v656) {
		goto L221
	} else {
		goto L270
	}
L270:
	;
	if base.Ui32(v656) < base.Ui32(v603) {
		goto L222
	} else {
		goto L271
	}
L271:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v585)+12))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	if v679 < v680 {
		goto L221
	} else {
		goto L272
	}
L272:
	;
	if v680 < v679 {
		goto L222
	} else {
		goto L273
	}
L273:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v585)+16))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v587)+16))
	if base.Ui32(v683) < base.Ui32(v684) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v686 = v585
	goto L276
L275:
	;
	v686 = v587
	goto L276
L276:
	;
	v687 = v686
	goto L223
L277:
	;
	if base.Ui32(v740) < base.Ui32(v741) {
		v814 = v741
		v816 = v743
		goto L279
	} else {
		goto L280
	}
L279:
	;
	if base.Ui32(v814) <= base.Ui32(v740) {
		goto L296
	} else {
		goto L297
	}
L280:
	;
	v755 = v741
	v757 = v743
	goto L281
L281:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v755)+8))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if base.Ui32(v765) < base.Ui32(v766) {
		v806 = v757
		goto L283
	} else {
		goto L284
	}
L282:
	;
	v814 = v809
	v816 = v806
	goto L279
L283:
	;
	v809 = v755 + int32(20)
	if base.Ui32(v809) <= base.Ui32(v740) {
		v755 = v809
		v757 = v806
		goto L281
	} else {
		goto L294
	}
L284:
	;
	if base.Ui32(v766) < base.Ui32(v765) {
		v814 = v755
		v816 = v757
		goto L279
	} else {
		goto L285
	}
L285:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if base.Ui32(v769) < base.Ui32(v770) {
		v806 = v757
		goto L283
	} else {
		goto L286
	}
L286:
	;
	if base.Ui32(v770) < base.Ui32(v769) {
		v814 = v755
		v816 = v757
		goto L279
	} else {
		goto L287
	}
L287:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if base.Ui32(v773) < base.Ui32(v774) {
		v806 = v757
		goto L283
	} else {
		goto L288
	}
L288:
	;
	if base.Ui32(v774) < base.Ui32(v773) {
		v814 = v755
		v816 = v757
		goto L279
	} else {
		goto L289
	}
L289:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v755)+12))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v777 < v778 {
		v806 = v757
		goto L283
	} else {
		goto L290
	}
L290:
	;
	if v778 < v777 {
		v814 = v755
		v816 = v757
		goto L279
	} else {
		goto L291
	}
L291:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v755)+16))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if base.Ui32(v781) < base.Ui32(v782) {
		v806 = v757
		goto L283
	} else {
		goto L292
	}
L292:
	;
	if base.Ui32(v782) < base.Ui32(v781) {
		v814 = v755
		v816 = v757
		goto L279
	} else {
		goto L293
	}
L293:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v757)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v785
	v787 = *(*int64)(unsafe.Add(mBase, uint32(v757)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v787
	v789 = *(*int64)(unsafe.Add(mBase, uint32(v757)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v789
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v755)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v757)+16)) = v791
	v793 = *(*int64)(unsafe.Add(mBase, uint32(v755)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v757)+8)) = v793
	v795 = *(*int64)(unsafe.Add(mBase, uint32(v755)))
	*(*int64)(unsafe.Add(mBase, uint32(v757))) = v795
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v755)+16)) = v797
	v799 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v755)+8)) = v799
	v801 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v755))) = v801
	v806 = v757 + int32(20)
	goto L283
L294:
	;
	goto L282
L295:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v814)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v1044
	v1046 = *(*int64)(unsafe.Add(mBase, uint32(v814)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v1046
	v1048 = *(*int64)(unsafe.Add(mBase, uint32(v814)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v1048
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v827)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v814)+16)) = v1050
	v1052 = *(*int64)(unsafe.Add(mBase, uint32(v827)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v814)+8)) = v1052
	v1054 = *(*int64)(unsafe.Add(mBase, uint32(v827)))
	*(*int64)(unsafe.Add(mBase, uint32(v814))) = v1054
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v827)+16)) = v1056
	v1058 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v827)+8)) = v1058
	v1060 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v827))) = v1060
	v1062 = int32(20)
	v740 = v827 - v1062
	v741 = v814 + v1062
	v743 = v816
	v745 = v832
	goto L277
L296:
	;
	v827 = v740
	v832 = v745
	goto L299
L297:
	;
	v886 = v740
	v891 = v745
	goto L298
L298:
	;
	v898 = int32(20)
	v899 = base.I32_div_s(v816-v18, v898)
	v902 = base.I32_div_s(v814-v816, v898)
	if v899 < v902 {
		goto L313
	} else {
		goto L314
	}
L299:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v827)+8))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if base.Ui32(v838) < base.Ui32(v839) {
		goto L295
	} else {
		goto L301
	}
L300:
	;
	v886 = v882
	v891 = v880
	goto L298
L301:
	;
	if base.Ui32(v839) < base.Ui32(v838) {
		v880 = v832
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v882 = v827 - int32(20)
	if base.Ui32(v814) <= base.Ui32(v882) {
		v827 = v882
		v832 = v880
		goto L299
	} else {
		goto L312
	}
L303:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v827)+4))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if base.Ui32(v842) < base.Ui32(v843) {
		goto L295
	} else {
		goto L304
	}
L304:
	;
	if base.Ui32(v843) < base.Ui32(v842) {
		v880 = v832
		goto L302
	} else {
		goto L305
	}
L305:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v827)))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if base.Ui32(v846) < base.Ui32(v847) {
		goto L295
	} else {
		goto L306
	}
L306:
	;
	if base.Ui32(v847) < base.Ui32(v846) {
		v880 = v832
		goto L302
	} else {
		goto L307
	}
L307:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v827)+12))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v850 < v851 {
		goto L295
	} else {
		goto L308
	}
L308:
	;
	if v851 < v850 {
		v880 = v832
		goto L302
	} else {
		goto L309
	}
L309:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v827)+16))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if base.Ui32(v854) < base.Ui32(v855) {
		goto L295
	} else {
		goto L310
	}
L310:
	;
	if base.Ui32(v855) < base.Ui32(v854) {
		v880 = v832
		goto L302
	} else {
		goto L311
	}
L311:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v827)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v858
	v860 = *(*int64)(unsafe.Add(mBase, uint32(v827)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v860
	v862 = *(*int64)(unsafe.Add(mBase, uint32(v827)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v862
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v832)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v827)+16)) = v864
	v866 = *(*int64)(unsafe.Add(mBase, uint32(v832)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v827)+8)) = v866
	v868 = *(*int64)(unsafe.Add(mBase, uint32(v832)))
	*(*int64)(unsafe.Add(mBase, uint32(v827))) = v868
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v832)+16)) = v870
	v872 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v832)+8)) = v872
	v874 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v832))) = v874
	v880 = v832 - int32(20)
	goto L302
L312:
	;
	goto L300
L313:
	;
	v904 = v899
	goto L315
L314:
	;
	v904 = v902
	goto L315
L315:
	;
	if v904 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v914 = int32(0)
	goto L319
L317:
	;
	goto L318
L318:
	;
	v961 = int32(20)
	v962 = base.I32_div_s(v891-v886, v961)
	v965 = base.I32_div_s(v48-v891, v961)
	v967 = v965 - int32(1)
	if v962 < v967 {
		goto L322
	} else {
		goto L323
	}
L319:
	;
	v923 = v914 * int32(20)
	v924 = v18 + v923
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v925
	v927 = *(*int64)(unsafe.Add(mBase, uint32(v924)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v927
	v929 = *(*int64)(unsafe.Add(mBase, uint32(v924)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v929
	v931 = v923 + (v814 + v904*int32(-20))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v931)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v924)+16)) = v932
	v934 = *(*int64)(unsafe.Add(mBase, uint32(v931)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v924)+8)) = v934
	v936 = *(*int64)(unsafe.Add(mBase, uint32(v931)))
	*(*int64)(unsafe.Add(mBase, uint32(v924))) = v936
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v931)+16)) = v938
	v940 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v931)+8)) = v940
	v942 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v931))) = v942
	v945 = v914 + int32(1)
	if v945 != v904 {
		v914 = v945
		goto L319
	} else {
		goto L321
	}
L320:
	;
	goto L318
L321:
	;
	goto L320
L322:
	;
	v969 = v962
	goto L324
L323:
	;
	v969 = v967
	goto L324
L324:
	;
	if v969 != 0 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v980 = int32(0)
	goto L328
L326:
	;
	goto L327
L327:
	;
	if base.Ui32(v902) <= base.Ui32(v962) {
		goto L331
	} else {
		goto L332
	}
L328:
	;
	v988 = v980 * int32(20)
	v989 = v814 + v988
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v989)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v990
	v992 = *(*int64)(unsafe.Add(mBase, uint32(v989)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v992
	v994 = *(*int64)(unsafe.Add(mBase, uint32(v989)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v994
	v996 = v988 + (v48 + v969*int32(-20))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+16)) = v997
	v999 = *(*int64)(unsafe.Add(mBase, uint32(v996)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v989)+8)) = v999
	v1001 = *(*int64)(unsafe.Add(mBase, uint32(v996)))
	*(*int64)(unsafe.Add(mBase, uint32(v989))) = v1001
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v996)+16)) = v1003
	v1005 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v996)+8)) = v1005
	v1007 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v996))) = v1007
	v1010 = v980 + int32(1)
	if v1010 != v969 {
		v980 = v1010
		goto L328
	} else {
		goto L330
	}
L329:
	;
	goto L327
L330:
	;
	goto L329
L331:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v902) {
		goto L334
	} else {
		goto L335
	}
L332:
	;
	goto L333
L333:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v962) {
		goto L338
	} else {
		goto L339
	}
L334:
	;
	F_sort_pending_writebacks(m, v18, v902)
	mBase = m.M
	goto L336
L335:
	;
	goto L336
L336:
	;
	if base.Ui32(v962) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	v18 = v48 + v962*int32(-20)
	v19 = v962
	goto L1
L338:
	;
	F_sort_pending_writebacks(m, v48+v962*int32(-20), v962)
	mBase = m.M
	goto L340
L339:
	;
	goto L340
L340:
	;
	if base.Ui32(int32(1)) < base.Ui32(v902) {
		v34 = v902
		goto L3
	} else {
		goto L341
	}
L341:
	;
	goto L5
}
func F_sortouts_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v8 < v11 {
		return int32(-1)
	} else {
		v15 = int32(1)
		if v11 < v8 {
			v28 = v15
			return v28
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6)+4)))
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			if v17 < v18 {
				return int32(-1)
			} else {
				if v18 < v17 {
					v28 = v15
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					if v24 < v25 {
						v28 = int32(-1)
					} else {
						v28 = base.B2i32(v25 < v24)
					}
				}
				return v28
			}
		}
	}
}
func F_soundex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var __phi62 int32
	_ = __phi62
	var v63 int32
	_ = v63
	var __phi63 int32
	_ = __phi63
	var v65 int32
	_ = v65
	var __phi65 int32
	_ = __phi65
	var v66 int32
	_ = v66
	var __phi66 int32
	_ = __phi66
	var v67 int32
	_ = v67
	var __phi67 int32
	_ = __phi67
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = F_text_to_cstring(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = v5 + int32(11)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v170 = F_cstring_to_text(m, v15)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	v54 = F_toupper(m, v25)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v54)
	v56 = int32(1)
	v58 = v5 + int32(12)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v59 != 0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v23 = v12
	v25 = v22
	goto L9
L7:
	;
	goto L8
L8:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v50
	goto L4
L9:
	;
	if base.Ui32(int32(229)) < base.Ui32((v25|int32(32)-int32(123))&int32(255)) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v39 != 0 {
		v23 = v23 + int32(1)
		v25 = v39
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v160)
	goto L4
L14:
	;
	__phi62 = v23
	__phi63 = v59
	__phi65 = v58
	__phi66 = v56
	__phi67 = v23 + int32(1)
	v62 = __phi62
	v63 = __phi63
	v65 = __phi65
	v66 = __phi66
	v67 = __phi67
	goto L17
L15:
	;
	v142 = v58
	v143 = v56
	goto L16
L16:
	;
	v148 = int32(4) - v143
	if v148 != 0 {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	if base.Ui32(int32(25)) < base.Ui32((v63|int32(32)-int32(97))&int32(255)) {
		v125 = v65
		v126 = v66
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if int32(3) < v126 {
		v155 = v125
		goto L13
	} else {
		goto L37
	}
L19:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v130 != 0 {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v80 = F_toupper(m, v63&int32(255))
	mBase = m.M
	v81 = base.I32_extend8_s(v80)
	v85 = base.B2i32(base.Ui32(int32(25)) < base.Ui32(v81-int32(65)))
	if v85 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v95 = F_toupper(m, v94)
	mBase = m.M
	v96 = base.I32_extend8_s(v95)
	if base.Ui32(v96-int32(65)) <= base.Ui32(int32(25)) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_soundex[0]))))
	v91 = v90
	goto L21
L23:
	;
	goto L24
L24:
	;
	v91 = v80
	goto L21
L25:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_soundex[0]))))
	v104 = v103
	goto L27
L26:
	;
	v104 = v95
	goto L27
L27:
	;
	if v91&int32(255) == v104&int32(255) {
		v125 = v65
		v126 = v66
		goto L19
	} else {
		goto L28
	}
L28:
	;
	if v85 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_soundex[0]))))
	v113 = v112
	goto L31
L30:
	;
	v113 = v80
	goto L31
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v113)
	if v113&int32(255) == int32(48) {
		v125 = v65
		v126 = v66
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v119 = int32(1)
	v125 = v65 + v119
	v126 = v66 + v119
	goto L19
L33:
	;
	if v126 < int32(4) {
		__phi62 = v67
		__phi63 = v130
		__phi65 = v125
		__phi66 = v126
		__phi67 = v67 + int32(1)
		v62 = __phi62
		v63 = __phi63
		v65 = __phi65
		v66 = __phi66
		v67 = __phi67
		goto L17
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L18
L36:
	;
	goto L35
L37:
	;
	v142 = v125
	v143 = v126
	goto L16
L38:
	;
	base.MemoryFill(m, v142, int32(48), v148)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v155 = v148 + v142
	goto L13
L41:
	;
	m.G0 = v5 + int32(16)
	return v170
}
func F_spgbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 float64
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v14 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			v20 = F_SpGistNewBuffer(m, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_SpGistNewBuffer(m, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_SpGistNewBuffer(m, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(_a_F_spgbuild_0)
						v28 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_spgbuild[0])) = v28 + int32(1)
						if v20 < int32(0) {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[1]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v35+(v20^int32(-1))<<(uint(int32(2))%32))))
							v49 = v41
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[2]))
							v49 = v43 + v20<<(uint(int32(13))%32) + int32(-8192)
						}
						F_PageInit(m, v49, int32(_a_F_spgbuild_1), int32(8))
						mBase = m.M
						v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+16)))
						v55 = v49 + v54
						v56 = int32(_a_F_spgbuild_2)
						*(*uint16)(unsafe.Add(mBase, uint32(v55)+6)) = uint16(v56)
						v58 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v55))) = uint16(v58)
						v60 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v49)+80)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+72)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+64)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+56)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+48)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v60
						*(*int32)(unsafe.Add(mBase, uint32(v49)+88)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = int64(-1173640210)
						v78 = int32(92)
						*(*uint16)(unsafe.Add(mBase, uint32(v49)+12)) = uint16(v78)
						v80 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v49)+84)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+68)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+60)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+52)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+44)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v49)+36)) = v80
						F_MarkBufferDirty(m, v20)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = int32(4)
							if v22 < int32(0) {
								v100 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[1]))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v100+(v22^int32(-1))<<(uint(int32(2))%32))))
								v114 = v106
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[2]))
								v114 = v108 + v22<<(uint(int32(13))%32) + int32(-8192)
							}
							F_PageInit(m, v114, int32(_a_F_spgbuild_1), int32(8))
							mBase = m.M
							v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+16)))
							v119 = v114 + v118
							v120 = int32(_a_F_spgbuild_2)
							*(*uint16)(unsafe.Add(mBase, uint32(v119)+6)) = uint16(v120)
							*(*uint16)(unsafe.Add(mBase, uint32(v119))) = uint16(v96)
							F_MarkBufferDirty(m, v22)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								v125 = int32(12)
								if v24 < int32(0) {
									v129 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[1]))
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+(v24^int32(-1))<<(uint(int32(2))%32))))
									v143 = v135
								} else {
									v137 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[2]))
									v143 = v137 + v24<<(uint(int32(13))%32) + int32(-8192)
								}
								F_PageInit(m, v143, int32(_a_F_spgbuild_1), int32(8))
								mBase = m.M
								v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+16)))
								v148 = v143 + v147
								v149 = int32(_a_F_spgbuild_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v148)+6)) = uint16(v149)
								*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v125)
								F_MarkBufferDirty(m, v24)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									v154 = int32(_a_F_spgbuild_0)
									v156 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[0]))
									*(*int32)(unsafe.Add(mBase, _c_F_spgbuild[0])) = v156 - int32(1)
									F_UnlockReleaseBuffer(m, v20)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int32(0)
									} else {
										F_UnlockReleaseBuffer(m, v22)
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return int32(0)
										} else {
											F_UnlockReleaseBuffer(m, v24)
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int32(0)
											} else {
												v167 = v11 + int32(8)
												F_initSpGistState(m, v167, l1)
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = int64(0)
													v172 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v172)
													v175 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[3]))
													v180 = F_AllocSetContextCreateInternal(m, v175, int32(_a_F_spgbuild_3), int32(0), int32(_a_F_spgbuild_1), int32(_a_F_spgbuild_4))
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v180
														v183 = int32(1)
														v184 = int32(0)
														v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
														v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+140))
														v192 = m.T0[v191].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v183, v184, v183, v184, int32(-1), int32(245), v167, v184)
														mBase = m.M
														v193 = m.ExcPending
														if v193 != 0 {
															return int32(0)
														} else {
															v194 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
															F_MemoryContextDelete(m, v194)
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																F_SpGistUpdateMetaPage(m, l1)
																mBase = m.M
																v198 = m.ExcPending
																if v198 != 0 {
																	return int32(0)
																} else {
																	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
																	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+118)))
																	if v200 != int32(112) {
																		v217 = F_palloc0(m, int32(16))
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return int32(0)
																		} else {
																			*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																			v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																			*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																			m.G0 = v11 + int32(112)
																			return v217
																		}
																	} else {
																		v204 = *(*int32)(unsafe.Add(mBase, _c_F_spgbuild[4]))
																		if v204 <= int32(0) {
																			v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
																			if v207 != 0 {
																				v217 = F_palloc0(m, int32(16))
																				mBase = m.M
																				v218 = m.ExcPending
																				if v218 != 0 {
																					return int32(0)
																				} else {
																					*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																					v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																					*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																					m.G0 = v11 + int32(112)
																					return v217
																				}
																			} else {
																				v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
																				if v208 != 0 {
																					v217 = F_palloc0(m, int32(16))
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																						v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																						*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																						m.G0 = v11 + int32(112)
																						return v217
																					}
																				} else {
																					v209 = int32(0)
																					v211 = F_RelationGetNumberOfBlocksInFork(m, l1, v209)
																					mBase = m.M
																					v212 = m.ExcPending
																					if v212 != 0 {
																						return int32(0)
																					} else {
																						F_log_newpage_range(m, l1, v209, v211, int32(1))
																						mBase = m.M
																						v215 = m.ExcPending
																						if v215 != 0 {
																							return int32(0)
																						} else {
																							v217 = F_palloc0(m, int32(16))
																							mBase = m.M
																							v218 = m.ExcPending
																							if v218 != 0 {
																								return int32(0)
																							} else {
																								*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																								v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																								*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																								m.G0 = v11 + int32(112)
																								return v217
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v209 = int32(0)
																			v211 = F_RelationGetNumberOfBlocksInFork(m, l1, v209)
																			mBase = m.M
																			v212 = m.ExcPending
																			if v212 != 0 {
																				return int32(0)
																			} else {
																				F_log_newpage_range(m, l1, v209, v211, int32(1))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					v217 = F_palloc0(m, int32(16))
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						*(*float64)(unsafe.Add(mBase, uint32(v217))) = v192
																						v220 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																						*(*float64)(unsafe.Add(mBase, uint32(v217)+8)) = base.F64_convert_i64_s(v220)
																						m.G0 = v11 + int32(112)
																						return v217
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
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v230 = m.ExcPending
			if v230 != 0 {
				return int32(0)
			} else {
				v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v231 + int32(4)
				F_errmsg_internal(m, int32(_a_F_spgbuild_5), v11)
				mBase = m.M
				v237 = m.ExcPending
				if v237 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_spgbuild_6), int32(84), int32(_a_F_spgbuild_7))
					mBase = m.M
					v242 = m.ExcPending
					if v242 != 0 {
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
func F_spgcanreturn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	if l1 <= int32(1) {
		v5 = F_spgGetCache(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+12)))
			v11 = v9
			return v11 & int32(1)
		}
	} else {
		v11 = int32(1)
		return v11 & int32(1)
	}
}
func F_spgendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+88))
	F_MemoryContextDelete(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+92))
	F_MemoryContextDelete(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+104))
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_pfree(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
	if v14 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+72))
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v14 == v18 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_FreeTupleDesc(m, v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	F_pfree(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v25 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v4)+120))
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+124))
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v4)+188))
	F_pfree(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v4)+192))
	F_pfree(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_pfree(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	F_pfree(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	return
}
func F_spggettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	if l1 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+208)) = uint8(v13)
	v16 = v12 + int32(3488)
	v18 = v12 + int32(_a_F_spggettuple_0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	v21 = v19
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L33
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	if v29 < v21 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return base.B2i32(v29 < v21)
L6:
	;
	v33 = v12 + v29*int32(6)
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+228)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+224))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v38+int32(2672)))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v42)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16+v44<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if int32(0) < v50 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if v74 <= int32(0) {
		v105 = v21
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18+v54<<(uint(int32(2))%32))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v12+int32(3080)))))
	F_index_store_float8_orderby_distances(m, l0, v53, v58, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+220)) = v68 + int32(1)
	return base.B2i32(v29 < v21)
L12:
	;
	return int32(0)
L13:
	;
	goto L11
L14:
	;
	if v105 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v77 = int32(0)
	if v21 <= v77 {
		v105 = v21
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v81 = v77
	v85 = v21
	goto L17
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v18+v81<<(uint(int32(2))%32))))
	if v92 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v105 = v96
	goto L14
L19:
	;
	F_pfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	v96 = v85
	goto L21
L21:
	;
	v98 = v81 + int32(1)
	if v98 < v96 {
		v81 = v98
		v85 = v96
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	v96 = v95
	goto L21
L23:
	;
	goto L18
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+216)) = int64(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_spgWalk(m, v147, v12, int32(0), int32(258))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L31
	}
L25:
	;
	v111 = int32(0)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+208)))
	if v112&int32(1) == v111 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v118 = v111
	goto L27
L27:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v16+v118<<(uint(int32(2))%32))))
	F_pfree(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	v133 = v118 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	if v133 < v134 {
		v118 = v133
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	if v152 != 0 {
		v21 = v152
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L5
L33:
	;
	F_errmsg_internal(m, int32(_a_F_spggettuple_1), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_spggettuple_2), int32(1031), int32(_a_F_spggettuple_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_split_pathtarget_at_srfs_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	v6 = l5
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(96)
	m.G0 = v20
	if l1 == l2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v20 + int32(96)
	return
L2:
	;
	if v149 != 0 {
		goto L47
	} else {
		goto L48
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v205
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = l1
	v28 = F_list_make1_impl(m, int32(1), v20+int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+60)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = l0
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v28
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v31
	v36 = F_list_make1_impl(m, int32(471), v20)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v205 = v36
	goto L3
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v42 = v41
	goto L12
L11:
	;
	v42 = int32(0)
	goto L12
L12:
	;
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v43
	v51 = F_list_make1_impl(m, int32(1), v20+int32(24))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v53
	v61 = F_list_make1_impl(m, int32(1), v20+int32(20))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v63 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v63
	v71 = F_list_make1_impl(m, int32(1), v20+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v71
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v76 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v77 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l1
	v176 = F_list_make1_impl(m, int32(1), v20+int32(12))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L44
	}
L19:
	;
	if v142 != 0 {
		goto L2
	} else {
		goto L43
	}
L20:
	;
	v142 = int32(0)
	v149 = v7
	goto L19
L21:
	;
	goto L22
L22:
	;
	v87 = int32(0)
	v91 = v7
	v94 = v7
	goto L23
L23:
	;
	v100 = v91 << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100+v101)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v105 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v142 = v130
	v149 = v132
	goto L19
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100+v105)))
	v108 = v107
	goto L27
L26:
	;
	v108 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v108
	v114 = F_split_pathtarget_walker(m, v103, v20+int32(56))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v117 = base.B2i32(v116 < v87)
	if v116 < v87 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v130 = v87
	v132 = v94
	goto L31
L31:
	;
	v134 = v91 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v134 < v135 {
		v87 = v130
		v91 = v134
		v94 = v132
		goto L23
	} else {
		goto L42
	}
L32:
	;
	v118 = v87
	goto L34
L33:
	;
	v118 = v116
	goto L34
L34:
	;
	v120 = base.B2i32(v116 <= v87) & v94
	if v116 < v87 {
		v129 = v120
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v130 = v118
	v132 = v129
	goto L31
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	switch v121 - int32(15) {
	case 0:
		goto L39
	default:
		goto L37
	case 2:
		goto L38
	}
L37:
	;
	v129 = int32(1)
	goto L35
L38:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+16)))
	if v127 != 0 {
		v129 = v120
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+12)))
	if v124 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v129 = v120
	goto L35
L41:
	;
	goto L37
L42:
	;
	goto L24
L43:
	;
	goto L18
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v176
	v179 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v179
	v186 = F_list_make1_impl(m, int32(471), v20+int32(8))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v205 = v186
	goto L3
L46:
	;
	v244 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v244
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v263 = v7
	v266 = v7
	goto L55
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v209 = F_lappend(m, v207, int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v223 = v142 << (uint(int32(2)) % 32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v226 = v223 + v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v229 = F_list_concat(m, v227, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L53
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v209
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v214 = F_lappend(m, v212, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v214
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v219 = F_lappend(m, v217, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v219
	v243 = v219
	goto L46
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = v229
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v234 = v233 + v223
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v237 = F_list_concat(m, v235, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v243 = v240
	goto L46
L55:
	;
	v267 = int32(0)
	if v249 == v267 {
		v277 = v267
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v278 = int32(0)
	if v248 == v278 {
		v287 = v278
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v271 <= v263 {
		v277 = int32(0)
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v277 = v273 + v263<<(uint(int32(2))%32)
	goto L57
L60:
	;
	if v243 == int32(0) {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v281 <= v263 {
		v287 = v278
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v287 = v283 + v263<<(uint(int32(2))%32)
	goto L60
L63:
	;
	v290 = int32(0)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if base.B2i32(v287 == v290)|(base.B2i32(v277 == v290)|base.B2i32(v294 <= v263)) != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	if v298 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if base.Ui32(v277+int32(4)) < base.Ui32(v305+v306<<(uint(int32(2))%32)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v312 = F_palloc0(m, int32(40))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L7
	} else {
		goto L69
	}
L67:
	;
	v608 = l1
	goto L68
L68:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v617 = F_lappend(m, v616, v608)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L7
	} else {
		goto L120
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = int32(277)
	if v301 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	if v367 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L71:
	;
	v318 = int32(0)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v319 <= v318 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v324 = v318
	goto L73
L73:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v324<<(uint(int32(2))%32))))
	F_add_sp_item_to_pathtarget(m, v312, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L7
	} else {
		goto L75
	}
L74:
	;
	goto L70
L75:
	;
	v347 = v324 + int32(1)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v347 < v348 {
		v324 = v347
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	if v479 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L78:
	;
	v371 = v287 + int32(4)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v378 = base.B2i32(base.Ui32(v371) < base.Ui32(v373+v374<<(uint(int32(2))%32)))
	if base.Ui32(v371) < base.Ui32(v373+v374<<(uint(int32(2))%32)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v379 = v371
	goto L81
L80:
	;
	v379 = int32(0)
	goto L81
L81:
	;
	if base.Ui32(v371) < base.Ui32(v373+v374<<(uint(int32(2))%32)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v383 = (v379 - v373) >> (uint(int32(2)) % 32)
	goto L84
L83:
	;
	v383 = v374
	goto L84
L84:
	;
	if v374 <= v383 {
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v390 = v383
	goto L86
L86:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402+v390<<(uint(int32(2))%32))))
	if v406 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L77
L88:
	;
	v459 = v390 + int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v459 < v460 {
		v390 = v459
		goto L86
	} else {
		goto L95
	}
L89:
	;
	v409 = int32(0)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v410 <= v409 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v415 = v409
	goto L91
L91:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v430+v415<<(uint(int32(2))%32))))
	F_add_sp_item_to_pathtarget(m, v312, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L7
	} else {
		goto L93
	}
L92:
	;
	goto L88
L93:
	;
	v438 = v415 + int32(1)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v438 < v439 {
		v415 = v438
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	goto L87
L96:
	;
	v597 = F_set_pathtarget_cost_width(m, l0, v312)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L7
	} else {
		goto L119
	}
L97:
	;
	v482 = int32(2)
	v486 = v298 + v263<<(uint(v482)%32) + int32(4)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v479)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	v493 = base.B2i32(base.Ui32(v486) < base.Ui32(v488+v489<<(uint(v482)%32)))
	if base.Ui32(v486) < base.Ui32(v488+v489<<(uint(v482)%32)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v494 = v486
	goto L100
L99:
	;
	v494 = int32(0)
	goto L100
L100:
	;
	if base.Ui32(v486) < base.Ui32(v488+v489<<(uint(v482)%32)) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v498 = (v494 - v488) >> (uint(int32(2)) % 32)
	goto L103
L102:
	;
	v498 = v489
	goto L103
L103:
	;
	if v489 <= v498 {
		goto L96
	} else {
		goto L104
	}
L104:
	;
	v507 = v498
	goto L105
L105:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v479)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517+v507<<(uint(int32(2))%32))))
	if v521 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L96
L107:
	;
	v577 = v507 + int32(1)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v577 < v578 {
		v507 = v577
		goto L105
	} else {
		goto L118
	}
L108:
	;
	v524 = int32(0)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v525 <= v524 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v530 = v524
	goto L110
L110:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v521)+12))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v545+v530<<(uint(int32(2))%32))))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v551 = F_list_member(m, v266, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L7
	} else {
		goto L112
	}
L111:
	;
	goto L107
L112:
	;
	if v551 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_add_sp_item_to_pathtarget(m, v312, v549)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L7
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v556 = v530 + int32(1)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v556 < v557 {
		v530 = v556
		goto L110
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	goto L111
L118:
	;
	goto L106
L119:
	;
	v608 = v312
	goto L68
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v617
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v623 = F_lappend_int(m, v620, base.B2i32(v301 != int32(0)))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v623
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	v263 = v263 + int32(1)
	v266 = v628
	goto L55
}
func F_str_tolower(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v66 = v4
		m.G0 = v10 + int32(16)
		return v66
	} else {
		if l2 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_str_tolower_0)
					F_errmsg(m, int32(_a_F_str_tolower_1), v10)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_str_tolower_2), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_str_tolower_3), int32(1655), int32(_a_F_str_tolower_4))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
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
		} else {
			v16 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
				if v20 == int32(1) {
					v23 = F_pnstrdup(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
						if v25 == int32(0) {
							v66 = v23
						} else {
							v28 = v23
							v30 = v25
							for {
								v35 = int32(255)
								v36 = v30 & v35
								if base.Ui32((v36-int32(65))&v35) < base.Ui32(int32(26)) {
									v45 = v36 | int32(32)
								} else {
									v45 = v36
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v45)
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
								if v47 != 0 {
									v28 = v28 + int32(1)
									v30 = v47
									continue
								} else {
									break
								}
								break
							}
							v66 = v23
						}
						m.G0 = v10 + int32(16)
						return v66
					}
				} else {
					v51 = l1 + int32(1)
					v52 = F_palloc(m, v51)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = F_pg_strlower(m, v52, v51, l0, l1, v16)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v57 = v54 + int32(1)
							if base.Ui32(v57) <= base.Ui32(v51) {
								v66 = v52
								m.G0 = v10 + int32(16)
								return v66
							} else {
								v59 = F_repalloc(m, v52, v57)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = F_pg_strlower(m, v59, v57, l0, l1, v16)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v66 = v59
										m.G0 = v10 + int32(16)
										return v66
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
func F_strict_word_similarity(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14019(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_strict_word_similarity_op(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14023(m, l0, int32(_a_F_strict_word_similarity_op_0), int32(3))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_strip_implicit_coercions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	v2 = l0
	goto L5
L3:
	;
	goto L4
L4:
	;
	v40 = int32(0)
	goto L1
L5:
	;
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	switch v3 - int32(15) {
	case 0:
		goto L13
	default:
		v40 = v2
		goto L1
	case 12:
		goto L12
	case 13:
		goto L11
	case 14:
		goto L10
	case 15:
		goto L9
	case 40:
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != 0 {
		v2 = v37
		goto L5
	} else {
		goto L20
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	if v31 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
	if v26 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L18
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	if v21 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	if v16 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	if v11 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	if v6 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v2)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v36 = v10
	goto L7
L15:
	;
	v36 = v2 + int32(4)
	goto L7
L16:
	;
	v36 = v2 + int32(4)
	goto L7
L17:
	;
	v36 = v2 + int32(4)
	goto L7
L18:
	;
	v36 = v2 + int32(4)
	goto L7
L19:
	;
	v36 = v2 + int32(4)
	goto L7
L20:
	;
	goto L6
}
func F_strlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	if l0&int32(3) == int32(0) {
		v26 = l0
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v51 - l0
L2:
	;
	v30 = v26
	goto L11
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 == int32(0) {
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
	v15 = l0
	goto L7
L7:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v51 = v19
	goto L1
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v45 = v30
	goto L14
L13:
	;
	goto L12
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v51 = v45
	goto L1
L16:
	;
	goto L15
}
func F_strpbrk(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v2 = int32(_a_F_strpbrk_0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*int8)(unsafe.Add(mBase, _c_F_strpbrk[0])))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v69 = v61 - l0 + l0
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v71 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	m.G0 = v8 + int32(32)
	goto L1
L3:
	;
	F___memset(m, v8, int32(0), int32(32))
	mBase = m.M
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strpbrk[0])))
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strpbrk[1])))
	if v11 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v12 = F___strchrnul(m, l0, v10)
	mBase = m.M
	v61 = v12
	goto L2
L7:
	;
	goto L6
L8:
	;
	v18 = v2
	v19 = v16
	goto L11
L9:
	;
	goto L10
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v40 == int32(0) {
		v61 = l0
		goto L2
	} else {
		goto L14
	}
L11:
	;
	v26 = v8 + int32(base.Ui32(v19)>>(uint(int32(3))%32))&int32(28)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27 | v28<<(uint(v19)%32)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v32 != 0 {
		v18 = v18 + v28
		v19 = v32
		goto L11
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	goto L12
L14:
	;
	v44 = l0
	v45 = v40
	goto L15
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v45)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v53)>>(uint(v45)%32))&int32(1) != 0 {
		v61 = v44
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v61 = v59
	goto L2
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v59 = v44 + int32(1)
	if v57 != 0 {
		v44 = v59
		v45 = v57
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v72 = v69
	goto L21
L20:
	;
	v72 = int32(0)
	goto L21
L21:
	;
	return v72
}
func F_strtod(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v145 int64
	_ = v145
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_strtox_1(m, v7, l0, l1, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return float64(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v23 = m.G0
		v25 = v23 - int32(32)
		m.G0 = v25
		v28 = v15 & int64(281474976710655)
		v32 = int64(base.Ui64(v15)>>(uint(int64(48))%64)) & int64(32767)
		v33 = base.I32_wrap_i64(v32)
		if base.Ui32(v33-int32(_a_F_strtod_0)) <= base.Ui32(int32(2045)) {
			v42 = v28<<(uint(int64(4))%64) | int64(base.Ui64(v14)>>(uint(int64(60))%64))
			v47 = v14 & int64(1152921504606846975)
			if base.Ui64(int64(576460752303423489)) <= base.Ui64(v47) {
				v57 = v42 + int64(1)
			} else {
				if v47 != int64(576460752303423488) {
					v57 = v42
				} else {
					v57 = v42&int64(1) + v42
				}
			}
			v60 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v57))
			if base.Ui64(int64(4503599627370495)) < base.Ui64(v57) {
				v61 = int64(0)
			} else {
				v61 = v57
			}
			v138 = v61
			v145 = base.I64_extend_i32_u(v60) + base.I64_extend_i32_u(v33-int32(_a_F_strtod_1))
		} else {
			if base.B2i32(v14|v28 == int64(0))|base.B2i32(v32 != int64(32767)) == int32(0) {
				v138 = v28<<(uint(int64(4))%64) | int64(base.Ui64(v14)>>(uint(int64(60))%64)) | int64(2251799813685248)
				v145 = int64(2047)
			} else {
				if base.Ui32(int32(_a_F_strtod_2)) < base.Ui32(v33) {
					v138 = int64(0)
					v145 = int64(2047)
				} else {
					v87 = base.B2i32(v32 == int64(0))
					if v32 == int64(0) {
						v88 = int32(_a_F_strtod_1)
					} else {
						v88 = int32(_a_F_strtod_0)
					}
					v89 = v88 - v33
					if int32(112) < v89 {
						v92 = int64(0)
						v138 = v92
						v145 = v92
					} else {
						if v32 == int64(0) {
							v96 = v28
						} else {
							v96 = v28 | int64(281474976710656)
						}
						if v33 != v88 {
							F___ashlti3(m, v25+int32(16), v14, v96, int32(128)-v89)
							mBase = m.M
							v104 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
							v109 = base.B2i32(v104|v105 != int64(0))
						} else {
							v109 = int32(0)
						}
						F___lshrti3(m, v25, v14, v96, v89)
						mBase = m.M
						v111 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
						v117 = v111<<(uint(int64(4))%64) | int64(base.Ui64(v114)>>(uint(int64(60))%64))
						v121 = base.I64_extend_i32_u(v109) | v114&int64(1152921504606846975)
						if base.Ui64(int64(576460752303423489)) <= base.Ui64(v121) {
							v131 = v117 + int64(1)
						} else {
							if v121 != int64(576460752303423488) {
								v131 = v117
							} else {
								v131 = v117&int64(1) + v117
							}
						}
						v135 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v131))
						if base.Ui64(int64(4503599627370495)) < base.Ui64(v131) {
							v136 = v131 ^ int64(4503599627370496)
						} else {
							v136 = v131
						}
						v138 = v136
						v145 = base.I64_extend_i32_u(v135)
					}
				}
			}
		}
		m.G0 = v25 + int32(32)
		m.G0 = v7 + int32(16)
		return base.F64_reinterpret_i64(v15&int64(-9223372036854775807-1) | v145<<(uint(int64(52))%64) | v138)
	}
}
func F_strtok_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	if l0 != 0 {
		v7 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(_a_F_strtok_r_0)
	v12 = m.G0
	v14 = v12 - int32(32)
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v15
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strtok_r[0])))
	if v23 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v4 != 0 {
		v7 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(0)
L4:
	;
	v92 = v91 + v7
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v91 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strtok_r[1])))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v7
	goto L11
L9:
	;
	goto L10
L10:
	;
	v41 = v8
	v42 = v23
	goto L14
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v37 == v23 {
		v31 = v31 + int32(1)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v91 = v31 - v7
	goto L4
L13:
	;
	goto L12
L14:
	;
	v49 = v14 + int32(base.Ui32(v42)>>(uint(int32(3))%32))&int32(28)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v51 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50 | v51<<(uint(v42)%32)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v55 != 0 {
		v41 = v41 + v51
		v42 = v55
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v58 == int32(0) {
		v81 = v7
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v91 = v81 - v7
	goto L4
L18:
	;
	v62 = v7
	v63 = v58
	goto L19
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(base.Ui32(v63)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v71)>>(uint(v63)%32))&int32(1) == int32(0) {
		v81 = v62
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v81 = v79
	goto L17
L21:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v79 = v62 + int32(1)
	if v77 != 0 {
		v62 = v79
		v63 = v77
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v96
	return v96
L24:
	;
	goto L25
L25:
	;
	v100 = int32(_a_F_strtok_r_0)
	v104 = m.G0
	v106 = v104 - int32(32)
	m.G0 = v106
	v108 = int32(*(*int8)(unsafe.Add(mBase, _c_F_strtok_r[0])))
	if v108 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v167 = v159 - v92 + v92
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v168 != 0 {
		goto L44
	} else {
		goto L45
	}
L27:
	;
	m.G0 = v106 + int32(32)
	goto L26
L28:
	;
	F___memset(m, v106, int32(0), int32(32))
	mBase = m.M
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strtok_r[0])))
	if v114 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_strtok_r[1])))
	if v109 != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v110 = F___strchrnul(m, v92, v108)
	mBase = m.M
	v159 = v110
	goto L27
L32:
	;
	goto L31
L33:
	;
	v116 = v100
	v117 = v114
	goto L36
L34:
	;
	goto L35
L35:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v138 == int32(0) {
		v159 = v92
		goto L27
	} else {
		goto L39
	}
L36:
	;
	v124 = v106 + int32(base.Ui32(v117)>>(uint(int32(3))%32))&int32(28)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v125 | v126<<(uint(v117)%32)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v130 != 0 {
		v116 = v116 + v126
		v117 = v130
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	goto L37
L39:
	;
	v142 = v92
	v143 = v138
	goto L40
L40:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v106+int32(base.Ui32(v143)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v151)>>(uint(v143)%32))&int32(1) != 0 {
		v159 = v142
		goto L27
	} else {
		goto L42
	}
L41:
	;
	v159 = v157
	goto L27
L42:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	v157 = v142 + int32(1)
	if v155 != 0 {
		v142 = v157
		v143 = v155
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v167 + int32(1)
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v172)
	return v92
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return v92
}
func F_strtoll(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(-9223372036854775807-1))
	return v5
}
func F_strtox_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v7 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(-1)
	v18 = v11 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v7
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = base.I64_extend_i32_s(v22 - v23)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v29-v23) <= v7) != 0 {
		v36 = v29
	} else {
		v36 = v23 + base.I32_wrap_i64(v7)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v36
	F___floatscan(m, v11, v18, l3, int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return
	} else {
		v41 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if l2 != 0 {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+136))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v43 + (l1 + (v44 - v45))
		} else {
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v41
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v42
		m.G0 = v11 + int32(160)
		return
	}
}
func F_subltree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_inner_subltree(m, v5, v9, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v13 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					return v11
				}
			} else {
				return v11
			}
		}
	}
}
func F_substitute_grouped_columns_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
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
	var v168 int32
	_ = v168
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	v3 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(32)
	return v862
L2:
	;
	v862 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v27 - int32(9) {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L6
	}
L5:
	;
	v852 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L20
	} else {
		goto L187
	}
L6:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v39 != int32(1) {
		v168 = v27
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v35 <= v34 {
		v862 = l0
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v30 == v31 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v31 < v30 {
		v862 = l0
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	goto L6
L12:
	;
	switch v168 - int32(6) {
	case 0:
		goto L43
	case 1, 2:
		v862 = l0
		goto L1
	default:
		goto L42
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v42 != 0 {
		v168 = v27
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v43 == int32(0) {
		v168 = v27
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v50 = int32(0)
	goto L17
L16:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v168 = v164
	goto L12
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v66 <= v50 {
		goto L16
	} else {
		goto L19
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+56))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v88 = v85 + v71<<(uint(int32(5))%32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(32))))
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88-int32(28)))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(24))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(20))))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(16))))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v105 = F_makeVar(m, v91, v94, v97, v100, v103, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L23
	}
L19:
	;
	v71 = v50 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v50<<(uint(int32(2))%32)+v72)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v76 = F_equal(m, l0, v75)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v76 == int32(0) {
		v50 = v71
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+36)) = v109
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+40)) = uint16(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+108))
	if v116 == int32(0) {
		v862 = v105
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v120 = int32(0)
	if v119 == v120 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v158 != 0 {
		v862 = v105
		goto L1
	} else {
		goto L38
	}
L26:
	;
	v158 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v126 <= int32(0) {
		v152 = v120
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v158 = v152
	goto L25
L30:
	;
	v129 = int32(0)
	if v129 < v126 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v132 = v126
	goto L33
L32:
	;
	v132 = v129
	goto L33
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v135 = int32(0)
	goto L34
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133+v135<<(uint(int32(2))%32))))
	v144 = base.B2i32(v143 == v82)
	if v143 == v82 {
		v152 = v144
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v152 = v144
	goto L29
L36:
	;
	v146 = v135 + int32(1)
	if v146 != v132 {
		v135 = v146
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v161 = F_bms_add_member(m, v159, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v161
	v862 = v105
	goto L1
L40:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v840 + int32(1)
	v846 = F_query_tree_mutator_impl(m, l0, int32(481), l1, int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L20
	} else {
		goto L186
	}
L41:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)+56))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)+16))
	v764 = v761 + v226<<(uint(int32(5))%32)
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v764-int32(32))))
	v770 = int32(*(*int16)(unsafe.Add(mBase, uint32(v764-int32(28)))))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v764-int32(24))))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v764-int32(20))))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v764-int32(16))))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v781 = F_makeVar(m, v767, v770, v773, v776, v779, v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L20
	} else {
		goto L168
	}
L42:
	;
	if v168 == int32(67) {
		goto L40
	} else {
		goto L166
	}
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v186 != v187 {
		v862 = l0
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v186 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v267 = int32(0)
	if v265 == v267 {
		goto L64
	} else {
		goto L65
	}
L46:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v191&int32(1) != 0 {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v194 == int32(0) {
		goto L45
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v197 <= int32(0) {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v200 = int32(0)
	if v200 < v197 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v204 = v197
	goto L54
L53:
	;
	v204 = v200
	goto L54
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v209 = v200
	goto L55
L55:
	;
	v226 = v209 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v205+v209<<(uint(int32(2))%32))))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	if v232 != int32(6) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L45
L57:
	;
	if v226 != v204 {
		v209 = v226
		goto L55
	} else {
		goto L62
	}
L58:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v235 != v236 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+8)))
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v238 != v239 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	if v241 == int32(0) {
		goto L41
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	goto L56
L63:
	;
	if v305 != 0 {
		v862 = l0
		goto L1
	} else {
		goto L76
	}
L64:
	;
	v305 = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if v273 <= int32(0) {
		v299 = v267
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v305 = v299
	goto L63
L68:
	;
	v276 = int32(0)
	if v276 < v273 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v279 = v273
	goto L71
L70:
	;
	v279 = v276
	goto L71
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v282 = int32(0)
	goto L72
L72:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v280+v282<<(uint(int32(2))%32))))
	v291 = base.B2i32(v290 == v266)
	if v290 == v266 {
		v299 = v291
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v299 = v291
	goto L67
L74:
	;
	v293 = v282 + int32(1)
	if v293 != v279 {
		v282 = v293
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v308+v309<<(uint(int32(2))%32)-int32(4))))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	if v316 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v701 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v702 = F_get_rte_attribute_name(m, v315, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L20
	} else {
		goto L150
	}
L78:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v315)+16))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v321 = v319 + int32(148)
	v322 = m.G0
	v324 = v322 - int32(16)
	m.G0 = v324
	v326 = m.G0
	v328 = v326 + int32(-64)
	m.G0 = v328
	v331 = v324 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = int32(0)
	v336 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L20
	} else {
		goto L81
	}
L79:
	;
	m.G0 = v324 + int32(16)
	if v668 == int32(0) {
		goto L77
	} else {
		goto L148
	}
L80:
	;
	if v476 == int32(0) {
		v668 = v3
		goto L79
	} else {
		goto L120
	}
L81:
	;
	v339 = v326 + int32(-48)
	F_ScanKeyInit(m, v339, int32(9), int32(3), int32(184), v317)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L20
	} else {
		goto L82
	}
L82:
	;
	v346 = int32(1)
	v349 = F_systable_beginscan(m, v336, int32(2665), v346, int32(0), v346, v339)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L20
	} else {
		goto L86
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L20
	} else {
		goto L117
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L20
	} else {
		goto L114
	}
L85:
	;
	F_systable_endscan(m, v349)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L20
	} else {
		goto L112
	}
L86:
	;
	v351 = F_systable_getnext(m, v349)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	if v351 == int32(0) {
		v476 = v3
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v360 = v351
	goto L89
L89:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+22)))
	v376 = v374 + v375
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+72)))
	if v377 == int32(112) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v476 = v3
	goto L85
L91:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+73)))
	if v380&int32(1) != 0 {
		v476 = v3
		goto L85
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v462 = F_systable_getnext(m, v349)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L20
	} else {
		goto L110
	}
L94:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v336)+52))
	v386 = F_heap_getattr_3(m, v360, v383, v326+int32(-49))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L20
	} else {
		goto L95
	}
L95:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+15)))
	if v388 == int32(1) {
		goto L84
	} else {
		goto L96
	}
L96:
	;
	v391 = F_pg_detoast_datum(m, v386)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L20
	} else {
		goto L97
	}
L97:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v393 != int32(1) {
		goto L83
	} else {
		goto L98
	}
L98:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	if v396 < int32(0) {
		goto L83
	} else {
		goto L99
	}
L99:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	if v399 != 0 {
		goto L83
	} else {
		goto L100
	}
L100:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	if v400 != int32(21) {
		goto L83
	} else {
		goto L101
	}
L101:
	;
	v403 = int32(0)
	if v396 == v403 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+22)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v457+v458)))
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v460
	v476 = v450
	goto L85
L103:
	;
	v450 = v3
	goto L102
L104:
	;
	goto L105
L105:
	;
	v418 = v403
	v420 = v3
	goto L106
L106:
	;
	v430 = int32(*(*int16)(unsafe.Add(mBase, uint32(v391+int32(24)+v418<<(uint(int32(1))%32)))))
	v433 = F_bms_add_member(m, v420, v430+int32(7))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L20
	} else {
		goto L108
	}
L107:
	;
	v450 = v433
	goto L102
L108:
	;
	v436 = v418 + int32(1)
	if v436 != v396 {
		v418 = v436
		v420 = v433
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	if v462 != 0 {
		v360 = v462
		goto L89
	} else {
		goto L111
	}
L111:
	;
	goto L90
L112:
	;
	F_relation_close(m, v336, int32(1))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L20
	} else {
		goto L113
	}
L113:
	;
	m.G0 = v328 - int32(-64)
	goto L80
L114:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+22)))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v495+v496)))
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = v498
	F_errmsg_internal(m, int32(_a_F_substitute_grouped_columns_mutator_0), v328)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_substitute_grouped_columns_mutator_1), int32(1499), int32(_a_F_substitute_grouped_columns_mutator_2))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L20
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errmsg_internal(m, int32(_a_F_substitute_grouped_columns_mutator_3), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L20
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_substitute_grouped_columns_mutator_1), int32(1506), int32(_a_F_substitute_grouped_columns_mutator_2))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L20
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	if v318 == int32(0) {
		v588 = v3
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v589 = int32(0)
	if v476 == v589 {
		goto L133
	} else {
		goto L134
	}
L122:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v526 <= int32(0) {
		v588 = v3
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v535 = int32(0)
	v548 = v3
	goto L124
L124:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v318)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549+v535<<(uint(int32(2))%32))))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	if v554 != int32(6) {
		v565 = v548
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v588 = v565
	goto L121
L126:
	;
	v567 = v535 + int32(1)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v567 < v568 {
		v535 = v567
		v548 = v565
		goto L124
	} else {
		goto L131
	}
L127:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	if v557 != v309 {
		v565 = v548
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v553)+28))
	if v559 != 0 {
		v565 = v548
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v560 = int32(*(*int16)(unsafe.Add(mBase, uint32(v553)+8)))
	v563 = F_bms_add_member(m, v548, v560+int32(7))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L20
	} else {
		goto L130
	}
L130:
	;
	v565 = v563
	goto L126
L131:
	;
	goto L125
L132:
	;
	if v642 == int32(0) {
		v668 = v3
		goto L79
	} else {
		goto L146
	}
L133:
	;
	v642 = int32(1)
	goto L132
L134:
	;
	goto L135
L135:
	;
	if v588 == int32(0) {
		v635 = v589
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v642 = v635
	goto L132
L137:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	if v599 < v598 {
		v635 = v589
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v601 = int32(1)
	if v598 <= v601 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v604 = v601
	goto L141
L140:
	;
	v604 = v598
	goto L141
L141:
	;
	v605 = int32(8)
	v610 = int32(0)
	goto L142
L142:
	;
	v617 = v610 << (uint(int32(2)) % 32)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v476+v605+v617)))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v588+v605+v617)))
	v624 = v619 & (v621 ^ int32(-1))
	v626 = base.B2i32(v624 == int32(0))
	if v624 != 0 {
		v635 = v626
		goto L136
	} else {
		goto L144
	}
L143:
	;
	v635 = v626
	goto L136
L144:
	;
	v628 = v610 + int32(1)
	if v628 != v604 {
		v610 = v628
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	v647 = F_lappend_oid(m, v645, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L20
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v647
	v668 = int32(1)
	goto L79
L148:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v678 = F_lappend_int(m, v676, v677)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L20
	} else {
		goto L149
	}
L149:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v680))) = v678
	v862 = l0
	goto L1
L150:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L20
	} else {
		goto L151
	}
L151:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L20
	} else {
		goto L152
	}
L152:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v704 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v713
	F_errmsg(m, int32(_a_F_substitute_grouped_columns_mutator_4), v22)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L20
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v713
	F_errmsg(m, int32(_a_F_substitute_grouped_columns_mutator_5), v22+int32(16))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L20
	} else {
		goto L163
	}
L156:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)))
	if v721 == int32(1) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	F_errdetail(m, int32(_a_F_substitute_grouped_columns_mutator_6), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L20
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_parser_errposition(m, v728, v729)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L20
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	F_errfinish(m, int32(_a_F_substitute_grouped_columns_mutator_7), int32(1556), int32(_a_F_substitute_grouped_columns_mutator_8))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L20
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_parser_errposition(m, v744, v745)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_substitute_grouped_columns_mutator_7), int32(1562), int32(_a_F_substitute_grouped_columns_mutator_8))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	v756 = F_expression_tree_mutator_impl(m, l0, int32(481), l1)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L20
	} else {
		goto L167
	}
L167:
	;
	v862 = v756
	goto L1
L168:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v764-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v781)+36)) = v785
	v789 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v764-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v781)+40)) = uint16(v789)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)+108))
	if v792 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v862 = v781
	goto L1
L170:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v796 = int32(0)
	if v795 == v796 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v834 != 0 {
		goto L169
	} else {
		goto L184
	}
L172:
	;
	v834 = int32(0)
	goto L171
L173:
	;
	goto L174
L174:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v802 <= int32(0) {
		v828 = v796
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v834 = v828
	goto L171
L176:
	;
	v805 = int32(0)
	if v805 < v802 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v808 = v802
	goto L179
L178:
	;
	v808 = v805
	goto L179
L179:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v795)+12))
	v811 = int32(0)
	goto L180
L180:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v809+v811<<(uint(int32(2))%32))))
	v820 = base.B2i32(v819 == v758)
	if v819 == v758 {
		v828 = v820
		goto L175
	} else {
		goto L182
	}
L181:
	;
	v828 = v820
	goto L175
L182:
	;
	v822 = v811 + int32(1)
	if v822 != v808 {
		v811 = v822
		goto L180
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v781)+24))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v760)+8))
	v837 = F_bms_add_member(m, v835, v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L20
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v781)+24)) = v837
	goto L169
L186:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v848 - int32(1)
	v862 = v846
	goto L1
L187:
	;
	v854 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v854)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v852)+28))
	v857 = F_substitute_grouped_columns_mutator(m, v856, l1)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L20
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v852)+28)) = v857
	v860 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v860)
	v862 = v852
	goto L1
}
func F_switchToPresortedPrefixMode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
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
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v270 int64
	_ = v270
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v287 int64
	_ = v287
	var v291 int64
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v371 int64
	_ = v371
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v388 int64
	_ = v388
	var v392 int64
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v413 int64
	_ = v413
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v435 int64
	_ = v435
	var v436 int64
	_ = v436
	var v437 int64
	_ = v437
	var v439 int64
	_ = v439
	v7 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v18 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v51 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v25 = int32(1)
	v28 = int32(2)
	v29 = v22 << (uint(v28) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_switchToPresortedPrefixMode[0]))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v44 = F_tuplesort_begin_heap(m, v17, v21-v22, v24+v22<<(uint(v25)%32), v29+v30, v32+v29, v34+v22, v37, int32(0), v39<<(uint(v25)%32)&v28)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_tuplesort_reset(m, v18)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v44
	goto L1
L7:
	;
	goto L1
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v57 = v55 - v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+236))
	if v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	goto L10
L10:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v86 <= int64(0) {
		v200 = v86
		v201 = v7
		goto L23
	} else {
		goto L24
	}
L11:
	;
	goto L10
L12:
	;
	goto L11
L13:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v54)+72)) = uint32(v57)
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+68)) = uint8(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+32))
	if v75 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if int64(1073741823) < v57 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if int64(1073741823) < v57 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+232))
	if v63 != int32(-1) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	goto L13
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	v78 = v77
	goto L22
L21:
	;
	v78 = v74
	goto L22
L22:
	;
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+32)) = v79
	goto L12
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v200 - v201
	if v200 == v201 {
		goto L58
	} else {
		goto L59
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v89 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	m.T0[v190].(func(*base.Module, int32))(m, v188)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L56
	}
L26:
	;
	v134 = int64(1)
	v135 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v135 < int64(2) {
		v200 = v135
		v201 = v134
		goto L23
	} else {
		goto L42
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v107 = int32(0)
	v109 = F_tuplesort_gettupleslot(m, v104, base.B2i32(v15 == int32(1)), v107, v89, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L32
	}
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v92&int32(2) != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_puttupleslot(m, v95, v89)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+32))
	m.T0[v101].(func(*base.Module, int32, int32))(m, v98, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v111 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v125 = F_isCurrentGroup(m, l0, v123, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L39
	}
L34:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	if v112&int32(2) == int32(0) {
		v123 = v111
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	m.T0[v119].(func(*base.Module, int32, int32))(m, v111, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v123 = v122
	goto L33
L39:
	;
	if v125 == int32(0) {
		v187 = v7
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_tuplesort_puttupleslot(m, v129, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L26
L42:
	;
	v147 = v134
	goto L43
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v149 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v152 = F_tuplesort_gettupleslot(m, v148, base.B2i32(v15 == int32(1)), v149, v150, v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L45
	}
L44:
	;
	v200 = v178
	v201 = v177
	goto L23
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v154 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v168 = F_isCurrentGroup(m, l0, v166, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L52
	}
L47:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
	if v155&int32(2) == int32(0) {
		v166 = v154
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+32))
	m.T0[v162].(func(*base.Module, int32, int32))(m, v154, v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v166 = v165
	goto L46
L52:
	;
	if v168 == int32(0) {
		v187 = v147
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_tuplesort_puttupleslot(m, v172, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v177 = v147 + int64(1)
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v177 < v178 {
		v147 = v177
		goto L43
	} else {
		goto L55
	}
L55:
	;
	goto L44
L56:
	;
	v193 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	v200 = v193
	v201 = v187
	goto L23
L57:
	;
	m.G0 = v11 + int32(16)
	return
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+32))
	m.T0[v208].(func(*base.Module, int32, int32))(m, v205, v206)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_performsort(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	m.T0[v215].(func(*base.Module, int32))(m, v213)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	goto L57
L63:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v221 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v432 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L65:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v224 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v339 + int64(1)
	v343 = int32(0)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v338)+128))
	if v347 == v343 {
		goto L97
	} else {
		goto L98
	}
L67:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v227 != int32(1) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_switchToPresortedPrefixMode[1]))
	v235 = v224 + v232*int32(96)
	v237 = v235 + int32(56)
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
	*(*int64)(unsafe.Add(mBase, uint32(v237))) = v238 + int64(1)
	v242 = int32(0)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v230)+128))
	if v246 == v242 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v307 {
	case 0:
		goto L91
	case 1:
		goto L90
	default:
		goto L89
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v284
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v230)+112))
	v291 = base.I64_div_s(v287+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v230)+124))
	switch v293 - int32(3) {
	case 0:
		goto L85
	case 1:
		v304 = v293
		goto L82
	case 2:
		goto L84
	default:
		goto L83
	}
L71:
	;
	if v263&int32(255) != base.B2i32(v246 != int32(0)) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v230)+96))
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v230)+88))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+120)))
	v263 = v252
	v264 = v249 - v250
	goto L71
L73:
	;
	goto L74
L74:
	;
	v253 = F_LogicalTapeSetBlocks(m, v246)
	mBase = m.M
	v255 = v253 << (uint(int64(13)) % 64)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+120)))
	if v257 != 0 {
		v263 = int32(1)
		v264 = v255
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+120)) = uint8(v258)
	*(*int64)(unsafe.Add(mBase, uint32(v230)+112)) = v255
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v230)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+124)) = v261
	v284 = v242
	goto L70
L76:
	;
	v284 = int32(1)
	goto L70
L77:
	;
	if v263&int32(1) != 0 {
		v284 = v242
		goto L70
	} else {
		goto L81
	}
L78:
	;
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v230)+112))
	if v264 <= v270 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+120)) = uint8(v263)
	*(*int64)(unsafe.Add(mBase, uint32(v230)+112)) = v264
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v230)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+124)) = v274
	if v263&int32(1) == int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v284 = v242
	goto L70
L81:
	;
	goto L76
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v304
	goto L69
L83:
	;
	v304 = int32(0)
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(8)
	goto L69
L85:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+69)))
	if v298 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v299 = int32(1)
	goto L88
L87:
	;
	v299 = int32(2)
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v299
	goto L69
L89:
	;
	v333 = v235 + int32(96)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v334 | v335
	goto L64
L90:
	;
	v320 = v235 + int32(88)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v321 + v322
	v326 = v235 + int32(80)
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v326)))
	if v321 <= v327 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	v309 = v235 + int32(72)
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v309)))
	*(*int64)(unsafe.Add(mBase, uint32(v309))) = v310 + v311
	v315 = v235 - int32(-64)
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	if v310 <= v316 {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v310
	goto L89
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v326))) = v321
	goto L89
L94:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v408 {
	case 0:
		goto L116
	case 1:
		goto L115
	default:
		goto L114
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v385
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v338)+112))
	v392 = base.I64_div_s(v388+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v338)+124))
	switch v394 - int32(3) {
	case 0:
		goto L110
	case 1:
		v405 = v394
		goto L107
	case 2:
		goto L109
	default:
		goto L108
	}
L96:
	;
	if v364&int32(255) != base.B2i32(v347 != int32(0)) {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v338)+96))
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v338)+88))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+120)))
	v364 = v353
	v365 = v350 - v351
	goto L96
L98:
	;
	goto L99
L99:
	;
	v354 = F_LogicalTapeSetBlocks(m, v347)
	mBase = m.M
	v356 = v354 << (uint(int64(13)) % 64)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+120)))
	if v358 != 0 {
		v364 = int32(1)
		v365 = v356
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v359 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+120)) = uint8(v359)
	*(*int64)(unsafe.Add(mBase, uint32(v338)+112)) = v356
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v338)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+124)) = v362
	v385 = v343
	goto L95
L101:
	;
	v385 = int32(1)
	goto L95
L102:
	;
	if v364&int32(1) != 0 {
		v385 = v343
		goto L95
	} else {
		goto L106
	}
L103:
	;
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v338)+112))
	if v365 <= v371 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+120)) = uint8(v364)
	*(*int64)(unsafe.Add(mBase, uint32(v338)+112)) = v365
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v338)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+124)) = v375
	if v364&int32(1) == int32(0) {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v385 = v343
	goto L95
L106:
	;
	goto L101
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v405
	goto L94
L108:
	;
	v405 = int32(0)
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(8)
	goto L94
L110:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+69)))
	if v399 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v400 = int32(1)
	goto L113
L112:
	;
	v400 = int32(2)
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v400
	goto L94
L114:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v424 | v425
	goto L64
L115:
	;
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v417 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v416 + v417
	v420 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	if v416 <= v420 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v410 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v409 + v410
	v413 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if v409 <= v413 {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v409
	goto L114
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v416
	goto L114
L119:
	;
	v435 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v436 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v437 = v436 + v201
	if v435 < v437 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(3)
	goto L57
L122:
	;
	v439 = v435
	goto L124
L123:
	;
	v439 = v437
	goto L124
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v439
	goto L121
}
func F_synchronize_slots(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
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
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int64
	_ = v366
	var v369 int64
	_ = v369
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v736 int64
	_ = v736
	var v737 int32
	_ = v737
	var v740 int64
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int64
	_ = v748
	var v750 int64
	_ = v750
	var v751 int64
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v769 int64
	_ = v769
	var v770 int32
	_ = v770
	var v772 int64
	_ = v772
	var v773 int64
	_ = v773
	var v778 int64
	_ = v778
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int64
	_ = v837
	var v838 int64
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v880 int64
	_ = v880
	var v882 int64
	_ = v882
	var v884 int64
	_ = v884
	var v886 int64
	_ = v886
	var v888 int64
	_ = v888
	var v890 int64
	_ = v890
	var v892 int64
	_ = v892
	var v894 int64
	_ = v894
	var v896 int32
	_ = v896
	var v898 int64
	_ = v898
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int64
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int64
	_ = v928
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v940 int64
	_ = v940
	var v943 int64
	_ = v943
	var v945 int64
	_ = v945
	var v948 int32
	_ = v948
	var v950 int64
	_ = v950
	var v951 int64
	_ = v951
	var v952 int64
	_ = v952
	var v953 int32
	_ = v953
	var v954 int64
	_ = v954
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1034 int32
	_ = v1034
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1076 int32
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1078 int64
	_ = v1078
	var v1081 int64
	_ = v1081
	var v1082 int64
	_ = v1082
	var v1085 int64
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1124 int32
	_ = v1124
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(240)
	m.G0 = v19
	v22 = *(*int64)(unsafe.Add(mBase, _c_F_synchronize_slots[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+160)) = v22
	v25 = *(*int64)(unsafe.Add(mBase, _c_F_synchronize_slots[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+152)) = v25
	v28 = *(*int64)(unsafe.Add(mBase, _c_F_synchronize_slots[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+144)) = v28
	v31 = *(*int64)(unsafe.Add(mBase, _c_F_synchronize_slots[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+136)) = v31
	v34 = *(*int64)(unsafe.Add(mBase, _c_F_synchronize_slots[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v34
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[5]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v40 = base.B2i32(v38 == int32(2))
	goto L1
L1:
	;
	if v40 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[6]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+60))
	v54 = m.T0[v53].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, int32(_a_F_synchronize_slots_0), int32(10), v19+int32(128))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L4
L7:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v1132 != 0 {
		goto L321
	} else {
		goto L322
	}
L8:
	;
	F_pfree(m, v1046)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L5
	} else {
		goto L320
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v56 == int32(2) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v61 = F_MakeTupleTableSlot(m, v59, int32(_a_F_synchronize_slots_1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L5
	} else {
		goto L317
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v66 = F_tuplestore_gettupleslot(m, v63, int32(1), int32(0), v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v66 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v75 = v2
	goto L18
L16:
	;
	v397 = v2
	goto L17
L17:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[7]))
	v411 = F_LWLockAcquire(m, v407+int32(_a_F_synchronize_slots_2), int32(1))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L135
	}
L18:
	;
	v85 = F_palloc0(m, int32(48))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v397 = v380
	goto L17
L20:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v87 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_slot_getsomeattrs_int(m, v61, int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v95 = F_text_to_cstring(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v95
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v98 <= int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_slot_getsomeattrs_int(m, v61, int32(2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v106 = F_text_to_cstring(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v106
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v109 <= int32(2) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_slot_getsomeattrs_int(m, v61, int32(3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v121 = int64(0)
	goto L37
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	v121 = v120
	goto L37
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v121
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v123 <= int32(3) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_slot_getsomeattrs_int(m, v61, int32(4))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
	if v130 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v135 = int64(0)
	goto L44
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v133)))
	v135 = v134
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v135
	v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v137 <= int32(4) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_slot_getsomeattrs_int(m, v61, int32(5))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
	if v147 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v148 = int32(0)
	goto L51
L50:
	;
	v148 = v145
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+40)) = v148
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v150 <= int32(5) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_slot_getsomeattrs_int(m, v61, int32(6))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+12)) = uint8(base.B2i32(v157 != int32(0)))
	v161 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v161 <= int32(6) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	F_slot_getsomeattrs_int(m, v61, int32(7))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
	if v168 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v173 = int64(0)
	goto L62
L61:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+24))
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v171)))
	v173 = v172
	goto L62
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v173
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v175 <= int32(7) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_slot_getsomeattrs_int(m, v61, int32(8))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+13)) = uint8(base.B2i32(v182 != int32(0)))
	v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v186 <= int32(8) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	F_slot_getsomeattrs_int(m, v61, int32(9))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+32))
	v194 = F_text_to_cstring(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L5
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v194
	v197 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+6)))
	if v197 <= int32(9) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_slot_getsomeattrs_int(m, v61, int32(10))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v203 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+9)))
	if v205 == v203 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+36))
	v210 = F_text_to_cstring(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	v363 = v203
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+44)) = v363
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
	if v366 == int64(0) {
		goto L125
	} else {
		goto L126
	}
L79:
	;
	v213 = int32(_a_F_synchronize_slots_3)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_synchronize_slots[8])))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if base.B2i32(v216 == int32(0))|base.B2i32(v216 != v219) != 0 {
		v237 = v216
		v238 = v219
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v363 = v362
	goto L78
L81:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	v362 = v361
	goto L80
L82:
	;
	if v237-v238 == int32(0) {
		v360 = int32(_a_F_synchronize_slots_4)
		goto L81
	} else {
		goto L89
	}
L83:
	;
	goto L82
L84:
	;
	v222 = v213
	v223 = v210
	goto L85
L85:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	if v227 == int32(0) {
		v237 = v227
		v238 = v226
		goto L83
	} else {
		goto L87
	}
L86:
	;
	v237 = v227
	v238 = v226
	goto L83
L87:
	;
	v230 = int32(1)
	if v227 == v226 {
		v222 = v222 + v230
		v223 = v223 + v230
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v243 = int32(_a_F_synchronize_slots_5)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_synchronize_slots[9])))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if base.B2i32(v246 == int32(0))|base.B2i32(v246 != v249) != 0 {
		v267 = v246
		v268 = v249
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v267-v268 == int32(0) {
		v360 = int32(_a_F_synchronize_slots_6)
		goto L81
	} else {
		goto L97
	}
L91:
	;
	goto L90
L92:
	;
	v252 = v243
	v253 = v210
	goto L93
L93:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	if v257 == int32(0) {
		v267 = v257
		v268 = v256
		goto L91
	} else {
		goto L95
	}
L94:
	;
	v267 = v257
	v268 = v256
	goto L91
L95:
	;
	v260 = int32(1)
	if v257 == v256 {
		v252 = v252 + v260
		v253 = v253 + v260
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v273 = int32(_a_F_synchronize_slots_7)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_synchronize_slots[10])))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if base.B2i32(v276 == int32(0))|base.B2i32(v276 != v279) != 0 {
		v297 = v276
		v298 = v279
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v297-v298 == int32(0) {
		v360 = int32(_a_F_synchronize_slots_8)
		goto L81
	} else {
		goto L105
	}
L99:
	;
	goto L98
L100:
	;
	v282 = v273
	v283 = v210
	goto L101
L101:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	if v287 == int32(0) {
		v297 = v287
		v298 = v286
		goto L99
	} else {
		goto L103
	}
L102:
	;
	v297 = v287
	v298 = v286
	goto L99
L103:
	;
	v290 = int32(1)
	if v287 == v286 {
		v282 = v282 + v290
		v283 = v283 + v290
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v303 = int32(_a_F_synchronize_slots_9)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_synchronize_slots[11])))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if base.B2i32(v306 == int32(0))|base.B2i32(v306 != v309) != 0 {
		v327 = v306
		v328 = v309
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v327-v328 == int32(0) {
		v360 = int32(_a_F_synchronize_slots_10)
		goto L81
	} else {
		goto L113
	}
L107:
	;
	goto L106
L108:
	;
	v312 = v303
	v313 = v210
	goto L109
L109:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+1)))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	if v317 == int32(0) {
		v327 = v317
		v328 = v316
		goto L107
	} else {
		goto L111
	}
L110:
	;
	v327 = v317
	v328 = v316
	goto L107
L111:
	;
	v320 = int32(1)
	if v317 == v316 {
		v312 = v312 + v320
		v313 = v313 + v320
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v332 = int32(_a_F_synchronize_slots_11)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_synchronize_slots[12])))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if base.B2i32(v335 == int32(0))|base.B2i32(v335 != v338) != 0 {
		v356 = v335
		v357 = v338
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v356-v357 != 0 {
		v362 = v203
		goto L80
	} else {
		goto L121
	}
L115:
	;
	goto L114
L116:
	;
	v341 = v332
	v342 = v210
	goto L117
L117:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	if v346 == int32(0) {
		v356 = v346
		v357 = v345
		goto L115
	} else {
		goto L119
	}
L118:
	;
	v356 = v346
	v357 = v345
	goto L115
L119:
	;
	v349 = int32(1)
	if v346 == v345 {
		v341 = v341 + v349
		v342 = v342 + v349
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v360 = int32(_a_F_synchronize_slots_12)
	goto L81
L122:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+12))
	m.T0[v382].(func(*base.Module, int32))(m, v61)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L132
	}
L123:
	;
	v378 = F_lappend(m, v75, v85)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L131
	}
L124:
	;
	F_pfree(m, v85)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L130
	}
L125:
	;
	if v363 != 0 {
		goto L123
	} else {
		goto L129
	}
L126:
	;
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
	if v369 == int64(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v85)+40))
	if v372|v363 == int32(0) {
		goto L124
	} else {
		goto L128
	}
L128:
	;
	goto L123
L129:
	;
	goto L124
L130:
	;
	v380 = v75
	goto L122
L131:
	;
	v380 = v378
	goto L122
L132:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v388 = F_tuplestore_gettupleslot(m, v385, int32(1), int32(0), v61)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	if v388 != 0 {
		v75 = v380
		goto L18
	} else {
		goto L134
	}
L134:
	;
	goto L19
L135:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[13]))
	if int32(0) < v414 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[14]))
	v420 = int32(0)
	v421 = v414
	v422 = v418
	v426 = v2
	goto L139
L137:
	;
	v463 = v2
	goto L138
L138:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[7]))
	F_LWLockRelease(m, v474+int32(_a_F_synchronize_slots_2))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L5
	} else {
		goto L146
	}
L139:
	;
	v438 = v422 + v420*int32(288)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+4)))
	if v439 != int32(1) {
		v451 = v421
		v452 = v422
		v453 = v426
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v463 = v453
	goto L138
L141:
	;
	v455 = v420 + int32(1)
	if v455 < v451 {
		v420 = v455
		v421 = v451
		v422 = v452
		v426 = v453
		goto L139
	} else {
		goto L145
	}
L142:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+201)))
	if v442 == int32(0) {
		v451 = v421
		v452 = v422
		v453 = v426
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v445 = F_lappend(m, v426, v438)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[13]))
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[14]))
	v451 = v448
	v452 = v450
	v453 = v445
	goto L141
L145:
	;
	goto L140
L146:
	;
	if v463 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	if v397 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L148:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v481 <= int32(0) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v493 = v2
	goto L150
L150:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500+v493<<(uint(int32(2))%32))))
	v506 = v504 + int32(24)
	if v397 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L147
L152:
	;
	v671 = v493 + int32(1)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v671 < v672 {
		v493 = v671
		goto L150
	} else {
		goto L194
	}
L153:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v504)+88))
	F_LockSharedObject(m, int32(1262), v597, int32(1))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L5
	} else {
		goto L177
	}
L154:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v509 <= int32(0) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v514 = int32(0)
	goto L156
L156:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v512+v514<<(uint(int32(2))%32))))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if base.B2i32(v537 == int32(0))|base.B2i32(v537 != v540) != 0 {
		v558 = v537
		v559 = v540
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = int32(1)
	if v564 != 0 {
		goto L169
	} else {
		goto L170
	}
L158:
	;
	if v558-v559 != 0 {
		goto L165
	} else {
		goto L166
	}
L159:
	;
	goto L158
L160:
	;
	v543 = v534
	v544 = v506
	goto L161
L161:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+1)))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+1)))
	if v548 == int32(0) {
		v558 = v548
		v559 = v547
		goto L159
	} else {
		goto L163
	}
L162:
	;
	v558 = v548
	v559 = v547
	goto L159
L163:
	;
	v551 = int32(1)
	if v548 == v547 {
		v543 = v543 + v551
		v544 = v544 + v551
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v562 = v514 + int32(1)
	if v562 != v509 {
		v514 = v562
		goto L156
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	goto L157
L168:
	;
	goto L153
L169:
	;
	F_s_lock(m, v504, int32(_a_F_synchronize_slots_13), int32(395), int32(_a_F_synchronize_slots_14))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L5
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v533)+44))
	if v572 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	goto L171
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = int32(0)
	goto L152
L174:
	;
	goto L175
L175:
	;
	v575 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v575
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v504)+112))
	if v577 == v575 {
		goto L152
	} else {
		goto L176
	}
L176:
	;
	goto L153
L177:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = int32(1)
	if v601 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	F_s_lock(m, v504, int32(_a_F_synchronize_slots_13), int32(461), int32(_a_F_synchronize_slots_15))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L5
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+4)))
	if v609 == int32(1) {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L180
L182:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v504)+88))
	F_UnlockSharedObject(m, int32(1262), v631, int32(1))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L5
	} else {
		goto L189
	}
L183:
	;
	v612 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v612
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+201)))
	if v614 == v612 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = int32(0)
	goto L182
L186:
	;
	F_ReplicationSlotAcquire(m, v506, int32(1), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	v621 = int32(_a_F_synchronize_slots_16)
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[15]))
	*(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[15])) = int32(0)
	F_ReplicationSlotDropPtr(m, v622)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	goto L182
L189:
	;
	v637 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L5
	} else {
		goto L190
	}
L190:
	;
	if v637 == int32(0) {
		goto L152
	} else {
		goto L191
	}
L191:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v504)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v506
	F_errmsg(m, int32(_a_F_synchronize_slots_17), v19+int32(96))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_13), int32(477), int32(_a_F_synchronize_slots_15))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L5
	} else {
		goto L193
	}
L193:
	;
	goto L152
L194:
	;
	goto L151
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L5
	} else {
		goto L314
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L5
	} else {
		goto L310
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L5
	} else {
		goto L306
	}
L198:
	;
	F_list_free_deep(m, v397)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L5
	} else {
		goto L304
	}
L199:
	;
	v1034 = int32(0)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v693 = int32(0)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v694 <= v693 {
		v1034 = v693
		goto L198
	} else {
		goto L202
	}
L202:
	;
	v706 = v693
	v709 = int32(0)
	goto L203
L203:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v717+v709<<(uint(int32(2))%32))))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	v724 = F_get_database_oid(m, v722, int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L5
	} else {
		goto L205
	}
L204:
	;
	v1034 = v1023
	goto L198
L205:
	;
	F_LockSharedObject(m, int32(1262), v724, int32(1))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	v729 = m.G0
	v731 = v729 - int32(16)
	m.G0 = v731
	v736 = F_GetWalRcvFlushRecPtr(m, int32(0), v731+int32(8))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v740 = F_GetXLogReplayRecPtr(m, v731+int32(12))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v731)+12))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	m.G0 = v731 + int32(16)
	if base.Ui64(v740) < base.Ui64(v736) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	F_UnlockSharedObject(m, int32(1262), v724, int32(1))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L5
	} else {
		goto L302
	}
L210:
	;
	v748 = v736
	goto L212
L211:
	;
	v748 = v740
	goto L212
L212:
	;
	if v742 == v743 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v750 = v748
	goto L215
L214:
	;
	v750 = v740
	goto L215
L215:
	;
	v751 = *(*int64)(unsafe.Add(mBase, uint32(v721)+24))
	if base.Ui64(v750) < base.Ui64(v751) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v753 = int32(0)
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[16]))
	if v757 == int32(7) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L218
L218:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	v790 = F_SearchNamedReplicationSlot(m, v788, int32(1))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L5
	} else {
		goto L227
	}
L219:
	;
	v760 = int32(15)
	goto L221
L220:
	;
	v760 = int32(21)
	goto L221
L221:
	;
	v762 = F_errstart(m, v760, int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L5
	} else {
		goto L222
	}
L222:
	;
	if v762 == int32(0) {
		v1013 = v753
		goto L209
	} else {
		goto L223
	}
L223:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	v769 = *(*int64)(unsafe.Add(mBase, uint32(v721)+24))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	*(*uint32)(unsafe.Add(mBase, uint32(v19+int32(16)))) = uint32(v750)
	v772 = int64(32)
	v773 = int64(base.Ui64(v750) >> (uint(v772) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+12)) = uint32(v773)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v770
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+4)) = uint32(v769)
	v778 = int64(base.Ui64(v769) >> (uint(v772) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19))) = uint32(v778)
	F_errmsg(m, int32(_a_F_synchronize_slots_18), v19)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L5
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_13), int32(649), int32(_a_F_synchronize_slots_19))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L5
	} else {
		goto L226
	}
L226:
	;
	v1013 = v753
	goto L209
L227:
	;
	if v790 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	*(*int32)(unsafe.Add(mBase, uint32(v790))) = int32(1)
	if v792 != 0 {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	goto L230
L230:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v721)+44))
	if v851 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L231:
	;
	F_s_lock(m, v790, int32(_a_F_synchronize_slots_13), int32(659), int32(_a_F_synchronize_slots_19))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L5
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v800 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v790))) = v800
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790)+201)))
	if v802 == v800 {
		goto L197
	} else {
		goto L235
	}
L234:
	;
	goto L233
L235:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	F_ReplicationSlotAcquire(m, v805, int32(1), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L5
	} else {
		goto L236
	}
L236:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v790)+112))
	if v810 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L5
	} else {
		goto L257
	}
L238:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v721)+44))
	if v811 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	*(*int32)(unsafe.Add(mBase, uint32(v790))) = int32(1)
	if v812 != 0 {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	goto L241
L241:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v790)+92))
	if v830 == int32(2) {
		goto L249
	} else {
		goto L250
	}
L242:
	;
	F_s_lock(m, v790, int32(_a_F_synchronize_slots_13), int32(698), int32(_a_F_synchronize_slots_19))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v721)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v790))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v790)+112)) = v820
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L5
	} else {
		goto L246
	}
L245:
	;
	goto L244
L246:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L5
	} else {
		goto L247
	}
L247:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v790)+112))
	if v828 != 0 {
		goto L237
	} else {
		goto L248
	}
L248:
	;
	goto L241
L249:
	;
	v833 = F_update_and_persist_local_synced_slot(m, v721, v724)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L5
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v837 = *(*int64)(unsafe.Add(mBase, uint32(v721)+24))
	v838 = *(*int64)(unsafe.Add(mBase, uint32(v790)+120))
	if base.Ui64(v837) < base.Ui64(v838) {
		goto L196
	} else {
		goto L254
	}
L252:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L5
	} else {
		goto L253
	}
L253:
	;
	v1013 = v833
	goto L209
L254:
	;
	v840 = int32(0)
	v842 = F_update_local_synced_slot(m, v721, v724, v840, v840)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L256
	}
L256:
	;
	v1013 = v842
	goto L209
L257:
	;
	v1013 = base.B2i32(v810 == int32(0))
	goto L209
L258:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	v855 = int32(1)
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721)+12)))
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721)+13)))
	F_ReplicationSlotCreate(m, v854, v855, int32(2), v857, v858, v855)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L5
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1013 = int32(0)
	goto L209
L261:
	;
	v863 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[15]))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v721)+4))
	v868 = F_strncpy(m, v19+int32(176), v866, int32(64))
	mBase = m.M
	v869 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v868)+63)) = uint8(v869)
	goto L262
L262:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	*(*int32)(unsafe.Add(mBase, uint32(v863))) = int32(1)
	if v871 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	F_s_lock(m, v863, int32(_a_F_synchronize_slots_13), int32(773), int32(_a_F_synchronize_slots_19))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L5
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v863)+88)) = v724
	v880 = *(*int64)(unsafe.Add(mBase, uint32(v19)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+137)) = v880
	v882 = *(*int64)(unsafe.Add(mBase, uint32(v19)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+145)) = v882
	v884 = *(*int64)(unsafe.Add(mBase, uint32(v19)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+153)) = v884
	v886 = *(*int64)(unsafe.Add(mBase, uint32(v19)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+161)) = v886
	v888 = *(*int64)(unsafe.Add(mBase, uint32(v19)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+169)) = v888
	v890 = *(*int64)(unsafe.Add(mBase, uint32(v19)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+177)) = v890
	v892 = *(*int64)(unsafe.Add(mBase, uint32(v19)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+185)) = v892
	v894 = *(*int64)(unsafe.Add(mBase, uint32(v19)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v863)+193)) = v894
	v896 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v863))) = v896
	v898 = *(*int64)(unsafe.Add(mBase, uint32(v721)+16))
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[15]))
	v902 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[7]))
	v906 = F_LWLockAcquire(m, v902+int32(_a_F_synchronize_slots_20), v896)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L5
	} else {
		goto L267
	}
L266:
	;
	goto L265
L267:
	;
	v908 = F_GetRedoRecPtr(m)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L5
	} else {
		goto L268
	}
L268:
	;
	v911 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[17]))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v911)+440)) = int32(1)
	if v912 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[17]))
	F_s_lock(m, v916+int32(440), int32(_a_F_synchronize_slots_21), int32(2685), int32(_a_F_synchronize_slots_22))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L5
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v925 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v925)+440)) = int32(0)
	v928 = *(*int64)(unsafe.Add(mBase, uint32(v925)+224))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v900)))
	*(*int32)(unsafe.Add(mBase, uint32(v900))) = int32(1)
	if v929 != 0 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	goto L271
L273:
	;
	F_s_lock(m, v900, int32(_a_F_synchronize_slots_13), int32(539), int32(_a_F_synchronize_slots_23))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L5
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v900))) = int32(0)
	if base.Ui64(v908) < base.Ui64(v928) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	goto L275
L277:
	;
	v940 = v908
	goto L279
L278:
	;
	v940 = v928
	goto L279
L279:
	;
	if v928 == int64(0) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v943 = v908
	goto L282
L281:
	;
	v943 = v940
	goto L282
L282:
	;
	if base.Ui64(v943) < base.Ui64(v898) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v945 = v898
	goto L285
L284:
	;
	v945 = v943
	goto L285
L285:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v900)+104)) = v945
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	v950 = int64(*(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[18])))
	v951 = *(*int64)(unsafe.Add(mBase, uint32(v900)+104))
	v952 = F_XLogGetLastRemovedSegno(m)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L5
	} else {
		goto L287
	}
L287:
	;
	v954 = base.I64_div_u_s(v951, v950)
	if base.Ui64(v954) <= base.Ui64(v952) {
		goto L195
	} else {
		goto L288
	}
L288:
	;
	v957 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[7]))
	F_LWLockRelease(m, v957+int32(_a_F_synchronize_slots_20))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L5
	} else {
		goto L289
	}
L289:
	;
	v963 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[7]))
	v967 = F_LWLockAcquire(m, v963+int32(_a_F_synchronize_slots_2), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L5
	} else {
		goto L290
	}
L290:
	;
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[7]))
	v974 = F_LWLockAcquire(m, v970+int32(512), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	v977 = F_GetOldestSafeDecodingTransactionId(m, int32(1))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L5
	} else {
		goto L292
	}
L292:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	*(*int32)(unsafe.Add(mBase, uint32(v863))) = int32(1)
	if v979 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	F_s_lock(m, v863, int32(_a_F_synchronize_slots_13), int32(783), int32(_a_F_synchronize_slots_19))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L5
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v863)+100)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(v863)+20)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(v863))) = int32(0)
	v991 = int32(1)
	F_ReplicationSlotsComputeRequiredXmin(m, v991)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L5
	} else {
		goto L297
	}
L296:
	;
	goto L295
L297:
	;
	v996 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[7]))
	F_LWLockRelease(m, v996+int32(512))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L5
	} else {
		goto L298
	}
L298:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, _c_F_synchronize_slots[7]))
	F_LWLockRelease(m, v1002+int32(_a_F_synchronize_slots_2))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L5
	} else {
		goto L299
	}
L299:
	;
	v1007 = F_update_and_persist_local_synced_slot(m, v721, v724)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L5
	} else {
		goto L300
	}
L300:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	v1013 = v991
	goto L209
L302:
	;
	v1023 = v1013 | v706
	v1025 = v709 + int32(1)
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v1025 < v1026 {
		v706 = v1023
		v709 = v1025
		goto L203
	} else {
		goto L303
	}
L303:
	;
	goto L204
L304:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	if v1046 != 0 {
		goto L8
	} else {
		goto L305
	}
L305:
	;
	goto L7
L306:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L5
	} else {
		goto L307
	}
L307:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v1054
	F_errmsg(m, int32(_a_F_synchronize_slots_24), v19+int32(48))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L5
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_13), int32(669), int32(_a_F_synchronize_slots_19))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L5
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v1070
	F_errmsg_internal(m, int32(_a_F_synchronize_slots_25), v19+int32(80))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L5
	} else {
		goto L311
	}
L311:
	;
	v1077 = *(*int64)(unsafe.Add(mBase, uint32(v790)+120))
	v1078 = *(*int64)(unsafe.Add(mBase, uint32(v721)+24))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+76)) = uint32(v1078)
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+68)) = uint32(v1077)
	v1081 = int64(32)
	v1082 = int64(base.Ui64(v1078) >> (uint(v1081) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+72)) = uint32(v1082)
	v1085 = int64(base.Ui64(v1077) >> (uint(v1081) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+64)) = uint32(v1085)
	F_errdetail_internal(m, int32(_a_F_synchronize_slots_26), v19-int32(-64))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L5
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_13), int32(739), int32(_a_F_synchronize_slots_19))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L5
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v900 + int32(24)
	F_errmsg_internal(m, int32(_a_F_synchronize_slots_27), v19+int32(32))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L5
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_13), int32(548), int32(_a_F_synchronize_slots_23))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L5
	} else {
		goto L316
	}
L316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L317:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v1118
	F_errmsg(m, int32(_a_F_synchronize_slots_28), v19+int32(112))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_synchronize_slots_13), int32(839), int32(_a_F_synchronize_slots_29))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L5
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	goto L7
L321:
	;
	F_tuplestore_end(m, v1132)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L5
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v1135 != 0 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	goto L323
L325:
	;
	F_FreeTupleDesc(m, v1135)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	F_pfree(m, v54)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L5
	} else {
		goto L329
	}
L328:
	;
	goto L327
L329:
	;
	if v40 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L5
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	m.G0 = v19 + int32(240)
	return v1034 & int32(1)
L333:
	;
	goto L332
}
