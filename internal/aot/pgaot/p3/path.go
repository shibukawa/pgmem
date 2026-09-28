package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_GetSearchPathMatcher(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	F_recomputeNamespacePath(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = int32(_a_F_GetSearchPathMatcher_0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetSearchPathMatcher[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetSearchPathMatcher[0])) = l0
	v14 = F_palloc0(m, int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetSearchPathMatcher[1]))
	v18 = F_list_copy(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v43
	v47 = *(*int64)(unsafe.Add(mBase, _c_F_GetSearchPathMatcher[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v47
	*(*int32)(unsafe.Add(mBase, _c_F_GetSearchPathMatcher[0])) = v10
	return v14
L5:
	;
	if v18 == int32(0) {
		v43 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = v18
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetSearchPathMatcher[3]))
	if v27 == v29 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v43 = int32(0)
	goto L4
L9:
	;
	v43 = v22
	goto L4
L10:
	;
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_GetSearchPathMatcher[4]))
	if v32 == v27 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v38 = F_list_delete_first(m, v22)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(v34)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v36)
	goto L12
L16:
	;
	if v38 != 0 {
		v22 = v38
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L8
}
func F_path_add(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v113 int32
	_ = v113
	var v117 float64
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			if v17 == int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				if v20 == int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v29 = v27 + v28
					if base.Ui32(int32(268435455)) < base.Ui32(v29) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_path_add_0), int32(0))
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_path_add_1), int32(_a_F_path_add_2), int32(_a_F_path_add_3))
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
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
						v33 = v29 << (uint(int32(4)) % 32)
						if v33 == int32(2147483632) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_path_add_0), int32(0))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_path_add_1), int32(_a_F_path_add_2), int32(_a_F_path_add_3))
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
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
							v37 = v33 + int32(16)
							v38 = F_palloc(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v38))) = v37 << (uint(int32(2)) % 32)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v43 + v44
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v48 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v48
								*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v47
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
								if v48 < v52 {
									v55 = int32(16)
									v59 = v48
									for {
										v68 = v59 << (uint(int32(4)) % 32)
										v69 = v38 + v55 + v68
										v70 = v68 + (v10 + v55)
										v71 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
										*(*float64)(unsafe.Add(mBase, uint32(v69))) = v71
										v73 = *(*float64)(unsafe.Add(mBase, uint32(v70)+8))
										*(*float64)(unsafe.Add(mBase, uint32(v69)+8)) = v73
										v76 = v59 + int32(1)
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										if v76 < v77 {
											v59 = v76
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								if int32(0) < v87 {
									v90 = int32(16)
									v95 = int32(0)
									for {
										v103 = int32(4)
										v104 = v95 << (uint(v103) % 32)
										v105 = v38 + v90 + v104
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										v110 = v104 + (v15 + v90)
										v111 = *(*float64)(unsafe.Add(mBase, uint32(v110)))
										*(*float64)(unsafe.Add(mBase, uint32(v105+v106<<(uint(v103)%32)))) = v111
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										v117 = *(*float64)(unsafe.Add(mBase, uint32(v110)+8))
										*(*float64)(unsafe.Add(mBase, uint32(v105+v113<<(uint(v103)%32))+8)) = v117
										v120 = v95 + int32(1)
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
										if v120 < v121 {
											v95 = v120
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								return base.I64_extend_i32_u(v38)
							}
						}
					}
				} else {
					v23 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
					return int64(0)
				}
			} else {
				v23 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
				return int64(0)
			}
		}
	}
}
func F_path_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v285 int64
	_ = v285
	v2 = int32(0)
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = int32(44)
	v18 = F___strchrnul(m, v16, v17)
	mBase = m.M
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v20 == v17 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v285
L2:
	;
	v258 = F_errsave_start(m, v15)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L36
	} else {
		goto L67
	}
L3:
	;
	if v24 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	v24 = v18
	goto L6
