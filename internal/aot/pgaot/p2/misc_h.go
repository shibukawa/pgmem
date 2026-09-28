package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_HoldPinnedPortals(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = v5 + int32(12)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_HoldPinnedPortals[0]))
	F_hash_seq_init(m, v8, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = F_hash_seq_search(m, v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L23
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v15 = v13
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v5 + int32(32)
	return
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+84)))
	if v18 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v32 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L17
	}
L12:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+85)))
	if v21 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v22 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v23 != int32(2) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	F_HoldPortal(m, v17)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+85)) = uint8(v28)
	goto L11
L17:
	;
	if v32 != 0 {
		v15 = v32
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_HoldPinnedPortals_0), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_HoldPinnedPortals_1), int32(1234), int32(_a_F_HoldPinnedPortals_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	F_errmsg_internal(m, int32(_a_F_HoldPinnedPortals_3), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_HoldPinnedPortals_1), int32(1238), int32(_a_F_HoldPinnedPortals_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_has_admin_privs_of_role(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14314(m, l0, l1, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_legal_joinclause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	if v13 == v3 {
		v113 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v113
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 <= int32(0) {
		v113 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v3
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v34 = int32(0)
	if base.B2i32(v27 == v34)|base.B2i32(v33 == v34) != 0 {
		v79 = v34
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v113 = v3
	goto L1
L6:
	;
	v103 = v24 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v103 < v104 {
		v24 = v103
		goto L4
	} else {
		goto L28
	}
L7:
	;
	if v79 != 0 {
		goto L6
	} else {
		goto L20
	}
L8:
	;
	goto L7
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v44 < v45 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = v44
	goto L12
L11:
	;
	v47 = v45
	goto L12
L12:
	;
	if v47 <= int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = int32(1)
	goto L15
L14:
	;
	v50 = v47
	goto L15
L15:
	;
	v51 = int32(8)
	v56 = int32(0)
	goto L16
L16:
	;
	v63 = v56 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33+v51+v63)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v27+v51+v63)))
	v68 = v65 & v67
	v70 = base.B2i32(v68 != int32(0))
	if v68 != 0 {
		v79 = v70
		goto L8
	} else {
		goto L18
	}
L17:
	;
	v79 = v70
	goto L8
L18:
	;
	v72 = v56 + int32(1)
	if v72 != v50 {
		v56 = v72
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v80 = F_have_relevant_joinclause(m, l0, l1, v32)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	if v80 == int32(0) {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v88 = F_bms_union(m, v86, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v94 = F_join_is_legal(m, l0, l1, v32, v88, v11+int32(12), v11+int32(11))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	F_bms_free(m, v88)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	if v94 == int32(0) {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v113 = int32(1)
	goto L1
L28:
	;
	goto L5
}
func F_has_parameter_privilege_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v13 = F_convert_any_priv_string(m, v10, int32(_a_F_has_parameter_privilege_name_0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_has_parameter_privilege_name[0]))
				v17 = F_text_to_cstring(m, v5)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = F_pg_parameter_aclcheck(m, v17, v16, v13)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(v19 == int32(0)))
					}
				}
			}
		}
	}
}
func F_has_privs_of_role(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14308(m, l0, l1, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_rolreplication(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v4 = F_superuser_arg(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v12 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(l0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				if v12 == int32(0) {
					return int32(0)
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v19)+73)))
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = v21
						return v25 & int32(1)
					}
				}
			}
		} else {
			v25 = int32(1)
			return v25 & int32(1)
		}
	}
}
func F_hashagg_reset_spill_state(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if int32(0) < v5 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	F_list_free_deep(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L14
	}
L4:
	;
	v10 = int32(0)
	goto L7
L5:
	;
	v30 = v4
	goto L6
L6:
	;
	F_pfree(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L13
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v15 = v12 + v10*int32(24)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_pfree(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v30 = v26
	goto L6
L9:
	;
	return
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	F_pfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v23 = v10 + int32(1)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v23 < v24 {
		v10 = v23
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	goto L3
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_LogicalTapeSetClose(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	return
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = int32(0)
	goto L17
}
func F_hashagg_spill_finish(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 float64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v103 int32
	_ = v103
	var v104 float64
	_ = v104
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v214 float64
	_ = v214
	var v216 float64
	_ = v216
	var v218 float64
	_ = v218
	var v231 float64
	_ = v231
	var v241 float64
	_ = v241
	var v252 float64
	_ = v252
	var v264 float64
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < v11 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v20 = v11
	v21 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_pfree(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L41
	} else {
		goto L47
	}
L7:
	;
	v28 = v21 << (uint(int32(3)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v28+v29)))
	if v31 != int64(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v21<<(uint(int32(2))%32))))
	v40 = v21 * int32(24)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v42 = v40 + v41
	v43 = int32(0)
	v50 = float64(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v293 = v20
	goto L11
L11:
	;
	v298 = v21 + int32(1)
	if v298 < v293 {
		v20 = v293
		v21 = v298
		goto L7
	} else {
		goto L46
	}
L12:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v265+v40)+16))
	F_pfree(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L41
	} else {
		goto L42
	}
L13:
	;
	v264 = v252
	goto L12
L14:
	;
	if base.F64_gt(v231, float64(1.4316557653333333e+08)) == int32(0) {
		v252 = v231
		goto L13
	} else {
		goto L40
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v52 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v42)+8))
	v218 = base.F64_div(v216, float64(0))
	if base.F64_le(v218, base.F64_mul(base.F64_convert_i32_u(v52), float64(2.5))) != 0 {
		v252 = v218
		goto L13
	} else {
		goto L39
	}
L18:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v42)+8))
	v118 = base.F64_div(v117, v115)
	v119 = base.F64_convert_i32_u(v52)
	if base.F64_le(v118, base.F64_mul(v119, float64(2.5))) == int32(0) {
		v231 = v118
		goto L14
	} else {
		goto L26
	}
L19:
	;
	v62 = v43
	v63 = v43
	v68 = v50
	goto L22
L20:
	;
	v92 = v43
	v98 = v50
	goto L21
L21:
	;
	v100 = float64(1)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v53))))
	v104 = F_scalbn(m, v100, v103)
	mBase = m.M
	v115 = base.F64_add(v98, base.F64_div(v100, v104))
	goto L18
L22:
	;
	v70 = float64(1)
	v71 = v62 + v53
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v73 = F_scalbn(m, v70, v72)
	mBase = m.M
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v77 = F_scalbn(m, v70, v76)
	mBase = m.M
	v82 = base.F64_add(base.F64_add(v68, base.F64_div(v70, v77)), base.F64_div(v70, v73))
	v83 = int32(2)
	v84 = v62 + v83
	v86 = v63 + v83
	if v86 != v52&int32(-2) {
		v62 = v84
		v63 = v86
		v68 = v82
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v52&int32(1) == int32(0) {
		v115 = v82
		goto L18
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v92 = v84
	v98 = v82
	goto L21
L26:
	;
	v126 = v52 & int32(3)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v128 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v52) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v203 == int32(0) {
		v252 = v118
		goto L13
	} else {
		goto L38
	}
L28:
	;
	v137 = int32(0)
	v138 = v128
	v139 = v128
	goto L31
L29:
	;
	v172 = v128
	v173 = v128
	goto L30
L30:
	;
	v182 = v172
	v183 = v173
	v186 = v128
	goto L35
L31:
	;
	v146 = v138 + v127
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	v148 = int32(0)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+2)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+3)))
	v162 = v139 + base.B2i32(v147 == v148) + base.B2i32(v151 == v148) + base.B2i32(v155 == v148) + base.B2i32(v159 == v148)
	v163 = int32(4)
	v164 = v138 + v163
	v166 = v137 + v163
	if v166 != v52&int32(-4) {
		v137 = v166
		v138 = v164
		v139 = v162
		goto L31
	} else {
		goto L33
	}
L32:
	;
	if v126 == int32(0) {
		v203 = v162
		goto L27
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v172 = v164
	v173 = v162
	goto L30
L35:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v127))))
	v194 = v183 + base.B2i32(v191 == int32(0))
	v195 = int32(1)
	v198 = v186 + v195
	if v198 != v126 {
		v182 = v182 + v195
		v183 = v194
		v186 = v198
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v203 = v194
	goto L27
L37:
	;
	goto L36
L38:
	;
	v214 = F_log(m, base.F64_div(v119, base.F64_convert_i32_s(v203)))
	mBase = m.M
	v264 = base.F64_mul(v214, v119)
	goto L12
L39:
	;
	v231 = v218
	goto L14
L40:
	;
	v241 = F_log(m, base.F64_add(base.F64_mul(v231, float64(-2.3283064365386963e-10)), float64(1)))
	mBase = m.M
	v252 = base.F64_mul(v241, float64(-4.294967296e+09))
	goto L13
L41:
	;
	return
L42:
	;
	F_LogicalTapeRewindForRead(m, v38, int32(_a_F_hashagg_spill_finish_0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v273+v28)))
	v277 = F_palloc0(m, int32(32))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v277)+24)) = v264
	*(*int64)(unsafe.Add(mBase, uint32(v277)+16)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = int32(32) - v15
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = l2
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v285 = F_lappend(m, v284, v277)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v285
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+336)) = v288 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v293 = v292
	goto L11
L46:
	;
	goto L8
L47:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_pfree(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_pfree(m, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L41
	} else {
		goto L49
	}
L49:
	;
	goto L3
}
func F_hashbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	v4 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(3316))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v11
			v13 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(-4294967296)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = v13
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v11)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v9
			return v4
		}
	}
}
func F_hashbpchar(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
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
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L127
	}
L4:
	;
	v15 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v19 = v17 & v15
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L122
	}
L7:
	;
	v20 = v15
	goto L9
L8:
	;
	v20 = int32(4)
	goto L9
L9:
	;
	v21 = v20 + v10
	if v17 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v54 = v48
	goto L21
