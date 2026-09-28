package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
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
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
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
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v389 int64
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
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
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	v13 = int64(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = int32(0)
	if v28 == v29 {
		v45 = v29
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v45&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	if v32 == int32(0) {
		v45 = v29
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v35 != int32(7) {
		v45 = v29
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v38 != int32(17) {
		v45 = v29
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+32)))
	v45 = v41 ^ int32(1)
	goto L4
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = F_get_fn_opclass_options(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v52 = int32(28)
	goto L11
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v54)
	v57 = base.I32_wrap_i64(v25) & int32(_a_F__ltree_consistent_0)
	switch v57 - int32(10) {
	case 0, 1:
		goto L20
	case 2, 3:
		goto L19
	case 4, 5:
		goto L18
	case 6, 7:
		goto L17
	default:
		goto L14
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v52 = v51
	goto L11
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L79
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L76
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L72
	}
L16:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v390 != v20 {
		goto L68
	} else {
		goto L69
	}
L17:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v229 = F_ArrayGetNItemsSafe(m, v226, v20+int32(16))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L44
	}
L18:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	if v209&int32(2) != 0 {
		v389 = int64(1)
		goto L16
	} else {
		goto L42
	}
L19:
	;
	v114 = int64(1)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	if v115&int32(2) != 0 {
		v389 = v114
		goto L16
	} else {
		goto L30
	}
L20:
	;
	v60 = int64(1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	if v61&int32(2) != 0 {
		v389 = v60
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v64 == int32(0) {
		v389 = v60
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v67 = int32(8)
	v74 = v20 + v67
	v75 = v64
	goto L23
L23:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74))))
	v89 = F_ltree_crc32_sz(m, v74+int32(2), v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v389 = v60
	goto L16
L25:
	;
	v91 = base.I32_rem_u_s(v89, v52<<(uint(int32(3))%32))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v67+int32(base.Ui32(v91)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v95)>>(uint(v91&int32(7))%32))&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v389 = int64(0)
	goto L16
L27:
	;
	goto L28
L28:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74))))
	v110 = int32(1)
	if v110 < v75 {
		v74 = v74 + (v104+int32(9))&int32(_a_F__ltree_consistent_1)
		v75 = v75 - v110
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v118 == int32(0) {
		v389 = v114
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v129 = v118
	v133 = v20 + int32(16)
	goto L32
L32:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+4)))
	if v140 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v389 = v114
	goto L16
L34:
	;
	v198 = int32(1)
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133))))
	if v198 < v129 {
		v129 = v129 - v198
		v133 = v133 + (v200+int32(7))&int32(_a_F__ltree_consistent_1)
		goto L32
	} else {
		goto L41
	}
L35:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+2)))
	if v143&int32(21) != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v149 = v133 + int32(16)
	v150 = v140
	goto L37
L37:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v162 = base.I32_rem_u_s(v161, v52<<(uint(int32(3))%32))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(8)+int32(base.Ui32(v162)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v166)>>(uint(v162&int32(7))%32))&int32(1) != 0 {
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v389 = int64(0)
	goto L16
L39:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+4)))
	v180 = int32(1)
	if v180 < v150 {
		v149 = v149 + (v172+int32(7))&int32(_a_F__ltree_consistent_1) + int32(8)
		v150 = v150 - v180
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	goto L33
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v52
	v213 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v53 + v213
	v222 = F_ltree_execute(m, v20+v213, v16+v213, int32(0), int32(_a_F__ltree_consistent_2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v389 = base.I64_extend_i32_u(v222)
	goto L16
L44:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(2) <= v231 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	v234 = F_array_contains_nulls(m, v20)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v234 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v229 <= int32(0) {
		v389 = v13
		goto L16
	} else {
		goto L48
	}
L48:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v238&int32(2) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v389 = int64(1)
	goto L16
L50:
	;
	if v225 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v247 = v225
	goto L53
L52:
	;
	v247 = (v226<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L53
L53:
	;
	v262 = v20 + v247
	v263 = v229
	goto L54
L54:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+4)))
	if v266 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L55:
	;
	goto L49
L56:
	;
	v272 = v266
	v277 = v262 + int32(16)
	goto L57
L57:
	;
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+4)))
	if v284 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L55
L59:
	;
	v353 = int32(1)
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277))))
	if v353 < v272 {
		v272 = v272 - v353
		v277 = v277 + (v355+int32(7))&int32(_a_F__ltree_consistent_1)
		goto L57
	} else {
		goto L67
	}
L60:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+2)))
	if v287&int32(21) != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v293 = v277 + int32(16)
	v294 = v284
	goto L62
