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
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
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
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v425 != v19 {
		goto L84
	} else {
		goto L85
	}
L14:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	if v322&int32(2) != 0 {
		v415 = v52
		goto L13
	} else {
		goto L72
	}
L15:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	if v307&int32(2) != 0 {
		v415 = v52
		goto L13
	} else {
		goto L70
	}
L16:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v131 = F_ArrayGetNItems(m, v128, v19+int32(16))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
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
		v415 = v52
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v60 == int32(0) {
		v415 = v52
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
	v415 = v105
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
	v415 = int32(0)
	goto L13
L25:
	;
	goto L26
L26:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70))))
	v105 = int32(1)
	if v105 < v73 {
		v70 = v70 + (v99+int32(9))&int32(131064)
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
	F_errmsg_internal(m, int32(504148), v15)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(514845), int32(541), int32(99233))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
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
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v133 < int32(2) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v136 = F_array_contains_nulls(m, v19)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L66
	}
L35:
	;
	if v136 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v131 <= int32(0) {
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
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L62
	}
L39:
	;
	v415 = int32(0)
	goto L13
L40:
	;
	goto L41
L41:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v143&int32(2) != 0 {
		v415 = v52
		goto L13
	} else {
		goto L42
	}
L42:
	;
	if v127 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v152 = v127
	goto L45
L44:
	;
	v152 = (v128<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L45
L45:
	;
	v167 = v19 + v152
	v169 = v131
	goto L46
L46:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+4)))
	if v170 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v415 = v256
	goto L13
L48:
	;
	v415 = int32(1)
	goto L13
L49:
	;
	goto L50
L50:
	;
	v177 = v170
	v179 = v167 + int32(16)
	goto L51
L51:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179)+4)))
	if v188 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L47
L53:
	;
	v256 = int32(1)
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179))))
	if v256 < v177 {
		v177 = v177 - v256
		v179 = v179 + (v258+int32(7))&int32(131064)
		goto L51
	} else {
		goto L61
	}
L54:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+2)))
	if v191&int32(21) != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v197 = v179 + int32(16)
	v198 = v188
	goto L56
L56:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v209 = base.I32_rem_u_s(v208, v50<<(uint(int32(3))%32))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+int32(8)+int32(base.Ui32(v209)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v213)>>(uint(v209&int32(7))%32))&int32(1) != 0 {
		goto L53
	} else {
		goto L58
	}
L57:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v239 = int32(1)
	if v239 < v169 {
		v167 = v167 + (int32(base.Ui32(v231)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v169 = v169 - v239
		goto L46
	} else {
		goto L60
	}
L58:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+4)))
	v227 = int32(1)
	if v227 < v198 {
		v197 = v197 + (v219+int32(7))&int32(131064) + int32(8)
		v198 = v198 - v227
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v415 = int32(0)
	goto L13
L61:
	;
	goto L52
L62:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(161982), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(514845), int32(493), int32(156058))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
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
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(327544), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(514845), int32(489), int32(156058))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
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
	v311 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v51 + v311
	v320 = F_ltree_execute(m, v19+v311, v15+v311, int32(0), int32(7439))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v415 = v320
	goto L13
L72:
	;
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v325 == int32(0) {
		v415 = v52
		goto L13
	} else {
		goto L73
	}
L73:
	;
	v337 = v19 + int32(16)
	v338 = v325
	goto L74
L74:
	;
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337)+4)))
	if v346 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v415 = v402
	goto L13
L76:
	;
	v402 = int32(1)
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337))))
	if v402 < v338 {
		v337 = v337 + (v404+int32(7))&int32(131064)
		v338 = v338 - v402
		goto L74
	} else {
		goto L83
	}
L77:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+2)))
	if v349&int32(21) != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v355 = v337 + int32(16)
	v356 = v346
	goto L79
L79:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v367 = base.I32_rem_u_s(v366, v50<<(uint(int32(3))%32))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+int32(8)+int32(base.Ui32(v367)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v371)>>(uint(v367&int32(7))%32))&int32(1) != 0 {
		goto L76
	} else {
		goto L81
	}
L80:
	;
	v415 = int32(0)
	goto L13