L5:
	;
	v24 = v2
	goto L6
L6:
	;
	goto L3
L7:
	;
	v29 = v24
	v31 = v2
	goto L8
L8:
	;
	v37 = int32(1)
	v41 = int32(44)
	v42 = F___strchrnul(m, v29+v37, v41)
	mBase = m.M
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v44 == v41 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v52 = int32(1)
	if v31&v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if v48 != 0 {
		v29 = v48
		v31 = v31 + v37
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v48 = v42
	goto L13
L12:
	;
	v48 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L9
L15:
	;
	v56 = int32(-1)
	goto L17
L16:
	;
	v56 = (v31 + int32(2)) >> (uint(v52) % 32)
	goto L17
L17:
	;
	if v56 <= int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v61 = v16
	goto L21
L19:
	;
	v112 = v56 << (uint(int32(4)) % 32)
	v113 = base.I32_div_s(v112, v56)
	if base.B2i32(v113 == int32(16))&base.B2i32(v112 != int32(2147483632)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v106
	v109 = v106
	v110 = v107
	goto L19
L21:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if base.Ui32(v69-int32(9)) < base.Ui32(int32(5)) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v61
	v82 = F_strlen(m, v61)
	mBase = m.M
	v89 = v82 + int32(1)
	goto L28
L23:
	;
	goto L22
L24:
	;
	v61 = v61 + int32(1)
	goto L21
L25:
	;
	switch v69 - int32(32) {
	case 0:
		goto L24
	default:
		v106 = v61
		v107 = v2
		goto L20
	case 8:
		goto L23
	}
L26:
	;
	if v101 != v61 {
		v109 = v61
		v110 = v2
		goto L19
	} else {
		goto L32
	}
L27:
	;
	goto L26
L28:
	;
	v91 = int32(0)
	if v89 == v91 {
		v101 = v91
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v101 = v96
	goto L27
L30:
	;
	v95 = v89 - int32(1)
	v96 = v61 + v95
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97 != int32(40) {
		v89 = v95
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v103 = int32(1)
	v106 = v61 + v103
	v107 = v103
	goto L20
L33:
	;
	v121 = F_errsave_start(m, v15)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v140 = v112 + int32(16)
	v141 = F_palloc(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L36
	} else {
		goto L42
	}
L36:
	;
	return int64(0)
L37:
	;
	if v121 == int32(0) {
		v285 = v10
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_path_in_0), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	F_errsave_finish(m, v15, int32(_a_F_path_in_1), int32(1487), int32(_a_F_path_in_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v285 = v10
	goto L1
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v140 << (uint(int32(2)) % 32)
	v155 = F_path_decode(m, v109, int32(1), v56, v141+int32(16), v13+int32(47), v13+int32(40), int32(_a_F_path_in_3), v16, v15)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v155 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v159)
	v285 = v10
	goto L1
L45:
	;
	goto L46
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v110 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v241 ^ int32(1)
	v285 = base.I64_extend_i32_u(v141)
	goto L1
L48:
	;
	v221 = F_errsave_start(m, v15)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L36
	} else {
		goto L62
	}
L49:
	;
	if v162&int32(255) != int32(41) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	v191 = v162
	goto L51
L51:
	;
	if v191&int32(255) == int32(0) {
		goto L47
	} else {
		goto L56
	}
L52:
	;
	v169 = v161
	goto L53
L53:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if base.B2i32(base.Ui32(v179-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v179 == int32(32)) != 0 {
		v169 = v169 + int32(1)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v191 = v179
	goto L51
L55:
	;
	goto L54
L56:
	;
	v201 = F_errsave_start(m, v15)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L36
	} else {
		goto L57
	}
L57:
	;
	if v201 == int32(0) {
		v285 = v10
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L36
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_path_in_3)
	F_errmsg(m, int32(_a_F_path_in_4), v13+int32(16))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L36
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, v15, int32(_a_F_path_in_1), int32(1512), int32(_a_F_path_in_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L36
	} else {
		goto L61
	}
L61:
	;
	v285 = v10
	goto L1
L62:
	;
	if v221 == int32(0) {
		v285 = v10
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L36
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_path_in_3)
	F_errmsg(m, int32(_a_F_path_in_4), v13+int32(32))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L36
	} else {
		goto L65
	}
L65:
	;
	F_errsave_finish(m, v15, int32(_a_F_path_in_1), int32(1504), int32(_a_F_path_in_2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L36
	} else {
		goto L66
	}
L66:
	;
	v285 = v10
	goto L1
L67:
	;
	if v258 == int32(0) {
		v285 = v10
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L36
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_path_in_3)
	F_errmsg(m, int32(_a_F_path_in_4), v13)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L36
	} else {
		goto L70
	}
L70:
	;
	F_errsave_finish(m, v15, int32(_a_F_path_in_1), int32(1467), int32(_a_F_path_in_2))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L36
	} else {
		goto L71
	}
L71:
	;
	v285 = v10
	goto L1
}
func F_path_n_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			return base.I64_extend_i32_u(base.B2i32(v11 == v12))
		}
	}
}
func F_path_sub_pt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return base.I64_extend_i32_u(v11)
L6:
	;
	v32 = v11 + int32(16) + v24<<(uint(int32(4))%32)
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
	v35 = base.F64_sub(v33, v34)
	v37 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v35), v37)|base.F64_eq(base.F64_abs(v33), v37)|base.F64_eq(base.F64_abs(v34), v37) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v50 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v52 = v35
	goto L10
