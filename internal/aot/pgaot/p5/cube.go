package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_cmp_v0(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v54 int32
	_ = v54
	var v58 float64
	_ = v58
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v65 int32
	_ = v65
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 float64
	_ = v119
	var v121 int32
	_ = v121
	var v125 float64
	_ = v125
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v132 int32
	_ = v132
	var v136 float64
	_ = v136
	var v139 float64
	_ = v139
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v157 int32
	_ = v157
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v202 int32
	_ = v202
	var v203 float64
	_ = v203
	var v210 float64
	_ = v210
	var v212 int32
	_ = v212
	var v216 float64
	_ = v216
	var v222 float64
	_ = v222
	var v228 float64
	_ = v228
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v257 float64
	_ = v257
	var v264 float64
	_ = v264
	var v266 int32
	_ = v266
	var v270 float64
	_ = v270
	var v276 float64
	_ = v276
	var v282 float64
	_ = v282
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v319 float64
	_ = v319
	var v326 float64
	_ = v326
	var v328 int32
	_ = v328
	var v332 float64
	_ = v332
	var v344 float64
	_ = v344
	var v348 float64
	_ = v348
	var v352 int32
	_ = v352
	var v365 int32
	_ = v365
	var v376 int32
	_ = v376
	var v377 float64
	_ = v377
	var v384 float64
	_ = v384
	var v386 int32
	_ = v386
	var v390 float64
	_ = v390
	var v402 float64
	_ = v402
	var v406 float64
	_ = v406
	var v411 int32
	_ = v411
	var v443 int32
	_ = v443
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = int32(2147483647)
	v21 = v19 & v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v24 = v22 & v20
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return int32(-1)
L2:
	;
	return v443
L3:
	;
	v443 = int32(1)
	goto L2
L4:
	;
	v26 = v21
	goto L6
L5:
	;
	v26 = v24
	goto L6
L6:
	;
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = int32(8)
	v44 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(v24) < base.Ui32(v21) {
		goto L44
	} else {
		goto L45
	}
L10:
	;
	v50 = v44 << (uint(int32(3)) % 32)
	v51 = l0 + v27 + v50
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
	v54 = base.B2i32(v19 < int32(0))
	if v19 < int32(0) {
		v61 = v52
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v93 = int32(8)
	v111 = int32(0)
	goto L27
L12:
	;
	v62 = v50 + (l1 + v27)
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
	v65 = base.B2i32(v22 < int32(0))
	if v22 < int32(0) {
		v72 = v63
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v51+v19<<(uint(int32(3))%32))))
	if base.F64_lt(v52, v58) != 0 {
		v61 = v52
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v61 = v58
	goto L12
L15:
	;
	if base.F64_gt(v61, v72) != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v62+v22<<(uint(int32(3))%32))))
	if base.F64_lt(v63, v69) != 0 {
		v72 = v63
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v72 = v69
	goto L15
L18:
	;
	if v19 < int32(0) {
		v79 = v52
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v22 < int32(0) {
		v86 = v63
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v51+v19<<(uint(int32(3))%32))))
	if base.F64_lt(v52, v77) != 0 {
		v79 = v52
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = v77
	goto L19
L22:
	;
	v88 = int32(-1)
	if base.F64_lt(v79, v86) != 0 {
		v443 = v88
		goto L2
	} else {
		goto L25
	}
L23:
	;
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v62+v22<<(uint(int32(3))%32))))
	if base.F64_lt(v63, v84) != 0 {
		v86 = v63
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v86 = v84
	goto L22
L25:
	;
	v91 = v44 + int32(1)
	if v91 != v26 {
		v44 = v91
		goto L10
	} else {
		goto L26
	}
L26:
	;
	goto L11
L27:
	;
	v117 = v111 << (uint(int32(3)) % 32)
	v118 = l0 + v93 + v117
	v119 = *(*float64)(unsafe.Add(mBase, uint32(v118)))
	v121 = base.B2i32(v19 < int32(0))
	if v19 < int32(0) {
		v128 = v119
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L9
L29:
	;
	v129 = v117 + (l1 + v93)
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v129)))
	v132 = base.B2i32(v22 < int32(0))
	if v22 < int32(0) {
		v139 = v130
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v118+v19<<(uint(int32(3))%32))))
	if base.F64_gt(v119, v125) != 0 {
		v128 = v119
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v128 = v125
	goto L29
L32:
	;
	if base.F64_gt(v128, v139) != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	v136 = *(*float64)(unsafe.Add(mBase, uint32(v129+v22<<(uint(int32(3))%32))))
	if base.F64_gt(v130, v136) != 0 {
		v139 = v130
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v139 = v136
	goto L32
L35:
	;
	if v19 < int32(0) {
		v146 = v119
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v22 < int32(0) {
		v153 = v130
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v118+v19<<(uint(int32(3))%32))))
	if base.F64_gt(v119, v144) != 0 {
		v146 = v119
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v146 = v144
	goto L36
L39:
	;
	if base.F64_lt(v146, v153) != 0 {
		v443 = v88
		goto L2
	} else {
		goto L42
	}
L40:
	;
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v129+v22<<(uint(int32(3))%32))))
	if base.F64_gt(v130, v151) != 0 {
		v153 = v130
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v153 = v151
	goto L39
L42:
	;
	v157 = v111 + int32(1)
	if v157 != v26 {
		v111 = v157
		goto L27
	} else {
		goto L43
	}
L43:
	;
	goto L28
L44:
	;
	v179 = l0 + int32(8)
	v182 = v26
	goto L47
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v24) <= base.Ui32(v21) {
		goto L77
	} else {
		goto L78
	}
