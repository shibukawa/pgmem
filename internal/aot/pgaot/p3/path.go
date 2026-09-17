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
func F_path_add(m *base.Module, l0 int32) int32 {
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
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
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
						v136 = m.ExcPending
						if v136 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_path_add_0), int32(0))
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_path_add_1), int32(_a_F_path_add_2), int32(_a_F_path_add_3))
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
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
						v33 = v29 << (uint(int32(4)) % 32)
						if v33 == int32(2147483632) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_path_add_0), int32(0))
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_path_add_1), int32(_a_F_path_add_2), int32(_a_F_path_add_3))
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
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
							v37 = v33 + int32(16)
							v38 = F_palloc(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
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
								return v38
							}
						}
					}
				} else {
					v23 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
					return int32(0)
				}
			} else {
				v23 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
				return int32(0)
			}
		}
	}
}
func F_path_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
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
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = int32(44)
	v17 = F___strchrnul(m, v15, v16)
	mBase = m.M
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 == v16 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v278
L2:
	;
	v257 = F_errsave_start(m, v14)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L36
	} else {
		goto L67
	}
L3:
	;
	if v23 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	v23 = v17
	goto L6
L5:
	;
	v23 = v2
	goto L6
L6:
	;
	goto L3
L7:
	;
	v28 = v23
	v30 = v2
	goto L8