L10:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
	v55 = base.F64_sub(v53, v54)
	v57 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v55), v57)|base.F64_eq(base.F64_abs(v53), v57)|base.F64_eq(base.F64_abs(v54), v57) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v52 = v50
	goto L10
L12:
	;
	v70 = v55
	goto L14
L13:
	;
	v68 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = v70
	*(*float64)(unsafe.Add(mBase, uint32(v32))) = v52
	v74 = v24 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v74 < v75 {
		v24 = v74
		goto L6
	} else {
		goto L16
	}
L15:
	;
	v70 = v68
	goto L14
L16:
	;
	goto L7
}
func F_setPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
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
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v156 int32
	_ = v156
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v208 int32
	_ = v208
	var v235 int32
	_ = v235
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v323 int32
	_ = v323
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v398 int32
	_ = v398
	var v412 int32
	_ = v412
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v436 int32
	_ = v436
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v488 int32
	_ = v488
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v512 int32
	_ = v512
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v563 int32
	_ = v563
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v927 int32
	_ = v927
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v973 int32
	_ = v973
	var v987 int32
	_ = v987
	var v1001 int32
	_ = v1001
	var v1016 int32
	_ = v1016
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	v9 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(192)
	m.G0 = v28
	F_check_stack_depth(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v32 = l2 + l5
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v33 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L280
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L1
	} else {
		goto L275
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L270
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L266
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L261
	}
L8:
	;
	v39 = F_JsonbIteratorNext(m, l0, v28-int32(-64), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L257
	}
L11:
	;
	m.G0 = v28 + int32(192)
	return
L12:
	;
	if l7&int32(32) != 0 {
		goto L252
	} else {
		goto L253
	}
L13:
	;
	F_pushJsonbValue(m, l4, int32(6), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L121
	}
