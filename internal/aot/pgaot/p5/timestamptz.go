package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_make_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+104))
	v8 = F_make_timestamp_internal(m, v2, v3, v4, v5, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = F_timestamp2timestamptz_safe(m, v8, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			return v13
		}
	}
}
func F_timestamptz_le_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_le_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_le_date[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(-1)
					} else {
						v32 = int32(1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(1)
				} else {
					v37 = int32(-1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(base.B2i32(int32(0) <= v41))
	}
}
func F_timestamptz_mi_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_interval_um_internal(m, v9, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v15 = F_timestamptz_pl_interval_internal(m, v8, v6, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 + int32(16)
			return v15
		}
	}
}
func F_timestamptz_pl_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_timestamptz_pl_interval_internal(m, v2, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_timestamptz_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_timestamptz2timestamp_safe(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(0) {
			v22 = v6
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(453) {
				v22 = v6
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)))
				if v16 != int32(1) {
					v22 = v6
				} else {
					v19 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
					v22 = int64(0)
				}
			}
		}
		return v22
	}
}
func F_timestamptz_to_char(m *base.Module, l0 int32) int64 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int64
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v17 == int32(1) {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
			if v23 == int32(18) {
				v26 = int32(16)
			} else {
				v26 = int32(0)
			}
			if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v33 = int32(4)
			} else {
				v33 = v26
			}
			v46 = v33
		} else {
			v34 = int32(1)
			if v17&v34 != 0 {
				v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui64(int64(1)) < base.Ui64(v11-int64(9223372036854775807)) {
			v52 = v46
		} else {
			v52 = int32(0)
		}
		if v52 == int32(0) {
			v55 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
			v180 = int64(0)
			m.G0 = v9 + int32(96)
			return v180
		} else {
			v58 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = int64(4294967297)
			v77 = F_timestamp2tm(m, v11, v9+int32(44), v9, v9+int32(88), v9+int32(92), int32(0))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int64(0)
			} else {
				if v77 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v188 = m.ExcPending
					if v188 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v191 = m.ExcPending
						if v191 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamptz_to_char_0), int32(0))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamptz_to_char_1), int32(_a_F_timestamptz_to_char_2), int32(_a_F_timestamptz_to_char_3))
								mBase = m.M
								v200 = m.ExcPending
								if v200 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v86 = base.B2i32(int32(2) < v80)
					if int32(2) < v80 {
						v87 = int32(_a_F_timestamptz_to_char_4)
					} else {
						v87 = int32(_a_F_timestamptz_to_char_5)
					}
					v88 = v87 + v79
					v93 = base.I32_div_s(v88, int32(4))
					v96 = base.I32_div_s(v88, int32(-100))
					v99 = base.I32_div_s(v88, int32(400))
					if int32(2) < v80 {
						v103 = int32(1)
					} else {
						v103 = int32(13)
					}
					v108 = base.I32_div_s((v103+v80)*int32(_a_F_timestamptz_to_char_6), int32(256))
					v111 = v81 + v88*int32(365) + v93 + v96 + v99 + v108 - int32(_a_F_timestamptz_to_char_7)
					v112 = int32(1)
					v115 = base.I32_rem_s(v111+v112, int32(7))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v115
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v126 = int32(_a_F_timestamptz_to_char_5) + v117
					v131 = base.I32_div_s(v126, int32(4))
					v134 = base.I32_div_s(v126, int32(-100))
					v137 = base.I32_div_s(v126, int32(400))
					v146 = base.I32_div_s(int32(_a_F_timestamptz_to_char_8), int32(256))
					v150 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v150
					v152 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+8)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v152
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v154
					v156 = *(*int64)(unsafe.Add(mBase, uint32(v9)+20))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v156
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v158
					v162 = v111 - (v112 + v126*int32(365) + v131 + v134 + v137 + v146 - int32(_a_F_timestamptz_to_char_7)) + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v162
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v162
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v169 = F_datetime_to_char_body(m, v9+int32(48), v13, int32(0), v168)
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return int64(0)
					} else {
						if v169 == int32(0) {
							v173 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v173)
							v180 = int64(0)
						} else {
							v180 = base.I64_extend_i32_u(v169)
						}
						m.G0 = v9 + int32(96)
						return v180
					}
				}
			}
		}
	}
}
func F_timestamptz_to_time_t(m *base.Module, l0 int64) int64 {
	var v3 int64
	_ = v3
	v3 = base.I64_div_s(l0, int64(1000000))
	return v3 + int64(946684800)
}
