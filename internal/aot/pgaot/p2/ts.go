package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TSConfigIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(74), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v148
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v148 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(42823), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(468864), int32(3241), int32(60933))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L47
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	if v87 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v76 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v76
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v76 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v76 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 < v91 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v90
	goto L36
L34:
	;
	goto L35
L35:
	;
	v138 = v3
	goto L14
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v96<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, _consts[287]))
	if v108 != v110 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v108 == v40 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v120 = v96 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v120 < v121 {
		v96 = v120
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v138 = int32(1)
	goto L14
L42:
	;
	goto L43
L43:
	;
	v115 = int32(0)
	v117 = F_SearchSysCacheExists(m, int32(73), v39+int32(4), v108, v115, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v117 != 0 {
		v138 = v3
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	goto L37
L47:
	;
	v148 = v138
	goto L1
}
func F_TSTemplateIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(80), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v148
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v148 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(46225), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(468864), int32(3095), int32(60954))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L47
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	if v87 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v76 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v76
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v76 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v76 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 < v91 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v90
	goto L36
L34:
	;
	goto L35
L35:
	;
	v138 = v3
	goto L14
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v96<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, _consts[287]))
	if v108 != v110 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v108 == v40 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v120 = v96 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v120 < v121 {
		v96 = v120
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v138 = int32(1)
	goto L14
L42:
	;
	goto L43
L43:
	;
	v115 = int32(0)
	v117 = F_SearchSysCacheExists(m, int32(79), v39+int32(4), v108, v115, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v117 != 0 {
		v138 = v3
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	goto L37
L47:
	;
	v148 = v138
	goto L1
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
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
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
	var v172 int32
	_ = v172
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
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
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
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
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
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
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
	v25 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
		goto L12
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v18 + int32(48)
	return v331
L8:
	;
	v331 = int32(0)
	goto L7
L9:
	;
	v304 = int32(1)
	v307 = F_palloc0(m, int32(16))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L78
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L75
	}
L11:
	;
	F_pfree(m, v34)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L74
	}
L12:
	;
	v34 = F_palloc0(m, int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	switch v49 - int32(1) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L9
	default:
		goto L10
	}
L15:
	;
	v36 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v36 != int32(1) {
		goto L11
	} else {
		goto L17
	}
L17:
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
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46
	v331 = v42
	goto L7
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = F_TS_execute_locations_recurse(m, l0+v84*int32(12), l1, l2, v18+int32(44))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L28
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v66 = F_TS_execute_locations_recurse(m, l0+v60*int32(12), l1, l2, v18+int32(44))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v56 = F_TS_execute_locations_recurse(m, l0+int32(12), l1, l2, v18+int32(44))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v331 = v56 ^ int32(1)
	goto L7
L23:
	;
	if v66 == int32(0) {
		v331 = v5
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v74 = F_TS_execute_locations_recurse(m, l0+int32(12), l1, l2, v18+int32(40))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v74 == int32(0) {
		v331 = v5
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v80 = F_list_concat(m, v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v80
	v331 = int32(1)
	goto L7
L28:
	;
	v96 = F_TS_execute_locations_recurse(m, l0+int32(12), l1, l2, v18+int32(40))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v90 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v100 = int32(0)
	if v96 == v100 {
		v331 = v100
		goto L7
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v105 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v104
	v331 = int32(1)
	goto L7
L35:
	;
	goto L36
L36:
	;
	if v104 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v111 <= int32(0) {
		v331 = int32(1)
		goto L7
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v105
	v331 = int32(1)
	goto L7
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v115 = v114
	v117 = v111
	v127 = v5
	goto L41
L41:
	;
	if int32(0) < v115 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v331 = v280
	goto L7
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v127<<(uint(int32(2))%32))))
	v151 = int32(0)
	goto L46
L44:
	;
	v265 = v115
	v267 = v117
	goto L45
L45:
	;
	v280 = int32(1)
	v282 = v127 + v280
	if v282 < v267 {
		v115 = v265
		v117 = v267
		v127 = v282
		goto L41
	} else {
		goto L73
	}
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v151<<(uint(int32(2))%32))))
	v159 = F_palloc0(m, int32(16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v265 = v262
	v267 = v264
	goto L45
L48:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v166 = int32(0)
	v168 = v166
	v172 = v166
	v179 = v161
	goto L49
L49:
	;
	if v179 <= v172 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	if v252 < v251 {
		goto L68
	} else {
		goto L69
	}
L51:
	;
	goto L50
L52:
	;
	if v222 == int32(0) {
		v168 = v221
		v172 = v223
		goto L49
	} else {
		goto L62
	}
L53:
	;
	v221 = v168 + int32(1)
	v222 = v217
	v223 = v172
	goto L52
L54:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v184 <= v168 {
		goto L51
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193+v172<<(uint(int32(1))%32)))))
	v199 = v197 & int32(16383)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v200 <= v168 {
		v213 = v168
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186+v168<<(uint(int32(1))%32)))))
	v217 = v190 & int32(16383)
	goto L53