L14:
	;
	v44 = l7 & int32(32)
	v45 = int32(0)
	if base.B2i32(v44 == v45)|base.B2i32(l3-int32(1) < l5) == v45 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	switch v39 - int32(2) {
	case 0, 1:
		goto L12
	case 2:
		goto L14
	default:
		goto L3
	case 4:
		goto L13
	}
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+80)))
	if v53&int32(1) != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_pushJsonbValue(m, l4, int32(4), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v28)+72))
	v61 = base.B2i32(l3 <= l5)
	if l3 <= l5 {
		v81 = v60
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if int32(0) <= v81 {
		v114 = v81
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v62 != 0 {
		v81 = v60
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1+l5<<(uint(int32(3))%32))))
	v67 = F_text_to_cstring(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_setPath[0])) = int32(0)
	v75 = F_strtol(m, v67, v28+int32(160), int32(10))
	mBase = m.M
	goto L25
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v28)+160))
	if v67 == v76 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v78 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_setPath[0]))
	if v80 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v81 = v75
	goto L21
L29:
	;
	v116 = l7 & int32(25)
	v117 = int32(0)
	v120 = l3 - int32(1)
	v121 = base.B2i32(l5 != v120)
	if base.Ui32(v60) < base.Ui32(v114) {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	if base.Ui32(v60) < base.Ui32(int32(0)-v81) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(l7) < base.Ui32(int32(64)) {
		v114 = int32(-2147483648)
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v114 = v81 + v60
	goto L29
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l5 + int32(1)
	F_errmsg(m, int32(_a_F_setPath_0), v28+int32(32))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_2), int32(_a_F_setPath_3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v125 = v60
	goto L41
L40:
	;
	v125 = v114
	goto L41
L41:
	;
	if int32(0) < v114 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v128 = v125
	goto L44
L43:
	;
	v128 = v114
	goto L44
L44:
	;
	if v44 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v129 = v114
	goto L47
L46:
	;
	v129 = v128
	goto L47
L47:
	;
	v134 = base.B2i32(v116 == v117) | (v121 | base.B2i32(v60 != v117)&base.B2i32(v129 != int32(-2147483648)))
	if v134 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v137 = int32(0)
	if v60|base.B2i32(v44 == v137)|base.B2i32(v129 <= v137) == v137 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v235 = base.B2i32(v134 == int32(0))
	if v60 != 0 {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = int32(0)
	v156 = v129
	goto L54
L52:
	;
	goto L53
L53:
	;
	F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L58
	}
L54:
	;
	F_pushJsonbValue(m, l4, int32(3), v28+int32(160))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L53
L56:
	;
	v177 = int32(1)
	if base.Ui32(v177) < base.Ui32(v156) {
		v156 = v156 - v177
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L50
L59:
	;
	v254 = int32(0)
	v258 = v235
	goto L62
L60:
	;
	v398 = v235
	goto L61
L61:
	;
	v412 = int32(0)
	if v398&int32(1)|(v121|base.B2i32(v116 == v412)) == v412 {
		goto L99
	} else {
		goto L100
	}
L62:
	;
	if v61|base.B2i32(v254 != v129) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v398 = v370
	goto L61
L64:
	;
	v383 = v254 + int32(1)
	if v383 != v60 {
		v254 = v383
		v258 = v370
		goto L62
	} else {
		goto L97
	}
L65:
	;
	if v121 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v298 = v28 + int32(128)
	v300 = F_JsonbIteratorNext(m, l0, v298, int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L83
	}
L68:
	;
	v279 = F_JsonbIteratorNext(m, l0, v28+int32(128), int32(1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	F_setPath(m, l0, l1, l2, l3, l4, l5+int32(1), l6, l7)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L82
	}
L71:
	;
	if l7&int32(9) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if l7&int32(24) != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	F_pushJsonbValue(m, l4, v279, v28+int32(128))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v288 = int32(1)
	if l7&int32(20) == int32(0) {
		v370 = v288
		goto L64
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v370 = v288
	goto L64
L82:
	;
	v370 = int32(1)
	goto L64
L83:
	;
	if base.Ui32(v300) < base.Ui32(int32(4)) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v305 = v298
	goto L86
L85:
	;
	v305 = int32(0)
	goto L86
L86:
	;
	F_pushJsonbValue(m, l4, v300, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v300&int32(-3) != int32(4) {
		v370 = v258
		goto L64
	} else {
		goto L88
	}
L88:
	;
	v323 = int32(1)
	goto L89
L89:
	;
	v339 = v28 + int32(128)
	v341 = F_JsonbIteratorNext(m, l0, v339, int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	v370 = v258
	goto L64
L91:
	;
	if base.Ui32(v341) < base.Ui32(int32(4)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v346 = v339
	goto L94
L93:
	;
	v346 = int32(0)
	goto L94
L94:
	;
	F_pushJsonbValue(m, l4, v341, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v350 = v341 & int32(-3)
	v356 = v323 + base.B2i32(v350 == int32(4)) - base.B2i32(v350 == int32(5))
	if v356 != 0 {
		v323 = v356
		goto L89
	} else {
		goto L96
	}
L96:
	;
	goto L90
L97:
	;
	goto L63
L98:
	;
	v592 = F_JsonbIteratorNext(m, l0, v28-int32(-64), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L119
	}
L99:
	;
	if base.B2i32(v44 == int32(0))|base.B2i32(base.Ui32(v129) <= base.Ui32(v60)) != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	if (base.B2i32(v44 == int32(0))|base.B2i32(v120 <= l5)|v398)&int32(1) != 0 {
		goto L98
	} else {
		goto L110
	}
L102:
	;
	F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L109
	}
L103:
	;
	v422 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v422
	v424 = v129 - v60
	if v424 <= v422 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v436 = v424
	goto L105
L105:
	;
	F_pushJsonbValue(m, l4, int32(3), v28+int32(160))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	goto L102
L107:
	;
	v457 = int32(1)
	if base.Ui32(v457) < base.Ui32(v436) {
		v436 = v436 - v457
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	goto L98
L110:
	;
	if v129 <= int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_push_path(m, l4, l5, l1, l2, l3, l6)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L118
	}
L112:
	;
	v498 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v498
	v500 = v129 - v60
	if v500 <= v498 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v512 = v500
	goto L114
L114:
	;
	F_pushJsonbValue(m, l4, int32(3), v28+int32(160))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	goto L111
L116:
	;
	v533 = int32(1)
	if base.Ui32(v533) < base.Ui32(v512) {
		v512 = v512 - v533
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	goto L98
L119:
	;
	F_pushJsonbValue(m, l4, v592, int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L11
L121:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v28)+72))
	v602 = int32(1)
	if l3 <= l5 {
		v612 = v9
		v613 = v602
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v615 = l3 - int32(1)
	v616 = base.B2i32(l5 != v615)
	v618 = l7 & int32(25)
	v619 = int32(0)
	if v601|(v616|base.B2i32(v618 == v619)) == v619 {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v604 != 0 {
		v612 = v9
		v613 = v602
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l1+l5<<(uint(int32(3))%32))))
	v610 = F_pg_detoast_datum_packed(m, v609)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v612 = v610
	v613 = int32(0)
	goto L122
L126:
	;
	v1016 = int32(0)
	if v1001|(base.B2i32(l7&int32(32) == v1016)|base.B2i32(v615 <= l5)) == v1016 {
		goto L231
	} else {
		goto L232
	}
L127:
	;
	v625 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v625
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	if v629&v625 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	if v601 == int32(0) {
		v1001 = v613
		goto L126
	} else {
		goto L146
	}
L130:
	;
	v632 = v625
	goto L132
L131:
	;
	v632 = int32(4)
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+172)) = v612 + v632
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	if v635 == int32(1) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+168)) = v664
	F_pushJsonbValue(m, l4, int32(1), v28+int32(160))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L144
	}
L134:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)))
	if v641 == int32(18) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v652 = int32(1)
	if v635&v652 != 0 {
		v664 = int32(base.Ui32(v635)>>(uint(v652)%32)) - v652
		goto L133
	} else {
		goto L143
	}
