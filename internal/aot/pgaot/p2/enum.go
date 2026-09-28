package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enum_cmp_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 float32
	_ = v167
	var v168 float32
	_ = v168
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l0 == l1 {
		v209 = int32(0)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L11
	} else {
		goto L64
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v209
L3:
	;
	if base.Ui32(l0) < base.Ui32(l1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = int32(-1)
	goto L6
L5:
	;
	v21 = int32(1)
	goto L6
L6:
	;
	if (l0|l1)&int32(1) == int32(0) {
		v209 = v21
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = F_SearchSysCache1(m, int32(23), base.I64_extend_i32_u(l0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v51 = v28
	goto L10
L10:
	;
	v52 = m.G0
	v54 = v52 - int32(32)
	m.G0 = v54
	if l0 == l1 {
		v177 = int32(0)
		goto L18
	} else {
		goto L19
	}
L11:
	;
	return int32(0)
L12:
	;
	if v33 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39+v40)+4))
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v46 = F_lookup_type_cache(m, v42, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v46
	v51 = v46
	goto L10
L16:
	;
	v209 = v177
	goto L2
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L60
	}
L18:
	;
	m.G0 = v54 + int32(32)
	goto L16
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+316))
	if v58 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_load_enum_cache_data(m, v51)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	v64 = v58
	goto L22
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if base.Ui32(l0) < base.Ui32(v65) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+316))
	v64 = v63
	goto L22
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v90 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	v67 = l0 - v65
	if v67 < int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v71 = F_bms_is_member(m, v67, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	if v71 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if base.Ui32(l1) < base.Ui32(v75) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v77 = l1 - v75
	if v77 < int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v81 = F_bms_is_member(m, v77, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	if v81 == int32(0) {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	if base.Ui32(l0) < base.Ui32(l1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v88 = int32(-1)
	goto L35
L34:
	;
	v88 = int32(1)
	goto L35
L35:
	;
	v177 = v88
	goto L18
L36:
	;
	v167 = *(*float32)(unsafe.Add(mBase, uint32(v163)+4))
	v168 = *(*float32)(unsafe.Add(mBase, uint32(v162)+4))
	if base.F32_lt(v167, v168) != 0 {
		v177 = int32(-1)
		goto L18
	} else {
		goto L59
	}
L37:
	;
	F_load_enum_cache_data(m, v51)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L11
	} else {
		goto L44
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = l0
	v95 = v54 + int32(24)
	v97 = v64 + int32(12)
	v100 = F_bsearch(m, v95, v97, v90, int32(8), int32(1822))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v102 <= int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = l1
	v108 = F_bsearch(m, v95, v97, v102, int32(8), int32(1822))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	if v100 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if v108 != 0 {
		v162 = v108
		v163 = v100
		goto L36
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v51)+316))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	if v119 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L11
	} else {
		goto L55
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = l0
	v124 = v54 + int32(24)
	v126 = v118 + int32(12)
	v129 = F_bsearch(m, v124, v126, v119, int32(8), int32(1822))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	if int32(0) < v131 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = l1
	v137 = F_bsearch(m, v124, v126, v131, int32(8), int32(1822))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v129 != 0 {
		goto L17
	} else {
		goto L54
	}
L51:
	;
	if v129 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	if v137 != 0 {
		v162 = v137
		v163 = v129
		goto L36
	} else {
		goto L53
	}
L53:
	;
	goto L17
L54:
	;
	goto L45
L55:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v150 = F_format_type_be(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = l0
	F_errmsg_internal(m, int32(_a_F_enum_cmp_internal_3), v54)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_enum_cmp_internal_4), int32(2746), int32(_a_F_enum_cmp_internal_5))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v177 = base.F32_gt(v167, v168)
	goto L18
L60:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v187 = F_format_type_be(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_enum_cmp_internal_3), v54+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_enum_cmp_internal_4), int32(2749), int32(_a_F_enum_cmp_internal_5))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(_a_F_enum_cmp_internal_0), v14)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_enum_cmp_internal_1), int32(292), int32(_a_F_enum_cmp_internal_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_enum_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = base.I32_wrap_i64(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v18 = F_pq_getmsgtext(m, v12, v13-v14, v8+int32(28))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = F_strlen(m, v18)
		mBase = m.M
		if base.Ui32(v22) < base.Ui32(int32(64)) {
			v29 = F_SearchSysCache2(m, int32(24), v10&int64(4294967295), base.I64_extend_i32_u(v18))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				if v29 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int64(0)
						} else {
							v73 = F_format_type_be(m, v11)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v18
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v73
								F_errmsg(m, int32(_a_F_enum_recv_0), v8+int32(16))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_enum_recv_1), int32(206), int32(_a_F_enum_recv_2))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					F_check_safe_enum_use(m, v29)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
						v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35+v36))))
						F_ReleaseCatCache(m, v29)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v18)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								m.G0 = v8 + int32(32)
								return v38
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					v54 = F_format_type_be(m, v11)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v18
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v54
						F_errmsg(m, int32(_a_F_enum_recv_0), v8)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_enum_recv_1), int32(196), int32(_a_F_enum_recv_2))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