L47:
	;
	v202 = v179 + v182<<(uint(int32(3))%32)
	v203 = *(*float64)(unsafe.Add(mBase, uint32(v202)))
	if base.B2i32(v19 < int32(0)) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v245 = v26
	goto L61
L49:
	;
	if base.F64_lt(v228, float64(0)) != 0 {
		goto L1
	} else {
		goto L59
	}
L50:
	;
	v210 = *(*float64)(unsafe.Add(mBase, uint32(v202+v21<<(uint(int32(3))%32))))
	if base.F64_lt(v203, v210) != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if base.F64_gt(v203, float64(0)) != 0 {
		goto L3
	} else {
		goto L58
	}
L53:
	;
	v212 = int32(0)
	goto L55
L54:
	;
	v212 = v19
	goto L55
L55:
	;
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v202+v212<<(uint(int32(3))%32))))
	if base.F64_gt(v216, float64(0)) != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v222 = *(*float64)(unsafe.Add(mBase, uint32(v202+v19<<(uint(int32(3))%32))))
	if base.F64_lt(v203, v222) == int32(0) {
		v228 = v222
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v228 = v203
	goto L49
L58:
	;
	v228 = v203
	goto L49
L59:
	;
	v232 = v182 + int32(1)
	if v232 != v21 {
		v182 = v232
		goto L47
	} else {
		goto L60
	}
L60:
	;
	goto L48
L61:
	;
	v256 = v179 + v245<<(uint(int32(3))%32)
	v257 = *(*float64)(unsafe.Add(mBase, uint32(v256)))
	if base.B2i32(v19 < int32(0)) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L1
L63:
	;
	if base.F64_lt(v282, float64(0)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L64:
	;
	v264 = *(*float64)(unsafe.Add(mBase, uint32(v256+v21<<(uint(int32(3))%32))))
	if base.F64_gt(v257, v264) != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if base.F64_gt(v257, float64(0)) != 0 {
		goto L3
	} else {
		goto L72
	}
L67:
	;
	v266 = int32(0)
	goto L69
L68:
	;
	v266 = v19
	goto L69
L69:
	;
	v270 = *(*float64)(unsafe.Add(mBase, uint32(v256+v266<<(uint(int32(3))%32))))
	if base.F64_gt(v270, float64(0)) != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	v276 = *(*float64)(unsafe.Add(mBase, uint32(v256+v19<<(uint(int32(3))%32))))
	if base.F64_gt(v257, v276) == int32(0) {
		v282 = v276
		goto L63
	} else {
		goto L71
	}
L71:
	;
	v282 = v257
	goto L63
L72:
	;
	v282 = v257
	goto L63
L73:
	;
	v287 = int32(1)
	v289 = v245 + v287
	if v289 == v21 {
		v443 = v287
		goto L2
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L62
L76:
	;
	v245 = v289
	goto L61
L77:
	;
	return int32(0)
L78:
	;
	goto L79
L79:
	;
	v295 = l1 + int32(8)
	v305 = v21
	goto L80
L80:
	;
	v318 = v295 + v305<<(uint(int32(3))%32)
	v319 = *(*float64)(unsafe.Add(mBase, uint32(v318)))
	if base.B2i32(v22 < int32(0)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v365 = v26
	goto L95
L82:
	;
	if base.F64_lt(v348, float64(0)) != 0 {
		goto L3
	} else {
		goto L93
	}
L83:
	;
	v344 = *(*float64)(unsafe.Add(mBase, uint32(v318+v22<<(uint(int32(3))%32))))
	if base.F64_lt(v319, v344) == int32(0) {
		v348 = v344
		goto L82
	} else {
		goto L92
	}
L84:
	;
	v326 = *(*float64)(unsafe.Add(mBase, uint32(v318+v24<<(uint(int32(3))%32))))
	if base.F64_lt(v319, v326) != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	if base.F64_gt(v319, float64(0)) == int32(0) {
		v348 = v319
		goto L82
	} else {
		goto L91
	}
L87:
	;
	v328 = int32(0)
	goto L89
L88:
	;
	v328 = v22
	goto L89
L89:
	;
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v318+v328<<(uint(int32(3))%32))))
	if base.F64_gt(v332, float64(0)) == int32(0) {
		goto L83
	} else {
		goto L90
	}
L90:
	;
	goto L1
L91:
	;
	goto L1
L92:
	;
	v348 = v319
	goto L82
L93:
	;
	v352 = v305 + int32(1)
	if v352 != v24 {
		v305 = v352
		goto L80
	} else {
		goto L94
	}
L94:
	;
	goto L81
L95:
	;
	v376 = v295 + v365<<(uint(int32(3))%32)
	v377 = *(*float64)(unsafe.Add(mBase, uint32(v376)))
	if base.B2i32(v22 < int32(0)) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v443 = int32(-1)
	goto L2
L97:
	;
	if base.F64_lt(v406, float64(0)) != 0 {
		goto L3
	} else {
		goto L108
	}
L98:
	;
	v402 = *(*float64)(unsafe.Add(mBase, uint32(v376+v22<<(uint(int32(3))%32))))
	if base.F64_gt(v377, v402) == int32(0) {
		v406 = v402
		goto L97
	} else {
		goto L107
	}
L99:
	;
	v384 = *(*float64)(unsafe.Add(mBase, uint32(v376+v24<<(uint(int32(3))%32))))
	if base.F64_gt(v377, v384) != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	if base.F64_gt(v377, float64(0)) == int32(0) {
		v406 = v377
		goto L97
	} else {
		goto L106
	}
L102:
	;
	v386 = int32(0)
	goto L104
L103:
	;
	v386 = v22
	goto L104
L104:
	;
	v390 = *(*float64)(unsafe.Add(mBase, uint32(v376+v386<<(uint(int32(3))%32))))
	if base.F64_gt(v390, float64(0)) == int32(0) {
		goto L98
	} else {
		goto L105
	}
L105:
	;
	goto L1
L106:
	;
	goto L1
L107:
	;
	v406 = v377
	goto L97
L108:
	;
	v411 = v365 + int32(1)
	if v24 != v411 {
		v365 = v411
		goto L95
	} else {
		goto L109
	}
L109:
	;
	goto L96
}
func F_cube_coord(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v14 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
					F_errmsg(m, int32(_a_F_cube_coord_0), v7)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_cube_coord_1), int32(1615), int32(_a_F_cube_coord_2))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
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
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if base.Ui32(v17<<(uint(int32(1))%32)) < base.Ui32(v14) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
						F_errmsg(m, int32(_a_F_cube_coord_0), v7)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cube_coord_1), int32(1615), int32(_a_F_cube_coord_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
				if v17 < int32(0) {
					v27 = base.I32_rem_u_s(v14-int32(1), v17&int32(2147483647))
					v36 = v10 + v27<<(uint(int32(3))%32) + int32(8)
				} else {
					v36 = v10 + v14<<(uint(int32(3))%32)
				}
				v37 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
				v38 = F_Float8GetDatum(m, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v38
				}
			}
		}
	}
}
func F_cube_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 float64
	_ = v53
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v88 int32
	_ = v88
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v111 int32
	_ = v111
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	var v134 int32
	_ = v134
	var v141 float64
	_ = v141
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v167 float64
	_ = v167
	var v168 int32
	_ = v168
	var v183 int32
	_ = v183
	var v184 float64
	_ = v184
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v194 int32
	_ = v194
	var v205 float64
	_ = v205
	var v207 float64
	_ = v207
	var v210 int32
	_ = v210
	var v218 float64
	_ = v218
	var v219 float64
	_ = v219
	var v221 float64
	_ = v221
	var v223 int32
	_ = v223
	var v230 float64
	_ = v230
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	v2 = float64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			v29 = int32(2147483647)
			v30 = v28 & v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			v33 = v31 & v29
			v34 = base.B2i32(base.Ui32(v30) < base.Ui32(v33))
			if base.Ui32(v30) < base.Ui32(v33) {
				v35 = v26
			} else {
				v35 = v21
			}
			if base.Ui32(v30) < base.Ui32(v33) {
				v36 = v21
			} else {
				v36 = v26
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
			v39 = v37 & int32(2147483647)
			if v39 == int32(0) {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v141 = v2
				v143 = v42
			} else {
				v43 = int32(8)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v53 = v2
				v58 = int32(0)
				for {
					v68 = v58 << (uint(int32(3)) % 32)
					v69 = v35 + v43 + v68
					v70 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
					v71 = v68 + (v36 + v43)
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v71)))
					if int32(0) <= v47 {
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v69+v47<<(uint(int32(3))%32))))
						v79 = v78
					} else {
						v79 = v70
					}
					if int32(0) <= v37 {
						v85 = *(*float64)(unsafe.Add(mBase, uint32(v71+v37<<(uint(int32(3))%32))))
						v86 = v85
					} else {
						v86 = v72
					}
					v88 = int32(0)
					if base.B2i32(base.F64_le(v79, v86) == v88)|base.B2i32(base.F64_le(v70, v72) == v88)|(base.B2i32(base.F64_le(v79, v72) == v88)|base.B2i32(base.F64_ge(v86, v70) == v88)) == v88 {
						if base.F64_gt(v86, v72) != 0 {
							v105 = v72
						} else {
							v105 = v86
						}
						if base.F64_lt(v79, v70) != 0 {
							v107 = v70
						} else {
							v107 = v79
						}
						v130 = base.F64_sub(v105, v107)
					} else {
						v111 = int32(0)
						if base.B2i32(base.F64_gt(v79, v86) == v111)|base.B2i32(base.F64_gt(v70, v72) == v111)|(base.B2i32(base.F64_gt(v79, v72) == v111)|base.B2i32(base.F64_lt(v86, v70) == v111)) != 0 {
							v130 = float64(0)
						} else {
							if base.F64_gt(v79, v70) != 0 {
								v126 = v70
							} else {
								v126 = v79
							}
							if base.F64_lt(v86, v72) != 0 {
								v128 = v72
							} else {
								v128 = v86
							}
							v130 = base.F64_sub(v126, v128)
						}
					}
					v132 = base.F64_add(base.F64_mul(v130, v130), v53)
					v134 = v58 + int32(1)
					if v134 != v39 {
						v53 = v132
						v58 = v134
						continue
					} else {
						break
					}
					break
				}
				v141 = v132
				v143 = v47
			}
			v156 = v143 & int32(2147483647)
			if base.Ui32(v39) < base.Ui32(v156) {
				v167 = v141
				v168 = v39
				for {
					v183 = v35 + int32(8) + v168<<(uint(int32(3))%32)
					v184 = *(*float64)(unsafe.Add(mBase, uint32(v183)))
					if base.B2i32(v143 < int32(0)) == int32(0) {
						v190 = *(*float64)(unsafe.Add(mBase, uint32(v183+v156<<(uint(int32(3))%32))))
						v191 = v190
					} else {
						v191 = v184
					}
					v192 = float64(0)
					v194 = int32(0)
					if base.B2i32(base.F64_le(v184, v192) == v194)|base.B2i32(base.F64_le(v191, v192) == v194) == v194 {
						if base.F64_lt(v191, v184) != 0 {
							v205 = v184
						} else {
							v205 = v191
						}
						v219 = base.F64_sub(float64(0), v205)
					} else {
						v207 = float64(0)
						v210 = int32(0)
						if base.B2i32(base.F64_gt(v184, v207) == v210)|base.B2i32(base.F64_gt(v191, v207) == v210) != 0 {
							v219 = v207
						} else {
							if base.F64_gt(v191, v184) != 0 {
								v218 = v184
							} else {
								v218 = v191
							}
							v219 = v218
						}
					}
					v221 = base.F64_add(base.F64_mul(v219, v219), v167)
					v223 = v168 + int32(1)
					if v223 != v156 {
						v167 = v221
						v168 = v223
						continue
					} else {
						break
					}
					break
				}
				v230 = v221
			} else {
				v230 = v141
			}
			v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(v30) < base.Ui32(v33) {
				if v244 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v248 = m.ExcPending
					if v248 != 0 {
						return int32(0)
					} else {
						v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v26 != v249 {
							F_pfree(m, v26)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								v259 = F_Float8GetDatum(m, base.F64_sqrt(v230))
								mBase = m.M
								v260 = m.ExcPending
								if v260 != 0 {
									return int32(0)
								} else {
									return v259
								}
							}
						} else {
							v259 = F_Float8GetDatum(m, base.F64_sqrt(v230))
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return int32(0)
							} else {
								return v259
							}
						}
					}
				} else {
					v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v26 != v249 {
						F_pfree(m, v26)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							v259 = F_Float8GetDatum(m, base.F64_sqrt(v230))
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return int32(0)
							} else {
								return v259
							}
						}
					} else {
						v259 = F_Float8GetDatum(m, base.F64_sqrt(v230))
						mBase = m.M
						v260 = m.ExcPending
						if v260 != 0 {
							return int32(0)
						} else {
							return v259
						}
					}
				}
			} else {
				if v244 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return int32(0)
					} else {
						v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v26 == v254 {
							v259 = F_Float8GetDatum(m, base.F64_sqrt(v230))
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return int32(0)
							} else {
								return v259
							}
						} else {
							F_pfree(m, v26)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								v259 = F_Float8GetDatum(m, base.F64_sqrt(v230))
								mBase = m.M
								v260 = m.ExcPending
								if v260 != 0 {
									return int32(0)
								} else {
									return v259
								}
							}
						}
					}
				} else {
					v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v26 == v254 {
						v259 = F_Float8GetDatum(m, base.F64_sqrt(v230))
						mBase = m.M
						v260 = m.ExcPending
						if v260 != 0 {
							return int32(0)
						} else {
							return v259
						}
					} else {
						F_pfree(m, v26)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							v259 = F_Float8GetDatum(m, base.F64_sqrt(v230))
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return int32(0)
							} else {
								return v259
							}
						}
					}
				}
			}
		}
	}
}
func F_cube_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_cube_union_v0(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return v13
							}
						} else {
							return v13
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return v13
						}
					} else {
						return v13
					}
				}
			}
		}
	}
}
func F_cube_ur_coord(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v28 float64
	_ = v28
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = float64(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v12 <= int32(0) {
			v38 = v6
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v17 = v15 & int32(2147483647)
			if base.Ui32(v17) < base.Ui32(v12) {
				v38 = v6
			} else {
				v21 = v8 + v12<<(uint(int32(3))%32)
				if v15 < int32(0) {
					v33 = v21
				} else {
					v24 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
					v28 = *(*float64)(unsafe.Add(mBase, uint32(v21+v17<<(uint(int32(3))%32))))
					if base.F64_gt(v24, v28) != 0 {
						v33 = v21
					} else {
						v33 = v21 + v15<<(uint(int32(3))%32)
					}
				}
				v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)))
				v38 = v34
			}
		}
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v39 != v8 {
			F_pfree(m, v8)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = F_Float8GetDatum(m, v38)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					return v43
				}
			}
		} else {
			v43 = F_Float8GetDatum(m, v38)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				return v43
			}
		}
	}
}
func F_cube_yylex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = l0
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = l1
	v72 = v68
	v81 = int32(0)
	goto L22
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v68 = v16
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(1)
	goto L7
