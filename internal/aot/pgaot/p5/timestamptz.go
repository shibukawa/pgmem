package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_make_timestamptz(m *base.Module, l0 int32) int32 {
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
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = F_make_timestamp_internal(m, v2, v3, v4, v5, v6, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = F_timestamp2timestamptz_opt_overflow(m, v9, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_Int64GetDatum(m, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func F_timestamptz_le_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(_a_F_timestamptz_le_date_0), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_le_date[0]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return int32(base.Ui32(v65^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_timestamptz_mi_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_interval_um_internal(m, v10, v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = F_timestamptz_pl_interval_internal(m, v9, v6, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_Int64GetDatum(m, v16)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v18
			}
		}
	}
}
func F_timestamptz_pl_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_timestamptz_pl_interval_internal(m, v3, v4, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_timestamptz_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = F_timestamptz2timestamp(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_Int64GetDatum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_timestamptz_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
		if v18 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v47 = v34
		} else {
			v35 = int32(1)
			if v18&v35 != 0 {
				v47 = int32(base.Ui32(v18)>>(uint(v35)%32)) - v35
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui64(int64(1)) < base.Ui64(v12-int64(9223372036854775807)) {
			v53 = v47
		} else {
			v53 = int32(0)
		}
		if v53 == int32(0) {
			v172 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
			v176 = int32(0)
			m.G0 = v9 + int32(96)
			return v176
		} else {
			v56 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = int64(4294967297)
			v75 = F_timestamp2tm(m, v12, v9+int32(44), v9, v9+int32(88), v9+int32(92), int32(0))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				if v75 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v188 = m.ExcPending
						if v188 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamptz_to_char_0), int32(0))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamptz_to_char_1), int32(4053), int32(_a_F_timestamptz_to_char_2))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v84 = base.B2i32(int32(2) < v78)
					if int32(2) < v78 {
						v85 = int32(_a_F_timestamptz_to_char_3)
					} else {
						v85 = int32(_a_F_timestamptz_to_char_4)
					}
					v86 = v85 + v77
					v91 = base.I32_div_s(v86, int32(4))
					v94 = base.I32_div_s(v86, int32(-100))
					v97 = base.I32_div_s(v86, int32(400))
					if int32(2) < v78 {
						v101 = int32(1)
					} else {
						v101 = int32(13)
					}
					v106 = base.I32_div_s((v101+v78)*int32(_a_F_timestamptz_to_char_5), int32(256))
					v109 = v79 + v86*int32(365) + v91 + v94 + v97 + v106 - int32(_a_F_timestamptz_to_char_6)
					v110 = int32(1)
					v113 = base.I32_rem_s(v109+v110, int32(7))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v113
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v124 = int32(_a_F_timestamptz_to_char_4) + v115
					v129 = base.I32_div_s(v124, int32(4))
					v132 = base.I32_div_s(v124, int32(-100))
					v135 = base.I32_div_s(v124, int32(400))
					v144 = base.I32_div_s(int32(_a_F_timestamptz_to_char_7), int32(256))
					v148 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v148
					v150 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+8)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v150
					v152 = *(*int64)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v152
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v9)+20))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v154
					v156 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v156
					v160 = v109 - (v110 + v124*int32(365) + v129 + v132 + v135 + v144 - int32(_a_F_timestamptz_to_char_6)) + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v160
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v160
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v167 = F_datetime_to_char_body(m, v9+int32(48), v14, int32(0), v166)
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						if v167 != 0 {
							v176 = v167
						} else {
							v172 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
							v176 = int32(0)
						}
						m.G0 = v9 + int32(96)
						return v176
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
