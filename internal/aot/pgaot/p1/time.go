package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EncodeTimeOnly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
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
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
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
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v7) {
		v22 = F_pg_ultoa_n(m, v7, l5)
		mBase = m.M
		if v8 <= v22 {
			v33 = l5 + v22
		} else {
			v25 = l5 + v8
			v27 = F_memmove(m, v25-v22, l5, v22)
			mBase = m.M
			v30 = F___memset(m, l5, int32(48), v8-v22)
			mBase = m.M
			v33 = v25
		}
	} else {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7<<(uint(int32(1))%32))+uint32(_consts[1077]))))
		*(*uint16)(unsafe.Add(mBase, uint32(l5))) = uint16(v18)
		v33 = l5 + int32(2)
	}
	v34 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v34)
	v37 = v33 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v38) {
		v53 = F_pg_ultoa_n(m, v38, v37)
		mBase = m.M
		if v39 <= v53 {
			v64 = v37 + v53
		} else {
			v56 = v33 + int32(3)
			v58 = F_memmove(m, v56-v53, v37, v53)
			mBase = m.M
			v61 = F___memset(m, v37, int32(48), v39-v53)
			mBase = m.M
			v64 = v56
		}
	} else {
		v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38<<(uint(int32(1))%32))+uint32(_consts[1077]))))
		*(*uint16)(unsafe.Add(mBase, uint32(v37))) = uint16(v49)
		v64 = v33 + int32(3)
	}
	v65 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v65)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = v69 >> (uint(int32(31)) % 32)
	v79 = F_pg_ultostr_zeropad(m, v64+int32(1), v69^v75-v75, int32(2))
	mBase = m.M
	if l1 == int32(0) {
		v187 = v79
	} else {
		v84 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v84)
		v87 = l1 >> (uint(int32(31)) % 32)
		v89 = l1 ^ v87 - v87
		v91 = base.I32_div_s(v89, int32(10))
		v94 = v91*int32(-10) + v89
		if v94 != 0 {
			v96 = v94 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)) = uint8(v96)
			v102 = v79 + int32(7)
		} else {
			v102 = v79 + int32(6)
		}
		v104 = base.I32_div_s(v89, int32(100))
		v107 = v104*int32(-10) + v91
		v108 = v94 | v107
		if v108 == int32(0) {
			v116 = v79 + int32(5)
		} else {
			v114 = v107 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v114)
			v116 = v102
		}
		v118 = base.I32_div_s(v89, int32(1000))
		v121 = v118*int32(-10) + v104
		v122 = v108 | v121
		if v122 == int32(0) {
			v130 = v79 + int32(4)
		} else {
			v128 = v121 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v128)
			v130 = v116
		}
		v132 = base.I32_div_s(v89, int32(10000))
		v135 = v132*int32(-10) + v118
		v136 = v122 | v135
		if v136 == int32(0) {
			v144 = v79 + int32(3)
		} else {
			v142 = v135 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v79)+3)) = uint8(v142)
			v144 = v130
		}
		v146 = base.I32_div_s(v89, int32(100000))
		v149 = v146*int32(-10) + v132
		v150 = v136 | v149
		if v150 == int32(0) {
			v158 = v79 + int32(2)
		} else {
			v156 = v149 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v79)+2)) = uint8(v156)
			v158 = v144
		}
		v160 = base.I32_div_s(v89, int32(1000000))
		v163 = v160*int32(-10) + v146
		if v150|v163 == int32(0) {
			v172 = v79 + int32(1)
		} else {
			v170 = v163 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)) = uint8(v170)
			v172 = v158
		}
		if base.Ui32(int32(19)) <= base.Ui32(v146+int32(9)) {
			v179 = F_pg_ultostr(m, v79+int32(1), v89)
			mBase = m.M
			v180 = v179
		} else {
			v180 = v172
		}
		v187 = v180
	}
	if l2 != 0 {
		if l3 <= int32(0) {
			v195 = int32(43)
		} else {
			v195 = int32(45)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v195)
		v198 = l3 >> (uint(int32(31)) % 32)
		v200 = l3 ^ v198 - v198
		v202 = base.I32_div_s(v200, int32(3600))
		v203 = int32(-60)
		v206 = base.I32_div_s(v200, int32(60))
		v207 = v202*v203 + v206
		v209 = v187 + int32(1)
		v212 = v206*v203 + v200
		if v212 != 0 {
			v213 = int32(2)
			v214 = F_pg_ultostr_zeropad(m, v209, v202, v213)
			mBase = m.M
			v215 = int32(58)
			*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v215)
			v220 = F_pg_ultostr_zeropad(m, v214+int32(1), v207, v213)
			mBase = m.M
			v227 = v220
			v228 = v212
			v229 = int32(58)
			*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v229)
			v234 = F_pg_ultostr_zeropad(m, v227+int32(1), v228, int32(2))
			mBase = m.M
			v235 = v234
		} else {
			v222 = F_pg_ultostr_zeropad(m, v209, v202, int32(2))
			mBase = m.M
			if l4 == int32(4) {
				v227 = v222
				v228 = v207
				v229 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v229)
				v234 = F_pg_ultostr_zeropad(m, v227+int32(1), v228, int32(2))
				mBase = m.M
				v235 = v234
			} else {
				if v207 == int32(0) {
					v235 = v222
				} else {
					v227 = v222
					v228 = v207
					v229 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v229)
					v234 = F_pg_ultostr_zeropad(m, v227+int32(1), v228, int32(2))
					mBase = m.M
					v235 = v234
				}
			}
		}
		v237 = v235
	} else {
		v237 = v187
	}
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v238)
	return
}
func F_time_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	v10 = m.G0
	v12 = v10 - int32(432)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_ParseDateTime(m, v16, v12+int32(240), int32(129), v12+int32(128), v12+int32(16), v12+int32(376))
	mBase = m.M
	if v26 == int32(0) {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+376))
		v44 = F_DecodeTimeOnly(m, v12+int32(128), v12+int32(16), v33, v12+int32(124), v12+int32(384), v12+int32(428), v12+int32(380), v12+int32(8))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			if v44 == int32(0) {
				v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+428)))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+384))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+388))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+392))
				v63 = int32(60)
				v72 = v59 + base.I64_extend_i32_s(v60+(v61+v62*v63)*v63)*int64(1000000)
				if base.Ui32(int32(6)) < base.Ui32(v14) {
					v95 = v72
				} else {
					v76 = v14 << (uint(int32(3)) % 32)
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[1069])))
					v82 = *(*int64)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[1070])))
					if int64(0) <= v72 {
						v85 = v72 + v82
						v86 = base.I64_rem_s(v85, v79)
						v95 = v85 - v86
					} else {
						v88 = v82 - v72
						v89 = base.I64_rem_s(v88, v79)
						v95 = v89 - v88
					}
				}
				v96 = F_Int64GetDatum(m, v95)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					v103 = v96
					m.G0 = v12 + int32(432)
					return v103
				}
			} else {
				v50 = v44
				F_DateTimeParseError(m, v50, v12+int32(8), v16, int32(393595), v15)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v56 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
					v103 = int32(0)
					m.G0 = v12 + int32(432)
					return v103
				}
			}
		}
	} else {
		v50 = v26
		F_DateTimeParseError(m, v50, v12+int32(8), v16, int32(393595), v15)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v56 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
			v103 = int32(0)
			m.G0 = v12 + int32(432)
			return v103
		}
	}
}
func F_time_part_common(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 float64
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v22 = int32(1)
	v23 = v16 + v22
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v28 = v26 & v22
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(32)
	return v330
L4:
	;
	v29 = v23
	goto L6
L5:
	;
	v29 = v16 + int32(4)
	goto L6
L6:
	;
	if v26 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v59 = F_downcase_truncate_identifier(m, v29, v57, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v32 = int32(4)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v34&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v47 = int32(1)
	if v28 != 0 {
		v57 = int32(base.Ui32(v26)>>(uint(v47)%32)) - v47
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v43 = v32
	goto L13
L12:
	;
	v43 = base.B2i32(v34 == int32(18)) << (uint(v32) % 32)
	goto L13
L13:
	;
	if v34 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = v32
	goto L16
L15:
	;
	v46 = v43
	goto L16
L16:
	;
	v57 = v46
	goto L7
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v62 = v13 + int32(28)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[1071]))
	if v69 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v130 == int32(31) {
		goto L38
	} else {
		goto L39
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1071])) = v113
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v113)+11)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v121
	v130 = v120
	goto L19
