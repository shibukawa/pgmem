package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecChooseHashTableSize(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v25 float64
	_ = v25
	var v27 int32
	_ = v27
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v43 int32
	_ = v43
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 float64
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v117 float64
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v166 float64
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v207 float64
	_ = v207
	var v209 float64
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	v20 = (l1 + int32(7)) & int32(-8)
	v25 = *(*float64)(unsafe.Add(mBase, _c_F_ExecChooseHashTableSize[0]))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ExecChooseHashTableSize[1]))
	v31 = base.F64_mul(base.F64_mul(v25, base.F64_convert_i32_s(v27)), float64(1024))
	v32 = float64(4.294967295e+09)
	if base.F64_lt(v31, v32) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = v31
	goto L3
L2:
	;
	v35 = v32
	goto L3
L3:
	;
	v36 = base.I32_trunc_sat_f64_u(v35)
	if base.F64_le(l0, float64(0)) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = float64(1000)
	goto L6
L5:
	;
	v40 = l0
	goto L6
L6:
	;
	v41 = base.F64_mul(v40, base.F64_convert_i32_s(v20+int32(24)))
	v43 = v20 + int32(68)
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v48 = base.F64_mul(base.F64_convert_i32_s(l4+int32(1)), base.F64_convert_i32_u(v36))
	v49 = float64(4.294967295e+09)
	if base.F64_lt(v48, v49) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v55 = v36
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v55
	if l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v52 = v48
	goto L12
L11:
	;
	v52 = v49
	goto L12
L12:
	;
	v55 = base.I32_trunc_sat_f64_u(v52)
	goto L9
L13:
	;
	v57 = base.I32_div_u_s(v55, v43)
	v58 = int32(50)
	v59 = base.I32_div_u_s(v57, v58)
	if base.Ui32(v58) <= base.Ui32(v57) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v66 = v55
	v67 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v67
	v69 = int32(1)
	v74 = base.F64_ceil(v40)
	v76 = int32(268435455)
	v78 = int32(base.Ui32(v66) >> (uint(int32(2)) % 32))
	if base.Ui32(v76) <= base.Ui32(v78) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v64 = v59 * v43
	goto L18
L17:
	;
	v64 = int32(0)
	goto L18
L18:
	;
	v66 = v55 - v64
	v67 = v59
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v291
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v285
	return
L20:
	;
	v81 = v76
	goto L22
L21:
	;
	v81 = v78
	goto L22
L22:
	;
	v83 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v81)) % 32))
	v84 = base.F64_convert_i32_u(v83)
	if base.F64_gt(v84, v74) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v86 = v74
	goto L25
L24:
	;
	v86 = v84
	goto L25
L25:
	;
	v87 = base.I32_trunc_sat_f64_s(v86)
	if v87 <= int32(1024) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v90 = int32(1024)
	goto L28
L27:
	;
	v90 = v87
	goto L28
L28:
	;
	if v90&(v90-int32(1)) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v97 = v69 << (uint(int32(32)-base.I32_clz(v90)) % 32)
	goto L31
L30:
	;
	v97 = v90
	goto L31
L31:
	;
	if base.F64_lt(base.F64_convert_i32_u(v66), base.F64_add(v41, base.F64_convert_i32_u(v97<<(uint(int32(2))%32)))) == int32(0) {
		v285 = v69
		v291 = v97
		goto L19
	} else {
		goto L32
	}
L32:
	;
	if l3 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v107 = *(*float64)(unsafe.Add(mBase, _c_F_ExecChooseHashTableSize[0]))
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ExecChooseHashTableSize[1]))
	v113 = base.F64_mul(base.F64_mul(v107, base.F64_convert_i32_s(v109)), float64(1024))
	v114 = float64(4.294967295e+09)
	if base.F64_lt(v113, v114) != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v166 = v84
	v167 = v66
	v171 = v83
	goto L35
L35:
	;
	v174 = v20 + int32(28)
	if base.Ui32(v174) < base.Ui32(v167) {
		goto L58
	} else {
		goto L59
	}
L36:
	;
	v117 = v113
	goto L38
L37:
	;
	v117 = v114
	goto L38