L58:
	;
	v221 = v213
	v222 = v199
	v223 = v172 + int32(1)
	goto L52
L59:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202+v168<<(uint(int32(1))%32)))))
	v208 = v206 & int32(16383)
	if base.Ui32(v199) < base.Ui32(v208) {
		v213 = v168
		goto L58
	} else {
		goto L60
	}
L60:
	;
	if v208 != v199 {
		v217 = v208
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v213 = v168 + int32(1)
	goto L58
L62:
	;
	if v159 == int32(0) {
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	if v229 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v232 = F_palloc(m, (v161+v162)<<(uint(int32(1))%32))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	v237 = v229
	goto L66
L66:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v239 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v238 + v239
	*(*uint16)(unsafe.Add(mBase, uint32(v237+v238<<(uint(v239)%32)))) = uint16(v222)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v168 = v221
	v172 = v223
	v179 = v246
	goto L49
L67:
	;
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+4)) = uint8(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = v232
	v237 = v232
	goto L66
L68:
	;
	v254 = v251
	goto L70
L69:
	;
	v254 = v252
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+12)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v257 = F_lappend(m, v256, v159)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v257
	v261 = v151 + int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v261 < v262 {
		v151 = v261
		goto L46
	} else {
		goto L72
	}
L72:
	;
	goto L47
L73:
	;
	goto L42
L74:
	;
	goto L8
L75:
	;
	v292 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v292
	F_errmsg_internal(m, int32(452798), v18+int32(16))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(465394), int32(2140), int32(338472))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v309 = F_TS_phrase_execute(m, l0, l1, int32(0), l2, v307)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v309 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+5)))
	if v313 != 0 {
		v331 = v304
		goto L7
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_pfree(m, v307)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v307
	v319 = F_list_make1_impl(m, int32(1), v18+int32(28))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v319
	v331 = v304
	goto L7
L85:
	;
	goto L8
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
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
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
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = l0
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1178]))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[1179]))
	if v45 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+76)) = int64(206158430212)
	v24 = F_hash_create(m, int32(375121), int32(8), v11+int32(60), int32(40))
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
	*(*int32)(unsafe.Add(mBase, _consts[1178])) = v24
	F_CacheRegisterSyscacheCallback(m, int32(76), int32(1612), v24)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[1178]))
	F_CacheRegisterSyscacheCallback(m, int32(80), int32(1612), v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[400]))
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
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L72
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L3
	} else {
		goto L69
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L3
	} else {
		goto L66
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L63
	}
L13:
	;
	m.G0 = v11 + int32(112)
	return v209
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[1178]))
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
		v209 = v45
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1179])) = v200
	v209 = v200
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
		v200 = v57
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
	if v118&int32(3) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[1178]))
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
	goto L38
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	v100 = F_AllocSetContextCreateInternal(m, v95, int32(16343), int32(0), int32(1024), int32(8192))
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
	goto L37
L37:
	;
	v117 = v100
	v118 = v92
	goto L30
L38:
	;
	F_MemoryContextReset(m, v107)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v114 = F_MemoryContextStrdup(m, v107, v68+int32(4))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+36)) = v114
	goto L41
L41:
	;
	v117 = v107
	v118 = v57
	goto L30
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v79)+72))
	if v159 != 0 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	v124 = v118 + int32(48)
	if base.Ui32(v124) <= base.Ui32(v118) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v140 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v118)+4)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v118)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v118)+36)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v118)+28)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v118)+20)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v118)+12)) = v140
	goto L42
L46:
	;
	v130 = v118 + int32(4)
	if base.Ui32(v130) < base.Ui32(v124) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v132 = v124
	goto L49
L48:
	;
	v132 = v130
	goto L49
L49:
	;
	v139 = F__emscripten_memset_bulkmem(m, v118, base.I32_extend8_s(int32(0)), (v118^int32(-1)+v132)&int32(-4)+int32(4))
	mBase = m.M
	goto L50
