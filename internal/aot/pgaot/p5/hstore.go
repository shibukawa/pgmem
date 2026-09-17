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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
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
				F_errmsg(m, int32(_a_F_hstoreCheckKeyLen_0), int32(0))
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_hstoreCheckKeyLen_1), int32(413), int32(_a_F_hstoreCheckKeyLen_2))
					v21 = m.ExcPending
					if v21 != 0 {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
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
				F_errmsg(m, int32(_a_F_hstoreCheckValLen_0), int32(0))
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_hstoreCheckValLen_1), int32(433), int32(_a_F_hstoreCheckValLen_2))
					v21 = m.ExcPending
					if v21 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(_a_F_hstore_exec_setup_0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = int32(_a_F_hstore_exec_setup_1)
	return
}
func F_hstore_fetchval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_hstoreUpgrade(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(1)
	v23 = v20 + v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v26 = v24 & v22
	if v24 == v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v56 = v54 & int32(268435455)
	if v56 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v32 == int32(18) {
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
	if v26 != 0 {
		v53 = int32(base.Ui32(v24)>>(uint(v43)%32)) - v43
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v35 = int32(16)
	goto L10
L9:
	;
	v35 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v42 = int32(4)
	goto L13
L12:
	;
	v42 = v35
	goto L13
L13:
	;
	v53 = v42
	goto L4
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v214)
	return int32(0)
L16:
	;
	if v26 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v23
	goto L19
L18:
	;
	v61 = v20 + int32(4)
	goto L19
L19:
	;
	v63 = v15 + int32(8)
	v75 = v56
	v76 = int32(0)
	goto L20
L20:
	;
	v83 = int32(base.Ui32(v75-v76)>>(uint(int32(1))%32)) + v76
	v86 = v63 + v83<<(uint(int32(3))%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v89 = v87 & int32(1073741823)
	if int32(0) <= v87 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	goto L15
L22:
	;
	v197 = base.B2i32(v193 < int32(0))
	if v193 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L23:
	;
	v109 = v108 + (v63 + v56<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v53) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	if base.Ui32(v53) < base.Ui32(v101) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v86-int32(4))))
	v96 = v94 & int32(1073741823)
	v97 = v89 - v96
	if v97 != v53 {
		v101 = v97
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v89 == v53 {
		v108 = int32(0)
		goto L23
	} else {
		goto L29
	}
L28:
	;
	v108 = v96
	goto L23
L29:
	;
	v101 = v89
	goto L24
L30:
	;
	v106 = int32(1)
	goto L32
L31:
	;
	v106 = int32(-1)
	goto L32
L32:
	;
	v193 = v106
	goto L22
L33:
	;
	if v171 != 0 {
		v193 = v171
		goto L22
	} else {
		goto L51
	}
L34:
	;
	v171 = int32(0)
	goto L33
L35:
	;
	v145 = v140
	v146 = v141
	v147 = v142
	goto L45
L36:
	;
	if (v109|v61)&int32(3) != 0 {
		v140 = v109
		v141 = v61
		v142 = v53
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v133 = v109
	v134 = v61
	v135 = v53
	goto L38
L38:
	;
	if v135 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v117 = v109
	v118 = v61
	v119 = v53
	goto L40
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v122 != v123 {
		v140 = v117
		v141 = v118
		v142 = v119
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v133 = v128
	v134 = v126
	v135 = v130
	goto L38
L42:
	;
	v125 = int32(4)
	v126 = v118 + v125
	v128 = v117 + v125
	v130 = v119 - v125
	if base.Ui32(int32(3)) < base.Ui32(v130) {
		v117 = v128
		v118 = v126
		v119 = v130
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v140 = v133
	v141 = v134
	v142 = v135
	goto L35
L45:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v150 == v151 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v171 = v150 - v151
	goto L33
L47:
	;
	v153 = int32(1)
	v158 = v147 - v153
	if v158 != 0 {
		v145 = v145 + v153
		v146 = v146 + v153
		v147 = v158
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	goto L34
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v172&int32(1073741824) != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	v180 = int32(0)
	v182 = base.B2i32(v180 <= v172)
	if v180 <= v172 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v183 = v89
	goto L55
L54:
	;
	v183 = v180
	goto L55
L55:
	;
	if v180 <= v172 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v188 = v172 - v89
	goto L58
L57:
	;
	v188 = v172 & int32(1073741823)
	goto L58
L58:
	;
	v189 = F_cstring_to_text_with_len(m, v63+v54<<(uint(int32(3))%32)&int32(2147483640)+v183, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	return v189
L60:
	;
	v198 = v83 + int32(1)
	goto L62
L61:
	;
	v198 = v76
	goto L62
L62:
	;
	if v193 < int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v199 = v75
	goto L65
L64:
	;
	v199 = v83
	goto L65
L65:
	;
	if v198 < v199 {
		v75 = v199
		v76 = v198
		goto L20
	} else {
		goto L66
	}
L66:
	;
	goto L21
}
func F_hstore_from_arrays(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
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
	var v140 int32
	_ = v140
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
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v16 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L12
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L12
	} else {
		goto L69
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L65
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L12
	} else {
		goto L61
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L12
	} else {
		goto L57
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L12
	} else {
		goto L53
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L12
	} else {
		goto L49
	}
L8:
	;
	m.G0 = v14 + int32(48)
	return v221
L9:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
	v221 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if int32(2) <= v27 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	F_deconstruct_array_builtin(m, v23, int32(25), v14+int32(40), v14+int32(36), v14+int32(32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if base.Ui32(int32(53687092)) <= base.Ui32(v39) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v42 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v95 = F_palloc(m, v91*int32(20))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L30
	}
L18:
	;
	v45 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v45
	v91 = v39
	goto L17
L19:
	;
	goto L20
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v51 = F_pg_detoast_datum(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if int32(2) <= v53 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v59 = int32(0)
	if base.B2i32(v53 != int32(1))&base.B2i32(v58 <= v59) == v59 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v53 != v58 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_deconstruct_array_builtin(m, v51, int32(25), v14+int32(28), v14+int32(24), v14+int32(20))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L29
	}
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v65 != v66 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v69 = v53 << (uint(int32(2)) % 32)
	v70 = int32(16)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+(v23+v70))))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v51+v70+v69)))
	if v73 != v77 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v91 = v89
	goto L17
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if int32(0) < v97 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v105 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	v204 = F_hstoreUniquePairs(m, v95, v97, v14+int32(44))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L47
	}
