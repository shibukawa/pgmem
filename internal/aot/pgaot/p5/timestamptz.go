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
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[329]))
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
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
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v48 = v35
		} else {
			v36 = int32(1)
			if v18&v36 != 0 {
				v48 = int32(base.Ui32(v18)>>(uint(v36)%32)) - v36
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui64(int64(1)) < base.Ui64(v12-int64(9223372036854775807)) {
			v54 = v48
		} else {
			v54 = int32(0)
		}
		if v54 == int32(0) {
			v57 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
			v181 = int32(0)
			m.G0 = v9 + int32(96)
			return v181
		} else {
			v59 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v9-int32(-64)))) = int64(4294967297)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v59
			v80 = F_timestamp2tm(m, v12, v9+int32(44), v9, v9+int32(88), v9+int32(92), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				if v80 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(401885), int32(0))
							mBase = m.M
							v197 = m.ExcPending
							if v197 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(497568), int32(4053), int32(229911))
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
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
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v89 = base.B2i32(int32(2) < v83)
					if int32(2) < v83 {
						v90 = int32(4800)
					} else {
						v90 = int32(4799)
					}
					v91 = v90 + v82
					v96 = base.I32_div_s(v91, int32(4))
					v99 = base.I32_div_s(v91, int32(-100))
					v102 = base.I32_div_s(v91, int32(400))
					if int32(2) < v83 {
						v106 = int32(1)
					} else {
						v106 = int32(13)
					}
					v111 = base.I32_div_s((v106+v83)*int32(7834), int32(256))
					v114 = v84 + v91*int32(365) + v96 + v99 + v102 + v111 - int32(32167)
					v115 = int32(1)
					v118 = base.I32_rem_s(v114+v115, int32(7))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v118
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v129 = int32(4799) + v120
					v134 = base.I32_div_s(v129, int32(4))
					v137 = base.I32_div_s(v129, int32(-100))
					v140 = base.I32_div_s(v129, int32(400))
					v149 = base.I32_div_s(int32(109676), int32(256))
					v153 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v153
					v155 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+8)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v155
					v157 = *(*int64)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v157
					v159 = *(*int64)(unsafe.Add(mBase, uint32(v9)+20))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v159
					v161 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v161
					v165 = v114 - (v115 + v129*int32(365) + v134 + v137 + v140 + v149 - int32(32167)) + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v165
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v165
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v172 = F_datetime_to_char_body(m, v9+int32(48), v14, int32(0), v171)
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						if v172 != 0 {
							v181 = v172
						} else {
							v174 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v174)
							v181 = int32(0)
						}
						m.G0 = v9 + int32(96)
						return v181
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
