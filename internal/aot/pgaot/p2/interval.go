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
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
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
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
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
	v110 = int32(_a_F_AdjustIntervalForTypmod_0)
	v111 = l1 & v110
	if v111 == v110 {
		goto L1
	} else {
		goto L42
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
	goto L15
L17:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v105 = base.I64_rem_s(v103, int64(60000000))
	v107 = v103 - v105
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v100 = base.I32_rem_s(v14, int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14 - v100
	v107 = v4
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
	if v38 == int32(2048) {
		goto L17
	} else {
		goto L36
	}
L21:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v74 = base.I64_rem_s(v72, int64(60000000))
	v107 = v72 - v74
	goto L16
L22:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v70 = base.I64_rem_s(v68, int64(60000000))
	v107 = v68 - v70
	goto L16
L23:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v66 = base.I64_rem_s(v64, int64(3600000000))
	v107 = v64 - v66
	goto L16
L24:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v62 = base.I64_rem_s(v60, int64(3600000000))
	v107 = v60 - v62
	goto L16
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v107 = v4
	goto L16
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v107 = v4
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
		v107 = v4
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
	if base.B2i32(v38 == int32(_a_F_AdjustIntervalForTypmod_6))|base.B2i32(v38 == int32(_a_F_AdjustIntervalForTypmod_7)) != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	if v38 != int32(_a_F_AdjustIntervalForTypmod_8) {
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
	v85 = m.ExcPending
	if v85 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_AdjustIntervalForTypmod_5), v12+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_AdjustIntervalForTypmod_2), int32(1489), int32(_a_F_AdjustIntervalForTypmod_3))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
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
	if base.Ui32(int32(7)) <= base.Ui32(v111) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v116 = F_errsave_start(m, l2)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L38
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v111<<(uint(int32(3))%32))+uint32(_c_F_AdjustIntervalForTypmod[0])))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if int64(0) <= v137 {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	if v116 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L38
	} else {
		goto L48
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(25769803776)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
	F_errmsg(m, int32(_a_F_AdjustIntervalForTypmod_1), v12)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, l2, int32(_a_F_AdjustIntervalForTypmod_2), int32(1498), int32(_a_F_AdjustIntervalForTypmod_3))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L38
	} else {
		goto L50
	}
L50:
	;
	goto L1
L51:
	;
	v140 = v136 + v137
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v140
	if base.B2i32(v136 < int64(0)) != base.B2i32(v140 < v137) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v168 = v137 - v136
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v168
	if base.B2i32(v168 < v137) != base.B2i32(int64(0) < v136) {
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v146 = F_errsave_start(m, l2)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L38
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v111<<(uint(int32(3))%32))+uint32(_c_F_AdjustIntervalForTypmod[1])))
	v165 = base.I64_rem_s(v140, v164)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v140 - v165
	goto L1
L57:
	;
	if v146 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L38
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_AdjustIntervalForTypmod_4), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L38
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, l2, int32(_a_F_AdjustIntervalForTypmod_2), int32(1507), int32(_a_F_AdjustIntervalForTypmod_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L38
	} else {
		goto L61
	}
L61:
	;
	goto L1
L62:
	;
	v174 = F_errsave_start(m, l2)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L38
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v111<<(uint(int32(3))%32))+uint32(_c_F_AdjustIntervalForTypmod[1])))
	v193 = base.I64_rem_s(v168, v192)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v168 - v193
	goto L1
L65:
	;
	if v174 == int32(0) {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L38
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_AdjustIntervalForTypmod_4), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L38
	} else {
		goto L68
	}
L68:
	;
	F_errsave_finish(m, l2, int32(_a_F_AdjustIntervalForTypmod_2), int32(1517), int32(_a_F_AdjustIntervalForTypmod_3))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L38
	} else {
		goto L69
	}
L69:
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
								F_errmsg(m, int32(_a_F_interval_avg_0), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_interval_avg_1), int32(_a_F_interval_avg_2), int32(_a_F_interval_avg_3))
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
						v63 = F_DirectFunctionCall2Coll(m, int32(1500), int32(0), v8+int32(8), v61)
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
						F_errmsg_internal(m, int32(_a_F_interval_avg_combine_0), int32(0))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_interval_avg_combine_1), int32(3992), int32(_a_F_interval_avg_combine_2))
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
					v54 = int32(_a_F_interval_avg_combine_3)
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_interval_avg_combine[0]))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_combine[0])) = v57
					v60 = F_palloc0(m, int32(40))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_combine[0])) = v55
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_DirectFunctionCall2Coll(m, int32(1460), v3, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = F_DirectFunctionCall2Coll(m, int32(2444), v3, v8, int32(_a_F_interval_dist_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v13 != 0 {
				v17 = F_DirectFunctionCall1Coll(m, int32(2448), int32(0), v8)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = v17
					return v19
				}
			} else {
				v19 = v8
				return v19
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
	v24 = F_DirectFunctionCall2Coll(m, int32(1270), int32(0), v6+int32(8), v23)
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
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
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_justify_days_0), int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_justify_days_1), int32(3081), int32(_a_F_interval_justify_days_2))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
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
					return v9
				}
			} else {
				return v9
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
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_justify_days_0), int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_justify_days_1), int32(3081), int32(_a_F_interval_justify_days_2))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
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
					return v9
				}
			} else {
				if v17 == int64(9223372036854775807) {
					return v9
				} else {
					v34 = base.I32_div_s(v15, int32(30))
					v35 = v13 + v34
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35
					v39 = v34*int32(-30) + v15
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v39
					if base.B2i32(v34 < int32(0)) != base.B2i32(v35 < v13) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_interval_justify_days_0), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_interval_justify_days_1), int32(3081), int32(_a_F_interval_justify_days_2))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
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
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_justify_hours_0), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_justify_hours_1), int32(3038), int32(_a_F_interval_justify_hours_2))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
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
					return v9
				}
			} else {
				return v9
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
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_justify_hours_0), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_justify_hours_1), int32(3038), int32(_a_F_interval_justify_hours_2))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
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
					return v9
				}
			} else {
				if v17 == int64(9223372036854775807) {
					return v9
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
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_interval_justify_hours_0), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_interval_justify_hours_1), int32(3038), int32(_a_F_interval_justify_hours_2))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
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
				v70 = int32(16711935)
				*(*int32)(unsafe.Add(mBase, uint32(v65+v66))) = base.I32_rotr(v61, int32(24))&v70 | base.I32_rotr(v61&v70, int32(8))
				v78 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v65 + v78
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				F_enlargeStringInfo(m, v8, v78)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
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
					return v102
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
