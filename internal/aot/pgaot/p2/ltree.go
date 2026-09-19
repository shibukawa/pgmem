package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = int32(0)
	if v26 == v27 {
		v43 = v27
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v43&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v30 == int32(0) {
		v43 = v27
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v33 != int32(7) {
		v43 = v27
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v36 != int32(17) {
		v43 = v27
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
	v43 = v39 ^ int32(1)
	goto L4
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = F_get_fn_opclass_options(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v50 = int32(28)
	goto L11
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v52)
	switch v24 - int32(10) {
	case 0, 1:
		goto L18
	case 2, 3:
		goto L14
	case 4, 5:
		goto L15
	case 6, 7:
		goto L16
	default:
		goto L17
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v50 = v49
	goto L11
L13:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v413 != v19 {
		goto L84
	} else {
		goto L85
	}
L14:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	if v310&int32(2) != 0 {
		v403 = v52
		goto L13
	} else {
		goto L72
	}
L15:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	if v295&int32(2) != 0 {
		v403 = v52
		goto L13
	} else {
		goto L70
	}
L16:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v127 = F_ArrayGetNItemsSafe(m, v124, v19+int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L31
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L18:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	if v57&int32(2) != 0 {
		v403 = v52
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v60 == int32(0) {
		v403 = v52
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v63 = int32(8)
	v70 = v19 + v63
	v73 = v60
	goto L21
L21:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70))))
	v84 = F_ltree_crc32_sz(m, v70+int32(2), v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v403 = v105
	goto L13
L23:
	;
	v86 = base.I32_rem_u_s(v84, v50<<(uint(int32(3))%32))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v63+int32(base.Ui32(v86)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v90)>>(uint(v86&int32(7))%32))&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v403 = int32(0)
	goto L13
L25:
	;
	goto L26
L26:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70))))
	v105 = int32(1)
	if v105 < v73 {
		v70 = v70 + (v99+int32(9))&int32(_a_F__ltree_consistent_0)
		v73 = v73 - v105
		goto L21
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v24
	F_errmsg_internal(m, int32(_a_F__ltree_consistent_1), v15)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F__ltree_consistent_2), int32(541), int32(_a_F__ltree_consistent_3))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v129 < int32(2) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v132 = F_array_contains_nulls(m, v19)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L66
	}
L35:
	;
	if v132 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v127 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L62
	}
L39:
	;
	v403 = int32(0)
	goto L13
L40:
	;
	goto L41
L41:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v139&int32(2) != 0 {
		v403 = v52
		goto L13
	} else {
		goto L42
	}
L42:
	;
	if v123 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v148 = v123
	goto L45
L44:
	;
	v148 = (v124<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L45
L45:
	;
	v161 = v19 + v148
	v163 = v127
	goto L46
L46:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+4)))
	if v166 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v403 = v252
	goto L13
L48:
	;
	v403 = int32(1)
	goto L13
L49:
	;
	goto L50
L50:
	;
	v173 = v166
	v175 = v161 + int32(16)
	goto L51
L51:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+4)))
	if v184 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L47
L53:
	;
	v252 = int32(1)
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175))))
	if v252 < v173 {
		v173 = v173 - v252
		v175 = v175 + (v254+int32(7))&int32(_a_F__ltree_consistent_0)
		goto L51
	} else {
		goto L61
	}
L54:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
	if v187&int32(21) != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v193 = v175 + int32(16)
	v194 = v184
	goto L56
L56:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v205 = base.I32_rem_u_s(v204, v50<<(uint(int32(3))%32))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+int32(8)+int32(base.Ui32(v205)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v209)>>(uint(v205&int32(7))%32))&int32(1) != 0 {
		goto L53
	} else {
		goto L58
	}
L57:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v235 = int32(1)
	if v235 < v163 {
		v161 = v161 + (int32(base.Ui32(v227)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v163 = v163 - v235
		goto L46
	} else {
		goto L60
	}
L58:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+4)))
	v223 = int32(1)
	if v223 < v194 {
		v193 = v193 + (v215+int32(7))&int32(_a_F__ltree_consistent_0) + int32(8)
		v194 = v194 - v223
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v403 = int32(0)
	goto L13
L61:
	;
	goto L52
L62:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F__ltree_consistent_4), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F__ltree_consistent_2), int32(493), int32(_a_F__ltree_consistent_5))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F__ltree_consistent_6), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F__ltree_consistent_2), int32(489), int32(_a_F__ltree_consistent_5))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v50
	v299 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v51 + v299
	v308 = F_ltree_execute(m, v19+v299, v15+v299, int32(0), int32(_a_F__ltree_consistent_7))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v403 = v308
	goto L13
L72:
	;
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v313 == int32(0) {
		v403 = v52
		goto L13
	} else {
		goto L73
	}
L73:
	;
	v325 = v19 + int32(16)
	v326 = v313
	goto L74
L74:
	;
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v325)+4)))
	if v334 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v403 = v390
	goto L13
L76:
	;
	v390 = int32(1)
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v325))))
	if v390 < v326 {
		v325 = v325 + (v392+int32(7))&int32(_a_F__ltree_consistent_0)
		v326 = v326 - v390
		goto L74
	} else {
		goto L83
	}