L8:
	;
	v35 = int32(1)
	v39 = int32(44)
	v40 = F___strchrnul(m, v28+v35, v39)
	mBase = m.M
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v42 == v39 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v50 = int32(1)
	if v30&v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if v46 != 0 {
		v28 = v46
		v30 = v30 + v35
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v46 = v40
	goto L13
L12:
	;
	v46 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L9
L15:
	;
	v54 = int32(-1)
	goto L17
L16:
	;
	v54 = (v30 + int32(2)) >> (uint(v50) % 32)
	goto L17
L17:
	;
	if v54 <= int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v59 = v15
	goto L21
L19:
	;
	v111 = v54 << (uint(int32(4)) % 32)
	v112 = base.I32_div_s(v111, v54)
	if base.B2i32(v112 == int32(16))&base.B2i32(v111 != int32(2147483632)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v105
	v108 = v105
	v109 = v106
	goto L19
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if base.Ui32(v66-int32(9)) < base.Ui32(int32(5)) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v59
	v81 = F_strlen(m, v59)
	mBase = m.M
	v88 = v81 + int32(1)
	goto L28
L23:
	;
	goto L22
L24:
	;
	v59 = v59 + int32(1)
	goto L21
L25:
	;
	v71 = int32(0)
	switch v66 - int32(32) {
	case 0:
		goto L24
	default:
		v105 = v59
		v106 = v71
		goto L20
	case 8:
		goto L23
	}
L26:
	;
	if v100 != v59 {
		v108 = v59
		v109 = v71
		goto L19
	} else {
		goto L32
	}
L27:
	;
	goto L26
L28:
	;
	v90 = int32(0)
	if v88 == v90 {
		v100 = v90
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v100 = v95
	goto L27
L30:
	;
	v94 = v88 - int32(1)
	v95 = v59 + v94
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v96 != int32(40) {
		v88 = v94
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v102 = int32(1)
	v105 = v59 + v102
	v106 = v102
	goto L20
L33:
	;
	v120 = int32(0)
	v121 = F_errsave_start(m, v14)
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
	v140 = v111 + int32(16)
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
	return int32(0)
L37:
	;
	if v121 == int32(0) {
		v278 = v120
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
	F_errsave_finish(m, v14, int32(_a_F_path_in_1), int32(1438), int32(_a_F_path_in_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v278 = v120
	goto L1
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v140 << (uint(int32(2)) % 32)
	v155 = F_path_decode(m, v108, int32(1), v54, v141+int32(16), v12+int32(47), v12+int32(40), int32(_a_F_path_in_3), v15, v14)
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
	v278 = int32(0)
	goto L1
L45:
	;
	goto L46
L46:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v109 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v242 ^ int32(1)
	v278 = v141
	goto L1
L48:
	;
	v221 = int32(0)
	v222 = F_errsave_start(m, v14)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L36
	} else {
		goto L62
	}
L49:
	;
	if v163&int32(255) != int32(41) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	v191 = v163
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
	v170 = v162
	goto L53
L53:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
	if base.B2i32(base.Ui32(v179-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v179 == int32(32)) != 0 {
		v170 = v170 + int32(1)
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
	v200 = int32(0)
	v201 = F_errsave_start(m, v14)
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
		v278 = v200
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_path_in_3)
	F_errmsg(m, int32(_a_F_path_in_4), v12+int32(16))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L36
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, v14, int32(_a_F_path_in_1), int32(1463), int32(_a_F_path_in_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L36
	} else {
		goto L61
	}
L61:
	;
	v278 = v200
	goto L1
L62:
	;
	if v222 == int32(0) {
		v278 = v221
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L36
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(_a_F_path_in_3)
	F_errmsg(m, int32(_a_F_path_in_4), v12+int32(32))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L36
	} else {
		goto L65
	}
L65:
	;
	F_errsave_finish(m, v14, int32(_a_F_path_in_1), int32(1455), int32(_a_F_path_in_2))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L36
	} else {
		goto L66
	}
L66:
	;
	v278 = v221
	goto L1
L67:
	;
	if v257 == int32(0) {
		v278 = v2
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L36
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_path_in_3)
	F_errmsg(m, int32(_a_F_path_in_4), v12)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L36
	} else {
		goto L70
	}
L70:
	;
	F_errsave_finish(m, v14, int32(_a_F_path_in_1), int32(1418), int32(_a_F_path_in_2))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L36
	} else {
		goto L71
	}
L71:
	;
	v278 = v2
	goto L1
}
func F_path_n_eq(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			return base.B2i32(v11 == v12)
		}
	}
}
func F_path_sub_pt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_copy(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L12
	}
L2:
	;
	return int32(0)
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if int32(0) < v16 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	return v12
L7:
	;
	v34 = v12 + int32(16) + v29<<(uint(int32(4))%32)
	v35 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
	v37 = base.F64_sub(v35, v36)
	v39 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39) == int32(0))&base.F64_ne(base.F64_abs(v36), v39) != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v34)+8))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	v53 = base.F64_sub(v51, v52)
	v55 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v53), v55)|base.F64_eq(base.F64_abs(v51), v55) == int32(0))&base.F64_ne(base.F64_abs(v52), v55) != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v34)+8)) = v53
	*(*float64)(unsafe.Add(mBase, uint32(v34))) = v37
	v70 = v29 + int32(1)
	if v70 != v16 {
		v29 = v70
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v158 int32
	_ = v158
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v237 int32
	_ = v237
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v325 int32
	_ = v325
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v438 int32
	_ = v438
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v514 int32
	_ = v514
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v565 int32
	_ = v565
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v931 int32
	_ = v931
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v976 int32
	_ = v976
	var v991 int32
	_ = v991
	var v1004 int32
	_ = v1004
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1121 int32
	_ = v1121
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	v9 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(144)
	m.G0 = v28
	F_check_stack_depth(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = l2 + l5
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v35 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L1
	} else {
		goto L280
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L1
	} else {
		goto L275
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L270
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L266
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L261
	}
L8:
	;
	v41 = F_JsonbIteratorNext(m, l0, v28-int32(-64), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
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
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L257
	}
L11:
	;
	m.G0 = v28 + int32(144)
	return v1121
L12:
	;
	if l7&int32(32) != 0 {
		goto L252
	} else {
		goto L253
	}
L13:
	;
	v601 = F_pushJsonbValue(m, l4, int32(6), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L121
	}
L14:
	;
	v46 = l7 & int32(32)
	v47 = int32(0)
	if base.B2i32(v46 == v47)|base.B2i32(l3-int32(1) < l5) == v47 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	switch v41 - int32(2) {
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
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+76)))
	if v55&int32(1) != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v60 = F_pushJsonbValue(m, l4, int32(4), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
	v63 = base.B2i32(l3 <= l5)
	if l3 <= l5 {
		v85 = v62
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if int32(0) <= v85 {
		v116 = v85
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v64 != 0 {
		v85 = v62
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1+l5<<(uint(int32(2))%32))))
	v69 = F_text_to_cstring(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_setPath[0])) = int32(0)
	v77 = F_strtol(m, v69, v28+int32(124), int32(10))
	mBase = m.M
	goto L25
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v28)+124))
	if v69 == v78 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v80 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_setPath[0]))
	if v82 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v85 = v77
	goto L21
