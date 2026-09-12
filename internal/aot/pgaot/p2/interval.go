package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AdjustIntervalForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v38 int32
	_ = v38
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	v4 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v14 != int32(2147483647) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return
L2:
	;
	v38 = int32(base.Ui32(l1) >> (uint(int32(16)) % 32))
	if base.Ui32(v38) <= base.Ui32(int32(2047)) {
		goto L27
	} else {
		goto L28
	}
L3:
	;
	if l1 < int32(0) {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	if v14 != int32(-2147483648) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v27 != int32(2147483647) {
		goto L3
	} else {
		goto L11
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v19 != int32(-2147483648) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if l1 < int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v24 != int64(-9223372036854775807-1) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	if l1 < int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v32 != int64(9223372036854775807) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	goto L2
L15:
	;
	v109 = int32(65535)
	v110 = l1 & v109
	if v110 == v109 {
		goto L1
	} else {
		goto L43
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v106
	goto L15
L17:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v104 = base.I64_rem_s(v102, int64(60000000))
	v106 = v102 - v104
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v99 = base.I32_rem_s(v14, int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14 - v99
	v106 = v4
	goto L16
L19:
	;
	if l1&int32(2080309248) == int32(402653184) {
		goto L15
	} else {
		goto L38
	}
L20:
	;
	if v38 == int32(2048) {
		goto L17
	} else {
		goto L37
	}
L21:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v73 = base.I64_rem_s(v71, int64(60000000))
	v106 = v71 - v73
	goto L16
L22:
	;
	v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v69 = base.I64_rem_s(v67, int64(60000000))
	v106 = v67 - v69
	goto L16
L23:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v65 = base.I64_rem_s(v63, int64(3600000000))
	v106 = v63 - v65
	goto L16
L24:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v61 = base.I64_rem_s(v59, int64(3600000000))
	v106 = v59 - v61
	goto L16
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v106 = v4
	goto L16
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v106 = v4
	goto L16
L27:
	;
	switch v38 - int32(2) {
	case 0:
		goto L26
	case 1, 3, 5:
		goto L19
	case 2:
		goto L18
	case 4:
		goto L25
	case 6:
		v106 = v4
		goto L16
	default:
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(v38) <= base.Ui32(int32(4095)) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	switch v38 - int32(1024) {
	case 0:
		goto L24
	default:
		goto L19
	case 8:
		goto L23
	}
L31:
	;
	switch v38 - int32(3072) {
	case 0:
		goto L21
	case 1, 2, 3, 4, 5, 6, 7:
		goto L19
	case 8:
		goto L22
	default:
		goto L20
	}
L32:
	;
	goto L33
L33:
	;
	if v38 == int32(4096) {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	if v38 == int32(7176) {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	if v38 != int32(32767) {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	goto L15
L37:
	;
	goto L19
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	F_errmsg_internal(m, int32(476238), v12+int32(16))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(484742), int32(1489), int32(412949))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	if base.Ui32(int32(7)) <= base.Ui32(v110) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v115 = F_errsave_start(m, l2)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L39
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v110<<(uint(int32(3))%32))+uint32(_consts[1147])))
	v138 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if int64(0) <= v138 {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	if v115 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L39
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(25769803776)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v110
	F_errmsg(m, int32(467693), v12)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	F_errsave_finish(m, l2, int32(484742), int32(1498), int32(412949))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	goto L1
L52:
	;
	v141 = v137 + v138
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v141
	if base.B2i32(v137 < int64(0)) != base.B2i32(v141 < v138) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v171 = v138 - v137
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v171
	if base.B2i32(v171 < v138) != base.B2i32(int64(0) < v137) {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	v147 = F_errsave_start(m, l2)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L39
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v110<<(uint(int32(3))%32))+uint32(_consts[1148])))
	v168 = base.I64_rem_s(v141, v167)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v141 - v168
	goto L1
L58:
	;
	if v147 == int32(0) {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L39
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(393566), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L39
	} else {
		goto L61
	}
L61:
	;
	F_errsave_finish(m, l2, int32(484742), int32(1507), int32(412949))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L39
	} else {
		goto L62
	}
L62:
	;
	goto L1
L63:
	;
	v177 = F_errsave_start(m, l2)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L39
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v110<<(uint(int32(3))%32))+uint32(_consts[1148])))
	v198 = base.I64_rem_s(v171, v197)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v171 - v198
	goto L1
