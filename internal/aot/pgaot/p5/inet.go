package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_abbrev(m *base.Module, l0 int32) int32 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
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
			return int32(0)
		} else {
			if v23 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_inet_abbrev_0), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_inet_abbrev_1), int32(1199), int32(_a_F_inet_abbrev_2))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
				v43 = F_cstring_to_text(m, v5)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 - int32(-64)
					return v43
				}
			}
		}
	}
}
func F_inet_gist_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v8 != v10 {
		v91 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v91)
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
func F_inet_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_network_recv(m, v2, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_inet_send(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_network_send(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_inet_spg_picksplit(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
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
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v20 = int32(1)
	v21 = v16 + v20
	v23 = v16 + int32(4)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v24&v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = v21
	goto L5
L4:
	;
	v27 = v23
	goto L5
L5:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v29 <= int32(1) {
		v157 = v28
		v164 = int32(0)
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v170 = F_palloc(m, v167<<(uint(int32(2))%32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L44
	}
L7:
	;
	v33 = int32(1)
	v34 = v28
	goto L8
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v33<<(uint(int32(2))%32))))
	v47 = F_pg_detoast_datum_packed(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v157 = int32(0)
	v164 = v65
	goto L6
L10:
	;
	v49 = int32(1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v51&v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = v49
	goto L13
L12:
	;
	v54 = int32(4)
	goto L13
L13:
	;
	v55 = v47 + v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v57 = int32(1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v61 = v59 & v57
	if v61 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = v57
	goto L16
L15:
	;
	v62 = int32(4)
	goto L16
L16:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v62))))
	v65 = base.B2i32(v56 != v64)
	if v56 != v64 {
		v157 = v34
		v164 = v65
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v66 = v21
	goto L20
L19:
	;
	v66 = v23
	goto L20
L20:
	;
	v67 = int32(2)
	v68 = v66 + v67
	v70 = v55 + v67
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v34 < v71 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v73 = v34
	goto L23
L22:
	;
	v73 = v71
	goto L23
L23:
	;
	v74 = int32(0)
	v79 = int32(8)
	v80 = base.I32_div_s(v73, v79)
	if v79 <= v73 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	if v149 != 0 {
		goto L40
	} else {
		goto L41
	}
L25:
	;
	v149 = v146 + v142<<(uint(int32(3))%32)
	goto L24
L26:
	;
	v128 = v119
	goto L37
L27:
	;
	v86 = v74
	goto L30
L28:
	;
	v103 = v74
	goto L29
L29:
	;
	v110 = v73 - v80<<(uint(int32(3))%32)
	if v110 == int32(0) {
		v142 = v103
		v146 = v74
		goto L25
	} else {
		goto L36
	}
L30:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v86))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v86))))
	if v92 != v94 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v103 = v80
	goto L29
L32:
	;
	v119 = int32(7)
	v120 = v86
	v122 = v92
	v123 = v94
	goto L26
L33:
	;
	goto L34
L34:
	;
	v98 = v86 + int32(1)
	if v98 != v80 {
		v86 = v98
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v103))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v103))))
	v119 = v110
	v120 = v103
	v122 = v116
	v123 = v114
	goto L26
