package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeTimezoneName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_strlen(m, l0)
	mBase = m.M
	v14 = F_downcase_truncate_identifier(m, l0, v12, v4)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v21 = v9 + int32(4)
		v22 = F_DecodeTimezoneAbbrev(m, v4, v14, v9+int32(12), l1, l2, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 != 0 {
				v24 = int32(0)
				F_DateTimeParseError(m, v22, v21, v24, v24, v24)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					if base.Ui32(v29-int32(5)) < base.Ui32(int32(2)) {
						v44 = v4
						m.G0 = v9 + int32(16)
						return v44
					} else {
						if v29 == int32(7) {
							v44 = int32(1)
							m.G0 = v9 + int32(16)
							return v44
						} else {
							v37 = F_pg_tzset(m, l0)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
								if v37 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
											F_errmsg(m, int32(_a_F_DecodeTimezoneName_0), v9)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_DecodeTimezoneName_1), int32(3330), int32(_a_F_DecodeTimezoneName_2))
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
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
									v44 = int32(2)
									m.G0 = v9 + int32(16)
									return v44
								}
							}
						}
					}
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				if base.Ui32(v29-int32(5)) < base.Ui32(int32(2)) {
					v44 = v4
					m.G0 = v9 + int32(16)
					return v44
				} else {
					if v29 == int32(7) {
						v44 = int32(1)
						m.G0 = v9 + int32(16)
						return v44
					} else {
						v37 = F_pg_tzset(m, l0)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
							if v37 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
										F_errmsg(m, int32(_a_F_DecodeTimezoneName_0), v9)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_DecodeTimezoneName_1), int32(3330), int32(_a_F_DecodeTimezoneName_2))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
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
								v44 = int32(2)
								m.G0 = v9 + int32(16)
								return v44
							}
						}
					}
				}
			}
		}
	}
}
func F_check_timezone(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 float64
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = v11
	v17 = int32(_a_F_check_timezone_0)
	v18 = int32(8)
	goto L2
L1:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v63 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v63 = int32(0)
	goto L1
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 == v22 {
		v44 = v21
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v46 = int32(1)
	if v44 != 0 {
		v16 = v16 + v46
		v17 = v17 + v46
		v18 = v18 - v46
		goto L2
	} else {
		goto L16
	}
L8:
	;
	if base.Ui32((v21-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v32 = v21 | int32(32)
	goto L11
L10:
	;
	v32 = v21
	goto L11
L11:
	;
	if base.Ui32((v22-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = v22 | int32(32)
	goto L14
L13:
	;
	v41 = v22
	goto L14
L14:
	;
	if v32 == v41 {
		v44 = v32
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v63 = v32 - v41
	goto L1
L16:
	;
	goto L6
L17:
	;
	m.G0 = v9 + int32(16)
	return v227
L18:
	;
	v217 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L30
	} else {
		goto L73
	}
L19:
	;
	if v195 != 0 {
		v212 = v195
		goto L18
	} else {
		goto L70
	}
L20:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
	v187 = base.I64_div_s(v185, int64(-1000000))
	v189 = F_pg_tzset_offset(m, base.I32_wrap_i64(v187))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L30
	} else {
		goto L68
	}
L21:
	;
	v69 = v64 + int32(8)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v145 = F_strtod(m, v64, v9+int32(12))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L30
	} else {
		goto L53
	}
L24:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if base.Ui32(v75-int32(9)) < base.Ui32(int32(5)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v88 = F_pstrdup(m, v69+int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	goto L25
L27:
	;
	v69 = v69 + int32(1)
	goto L24
L28:
	;
	v80 = int32(0)
	switch v75 - int32(32) {
	case 0:
		goto L27
	default:
		v227 = v80
		goto L17
	case 7:
		goto L26
	}
L29:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v105)
	v111 = F_DirectFunctionCall3Coll(m, int32(585), v105, v88, v105, int32(-1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L30
	} else {
		goto L41
	}
L30:
	;
	return int32(0)
L31:
	;
	v92 = int32(39)
	v93 = F___strchrnul(m, v88, v92)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v95 == v92 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v99 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v99 = v93
	goto L35
L34:
	;
	v99 = int32(0)
	goto L35
L35:
	;
	goto L32
L36:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v100 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v88)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L30
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v227 = v80
	goto L17
L41:
	;
	F_pfree(m, v88)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	if v115 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v117
	goto L46
L44:
	;
	goto L45
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	if v128 == int32(0) {
		goto L20
	} else {
		goto L49
	}
L46:
	;
	v123 = F_format_elog_string(m, int32(_a_F_check_timezone_1), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v123
	F_pfree(m, v111)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v227 = v80
	goto L17
L49:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v132
	goto L50
L50:
	;
	v138 = F_format_elog_string(m, int32(_a_F_check_timezone_2), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L30
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v138
	F_pfree(m, v111)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L30
	} else {
		goto L52
	}
L52:
	;
	v227 = v80
	goto L17
L53:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v147 == v148 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v156 = F_pg_tzset(m, v148)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L30
	} else {
		goto L58
	}
L55:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v150 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v154 = F_pg_tzset_offset(m, base.I32_trunc_sat_f64_s(base.F64_mul(v145, float64(-3600))))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L30
	} else {
		goto L57
	}
L57:
	;
	v195 = v154
	goto L19
L58:
	;
	if v156 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v227 = int32(0)
	goto L17
L60:
	;
	goto L61
L61:
	;
	v161 = F_pg_tz_acceptable(m, v156)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	if v161 != 0 {
		v212 = v156
		goto L18
	} else {
		goto L63
	}
L63:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v165
	goto L64
L64:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v168
	v172 = F_format_elog_string(m, int32(_a_F_check_timezone_3), v9)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L30
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[3])) = v172
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v176
	goto L66
L66:
	;
	v182 = F_format_elog_string(m, int32(_a_F_check_timezone_4), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L30
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v182
	v227 = int32(0)
	goto L17
L68:
	;
	F_pfree(m, v111)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	v195 = v189
	goto L19
L70:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v201
	goto L71
L71:
	;
	v207 = F_format_elog_string(m, int32(_a_F_check_timezone_5), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L30
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v207
	v227 = int32(0)
	goto L17
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v217
	if v217 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v227 = int32(0)
	goto L17
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v212
	v227 = int32(1)
	goto L17
}
