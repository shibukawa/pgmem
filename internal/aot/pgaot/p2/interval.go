package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AdjustIntervalForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v210 int32
	_ = v210
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
	return v210
L2:
	;
	v42 = int32(base.Ui32(l1) >> (uint(int32(16)) % 32))
	if base.Ui32(v42) <= base.Ui32(int32(2047)) {
		goto L27
	} else {
		goto L28
	}
L3:
	;
	if int32(0) <= l1 {
		goto L2
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v28 != int32(2147483647) {
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
	v22 = int32(1)
	if l1 < int32(0) {
		v210 = v22
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v25 != int64(-9223372036854775807-1) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v210 = v22
	goto L1
L11:
	;
	v31 = int32(1)
	if l1 < int32(0) {
		v210 = v31
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v34 == int64(9223372036854775807) {
		v210 = v31
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L2
L14:
	;
	v210 = int32(1)
	goto L1
L15:
	;
	v116 = int32(1)
	v117 = int32(_a_F_AdjustIntervalForTypmod_0)
	v118 = l1 & v117
	if v118 == v117 {
		v210 = v116
		goto L1
	} else {
		goto L42
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v113
	goto L15
L17:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v111 = base.I64_rem_s(v109, int64(60000000))
	v113 = v109 - v111
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v106 = base.I32_rem_s(v14, int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14 - v106
	v113 = v4
	goto L16
L19:
	;
	if l1&int32(2080309248) == int32(402653184) {
		goto L15
	} else {
		goto L37
	}
L20:
	;
	if v42 == int32(2048) {
		goto L17
	} else {
		goto L36
	}
L21:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v78 = base.I64_rem_s(v76, int64(60000000))
	v113 = v76 - v78
	goto L16
L22:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v74 = base.I64_rem_s(v72, int64(60000000))
	v113 = v72 - v74
	goto L16
L23:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v70 = base.I64_rem_s(v68, int64(3600000000))
	v113 = v68 - v70
	goto L16
L24:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v66 = base.I64_rem_s(v64, int64(3600000000))
	v113 = v64 - v66
	goto L16
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v113 = v4
	goto L16
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v113 = v4
	goto L16
L27:
	;
	switch v42 - int32(2) {
	case 0:
		goto L26
	case 1, 3, 5:
		goto L19
	case 2:
		goto L18
	case 4:
		goto L25
	case 6:
		v113 = v4
		goto L16
	default:
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(v42) <= base.Ui32(int32(4095)) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	switch v42 - int32(1024) {
	case 0:
		goto L24
	default:
		goto L19
	case 8:
		goto L23
	}
L31:
	;
	switch v42 - int32(3072) {
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
	if base.B2i32(v42 == int32(_a_F_AdjustIntervalForTypmod_6))|base.B2i32(v42 == int32(_a_F_AdjustIntervalForTypmod_7)) != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	if v42 != int32(_a_F_AdjustIntervalForTypmod_8) {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	goto L15
L36:
	;
	goto L19
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_AdjustIntervalForTypmod_5), v12+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_AdjustIntervalForTypmod_2), int32(1492), int32(_a_F_AdjustIntervalForTypmod_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	if base.Ui32(int32(7)) <= base.Ui32(v118) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v123 = int32(0)
	v124 = F_errsave_start(m, l2)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L38
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v118<<(uint(int32(3))%32))+uint32(_c_F_AdjustIntervalForTypmod[0])))
	v145 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if int64(0) <= v145 {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	if v124 == int32(0) {
		v210 = v123
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L38
	} else {
		goto L48
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(25769803776)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v118
	F_errmsg(m, int32(_a_F_AdjustIntervalForTypmod_1), v12)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, l2, int32(_a_F_AdjustIntervalForTypmod_2), int32(1501), int32(_a_F_AdjustIntervalForTypmod_3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L38
	} else {
		goto L50
	}
L50:
	;
	v210 = v123
	goto L1
L51:
	;
	v148 = v144 + v145
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v148
	if base.B2i32(v144 < int64(0)) != base.B2i32(v148 < v145) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v177 = v145 - v144
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v177
	if base.B2i32(v177 < v145) != base.B2i32(int64(0) < v144) {
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v154 = int32(0)
	v155 = F_errsave_start(m, l2)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L38
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v118<<(uint(int32(3))%32))+uint32(_c_F_AdjustIntervalForTypmod[1])))
	v174 = base.I64_rem_s(v148, v173)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v148 - v174
	v210 = v116
	goto L1
L57:
	;
	if v155 == int32(0) {
		v210 = v154
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L38
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_AdjustIntervalForTypmod_4), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L38
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, l2, int32(_a_F_AdjustIntervalForTypmod_2), int32(1510), int32(_a_F_AdjustIntervalForTypmod_3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L38
	} else {
		goto L61
	}
L61:
	;
	v210 = v154
	goto L1
L62:
	;
	v183 = int32(0)
	v184 = F_errsave_start(m, l2)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L38
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v118<<(uint(int32(3))%32))+uint32(_c_F_AdjustIntervalForTypmod[1])))
	v203 = base.I64_rem_s(v177, v202)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v177 - v203
	v210 = v116
	goto L1
L65:
	;
	if v184 == int32(0) {
		v210 = v183
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L38
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_AdjustIntervalForTypmod_4), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L38
	} else {
		goto L68
	}
L68:
	;
	F_errsave_finish(m, l2, int32(_a_F_AdjustIntervalForTypmod_2), int32(1520), int32(_a_F_AdjustIntervalForTypmod_3))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L38
	} else {
		goto L69
	}
L69:
	;
	v210 = v183
	goto L1
}
func F_interval_avg(m *base.Module, l0 int32) int64 {
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
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7 != 0 {
		v22 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v8 == int32(0) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			return int64(0)
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
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_interval_avg_0), int32(0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_interval_avg_1), int32(_a_F_interval_avg_2), int32(_a_F_interval_avg_3))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
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
						v39 = F_palloc(m, int32(16))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
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
							return base.I64_extend_i32_u(v39)
						}
					}
				} else {
					v64 = F_DirectFunctionCall2Coll(m, int32(1713), int32(0), base.I64_extend_i32_u(v8+int32(8)), base.I64_reinterpret_f64(base.F64_convert_i64_s(v12)))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int64(0)
					} else {
						return v64
					}
				}
			} else {
				v22 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
				return int64(0)
			}
		}
	}
}
func F_interval_avg_combine(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
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
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v9 == v2 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = v12
	} else {
		v13 = v2
	}
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v14 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v17 != 0 {
			if v13 == int32(0) {
				v25 = v7 + int32(12)
				v26 = int32(0)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v27 == v26 {
					v44 = int32(0)
					if v25 == v44 {
						v52 = v44
					} else {
						v47 = v44
						v48 = v26
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = v47
						v52 = v48
					}
					v55 = v52
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					switch v30 - int32(435) {
					case 0:
						if v25 == int32(0) {
							v55 = int32(1)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+168))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
							v47 = v37
							v48 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v47
							v52 = v48
							v55 = v52
						}
					case 1:
						if v25 == int32(0) {
							v55 = int32(2)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v27)+376))
							v47 = v42
							v48 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v47
							v52 = v48
							v55 = v52
						}
					default:
						v44 = int32(0)
						if v25 == v44 {
							v52 = v44
						} else {
							v47 = v44
							v48 = v26
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v47
							v52 = v48
						}
						v55 = v52
					}
				}
				if v55 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_interval_avg_combine_0), int32(0))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_interval_avg_combine_1), int32(3993), int32(_a_F_interval_avg_combine_2))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v58 = int32(_a_F_interval_avg_combine_3)
					v59 = *(*int32)(unsafe.Add(mBase, _c_F_interval_avg_combine[0]))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_combine[0])) = v61
					v64 = F_palloc0(m, int32(40))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_combine[0])) = v59
						v70 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
						*(*int64)(unsafe.Add(mBase, uint32(v64))) = v70
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v64)+24)) = v72
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v76
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = v78
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = v80
						v104 = v64
						m.G0 = v7 + int32(16)
						return base.I64_extend_i32_u(v104)
					}
				}
			} else {
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
				v83 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = v82 + v83
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
				v87 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v86 + v87
				v90 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
				v91 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v90 + v91
				v94 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				if v94 <= int64(0) {
					v104 = v13
					m.G0 = v7 + int32(16)
					return base.I64_extend_i32_u(v104)
				} else {
					v97 = int32(8)
					v98 = v13 + v97
					F_finite_interval_pl(m, v98, v17+v97, v98)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						v104 = v13
						m.G0 = v7 + int32(16)
						return base.I64_extend_i32_u(v104)
					}
				}
			}
		} else {
			if v13 != 0 {
				v104 = v13
			} else {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				v104 = int32(0)
			}
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v104)
		}
	} else {
		if v13 != 0 {
			v104 = v13
		} else {
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v104 = int32(0)
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v104)
	}
}
func F_interval_dist(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	v4 = int32(0)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = F_DirectFunctionCall2Coll(m, int32(1606), v4, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v14 = v9 & int64(4294967295)
		v16 = F_DirectFunctionCall2Coll(m, int32(2657), v4, v14, int64(4165696))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			if v16 != int64(0) {
				v22 = F_DirectFunctionCall1Coll(m, int32(2661), int32(0), v14)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					v26 = v22 & int64(4294967295)
					return v26
				}
			} else {
				v26 = v14
				return v26
			}
		}
	}
}
func F_interval_hash_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+12)))
	v9 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+8)))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_DirectFunctionCall2Coll(m, int32(1398), int32(0), (v6*int64(30)+v9)*int64(86400000000)+v13, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		return v16
	}
}
func F_interval_justify_days(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_palloc(m, int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v15
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v17
		if v13 != int32(2147483647) {
			v21 = int32(-2147483648)
			if base.B2i32(v13 != v21)|base.B2i32(v15 != v21)|base.B2i32(v17 != int64(-9223372036854775807-1)) != 0 {
				v34 = base.I32_div_s(v15, int32(30))
				v35 = v13 + v34
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35
				v39 = v34*int32(-30) + v15
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39
				if base.B2i32(v34 < int32(0)) != base.B2i32(v35 < v13) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_justify_days_0), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_justify_days_1), int32(3076), int32(_a_F_interval_justify_days_2))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
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
					if int32(0) < v35 {
						if int32(0) <= v39 {
						} else {
							v58 = int32(-1)
							v59 = int32(30)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35 + v58
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39 + v59
						}
					} else {
						v52 = int32(0)
						if base.B2i32(v35 == v52)|base.B2i32(v39 <= v52) != 0 {
						} else {
							v58 = int32(1)
							v59 = int32(-30)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35 + v58
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39 + v59
						}
					}
					return base.I64_extend_i32_u(v9)
				}
			} else {
				return base.I64_extend_i32_u(v9)
			}
		} else {
			if v15 != int32(2147483647) {
				v34 = base.I32_div_s(v15, int32(30))
				v35 = v13 + v34
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35
				v39 = v34*int32(-30) + v15
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39
				if base.B2i32(v34 < int32(0)) != base.B2i32(v35 < v13) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_justify_days_0), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_justify_days_1), int32(3076), int32(_a_F_interval_justify_days_2))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
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
					if int32(0) < v35 {
						if int32(0) <= v39 {
						} else {
							v58 = int32(-1)
							v59 = int32(30)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35 + v58
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39 + v59
						}
					} else {
						v52 = int32(0)
						if base.B2i32(v35 == v52)|base.B2i32(v39 <= v52) != 0 {
						} else {
							v58 = int32(1)
							v59 = int32(-30)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35 + v58
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39 + v59
						}
					}
					return base.I64_extend_i32_u(v9)
				}
			} else {
				if v17 == int64(9223372036854775807) {
					return base.I64_extend_i32_u(v9)
				} else {
					v34 = base.I32_div_s(v15, int32(30))
					v35 = v13 + v34
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35
					v39 = v34*int32(-30) + v15
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39
					if base.B2i32(v34 < int32(0)) != base.B2i32(v35 < v13) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_interval_justify_days_0), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_interval_justify_days_1), int32(3076), int32(_a_F_interval_justify_days_2))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
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
						if int32(0) < v35 {
							if int32(0) <= v39 {
							} else {
								v58 = int32(-1)
								v59 = int32(30)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35 + v58
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39 + v59
							}
						} else {
							v52 = int32(0)
							if base.B2i32(v35 == v52)|base.B2i32(v39 <= v52) != 0 {
							} else {
								v58 = int32(1)
								v59 = int32(-30)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35 + v58
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39 + v59
							}
						}
						return base.I64_extend_i32_u(v9)
					}
				}
			}
		}
	}
}
func F_interval_justify_hours(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v34 int64
	_ = v34
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_palloc(m, int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v15
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v17
		if v13 != int32(2147483647) {
			v21 = int32(-2147483648)
			if base.B2i32(v13 != v21)|base.B2i32(v15 != v21)|base.B2i32(v17 != int64(-9223372036854775807-1)) != 0 {
				v34 = base.I64_div_s(v17, int64(86400000000))
				if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
					v41 = v34*int64(-86400000000) + v17
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41
					v43 = v41
				} else {
					v43 = v17
				}
				v45 = v15 + base.I32_wrap_i64(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45
				if base.B2i32(v45 < v15) != base.B2i32(v34 < int64(0)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_justify_hours_0), int32(0))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_justify_hours_1), int32(3033), int32(_a_F_interval_justify_hours_2))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
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
					if int32(0) < v45 {
						if int64(0) <= v43 {
						} else {
							v64 = int32(-1)
							v65 = int64(86400000000)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45 + v64
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v43 + v65
						}
					} else {
						if base.B2i32(v45 == int32(0))|base.B2i32(v43 <= int64(0)) != 0 {
						} else {
							v64 = int32(1)
							v65 = int64(-86400000000)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45 + v64
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v43 + v65
						}
					}
					return base.I64_extend_i32_u(v9)
				}
			} else {
				return base.I64_extend_i32_u(v9)
			}
		} else {
			if v15 != int32(2147483647) {
				v34 = base.I64_div_s(v17, int64(86400000000))
				if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
					v41 = v34*int64(-86400000000) + v17
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41
					v43 = v41
				} else {
					v43 = v17
				}
				v45 = v15 + base.I32_wrap_i64(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45
				if base.B2i32(v45 < v15) != base.B2i32(v34 < int64(0)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_justify_hours_0), int32(0))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_justify_hours_1), int32(3033), int32(_a_F_interval_justify_hours_2))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
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
					if int32(0) < v45 {
						if int64(0) <= v43 {
						} else {
							v64 = int32(-1)
							v65 = int64(86400000000)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45 + v64
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v43 + v65
						}
					} else {
						if base.B2i32(v45 == int32(0))|base.B2i32(v43 <= int64(0)) != 0 {
						} else {
							v64 = int32(1)
							v65 = int64(-86400000000)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45 + v64
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v43 + v65
						}
					}
					return base.I64_extend_i32_u(v9)
				}
			} else {
				if v17 == int64(9223372036854775807) {
					return base.I64_extend_i32_u(v9)
				} else {
					v34 = base.I64_div_s(v17, int64(86400000000))
					if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
						v41 = v34*int64(-86400000000) + v17
						*(*int64)(unsafe.Add(mBase, uint32(v9))) = v41
						v43 = v41
					} else {
						v43 = v17
					}
					v45 = v15 + base.I32_wrap_i64(v34)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45
					if base.B2i32(v45 < v15) != base.B2i32(v34 < int64(0)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_interval_justify_hours_0), int32(0))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_interval_justify_hours_1), int32(3033), int32(_a_F_interval_justify_hours_2))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
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
						if int32(0) < v45 {
							if int64(0) <= v43 {
							} else {
								v64 = int32(-1)
								v65 = int64(86400000000)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45 + v64
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v43 + v65
							}
						} else {
							if base.B2i32(v45 == int32(0))|base.B2i32(v43 <= int64(0)) != 0 {
							} else {
								v64 = int32(1)
								v65 = int64(-86400000000)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v45 + v64
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v43 + v65
							}
						}
						return base.I64_extend_i32_u(v9)
					}
				}
			}
		}
	}
}
func F_interval_scale(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v17 = F_AdjustIntervalForTypmod(m, v8, v5, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if v17 == int32(0) {
				v21 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v8)
			}
		}
	}
}
func F_interval_send(m *base.Module, l0 int32) int64 {
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
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		F_enlargeStringInfo(m, v8, int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
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
				return int64(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v70 = int32(16711935)
				*(*int32)(unsafe.Add(mBase, uint32(v65+v66))) = base.I32_rotr(v61, int32(24))&v70 | base.I32_rotr(v61&v70, int32(8))
				v78 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v65 + v78
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				F_enlargeStringInfo(m, v8, v78)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v90 = int32(16711935)
					*(*int32)(unsafe.Add(mBase, uint32(v85+v86))) = base.I32_rotr(v81, int32(24))&v90 | base.I32_rotr(v81&v90, int32(8))
					v99 = v85 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v99
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					*(*int32)(unsafe.Add(mBase, uint32(v102))) = v99 << (uint(int32(2)) % 32)
					m.G0 = v8 + int32(16)
					return base.I64_extend_i32_u(v102)
				}
			}
		}
	}
}
func F_interval_um(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_palloc(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		F_interval_um_internal(m, v2, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v4)
		}
	}
}
