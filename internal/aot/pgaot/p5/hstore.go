package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstoreCheckKeyLen(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	if base.Ui32(int32(1073741824)) <= base.Ui32(l0) {
		F_errstart_cold(m, int32(21), int32(0))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16777346))
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(22600), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(516125), int32(413), int32(293763))
					v25 = m.ExcPending
					if v25 != 0 {
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
		return l0
	}
}
func F_hstoreCheckValLen(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	if base.Ui32(int32(1073741824)) <= base.Ui32(l0) {
		F_errstart_cold(m, int32(21), int32(0))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16777346))
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(360669), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(516125), int32(433), int32(293781))
					v25 = m.ExcPending
					if v25 != 0 {
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
		return l0
	}
}
func F_hstore_akeys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_hstoreUpgrade(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v15 = v13 & int32(268435455)
	if v15 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v23 = v9 + int32(8)
	v26 = v23 + v15<<(uint(int32(3))%32)
	v30 = F_palloc(m, v15<<(uint(int32(2))%32))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	return v19
L7:
	;
	v32 = int32(0)
	goto L8
L8:
	;
	v44 = v23 + v32<<(uint(int32(3))%32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v69 = F_construct_array_builtin(m, v30, v15, int32(25))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	v62 = F_cstring_to_text_with_len(m, v61, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v59 = v45 & int32(1073741823)
	v61 = v26
	goto L10
L12:
	;
	goto L13
L13:
	;
	v50 = int32(1073741823)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44-int32(4))))
	v56 = v54 & v50
	v59 = v45&v50 - v56
	v61 = v56 + v26
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+v32<<(uint(int32(2))%32)))) = v62
	v66 = v32 + int32(1)
	if v66 != v15 {
		v32 = v66
		goto L8
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	return v69
}
func F_hstore_exec_setup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(7351)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = int32(7352)
	return
}
func F_hstore_fetchval(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_hstoreUpgrade(m, v13)
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
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(1)
	v22 = v19 + v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v25 = v23 & v21
	if v23 == v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v56 = v54 & int32(268435455)
	if v56 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v28 = int32(4)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v30&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v39 = v28
	goto L10
L9:
	;
	v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
	goto L10
L10:
	;
	if v30 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v42 = v28
	goto L13
L12:
	;
	v42 = v39
	goto L13
L13:
	;
	v53 = v42
	goto L4
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v178-int32(4))))
	v223 = v221 & int32(1073741823)
	v226 = F_cstring_to_text_with_len(m, v186+v223, v179-v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L64
	}
L16:
	;
	v215 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v215)
	return int32(0)
L17:
	;
	if v25 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = v22
	goto L20
L19:
	;
	v61 = v19 + int32(4)
	goto L20
L20:
	;
	v63 = v14 + int32(8)
	v70 = v56
	v71 = int32(0)
	goto L21