L37:
	;
	if int32(base.Ui32(v122^v123)>>(uint(int32(8)-v128)%32)) != 0 {
		v128 = v128 - int32(1)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v142 = v120
	v146 = v128
	goto L25
L39:
	;
	goto L38
L40:
	;
	v151 = v33 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v152 <= v151 {
		v157 = v149
		v164 = v65
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L9
L43:
	;
	v33 = v151
	v34 = v149
	goto L8
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v170
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v176 = F_palloc(m, v173<<(uint(int32(2))%32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v176
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	return int32(0)
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(2)
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v181)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v184 <= v181 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v224)
	v226 = F_cidr_set_masklen_internal(m, v16, v157)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L58
	}
L50:
	;
	v188 = v181
	goto L51
L51:
	;
	v198 = v188 << (uint(int32(2)) % 32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v198+v199)))
	v202 = F_pg_detoast_datum_packed(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L46
L53:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v206 = int32(1)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v208&v206 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v211 = v206
	goto L56
L55:
	;
	v211 = int32(4)
	goto L56
L56:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+v211))))
	*(*int32)(unsafe.Add(mBase, uint32(v204+v198))) = base.B2i32(v213 != int32(2))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v217+v198))) = v202
	v221 = v188 + int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v221 < v222 {
		v188 = v221
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v226
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v231 <= int32(0) {
		goto L46
	} else {
		goto L59
	}
L59:
	;
	v235 = base.I32_div_s(v157, int32(8))
	v246 = int32(0)
	goto L60
L60:
	;
	v256 = v246 << (uint(int32(2)) % 32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256+v257)))
	v260 = F_pg_detoast_datum_packed(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L46
L62:
	;
	v262 = int32(1)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v264&v262 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v267 = v262
	goto L65
L64:
	;
	v267 = int32(4)
	goto L65
L65:
	;
	v268 = v260 + v267
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v269 == int32(2) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v272 = int32(32)
	goto L68
L67:
	;
	v272 = int32(128)
	goto L68
L68:
	;
	if v157 < v272 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v235)+2)))
	v279 = int32(base.Ui32(v275)>>(uint(v235<<(uint(int32(3))%32)-v157+int32(7))%32)) & int32(1)
	goto L71
L70:
	;
	v279 = int32(0)
	goto L71
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	if v157 < v284 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v286 = v279 | int32(2)
	goto L74
L73:
	;
	v286 = v279
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280+v256))) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v288+v256))) = v260
	v292 = v246 + int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v292 < v293 {
		v246 = v292
		goto L60
	} else {
		goto L75
	}
L75:
	;
	goto L61
}
func F_inet_to_cidr(m *base.Module, l0 int32) int32 {
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
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
				return int32(0)
			} else {
				v35 = int32(1)
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
				v39 = v37 & v35
				if v39 != 0 {
					v40 = v35
				} else {
					v40 = int32(4)
				}
				v42 = int32(1)
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v44&v42 != 0 {
					v47 = v42
				} else {
					v47 = int32(4)
				}
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v47))))
				*(*uint8)(unsafe.Add(mBase, uint32(v33+v40))) = uint8(v49)
				v52 = v33 + int32(1)
				v54 = v33 + int32(4)
				if v39 != 0 {
					v55 = v52
				} else {
					v55 = v54
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)) = uint8(v30)
				if v30 == int32(0) {
				} else {
					v64 = int32(base.Ui32((v30+int32(7))&int32(248)) >> (uint(int32(3)) % 32))
					if v64 != 0 {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
						if v65&int32(1) != 0 {
							v68 = v52
						} else {
							v68 = v54
						}
						v71 = int32(1)
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if v75&v71 != 0 {
							v78 = v15 + v71
						} else {
							v78 = v15 + int32(4)
						}
						base.MemoryCopy(m, v68+int32(2), v78+int32(2), v64)
					} else {
					}
					v83 = v30 & int32(7)
					if v83 == int32(0) {
					} else {
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
						if v88&int32(1) != 0 {
							v91 = v52
						} else {
							v91 = v54
						}
						v92 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) + v91
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+2)))
						v96 = v93 & (int32(-256) >> (uint(v83) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(v92)+2)) = uint8(v96)
					}
				}
				v103 = int32(1)
				v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
				if v105&v103 != 0 {
					v108 = v103
				} else {
					v108 = int32(4)
				}
				v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v108))))
				if v110 == int32(2) {
					v113 = int32(40)
				} else {
					v113 = int32(88)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = v113
				m.G0 = v10 + int32(16)
				return v33
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v30
				F_errmsg_internal(m, int32(_a_F_inet_to_cidr_0), v10)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_inet_to_cidr_1), int32(316), int32(_a_F_inet_to_cidr_2))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
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
