package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enum_cmp_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
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
	var v166 float32
	_ = v166
	var v167 float32
	_ = v167
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == l1 {
		v207 = int32(0)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L64
	}
L2:
	;
	m.G0 = v13 + int32(16)
	return v207
L3:
	;
	if base.Ui32(l0) < base.Ui32(l1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int32(-1)
	goto L6
L5:
	;
	v20 = int32(1)
	goto L6
L6:
	;
	if (l0|l1)&int32(1) == int32(0) {
		v207 = v20
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = F_SearchSysCache1(m, int32(23), l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v49 = v27
	goto L10
L10:
	;
	v50 = m.G0
	v52 = v50 - int32(32)
	m.G0 = v52
	if l0 == l1 {
		v175 = int32(0)
		goto L18
	} else {
		goto L19
	}
L11:
	;
	return int32(0)
L12:
	;
	if v31 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)+4))
	F_ReleaseCatCache(m, v31)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v44 = F_lookup_type_cache(m, v40, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v44
	v49 = v44
	goto L10
L16:
	;
	v207 = v175
	goto L2
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L11
	} else {
		goto L60
	}
L18:
	;
	m.G0 = v52 + int32(32)
	goto L16
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+316))
	if v56 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_load_enum_cache_data(m, v49)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	v62 = v56
	goto L22
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if base.Ui32(l0) < base.Ui32(v63) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v49)+316))
	v62 = v61
	goto L22
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v88 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	v65 = l0 - v63
	if v65 < int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v69 = F_bms_is_member(m, v65, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	if v69 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if base.Ui32(l1) < base.Ui32(v73) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v75 = l1 - v73
	if v75 < int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v79 = F_bms_is_member(m, v75, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	if v79 == int32(0) {
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
	v86 = int32(-1)
	goto L35
L34:
	;
	v86 = int32(1)
	goto L35
L35:
	;
	v175 = v86
	goto L18
L36:
	;
	v166 = *(*float32)(unsafe.Add(mBase, uint32(v163)+4))
	v167 = *(*float32)(unsafe.Add(mBase, uint32(v162)+4))
	if base.F32_lt(v166, v167) != 0 {
		v175 = int32(-1)
		goto L18
	} else {
		goto L59
	}
L37:
	;
	F_load_enum_cache_data(m, v49)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L11
	} else {
		goto L44
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = l0
	v95 = v62 + int32(12)
	v98 = F_bsearch(m, v52+int32(24), v95, v88, int32(8), int32(1627))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v100 <= int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = l1
	v108 = F_bsearch(m, v52+int32(24), v95, v100, int32(8), int32(1627))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	if v98 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if v108 != 0 {
		v162 = v108
		v163 = v98
		goto L36
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v49)+316))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if v118 <= int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = l0
	v125 = v117 + int32(12)
	v128 = F_bsearch(m, v52+int32(24), v125, v118, int32(8), int32(1627))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if int32(0) < v130 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = l1
	v138 = F_bsearch(m, v52+int32(24), v125, v130, int32(8), int32(1627))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v128 != 0 {
		goto L17
	} else {
		goto L54
	}
L51:
	;
	if v128 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	if v138 != 0 {
		v162 = v138
		v163 = v128
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
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = l0
	F_errmsg_internal(m, int32(196650), v52)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(525524), int32(2719), int32(301828))
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
	v175 = base.F32_gt(v166, v167)
	goto L18
L60:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v185 = F_format_type_be(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = l1
	F_errmsg_internal(m, int32(196650), v52+int32(16))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(525524), int32(2722), int32(301828))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
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
	v218 = m.ExcPending
	if v218 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg(m, int32(63735), v13)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(523322), int32(292), int32(327321))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
func F_enum_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v16 = F_pq_getmsgtext(m, v10, v11-v12, v7+int32(28))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = F_strlen(m, v16)
		mBase = m.M
		if base.Ui32(v20) < base.Ui32(int32(64)) {
			v24 = F_SearchSysCache2(m, int32(24), v9, v16)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = F_format_type_be(m, v9)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v16
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v68
								F_errmsg(m, int32(762057), v7+int32(16))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(523322), int32(206), int32(38013))
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
					}
				} else {
					F_check_safe_enum_use(m, v24)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31)))
						F_ReleaseCatCache(m, v24)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v16)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v33
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = F_format_type_be(m, v9)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v16
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v49
						F_errmsg(m, int32(762057), v7)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523322), int32(196), int32(38013))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
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