L11:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v27 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v38 = int32(1)
	if v19 != 0 {
		v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v30 = int32(16)
	goto L16
L15:
	;
	v30 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v37 = int32(4)
	goto L19
L18:
	;
	v37 = v30
	goto L19
L19:
	;
	v48 = v37
	goto L10
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	if v54 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v70 = F_pg_newlocale_from_collation(m, v14)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v68 = v48 & (v48 >> (uint(int32(31)) % 32))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v63 = v54 - int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v63))))
	if v65 == int32(32) {
		v54 = v63
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v68 = v54
	goto L23
L28:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v623 != v10 {
		goto L118
	} else {
		goto L119
	}
L29:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v72 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v80 = v68 - int32(1636608432)
	if v21&int32(3) != 0 {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	goto L32
L32:
	;
	v339 = int32(0)
	v341 = F_pg_strnxfrm(m, v339, v339, v21, v68, v70)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L73
	}
L33:
	;
	v619 = v334 ^ v326 - base.I32_rotl(v334, int32(24))
	goto L28
L34:
	;
	v312 = int32(14)
	v314 = v308 ^ v309 - base.I32_rotl(v308, v312)
	v318 = v314 ^ v307 - base.I32_rotl(v314, int32(11))
	v322 = v318 ^ v308 - base.I32_rotl(v318, int32(25))
	v326 = v322 ^ v314 - base.I32_rotl(v322, int32(16))
	v330 = v326 ^ v318 - base.I32_rotl(v326, int32(4))
	v334 = v330 ^ v322 - base.I32_rotl(v330, v312)
	goto L33
L35:
	;
	switch v238 - int32(1) {
	case 0:
		v300 = v239
		v301 = v240
		v302 = v241
		goto L62
	case 1:
		v293 = v239
		v294 = v240
		v295 = v241
		goto L63
	case 2:
		v286 = v239
		v287 = v240
		v288 = v241
		goto L64
	case 3:
		v280 = v240
		v281 = v241
		goto L65
	case 4:
		v276 = v240
		v277 = v241
		goto L66
	case 5:
		v270 = v240
		v271 = v241
		goto L67
	case 6:
		v264 = v240
		v265 = v241
		goto L68
	case 7:
		v259 = v241
		goto L69
	case 8:
		v254 = v241
		goto L70
	case 9:
		v249 = v241
		goto L71
	case 10:
		goto L72
	default:
		v307 = v239
		v308 = v240
		v309 = v241
		goto L34
	}
L36:
	;
	v189 = v21
	v190 = v68
	v191 = v80
	v192 = v80
	v193 = v80
	goto L59
L37:
	;
	if base.Ui32(int32(11)) < base.Ui32(v68) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v68) < base.Ui32(int32(12)) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v237 = v21
	v238 = v68
	v239 = v80
	v240 = v80
	v241 = v80
	goto L35
L41:
	;
	switch v136 - int32(1) {
	case 0:
		v186 = v137
		goto L48
	case 1:
		v181 = v137
		goto L49
	case 2:
		goto L50
	case 3:
		v174 = v138
		goto L51
	case 4:
		v171 = v138
		goto L52
	case 5:
		v166 = v138
		goto L53
	case 6:
		goto L54
	case 7:
		v157 = v139
		goto L55
	case 8:
		v152 = v139
		goto L56
	case 9:
		v147 = v139
		goto L57
	case 10:
		goto L58
	default:
		v307 = v137
		v308 = v138
		v309 = v139
		goto L34
	}
L42:
	;
	v135 = v21
	v136 = v68
	v137 = v80
	v138 = v80
	v139 = v80
	goto L41
L43:
	;
	goto L44
L44:
	;
	v87 = v21
	v88 = v68
	v89 = v80
	v90 = v80
	v91 = v80
	goto L45
L45:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v94 = v93 + v90
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v98 = v97 + v91
	v100 = int32(4)
	v102 = v95 + v89 - v98 ^ base.I32_rotl(v98, v100)
	v106 = v94 - v102 ^ base.I32_rotl(v102, int32(6))
	v107 = v98 + v94
	v108 = v102 + v107
	v109 = v106 + v108
	v113 = v107 - v106 ^ base.I32_rotl(v106, int32(8))
	v117 = v108 - v113 ^ base.I32_rotl(v113, int32(16))
	v121 = v109 - v117 ^ base.I32_rotl(v117, int32(19))
	v122 = v113 + v109
	v123 = v117 + v122
	v124 = v121 + v123
	v128 = v122 - v121 ^ base.I32_rotl(v121, v100)
	v129 = int32(12)
	v130 = v87 + v129
	v132 = v88 - v129
	if base.Ui32(int32(11)) < base.Ui32(v132) {
		v87 = v130
		v88 = v132
		v89 = v123
		v90 = v124
		v91 = v128
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v135 = v130
	v136 = v132
	v137 = v123
	v138 = v124
	v139 = v128
	goto L41
L47:
	;
	goto L46
L48:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v307 = v186 + v187
	v308 = v138
	v309 = v139
	goto L34
L49:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	v186 = v182<<(uint(int32(8))%32) + v181
	goto L48
L50:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
	v181 = v177<<(uint(int32(16))%32) + v137
	goto L49
L51:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v307 = v175 + v137
	v308 = v174
	v309 = v139
	goto L34
L52:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+4)))
	v174 = v171 + v172
	goto L51
L53:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+5)))
	v171 = v167<<(uint(int32(8))%32) + v166
	goto L52
L54:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+6)))
	v166 = v162<<(uint(int32(16))%32) + v138
	goto L53
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v307 = v158 + v137
	v308 = v160 + v138
	v309 = v157
	goto L34
L56:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+8)))
	v157 = v153<<(uint(int32(8))%32) + v152
	goto L55
L57:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+9)))
	v152 = v148<<(uint(int32(16))%32) + v147
	goto L56
L58:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+10)))
	v147 = v143<<(uint(int32(24))%32) + v139
	goto L57
L59:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v196 = v195 + v192
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v200 = v199 + v193
	v202 = int32(4)
	v204 = v197 + v191 - v200 ^ base.I32_rotl(v200, v202)
	v208 = v196 - v204 ^ base.I32_rotl(v204, int32(6))
	v209 = v200 + v196
	v210 = v204 + v209
	v211 = v208 + v210
	v215 = v209 - v208 ^ base.I32_rotl(v208, int32(8))
	v219 = v210 - v215 ^ base.I32_rotl(v215, int32(16))
	v223 = v211 - v219 ^ base.I32_rotl(v219, int32(19))
	v224 = v215 + v211
	v225 = v219 + v224
	v226 = v223 + v225
	v230 = v224 - v223 ^ base.I32_rotl(v223, v202)
	v231 = int32(12)
	v232 = v189 + v231
	v234 = v190 - v231
	if base.Ui32(int32(11)) < base.Ui32(v234) {
		v189 = v232
		v190 = v234
		v191 = v225
		v192 = v226
		v193 = v230
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v237 = v232
	v238 = v234
	v239 = v225
	v240 = v226
	v241 = v230
	goto L35
L61:
	;
	goto L60
L62:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v307 = v300 + v303
	v308 = v301
	v309 = v302
	goto L34
L63:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)))
	v300 = v296<<(uint(int32(8))%32) + v293
	v301 = v294
	v302 = v295
	goto L62
L64:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+2)))
	v293 = v289<<(uint(int32(16))%32) + v286
	v294 = v287
	v295 = v288
	goto L63
L65:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+3)))
	v286 = v282<<(uint(int32(24))%32) + v239
	v287 = v280
	v288 = v281
	goto L64
L66:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+4)))
	v280 = v276 + v278
	v281 = v277
	goto L65
L67:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+5)))
	v276 = v272<<(uint(int32(8))%32) + v270
	v277 = v271
	goto L66
L68:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+6)))
	v270 = v266<<(uint(int32(16))%32) + v264
	v271 = v265
	goto L67
L69:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+7)))
	v264 = v260<<(uint(int32(24))%32) + v240
	v265 = v259
	goto L68
L70:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+8)))
	v259 = v255<<(uint(int32(8))%32) + v254
	goto L69
L71:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+9)))
	v254 = v250<<(uint(int32(16))%32) + v249
	goto L70
L72:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+10)))
	v249 = v245<<(uint(int32(24))%32) + v241
	goto L71
