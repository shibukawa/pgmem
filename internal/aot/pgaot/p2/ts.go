package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommitTsShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(_a_F_CommitTsShmemInit_0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTsShmemInit[0]))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTsShmemInit[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)) = uint8(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)) = uint16(v4)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(-9223372036854775807 - 1)
	return
}
func F_TSConfigIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn14223(m, l0, l1, int32(73), int32(_a_F_TSConfigIsVisibleExt_0), int32(3310), int32(_a_F_TSConfigIsVisibleExt_1), int32(74))
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
	v8 = Fn14223(m, l0, l1, int32(79), int32(_a_F_TSTemplateIsVisibleExt_0), int32(3164), int32(_a_F_TSTemplateIsVisibleExt_1), int32(80))
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
	F_errfinish(m, int32(_a_F_TS_execute_locations_recurse_2), int32(2144), int32(_a_F_TS_execute_locations_recurse_3))
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
func F_commit_ts_errdetail_for_io_error(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14249(m, l0, int32(_a_F_commit_ts_errdetail_for_io_error_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_lookup_ts_dictionary_cache(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int64
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
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
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = l0
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0]))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[1]))
	if v48 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = int64(206158430212)
	v26 = F_hash_create(m, int32(_a_F_lookup_ts_dictionary_cache_9), int64(8), v13+int32(56), int32(40))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0])) = v26
	F_CacheRegisterSyscacheCallback(m, int32(76), int32(1811), base.I64_extend_i32_u(v26))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v39 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0])))
	F_CacheRegisterSyscacheCallback(m, int32(80), int32(1811), v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[2]))
	if v43 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
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
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L62
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L3
	} else {
		goto L59
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L3
	} else {
		goto L56
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L53
	}
L13:
	;
	m.G0 = v13 + int32(112)
	return v207
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0]))
	v58 = int32(0)
	v60 = F_hash_search(m, v55, v13+int32(108), v58, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L19
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v51 != l0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
	if v53 != 0 {
		v207 = v48
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[1])) = v196
	v207 = v196
	goto L13
L19:
	;
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
	if v62 != 0 {
		v196 = v60
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+108)))
	v65 = F_SearchSysCache1(m, int32(76), v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	if v65 == int32(0) {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+22)))
	v71 = v69 + v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	if v72 == int32(0) {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v77 = F_SearchSysCache1(m, int32(80), base.I64_extend_i32_u(v72))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if v77 == int32(0) {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+22)))
	v83 = v81 + v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
	if v84 == int32(0) {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v60 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v149 = F_MemoryContextStrdup(m, v144, v71+int32(4))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L40
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = int32(0)
	F_MemoryContextReset(m, v122)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L39
	}
L32:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[2]))
	v137 = F_AllocSetContextCreateInternal(m, v132, int32(_a_F_lookup_ts_dictionary_cache_6), int32(0), int32(1024), int32(_a_F_lookup_ts_dictionary_cache_7))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L38
	}
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[0]))
	v96 = F_hash_search(m, v90, v13+int32(108), int32(1), v13+int32(56))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+4)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v60)+12)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v60)+20)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v60)+28)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v60)+44)) = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+36)) = v112
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v125
	if v122 != 0 {
		goto L31
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+44)) = int32(0)
	v100 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+36)) = v100
	*(*int64)(unsafe.Add(mBase, uint32(v96)+28)) = v100
	*(*int64)(unsafe.Add(mBase, uint32(v96)+20)) = v100
	*(*int64)(unsafe.Add(mBase, uint32(v96)+12)) = v100
	*(*int64)(unsafe.Add(mBase, uint32(v96)+4)) = v100
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v110
	v129 = v96
	goto L32
L37:
	;
	v129 = v60
	goto L32
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+40)) = v137
	v144 = v137
	v145 = v129
	goto L30
L39:
	;
	v144 = v122
	v145 = v60
	goto L30
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+36)) = v149
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
	if v154 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v155 = int32(_a_F_lookup_ts_dictionary_cache_8)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[3]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v145)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[3])) = v158
	v164 = F_SysCacheGetAttr(m, int32(76), v65, int32(6), v13+int32(56))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_ReleaseCatCache(m, v77)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L50
	}
L44:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+56)))
	if v166 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v169 = F_deserialize_deflist(m, v164)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L48
	}
L46:
	;
	v172 = int64(0)
	goto L47
L47:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
	v175 = F_OidFunctionCall1Coll(m, v173, int32(0), v172)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L49
	}
L48:
	;
	v172 = base.I64_extend_i32_u(v169)
	goto L47
L49:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v145)+44)) = uint32(v175)
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_dictionary_cache[3])) = v156
	goto L43