L137:
	;
	v644 = int32(16)
	goto L139
L138:
	;
	v644 = int32(0)
	goto L139
L139:
	;
	if base.Ui32((v641-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v651 = int32(4)
	goto L142
L141:
	;
	v651 = v644
	goto L142
L142:
	;
	v664 = v651
	goto L133
L143:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v664 = int32(base.Ui32(v658)>>(uint(int32(2))%32)) - int32(4)
	goto L133
L144:
	;
	F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v1001 = v613
	goto L126
L146:
	;
	v680 = int32(1)
	v685 = int32(0)
	v702 = v613
	v707 = v9
	goto L147
L147:
	;
	v718 = F_JsonbIteratorNext(m, l0, v28+int32(160), int32(1))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v1001 = v973
	goto L126
L149:
	;
	if v702 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v987 = v707 + int32(1)
	if v987 != v601 {
		v702 = v973
		v707 = v987
		goto L147
	} else {
		goto L230
	}
L151:
	;
	F_pushJsonbValue(m, l4, v718, v28+int32(160))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L215
	}
L152:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	if v720 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	if base.B2i32(v707 != v601-v680)|base.B2i32(base.B2i32(l5 == v615)&base.B2i32(v618 != v685) == v685) != 0 {
		goto L151
	} else {
		goto L198
	}
L154:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
	if v749 != v750 {
		goto L153
	} else {
		goto L165
	}
L155:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)))
	if v726 == int32(18) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	goto L157