L77:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+2)))
	if v337&int32(21) != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v343 = v325 + int32(16)
	v344 = v334
	goto L79
L79:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v355 = base.I32_rem_u_s(v354, v50<<(uint(int32(3))%32))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+int32(8)+int32(base.Ui32(v355)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v359)>>(uint(v355&int32(7))%32))&int32(1) != 0 {
		goto L76
	} else {
		goto L81
	}
L80:
	;
	v403 = int32(0)
	goto L13
L81:
	;
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343)+4)))
	v373 = int32(1)
	if v373 < v344 {
		v343 = v343 + (v365+int32(7))&int32(_a_F__ltree_consistent_0) + int32(8)
		v344 = v344 - v373
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	goto L75
L84:
	;
	F_pfree(m, v19)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	m.G0 = v15 + int32(16)
	return v403
L87:
	;
	goto L86
}
func F__ltree_extract_isparent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13844(m, l0, int32(_a_F__ltree_extract_isparent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F__ltree_extract_risparent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13844(m, l0, int32(_a_F__ltree_extract_risparent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F__ltree_gist_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F__ltree_gist_options_0), int32(_a_F__ltree_gist_options_1), int32(28), int32(1), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return int32(0)
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
	var v90 int32
	_ = v90
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
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
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
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v192 = int32(base.Ui32(v190) >> (uint(int32(2)) % 32))
	if v192 == int32(0) {
		goto L12
	} else {
		goto L63
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
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v167&int32(2) != 0 {
		goto L56
	} else {
		goto L57
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
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v161 | int32(4)
	return v24
L36:
	;
	v65 = int32(0)
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if base.B2i32(v72 == v65)|base.B2i32(v75 == v65) != 0 {
		v139 = v72
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v160 != 0 {
		goto L32
	} else {
		goto L55
	}
L38:
	;
	v160 = (v139 + int32(1)) * (v72 - v75) * int32(10)
	goto L37
L39:
	;
	v79 = int32(8)
	v83 = l3 + v79
	v84 = l4 + v79
	v87 = v72
	v90 = v75
	goto L40
L40:
	;
	v92 = int32(2)
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83))))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
	if base.Ui32(v96) < base.Ui32(v97) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v139 = v119
	goto L38
L42:
	;
	v119 = v87 - int32(1)
	if v87 < int32(2) {
		v139 = v119
		goto L38
	} else {
		goto L53
	}
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
	if v100 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v96 == v97 {
		goto L42
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v100 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v104 = int32(10)
	v160 = (v87*v104 + v104) * (v96 - v97)
	goto L37
L50:
	;
	v116 = int32(-10)
	goto L52
L51:
	;
	v116 = int32(10)
	goto L52
L52:
	;
	v160 = (v87 + int32(1)) * v116
	goto L37
L53:
	;
	v122 = int32(9)
	v124 = int32(_a_F_ltree_gist_alloc_0)
	v132 = int32(1)
	if v132 < v90 {
		v83 = v83 + (v96+v122)&v124
		v84 = v84 + (v97+v122)&v124
		v87 = v119
		v90 = v90 - v132
		goto L40
	} else {
		goto L54
	}
L54:
	;
	goto L41
L55:
	;
	goto L35
L56:
	;
	v170 = int32(0)
	goto L58
L57:
	;
	v170 = l2
	goto L58
L58:
	;
	v171 = v48 + v170
	if v167&int32(4) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v180 = v171 + int32(base.Ui32(v176)>>(uint(int32(2))%32))
	goto L61
L60:
	;
	v180 = v171
	goto L61
L61:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v183 = int32(base.Ui32(v181) >> (uint(int32(2)) % 32))
	if v183 == int32(0) {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	base.MemoryCopy(m, v180, l4, v183)
	return v24
L63:
	;
	base.MemoryCopy(m, v24+int32(8), l3, v192)
	goto L12
}
func F_ltree_gt(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
		v152 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v180 != v14 {
		goto L40
	} else {
		goto L41
	}
L5:
	;
	v179 = (v152 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	v28 = int32(8)
	v33 = v22
	v37 = v14 + v28
	v38 = v19 + v28
	v43 = v21
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v37 + v44
	v47 = v38 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v152 = v132
	goto L5
L9:
	;
	v132 = v33 - int32(1)
	if v33 < int32(2) {
		v152 = v132
		goto L5
	} else {
		goto L38
	}
L10:
	;
	v51 = v48
	goto L12
L11:
	;
	v51 = v49
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v113 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	v113 = int32(0)
	goto L13
L15:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L25
L16:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L18
L18:
	;
	if v77 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L20
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L18
L22:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L15
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v113 = v92 - v93
	goto L13
L27:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	if v48 == v49 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v113 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v117 = int32(10)
	v179 = (v33*v117 + v117) * (v48 - v49)
	goto L4
L35:
	;
	v129 = int32(-10)
	goto L37
L36:
	;
	v129 = int32(10)
	goto L37
L37:
	;
	v179 = (v33 + int32(1)) * v129
	goto L4
L38:
	;
	v135 = int32(9)
	v137 = int32(_a_F_ltree_gt_0)
	v145 = int32(1)
	if v145 < v43 {
		v33 = v132
		v37 = v37 + (v48+v135)&v137
		v38 = v38 + (v49+v135)&v137
		v43 = v43 - v145
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L8
L40:
	;
	F_pfree(m, v14)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v184 != v19 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_pfree(m, v19)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	return base.B2i32(int32(0) < v179)
L47:
	;
	goto L46
}
func F_ltree_in(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_parse_ltree(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v12 = v5
		} else {
			v9 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
			v12 = int32(0)
		}
		return v12
	}
}
func F_ltree_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v645 int32
	_ = v645
	var v652 int32
	_ = v652
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v932 int32
	_ = v932
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v1017 int32
	_ = v1017
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	v2 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v26 == v2 {
		v43 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v43&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v30 == int32(0) {
		v43 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v33 != int32(7) {
		v43 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v36 != int32(17) {
		v43 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
	v43 = v39 ^ int32(1)
	goto L2
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = F_get_fn_opclass_options(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v52 = int32(8)
	goto L9
L9:
	;
	v53 = F_palloc0(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v52 = v51
	goto L9
L12:
	;
	v55 = F_palloc0(m, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v58 = int32(1)
	v61 = (v57 - v58) & int32(_a_F_ltree_picksplit_0)
	v65 = v61<<(uint(v58)%32) + int32(4)
	v66 = F_palloc(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v66
	v69 = F_palloc(m, v65)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v71 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v69
	v80 = F_palloc(m, v61<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v57&int32(_a_F_ltree_picksplit_0) == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v956 = int32(1)
	v959 = base.B2i32(v52 <= int32(0))
	if v959|v940 != 0 {
		v995 = v956
		goto L152
	} else {
		goto L153
	}
L18:
	;
	v86 = int32(8)
	F_pg_qsort(m, v80+v86, v61, v86, int32(_a_F_ltree_picksplit_1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v95 = int32(1)
	goto L22
L21:
	;
	v940 = v2
	v942 = v2
	v945 = v2
	v947 = v2
	goto L17
L22:
	;
	v117 = int32(3)
	v119 = v80 + v95<<(uint(v117)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v95
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(4)+v95<<(uint(int32(4))%32))))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v126&v117 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v139 = int32(8)
	F_pg_qsort(m, v80+v139, v61, v139, int32(_a_F_ltree_picksplit_1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L10
	} else {
		goto L28
	}
L24:
	;
	v129 = int32(0)
	goto L26
L25:
	;
	v129 = v52
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v124 + v129 + int32(8)
	v137 = (v95 + int32(1)) & int32(_a_F_ltree_picksplit_0)
	if base.Ui32(v137) <= base.Ui32(v61) {
		v95 = v137
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v146 = v52 & int32(2147483644)
	v147 = int32(3)
	v148 = v52 & v147
	v150 = v52 << (uint(v147) % 32)
	v151 = int32(1)
	v162 = v2
	v164 = v2
	v167 = v2
	v169 = v2
	v173 = v151
	goto L29
L29:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v80+v173<<(uint(int32(3))%32))))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(4)+v181<<(uint(int32(4))%32))))
	if base.Ui32(v173) <= base.Ui32(int32(base.Ui32(v61)>>(uint(v151)%32))) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v940 = v913
	v942 = v915
	v945 = v918
	v947 = v920
	goto L17
L31:
	;
	v932 = (v173 + int32(1)) & int32(_a_F_ltree_picksplit_0)
	if base.Ui32(v932) <= base.Ui32(v61) {
		v162 = v913
		v164 = v915
		v167 = v918
		v169 = v920
		v173 = v932
		goto L29
	} else {
		goto L151
	}
L32:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v189 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v187+v188<<(uint(v189)%32)))) = uint16(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v188 + v189
	if v167 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v549 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v547+v548<<(uint(v549)%32)))) = uint16(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v548 + v549
	if v169 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L35:
	;
	if v337&int32(1) != 0 {
		goto L70
	} else {
		goto L71
	}
L36:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v319&int32(1) != 0 {
		goto L63
	} else {
		goto L64
	}
L37:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v200&int32(1) != 0 {
		v217 = v185 + int32(8)
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v218 = int32(0)
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+4)))
	if base.B2i32(v225 == v218)|base.B2i32(v228 == v218) != 0 {
		v292 = v225
		goto L45
	} else {
		goto L46
	}
L39:
	;
	if v200&int32(2) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v206 = int32(0)
	goto L42
L41:
	;
	v206 = v52
	goto L42
L42:
	;
	v209 = v185 + v206 + int32(8)
	if v200&int32(4) != 0 {
		v217 = v209
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v217 = v209 + int32(base.Ui32(v212)>>(uint(int32(2))%32))
	goto L38
L44:
	;
	if int32(0) < v313 {
		goto L36
	} else {
		goto L62
	}
L45:
	;
	v313 = (v292 + int32(1)) * (v225 - v228) * int32(10)
	goto L44
L46:
	;
	v232 = int32(8)
	v236 = v217 + v232
	v237 = v167 + v232
	v240 = v225
	v243 = v228
	goto L47
L47:
	;
	v245 = int32(2)
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237))))
	if base.Ui32(v249) < base.Ui32(v250) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v292 = v272
	goto L45
