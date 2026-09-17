package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tidhash_insert_hash_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v98 int64
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	goto L1
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v28) <= base.Ui32(v27) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L4:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v241 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v240 + v241
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+4)) = uint16(v244)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v246
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+6)) = uint8(v241)
	v250 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v250)
	return v231
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L10
	} else {
		goto L51
	}
L6:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v30 == int64(4294967296) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = v42 & l2
	v46 = v41 + v43<<(uint(int32(3))%32)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+6)))
	if v47 == v40 {
		v231 = v46
		goto L4
	} else {
		goto L12
	}
L9:
	;
	F_tidhash_grow(m, l0, v30<<(uint(int64(1))%64))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L8
L12:
	;
	v54 = v46
	v56 = v40
	v57 = v43
	goto L13
L13:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+2)))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
	v65 = int32(16)
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v63|v64<<(uint(v65)%32) == v68|v69<<(uint(v65)%32) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v231 = v212
	goto L4
L15:
	;
	if v79 != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	goto L15
L17:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v75 == v76 {
		v79 = int32(1)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v79 = int32(0)
	goto L16
L20:
	;
	goto L19
L21:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v80)
	return v54
L22:
	;
	goto L23
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v84 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
	v86 = v84 << (uint(int64(32)) % 64)
	v87 = int64(33)
	v89 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v54))))
	v93 = (int64(base.Ui64(v86)>>(uint(v87)%64)) ^ (v89 | v86)) * int64(-49064778989728563)
	v98 = (int64(base.Ui64(v93)>>(uint(v87)%64)) ^ v93) * int64(-4265267296055464877)
	v103 = v83 & base.I32_wrap_i64(int64(base.Ui64(v98)>>(uint(v87)%64))^v98)
	if base.Ui32(v57) < base.Ui32(v103) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v107 = v57 + v105
	goto L26
L25:
	;
	v107 = v57
	goto L26
L26:
	;
	v110 = (v57 + int32(1)) & v83
	if base.Ui32(v107-v103) < base.Ui32(v56) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v116 = v41 + v110<<(uint(int32(3))%32)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
	if v117 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v200 = v56 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v200) {
		goto L46
	} else {
		goto L47
	}
L30:
	;
	v123 = v110
	v124 = int32(0)
	goto L33
L31:
	;
	v154 = v110
	v157 = v116
	goto L32
L32:
	;
	if v154 != v57 {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	if base.Ui32(int32(150)) <= base.Ui32(v124) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v154 = v144
	v157 = v147
	goto L32
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v133), base.F64_convert_i64_u(v135)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v140 = int32(1)
	v144 = (v123 + v140) & v83
	v147 = v41 + v144<<(uint(int32(3))%32)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+6)))
	if v148 != 0 {
		v123 = v144
		v124 = v124 + v140
		goto L33
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	goto L34
L40:
	;
	v168 = v154
	v171 = v157
	goto L43
L41:
	;
	goto L42
L42:
	;
	v231 = v54
	goto L4
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v179 = v176 & (v168 - int32(1))
	v182 = v41 + v179<<(uint(int32(3))%32)
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(v171))) = v183
	if v179 != v57 {
		v168 = v179
		v171 = v182
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v205 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v203), base.F64_convert_i64_u(v205)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v212 = v41 + v110<<(uint(int32(3))%32)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)))
	if v213 != 0 {
		v54 = v212
		v56 = v200
		v57 = v110
		goto L13
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	goto L14
L51:
	;
	F_errmsg_internal(m, int32(_a_F_tidhash_insert_hash_internal_0), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_tidhash_insert_hash_internal_1), int32(630), int32(_a_F_tidhash_insert_hash_internal_2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tidhash_stat(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 float64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 float64
	_ = v152
	var v156 int64
	_ = v156
	var v161 float64
	_ = v161
	var v162 float64
	_ = v162
	var v163 float64
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v2 = int32(0)
	v13 = float64(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = F_palloc0(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if base.B2i32(v25 == int64(0)) == int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v35 = v2
			v39 = v2
			v41 = v2
			for {
				v49 = v31 + v35<<(uint(int32(3))%32)
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+6)))
				if v50 == int32(1) {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v54 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v49)+4)))
					v56 = v54 << (uint(int64(32)) % 64)
					v57 = int64(33)
					v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v49))))
					v63 = (int64(base.Ui64(v56)>>(uint(v57)%64)) ^ (v59 | v56)) * int64(-49064778989728563)
					v68 = (int64(base.Ui64(v63)>>(uint(v57)%64)) ^ v63) * int64(-4265267296055464877)
					v73 = v53 & base.I32_wrap_i64(int64(base.Ui64(v68)>>(uint(v57)%64))^v68)
					v76 = v23 + v73<<(uint(int32(2))%32)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
					*(*int32)(unsafe.Add(mBase, uint32(v76))) = v77 + int32(1)
					if base.Ui32(v35) < base.Ui32(v73) {
						v84 = base.I32_wrap_i64(v25)
					} else {
						v84 = int32(0)
					}
					v85 = v35 - v73 + v84
					if base.Ui32(v41) < base.Ui32(v85) {
						v87 = v85
					} else {
						v87 = v41
					}
					v90 = v85 + v39
					v92 = v87
				} else {
					v90 = v39
					v92 = v41
				}
				v95 = v35 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v95)) < base.Ui64(v25) {
					v35 = v95
					v39 = v90
					v41 = v92
					continue
				} else {
					break
				}
				break
			}
			v98 = int32(0)
			v102 = v98
			v104 = v98
			v106 = v98
			for {
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v23+v106<<(uint(int32(2))%32))))
				v121 = v119 - int32(1)
				if base.Ui32(v104) < base.Ui32(v121) {
					v123 = v121
				} else {
					v123 = v104
				}
				if v119 != 0 {
					v124 = v123
				} else {
					v124 = v104
				}
				v128 = v119 - base.B2i32(v119 != int32(0)) + v102
				v130 = v106 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v130)) < base.Ui64(v25) {
					v102 = v128
					v104 = v124
					v106 = v130
					continue
				} else {
					break
				}
				break
			}
			v134 = v128
			v136 = v124
			v140 = v90
			v142 = v92
		} else {
			v134 = v2
			v136 = v2
			v140 = v2
			v142 = v2
		}
		F_pfree(m, v23)
		mBase = m.M
		v149 = m.ExcPending
		if v149 != 0 {
			return
		} else {
			v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v150 != 0 {
				v152 = base.F64_convert_i32_u(v150)
				v156 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				v161 = base.F64_div(base.F64_convert_i32_u(v134), v152)
				v162 = base.F64_div(base.F64_convert_i32_u(v140), v152)
				v163 = base.F64_div(v152, base.F64_convert_i64_u(v156))
			} else {
				v161 = v13
				v162 = v13
				v163 = float64(0)
			}
			v166 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v167 = m.ExcPending
			if v167 != 0 {
				return
			} else {
				if v166 != 0 {
					v168 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v18)+48)) = v161
					*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v136
					*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v134
					*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v162
					*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v142
					*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v140
					*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v163
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v169
					*(*int64)(unsafe.Add(mBase, uint32(v18))) = v168
					F_errmsg_internal(m, int32(_a_F_tidhash_stat_0), v18)
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_tidhash_stat_1), int32(1144), int32(_a_F_tidhash_stat_2))
						mBase = m.M
						v186 = m.ExcPending
						if v186 != 0 {
							return
						} else {
							m.G0 = v18 - int32(-64)
							return
						}
					}
				} else {
					m.G0 = v18 - int32(-64)
					return
				}
			}
		}
	}
}