L81:
	;
	v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v355)+4)))
	v385 = int32(1)
	if v385 < v356 {
		v355 = v355 + (v377+int32(7))&int32(131064) + int32(8)
		v356 = v356 - v385
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
	v428 = m.ExcPending
	if v428 != 0 {
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
	return v415
L87:
	;
	goto L86
}
func F__ltree_extract_isparent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v22 = F_array_iterator(m, v12, int32(5638), v18, v9+int32(12))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v26 != v12 {
						F_pfree(m, v12)
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
									v34 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
									v56 = int32(0)
									m.G0 = v9 + int32(16)
									return v56
								}
							} else {
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v56 = int32(0)
								m.G0 = v9 + int32(16)
								return v56
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
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v56 = int32(0)
								m.G0 = v9 + int32(16)
								return v56
							}
						} else {
							v34 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
							v56 = int32(0)
							m.G0 = v9 + int32(16)
							return v56
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v40 = F_palloc0(m, int32(base.Ui32(v37)>>(uint(int32(2))%32)))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
						if v45 != 0 {
							v46 = F__emscripten_memcpy_bulkmem(m, v40, v42, v45)
							mBase = m.M
						} else {
						}
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v48 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v18 == v52 {
									v56 = v40
									m.G0 = v9 + int32(16)
									return v56
								} else {
									F_pfree(m, v18)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = v40
										m.G0 = v9 + int32(16)
										return v56
									}
								}
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v18 == v52 {
								v56 = v40
								m.G0 = v9 + int32(16)
								return v56
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v40
									m.G0 = v9 + int32(16)
									return v56
								}
							}
						}
					}
				}
			}
		}
	}
}
func F__ltree_extract_risparent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v22 = F_array_iterator(m, v12, int32(5648), v18, v9+int32(12))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v26 != v12 {
						F_pfree(m, v12)
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
									v34 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
									v56 = int32(0)
									m.G0 = v9 + int32(16)
									return v56
								}
							} else {
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v56 = int32(0)
								m.G0 = v9 + int32(16)
								return v56
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
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v56 = int32(0)
								m.G0 = v9 + int32(16)
								return v56
							}
						} else {
							v34 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
							v56 = int32(0)
							m.G0 = v9 + int32(16)
							return v56
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v40 = F_palloc0(m, int32(base.Ui32(v37)>>(uint(int32(2))%32)))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
						if v45 != 0 {
							v46 = F__emscripten_memcpy_bulkmem(m, v40, v42, v45)
							mBase = m.M
						} else {
						}
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v48 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v18 == v52 {
									v56 = v40
									m.G0 = v9 + int32(16)
									return v56
								} else {
									F_pfree(m, v18)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = v40
										m.G0 = v9 + int32(16)
										return v56
									}
								}
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v18 == v52 {
								v56 = v40
								m.G0 = v9 + int32(16)
								return v56
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v40
									m.G0 = v9 + int32(16)
									return v56
								}
							}
						}
					}
				}
			}
		}
	}
}
func F__ltree_gist_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(294915), int32(335963), int32(28), int32(1), int32(2024))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
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
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v186 = int32(base.Ui32(v184) >> (uint(int32(2)) % 32))
	if v186 != 0 {
		goto L72
	} else {
		goto L73
	}
L16:
	;
	if l3 == int32(0) {
		goto L12
	} else {
		goto L28
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
	if l2 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	v41 = F__emscripten_memset_bulkmem(m, v36, base.I32_extend8_s(int32(0)), l2)
	mBase = m.M
	goto L27
L23:
	;
	goto L16
L24:
	;
	v37 = F__emscripten_memcpy_bulkmem(m, v36, l1, l2)
	mBase = m.M
	goto L26
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L16
L28:
	;
	v46 = v24 + int32(8)
	if l0 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v48 = int32(0)
	goto L31
L30:
	;
	v48 = l2
	goto L31
L31:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v52 = int32(base.Ui32(v50) >> (uint(int32(2)) % 32))
	if v52 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if l4 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v53 = F__emscripten_memcpy_bulkmem(m, v46+v48, l3, v52)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v162&int32(2) != 0 {
		goto L61
	} else {
		goto L62
	}
L37:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v156 | int32(4)
	return v24
L38:
	;
	if l3 == l4 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if v58 != v59 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v69 == int32(0) {
		v133 = v69
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v155 != 0 {
		goto L36
	} else {
		goto L60
	}
L42:
	;
	v155 = (v133 + int32(1)) * (v69 - v68) * int32(10)
	goto L41
L43:
	;
	if v68 == int32(0) {
		v133 = v69
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v74 = int32(8)
	v78 = l3 + v74
	v79 = l4 + v74
	v82 = v69
	v86 = v68
	goto L45
L45:
	;
	v87 = int32(2)
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
	if base.Ui32(v91) < base.Ui32(v92) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v133 = v114
	goto L42
L47:
	;
	v114 = v82 - int32(1)
	if v82 < int32(2) {
		v133 = v114
		goto L42
	} else {
		goto L58
	}
L48:
	;
	v94 = v91
	goto L50
L49:
	;
	v94 = v92
	goto L50
L50:
	;
	v95 = F_memcmp(m, v78+v87, v79+v87, v94)
	mBase = m.M
	if v95 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v91 == v92 {
		goto L47
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v95 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v99 = int32(10)
	v155 = (v82*v99 + v99) * (v91 - v92)
	goto L41
L55:
	;
	v111 = int32(-10)
	goto L57
L56:
	;
	v111 = int32(10)
	goto L57
L57:
	;
	v155 = (v82 + int32(1)) * v111
	goto L41
L58:
	;
	v117 = int32(9)
	v119 = int32(131064)
	v127 = int32(1)
	if v127 < v86 {
		v78 = v78 + (v91+v117)&v119
		v79 = v79 + (v92+v117)&v119
		v82 = v114
		v86 = v86 - v127
		goto L45
	} else {
		goto L59
	}
L59:
	;
	goto L46
L60:
	;
	goto L37
L61:
	;
	v165 = int32(0)
	goto L63
L62:
	;
	v165 = l2
	goto L63
L63:
	;
	v166 = v46 + v165
	if v162&int32(4) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v173 = v166
	goto L66
L65:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v173 = v166 + int32(base.Ui32(v169)>>(uint(int32(2))%32))
	goto L66
L66:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v176 = int32(base.Ui32(v174) >> (uint(int32(2)) % 32))
	if v176 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	return v24
L68:
	;
	v177 = F__emscripten_memcpy_bulkmem(m, v173, l4, v176)
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	goto L12
L72:
	;
	v187 = F__emscripten_memcpy_bulkmem(m, v24+int32(8), l3, v186)
	mBase = m.M
	goto L74
L73:
	;
	goto L74
L74:
	;
	goto L71
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
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
	if v22 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v179 != v14 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v178 = (v149 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	if v21 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v27 = int32(8)
	v33 = v14 + v27
	v35 = v22
	v38 = v19 + v27
	v42 = v21
	goto L8
L8:
	;
	v43 = int32(2)
	v44 = v33 + v43
	v46 = v38 + v43
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33))))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v47) < base.Ui32(v48) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v149 = v131
	goto L5
