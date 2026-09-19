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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = int32(2)
	if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v7)) == int32(0) {
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7<<(uint(int32(1))%32))+uint32(_c_F_EncodeTimeOnly[0]))))
		*(*uint16)(unsafe.Add(mBase, uint32(l5))) = uint16(v19)
		v34 = l5 + int32(2)
	} else {
		v23 = F_pg_ultoa_n(m, v7, l5)
		mBase = m.M
		if v8 <= v23 {
			v34 = l5 + v23
		} else {
			v26 = l5 + v8
			if v23 != 0 {
				base.MemoryCopy(m, v26-v23, l5, v23)
			} else {
			}
			v29 = v8 - v23
			if v29 != 0 {
				base.MemoryFill(m, l5, int32(48), v29)
			} else {
			}
			v34 = v26
		}
	}
	v35 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v35)
	v38 = v34 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = int32(2)
	if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v39)) == int32(0) {
		v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39<<(uint(int32(1))%32))+uint32(_c_F_EncodeTimeOnly[0]))))
		*(*uint16)(unsafe.Add(mBase, uint32(v38))) = uint16(v51)
		v66 = v34 + int32(3)
	} else {
		v55 = F_pg_ultoa_n(m, v39, v38)
		mBase = m.M
		if v40 <= v55 {
			v66 = v38 + v55
		} else {
			v58 = v34 + int32(3)
			if v55 != 0 {
				base.MemoryCopy(m, v58-v55, v38, v55)
			} else {
			}
			v61 = v40 - v55
			if v61 != 0 {
				base.MemoryFill(m, v38, int32(48), v61)
			} else {
			}
			v66 = v58
		}
	}
	v67 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v67)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = v71 >> (uint(int32(31)) % 32)
	v80 = F_pg_ultostr_zeropad(m, v66+int32(1), v71^v76-v76, int32(2))
	mBase = m.M
	if l1 == int32(0) {
		v187 = v80
	} else {
		v85 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v85)
		v88 = l1 >> (uint(int32(31)) % 32)
		v90 = l1 ^ v88 - v88
		v92 = base.I32_div_s(v90, int32(10))
		v95 = v92*int32(-10) + v90
		if v95 != 0 {
			v97 = v95 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v80)+6)) = uint8(v97)
			v103 = v80 + int32(7)
		} else {
			v103 = v80 + int32(6)
		}
		v105 = base.I32_div_s(v90, int32(100))
		v108 = v105*int32(-10) + v92
		v109 = v95 | v108
		if v109 == int32(0) {
			v117 = v80 + int32(5)
		} else {
			v115 = v108 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v80)+5)) = uint8(v115)
			v117 = v103
		}
		v119 = base.I32_div_s(v90, int32(1000))
		v122 = v105 + v119*int32(-10)
		v123 = v109 | v122
		if v123 == int32(0) {
			v131 = v80 + int32(4)
		} else {
			v129 = v122 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)) = uint8(v129)
			v131 = v117
		}
		v133 = base.I32_div_s(v90, int32(_a_F_EncodeTimeOnly_0))
		v136 = v119 + v133*int32(-10)
		v137 = v123 | v136
		if v137 == int32(0) {
			v145 = v80 + int32(3)
		} else {
			v143 = v136 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v80)+3)) = uint8(v143)
			v145 = v131
		}
		v147 = base.I32_div_s(v90, int32(_a_F_EncodeTimeOnly_1))
		v150 = v133 + v147*int32(-10)
		v151 = v137 | v150
		if v151 == int32(0) {
			v159 = v80 + int32(2)
		} else {
			v157 = v150 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v80)+2)) = uint8(v157)
			v159 = v145
		}
		v161 = base.I32_div_s(v90, int32(_a_F_EncodeTimeOnly_2))
		v164 = v161*int32(-10) + v147
		if v151|v164 == int32(0) {
			v173 = v80 + int32(1)
		} else {
			v171 = v164 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)) = uint8(v171)
			v173 = v159
		}
		if base.Ui32(int32(19)) <= base.Ui32(v147+int32(9)) {
			v180 = F_pg_ultostr(m, v80+int32(1), v90)
			mBase = m.M
			v181 = v180
		} else {
			v181 = v173
		}
		v187 = v181
	}
	if l2 != 0 {
		if l3 <= int32(0) {
			v196 = int32(43)
		} else {
			v196 = int32(45)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v196)
		v199 = l3 >> (uint(int32(31)) % 32)
		v201 = l3 ^ v199 - v199
		v203 = base.I32_div_s(v201, int32(3600))
		v204 = int32(-60)
		v207 = base.I32_div_s(v201, int32(60))
		v208 = v203*v204 + v207
		v210 = v187 + int32(1)
		v213 = v207*v204 + v201
		if v213 != 0 {
			v214 = int32(2)
			v215 = F_pg_ultostr_zeropad(m, v210, v203, v214)
			mBase = m.M
			v216 = int32(58)
			*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v216)
			v221 = F_pg_ultostr_zeropad(m, v215+int32(1), v208, v214)
			mBase = m.M
			v228 = v221
			v229 = v213
			v230 = int32(58)
			*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v230)
			v235 = F_pg_ultostr_zeropad(m, v228+int32(1), v229, int32(2))
			mBase = m.M
			v236 = v235
		} else {
			v223 = F_pg_ultostr_zeropad(m, v210, v203, int32(2))
			mBase = m.M
			if l4 == int32(4) {
				v228 = v223
				v229 = v208
				v230 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v230)
				v235 = F_pg_ultostr_zeropad(m, v228+int32(1), v229, int32(2))
				mBase = m.M
				v236 = v235
			} else {
				if v208 == int32(0) {
					v236 = v223
				} else {
					v228 = v223
					v229 = v208
					v230 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v230)
					v235 = F_pg_ultostr_zeropad(m, v228+int32(1), v229, int32(2))
					mBase = m.M
					v236 = v235
				}
			}
		}
		v238 = v236
	} else {
		v238 = v187
	}
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v238))) = uint8(v239)
	return
}
func F_time_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	v12 = m.G0
	v14 = v12 - int32(432)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = v14 + int32(128)
	v25 = v14 + int32(16)
	v28 = F_ParseDateTime(m, v18, v14+int32(240), int32(129), v23, v25, v14+int32(376))
	mBase = m.M
	if v28 == int32(0) {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+376))
		v42 = F_DecodeTimeOnly(m, v23, v25, v31, v14+int32(124), v14+int32(384), v14+int32(428), v14+int32(380), v14+int32(8))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			if v42 == int32(0) {
				v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+428)))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v14)+384))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+388))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+392))
				v61 = int32(60)
				v70 = v57 + base.I64_extend_i32_s(v58+(v59+v60*v61)*v61)*int64(1000000)
				if base.Ui32(int32(6)) < base.Ui32(v17) {
					v89 = v70
				} else {
					v74 = v17 << (uint(int32(3)) % 32)
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_time_in[0])))
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_time_in[1])))
					if int64(0) <= v70 {
						v79 = v70 + v76
						v80 = base.I64_rem_s(v79, v75)
						v89 = v79 - v80
					} else {
						v82 = v76 - v70
						v83 = base.I64_rem_s(v82, v75)
						v89 = v83 - v82
					}
				}
				v90 = F_Int64GetDatum(m, v89)
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					v97 = v90
					m.G0 = v14 + int32(432)
					return v97
				}
			} else {
				v48 = v42
				F_DateTimeParseError(m, v48, v14+int32(8), v18, int32(_a_F_time_in_0), v16)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
					v97 = int32(0)
					m.G0 = v14 + int32(432)
					return v97
				}
			}
		}
	} else {
		v48 = v28
		F_DateTimeParseError(m, v48, v14+int32(8), v18, int32(_a_F_time_in_0), v16)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v54 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
			v97 = int32(0)
			m.G0 = v14 + int32(432)
			return v97
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 float64
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
		v22 = int32(1)
		v23 = v16 + v22
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
		v28 = v26 & v22
		if v28 != 0 {
			v29 = v23
		} else {
			v29 = v16 + int32(4)
		}
		if v26 == int32(1) {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			if v35 == int32(18) {
				v38 = int32(16)
			} else {
				v38 = int32(0)
			}
			if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v45 = int32(4)
			} else {
				v45 = v38
			}
			v56 = v45
		} else {
			v46 = int32(1)
			if v28 != 0 {
				v56 = int32(base.Ui32(v26)>>(uint(v46)%32)) - v46
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v58 = F_downcase_truncate_identifier(m, v29, v56, int32(0))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			v61 = v13 + int32(28)
			v65 = Fn13826(m, v58, v61, int32(_a_F_time_part_common_0), int32(_a_F_time_part_common_1), int32(_a_F_time_part_common_2))
			mBase = m.M
			if v65 == int32(31) {
				v71 = Fn13826(m, v58, v61, int32(_a_F_time_part_common_3), int32(_a_F_time_part_common_4), int32(_a_F_time_part_common_5))
				mBase = m.M
				v72 = v71
			} else {
				v72 = v65
			}
			if v72 == int32(17) {
				v76 = base.I64_div_s(v21, int64(3600000000))
				v77 = base.I64_extend32_s(v76)
				v80 = v77*int64(-3600000000) + v21
				v82 = base.I64_div_s(v80, int64(60000000))
				v83 = base.I64_extend32_s(v82)
				v86 = v83*int64(-60000000) + v80
				v88 = base.I64_div_s(v86, int64(1000000))
				v91 = v88*int64(4293967296) + v86
				v92 = base.I32_wrap_i64(v91)
				v93 = base.I32_wrap_i64(v88)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
				switch v94 - int32(18) {
				case 0:
					if l1 != 0 {
						v120 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v91)+base.I64_extend32_s(v88)*int64(1000000), int32(6))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							v199 = v120
							m.G0 = v13 + int32(32)
							return v199
						}
					} else {
						v127 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v92), float64(1e+06)), base.F64_convert_i32_s(v93)))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int32(0)
						} else {
							v199 = v127
							m.G0 = v13 + int32(32)
							return v199
						}
					}
				case 1:
					v187 = v83
					if l1 != 0 {
						v188 = F_int64_to_numeric(m, v187)
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int32(0)
						} else {
							v199 = v188
							m.G0 = v13 + int32(32)
							return v199
						}
					} else {
						v191 = F_Float8GetDatum(m, base.F64_convert_i64_s(v187))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return int32(0)
						} else {
							v199 = v191
							m.G0 = v13 + int32(32)
							return v199
						}
					}
				case 2:
					v187 = v77
					if l1 != 0 {
						v188 = F_int64_to_numeric(m, v187)
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int32(0)
						} else {
							v199 = v188
							m.G0 = v13 + int32(32)
							return v199
						}
					} else {
						v191 = F_Float8GetDatum(m, base.F64_convert_i64_s(v187))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return int32(0)
						} else {
							v199 = v191
							m.G0 = v13 + int32(32)
							return v199
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v132 = m.ExcPending
					if v132 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return int32(0)
						} else {
							v137 = F_format_type_be(m, int32(1083))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v137
								*(*int32)(unsafe.Add(mBase, uint32(v13))) = v58
								F_errmsg(m, int32(_a_F_time_part_common_6), v13)
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_time_part_common_7), int32(2281), int32(_a_F_time_part_common_8))
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
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
				case 11:
					if l1 != 0 {
						v103 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v91)+base.I64_extend32_s(v88)*int64(1000000), int32(3))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v199 = v103
							m.G0 = v13 + int32(32)
							return v199
						}
					} else {
						v106 = float64(1000)
						v112 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v93), v106), base.F64_div(base.F64_convert_i32_s(v92), v106)))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							v199 = v112
							m.G0 = v13 + int32(32)
							return v199
						}
					}
				case 12:
					v187 = base.I64_extend32_s(v91) + base.I64_extend32_s(v88)*int64(1000000)
					if l1 != 0 {
						v188 = F_int64_to_numeric(m, v187)
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int32(0)
						} else {
							v199 = v188
							m.G0 = v13 + int32(32)
							return v199
						}
					} else {
						v191 = F_Float8GetDatum(m, base.F64_convert_i64_s(v187))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return int32(0)
						} else {
							v199 = v191
							m.G0 = v13 + int32(32)
							return v199
						}
					}
				}
			} else {
				if v72 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							v168 = F_format_type_be(m, int32(1083))
							mBase = m.M
							v169 = m.ExcPending
							if v169 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v168
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v58
								F_errmsg(m, int32(_a_F_time_part_common_9), v13+int32(16))
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_time_part_common_7), int32(2297), int32(_a_F_time_part_common_8))
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
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
				} else {
					v149 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
					if v149 != int32(11) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								v168 = F_format_type_be(m, int32(1083))
								mBase = m.M
								v169 = m.ExcPending
								if v169 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v168
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v58
									F_errmsg(m, int32(_a_F_time_part_common_9), v13+int32(16))
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_time_part_common_7), int32(2297), int32(_a_F_time_part_common_8))
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
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
					} else {
						if l1 != 0 {
							v153 = F_int64_div_fast_to_numeric(m, v21, int32(6))
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int32(0)
							} else {
								v199 = v153
								m.G0 = v13 + int32(32)
								return v199
							}
						} else {
							v158 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v21), float64(1e+06)))
							mBase = m.M
							v159 = m.ExcPending
							if v159 != 0 {
								return int32(0)
							} else {
								v199 = v158
								m.G0 = v13 + int32(32)
								return v199
							}
						}
					}
				}
			}
		}
	}
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
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
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
			v37 = v2
			v43 = v37
			return v43
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v37 = v2
				v43 = v37
				return v43
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					v23 = int32(0)
					if base.B2i32(base.B2i32(v22 < v23)|base.B2i32(v22 == int32(6)) == v23)&base.B2i32(base.Ui32(v22) < base.Ui32(v18)) != 0 {
						v37 = v2
						v43 = v37
						return v43
					} else {
						v32 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v37 = v32
							v43 = v37
							return v43
						}
					}
				}
			}
		}
	} else {
		v43 = int32(0)
		return v43
	}
}