L157:
	;
	v737 = int32(1)
	if v720&v737 != 0 {
		v749 = int32(base.Ui32(v720)>>(uint(v737)%32)) - v737
		goto L154
	} else {
		goto L164
	}
L158:
	;
	v729 = int32(16)
	goto L160
L159:
	;
	v729 = int32(0)
	goto L160
L160:
	;
	if base.Ui32((v726-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v736 = int32(4)
	goto L163
L162:
	;
	v736 = v729
	goto L163
L163:
	;
	v749 = v736
	goto L154
L164:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v749 = int32(base.Ui32(v743)>>(uint(int32(2))%32)) - int32(4)
	goto L154
L165:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v28)+172))
	v753 = int32(1)
	if v720&v753 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v757 = v753
	goto L168
L167:
	;
	v757 = int32(4)
	goto L168
L168:
	;
	v758 = v612 + v757
	if base.Ui32(int32(4)) <= base.Ui32(v750) {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	if v820 != 0 {
		goto L153
	} else {
		goto L187
	}
L170:
	;
	v820 = int32(0)
	goto L169
L171:
	;
	v794 = v789
	v795 = v790
	v796 = v791
	goto L181
L172:
	;
	if (v752|v758)&int32(3) != 0 {
		v789 = v752
		v790 = v758
		v791 = v750
		goto L171
	} else {
		goto L175
	}
L173:
	;
	v782 = v752
	v783 = v758
	v784 = v750
	goto L174
L174:
	;
	if v784 == int32(0) {
		goto L170
	} else {
		goto L180
	}
L175:
	;
	v766 = v752
	v767 = v758
	v768 = v750
	goto L176
L176:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v766)))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	if v771 != v772 {
		v789 = v766
		v790 = v767
		v791 = v768
		goto L171
	} else {
		goto L178
	}
L177:
	;
	v782 = v777
	v783 = v775
	v784 = v779
	goto L174
L178:
	;
	v774 = int32(4)
	v775 = v767 + v774
	v777 = v766 + v774
	v779 = v768 - v774
	if base.Ui32(int32(3)) < base.Ui32(v779) {
		v766 = v777
		v767 = v775
		v768 = v779
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v789 = v782
	v790 = v783
	v791 = v784
	goto L171
L181:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v795))))
	if v799 == v800 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v820 = v799 - v800
	goto L169