L10:
	;
	v131 = v35 - int32(1)
	if v35 < int32(2) {
		v149 = v131
		goto L5
	} else {
		goto L39
	}
L11:
	;
	v50 = v47
	goto L13
L12:
	;
	v50 = v48
	goto L13
L13:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v112 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v112 = int32(0)
	goto L14
L16:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L26
L17:
	;
	if (v44|v46)&int32(3) != 0 {
		v81 = v44
		v82 = v46
		v83 = v50
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v74 = v44
	v75 = v46
	v76 = v50
	goto L19
L19:
	;
	if v76 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v58 = v44
	v59 = v46
	v60 = v50
	goto L21
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v63 != v64 {
		v81 = v58
		v82 = v59
		v83 = v60
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v74 = v69
	v75 = v67
	v76 = v71
	goto L19
L23:
	;
	v66 = int32(4)
	v67 = v59 + v66
	v69 = v58 + v66
	v71 = v60 - v66
	if base.Ui32(int32(3)) < base.Ui32(v71) {
		v58 = v69
		v59 = v67
		v60 = v71
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L16
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 == v92 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v112 = v91 - v92
	goto L14
L28:
	;
	v94 = int32(1)
	v99 = v88 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v87 + v94
		v88 = v99
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	if v47 == v48 {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v112 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v116 = int32(10)
	v178 = (v35*v116 + v116) * (v47 - v48)
	goto L4
L36:
	;
	v128 = int32(-10)
	goto L38
L37:
	;
	v128 = int32(10)
	goto L38
L38:
	;
	v178 = (v35 + int32(1)) * v128
	goto L4
L39:
	;
	v134 = int32(9)
	v136 = int32(131064)
	v144 = int32(1)
	if v144 < v42 {
		v33 = v33 + (v47+v134)&v136
		v35 = v131
		v38 = v38 + (v48+v134)&v136
		v42 = v42 - v144
		goto L8
	} else {
		goto L40
	}
L40:
	;
	goto L9
L41:
	;
	F_pfree(m, v14)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v183 != v19 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v19)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	return base.B2i32(int32(0) < v178)
L48:
	;
	goto L47
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
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
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v490 int32
	_ = v490
	var v501 int32
	_ = v501
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v753 int32
	_ = v753
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v849 int32
	_ = v849
	var v861 int32
	_ = v861
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v930 int32
	_ = v930
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v1005 int32
	_ = v1005
	var v1018 int32
	_ = v1018
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
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
	v61 = (v57 - v58) & int32(65535)
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
	if v57&int32(65535) == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v954 = int32(1)
	v957 = base.B2i32(v52 <= int32(0))
	if v52 <= int32(0) {
		v1005 = v954
		goto L156
	} else {
		goto L157
	}
L18:
	;
	v86 = int32(8)
	F_pg_qsort(m, v80+v86, v61, v86, int32(7440))
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
	v941 = v2
	v942 = v2
	v943 = v2
	v944 = v2
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
	F_pg_qsort(m, v80+v139, v61, v139, int32(7440))
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
	v137 = (v95 + int32(1)) & int32(65535)
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
	v165 = v2
	v166 = v2
	v167 = v2
	v168 = v2
	v174 = v151
	goto L29
L29:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v80+v174<<(uint(int32(3))%32))))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(4)+v181<<(uint(int32(4))%32))))
	if base.Ui32(v174) <= base.Ui32(int32(base.Ui32(v61)>>(uint(v151)%32))) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v941 = v914
	v942 = v915
	v943 = v916
	v944 = v917
	goto L17
L31:
	;
	v930 = (v174 + int32(1)) & int32(65535)
	if base.Ui32(v930) <= base.Ui32(v61) {
		v165 = v914
		v166 = v915
		v167 = v916
		v168 = v917
		v174 = v930
		goto L29
	} else {
		goto L155
	}
L32:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v189 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v187+v188<<(uint(v189)%32)))) = uint16(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v188 + v189
	if v166 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v548 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v546+v547<<(uint(v548)%32)))) = uint16(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v547 + v548
	if v165 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L35:
	;
	if v336&int32(1) != 0 {
		goto L71
	} else {
		goto L72
	}
L36:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v318&int32(1) != 0 {
		goto L64
	} else {
		goto L65
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
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	if v226 == int32(0) {
		v290 = v226
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
	if int32(0) < v312 {
		goto L36
	} else {
		goto L63
	}
L45:
	;
	v312 = (v290 + int32(1)) * (v226 - v225) * int32(10)
	goto L44
L46:
	;
	if v225 == int32(0) {
		v290 = v226
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v231 = int32(8)
	v235 = v217 + v231
	v236 = v166 + v231
	v239 = v226
	v243 = v225
	goto L48
L48:
	;
	v244 = int32(2)
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v235))))
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
	if base.Ui32(v248) < base.Ui32(v249) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v290 = v271
	goto L45