L6:
	;
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v24 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_cube_yylex[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v28
	goto L10
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_cube_yylex[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v34
	goto L13
L12:
	;
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v65)
	v68 = v60
	goto L1
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(int32(2))%32))))
	if v41 != 0 {
		v57 = v41
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_cube_yyensure_buffer_stack(m, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	return int32(0)
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v49 = F_cube_yy_create_buffer(m, v47, int32(_a_F_cube_yylex_0), l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(2))%32)))) = v49
	v57 = v49
	goto L14
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
	v86 = v70
	v87 = v84
	v88 = v72
	v89 = v72
	v97 = v81
	goto L24
L24:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_cube_yylex[2]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v87))%64)&int64(101125980167) == int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v87
	goto L28
L27:
	;
	goto L28
L28:
	;
	v111 = int32(1)
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87<<(uint(v111)%32))+uint32(_c_F_cube_yylex[3]))))
	v116 = v115 + v101
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116<<(uint(v111)%32))+uint32(_c_F_cube_yylex[4]))))
	if v121 != v87 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v125 = v87
	goto L32
L30:
	;
	v154 = v116
	goto L31
L31:
	;
	v167 = int32(1)
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v154<<(uint(v167)%32))+uint32(_c_F_cube_yylex[5]))))
	if v173 != int32(36) {
		v87 = v173
		v89 = v89 + v167
		goto L24
	} else {
		goto L35
	}