L50:
	;
	F_ReleaseCatCache(m, v65)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v145)+40))
	F_fmgr_info_cxt(m, v187, v145+int32(12), v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+4)) = uint8(v193)
	v196 = v145
	goto L18
L53:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v223
	F_errmsg_internal(m, int32(_a_F_lookup_ts_dictionary_cache_0), v13)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_dictionary_cache_1), int32(257), int32(_a_F_lookup_ts_dictionary_cache_2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v237
	F_errmsg_internal(m, int32(_a_F_lookup_ts_dictionary_cache_3), v13+int32(16))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_dictionary_cache_1), int32(264), int32(_a_F_lookup_ts_dictionary_cache_2))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
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
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v253
	F_errmsg_internal(m, int32(_a_F_lookup_ts_dictionary_cache_4), v13+int32(32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_dictionary_cache_1), int32(273), int32(_a_F_lookup_ts_dictionary_cache_2))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v269
	F_errmsg_internal(m, int32(_a_F_lookup_ts_dictionary_cache_5), v13+int32(48))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_dictionary_cache_1), int32(281), int32(_a_F_lookup_ts_dictionary_cache_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
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
func F_ts_headline_opt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v12 = F_DirectFunctionCall4Coll(m, int32(1286), int32(0), base.I64_extend_i32_u(v4), v9, v10, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			return v12
		}
	}
}
func F_ts_match_tt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v2 = int32(0)
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_DirectFunctionCall1Coll(m, int32(1738), v2, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v15 = F_pg_detoast_datum(m, base.I32_wrap_i64(v10))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			v21 = F_DirectFunctionCall1Coll(m, int32(1739), int32(0), v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v25 = F_DirectFunctionCall2Coll(m, int32(1736), v2, base.I64_extend_i32_u(v15), v21&int64(4294967295))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					F_pfree(m, v15)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						F_pfree(m, base.I32_wrap_i64(v21))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v25 != int64(0)))
						}
					}
				}
			}
		}
	}
}
func F_ts_parse_byname(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v15 = F_pg_detoast_datum_packed(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				v17 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = F_textToQualifiedNameList(m, v10)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						v22 = F_get_ts_parser_oid(m, v19, int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							F_prs_setup_firstcall(m, v17, l0, v22, v15)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int64(0)
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
								v30 = F_prs_process_call(m, v29)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int64(0)
								} else {
									if v30 != int64(0) {
										v34 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
										*(*int64)(unsafe.Add(mBase, uint32(v29))) = v34 + int64(1)
										v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(1)
										return v30
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int64(0)
										} else {
											v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = int32(2)
											v47 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
											return v30
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
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
		v30 = F_prs_process_call(m, v29)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int64(0)
		} else {
			if v30 != int64(0) {
				v34 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
				*(*int64)(unsafe.Add(mBase, uint32(v29))) = v34 + int64(1)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(1)
				return v30
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = int32(2)
					v47 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
					return v30
				}
			}
		}
	}
}
func F_ts_rank_wttf(m *base.Module, l0 int32) int64 {
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			F_getWeights(m, v13, v10)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = F_calc_rank(m, v10, v18, v21, v20)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v26 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v30 != v18 {
								F_pfree(m, v18)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int64(0)
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v34 != v21 {
										F_pfree(m, v21)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int64(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
										}
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
									}
								}
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v34 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v30 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v34 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							if v34 != v21 {
								F_pfree(m, v21)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
								}
							} else {
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_s(base.I32_reinterpret_f32(v24))
							}
						}
					}
				}
			}
		}
	}
}
func F_ts_rankcd_tt(m *base.Module, l0 int32) int64 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_calc_rank_cd(m, int32(_a_F_ts_rankcd_tt_0), v7, v11, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v15 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(base.I32_reinterpret_f32(v13))
						}
					} else {
						return base.I64_extend_i32_s(base.I32_reinterpret_f32(v13))
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v19 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_s(base.I32_reinterpret_f32(v13))
					}
				} else {
					return base.I64_extend_i32_s(base.I32_reinterpret_f32(v13))
				}
			}
		}
	}
}
func F_ts_rankcd_ttf(m *base.Module, l0 int32) int64 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v13 = F_calc_rank_cd(m, int32(_a_F_ts_rankcd_ttf_0), v7, v11, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v15 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(base.I32_reinterpret_f32(v13))
						}
					} else {
						return base.I64_extend_i32_s(base.I32_reinterpret_f32(v13))
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v19 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_s(base.I32_reinterpret_f32(v13))
					}
				} else {
					return base.I64_extend_i32_s(base.I32_reinterpret_f32(v13))
				}
			}
		}
	}
}