L73:
	;
	v344 = v341 + int32(1)
	v345 = F_palloc(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v347 = F_pg_strnxfrm(m, v345, v344, v21, v68, v70)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if base.Ui32(v341) < base.Ui32(v347) {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	v351 = v347 + int32(1)
	v357 = v351 - int32(1636608432)
	if v345&int32(3) != 0 {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	F_pfree(m, v345)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L117
	}
L78:
	;
	v589 = int32(14)
	v591 = v585 ^ v586 - base.I32_rotl(v585, v589)
	v595 = v591 ^ v584 - base.I32_rotl(v591, int32(11))
	v599 = v595 ^ v585 - base.I32_rotl(v595, int32(25))
	v603 = v599 ^ v591 - base.I32_rotl(v599, int32(16))
	v607 = v603 ^ v595 - base.I32_rotl(v603, int32(4))
	v611 = v607 ^ v599 - base.I32_rotl(v607, v589)
	goto L77
L79:
	;
	switch v515 - int32(1) {
	case 0:
		v577 = v516
		v578 = v517
		v579 = v518
		goto L106
	case 1:
		v570 = v516
		v571 = v517
		v572 = v518
		goto L107
	case 2:
		v563 = v516
		v564 = v517
		v565 = v518
		goto L108
	case 3:
		v557 = v517
		v558 = v518
		goto L109
	case 4:
		v553 = v517
		v554 = v518
		goto L110
	case 5:
		v547 = v517
		v548 = v518
		goto L111
	case 6:
		v541 = v517
		v542 = v518
		goto L112
	case 7:
		v536 = v518
		goto L113
	case 8:
		v531 = v518
		goto L114
	case 9:
		v526 = v518
		goto L115
	case 10:
		goto L116
	default:
		v584 = v516
		v585 = v517
		v586 = v518
		goto L78
	}
L80:
	;
	v466 = v345
	v467 = v351
	v468 = v357
	v469 = v357
	v470 = v357
	goto L103
L81:
	;
	if base.Ui32(int32(11)) < base.Ui32(v351) {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(v351) < base.Ui32(int32(12)) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v514 = v345
	v515 = v351
	v516 = v357
	v517 = v357
	v518 = v357
	goto L79
L85:
	;
	switch v413 - int32(1) {
	case 0:
		v463 = v414
		goto L92
	case 1:
		v458 = v414
		goto L93
	case 2:
		goto L94
	case 3:
		v451 = v415
		goto L95
	case 4:
		v448 = v415
		goto L96
	case 5:
		v443 = v415
		goto L97
	case 6:
		goto L98
	case 7:
		v434 = v416
		goto L99
	case 8:
		v429 = v416
		goto L100
	case 9:
		v424 = v416
		goto L101
	case 10:
		goto L102
	default:
		v584 = v414
		v585 = v415
		v586 = v416
		goto L78
	}
L86:
	;
	v412 = v345
	v413 = v351
	v414 = v357
	v415 = v357
	v416 = v357
	goto L85
L87:
	;
	goto L88
L88:
	;
	v364 = v345
	v365 = v351
	v366 = v357
	v367 = v357
	v368 = v357
	goto L89
L89:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	v371 = v370 + v367
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v364)+8))
	v375 = v374 + v368
	v377 = int32(4)
	v379 = v372 + v366 - v375 ^ base.I32_rotl(v375, v377)
	v383 = v371 - v379 ^ base.I32_rotl(v379, int32(6))
	v384 = v375 + v371
	v385 = v379 + v384
	v386 = v383 + v385
	v390 = v384 - v383 ^ base.I32_rotl(v383, int32(8))
	v394 = v385 - v390 ^ base.I32_rotl(v390, int32(16))
	v398 = v386 - v394 ^ base.I32_rotl(v394, int32(19))
	v399 = v390 + v386
	v400 = v394 + v399
	v401 = v398 + v400
	v405 = v399 - v398 ^ base.I32_rotl(v398, v377)
	v406 = int32(12)
	v407 = v364 + v406
	v409 = v365 - v406
	if base.Ui32(int32(11)) < base.Ui32(v409) {
		v364 = v407
		v365 = v409
		v366 = v400
		v367 = v401
		v368 = v405
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v412 = v407
	v413 = v409
	v414 = v400
	v415 = v401
	v416 = v405
	goto L85
L91:
	;
	goto L90
L92:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	v584 = v463 + v464
	v585 = v415
	v586 = v416
	goto L78
L93:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+1)))
	v463 = v459<<(uint(int32(8))%32) + v458
	goto L92
L94:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+2)))
	v458 = v454<<(uint(int32(16))%32) + v414
	goto L93
L95:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v584 = v452 + v414
	v585 = v451
	v586 = v416
	goto L78
L96:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+4)))
	v451 = v448 + v449
	goto L95
L97:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+5)))
	v448 = v444<<(uint(int32(8))%32) + v443
	goto L96
L98:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+6)))
	v443 = v439<<(uint(int32(16))%32) + v415
	goto L97
L99:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	v584 = v435 + v414
	v585 = v437 + v415
	v586 = v434
	goto L78
L100:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+8)))
	v434 = v430<<(uint(int32(8))%32) + v429
	goto L99
L101:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+9)))
	v429 = v425<<(uint(int32(16))%32) + v424
	goto L100
L102:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+10)))
	v424 = v420<<(uint(int32(24))%32) + v416
	goto L101
L103:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	v473 = v472 + v469
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v466)+8))
	v477 = v476 + v470
	v479 = int32(4)
	v481 = v474 + v468 - v477 ^ base.I32_rotl(v477, v479)
	v485 = v473 - v481 ^ base.I32_rotl(v481, int32(6))
	v486 = v477 + v473
	v487 = v481 + v486
	v488 = v485 + v487
	v492 = v486 - v485 ^ base.I32_rotl(v485, int32(8))
	v496 = v487 - v492 ^ base.I32_rotl(v492, int32(16))
	v500 = v488 - v496 ^ base.I32_rotl(v496, int32(19))
	v501 = v492 + v488
	v502 = v496 + v501
	v503 = v500 + v502
	v507 = v501 - v500 ^ base.I32_rotl(v500, v479)
	v508 = int32(12)
	v509 = v466 + v508
	v511 = v467 - v508
	if base.Ui32(int32(11)) < base.Ui32(v511) {
		v466 = v509
		v467 = v511
		v468 = v502
		v469 = v503
		v470 = v507
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v514 = v509
	v515 = v511
	v516 = v502
	v517 = v503
	v518 = v507
	goto L79
L105:
	;
	goto L104
L106:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	v584 = v577 + v580
	v585 = v578
	v586 = v579
	goto L78
L107:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
	v577 = v573<<(uint(int32(8))%32) + v570
	v578 = v571
	v579 = v572
	goto L106
L108:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+2)))
	v570 = v566<<(uint(int32(16))%32) + v563
	v571 = v564
	v572 = v565
	goto L107
L109:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+3)))
	v563 = v559<<(uint(int32(24))%32) + v516
	v564 = v557
	v565 = v558
	goto L108
L110:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+4)))
	v557 = v553 + v555
	v558 = v554
	goto L109
L111:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+5)))
	v553 = v549<<(uint(int32(8))%32) + v547
	v554 = v548
	goto L110
L112:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+6)))
	v547 = v543<<(uint(int32(16))%32) + v541
	v548 = v542
	goto L111
L113:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+7)))
	v541 = v537<<(uint(int32(24))%32) + v517
	v542 = v536
	goto L112
L114:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+8)))
	v536 = v532<<(uint(int32(8))%32) + v531
	goto L113
L115:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+9)))
	v531 = v527<<(uint(int32(16))%32) + v526
	goto L114
L116:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+10)))
	v526 = v522<<(uint(int32(24))%32) + v518
	goto L115
L117:
	;
	v619 = v611 ^ v603 - base.I32_rotl(v611, int32(24))
	goto L28
L118:
	;
	F_pfree(m, v10)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	return base.I64_extend_i32_u(v619)
L121:
	;
	goto L120
