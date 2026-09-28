package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_abbrev(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v3 = m.G0
	v5 = v3 + int32(-64)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v14&v12 != 0 {
			v17 = v12
		} else {
			v17 = int32(4)
		}
		v18 = v8 + v17
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
		v23 = F_pg_inet_net_ntop(m, v19, v18+int32(2), v22, v5)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v23 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_inet_abbrev_0), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_inet_abbrev_1), int32(1167), int32(_a_F_inet_abbrev_2))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v43 = F_cstring_to_text(m, v5)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					m.G0 = v5 - int32(-64)
					return base.I64_extend_i32_u(v43)
				}
			}
		}
	}
}
func F_inet_gist_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	v2 = int32(0)
	v6 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v8 != v10 {
		v91 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v6)))) = uint8(v91)
	return v6
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)))
	if v12 != v13 {
		v91 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+3)))
	if v15 != v16 {
		v91 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = int32(4)
	v19 = v7 + v18
	v21 = v9 + v18
	if v8 == int32(3) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = int32(16)
	goto L7
L6:
	;
	v26 = v18
	goto L7
L7:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v26) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v91 = base.B2i32(v88 == int32(0))
	goto L1
L9:
	;
	v88 = int32(0)
	goto L8
L10:
	;
	v62 = v57
	v63 = v58
	v64 = v59
	goto L20
L11:
	;
	if (v19|v21)&int32(3) != 0 {
		v57 = v19
		v58 = v21
		v59 = v26
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v50 = v19
	v51 = v21
	v52 = v26
	goto L13
L13:
	;
	if v52 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v34 = v19
	v35 = v21
	v36 = v26
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v39 != v40 {
		v57 = v34
		v58 = v35
		v59 = v36
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v50 = v45
	v51 = v43
	v52 = v47
	goto L13
L17:
	;
	v42 = int32(4)
	v43 = v35 + v42
	v45 = v34 + v42
	v47 = v36 - v42
	if base.Ui32(int32(3)) < base.Ui32(v47) {
		v34 = v45
		v35 = v43
		v36 = v47
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v57 = v50
	v58 = v51
	v59 = v52
	goto L10
L20:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 == v68 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v88 = v67 - v68
	goto L8
L22:
	;
	v70 = int32(1)
	v75 = v64 - v70
	if v75 != 0 {
		v62 = v62 + v70
		v63 = v63 + v70
		v64 = v75
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L9
}
func F_inet_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_network_recv(m, v2, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_inet_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_network_send(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_inet_spg_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v22&v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = v20
	goto L5
L4:
	;
	v25 = int32(4)
	goto L5
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v25)+1)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v28 <= int32(1) {
		v155 = v27
		v160 = int32(0)
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v167 = F_palloc_mul(m, int32(4), v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L41
	}
L7:
	;
	v32 = int32(1)
	v33 = v27
	goto L8
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v32<<(uint(int32(3))%32))))
	v46 = F_pg_detoast_datum_packed(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v155 = int32(0)
	v160 = v64
	goto L6
L10:
	;
	v48 = int32(1)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50&v48 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v48
	goto L13
L12:
	;
	v53 = int32(4)
	goto L13
L13:
	;
	v54 = v46 + v53
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v56 = int32(1)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v58&v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v56
	goto L16
L15:
	;
	v61 = int32(4)
	goto L16
L16:
	;
	v62 = v16 + v61
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v64 = base.B2i32(v55 != v63)
	if v55 != v63 {
		v155 = v33
		v160 = v64
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v65 = int32(2)
	v66 = v62 + v65
	v68 = v54 + v65
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v33 < v69 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v71 = v33
	goto L20
L19:
	;
	v71 = v69
	goto L20
L20:
	;
	v72 = int32(0)
	v77 = int32(8)
	v78 = base.I32_div_s(v71, v77)
	if v77 <= v71 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if v147 != 0 {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v147 = v144 + v140<<(uint(int32(3))%32)
	goto L21
L23:
	;
	v126 = v117
	goto L34
L24:
	;
	v84 = v72
	goto L27
L25:
	;
	v101 = v72
	goto L26
L26:
	;
	v108 = v71 - v78<<(uint(int32(3))%32)
	if v108 == int32(0) {
		v140 = v101
		v144 = v72
		goto L22
	} else {
		goto L33
	}
L27:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v84))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v84))))
	if v90 != v92 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v101 = v78
	goto L26
L29:
	;
	v117 = int32(7)
	v118 = v84
	v120 = v90
	v121 = v92
	goto L23
L30:
	;
	goto L31
L31:
	;
	v96 = v84 + int32(1)
	if v96 != v78 {
		v84 = v96
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v101))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v101))))
	v117 = v108
	v118 = v101
	v120 = v114
	v121 = v112
	goto L23