L62:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v306 = base.I32_rem_u_s(v305, v52<<(uint(int32(3))%32))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(8)+int32(base.Ui32(v306)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v310)>>(uint(v306&int32(7))%32))&int32(1) != 0 {
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v336 = int32(1)
	if v336 < v263 {
		v262 = v262 + (int32(base.Ui32(v328)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v263 = v263 - v336
		goto L54
	} else {
		goto L66
	}
L64:
	;
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v293)+4)))
	v324 = int32(1)
	if v324 < v294 {
		v293 = v293 + (v316+int32(7))&int32(_a_F__ltree_consistent_1) + int32(8)
		v294 = v294 - v324
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v389 = v13
	goto L16
L67:
	;
	goto L58
L68:
	;
	F_pfree(m, v20)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	m.G0 = v16 + int32(16)
	return v389
L71:
	;
	goto L70
L72:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F__ltree_consistent_3), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F__ltree_consistent_4), int32(487), int32(_a_F__ltree_consistent_5))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v57
	F_errmsg_internal(m, int32(_a_F__ltree_consistent_6), v16)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F__ltree_consistent_4), int32(540), int32(_a_F__ltree_consistent_7))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errmsg(m, int32(_a_F__ltree_consistent_8), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F__ltree_consistent_4), int32(491), int32(_a_F__ltree_consistent_5))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltree_extract_isparent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14225(m, l0, int32(_a_F__ltree_extract_isparent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F__ltree_extract_risparent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14225(m, l0, int32(_a_F__ltree_extract_risparent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F__ltree_gist_options(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F__ltree_gist_options_0), int32(_a_F__ltree_gist_options_1), int32(28), int32(1), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_ltree_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v5 = int32(0)
	F_check_stack_depth(m)
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
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v12 == int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return (v85 ^ v86) & int32(1)
L4:
	;
	v15 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l1, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v17 = l0
	v19 = l2
	v22 = v5
	goto L8
L7:
	;
	v85 = v15
	v86 = v5
	goto L3
L8:
	;
	v26 = v17
	goto L10
L9:
	;
	v79 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l1, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	switch v33 - int32(33) {
	case 0:
		goto L16
	default:
		goto L14
	case 5:
		goto L15
	}
L11:
	;
	goto L9
L12:
	;
	goto L11
L13:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L24
	}
L14:
	;
	v58 = int32(1)
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+2)))
	v65 = F_ltree_execute(m, v26+v59*int32(12), l1, v19&v58, l3)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+2)))
	v55 = F_ltree_execute(m, v26+v51*int32(12), l1, v19&int32(1), l3)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L16:
	;
	v36 = int32(1)
	if v19&v36 == int32(0) {
		v85 = v36
		v86 = v22
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v41 = int32(1)
	v43 = v22 ^ v41
	F_check_stack_depth(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+12)))
	v48 = v26 + int32(12)
	if v46 != int32(2) {
		v17 = v48
		v19 = v41
		v22 = v43
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v75 = v48
	v78 = v43
	goto L12
L20:
	;
	if v55 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v85 = int32(0)
	v86 = v22
	goto L3
L22:
	;
	if v65 != 0 {
		v85 = v58
		v86 = v22
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L13
L24:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+12)))
	v72 = v26 + int32(12)
	if v70 != int32(2) {
		v26 = v72
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v75 = v72
	v78 = v22
	goto L12
L26:
	;
	v85 = v79
	v86 = v78
	goto L3
}
func F_ltree_gist_alloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l4 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	v19 = int32(8)
	if l0 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v12 = int32(base.Ui32(v8) >> (uint(int32(2)) % 32))
	goto L6
L5:
	;
	v12 = int32(0)
	goto L6
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v18 = v12 + int32(base.Ui32(v13)>>(uint(int32(2))%32))
	goto L3
L7:
	;
	v22 = v19
	goto L9
L8:
	;
	v22 = l2 + v19
	goto L9
L9:
	;
	v23 = v18 + v22
	v24 = F_palloc(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v23 << (uint(int32(2)) % 32)
	if l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return v24
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(0)
	if l0 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v181 = int32(base.Ui32(v179) >> (uint(int32(2)) % 32))
	if v181 == int32(0) {
		goto L12
	} else {
		goto L60
	}
L16:
	;
	if l3 == int32(0) {
		goto L12
	} else {
		goto L25
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(2)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v36 = v24 + int32(8)
	if l1 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if l2 == int32(0) {
		goto L16
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if l2 == int32(0) {
		goto L16
	} else {
		goto L24
	}
L23:
	;
	base.MemoryCopy(m, v36, l1, l2)
	goto L16
L24:
	;
	base.MemoryFill(m, v36, int32(0), l2)
	goto L16
L25:
	;
	v48 = v24 + int32(8)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v51 = int32(base.Ui32(v49) >> (uint(int32(2)) % 32))
	if v51 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if l0 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v56 = int32(0)
	if base.B2i32(l4 == v56)|base.B2i32(l3 == l4) == v56 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v53 = int32(0)
	goto L31
L30:
	;
	v53 = l2
	goto L31
L31:
	;
	base.MemoryCopy(m, v48+v53, l3, v51)
	goto L28
L32:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v156&int32(2) != 0 {
		goto L53
	} else {
		goto L54
	}
L33:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if v62 != v63 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v150 | int32(4)
	return v24
L36:
	;
	v65 = int32(0)
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if base.B2i32(v72 == v65)|base.B2i32(v75 == v65) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if v149 != 0 {
		goto L32
	} else {
		goto L52
	}
L38:
	;
	v149 = v139
	goto L37
L39:
	;
	v139 = v72 - v75
	goto L38
L40:
	;
	v79 = int32(8)
	v83 = l3 + v79
	v84 = l4 + v79
	v87 = v75
	v88 = v72
	goto L41
L41:
	;
	v92 = int32(2)
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83))))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
	if base.Ui32(v96) < base.Ui32(v97) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L39
L43:
	;
	v99 = v96
	goto L45
L44:
	;
	v99 = v97
	goto L45
L45:
	;
	v100 = F_memcmp(m, v83+v92, v84+v92, v99)
	mBase = m.M
	if v100 != 0 {
		v139 = v100
		goto L38
	} else {
		goto L46
	}
L46:
	;
	if v96 != v97 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v149 = v96 - v97
	goto L37
L48:
	;
	goto L49