L49:
	;
	v272 = v240 - int32(1)
	if v240 < int32(2) {
		v292 = v272
		goto L45
	} else {
		goto L60
	}
L50:
	;
	v252 = v249
	goto L52
L51:
	;
	v252 = v250
	goto L52
L52:
	;
	v253 = F_memcmp(m, v236+v245, v237+v245, v252)
	mBase = m.M
	if v253 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v249 == v250 {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v253 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v257 = int32(10)
	v313 = (v240*v257 + v257) * (v249 - v250)
	goto L44
L57:
	;
	v269 = int32(-10)
	goto L59
L58:
	;
	v269 = int32(10)
	goto L59
L59:
	;
	v313 = (v240 + int32(1)) * v269
	goto L44
L60:
	;
	v275 = int32(9)
	v277 = int32(_a_F_ltree_picksplit_2)
	v285 = int32(1)
	if v285 < v243 {
		v236 = v236 + (v249+v275)&v277
		v237 = v237 + (v250+v275)&v277
		v240 = v272
		v243 = v243 - v285
		goto L47
	} else {
		goto L61
	}
L61:
	;
	goto L48
L62:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v337 = v316
	v339 = v167
	goto L35
L63:
	;
	v337 = v319
	v339 = v185 + int32(8)
	goto L35
L64:
	;
	goto L65
L65:
	;
	if v319&int32(2) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v327 = int32(0)
	goto L68
L67:
	;
	v327 = v52
	goto L68
L68:
	;
	v330 = v185 + v327 + int32(8)
	if v319&int32(4) != 0 {
		v337 = v319
		v339 = v330
		goto L35
	} else {
		goto L69
	}
L69:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v337 = v319
	v339 = v330 + int32(base.Ui32(v333)>>(uint(int32(2))%32))
	goto L35
L70:
	;
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+12)))
	if v342 == int32(0) {
		v913 = v162
		v915 = v164
		v918 = v339
		v920 = v169
		goto L31
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v395 = int32(1)
	if (v162|int32(base.Ui32(v337)>>(uint(v395)%32)))&v395 != 0 {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v347 = v185 + int32(16)
	v348 = v342
	goto L74
L74:
	;
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347))))
	v372 = F_ltree_crc32_sz(m, v347+int32(2), v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L10
	} else {
		goto L76
	}