L50:
	;
	v271 = v239 - int32(1)
	if v239 < int32(2) {
		v290 = v271
		goto L45
	} else {
		goto L61
	}
L51:
	;
	v251 = v248
	goto L53
L52:
	;
	v251 = v249
	goto L53
L53:
	;
	v252 = F_memcmp(m, v235+v244, v236+v244, v251)
	mBase = m.M
	if v252 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v248 == v249 {
		goto L50
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v252 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v256 = int32(10)
	v312 = (v239*v256 + v256) * (v248 - v249)
	goto L44
L58:
	;
	v268 = int32(-10)
	goto L60
L59:
	;
	v268 = int32(10)
	goto L60
L60:
	;
	v312 = (v239 + int32(1)) * v268
	goto L44
L61:
	;
	v274 = int32(9)
	v276 = int32(131064)
	v284 = int32(1)
	if v284 < v243 {
		v235 = v235 + (v248+v274)&v276
		v236 = v236 + (v249+v274)&v276
		v239 = v271
		v243 = v243 - v284
		goto L48
	} else {
		goto L62
	}
L62:
	;
	goto L49
L63:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v336 = v315
	v338 = v166
	goto L35
L64:
	;
	v336 = v318
	v338 = v185 + int32(8)
	goto L35
L65:
	;
	goto L66
L66:
	;
	if v318&int32(2) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v326 = int32(0)
	goto L69
L68:
	;
	v326 = v52
	goto L69
L69:
	;
	v329 = v185 + v326 + int32(8)
	if v318&int32(4) != 0 {
		v336 = v318
		v338 = v329
		goto L35
	} else {
		goto L70
	}
L70:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v336 = v318
	v338 = v329 + int32(base.Ui32(v332)>>(uint(int32(2))%32))
	goto L35
L71:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+12)))
	if v341 == int32(0) {
		v914 = v165
		v915 = v338
		v916 = v167
		v917 = v168
		goto L31
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v394 = int32(1)
	if (v167|int32(base.Ui32(v336)>>(uint(v394)%32)))&v394 != 0 {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	v346 = v185 + int32(16)
	v347 = v341
	goto L75
L75:
	;
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v346))))
	v371 = F_ltree_crc32_sz(m, v346+int32(2), v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L10
	} else {
		goto L77
	}
L76:
	;
	v914 = v165
	v915 = v338
	v916 = v167
	v917 = v168
	goto L31
L77:
	;
	v373 = base.I32_rem_u_s(v371, v150)
	v376 = v53 + int32(base.Ui32(v373)>>(uint(int32(3))%32))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	v378 = int32(1)
	v382 = v377 | v378<<(uint(v373&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v376))) = uint8(v382)
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v346))))
	if base.Ui32(v378) < base.Ui32(v347) {
		v346 = v346 + (v384+int32(9))&int32(131064)
		v347 = v347 - v378
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v914 = v165
	v915 = v338
	v916 = int32(1)
	v917 = v168
	goto L31
L80:
	;
	goto L81
L81:
	;
	if int32(0) < v52 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v403 = v185 + int32(8)
	v404 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v52) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v914 = v165
	v915 = v338
	v916 = int32(0)
	v917 = v168
	goto L31
L85:
	;
	v409 = v404
	v411 = v404
	goto L88
L86:
	;
	v466 = v404
	goto L87
L87:
	;
	if v148 == int32(0) {
		v914 = v165
		v915 = v338
		v916 = v404
		v917 = v168
		goto L31
	} else {
		goto L91
	}
L88:
	;
	v431 = v409 + v53
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431))))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409+v403))))
	v435 = v432 | v434
	*(*uint8)(unsafe.Add(mBase, uint32(v431))) = uint8(v435)
	v438 = v409 | int32(1)
	v439 = v53 + v438
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439))))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v438))))
	v443 = v440 | v442
	*(*uint8)(unsafe.Add(mBase, uint32(v439))) = uint8(v443)
	v446 = v409 | int32(2)
	v447 = v53 + v446
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447))))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v446))))
	v451 = v448 | v450
	*(*uint8)(unsafe.Add(mBase, uint32(v447))) = uint8(v451)
	v454 = v409 | int32(3)
	v455 = v53 + v454
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455))))
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v454))))
	v459 = v456 | v458
	*(*uint8)(unsafe.Add(mBase, uint32(v455))) = uint8(v459)
	v461 = int32(4)
	v462 = v409 + v461
	v464 = v411 + v461
	if v464 != v146 {
		v409 = v462
		v411 = v464
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v466 = v462
	goto L87
L90:
	;
	goto L89
L91:
	;
	v490 = v466
	v501 = v404
	goto L92
L92:
	;
	v512 = v490 + v53
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490+v403))))
	v516 = v513 | v515
	*(*uint8)(unsafe.Add(mBase, uint32(v512))) = uint8(v516)
	v518 = int32(1)
	v521 = v501 + v518
	if v521 != v148 {
		v490 = v490 + v518
		v501 = v521
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L84
L94:
	;
	goto L93
L95:
	;
	if v695&int32(1) != 0 {
		goto L131
	} else {
		goto L132
	}
L96:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v677&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L97:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v559&int32(1) != 0 {
		v576 = v185 + int32(8)
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+4)))
	v585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v576)+4)))
	if v585 == int32(0) {
		v649 = v585
		goto L105
	} else {
		goto L106
	}