L32:
	;
	v136 = int32(1)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125<<(uint(v136)%32))+uint32(_c_F_cube_yylex[6]))))
	v141 = base.I32_extend16_s(v140)
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v141<<(uint(v136)%32))+uint32(_c_F_cube_yylex[3]))))
	v147 = v146 + v101
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147<<(uint(v136)%32))+uint32(_c_F_cube_yylex[4]))))
	if v140 != v152 {
		v125 = v141
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v154 = v147
	goto L31
L34:
	;
	goto L33
L35:
	;
	v181 = v88
	goto L36
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v86)+64))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v86)+68))
	v193 = v189
	v196 = v181
	v197 = v190
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+80)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v197 - v196
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)) = uint8(v207)
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v209)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v197
	v216 = int32(*(*int16)(unsafe.Add(mBase, uint32(v193<<(uint(int32(1))%32))+uint32(_c_F_cube_yylex[7]))))
	v217 = v216
	v220 = v197
	goto L40
L40:
	;
	if v217 != int32(12) {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v1174
	*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = int32(0)
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	v1191 = base.I32_div_s(v1187-int32(1), int32(2))
	v217 = v1191 + int32(13)
	v220 = v1174
	goto L40
L43:
	;
	F_yy_fatal_error_6(m, int32(_a_F_cube_yylex_1))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L19
	} else {
		goto L229
	}
