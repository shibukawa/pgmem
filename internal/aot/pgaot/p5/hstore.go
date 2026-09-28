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
func F_hstore_akeys(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_hstoreUpgrade(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v16 = v14 & int32(268435455)
	if v16 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v25 = v10 + int32(8)
	v27 = v16 << (uint(int32(3)) % 32)
	v28 = v25 + v27
	v30 = F_palloc(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	return base.I64_extend_i32_u(v20)
L7:
	;
	v32 = int32(0)
	goto L8
L8:
	;
	v41 = v32 << (uint(int32(3)) % 32)
	v42 = v25 + v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v69 = F_construct_array_builtin(m, v30, v16, int32(25))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	v61 = F_cstring_to_text_with_len(m, v59, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v58 = v43 & int32(1073741823)
	v59 = v28
	goto L10
L12:
	;
	goto L13
L13:
	;
	v48 = int32(1073741823)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42-int32(4))))
	v54 = v52 & v48
	v58 = v43&v48 - v54
	v59 = v54 + v28
	goto L10
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30+v41))) = base.I64_extend_i32_u(v61)
	v66 = v32 + int32(1)
	if v66 != v16 {
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
	return base.I64_extend_i32_u(v69)
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
func F_hstore_fetchval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int64
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
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
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v23 = int32(1)
	v24 = v22 & v23
	if v22 == v23 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v54 = v52 & int32(268435455)
	if v54 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v30 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v24 != 0 {
		v51 = int32(base.Ui32(v22)>>(uint(v41)%32)) - v41
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v33 = int32(16)
	goto L10
L9:
	;
	v33 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = int32(4)
	goto L13
L12:
	;
	v40 = v33
	goto L13
L13:
	;
	v51 = v40
	goto L4
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v214)
	return int64(0)
L16:
	;
	if v24 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v59 = int32(1)
	goto L19
L18:
	;
	v59 = int32(4)
	goto L19
L19:
	;
	v60 = v20 + v59
	v62 = v15 + int32(8)
	v73 = v54
	v74 = int32(0)
	goto L20
L20:
	;
	v82 = int32(base.Ui32(v73-v74)>>(uint(int32(1))%32)) + v74
	v85 = v62 + v82<<(uint(int32(3))%32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v88 = v86 & int32(1073741823)
	if int32(0) <= v86 {
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
	v108 = v107 + (v62 + v54<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	if base.Ui32(v51) < base.Ui32(v100) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v85-int32(4))))
	v95 = v93 & int32(1073741823)
	v96 = v88 - v95
	if v96 != v51 {
		v100 = v96
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v88 == v51 {
		v107 = int32(0)
		goto L23
	} else {
		goto L29
	}
L28:
	;
	v107 = v95
	goto L23
L29:
	;
	v100 = v88
	goto L24
L30:
	;
	v105 = int32(1)
	goto L32
L31:
	;
	v105 = int32(-1)
	goto L32
L32:
	;
	v193 = v105
	goto L22
L33:
	;
	if v170 != 0 {
		v193 = v170
		goto L22
	} else {
		goto L51
	}
L34:
	;
	v170 = int32(0)
	goto L33
L35:
	;
	v144 = v139
	v145 = v140
	v146 = v141
	goto L45
L36:
	;
	if (v108|v60)&int32(3) != 0 {
		v139 = v108
		v140 = v60
		v141 = v51
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v132 = v108
	v133 = v60
	v134 = v51
	goto L38
L38:
	;
	if v134 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v116 = v108
	v117 = v60
	v118 = v51
	goto L40
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v121 != v122 {
		v139 = v116
		v140 = v117
		v141 = v118
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v132 = v127
	v133 = v125
	v134 = v129
	goto L38
L42:
	;
	v124 = int32(4)
	v125 = v117 + v124
	v127 = v116 + v124
	v129 = v118 - v124
	if base.Ui32(int32(3)) < base.Ui32(v129) {
		v116 = v127
		v117 = v125
		v118 = v129
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v139 = v132
	v140 = v133
	v141 = v134
	goto L35
L45:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v149 == v150 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v170 = v149 - v150
	goto L33
L47:
	;
	v152 = int32(1)
	v157 = v146 - v152
	if v157 != 0 {
		v144 = v144 + v152
		v145 = v145 + v152
		v146 = v157
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
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v171&int32(1073741824) != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	v179 = int32(0)
	v181 = base.B2i32(v179 <= v171)
	if v179 <= v171 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v182 = v88
	goto L55
L54:
	;
	v182 = v179
	goto L55
L55:
	;
	if v179 <= v171 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v187 = v171 - v88
	goto L58
L57:
	;
	v187 = v171 & int32(1073741823)
	goto L58
L58:
	;
	v188 = F_cstring_to_text_with_len(m, v62+v52<<(uint(int32(3))%32)&int32(2147483640)+v182, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	return base.I64_extend_i32_u(v188)
L60:
	;
	v198 = v82 + int32(1)
	goto L62
L61:
	;
	v198 = v74
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
	v199 = v73
	goto L65
L64:
	;
	v199 = v82
	goto L65
L65:
	;
	if v198 < v199 {
		v73 = v199
		v74 = v198
		goto L20
	} else {
		goto L66
	}
L66:
	;
	goto L21
}
func F_hstore_from_arrays(m *base.Module, l0 int32) int64 {
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v225 int64
	_ = v225
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v17 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L12
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L12
	} else {
		goto L69
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L12
	} else {
		goto L65
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L12
	} else {
		goto L61
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L12
	} else {
		goto L57
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L12
	} else {
		goto L53
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L12
	} else {
		goto L49
	}
L8:
	;
	m.G0 = v15 + int32(48)
	return v225
L9:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
	v225 = int64(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int64(0)
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(2) <= v28 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	F_deconstruct_array_builtin(m, v24, int32(25), v15+int32(40), v15+int32(36), v15+int32(32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if base.Ui32(int32(53687092)) <= base.Ui32(v40) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v43 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v96 = F_palloc(m, v92*int32(20))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L12
	} else {
		goto L30
	}
L18:
	;
	v46 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v46
	v92 = v40
	goto L17
L19:
	;
	goto L20
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v52 = F_pg_detoast_datum(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if int32(2) <= v54 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v60 = int32(0)
	if base.B2i32(v54 != int32(1))&base.B2i32(v59 <= v60) == v60 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v54 != v59 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_deconstruct_array_builtin(m, v52, int32(25), v15+int32(28), v15+int32(24), v15+int32(20))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L12
	} else {
		goto L29
	}
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v66 != v67 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v70 = v54 << (uint(int32(2)) % 32)
	v71 = int32(16)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v24+v71))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v52+v71+v70)))
	if v74 != v78 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v92 = v90
	goto L17
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if int32(0) < v98 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v106 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	v206 = F_hstoreUniquePairs(m, v96, v98, v15+int32(44))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L12
	} else {
		goto L47
	}
L34:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+v104))))
	if v119 == int32(1) {
		goto L3
	} else {
		goto L36
	}