L34:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v103))))
	if v117 == int32(1) {
		goto L3
	} else {
		goto L36
	}
L35:
	;
	goto L33
L36:
	;
	if v102 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v183 = v95 + v105*int32(20)
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+17)) = uint8(v184)
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+16)) = uint8(v178)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v179
	v189 = v105 + int32(1)
	if v189 != v97 {
		v105 = v189
		goto L34
	} else {
		goto L46
	}
L38:
	;
	v148 = v95 + v105*int32(20)
	v149 = int32(2)
	v150 = v105 << (uint(v149) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v101+v150)))
	v153 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v152 + v153
	v156 = v150 + v100
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v157 + v153
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v165 = int32(base.Ui32(v161)>>(uint(v149)%32)) - v153
	if base.Ui32(int32(1073741824)) <= base.Ui32(v165) {
		goto L1
	} else {
		goto L44
	}
L39:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v102))))
	if v121 != int32(1) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v124 = int32(2)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v101+v105<<(uint(v124)%32))))
	v130 = v95 + v105*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = int32(0)
	v133 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v127 + v133
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v140 = int32(base.Ui32(v136)>>(uint(v124)%32)) - v133
	if base.Ui32(int32(1073741824)) <= base.Ui32(v140) {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v140
	v178 = int32(1)
	v179 = int32(4)
	goto L37
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+8)) = v165
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v175 = int32(base.Ui32(v171)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v175) {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v178 = int32(0)
	v179 = v175
	goto L37
L46:
	;
	goto L35
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v204
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v208 = F_hstorePairs(m, v95, v204, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v221 = v208
	goto L8
L49:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_0), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(633), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(53687091)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v251
	F_errmsg(m, int32(_a_F_hstore_from_arrays_3), v14)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(642), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L12
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_0), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(662), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_4), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(670), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_5), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(684), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L12
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
	F_errcode(m, int32(16777346))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_6), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(433), int32(_a_F_hstore_from_arrays_7))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_8), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(413), int32(_a_F_hstore_from_arrays_9))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hstore_svals(m *base.Module, l0 int32) int32 {
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = F_hstoreUpgrade(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = int32(_a_F_hstore_svals_0)
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_hstore_svals[0]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_hstore_svals[0])) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v27 = F_palloc(m, int32(base.Ui32(v24)>>(uint(int32(2))%32)))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v31 = int32(base.Ui32(v29) >> (uint(int32(2)) % 32))
					if v31 != 0 {
						base.MemoryCopy(m, v27, v13, v31)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v27
					*(*int32)(unsafe.Add(mBase, _c_F_hstore_svals[0])) = v20
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v46 = v44 & int32(268435455)
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
					v48 = base.I32_wrap_i64(v47)
					if base.Ui32(v48) < base.Ui32(v46) {
						v51 = v43 + int32(8)
						v54 = v51 + v48<<(uint(int32(3))%32)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
						if v55&int32(1073741824) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v42))) = v47 + int64(1)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v62 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v62
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v62)
							return int32(0)
						} else {
							v70 = v51 + v46<<(uint(int32(3))%32)
							if v55 < int32(0) {
								v80 = v55 & int32(1073741823)
								v81 = v70
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
								v77 = v75 & int32(1073741823)
								v80 = v55 - v77
								v81 = v70 + v77
							}
							v83 = F_cstring_to_text_with_len(m, v81, v80)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
								*(*int64)(unsafe.Add(mBase, uint32(v42))) = v85 + int64(1)
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = int32(1)
								return v83
							}
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = int32(2)
							v98 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v98)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
		v46 = v44 & int32(268435455)
		v47 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
		v48 = base.I32_wrap_i64(v47)
		if base.Ui32(v48) < base.Ui32(v46) {
			v51 = v43 + int32(8)
			v54 = v51 + v48<<(uint(int32(3))%32)
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
			if v55&int32(1073741824) != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v42))) = v47 + int64(1)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v62 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v62
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v62)
				return int32(0)
			} else {
				v70 = v51 + v46<<(uint(int32(3))%32)
				if v55 < int32(0) {
					v80 = v55 & int32(1073741823)
					v81 = v70
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
					v77 = v75 & int32(1073741823)
					v80 = v55 - v77
					v81 = v70 + v77
				}
				v83 = F_cstring_to_text_with_len(m, v81, v80)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v85 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
					*(*int64)(unsafe.Add(mBase, uint32(v42))) = v85 + int64(1)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = int32(1)
					return v83
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = int32(2)
				v98 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v98)
				return int32(0)
			}
		}
	}
}