L49:
	;
	if v88 < int32(2) {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	v105 = int32(1)
	v107 = int32(9)
	v109 = int32(_a_F_ltree_gist_alloc_0)
	if v105 < v87 {
		v83 = v83 + (v96+v107)&v109
		v84 = v84 + (v97+v107)&v109
		v87 = v87 - v105
		v88 = v88 - v105
		goto L41
	} else {
		goto L51
	}
L51:
	;
	goto L42
L52:
	;
	goto L35
L53:
	;
	v159 = int32(0)
	goto L55
L54:
	;
	v159 = l2
	goto L55
L55:
	;
	v160 = v48 + v159
	if v156&int32(4) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v169 = v160 + int32(base.Ui32(v165)>>(uint(int32(2))%32))
	goto L58
L57:
	;
	v169 = v160
	goto L58
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v172 = int32(base.Ui32(v170) >> (uint(int32(2)) % 32))
	if v172 == int32(0) {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	base.MemoryCopy(m, v169, l4, v172)
	return v24
L60:
	;
	base.MemoryCopy(m, v24+int32(8), l3, v181)
	goto L12
}
func F_ltree_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v23 = int32(0)
	if base.B2i32(v22 == v23)|base.B2i32(v21 == v23) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v159 != v14 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v157 = v22 - v21
	goto L4
L6:
	;
	v28 = int32(8)
	v35 = v14 + v28
	v36 = v19 + v28
	v40 = v21
	v41 = v22
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v35 + v44
	v47 = v36 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v51 = v48
	goto L11
L10:
	;
	v51 = v49
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v113 != 0 {
		v157 = v113
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v113 = int32(0)
	goto L12
L14:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L24
L15:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L17
L17:
	;
	if v77 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L17
L21:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L14
L24:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v113 = v92 - v93
	goto L12
L26:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L13
L30:
	;
	if v48 != v49 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v157 = v48 - v49
	goto L4
L32:
	;
	goto L33
L33:
	;
	if v41 < int32(2) {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v118 = int32(1)
	v120 = int32(9)
	v122 = int32(_a_F_ltree_gt_0)
	if v118 < v40 {
		v35 = v35 + (v48+v120)&v122
		v36 = v36 + (v49+v120)&v122
		v40 = v40 - v118
		v41 = v41 - v118
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L8
L36:
	;
	F_pfree(m, v14)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v163 != v19 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v19)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return base.I64_extend_i32_u(base.B2i32(int32(0) < v157))
L43:
	;
	goto L42
}
func F_ltree_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_parse_ltree(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v5)
		}
	}
}
func F_ltree_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
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
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
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
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v649 int32
	_ = v649
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v742 int32
	_ = v742
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
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
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v924 int32
	_ = v924
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v1014 int32
	_ = v1014
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	v2 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = base.I32_wrap_i64(v25)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v28 == v2 {
		v45 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v45&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	if v32 == int32(0) {
		v45 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v35 != int32(7) {
		v45 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v38 != int32(17) {
		v45 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+32)))
	v45 = v41 ^ int32(1)
	goto L2
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = F_get_fn_opclass_options(m, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v54 = int32(8)
	goto L9
L9:
	;
	v55 = F_palloc0(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int64(0)
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v54 = v53
	goto L9
L12:
	;
	v57 = F_palloc0(m, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v60 = int32(1)
	v63 = (v59 - v60) & int32(_a_F_ltree_picksplit_0)
	v67 = v63<<(uint(v60)%32) + int32(4)
	v68 = F_palloc(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v68
	v71 = F_palloc(m, v67)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v71
	v81 = F_palloc_mul(m, int32(8), v63+int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v59&int32(_a_F_ltree_picksplit_0) == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v949 = int32(1)
	v952 = base.B2i32(v54 <= int32(0))
	if v952|v932 != 0 {
		v989 = v949
		goto L146
	} else {
		goto L147
	}
L18:
	;
	v87 = int32(8)
	F_pg_qsort(m, v81+v87, v63, v87, int32(_a_F_ltree_picksplit_1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v96 = int32(1)
	goto L22
L21:
	;
	v932 = v2
	v934 = v2
	v937 = v2
	v939 = v2
	goto L17
L22:
	;
	v119 = int32(3)
	v121 = v81 + v96<<(uint(v119)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v96
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(8)+v96*int32(24))))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v128&v119 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v141 = int32(8)
	F_pg_qsort(m, v81+v141, v63, v141, int32(_a_F_ltree_picksplit_1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L28
	}
L24:
	;
	v131 = int32(0)
	goto L26
L25:
	;
	v131 = v54
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v126 + v131 + int32(8)
	v139 = (v96 + int32(1)) & int32(_a_F_ltree_picksplit_0)
	if base.Ui32(v139) <= base.Ui32(v63) {
		v96 = v139
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v148 = v54 & int32(2147483644)
	v149 = int32(3)
	v150 = v54 & v149
	v152 = v54 << (uint(v149) % 32)
	v153 = int32(1)
	v164 = v2
	v166 = v2
	v169 = v2
	v171 = v2
	v175 = v153
	goto L29
L29:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v81+v175<<(uint(int32(3))%32))))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(8)+v184*int32(24))))
	if base.Ui32(v175) <= base.Ui32(int32(base.Ui32(v63)>>(uint(v153)%32))) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v932 = v904
	v934 = v906
	v937 = v909
	v939 = v911
	goto L17
L31:
	;
	v924 = (v175 + int32(1)) & int32(_a_F_ltree_picksplit_0)
	if base.Ui32(v924) <= base.Ui32(v63) {
		v164 = v904
		v166 = v906
		v169 = v909
		v171 = v911
		v175 = v924
		goto L29
	} else {
		goto L145
	}
L32:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v192 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v190+v191<<(uint(v192)%32)))) = uint16(v184)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v191 + v192
	if v169 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v546 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v544+v545<<(uint(v546)%32)))) = uint16(v184)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v545 + v546
	if v171 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L35:
	;
	if v329&int32(1) != 0 {
		goto L67
	} else {
		goto L68
	}
