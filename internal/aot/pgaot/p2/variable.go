package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExtractSetVariableArgs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v3 {
	case 0:
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v6 = F_flatten_set_variable_args(m, v4, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	default:
		v16 = int32(0)
		return v16
	case 2:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = int32(0)
		v14 = F_GetConfigOptionByName(m, v11, v12, v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = v14
			return v16
		}
	}
}
func F_flatten_set_variable_args(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L56
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L7
	} else {
		goto L52
	}
L3:
	;
	v15 = F_find_option(m, l0, int32(0), int32(1), int32(19))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v156 = int32(0)
	goto L5
L5:
	;
	m.G0 = v10 + int32(96)
	return v156
L6:
	;
	F_initStringInfo(m, v10+int32(80))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L14
	}
L7:
	;
	return int32(0)
L8:
	;
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v19&int32(1) != 0 {
		v26 = v19
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v22 = int32(0)
	goto L11
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v23 != int32(1) {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	v22 = v19
	goto L11
L13:
	;
	v26 = v22
	goto L6
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v31 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v40 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v156 = v147
	goto L5
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32))))
	if v40&int32(1073741823) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	F_appendStringInfoString(m, v10+int32(80), int32(719931))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v57 == int32(73) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = v61
	v64 = v60
	v65 = v62
	goto L26
L25:
	;
	v63 = v48
	v64 = int32(0)
	v65 = v57
	goto L26
L26:
	;
	if v65 != int32(72) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	switch v68 - int32(465) {
	case 0:
		goto L29
	case 1:
		goto L32
	default:
		goto L30
	case 3:
		goto L31
	}
L28:
	;
	v137 = v40 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v137 < v138 {
		v40 = v137
		goto L18
	} else {
		goto L51
	}
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v126
	F_appendStringInfo(m, v10+int32(80), int32(479600), v10+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L50
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L47
	}
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v64 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	F_appendStringInfoString(m, v10+int32(80), v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	F_typenameTypeIdAndMod(m, int32(0), v64, v10+int32(76), v10+int32(72))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v26&int32(2) != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v85 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	v90 = F_DirectFunctionCall3Coll(m, int32(584), v85, v76, v85, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v92 = F_DirectFunctionCall1Coll(m, int32(1288), v85, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v92
	F_appendStringInfo(m, v10+int32(80), int32(661176), v10+int32(32))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	goto L28
L41:
	;
	v104 = F_quote_identifier(m, v76)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_appendStringInfoString(m, v10+int32(80), v76)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L46
	}
L44:
	;
	F_appendStringInfoString(m, v10+int32(80), v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L28
L46:
	;
	goto L28
L47:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v116
	F_errmsg_internal(m, int32(477341), v10)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(486063), int32(300), int32(152710))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	goto L28
L51:
	;
	goto L19
L52:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l0
	F_errmsg(m, int32(92627), v10-int32(-64))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(486063), int32(218), int32(152710))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v183
	F_errmsg_internal(m, int32(477341), v10+int32(48))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(486063), int32(246), int32(152710))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_variable_numdistinct(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float32
	_ = v12
	var v14 float32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v51 float64
	_ = v51
	var v59 float64
	_ = v59
	var v63 float64
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 float64
	_ = v72
	var v75 int32
	_ = v75
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v92 float64
	_ = v92
	var v96 float64
	_ = v96
	var v101 float64
	_ = v101
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	v3 = int32(0)
	v4 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v3)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
		v11 = v9 + v10
		v12 = *(*float32)(unsafe.Add(mBase, uint32(v11)+8))
		v14 = *(*float32)(unsafe.Add(mBase, uint32(v11)+16))
		v41 = base.F64_promote_f32(v14)
		v42 = base.F64_promote_f32(v12)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v16 == int32(16) {
			v41 = float64(2)
			v42 = v4
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v20 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v27 == int32(0) {
					v41 = float64(0)
					v42 = v4
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					if v30 != int32(6) {
						v41 = float64(0)
						v42 = v4
					} else {
						v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
						switch v34 - int32(65530) {
						case 0:
							v41 = float64(1)
							v42 = v4
						default:
							v41 = float64(0)
							v42 = v4
						case 5:
							v41 = float64(-1)
							v42 = v4
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
				if v23 != int32(5) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v27 == int32(0) {
						v41 = float64(0)
						v42 = v4
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						if v30 != int32(6) {
							v41 = float64(0)
							v42 = v4
						} else {
							v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
							switch v34 - int32(65530) {
							case 0:
								v41 = float64(1)
								v42 = v4
							default:
								v41 = float64(0)
								v42 = v4
							case 5:
								v41 = float64(-1)
								v42 = v4
							}
						}
					}
				} else {
					v41 = float64(-1)
					v42 = v4
				}
			}
		}
	}
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v46 != 0 {
		v47 = base.F64_neg(base.F64_sub(float64(1), v42))
	} else {
		v47 = v41
	}
	if base.F64_gt(v47, float64(0)) != 0 {
		v51 = float64(1e+100)
		if base.F64_gt(v47, v51) != 0 {
			v63 = v51
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v47)&int64(9223372036854775807)) {
				v63 = v51
			} else {
				v59 = float64(1)
				if base.F64_le(v47, v59) != 0 {
					v63 = v59
				} else {
					v63 = base.F64_nearest(v47)
				}
			}
		}
		return v63
	} else {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v65 == int32(0) {
			v68 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v68)
			return float64(200)
		} else {
			v72 = *(*float64)(unsafe.Add(mBase, uint32(v65)+120))
			if base.F64_le(v72, float64(0)) != 0 {
				v75 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v75)
				return float64(200)
			} else {
				if base.F64_lt(v47, float64(0)) != 0 {
					v82 = base.F64_mul(v72, base.F64_neg(v47))
					v84 = float64(1e+100)
					if base.F64_gt(v82, v84) != 0 {
						v96 = v84
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v82)&int64(9223372036854775807)) {
							v96 = v84
						} else {
							v92 = float64(1)
							if base.F64_le(v82, v92) != 0 {
								v96 = v92
							} else {
								v96 = base.F64_nearest(v82)
							}
						}
					}
					return v96
				} else {
					if base.F64_lt(v72, float64(200)) != 0 {
						v101 = float64(1e+100)
						if base.F64_gt(v72, v101) != 0 {
							v113 = v101
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v72)&int64(9223372036854775807)) {
								v113 = v101
							} else {
								v109 = float64(1)
								if base.F64_le(v72, v109) != 0 {
									v113 = v109
								} else {
									v113 = base.F64_nearest(v72)
								}
							}
						}
						return v113
					} else {
						v115 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v115)
						return float64(200)
					}
				}
			}
		}
	}
}