L122:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(_a_F_hashbpchar_0), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errhint(m, int32(_a_F_hashbpchar_1), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_hashbpchar_2), int32(1004), int32(_a_F_hashbpchar_3))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errmsg_internal(m, int32(_a_F_hashbpchar_4), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_hashbpchar_2), int32(1028), int32(_a_F_hashbpchar_3))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashchar(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
	v7 = int32(711645284)
	v10 = v2 - int32(1636608428) ^ v7 - int32(1455628627)
	v15 = v10 ^ int32(-1636608428) - base.I32_rotl(v10, int32(25))
	v20 = v15 ^ v7 - base.I32_rotl(v15, int32(16))
	v24 = v20 ^ v10 - base.I32_rotl(v20, int32(4))
	v28 = v24 ^ v15 - base.I32_rotl(v24, int32(14))
	return base.I64_extend_i32_u(v28 ^ v20 - base.I32_rotl(v28, int32(24)))
}
func F_hashfloat4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 float32
	_ = v10
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v290 int64
	_ = v290
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F32_ne(v10, float32(0)) != 0 {
		v14 = base.F64_promote_f32(v10)
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) {
			v20 = math.Float64frombits(uint64(0x7ff8000000000000))
		} else {
			v20 = v14
		}
		*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v20
		v23 = v8 + int32(8)
		v30 = int32(-1636608424)
		if v23&int32(3) != 0 {
			switch int32(7) {
			case 0:
				v250 = v30
				v251 = v30
				v252 = v30
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 1:
				v243 = v30
				v244 = v30
				v245 = v30
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 2:
				v236 = v30
				v237 = v30
				v238 = v30
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 3:
				v230 = v30
				v231 = v30
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
				v236 = v232<<(uint(int32(24))%32) + v30
				v237 = v230
				v238 = v231
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 4:
				v226 = v30
				v227 = v30
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v230 = v226 + v228
				v231 = v227
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
				v236 = v232<<(uint(int32(24))%32) + v30
				v237 = v230
				v238 = v231
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 5:
				v220 = v30
				v221 = v30
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)))
				v226 = v222<<(uint(int32(8))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v230 = v226 + v228
				v231 = v227
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
				v236 = v232<<(uint(int32(24))%32) + v30
				v237 = v230
				v238 = v231
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 6:
				v214 = v30
				v215 = v30
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+6)))
				v220 = v216<<(uint(int32(16))%32) + v214
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)))
				v226 = v222<<(uint(int32(8))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v230 = v226 + v228
				v231 = v227
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
				v236 = v232<<(uint(int32(24))%32) + v30
				v237 = v230
				v238 = v231
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 7:
				v209 = v30
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)))
				v214 = v210<<(uint(int32(24))%32) + v30
				v215 = v209
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+6)))
				v220 = v216<<(uint(int32(16))%32) + v214
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)))
				v226 = v222<<(uint(int32(8))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v230 = v226 + v228
				v231 = v227
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
				v236 = v232<<(uint(int32(24))%32) + v30
				v237 = v230
				v238 = v231
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 8:
				v204 = v30
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
				v209 = v205<<(uint(int32(8))%32) + v204
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)))
				v214 = v210<<(uint(int32(24))%32) + v30
				v215 = v209
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+6)))
				v220 = v216<<(uint(int32(16))%32) + v214
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)))
				v226 = v222<<(uint(int32(8))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v230 = v226 + v228
				v231 = v227
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
				v236 = v232<<(uint(int32(24))%32) + v30
				v237 = v230
				v238 = v231
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 9:
				v199 = v30
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+9)))
				v204 = v200<<(uint(int32(16))%32) + v199
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
				v209 = v205<<(uint(int32(8))%32) + v204
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)))
				v214 = v210<<(uint(int32(24))%32) + v30
				v215 = v209
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+6)))
				v220 = v216<<(uint(int32(16))%32) + v214
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)))
				v226 = v222<<(uint(int32(8))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v230 = v226 + v228
				v231 = v227
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
				v236 = v232<<(uint(int32(24))%32) + v30
				v237 = v230
				v238 = v231
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			case 10:
				v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+10)))
				v199 = v195<<(uint(int32(24))%32) + v30
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+9)))
				v204 = v200<<(uint(int32(16))%32) + v199
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
				v209 = v205<<(uint(int32(8))%32) + v204
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)))
				v214 = v210<<(uint(int32(24))%32) + v30
				v215 = v209
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+6)))
				v220 = v216<<(uint(int32(16))%32) + v214
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)))
				v226 = v222<<(uint(int32(8))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v230 = v226 + v228
				v231 = v227
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
				v236 = v232<<(uint(int32(24))%32) + v30
				v237 = v230
				v238 = v231
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v243 = v239<<(uint(int32(16))%32) + v236
				v244 = v237
				v245 = v238
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v250 = v246<<(uint(int32(8))%32) + v243
				v251 = v244
				v252 = v245
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v250 + v253
				v258 = v251
				v259 = v252
			default:
				v257 = v30
				v258 = v30
				v259 = v30
			}
		} else {
			switch int32(7) {
			case 0:
				v136 = v30
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v136 + v137
				v258 = v30
				v259 = v30
			case 1:
				v131 = v30
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v136 = v132<<(uint(int32(8))%32) + v131
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v136 + v137
				v258 = v30
				v259 = v30
			case 2:
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
				v131 = v127<<(uint(int32(16))%32) + v30
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v136 = v132<<(uint(int32(8))%32) + v131
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v257 = v136 + v137
				v258 = v30
				v259 = v30
			case 3:
				v124 = v30
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v257 = v125 + v30
				v258 = v124
				v259 = v30
			case 4:
				v121 = v30
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v124 = v121 + v122
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v257 = v125 + v30
				v258 = v124
				v259 = v30
			case 5:
				v116 = v30
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)))
				v121 = v117<<(uint(int32(8))%32) + v116
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v124 = v121 + v122
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v257 = v125 + v30
				v258 = v124
				v259 = v30
			case 6:
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+6)))
				v116 = v112<<(uint(int32(16))%32) + v30
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)))
				v121 = v117<<(uint(int32(8))%32) + v116
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
				v124 = v121 + v122
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v257 = v125 + v30
				v258 = v124
				v259 = v30
			case 7:
				v107 = v30
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v257 = v108 + v30
				v258 = v110 + v30
				v259 = v107
			case 8:
				v102 = v30
				v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
				v107 = v103<<(uint(int32(8))%32) + v102
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v257 = v108 + v30
				v258 = v110 + v30
				v259 = v107
			case 9:
				v97 = v30
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+9)))
				v102 = v98<<(uint(int32(16))%32) + v97
				v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
				v107 = v103<<(uint(int32(8))%32) + v102
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v257 = v108 + v30
				v258 = v110 + v30
				v259 = v107
			case 10:
				v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+10)))
				v97 = v93<<(uint(int32(24))%32) + v30
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+9)))
				v102 = v98<<(uint(int32(16))%32) + v97
				v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
				v107 = v103<<(uint(int32(8))%32) + v102
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v257 = v108 + v30
				v258 = v110 + v30
				v259 = v107
			default:
				v257 = v30
				v258 = v30
				v259 = v30
			}
		}
		v262 = int32(14)
		v264 = v258 ^ v259 - base.I32_rotl(v258, v262)
		v268 = v264 ^ v257 - base.I32_rotl(v264, int32(11))
		v272 = v268 ^ v258 - base.I32_rotl(v268, int32(25))
		v276 = v272 ^ v264 - base.I32_rotl(v272, int32(16))
		v280 = v276 ^ v268 - base.I32_rotl(v276, int32(4))
		v284 = v280 ^ v272 - base.I32_rotl(v280, v262)
		v290 = base.I64_extend_i32_u(v284 ^ v276 - base.I32_rotl(v284, int32(24)))
	} else {
		v290 = int64(0)
	}
	m.G0 = v8 + int32(16)
	return v290
}
func F_hashint4extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v3 == int64(0) {
		v10 = int32(-1636608428)
		v49 = v10
		v50 = v10
		v53 = int32(0)
	} else {
		v13 = base.I32_wrap_i64(v3)
		v15 = v13 + int32(1021750440)
		v20 = base.I32_wrap_i64(int64(base.Ui64(v3)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v26 = v13 - v20 - int32(1636608428) ^ base.I32_rotl(v20, int32(6))
		v30 = v15 - v26 ^ base.I32_rotl(v26, int32(8))
		v31 = v20 + v15
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v37 + v33
		v47 = v41 + v46
		v49 = v47
		v50 = v46
		v53 = v33 - v41 ^ base.I32_rotl(v41, int32(4)) ^ v47
	}
	v54 = int32(14)
	v56 = v53 - base.I32_rotl(v49, v54)
	v61 = v56 ^ (v2 + v50) - base.I32_rotl(v56, int32(11))
	v65 = v49 ^ v61 - base.I32_rotl(v61, int32(25))
	v69 = v65 ^ v56 - base.I32_rotl(v65, int32(16))
	v73 = v69 ^ v61 - base.I32_rotl(v69, int32(4))
	v77 = v73 ^ v65 - base.I32_rotl(v73, v54)
	return base.I64_extend_i32_u(v77)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v77^v69-base.I32_rotl(v77, int32(24)))
}
func F_hemdistcache_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int64
	_ = v149
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v231 int64
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int64
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v263 int64
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int64
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int64
	_ = v292
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int64
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int64
	_ = v364
	var v374 int64
	_ = v374
	var v383 int64
	_ = v383
	var v395 int64
	_ = v395
	v9 = int64(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(1) {
		if v10&int32(1) != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if int32(7) < l2 {
				v231 = int64(0)
				v232 = int32(0)
				if l2 == v232 {
					v302 = int64(0)
				} else {
					v239 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v244 = v20
						v246 = v231
						v249 = v232
						for {
							v250 = int32(4)
							v251 = v244 + v250
							v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+3)))
							v253 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v252)+uint32(_c_F_hemdistcache_1[0]))))
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+2)))
							v255 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v254)+uint32(_c_F_hemdistcache_1[0]))))
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
							v257 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v256)+uint32(_c_F_hemdistcache_1[0]))))
							v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
							v259 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v258)+uint32(_c_F_hemdistcache_1[0]))))
							v263 = v253 + (v255 + (v257 + (v246 + v259)))
							v265 = v249 + v250
							if v265 != l2&int32(-4) {
								v244 = v251
								v246 = v263
								v249 = v265
								continue
							} else {
								break
							}
							break
						}
						if v239 == int32(0) {
							v292 = v263
						} else {
							v269 = v251
							v271 = v263
							v276 = v269
							v277 = int32(0)
							v278 = v271
							for {
								v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
								v283 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_hemdistcache_1[0]))))
								v284 = v278 + v283
								v285 = int32(1)
								v288 = v277 + v285
								if v288 != v239 {
									v276 = v276 + v285
									v277 = v288
									v278 = v284
									continue
								} else {
									break
								}
								break
							}
							v292 = v284
						}
					} else {
						v269 = v20
						v271 = v231
						v276 = v269
						v277 = int32(0)
						v278 = v271
						for {
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
							v283 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_hemdistcache_1[0]))))
							v284 = v278 + v283
							v285 = int32(1)
							v288 = v277 + v285
							if v288 != v239 {
								v276 = v276 + v285
								v277 = v288
								v278 = v284
								continue
							} else {
								break
							}
							break
						}
						v292 = v284
					}
					v302 = v292
				}
				v395 = v302
			} else {
				if l2 == int32(0) {
					v395 = v9
				} else {
					v26 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v32 = v20
						v34 = int32(0)
						v40 = v9
						for {
							v41 = int32(4)
							v42 = v32 + v41
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+3)))
							v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_hemdistcache_1[0]))))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+2)))
							v46 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_hemdistcache_1[0]))))
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
							v48 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_hemdistcache_1[0]))))
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
							v50 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+uint32(_c_F_hemdistcache_1[0]))))
							v54 = v44 + (v46 + (v48 + (v40 + v50)))
							v56 = v34 + v41
							if v56 != l2&int32(-4) {
								v32 = v42
								v34 = v56
								v40 = v54
								continue
							} else {
								break
							}
							break
						}
						if v26 == int32(0) {
							v395 = v54
						} else {
							v60 = v42
							v68 = v54
							v70 = v60
							v71 = int32(0)
							v78 = v68
							for {
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
								v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_hemdistcache_1[0]))))
								v81 = v78 + v80
								v82 = int32(1)
								v85 = v71 + v82
								if v85 != v26 {
									v70 = v70 + v82
									v71 = v85
									v78 = v81
									continue
								} else {
									break
								}
								break
							}
							v395 = v81
						}
					} else {
						v60 = v20
						v68 = v9
						v70 = v60
						v71 = int32(0)
						v78 = v68
						for {
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
							v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_hemdistcache_1[0]))))
							v81 = v78 + v80
							v82 = int32(1)
							v85 = v71 + v82
							if v85 != v26 {
								v70 = v70 + v82
								v71 = v85
								v78 = v81
								continue
							} else {
								break
							}
							break
						}
						v395 = v81
					}
				}
			}
			return l2<<(uint(int32(3))%32) - base.I32_wrap_i64(v395)
		}
	} else {
		if v10&int32(1) != 0 {
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if int32(7) < l2 {
				v303 = int64(0)
				v304 = int32(0)
				if l2 == v304 {
					v374 = int64(0)
				} else {
					v311 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v316 = v91
						v318 = v303
						v321 = v304
						for {
							v322 = int32(4)
							v323 = v316 + v322
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+3)))
							v325 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v324)+uint32(_c_F_hemdistcache_1[0]))))
							v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+2)))
							v327 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v326)+uint32(_c_F_hemdistcache_1[0]))))
							v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+1)))
							v329 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v328)+uint32(_c_F_hemdistcache_1[0]))))
							v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
							v331 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v330)+uint32(_c_F_hemdistcache_1[0]))))
							v335 = v325 + (v327 + (v329 + (v318 + v331)))
							v337 = v321 + v322
							if v337 != l2&int32(-4) {
								v316 = v323
								v318 = v335
								v321 = v337
								continue
							} else {
								break
							}
							break
						}
						if v311 == int32(0) {
							v364 = v335
						} else {
							v341 = v323
							v343 = v335
							v348 = v341
							v349 = int32(0)
							v350 = v343
							for {
								v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
								v355 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v354)+uint32(_c_F_hemdistcache_1[0]))))
								v356 = v350 + v355
								v357 = int32(1)
								v360 = v349 + v357
								if v360 != v311 {
									v348 = v348 + v357
									v349 = v360
									v350 = v356
									continue
								} else {
									break
								}
								break
							}
							v364 = v356
						}
					} else {
						v341 = v91
						v343 = v303
						v348 = v341
						v349 = int32(0)
						v350 = v343
						for {
							v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
							v355 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v354)+uint32(_c_F_hemdistcache_1[0]))))
							v356 = v350 + v355
							v357 = int32(1)
							v360 = v349 + v357
							if v360 != v311 {
								v348 = v348 + v357
								v349 = v360
								v350 = v356
								continue
							} else {
								break
							}
							break
						}
						v364 = v356
					}
					v374 = v364
				}
				v383 = v374
			} else {
				if l2 == int32(0) {
					v383 = v9
				} else {
					v97 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v103 = v91
						v105 = int32(0)
						v111 = v9
						for {
							v112 = int32(4)
							v113 = v103 + v112
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+3)))
							v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_hemdistcache_1[0]))))
							v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+2)))
							v117 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_hemdistcache_1[0]))))
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
							v119 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v118)+uint32(_c_F_hemdistcache_1[0]))))
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
							v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_hemdistcache_1[0]))))
							v125 = v115 + (v117 + (v119 + (v111 + v121)))
							v127 = v105 + v112
							if v127 != l2&int32(-4) {
								v103 = v113
								v105 = v127
								v111 = v125
								continue
							} else {
								break
							}
							break
						}
						if v97 == int32(0) {
							v383 = v125
						} else {
							v131 = v113
							v139 = v125
							v141 = v131
							v142 = int32(0)
							v149 = v139
							for {
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
								v151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v150)+uint32(_c_F_hemdistcache_1[0]))))
								v152 = v149 + v151
								v153 = int32(1)
								v156 = v142 + v153
								if v156 != v97 {
									v141 = v141 + v153
									v142 = v156
									v149 = v152
									continue
								} else {
									break
								}
								break
							}
							v383 = v152
						}
					} else {
						v131 = v91
						v139 = v9
						v141 = v131
						v142 = int32(0)
						v149 = v139
						for {
							v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
							v151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v150)+uint32(_c_F_hemdistcache_1[0]))))
							v152 = v149 + v151
							v153 = int32(1)
							v156 = v142 + v153
							if v156 != v97 {
								v141 = v141 + v153
								v142 = v156
								v149 = v152
								continue
							} else {
								break
							}
							break
						}
						v383 = v152
					}
				}
			}
			return l2<<(uint(int32(3))%32) - base.I32_wrap_i64(v383)
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v164 = int32(0)
				if l2 != int32(1) {
					v173 = v164
					v174 = v164
					v175 = int32(0)
					for {
						v183 = v173 | int32(1)
						v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v183))))
						v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v183))))
						v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185^v187)+uint32(_c_F_hemdistcache_1[0]))))
						v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v162))))
						v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v163))))
						v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191^v193)+uint32(_c_F_hemdistcache_1[0]))))
						v197 = v189 + (v174 + v195)
						v198 = int32(2)
						v199 = v173 + v198
						v201 = v175 + v198
						if v201 != l2&int32(2147483646) {
							v173 = v199
							v174 = v197
							v175 = v201
							continue
						} else {
							break
						}
						break
					}
					if l2&int32(1) == int32(0) {
						v222 = v197
					} else {
						v205 = v199
						v206 = v197
						v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v162))))
						v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v163))))
						v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215^v217)+uint32(_c_F_hemdistcache_1[0]))))
						v222 = v206 + v219
					}
				} else {
					v205 = v164
					v206 = v164
					v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v162))))
					v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v163))))
					v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215^v217)+uint32(_c_F_hemdistcache_1[0]))))
					v222 = v206 + v219
				}
				return v222
			}
		}
	}
}
func F_hex_dec_len(m *base.Module, l0 int32, l1 int32) int64 {
	return base.I64_extend_i32_u(int32(base.Ui32(l1) >> (uint(int32(1)) % 32)))
}
func F_hk_depth_search(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v96 int32
	_ = v96
	v2 = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v2<<(uint(int32(2))%32))))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17))))
	v20 = v18
	goto L3