L36:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v311&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L37:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v203&int32(1) != 0 {
		v220 = v188 + int32(8)
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v221 = int32(0)
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	if base.B2i32(v228 == v221)|base.B2i32(v231 == v221) != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	if v203&int32(2) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v209 = int32(0)
	goto L42
L41:
	;
	v209 = v54
	goto L42
L42:
	;
	v212 = v188 + v209 + int32(8)
	if v203&int32(4) != 0 {
		v220 = v212
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v220 = v212 + int32(base.Ui32(v215)>>(uint(int32(2))%32))
	goto L38
L44:
	;
	if int32(0) < v305 {
		goto L36
	} else {
		goto L59
	}
L45:
	;
	v305 = v295
	goto L44
L46:
	;
	v295 = v228 - v231
	goto L45
L47:
	;
	v235 = int32(8)
	v239 = v220 + v235
	v240 = v169 + v235
	v243 = v231
	v244 = v228
	goto L48
L48:
	;
	v248 = int32(2)
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239))))
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240))))
	if base.Ui32(v252) < base.Ui32(v253) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L46
L50:
	;
	v255 = v252
	goto L52
L51:
	;
	v255 = v253
	goto L52
L52:
	;
	v256 = F_memcmp(m, v239+v248, v240+v248, v255)
	mBase = m.M
	if v256 != 0 {
		v295 = v256
		goto L45
	} else {
		goto L53
	}
L53:
	;
	if v252 != v253 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v305 = v252 - v253
	goto L44
L55:
	;
	goto L56
L56:
	;
	if v244 < int32(2) {
		goto L46
	} else {
		goto L57
	}
L57:
	;
	v261 = int32(1)
	v263 = int32(9)
	v265 = int32(_a_F_ltree_picksplit_2)
	if v261 < v243 {
		v239 = v239 + (v252+v263)&v265
		v240 = v240 + (v253+v263)&v265
		v243 = v243 - v261
		v244 = v244 - v261
		goto L48
	} else {
		goto L58
	}
L58:
	;
	goto L49
L59:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v329 = v308
	v331 = v169
	goto L35
L60:
	;
	v329 = v311
	v331 = v188 + int32(8)
	goto L35
L61:
	;
	goto L62
L62:
	;
	if v311&int32(2) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v319 = int32(0)
	goto L65
L64:
	;
	v319 = v54
	goto L65
L65:
	;
	v322 = v188 + v319 + int32(8)
	if v311&int32(4) != 0 {
		v329 = v311
		v331 = v322
		goto L35
	} else {
		goto L66
	}
L66:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v329 = v311
	v331 = v322 + int32(base.Ui32(v325)>>(uint(int32(2))%32))
	goto L35
L67:
	;
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+12)))
	if v334 == int32(0) {
		v904 = v164
		v906 = v166
		v909 = v331
		v911 = v171
		goto L31
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v388 = int32(1)
	if (v164|int32(base.Ui32(v329)>>(uint(v388)%32)))&v388 != 0 {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v339 = v188 + int32(16)
	v340 = v334
	goto L71
L71:
	;
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339))))
	v365 = F_ltree_crc32_sz(m, v339+int32(2), v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L10
	} else {
		goto L73
	}
L72:
	;
	v904 = v164
	v906 = v166
	v909 = v331
	v911 = v171
	goto L31
L73:
	;
	v367 = base.I32_rem_u_s(v365, v152)
	v370 = v55 + int32(base.Ui32(v367)>>(uint(int32(3))%32))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	v372 = int32(1)
	v376 = v371 | v372<<(uint(v367&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v370))) = uint8(v376)
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339))))
	if base.Ui32(v372) < base.Ui32(v340) {
		v339 = v339 + (v378+int32(9))&int32(_a_F_ltree_picksplit_2)
		v340 = v340 - v372
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v904 = int32(1)
	v906 = v166
	v909 = v331
	v911 = v171
	goto L31
L76:
	;
	goto L77
L77:
	;
	if v54 <= int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v904 = int32(0)
	v906 = v166
	v909 = v331
	v911 = v171
	goto L31
L79:
	;
	v397 = v188 + int32(8)
	v398 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v54) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v403 = v398
	v404 = v398
	goto L83
L81:
	;
	v464 = v398
	goto L82
L82:
	;
	v487 = v464
	v490 = v398
	goto L87
L83:
	;
	v426 = v403 + v55
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v397))))
	v430 = v427 | v429
	*(*uint8)(unsafe.Add(mBase, uint32(v426))) = uint8(v430)
	v433 = v403 | int32(1)
	v434 = v55 + v433
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v433))))
	v438 = v435 | v437
	*(*uint8)(unsafe.Add(mBase, uint32(v434))) = uint8(v438)
	v441 = v403 | int32(2)
	v442 = v55 + v441
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v441))))
	v446 = v443 | v445
	*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v446)
	v449 = v403 | int32(3)
	v450 = v55 + v449
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v449))))
	v454 = v451 | v453
	*(*uint8)(unsafe.Add(mBase, uint32(v450))) = uint8(v454)
	v456 = int32(4)
	v457 = v403 + v456
	v459 = v404 + v456
	if v459 != v148 {
		v403 = v457
		v404 = v459
		goto L83
	} else {
		goto L85
	}