L29:
	;
	v118 = l7 & int32(25)
	v119 = int32(0)
	v122 = l3 - int32(1)
	v123 = base.B2i32(l5 != v122)
	if base.Ui32(v62) < base.Ui32(v116) {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	if base.Ui32(v62) < base.Ui32(int32(0)-v85) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(l7) < base.Ui32(int32(64)) {
		v116 = int32(-2147483648)
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v116 = v85 + v62
	goto L29
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l5 + int32(1)
	F_errmsg(m, int32(_a_F_setPath_0), v28+int32(32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_2), int32(_a_F_setPath_3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
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
	v127 = v62
	goto L41
L40:
	;
	v127 = v116
	goto L41
L41:
	;
	if int32(0) < v116 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v130 = v127
	goto L44
L43:
	;
	v130 = v116
	goto L44
L44:
	;
	if v46 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v131 = v116
	goto L47
L46:
	;
	v131 = v130
	goto L47
L47:
	;
	v136 = base.B2i32(v118 == v119) | (v123 | base.B2i32(v62 != v119)&base.B2i32(v131 != int32(-2147483648)))
	if v136 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v139 = int32(0)
	if v62|base.B2i32(v46 == v139)|base.B2i32(v131 <= v139) == v139 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v237 = base.B2i32(v136 == int32(0))
	if v62 != 0 {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = int32(0)
	v158 = v131
	goto L54
L52:
	;
	goto L53
L53:
	;
	v209 = F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L58
	}
L54:
	;
	v177 = F_pushJsonbValue(m, l4, int32(3), v28+int32(124))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L53
L56:
	;
	v179 = int32(1)
	if base.Ui32(v179) < base.Ui32(v158) {
		v158 = v158 - v179
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
	v256 = v237
	v259 = int32(0)
	goto L62
L60:
	;
	v396 = v237
	goto L61
L61:
	;
	v414 = int32(0)
	if v396&int32(1)|(v123|base.B2i32(v118 == v414)) == v414 {
		goto L99
	} else {
		goto L100
	}
L62:
	;
	if v63|base.B2i32(v131 != v259) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v396 = v368
	goto L61
L64:
	;
	v385 = v259 + int32(1)
	if v385 != v62 {
		v256 = v368
		v259 = v385
		goto L62
	} else {
		goto L97
	}
L65:
	;
	if v123 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v300 = v28 + int32(104)
	v302 = F_JsonbIteratorNext(m, l0, v300, int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L83
	}
L68:
	;
	v281 = F_JsonbIteratorNext(m, l0, v28+int32(104), int32(1))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v296 = F_setPath(m, l0, l1, l2, l3, l4, l5+int32(1), l6, l7)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
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
	v284 = F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
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
	v288 = F_pushJsonbValue(m, l4, v281, v28+int32(104))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v290 = int32(1)
	if l7&int32(20) == int32(0) {
		v368 = v290
		goto L64
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v294 = F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v368 = v290
	goto L64
L82:
	;
	v368 = int32(1)
	goto L64
L83:
	;
	if base.Ui32(v302) < base.Ui32(int32(4)) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v307 = v300
	goto L86
L85:
	;
	v307 = int32(0)
	goto L86
L86:
	;
	v308 = F_pushJsonbValue(m, l4, v302, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v302&int32(-3) != int32(4) {
		v368 = v256
		goto L64
	} else {
		goto L88
	}
L88:
	;
	v325 = int32(1)
	goto L89
L89:
	;
	v341 = v28 + int32(104)
	v343 = F_JsonbIteratorNext(m, l0, v341, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	v368 = v256
	goto L64
L91:
	;
	if base.Ui32(v343) < base.Ui32(int32(4)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v348 = v341
	goto L94
L93:
	;
	v348 = int32(0)
	goto L94
L94:
	;
	v349 = F_pushJsonbValue(m, l4, v343, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v352 = v343 & int32(-3)
	v358 = v325 + base.B2i32(v352 == int32(4)) - base.B2i32(v352 == int32(5))
	if v358 != 0 {
		v325 = v358
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
	v594 = F_JsonbIteratorNext(m, l0, v28-int32(-64), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L119
	}
L99:
	;
	if base.B2i32(v46 == int32(0))|base.B2i32(base.Ui32(v131) <= base.Ui32(v62)) != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	if (base.B2i32(v46 == int32(0))|base.B2i32(v122 <= l5)|v396)&int32(1) != 0 {
		goto L98
	} else {
		goto L110
	}
L102:
	;
	v489 = F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L109
	}
L103:
	;
	v424 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v424
	v426 = v131 - v62
	if v426 <= v424 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v438 = v426
	goto L105
L105:
	;
	v457 = F_pushJsonbValue(m, l4, int32(3), v28+int32(124))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	goto L102
L107:
	;
	v459 = int32(1)
	if base.Ui32(v459) < base.Ui32(v438) {
		v438 = v438 - v459
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
	if v131 <= int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_push_path(m, l4, l5, l1, l2, l3, l6)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L118
	}
L112:
	;
	v500 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v500
	v502 = v131 - v62
	if v502 <= v500 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v514 = v502
	goto L114
L114:
	;
	v533 = F_pushJsonbValue(m, l4, int32(3), v28+int32(124))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	goto L111
L116:
	;
	v535 = int32(1)
	if base.Ui32(v535) < base.Ui32(v514) {
		v514 = v514 - v535
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
	v597 = F_pushJsonbValue(m, l4, v594, int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v1121 = v597
	goto L11
L121:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
	v604 = int32(1)
	if l3 <= l5 {
		v614 = v9
		v615 = v604
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v617 = l3 - int32(1)
	v618 = base.B2i32(l5 != v617)
	v620 = l7 & int32(25)
	v621 = int32(0)
	if v603|(v618|base.B2i32(v620 == v621)) == v621 {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v606 != 0 {
		v614 = v9
		v615 = v604
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l1+l5<<(uint(int32(2))%32))))
	v612 = F_pg_detoast_datum_packed(m, v611)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v614 = v612
	v615 = int32(0)
	goto L122
L126:
	;
	v1020 = int32(0)
	if v1004|(base.B2i32(l7&int32(32) == v1020)|base.B2i32(v617 <= l5)) == v1020 {
		goto L231
	} else {
		goto L232
	}
L127:
	;
	v627 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v627
	v630 = v614 + v627
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v633&v627 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	if v603 == int32(0) {
		v1004 = v615
		goto L126
	} else {
		goto L146
	}
L130:
	;
	v636 = v630
	goto L132
L131:
	;
	v636 = v614 + int32(4)
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+132)) = v636
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v638 == int32(1) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v667
	v672 = F_pushJsonbValue(m, l4, int32(1), v28+int32(124))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L144
	}
L134:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
	if v644 == int32(18) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v655 = int32(1)
	if v638&v655 != 0 {
		v667 = int32(base.Ui32(v638)>>(uint(v655)%32)) - v655
		goto L133
	} else {
		goto L143
	}
L137:
	;
	v647 = int32(16)
	goto L139
L138:
	;
	v647 = int32(0)
	goto L139
L139:
	;
	if base.Ui32((v644-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v654 = int32(4)
	goto L142
L141:
	;
	v654 = v647
	goto L142
L142:
	;
	v667 = v654
	goto L133
L143:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v667 = int32(base.Ui32(v661)>>(uint(int32(2))%32)) - int32(4)
	goto L133
L144:
	;
	v675 = F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v1004 = v615
	goto L126
L146:
	;
	v683 = int32(1)
	v690 = v614 + v683
	v692 = int32(0)
	v708 = v615
	v709 = v9
	goto L147
L147:
	;
	v725 = F_JsonbIteratorNext(m, l0, v28+int32(124), int32(1))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v1004 = v976
	goto L126
L149:
	;
	if v708 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v991 = v709 + int32(1)
	if v991 != v603 {
		v708 = v976
		v709 = v991
		goto L147
	} else {
		goto L230
	}
L151:
	;
	v903 = F_pushJsonbValue(m, l4, v725, v28+int32(124))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L215
	}
L152:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v727 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	if base.B2i32(v709 != v603-v683)|base.B2i32(base.B2i32(l5 == v617)&base.B2i32(v620 != v692) == v692) != 0 {
		goto L151
	} else {
		goto L198
	}
L154:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v28)+128))
	if v756 != v757 {
		goto L153
	} else {
		goto L165
	}
L155:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690))))
	if v733 == int32(18) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	goto L157
L157:
	;
	v744 = int32(1)
	if v727&v744 != 0 {
		v756 = int32(base.Ui32(v727)>>(uint(v744)%32)) - v744
		goto L154
	} else {
		goto L164
	}
L158:
	;
	v736 = int32(16)
	goto L160
L159:
	;
	v736 = int32(0)
	goto L160
L160:
	;
	if base.Ui32((v733-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v743 = int32(4)
	goto L163
L162:
	;
	v743 = v736
	goto L163
L163:
	;
	v756 = v743
	goto L154
L164:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v756 = int32(base.Ui32(v750)>>(uint(int32(2))%32)) - int32(4)
	goto L154
L165:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v28)+132))
	v760 = int32(1)
	if v727&v760 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v764 = v760
	goto L168
L167:
	;
	v764 = int32(4)
	goto L168
L168:
	;
	v765 = v614 + v764
	if base.Ui32(int32(4)) <= base.Ui32(v757) {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	if v827 != 0 {
		goto L153
	} else {
		goto L187
	}
L170:
	;
	v827 = int32(0)
	goto L169
L171:
	;
	v801 = v796
	v802 = v797
	v803 = v798
	goto L181
L172:
	;
	if (v759|v765)&int32(3) != 0 {
		v796 = v759
		v797 = v765
		v798 = v757
		goto L171
	} else {
		goto L175
	}
L173:
	;
	v789 = v759
	v790 = v765
	v791 = v757
	goto L174
L174:
	;
	if v791 == int32(0) {
		goto L170
	} else {
		goto L180
	}
L175:
	;
	v773 = v759
	v774 = v765
	v775 = v757
	goto L176
L176:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v773)))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	if v778 != v779 {
		v796 = v773
		v797 = v774
		v798 = v775
		goto L171
	} else {
		goto L178
	}
