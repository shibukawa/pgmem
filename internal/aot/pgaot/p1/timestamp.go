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
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
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
							F_errmsg(m, int32(501674), v10)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_errsave_finish(m, l2, int32(521224), int32(397), int32(443654))
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
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[1143])))
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[1144])))
				if int64(0) <= v12 {
					v49 = v12 + v46
					v50 = base.I64_rem_s(v49, v43)
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v49 - v50
				} else {
					v53 = v46 - v12
					v54 = base.I64_rem_s(v53, v43)
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v54 - v53
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
			v172 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v172)
			v175 = *(*int64)(unsafe.Add(mBase, _consts[1063]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v175
		} else {
			v18 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1060])))
			*(*uint16)(unsafe.Add(mBase, uint32(v7)+8)) = uint16(v18)
			v21 = *(*int64)(unsafe.Add(mBase, _consts[1061]))
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
			v51 = v39 + int32(2483589)
			v52 = int32(146097)
			v53 = base.I32_div_u_s(v51, v52)
			v54 = int32(3)
			v60 = int32(2)
			v65 = base.I32_div_u_s((v53*int32(1073595727)+v51)<<(uint(v60)%32)|v54, v52)
			v68 = v39 + int32(2451545) + v53*v54 + v65 + int32(32104)
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
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(152)))) = v88 + v70<<(uint(int32(2))%32) - int32(4800)
			v96 = v86 + int32(123)
			v100 = int32(base.Ui32(v96*int32(2141)) >> (uint(int32(16)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(144)))) = v96 - int32(base.Ui32(v100*int32(7834))>>(uint(int32(8))%32))
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
			v150 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
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
					F_errmsg(m, int32(422473), int32(0))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(521224), int32(250), int32(73060))
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
				v50 = v38 + int32(2483589)
				v51 = int32(146097)
				v52 = base.I32_div_u_s(v50, v51)
				v53 = int32(3)
				v59 = int32(2)
				v64 = base.I32_div_u_s((v52*int32(1073595727)+v50)<<(uint(v59)%32)|v53, v51)
				v67 = v38 + int32(2451545) + v52*v53 + v64 + int32(32104)
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
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-32)))) = v87 + v69<<(uint(int32(2))%32) - int32(4800)
				v95 = v85 + int32(123)
				v99 = int32(base.Ui32(v95*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-40)))) = v95 - int32(base.Ui32(v99*int32(7834))>>(uint(int32(8))%32))
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
							F_errmsg(m, int32(422473), int32(0))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521224), int32(282), int32(37964))
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
						F_errmsg(m, int32(422473), int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(521224), int32(282), int32(37964))
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v624 int64
	_ = v624
	var v633 int64
	_ = v633
	var v634 int64
	_ = v634
	var v636 int64
	_ = v636
	var v639 int64
	_ = v639
	var v640 int64
	_ = v640
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v647 int64
	_ = v647
	var v654 int64
	_ = v654
	var v665 int64
	_ = v665
	var v666 int64
	_ = v666
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v683 int64
	_ = v683
	var v686 int64
	_ = v686
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v740 int64
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
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
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L159
	}
L4:
	;
	v742 = F_Int64GetDatum(m, v740)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L158
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
	v58 = F_downcase_truncate_identifier(m, v28, v56, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v31 = int32(4)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v33&int32(254) == int32(2) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v46 = int32(1)
	if v27 != 0 {
		v56 = int32(base.Ui32(v25)>>(uint(v46)%32)) - v46
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v42 = v31
	goto L14
L13:
	;
	v42 = base.B2i32(v33 == int32(18)) << (uint(v31) % 32)
	goto L14
L14:
	;
	if v33 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v45 = v31
	goto L17
L16:
	;
	v45 = v42
	goto L17
L17:
	;
	v56 = v45
	goto L8
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	v61 = v12 + int32(108)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[1071]))
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v129 == int32(17) {
		goto L39
	} else {
		goto L40
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1071])) = v112
	v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112)+11)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v120
	v129 = v119
	goto L20
L22:
	;
	v70 = F_strncmp(m, v58, v68, int32(10))
	mBase = m.M
	if v70 == int32(0) {
		v112 = v68
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v58))))
	v79 = int32(1696416)
	v81 = int32(1697376)
	goto L26
L25:
	;
	goto L24
L26:
	;
	v88 = v79 + (v81-v79)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v88))))
	v90 = v73 - v89
	if v90 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(0)
	v129 = int32(31)
	goto L20
L28:
	;
	v94 = F_strncmp(m, v58, v88, int32(10))
	mBase = m.M
	if v94 == int32(0) {
		v112 = v88
		goto L21
	} else {
		goto L31
	}