L84:
	;
	if v150 == int32(0) {
		goto L78
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v464 = v457
	goto L82
L87:
	;
	v509 = v487 + v55
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487+v397))))
	v513 = v510 | v512
	*(*uint8)(unsafe.Add(mBase, uint32(v509))) = uint8(v513)
	v515 = int32(1)
	v518 = v490 + v515
	if v518 != v150 {
		v487 = v487 + v515
		v490 = v518
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L78
L89:
	;
	goto L88
L90:
	;
	if v683&int32(1) != 0 {
		goto L122
	} else {
		goto L123
	}
L91:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v665&int32(1) != 0 {
		goto L115
	} else {
		goto L116
	}
L92:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v557&int32(1) != 0 {
		v574 = v188 + int32(8)
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v575 = int32(0)
	v582 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v574)+4)))
	v585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+4)))
	if base.B2i32(v582 == v575)|base.B2i32(v585 == v575) != 0 {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	if v557&int32(2) != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v563 = int32(0)
	goto L97
L96:
	;
	v563 = v54
	goto L97
L97:
	;
	v566 = v188 + v563 + int32(8)
	if v557&int32(4) != 0 {
		v574 = v566
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v574 = v566 + int32(base.Ui32(v569)>>(uint(int32(2))%32))
	goto L93
L99:
	;
	if int32(0) < v659 {
		goto L91
	} else {
		goto L114
	}
L100:
	;
	v659 = v649
	goto L99
L101:
	;
	v649 = v582 - v585
	goto L100
L102:
	;
	v589 = int32(8)
	v593 = v574 + v589
	v594 = v171 + v589
	v597 = v585
	v598 = v582
	goto L103
L103:
	;
	v602 = int32(2)
	v606 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v593))))
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v594))))
	if base.Ui32(v606) < base.Ui32(v607) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L101
L105:
	;
	v609 = v606
	goto L107
L106:
	;
	v609 = v607
	goto L107
L107:
	;
	v610 = F_memcmp(m, v593+v602, v594+v602, v609)
	mBase = m.M
	if v610 != 0 {
		v649 = v610
		goto L100
	} else {
		goto L108
	}
L108:
	;
	if v606 != v607 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v659 = v606 - v607
	goto L99
L110:
	;
	goto L111
L111:
	;
	if v598 < int32(2) {
		goto L101
	} else {
		goto L112
	}
L112:
	;
	v615 = int32(1)
	v617 = int32(9)
	v619 = int32(_a_F_ltree_picksplit_2)
	if v615 < v597 {
		v593 = v593 + (v606+v617)&v619
		v594 = v594 + (v607+v617)&v619
		v597 = v597 - v615
		v598 = v598 - v615
		goto L103
	} else {
		goto L113
	}
L113:
	;
	goto L104
L114:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v683 = v662
	v685 = v171
	goto L90
L115:
	;
	v683 = v665
	v685 = v188 + int32(8)
	goto L90
L116:
	;
	goto L117
L117:
	;
	if v665&int32(2) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v673 = int32(0)
	goto L120
L119:
	;
	v673 = v54
	goto L120
L120:
	;
	v676 = v188 + v673 + int32(8)
	if v665&int32(4) != 0 {
		v683 = v665
		v685 = v676
		goto L90
	} else {
		goto L121
	}
L121:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v676)))
	v683 = v665
	v685 = v676 + int32(base.Ui32(v679)>>(uint(int32(2))%32))
	goto L90
L122:
	;
	v688 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+12)))
	if v688 == int32(0) {
		v904 = v164
		v906 = v166
		v909 = v169
		v911 = v685
		goto L31
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v742 = int32(1)
	if (v166|int32(base.Ui32(v683)>>(uint(v742)%32)))&v742 != 0 {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v693 = v188 + int32(16)
	v694 = v688
	goto L126
L126:
	;
	v718 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v693))))
	v719 = F_ltree_crc32_sz(m, v693+int32(2), v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L10
	} else {
		goto L128
	}
L127:
	;
	v904 = v164
	v906 = v166
	v909 = v169
	v911 = v685
	goto L31
L128:
	;
	v721 = base.I32_rem_u_s(v719, v152)
	v724 = v57 + int32(base.Ui32(v721)>>(uint(int32(3))%32))
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724))))
	v726 = int32(1)
	v730 = v725 | v726<<(uint(v721&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v724))) = uint8(v730)
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v693))))
	if base.Ui32(v726) < base.Ui32(v694) {
		v693 = v693 + (v732+int32(9))&int32(_a_F_ltree_picksplit_2)
		v694 = v694 - v726
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v904 = v164
	v906 = int32(1)
	v909 = v169
	v911 = v685
	goto L31
L131:
	;
	goto L132
L132:
	;
	if v54 <= int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v904 = v164
	v906 = int32(0)
	v909 = v169
	v911 = v685
	goto L31
L134:
	;
	v751 = v188 + int32(8)
	v752 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v54) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v757 = v752
	v758 = v752
	goto L138
L136:
	;
	v818 = v752
	goto L137
L137:
	;
	v841 = v818
	v844 = v752
	goto L142