L75:
	;
	v913 = v162
	v915 = v164
	v918 = v339
	v920 = v169
	goto L31
L76:
	;
	v374 = base.I32_rem_u_s(v372, v150)
	v377 = v53 + int32(base.Ui32(v374)>>(uint(int32(3))%32))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	v379 = int32(1)
	v383 = v378 | v379<<(uint(v374&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v377))) = uint8(v383)
	v385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347))))
	if base.Ui32(v379) < base.Ui32(v348) {
		v347 = v347 + (v385+int32(9))&int32(_a_F_ltree_picksplit_2)
		v348 = v348 - v379
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v913 = int32(1)
	v915 = v164
	v918 = v339
	v920 = v169
	goto L31
L79:
	;
	goto L80
L80:
	;
	if v52 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v913 = int32(0)
	v915 = v164
	v918 = v339
	v920 = v169
	goto L31
L82:
	;
	v404 = v185 + int32(8)
	v405 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v52) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v410 = v405
	v411 = v405
	goto L86
L84:
	;
	v470 = v405
	goto L85
L85:
	;
	v492 = v470
	v496 = v405
	goto L90
L86:
	;
	v432 = v410 + v53
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v404))))
	v436 = v433 | v435
	*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v436)
	v439 = v410 | int32(1)
	v440 = v53 + v439
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v439))))
	v444 = v441 | v443
	*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v444)
	v447 = v410 | int32(2)
	v448 = v53 + v447
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v447))))
	v452 = v449 | v451
	*(*uint8)(unsafe.Add(mBase, uint32(v448))) = uint8(v452)
	v455 = v410 | int32(3)
	v456 = v53 + v455
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v455))))
	v460 = v457 | v459
	*(*uint8)(unsafe.Add(mBase, uint32(v456))) = uint8(v460)
	v462 = int32(4)
	v463 = v410 + v462
	v465 = v411 + v462
	if v465 != v146 {
		v410 = v463
		v411 = v465
		goto L86
	} else {
		goto L88
	}
L87:
	;
	if v148 == int32(0) {
		goto L81
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v470 = v463
	goto L85
L90:
	;
	v513 = v492 + v53
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492+v404))))
	v517 = v514 | v516
	*(*uint8)(unsafe.Add(mBase, uint32(v513))) = uint8(v517)
	v519 = int32(1)
	v522 = v496 + v519
	if v522 != v148 {
		v492 = v492 + v519
		v496 = v522
		goto L90
	} else {
		goto L92
	}
L91:
	;
	goto L81
L92:
	;
	goto L91
L93:
	;
	if v697&int32(1) != 0 {
		goto L128
	} else {
		goto L129
	}
L94:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v679&int32(1) != 0 {
		goto L121
	} else {
		goto L122
	}
L95:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v560&int32(1) != 0 {
		v577 = v185 + int32(8)
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v578 = int32(0)
	v585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+4)))
	v588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	if base.B2i32(v585 == v578)|base.B2i32(v588 == v578) != 0 {
		v652 = v585
		goto L103
	} else {
		goto L104
	}
