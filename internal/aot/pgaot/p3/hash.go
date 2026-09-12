package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecChooseHashTableSize(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v27 float64
	_ = v27
	var v29 int32
	_ = v29
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 float64
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 float64
	_ = v88
	var v92 float64
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v130 float64
	_ = v130
	var v132 int32
	_ = v132
	var v136 float64
	_ = v136
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 float64
	_ = v177
	var v179 float64
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v203 float64
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v244 float64
	_ = v244
	var v246 float64
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	v23 = (l1 + int32(7)) & int32(-8)
	v27 = *(*float64)(unsafe.Add(mBase, _consts[528]))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[523]))
	v33 = base.F64_mul(base.F64_mul(v27, base.F64_convert_i32_s(v29)), float64(1024))
	v34 = float64(4.294967295e+09)
	if base.F64_lt(v33, v34) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if base.F64_le(l0, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	if l3 == int32(0) {
		v65 = v45
		goto L1
	} else {
		goto L9
	}
L3:
	;
	v37 = v33
	goto L5
L4:
	;
	v37 = v34
	goto L5
L5:
	;
	if base.F64_lt(v37, float64(4.294967296e+09))&base.F64_ge(v37, float64(0)) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v43 = base.I32_trunc_f64_u(v37)
	v45 = v43
	goto L2
L7:
	;
	goto L8
L8:
	;
	v45 = int32(0)
	goto L2
L9:
	;
	v52 = base.F64_mul(base.F64_convert_i32_s(l4+int32(1)), base.F64_convert_i32_u(v45))
	v53 = float64(4.294967295e+09)
	if base.F64_lt(v52, v53) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v56 = v52
	goto L12
L11:
	;
	v56 = v53
	goto L12
L12:
	;
	if base.F64_lt(v56, float64(4.294967296e+09))&base.F64_ge(v56, float64(0)) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v62 = base.I32_trunc_f64_u(v56)
	v65 = v62
	goto L1
L14:
	;
	goto L15
L15:
	;
	v65 = int32(0)
	goto L1
L16:
	;
	v67 = float64(1000)
	goto L18
L17:
	;
	v67 = l0
	goto L18
L18:
	;
	v70 = v23 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v65
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v73 = base.I32_div_u_s(v65, v70)
	v74 = int32(50)
	v75 = base.I32_div_u_s(v73, v74)
	if base.Ui32(v74) <= base.Ui32(v73) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v82 = v65
	v83 = int32(0)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v83
	v86 = int32(1)
	v88 = base.F64_mul(v67, base.F64_convert_i32_s(v23+int32(24)))
	v92 = base.F64_ceil(v67)
	v94 = int32(268435455)
	v96 = int32(base.Ui32(v82) >> (uint(int32(2)) % 32))
	if base.Ui32(v94) <= base.Ui32(v96) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v80 = v75 * v70
	goto L24
L23:
	;
	v80 = int32(0)
	goto L24
L24:
	;
	v82 = v65 - v80
	v83 = v75
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v327
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v333
	return
L26:
	;
	if v110 <= int32(1024) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	v99 = v94
	goto L29
L28:
	;
	v99 = v96
	goto L29
L29:
	;
	v101 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v99)) % 32))
	v102 = base.F64_convert_i32_u(v101)
	if base.F64_gt(v102, v92) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v104 = v92
	goto L32
L31:
	;
	v104 = v102
	goto L32
L32:
	;
	if base.F64_lt(base.F64_abs(v104), float64(2.147483648e+09)) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v108 = base.I32_trunc_f64_s(v104)
	v110 = v108
	goto L26
L34:
	;
	goto L35
L35:
	;
	v110 = int32(-2147483648)
	goto L26
L36:
	;
	v113 = int32(1024)
	goto L38
L37:
	;
	v113 = v110
	goto L38
L38:
	;
	if v113&(v113-int32(1)) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v120 = v86 << (uint(int32(32)-base.I32_clz(v113)) % 32)
	goto L41
L40:
	;
	v120 = v113
	goto L41
L41:
	;
	if base.F64_lt(base.F64_convert_i32_u(v82), base.F64_add(v88, base.F64_convert_i32_u(v120<<(uint(int32(2))%32)))) == int32(0) {
		v327 = v120
		v333 = v86
		goto L25
	} else {
		goto L42
	}
L42:
	;
	if l3 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = *(*float64)(unsafe.Add(mBase, _consts[528]))
	v132 = *(*int32)(unsafe.Add(mBase, _consts[523]))
	v136 = base.F64_mul(base.F64_mul(v130, base.F64_convert_i32_s(v132)), float64(1024))
	v137 = float64(4.294967295e+09)
	if base.F64_lt(v136, v137) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v203 = v102
	v204 = v82
	v208 = v101
	goto L45