L21:
	;
	v82 = base.I32_div_s(v70-v71, int32(2))
	v83 = v82 + v71
	v86 = v63 + v83<<(uint(int32(3))%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v89 = v87 & int32(1073741823)
	if int32(0) <= v87 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L16
L23:
	;
	v199 = base.B2i32(v194 < int32(0))
	if v194 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L24:
	;
	v109 = v108 + (v63 + v56<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v53) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	if base.Ui32(v53) < base.Ui32(v101) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v86-int32(4))))
	v96 = v94 & int32(1073741823)
	v97 = v89 - v96
	if v97 != v53 {
		v101 = v97
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v89 == v53 {
		v108 = int32(0)
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v108 = v96
	goto L24
L30:
	;
	v101 = v89
	goto L25
L31:
	;
	v106 = int32(1)
	goto L33
L32:
	;
	v106 = int32(-1)
	goto L33
L33:
	;
	v194 = v106
	goto L23
L34:
	;
	if v171 != 0 {
		v194 = v171
		goto L23
	} else {
		goto L52
	}
L35:
	;
	v171 = int32(0)
	goto L34
L36:
	;
	v145 = v140
	v146 = v141
	v147 = v142
	goto L46
L37:
	;
	if (v109|v61)&int32(3) != 0 {
		v140 = v109
		v141 = v61
		v142 = v53
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v133 = v109
	v134 = v61
	v135 = v53
	goto L39
L39:
	;
	if v135 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v117 = v109
	v118 = v61
	v119 = v53
	goto L41
L41:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v122 != v123 {
		v140 = v117
		v141 = v118
		v142 = v119
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v133 = v128
	v134 = v126
	v135 = v130
	goto L39
L43:
	;
	v125 = int32(4)
	v126 = v118 + v125
	v128 = v117 + v125
	v130 = v119 - v125
	if base.Ui32(int32(3)) < base.Ui32(v130) {
		v117 = v128
		v118 = v126
		v119 = v130
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v140 = v133
	v141 = v134
	v142 = v135
	goto L36
L46:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v150 == v151 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v171 = v150 - v151
	goto L34
L48:
	;
	v153 = int32(1)
	v158 = v147 - v153
	if v158 != 0 {
		v145 = v145 + v153
		v146 = v146 + v153
		v147 = v158
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	if v83 < int32(0) {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v178 = v63 + v83<<(uint(int32(3))%32) + int32(4)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v179&int32(1073741824) != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	v186 = v63 + v54<<(uint(int32(3))%32)&int32(2147483640)
	if int32(0) <= v179 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	v191 = F_cstring_to_text_with_len(m, v186, v179&int32(1073741823))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	return v191
L57:
	;
	v200 = v83 + int32(1)
	goto L59
L58:
	;
	v200 = v71
	goto L59
L59:
	;
	if v194 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v201 = v70
	goto L62
L61:
	;
	v201 = v83
	goto L62
L62:
	;
	if v200 < v201 {
		v70 = v201
		v71 = v200
		goto L21
	} else {
		goto L63
	}
L63:
	;
	goto L22
L64:
	;
	return v226
}
func F_hstore_from_arrays(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
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
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v17 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L78
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L13
	} else {
		goto L74
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L13
	} else {
		goto L70
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L13
	} else {
		goto L66
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L13
	} else {
		goto L62
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L13
	} else {
		goto L58
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L54
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L50
	}
L9:
	;
	m.G0 = v15 + int32(48)
	return v227
L10:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
	v227 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(2) <= v28 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	F_deconstruct_array_builtin(m, v24, int32(25), v15+int32(40), v15+int32(36), v15+int32(32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if base.Ui32(int32(53687092)) <= base.Ui32(v40) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v43 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v97 = F_palloc(m, v93*int32(20))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L31
	}
L19:
	;
	v46 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v46
	v93 = v40
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v52 = F_pg_detoast_datum(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if int32(2) <= v54 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v60 = int32(0)
	if base.B2i32(v54 != int32(1))&base.B2i32(v59 <= v60) == v60 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v54 != v59 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	F_deconstruct_array_builtin(m, v52, int32(25), v15+int32(28), v15+int32(24), v15+int32(20))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L13
	} else {
		goto L30
	}
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v66 != v67 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v69 = int32(16)
	v72 = v54 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24+v69+v72)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v52+v69+v72)))
	if v74 != v78 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v93 = v90
	goto L18
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if int32(0) < v99 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v107 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	v209 = F_hstoreUniquePairs(m, v97, v99, v15+int32(44))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L13
	} else {
		goto L48
	}
L35:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v105))))
	if v120 == int32(1) {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	if v104 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v187 = v97 + v107*int32(20)
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+17)) = uint8(v188)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+16)) = uint8(v182)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+12)) = v183
	v193 = v107 + int32(1)
	if v193 != v99 {
		v107 = v193
		goto L35
	} else {
		goto L47
	}
L39:
	;
	v151 = v97 + v107*int32(20)
	v152 = int32(2)
	v153 = v107 << (uint(v152) % 32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v103+v153)))
	v156 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v155 + v156
	v159 = v153 + v102
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v160 + v156
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v168 = int32(base.Ui32(v164)>>(uint(v152)%32)) - v156
	if base.Ui32(int32(1073741824)) <= base.Ui32(v168) {
		goto L2
	} else {
		goto L45
	}