L97:
	;
	if v560&int32(2) != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v566 = int32(0)
	goto L100
L99:
	;
	v566 = v52
	goto L100
L100:
	;
	v569 = v185 + v566 + int32(8)
	if v560&int32(4) != 0 {
		v577 = v569
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v577 = v569 + int32(base.Ui32(v572)>>(uint(int32(2))%32))
	goto L96
L102:
	;
	if int32(0) < v673 {
		goto L94
	} else {
		goto L120
	}
L103:
	;
	v673 = (v652 + int32(1)) * (v585 - v588) * int32(10)
	goto L102
L104:
	;
	v592 = int32(8)
	v596 = v577 + v592
	v597 = v169 + v592
	v600 = v585
	v603 = v588
	goto L105
L105:
	;
	v605 = int32(2)
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v596))))
	v610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597))))
	if base.Ui32(v609) < base.Ui32(v610) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v652 = v632
	goto L103
L107:
	;
	v632 = v600 - int32(1)
	if v600 < int32(2) {
		v652 = v632
		goto L103
	} else {
		goto L118
	}
L108:
	;
	v612 = v609
	goto L110
L109:
	;
	v612 = v610
	goto L110
L110:
	;
	v613 = F_memcmp(m, v596+v605, v597+v605, v612)
	mBase = m.M
	if v613 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v609 == v610 {
		goto L107
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v613 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v617 = int32(10)
	v673 = (v600*v617 + v617) * (v609 - v610)
	goto L102
L115:
	;
	v629 = int32(-10)
	goto L117
L116:
	;
	v629 = int32(10)
	goto L117
L117:
	;
	v673 = (v600 + int32(1)) * v629
	goto L102
L118:
	;
	v635 = int32(9)
	v637 = int32(_a_F_ltree_picksplit_2)
	v645 = int32(1)
	if v645 < v603 {
		v596 = v596 + (v609+v635)&v637
		v597 = v597 + (v610+v635)&v637
		v600 = v632
		v603 = v603 - v645
		goto L105
	} else {
		goto L119
	}
L119:
	;
	goto L106
L120:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v697 = v676
	v699 = v169
	goto L93
L121:
	;
	v697 = v679
	v699 = v185 + int32(8)
	goto L93
L122:
	;
	goto L123
L123:
	;
	if v679&int32(2) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v687 = int32(0)
	goto L126
L125:
	;
	v687 = v52
	goto L126
L126:
	;
	v690 = v185 + v687 + int32(8)
	if v679&int32(4) != 0 {
		v697 = v679
		v699 = v690
		goto L93
	} else {
		goto L127
	}
L127:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	v697 = v679
	v699 = v690 + int32(base.Ui32(v693)>>(uint(int32(2))%32))
	goto L93
L128:
	;
	v702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+12)))
	if v702 == int32(0) {
		v913 = v162
		v915 = v164
		v918 = v167
		v920 = v699
		goto L31
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v755 = int32(1)
	if (v164|int32(base.Ui32(v697)>>(uint(v755)%32)))&v755 != 0 {
		goto L136
	} else {
		goto L137
	}
L131:
	;
	v707 = v185 + int32(16)
	v708 = v702
	goto L132
L132:
	;
	v731 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v707))))
	v732 = F_ltree_crc32_sz(m, v707+int32(2), v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L10
	} else {
		goto L134
	}
L133:
	;
	v913 = v162
	v915 = v164
	v918 = v167
	v920 = v699
	goto L31
L134:
	;
	v734 = base.I32_rem_u_s(v732, v150)
	v737 = v55 + int32(base.Ui32(v734)>>(uint(int32(3))%32))
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737))))
	v739 = int32(1)
	v743 = v738 | v739<<(uint(v734&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v737))) = uint8(v743)
	v745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v707))))
	if base.Ui32(v739) < base.Ui32(v708) {
		v707 = v707 + (v745+int32(9))&int32(_a_F_ltree_picksplit_2)
		v708 = v708 - v739
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v913 = v162
	v915 = int32(1)
	v918 = v167
	v920 = v699
	goto L31
L137:
	;
	goto L138
L138:
	;
	if v52 <= int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v913 = v162
	v915 = int32(0)
	v918 = v167
	v920 = v699
	goto L31
L140:
	;
	v764 = v185 + int32(8)
	v765 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v52) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v770 = v765
	v771 = v765
	goto L144
L142:
	;
	v830 = v765
	goto L143
L143:
	;
	v852 = v830
	v856 = v765
	goto L148
L144:
	;
	v792 = v770 + v55
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792))))
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770+v764))))
	v796 = v793 | v795
	*(*uint8)(unsafe.Add(mBase, uint32(v792))) = uint8(v796)
	v799 = v770 | int32(1)
	v800 = v55 + v799
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800))))
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764+v799))))
	v804 = v801 | v803
	*(*uint8)(unsafe.Add(mBase, uint32(v800))) = uint8(v804)
	v807 = v770 | int32(2)
	v808 = v55 + v807
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v808))))
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764+v807))))
	v812 = v809 | v811
	*(*uint8)(unsafe.Add(mBase, uint32(v808))) = uint8(v812)
	v815 = v770 | int32(3)
	v816 = v55 + v815
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816))))
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764+v815))))
	v820 = v817 | v819
	*(*uint8)(unsafe.Add(mBase, uint32(v816))) = uint8(v820)
	v822 = int32(4)
	v823 = v770 + v822
	v825 = v771 + v822
	if v825 != v146 {
		v770 = v823
		v771 = v825
		goto L144
	} else {
		goto L146
	}
