package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TSConfigIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn13840(m, l0, l1, int32(73), int32(_a_F_TSConfigIsVisibleExt_0), int32(3241), int32(_a_F_TSConfigIsVisibleExt_1), int32(74))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_TSTemplateIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn13840(m, l0, l1, int32(79), int32(_a_F_TSTemplateIsVisibleExt_0), int32(3095), int32(_a_F_TSTemplateIsVisibleExt_1), int32(80))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_TS_execute_locations_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
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
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	F_check_stack_depth(m)
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
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_TS_execute_locations_recurse[0]))
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v30 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v18 + int32(48)
	return v330
L8:
	;
	F_pfree(m, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L83
	}
L9:
	;
	v302 = int32(1)
	v305 = F_palloc0(m, int32(16))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L76
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L73
	}
L11:
	;
	v34 = F_palloc0(m, int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	switch v49 - int32(1) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L18
	case 3:
		goto L9
	default:
		goto L10
	}
L14:
	;
	v36 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v36 != int32(1) {
		v320 = v34
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v34
	v42 = int32(1)
	v46 = F_list_make1_impl(m, v42, v18+int32(12))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46
	v330 = v42
	goto L7
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = F_TS_execute_locations_recurse(m, l0+v84*int32(12), l1, l2, v18+int32(44))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v66 = F_TS_execute_locations_recurse(m, l0+v60*int32(12), l1, l2, v18+int32(44))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v56 = F_TS_execute_locations_recurse(m, l0+int32(12), l1, l2, v18+int32(44))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v330 = v56 ^ int32(1)
	goto L7
L22:
	;
	if v66 == int32(0) {
		v330 = v5
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v74 = F_TS_execute_locations_recurse(m, l0+int32(12), l1, l2, v18+int32(40))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v74 == int32(0) {
		v330 = v5
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v80 = F_list_concat(m, v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v80
	v330 = int32(1)
	goto L7
L27:
	;
	v96 = F_TS_execute_locations_recurse(m, l0+int32(12), l1, l2, v18+int32(40))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v90 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v100 = int32(0)
	if v96 == v100 {
		v330 = v100
		goto L7
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v105 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v104
	v330 = int32(1)
	goto L7
L34:
	;
	goto L35
L35:
	;
	if v104 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v111 <= int32(0) {
		v330 = int32(1)
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v105
	v330 = int32(1)
	goto L7
L39:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v115 = v114
	v117 = v111
	v128 = v5
	goto L40
L40:
	;
	if int32(0) < v115 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v330 = v280
	goto L7
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v128<<(uint(int32(2))%32))))
	v150 = int32(0)
	goto L45
L43:
	;
	v265 = v115
	v267 = v117
	goto L44
L44:
	;
	v280 = int32(1)
	v282 = v128 + v280
	if v282 < v267 {
		v115 = v265
		v117 = v267
		v128 = v282
		goto L40
	} else {
		goto L72
	}
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v150<<(uint(int32(2))%32))))
	v159 = F_palloc0(m, int32(16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v265 = v262
	v267 = v264
	goto L44
L47:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v166 = int32(0)
	v168 = v166
	v173 = v166
	v179 = v161
	goto L48
L48:
	;
	if v179 <= v173 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	if v252 < v251 {
		goto L67
	} else {
		goto L68
	}
L50:
	;
	goto L49
L51:
	;
	if v223 == int32(0) {
		v168 = v221
		v173 = v224
		goto L48
	} else {
		goto L61
	}
L52:
	;
	v221 = v168 + int32(1)
	v223 = v218
	v224 = v173
	goto L51
L53:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v184 <= v168 {
		goto L50
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193+v173<<(uint(int32(1))%32)))))
	v199 = v197 & int32(_a_F_TS_execute_locations_recurse_0)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v200 <= v168 {
		v213 = v168
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186+v168<<(uint(int32(1))%32)))))
	v218 = v190 & int32(_a_F_TS_execute_locations_recurse_0)
	goto L52
L57:
	;
	v221 = v213
	v223 = v199
	v224 = v173 + int32(1)
	goto L51
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202+v168<<(uint(int32(1))%32)))))
	v208 = v206 & int32(_a_F_TS_execute_locations_recurse_0)
	if base.Ui32(v199) < base.Ui32(v208) {
		v213 = v168
		goto L57
	} else {
		goto L59
	}