L40:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v104))))
	if v124 != int32(1) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v127 = int32(2)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v103+v107<<(uint(v127)%32))))
	v133 = v97 + v107*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = int32(0)
	v136 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v130 + v136
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v143 = int32(base.Ui32(v139)>>(uint(v127)%32)) - v136
	if base.Ui32(int32(1073741824)) <= base.Ui32(v143) {
		goto L3
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v143
	v182 = int32(1)
	v183 = int32(4)
	goto L38
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v168
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v178 = int32(base.Ui32(v174)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v178) {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v182 = int32(0)
	v183 = v178
	goto L38
L47:
	;
	goto L36
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v209
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v213 = F_hstorePairs(m, v97, v209, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v227 = v213
	goto L9
L50:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(125349), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(516125), int32(633), int32(120156))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
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
	F_errcode(m, int32(261))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(53687091)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v261
	F_errmsg(m, int32(705224), v15)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(516125), int32(642), int32(120156))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(125349), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(516125), int32(662), int32(120156))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L13
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(181354), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(516125), int32(670), int32(120156))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L13
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(22631), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(516125), int32(684), int32(120156))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L13
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
	F_errcode(m, int32(16777346))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(22600), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(516125), int32(413), int32(293763))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(22600), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(516125), int32(413), int32(293763))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L13
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
	F_errcode(m, int32(16777346))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(360669), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(516125), int32(433), int32(293781))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hstore_svals(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v10 == int32(0) {
		v13 = int32(4549024)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = F_hstoreUpgrade(m, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[10]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
				*(*int32)(unsafe.Add(mBase, _consts[10])) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v27 = F_palloc(m, int32(base.Ui32(v24)>>(uint(int32(2))%32)))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v31 = int32(base.Ui32(v29) >> (uint(int32(2)) % 32))
					if v31 != 0 {
						v32 = F__emscripten_memcpy_bulkmem(m, v27, v15, v31)
						mBase = m.M
						v33 = v32
					} else {
						v33 = v27
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v33
					*(*int32)(unsafe.Add(mBase, _consts[10])) = v21
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
					v45 = v43 & int32(268435455)
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
					v47 = base.I32_wrap_i64(v46)
					if base.Ui32(v47) < base.Ui32(v45) {
						v50 = v42 + int32(8)
						v55 = v50 + v47<<(uint(int32(3))%32) + int32(4)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						if v56&int32(1073741824) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v41))) = v46 + int64(1)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v63 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = v63
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
							return int32(0)
						} else {
							v71 = v50 + v45<<(uint(int32(3))%32)
							if v56 < int32(0) {
								v83 = v56 & int32(1073741823)
								v84 = v71
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(4))))
								v80 = v78 & int32(1073741823)
								v83 = v56 - v80
								v84 = v71 + v80
							}
							v86 = F_cstring_to_text_with_len(m, v84, v83)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								v88 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
								*(*int64)(unsafe.Add(mBase, uint32(v41))) = v88 + int64(1)
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(1)
								return v86
							}
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
							v101 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
		v45 = v43 & int32(268435455)
		v46 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
		v47 = base.I32_wrap_i64(v46)
		if base.Ui32(v47) < base.Ui32(v45) {
			v50 = v42 + int32(8)
			v55 = v50 + v47<<(uint(int32(3))%32) + int32(4)
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
			if v56&int32(1073741824) != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = v46 + int64(1)
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v63 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = v63
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
				return int32(0)
			} else {
				v71 = v50 + v45<<(uint(int32(3))%32)
				if v56 < int32(0) {
					v83 = v56 & int32(1073741823)
					v84 = v71
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(4))))
					v80 = v78 & int32(1073741823)
					v83 = v56 - v80
					v84 = v71 + v80
				}
				v86 = F_cstring_to_text_with_len(m, v84, v83)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
					*(*int64)(unsafe.Add(mBase, uint32(v41))) = v88 + int64(1)
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(1)
					return v86
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return int32(0)
			} else {
				v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
				v101 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
				return int32(0)
			}
		}
	}
}