L145:
	;
	if v148 == int32(0) {
		goto L139
	} else {
		goto L147
	}
L146:
	;
	goto L145
L147:
	;
	v830 = v823
	goto L143
L148:
	;
	v873 = v852 + v55
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873))))
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852+v764))))
	v877 = v874 | v876
	*(*uint8)(unsafe.Add(mBase, uint32(v873))) = uint8(v877)
	v879 = int32(1)
	v882 = v856 + v879
	if v882 != v148 {
		v852 = v852 + v879
		v856 = v882
		goto L148
	} else {
		goto L150
	}
L149:
	;
	goto L139
L150:
	;
	goto L149
L151:
	;
	goto L30
L152:
	;
	if v959|v942 != 0 {
		v1049 = v956
		goto L158
	} else {
		goto L159
	}
L153:
	;
	v962 = int32(0)
	goto L154
L154:
	;
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962+v53))))
	v986 = int32(255)
	v987 = base.B2i32(v985 == v986)
	if v985 != v986 {
		v995 = v987
		goto L152
	} else {
		goto L156
	}
L155:
	;
	v995 = v987
	goto L152
L156:
	;
	v991 = v962 + int32(1)
	if v991 != v52 {
		v962 = v991
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v1070 = int32(4)
	v1071 = v24 + v1070
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1071+v1072<<(uint(v1070)%32))))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1076)+4))
	if v1078&int32(3) != 0 {
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v1017 = int32(0)
	goto L160
L160:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1017+v55))))
	v1041 = int32(255)
	v1042 = base.B2i32(v1040 == v1041)
	if v1040 != v1041 {
		v1049 = v1042
		goto L158
	} else {
		goto L162
	}
L161:
	;
	v1049 = v1042
	goto L158
L162:
	;
	v1046 = v1017 + int32(1)
	if v1046 != v52 {
		v1017 = v1046
		goto L160
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	v1081 = int32(0)
	goto L166
L165:
	;
	v1081 = v52
	goto L166
L166:
	;
	v1085 = F_ltree_gist_alloc(m, v995, v53, v52, v1076+v1081+int32(8), v945)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v80+v61<<(uint(int32(2))%32)&int32(_a_F_ltree_picksplit_3))+8))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1071+v1092<<(uint(int32(4))%32))))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	if v1098&int32(3) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v1101 = int32(0)
	goto L170
L169:
	;
	v1101 = v52
	goto L170
L170:
	;
	v1105 = F_ltree_gist_alloc(m, v1049, v55, v52, v1096+v1101+int32(8), v947)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	F_pfree(m, v53)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	F_pfree(m, v55)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1085
	return v23
}
func F_ltree_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
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
	var v192 int32
	_ = v192
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
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
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
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v515 int32
	_ = v515
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v774 int32
	_ = v774
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v817 int32
	_ = v817
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v870 int32
	_ = v870
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	v2 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
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
	return int32(0)
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
	v808 = v2
	v810 = v2
	v817 = v2
	goto L15
L15:
	;
	if v817&int32(1)|base.B2i32(v51 <= int32(0)) != 0 {
		v870 = int32(1)
		goto L145
	} else {
		goto L146
	}
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(4)+v80<<(uint(int32(4))%32))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v91&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v808 = v783
	v810 = v785
	v817 = v792
	goto L15
L18:
	;
	v799 = v80 + int32(1)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v799 < v800 {
		v72 = v783
		v74 = v785
		v80 = v799
		v81 = v792
		goto L16
	} else {
		goto L144
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
	v368 = v91 & int32(2)
	v369 = int32(1)
	v371 = v81 | int32(base.Ui32(v368)>>(uint(v369)%32))
	if v371&v369 != 0 {
		goto L75
	} else {
		goto L76
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
		goto L52
	} else {
		goto L53
	}
L30:
	;
	v167 = int32(0)
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)))
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	if base.B2i32(v174 == v167)|base.B2i32(v177 == v167) != 0 {
		v241 = v174
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v265 = v166
	goto L29
L33:
	;
	if v262 <= int32(0) {
		v265 = v74
		goto L29
	} else {
		goto L51
	}
L34:
	;
	v262 = (v241 + int32(1)) * (v174 - v177) * int32(10)
	goto L33
L35:
	;
	v185 = v74 + int32(8)
	v186 = v90 + int32(16)
	v189 = v174
	v192 = v177
	goto L36
L36:
	;
	v194 = int32(2)
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185))))
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186))))
	if base.Ui32(v198) < base.Ui32(v199) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v241 = v221
	goto L34