L59:
	;
	if v199 != v208 {
		v218 = v208
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v213 = v168 + int32(1)
	goto L57
L61:
	;
	if v159 == int32(0) {
		goto L50
	} else {
		goto L62
	}
L62:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	if v229 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v232 = F_palloc(m, (v161+v162)<<(uint(int32(1))%32))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	v237 = v229
	goto L65
L65:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v239 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v238 + v239
	*(*uint16)(unsafe.Add(mBase, uint32(v237+v238<<(uint(v239)%32)))) = uint16(v223)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v168 = v221
	v173 = v224
	v179 = v246
	goto L48
L66:
	;
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+4)) = uint8(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = v232
	v237 = v232
	goto L65
L67:
	;
	v254 = v251
	goto L69
L68:
	;
	v254 = v252
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+12)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v257 = F_lappend(m, v256, v159)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v257
	v261 = v150 + int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v261 < v262 {
		v150 = v261
		goto L45
	} else {
		goto L71
	}
L71:
	;
	goto L46
L72:
	;
	goto L41
L73:
	;
	v290 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v290
	F_errmsg_internal(m, int32(_a_F_TS_execute_locations_recurse_1), v18+int32(16))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_TS_execute_locations_recurse_2), int32(2140), int32(_a_F_TS_execute_locations_recurse_3))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
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
	v307 = F_TS_phrase_execute(m, l0, l1, int32(0), l2, v305)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v307 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+5)))
	if v311 != 0 {
		v330 = v302
		goto L7
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v320 = v305
	goto L8
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v305
	v317 = F_list_make1_impl(m, int32(1), v18+int32(28))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v317
	v330 = v302
	goto L7
L83:
	;
	v330 = int32(0)
	goto L7
}
func F_lookup_ts_dictionary_cache(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
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
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = l0
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0]))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[1]))
	if v45 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+76)) = int64(206158430212)
	v24 = F_hash_create(m, int32(_a_F_lookup_ts_dictionary_cache_9), int32(8), v11+int32(60), int32(40))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0])) = v24
	F_CacheRegisterSyscacheCallback(m, int32(76), int32(1597), v24)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0]))
	F_CacheRegisterSyscacheCallback(m, int32(80), int32(1597), v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[2]))
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L3
	} else {
		goto L60
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L3
	} else {
		goto L57
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L3
	} else {
		goto L54
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L3
	} else {
		goto L51
	}
L13:
	;
	m.G0 = v11 + int32(112)
	return v185
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0]))
	v55 = int32(0)
	v57 = F_hash_search(m, v52, v11+int32(108), v55, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L19
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v48 != l0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+4)))
	if v50 != 0 {
		v185 = v45
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[1])) = v176
	v185 = v176
	goto L13
L19:
	;
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)))
	if v59 != 0 {
		v176 = v57
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v62 = F_SearchSysCache1(m, int32(76), v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	if v62 == int32(0) {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v68 = v66 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+76))
	if v69 == int32(0) {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v73 = F_SearchSysCache1(m, int32(80), v69)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if v73 == int32(0) {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+22)))
	v79 = v77 + v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	if v80 == int32(0) {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v57 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+44)) = int32(0)
	v121 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v118)+36)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v118)+28)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v118)+20)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v118)+12)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v118)+4)) = v121
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v131
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v79)+72))
	if v136 != 0 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0]))
	v92 = F_hash_search(m, v86, v11+int32(108), int32(1), v11+int32(60))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = int32(0)
	F_MemoryContextReset(m, v107)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L37
	}
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[2]))
	v100 = F_AllocSetContextCreateInternal(m, v95, int32(_a_F_lookup_ts_dictionary_cache_6), int32(0), int32(1024), int32(_a_F_lookup_ts_dictionary_cache_7))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v104 = F_MemoryContextStrdup(m, v100, v68+int32(4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+36)) = v104
	v117 = v100
	v118 = v92
	goto L30
L37:
	;
	v114 = F_MemoryContextStrdup(m, v107, v68+int32(4))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = v114
	v117 = v107
	v118 = v57
	goto L30
L39:
	;
	v137 = int32(_a_F_lookup_ts_dictionary_cache_8)
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[3])) = v117
	v145 = F_SysCacheGetAttr(m, int32(76), v62, int32(6), v11+int32(60))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_ReleaseCatCache(m, v73)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L48
	}
L42:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)))
	if v147 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v150 = F_deserialize_deflist(m, v145)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	v152 = int32(0)
	goto L45
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v79)+72))
	v155 = F_OidFunctionCall1Coll(m, v153, int32(0), v152)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L47
	}
L46:
	;
	v152 = v150
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+44)) = v155
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[3])) = v138
	goto L41
L48:
	;
	F_ReleaseCatCache(m, v62)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	F_fmgr_info_cxt(m, v167, v118+int32(12), v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)) = uint8(v173)
	v176 = v118
	goto L18