L45:
	;
	v210 = int32(1)
	v215 = v23 + int32(28)
	if base.Ui32(v215) < base.Ui32(v204) {
		goto L78
	} else {
		goto L79
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v148
	if l2 != 0 {
		goto L53
	} else {
		goto L54
	}
L47:
	;
	v140 = v136
	goto L49
L48:
	;
	v140 = v137
	goto L49
L49:
	;
	if base.F64_lt(v140, float64(4.294967296e+09))&base.F64_ge(v140, float64(0)) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v146 = base.I32_trunc_f64_u(v140)
	v148 = v146
	goto L46
L51:
	;
	goto L52
L52:
	;
	v148 = int32(0)
	goto L46
L53:
	;
	v150 = base.I32_div_u_s(v148, v70)
	v151 = int32(50)
	v152 = base.I32_div_u_s(v150, v151)
	if base.Ui32(v151) <= base.Ui32(v150) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v159 = v148
	v160 = int32(0)
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v160
	v163 = int32(1)
	v169 = int32(268435455)
	v171 = int32(base.Ui32(v159) >> (uint(int32(2)) % 32))
	if base.Ui32(v169) <= base.Ui32(v171) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v157 = v152 * v70
	goto L58
L57:
	;
	v157 = int32(0)
	goto L58
L58:
	;
	v159 = v148 - v157
	v160 = v152
	goto L55
L59:
	;
	if v185 <= int32(1024) {
		goto L69
	} else {
		goto L70
	}
L60:
	;
	v174 = v169
	goto L62
L61:
	;
	v174 = v171
	goto L62
L62:
	;
	v176 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v174)) % 32))
	v177 = base.F64_convert_i32_u(v176)
	if base.F64_gt(v177, v92) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v179 = v92
	goto L65
L64:
	;
	v179 = v177
	goto L65