L99:
	;
	if v559&int32(2) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v565 = int32(0)
	goto L102
L101:
	;
	v565 = v52
	goto L102
L102:
	;
	v568 = v185 + v565 + int32(8)
	if v559&int32(4) != 0 {
		v576 = v568
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v576 = v568 + int32(base.Ui32(v571)>>(uint(int32(2))%32))
	goto L98
L104:
	;
	if int32(0) < v671 {
		goto L96
	} else {
		goto L123
	}
L105:
	;
	v671 = (v649 + int32(1)) * (v585 - v584) * int32(10)
	goto L104
L106:
	;
	if v584 == int32(0) {
		v649 = v585
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v590 = int32(8)
	v594 = v576 + v590
	v595 = v165 + v590
	v598 = v585
	v602 = v584
	goto L108
L108:
	;
	v603 = int32(2)
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v594))))
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v595))))
	if base.Ui32(v607) < base.Ui32(v608) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v649 = v630
	goto L105
L110:
	;
	v630 = v598 - int32(1)
	if v598 < int32(2) {
		v649 = v630
		goto L105
	} else {
		goto L121
	}
L111:
	;
	v610 = v607
	goto L113
L112:
	;
	v610 = v608
	goto L113
L113:
	;
	v611 = F_memcmp(m, v594+v603, v595+v603, v610)
	mBase = m.M
	if v611 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	if v607 == v608 {
		goto L110
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if v611 < int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v615 = int32(10)
	v671 = (v598*v615 + v615) * (v607 - v608)
	goto L104
L118:
	;
	v627 = int32(-10)
	goto L120
L119:
	;
	v627 = int32(10)
	goto L120
L120:
	;
	v671 = (v598 + int32(1)) * v627
	goto L104
L121:
	;
	v633 = int32(9)
	v635 = int32(131064)
	v643 = int32(1)
	if v643 < v602 {
		v594 = v594 + (v607+v633)&v635
		v595 = v595 + (v608+v633)&v635
		v598 = v630
		v602 = v602 - v643
		goto L108
	} else {
		goto L122
	}
L122:
	;
	goto L109
L123:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v695 = v674
	v697 = v165
	goto L95
L124:
	;
	v695 = v677
	v697 = v185 + int32(8)
	goto L95
L125:
	;
	goto L126
L126:
	;
	if v677&int32(2) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v685 = int32(0)
	goto L129
L128:
	;
	v685 = v52
	goto L129
L129:
	;
	v688 = v185 + v685 + int32(8)
	if v677&int32(4) != 0 {
		v695 = v677
		v697 = v688
		goto L95
	} else {
		goto L130
	}
L130:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	v695 = v677
	v697 = v688 + int32(base.Ui32(v691)>>(uint(int32(2))%32))
	goto L95
L131:
	;
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+12)))
	if v700 == int32(0) {
		v914 = v697
		v915 = v166
		v916 = v167
		v917 = v168
		goto L31
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v753 = int32(1)
	if (v168|int32(base.Ui32(v695)>>(uint(v753)%32)))&v753 != 0 {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v705 = v185 + int32(16)
	v706 = v700
	goto L135
L135:
	;
	v729 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v705))))
	v730 = F_ltree_crc32_sz(m, v705+int32(2), v729)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L10
	} else {
		goto L137
	}
L136:
	;
	v914 = v697
	v915 = v166
	v916 = v167
	v917 = v168
	goto L31
L137:
	;
	v732 = base.I32_rem_u_s(v730, v150)
	v735 = v55 + int32(base.Ui32(v732)>>(uint(int32(3))%32))
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	v737 = int32(1)
	v741 = v736 | v737<<(uint(v732&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v735))) = uint8(v741)
	v743 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v705))))
	if base.Ui32(v737) < base.Ui32(v706) {
		v705 = v705 + (v743+int32(9))&int32(131064)
		v706 = v706 - v737
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v914 = v697
	v915 = v166
	v916 = v167
	v917 = int32(1)
	goto L31
L140:
	;
	goto L141
L141:
	;
	if int32(0) < v52 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v762 = v185 + int32(8)
	v763 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v52) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	v914 = v697
	v915 = v166
	v916 = v167
	v917 = int32(0)
	goto L31
L145:
	;
	v768 = v763
	v770 = v763
	goto L148
L146:
	;
	v825 = v763
	goto L147
L147:
	;
	if v148 == int32(0) {
		v914 = v697
		v915 = v166
		v916 = v167
		v917 = v763
		goto L31
	} else {
		goto L151
	}
L148:
	;
	v790 = v768 + v55
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790))))
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768+v762))))
	v794 = v791 | v793
	*(*uint8)(unsafe.Add(mBase, uint32(v790))) = uint8(v794)
	v797 = v768 | int32(1)
	v798 = v55 + v797
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762+v797))))
	v802 = v799 | v801
	*(*uint8)(unsafe.Add(mBase, uint32(v798))) = uint8(v802)
	v805 = v768 | int32(2)
	v806 = v55 + v805
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806))))
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762+v805))))
	v810 = v807 | v809
	*(*uint8)(unsafe.Add(mBase, uint32(v806))) = uint8(v810)
	v813 = v768 | int32(3)
	v814 = v55 + v813
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v814))))
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762+v813))))
	v818 = v815 | v817
	*(*uint8)(unsafe.Add(mBase, uint32(v814))) = uint8(v818)
	v820 = int32(4)
	v821 = v768 + v820
	v823 = v770 + v820
	if v823 != v146 {
		v768 = v821
		v770 = v823
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v825 = v821
	goto L147
L150:
	;
	goto L149
L151:
	;
	v849 = v825
	v861 = v763
	goto L152
L152:
	;
	v871 = v849 + v55
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871))))
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849+v762))))
	v875 = v872 | v874
	*(*uint8)(unsafe.Add(mBase, uint32(v871))) = uint8(v875)
	v877 = int32(1)
	v880 = v861 + v877
	if v880 != v148 {
		v849 = v849 + v877
		v861 = v880
		goto L152
	} else {
		goto L154
	}