L177:
	;
	v789 = v784
	v790 = v782
	v791 = v786
	goto L174
L178:
	;
	v781 = int32(4)
	v782 = v774 + v781
	v784 = v773 + v781
	v786 = v775 - v781
	if base.Ui32(int32(3)) < base.Ui32(v786) {
		v773 = v784
		v774 = v782
		v775 = v786
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v796 = v789
	v797 = v790
	v798 = v791
	goto L171
L181:
	;
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802))))
	if v806 == v807 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v827 = v806 - v807
	goto L169
L183:
	;
	v809 = int32(1)
	v814 = v803 - v809
	if v814 != 0 {
		v801 = v801 + v809
		v802 = v802 + v809
		v803 = v814
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
	if v618 == int32(0) {
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
	v846 = F_pushJsonbValue(m, l4, v725, v28+int32(124))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L196
	}
L191:
	;
	v830 = int32(1)
	v834 = F_JsonbIteratorNext(m, l0, v28+int32(104), v830)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	if l7&int32(2) != 0 {
		v976 = v830
		goto L150
	} else {
		goto L193
	}
L193:
	;
	v839 = F_pushJsonbValue(m, l4, int32(1), v28+int32(124))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v842 = F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v976 = v830
	goto L150
L196:
	;
	v848 = F_setPath(m, l0, l1, l2, l3, l4, l5+v683, l6, l7)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v976 = int32(1)
	goto L150
L198:
	;
	v853 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v853
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v855&v853 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v858 = v690
	goto L201
L200:
	;
	v858 = v614 + int32(4)
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v858
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v860 == int32(1) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v889
	v894 = F_pushJsonbValue(m, l4, int32(1), v28+int32(84))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L213
	}