L38:
	;
	v118 = base.I32_trunc_sat_f64_u(v117)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v118
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v120 = base.I32_div_u_s(v118, v43)
	v121 = int32(50)
	v122 = base.I32_div_u_s(v120, v121)
	if base.Ui32(v121) <= base.Ui32(v120) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v129 = v118
	v130 = int32(0)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v130
	v137 = int32(268435455)
	v139 = int32(base.Ui32(v129) >> (uint(int32(2)) % 32))
	if base.Ui32(v137) <= base.Ui32(v139) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v127 = v122 * v43
	goto L44
L43:
	;
	v127 = int32(0)
	goto L44
L44:
	;
	v129 = v118 - v127
	v130 = v122
	goto L41
L45:
	;
	v142 = v137
	goto L47
L46:
	;
	v142 = v139
	goto L47
L47:
	;
	v144 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v142)) % 32))
	v145 = base.F64_convert_i32_u(v144)
	if base.F64_gt(v145, v74) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v147 = v74
	goto L50
L49:
	;
	v147 = v145
	goto L50
L50:
	;
	v148 = base.I32_trunc_sat_f64_s(v147)
	if v148 <= int32(1024) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v151 = int32(1024)
	goto L53
L52:
	;
	v151 = v148
	goto L53
L53:
	;
	if v151&(v151-int32(1)) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v158 = int32(1) << (uint(int32(32)-base.I32_clz(v151)) % 32)
	goto L56
L55:
	;
	v158 = v151
	goto L56
L56:
	;
	if base.F64_lt(base.F64_convert_i32_u(v129), base.F64_add(v41, base.F64_convert_i32_u(v158<<(uint(int32(2))%32)))) == int32(0) {
		v285 = v69
		v291 = v158
		goto L19
	} else {
		goto L57
	}
L57:
	;
	v166 = v145
	v167 = v129
	v171 = v144
	goto L35
L58:
	;
	v176 = int32(1)
	v178 = base.I32_div_u_s(v167, v174)
	if v178&(v178-v176) != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v189 = int32(1)
	goto L60