L50:
	;
	goto L42
L51:
	;
	v160 = int32(4425280)
	v161 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v117
	v168 = F_SysCacheGetAttr(m, int32(76), v62, int32(6), v11+int32(60))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_ReleaseCatCache(m, v73)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L60
	}
L54:
	;
	v170 = int32(0)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)))
	if v171 == v170 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v174 = F_deserialize_deflist(m, v168)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	v176 = v170
	goto L57
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v79)+72))
	v179 = F_OidFunctionCall1Coll(m, v177, int32(0), v176)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L59
	}
L58:
	;
	v176 = v174
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+44)) = v179
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v161
	goto L53
L60:
	;
	F_ReleaseCatCache(m, v62)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	F_fmgr_info_cxt(m, v191, v118+int32(12), v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	v197 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)) = uint8(v197)
	v200 = v118
	goto L18
L63:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v223
	F_errmsg_internal(m, int32(36577), v11)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(468598), int32(256), int32(374973))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
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
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v237
	F_errmsg_internal(m, int32(333003), v11+int32(16))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(468598), int32(263), int32(374973))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
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
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v68)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v253
	F_errmsg_internal(m, int32(46225), v11+int32(32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(468598), int32(272), int32(374973))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
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
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v269
	F_errmsg_internal(m, int32(397059), v11+int32(48))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(468598), int32(280), int32(374973))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	v2 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 <= int32(0) {
		return int64(0)
	} else {
		v11 = int32(1)
		v14 = l0 + int32(8)
		if v6 == v11 {
			v46 = v14
			v47 = v2
		} else {
			v20 = v14
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
			v46 = v42
			v47 = v40
		}
		if v6&v11 == int32(0) {
			v60 = v47
		} else {
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
			if v53 != int32(1) {
				v60 = v47
			} else {
				v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v46)+4)))
				v60 = int64(1)<<(uint(v57)%64) | v47
			}
		}
		return v60
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
		v11 = F_DirectFunctionCall4Coll(m, int32(1186), int32(0), v4, v8, v9, v10)
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
	v9 = F_DirectFunctionCall1Coll(m, int32(1540), v2, v8)
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
			v18 = F_DirectFunctionCall1Coll(m, int32(1541), int32(0), v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = F_DirectFunctionCall2Coll(m, int32(1538), v2, v13, v18)
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v27 float32
	_ = v27
	var v28 int32
	_ = v28
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
	var v40 int32
	_ = v40
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v19 = l0 + int32(28)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v21 = F_pg_detoast_datum(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_getWeights(m, v14, v11)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_calc_rank(m, v11, v21, v24, v23)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v29 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							if v33 != v21 {
								F_pfree(m, v21)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v37 != v24 {
										F_pfree(m, v24)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(16)
											return base.I32_reinterpret_f32(v27)
										}
									} else {
										m.G0 = v11 + int32(16)
										return base.I32_reinterpret_f32(v27)
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v37 != v24 {
									F_pfree(m, v24)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return base.I32_reinterpret_f32(v27)
									}
								} else {
									m.G0 = v11 + int32(16)
									return base.I32_reinterpret_f32(v27)
								}
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						if v33 != v21 {
							F_pfree(m, v21)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v37 != v24 {
									F_pfree(m, v24)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return base.I32_reinterpret_f32(v27)
									}
								} else {
									m.G0 = v11 + int32(16)
									return base.I32_reinterpret_f32(v27)
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v37 != v24 {
								F_pfree(m, v24)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(16)
									return base.I32_reinterpret_f32(v27)
								}
							} else {
								m.G0 = v11 + int32(16)
								return base.I32_reinterpret_f32(v27)
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 float32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v16 = F_calc_rank_cd(m, int32(1689904), v8, v14, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v18 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v16)
						}
					} else {
						return base.I32_reinterpret_f32(v16)
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v22 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return base.I32_reinterpret_f32(v16)
					}
				} else {
					return base.I32_reinterpret_f32(v16)
				}
			}
		}
	}
}
func F_ts_rankcd_ttf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 float32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v16 = F_calc_rank_cd(m, int32(1689904), v8, v14, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v18 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v16)
						}
					} else {
						return base.I32_reinterpret_f32(v16)
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v22 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return base.I32_reinterpret_f32(v16)
					}
				} else {
					return base.I32_reinterpret_f32(v16)
				}
			}
		}
	}
}