L2:
	;
	v20 = int32(0)
	goto L3
L3:
	;
	if v2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(1)
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = v25 + v2<<(uint(int32(1))%32)
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	if v29 != int32(_a_F_hk_depth_search_0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_check_stack_depth(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	return int32(0)
L10:
	;
	return int32(0)
L11:
	;
	if int32(0) < v20 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = v20
	goto L15
L13:
	;
	goto L14
L14:
	;
	v96 = int32(_a_F_hk_depth_search_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v28))) = uint16(v96)
	goto L9
L15:
	;
	v56 = int32(1)
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17+v46<<(uint(v56)%32)))))
	v62 = v32 + v59<<(uint(v56)%32)
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62))))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v63<<(uint(v56)%32)))))
	if v67 != (v29+int32(1))&int32(_a_F_hk_depth_search_1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v80 = int32(1)
	if v80 < v46 {
		v46 = v46 - v80
		goto L15
	} else {
		goto L21
	}
L18:
	;
	v69 = F_hk_depth_search(m, l0, v63)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	if v69 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v2)
	v74 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33+v2<<(uint(v74)%32)))) = uint16(v59)
	return v74
L21:
	;
	goto L16
}
func F_hnswbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(208)
	m.G0 = v7
	F_BuildIndex_1(m, l0, l1, l2, v7, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = F_palloc(m, int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v7)+40))
			*(*float64)(unsafe.Add(mBase, uint32(v15))) = v17
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v7)+32))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v19
			m.G0 = v7 + int32(208)
			return v15
		}
	}
}
func F_hnswgettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 float64
	_ = v424
	var v425 float64
	_ = v425
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(_a_F_hnswgettuple_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	*(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[0])) = v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	if v24 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L11
	} else {
		goto L101
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L11
	} else {
		goto L98
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+272))
	if v28 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	goto L44
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+268)))
	if v31 != int32(1) {
		v47 = v27
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v40 = v28
	v41 = v27
	goto L9
L9:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+16)) = v42 + int64(1)
	v47 = v41
	goto L6
L10:
	;
	F_pgstat_assoc_relation(m, v27)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+272))
	v40 = v39
	v41 = v38
	goto L9
L13:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = v49 + int64(1)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v53 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v58&int32(1) != 0 {
		v72 = v47
		v73 = int64(0)
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_LockPage(m, v72, int32(1), int32(5))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L22
	}
L19:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v53)+48))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+68))
	if v63 == int32(0) {
		v72 = v47
		v73 = v61
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	v68 = F_HnswNormValue(m, v66, v67, v61)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = v70
	v73 = v68
	goto L18
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HnswGetMetaPageInfo(m, v79, v16+int32(12), v16+int32(8))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v78)+24)) = v73
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v89 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v90 = int32(0)
	v92 = v78 + int32(24)
	v94 = v78 - int32(-64)
	v96 = F_HnswEntryCandidate(m, v90, v89, v92, v79, v94, v90)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L27
	}
L25:
	;
	v180 = int32(0)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_UnlockPage(m, v182, int32(1), int32(5))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L40
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v96
	v101 = F_list_make1_impl(m, int32(1), v16)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+65)))
	if v104 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v106 = v104
	v107 = v101
	goto L32
L30:
	;
	v135 = v101
	goto L31
L31:
	;
	v146 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[1]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[2]))
	if v159 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v118 = int32(1)
	v120 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v129 = F_HnswSearchLayer(m, v120, v92, v107, v118, v106, v79, v94, v122, v120, v120, v120, v120, v118, v120)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	v135 = v129
	goto L31
L34:
	;
	if base.Ui32(v118) < base.Ui32(v106) {
		v106 = v106 - int32(1)
		v107 = v129
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v160 = v78 + int32(16)
	goto L38
L37:
	;
	v160 = v146
	goto L38
L38:
	;
	v164 = F_HnswSearchLayer(m, v146, v92, v135, v148, v146, v79, v94, v150, v146, v146, v78+int32(12), v160, int32(1), v78+int32(40))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v180 = v164
	goto L26
L40:
	;
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v187)
	goto L5
L41:
	;
	m.G0 = v16 + int32(16)
	return v476
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[0])) = v19
	v476 = int32(0)
	goto L41
L43:
	;
	v435 = v402 + v418&int32(255)*int32(6) + int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[0])) = v19
	v438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v435)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v438)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v440
	v442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v442)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v442)
	v476 = int32(1)
	goto L41
L44:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v215 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+48)) = v424
	goto L43
L46:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v395+v384<<(uint(int32(2))%32)-int32(4))))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+24))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+64)))
	if v403 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L47:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v216 != 0 {
		v383 = v215
		v384 = v216
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[2]))
	if v219 == int32(0) {
		goto L42
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v222 == int32(0) {
		goto L42
	} else {
		goto L52
	}
L52:
	;
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	v227 = int64(*(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[3])))
	if v225 < v227 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	if v365 == int32(0) {
		goto L42
	} else {
		goto L87
	}