L138:
	;
	v780 = v757 + v57
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780))))
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757+v751))))
	v784 = v781 | v783
	*(*uint8)(unsafe.Add(mBase, uint32(v780))) = uint8(v784)
	v787 = v757 | int32(1)
	v788 = v57 + v787
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788))))
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751+v787))))
	v792 = v789 | v791
	*(*uint8)(unsafe.Add(mBase, uint32(v788))) = uint8(v792)
	v795 = v757 | int32(2)
	v796 = v57 + v795
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796))))
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751+v795))))
	v800 = v797 | v799
	*(*uint8)(unsafe.Add(mBase, uint32(v796))) = uint8(v800)
	v803 = v757 | int32(3)
	v804 = v57 + v803
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751+v803))))
	v808 = v805 | v807
	*(*uint8)(unsafe.Add(mBase, uint32(v804))) = uint8(v808)
	v810 = int32(4)
	v811 = v757 + v810
	v813 = v758 + v810
	if v813 != v148 {
		v757 = v811
		v758 = v813
		goto L138
	} else {
		goto L140
	}
L139:
	;
	if v150 == int32(0) {
		goto L133
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v818 = v811
	goto L137
L142:
	;
	v863 = v841 + v57
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841+v751))))
	v867 = v864 | v866
	*(*uint8)(unsafe.Add(mBase, uint32(v863))) = uint8(v867)
	v869 = int32(1)
	v872 = v844 + v869
	if v872 != v150 {
		v841 = v841 + v869
		v844 = v872
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L133
L144:
	;
	goto L143
L145:
	;
	goto L30
L146:
	;
	if v952|v934 != 0 {
		v1047 = v949
		goto L152
	} else {
		goto L153
	}
L147:
	;
	v955 = int32(0)
	goto L148
L148:
	;
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955+v55))))
	v980 = int32(255)
	v981 = base.B2i32(v979 == v980)
	if v979 != v980 {
		v989 = v981
		goto L146
	} else {
		goto L150
	}
L149:
	;
	v989 = v981
	goto L146
L150:
	;
	v985 = v955 + int32(1)
	if v985 != v54 {
		v955 = v985
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v1070 = v24 + int32(8)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1070+v1071*int32(24))))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+4))
	if v1077&int32(3) != 0 {
		goto L158
	} else {
		goto L159
	}
L153:
	;
	v1014 = int32(0)
	goto L154
L154:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014+v57))))
	v1039 = int32(255)
	v1040 = base.B2i32(v1038 == v1039)
	if v1038 != v1039 {
		v1047 = v1040
		goto L152
	} else {
		goto L156
	}
L155:
	;
	v1047 = v1040
	goto L152
L156:
	;
	v1044 = v1014 + int32(1)
	if v1044 != v54 {
		v1014 = v1044
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v1080 = int32(0)
	goto L160
L159:
	;
	v1080 = v54
	goto L160
L160:
	;
	v1084 = F_ltree_gist_alloc(m, v989, v55, v54, v1075+v1080+int32(8), v937)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L10
	} else {
		goto L161
	}
L161:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v81+v63<<(uint(int32(2))%32)&int32(_a_F_ltree_picksplit_3))+8))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1070+v1091*int32(24))))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+4))
	if v1097&int32(3) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1100 = int32(0)
	goto L164
L163:
	;
	v1100 = v54
	goto L164
L164:
	;
	v1104 = F_ltree_gist_alloc(m, v1047, v57, v54, v1095+v1100+int32(8), v939)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L10
	} else {
		goto L165
	}
L165:
	;
	F_pfree(m, v55)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L10
	} else {
		goto L166
	}
L166:
	;
	F_pfree(m, v57)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = base.I64_extend_i32_u(v1104)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = base.I64_extend_i32_u(v1084)
	return v25 & int64(4294967295)
}
func F_ltree_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v493 int32
	_ = v493
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v617 int32
	_ = v617
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v720 int32
	_ = v720
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v826 int32
	_ = v826
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	v2 = int32(0)
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v25 == v2 {
		v42 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v42&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	if v29 == int32(0) {
		v42 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32 != int32(7) {
		v42 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v35 != int32(17) {
		v42 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+32)))
	v42 = v38 ^ int32(1)
	goto L2
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = F_get_fn_opclass_options(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v51 = int32(8)
	goto L9
L9:
	;
	v52 = F_palloc0(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int64(0)
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v51 = v50
	goto L9
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if int32(0) < v55 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v60 = int32(3)
	v61 = v51 & v60
	v72 = v2
	v74 = v2
	v80 = v2
	v81 = v2
	goto L16
L14:
	;
	v764 = v2
	v766 = v2
	v773 = v2
	goto L15
L15:
	;
	if v773&int32(1)|base.B2i32(v51 <= int32(0)) != 0 {
		v826 = int32(1)
		goto L133
	} else {
		goto L134
	}
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(8)+v80*int32(24))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v91&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v764 = v739
	v766 = v741
	v773 = v748
	goto L15
L18:
	;
	v755 = v80 + int32(1)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v755 < v756 {
		v72 = v739
		v74 = v741
		v80 = v755
		v81 = v748
		goto L16
	} else {
		goto L132
	}
L19:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+12)))
	if v94 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v346 = v91 & int32(2)
	v347 = int32(1)
	v349 = v81 | int32(base.Ui32(v346)>>(uint(v347)%32))
	if v349&v347 != 0 {
		goto L69
	} else {
		goto L70
	}
L22:
	;
	v97 = v90 + int32(16)
	v98 = v94
	goto L25
L23:
	;
	goto L24
L24:
	;
	v166 = v90 + int32(8)
	if v74 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	v121 = F_ltree_crc32_sz(m, v97+int32(2), v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L27
	}
L26:
	;
	goto L24
