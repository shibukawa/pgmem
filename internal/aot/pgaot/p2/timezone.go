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
												F_errfinish(m, int32(_a_F_DecodeTimezoneName_1), int32(3351), int32(_a_F_DecodeTimezoneName_2))
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
											F_errfinish(m, int32(_a_F_DecodeTimezoneName_1), int32(3351), int32(_a_F_DecodeTimezoneName_2))
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 float64
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = v12
	v18 = int32(_a_F_check_timezone_0)
	v19 = int32(8)
	goto L2
L1:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v64 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v64 = int32(0)
	goto L1
L4:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 == v23 {
		v45 = v22
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
	v47 = int32(1)
	if v45 != 0 {
		v17 = v17 + v47
		v18 = v18 + v47
		v19 = v19 - v47
		goto L2
	} else {
		goto L16
	}
L8:
	;
	if base.Ui32((v22-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v33 = v22 | int32(32)
	goto L11
L10:
	;
	v33 = v22
	goto L11
L11:
	;
	if base.Ui32((v23-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v23 | int32(32)
	goto L14
L13:
	;
	v42 = v23
	goto L14
L14:
	;
	if v33 == v42 {
		v45 = v33
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v64 = v33 - v42
	goto L1
L16:
	;
	goto L6
L17:
	;
	m.G0 = v10 + int32(16)
	return v233
L18:
	;
	v223 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L30
	} else {
		goto L73
	}
L19:
	;
	if v199 != 0 {
		v217 = v199
		goto L18
	} else {
		goto L70
	}
L20:
	;
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	v191 = base.I64_div_s(v189, int64(-1000000))
	v193 = F_pg_tzset_offset(m, base.I32_wrap_i64(v191))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L30
	} else {
		goto L68
	}
L21:
	;
	v70 = v65 + int32(8)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v149 = F_strtod(m, v65, v10+int32(12))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L30
	} else {
		goto L53
	}
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if base.Ui32(v77-int32(9)) < base.Ui32(int32(5)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v90 = F_pstrdup(m, v70+int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	goto L25
L27:
	;
	v70 = v70 + int32(1)
	goto L24
L28:
	;
	v82 = int32(0)
	switch v77 - int32(32) {
	case 0:
		goto L27
	default:
		v233 = v82
		goto L17
	case 7:
		goto L26
	}
L29:
	;
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v107)
	v114 = F_DirectFunctionCall3Coll(m, int32(631), v107, base.I64_extend_i32_u(v90), int64(0), int64(-1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L30
	} else {
		goto L41
	}
L30:
	;
	return int32(0)
L31:
	;
	v94 = int32(39)
	v95 = F___strchrnul(m, v90, v94)
	mBase = m.M
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v97 == v94 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v101 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v101 = v95
	goto L35
L34:
	;
	v101 = int32(0)
	goto L35
L35:
	;
	goto L32
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v102 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v90)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L30
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v233 = v82
	goto L17
L41:
	;
	F_pfree(m, v90)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v118 = base.I32_wrap_i64(v114)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	if v119 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v121
	goto L46
L44:
	;
	goto L45
L45:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	if v132 == int32(0) {
		goto L20
	} else {
		goto L49
	}
L46:
	;
	v127 = F_format_elog_string(m, int32(_a_F_check_timezone_1), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v127
	F_pfree(m, v118)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v233 = v82
	goto L17
L49:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v136
	goto L50
L50:
	;
	v142 = F_format_elog_string(m, int32(_a_F_check_timezone_2), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L30
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v142
	F_pfree(m, v118)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L30
	} else {
		goto L52
	}
L52:
	;
	v233 = v82
	goto L17
L53:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v151 == v152 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v160 = F_pg_tzset(m, v152)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L30
	} else {
		goto L58
	}
L55:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v154 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v158 = F_pg_tzset_offset(m, base.I32_trunc_sat_f64_s(base.F64_mul(v149, float64(-3600))))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L30
	} else {
		goto L57
	}
L57:
	;
	v199 = v158
	goto L19
L58:
	;
	if v160 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v233 = int32(0)
	goto L17
L60:
	;
	goto L61
L61:
	;
	v165 = F_pg_tz_acceptable(m, v160)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	if v165 != 0 {
		v217 = v160
		goto L18
	} else {
		goto L63
	}
L63:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v169
	goto L64
L64:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v172
	v176 = F_format_elog_string(m, int32(_a_F_check_timezone_3), v10)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L30
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[3])) = v176
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v180
	goto L66
L66:
	;
	v186 = F_format_elog_string(m, int32(_a_F_check_timezone_4), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L30
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v186
	v233 = int32(0)
	goto L17
L68:
	;
	F_pfree(m, v118)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	v199 = v193
	goto L19
L70:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v206
	goto L71
L71:
	;
	v212 = F_format_elog_string(m, int32(_a_F_check_timezone_5), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L30
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v212
	v233 = int32(0)
	goto L17
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v223
	if v223 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v233 = int32(0)
	goto L17
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v217
	v233 = int32(1)
	goto L17
}
