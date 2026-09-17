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
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v16 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_list_sort(m, v229, int32(479))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L11
	} else {
		goto L67
	}
L5:
	;
	return v306
L6:
	;
	v24 = v4
	v27 = v4
	v30 = float64(1)
	goto L9
L7:
	;
	v62 = v4
	goto L8
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v24<<(uint(int32(2))%32))))
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
	v62 = v50
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
	v47 = base.F64_mul(v30, v46)
	if base.B2i32(int32(0) <= l2)&base.F64_lt(base.F64_convert_i32_u(l2), v47) != 0 {
		v306 = v4
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v50 = F_lappend(m, v27, v39)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v53 = v24 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v53 < v54 {
		v24 = v53
		v27 = v50
		v30 = v47
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v236 = int32(0)
	if base.B2i32(l1 == v236)|base.B2i32(v229 == v236) != 0 {
		goto L4
	} else {
		goto L50
	}
L20:
	;
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v69 < v70 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v115 = v4
	goto L22
L22:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v122 < int32(2) {
		v229 = v115
		goto L19
	} else {
		goto L32
	}
L23:
	;
	v76 = v69
	v77 = v4
	goto L26
L24:
	;
	v102 = v4
	goto L25
L25:
	;
	if v62 == int32(0) {
		v229 = v102
		goto L19
	} else {
		goto L31
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v76<<(uint(int32(2))%32))))
	v90 = F_list_union_int(m, int32(0), v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L11
	} else {
		goto L28
	}
L27:
	;
	v102 = v92
	goto L25
L28:
	;
	v92 = F_lappend(m, v77, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v95 = v76 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v95 < v96 {
		v76 = v95
		v77 = v92
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v115 = v102
	goto L22
L32:
	;
	v130 = v115
	v133 = int32(1)
	goto L33
L33:
	;
	v137 = int32(0)
	if v130 == v137 {
		v214 = v137
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v229 = v214
	goto L19
L35:
	;
	v222 = v133 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v222 < v223 {
		v130 = v214
		v133 = v222
		goto L33
	} else {
		goto L49
	}
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v140 <= int32(0) {
		v214 = v137
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143+v133<<(uint(int32(2))%32))))
	v151 = int32(0)
	v153 = v137
	goto L38
L38:
	;
	if v147 == int32(0) {
		v199 = v153
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v214 = v199
	goto L35
L40:
	;
	v207 = v151 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v207 < v208 {
		v151 = v207
		v153 = v199
		goto L38
	} else {
		goto L48
	}
L41:
	;
	v162 = int32(0)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v163 <= v162 {
		v199 = v153
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166+v151<<(uint(int32(2))%32))))
	v174 = v162
	v175 = v153
	goto L43
L43:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182+v174<<(uint(int32(2))%32))))
	v187 = F_list_union_int(m, v170, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L45
	}
L44:
	;
	v199 = v189
	goto L40
L45:
	;
	v189 = F_lappend(m, v175, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v192 = v174 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v192 < v193 {
		v174 = v192
		v175 = v189
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	goto L39
L49:
	;
	goto L34
L50:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v241 < int32(2) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v248 = int32(0)
	goto L52
L52:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256+v248<<(uint(int32(2))%32))))
	F_list_sort(m, v260, int32(477))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L11
	} else {
		goto L54
	}
L53:
	;
	F_list_sort(m, v229, int32(478))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L11
	} else {
		goto L56
	}
L54:
	;
	v265 = v248 + int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v265 < v266 {
		v248 = v265
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v274 = v272
	v277 = int32(1)
	v278 = v229
	goto L57
L57:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v285 <= v277 {
		v306 = v278
		goto L5
	} else {
		goto L59
	}
L58:
	;
	v306 = v301
	goto L5
L59:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v290 = v287 + v277<<(uint(int32(2))%32)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v292 = F_equal(m, v291, v274)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L11
	} else {
		goto L61
	}
L60:
	;
	if v301 != 0 {
		v274 = v299
		v277 = v300
		v278 = v301
		goto L57
	} else {
		goto L66
	}
L61:
	;
	if v292 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v294 = F_list_delete_nth_cell(m, v278, v277)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L11
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v299 = v298
	v300 = v277 + int32(1)
	v301 = v278
	goto L60
L65:
	;
	v299 = v274
	v300 = v277
	v301 = v294
	goto L60
L66:
	;
	goto L58
L67:
	;
	return v229
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
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
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = F_palloc(m, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = F_palloc(m, v19<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	return v14
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v24 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L23
	}
L8:
	;
	v31 = v3
	goto L11
L9:
	;
	goto L10
L10:
	;
	return v22
L11:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v39 <= int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v43 = v31 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v52 = int32(0)
	goto L15
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v48+v52<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	if v47 != v64 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v71 = F_exprCollation(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L21
	}
L17:
	;
	v67 = v52 + int32(1)
	if v67 != v39 {
		v52 = v67
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L7
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+v43))) = v71
	v75 = v31 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v75 < v76 {
		v31 = v75
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L12
L23:
	;
	F_errmsg_internal(m, int32(_a_F_extract_grouping_collations_0), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_extract_grouping_collations_1), int32(366), int32(_a_F_extract_grouping_collations_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v2 = int32(0)
	if l0 == v2 {
		v8 = F_palloc(m, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v8
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = F_palloc(m, v13<<(uint(int32(2))%32))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if int32(0) < v18 {
				v23 = v2
				for {
					v26 = v23 << (uint(int32(2)) % 32)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26)))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v16+v26))) = v31
					v34 = v23 + int32(1)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v34 < v35 {
						v23 = v34
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return v16
		}
	}
}
