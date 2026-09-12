package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_expand_grouping_sets(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 float64
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
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
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	v4 = int32(0)
	if l0 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v17 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_list_sort(m, v240, int32(479))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L11
	} else {
		goto L69
	}
L5:
	;
	return v321
L6:
	;
	v25 = v4
	v29 = v4
	v32 = float64(1)
	goto L9
L7:
	;
	v65 = v4
	goto L8
L8:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v25<<(uint(int32(2))%32))))
	v39 = F_expand_groupingset_node(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v65 = v52
	goto L8
L11:
	;
	return int32(0)
L12:
	;
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v46 = base.F64_convert_i32_s(v43)
	goto L15
L14:
	;
	v46 = float64(0)
	goto L15
L15:
	;
	v47 = base.F64_mul(v32, v46)
	if base.B2i32(int32(0) <= l2)&base.F64_gt(v47, base.F64_convert_i32_u(l2)) != 0 {
		v321 = v4
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v52 = F_lappend(m, v29, v39)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v55 = v25 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v55 < v56 {
		v25 = v55
		v29 = v52
		v32 = v47
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L51
	}
L20:
	;
	v72 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v73 <= v72 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v120 = v4
	goto L22
L22:
	;
	v128 = int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v129 <= v128 {
		v240 = v120
		goto L19
	} else {
		goto L33
	}
L23:
	;
	if v65 == int32(0) {
		v240 = v106
		goto L19
	} else {
		goto L32
	}
L24:
	;
	v106 = v4
	goto L23
L25:
	;
	goto L26
L26:
	;
	v79 = v72
	v80 = v4
	goto L27
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v79<<(uint(int32(2))%32))))
	v94 = F_list_union_int(m, int32(0), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	v106 = v96
	goto L23
L29:
	;
	v96 = F_lappend(m, v80, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v99 = v79 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v99 < v100 {
		v79 = v99
		v80 = v96
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v120 = v106
	goto L22
L33:
	;
	v136 = v120
	v141 = v128
	goto L34
L34:
	;
	v144 = int32(0)
	if v136 == v144 {
		v224 = v144
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v240 = v224
	goto L19
L36:
	;
	v233 = v141 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v233 < v234 {
		v136 = v224
		v141 = v233
		goto L34
	} else {
		goto L50
	}
L37:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v147 <= int32(0) {
		v224 = v144
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+v141<<(uint(int32(2))%32))))
	v158 = int32(0)
	v160 = v144
	goto L39
L39:
	;
	if v154 == int32(0) {
		v208 = v160
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v224 = v208
	goto L36
L41:
	;
	v217 = v158 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v217 < v218 {
		v158 = v217
		v160 = v208
		goto L39
	} else {
		goto L49
	}
L42:
	;
	v170 = int32(0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v171 <= v170 {
		v208 = v160
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174+v158<<(uint(int32(2))%32))))
	v182 = v170
	v183 = v160
	goto L44
L44:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v182<<(uint(int32(2))%32))))
	v196 = F_list_union_int(m, v178, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	v208 = v198
	goto L41
L46:
	;
	v198 = F_lappend(m, v183, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v201 = v182 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v201 < v202 {
		v182 = v201
		v183 = v198
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L40
L50:
	;
	goto L35
L51:
	;
	if v240 == int32(0) {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v252 < int32(2) {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v259 = int32(0)
	goto L54
L54:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268+v259<<(uint(int32(2))%32))))
	F_list_sort(m, v272, int32(477))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L11
	} else {
		goto L56
	}
L55:
	;
	F_list_sort(m, v240, int32(478))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L11
	} else {
		goto L58
	}
L56:
	;
	v277 = v259 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v277 < v278 {
		v259 = v277
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v289 = int32(1)
	v291 = v284
	v292 = v240
	goto L59
L59:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v298 <= v289 {
		v321 = v292
		goto L5
	} else {
		goto L61
	}
L60:
	;
	v321 = v314
	goto L5
L61:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v303 = v300 + v289<<(uint(int32(2))%32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v305 = F_equal(m, v304, v291)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L11
	} else {
		goto L63
	}
L62:
	;
	if v314 != 0 {
		v289 = v312
		v291 = v313
		v292 = v314
		goto L59
	} else {
		goto L68
	}
L63:
	;
	if v305 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v307 = F_list_delete_nth_cell(m, v292, v289)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L11
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v312 = v289 + int32(1)
	v313 = v311
	v314 = v292
	goto L62
L67:
	;
	v312 = v289
	v313 = v291
	v314 = v307
	goto L62
L68:
	;
	goto L60
L69:
	;
	return v240
}
func F_extract_grouping_collations(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v102
L2:
	;
	v14 = F_palloc(m, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = F_palloc(m, v18<<(uint(int32(2))%32))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v102 = v14
	goto L1
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23 <= int32(0) {
		v102 = v21
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v29 = v3
	goto L9
L9:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L22
	}
L11:
	;
	goto L10
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v38 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v42 = v29 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42+v43)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v53 = int32(0)
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v47+v53<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	if v46 != v63 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v70 = F_exprCollation(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L20
	}
L16:
	;
	v66 = v53 + int32(1)
	if v66 != v38 {
		v53 = v66
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	goto L11
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+v42))) = v70
	v74 = v29 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v74 < v75 {
		v29 = v74
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v102 = v21
	goto L1
L22:
	;
	F_errmsg_internal(m, int32(72784), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(487644), int32(366), int32(379475))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_extract_grouping_ops(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	v2 = int32(0)
	if l0 == v2 {
		v8 = F_palloc(m, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v37 = v8
			return v37
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = F_palloc(m, v12<<(uint(int32(2))%32))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 <= int32(0) {
				v37 = v15
			} else {
				v22 = v2
				for {
					v25 = v22 << (uint(int32(2)) % 32)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+v25)))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v15+v25))) = v30
					v33 = v22 + int32(1)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v33 < v34 {
						v22 = v33
						continue
					} else {
						break
					}
					break
				}
				v37 = v15
			}
			return v37
		}
	}
}