L21:
	;
	v71 = F_strncmp(m, v59, v69, int32(10))
	mBase = m.M
	if v71 == int32(0) {
		v113 = v69
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59))))
	v80 = int32(1692320)
	v82 = int32(1693280)
	goto L25
L24:
	;
	goto L23
L25:
	;
	v89 = v80 + (v82-v80)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v90 = int32(*(*int8)(unsafe.Add(mBase, uint32(v89))))
	v91 = v74 - v90
	if v91 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(0)
	v130 = int32(31)
	goto L19
L27:
	;
	v95 = F_strncmp(m, v59, v89, int32(10))
	mBase = m.M
	if v95 == int32(0) {
		v113 = v89
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v98 = v91
	goto L29
L29:
	;
	v102 = base.B2i32(v98 < int32(0))
	if v98 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v98 = v95
	goto L29
L31:
	;
	v103 = v89 - int32(16)
	goto L33
L32:
	;
	v103 = v82
	goto L33
L33:
	;
	if v98 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v106 = v80
	goto L36
L35:
	;
	v106 = v89 + int32(16)
	goto L36
L36:
	;
	if base.Ui32(v106) <= base.Ui32(v103) {
		v80 = v106
		v82 = v103
		goto L25
	} else {
		goto L37
	}
L37:
	;
	goto L26
L38:
	;
	v134 = v13 + int32(28)
	v141 = *(*int32)(unsafe.Add(mBase, _consts[1072]))
	if v141 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v203 = v130
	goto L40
L40:
	;
	if v203 == int32(17) {
		goto L62
	} else {
		goto L63
	}
L41:
	;
	v203 = v202
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1072])) = v185
	v192 = int32(*(*int8)(unsafe.Add(mBase, uint32(v185)+11)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v193
	v202 = v192
	goto L41
L43:
	;
	v143 = F_strncmp(m, v59, v141, int32(10))
	mBase = m.M
	if v143 == int32(0) {
		v185 = v141
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v146 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59))))
	v152 = int32(1691168)
	v154 = int32(1692304)
	goto L47