L44:
	;
	F_yy_fatal_error_6(m, int32(_a_F_cube_yylex_2))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L19
	} else {
		goto L228
	}
L45:
	;
	F_yy_fatal_error_6(m, int32(_a_F_cube_yylex_3))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L19
	} else {
		goto L227
	}
L46:
	;
	F_yy_fatal_error_6(m, int32(_a_F_cube_yylex_4))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L19
	} else {
		goto L226
	}
L47:
	;
	return v1155
L48:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1151))) = v1152
	v1155 = int32(258)
	goto L47
L49:
	;
	F_yy_fatal_error_6(m, int32(_a_F_cube_yylex_5))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L19
	} else {
		goto L225
	}
L50:
	;
	switch v217 {
	case 0:
		goto L60
	case 1, 2, 3:
		goto L48
	case 4:
		goto L59
	case 5:
		goto L58
	case 6:
		goto L57
	case 7:
		goto L56
	case 8:
		goto L55
	case 9:
		v70 = v86
		v72 = v220
		v81 = v97
		goto L22
	case 10:
		goto L54
	case 11:
		goto L53
	default:
		goto L49
	case 13:
		v1155 = v97
		goto L47
	}
L51:
	;
	goto L52
L52:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v266)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v268+v269<<(uint(int32(2))%32))))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+44))
	if v274 != 0 {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	F_yy_fatal_error_6(m, int32(_a_F_cube_yylex_6))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L19
	} else {
		goto L61
	}
L54:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v260 = int32(*(*int8)(unsafe.Add(mBase, uint32(v259))))
	return v260
L55:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = int32(_a_F_cube_yylex_7)
	return int32(263)
L56:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = int32(_a_F_cube_yylex_8)
	return int32(260)
L57:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = int32(_a_F_cube_yylex_9)
	return int32(259)
L58:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = int32(_a_F_cube_yylex_8)
	return int32(262)
L59:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(_a_F_cube_yylex_9)
	return int32(261)
L60:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v232)
	v181 = v196
	goto L36
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v288 = v287 + v284
	if base.Ui32(v286) <= base.Ui32(v288) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v284 = v275
	v285 = v274
	goto L62
L64:
	;
	goto L65
L65:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v279 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v273)+44)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v278
	v284 = v276
	v285 = v279
	goto L62
L66:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v293 = v265 ^ int32(-1) + v197
	v294 = v290 + v293
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v294
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if int32(0) < v293 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	if base.Ui32(v288+int32(1)) < base.Ui32(v286) {
		goto L46
	} else {
		goto L98
	}
L69:
	;
	v301 = v296
	v303 = v290
	goto L72
L70:
	;
	v393 = v296
	goto L71
L71:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v393))%64)&int64(101125980167) == int64(0) {
		goto L87
	} else {
		goto L88
	}
L72:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v312 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v393 = v387
	goto L71
L74:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+uint32(_c_F_cube_yylex[2]))))
	v317 = v315
	goto L76
L75:
	;
	v317 = int32(1)
	goto L76
L76:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v301))%64)&int64(101125980167) == int64(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v301
	goto L79
L78:
	;
	goto L79
L79:
	;
	v327 = int32(1)
	v331 = int32(*(*int16)(unsafe.Add(mBase, uint32(v301<<(uint(v327)%32))+uint32(_c_F_cube_yylex[3]))))
	v332 = v317 + v331
	v337 = int32(*(*int16)(unsafe.Add(mBase, uint32(v332<<(uint(v327)%32))+uint32(_c_F_cube_yylex[4]))))
	if v337 != v301 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v341 = v301
	goto L83
L81:
	;
	v370 = v332
	goto L82
L82:
	;
	v383 = int32(1)
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v370<<(uint(v383)%32))+uint32(_c_F_cube_yylex[5]))))
	v389 = v303 + v383
	if v389 != v294 {
		v301 = v387
		v303 = v389
		goto L72
	} else {
		goto L86
	}