L183:
	;
	v802 = int32(1)
	v807 = v796 - v802
	if v807 != 0 {
		v794 = v794 + v802
		v795 = v795 + v802
		v796 = v807
		goto L181
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	goto L182
L186:
	;
	goto L170
L187:
	;
	if v616 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	if l7&int32(24) != 0 {
		goto L5
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	F_pushJsonbValue(m, l4, v718, v28+int32(160))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L196
	}
L191:
	;
	v823 = int32(1)
	v827 = F_JsonbIteratorNext(m, l0, v28+int32(128), v823)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	if l7&int32(2) != 0 {
		v973 = v823
		goto L150
	} else {
		goto L193
	}
L193:
	;
	F_pushJsonbValue(m, l4, int32(1), v28+int32(160))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v973 = v823
	goto L150
L196:
	;
	F_setPath(m, l0, l1, l2, l3, l4, l5+v680, l6, l7)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v973 = int32(1)
	goto L150
L198:
	;
	v846 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v846
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	if v850&v846 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v853 = v846
	goto L201
L200:
	;
	v853 = int32(4)
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v612 + v853
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	if v856 == int32(1) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v885
	F_pushJsonbValue(m, l4, int32(1), v28+int32(96))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L213
	}
L203:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)))
	if v862 == int32(18) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v873 = int32(1)
	if v856&v873 != 0 {
		v885 = int32(base.Ui32(v856)>>(uint(v873)%32)) - v873
		goto L202
	} else {
		goto L212
	}
L206:
	;
	v865 = int32(16)
	goto L208
L207:
	;
	v865 = int32(0)
	goto L208
L208:
	;
	if base.Ui32((v862-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v872 = int32(4)
	goto L211
L210:
	;
	v872 = v865
	goto L211
L211:
	;
	v885 = v872
	goto L202
L212:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v885 = int32(base.Ui32(v879)>>(uint(int32(2))%32)) - int32(4)
	goto L202
L213:
	;
	F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	goto L151
L215:
	;
	v902 = v28 + int32(128)
	v904 = F_JsonbIteratorNext(m, l0, v902, int32(0))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	if base.Ui32(v904) < base.Ui32(int32(4)) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v909 = v902
	goto L219
L218:
	;
	v909 = int32(0)
	goto L219
L219:
	;
	F_pushJsonbValue(m, l4, v904, v909)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	if v904&int32(-3) != int32(4) {
		v973 = v702
		goto L150
	} else {
		goto L221
	}
L221:
	;
	v927 = int32(1)
	goto L222
L222:
	;
	v943 = v28 + int32(128)
	v945 = F_JsonbIteratorNext(m, l0, v943, int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L224
	}
L223:
	;
	v973 = v702
	goto L150
L224:
	;
	if base.Ui32(v945) < base.Ui32(int32(4)) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v950 = v943
	goto L227
L226:
	;
	v950 = int32(0)
	goto L227
L227:
	;
	F_pushJsonbValue(m, l4, v945, v950)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v954 = v945 & int32(-3)
	v960 = v927 + base.B2i32(v954 == int32(4)) - base.B2i32(v954 == int32(5))
	if v960 != 0 {
		v927 = v960
		goto L222
	} else {
		goto L229
	}
L229:
	;
	goto L223
L230:
	;
	goto L148
L231:
	;
	v1023 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v1023
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
	if v1027&v1023 != 0 {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	goto L233
L233:
	;
	v1074 = F_JsonbIteratorNext(m, l0, v28-int32(-64), int32(1))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L250
	}
L234:
	;
	v1030 = v1023
	goto L236
L235:
	;
	v1030 = int32(4)
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v612 + v1030
	if v1027 == int32(1) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v1061
	F_pushJsonbValue(m, l4, int32(1), v28+int32(96))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L248
	}
L238:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)))
	if v1038 == int32(18) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	v1049 = int32(1)
	if v1027&v1049 != 0 {
		v1061 = int32(base.Ui32(v1027)>>(uint(v1049)%32)) - v1049
		goto L237
	} else {
		goto L247
	}
