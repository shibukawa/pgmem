package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AdjustTimestampForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v62 int32
	_ = v62
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(v14-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v62 = v13
		m.G0 = v11 + int32(16)
		return v62
	} else {
		switch l1 + int32(1) {
		case 0, 7:
			v62 = v13
			m.G0 = v11 + int32(16)
			return v62
		default:
			if base.Ui32(int32(7)) <= base.Ui32(l1) {
				v23 = int32(0)
				v24 = F_errsave_start(m, l2)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						v62 = v23
						m.G0 = v11 + int32(16)
						return v62
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(25769803776)
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
							F_errmsg(m, int32(_a_F_AdjustTimestampForTypmod_0), v11)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, l2, int32(_a_F_AdjustTimestampForTypmod_1), int32(392), int32(_a_F_AdjustTimestampForTypmod_2))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v62 = v23
									m.G0 = v11 + int32(16)
									return v62
								}
							}
						}
					}
				}
			} else {
				v45 = l1 << (uint(int32(3)) % 32)
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_AdjustTimestampForTypmod[0])))
				v47 = *(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_AdjustTimestampForTypmod[1])))
				if int64(0) <= v14 {
					v50 = v14 + v47
					v51 = base.I64_rem_s(v50, v46)
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50 - v51
					v62 = v13
				} else {
					v54 = v47 - v14
					v55 = base.I64_rem_s(v54, v46)
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v55 - v54
					v62 = v13
				}
				m.G0 = v11 + int32(16)
				return v62
			}
		}
	}
}
func F_timestamp_ne_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_date_cmp_timestamp_internal(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(v4 != int32(0)))
	}
}
func F_timestamp_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int64
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v9-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		if v9 != int64(-9223372036854775807-1) {
			v171 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_timestamp_out[0])))
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v171)
			v174 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_out[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v174
		} else {
			v17 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_timestamp_out[2])))
			*(*uint16)(unsafe.Add(mBase, uint32(v7)+8)) = uint16(v17)
			v20 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_out[3]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
		}
		v178 = F_pstrdup(m, v7)
		mBase = m.M
		v179 = m.ExcPending
		if v179 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(176)
			return base.I64_extend_i32_u(v178)
		}
	} else {
		v23 = base.I64_div_s(v9, int64(86400000000))
		if base.Ui64(int64(172799999999)) <= base.Ui64(v9+int64(86399999999)) {
			v31 = v23 * int64(-86400000000)
		} else {
			v31 = int64(0)
		}
		v32 = v31 + v9
		v35 = v32>>(uint(int64(63))%64) + v23
		if int64(-2451545) <= v35 {
			v38 = base.I32_wrap_i64(v35)
			v50 = v38 + int32(_a_F_timestamp_out_0)
			v51 = int32(_a_F_timestamp_out_1)
			v52 = base.I32_div_u_s(v50, v51)
			v53 = int32(3)
			v59 = int32(2)
			v64 = base.I32_div_u_s((v52*int32(1073595727)+v50)<<(uint(v59)%32)|v53, v51)
			v67 = v38 + int32(_a_F_timestamp_out_2) + v52*v53 + v64 + int32(_a_F_timestamp_out_3)
			v68 = int32(1461)
			v69 = base.I32_div_u_s(v67, v68)
			v72 = v69*int32(-1461) + v67
			v74 = v72 << (uint(v59) % 32)
			if base.Ui32(v68) <= base.Ui32(v74) {
				v80 = base.I32_rem_u_s(v72+int32(305), int32(365))
				v85 = v80
			} else {
				v84 = base.I32_rem_u_s(v72+int32(306), int32(366))
				v85 = v84
			}
			v87 = base.I32_div_u_s(v74, int32(1461))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(152)))) = v87 + v69<<(uint(int32(2))%32) - int32(_a_F_timestamp_out_4)
			v95 = v85 + int32(123)
			v99 = int32(base.Ui32(v95*int32(2141)) >> (uint(int32(16)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(144)))) = v95 - int32(base.Ui32(v99*int32(_a_F_timestamp_out_5))>>(uint(int32(8))%32))
			v109 = base.I32_rem_u_s(v99+int32(10), int32(12))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(148)))) = v109 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+164)) = int64(4294967295)
			if v32 < int64(0) {
				v121 = v32 + int64(86400000000)
			} else {
				v121 = v32
			}
			v123 = base.I64_div_s(v121, int64(3600000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v7)+140)) = uint32(v123)
			v128 = base.I64_extend32_s(v123)*int64(-3600000000) + v121
			v130 = base.I64_div_s(v128, int64(60000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v7)+136)) = uint32(v130)
			v135 = base.I64_extend32_s(v130)*int64(-60000000) + v128
			v137 = base.I64_div_s(v135, int64(1000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v7)+132)) = uint32(v137)
			v145 = int32(0)
			v149 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_out[4]))
			F_EncodeDateTime(m, v7+int32(132), base.I32_wrap_i64(v137*int64(4293967296)+v135), v145, v145, v145, v149, v7)
			mBase = m.M
			v153 = m.ExcPending
			if v153 != 0 {
				return int64(0)
			} else {
				v178 = F_pstrdup(m, v7)
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int64(0)
				} else {
					m.G0 = v7 + int32(176)
					return base.I64_extend_i32_u(v178)
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v160 = m.ExcPending
				if v160 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_timestamp_out_6), int32(0))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_timestamp_out_7), int32(243), int32(_a_F_timestamp_out_8))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
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
}
func F_timestamp_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v23 int64
	_ = v23
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v137 int64
	_ = v137
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pq_getmsgint64(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v13
		if base.Ui64(v13-int64(9223372036854775807)) < base.Ui64(int64(2)) {
			v167 = F_AdjustTimestampForTypmod(m, v7+int32(-8), base.I32_wrap_i64(v11), int32(0))
			mBase = m.M
			v168 = m.ExcPending
			if v168 != 0 {
				return int64(0)
			} else {
				v169 = *(*int64)(unsafe.Add(mBase, uint32(v9)+56))
				m.G0 = v9 - int32(-64)
				return v169
			}
		} else {
			v23 = base.I64_div_s(v13, int64(86400000000))
			if base.Ui64(int64(172799999999)) <= base.Ui64(v13+int64(86399999999)) {
				v31 = v23 * int64(-86400000000)
			} else {
				v31 = int64(0)
			}
			v32 = v31 + v13
			v35 = v32>>(uint(int64(63))%64) + v23
			if int64(-2451545) <= v35 {
				v38 = base.I32_wrap_i64(v35)
				v50 = v38 + int32(_a_F_timestamp_recv_0)
				v51 = int32(_a_F_timestamp_recv_1)
				v52 = base.I32_div_u_s(v50, v51)
				v53 = int32(3)
				v59 = int32(2)
				v64 = base.I32_div_u_s((v52*int32(1073595727)+v50)<<(uint(v59)%32)|v53, v51)
				v67 = v38 + int32(_a_F_timestamp_recv_2) + v52*v53 + v64 + int32(_a_F_timestamp_recv_3)
				v68 = int32(1461)
				v69 = base.I32_div_u_s(v67, v68)
				v72 = v69*int32(-1461) + v67
				v74 = v72 << (uint(v59) % 32)
				if base.Ui32(v68) <= base.Ui32(v74) {
					v80 = base.I32_rem_u_s(v72+int32(305), int32(365))
					v85 = v80
				} else {
					v84 = base.I32_rem_u_s(v72+int32(306), int32(366))
					v85 = v84
				}
				v87 = base.I32_div_u_s(v74, int32(1461))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-32)))) = v87 + v69<<(uint(int32(2))%32) - int32(_a_F_timestamp_recv_4)
				v95 = v85 + int32(123)
				v99 = int32(base.Ui32(v95*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-40)))) = v95 - int32(base.Ui32(v99*int32(_a_F_timestamp_recv_5))>>(uint(int32(8))%32))
				v109 = base.I32_rem_u_s(v99+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-36)))) = v109 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v9)+44)) = int64(4294967295)
				if v32 < int64(0) {
					v121 = v32 + int64(86400000000)
				} else {
					v121 = v32
				}
				v123 = base.I64_div_s(v121, int64(3600000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v123)
				v128 = base.I64_extend32_s(v123)*int64(-3600000000) + v121
				v130 = base.I64_div_s(v128, int64(60000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v130)
				v137 = base.I64_div_s(base.I64_extend32_s(v130)*int64(-60000000)+v128, int64(1000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v137)
				if base.Ui64(v13+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v167 = F_AdjustTimestampForTypmod(m, v7+int32(-8), base.I32_wrap_i64(v11), int32(0))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int64(0)
					} else {
						v169 = *(*int64)(unsafe.Add(mBase, uint32(v9)+56))
						m.G0 = v9 - int32(-64)
						return v169
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_recv_6), int32(0))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_recv_7), int32(275), int32(_a_F_timestamp_recv_8))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
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
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamp_recv_6), int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamp_recv_7), int32(275), int32(_a_F_timestamp_recv_8))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
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
	}
}
func F_timestamp_trunc(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v561 int64
	_ = v561
	var v570 int64
	_ = v570
	var v571 int64
	_ = v571
	var v573 int64
	_ = v573
	var v576 int64
	_ = v576
	var v577 int64
	_ = v577
	var v579 int64
	_ = v579
	var v580 int64
	_ = v580
	var v584 int64
	_ = v584
	var v591 int64
	_ = v591
	var v602 int64
	_ = v602
	var v603 int64
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v620 int64
	_ = v620
	var v623 int64
	_ = v623
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v677 int64
	_ = v677
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = int32(1)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v23 = v21 & v19
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = v19
	goto L5
L4:
	;
	v24 = int32(4)
	goto L5
L5:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v21 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L139
	}