L153:
	;
	goto L144
L154:
	;
	goto L153
L155:
	;
	goto L30
L156:
	;
	if (v957|v944)&int32(1) != 0 {
		v1051 = v954
		goto L163
	} else {
		goto L164
	}
L157:
	;
	if v943&int32(1) != 0 {
		v1005 = v954
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v961 = int32(0)
	goto L159
L159:
	;
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v961+v53))))
	v985 = int32(255)
	v986 = base.B2i32(v984 == v985)
	if v984 != v985 {
		v1005 = v986
		goto L156
	} else {
		goto L161
	}
L160:
	;
	v1005 = v986
	goto L156
L161:
	;
	v990 = v961 + int32(1)
	if v990 != v52 {
		v961 = v990
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v1071 = int32(4)
	v1072 = v24 + v1071
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1072+v1073<<(uint(v1071)%32))))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+4))
	if v1079&int32(3) != 0 {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	v1018 = int32(0)
	goto L165
L165:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1018+v55))))
	v1042 = int32(255)
	v1043 = base.B2i32(v1041 == v1042)
	if v1041 != v1042 {
		v1051 = v1043
		goto L163
	} else {
		goto L167
	}
L166:
	;
	v1051 = v1043
	goto L163
L167:
	;
	v1047 = v1018 + int32(1)
	if v1047 != v52 {
		v1018 = v1047
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	v1082 = int32(0)
	goto L171
L170:
	;
	v1082 = v52
	goto L171
L171:
	;
	v1086 = F_ltree_gist_alloc(m, v1005, v53, v52, v1077+v1082+int32(8), v942)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v80+v61<<(uint(int32(2))%32)&int32(262136))+8))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1072+v1093<<(uint(int32(4))%32))))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	if v1099&int32(3) != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1102 = int32(0)
	goto L175
L174:
	;
	v1102 = v52
	goto L175
L175:
	;
	v1106 = F_ltree_gist_alloc(m, v1051, v55, v52, v1097+v1102+int32(8), v941)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L10
	} else {
		goto L176
	}
L176:
	;
	F_pfree(m, v53)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	F_pfree(m, v55)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L10
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1086
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
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
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v511 int32
	_ = v511
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v834 int32
	_ = v834
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
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
	v54 = int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v55 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v51 <= int32(0) {
		v866 = v54
		goto L151
	} else {
		goto L152
	}
L14:
	;
	v799 = v2
	v801 = v2
	v807 = v2
	goto L13
L15:
	;
	goto L16
L16:
	;
	v60 = int32(3)
	v61 = v51 & v60
	v69 = v2
	v71 = v2
	v77 = v2
	v80 = v2
	goto L17
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(4)+v80<<(uint(int32(4))%32))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v91&int32(1) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v799 = v774
	v801 = v776
	v807 = v782
	goto L13
L19:
	;
	v793 = v80 + int32(1)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v793 < v794 {
		v69 = v774
		v71 = v776
		v77 = v782
		v80 = v793
		goto L17
	} else {
		goto L150
	}
L20:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+12)))
	if v94 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v366 = v91 & int32(2)
	v367 = int32(1)
	v369 = v77 | int32(base.Ui32(v366)>>(uint(v367)%32))
	if v369&v367 != 0 {
		goto L78
	} else {
		goto L79
	}
L23:
	;
	v97 = v90 + int32(16)
	v99 = v94
	goto L26
L24:
	;
	goto L25