L83:
	;
	v352 = int32(1)
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341<<(uint(v352)%32))+uint32(_c_F_cube_yylex[6]))))
	v357 = base.I32_extend16_s(v356)
	v362 = int32(*(*int16)(unsafe.Add(mBase, uint32(v357<<(uint(v352)%32))+uint32(_c_F_cube_yylex[3]))))
	v363 = v317 + v362
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363<<(uint(v352)%32))+uint32(_c_F_cube_yylex[4]))))
	if v356 != v368 {
		v341 = v357
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v370 = v363
	goto L82
L85:
	;
	goto L84
L86:
	;
	goto L73
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v393
	goto L89
L88:
	;
	goto L89
L89:
	;
	v413 = int32(1)
	v417 = int32(*(*int16)(unsafe.Add(mBase, uint32(v393<<(uint(v413)%32))+uint32(_c_F_cube_yylex[3]))))
	v419 = v417 + v413
	v424 = int32(*(*int16)(unsafe.Add(mBase, uint32(v419<<(uint(v413)%32))+uint32(_c_F_cube_yylex[4]))))
	if v424 != v393 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v428 = v393
	goto L93
L91:
	;
	v458 = v419
	goto L92
L92:
	;
	if v458 == int32(0) {
		v181 = v290
		goto L36
	} else {
		goto L96
	}
L93:
	;
	v439 = int32(1)
	v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428<<(uint(v439)%32))+uint32(_c_F_cube_yylex[6]))))
	v444 = base.I32_extend16_s(v443)
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v444<<(uint(v439)%32))+uint32(_c_F_cube_yylex[3]))))
	v451 = v449 + v439
	v456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v451<<(uint(v439)%32))+uint32(_c_F_cube_yylex[4]))))
	if v443 != v456 {
		v428 = v444
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v458 = v451
	goto L92
L95:
	;
	goto L94
L96:
	;
	v477 = int32(*(*int16)(unsafe.Add(mBase, uint32(v458<<(uint(int32(1))%32))+uint32(_c_F_cube_yylex[5]))))
	if v477 == int32(36) {
		v181 = v290
		goto L36
	} else {
		goto L97
	}
L97:
	;
	v481 = v294 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v481
	v87 = v477
	v88 = v290
	v89 = v481
	goto L24
L98:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v273)+40))
	if v487 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v1052 = v1039 + v1043
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v1052
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if base.Ui32(v1052) <= base.Ui32(v1042) {
		v193 = v1054
		v196 = v1042
		v197 = v1052
		goto L38
	} else {
		goto L209
	}
L101:
	;
	if v286-v486 != int32(1) {
		v1039 = v287
		v1042 = v486
		v1043 = v284
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v495 = v486 ^ int32(-1) + v286
	if int32(0) < v495 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v1174 = v486
	goto L42
L105:
	;
	v880 = v866 + v495
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v868)+12))
	if v880 <= v881 {
		goto L180
	} else {
		goto L181
	}
L106:
	;
	if v495 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L107:
	;
	v498 = int32(7)
	v499 = v495 & v498
	if base.Ui32(v286-v486-int32(2)) < base.Ui32(v498) {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	v604 = v273
	v615 = v285
	goto L109
L109:
	;
	if v615 == int32(2) {
		goto L122
	} else {
		goto L123
	}
L110:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v595+v596<<(uint(int32(2))%32))))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+44))
	v604 = v600
	v615 = v601
	goto L109
L111:
	;
	v560 = v546
	v562 = v548
	v567 = int32(0)
	goto L119
L112:
	;
	v546 = v287
	v548 = v486
	goto L111
L113:
	;
	goto L114
L114:
	;
	v508 = v287
	v510 = v486
	v515 = int32(0)
	goto L115
L115:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510))))
	*(*uint8)(unsafe.Add(mBase, uint32(v508))) = uint8(v521)
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)) = uint8(v523)
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+2)) = uint8(v525)
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+3)) = uint8(v527)
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+4)) = uint8(v529)
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+5)) = uint8(v531)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+6)) = uint8(v533)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+7)) = uint8(v535)
	v537 = int32(8)
	v538 = v508 + v537
	v540 = v510 + v537
	v542 = v515 + v537
	if v542 != v495&int32(2147483640) {
		v508 = v538
		v510 = v540
		v515 = v542
		goto L115
	} else {
		goto L117
	}
L116:
	;
	if v499 == int32(0) {
		goto L110
	} else {
		goto L118
	}
L117:
	;
	goto L116
L118:
	;
	v546 = v538
	v548 = v540
	goto L111
L119:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	*(*uint8)(unsafe.Add(mBase, uint32(v560))) = uint8(v573)
	v575 = int32(1)
	v580 = v567 + v575
	if v580 != v499 {
		v560 = v560 + v575
		v562 = v562 + v575
		v567 = v580
		goto L119
	} else {
		goto L121
	}
L120:
	;
	goto L110
L121:
	;
	goto L120
L122:
	;
	v618 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v618
	*(*int32)(unsafe.Add(mBase, uint32(v604)+16)) = v618
	v838 = v604
	goto L106
L123:
	;
	goto L124
L124:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v604)+12))
	v623 = v486 - v286
	v624 = v622 + v623
	if v624 <= int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v628 = v622
	v630 = v604
	v632 = v627
	goto L128