L38:
	;
	v221 = v189 - int32(1)
	if v189 < int32(2) {
		v241 = v221
		goto L34
	} else {
		goto L49
	}
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
	if v202 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v198 == v199 {
		goto L38
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v202 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v206 = int32(10)
	v262 = (v189*v206 + v206) * (v198 - v199)
	goto L33
L46:
	;
	v218 = int32(-10)
	goto L48
L47:
	;
	v218 = int32(10)
	goto L48
L48:
	;
	v262 = (v189 + int32(1)) * v218
	goto L33
L49:
	;
	v224 = int32(9)
	v226 = int32(_a_F_ltree_union_0)
	v234 = int32(1)
	if v234 < v192 {
		v185 = v185 + (v198+v224)&v226
		v186 = v186 + (v199+v224)&v226
		v189 = v221
		v192 = v192 - v234
		goto L36
	} else {
		goto L50
	}
L50:
	;
	goto L37
L51:
	;
	goto L32
L52:
	;
	v783 = v166
	v785 = v265
	v792 = v81
	goto L18
L53:
	;
	goto L54
L54:
	;
	v268 = int32(0)
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	if base.B2i32(v275 == v268)|base.B2i32(v278 == v268) != 0 {
		v342 = v275
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if int32(0) <= v363 {
		v783 = v72
		v785 = v265
		v792 = v81
		goto L18
	} else {
		goto L73
	}
L56:
	;
	v363 = (v342 + int32(1)) * (v275 - v278) * int32(10)
	goto L55
L57:
	;
	v286 = v72 + int32(8)
	v287 = v90 + int32(16)
	v290 = v275
	v293 = v278
	goto L58
L58:
	;
	v295 = int32(2)
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v286))))
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287))))
	if base.Ui32(v299) < base.Ui32(v300) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v342 = v322
	goto L56
L60:
	;
	v322 = v290 - int32(1)
	if v290 < int32(2) {
		v342 = v322
		goto L56
	} else {
		goto L71
	}
L61:
	;
	v302 = v299
	goto L63
L62:
	;
	v302 = v300
	goto L63
L63:
	;
	v303 = F_memcmp(m, v286+v295, v287+v295, v302)
	mBase = m.M
	if v303 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v299 == v300 {
		goto L60
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v303 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v307 = int32(10)
	v363 = (v290*v307 + v307) * (v299 - v300)
	goto L55
L68:
	;
	v319 = int32(-10)
	goto L70
L69:
	;
	v319 = int32(10)
	goto L70
L70:
	;
	v363 = (v290 + int32(1)) * v319
	goto L55
L71:
	;
	v325 = int32(9)
	v327 = int32(_a_F_ltree_union_0)
	v335 = int32(1)
	if v335 < v293 {
		v286 = v286 + (v299+v325)&v327
		v287 = v287 + (v300+v325)&v327
		v290 = v322
		v293 = v293 - v335
		goto L58
	} else {
		goto L72
	}
L72:
	;
	goto L59
L73:
	;
	v783 = v166
	v785 = v265
	v792 = v81
	goto L18
L74:
	;
	v563 = v90 + int32(8)
	v564 = v563 + v541
	if v74 != 0 {
		goto L94
	} else {
		goto L95
	}
L75:
	;
	v539 = v368
	goto L77
L76:
	;
	if v51 <= int32(0) {
		v541 = v51
		goto L74
	} else {
		goto L78
	}
L77:
	;
	if v539 != 0 {
		goto L90
	} else {
		goto L91
	}
L78:
	;
	v377 = v90 + int32(8)
	v378 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v539 = v515 & int32(2)
	goto L77
L80:
	;
	v383 = v378
	v384 = v378
	goto L83
L81:
	;
	v442 = v378
	goto L82
L82:
	;
	v463 = v442
	v472 = v378
	goto L87
L83:
	;
	v404 = v383 + v52
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+v377))))
	v408 = v405 | v407
	*(*uint8)(unsafe.Add(mBase, uint32(v404))) = uint8(v408)
	v411 = v383 | int32(1)
	v412 = v52 + v411
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377+v411))))
	v416 = v413 | v415
	*(*uint8)(unsafe.Add(mBase, uint32(v412))) = uint8(v416)
	v419 = v383 | int32(2)
	v420 = v52 + v419
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377+v419))))
	v424 = v421 | v423
	*(*uint8)(unsafe.Add(mBase, uint32(v420))) = uint8(v424)
	v427 = v383 | int32(3)
	v428 = v52 + v427
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377+v427))))
	v432 = v429 | v431
	*(*uint8)(unsafe.Add(mBase, uint32(v428))) = uint8(v432)
	v434 = int32(4)
	v435 = v383 + v434
	v437 = v384 + v434
	if v437 != v51&int32(2147483644) {
		v383 = v435
		v384 = v437
		goto L83
	} else {
		goto L85
	}
L84:
	;
	if v61 == int32(0) {
		goto L79
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v442 = v435
	goto L82
L87:
	;
	v483 = v463 + v52
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v377))))
	v487 = v484 | v486
	*(*uint8)(unsafe.Add(mBase, uint32(v483))) = uint8(v487)
	v489 = int32(1)
	v492 = v472 + v489
	if v492 != v61 {
		v463 = v463 + v489
		v472 = v492
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L79
L89:
	;
	goto L88
L90:
	;
	v540 = int32(0)
	goto L92
L91:
	;
	v540 = v51
	goto L92
L92:
	;
	v541 = v540
	goto L74
L93:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v665&int32(2) != 0 {
		goto L116
	} else {
		goto L117
	}
L94:
	;
	v565 = int32(0)
	v572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)))
	v575 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v564)+4)))
	if base.B2i32(v572 == v565)|base.B2i32(v575 == v565) != 0 {
		v639 = v572
		goto L98
	} else {
		goto L99
	}