L25:
	;
	v166 = v90 + int32(8)
	if v71 != 0 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	v121 = F_ltree_crc32_sz(m, v97+int32(2), v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	v123 = base.I32_rem_u_s(v121, v51<<(uint(v60)%32))
	v126 = v52 + int32(base.Ui32(v123)>>(uint(int32(3))%32))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v128 = int32(1)
	v132 = v127 | v128<<(uint(v123&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v132)
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	if base.Ui32(v128) < base.Ui32(v99) {
		v97 = v97 + (v134+int32(9))&int32(131064)
		v99 = v99 - v128
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v69 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L31:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v175 == int32(0) {
		v239 = v175
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v264 = v166
	goto L30
L34:
	;
	if v261 <= int32(0) {
		v264 = v71
		goto L30
	} else {
		goto L53
	}
L35:
	;
	v261 = (v239 + int32(1)) * (v175 - v174) * int32(10)
	goto L34
L36:
	;
	if v174 == int32(0) {
		v239 = v175
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v184 = v71 + int32(8)
	v185 = v90 + int32(16)
	v188 = v175
	v192 = v174
	goto L38
L38:
	;
	v193 = int32(2)
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184))))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32(v197) < base.Ui32(v198) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v239 = v220
	goto L35
L40:
	;
	v220 = v188 - int32(1)
	if v188 < int32(2) {
		v239 = v220
		goto L35
	} else {
		goto L51
	}
L41:
	;
	v200 = v197
	goto L43
L42:
	;
	v200 = v198
	goto L43
L43:
	;
	v201 = F_memcmp(m, v184+v193, v185+v193, v200)
	mBase = m.M
	if v201 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v197 == v198 {
		goto L40
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v201 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v205 = int32(10)
	v261 = (v188*v205 + v205) * (v197 - v198)
	goto L34
L48:
	;
	v217 = int32(-10)
	goto L50
L49:
	;
	v217 = int32(10)
	goto L50
L50:
	;
	v261 = (v188 + int32(1)) * v217
	goto L34
L51:
	;
	v223 = int32(9)
	v225 = int32(131064)
	v233 = int32(1)
	if v233 < v192 {
		v184 = v184 + (v197+v223)&v225
		v185 = v185 + (v198+v223)&v225
		v188 = v220
		v192 = v192 - v233
		goto L38
	} else {
		goto L52
	}
L52:
	;
	goto L39
L53:
	;
	goto L33
L54:
	;
	v774 = v166
	v776 = v264
	v782 = v77
	goto L19
L55:
	;
	goto L56
L56:
	;
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	if v275 == int32(0) {
		v339 = v275
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if int32(0) <= v361 {
		v774 = v69
		v776 = v264
		v782 = v77
		goto L19
	} else {
		goto L76
	}
L58:
	;
	v361 = (v339 + int32(1)) * (v275 - v274) * int32(10)
	goto L57
L59:
	;
	if v274 == int32(0) {
		v339 = v275
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v284 = v69 + int32(8)
	v285 = v90 + int32(16)
	v288 = v275
	v292 = v274
	goto L61
L61:
	;
	v293 = int32(2)
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284))))
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285))))
	if base.Ui32(v297) < base.Ui32(v298) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v339 = v320
	goto L58
L63:
	;
	v320 = v288 - int32(1)
	if v288 < int32(2) {
		v339 = v320
		goto L58
	} else {
		goto L74
	}
L64:
	;
	v300 = v297
	goto L66
L65:
	;
	v300 = v298
	goto L66
L66:
	;
	v301 = F_memcmp(m, v284+v293, v285+v293, v300)
	mBase = m.M
	if v301 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v297 == v298 {
		goto L63
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v301 < int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v305 = int32(10)
	v361 = (v288*v305 + v305) * (v297 - v298)
	goto L57
L71:
	;
	v317 = int32(-10)
	goto L73
L72:
	;
	v317 = int32(10)
	goto L73
L73:
	;
	v361 = (v288 + int32(1)) * v317
	goto L57
L74:
	;
	v323 = int32(9)
	v325 = int32(131064)
	v333 = int32(1)
	if v333 < v292 {
		v284 = v284 + (v297+v323)&v325
		v285 = v285 + (v298+v323)&v325
		v288 = v320
		v292 = v292 - v333
		goto L61
	} else {
		goto L75
	}
L75:
	;
	goto L62
L76:
	;
	v774 = v166
	v776 = v264
	v782 = v77
	goto L19
L77:
	;
	v559 = v90 + int32(8)
	v560 = v559 + v537
	if v71 != 0 {
		goto L98
	} else {
		goto L99
	}
L78:
	;
	v535 = v366
	goto L80
L79:
	;
	if v51 <= int32(0) {
		v537 = v51
		goto L77
	} else {
		goto L81
	}
L80:
	;
	if v535 != 0 {
		goto L94
	} else {
		goto L95
	}
L81:
	;
	v375 = v90 + int32(8)
	v376 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v381 = v376
	v390 = v376
	goto L85
L83:
	;
	v437 = v376
	goto L84
L84:
	;
	if v61 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v402 = v381 + v52
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381+v375))))
	v406 = v403 | v405
	*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v406)
	v409 = v381 | int32(1)
	v410 = v52 + v409
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375+v409))))
	v414 = v411 | v413
	*(*uint8)(unsafe.Add(mBase, uint32(v410))) = uint8(v414)
	v417 = v381 | int32(2)
	v418 = v52 + v417
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375+v417))))
	v422 = v419 | v421
	*(*uint8)(unsafe.Add(mBase, uint32(v418))) = uint8(v422)
	v425 = v381 | int32(3)
	v426 = v52 + v425
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375+v425))))
	v430 = v427 | v429
	*(*uint8)(unsafe.Add(mBase, uint32(v426))) = uint8(v430)
	v432 = int32(4)
	v433 = v381 + v432
	v435 = v390 + v432
	if v435 != v51&int32(2147483644) {
		v381 = v433
		v390 = v435
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v437 = v433
	goto L84
L87:
	;
	goto L86
L88:
	;
	v458 = v437
	v464 = v376
	goto L91
L89:
	;
	goto L90
L90:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v535 = v511 & int32(2)
	goto L80
L91:
	;
	v479 = v458 + v52
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+v375))))
	v483 = v480 | v482
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v483)
	v485 = int32(1)
	v488 = v464 + v485
	if v488 != v61 {
		v458 = v458 + v485
		v464 = v488
		goto L91
	} else {
		goto L93
	}
L92:
	;
	goto L90
L93:
	;
	goto L92
L94:
	;
	v536 = int32(0)
	goto L96