L203:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690))))
	if v866 == int32(18) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v877 = int32(1)
	if v860&v877 != 0 {
		v889 = int32(base.Ui32(v860)>>(uint(v877)%32)) - v877
		goto L202
	} else {
		goto L212
	}
L206:
	;
	v869 = int32(16)
	goto L208
L207:
	;
	v869 = int32(0)
	goto L208
L208:
	;
	if base.Ui32((v866-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v876 = int32(4)
	goto L211
L210:
	;
	v876 = v869
	goto L211
L211:
	;
	v889 = v876
	goto L202
L212:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v889 = int32(base.Ui32(v883)>>(uint(int32(2))%32)) - int32(4)
	goto L202
L213:
	;
	v897 = F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	goto L151
L215:
	;
	v906 = v28 + int32(104)
	v908 = F_JsonbIteratorNext(m, l0, v906, int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	if base.Ui32(v908) < base.Ui32(int32(4)) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v913 = v906
	goto L219
L218:
	;
	v913 = int32(0)
	goto L219
L219:
	;
	v914 = F_pushJsonbValue(m, l4, v908, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	if v908&int32(-3) != int32(4) {
		v976 = v708
		goto L150
	} else {
		goto L221
	}
L221:
	;
	v931 = int32(1)
	goto L222
L222:
	;
	v947 = v28 + int32(104)
	v949 = F_JsonbIteratorNext(m, l0, v947, int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L224
	}
L223:
	;
	v976 = v708
	goto L150
L224:
	;
	if base.Ui32(v949) < base.Ui32(int32(4)) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v954 = v947
	goto L227
L226:
	;
	v954 = int32(0)
	goto L227
L227:
	;
	v955 = F_pushJsonbValue(m, l4, v949, v954)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v958 = v949 & int32(-3)
	v964 = v931 + base.B2i32(v958 == int32(4)) - base.B2i32(v958 == int32(5))
	if v964 != 0 {
		v931 = v964
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
	v1027 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v1027
	v1030 = v614 + v1027
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v1033&v1027 != 0 {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	goto L233
L233:
	;
	v1080 = F_JsonbIteratorNext(m, l0, v28-int32(-64), int32(1))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
	} else {
		goto L250
	}
L234:
	;
	v1036 = v1030
	goto L236
L235:
	;
	v1036 = v614 + int32(4)
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v1036
	if v1033 == int32(1) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v1066
	v1071 = F_pushJsonbValue(m, l4, int32(1), v28+int32(84))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L248
	}
L238:
	;
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1030))))
	if v1043 == int32(18) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	v1054 = int32(1)
	if v1033&v1054 != 0 {
		v1066 = int32(base.Ui32(v1033)>>(uint(v1054)%32)) - v1054
		goto L237
	} else {
		goto L247
	}
