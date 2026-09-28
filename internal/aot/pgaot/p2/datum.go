package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DatumGetAnyArrayP(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = base.I32_wrap_i64(l0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v4 != int32(1) {
		v15 = F_pg_detoast_datum(m, v3)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v15
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
		if v7&int32(254) != int32(2) {
			v15 = F_pg_detoast_datum(m, v3)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v15
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
			return v13
		}
	}
}
func F_datumCopy(m *base.Module, l0 int64, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int64
	_ = v65
	if l1 != 0 {
		v65 = l0
		return v65
	} else {
		if l2 == int32(-1) {
			v7 = base.I32_wrap_i64(l0)
			v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if v8 == int32(1) {
				v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
				if v11&int32(254) == int32(2) {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
					v18 = F_EOH_get_flat_size(m, v17)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int64(0)
					} else {
						v22 = F_palloc(m, v18)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							F_EOH_flatten_into(m, v17, v22, v18)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v22)
							}
						}
					}
				} else {
					v29 = int32(18)
					if v11 == v29 {
						v33 = v29
					} else {
						v33 = int32(2)
					}
					if base.Ui32((v11-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v40 = int32(6)
					} else {
						v40 = v33
					}
					v49 = v40
					v50 = F_palloc(m, v49)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						if v49 != 0 {
							base.MemoryCopy(m, v50, v7, v49)
						} else {
						}
						return base.I64_extend_i32_u(v50)
					}
				}
			} else {
				v41 = int32(1)
				if v8&v41 != 0 {
					v49 = int32(base.Ui32(v8) >> (uint(v41) % 32))
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v49 = int32(base.Ui32(v45) >> (uint(int32(2)) % 32))
				}
				v50 = F_palloc(m, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					if v49 != 0 {
						base.MemoryCopy(m, v50, v7, v49)
					} else {
					}
					return base.I64_extend_i32_u(v50)
				}
			}
		} else {
			v56 = F_datumGetSize(m, l0, int32(0), l2)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int64(0)
			} else {
				v58 = F_palloc(m, v56)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					if v56 != 0 {
						base.MemoryCopy(m, v58, base.I32_wrap_i64(l0), v56)
					} else {
					}
					v65 = base.I64_extend_i32_u(v58)
					return v65
				}
			}
		}
	}
}
func F_datum_image_eq(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
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
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v289
L2:
	;
	v282 = int32(255)
	v289 = base.B2i32(base.I32_wrap_i64(l0)&v282 == base.I32_wrap_i64(l1)&v282)
	goto L1
L3:
	;
	switch l3 - int32(1) {
	case 0:
		goto L2
	case 1:
		goto L8
	default:
		goto L6
	case 3:
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if int32(0) < l3 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v289 = base.B2i32(l0 == l1)
	goto L1
L7:
	;
	v289 = base.B2i32(base.I32_wrap_i64(l0) == base.I32_wrap_i64(l1))
	goto L1
L8:
	;
	v16 = int32(_a_F_datum_image_eq_0)
	v289 = base.B2i32(base.I32_wrap_i64(l0)&v16 == base.I32_wrap_i64(l1)&v16)
	goto L1
L9:
	;
	v28 = base.I32_wrap_i64(l0)
	v29 = base.I32_wrap_i64(l1)
	if base.Ui32(int32(4)) <= base.Ui32(l3) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	goto L11
L11:
	;
	switch l3 + int32(2) {
	case 0:
		goto L31
	case 1:
		goto L32
	default:
		goto L30
	}
L12:
	;
	v289 = base.B2i32(v91 == int32(0))
	goto L1
L13:
	;
	v91 = int32(0)
	goto L12
L14:
	;
	v65 = v60
	v66 = v61
	v67 = v62
	goto L24
L15:
	;
	if (v28|v29)&int32(3) != 0 {
		v60 = v28
		v61 = v29
		v62 = l3
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v53 = v28
	v54 = v29
	v55 = l3
	goto L17
L17:
	;
	if v55 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v37 = v28
	v38 = v29
	v39 = l3
	goto L19
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v42 != v43 {
		v60 = v37
		v61 = v38
		v62 = v39
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v53 = v48
	v54 = v46
	v55 = v50
	goto L17
L21:
	;
	v45 = int32(4)
	v46 = v38 + v45
	v48 = v37 + v45
	v50 = v39 - v45
	if base.Ui32(int32(3)) < base.Ui32(v50) {
		v37 = v48
		v38 = v46
		v39 = v50
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v60 = v53
	v61 = v54
	v62 = v55
	goto L14
L24:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v70 == v71 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v91 = v70 - v71
	goto L12
L26:
	;
	v73 = int32(1)
	v78 = v67 - v73
	if v78 != 0 {
		v65 = v65 + v73
		v66 = v66 + v73
		v67 = v78
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L33
	} else {
		goto L92
	}
L31:
	;
	v196 = base.I32_wrap_i64(l0)
	v197 = F_strlen(m, v196)
	mBase = m.M
	v198 = base.I32_wrap_i64(l1)
	v199 = F_strlen(m, v198)
	mBase = m.M
	if v197 != v199 {
		goto L71
	} else {
		goto L72
	}
L32:
	;
	v96 = F_toast_raw_datum_size(m, l0)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	v100 = F_toast_raw_datum_size(m, l1)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if v96 != v100 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v289 = int32(0)
	goto L1
L37:
	;
	goto L38
L38:
	;
	v104 = base.I32_wrap_i64(l0)
	v105 = F_pg_detoast_datum_packed(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v107 = base.I32_wrap_i64(l1)
	v108 = F_pg_detoast_datum_packed(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v110 = int32(1)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v112&v110 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v115 = v110
	goto L43
L42:
	;
	v115 = int32(4)
	goto L43
L43:
	;
	v116 = v105 + v115
	v117 = int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v119&v117 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v122 = v117
	goto L46
L45:
	;
	v122 = int32(4)
	goto L46
L46:
	;
	v123 = v108 + v122
	v124 = int32(4)
	v125 = v96 - v124
	if base.Ui32(v124) <= base.Ui32(v125) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	if v105 != v104 {
		goto L65
	} else {
		goto L66
	}
L48:
	;
	v187 = int32(0)
	goto L47
L49:
	;
	v161 = v156
	v162 = v157
	v163 = v158
	goto L59
L50:
	;
	if (v116|v123)&int32(3) != 0 {
		v156 = v116
		v157 = v123
		v158 = v125
		goto L49
	} else {
		goto L53
	}
L51:
	;
	v149 = v116
	v150 = v123
	v151 = v125
	goto L52
L52:
	;
	if v151 == int32(0) {
		goto L48
	} else {
		goto L58
	}
L53:
	;
	v133 = v116
	v134 = v123
	v135 = v125
	goto L54
L54:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v138 != v139 {
		v156 = v133
		v157 = v134
		v158 = v135
		goto L49
	} else {
		goto L56
	}
L55:
	;
	v149 = v144
	v150 = v142
	v151 = v146
	goto L52
L56:
	;
	v141 = int32(4)
	v142 = v134 + v141
	v144 = v133 + v141
	v146 = v135 - v141
	if base.Ui32(int32(3)) < base.Ui32(v146) {
		v133 = v144
		v134 = v142
		v135 = v146
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v156 = v149
	v157 = v150
	v158 = v151
	goto L49
L59:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v166 == v167 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v187 = v166 - v167
	goto L47
L61:
	;
	v169 = int32(1)
	v174 = v163 - v169
	if v174 != 0 {
		v161 = v161 + v169
		v162 = v162 + v169
		v163 = v174
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L48
L65:
	;
	F_pfree(m, v105)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L33
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v192 = base.B2i32(v187 == int32(0))
	if v108 == v107 {
		v289 = v192
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	F_pfree(m, v108)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L33
	} else {
		goto L70
	}
L70:
	;
	v289 = v192
	goto L1
L71:
	;
	v289 = int32(0)
	goto L1
L72:
	;
	goto L73
L73:
	;
	v203 = v197 + int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v203) {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v289 = base.B2i32(v265 == int32(0))
	goto L1
L75:
	;
	v265 = int32(0)
	goto L74
L76:
	;
	v239 = v234
	v240 = v235
	v241 = v236
	goto L86
L77:
	;
	if (v196|v198)&int32(3) != 0 {
		v234 = v196
		v235 = v198
		v236 = v203
		goto L76
	} else {
		goto L80
	}
L78:
	;
	v227 = v196
	v228 = v198
	v229 = v203
	goto L79
L79:
	;
	if v229 == int32(0) {
		goto L75
	} else {
		goto L85
	}
L80:
	;
	v211 = v196
	v212 = v198
	v213 = v203
	goto L81
L81:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v216 != v217 {
		v234 = v211
		v235 = v212
		v236 = v213
		goto L76
	} else {
		goto L83
	}
L82:
	;
	v227 = v222
	v228 = v220
	v229 = v224
	goto L79
L83:
	;
	v219 = int32(4)
	v220 = v212 + v219
	v222 = v211 + v219
	v224 = v213 - v219
	if base.Ui32(int32(3)) < base.Ui32(v224) {
		v211 = v222
		v212 = v220
		v213 = v224
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v234 = v227
	v235 = v228
	v236 = v229
	goto L76
L86:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v244 == v245 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v265 = v244 - v245
	goto L74
L88:
	;
	v247 = int32(1)
	v252 = v241 - v247
	if v252 != 0 {
		v239 = v239 + v247
		v240 = v240 + v247
		v241 = v252
		goto L86
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	goto L87
L91:
	;
	goto L75
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
	F_errmsg_internal(m, int32(_a_F_datum_image_eq_1), v11)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L33
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_datum_image_eq_2), int32(342), int32(_a_F_datum_image_eq_3))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L33
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