L126:
	;
	v683 = v604
	v688 = v624
	goto L127
L127:
	;
	v694 = int32(16777216)
	if base.Ui32(v694) <= base.Ui32(v688) {
		goto L144
	} else {
		goto L145
	}
L128:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v630)+20))
	if v641 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v683 = v676
	v688 = v678
	goto L127
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+4)) = int32(0)
	goto L43
L131:
	;
	goto L132
L132:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	v647 = int32(0)
	if v628 <= v647 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v656 = v628 - int32(base.Ui32(v647-v628)>>(uint(int32(3))%32))
	goto L135
L134:
	;
	v656 = v628 << (uint(int32(1)) % 32)
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+12)) = v656
	v659 = v656 + int32(2)
	if v646 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+4)) = v664
	if v664 == int32(0) {
		goto L43
	} else {
		goto L142
	}
L137:
	;
	v660 = F_repalloc(m, v646, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L19
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v662 = F_palloc(m, v659)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L19
	} else {
		goto L141
	}
L140:
	;
	v664 = v660
	goto L136
L141:
	;
	v664 = v662
	goto L136
L142:
	;
	v669 = v664 + (v632 - v646)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v669
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v671+v672<<(uint(int32(2))%32))))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v676)+12))
	v678 = v677 + v623
	if v678 <= int32(0) {
		v628 = v677
		v630 = v676
		v632 = v669
		goto L128
	} else {
		goto L143
	}
L143:
	;
	goto L129
L144:
	;
	v697 = v694
	goto L146
L145:
	;
	v697 = v688
	goto L146
L146:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v683)+24))
	if v699 != 0 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v826
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v828+v829<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v833)+16)) = v826
	if v826 != 0 {
		v866 = v826
		v868 = v833
		v879 = int32(0)
		goto L105
	} else {
		goto L174
	}
L148:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v813+v814<<(uint(int32(2))%32))))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v818)+4))
	v822 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v819+v495+v731))) = uint8(v822)
	v826 = v731 + int32(1)
	goto L147
L149:
	;
	v700 = int32(0)
	goto L153
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cube_yylex[8])) = int32(0)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v683)+4))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v750 = F_fread(m, v746+v495, int32(1), v697, v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L19
	} else {
		goto L162
	}
L152:
	;
	switch v717 {
	case 0:
		goto L158
	default:
		v826 = v731
		goto L147
	case 11:
		goto L148
	}
L153:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v714 = F_do_getc(m, v713)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L19
	} else {
		goto L156
	}
L154:
	;
	v731 = v697
	goto L152
L155:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v718+v719<<(uint(int32(2))%32))))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v724+v495+v700))) = uint8(v714)
	v729 = v700 + int32(1)
	if v729 != v697 {
		v700 = v729
		goto L153
	} else {
		goto L157
	}
L156:
	;
	v717 = v714 + int32(1)
	switch v717 {
	case 0, 11:
		v731 = v700
		goto L152
	default:
		goto L155
	}
L157:
	;
	goto L154
L158:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	goto L159
L159:
	;
	if int32(base.Ui32(v733)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v826 = v731
		goto L147
	} else {
		goto L160
	}
L160:
	;
	F_yy_fatal_error_6(m, int32(_a_F_cube_yylex_3))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	v752 = v750
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v752
	if v752 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v805+v806<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v810)+16)) = v752
	v866 = v752
	v868 = v810
	v879 = int32(0)
	goto L105
L165:
	;
	goto L164
L166:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v766)))
	goto L167
L167:
	;
	if int32(base.Ui32(v767)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v774+v775<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v779)+16)) = int32(0)
	v838 = v779
	goto L106
L169:
	;
	goto L170
L170:
	;
	v783 = *(*int32)(unsafe.Add(mBase, _c_F_cube_yylex[8]))
	if v783 != int32(27) {
		goto L45
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cube_yylex[8])) = int32(0)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v766)))
	*(*int32)(unsafe.Add(mBase, uint32(v766))) = v789 & int32(-49)
	goto L172
L172:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v793+v794<<(uint(int32(2))%32))))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v798)+4))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v803 = F_fread(m, v799+v495, int32(1), v697, v802)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L19
	} else {
		goto L173
	}
L173:
	;
	v752 = v803
	goto L163
L174:
	;
	v838 = v833
	goto L106
L175:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	F_cube_yyrestart(m, v851, v86)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L19
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v862 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v838)+44)) = v862
	v866 = int32(0)
	v868 = v838
	v879 = v862
	goto L105
L178:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v854+v855<<(uint(int32(2))%32))))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v866 = v860
	v868 = v859
	v879 = int32(1)
	goto L105
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v910
	v913 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v907+v910))) = uint8(v913)
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v917 = int32(2)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v915+v916<<(uint(v917)%32))))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v920)+4))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v921+v922)+1)) = uint8(v913)
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v926+v927<<(uint(v917)%32))))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v931)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+80)) = v932
	if v879 == int32(1) {
		v1174 = v932
		goto L42
	} else {
		goto L190
	}
L180:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	v907 = v883
	v910 = v880
	goto L179
L181:
	;
	goto L182
L182:
	;
	v886 = v880 + v866>>(uint(int32(1))%32)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	if v887 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v893+v894<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v898)+4)) = v892
	if v892 == int32(0) {
		goto L44
	} else {
		goto L189
	}
