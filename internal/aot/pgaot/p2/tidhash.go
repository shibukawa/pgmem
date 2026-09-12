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
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
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
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int64
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	goto L1
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v28) <= base.Ui32(v27) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L4:
	;
	v273 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v273)
	return v264
L5:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v251 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v250 + v251
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v241)+4)) = uint16(v254)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v256
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+6)) = uint8(v251)
	v264 = v241
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L11
	} else {
		goto L52
	}
L7:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v30 == int64(4294967296) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = v42 & l2
	v46 = v41 + v43<<(uint(int32(3))%32)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+6)))
	if v47 == v40 {
		v241 = v46
		goto L5
	} else {
		goto L13
	}
L10:
	;
	F_tidhash_grow(m, l0, v30<<(uint(int64(1))%64))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L9
L13:
	;
	v54 = v46
	v55 = v40
	v57 = v43
	goto L14
L14:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+2)))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
	v65 = int32(16)
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v63|v64<<(uint(v65)%32) == v68|v69<<(uint(v65)%32) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v241 = v222
	goto L5
L16:
	;
	if v79 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v75 == v76 {
		v79 = int32(1)
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v79 = int32(0)
	goto L17
L21:
	;
	goto L20
L22:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v80)
	return v54
L23:
	;
	goto L24
L24:
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
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v107 = v57 + v105
	goto L27
L26:
	;
	v107 = v57
	goto L27
L27:
	;
	v110 = (v57 + int32(1)) & v83
	if base.Ui32(v107-v103) < base.Ui32(v55) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v116 = v41 + v110<<(uint(int32(3))%32)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
	if v117 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v210 = v55 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v210) {
		goto L47
	} else {
		goto L48
	}
L31:
	;
	v123 = int32(0)
	v124 = v110
	goto L34
L32:
	;
	v155 = v110
	v157 = v116
	goto L33
L33:
	;
	if v155 != v57 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	if base.Ui32(int32(150)) <= base.Ui32(v123) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v155 = v144
	v157 = v147
	goto L33
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v133), base.F64_convert_i64_u(v135)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v140 = int32(1)
	v144 = (v124 + v140) & v83
	v147 = v41 + v144<<(uint(int32(3))%32)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+6)))
	if v148 != 0 {
		v123 = v123 + v140
		v124 = v144
		goto L34
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	goto L35
L41:
	;
	v169 = v155
	v171 = v157
	goto L44
L42:
	;
	goto L43
L43:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v200 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + v200
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)) = uint16(v203)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v205
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+6)) = uint8(v200)
	v264 = v54
	goto L4
L44:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v179 = v176 & (v169 - int32(1))
	v182 = v41 + v179<<(uint(int32(3))%32)
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(v171))) = v183
	if v179 != v57 {
		v169 = v179
		v171 = v182
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	goto L45
L47:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v213), base.F64_convert_i64_u(v215)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v222 = v41 + v110<<(uint(int32(3))%32)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+6)))
	if v223 != 0 {
		v54 = v222
		v55 = v210
		v57 = v110
		goto L14
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	goto L15
L52:
	;
	F_errmsg_internal(m, int32(479873), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(340209), int32(630), int32(324085))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v61 int64
	_ = v61
	var v66 int64
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 float64
	_ = v153
	var v157 int64
	_ = v157
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
		if v25 == int64(0) {
			v133 = v2
			v135 = v2
			v136 = v2
			v137 = v2
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v32 = v2
			v35 = v2
			v36 = v2
			for {
				v47 = v29 + v32<<(uint(int32(3))%32)
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+6)))
				if v48 == int32(1) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v52 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)))
					v54 = v52 << (uint(int64(32)) % 64)
					v55 = int64(33)
					v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v47))))
					v61 = (int64(base.Ui64(v54)>>(uint(v55)%64)) ^ (v57 | v54)) * int64(-49064778989728563)
					v66 = (int64(base.Ui64(v61)>>(uint(v55)%64)) ^ v61) * int64(-4265267296055464877)
					v71 = v51 & base.I32_wrap_i64(int64(base.Ui64(v66)>>(uint(v55)%64))^v66)
					v74 = v23 + v71<<(uint(int32(2))%32)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
					*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75 + int32(1)
					if base.Ui32(v32) < base.Ui32(v71) {
						v82 = base.I32_wrap_i64(v25)
					} else {
						v82 = int32(0)
					}
					v83 = v32 - v71 + v82
					if base.Ui32(v36) < base.Ui32(v83) {
						v85 = v83
					} else {
						v85 = v36
					}
					v88 = v83 + v35
					v89 = v85
				} else {
					v88 = v35
					v89 = v36
				}
				v93 = v32 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v93)) < base.Ui64(v25) {
					v32 = v93
					v35 = v88
					v36 = v89
					continue
				} else {
					break
				}
				break
			}
			v96 = int32(0)
			v101 = v96
			v103 = v96
			v106 = v96
			for {
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v23+v106<<(uint(int32(2))%32))))
				v119 = v117 - int32(1)
				if base.Ui32(v101) < base.Ui32(v119) {
					v121 = v119
				} else {
					v121 = v101
				}
				if v117 != 0 {
					v122 = v121
				} else {
					v122 = v101
				}
				if base.Ui32(v119) <= base.Ui32(v117) {
					v125 = v119
				} else {
					v125 = int32(0)
				}
				v126 = v125 + v103
				v128 = v106 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v128)) < base.Ui64(v25) {
					v101 = v122
					v103 = v126
					v106 = v128
					continue
				} else {
					break
				}
				break
			}
			v133 = v122
			v135 = v126
			v136 = v88
			v137 = v89
		}
		F_pfree(m, v23)
		mBase = m.M
		v147 = m.ExcPending
		if v147 != 0 {
			return
		} else {
			v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v148 == int32(0) {
				v161 = v13
				v162 = v13
				v163 = float64(0)
			} else {
				v153 = base.F64_convert_i32_u(v148)
				v157 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				v161 = base.F64_div(base.F64_convert_i32_u(v135), v153)
				v162 = base.F64_div(base.F64_convert_i32_u(v136), v153)
				v163 = base.F64_div(v153, base.F64_convert_i64_u(v157))
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
					*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v133
					*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v135
					*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v162
					*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v137
					*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v136
					*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v163
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v169
					*(*int64)(unsafe.Add(mBase, uint32(v18))) = v168
					F_errmsg_internal(m, int32(354062), v18)
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return
					} else {
						F_errfinish(m, int32(340209), int32(1144), int32(118608))
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