L54:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_LockPage(m, v272, int32(1), int32(5))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L11
	} else {
		goto L73
	}
L55:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)+8))
	goto L59
L56:
	;
	v260 = v222
	goto L57
L57:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	if v261 == int32(0) {
		goto L42
	} else {
		goto L70
	}
L58:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if base.Ui32(v233) <= base.Ui32(v257) {
		goto L54
	} else {
		goto L69
	}
L59:
	;
	goto L58
L69:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v260 = v259
	goto L57
L70:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v265 = F_pairingheap_remove_first(m, v260)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	v269 = F_lappend(m, v264, v265-int32(12))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v269
	v365 = v269
	goto L53
L73:
	;
	v277 = int32(0)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+16))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	if v281 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v283 = v279 + int32(16)
	v284 = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[1]))
	if v286 <= v284 {
		v316 = v277
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v346 = v277
	v351 = v278
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v346
	F_UnlockPage(m, v351, int32(1), int32(5))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L11
	} else {
		goto L86
	}
L77:
	;
	v328 = int32(0)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v279)+32))
	v342 = F_HnswSearchLayer(m, v328, v279+int32(24), v316, v286, v328, v278, v279-int32(-64), v334, v328, v328, v279+int32(12), v283, v328, v279+int32(40))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L11
	} else {
		goto L85
	}
L78:
	;
	v290 = v277
	v291 = v284
	goto L79
L79:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+8))
	if v303 == int32(0) {
		v316 = v290
		goto L77
	} else {
		goto L81
	}
L80:
	;
	v316 = v310
	goto L77
L81:
	;
	v306 = F_pairingheap_remove_first(m, v302)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	v310 = F_lappend(m, v290, v306-int32(12))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	v313 = v291 + int32(1)
	if v313 != v286 {
		v290 = v310
		v291 = v313
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v346 = v342
	v351 = v344
	goto L76
L86:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v365 = v363
	goto L53
L87:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v379 == int32(0) {
		goto L42
	} else {
		goto L88
	}
L88:
	;
	v383 = v365
	v384 = v379
	goto L46
L89:
	;
	v406 = F_list_delete_last(m, v383)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L11
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v418 = v403 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v402)+64)) = uint8(v418)
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[2]))
	if v421 != int32(2) {
		goto L43
	} else {
		goto L96
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v406
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[2]))
	if v410 == int32(0) {
		goto L44
	} else {
		goto L93
	}
L93:
	;
	F_pfree(m, v402)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	F_pfree(m, v401)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	goto L44
L96:
	;
	v424 = *(*float64)(unsafe.Add(mBase, uint32(v401)+32))
	v425 = *(*float64)(unsafe.Add(mBase, uint32(v21)+48))
	if base.F64_lt(v424, v425) != 0 {
		goto L44
	} else {
		goto L97
	}
L97:
	;
	goto L45
L98:
	;
	F_errmsg_internal(m, int32(_a_F_hnswgettuple_1), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_hnswgettuple_2), int32(214), int32(_a_F_hnswgettuple_3))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errmsg_internal(m, int32(_a_F_hnswgettuple_4), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_hnswgettuple_2), int32(219), int32(_a_F_hnswgettuple_3))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hnswrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+48)) = int64(-4503599627370496)
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = v9
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	F_MemoryContextReset(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if l1 == int32(0) {
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v20 <= int32(0) {
			} else {
				v24 = v20 * int32(56)
				if v24 == int32(0) {
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					base.MemoryCopy(m, v27, l1, v24)
				}
			}
		}
		if l3 == int32(0) {
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v32 <= int32(0) {
			} else {
				v36 = v32 * int32(56)
				if v36 == int32(0) {
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					base.MemoryCopy(m, v39, l3, v36)
				}
			}
		}
		return
	}
}
func F_hungarian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L5
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v374
	v378 = v374 - int32(1)
	if v378 <= v7 {
		goto L82
	} else {
		goto L83
	}
L2:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v369 + v368
	goto L1
L3:
	;
	if v125 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	v125 = v118
	goto L3
L5:
	;
	if v5 <= v7 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v118 = int32(0)
	goto L4
L7:
	;
	v125 = int32(-1)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v36 = int32(1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v21))))
	if base.Ui32(v38) < base.Ui32(int32(192)) {
		v95 = v38
		v96 = v36
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if int32(369) < v95 {
		v118 = v96
		goto L4
	} else {
		goto L23
	}
L11:
	;
	v42 = v7 + int32(1)
	if v42 == v5 {
		v95 = v38
		v96 = v36
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v21))))
	v47 = v45 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v38) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v21))))
	v63 = v61 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v38) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v51 = v7 + int32(2)
	if v51 != v5 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v95 = v38<<(uint(int32(6))%32)&int32(1984) | v47
	v96 = int32(2)
	goto L10
L17:
	;
	goto L16
L18:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v67))))
	v95 = v80&int32(63) | (v38<<(uint(int32(18))%32)&int32(_a_F_hungarian_UTF_8_stem_0) | v47<<(uint(int32(12))%32) | v63<<(uint(int32(6))%32))
	v96 = int32(4)
	goto L10
L19:
	;
	v67 = v7 + int32(3)
	if v67 != v5 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v95 = v38<<(uint(int32(12))%32)&int32(_a_F_hungarian_UTF_8_stem_1) | v47<<(uint(int32(6))%32) | v63
	v96 = int32(3)
	goto L10
L22:
	;
	goto L21
L23:
	;
	v100 = v95 - int32(97)
	if v100 < int32(0) {
		v118 = v96
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v100)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_UTF_8_stem[0]))))
	if int32(base.Ui32(v106)>>(uint(v100&int32(7))%32))&int32(1) == int32(0) {
		v118 = v96
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96 + v7
	goto L26
L26:
	;
	goto L6
L27:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = v139
	goto L32
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v270 = v7
	goto L57
L30:
	;
	if int32(0) <= v245 {
		v368 = v245
		goto L2
	} else {
		goto L54
	}
L31:
	;
	v245 = v216
	goto L30
L32:
	;
	if v140 <= v149 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v245 = int32(-1)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v156 = int32(1)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v141))))
	if base.Ui32(v158) < base.Ui32(int32(192)) {
		v215 = v158
		v216 = v156
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if int32(369) < v215 {
		goto L31
	} else {
		goto L50
	}
L38:
	;
	v162 = v149 + int32(1)
	if v162 == v140 {
		v215 = v158
		v216 = v156
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v141))))
	v167 = v165 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v158) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v141))))
	v183 = v181 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v158) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v171 = v149 + int32(2)
	if v171 != v140 {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v215 = v158<<(uint(int32(6))%32)&int32(1984) | v167
	v216 = int32(2)
	goto L37
L44:
	;
	goto L43
L45:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v187))))
	v215 = v200&int32(63) | (v158<<(uint(int32(18))%32)&int32(_a_F_hungarian_UTF_8_stem_0) | v167<<(uint(int32(12))%32) | v183<<(uint(int32(6))%32))
	v216 = int32(4)
	goto L37
L46:
	;
	v187 = v149 + int32(3)
	if v187 != v140 {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v215 = v158<<(uint(int32(12))%32)&int32(_a_F_hungarian_UTF_8_stem_1) | v167<<(uint(int32(6))%32) | v183
	v216 = int32(3)
	goto L37
L49:
	;
	goto L48
L50:
	;
	v220 = v215 - int32(97)
	if v220 < int32(0) {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v220)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_UTF_8_stem[0]))))
	if int32(base.Ui32(v226)>>(uint(v220&int32(7))%32))&int32(1) == int32(0) {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	v234 = v216 + v149
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v234
	v149 = v234
	goto L32
L54:
	;
	goto L1
L55:
	;
	if v365 < int32(0) {
		goto L1
	} else {
		goto L80
	}
L56:
	;
	v365 = v337
	goto L55
L57:
	;
	if v261 <= v270 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v365 = int32(-1)
	goto L55
L60:
	;
	goto L61
L61:
	;
	v277 = int32(1)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+v262))))
	if base.Ui32(v279) < base.Ui32(int32(192)) {
		v336 = v279
		v337 = v277
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if int32(369) < v336 {
		goto L75
	} else {
		goto L76
	}
L63:
	;
	v283 = v270 + int32(1)
	if v283 == v261 {
		v336 = v279
		v337 = v277
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v262))))
	v288 = v286 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v279) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v262))))
	v304 = v302 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v279) {
		goto L71
	} else {
		goto L72
	}
L66:
	;
	v292 = v270 + int32(2)
	if v292 != v261 {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v336 = v279<<(uint(int32(6))%32)&int32(1984) | v288
	v337 = int32(2)
	goto L62
L69:
	;
	goto L68
L70:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v308))))
	v336 = v321&int32(63) | (v279<<(uint(int32(18))%32)&int32(_a_F_hungarian_UTF_8_stem_0) | v288<<(uint(int32(12))%32) | v304<<(uint(int32(6))%32))
	v337 = int32(4)
	goto L62
L71:
	;
	v308 = v270 + int32(3)
	if v308 != v261 {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v336 = v279<<(uint(int32(12))%32)&int32(_a_F_hungarian_UTF_8_stem_1) | v288<<(uint(int32(6))%32) | v304
	v337 = int32(3)
	goto L62
L74:
	;
	goto L73
L75:
	;
	v354 = v337 + v270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v270 = v354
	goto L57
L76:
	;
	v341 = v336 - int32(97)
	if v341 < int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v341)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_UTF_8_stem[0]))))
	if int32(base.Ui32(v347)>>(uint(v341&int32(7))%32))&int32(1) != 0 {
		goto L56
	} else {
		goto L78
	}
L78:
	;
	goto L75
L80:
	;
	v368 = v365
	goto L2
L81:
	;
	return v898
L82:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v465
	v471 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_2), int32(44), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L85
	} else {
		goto L103
	}