L60:
	;
	v190 = int32(1)
	v191 = int32(32)
	if v189&(v189-v190) != 0 {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	v185 = v176 << (uint(int32(32)-base.I32_clz(v178)) % 32)
	goto L63
L62:
	;
	v185 = v178
	goto L63
L63:
	;
	if base.Ui32(v185) < base.Ui32(v171) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v187 = v185
	goto L66
L65:
	;
	v187 = v171
	goto L66
L66:
	;
	v189 = v187
	goto L60
L67:
	;
	v285 = v266
	v291 = v274
	goto L19
L68:
	;
	v201 = v190 << (uint(v191-base.I32_clz(v189)) % 32)
	goto L70
L69:
	;
	v201 = v189
	goto L70
L70:
	;
	v207 = base.F64_ceil(base.F64_div(v41, base.F64_convert_i32_u(v167-v201<<(uint(int32(2))%32))))
	if base.F64_gt(v166, v207) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v209 = v207
	goto L73
L72:
	;
	v209 = v166
	goto L73
L73:
	;
	v210 = base.I32_trunc_sat_f64_s(v209)
	if v210 <= int32(2) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v213 = int32(2)
	goto L76
L75:
	;
	v213 = v210
	goto L76
L76:
	;
	if v213&(v213-int32(1)) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v220 = v190 << (uint(v191-base.I32_clz(v213)) % 32)
	goto L79
L78:
	;
	v220 = v213
	goto L79
L79:
	;
	if base.B2i32(v220 < int32(2))|base.B2i32(base.Ui32(int32(134217727)) < base.Ui32(v201)) != 0 {
		v266 = v220
		v274 = v201
		goto L67
	} else {
		goto L80
	}
L80:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v228 = v220
	v230 = v226
	v236 = v201
	goto L81
L81:
	;
	if base.B2i32(v230 < int32(0))|base.B2i32(base.Ui32(v228) < base.Ui32(int32(base.Ui32(v230)>>(uint(int32(13))%32)))) != 0 {
		v266 = v228
		v274 = v236
		goto L67
	} else {
		goto L83
	}
L82:
	;
	v285 = v258
	v291 = v260
	goto L19
L83:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v250 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v249 << (uint(v250) % 32)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v255 = v253 << (uint(v250) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v255
	v258 = int32(base.Ui32(v228) >> (uint(v250) % 32))
	v260 = v236 << (uint(v250) % 32)
	if base.Ui32(v228) < base.Ui32(int32(4)) {
		v285 = v258
		v291 = v260
		goto L19
	} else {
		goto L84
	}
L84:
	;
	if base.Ui32(v236) < base.Ui32(int32(67108864)) {
		v228 = v258
		v230 = v255
		v236 = v260
		goto L81
	} else {
		goto L85
	}
L85:
	;
	goto L82
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if l2 == int32(0) {
		if v6 != 0 {
			v71 = v6
			return v71
		} else {
			v14 = l0 + int32(256)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			v17 = F_MemoryContextAlloc(m, v15, int32(_a_F__hash_getcachedmetap_0))
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
							v42 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[0]))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
							v56 = v48
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[1]))
							v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
						}
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						if v57 != 0 {
							v59 = v57
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
							v59 = v22
						}
						base.MemoryCopy(m, v59, v56+int32(24), int32(_a_F__hash_getcachedmetap_0))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_LockBuffer(m, v64, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v71 = v68
							return v71
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
									v42 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[0]))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
									v56 = v48
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[1]))
									v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
								}
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								if v57 != 0 {
									v59 = v57
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
									v59 = v22
								}
								base.MemoryCopy(m, v59, v56+int32(24), int32(_a_F__hash_getcachedmetap_0))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								F_LockBuffer(m, v64, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v71 = v68
									return v71
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
						v42 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[0]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
						v56 = v48
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[1]))
						v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
					}
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
					if v57 != 0 {
						v59 = v57
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
						v59 = v22
					}
					base.MemoryCopy(m, v59, v56+int32(24), int32(_a_F__hash_getcachedmetap_0))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_LockBuffer(m, v64, int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						v71 = v68
						return v71
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
								v42 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[0]))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
								v56 = v48
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[1]))
								v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
							}
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							if v57 != 0 {
								v59 = v57
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
								v59 = v22
							}
							base.MemoryCopy(m, v59, v56+int32(24), int32(_a_F__hash_getcachedmetap_0))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							F_LockBuffer(m, v64, int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								v71 = v68
								return v71
							}
						}
					}
				}
			}
		} else {
			v14 = v12
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			v17 = F_MemoryContextAlloc(m, v15, int32(_a_F__hash_getcachedmetap_0))
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
							v42 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[0]))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
							v56 = v48
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[1]))
							v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
						}
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						if v57 != 0 {
							v59 = v57
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
							v59 = v22
						}
						base.MemoryCopy(m, v59, v56+int32(24), int32(_a_F__hash_getcachedmetap_0))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_LockBuffer(m, v64, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v71 = v68
							return v71
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
									v42 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[0]))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v38^int32(-1))<<(uint(int32(2))%32))))
									v56 = v48
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getcachedmetap[1]))
									v56 = v50 + v38<<(uint(int32(13))%32) + int32(-8192)
								}
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								if v57 != 0 {
									v59 = v57
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
									v59 = v22
								}
								base.MemoryCopy(m, v59, v56+int32(24), int32(_a_F__hash_getcachedmetap_0))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								F_LockBuffer(m, v64, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v71 = v68
									return v71
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
	var v36 int32
	_ = v36
	if l0 < int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F__hash_initbitmapbuffer[0]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7+(l0^int32(-1))<<(uint(int32(2))%32))))
		v21 = v13
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F__hash_initbitmapbuffer[1]))
		v21 = v15 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if l2 != 0 {
		F_PageInit(m, v21, int32(_a_F__hash_initbitmapbuffer_0), int32(16))
		mBase = m.M
	} else {
	}
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v26 = v21 + v25
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = int64(-36028775544127489)
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = int64(-1)
	if l1 != 0 {
		base.MemoryFill(m, v21+int32(24), int32(255), l1)
	} else {
	}
	v36 = l1 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+12)) = uint16(v36)
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v4 = l3
	if l0 < int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F__hash_initbuf[0]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+(l0^int32(-1))<<(uint(int32(2))%32))))
		v22 = v14
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F__hash_initbuf[1]))
		v22 = v16 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v23 = int32(_a_F__hash_initbuf_0)
	v25 = int32(0)
	if v25|(v22&int32(3)|int32(1)) == v25 {
		v41 = v22 + v23
		v43 = v22 + int32(4)
		if base.Ui32(v43) < base.Ui32(v41) {
			v45 = v41
		} else {
			v45 = v43
		}
		v50 = (v22^int32(-1)+v45)&int32(-4) + int32(4)
		if v50 == int32(0) {
		} else {
			base.MemoryFill(m, v22, int32(0), v50)
		}
	} else {
		base.MemoryFill(m, v22, int32(0), v23)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v22)+10)) = int32(_a_F__hash_initbuf_1)
	v64 = int32(_a_F__hash_initbuf_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+18)) = uint16(v64)
	v70 = int32(_a_F__hash_initbuf_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)) = uint16(v70)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)) = uint16(v70)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)))
	v74 = v22 + v73
	v75 = int32(_a_F__hash_initbuf_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
	*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = l1
	return
}
func F__hash_pageinit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v2 = int32(_a_F__hash_pageinit_0)
	v4 = int32(0)
	if v4|(l0&int32(3)|int32(1)) == v4 {
		v20 = l0 + v2
		v22 = l0 + int32(4)
		if base.Ui32(v22) < base.Ui32(v20) {
			v24 = v20
		} else {
			v24 = v22
		}
		v29 = (l0^int32(-1)+v24)&int32(-4) + int32(4)
		if v29 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v29)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v2)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F__hash_pageinit_1)
	v43 = int32(_a_F__hash_pageinit_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v43)
	v49 = int32(_a_F__hash_pageinit_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v49)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v49)
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
				v26 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[0]))
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
									v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
									v90 = v82
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
											v49 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[3]))
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
											v64 = v55
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[4]))
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
												v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
												v90 = v82
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
								v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
								v90 = v82
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
										v49 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[3]))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
										v64 = v55
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[4]))
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
											v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
											v90 = v82
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
				v26 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[0]))
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
									v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
									v90 = v82
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
											v49 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[3]))
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
											v64 = v55
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[4]))
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
												v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
												v90 = v82
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
								v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
								v90 = v82
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
										v49 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[3]))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
										v64 = v55
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[4]))
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
											v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
											v90 = v82
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
			v26 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[0]))
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
								v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
								v90 = v82
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
										v49 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[3]))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
										v64 = v55
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[4]))
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
											v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
											v90 = v82
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
							v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
							v90 = v82
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
									v49 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[3]))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v45^int32(-1))<<(uint(int32(6))%32))+16))
									v64 = v55
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[4]))
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
										v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[1]))
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v71^int32(-1))<<(uint(int32(2))%32))))
										v90 = v82
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_readnext[2]))
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
func F_hash_bitmap_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
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
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
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
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+78)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+76)) = uint16(v2)
	v21 = F_superuser(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L7
	} else {
		goto L87
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L79
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L75
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L7
	} else {
		goto L71
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L7
	} else {
		goto L67
	}