L95:
	;
	v536 = v51
	goto L96
L96:
	;
	v537 = v536
	goto L77
L97:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v660&int32(2) != 0 {
		goto L121
	} else {
		goto L122
	}
L98:
	;
	v568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v560)+4)))
	v569 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v569 == int32(0) {
		v633 = v569
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v658 = v560
	goto L97
L101:
	;
	if v655 <= int32(0) {
		v658 = v71
		goto L97
	} else {
		goto L120
	}
L102:
	;
	v655 = (v633 + int32(1)) * (v569 - v568) * int32(10)
	goto L101
L103:
	;
	if v568 == int32(0) {
		v633 = v569
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v574 = int32(8)
	v578 = v71 + v574
	v579 = v560 + v574
	v582 = v569
	v586 = v568
	goto L105
L105:
	;
	v587 = int32(2)
	v591 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v578))))
	v592 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579))))
	if base.Ui32(v591) < base.Ui32(v592) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v633 = v614
	goto L102
L107:
	;
	v614 = v582 - int32(1)
	if v582 < int32(2) {
		v633 = v614
		goto L102
	} else {
		goto L118
	}
L108:
	;
	v594 = v591
	goto L110
L109:
	;
	v594 = v592
	goto L110
L110:
	;
	v595 = F_memcmp(m, v578+v587, v579+v587, v594)
	mBase = m.M
	if v595 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v591 == v592 {
		goto L107
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v595 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v599 = int32(10)
	v655 = (v582*v599 + v599) * (v591 - v592)
	goto L101
L115:
	;
	v611 = int32(-10)
	goto L117
L116:
	;
	v611 = int32(10)
	goto L117
L117:
	;
	v655 = (v582 + int32(1)) * v611
	goto L101
L118:
	;
	v617 = int32(9)
	v619 = int32(131064)
	v627 = int32(1)
	if v627 < v586 {
		v578 = v578 + (v591+v617)&v619
		v579 = v579 + (v592+v617)&v619
		v582 = v614
		v586 = v586 - v627
		goto L105
	} else {
		goto L119
	}
L119:
	;
	goto L106
L120:
	;
	goto L100
L121:
	;
	v663 = int32(0)
	goto L123
L122:
	;
	v663 = v51
	goto L123
L123:
	;
	v664 = v559 + v663
	if v660&int32(4) == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v664)))
	v673 = v664 + int32(base.Ui32(v669)>>(uint(int32(2))%32))
	goto L126
L125:
	;
	v673 = v664
	goto L126
L126:
	;
	if v69 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v673)+4)))
	v682 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	if v682 == int32(0) {
		v746 = v682
		goto L131
	} else {
		goto L132
	}
L128:
	;
	goto L129
L129:
	;
	v774 = v673
	v776 = v658
	v782 = v369
	goto L19
L130:
	;
	if int32(0) <= v768 {
		v774 = v69
		v776 = v658
		v782 = v369
		goto L19
	} else {
		goto L149
	}
L131:
	;
	v768 = (v746 + int32(1)) * (v682 - v681) * int32(10)
	goto L130
L132:
	;
	if v681 == int32(0) {
		v746 = v682
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v687 = int32(8)
	v691 = v69 + v687
	v692 = v673 + v687
	v695 = v682
	v699 = v681
	goto L134
L134:
	;
	v700 = int32(2)
	v704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v691))))
	v705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v692))))
	if base.Ui32(v704) < base.Ui32(v705) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v746 = v727
	goto L131
L136:
	;
	v727 = v695 - int32(1)
	if v695 < int32(2) {
		v746 = v727
		goto L131
	} else {
		goto L147
	}
L137:
	;
	v707 = v704
	goto L139
L138:
	;
	v707 = v705
	goto L139
L139:
	;
	v708 = F_memcmp(m, v691+v700, v692+v700, v707)
	mBase = m.M
	if v708 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v704 == v705 {
		goto L136
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if v708 < int32(0) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v712 = int32(10)
	v768 = (v695*v712 + v712) * (v704 - v705)
	goto L130
L144:
	;
	v724 = int32(-10)
	goto L146
L145:
	;
	v724 = int32(10)
	goto L146
L146:
	;
	v768 = (v695 + int32(1)) * v724
	goto L130
L147:
	;
	v730 = int32(9)
	v732 = int32(131064)
	v740 = int32(1)
	if v740 < v699 {
		v691 = v691 + (v704+v730)&v732
		v692 = v692 + (v705+v730)&v732
		v695 = v727
		v699 = v699 - v740
		goto L134
	} else {
		goto L148
	}
L148:
	;
	goto L135
L149:
	;
	goto L129
L150:
	;
	goto L18
L151:
	;
	v872 = F_ltree_gist_alloc(m, v866, v52, v51, v801, v799)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L10
	} else {
		goto L158
	}
L152:
	;
	if v807&int32(1) != 0 {
		v866 = v54
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v834 = v2
	goto L154
L154:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v834))))
	v844 = int32(255)
	v845 = base.B2i32(v843 == v844)
	if v843 != v844 {
		v866 = v845
		goto L151
	} else {
		goto L156
	}
L155:
	;
	v866 = v845
	goto L151
L156:
	;
	v849 = v834 + int32(1)
	if v849 != v51 {
		v834 = v849
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(base.Ui32(v874) >> (uint(int32(2)) % 32))
	return v872
}