L34:
	;
	if int32(base.Ui32(v120^v121)>>(uint(int32(8)-v126)%32)) != 0 {
		v126 = v126 - int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v140 = v118
	v144 = v126
	goto L22
L36:
	;
	goto L35
L37:
	;
	v149 = v32 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v150 <= v149 {
		v155 = v147
		v160 = v64
		goto L6
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	goto L9
L40:
	;
	v32 = v149
	v33 = v147
	goto L8
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v167
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v172 = F_palloc_mul(m, int32(8), v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v172
	if v160 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	return int64(0)
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(2)
	v177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v177)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v180 <= v177 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v223 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v223)
	v225 = F_cidr_set_masklen_internal(m, v16, v155)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L55
	}
L47:
	;
	v184 = v177
	goto L48
L48:
	;
	v194 = v184 << (uint(int32(3)) % 32)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194+v195)))
	v198 = F_pg_detoast_datum_packed(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	goto L43
L50:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v204 = int32(1)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v206&v204 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v209 = v204
	goto L53
L52:
	;
	v209 = int32(4)
	goto L53
L53:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v209))))
	*(*int32)(unsafe.Add(mBase, uint32(v200+v184<<(uint(int32(2))%32)))) = base.B2i32(v211 != int32(2))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v215+v194))) = base.I64_extend_i32_u(v198)
	v220 = v184 + int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v220 < v221 {
		v184 = v220
		goto L48
	} else {
		goto L54
	}
L54:
	;
	goto L49
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = base.I64_extend_i32_u(v225)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v231 <= int32(0) {
		goto L43
	} else {
		goto L56
	}
L56:
	;
	v235 = base.I32_div_s(v155, int32(8))
	v243 = int32(0)
	goto L57
L57:
	;
	v256 = v243 << (uint(int32(3)) % 32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256+v257)))
	v260 = F_pg_detoast_datum_packed(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L43
L59:
	;
	v262 = int32(1)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v264&v262 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v267 = v262
	goto L62
L61:
	;
	v267 = int32(4)
	goto L62
L62:
	;
	v268 = v260 + v267
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v269 == int32(2) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v272 = int32(32)
	goto L65
L64:
	;
	v272 = int32(128)
	goto L65
L65:
	;
	if v155 < v272 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v235)+2)))
	v279 = int32(base.Ui32(v275)>>(uint(v235<<(uint(int32(3))%32)-v155+int32(7))%32)) & int32(1)
	goto L68
L67:
	;
	v279 = int32(0)
	goto L68
L68:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v281 = int32(2)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	if v155 < v286 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v288 = v279 | v281
	goto L71
L70:
	;
	v288 = v279
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280+v243<<(uint(v281)%32)))) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v290+v256))) = base.I64_extend_i32_u(v260)
	v295 = v243 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v295 < v296 {
		v243 = v295
		goto L57
	} else {
		goto L72
	}
L72:
	;
	goto L58
}
func F_inet_to_cidr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(1)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		if v21&v19 != 0 {
			v24 = v19
		} else {
			v24 = int32(4)
		}
		v25 = v15 + v24
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
		if v26 == int32(2) {
			v29 = int32(32)
		} else {
			v29 = int32(128)
		}
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
		if base.Ui32(v30) <= base.Ui32(v29) {
			v33 = F_palloc0(m, int32(22))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				v35 = int32(1)
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v37&v35 != 0 {
					v40 = v35
				} else {
					v40 = int32(4)
				}
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v40))))
				v43 = int32(1)
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
				if v45&v43 != 0 {
					v48 = v43
				} else {
					v48 = int32(4)
				}
				v49 = v33 + v48
				*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)) = uint8(v30)
				*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v42)
				if v30 == int32(0) {
				} else {
					v55 = v49 + int32(2)
					v61 = int32(base.Ui32((v30+int32(7))&int32(248)) >> (uint(int32(3)) % 32))
					if v61 != 0 {
						v62 = int32(1)
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if v64&v62 != 0 {
							v67 = v62
						} else {
							v67 = int32(4)
						}
						base.MemoryCopy(m, v55, v15+v67+int32(2), v61)
					} else {
					}
					v73 = v30 & int32(7)
					if v73 == int32(0) {
					} else {
						v78 = v55 + int32(base.Ui32(v30)>>(uint(int32(3))%32))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
						v82 = v79 & (int32(-256) >> (uint(v73) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v82)
					}
				}
				if v42 == int32(2) {
					v92 = int32(40)
				} else {
					v92 = int32(88)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = v92
				m.G0 = v10 + int32(16)
				return base.I64_extend_i32_u(v33)
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v30
				F_errmsg_internal(m, int32(_a_F_inet_to_cidr_0), v10)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_inet_to_cidr_1), int32(314), int32(_a_F_inet_to_cidr_2))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