L7:
	;
	return int32(0)
L8:
	;
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = F_relation_open(m, v14, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L63
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+119)))
	if v29 != int32(105) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	if v32 != int32(405) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+118)))
	if v35 == int32(116) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	if v38 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if base.Ui64(int64(4294967295)) <= base.Ui64(v16) {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v44 = F_RelationGetNumberOfBlocksInFork(m, v26, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if base.Ui64(base.I64_extend_i32_u(v44)) <= base.Ui64(v16) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v51 = F__hash_getbuf(m, v26, int32(0), int32(1), int32(8))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	if v16 != int64(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	if v51 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_hash_bitmap_info[0]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+(v51^int32(-1))<<(uint(int32(2))%32))))
	v70 = v62
	goto L22
L25:
	;
	goto L26
L26:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_hash_bitmap_info[1]))
	v70 = v64 + v51<<(uint(int32(13))%32) + int32(-8192)
	goto L22
L27:
	;
	v141 = base.I32_wrap_i64(v16)
	v142 = F__hash_ovflblkno_to_bitno(m, v70+int32(24), v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L7
	} else {
		goto L47
	}
L28:
	;
	v99 = v2
	goto L37
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	if v73 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L33
	}
L32:
	;
	goto L28
L33:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	F_errmsg(m, int32(_a_F_hash_bitmap_info_0), v12)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(454), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v70+int32(468)+v99<<(uint(int32(2))%32)))))
	if v107 != v16 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L43
	}
