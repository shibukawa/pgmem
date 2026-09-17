package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AdjustTimestampForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(v12-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		m.G0 = v10 + int32(16)
		return
	} else {
		switch l1 + int32(1) {
		case 0, 7:
			m.G0 = v10 + int32(16)
			return
		default:
			if base.Ui32(int32(7)) <= base.Ui32(l1) {
				v21 = F_errsave_start(m, l2)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					if v21 == int32(0) {
						m.G0 = v10 + int32(16)
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(25769803776)
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
							F_errmsg(m, int32(_a_F_AdjustTimestampForTypmod_0), v10)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_errsave_finish(m, l2, int32(_a_F_AdjustTimestampForTypmod_1), int32(397), int32(_a_F_AdjustTimestampForTypmod_2))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v40 = l1 << (uint(int32(3)) % 32)
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_AdjustTimestampForTypmod[0])))
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_AdjustTimestampForTypmod[1])))
				if int64(0) <= v12 {
					v45 = v12 + v42
					v46 = base.I64_rem_s(v45, v41)
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v45 - v46
				} else {
					v49 = v42 - v12
					v50 = base.I64_rem_s(v49, v41)
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50 - v49
				}
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
func F_timestamp_ne_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) != int32(0))
	} else {
		if v6 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) != int32(0))
		} else {
			if int32(106751982) < v6 {
				v24 = int32(1)
			} else {
				v18 = base.I64_extend_i32_s(v6) * int64(86400000000)
				v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) != int32(0))
			}
		}
	}
	return v24
}
func F_timestamp_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui64(v10-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		if v10 != int64(-9223372036854775807-1) {
			v172 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_timestamp_out[0])))
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v172)
			v175 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_out[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v175
		} else {
			v18 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_timestamp_out[2])))
			*(*uint16)(unsafe.Add(mBase, uint32(v7)+8)) = uint16(v18)
			v21 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_out[3]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v21
		}
		v179 = F_pstrdup(m, v7)
		mBase = m.M
		v180 = m.ExcPending
		if v180 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(176)
			return v179
		}
	} else {
		v24 = base.I64_div_s(v10, int64(86400000000))
		if base.Ui64(int64(172799999999)) <= base.Ui64(v10+int64(86399999999)) {
			v32 = v24 * int64(-86400000000)
		} else {
			v32 = int64(0)
		}
		v33 = v32 + v10
		v36 = v33>>(uint(int64(63))%64) + v24
		if int64(-2451545) <= v36 {
			v39 = base.I32_wrap_i64(v36)
			v51 = v39 + int32(_a_F_timestamp_out_0)
			v52 = int32(_a_F_timestamp_out_1)
			v53 = base.I32_div_u_s(v51, v52)
			v54 = int32(3)
			v60 = int32(2)
			v65 = base.I32_div_u_s((v53*int32(1073595727)+v51)<<(uint(v60)%32)|v54, v52)
			v68 = v39 + int32(_a_F_timestamp_out_2) + v53*v54 + v65 + int32(_a_F_timestamp_out_3)
			v69 = int32(1461)
			v70 = base.I32_div_u_s(v68, v69)
			v73 = v70*int32(-1461) + v68
			v75 = v73 << (uint(v60) % 32)
			if base.Ui32(v69) <= base.Ui32(v75) {
				v81 = base.I32_rem_u_s(v73+int32(305), int32(365))
				v86 = v81
			} else {
				v85 = base.I32_rem_u_s(v73+int32(306), int32(366))
				v86 = v85
			}
			v88 = base.I32_div_u_s(v75, int32(1461))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(152)))) = v88 + v70<<(uint(int32(2))%32) - int32(_a_F_timestamp_out_4)
			v96 = v86 + int32(123)
			v100 = int32(base.Ui32(v96*int32(2141)) >> (uint(int32(16)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(144)))) = v96 - int32(base.Ui32(v100*int32(_a_F_timestamp_out_5))>>(uint(int32(8))%32))
			v110 = base.I32_rem_u_s(v100+int32(10), int32(12))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(148)))) = v110 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+164)) = int64(4294967295)
			if v33 < int64(0) {
				v122 = v33 + int64(86400000000)
			} else {
				v122 = v33
			}
			v124 = base.I64_div_s(v122, int64(3600000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v7)+140)) = uint32(v124)
			v129 = base.I64_extend32_s(v124)*int64(-3600000000) + v122
			v131 = base.I64_div_s(v129, int64(60000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v7)+136)) = uint32(v131)
			v136 = base.I64_extend32_s(v131)*int64(-60000000) + v129
			v138 = base.I64_div_s(v136, int64(1000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v7)+132)) = uint32(v138)
			v146 = int32(0)
			v150 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_out[4]))
			F_EncodeDateTime(m, v7+int32(132), base.I32_wrap_i64(v138*int64(4293967296)+v136), v146, v146, v146, v150, v7)
			mBase = m.M
			v154 = m.ExcPending
			if v154 != 0 {
				return int32(0)
			} else {
				v179 = F_pstrdup(m, v7)
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(176)
					return v179
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v158 = m.ExcPending
			if v158 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_timestamp_out_6), int32(0))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_timestamp_out_7), int32(250), int32(_a_F_timestamp_out_8))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
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
func F_timestamp_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
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
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pq_getmsgint64(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v13
		if base.Ui64(v13-int64(9223372036854775807)) < base.Ui64(int64(2)) {
			F_AdjustTimestampForTypmod(m, v7+int32(-8), v11, int32(0))
			mBase = m.M
			v167 = m.ExcPending
			if v167 != 0 {
				return int32(0)
			} else {
				v168 = *(*int64)(unsafe.Add(mBase, uint32(v9)+56))
				v169 = F_Int64GetDatum(m, v168)
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 - int32(-64)
					return v169
				}
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
					F_AdjustTimestampForTypmod(m, v7+int32(-8), v11, int32(0))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						v168 = *(*int64)(unsafe.Add(mBase, uint32(v9)+56))
						v169 = F_Int64GetDatum(m, v168)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 - int32(-64)
							return v169
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_recv_6), int32(0))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_recv_7), int32(282), int32(_a_F_timestamp_recv_8))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamp_recv_6), int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamp_recv_7), int32(282), int32(_a_F_timestamp_recv_8))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
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
func F_timestamp_trunc(m *base.Module, l0 int32) int32 {
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
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v563 int64
	_ = v563
	var v572 int64
	_ = v572
	var v573 int64
	_ = v573
	var v575 int64
	_ = v575
	var v578 int64
	_ = v578
	var v579 int64
	_ = v579
	var v581 int64
	_ = v581
	var v582 int64
	_ = v582
	var v586 int64
	_ = v586
	var v593 int64
	_ = v593
	var v604 int64
	_ = v604
	var v605 int64
	_ = v605
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v622 int64
	_ = v622
	var v625 int64
	_ = v625
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v679 int64
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v21 = int32(1)
	v22 = v15 + v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v27 = v25 & v21
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L140
	}
