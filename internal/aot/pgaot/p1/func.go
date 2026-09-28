package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LookupFuncName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v21 = F_FuncnameGetCandidates(m, l0, l1, v5, v5, v5, v5, l3, v13+int32(44))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v192
L2:
	;
	v175 = F_NameListToString(m, l0)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L55
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L49
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L43
	}
L5:
	;
	return int32(0)
L6:
	;
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = l1 << (uint(int32(2)) % 32)
	v34 = v21
	v35 = v5
	goto L10
L8:
	;
	goto L9
L9:
	;
	if l3 != 0 {
		v192 = v5
		goto L1
	} else {
		goto L42
	}
L10:
	;
	if base.B2i32(l1 <= int32(0)) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if l3 != 0 {
		v192 = v113
		goto L1
	} else {
		goto L40
	}
L12:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v115 != 0 {
		v34 = v115
		v35 = v113
		goto L10
	} else {
		goto L39
	}
L13:
	;
	v42 = v34 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v26) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	goto L15
L15:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v105 == int32(0) {
		goto L3
	} else {
		goto L35
	}
L16:
	;
	if v104 != 0 {
		v113 = v35
		goto L12
	} else {
		goto L34
	}
L17:
	;
	v104 = int32(0)
	goto L16
L18:
	;
	v78 = v73
	v79 = v74
	v80 = v75
	goto L28
L19:
	;
	if (l2|v42)&int32(3) != 0 {
		v73 = l2
		v74 = v42
		v75 = v26
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v66 = l2
	v67 = v42
	v68 = v26
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v50 = l2
	v51 = v42
	v52 = v26
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v55 != v56 {
		v73 = v50
		v74 = v51
		v75 = v52
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v66 = v61
	v67 = v59
	v68 = v63
	goto L21
L25:
	;
	v58 = int32(4)
	v59 = v51 + v58
	v61 = v50 + v58
	v63 = v52 - v58
	if base.Ui32(int32(3)) < base.Ui32(v63) {
		v50 = v61
		v51 = v59
		v52 = v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v73 = v66
	v74 = v67
	v75 = v68
	goto L18
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == v84 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v104 = v83 - v84
	goto L16
L30:
	;
	v86 = int32(1)
	v91 = v80 - v86
	if v91 != 0 {
		v78 = v78 + v86
		v79 = v79 + v86
		v80 = v91
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	goto L15
L35:
	;
	v108 = F_get_func_prokind(m, v105)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if v108 == int32(112) {
		v113 = v35
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v35 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v113 = v112
	goto L12
L39:
	;
	goto L11
L40:
	;
	if v113 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v192 = v113
	goto L1
L42:
	;
	goto L4
L43:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if l1 < int32(0) {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v138 = F_func_signature_string(m, l0, l1, int32(0), l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v138
	F_errmsg(m, int32(_a_F_LookupFuncName_0), v13+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_LookupFuncName_1), int32(2361), int32(_a_F_LookupFuncName_2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v158 = F_NameListToString(m, l0)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v158
	F_errmsg(m, int32(_a_F_LookupFuncName_3), v13+int32(32))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_errhint(m, int32(_a_F_LookupFuncName_4), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_LookupFuncName_1), int32(2370), int32(_a_F_LookupFuncName_2))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v175
	F_errmsg(m, int32(_a_F_LookupFuncName_5), v13)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_LookupFuncName_1), int32(2355), int32(_a_F_LookupFuncName_2))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_lookup_failure_details(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	if l0&int32(8) == int32(0) {
		if l0&int32(1) != 0 {
			return
		} else {
			if l0&int32(4) == int32(0) {
				if l2 != 0 {
					v16 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_0), int32(0))
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						return
					}
				} else {
					v20 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_1), int32(0))
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				if l2 != 0 {
					v24 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_2), int32(0))
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				} else {
					v28 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_3), int32(0))
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		if l0&int32(16) == int32(0) {
			if l2 != 0 {
				v36 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_4), int32(0))
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					return
				}
			} else {
				v40 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_5), int32(0))
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v42 = int32(0)
			if base.B2i32(l1 == v42)|l0&int32(32) == v42 {
				if l2 != 0 {
					v51 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_6), int32(0))
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						return
					}
				} else {
					v55 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_7), int32(0))
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v57 = int32(0)
				if base.B2i32(l1 == v57)|l0&int32(64) == v57 {
					v66 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_8), int32(0))
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						return
					}
				} else {
					v68 = int32(0)
					if base.B2i32(l1 == v68)|l0&int32(128) == v68 {
						v77 = F_errdetail(m, int32(_a_F_func_lookup_failure_details_9), int32(0))
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							return
						}
					} else {
						v79 = int32(0)
						if base.B2i32(l1 == v79)|l0&int32(256) == v79 {
							F_errhint(m, int32(_a_F_func_lookup_failure_details_10), int32(0))
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								return
							}
						} else {
							if l0&int32(512) != 0 {
								F_errhint(m, int32(_a_F_func_lookup_failure_details_11), int32(0))
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									return
								}
							} else {
								if l2 != 0 {
									v98 = int32(_a_F_func_lookup_failure_details_12)
								} else {
									v98 = int32(_a_F_func_lookup_failure_details_13)
								}
								v100 = F_errdetail(m, v98, int32(0))
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_func_lookup_failure_details_14), int32(0))
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_get_func_arg_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
	v21 = F_SysCacheGetAttr(m, int32(47), l0, int32(21), v13+int32(15))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L57
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L54
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L51
	}