L95:
	;
	goto L96
L96:
	;
	v663 = v564
	goto L93
L97:
	;
	if v660 <= int32(0) {
		v663 = v74
		goto L93
	} else {
		goto L115
	}
L98:
	;
	v660 = (v639 + int32(1)) * (v572 - v575) * int32(10)
	goto L97
L99:
	;
	v579 = int32(8)
	v583 = v74 + v579
	v584 = v564 + v579
	v587 = v572
	v590 = v575
	goto L100
L100:
	;
	v592 = int32(2)
	v596 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	v597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v584))))
	if base.Ui32(v596) < base.Ui32(v597) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v639 = v619
	goto L98
L102:
	;
	v619 = v587 - int32(1)
	if v587 < int32(2) {
		v639 = v619
		goto L98
	} else {
		goto L113
	}
L103:
	;
	v599 = v596
	goto L105
L104:
	;
	v599 = v597
	goto L105
L105:
	;
	v600 = F_memcmp(m, v583+v592, v584+v592, v599)
	mBase = m.M
	if v600 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	if v596 == v597 {
		goto L102
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v600 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v604 = int32(10)
	v660 = (v587*v604 + v604) * (v596 - v597)
	goto L97
L110:
	;
	v616 = int32(-10)
	goto L112
L111:
	;
	v616 = int32(10)
	goto L112
L112:
	;
	v660 = (v587 + int32(1)) * v616
	goto L97
L113:
	;
	v622 = int32(9)
	v624 = int32(_a_F_ltree_union_0)
	v632 = int32(1)
	if v632 < v590 {
		v583 = v583 + (v596+v622)&v624
		v584 = v584 + (v597+v622)&v624
		v587 = v619
		v590 = v590 - v632
		goto L100
	} else {
		goto L114
	}
L114:
	;
	goto L101
L115:
	;
	goto L96
L116:
	;
	v668 = int32(0)
	goto L118
L117:
	;
	v668 = v51
	goto L118
L118:
	;
	v669 = v563 + v668
	if v665&int32(4) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	v678 = v669 + int32(base.Ui32(v674)>>(uint(int32(2))%32))
	goto L121
L120:
	;
	v678 = v669
	goto L121
L121:
	;
	if v72 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v679 = int32(0)
	v686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	v689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v678)+4)))
	if base.B2i32(v686 == v679)|base.B2i32(v689 == v679) != 0 {
		v753 = v686
		goto L126
	} else {
		goto L127
	}
L123:
	;
	goto L124
L124:
	;
	v783 = v678
	v785 = v663
	v792 = v371
	goto L18
L125:
	;
	if int32(0) <= v774 {
		v783 = v72
		v785 = v663
		v792 = v371
		goto L18
	} else {
		goto L143
	}
L126:
	;
	v774 = (v753 + int32(1)) * (v686 - v689) * int32(10)
	goto L125
L127:
	;
	v693 = int32(8)
	v697 = v72 + v693
	v698 = v678 + v693
	v701 = v686
	v704 = v689
	goto L128
L128:
	;
	v706 = int32(2)
	v710 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v697))))
	v711 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v698))))
	if base.Ui32(v710) < base.Ui32(v711) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v753 = v733
	goto L126
L130:
	;
	v733 = v701 - int32(1)
	if v701 < int32(2) {
		v753 = v733
		goto L126
	} else {
		goto L141
	}
L131:
	;
	v713 = v710
	goto L133
L132:
	;
	v713 = v711
	goto L133
L133:
	;
	v714 = F_memcmp(m, v697+v706, v698+v706, v713)
	mBase = m.M
	if v714 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	if v710 == v711 {
		goto L130
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if v714 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v718 = int32(10)
	v774 = (v701*v718 + v718) * (v710 - v711)
	goto L125
L138:
	;
	v730 = int32(-10)
	goto L140
L139:
	;
	v730 = int32(10)
	goto L140
L140:
	;
	v774 = (v701 + int32(1)) * v730
	goto L125
L141:
	;
	v736 = int32(9)
	v738 = int32(_a_F_ltree_union_0)
	v746 = int32(1)
	if v746 < v704 {
		v697 = v697 + (v710+v736)&v738
		v698 = v698 + (v711+v736)&v738
		v701 = v733
		v704 = v704 - v746
		goto L128
	} else {
		goto L142
	}
L142:
	;
	goto L129
L143:
	;
	goto L124
L144:
	;
	goto L17
L145:
	;
	v879 = F_ltree_gist_alloc(m, v870, v52, v51, v810, v808)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L10
	} else {
		goto L151
	}
L146:
	;
	v844 = v2
	goto L147
L147:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v844))))
	v851 = int32(255)
	v852 = base.B2i32(v850 == v851)
	if v850 != v851 {
		v870 = v852
		goto L145
	} else {
		goto L149
	}
L148:
	;
	v870 = v852
	goto L145
L149:
	;
	v856 = v844 + int32(1)
	if v856 != v51 {
		v844 = v856
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(base.Ui32(v881) >> (uint(int32(2)) % 32))
	return v879
}