L66:
	;
	if v177 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L39
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(393566), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L39
	} else {
		goto L69
	}
L69:
	;
	F_errsave_finish(m, l2, int32(484742), int32(1517), int32(412949))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L39
	} else {
		goto L70
	}
L70:
	;
	goto L1
}
func F_interval_avg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int64
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v7 != 0 {
		v22 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v8 == int32(0) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			return int32(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
			if v11+v12 != int64(0)-v15 {
				v26 = int64(0)
				if base.B2i32(v11 <= v26)&base.B2i32(v15 <= v26) == int32(0) {
					v33 = int64(0)
					if base.B2i32(v33 < v11)&base.B2i32(v33 < v15) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(393566), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(484742), int32(4248), int32(319426))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
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
						v39 = F_palloc(m, int32(16))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v45 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
							v47 = base.B2i32(int64(0) < v45)
							if int64(0) < v45 {
								v48 = int32(2147483647)
							} else {
								v48 = int32(-2147483648)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v48
							if int64(0) < v45 {
								v53 = int64(9223372036854775807)
							} else {
								v53 = int64(-9223372036854775807 - 1)
							}
							*(*int64)(unsafe.Add(mBase, uint32(v39))) = v53
							return v39
						}
					}
				} else {
					v61 = F_Float8GetDatum(m, base.F64_convert_i64_s(v12))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = F_DirectFunctionCall2Coll(m, int32(1515), int32(0), v8+int32(8), v61)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							return v63
						}
					}
				}
			} else {
				v22 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
				return int32(0)
			}
		}
	}
}
func F_interval_avg_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v9 == v2 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = v12
	} else {
		v13 = v2
	}
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v14 != 0 {
		v100 = v13
		m.G0 = v7 + int32(16)
		return v100
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v15 == int32(0) {
			v100 = v13
			m.G0 = v7 + int32(16)
			return v100
		} else {
			if v13 == int32(0) {
				v21 = v7 + int32(12)
				v22 = int32(0)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v23 == v22 {
					v40 = int32(0)
					if v21 == v40 {
						v48 = v40
					} else {
						v43 = v40
						v44 = v22
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v43
						v48 = v44
					}
					v51 = v48
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					switch v26 - int32(429) {
					case 0:
						if v21 == int32(0) {
							v51 = int32(1)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
							v43 = v33
							v44 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v43
							v48 = v44
							v51 = v48
						}
					case 1:
						if v21 == int32(0) {
							v51 = int32(2)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+368))
							v43 = v38
							v44 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v43
							v48 = v44
							v51 = v48
						}
					default:
						v40 = int32(0)
						if v21 == v40 {
							v48 = v40
						} else {
							v43 = v40
							v44 = v22
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v43
							v48 = v44
						}
						v51 = v48
					}
				}
				if v51 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(59977), int32(0))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(484742), int32(3992), int32(346037))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v54 = int32(4464496)
					v55 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v57
					v60 = F_palloc0(m, int32(40))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v55
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v60))) = v66
						v68 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v68
						v70 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v70
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v72
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v74
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v76
						v100 = v60
						m.G0 = v7 + int32(16)
						return v100
					}
				}
			} else {
				v78 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = v78 + v79
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
				v83 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v82 + v83
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
				v87 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v86 + v87
				v90 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
				if v90 <= int64(0) {
					v100 = v13
					m.G0 = v7 + int32(16)
					return v100
				} else {
					v93 = int32(8)
					v94 = v13 + v93
					F_finite_interval_pl(m, v94, v15+v93, v94)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						v100 = v13
						m.G0 = v7 + int32(16)
						return v100
					}
				}
			}
		}
	}
}
func F_interval_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = F_DirectFunctionCall2Coll(m, int32(1475), v2, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_DirectFunctionCall2Coll(m, int32(2459), v2, v9, int32(4039552))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v19 = F_DirectFunctionCall1Coll(m, int32(2463), int32(0), v9)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = v19
					return v21
				}
			} else {
				v21 = v9
				return v21
			}
		}
	}
}
func F_interval_hash_extended(m *base.Module, l0 int32) int32 {
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
	var v12 int64
	_ = v12
	var v16 int64
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+12)))
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = (v9*int64(30)+v12)*int64(86400000000) + v16
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = F_DirectFunctionCall2Coll(m, int32(1285), int32(0), v6+int32(8), v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v24
	}
}
func F_interval_justify_days(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_palloc(m, int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v15
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v17
		if v13 != int32(2147483647) {
			if v13 != int32(-2147483648) {
				v32 = base.I32_div_s(v15, int32(30))
				v33 = v13 + v32
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
				v37 = v32*int32(-30) + v15
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37
				if base.B2i32(v32 < int32(0)) != base.B2i32(v33 < v13) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(393566), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(484742), int32(3081), int32(110600))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
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
					if int32(0) < v33 {
						if int32(0) <= v37 {
						} else {
							v55 = int32(-1)
							v56 = int32(30)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
						}
					} else {
						if v33 == int32(0) {
						} else {
							if v37 <= int32(0) {
							} else {
								v55 = int32(1)
								v56 = int32(-30)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
							}
						}
					}
					return v9
				}
			} else {
				if v15 != int32(-2147483648) {
					v32 = base.I32_div_s(v15, int32(30))
					v33 = v13 + v32
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
					v37 = v32*int32(-30) + v15
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37
					if base.B2i32(v32 < int32(0)) != base.B2i32(v33 < v13) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(393566), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(484742), int32(3081), int32(110600))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
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
						if int32(0) < v33 {
							if int32(0) <= v37 {
							} else {
								v55 = int32(-1)
								v56 = int32(30)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
							}
						} else {
							if v33 == int32(0) {
							} else {
								if v37 <= int32(0) {
								} else {
									v55 = int32(1)
									v56 = int32(-30)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
								}
							}
						}
						return v9
					}
				} else {
					if v17 != int64(-9223372036854775807-1) {
						v32 = base.I32_div_s(v15, int32(30))
						v33 = v13 + v32
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
						v37 = v32*int32(-30) + v15
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37
						if base.B2i32(v32 < int32(0)) != base.B2i32(v33 < v13) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(393566), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(484742), int32(3081), int32(110600))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
							if int32(0) < v33 {
								if int32(0) <= v37 {
								} else {
									v55 = int32(-1)
									v56 = int32(30)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
								}
							} else {
								if v33 == int32(0) {
								} else {
									if v37 <= int32(0) {
									} else {
										v55 = int32(1)
										v56 = int32(-30)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
										*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
									}
								}
							}
							return v9
						}
					} else {
						return v9
					}
				}
			}
		} else {
			if v15 != int32(2147483647) {
				v32 = base.I32_div_s(v15, int32(30))
				v33 = v13 + v32
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
				v37 = v32*int32(-30) + v15
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37
				if base.B2i32(v32 < int32(0)) != base.B2i32(v33 < v13) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(393566), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(484742), int32(3081), int32(110600))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
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
					if int32(0) < v33 {
						if int32(0) <= v37 {
						} else {
							v55 = int32(-1)
							v56 = int32(30)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
						}
					} else {
						if v33 == int32(0) {
						} else {
							if v37 <= int32(0) {
							} else {
								v55 = int32(1)
								v56 = int32(-30)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
							}
						}
					}
					return v9
				}
			} else {
				if v17 == int64(9223372036854775807) {
					return v9
				} else {
					v32 = base.I32_div_s(v15, int32(30))
					v33 = v13 + v32
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
					v37 = v32*int32(-30) + v15
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37
					if base.B2i32(v32 < int32(0)) != base.B2i32(v33 < v13) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(393566), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(484742), int32(3081), int32(110600))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
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
						if int32(0) < v33 {
							if int32(0) <= v37 {
							} else {
								v55 = int32(-1)
								v56 = int32(30)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
							}
						} else {
							if v33 == int32(0) {
							} else {
								if v37 <= int32(0) {
								} else {
									v55 = int32(1)
									v56 = int32(-30)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33 + v55
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v37 + v56
								}
							}
						}
						return v9
					}
				}
			}
		}
	}
}
func F_interval_justify_hours(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v32 int64
	_ = v32
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_palloc(m, int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v15
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v17
		if v13 != int32(2147483647) {
			if v13 != int32(-2147483648) {
				v32 = base.I64_div_s(v17, int64(86400000000))
				if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
					v39 = v32*int64(-86400000000) + v17
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v39
					v41 = v39
				} else {
					v41 = v17
				}
				v42 = base.I32_wrap_i64(v32)
				v43 = v15 + v42
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43
				if base.B2i32(v42 < int32(0)) != base.B2i32(v43 < v15) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(393566), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(484742), int32(3038), int32(127759))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
					if int32(0) < v43 {
						if int64(0) <= v41 {
						} else {
							v61 = int32(-1)
							v62 = int64(86400000000)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
						}
					} else {
						if v43 == int32(0) {
						} else {
							if v41 <= int64(0) {
							} else {
								v61 = int32(1)
								v62 = int64(-86400000000)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
							}
						}
					}
					return v9
				}
			} else {
				if v15 != int32(-2147483648) {
					v32 = base.I64_div_s(v17, int64(86400000000))
					if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
						v39 = v32*int64(-86400000000) + v17
						*(*int64)(unsafe.Add(mBase, uint32(v9))) = v39
						v41 = v39
					} else {
						v41 = v17
					}
					v42 = base.I32_wrap_i64(v32)
					v43 = v15 + v42
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43
					if base.B2i32(v42 < int32(0)) != base.B2i32(v43 < v15) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(393566), int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(484742), int32(3038), int32(127759))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
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
						if int32(0) < v43 {
							if int64(0) <= v41 {
							} else {
								v61 = int32(-1)
								v62 = int64(86400000000)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
							}
						} else {
							if v43 == int32(0) {
							} else {
								if v41 <= int64(0) {
								} else {
									v61 = int32(1)
									v62 = int64(-86400000000)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
									*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
								}
							}
						}
						return v9
					}
				} else {
					if v17 != int64(-9223372036854775807-1) {
						v32 = base.I64_div_s(v17, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
							v39 = v32*int64(-86400000000) + v17
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v39
							v41 = v39
						} else {
							v41 = v17
						}
						v42 = base.I32_wrap_i64(v32)
						v43 = v15 + v42
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43
						if base.B2i32(v42 < int32(0)) != base.B2i32(v43 < v15) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(393566), int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(484742), int32(3038), int32(127759))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
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
							if int32(0) < v43 {
								if int64(0) <= v41 {
								} else {
									v61 = int32(-1)
									v62 = int64(86400000000)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
									*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
								}
							} else {
								if v43 == int32(0) {
								} else {
									if v41 <= int64(0) {
									} else {
										v61 = int32(1)
										v62 = int64(-86400000000)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
										*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
									}
								}
							}
							return v9
						}
					} else {
						return v9
					}
				}
			}
		} else {
			if v15 != int32(2147483647) {
				v32 = base.I64_div_s(v17, int64(86400000000))
				if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
					v39 = v32*int64(-86400000000) + v17
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v39
					v41 = v39
				} else {
					v41 = v17
				}
				v42 = base.I32_wrap_i64(v32)
				v43 = v15 + v42
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43
				if base.B2i32(v42 < int32(0)) != base.B2i32(v43 < v15) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(393566), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(484742), int32(3038), int32(127759))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
					if int32(0) < v43 {
						if int64(0) <= v41 {
						} else {
							v61 = int32(-1)
							v62 = int64(86400000000)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
						}
					} else {
						if v43 == int32(0) {
						} else {
							if v41 <= int64(0) {
							} else {
								v61 = int32(1)
								v62 = int64(-86400000000)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
							}
						}
					}
					return v9
				}
			} else {
				if v17 == int64(9223372036854775807) {
					return v9
				} else {
					v32 = base.I64_div_s(v17, int64(86400000000))
					if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
						v39 = v32*int64(-86400000000) + v17
						*(*int64)(unsafe.Add(mBase, uint32(v9))) = v39
						v41 = v39
					} else {
						v41 = v17
					}
					v42 = base.I32_wrap_i64(v32)
					v43 = v15 + v42
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43
					if base.B2i32(v42 < int32(0)) != base.B2i32(v43 < v15) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(393566), int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(484742), int32(3038), int32(127759))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
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
						if int32(0) < v43 {
							if int64(0) <= v41 {
							} else {
								v61 = int32(-1)
								v62 = int64(86400000000)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
							}
						} else {
							if v43 == int32(0) {
							} else {
								if v41 <= int64(0) {
								} else {
									v61 = int32(1)
									v62 = int64(-86400000000)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v43 + v61
									*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41 + v62
								}
							}
						}
						return v9
					}
				}
			}
		}
	}
}
func F_interval_scale(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v13
		F_AdjustIntervalForTypmod(m, v7, v4, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_interval_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		F_enlargeStringInfo(m, v8, int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v22 = int64(56)
			v24 = int64(65280)
			v26 = int64(40)
			v29 = int64(16711680)
			v31 = int64(24)
			v33 = int64(4278190080)
			v35 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v19+v20))) = v15<<(uint(v22)%64) | v15&v24<<(uint(v26)%64) | (v15&v29<<(uint(v31)%64) | v15&v33<<(uint(v35)%64)) | (int64(base.Ui64(v15)>>(uint(v35)%64))&v33 | int64(base.Ui64(v15)>>(uint(v31)%64))&v29 | (int64(base.Ui64(v15)>>(uint(v26)%64))&v24 | int64(base.Ui64(v15)>>(uint(v22)%64))))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v19 + int32(8)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			F_enlargeStringInfo(m, v8, int32(4))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v68 = int32(24)
				v70 = int32(65280)
				v72 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v65+v66))) = v61<<(uint(v68)%32) | v61&v70<<(uint(v72)%32) | (int32(base.Ui32(v61)>>(uint(v72)%32))&v70 | int32(base.Ui32(v61)>>(uint(v68)%32)))
				v84 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v65 + v84
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				F_enlargeStringInfo(m, v8, v84)
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v94 = int32(24)
					v96 = int32(65280)
					v98 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v91+v92))) = v87<<(uint(v94)%32) | v87&v96<<(uint(v98)%32) | (int32(base.Ui32(v87)>>(uint(v98)%32))&v96 | int32(base.Ui32(v87)>>(uint(v94)%32)))
					v111 = v91 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v111
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					*(*int32)(unsafe.Add(mBase, uint32(v114))) = v111 << (uint(int32(2)) % 32)
					m.G0 = v8 + int32(16)
					return v114
				}
			}
		}
	}
}
func F_interval_um(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_palloc(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_interval_um_internal(m, v2, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v4
		}
	}
}