L65:
	;
	if base.F64_lt(base.F64_abs(v179), float64(2.147483648e+09)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v183 = base.I32_trunc_f64_s(v179)
	v185 = v183
	goto L59
L67:
	;
	goto L68
L68:
	;
	v185 = int32(-2147483648)
	goto L59
L69:
	;
	v188 = int32(1024)
	goto L71
L70:
	;
	v188 = v185
	goto L71
L71:
	;
	if v188&(v188-int32(1)) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v195 = v163 << (uint(int32(32)-base.I32_clz(v188)) % 32)
	goto L74
L73:
	;
	v195 = v188
	goto L74
L74:
	;
	if base.F64_lt(base.F64_convert_i32_u(v159), base.F64_add(v88, base.F64_convert_i32_u(v195<<(uint(int32(2))%32)))) == int32(0) {
		v327 = v195
		v333 = v163
		goto L25
	} else {
		goto L75
	}
L75:
	;
	v203 = v177
	v204 = v159
	v208 = v176
	goto L45
L76:
	;
	v327 = v315
	v333 = v307
	goto L25
L77:
	;
	if v252 <= int32(2) {
		goto L96
	} else {
		goto L97
	}
L78:
	;
	v217 = int32(1)
	v219 = base.I32_div_u_s(v204, v215)
	if v219&(v219-v217) != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v227 = v210
	goto L80
L80:
	;
	if base.Ui32(v227) < base.Ui32(v208) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v226 = v217 << (uint(int32(32)-base.I32_clz(v219)) % 32)
	goto L83
L82:
	;
	v226 = v219
	goto L83
L83:
	;
	v227 = v226
	goto L80
L84:
	;
	v231 = v227
	goto L86
L85:
	;
	v231 = v208
	goto L86
L86:
	;
	if v231&(v231-int32(1)) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v238 = int32(1) << (uint(int32(32)-base.I32_clz(v231)) % 32)
	goto L89
L88:
	;
	v238 = v231
	goto L89
L89:
	;
	v244 = base.F64_ceil(base.F64_div(v88, base.F64_convert_i32_u(v204-v238<<(uint(int32(2))%32))))
	if base.F64_gt(v203, v244) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v246 = v244
	goto L92
L91:
	;
	v246 = v203
	goto L92
L92:
	;
	if base.F64_lt(base.F64_abs(v246), float64(2.147483648e+09)) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v250 = base.I32_trunc_f64_s(v246)
	v252 = v250
	goto L77
L94:
	;
	goto L95
L95:
	;
	v252 = int32(-2147483648)
	goto L77
L96:
	;
	v255 = int32(2)
	goto L98
L97:
	;
	v255 = v252
	goto L98
L98:
	;
	if v255&(v255-int32(1)) != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v262 = v210 << (uint(int32(32)-base.I32_clz(v255)) % 32)
	goto L101
L100:
	;
	v262 = v255
	goto L101
L101:
	;
	if v262 < int32(2) {
		v307 = v262
		v315 = v238
		goto L76
	} else {
		goto L102
	}
L102:
	;
	if base.Ui32(int32(134217727)) < base.Ui32(v238) {
		v307 = v262
		v315 = v238
		goto L76
	} else {
		goto L103
	}
L103:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v269 = v262
	v277 = v238
	v279 = v267
	goto L104
L104:
	;
	if v279 < int32(0) {
		v307 = v269
		v315 = v277
		goto L76
	} else {
		goto L106
	}
L105:
	;
	v327 = v301
	v333 = v299
	goto L25
L106:
	;
	if base.Ui32(v269) < base.Ui32(int32(base.Ui32(v279)>>(uint(int32(13))%32))) {
		v307 = v269
		v315 = v277
		goto L76
	} else {
		goto L107
	}
L107:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v291 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v290 << (uint(v291) % 32)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v296 = v294 << (uint(v291) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v296
	v299 = int32(base.Ui32(v269) >> (uint(v291) % 32))
	v301 = v277 << (uint(v291) % 32)
	if base.Ui32(v269) < base.Ui32(int32(4)) {
		v327 = v301
		v333 = v299
		goto L25
	} else {
		goto L108
	}
L108:
	;
	if base.Ui32(v277) < base.Ui32(int32(67108864)) {
		v269 = v299
		v277 = v301
		v279 = v296
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L105
}
func F__hash_checkqual(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(1)
}
func F__hash_dropscanbuf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 == int32(0) {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v12 == v10 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v21 != 0 {
				F_ReleaseBuffer(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
					return
				}
			} else {
				v24 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
				return
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v12 == v15 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v21 != 0 {
					F_ReleaseBuffer(m, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
						return
					}
				} else {
					v24 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
					return
				}
			} else {
				F_ReleaseBuffer(m, v12)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v21 != 0 {
						F_ReleaseBuffer(m, v21)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v24 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
							return
						}
					} else {
						v24 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
						return
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v3 == v6 {
			v10 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v12 == v10 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v21 != 0 {
					F_ReleaseBuffer(m, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
						return
					}
				} else {
					v24 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
					return
				}
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v12 == v15 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v21 != 0 {
						F_ReleaseBuffer(m, v21)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v24 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
							return
						}
					} else {
						v24 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
						return
					}
				} else {
					F_ReleaseBuffer(m, v12)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v21 != 0 {
							F_ReleaseBuffer(m, v21)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return
							} else {
								v24 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
								return
							}
						} else {
							v24 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
							return
						}
					}
				}
			}
		} else {
			F_ReleaseBuffer(m, v3)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v10 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v12 == v10 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v21 != 0 {
						F_ReleaseBuffer(m, v21)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v24 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
							return
						}
					} else {
						v24 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
						return
					}
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v12 == v15 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v21 != 0 {
							F_ReleaseBuffer(m, v21)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return
							} else {
								v24 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
								return
							}
						} else {
							v24 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
							return
						}
					} else {
						F_ReleaseBuffer(m, v12)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v21 != 0 {
								F_ReleaseBuffer(m, v21)
								mBase = m.M
								v23 = m.ExcPending
								if v23 != 0 {
									return
								} else {
									v24 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
									return
								}
							} else {
								v24 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v24)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
								return
							}
						}
					}
				}
			}
		}
	}
}
func F__hash_getcachedmetap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if l2 == int32(0) {
		if v6 != 0 {
			v72 = v6
			return v72
		} else {
			v14 = l0 + int32(256)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			v17 = F_MemoryContextAlloc(m, v15, int32(4544))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = v14
				v22 = v17
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v23 != 0 {
					F_LockBuffer(m, v23, int32(1))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v38 = v27
						if v38 < int32(0) {
							v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
							v56 = v48
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _consts[10]))
							v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
						}
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						if v57 != 0 {
							v59 = v57
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
							v59 = v22
						}
						v63 = F__emscripten_memcpy_bulkmem(m, v59, v56+int32(24), int32(4544))
						mBase = m.M
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_LockBuffer(m, v65, int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v72 = v69
							return v72
						}
					}
				} else {
					v29 = F_ReadBuffer(m, l0, int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_LockBuffer(m, v29, int32(1))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F__hash_checkpage(m, l0, v29, int32(8))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
								v38 = v29
								if v38 < int32(0) {
									v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
									v56 = v48
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, _consts[10]))
									v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
								}
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								if v57 != 0 {
									v59 = v57
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
									v59 = v22
								}
								v63 = F__emscripten_memcpy_bulkmem(m, v59, v56+int32(24), int32(4544))
								mBase = m.M
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								F_LockBuffer(m, v65, int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v72 = v69
									return v72
								}
							}
						}
					}
				}
			}
		}
	} else {
		v12 = l0 + int32(256)
		if v6 != 0 {
			v21 = v12
			v22 = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v23 != 0 {
				F_LockBuffer(m, v23, int32(1))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v38 = v27
					if v38 < int32(0) {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
						v56 = v48
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
					}
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
					if v57 != 0 {
						v59 = v57
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
						v59 = v22
					}
					v63 = F__emscripten_memcpy_bulkmem(m, v59, v56+int32(24), int32(4544))
					mBase = m.M
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_LockBuffer(m, v65, int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						v72 = v69
						return v72
					}
				}
			} else {
				v29 = F_ReadBuffer(m, l0, int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_LockBuffer(m, v29, int32(1))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F__hash_checkpage(m, l0, v29, int32(8))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
							v38 = v29
							if v38 < int32(0) {
								v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
								v56 = v48
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
							}
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							if v57 != 0 {
								v59 = v57
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
								v59 = v22
							}
							v63 = F__emscripten_memcpy_bulkmem(m, v59, v56+int32(24), int32(4544))
							mBase = m.M
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							F_LockBuffer(m, v65, int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								v72 = v69
								return v72
							}
						}
					}
				}
			}
		} else {
			v14 = v12
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			v17 = F_MemoryContextAlloc(m, v15, int32(4544))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = v14
				v22 = v17
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v23 != 0 {
					F_LockBuffer(m, v23, int32(1))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v38 = v27
						if v38 < int32(0) {
							v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
							v56 = v48
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _consts[10]))
							v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
						}
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						if v57 != 0 {
							v59 = v57
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
							v59 = v22
						}
						v63 = F__emscripten_memcpy_bulkmem(m, v59, v56+int32(24), int32(4544))
						mBase = m.M
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_LockBuffer(m, v65, int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v72 = v69
							return v72
						}
					}
				} else {
					v29 = F_ReadBuffer(m, l0, int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_LockBuffer(m, v29, int32(1))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F__hash_checkpage(m, l0, v29, int32(8))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
								v38 = v29
								if v38 < int32(0) {
									v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
									v56 = v48
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, _consts[10]))
									v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
								}
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								if v57 != 0 {
									v59 = v57
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
									v59 = v22
								}
								v63 = F__emscripten_memcpy_bulkmem(m, v59, v56+int32(24), int32(4544))
								mBase = m.M
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								F_LockBuffer(m, v65, int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v72 = v69
									return v72
								}
							}
						}
					}
				}
			}
		}
	}
}
func F__hash_hashkey2bucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v6 = l0 & l2
	if base.Ui32(v6) <= base.Ui32(l1) {
		v8 = int32(-1)
	} else {
		v8 = l3
	}
	return v8 & v6
}
func F__hash_initbitmapbuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	if l0 < int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7+(l0^int32(-1))<<(uint(int32(2))%32))))
		v21 = v13
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[10]))
		v21 = v15 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if l2 != 0 {
		F_PageInit(m, v21, int32(8192), int32(16))
		mBase = m.M
	} else {
	}
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v26 = v21 + v25
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = int64(-36028775544127489)
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = int64(-1)
	v35 = F__emscripten_memset_bulkmem(m, v21+int32(24), base.I32_extend8_s(int32(255)), l1)
	mBase = m.M
	v37 = l1 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+12)) = uint16(v37)
	return
}
func F__hash_initbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v4 = l3
	if l0 < int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+(l0^int32(-1))<<(uint(int32(2))%32))))
		v22 = v14
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[10]))
		v22 = v16 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if v22&int32(3) != 0 {
	} else {
	}
	v49 = F___memset(m, v22, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v22)+10)) = int32(1572864)
	v55 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+18)) = uint16(v55)
	v61 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)) = uint16(v61)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)) = uint16(v61)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)))
	v65 = v22 + v64
	v66 = int32(65408)
	*(*uint16)(unsafe.Add(mBase, uint32(v65)+14)) = uint16(v66)
	*(*uint16)(unsafe.Add(mBase, uint32(v65)+12)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = l1
	return
}
func F__hash_pageinit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	if l0&int32(3) != 0 {
	} else {
	}
	v28 = F___memset(m, l0, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(1572864)
	v34 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v34)
	v40 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v40)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v40)
	return
}
func F__hash_readnext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v12 != v14 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v12 != v16 {
			F_UnlockReleaseBuffer(m, v12)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
				v26 = *(*int32)(unsafe.Add(mBase, _consts[8]))
				if v26 != 0 {
					F_ProcessInterrupts(m)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if v11 != int32(-1) {
							v31 = int32(1)
							v33 = F__hash_getbuf(m, v9, v11, v31, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
								v71 = v33
								if v71 < int32(0) {
									v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
									v90 = v82
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
									v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
								v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
								return
							}
						} else {
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)))
							if v36 != int32(1) {
								return
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
								if v39 != 0 {
									return
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
									F_LockBuffer(m, v40, int32(1))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										if v45 < int32(0) {
											v49 = *(*int32)(unsafe.Add(mBase, _consts[6]))
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
											v64 = v55
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, _consts[7]))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+v45<<(uint(int32(6))%32)+int32(-64))+16))
											v64 = v63
										}
										v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_PredicateLockPage(m, v9, v64, v65)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											v68 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v68)
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v71 = v70
											if v71 < int32(0) {
												v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
												v90 = v82
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
												v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
											}
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
											v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
											return
										}
									}
								}
							}
						}
					}
				} else {
					if v11 != int32(-1) {
						v31 = int32(1)
						v33 = F__hash_getbuf(m, v9, v11, v31, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
							v71 = v33
							if v71 < int32(0) {
								v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
								v90 = v82
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
							v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
							return
						}
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)))
						if v36 != int32(1) {
							return
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
							if v39 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
								F_LockBuffer(m, v40, int32(1))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v45 < int32(0) {
										v49 = *(*int32)(unsafe.Add(mBase, _consts[6]))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
										v64 = v55
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+v45<<(uint(int32(6))%32)+int32(-64))+16))
										v64 = v63
									}
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									F_PredicateLockPage(m, v9, v64, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v68 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v68)
										v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v71 = v70
										if v71 < int32(0) {
											v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
											v90 = v82
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
											v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
										v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_LockBuffer(m, v12, int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
				v26 = *(*int32)(unsafe.Add(mBase, _consts[8]))
				if v26 != 0 {
					F_ProcessInterrupts(m)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if v11 != int32(-1) {
							v31 = int32(1)
							v33 = F__hash_getbuf(m, v9, v11, v31, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
								v71 = v33
								if v71 < int32(0) {
									v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
									v90 = v82
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
									v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
								v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
								return
							}
						} else {
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)))
							if v36 != int32(1) {
								return
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
								if v39 != 0 {
									return
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
									F_LockBuffer(m, v40, int32(1))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										if v45 < int32(0) {
											v49 = *(*int32)(unsafe.Add(mBase, _consts[6]))
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
											v64 = v55
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, _consts[7]))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+v45<<(uint(int32(6))%32)+int32(-64))+16))
											v64 = v63
										}
										v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_PredicateLockPage(m, v9, v64, v65)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											v68 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v68)
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v71 = v70
											if v71 < int32(0) {
												v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
												v90 = v82
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
												v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
											}
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
											v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
											return
										}
									}
								}
							}
						}
					}
				} else {
					if v11 != int32(-1) {
						v31 = int32(1)
						v33 = F__hash_getbuf(m, v9, v11, v31, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
							v71 = v33
							if v71 < int32(0) {
								v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
								v90 = v82
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
							v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
							return
						}
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)))
						if v36 != int32(1) {
							return
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
							if v39 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
								F_LockBuffer(m, v40, int32(1))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v45 < int32(0) {
										v49 = *(*int32)(unsafe.Add(mBase, _consts[6]))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
										v64 = v55
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+v45<<(uint(int32(6))%32)+int32(-64))+16))
										v64 = v63
									}
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									F_PredicateLockPage(m, v9, v64, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v68 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v68)
										v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v71 = v70
										if v71 < int32(0) {
											v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
											v90 = v82
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
											v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
										v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_LockBuffer(m, v12, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			v26 = *(*int32)(unsafe.Add(mBase, _consts[8]))
			if v26 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					if v11 != int32(-1) {
						v31 = int32(1)
						v33 = F__hash_getbuf(m, v9, v11, v31, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
							v71 = v33
							if v71 < int32(0) {
								v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
								v90 = v82
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
							v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
							return
						}
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)))
						if v36 != int32(1) {
							return
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
							if v39 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
								F_LockBuffer(m, v40, int32(1))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v45 < int32(0) {
										v49 = *(*int32)(unsafe.Add(mBase, _consts[6]))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
										v64 = v55
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+v45<<(uint(int32(6))%32)+int32(-64))+16))
										v64 = v63
									}
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									F_PredicateLockPage(m, v9, v64, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v68 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v68)
										v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v71 = v70
										if v71 < int32(0) {
											v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
											v90 = v82
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
											v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
										v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
										return
									}
								}
							}
						}
					}
				}
			} else {
				if v11 != int32(-1) {
					v31 = int32(1)
					v33 = F__hash_getbuf(m, v9, v11, v31, v31)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
						v71 = v33
						if v71 < int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
							v90 = v82
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
							v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
						v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
						return
					}
				} else {
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)))
					if v36 != int32(1) {
						return
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
						if v39 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
							F_LockBuffer(m, v40, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v45 < int32(0) {
									v49 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
									v64 = v55
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, _consts[7]))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+v45<<(uint(int32(6))%32)+int32(-64))+16))
									v64 = v63
								}
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								F_PredicateLockPage(m, v9, v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									v68 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v68)
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v71 = v70
									if v71 < int32(0) {
										v76 = *(*int32)(unsafe.Add(mBase, _consts[9]))
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
										v90 = v82
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
										v90 = v84 + v71<<(uint(int32(13))%32) + int32(-8192)
									}
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
									v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v90 + v92
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_get_hash_value(m *base.Module, l0 int32, l1 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = m.T0[v4].(func(*base.Module, int32, int32) int32)(m, l1, v3)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_hash_destroy(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_MemoryContextDelete(m, v2)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_hash_estimate_size(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v6 = int32(1)
	v8 = int32(32)
	v12 = int32(1073741823)
	if v12 <= l0 {
		v15 = v12
	} else {
		v15 = l0
	}
	v16 = int32(1)
	if base.Ui32(v15) <= base.Ui32(v16) {
		v29 = v6
	} else {
		v29 = int32(base.Ui32(int32(-1)<<(uint(v8-base.I32_clz(v15-v16))%32)^int32(-1))>>(uint(int32(8))%32)) + v16
	}
	v30 = int32(1)
	if base.Ui32(v29) <= base.Ui32(v30) {
		v37 = v6
	} else {
		v37 = v6 << (uint(v8-base.I32_clz(v29-v30)) % 32)
	}
	v41 = int32(256)
	for {
		if v41 < v37 {
			v41 = v41 << (uint(int32(1)) % 32)
			continue
		} else {
			break
		}
		break
	}
	v52 = (l1+int32(7))&int32(-8) + int32(8)
	v56 = F_mul_size(m, v41, int32(4))
	v59 = m.ExcPending
	if v59 != 0 {
		return int32(0)
	} else {
		v60 = F_add_size(m, int32(432), v56)
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			v63 = F_mul_size(m, v37, int32(1024))
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = F_add_size(m, v60, v63)
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v69 = int32(128)
					for {
						v73 = v69 << (uint(int32(1)) % 32)
						v74 = base.I32_div_u_s(v73, v52)
						if v74 < int32(32) {
							v69 = v73
							continue
						} else {
							break
						}
						break
					}
					v77 = int32(1)
					v79 = base.I32_div_s(l0-v77, v74)
					v82 = F_mul_size(m, v74, v52)
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						v84 = F_mul_size(m, v79+v77, v82)
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							v86 = F_add_size(m, v65, v84)
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								return v86
							}
						}
					}
				}
			}
		}
	}
}
func F_hash_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	if base.Ui32(l0) <= base.Ui32(int32(207)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_consts[75])))
		v12 = v11
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_hash_ltree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
		if v12 == int32(0) {
			v307 = int32(1)
		} else {
			v20 = v8 + int32(8)
			v21 = int32(1)
			v23 = v12
			for {
				v26 = v20 + int32(2)
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20))))
				v33 = v27 - int32(1636608432)
				if v26&int32(3) != 0 {
					if base.Ui32(int32(11)) < base.Ui32(v27) {
						v142 = v26
						v143 = v27
						v144 = v33
						v145 = v33
						v146 = v33
						for {
							v148 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
							v149 = v148 + v145
							v150 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
							v153 = v152 + v146
							v155 = int32(4)
							v157 = v150 + v144 - v153 ^ base.I32_rotl(v153, v155)
							v161 = v149 - v157 ^ base.I32_rotl(v157, int32(6))
							v162 = v153 + v149
							v163 = v157 + v162
							v164 = v161 + v163
							v168 = v162 - v161 ^ base.I32_rotl(v161, int32(8))
							v172 = v163 - v168 ^ base.I32_rotl(v168, int32(16))
							v176 = v164 - v172 ^ base.I32_rotl(v172, int32(19))
							v177 = v168 + v164
							v178 = v172 + v177
							v179 = v176 + v178
							v183 = v177 - v176 ^ base.I32_rotl(v176, v155)
							v184 = int32(12)
							v185 = v142 + v184
							v187 = v143 - v184
							if base.Ui32(int32(11)) < base.Ui32(v187) {
								v142 = v185
								v143 = v187
								v144 = v178
								v145 = v179
								v146 = v183
								continue
							} else {
								break
							}
							break
						}
						v190 = v185
						v191 = v187
						v192 = v178
						v193 = v179
						v194 = v183
					} else {
						v190 = v26
						v191 = v27
						v192 = v33
						v193 = v33
						v194 = v33
					}
					switch v191 - int32(1) {
					case 0:
						v253 = v192
						v254 = v193
						v255 = v194
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 1:
						v246 = v192
						v247 = v193
						v248 = v194
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 2:
						v239 = v192
						v240 = v193
						v241 = v194
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 3:
						v233 = v193
						v234 = v194
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
						v239 = v235<<(uint(int32(24))%32) + v192
						v240 = v233
						v241 = v234
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 4:
						v229 = v193
						v230 = v194
						v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
						v233 = v229 + v231
						v234 = v230
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
						v239 = v235<<(uint(int32(24))%32) + v192
						v240 = v233
						v241 = v234
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 5:
						v223 = v193
						v224 = v194
						v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+5)))
						v229 = v225<<(uint(int32(8))%32) + v223
						v230 = v224
						v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
						v233 = v229 + v231
						v234 = v230
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
						v239 = v235<<(uint(int32(24))%32) + v192
						v240 = v233
						v241 = v234
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 6:
						v217 = v193
						v218 = v194
						v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+6)))
						v223 = v219<<(uint(int32(16))%32) + v217
						v224 = v218
						v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+5)))
						v229 = v225<<(uint(int32(8))%32) + v223
						v230 = v224
						v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
						v233 = v229 + v231
						v234 = v230
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
						v239 = v235<<(uint(int32(24))%32) + v192
						v240 = v233
						v241 = v234
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 7:
						v212 = v194
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+7)))
						v217 = v213<<(uint(int32(24))%32) + v193
						v218 = v212
						v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+6)))
						v223 = v219<<(uint(int32(16))%32) + v217
						v224 = v218
						v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+5)))
						v229 = v225<<(uint(int32(8))%32) + v223
						v230 = v224
						v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
						v233 = v229 + v231
						v234 = v230
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
						v239 = v235<<(uint(int32(24))%32) + v192
						v240 = v233
						v241 = v234
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 8:
						v207 = v194
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+8)))
						v212 = v208<<(uint(int32(8))%32) + v207
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+7)))
						v217 = v213<<(uint(int32(24))%32) + v193
						v218 = v212
						v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+6)))
						v223 = v219<<(uint(int32(16))%32) + v217
						v224 = v218
						v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+5)))
						v229 = v225<<(uint(int32(8))%32) + v223
						v230 = v224
						v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
						v233 = v229 + v231
						v234 = v230
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
						v239 = v235<<(uint(int32(24))%32) + v192
						v240 = v233
						v241 = v234
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 9:
						v202 = v194
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+9)))
						v207 = v203<<(uint(int32(16))%32) + v202
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+8)))
						v212 = v208<<(uint(int32(8))%32) + v207
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+7)))
						v217 = v213<<(uint(int32(24))%32) + v193
						v218 = v212
						v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+6)))
						v223 = v219<<(uint(int32(16))%32) + v217
						v224 = v218
						v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+5)))
						v229 = v225<<(uint(int32(8))%32) + v223
						v230 = v224
						v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
						v233 = v229 + v231
						v234 = v230
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
						v239 = v235<<(uint(int32(24))%32) + v192
						v240 = v233
						v241 = v234
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					case 10:
						v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+10)))
						v202 = v198<<(uint(int32(24))%32) + v194
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+9)))
						v207 = v203<<(uint(int32(16))%32) + v202
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+8)))
						v212 = v208<<(uint(int32(8))%32) + v207
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+7)))
						v217 = v213<<(uint(int32(24))%32) + v193
						v218 = v212
						v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+6)))
						v223 = v219<<(uint(int32(16))%32) + v217
						v224 = v218
						v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+5)))
						v229 = v225<<(uint(int32(8))%32) + v223
						v230 = v224
						v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
						v233 = v229 + v231
						v234 = v230
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
						v239 = v235<<(uint(int32(24))%32) + v192
						v240 = v233
						v241 = v234
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)))
						v246 = v242<<(uint(int32(16))%32) + v239
						v247 = v240
						v248 = v241
						v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
						v253 = v249<<(uint(int32(8))%32) + v246
						v254 = v247
						v255 = v248
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
						v260 = v253 + v256
						v261 = v254
						v262 = v255
					default:
						v260 = v192
						v261 = v193
						v262 = v194
					}
				} else {
					if base.Ui32(v27) < base.Ui32(int32(12)) {
						v88 = v26
						v89 = v27
						v90 = v33
						v91 = v33
						v92 = v33
					} else {
						v40 = v26
						v41 = v27
						v42 = v33
						v43 = v33
						v44 = v33
						for {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							v47 = v46 + v43
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
							v51 = v50 + v44
							v53 = int32(4)
							v55 = v48 + v42 - v51 ^ base.I32_rotl(v51, v53)
							v59 = v47 - v55 ^ base.I32_rotl(v55, int32(6))
							v60 = v51 + v47
							v61 = v55 + v60
							v62 = v59 + v61
							v66 = v60 - v59 ^ base.I32_rotl(v59, int32(8))
							v70 = v61 - v66 ^ base.I32_rotl(v66, int32(16))
							v74 = v62 - v70 ^ base.I32_rotl(v70, int32(19))
							v75 = v66 + v62
							v76 = v70 + v75
							v77 = v74 + v76
							v81 = v75 - v74 ^ base.I32_rotl(v74, v53)
							v82 = int32(12)
							v83 = v40 + v82
							v85 = v41 - v82
							if base.Ui32(int32(11)) < base.Ui32(v85) {
								v40 = v83
								v41 = v85
								v42 = v76
								v43 = v77
								v44 = v81
								continue
							} else {
								break
							}
							break
						}
						v88 = v83
						v89 = v85
						v90 = v76
						v91 = v77
						v92 = v81
					}
					switch v89 - int32(1) {
					case 0:
						v139 = v90
						v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
						v260 = v139 + v140
						v261 = v91
						v262 = v92
					case 1:
						v134 = v90
						v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
						v139 = v135<<(uint(int32(8))%32) + v134
						v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
						v260 = v139 + v140
						v261 = v91
						v262 = v92
					case 2:
						v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+2)))
						v134 = v130<<(uint(int32(16))%32) + v90
						v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
						v139 = v135<<(uint(int32(8))%32) + v134
						v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
						v260 = v139 + v140
						v261 = v91
						v262 = v92
					case 3:
						v127 = v91
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v260 = v128 + v90
						v261 = v127
						v262 = v92
					case 4:
						v124 = v91
						v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
						v127 = v124 + v125
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v260 = v128 + v90
						v261 = v127
						v262 = v92
					case 5:
						v119 = v91
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
						v124 = v120<<(uint(int32(8))%32) + v119
						v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
						v127 = v124 + v125
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v260 = v128 + v90
						v261 = v127
						v262 = v92
					case 6:
						v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+6)))
						v119 = v115<<(uint(int32(16))%32) + v91
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
						v124 = v120<<(uint(int32(8))%32) + v119
						v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
						v127 = v124 + v125
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v260 = v128 + v90
						v261 = v127
						v262 = v92
					case 7:
						v110 = v92
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
						v260 = v111 + v90
						v261 = v113 + v91
						v262 = v110
					case 8:
						v105 = v92
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+8)))
						v110 = v106<<(uint(int32(8))%32) + v105
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
						v260 = v111 + v90
						v261 = v113 + v91
						v262 = v110
					case 9:
						v100 = v92
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+9)))
						v105 = v101<<(uint(int32(16))%32) + v100
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+8)))
						v110 = v106<<(uint(int32(8))%32) + v105
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
						v260 = v111 + v90
						v261 = v113 + v91
						v262 = v110
					case 10:
						v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+10)))
						v100 = v96<<(uint(int32(24))%32) + v92
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+9)))
						v105 = v101<<(uint(int32(16))%32) + v100
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+8)))
						v110 = v106<<(uint(int32(8))%32) + v105
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
						v260 = v111 + v90
						v261 = v113 + v91
						v262 = v110
					default:
						v260 = v90
						v261 = v91
						v262 = v92
					}
				}
				v265 = int32(14)
				v267 = v261 ^ v262 - base.I32_rotl(v261, v265)
				v271 = v267 ^ v260 - base.I32_rotl(v267, int32(11))
				v275 = v271 ^ v261 - base.I32_rotl(v271, int32(25))
				v279 = v275 ^ v267 - base.I32_rotl(v275, int32(16))
				v283 = v279 ^ v271 - base.I32_rotl(v279, int32(4))
				v287 = v283 ^ v275 - base.I32_rotl(v283, v265)
				v294 = v287 ^ v279 - base.I32_rotl(v287, int32(24)) + v21*int32(31)
				v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20))))
				v301 = int32(1)
				if base.Ui32(v301) < base.Ui32(v23) {
					v20 = v20 + (v295+int32(9))&int32(131064)
					v21 = v294
					v23 = v23 - v301
					continue
				} else {
					break
				}
				break
			}
			v307 = v294
		}
		v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v311 != v8 {
			F_pfree(m, v8)
			mBase = m.M
			v314 = m.ExcPending
			if v314 != 0 {
				return int32(0)
			} else {
				return v307
			}
		} else {
			return v307
		}
	}
}
func F_hash_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v12 = v10 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v12)
	F_mask_unused_space(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v17 = l0 + v16
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)))
		v20 = v18 & int32(15)
		if v20 == int32(0) {
			v25 = int32(0)
			v27 = F___memset(m, l0+int32(24), v25, int32(8168))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
		} else {
			if base.Ui32(int32(2)) < base.Ui32(v20) {
			} else {
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				if base.Ui32(v39) < base.Ui32(int32(25)) {
				} else {
					v45 = int32(base.Ui32(v39+int32(262120)) >> (uint(int32(2)) % 32))
					if v45&int32(65535) == int32(0) {
					} else {
						v50 = int32(1)
						v51 = int32(2)
						v55 = (v45 + v50) & int32(65535)
						if base.Ui32(v55) <= base.Ui32(v51) {
							v58 = v51
						} else {
							v58 = v55
						}
						v59 = int32(1)
						v60 = v58 - v59
						v64 = l0 + int32(24)
						v65 = int32(0)
						if base.Ui32(int32(3)) <= base.Ui32(v55) {
							v70 = v65
							v71 = v50
							for {
								v80 = v71<<(uint(int32(2))%32) + v64
								v82 = v80 - int32(4)
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
								if v83&int32(98304) != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v82))) = v83 & int32(-98305)
								} else {
								}
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								if v89&int32(98304) != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v80))) = v89 & int32(-98305)
								} else {
								}
								v95 = int32(2)
								v98 = v70 + v95
								if v98 != v60&int32(-2) {
									v70 = v98
									v71 = v71 + v95
									continue
								} else {
									break
								}
								break
							}
							v102 = v71 + int32(1)
						} else {
							v102 = v65
						}
						if v60&v59 == int32(0) {
						} else {
							v114 = v64 + v102<<(uint(int32(2))%32)
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
							if v115&int32(98304) == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v114))) = v115 & int32(-98305)
							}
						}
					}
				}
			}
		}
		v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)))
		v133 = v131 & int32(65407)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)) = uint16(v133)
		return
	}
}
func F_hash_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	switch v9 {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v19
				F_errmsg(m, int32(225990), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(485799), int32(3950), int32(225823))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
		m.G0 = v6 + int32(16)
		return int32(0)
	default:
		m.G0 = v6 + int32(16)
		return int32(0)
	}
}