L39:
	;
	v110 = v99 + int32(1)
	if v73 != v110 {
		v99 = v110
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L27
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+32)) = uint32(v16)
	F_errmsg(m, int32(_a_F_hash_bitmap_info_0), v12+int32(32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(460), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+46)))
	v145 = int32(base.Ui32(v142) >> (uint(v144) % 32))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	if base.Ui32(v146) <= base.Ui32(v145) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+44)))
	v153 = (v148<<(uint(int32(3))%32) - int32(1)) & v142
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v70+v145<<(uint(int32(2))%32))+468))
	F_UnlockReleaseBuffer(m, v51)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v162 = F__hash_getbuf(m, v26, v157, int32(1), int32(4))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L51
	}
L50:
	;
	v183 = base.I32_div_s(v153, int32(32))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181+v183<<(uint(int32(2))%32))+24))
	F_UnlockReleaseBuffer(m, v162)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L55
	}
L51:
	;
	if v162 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_hash_bitmap_info[0]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167+(v162^int32(-1))<<(uint(int32(2))%32))))
	v181 = v173
	goto L50
L53:
	;
	goto L54
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_hash_bitmap_info[1]))
	v181 = v175 + v162<<(uint(int32(13))%32) + int32(-8192)
	goto L50
L55:
	;
	F_relation_close(m, v26, int32(1))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v196 = F_get_call_result_type(m, l0, int32(0), v12+int32(92))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	if v196 != int32(1) {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v201 = F_BlessTupleDesc(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v201
	v205 = F_Int64GetDatum(m, base.I64_extend_i32_u(v157))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = int32(base.Ui32(v187)>>(uint(v153)%32)) & int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v218 = F_heap_form_tuple(m, v213, v12+int32(80), v12+int32(76))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
	v221 = F_HeapTupleHeaderGetDatum(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	m.G0 = v12 + int32(96)
	return v221
L63:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_hash_bitmap_info_3), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(416), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = int32(_a_F_hash_bitmap_info_4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v250 + int32(4)
	F_errmsg(m, int32(_a_F_hash_bitmap_info_5), v12-int32(-64))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(424), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_errmsg(m, int32(_a_F_hash_bitmap_info_6), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(429), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(_a_F_hash_bitmap_info_7), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(434), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v305 + int32(4)
	F_errmsg(m, int32(_a_F_hash_bitmap_info_8), v12+int32(48))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(440), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v141
	F_errmsg(m, int32(_a_F_hash_bitmap_info_0), v12+int32(16))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(475), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errmsg_internal(m, int32(_a_F_hash_bitmap_info_9), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_hash_bitmap_info_1), int32(493), int32(_a_F_hash_bitmap_info_2))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v3 = int32(0)
	v11 = int32(1073741823)
	if v11 <= l0 {
		v14 = v11
	} else {
		v14 = l0
	}
	v26 = int32(1) << (uint(v3-base.I32_clz(int32(base.Ui32(int32(-1)<<(uint(v3-base.I32_clz(v14-int32(1)))%32)^int32(-1))>>(uint(int32(8))%32)))) % 32)
	v30 = int32(256)
	for {
		if v30 < v26 {
			v30 = v30 << (uint(int32(1)) % 32)
			continue
		} else {
			break
		}
		break
	}
	v42 = (l1+int32(7))&int32(-8) + int32(8)
	v46 = F_mul_size(m, v30, int32(4))
	v49 = m.ExcPending
	if v49 != 0 {
		return int32(0)
	} else {
		v50 = F_add_size(m, int32(432), v46)
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			v53 = F_mul_size(m, v26, int32(1024))
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v55 = F_add_size(m, v50, v53)
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v59 = int32(128)
					for {
						v64 = v59 << (uint(int32(1)) % 32)
						v65 = base.I32_div_u_s(v64, v42)
						if base.Ui32(v65) < base.Ui32(int32(32)) {
							v59 = v64
							continue
						} else {
							break
						}
						break
					}
					v68 = int32(1)
					v70 = base.I32_div_s(l0-v68, v65)
					v73 = F_mul_size(m, v65, v42)
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v75 = F_mul_size(m, v70+v68, v73)
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							v77 = F_add_size(m, v55, v75)
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								return v77
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(207)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_hash_identify[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
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
					v20 = v20 + (v295+int32(9))&int32(_a_F_hash_ltree_0)
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
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v12 = v10 & int32(_a_F_hash_mask_0)
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
			base.MemoryFill(m, l0+int32(24), v25, int32(_a_F_hash_mask_1))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
		} else {
			if base.Ui32(int32(2)) < base.Ui32(v20) {
			} else {
				v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				if base.Ui32(v38) < base.Ui32(int32(25)) {
				} else {
					v44 = int32(base.Ui32(v38+int32(_a_F_hash_mask_2)) >> (uint(int32(2)) % 32))
					if v44&int32(_a_F_hash_mask_3) == int32(0) {
					} else {
						v49 = int32(1)
						v51 = l0 + int32(20)
						v55 = (v44 + v49) & int32(_a_F_hash_mask_3)
						if base.Ui32(int32(3)) <= base.Ui32(v55) {
							v58 = int32(2)
							if base.Ui32(v55) <= base.Ui32(v58) {
								v61 = v58
							} else {
								v61 = v55
							}
							v62 = int32(1)
							v63 = v61 - v62
							v71 = v62
							v72 = int32(0)
							for {
								v79 = v51 + v71<<(uint(int32(2))%32)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								if v80&int32(_a_F_hash_mask_4) != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v79))) = v80 & int32(-98305)
								} else {
								}
								v87 = v79 + int32(4)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
								if v88&int32(_a_F_hash_mask_4) != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v87))) = v88 & int32(-98305)
								} else {
								}
								v94 = int32(2)
								v95 = v71 + v94
								v97 = v72 + v94
								if v97 != v63&int32(-2) {
									v71 = v95
									v72 = v97
									continue
								} else {
									break
								}
								break
							}
							if v63&v62 == int32(0) {
							} else {
								v102 = v95
								v110 = v51 + v102<<(uint(int32(2))%32)
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
								if v111&int32(_a_F_hash_mask_4) == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v110))) = v111 & int32(-98305)
								}
							}
						} else {
							v102 = v49
							v110 = v51 + v102<<(uint(int32(2))%32)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
							if v111&int32(_a_F_hash_mask_4) == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v110))) = v111 & int32(-98305)
							}
						}
					}
				}
			}
		}
		v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)))
		v128 = v126 & int32(_a_F_hash_mask_5)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)) = uint16(v128)
		return
	}
}
func F_hash_page_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_superuser(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if v8 != 0 {
				v11 = F_verify_hash_page(m, v4, int32(0))
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return int32(0)
				} else {
					v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+14)))
					if v13 == int32(0) {
						v17 = F_cstring_to_text(m, int32(_a_F_hash_page_type_0))
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return int32(0)
						} else {
							return v17
						}
					} else {
						v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)))
						v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+v20)+12)))
						v24 = v22 & int32(15)
						v26 = v24 - int32(1)
						if base.Ui32(v24^v26) <= base.Ui32(v26) {
							v30 = F_cstring_to_text(m, int32(_a_F_hash_page_type_0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								return v30
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_ctz(v24)<<(uint(int32(2))%32))+uint32(_c_F_hash_page_type[0])))
							v39 = F_cstring_to_text(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								return v39
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_hash_page_type_1), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_hash_page_type_2), int32(203), int32(_a_F_hash_page_type_3))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
				F_errmsg(m, int32(_a_F_hash_scalar_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_hash_scalar_1), int32(3950), int32(_a_F_hash_scalar_2))
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