L241:
	;
	v1041 = int32(16)
	goto L243
L242:
	;
	v1041 = int32(0)
	goto L243
L243:
	;
	if base.Ui32((v1038-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1048 = int32(4)
	goto L246
L245:
	;
	v1048 = v1041
	goto L246
L246:
	;
	v1061 = v1048
	goto L237
L247:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v1061 = int32(base.Ui32(v1055)>>(uint(int32(2))%32)) - int32(4)
	goto L237
L248:
	;
	F_push_path(m, l4, l5, l1, l2, l3, l6)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	goto L233
L250:
	;
	F_pushJsonbValue(m, l4, v1074, int32(0))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	goto L11
L252:
	;
	v1085 = base.B2i32(l5 <= l3-int32(1))
	goto L254
L253:
	;
	v1085 = int32(0)
	goto L254
L254:
	;
	if v1085 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	F_pushJsonbValue(m, l4, v39, v28-int32(-64))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	goto L11
L257:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l5 + int32(1)
	F_errmsg(m, int32(_a_F_setPath_4), v28)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_5), int32(_a_F_setPath_6))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(_a_F_setPath_7), int32(0))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v1149 = F_errdetail(m, int32(_a_F_setPath_8), int32(0))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_9), int32(_a_F_setPath_6))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = l5 + int32(1)
	F_errmsg(m, int32(_a_F_setPath_10), v28+int32(48))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_11), int32(_a_F_setPath_3))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_errmsg(m, int32(_a_F_setPath_7), int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_errhint(m, int32(_a_F_setPath_12), int32(0))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_13), int32(_a_F_setPath_14))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	F_errmsg(m, int32(_a_F_setPath_7), int32(0))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	v1210 = F_errdetail(m, int32(_a_F_setPath_8), int32(0))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_15), int32(_a_F_setPath_6))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v39
	F_errmsg_internal(m, int32(_a_F_setPath_16), v28+int32(16))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_17), int32(_a_F_setPath_6))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_substitute_path_macro(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 != int32(36) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v9 + int32(32)
	return v80
L3:
	;
	v14 = F_pstrdup(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v19 = Fn14265(m, l0, int32(47))
	mBase = m.M
	goto L8
L6:
	;
	return int32(0)
L7:
	;
	v80 = v14
	goto L2
L8:
	;
	if v19 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v22 = F_strlen(m, l0)
	mBase = m.M
	v24 = v22 + l0
	goto L11
L10:
	;
	v24 = v19
	goto L11
L11:
	;
	v25 = F_strlen(m, l1)
	mBase = m.M
	if v25 != v24-l0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v25 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v72 != 0 {
		goto L1
	} else {
		goto L26
	}
L14:
	;
	v72 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v33 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v34 = l0
	v35 = l1
	v36 = v25
	v37 = v33
	goto L21
L18:
	;
	v60 = l1
	v64 = int32(0)
	goto L19
L19:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v72 = v64 - v65
	goto L13
L20:
	;
	v60 = v55
	v64 = v57
	goto L19
L21:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if base.B2i32(v37 != v39)|base.B2i32(v39 == int32(0)) != 0 {
		v55 = v35
		v57 = v37
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v55 = v49
	v57 = int32(0)
	goto L20
L23:
	;
	v45 = v36 - int32(1)
	if v45 == int32(0) {
		v55 = v35
		v57 = v37
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v48 = int32(1)
	v49 = v35 + v48
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v50 != 0 {
		v34 = v34 + v48
		v35 = v49
		v36 = v45
		v37 = v50
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
	v76 = F_psprintf(m, int32(_a_F_substitute_path_macro_0), v9)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v80 = v76
	goto L2
L28:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	F_errmsg(m, int32(_a_F_substitute_path_macro_1), v9+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_substitute_path_macro_2), int32(555), int32(_a_F_substitute_path_macro_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