L29:
	;
	v97 = v90
	goto L30
L30:
	;
	v101 = base.B2i32(v97 < int32(0))
	if v97 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v97 = v94
	goto L30
L32:
	;
	v102 = v88 - int32(16)
	goto L34
L33:
	;
	v102 = v81
	goto L34
L34:
	;
	if v97 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v105 = v79
	goto L37
L36:
	;
	v105 = v88 + int32(16)
	goto L37
L37:
	;
	if base.Ui32(v105) <= base.Ui32(v102) {
		v79 = v105
		v81 = v102
		goto L26
	} else {
		goto L38
	}
L38:
	;
	goto L27
L39:
	;
	if base.Ui64(v20-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L153
	}
L42:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if base.Ui32(v136-int32(18)) < base.Ui32(int32(13)) {
		v740 = v20
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v162 = base.I64_div_s(v20, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v20+int64(86399999999)) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v149 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v58
	F_errmsg(m, int32(201084), v12)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(521224), int32(4725), int32(513618))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v170 = v162 * int64(-86400000000)
	goto L53
L52:
	;
	v170 = int64(0)
	goto L53
L53:
	;
	v171 = v170 + v20
	v174 = v171>>(uint(int64(63))%64) + v162
	if v174 <= int64(-2451546) {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v177 = base.I32_wrap_i64(v174)
	v181 = v12 + int32(84)
	v183 = v12 + int32(80)
	v185 = v12 + int32(76)
	v189 = v177 + int32(2483589)
	v190 = int32(146097)
	v191 = base.I32_div_u_s(v189, v190)
	v192 = int32(3)
	v198 = int32(2)
	v203 = base.I32_div_u_s((v191*int32(1073595727)+v189)<<(uint(v198)%32)|v192, v190)
	v206 = v177 + int32(2451545) + v191*v192 + v203 + int32(32104)
	v207 = int32(1461)
	v208 = base.I32_div_u_s(v206, v207)
	v211 = v208*int32(-1461) + v206
	v213 = v211 << (uint(v198) % 32)
	if base.Ui32(v207) <= base.Ui32(v213) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = int64(4294967295)
	if v171 < int64(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v226 = base.I32_div_u_s(v213, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v226 + v208<<(uint(int32(2))%32) - int32(4800)
	v234 = v224 + int32(123)
	v238 = int32(base.Ui32(v234*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v234 - int32(base.Ui32(v238*int32(7834))>>(uint(int32(8))%32))
	v248 = base.I32_rem_u_s(v238+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v248 + int32(1)
	goto L55
L57:
	;
	v219 = base.I32_rem_u_s(v211+int32(305), int32(365))
	v224 = v219
	goto L56
L58:
	;
	goto L59
L59:
	;
	v223 = base.I32_rem_u_s(v211+int32(306), int32(366))
	v224 = v223
	goto L56
L60:
	;
	v260 = v171 + int64(86400000000)
	goto L62
L61:
	;
	v260 = v171
	goto L62
L62:
	;
	v262 = base.I64_div_s(v260, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+72)) = uint32(v262)
	v267 = base.I64_extend32_s(v262)*int64(-3600000000) + v260
	v269 = base.I64_div_s(v267, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+68)) = uint32(v269)
	v274 = base.I64_extend32_s(v269)*int64(-60000000) + v267
	v276 = base.I64_div_s(v274, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+64)) = uint32(v276)
	v281 = base.I32_wrap_i64(v276*int64(4293967296) + v274)
	v282 = int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	switch v283 - int32(18) {
	case 0:
		goto L66
	case 1:
		goto L67
	case 2:
		goto L68
	case 3:
		goto L69
	case 4:
		goto L76
	case 5:
		goto L70
	case 6:
		goto L78
	case 7:
		v518 = v282
		goto L71
	case 8:
		goto L73
	case 9:
		goto L77
	case 10:
		goto L75
	case 11:
		goto L64
	case 12:
		v569 = v281
		goto L63
	default:
		goto L65
	}
L63:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if v572 <= int32(-4713) {
		goto L128
	} else {
		goto L129
	}
L64:
	;
	v567 = base.I32_rem_s(v281, int32(1000))
	v569 = v281 - v567
	goto L63
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L121
	}
L66:
	;
	v569 = int32(0)
	goto L63
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = int32(0)
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = int32(0)
	goto L67
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = int32(0)
	goto L68
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = int32(1)
	goto L69
L71:
	;
	v522 = base.I32_rem_s(v518-int32(1), int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v518 - v522
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v515
	v518 = v282
	goto L71
L73:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if int32(0) < v503 {
		goto L118
	} else {
		goto L119
	}
L74:
	;
	if int32(0) < v488 {
		goto L115
	} else {
		goto L116
	}
L75:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if int32(0) < v472 {
		goto L112
	} else {
		goto L113
	}
L76:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v292 = F_date2j(m, v288, v289, v290)
	mBase = m.M
	v293 = int32(1)
	v295 = F_date2j(m, v288, v293, int32(4))
	mBase = m.M
	v298 = F_j2day(m, v295-v293)
	mBase = m.M
	if v292 < v295-v298 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v488 = v287
	goto L74
L78:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v518 = v286
	goto L71
L79:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v332 < int32(52) {
		goto L91
	} else {
		goto L92
	}
L80:
	;
	v301 = int32(1)
	v305 = F_date2j(m, v288-v301, v301, int32(4))
	mBase = m.M
	v308 = F_j2day(m, v305-v301)
	mBase = m.M
	v309 = v305
	v310 = v308
	goto L82
L81:
	;
	v309 = v295
	v310 = v298
	goto L82
L82:
	;
	v312 = v310 - v309 + v292
	if int32(357) <= v312 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v315 = int32(1)
	v319 = F_date2j(m, v288+v315, v315, int32(4))
	mBase = m.M
	v322 = F_j2day(m, v319-v315)
	mBase = m.M
	v323 = v319 - v322
	if v292 < v323 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v328 = v312
	goto L85
L85:
	;
	v330 = base.I32_div_s(v328, int32(7))
	v332 = v330 + int32(1)
	goto L79
L86:
	;
	v326 = v312
	goto L88
L87:
	;
	v326 = v292 - v323
	goto L88
L88:
	;
	v328 = v326
	goto L85
L89:
	;
	goto L98
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v349
	v351 = v349
	goto L89
L91:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	if int32(1) < v332 {
		v351 = v341
		goto L89
	} else {
		goto L94
	}
L92:
	;
	if v333 != int32(1) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v349 = v338 - int32(1)
	goto L90
L94:
	;
	if v333 != int32(12) {
		v351 = v341
		goto L89
	} else {
		goto L95
	}
L95:
	;
	v349 = v341 + int32(1)
	goto L90
L96:
	;
	v387 = int32(1)
	v391 = int32(7)
	v392 = base.I32_rem_s(v385-v387+v387, v391)
	if v392 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	goto L99
L99:
	;
	v362 = int32(4799) + v351
	v367 = base.I32_div_s(v362, int32(4))
	v370 = base.I32_div_s(v362, int32(-100))
	v373 = base.I32_div_s(v362, int32(400))
	goto L101
L101:
	;
	goto L102
L102:
	;
	v382 = base.I32_div_s(int32(109676), int32(256))
	v385 = int32(4) + v362*int32(365) + v367 + v370 + v373 + v382 - int32(32167)
	goto L96
L103:
	;
	v400 = v332*int32(7) + v385 - v397 - int32(7)
	v404 = v400 + int32(32044)
	v405 = int32(146097)
	v406 = base.I32_div_u_s(v404, v405)
	v407 = int32(3)
	v413 = int32(2)
	v418 = base.I32_div_u_s((v406*int32(1073595727)+v404)<<(uint(v413)%32)|v407, v405)
	v421 = v400 + v406*v407 + v418 + int32(32104)
	v422 = int32(1461)
	v423 = base.I32_div_u_s(v421, v422)
	v426 = v423*int32(-1461) + v421
	v428 = v426 << (uint(v413) % 32)
	if base.Ui32(v422) <= base.Ui32(v428) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v397 = v392 + v391
	goto L106
L105:
	;
	v397 = v392
	goto L106
L106:
	;
	goto L103
L107:
	;
	v467 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(0)
	v569 = v467
	goto L63
L108:
	;
	v441 = base.I32_div_u_s(v428, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v441 + v423<<(uint(int32(2))%32) - int32(4800)
	v449 = v439 + int32(123)
	v453 = int32(base.Ui32(v449*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v449 - int32(base.Ui32(v453*int32(7834))>>(uint(int32(8))%32))
	v463 = base.I32_rem_u_s(v453+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v463 + int32(1)
	goto L107
L109:
	;
	v434 = base.I32_rem_u_s(v426+int32(305), int32(365))
	v439 = v434
	goto L108
L110:
	;
	goto L111
L111:
	;
	v438 = base.I32_rem_u_s(v426+int32(306), int32(366))
	v439 = v438
	goto L108
L112:
	;
	v478 = base.I32_rem_s(v472+int32(999), int32(1000))
	v488 = v472 - v478
	goto L74
L113:
	;
	goto L114
L114:
	;
	v480 = int32(1000)
	v483 = base.I32_rem_s(v480-v472, v480)
	v488 = v472 + v483 - int32(999)
	goto L74
L115:
	;
	v494 = base.I32_rem_s(v488+int32(99), int32(100))
	v515 = v488 - v494
	goto L72
L116:
	;
	goto L117
L117:
	;
	v496 = int32(100)
	v499 = base.I32_rem_s(v496-v488, v496)
	v515 = v488 + v499 - int32(99)
	goto L72
L118:
	;
	v507 = base.I32_rem_u_s(v503, int32(10))
	v515 = v503 - v507
	goto L72
L119:
	;
	goto L120
L120:
	;
	v510 = int32(9) - v503
	v512 = base.I32_rem_s(v510, int32(10))
	v515 = v512 - v510
	goto L72
L121:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v552 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v552
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v58
	F_errmsg(m, int32(201084), v12+int32(16))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(521224), int32(4816), int32(513618))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L149
	}
L127:
	;
	v590 = v12 + int32(32)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v596 = base.B2i32(int32(2) < v588)
	if int32(2) < v588 {
		goto L139
	} else {
		goto L140
	}
L128:
	;
	if v572 != int32(-4713) {
		goto L126
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if v572 <= int32(5874897) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if int32(10) < v577 {
		v588 = v577
		goto L127
	} else {
		goto L132
	}
L132:
	;
	goto L126
L133:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v588 = v582
	goto L127
L134:
	;
	goto L135
L135:
	;
	if v572 != int32(5874898) {
		goto L126
	} else {
		goto L136
	}
L136:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if int32(5) < v585 {
		goto L126
	} else {
		goto L137
	}
L137:
	;
	v588 = v585
	goto L127
L138:
	;
	v624 = base.I64_extend_i32_s(v591 + v598*int32(365) + v603 + v606 + v609 + v618 - int32(32167) - int32(2451545))
	v633 = int64(32)
	v634 = int64(20)
	v636 = int64(base.Ui64(v624) >> (uint(v633) % 64))
	v639 = int64(4294967295)
	v640 = int64(500654080)
	v642 = v624 & v639
	v643 = v640 * v642
	v647 = int64(base.Ui64(v643)>>(uint(v633)%64)) + v640*v636
	v654 = v642*v634 + v647&v639
	*(*int64)(unsafe.Add(mBase, uint32(v590)+8)) = v624*int64(0) + v624>>(uint(int64(63))%64)*int64(86400000000) + v634*v636 + int64(base.Ui64(v647)>>(uint(v633)%64)) + int64(base.Ui64(v654)>>(uint(v633)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v590))) = v643&v639 | v654<<(uint(v633)%64)
	goto L145
L139:
	;
	v597 = int32(4800)
	goto L141
L140:
	;
	v597 = int32(4799)
	goto L141
L141:
	;
	v598 = v597 + v572
	v603 = base.I32_div_s(v598, int32(4))
	v606 = base.I32_div_s(v598, int32(-100))
	v609 = base.I32_div_s(v598, int32(400))
	if int32(2) < v588 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v613 = int32(1)
	goto L144
L143:
	;
	v613 = int32(13)
	goto L144
L144:
	;
	v618 = base.I32_div_s((v613+v588)*int32(7834), int32(256))
	goto L138
L145:
	;
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	if v665 != v666>>(uint(int64(63))%64) {
		goto L126
	} else {
		goto L146
	}
L146:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	v674 = int32(60)
	v683 = base.I64_extend_i32_s(v569) + base.I64_extend_i32_s(v671+(v672+v673*v674)*v674)*int64(1000000)
	v686 = v683 + v666
	if base.B2i32(v683 < int64(0))^base.B2i32(v686 < v666) != 0 {
		goto L126
	} else {
		goto L147
	}
L147:
	;
	if base.Ui64(int64(9011559254509551615)) < base.Ui64(v686-int64(9223371331200000000)) {
		v740 = v686
		goto L4
	} else {
		goto L148
	}
L148:
	;
	goto L126
L149:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(422473), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(521224), int32(4823), int32(513618))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v720 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v58
	F_errmsg(m, int32(201047), v12+int32(48))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(521224), int32(4830), int32(513618))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	m.G0 = v12 + int32(112)
	return v742
L159:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errmsg(m, int32(422473), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(521224), int32(4733), int32(513618))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
