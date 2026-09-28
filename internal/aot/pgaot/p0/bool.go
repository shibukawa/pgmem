package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bool_accum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v11 != 0 {
			v57 = v11
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v58 != 0 {
			} else {
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
				*(*int64)(unsafe.Add(mBase, uint32(v57))) = v59 + int64(1)
				v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				if v63 == int64(0) {
				} else {
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v66 + int64(1)
				}
			}
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v57)
		} else {
			v14 = v6 + int32(12)
			v15 = int32(0)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v16 == v15 {
				v33 = int32(0)
				if v14 == v33 {
					v41 = v33
				} else {
					v36 = v33
					v37 = v15
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
					v41 = v37
				}
				v44 = v41
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				switch v19 - int32(435) {
				case 0:
					if v14 == int32(0) {
						v44 = int32(1)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
						v36 = v26
						v37 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
						v41 = v37
						v44 = v41
					}
				case 1:
					if v14 == int32(0) {
						v44 = int32(2)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+376))
						v36 = v31
						v37 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
						v41 = v37
						v44 = v41
					}
				default:
					v33 = int32(0)
					if v14 == v33 {
						v41 = v33
					} else {
						v36 = v33
						v37 = v15
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
						v41 = v37
					}
					v44 = v41
				}
			}
			if v44 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_bool_accum_0), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_bool_accum_1), int32(330), int32(_a_F_bool_accum_2))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v49 = F_MemoryContextAlloc(m, v47, int32(16))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					v53 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v53
					*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53
					v57 = v49
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v58 != 0 {
					} else {
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
						*(*int64)(unsafe.Add(mBase, uint32(v57))) = v59 + int64(1)
						v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						if v63 == int64(0) {
						} else {
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v66 + int64(1)
						}
					}
					m.G0 = v6 + int32(16)
					return base.I64_extend_i32_u(v57)
				}
			}
		}
	} else {
		v14 = v6 + int32(12)
		v15 = int32(0)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v16 == v15 {
			v33 = int32(0)
			if v14 == v33 {
				v41 = v33
			} else {
				v36 = v33
				v37 = v15
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
			}
			v44 = v41
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			switch v19 - int32(435) {
			case 0:
				if v14 == int32(0) {
					v44 = int32(1)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
					v36 = v26
					v37 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
					v41 = v37
					v44 = v41
				}
			case 1:
				if v14 == int32(0) {
					v44 = int32(2)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+376))
					v36 = v31
					v37 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
					v41 = v37
					v44 = v41
				}
			default:
				v33 = int32(0)
				if v14 == v33 {
					v41 = v33
				} else {
					v36 = v33
					v37 = v15
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
					v41 = v37
				}
				v44 = v41
			}
		}
		if v44 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_bool_accum_0), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_bool_accum_1), int32(330), int32(_a_F_bool_accum_2))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v49 = F_MemoryContextAlloc(m, v47, int32(16))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				v53 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v53
				*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53
				v57 = v49
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
				if v58 != 0 {
				} else {
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
					*(*int64)(unsafe.Add(mBase, uint32(v57))) = v59 + int64(1)
					v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					if v63 == int64(0) {
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v66 + int64(1)
					}
				}
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v57)
			}
		}
	}
}
func F_bool_alltrue(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 != 0 {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		return int64(0)
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int64(0)
		} else {
			v8 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int64(0) {
				v17 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
				return base.I64_extend_i32_u(base.B2i32(v17 == v8))
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int64(0)
			}
		}
	}
}
func F_executeBoolItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	F_check_stack_depth(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L69
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v17 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v20 - int32(4) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L9
	case 4, 5, 6, 7, 8, 9:
		goto L14
	default:
		goto L10
	case 26:
		goto L11
	case 37:
		goto L13
	case 38:
		goto L12
	}
L7:
	;
	goto L6
L8:
	;
	m.G0 = v9 + int32(128)
	return v199
L9:
	;
	v188 = v9 + int32(100)
	F_jspGetArg(m, l1, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L67
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L64
	}