L4:
	;
	v681 = F_Int64GetDatum(m, v679)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L139
	}
L5:
	;
	v28 = v22
	goto L7
L6:
	;
	v28 = v15 + int32(4)
	goto L7
L7:
	;
	if v25 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v57 = F_downcase_truncate_identifier(m, v28, v55, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v34 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v45 = int32(1)
	if v27 != 0 {
		v55 = int32(base.Ui32(v25)>>(uint(v45)%32)) - v45
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v37 = int32(16)
	goto L14
L13:
	;
	v37 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v44 = int32(4)
	goto L17
L16:
	;
	v44 = v37
	goto L17
L17:
	;
	v55 = v44
	goto L8
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	v64 = Fn13825(m, v57, v12+int32(108), int32(_a_F_timestamp_trunc_0), int32(_a_F_timestamp_trunc_1), int32(_a_F_timestamp_trunc_2))
	mBase = m.M
	goto L20
L20:
	;
	if v64 == int32(17) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if base.Ui64(v20-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
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
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L134
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if base.Ui32(v71-int32(18)) < base.Ui32(int32(13)) {
		v679 = v20
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v97 = base.I64_div_s(v20, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v20+int64(86399999999)) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v84 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
	F_errmsg(m, int32(_a_F_timestamp_trunc_3), v12)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_5), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
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
	v105 = v97 * int64(-86400000000)
	goto L35
L34:
	;
	v105 = int64(0)
	goto L35
L35:
	;
	v106 = v105 + v20
	v109 = v106>>(uint(int64(63))%64) + v97
	if v109 <= int64(-2451546) {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v112 = base.I32_wrap_i64(v109)
	v116 = v12 + int32(84)
	v118 = v12 + int32(80)
	v120 = v12 + int32(76)
	v124 = v112 + int32(_a_F_timestamp_trunc_9)
	v125 = int32(_a_F_timestamp_trunc_10)
	v126 = base.I32_div_u_s(v124, v125)
	v127 = int32(3)
	v133 = int32(2)
	v138 = base.I32_div_u_s((v126*int32(1073595727)+v124)<<(uint(v133)%32)|v127, v125)
	v141 = v112 + int32(_a_F_timestamp_trunc_11) + v126*v127 + v138 + int32(_a_F_timestamp_trunc_12)
	v142 = int32(1461)
	v143 = base.I32_div_u_s(v141, v142)
	v146 = v143*int32(-1461) + v141
	v148 = v146 << (uint(v133) % 32)
	if base.Ui32(v142) <= base.Ui32(v148) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = int64(4294967295)
	if v106 < int64(0) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v161 = base.I32_div_u_s(v148, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v161 + v143<<(uint(int32(2))%32) - int32(_a_F_timestamp_trunc_13)
	v169 = v159 + int32(123)
	v173 = int32(base.Ui32(v169*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v169 - int32(base.Ui32(v173*int32(_a_F_timestamp_trunc_14))>>(uint(int32(8))%32))
	v183 = base.I32_rem_u_s(v173+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v183 + int32(1)
	goto L37
L39:
	;
	v154 = base.I32_rem_u_s(v146+int32(305), int32(365))
	v159 = v154
	goto L38
L40:
	;
	goto L41
L41:
	;
	v158 = base.I32_rem_u_s(v146+int32(306), int32(366))
	v159 = v158
	goto L38
L42:
	;
	v195 = v106 + int64(86400000000)
	goto L44
L43:
	;
	v195 = v106
	goto L44
L44:
	;
	v197 = base.I64_div_s(v195, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+72)) = uint32(v197)
	v202 = base.I64_extend32_s(v197)*int64(-3600000000) + v195
	v204 = base.I64_div_s(v202, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+68)) = uint32(v204)
	v209 = base.I64_extend32_s(v204)*int64(-60000000) + v202
	v211 = base.I64_div_s(v209, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+64)) = uint32(v211)
	v216 = base.I32_wrap_i64(v211*int64(4293967296) + v209)
	v217 = int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	switch v218 - int32(18) {
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
		v457 = v217
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
		v508 = v216
		goto L45
	default:
		goto L47
	}
L45:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if v511 <= int32(-4713) {
		goto L109
	} else {
		goto L110
	}
L46:
	;
	v506 = base.I32_rem_s(v216, int32(1000))
	v508 = v216 - v506
	goto L45
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L102
	}
L48:
	;
	v508 = int32(0)
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
	v461 = base.I32_rem_s(v457-int32(1), int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v457 - v461
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v454
	v457 = v217
	goto L53
L55:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if int32(0) < v442 {
		goto L99
	} else {
		goto L100
	}
L56:
	;
	if int32(0) < v427 {
		goto L96
	} else {
		goto L97
	}
L57:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if int32(0) < v411 {
		goto L93
	} else {
		goto L94
	}
L58:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v227 = F_date2j(m, v223, v224, v225)
	mBase = m.M
	v228 = int32(1)
	v230 = F_date2j(m, v223, v228, int32(4))
	mBase = m.M
	v233 = F_j2day(m, v230-v228)
	mBase = m.M
	if v227 < v230-v233 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v427 = v222
	goto L56
L60:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v457 = v221
	goto L53
L61:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if base.B2i32(v268 != int32(1))|base.B2i32(v267 < int32(52)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	v236 = int32(1)
	v240 = F_date2j(m, v223-v236, v236, int32(4))
	mBase = m.M
	v243 = F_j2day(m, v240-v236)
	mBase = m.M
	v244 = v240
	v245 = v243
	goto L64
L63:
	;
	v244 = v230
	v245 = v233
	goto L64
L64:
	;
	v247 = v245 - v244 + v227
	if int32(357) <= v247 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v250 = int32(1)
	v254 = F_date2j(m, v223+v250, v250, int32(4))
	mBase = m.M
	v257 = F_j2day(m, v254-v250)
	mBase = m.M
	v258 = v254 - v257
	if v227 < v258 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v263 = v247
	goto L67
L67:
	;
	v265 = base.I32_div_s(v263, int32(7))
	v267 = v265 + int32(1)
	goto L61
L68:
	;
	v261 = v247
	goto L70
L69:
	;
	v261 = v227 - v258
	goto L70
L70:
	;
	v263 = v261
	goto L67
L71:
	;
	goto L79
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v288
	v290 = v288
	goto L71
L73:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v288 = v276 - int32(1)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if base.B2i32(v268 != int32(12))|base.B2i32(int32(1) < v267) != 0 {
		v290 = v279
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v288 = v279 + int32(1)
	goto L72
L77:
	;
	v323 = int32(7)
	v326 = int32(1)
	v331 = base.I32_rem_s(v322-v326+v326, v323)
	if v331 < int32(0) {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	goto L80
L80:
	;
	v299 = int32(_a_F_timestamp_trunc_16) + v290
	v304 = base.I32_div_s(v299, int32(4))
	v307 = base.I32_div_s(v299, int32(-100))
	v310 = base.I32_div_s(v299, int32(400))
	goto L82
L82:
	;
	goto L83
L83:
	;
	v319 = base.I32_div_s(int32(_a_F_timestamp_trunc_20), int32(256))
	v322 = int32(4) + v299*int32(365) + v304 + v307 + v310 + v319 - int32(_a_F_timestamp_trunc_17)
	goto L77
L84:
	;
	v339 = v322 + v267*v323 - v336 - int32(7)
	v343 = v339 + int32(_a_F_timestamp_trunc_21)
	v344 = int32(_a_F_timestamp_trunc_10)
	v345 = base.I32_div_u_s(v343, v344)
	v346 = int32(3)
	v352 = int32(2)
	v357 = base.I32_div_u_s((v345*int32(1073595727)+v343)<<(uint(v352)%32)|v346, v344)
	v360 = v339 + v345*v346 + v357 + int32(_a_F_timestamp_trunc_12)
	v361 = int32(1461)
	v362 = base.I32_div_u_s(v360, v361)
	v365 = v362*int32(-1461) + v360
	v367 = v365 << (uint(v352) % 32)
	if base.Ui32(v361) <= base.Ui32(v367) {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v336 = v331 + v323
	goto L87
L86:
	;
	v336 = v331
	goto L87
L87:
	;
	goto L84
L88:
	;
	v406 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v406
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(0)
	v508 = v406
	goto L45
L89:
	;
	v380 = base.I32_div_u_s(v367, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v380 + v362<<(uint(int32(2))%32) - int32(_a_F_timestamp_trunc_13)
	v388 = v378 + int32(123)
	v392 = int32(base.Ui32(v388*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v388 - int32(base.Ui32(v392*int32(_a_F_timestamp_trunc_14))>>(uint(int32(8))%32))
	v402 = base.I32_rem_u_s(v392+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v402 + int32(1)
	goto L88
L90:
	;
	v373 = base.I32_rem_u_s(v365+int32(305), int32(365))
	v378 = v373
	goto L89
L91:
	;
	goto L92
L92:
	;
	v377 = base.I32_rem_u_s(v365+int32(306), int32(366))
	v378 = v377
	goto L89
L93:
	;
	v417 = base.I32_rem_s(v411+int32(999), int32(1000))
	v427 = v411 - v417
	goto L56
L94:
	;
	goto L95
L95:
	;
	v419 = int32(1000)
	v422 = base.I32_rem_s(v419-v411, v419)
	v427 = v411 + v422 - int32(999)
	goto L56
L96:
	;
	v433 = base.I32_rem_s(v427+int32(99), int32(100))
	v454 = v427 - v433
	goto L54
L97:
	;
	goto L98
L98:
	;
	v435 = int32(100)
	v438 = base.I32_rem_s(v435-v427, v435)
	v454 = v427 + v438 - int32(99)
	goto L54
L99:
	;
	v446 = base.I32_rem_u_s(v442, int32(10))
	v454 = v442 - v446
	goto L54
L100:
	;
	goto L101
L101:
	;
	v449 = int32(9) - v442
	v451 = base.I32_rem_s(v449, int32(10))
	v454 = v451 - v449
	goto L54
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v491 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v57
	F_errmsg(m, int32(_a_F_timestamp_trunc_3), v12+int32(16))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_22), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
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
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L130
	}
L108:
	;
	v529 = v12 + int32(32)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v535 = base.B2i32(int32(2) < v527)
	if int32(2) < v527 {
		goto L120
	} else {
		goto L121
	}
L109:
	;
	if v511 != int32(-4713) {
		goto L107
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v511 <= int32(_a_F_timestamp_trunc_18) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if int32(10) < v516 {
		v527 = v516
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L107
L114:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v527 = v521
	goto L108
L115:
	;
	goto L116
L116:
	;
	if v511 != int32(_a_F_timestamp_trunc_19) {
		goto L107
	} else {
		goto L117
	}
L117:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if int32(5) < v524 {
		goto L107
	} else {
		goto L118
	}
L118:
	;
	v527 = v524
	goto L108
L119:
	;
	v563 = base.I64_extend_i32_s(v530 + v537*int32(365) + v542 + v545 + v548 + v557 - int32(_a_F_timestamp_trunc_17) - int32(_a_F_timestamp_trunc_11))
	v572 = int64(32)
	v573 = int64(20)
	v575 = int64(base.Ui64(v563) >> (uint(v572) % 64))
	v578 = int64(4294967295)
	v579 = int64(500654080)
	v581 = v563 & v578
	v582 = v579 * v581
	v586 = int64(base.Ui64(v582)>>(uint(v572)%64)) + v579*v575
	v593 = v581*v573 + v586&v578
	*(*int64)(unsafe.Add(mBase, uint32(v529)+8)) = v563*int64(0) + v563>>(uint(int64(63))%64)*int64(86400000000) + v573*v575 + int64(base.Ui64(v586)>>(uint(v572)%64)) + int64(base.Ui64(v593)>>(uint(v572)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v529))) = v582&v578 | v593<<(uint(v572)%64)
	goto L126
L120:
	;
	v536 = int32(_a_F_timestamp_trunc_13)
	goto L122
L121:
	;
	v536 = int32(_a_F_timestamp_trunc_16)
	goto L122
L122:
	;
	v537 = v536 + v511
	v542 = base.I32_div_s(v537, int32(4))
	v545 = base.I32_div_s(v537, int32(-100))
	v548 = base.I32_div_s(v537, int32(400))
	if int32(2) < v527 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v552 = int32(1)
	goto L125
L124:
	;
	v552 = int32(13)
	goto L125
L125:
	;
	v557 = base.I32_div_s((v552+v527)*int32(_a_F_timestamp_trunc_14), int32(256))
	goto L119
L126:
	;
	v604 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	v605 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	if v604 != v605>>(uint(int64(63))%64) {
		goto L107
	} else {
		goto L127
	}
L127:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	v613 = int32(60)
	v622 = base.I64_extend_i32_s(v508) + base.I64_extend_i32_s(v610+(v611+v612*v613)*v613)*int64(1000000)
	v625 = v605 + v622
	if base.B2i32(v622 < int64(0))^base.B2i32(v625 < v605) != 0 {
		goto L107
	} else {
		goto L128
	}
L128:
	;
	if base.Ui64(int64(9011559254509551615)) < base.Ui64(v625-int64(9223371331200000000)) {
		v679 = v625
		goto L4
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
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_timestamp_trunc_7), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_15), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
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
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v659 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v659
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v57
	F_errmsg(m, int32(_a_F_timestamp_trunc_23), v12+int32(48))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_24), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
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
	m.G0 = v12 + int32(112)
	return v681
L140:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(_a_F_timestamp_trunc_7), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_timestamp_trunc_4), int32(_a_F_timestamp_trunc_8), int32(_a_F_timestamp_trunc_6))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