L27:
	;
	v123 = base.I32_rem_u_s(v121, v51<<(uint(v60)%32))
	v126 = v52 + int32(base.Ui32(v123)>>(uint(int32(3))%32))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v128 = int32(1)
	v132 = v127 | v128<<(uint(v123&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v132)
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	if base.Ui32(v128) < base.Ui32(v98) {
		v97 = v97 + (v134+int32(9))&int32(_a_F_ltree_union_0)
		v98 = v98 - v128
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v72 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L30:
	;
	v167 = int32(0)
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)))
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	if base.B2i32(v174 == v167)|base.B2i32(v177 == v167) != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	goto L32
L32:
	;
	v254 = v166
	goto L29
L33:
	;
	if v251 <= int32(0) {
		v254 = v74
		goto L29
	} else {
		goto L48
	}
L34:
	;
	v251 = v241
	goto L33
L35:
	;
	v241 = v174 - v177
	goto L34
L36:
	;
	v185 = v74 + int32(8)
	v186 = v90 + int32(16)
	v189 = v177
	v190 = v174
	goto L37
L37:
	;
	v194 = int32(2)
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185))))
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186))))
	if base.Ui32(v198) < base.Ui32(v199) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L35
L39:
	;
	v201 = v198
	goto L41
L40:
	;
	v201 = v199
	goto L41
L41:
	;
	v202 = F_memcmp(m, v185+v194, v186+v194, v201)
	mBase = m.M
	if v202 != 0 {
		v241 = v202
		goto L34
	} else {
		goto L42
	}
L42:
	;
	if v198 != v199 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v251 = v198 - v199
	goto L33
L44:
	;
	goto L45
L45:
	;
	if v190 < int32(2) {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	v207 = int32(1)
	v209 = int32(9)
	v211 = int32(_a_F_ltree_union_0)
	if v207 < v189 {
		v185 = v185 + (v198+v209)&v211
		v186 = v186 + (v199+v209)&v211
		v189 = v189 - v207
		v190 = v190 - v207
		goto L37
	} else {
		goto L47
	}
L47:
	;
	goto L38
L48:
	;
	goto L32
L49:
	;
	v739 = v166
	v741 = v254
	v748 = v81
	goto L18
L50:
	;
	goto L51
L51:
	;
	v257 = int32(0)
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	if base.B2i32(v264 == v257)|base.B2i32(v267 == v257) != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if int32(0) <= v341 {
		v739 = v72
		v741 = v254
		v748 = v81
		goto L18
	} else {
		goto L67
	}
L53:
	;
	v341 = v331
	goto L52
L54:
	;
	v331 = v264 - v267
	goto L53
L55:
	;
	v275 = v72 + int32(8)
	v276 = v90 + int32(16)
	v279 = v267
	v280 = v264
	goto L56
L56:
	;
	v284 = int32(2)
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275))))
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v276))))
	if base.Ui32(v288) < base.Ui32(v289) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L54
L58:
	;
	v291 = v288
	goto L60
L59:
	;
	v291 = v289
	goto L60
L60:
	;
	v292 = F_memcmp(m, v275+v284, v276+v284, v291)
	mBase = m.M
	if v292 != 0 {
		v331 = v292
		goto L53
	} else {
		goto L61
	}
L61:
	;
	if v288 != v289 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v341 = v288 - v289
	goto L52
L63:
	;
	goto L64
L64:
	;
	if v280 < int32(2) {
		goto L54
	} else {
		goto L65
	}
L65:
	;
	v297 = int32(1)
	v299 = int32(9)
	v301 = int32(_a_F_ltree_union_0)
	if v297 < v279 {
		v275 = v275 + (v288+v299)&v301
		v276 = v276 + (v289+v299)&v301
		v279 = v279 - v297
		v280 = v280 - v297
		goto L56
	} else {
		goto L66
	}
L66:
	;
	goto L57
L67:
	;
	v739 = v166
	v741 = v254
	v748 = v81
	goto L18
L68:
	;
	v541 = v90 + int32(8)
	v542 = v541 + v519
	if v74 != 0 {
		goto L88
	} else {
		goto L89
	}
L69:
	;
	v517 = v346
	goto L71
L70:
	;
	if v51 <= int32(0) {
		v519 = v51
		goto L68
	} else {
		goto L72
	}
L71:
	;
	if v517 != 0 {
		goto L84
	} else {
		goto L85
	}
L72:
	;
	v355 = v90 + int32(8)
	v356 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v517 = v493 & int32(2)
	goto L71
L74:
	;
	v361 = v356
	v362 = v356
	goto L77
L75:
	;
	v420 = v356
	goto L76
L76:
	;
	v441 = v420
	v450 = v356
	goto L81
L77:
	;
	v382 = v361 + v52
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361+v355))))
	v386 = v383 | v385
	*(*uint8)(unsafe.Add(mBase, uint32(v382))) = uint8(v386)
	v389 = v361 | int32(1)
	v390 = v52 + v389
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v389))))
	v394 = v391 | v393
	*(*uint8)(unsafe.Add(mBase, uint32(v390))) = uint8(v394)
	v397 = v361 | int32(2)
	v398 = v52 + v397
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v397))))
	v402 = v399 | v401
	*(*uint8)(unsafe.Add(mBase, uint32(v398))) = uint8(v402)
	v405 = v361 | int32(3)
	v406 = v52 + v405
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v405))))
	v410 = v407 | v409
	*(*uint8)(unsafe.Add(mBase, uint32(v406))) = uint8(v410)
	v412 = int32(4)
	v413 = v361 + v412
	v415 = v362 + v412
	if v415 != v51&int32(2147483644) {
		v361 = v413
		v362 = v415
		goto L77
	} else {
		goto L79
	}