L241:
	;
	v1046 = int32(16)
	goto L243
L242:
	;
	v1046 = int32(0)
	goto L243
L243:
	;
	if base.Ui32((v1043-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1053 = int32(4)
	goto L246
L245:
	;
	v1053 = v1046
	goto L246
L246:
	;
	v1066 = v1053
	goto L237
L247:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v1066 = int32(base.Ui32(v1060)>>(uint(int32(2))%32)) - int32(4)
	goto L237
L248:
	;
	F_push_path(m, l4, l5, l1, l2, l3, l6)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	goto L233
L250:
	;
	v1083 = F_pushJsonbValue(m, l4, v1080, int32(0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1121 = v1083
	goto L11
L252:
	;
	v1091 = base.B2i32(l5 <= l3-int32(1))
	goto L254
L253:
	;
	v1091 = int32(0)
	goto L254
L254:
	;
	if v1091 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	v1094 = F_pushJsonbValue(m, l4, v41, v28-int32(-64))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v1121 = v1094
	goto L11
L257:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l5 + int32(1)
	F_errmsg(m, int32(_a_F_setPath_4), v28)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_5), int32(_a_F_setPath_6))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
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
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(_a_F_setPath_7), int32(0))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_errdetail(m, int32(_a_F_setPath_8), int32(0))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_9), int32(_a_F_setPath_6))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
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
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = l5 + int32(1)
	F_errmsg(m, int32(_a_F_setPath_10), v28+int32(48))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_11), int32(_a_F_setPath_3))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
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
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_errmsg(m, int32(_a_F_setPath_7), int32(0))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_errhint(m, int32(_a_F_setPath_12), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_13), int32(_a_F_setPath_14))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
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
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	F_errmsg(m, int32(_a_F_setPath_7), int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_errdetail(m, int32(_a_F_setPath_8), int32(0))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_15), int32(_a_F_setPath_6))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v41
	F_errmsg_internal(m, int32(_a_F_setPath_16), v28+int32(16))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_setPath_1), int32(_a_F_setPath_17), int32(_a_F_setPath_6))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
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
	v19 = Fn13878(m, l0, int32(47))
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