L46:
	;
	goto L45
L47:
	;
	v161 = v152 + (v154-v152)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v162 = int32(*(*int8)(unsafe.Add(mBase, uint32(v161))))
	v163 = v146 - v162
	if v163 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(0)
	v202 = int32(31)
	goto L41
L49:
	;
	v167 = F_strncmp(m, v59, v161, int32(10))
	mBase = m.M
	if v167 == int32(0) {
		v185 = v161
		goto L42
	} else {
		goto L52
	}
L50:
	;
	v170 = v163
	goto L51
L51:
	;
	v174 = base.B2i32(v170 < int32(0))
	if v170 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v170 = v167
	goto L51
L53:
	;
	v175 = v161 - int32(16)
	goto L55
L54:
	;
	v175 = v154
	goto L55
L55:
	;
	if v170 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v178 = v152
	goto L58
L57:
	;
	v178 = v161 + int32(16)
	goto L58
L58:
	;
	if base.Ui32(v178) <= base.Ui32(v175) {
		v152 = v178
		v154 = v175
		goto L47
	} else {
		goto L59
	}
L59:
	;
	goto L48
L60:
	;
	if l1 != 0 {
		goto L97
	} else {
		goto L98
	}
L61:
	;
	v318 = base.I64_extend32_s(v222) + base.I64_extend32_s(v219)*int64(1000000)
	goto L60
L62:
	;
	v207 = base.I64_div_s(v21, int64(3600000000))
	v208 = base.I64_extend32_s(v207)
	v211 = v208*int64(-3600000000) + v21
	v213 = base.I64_div_s(v211, int64(60000000))
	v214 = base.I64_extend32_s(v213)
	v217 = v214*int64(-60000000) + v211
	v219 = base.I64_div_s(v217, int64(1000000))
	v222 = v219*int64(4293967296) + v217
	v223 = base.I32_wrap_i64(v222)
	v224 = base.I32_wrap_i64(v219)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	switch v225 - int32(18) {
	case 0:
		goto L67
	case 1:
		v318 = v214
		goto L60
	case 2:
		goto L66
	default:
		goto L65
	case 11:
		goto L68
	case 12:
		goto L61
	}
L63:
	;
	goto L64
L64:
	;
	if v203 != 0 {
		goto L84
	} else {
		goto L85
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L79
	}
L66:
	;
	v318 = v208
	goto L60
L67:
	;
	if l1 != 0 {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	if l1 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v234 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v222)+base.I64_extend32_s(v219)*int64(1000000), int32(3))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v237 = float64(1000)
	v243 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v224), v237), base.F64_div(base.F64_convert_i32_s(v223), v237)))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	v330 = v234
	goto L3
L73:
	;
	v330 = v243
	goto L3
L74:
	;
	v251 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v222)+base.I64_extend32_s(v219)*int64(1000000), int32(6))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v258 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v223), float64(1e+06)), base.F64_convert_i32_s(v224)))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v330 = v251
	goto L3
L78:
	;
	v330 = v258
	goto L3
L79:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v268 = F_format_type_be(m, int32(1083))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v59
	F_errmsg(m, int32(200330), v13)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(522212), int32(2281), int32(257804))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L92
	}
L85:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v280 != int32(11) {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	if l1 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v284 = F_int64_div_fast_to_numeric(m, v21, int32(6))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v289 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v21), float64(1e+06)))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	v330 = v284
	goto L3
L91:
	;
	v330 = v289
	goto L3
L92:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v299 = F_format_type_be(m, int32(1083))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v59
	F_errmsg(m, int32(200293), v13+int32(16))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(522212), int32(2297), int32(257804))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v319 = F_int64_to_numeric(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v322 = F_Float8GetDatum(m, base.F64_convert_i64_s(v318))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	v330 = v319
	goto L3
L101:
	;
	v330 = v322
	goto L3
}
func F_time_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 == int32(457) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v32 = v2
			v39 = v32
			return v39
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v32 = v2
				v39 = v32
				return v39
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					if v22 < int32(0) {
						v28 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v32 = v28
							v39 = v32
							return v39
						}
					} else {
						if v22 == int32(6) {
							v28 = F_relabel_to_typmod(m, v17, v22)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								v32 = v28
								v39 = v32
								return v39
							}
						} else {
							if base.Ui32(v22) < base.Ui32(v18) {
								v32 = v2
								v39 = v32
								return v39
							} else {
								v28 = F_relabel_to_typmod(m, v17, v22)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int32(0)
								} else {
									v32 = v28
									v39 = v32
									return v39
								}
							}
						}
					}
				}
			}
		}
	} else {
		v39 = int32(0)
		return v39
	}
}