L78:
	;
	if v61 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v420 = v413
	goto L76
L81:
	;
	v461 = v441 + v52
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441+v355))))
	v465 = v462 | v464
	*(*uint8)(unsafe.Add(mBase, uint32(v461))) = uint8(v465)
	v467 = int32(1)
	v470 = v450 + v467
	if v470 != v61 {
		v441 = v441 + v467
		v450 = v470
		goto L81
	} else {
		goto L83
	}
L82:
	;
	goto L73
L83:
	;
	goto L82
L84:
	;
	v518 = int32(0)
	goto L86
L85:
	;
	v518 = v51
	goto L86
L86:
	;
	v519 = v518
	goto L68
L87:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v632&int32(2) != 0 {
		goto L107
	} else {
		goto L108
	}
L88:
	;
	v543 = int32(0)
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)))
	v553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v542)+4)))
	if base.B2i32(v550 == v543)|base.B2i32(v553 == v543) != 0 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	goto L90
L90:
	;
	v630 = v542
	goto L87
L91:
	;
	if v627 <= int32(0) {
		v630 = v74
		goto L87
	} else {
		goto L106
	}
L92:
	;
	v627 = v617
	goto L91
L93:
	;
	v617 = v550 - v553
	goto L92
L94:
	;
	v557 = int32(8)
	v561 = v74 + v557
	v562 = v542 + v557
	v565 = v553
	v566 = v550
	goto L95
L95:
	;
	v570 = int32(2)
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v561))))
	v575 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v562))))
	if base.Ui32(v574) < base.Ui32(v575) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L93
L97:
	;
	v577 = v574
	goto L99
L98:
	;
	v577 = v575
	goto L99
L99:
	;
	v578 = F_memcmp(m, v561+v570, v562+v570, v577)
	mBase = m.M
	if v578 != 0 {
		v617 = v578
		goto L92
	} else {
		goto L100
	}
L100:
	;
	if v574 != v575 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v627 = v574 - v575
	goto L91
L102:
	;
	goto L103
L103:
	;
	if v566 < int32(2) {
		goto L93
	} else {
		goto L104
	}
L104:
	;
	v583 = int32(1)
	v585 = int32(9)
	v587 = int32(_a_F_ltree_union_0)
	if v583 < v565 {
		v561 = v561 + (v574+v585)&v587
		v562 = v562 + (v575+v585)&v587
		v565 = v565 - v583
		v566 = v566 - v583
		goto L95
	} else {
		goto L105
	}
L105:
	;
	goto L96
L106:
	;
	goto L90
L107:
	;
	v635 = int32(0)
	goto L109
L108:
	;
	v635 = v51
	goto L109
L109:
	;
	v636 = v541 + v635
	if v632&int32(4) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	v645 = v636 + int32(base.Ui32(v641)>>(uint(int32(2))%32))
	goto L112
L111:
	;
	v645 = v636
	goto L112
L112:
	;
	if v72 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v646 = int32(0)
	v653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	v656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v645)+4)))
	if base.B2i32(v653 == v646)|base.B2i32(v656 == v646) != 0 {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	goto L115
L115:
	;
	v739 = v645
	v741 = v630
	v748 = v349
	goto L18
L116:
	;
	if int32(0) <= v730 {
		v739 = v72
		v741 = v630
		v748 = v349
		goto L18
	} else {
		goto L131
	}
L117:
	;
	v730 = v720
	goto L116
L118:
	;
	v720 = v653 - v656
	goto L117
L119:
	;
	v660 = int32(8)
	v664 = v72 + v660
	v665 = v645 + v660
	v668 = v656
	v669 = v653
	goto L120
L120:
	;
	v673 = int32(2)
	v677 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v664))))
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v665))))
	if base.Ui32(v677) < base.Ui32(v678) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L118
L122:
	;
	v680 = v677
	goto L124
L123:
	;
	v680 = v678
	goto L124
L124:
	;
	v681 = F_memcmp(m, v664+v673, v665+v673, v680)
	mBase = m.M
	if v681 != 0 {
		v720 = v681
		goto L117
	} else {
		goto L125
	}
L125:
	;
	if v677 != v678 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v730 = v677 - v678
	goto L116
L127:
	;
	goto L128
L128:
	;
	if v669 < int32(2) {
		goto L118
	} else {
		goto L129
	}
L129:
	;
	v686 = int32(1)
	v688 = int32(9)
	v690 = int32(_a_F_ltree_union_0)
	if v686 < v668 {
		v664 = v664 + (v677+v688)&v690
		v665 = v665 + (v678+v688)&v690
		v668 = v668 - v686
		v669 = v669 - v686
		goto L120
	} else {
		goto L130
	}
L130:
	;
	goto L121
L131:
	;
	goto L115
L132:
	;
	goto L17
L133:
	;
	v836 = F_ltree_gist_alloc(m, v826, v52, v51, v766, v764)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L10
	} else {
		goto L139
	}
L134:
	;
	v800 = v2
	goto L135
L135:
	;
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v800))))
	v807 = int32(255)
	v808 = base.B2i32(v806 == v807)
	if v806 != v807 {
		v826 = v808
		goto L133
	} else {
		goto L137
	}
L136:
	;
	v826 = v808
	goto L133
L137:
	;
	v812 = v800 + int32(1)
	if v812 != v51 {
		v800 = v812
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v836)))
	*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v22)))) = int32(base.Ui32(v838) >> (uint(int32(2)) % 32))
	return base.I64_extend_i32_u(v836)
}