L11:
	;
	v114 = v9 + int32(100)
	F_jspGetArg(m, l1, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L45
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(0)
	v101 = v9 + int32(100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v101, v102, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L43
	}
L13:
	;
	v86 = v9 + int32(100)
	F_jspGetArg(m, l1, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L40
	}
L14:
	;
	v74 = v9 + int32(100)
	F_jspGetArg(m, l1, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L37
	}
L15:
	;
	v62 = v9 + int32(100)
	F_jspGetArg(m, l1, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L34
	}
L16:
	;
	v44 = v9 + int32(100)
	F_jspGetArg(m, l1, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L26
	}
L17:
	;
	v24 = v9 + int32(100)
	F_jspGetArg(m, l1, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v27 = int32(0)
	v29 = F_executeBoolItem(m, l0, v24, l2, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v29 == int32(0) {
		v199 = v27
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v34 = v9 + int32(16)
	F_jspGetRightArg(m, l1, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v38 = F_executeBoolItem(m, l0, v34, l2, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v38 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v42 = v29
	goto L25
L24:
	;
	v42 = v38
	goto L25
L25:
	;
	v199 = v42
	goto L8
L26:
	;
	v49 = F_executeBoolItem(m, l0, v44, l2, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v49 == int32(1) {
		v199 = int32(1)
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v54 = v9 + int32(16)
	F_jspGetRightArg(m, l1, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v58 = F_executeBoolItem(m, l0, v54, l2, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v58 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v60 = v58
	goto L33
L32:
	;
	v60 = v49
	goto L33
L33:
	;
	v199 = v60
	goto L8
L34:
	;
	v67 = F_executeBoolItem(m, l0, v62, l2, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v67 == int32(2) {
		v199 = int32(2)
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v199 = base.B2i32(v67 != int32(1))
	goto L8
L37:
	;
	v78 = v9 + int32(16)
	F_jspGetRightArg(m, l1, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v83 = F_executePredicate(m, l0, l1, v74, v78, l2, int32(1), int32(1571), l0)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v199 = v83
	goto L8
L40:
	;
	v90 = v9 + int32(16)
	F_jspGetRightArg(m, l1, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v93 = int32(0)
	v96 = F_executePredicate(m, l0, l1, v86, v90, l2, v93, int32(1572), v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v199 = v96
	goto L8
L43:
	;
	v106 = int32(0)
	v111 = F_executePredicate(m, l0, l1, v101, v106, l2, v106, int32(1573), v9+int32(16))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v199 = v111
	goto L8
L45:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v117 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v120
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(8589934592)
	v125 = v9 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v120)
	v131 = F_executeItemOptUnwrapTarget(m, l0, v114, l2, v125, v120)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	v158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v158)
	v164 = F_executeItemOptUnwrapTarget(m, l0, v9+int32(100), l2, v158, int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L60
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v127)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v135 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v137 = v135
	goto L53
L51:
	;
	goto L52
L52:
	;
	v151 = int32(2)
	if v131 == v151 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	F_pfree(m, v137)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L52
L55:
	;
	if v142 != 0 {
		v137 = v142
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v156 = v151
	goto L59
L58:
	;
	v156 = base.B2i32(v134 != int32(0))
	goto L59
L59:
	;
	v199 = v156
	goto L8
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v157)
	v167 = int32(2)
	if v164 == v167 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v172 = v167
	goto L63
L62:
	;
	v172 = base.B2i32(v164 == int32(0))
	goto L63
L63:
	;
	v199 = v172
	goto L8
L64:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v177
	F_errmsg_internal(m, int32(_a_F_executeBoolItem_0), v9)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_executeBoolItem_1), int32(1945), int32(_a_F_executeBoolItem_2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v192 = F_executeBoolItem(m, l0, v188, l2, int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v199 = base.B2i32(v192 == int32(2))
	goto L8
L69:
	;
	F_errmsg_internal(m, int32(_a_F_executeBoolItem_3), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_executeBoolItem_1), int32(1824), int32(_a_F_executeBoolItem_2))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