L35:
	;
	goto L33
L36:
	;
	if v103 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v184 = v96 + v106*int32(20)
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v184)+17)) = uint8(v185)
	*(*uint8)(unsafe.Add(mBase, uint32(v184)+16)) = uint8(v179)
	*(*int32)(unsafe.Add(mBase, uint32(v184)+12)) = v180
	v190 = v106 + int32(1)
	if v190 != v98 {
		v106 = v190
		goto L34
	} else {
		goto L46
	}
L38:
	;
	v150 = v96 + v106*int32(20)
	v152 = v106 << (uint(int32(3)) % 32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v102+v152)))
	v155 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v154 + v155
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v152+v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v159 + v155
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v167 = int32(base.Ui32(v163)>>(uint(int32(2))%32)) - v155
	if base.Ui32(int32(1073741824)) <= base.Ui32(v167) {
		goto L1
	} else {
		goto L44
	}
L39:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+v103))))
	if v123 != int32(1) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v102+v106<<(uint(int32(3))%32))))
	v132 = v96 + v106*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = int32(0)
	v135 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v129 + v135
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v142 = int32(base.Ui32(v138)>>(uint(int32(2))%32)) - v135
	if base.Ui32(int32(1073741824)) <= base.Ui32(v142) {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v142
	v179 = int32(1)
	v180 = int32(4)
	goto L37
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v167
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v176 = int32(base.Ui32(v172)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v176) {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v179 = int32(0)
	v180 = v176
	goto L37
L46:
	;
	goto L35
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v206
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v210 = F_hstorePairs(m, v96, v206, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v225 = base.I64_extend_i32_u(v210)
	goto L8
L49:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_0), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(633), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
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
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(53687091)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v255
	F_errmsg(m, int32(_a_F_hstore_from_arrays_3), v15)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(642), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
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
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_0), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(662), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
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
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_4), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(670), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
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
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_5), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(684), int32(_a_F_hstore_from_arrays_2))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
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
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_6), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(433), int32(_a_F_hstore_from_arrays_7))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
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
	v339 = m.ExcPending
	if v339 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_hstore_from_arrays_8), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_hstore_from_arrays_1), int32(413), int32(_a_F_hstore_from_arrays_9))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
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
func F_hstore_svals(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int64
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = F_hstoreUpgrade(m, v12)
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
				v19 = int32(_a_F_hstore_svals_0)
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_hstore_svals[0]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_hstore_svals[0])) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v27 = F_palloc(m, int32(base.Ui32(v24)>>(uint(int32(2))%32)))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
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
							return int64(0)
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
								return int64(0)
							} else {
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
								*(*int64)(unsafe.Add(mBase, uint32(v42))) = v85 + int64(1)
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = int32(1)
								return base.I64_extend_i32_u(v83)
							}
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int64(0)
						} else {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v96)+20)) = int32(2)
							v99 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v99)
							return int64(0)
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
				return int64(0)
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
					return int64(0)
				} else {
					v85 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
					*(*int64)(unsafe.Add(mBase, uint32(v42))) = v85 + int64(1)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = int32(1)
					return base.I64_extend_i32_u(v83)
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int64(0)
			} else {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v96)+20)) = int32(2)
				v99 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v99)
				return int64(0)
			}
		}
	}
}