L83:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380+v378))))
	if v382 != int32(108) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v388 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_3), int32(2), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	return int32(0)
L86:
	;
	if v388 == int32(0) {
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v394
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v394 < v396 {
		goto L82
	} else {
		goto L88
	}
L88:
	;
	v399 = v394 - int32(1)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v399 <= v400 {
		goto L82
	} else {
		goto L89
	}
L89:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v399))))
	if base.B2i32(v404&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v404)%32)&int32(106790108) == int32(0)) != 0 {
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v420 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_4), int32(23), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	if v420 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L92:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v424 + (v394 - v416)
	v428 = F_slice_del(m, l0)
	mBase = m.M
	if v428 < int32(0) {
		v898 = v428
		goto L81
	} else {
		goto L93
	}
L93:
	;
	v431 = int32(0)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v437 = F_skip_b_utf8(m, v433, v434, v435, int32(1))
	mBase = m.M
	if v437 < v431 {
		v460 = v431
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v460 < int32(0) {
		v898 = v460
		goto L81
	} else {
		goto L101
	}
L95:
	;
	goto L94
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v437
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v445 = F_skip_b_utf8(m, v442, v437, v443, int32(1))
	mBase = m.M
	if v445 < int32(0) {
		v460 = v431
		goto L95
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v445
	v451 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v451 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v457 = int32(1)
	goto L100
L99:
	;
	v457 = v451 >> (uint(int32(31)) % 32) & v451
	goto L100
L100:
	;
	v460 = v457
	goto L95
L101:
	;
	goto L82
L102:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v520
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v520-int32(2) <= v523 {
		goto L118
	} else {
		goto L119
	}
L103:
	;
	if v471 == int32(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v475 < v477 {
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v479 = F_slice_del(m, l0)
	mBase = m.M
	if v479 < int32(0) {
		v898 = v479
		goto L81
	} else {
		goto L106
	}
L106:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v482
	v485 = v482 - int32(1)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v485 <= v486 {
		goto L102
	} else {
		goto L107
	}
L107:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488+v485))))
	switch v490 - int32(161) {
	case 0, 8:
		goto L108
	default:
		goto L102
	}
L108:
	;
	v496 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_5), int32(2), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L85
	} else {
		goto L109
	}
L109:
	;
	if v496 == int32(0) {
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v500
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v500 < v502 {
		goto L102
	} else {
		goto L111
	}
L111:
	;
	switch v496 - int32(1) {
	case 0:
		goto L113
	case 1:
		goto L112
	default:
		goto L102
	}
L112:
	;
	v514 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_6))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L85
	} else {
		goto L116
	}
L113:
	;
	v508 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_7))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L85
	} else {
		goto L114
	}
L114:
	;
	if int32(0) <= v508 {
		goto L102
	} else {
		goto L115
	}
L115:
	;
	v898 = v508
	goto L81
L116:
	;
	if v514 < int32(0) {
		v898 = v514
		goto L81
	} else {
		goto L117
	}
L117:
	;
	goto L102
L118:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v561
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v561-int32(3) <= v564 {
		goto L130
	} else {
		goto L131
	}
L119:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527+v520-int32(1)))))
	switch v531 - int32(110) {
	case 0, 6:
		goto L120
	default:
		goto L118
	}
L120:
	;
	v537 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_8), int32(3), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L85
	} else {
		goto L121
	}
L121:
	;
	if v537 == int32(0) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v541
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v541 < v543 {
		goto L118
	} else {
		goto L123
	}
L123:
	;
	switch v537 - int32(1) {
	case 0:
		goto L125
	case 1:
		goto L124
	default:
		goto L118
	}
L124:
	;
	v555 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_9))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L85
	} else {
		goto L128
	}
L125:
	;
	v549 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_10))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L85
	} else {
		goto L126
	}
L126:
	;
	if int32(0) <= v549 {
		goto L118
	} else {
		goto L127
	}
L127:
	;
	v898 = v549
	goto L81
L128:
	;
	if v555 < int32(0) {
		v898 = v555
		goto L81
	} else {
		goto L129
	}
L129:
	;
	goto L118
L130:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v605
	v610 = v605 - int32(1)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v610 <= v611 {
		v694 = v2
		goto L144
	} else {
		goto L145
	}
L131:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+v561-int32(1)))))
	if v572 != int32(108) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v578 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_11), int32(6), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L85
	} else {
		goto L133
	}
L133:
	;
	if v578 == int32(0) {
		goto L130
	} else {
		goto L134
	}
L134:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v582
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v582 < v584 {
		goto L130
	} else {
		goto L135
	}
L135:
	;
	switch v578 - int32(1) {
	case 0:
		goto L138
	case 1:
		goto L137
	case 2:
		goto L136
	default:
		goto L130
	}
L136:
	;
	v599 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_12))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L85
	} else {
		goto L142
	}
L137:
	;
	v593 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_13))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L85
	} else {
		goto L140
	}
L138:
	;
	v588 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v588 {
		goto L130
	} else {
		goto L139
	}
L139:
	;
	v898 = v588
	goto L81
L140:
	;
	if int32(0) <= v593 {
		goto L130
	} else {
		goto L141
	}
L141:
	;
	v898 = v593
	goto L81
L142:
	;
	if v599 < int32(0) {
		v898 = v599
		goto L81
	} else {
		goto L143
	}
L143:
	;
	goto L130
L144:
	;
	if v694 < int32(0) {
		v898 = v694
		goto L81
	} else {
		goto L162
	}
L145:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613+v610))))
	switch v615 - int32(161) {
	case 0, 8:
		goto L146
	default:
		v694 = v2
		goto L144
	}
L146:
	;
	v621 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_14), int32(2), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L85
	} else {
		goto L147
	}
L147:
	;
	if v621 == int32(0) {
		v694 = v2
		goto L144
	} else {
		goto L148
	}
L148:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v625
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v625 < v627 {
		v694 = v2
		goto L144
	} else {
		goto L149
	}
L149:
	;
	v630 = v625 - int32(1)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v630 <= v631 {
		v694 = v2
		goto L144
	} else {
		goto L150
	}
L150:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633+v630))))
	if base.B2i32(v635&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v635)%32)&int32(106790108) == int32(0)) != 0 {
		v694 = v2
		goto L144
	} else {
		goto L151
	}
L151:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v651 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_4), int32(23), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L85
	} else {
		goto L152
	}
L152:
	;
	if v651 == int32(0) {
		v694 = v2
		goto L144
	} else {
		goto L153
	}
L153:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v655 + (v625 - v647)
	v659 = F_slice_del(m, l0)
	mBase = m.M
	if v659 < int32(0) {
		v694 = v659
		goto L144
	} else {
		goto L154
	}
L154:
	;
	v662 = int32(0)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v668 = F_skip_b_utf8(m, v664, v665, v666, int32(1))
	mBase = m.M
	if v668 < v662 {
		v691 = v662
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v694 = v691
	goto L144
L156:
	;
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v668
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v668
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v676 = F_skip_b_utf8(m, v673, v668, v674, int32(1))
	mBase = m.M
	if v676 < int32(0) {
		v691 = v662
		goto L156
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v676
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v676
	v682 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v682 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v688 = int32(1)
	goto L161
L160:
	;
	v688 = v682 >> (uint(int32(31)) % 32) & v682
	goto L161
L161:
	;
	v691 = v688
	goto L156
L162:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v697
	v699 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v697
	v703 = v697 - int32(1)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v703 <= v704 {
		v744 = v699
		goto L163
	} else {
		goto L164
	}
L163:
	;
	if v744 < int32(0) {
		v898 = v744
		goto L81
	} else {
		goto L178
	}
L164:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706+v703))))
	if base.B2i32(v708 != int32(169))&base.B2i32(v708 != int32(105)) != 0 {
		v744 = v699
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v717 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_15), int32(12), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L85
	} else {
		goto L166
	}
L166:
	;
	if v717 == int32(0) {
		v744 = v699
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v721
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v721 < v723 {
		v744 = v699
		goto L163
	} else {
		goto L168
	}
L168:
	;
	switch v717 - int32(1) {
	case 0:
		goto L172
	case 1:
		goto L171
	case 2:
		goto L170
	default:
		goto L169
	}
L169:
	;
	v744 = int32(1)
	goto L163
L170:
	;
	v738 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_16))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L85
	} else {
		goto L176
	}
L171:
	;
	v732 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_17))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L85
	} else {
		goto L174
	}
L172:
	;
	v727 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v727 {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v744 = v727
	goto L163
L174:
	;
	if int32(0) <= v732 {
		goto L169
	} else {
		goto L175
	}
L175:
	;
	v744 = v732
	goto L163
L176:
	;
	if v738 < int32(0) {
		v744 = v738
		goto L163
	} else {
		goto L177
	}
L177:
	;
	goto L169
L178:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v749
	v751 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v749
	v757 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_18), int32(31), v751)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L85
	} else {
		goto L180
	}
L179:
	;
	if v784 < int32(0) {
		v898 = v784
		goto L81
	} else {
		goto L192
	}
L180:
	;
	if v757 == int32(0) {
		v784 = v751
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v761
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v761 < v763 {
		v784 = v751
		goto L179
	} else {
		goto L182
	}
L182:
	;
	switch v757 - int32(1) {
	case 0:
		goto L186
	case 1:
		goto L185
	case 2:
		goto L184
	default:
		goto L183
	}
L183:
	;
	v784 = int32(1)
	goto L179
L184:
	;
	v778 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_19))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L85
	} else {
		goto L190
	}
L185:
	;
	v772 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_20))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L85
	} else {
		goto L188
	}
