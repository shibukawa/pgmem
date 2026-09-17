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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 float32
	_ = v166
	var v167 float32
	_ = v167
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l0 == l1 {
		v208 = int32(0)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L11
	} else {
		goto L64
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v208
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
		v208 = v21
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
	v32 = F_SearchSysCache1(m, int32(23), l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v50 = v28
	goto L10
L10:
	;
	v51 = m.G0
	v53 = v51 - int32(32)
	m.G0 = v53
	if l0 == l1 {
		v176 = int32(0)
		goto L18
	} else {
		goto L19
	}
L11:
	;
	return int32(0)
L12:
	;
	if v32 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38+v39)+4))
	F_ReleaseCatCache(m, v32)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v45 = F_lookup_type_cache(m, v41, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v45
	v50 = v45
	goto L10
L16:
	;
	v208 = v176
	goto L2
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L60
	}
L18:
	;
	m.G0 = v53 + int32(32)
	goto L16
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+316))
	if v57 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_load_enum_cache_data(m, v50)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	v63 = v57
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if base.Ui32(l0) < base.Ui32(v64) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)+316))
	v63 = v62
	goto L22
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v89 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	v66 = l0 - v64
	if v66 < int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v70 = F_bms_is_member(m, v66, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	if v70 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if base.Ui32(l1) < base.Ui32(v74) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v76 = l1 - v74
	if v76 < int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v80 = F_bms_is_member(m, v76, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	if v80 == int32(0) {
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
	v87 = int32(-1)
	goto L35
L34:
	;
	v87 = int32(1)
	goto L35
L35:
	;
	v176 = v87
	goto L18
L36:
	;
	v166 = *(*float32)(unsafe.Add(mBase, uint32(v162)+4))
	v167 = *(*float32)(unsafe.Add(mBase, uint32(v161)+4))
	if base.F32_lt(v166, v167) != 0 {
		v176 = int32(-1)
		goto L18
	} else {
		goto L59
	}
L37:
	;
	F_load_enum_cache_data(m, v50)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L11
	} else {
		goto L44
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = l0
	v94 = v53 + int32(24)
	v96 = v63 + int32(12)
	v99 = F_bsearch(m, v94, v96, v89, int32(8), int32(1608))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v101 <= int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = l1
	v107 = F_bsearch(m, v94, v96, v101, int32(8), int32(1608))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	if v99 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if v107 != 0 {
		v161 = v107
		v162 = v99
		goto L36
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v50)+316))
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
	v147 = m.ExcPending
	if v147 != 0 {
		goto L11
	} else {
		goto L55
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = l0
	v123 = v53 + int32(24)
	v125 = v117 + int32(12)
	v128 = F_bsearch(m, v123, v125, v118, int32(8), int32(1608))
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
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = l1
	v136 = F_bsearch(m, v123, v125, v130, int32(8), int32(1608))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
	if v136 != 0 {
		v161 = v136
		v162 = v128
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
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v149 = F_format_type_be(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = l0
	F_errmsg_internal(m, int32(_a_F_enum_cmp_internal_3), v53)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_enum_cmp_internal_4), int32(2719), int32(_a_F_enum_cmp_internal_5))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	v176 = base.F32_gt(v166, v167)
	goto L18
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v186 = F_format_type_be(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_enum_cmp_internal_3), v53+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_enum_cmp_internal_4), int32(2722), int32(_a_F_enum_cmp_internal_5))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
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
	v219 = m.ExcPending
	if v219 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(_a_F_enum_cmp_internal_0), v14)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_enum_cmp_internal_1), int32(292), int32(_a_F_enum_cmp_internal_2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
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
								F_errmsg(m, int32(_a_F_enum_recv_0), v7+int32(16))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_enum_recv_1), int32(206), int32(_a_F_enum_recv_2))
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
						F_errmsg(m, int32(_a_F_enum_recv_0), v7)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_enum_recv_1), int32(196), int32(_a_F_enum_recv_2))
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
