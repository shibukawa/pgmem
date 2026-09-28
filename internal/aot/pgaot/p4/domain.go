package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitDomainConstraintRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	v4 = l3
	v5 = int32(0)
	v9 = F_lookup_type_cache(m, l0, int32(_a_F_InitDomainConstraintRef_0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(1818)
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v20 = l1 + int32(20)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v21
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v15)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v20
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+308))
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v30 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v29 + v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v35 != v30 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v97 = v5
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v97
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34
	return
L8:
	;
	goto L9
L9:
	;
	v39 = int32(_a_F_InitDomainConstraintRef_1)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_InitDomainConstraintRef[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_InitDomainConstraintRef[0])) = v42
	if v34 == int32(0) {
		v88 = v5
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitDomainConstraintRef[0])) = v40
	v97 = v88
	goto L6
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v46 <= int32(0) {
		v88 = v5
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v50 = int32(0)
	v55 = v5
	goto L13
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v50<<(uint(int32(2))%32))))
	v63 = F_palloc0(m, int32(20))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v88 = v77
	goto L10
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(399)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v71
	v74 = F_ExecInitExpr(m, v71, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v74
	v77 = F_lappend(m, v55, v63)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v80 = v50 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v80 < v81 {
		v50 = v80
		v55 = v77
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
}
func F_domain_check_input(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
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
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v140 int32
	_ = v140
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int64
	_ = v244
	var v245 int64
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
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
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	v2 = l1
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v22 = l2 + int32(44)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+314)))
	if v24&int32(8) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+308))
	if v32 == v33 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+13)))
	if v27 != int32(100) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_load_domaintype_info(m, v23)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v161 == int32(0) {
		v310 = v20
		goto L31
	} else {
		goto L32
	}
L7:
	;
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v35
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v41 = v39 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v41
	if v41 <= v35 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v49 = v33
	goto L10
L10:
	;
	if v49 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	F_MemoryContextDelete(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+308))
	v49 = v48
	goto L10
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v54 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v53 + v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+12)))
	if v58 != v54 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v140
	goto L6
L17:
	;
	v140 = v57
	goto L16
L18:
	;
	goto L19
L19:
	;
	v61 = int32(_a_F_domain_check_input_0)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_domain_check_input[0]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_domain_check_input[0])) = v64
	if v57 == int32(0) {
		v123 = v5
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_domain_check_input[0])) = v62
	v140 = v123
	goto L16
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v68 <= int32(0) {
		v123 = v5
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v81 = int32(0)
	v82 = v5
	goto L23
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v81<<(uint(int32(2))%32))))
	v93 = F_palloc0(m, int32(20))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	v123 = v107
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(399)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v101
	v104 = F_ExecInitExpr(m, v101, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v104
	v107 = F_lappend(m, v82, v93)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v110 = v81 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v110 < v111 {
		v81 = v110
		v82 = v107
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	m.G0 = v18 + int32(48)
	return
L30:
	;
	F_ReScanExprContext(m, v327)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L74
	}
L31:
	;
	if v310 == int32(0) {
		goto L29
	} else {
		goto L73
	}
L32:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v164 <= int32(0) {
		v310 = v20
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v175 = v20
	v179 = v5
	goto L34
L34:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182+v179<<(uint(int32(2))%32))))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	switch v187 {
	case 0:
		goto L39
	case 1:
		goto L38
	default:
		goto L37
	}
L35:
	;
	v310 = v296
	goto L31
L36:
	;
	v299 = v179 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v299 < v300 {
		v175 = v296
		v179 = v299
		goto L34
	} else {
		goto L72
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L69
	}
L38:
	;
	if v175 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L39:
	;
	if v2 == int32(0) {
		v296 = v175
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v190 = F_errsave_start(m, l3)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v190 == int32(0) {
		v310 = v175
		goto L31
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v198 = F_format_type_be(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v198
	F_errmsg(m, int32(_a_F_domain_check_input_1), v18+int32(16))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_errdatatype(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errsave_finish(m, l3, int32(_a_F_domain_check_input_2), int32(160), int32(_a_F_domain_check_input_3))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v310 = v175
	goto L31
L48:
	;
	v216 = int32(_a_F_domain_check_input_0)
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_domain_check_input[0]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	*(*int32)(unsafe.Add(mBase, _c_F_domain_check_input[0])) = v219
	v221 = F_CreateStandaloneExprContext(m)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	v227 = v175
	goto L50
L50:
	;
	if v2 != 0 {
		v245 = l0
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_domain_check_input[0])) = v217
	*(*int32)(unsafe.Add(mBase, uint32(l2)+76)) = v221
	v227 = v221
	goto L50
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v227)+64)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(v227)+56)) = v245
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v249 = F_ExecCheck(m, v248, v227)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L59
	}
L53:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+8)))
	if v229 != int32(_a_F_domain_check_input_4) {
		v245 = l0
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v233 = base.I32_wrap_i64(l0)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v234 != int32(1) {
		v244 = l0
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v245 = v244
	goto L52
L56:
	;
	goto L55
L57:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	if v237 != int32(3) {
		v244 = l0
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v233)+2))
	v244 = base.I64_extend_i32_u(v240 + int32(18))
	goto L56
L59:
	;
	if v249 != 0 {
		v296 = v227
		goto L36
	} else {
		goto L60
	}
L60:
	;
	v251 = F_errsave_start(m, l3)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v251 == int32(0) {
		v327 = v227
		goto L30
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v259 = F_format_type_be(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v259
	F_errmsg(m, int32(_a_F_domain_check_input_5), v18+int32(32))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_errdatatype(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_err_generic_string(m, int32(110), v269)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errsave_finish(m, l3, int32(_a_F_domain_check_input_2), int32(200), int32(_a_F_domain_check_input_3))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v327 = v227
	goto L30
L69:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v285
	F_errmsg_internal(m, int32(_a_F_domain_check_input_6), v18)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_domain_check_input_2), int32(207), int32(_a_F_domain_check_input_3))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
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
	goto L35
L73:
	;
	v327 = v310
	goto L30
L74:
	;
	goto L29
}