L6:
	;
	v84 = F_SysCacheGetAttr(m, int32(47), l0, int32(23), v13+int32(15))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L7:
	;
	v29 = F_pg_detoast_datum(m, base.I32_wrap_i64(v21))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v63 = v15 + v16
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+128))
	v65 = F_palloc_mul(m, int32(4), v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v31 != int32(1) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v34 < int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if v37 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v38 != int32(26) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v42 = F_palloc_mul(m, int32(4), v34)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if v45 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v55 = (v48<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L17:
	;
	v55 = v45
	goto L18
L18:
	;
	v57 = v34 << (uint(int32(2)) % 32)
	if v57 == int32(0) {
		v76 = v34
		goto L6
	} else {
		goto L19
	}
L19:
	;
	base.MemoryCopy(m, v42, v55+v29, v57)
	v76 = v34
	goto L6
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v65
	v69 = v64 << (uint(int32(2)) % 32)
	if v69 == int32(0) {
		v76 = v64
		goto L6
	} else {
		goto L21
	}
L21:
	;
	base.MemoryCopy(m, v65, v63+int32(136), v69)
	v76 = v64
	goto L6
L22:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v86 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v150 = F_SysCacheGetAttr(m, int32(47), l0, int32(22), v13+int32(15))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v93 = F_pg_detoast_datum(m, base.I32_wrap_i64(v84))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_deconstruct_array_builtin(m, v93, int32(25), v13+int32(8), int32(0), v13+int32(4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v103 != v76 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v106 = F_palloc_mul(m, int32(4), v76)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v106
	if v76 <= int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v112 = int32(0)
	goto L32
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v112<<(uint(int32(3))%32))))
	v126 = F_text_to_cstring(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L23
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v128+v112<<(uint(int32(2))%32)))) = v126
	v134 = v112 + int32(1)
	if v134 != v76 {
		v112 = v134
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v152 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	m.G0 = v13 + int32(16)
	return v76
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v158 = F_pg_detoast_datum(m, base.I32_wrap_i64(v150))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v160 != int32(1) {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	if v163 != v76 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v165 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	if v166 != int32(18) {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v169 = F_palloc(m, v76)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v169
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v172 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v182 = (v175<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L49
L48:
	;
	v182 = v172
	goto L49
L49:
	;
	if v76 == int32(0) {
		goto L37
	} else {
		goto L50
	}
L50:
	;
	base.MemoryCopy(m, v169, v158+v182, v76)
	goto L37
L51:
	;
	F_errmsg_internal(m, int32(_a_F_get_func_arg_info_0), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_get_func_arg_info_1), int32(1414), int32(_a_F_get_func_arg_info_2))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errmsg_internal(m, int32(_a_F_get_func_arg_info_3), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_get_func_arg_info_1), int32(1441), int32(_a_F_get_func_arg_info_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v76
	F_errmsg_internal(m, int32(_a_F_get_func_arg_info_4), v13)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_get_func_arg_info_1), int32(1461), int32(_a_F_get_func_arg_info_2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_func_namespace(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14289(m, l0, int32(47))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_func_support(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14287(m, l0, int32(47))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