L51:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v199
	F_errmsg_internal(m, int32(_a_F_lookup_ts_dictionary_cache_0), v11)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_dictionary_cache_1), int32(256), int32(_a_F_lookup_ts_dictionary_cache_2))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
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
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v213
	F_errmsg_internal(m, int32(_a_F_lookup_ts_dictionary_cache_3), v11+int32(16))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_dictionary_cache_1), int32(263), int32(_a_F_lookup_ts_dictionary_cache_2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v68)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v229
	F_errmsg_internal(m, int32(_a_F_lookup_ts_dictionary_cache_4), v11+int32(32))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_dictionary_cache_1), int32(272), int32(_a_F_lookup_ts_dictionary_cache_2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v245
	F_errmsg_internal(m, int32(_a_F_lookup_ts_dictionary_cache_5), v11+int32(48))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_dictionary_cache_1), int32(280), int32(_a_F_lookup_ts_dictionary_cache_2))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makeTSQuerySign(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v61 int64
	_ = v61
	v2 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 <= int32(0) {
		return int64(0)
	} else {
		v12 = l0 + int32(8)
		if v6 != int32(1) {
			v20 = v12
			v21 = v2
			v22 = int32(0)
			for {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				if v25 == int32(1) {
					v29 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+4)))
					v32 = int64(1)<<(uint(v29)%64) | v21
				} else {
					v32 = v21
				}
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+12)))
				if v33 == int32(1) {
					v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+16)))
					v40 = int64(1)<<(uint(v37)%64) | v32
				} else {
					v40 = v32
				}
				v42 = v20 + int32(24)
				v44 = v22 + int32(2)
				if v44 != v6&int32(2147483646) {
					v20 = v42
					v21 = v40
					v22 = v44
					continue
				} else {
					break
				}
				break
			}
			if v6&int32(1) == int32(0) {
				v61 = v40
			} else {
				v48 = v42
				v49 = v40
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
				if v53 != int32(1) {
					v61 = v49
				} else {
					v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+4)))
					v61 = int64(1)<<(uint(v57)%64) | v49
				}
			}
		} else {
			v48 = v12
			v49 = v2
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
			if v53 != int32(1) {
				v61 = v49
			} else {
				v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+4)))
				v61 = int64(1)<<(uint(v57)%64) | v49
			}
		}
		return v61
	}
}
func F_ts_headline_opt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_DirectFunctionCall4Coll(m, int32(1171), int32(0), v4, v8, v9, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_ts_match_tt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v17 int32
	_ = v17
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
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_DirectFunctionCall1Coll(m, int32(1525), v2, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_pg_detoast_datum(m, v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v18 = F_DirectFunctionCall1Coll(m, int32(1526), int32(0), v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = F_DirectFunctionCall2Coll(m, int32(1523), v2, v13, v18)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v13)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v18)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v20 != int32(0))
						}
					}
				}
			}
		}
	}
}
func F_ts_parse_byname(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v5 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v14 = F_pg_detoast_datum_packed(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = F_textToQualifiedNameList(m, v9)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v21 = F_get_ts_parser_oid(m, v18, int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							F_prs_setup_firstcall(m, v16, l0, v21, v14)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
								v29 = F_prs_process_call(m, v28)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									if v29 != 0 {
										v31 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
										*(*int64)(unsafe.Add(mBase, uint32(v28))) = v31 + int64(1)
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(1)
										return v29
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = int32(2)
											v44 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
											return v29
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
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
		v29 = F_prs_process_call(m, v28)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
				*(*int64)(unsafe.Add(mBase, uint32(v28))) = v31 + int64(1)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(1)
				return v29
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = int32(2)
					v44 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
					return v29
				}
			}
		}
	}
}
func F_ts_rank_wttf(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 float32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_getWeights(m, v13, v10)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_calc_rank(m, v10, v18, v21, v20)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v26 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v30 != v18 {
								F_pfree(m, v18)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v34 != v21 {
										F_pfree(m, v21)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I32_reinterpret_f32(v24)
										}
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v24)
									}
								}
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v34 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v24)
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v24)
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v30 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v34 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v24)
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v24)
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v34 != v21 {
								F_pfree(m, v21)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v24)
								}
							} else {
								m.G0 = v10 + int32(16)
								return base.I32_reinterpret_f32(v24)
							}
						}
					}
				}
			}
		}
	}
}
func F_ts_rankcd_tt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 float32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_calc_rank_cd(m, int32(_a_F_ts_rankcd_tt_0), v7, v11, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v15 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v13)
						}
					} else {
						return base.I32_reinterpret_f32(v13)
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v19 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return base.I32_reinterpret_f32(v13)
					}
				} else {
					return base.I32_reinterpret_f32(v13)
				}
			}
		}
	}
}
func F_ts_rankcd_ttf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v13 = F_calc_rank_cd(m, int32(_a_F_ts_rankcd_ttf_0), v7, v11, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v15 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v13)
						}
					} else {
						return base.I32_reinterpret_f32(v13)
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v19 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return base.I32_reinterpret_f32(v13)
					}
				} else {
					return base.I32_reinterpret_f32(v13)
				}
			}
		}
	}
}