L184:
	;
	v888 = F_repalloc(m, v887, v886)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L19
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v890 = F_palloc(m, v886)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L19
	} else {
		goto L188
	}
L187:
	;
	v892 = v888
	goto L183
L188:
	;
	v892 = v890
	goto L183
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v898)+12)) = v886 - int32(2)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v907 = v892
	v910 = v905 + v495
	goto L179
L190:
	;
	switch v879 - int32(1) {
	case 0:
		goto L99
	case 1:
		goto L191
	default:
		goto L192
	}
L191:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v1039 = v932
	v1042 = v932
	v1043 = v1038
	goto L100
L192:
	;
	v940 = v265 ^ int32(-1) + v197
	v941 = v932 + v940
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v941
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if v940 <= int32(0) {
		v87 = v943
		v88 = v932
		v89 = v941
		goto L24
	} else {
		goto L193
	}
L193:
	;
	v948 = v943
	v951 = v932
	goto L194
L194:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951))))
	if v959 != 0 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v87 = v1034
	v88 = v932
	v89 = v941
	goto L24
L196:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959)+uint32(_c_F_cube_yylex[2]))))
	v964 = v962
	goto L198
L197:
	;
	v964 = int32(1)
	goto L198
L198:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v948))%64)&int64(101125980167) == int64(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v948
	goto L201
L200:
	;
	goto L201
L201:
	;
	v974 = int32(1)
	v978 = int32(*(*int16)(unsafe.Add(mBase, uint32(v948<<(uint(v974)%32))+uint32(_c_F_cube_yylex[3]))))
	v979 = v964 + v978
	v984 = int32(*(*int16)(unsafe.Add(mBase, uint32(v979<<(uint(v974)%32))+uint32(_c_F_cube_yylex[4]))))
	if v984 != v948 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v988 = v948
	goto L205
L203:
	;
	v1017 = v979
	goto L204
L204:
	;
	v1030 = int32(1)
	v1034 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1017<<(uint(v1030)%32))+uint32(_c_F_cube_yylex[5]))))
	v1036 = v951 + v1030
	if v941 != v1036 {
		v948 = v1034
		v951 = v1036
		goto L194
	} else {
		goto L208
	}
L205:
	;
	v999 = int32(1)
	v1003 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v988<<(uint(v999)%32))+uint32(_c_F_cube_yylex[6]))))
	v1004 = base.I32_extend16_s(v1003)
	v1009 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1004<<(uint(v999)%32))+uint32(_c_F_cube_yylex[3]))))
	v1010 = v964 + v1009
	v1015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1010<<(uint(v999)%32))+uint32(_c_F_cube_yylex[4]))))
	if v1003 != v1015 {
		v988 = v1004
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v1017 = v1010
	goto L204
L207:
	;
	goto L206
L208:
	;
	goto L195
L209:
	;
	v1058 = v1054
	v1060 = v1042
	goto L210
L210:
	;
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060))))
	if v1069 != 0 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v193 = v1144
	v196 = v1042
	v197 = v1052
	goto L38
L212:
	;
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1069)+uint32(_c_F_cube_yylex[2]))))
	v1074 = v1072
	goto L214
L213:
	;
	v1074 = int32(1)
	goto L214
L214:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v1058))%64)&int64(101125980167) == int64(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v1060
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v1058
	goto L217
L216:
	;
	goto L217
L217:
	;
	v1084 = int32(1)
	v1088 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1058<<(uint(v1084)%32))+uint32(_c_F_cube_yylex[3]))))
	v1089 = v1074 + v1088
	v1094 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1089<<(uint(v1084)%32))+uint32(_c_F_cube_yylex[4]))))
	if v1094 != v1058 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1098 = v1058
	goto L221
L219:
	;
	v1127 = v1089
	goto L220
L220:
	;
	v1140 = int32(1)
	v1144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1127<<(uint(v1140)%32))+uint32(_c_F_cube_yylex[5]))))
	v1146 = v1060 + v1140
	if v1146 != v1052 {
		v1058 = v1144
		v1060 = v1146
		goto L210
	} else {
		goto L224
	}
L221:
	;
	v1109 = int32(1)
	v1113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1098<<(uint(v1109)%32))+uint32(_c_F_cube_yylex[6]))))
	v1114 = base.I32_extend16_s(v1113)
	v1119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1114<<(uint(v1109)%32))+uint32(_c_F_cube_yylex[3]))))
	v1120 = v1074 + v1119
	v1125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1120<<(uint(v1109)%32))+uint32(_c_F_cube_yylex[4]))))
	if v1113 != v1125 {
		v1098 = v1114
		goto L221
	} else {
		goto L223
	}
L222:
	;
	v1127 = v1120
	goto L220
L223:
	;
	goto L222
L224:
	;
	goto L211
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cube_yylex_init_extra(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	if l1 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_cube_yylex_init_extra[0])) = int32(28)
		return int32(1)
	} else {
		v12 = F_palloc(m, int32(96))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
			if v12 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_cube_yylex_init_extra[0])) = int32(48)
				return int32(1)
			} else {
				v24 = int32(0)
				base.MemoryFill(m, v12, v24, int32(96))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v24
				v32 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v29)+52)) = v32
				*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v24
				*(*int64)(unsafe.Add(mBase, uint32(v29)+36)) = v32
				*(*int64)(unsafe.Add(mBase, uint32(v29)+4)) = v32
				*(*int64)(unsafe.Add(mBase, uint32(v29)+12)) = v32
				*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v24
				return v24
			}
		}
	}
}
func F_cube_yyset_extra(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
	return
}