L7:
	;
	m.G0 = v12 + int32(112)
	return v677
L8:
	;
	v55 = F_downcase_truncate_identifier(m, v15+v24, v53, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v32 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v43 = int32(1)
	if v23 != 0 {
		v53 = int32(base.Ui32(v21)>>(uint(v43)%32)) - v43
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v35 = int32(16)
	goto L14
L13:
	;
	v35 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = int32(4)
	goto L17
L16:
	;
	v42 = v35
	goto L17
L17:
	;
	v53 = v42
	goto L8
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	v62 = Fn14210(m, v55, v12+int32(108), int32(_a_F_timestamp_trunc_0), int32(_a_F_timestamp_trunc_1), int32(_a_F_timestamp_trunc_2))
	mBase = m.M
	goto L20
L20:
	;
	if v62 == int32(17) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if base.Ui64(v26-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L134
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if base.Ui32(v69-int32(18)) < base.Ui32(int32(13)) {
		v677 = v26
		goto L7
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v95 = base.I64_div_s(v26, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v26+int64(86399999999)) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v82 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v55
	F_errmsg(m, int32(_a_F_timestamp_trunc_3), v12)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_5), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v103 = v95 * int64(-86400000000)
	goto L35
L34:
	;
	v103 = int64(0)
	goto L35
L35:
	;
	v104 = v103 + v26
	v107 = v104>>(uint(int64(63))%64) + v95
	if v107 <= int64(-2451546) {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v110 = base.I32_wrap_i64(v107)
	v114 = v12 + int32(84)
	v116 = v12 + int32(80)
	v118 = v12 + int32(76)
	v122 = v110 + int32(_a_F_timestamp_trunc_9)
	v123 = int32(_a_F_timestamp_trunc_10)
	v124 = base.I32_div_u_s(v122, v123)
	v125 = int32(3)
	v131 = int32(2)
	v136 = base.I32_div_u_s((v124*int32(1073595727)+v122)<<(uint(v131)%32)|v125, v123)
	v139 = v110 + int32(_a_F_timestamp_trunc_11) + v124*v125 + v136 + int32(_a_F_timestamp_trunc_12)
	v140 = int32(1461)
	v141 = base.I32_div_u_s(v139, v140)
	v144 = v141*int32(-1461) + v139
	v146 = v144 << (uint(v131) % 32)
	if base.Ui32(v140) <= base.Ui32(v146) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = int64(4294967295)
	if v104 < int64(0) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v159 = base.I32_div_u_s(v146, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v159 + v141<<(uint(int32(2))%32) - int32(_a_F_timestamp_trunc_13)
	v167 = v157 + int32(123)
	v171 = int32(base.Ui32(v167*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v167 - int32(base.Ui32(v171*int32(_a_F_timestamp_trunc_14))>>(uint(int32(8))%32))
	v181 = base.I32_rem_u_s(v171+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v181 + int32(1)
	goto L37
L39:
	;
	v152 = base.I32_rem_u_s(v144+int32(305), int32(365))
	v157 = v152
	goto L38
L40:
	;
	goto L41
L41:
	;
	v156 = base.I32_rem_u_s(v144+int32(306), int32(366))
	v157 = v156
	goto L38
L42:
	;
	v193 = v104 + int64(86400000000)
	goto L44
L43:
	;
	v193 = v104
	goto L44
L44:
	;
	v195 = base.I64_div_s(v193, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+72)) = uint32(v195)
	v200 = base.I64_extend32_s(v195)*int64(-3600000000) + v193
	v202 = base.I64_div_s(v200, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+68)) = uint32(v202)
	v207 = base.I64_extend32_s(v202)*int64(-60000000) + v200
	v209 = base.I64_div_s(v207, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+64)) = uint32(v209)
	v214 = base.I32_wrap_i64(v209*int64(4293967296) + v207)
	v215 = int32(1)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	switch v216 - int32(18) {
	case 0:
		goto L48
	case 1:
		goto L49
	case 2:
		goto L50
	case 3:
		goto L51
	case 4:
		goto L58
	case 5:
		goto L52
	case 6:
		goto L60
	case 7:
		v455 = v215
		goto L53
	case 8:
		goto L55
	case 9:
		goto L59
	case 10:
		goto L57
	case 11:
		goto L46
	case 12:
		v506 = v214
		goto L45
	default:
		goto L47
	}
L45:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if v509 <= int32(-4713) {
		goto L109
	} else {
		goto L110
	}
L46:
	;
	v504 = base.I32_rem_s(v214, int32(1000))
	v506 = v214 - v504
	goto L45
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L102
	}
L48:
	;
	v506 = int32(0)
	goto L45
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = int32(0)
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = int32(0)
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = int32(0)
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = int32(1)
	goto L51
L53:
	;
	v459 = base.I32_rem_s(v455-int32(1), int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v455 - v459
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v452
	v455 = v215
	goto L53
L55:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if int32(0) < v440 {
		goto L99
	} else {
		goto L100
	}
L56:
	;
	if int32(0) < v425 {
		goto L96
	} else {
		goto L97
	}
L57:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if int32(0) < v409 {
		goto L93
	} else {
		goto L94
	}
L58:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v225 = F_date2j(m, v221, v222, v223)
	mBase = m.M
	v226 = int32(1)
	v228 = F_date2j(m, v221, v226, int32(4))
	mBase = m.M
	v231 = F_j2day(m, v228-v226)
	mBase = m.M
	if v225 < v228-v231 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v425 = v220
	goto L56
L60:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v455 = v219
	goto L53
L61:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if base.B2i32(v266 != int32(1))|base.B2i32(v265 < int32(52)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	v234 = int32(1)
	v238 = F_date2j(m, v221-v234, v234, int32(4))
	mBase = m.M
	v241 = F_j2day(m, v238-v234)
	mBase = m.M
	v242 = v238
	v243 = v241
	goto L64
L63:
	;
	v242 = v228
	v243 = v231
	goto L64
L64:
	;
	v245 = v243 - v242 + v225
	if int32(357) <= v245 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v248 = int32(1)
	v252 = F_date2j(m, v221+v248, v248, int32(4))
	mBase = m.M
	v255 = F_j2day(m, v252-v248)
	mBase = m.M
	v256 = v252 - v255
	if v225 < v256 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v261 = v245
	goto L67
L67:
	;
	v263 = base.I32_div_s(v261, int32(7))
	v265 = v263 + int32(1)
	goto L61
L68:
	;
	v259 = v245
	goto L70
L69:
	;
	v259 = v225 - v256
	goto L70
L70:
	;
	v261 = v259
	goto L67
L71:
	;
	goto L79
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v286
	v288 = v286
	goto L71
L73:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v286 = v274 - int32(1)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if base.B2i32(v266 != int32(12))|base.B2i32(int32(1) < v265) != 0 {
		v288 = v277
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v286 = v277 + int32(1)
	goto L72
L77:
	;
	v321 = int32(7)
	v324 = int32(1)
	v329 = base.I32_rem_s(v320-v324+v324, v321)
	if v329 < int32(0) {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	goto L80
L80:
	;
	v297 = int32(_a_F_timestamp_trunc_16) + v288
	v302 = base.I32_div_s(v297, int32(4))
	v305 = base.I32_div_s(v297, int32(-100))
	v308 = base.I32_div_s(v297, int32(400))
	goto L82
L82:
	;
	goto L83
L83:
	;
	v317 = base.I32_div_s(int32(_a_F_timestamp_trunc_20), int32(256))
	v320 = int32(4) + v297*int32(365) + v302 + v305 + v308 + v317 - int32(_a_F_timestamp_trunc_17)
	goto L77
L84:
	;
	v337 = v320 + v265*v321 - v334 - int32(7)
	v341 = v337 + int32(_a_F_timestamp_trunc_21)
	v342 = int32(_a_F_timestamp_trunc_10)
	v343 = base.I32_div_u_s(v341, v342)
	v344 = int32(3)
	v350 = int32(2)
	v355 = base.I32_div_u_s((v343*int32(1073595727)+v341)<<(uint(v350)%32)|v344, v342)
	v358 = v337 + v343*v344 + v355 + int32(_a_F_timestamp_trunc_12)
	v359 = int32(1461)
	v360 = base.I32_div_u_s(v358, v359)
	v363 = v360*int32(-1461) + v358
	v365 = v363 << (uint(v350) % 32)
	if base.Ui32(v359) <= base.Ui32(v365) {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v334 = v329 + v321
	goto L87
L86:
	;
	v334 = v329
	goto L87
L87:
	;
	goto L84
L88:
	;
	v404 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v404
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(0)
	v506 = v404
	goto L45
L89:
	;
	v378 = base.I32_div_u_s(v365, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v378 + v360<<(uint(int32(2))%32) - int32(_a_F_timestamp_trunc_13)
	v386 = v376 + int32(123)
	v390 = int32(base.Ui32(v386*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v386 - int32(base.Ui32(v390*int32(_a_F_timestamp_trunc_14))>>(uint(int32(8))%32))
	v400 = base.I32_rem_u_s(v390+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v400 + int32(1)
	goto L88
L90:
	;
	v371 = base.I32_rem_u_s(v363+int32(305), int32(365))
	v376 = v371
	goto L89
L91:
	;
	goto L92
L92:
	;
	v375 = base.I32_rem_u_s(v363+int32(306), int32(366))
	v376 = v375
	goto L89
L93:
	;
	v415 = base.I32_rem_s(v409+int32(999), int32(1000))
	v425 = v409 - v415
	goto L56
L94:
	;
	goto L95
L95:
	;
	v417 = int32(1000)
	v420 = base.I32_rem_s(v417-v409, v417)
	v425 = v409 + v420 - int32(999)
	goto L56
L96:
	;
	v431 = base.I32_rem_s(v425+int32(99), int32(100))
	v452 = v425 - v431
	goto L54
L97:
	;
	goto L98
L98:
	;
	v433 = int32(100)
	v436 = base.I32_rem_s(v433-v425, v433)
	v452 = v425 + v436 - int32(99)
	goto L54
L99:
	;
	v444 = base.I32_rem_u_s(v440, int32(10))
	v452 = v440 - v444
	goto L54
L100:
	;
	goto L101
L101:
	;
	v447 = int32(9) - v440
	v449 = base.I32_rem_s(v447, int32(10))
	v452 = v449 - v447
	goto L54
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v489 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
	F_errmsg(m, int32(_a_F_timestamp_trunc_3), v12+int32(16))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_22), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L130
	}
L108:
	;
	v527 = v12 + int32(32)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v533 = base.B2i32(int32(2) < v525)
	if int32(2) < v525 {
		goto L120
	} else {
		goto L121
	}
L109:
	;
	if v509 != int32(-4713) {
		goto L107
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v509 <= int32(_a_F_timestamp_trunc_18) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if int32(10) < v514 {
		v525 = v514
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L107
L114:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v525 = v519
	goto L108
L115:
	;
	goto L116
L116:
	;
	if v509 != int32(_a_F_timestamp_trunc_19) {
		goto L107
	} else {
		goto L117
	}
L117:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if int32(5) < v522 {
		goto L107
	} else {
		goto L118
	}
L118:
	;
	v525 = v522
	goto L108
L119:
	;
	v561 = base.I64_extend_i32_s(v528 + v535*int32(365) + v540 + v543 + v546 + v555 - int32(_a_F_timestamp_trunc_17) - int32(_a_F_timestamp_trunc_11))
	v570 = int64(32)
	v571 = int64(20)
	v573 = int64(base.Ui64(v561) >> (uint(v570) % 64))
	v576 = int64(4294967295)
	v577 = int64(500654080)
	v579 = v561 & v576
	v580 = v577 * v579
	v584 = int64(base.Ui64(v580)>>(uint(v570)%64)) + v577*v573
	v591 = v579*v571 + v584&v576
	*(*int64)(unsafe.Add(mBase, uint32(v527)+8)) = v561*int64(0) + v561>>(uint(int64(63))%64)*int64(86400000000) + v571*v573 + int64(base.Ui64(v584)>>(uint(v570)%64)) + int64(base.Ui64(v591)>>(uint(v570)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v527))) = v580&v576 | v591<<(uint(v570)%64)
	goto L126
L120:
	;
	v534 = int32(_a_F_timestamp_trunc_13)
	goto L122
L121:
	;
	v534 = int32(_a_F_timestamp_trunc_16)
	goto L122
L122:
	;
	v535 = v534 + v509
	v540 = base.I32_div_s(v535, int32(4))
	v543 = base.I32_div_s(v535, int32(-100))
	v546 = base.I32_div_s(v535, int32(400))
	if int32(2) < v525 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v550 = int32(1)
	goto L125
L124:
	;
	v550 = int32(13)
	goto L125
L125:
	;
	v555 = base.I32_div_s((v550+v525)*int32(_a_F_timestamp_trunc_14), int32(256))
	goto L119
L126:
	;
	v602 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	if v602 != v603>>(uint(int64(63))%64) {
		goto L107
	} else {
		goto L127
	}
L127:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	v611 = int32(60)
	v620 = base.I64_extend_i32_s(v506) + base.I64_extend_i32_s(v608+(v609+v610*v611)*v611)*int64(1000000)
	v623 = v603 + v620
	if base.B2i32(v620 < int64(0))^base.B2i32(v623 < v603) != 0 {
		goto L107
	} else {
		goto L128
	}
L128:
	;
	if base.Ui64(int64(9011559254509551615)) < base.Ui64(v623-int64(9223371331200000000)) {
		v677 = v623
		goto L7
	} else {
		goto L129
	}
L129:
	;
	goto L107
L130:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_timestamp_trunc_7), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_15), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v657 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v657
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v55
	F_errmsg(m, int32(_a_F_timestamp_trunc_23), v12+int32(48))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_24), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_timestamp_trunc_7), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_8), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