L186:
	;
	v767 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v767 {
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v784 = v767
	goto L179
L188:
	;
	if int32(0) <= v772 {
		goto L183
	} else {
		goto L189
	}
L189:
	;
	v784 = v772
	goto L179
L190:
	;
	if v778 < int32(0) {
		v784 = v778
		goto L179
	} else {
		goto L191
	}
L191:
	;
	goto L183
L192:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v788
	v790 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v788
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v788 <= v793 {
		v841 = v790
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if v841 < int32(0) {
		v898 = v841
		goto L81
	} else {
		goto L208
	}
L194:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v797 = int32(1)
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v795+v788-v797))))
	if base.B2i32(v799&int32(224) != int32(96))|base.B2i32(v797<<(uint(v799)%32)&int32(_a_F_hungarian_UTF_8_stem_21) == int32(0)) != 0 {
		v841 = v790
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v814 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_22), int32(42), int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L85
	} else {
		goto L196
	}
L196:
	;
	if v814 == int32(0) {
		v841 = v790
		goto L193
	} else {
		goto L197
	}
L197:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v818
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v818 < v820 {
		v841 = v790
		goto L193
	} else {
		goto L198
	}
L198:
	;
	switch v814 - int32(1) {
	case 0:
		goto L202
	case 1:
		goto L201
	case 2:
		goto L200
	default:
		goto L199
	}
L199:
	;
	v841 = int32(1)
	goto L193
L200:
	;
	v835 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_23))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L85
	} else {
		goto L206
	}
L201:
	;
	v829 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_24))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L85
	} else {
		goto L204
	}
L202:
	;
	v824 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v824 {
		goto L199
	} else {
		goto L203
	}
L203:
	;
	v841 = v824
	goto L193
L204:
	;
	if int32(0) <= v829 {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	v841 = v829
	goto L193
L206:
	;
	if v835 < int32(0) {
		v841 = v835
		goto L193
	} else {
		goto L207
	}
L207:
	;
	goto L199
L208:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v846
	v848 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v846
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v846 <= v851 {
		v890 = v848
		goto L209
	} else {
		goto L210
	}
L209:
	;
	if v890 < int32(0) {
		v898 = v890
		goto L81
	} else {
		goto L224
	}
L210:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853+v846-int32(1)))))
	if v857 != int32(107) {
		v890 = v848
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v863 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_25), int32(7), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L85
	} else {
		goto L212
	}
L212:
	;
	if v863 == int32(0) {
		v890 = v848
		goto L209
	} else {
		goto L213
	}
L213:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v867
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v867 < v869 {
		v890 = v848
		goto L209
	} else {
		goto L214
	}
L214:
	;
	switch v863 - int32(1) {
	case 0:
		goto L218
	case 1:
		goto L217
	case 2:
		goto L216
	default:
		goto L215
	}
L215:
	;
	v890 = int32(1)
	goto L209
L216:
	;
	v885 = F_slice_del(m, l0)
	mBase = m.M
	if v885 < int32(0) {
		v890 = v885
		goto L209
	} else {
		goto L223
	}
L217:
	;
	v881 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_26))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L85
	} else {
		goto L221
	}
L218:
	;
	v875 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_27))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L85
	} else {
		goto L219
	}
L219:
	;
	if int32(0) <= v875 {
		goto L215
	} else {
		goto L220
	}
L220:
	;
	v890 = v875
	goto L209
L221:
	;
	if int32(0) <= v881 {
		goto L215
	} else {
		goto L222
	}
L222:
	;
	v890 = v881
	goto L209
L223:
	;
	goto L215
L224:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v895
	v898 = int32(1)
	goto L81
}
func F_hypothetical_dense_rank_final(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var __phi247 int32
	_ = __phi247
	var v248 int32
	_ = v248
	var __phi248 int32
	_ = __phi248
	var v258 int64
	_ = v258
	var __phi258 int64
	_ = __phi258
	var v259 int64
	_ = v259
	var __phi259 int64
	_ = __phi259
	var v260 int64
	_ = v260
	var __phi260 int64
	_ = __phi260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int64
	_ = v272
	var v279 int32
	_ = v279
	var v282 int64
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int64
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v365 int64
	_ = v365
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	v2 = int32(0)
	v14 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v14
	v24 = int64(1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v25 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L8
	} else {
		goto L63
	}
L2:
	;
	v29 = l0 + int32(24)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v365 = v24
	goto L4
L4:
	;
	m.G0 = v19 + int32(16)
	return v365
L5:
	;
	v35 = int32(_a_F_hypothetical_dense_rank_final_0)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v38
	v40 = F_CreateStandaloneExprContext(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v52 = v32
	goto L7
L7:
	;
	v53 = int32(1)
	v54 = v21 - v53
	if v54&v53 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	return int64(0)
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v40
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v36
	v52 = v47
	goto L7
L10:
	;
	v58 = v54 >> (uint(int32(1)) % 32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	F_hypothetical_check_argtypes(m, l0, v58, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+48))
	if v64 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v90 = v64
	v91 = v63
	goto L14
L13:
	;
	v65 = int32(_a_F_hypothetical_dense_rank_final_0)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+36))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
	v79 = F_execTuplesMatchPrepare(m, v73, v68-int32(1), v67, v76, v77, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	m.T0[v94].(func(*base.Module, int32))(m, v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v66
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+48)) = v79
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v90 = v79
	v91 = v85
	goto L14
L16:
	;
	v97 = int32(0)
	if v58 <= v97 {
		v203 = v97
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v204+v203<<(uint(int32(3))%32)))) = int64(-1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	v212 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v210+v203))) = uint8(v212)
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+4)))
	v216 = v214 & int32(_a_F_hypothetical_dense_rank_final_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v92)+4)) = uint16(v216)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	*(*uint16)(unsafe.Add(mBase, uint32(v92)+6)) = uint16(v219)
	goto L26
L18:
	;
	v100 = int32(0)
	if v54 != int32(2) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v108 = v100
	v117 = v2
	goto L22
L20:
	;
	v159 = v100
	goto L21
L21:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v180 = v29 + v159<<(uint(int32(4))%32)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v180)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v174+v159<<(uint(int32(3))%32)))) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v183+v159))) = uint8(v185)
	v203 = v58
	goto L17
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v124 = int32(3)
	v128 = v108 | int32(1)
	v129 = int32(4)
	v131 = v29 + v128<<(uint(v129)%32)
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
	*(*int64)(unsafe.Add(mBase, uint32(v123+v108<<(uint(v124)%32)))) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v134+v108))) = uint8(v136)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v142 = int32(2)
	v143 = v108 + v142
	v146 = v29 + v143<<(uint(v129)%32)
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
	*(*int64)(unsafe.Add(mBase, uint32(v138+v128<<(uint(v124)%32)))) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v128+v149))) = uint8(v151)
	v154 = v117 + v142
	if v154 != v58&int32(2147483646) {
		v108 = v143
		v117 = v154
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v58&int32(1) == int32(0) {
		v203 = v58
		goto L17
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v159 = v143
	goto L21
L26:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	F_tuplesort_puttupleslot(m, v221, v92)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	F_tuplesort_performsort(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)) = uint8(v227)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	v232 = F_MakeSingleTupleTableSlot(m, v230, int32(_a_F_hypothetical_dense_rank_final_2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v235 = int32(1)
	v237 = F_tuplesort_gettupleslot(m, v234, v235, v235, v92, v19)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	m.T0[v342].(func(*base.Module, int32))(m, v325)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L60
	}
L31:
	;
	if v237 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v325 = v92
	v326 = v232
	v338 = v24
	v339 = v14
	goto L30
L33:
	;
	goto L34
L34:
	;
	__phi247 = v92
	__phi248 = v232
	__phi258 = v24
	__phi259 = v14
	__phi260 = v14
	v247 = __phi247
	v248 = __phi248
	v258 = __phi258
	v259 = __phi259
	v260 = __phi260
	goto L35
L35:
	;
	v261 = int32(*(*int16)(unsafe.Add(mBase, uint32(v247)+6)))
	if v261 <= v58 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v325 = v248
	v326 = v247
	v338 = v319
	v339 = v311
	goto L30
L37:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	m.T0[v264].(func(*base.Module, int32, int32))(m, v247, v58+int32(1))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L8
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v247)+20))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v58))))
	if v269 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v247
	if v248 == int32(0) {
		v311 = v259
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
	v272 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v270+v58<<(uint(int32(3))%32)))))
	if v272 == int64(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v325 = v247
	v326 = v248
	v338 = v258
	v339 = v259
	goto L30
L44:
	;
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[1]))
	if v315 != 0 {
		goto L54
	} else {
		goto L55
	}
L45:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+4)))
	if v279&int32(2) != 0 {
		v311 = v259
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	if v282 != v260 {
		v311 = v259
		goto L44
	} else {
		goto L47
	}
L47:
	;
	if v90 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	F_MemoryContextReset(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v291 = int32(_a_F_hypothetical_dense_rank_final_0)
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v294
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
	v299 = m.T0[v298].(func(*base.Module, int32, int32, int32) int64)(m, v90, v52, v19+int32(15))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L8
	} else {
		goto L52
	}
L51:
	;
	v311 = v259 + int64(1)
	goto L44
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v292
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	F_MemoryContextReset(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v311 = v259 + base.I64_extend_i32_u(base.B2i32(v299 != int64(0)))
	goto L44
L54:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v319 = v258 + int64(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v321 = int32(1)
	v323 = F_tuplesort_gettupleslot(m, v320, v321, v321, v248, v19)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L8
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	if v323 != 0 {
		__phi247 = v248
		__phi248 = v247
		__phi258 = v319
		__phi259 = v311
		__phi260 = v313
		v247 = __phi247
		v248 = __phi248
		v258 = __phi258
		v259 = __phi259
		v260 = __phi260
		goto L35
	} else {
		goto L59
	}
L59:
	;
	goto L36
L60:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	m.T0[v346].(func(*base.Module, int32))(m, v326)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L61
	}
L61:
	;
	F_ExecDropSingleTupleTableSlot(m, v232)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v365 = v338 - v339
	goto L4
L63:
	;
	F_errmsg_internal(m, int32(_a_F_hypothetical_dense_rank_final_3), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_hypothetical_dense_rank_final_4), int32(1333), int32(_a_F_hypothetical_dense_rank_final_5))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
