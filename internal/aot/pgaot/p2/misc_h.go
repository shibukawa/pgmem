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
	F_errfinish(m, int32(_a_F_HoldPinnedPortals_1), int32(1232), int32(_a_F_HoldPinnedPortals_2))
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
	F_errfinish(m, int32(_a_F_HoldPinnedPortals_1), int32(1236), int32(_a_F_HoldPinnedPortals_2))
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
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
func F_has_parameter_privilege_name(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = F_convert_any_priv_string(m, v10, int32(_a_F_has_parameter_privilege_name_0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_has_parameter_privilege_name[0]))
				v17 = F_text_to_cstring(m, v5)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = F_pg_parameter_aclcheck(m, v17, v16, v13)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v19 == int32(0))
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
	v4 = Fn13919(m, l0, l1, int32(1))
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v4 = F_superuser_arg(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v11 = F_SearchSysCache1(m, int32(11), l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					return int32(0)
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v18)+73)))
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = v20
						return v24 & int32(1)
					}
				}
			}
		} else {
			v24 = int32(1)
			return v24 & int32(1)
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
func F_hashbpchar(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
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
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L130
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
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L125
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
	v23 = v10 + int32(1)
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = v23
	goto L12
L11:
	;
	v26 = v10 + int32(4)
	goto L12
L12:
	;
	if v17 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v58 = v53
	goto L24
L14:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v32 == int32(18) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v43 = int32(1)
	if v19 != 0 {
		v53 = int32(base.Ui32(v17)>>(uint(v43)%32)) - v43
		goto L13
	} else {
		goto L23
	}
L17:
	;
	v35 = int32(16)
	goto L19
L18:
	;
	v35 = int32(0)
	goto L19
L19:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v42 = int32(4)
	goto L22
L21:
	;
	v42 = v35
	goto L22
L22:
	;
	v53 = v42
	goto L13
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L13
L24:
	;
	if v58 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v75 = F_pg_newlocale_from_collation(m, v14)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v74 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v68 = v58 - int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v68))))
	if v70 == int32(32) {
		v58 = v68
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v74 = v58
	goto L26
L31:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v625 != v10 {
		goto L121
	} else {
		goto L122
	}
L32:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v77 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v85 = v74 - int32(1636608432)
	if v21&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	goto L35
L35:
	;
	v344 = int32(0)
	v346 = F_pg_strnxfrm(m, v344, v344, v21, v74, v75)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L76
	}
L36:
	;
	v621 = v339 ^ v331 - base.I32_rotl(v339, int32(24))
	goto L31
L37:
	;
	v317 = int32(14)
	v319 = v313 ^ v314 - base.I32_rotl(v313, v317)
	v323 = v319 ^ v312 - base.I32_rotl(v319, int32(11))
	v327 = v323 ^ v313 - base.I32_rotl(v323, int32(25))
	v331 = v327 ^ v319 - base.I32_rotl(v327, int32(16))
	v335 = v331 ^ v323 - base.I32_rotl(v331, int32(4))
	v339 = v335 ^ v327 - base.I32_rotl(v335, v317)
	goto L36
L38:
	;
	switch v243 - int32(1) {
	case 0:
		v305 = v244
		v306 = v245
		v307 = v246
		goto L65
	case 1:
		v298 = v244
		v299 = v245
		v300 = v246
		goto L66
	case 2:
		v291 = v244
		v292 = v245
		v293 = v246
		goto L67
	case 3:
		v285 = v245
		v286 = v246
		goto L68
	case 4:
		v281 = v245
		v282 = v246
		goto L69
	case 5:
		v275 = v245
		v276 = v246
		goto L70
	case 6:
		v269 = v245
		v270 = v246
		goto L71
	case 7:
		v264 = v246
		goto L72
	case 8:
		v259 = v246
		goto L73
	case 9:
		v254 = v246
		goto L74
	case 10:
		goto L75
	default:
		v312 = v244
		v313 = v245
		v314 = v246
		goto L37
	}
L39:
	;
	v194 = v21
	v195 = v74
	v196 = v85
	v197 = v85
	v198 = v85
	goto L62
L40:
	;
	if base.Ui32(int32(11)) < base.Ui32(v74) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(v74) < base.Ui32(int32(12)) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v242 = v21
	v243 = v74
	v244 = v85
	v245 = v85
	v246 = v85
	goto L38
L44:
	;
	switch v141 - int32(1) {
	case 0:
		v191 = v142
		goto L51
	case 1:
		v186 = v142
		goto L52
	case 2:
		goto L53
	case 3:
		v179 = v143
		goto L54
	case 4:
		v176 = v143
		goto L55
	case 5:
		v171 = v143
		goto L56
	case 6:
		goto L57
	case 7:
		v162 = v144
		goto L58
	case 8:
		v157 = v144
		goto L59
	case 9:
		v152 = v144
		goto L60
	case 10:
		goto L61
	default:
		v312 = v142
		v313 = v143
		v314 = v144
		goto L37
	}
L45:
	;
	v140 = v21
	v141 = v74
	v142 = v85
	v143 = v85
	v144 = v85
	goto L44
L46:
	;
	goto L47
L47:
	;
	v92 = v21
	v93 = v74
	v94 = v85
	v95 = v85
	v96 = v85
	goto L48
L48:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v99 = v98 + v95
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v103 = v102 + v96
	v105 = int32(4)
	v107 = v100 + v94 - v103 ^ base.I32_rotl(v103, v105)
	v111 = v99 - v107 ^ base.I32_rotl(v107, int32(6))
	v112 = v103 + v99
	v113 = v107 + v112
	v114 = v111 + v113
	v118 = v112 - v111 ^ base.I32_rotl(v111, int32(8))
	v122 = v113 - v118 ^ base.I32_rotl(v118, int32(16))
	v126 = v114 - v122 ^ base.I32_rotl(v122, int32(19))
	v127 = v118 + v114
	v128 = v122 + v127
	v129 = v126 + v128
	v133 = v127 - v126 ^ base.I32_rotl(v126, v105)
	v134 = int32(12)
	v135 = v92 + v134
	v137 = v93 - v134
	if base.Ui32(int32(11)) < base.Ui32(v137) {
		v92 = v135
		v93 = v137
		v94 = v128
		v95 = v129
		v96 = v133
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v140 = v135
	v141 = v137
	v142 = v128
	v143 = v129
	v144 = v133
	goto L44
L50:
	;
	goto L49
L51:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v312 = v191 + v192
	v313 = v143
	v314 = v144
	goto L37
L52:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)))
	v191 = v187<<(uint(int32(8))%32) + v186
	goto L51
L53:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+2)))
	v186 = v182<<(uint(int32(16))%32) + v142
	goto L52
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v312 = v180 + v142
	v313 = v179
	v314 = v144
	goto L37
L55:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+4)))
	v179 = v176 + v177
	goto L54
L56:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+5)))
	v176 = v172<<(uint(int32(8))%32) + v171
	goto L55
L57:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+6)))
	v171 = v167<<(uint(int32(16))%32) + v143
	goto L56
L58:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v312 = v163 + v142
	v313 = v165 + v143
	v314 = v162
	goto L37
L59:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+8)))
	v162 = v158<<(uint(int32(8))%32) + v157
	goto L58
L60:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+9)))
	v157 = v153<<(uint(int32(16))%32) + v152
	goto L59
L61:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+10)))
	v152 = v148<<(uint(int32(24))%32) + v144
	goto L60
L62:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	v201 = v200 + v197
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v205 = v204 + v198
	v207 = int32(4)
	v209 = v202 + v196 - v205 ^ base.I32_rotl(v205, v207)
	v213 = v201 - v209 ^ base.I32_rotl(v209, int32(6))
	v214 = v205 + v201
	v215 = v209 + v214
	v216 = v213 + v215
	v220 = v214 - v213 ^ base.I32_rotl(v213, int32(8))
	v224 = v215 - v220 ^ base.I32_rotl(v220, int32(16))
	v228 = v216 - v224 ^ base.I32_rotl(v224, int32(19))
	v229 = v220 + v216
	v230 = v224 + v229
	v231 = v228 + v230
	v235 = v229 - v228 ^ base.I32_rotl(v228, v207)
	v236 = int32(12)
	v237 = v194 + v236
	v239 = v195 - v236
	if base.Ui32(int32(11)) < base.Ui32(v239) {
		v194 = v237
		v195 = v239
		v196 = v230
		v197 = v231
		v198 = v235
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v242 = v237
	v243 = v239
	v244 = v230
	v245 = v231
	v246 = v235
	goto L38
L64:
	;
	goto L63
L65:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v312 = v305 + v308
	v313 = v306
	v314 = v307
	goto L37
L66:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	v305 = v301<<(uint(int32(8))%32) + v298
	v306 = v299
	v307 = v300
	goto L65
L67:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+2)))
	v298 = v294<<(uint(int32(16))%32) + v291
	v299 = v292
	v300 = v293
	goto L66
L68:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+3)))
	v291 = v287<<(uint(int32(24))%32) + v244
	v292 = v285
	v293 = v286
	goto L67
L69:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+4)))
	v285 = v281 + v283
	v286 = v282
	goto L68
L70:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+5)))
	v281 = v277<<(uint(int32(8))%32) + v275
	v282 = v276
	goto L69
L71:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+6)))
	v275 = v271<<(uint(int32(16))%32) + v269
	v276 = v270
	goto L70
L72:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+7)))
	v269 = v265<<(uint(int32(24))%32) + v245
	v270 = v264
	goto L71
L73:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+8)))
	v264 = v260<<(uint(int32(8))%32) + v259
	goto L72
L74:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+9)))
	v259 = v255<<(uint(int32(16))%32) + v254
	goto L73
L75:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+10)))
	v254 = v250<<(uint(int32(24))%32) + v246
	goto L74
L76:
	;
	v349 = v346 + int32(1)
	v350 = F_palloc(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v352 = F_pg_strnxfrm(m, v350, v349, v21, v74, v75)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(v346) < base.Ui32(v352) {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v360 = v349 - int32(1636608432)
	if v350&int32(3) != 0 {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	F_pfree(m, v350)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L120
	}
L81:
	;
	v592 = int32(14)
	v594 = v588 ^ v589 - base.I32_rotl(v588, v592)
	v598 = v594 ^ v587 - base.I32_rotl(v594, int32(11))
	v602 = v598 ^ v588 - base.I32_rotl(v598, int32(25))
	v606 = v602 ^ v594 - base.I32_rotl(v602, int32(16))
	v610 = v606 ^ v598 - base.I32_rotl(v606, int32(4))
	v614 = v610 ^ v602 - base.I32_rotl(v610, v592)
	goto L80
L82:
	;
	switch v518 - int32(1) {
	case 0:
		v580 = v519
		v581 = v520
		v582 = v521
		goto L109
	case 1:
		v573 = v519
		v574 = v520
		v575 = v521
		goto L110
	case 2:
		v566 = v519
		v567 = v520
		v568 = v521
		goto L111
	case 3:
		v560 = v520
		v561 = v521
		goto L112
	case 4:
		v556 = v520
		v557 = v521
		goto L113
	case 5:
		v550 = v520
		v551 = v521
		goto L114
	case 6:
		v544 = v520
		v545 = v521
		goto L115
	case 7:
		v539 = v521
		goto L116
	case 8:
		v534 = v521
		goto L117
	case 9:
		v529 = v521
		goto L118
	case 10:
		goto L119
	default:
		v587 = v519
		v588 = v520
		v589 = v521
		goto L81
	}
L83:
	;
	v469 = v350
	v470 = v349
	v471 = v360
	v472 = v360
	v473 = v360
	goto L106
L84:
	;
	if base.Ui32(int32(11)) < base.Ui32(v349) {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(v349) < base.Ui32(int32(12)) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v517 = v350
	v518 = v349
	v519 = v360
	v520 = v360
	v521 = v360
	goto L82
L88:
	;
	switch v416 - int32(1) {
	case 0:
		v466 = v417
		goto L95
	case 1:
		v461 = v417
		goto L96
	case 2:
		goto L97
	case 3:
		v454 = v418
		goto L98
	case 4:
		v451 = v418
		goto L99
	case 5:
		v446 = v418
		goto L100
	case 6:
		goto L101
	case 7:
		v437 = v419
		goto L102
	case 8:
		v432 = v419
		goto L103
	case 9:
		v427 = v419
		goto L104
	case 10:
		goto L105
	default:
		v587 = v417
		v588 = v418
		v589 = v419
		goto L81
	}
L89:
	;
	v415 = v350
	v416 = v349
	v417 = v360
	v418 = v360
	v419 = v360
	goto L88
L90:
	;
	goto L91
L91:
	;
	v367 = v350
	v368 = v349
	v369 = v360
	v370 = v360
	v371 = v360
	goto L92
L92:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v374 = v373 + v370
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	v378 = v377 + v371
	v380 = int32(4)
	v382 = v375 + v369 - v378 ^ base.I32_rotl(v378, v380)
	v386 = v374 - v382 ^ base.I32_rotl(v382, int32(6))
	v387 = v378 + v374
	v388 = v382 + v387
	v389 = v386 + v388
	v393 = v387 - v386 ^ base.I32_rotl(v386, int32(8))
	v397 = v388 - v393 ^ base.I32_rotl(v393, int32(16))
	v401 = v389 - v397 ^ base.I32_rotl(v397, int32(19))
	v402 = v393 + v389
	v403 = v397 + v402
	v404 = v401 + v403
	v408 = v402 - v401 ^ base.I32_rotl(v401, v380)
	v409 = int32(12)
	v410 = v367 + v409
	v412 = v368 - v409
	if base.Ui32(int32(11)) < base.Ui32(v412) {
		v367 = v410
		v368 = v412
		v369 = v403
		v370 = v404
		v371 = v408
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v415 = v410
	v416 = v412
	v417 = v403
	v418 = v404
	v419 = v408
	goto L88
L94:
	;
	goto L93
L95:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	v587 = v466 + v467
	v588 = v418
	v589 = v419
	goto L81
L96:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+1)))
	v466 = v462<<(uint(int32(8))%32) + v461
	goto L95
L97:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+2)))
	v461 = v457<<(uint(int32(16))%32) + v417
	goto L96
L98:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v587 = v455 + v417
	v588 = v454
	v589 = v419
	goto L81
L99:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+4)))
	v454 = v451 + v452
	goto L98
L100:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+5)))
	v451 = v447<<(uint(int32(8))%32) + v446
	goto L99
L101:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+6)))
	v446 = v442<<(uint(int32(16))%32) + v418
	goto L100
L102:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	v587 = v438 + v417
	v588 = v440 + v418
	v589 = v437
	goto L81
L103:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+8)))
	v437 = v433<<(uint(int32(8))%32) + v432
	goto L102
L104:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+9)))
	v432 = v428<<(uint(int32(16))%32) + v427
	goto L103
L105:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+10)))
	v427 = v423<<(uint(int32(24))%32) + v419
	goto L104
L106:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	v476 = v475 + v472
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v469)+8))
	v480 = v479 + v473
	v482 = int32(4)
	v484 = v477 + v471 - v480 ^ base.I32_rotl(v480, v482)
	v488 = v476 - v484 ^ base.I32_rotl(v484, int32(6))
	v489 = v480 + v476
	v490 = v484 + v489
	v491 = v488 + v490
	v495 = v489 - v488 ^ base.I32_rotl(v488, int32(8))
	v499 = v490 - v495 ^ base.I32_rotl(v495, int32(16))
	v503 = v491 - v499 ^ base.I32_rotl(v499, int32(19))
	v504 = v495 + v491
	v505 = v499 + v504
	v506 = v503 + v505
	v510 = v504 - v503 ^ base.I32_rotl(v503, v482)
	v511 = int32(12)
	v512 = v469 + v511
	v514 = v470 - v511
	if base.Ui32(int32(11)) < base.Ui32(v514) {
		v469 = v512
		v470 = v514
		v471 = v505
		v472 = v506
		v473 = v510
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v517 = v512
	v518 = v514
	v519 = v505
	v520 = v506
	v521 = v510
	goto L82
L108:
	;
	goto L107
L109:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	v587 = v580 + v583
	v588 = v581
	v589 = v582
	goto L81
L110:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+1)))
	v580 = v576<<(uint(int32(8))%32) + v573
	v581 = v574
	v582 = v575
	goto L109
L111:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+2)))
	v573 = v569<<(uint(int32(16))%32) + v566
	v574 = v567
	v575 = v568
	goto L110
L112:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+3)))
	v566 = v562<<(uint(int32(24))%32) + v519
	v567 = v560
	v568 = v561
	goto L111
L113:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+4)))
	v560 = v556 + v558
	v561 = v557
	goto L112
L114:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+5)))
	v556 = v552<<(uint(int32(8))%32) + v550
	v557 = v551
	goto L113
L115:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+6)))
	v550 = v546<<(uint(int32(16))%32) + v544
	v551 = v545
	goto L114
L116:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+7)))
	v544 = v540<<(uint(int32(24))%32) + v520
	v545 = v539
	goto L115
L117:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+8)))
	v539 = v535<<(uint(int32(8))%32) + v534
	goto L116
L118:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+9)))
	v534 = v530<<(uint(int32(16))%32) + v529
	goto L117
L119:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+10)))
	v529 = v525<<(uint(int32(24))%32) + v521
	goto L118
L120:
	;
	v621 = v614 ^ v606 - base.I32_rotl(v614, int32(24))
	goto L31
L121:
	;
	F_pfree(m, v10)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	return v621
L124:
	;
	goto L123
L125:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_hashbpchar_0), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errhint(m, int32(_a_F_hashbpchar_1), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_hashbpchar_2), int32(1001), int32(_a_F_hashbpchar_3))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errmsg_internal(m, int32(_a_F_hashbpchar_4), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_hashbpchar_2), int32(1025), int32(_a_F_hashbpchar_3))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashchar(m *base.Module, l0 int32) int32 {
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
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+20)))
	v7 = int32(711645284)
	v10 = v2 - int32(1636608428) ^ v7 - int32(1455628627)
	v15 = v10 ^ int32(-1636608428) - base.I32_rotl(v10, int32(25))
	v20 = v15 ^ v7 - base.I32_rotl(v15, int32(16))
	v24 = v20 ^ v10 - base.I32_rotl(v20, int32(4))
	v28 = v24 ^ v15 - base.I32_rotl(v24, int32(14))
	return v28 ^ v20 - base.I32_rotl(v28, int32(24))
}
func F_hashfloat4(m *base.Module, l0 int32) int32 {
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
	var v289 int32
	_ = v289
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
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
		v289 = v284 ^ v276 - base.I32_rotl(v284, int32(24))
	} else {
		v289 = int32(0)
	}
	m.G0 = v8 + int32(16)
	return v289
}
func F_hashint4extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		v11 = int32(-1636608428)
		v50 = v11
		v51 = v11
		v54 = int32(0)
	} else {
		v14 = base.I32_wrap_i64(v4)
		v16 = v14 + int32(1021750440)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v27 = v14 - v21 - int32(1636608428) ^ base.I32_rotl(v21, int32(6))
		v31 = v16 - v27 ^ base.I32_rotl(v27, int32(8))
		v32 = v21 + v16
		v33 = v27 + v32
		v34 = v31 + v33
		v38 = v32 - v31 ^ base.I32_rotl(v31, int32(16))
		v42 = v33 - v38 ^ base.I32_rotl(v38, int32(19))
		v47 = v38 + v34
		v48 = v42 + v47
		v50 = v48
		v51 = v47
		v54 = v34 - v42 ^ base.I32_rotl(v42, int32(4)) ^ v48
	}
	v55 = int32(14)
	v57 = v54 - base.I32_rotl(v50, v55)
	v62 = v57 ^ (v2 + v51) - base.I32_rotl(v57, int32(11))
	v66 = v50 ^ v62 - base.I32_rotl(v62, int32(25))
	v70 = v66 ^ v57 - base.I32_rotl(v66, int32(16))
	v74 = v70 ^ v62 - base.I32_rotl(v70, int32(4))
	v78 = v74 ^ v66 - base.I32_rotl(v74, v55)
	v88 = F_Int64GetDatum(m, base.I64_extend_i32_u(v78)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v78^v70-base.I32_rotl(v78, int32(24))))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		return int32(0)
	} else {
		return v88
	}
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
	var v18 int32
	_ = v18
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
	var v89 int32
	_ = v89
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
	var v235 int64
	_ = v235
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v282 int64
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int64
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int64
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int64
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int64
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int64
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int64
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int64
	_ = v360
	var v361 int32
	_ = v361
	var v362 int64
	_ = v362
	var v363 int32
	_ = v363
	var v364 int64
	_ = v364
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v383 int64
	_ = v383
	var v388 int64
	_ = v388
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int64
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v435 int64
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int64
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int64
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int64
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int64
	_ = v473
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int64
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v494 int64
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int64
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int64
	_ = v513
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v516 int32
	_ = v516
	var v517 int64
	_ = v517
	var v518 int32
	_ = v518
	var v519 int64
	_ = v519
	var v520 int32
	_ = v520
	var v521 int64
	_ = v521
	var v525 int64
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v536 int64
	_ = v536
	var v545 int64
	_ = v545
	var v557 int64
	_ = v557
	v9 = int64(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(1) {
		if v10&int32(1) != 0 {
			return int32(0)
		} else {
			v18 = int32(3)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v18 < l2 {
				v235 = int64(0)
				if base.B2i32(v20 != (v20+int32(3))&int32(-4))|base.B2i32(l2 < int32(4)) != 0 {
					v314 = v20
					v315 = l2
					v320 = v235
				} else {
					v245 = l2 - int32(4)
					v249 = int32(base.Ui32(v245)>>(uint(int32(2))%32)) + int32(1)
					v251 = v249 & int32(3)
					if base.Ui32(int32(12)) <= base.Ui32(v245) {
						v256 = v20
						v257 = l2
						v260 = int32(0)
						v262 = v235
						for {
							v263 = int32(16)
							v264 = v257 - v263
							v266 = v256 + v263
							v267 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
							v270 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
							v273 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
							v276 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
							v282 = base.I64_extend_i32_u(base.I32_popcnt(v267)) + (base.I64_extend_i32_u(base.I32_popcnt(v270)) + (base.I64_extend_i32_u(base.I32_popcnt(v273)) + (v262 + base.I64_extend_i32_u(base.I32_popcnt(v276)))))
							v284 = v260 + int32(4)
							if v284 != v249&int32(2147483644) {
								v256 = v266
								v257 = v264
								v260 = v284
								v262 = v282
								continue
							} else {
								break
							}
							break
						}
						if v251 == int32(0) {
							v314 = v266
							v315 = v264
							v320 = v282
						} else {
							v288 = v266
							v289 = v264
							v294 = v282
							v296 = v288
							v297 = v289
							v298 = int32(0)
							v302 = v294
							for {
								v303 = int32(4)
								v304 = v297 - v303
								v306 = v296 + v303
								v307 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
								v310 = v302 + base.I64_extend_i32_u(base.I32_popcnt(v307))
								v312 = v298 + int32(1)
								if v312 != v251 {
									v296 = v306
									v297 = v304
									v298 = v312
									v302 = v310
									continue
								} else {
									break
								}
								break
							}
							v314 = v306
							v315 = v304
							v320 = v310
						}
					} else {
						v288 = v20
						v289 = l2
						v294 = v235
						v296 = v288
						v297 = v289
						v298 = int32(0)
						v302 = v294
						for {
							v303 = int32(4)
							v304 = v297 - v303
							v306 = v296 + v303
							v307 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
							v310 = v302 + base.I64_extend_i32_u(base.I32_popcnt(v307))
							v312 = v298 + int32(1)
							if v312 != v251 {
								v296 = v306
								v297 = v304
								v298 = v312
								v302 = v310
								continue
							} else {
								break
							}
							break
						}
						v314 = v306
						v315 = v304
						v320 = v310
					}
				}
				if v315 == int32(0) {
					v383 = v320
				} else {
					v324 = v315 & int32(3)
					if v324 == int32(0) {
						v345 = v314
						v347 = v315
						v351 = v320
					} else {
						v328 = v314
						v330 = v315
						v332 = int32(0)
						v334 = v320
						for {
							v335 = int32(1)
							v336 = v328 + v335
							v338 = v330 - v335
							v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
							v340 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v339)+uint32(_c_F_hemdistcache_1[0]))))
							v341 = v334 + v340
							v343 = v332 + v335
							if v343 != v324 {
								v328 = v336
								v330 = v338
								v332 = v343
								v334 = v341
								continue
							} else {
								break
							}
							break
						}
						v345 = v336
						v347 = v338
						v351 = v341
					}
					if base.Ui32(v315) < base.Ui32(int32(4)) {
						v383 = v351
					} else {
						v354 = v345
						v356 = v347
						v360 = v351
						for {
							v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+3)))
							v362 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v361)+uint32(_c_F_hemdistcache_1[0]))))
							v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+2)))
							v364 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_hemdistcache_1[0]))))
							v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+1)))
							v366 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v365)+uint32(_c_F_hemdistcache_1[0]))))
							v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
							v368 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v367)+uint32(_c_F_hemdistcache_1[0]))))
							v372 = v362 + (v364 + (v366 + (v360 + v368)))
							v373 = int32(4)
							v376 = v356 - v373
							if v376 != 0 {
								v354 = v354 + v373
								v356 = v376
								v360 = v372
								continue
							} else {
								break
							}
							break
						}
						v383 = v372
					}
				}
				v557 = v383
			} else {
				if l2 == int32(0) {
					v557 = v9
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
							v557 = v54
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
							v557 = v81
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
						v557 = v81
					}
				}
			}
			return l2<<(uint(v18)%32) - base.I32_wrap_i64(v557)
		}
	} else {
		if v10&int32(1) != 0 {
			v89 = int32(3)
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v89 < l2 {
				v388 = int64(0)
				if base.B2i32(v91 != (v91+int32(3))&int32(-4))|base.B2i32(l2 < int32(4)) != 0 {
					v467 = v91
					v468 = l2
					v473 = v388
				} else {
					v398 = l2 - int32(4)
					v402 = int32(base.Ui32(v398)>>(uint(int32(2))%32)) + int32(1)
					v404 = v402 & int32(3)
					if base.Ui32(int32(12)) <= base.Ui32(v398) {
						v409 = v91
						v410 = l2
						v413 = int32(0)
						v415 = v388
						for {
							v416 = int32(16)
							v417 = v410 - v416
							v419 = v409 + v416
							v420 = *(*int32)(unsafe.Add(mBase, uint32(v409)+12))
							v423 = *(*int32)(unsafe.Add(mBase, uint32(v409)+8))
							v426 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
							v429 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
							v435 = base.I64_extend_i32_u(base.I32_popcnt(v420)) + (base.I64_extend_i32_u(base.I32_popcnt(v423)) + (base.I64_extend_i32_u(base.I32_popcnt(v426)) + (v415 + base.I64_extend_i32_u(base.I32_popcnt(v429)))))
							v437 = v413 + int32(4)
							if v437 != v402&int32(2147483644) {
								v409 = v419
								v410 = v417
								v413 = v437
								v415 = v435
								continue
							} else {
								break
							}
							break
						}
						if v404 == int32(0) {
							v467 = v419
							v468 = v417
							v473 = v435
						} else {
							v441 = v419
							v442 = v417
							v447 = v435
							v449 = v441
							v450 = v442
							v451 = int32(0)
							v455 = v447
							for {
								v456 = int32(4)
								v457 = v450 - v456
								v459 = v449 + v456
								v460 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
								v463 = v455 + base.I64_extend_i32_u(base.I32_popcnt(v460))
								v465 = v451 + int32(1)
								if v465 != v404 {
									v449 = v459
									v450 = v457
									v451 = v465
									v455 = v463
									continue
								} else {
									break
								}
								break
							}
							v467 = v459
							v468 = v457
							v473 = v463
						}
					} else {
						v441 = v91
						v442 = l2
						v447 = v388
						v449 = v441
						v450 = v442
						v451 = int32(0)
						v455 = v447
						for {
							v456 = int32(4)
							v457 = v450 - v456
							v459 = v449 + v456
							v460 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
							v463 = v455 + base.I64_extend_i32_u(base.I32_popcnt(v460))
							v465 = v451 + int32(1)
							if v465 != v404 {
								v449 = v459
								v450 = v457
								v451 = v465
								v455 = v463
								continue
							} else {
								break
							}
							break
						}
						v467 = v459
						v468 = v457
						v473 = v463
					}
				}
				if v468 == int32(0) {
					v536 = v473
				} else {
					v477 = v468 & int32(3)
					if v477 == int32(0) {
						v498 = v467
						v500 = v468
						v504 = v473
					} else {
						v481 = v467
						v483 = v468
						v485 = int32(0)
						v487 = v473
						for {
							v488 = int32(1)
							v489 = v481 + v488
							v491 = v483 - v488
							v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
							v493 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v492)+uint32(_c_F_hemdistcache_1[0]))))
							v494 = v487 + v493
							v496 = v485 + v488
							if v496 != v477 {
								v481 = v489
								v483 = v491
								v485 = v496
								v487 = v494
								continue
							} else {
								break
							}
							break
						}
						v498 = v489
						v500 = v491
						v504 = v494
					}
					if base.Ui32(v468) < base.Ui32(int32(4)) {
						v536 = v504
					} else {
						v507 = v498
						v509 = v500
						v513 = v504
						for {
							v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+3)))
							v515 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v514)+uint32(_c_F_hemdistcache_1[0]))))
							v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+2)))
							v517 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v516)+uint32(_c_F_hemdistcache_1[0]))))
							v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+1)))
							v519 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v518)+uint32(_c_F_hemdistcache_1[0]))))
							v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
							v521 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v520)+uint32(_c_F_hemdistcache_1[0]))))
							v525 = v515 + (v517 + (v519 + (v513 + v521)))
							v526 = int32(4)
							v529 = v509 - v526
							if v529 != 0 {
								v507 = v507 + v526
								v509 = v529
								v513 = v525
								continue
							} else {
								break
							}
							break
						}
						v536 = v525
					}
				}
				v545 = v536
			} else {
				if l2 == int32(0) {
					v545 = v9
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
							v545 = v125
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
							v545 = v152
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
						v545 = v152
					}
				}
			}
			return l2<<(uint(v89)%32) - base.I32_wrap_i64(v545)
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 float64
	_ = v426
	var v427 float64
	_ = v427
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(_a_F_hnswgettuple_0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[0]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	*(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[0])) = v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
	if v23 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L10
	} else {
		goto L103
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	goto L49
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+268)))
	if v30 != int32(1) {
		v46 = v26
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v39 = v27
	v40 = v26
	goto L8
L8:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = v41 + int64(1)
	v46 = v40
	goto L5
L9:
	;
	F_pgstat_assoc_relation(m, v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
	v39 = v38
	v40 = v37
	goto L8
L12:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = v48 + int64(1)
	goto L14
L13:
	;
	goto L14
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v52 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	switch v56 {
	case 0, 5:
		goto L16
	default:
		goto L17
	}
L16:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v70&int32(1) != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_errmsg_internal(m, int32(_a_F_hnswgettuple_1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_hnswgettuple_2), int32(219), int32(_a_F_hnswgettuple_3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	F_LockPage(m, v84, int32(1), int32(5))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L27
	}
L22:
	;
	v84 = v46
	v85 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v52)+44))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+60))
	if v75 == int32(0) {
		v84 = v46
		v85 = v73
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
	v80 = F_HnswNormValue(m, v78, v79, v73)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v84 = v82
	v85 = v80
	goto L21
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HnswGetMetaPageInfo(m, v91, v15+int32(12), v15+int32(8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = v85
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v101 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v102 = int32(0)
	v104 = v90 + int32(20)
	v106 = v90 + int32(56)
	v108 = F_HnswEntryCandidate(m, v102, v101, v104, v91, v106, v102)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L32
	}
L30:
	;
	v189 = int32(0)
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_UnlockPage(m, v191, int32(1), int32(5))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L45
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v108
	v113 = F_list_make1_impl(m, int32(1), v15)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+65)))
	if v116 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v118 = v116
	v120 = v113
	goto L37
L35:
	;
	v147 = v113
	goto L36
L36:
	;
	v156 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[1]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[2]))
	if v169 != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v129 = int32(1)
	v131 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v140 = F_HnswSearchLayer(m, v131, v104, v120, v129, v118, v91, v106, v133, v131, v131, v131, v131, v129, v131)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	v147 = v140
	goto L36
L39:
	;
	if base.Ui32(v129) < base.Ui32(v118) {
		v118 = v118 - int32(1)
		v120 = v140
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v170 = v90 + int32(16)
	goto L43
L42:
	;
	v170 = v156
	goto L43
L43:
	;
	v174 = F_HnswSearchLayer(m, v156, v104, v147, v158, v156, v91, v106, v160, v156, v156, v90+int32(12), v170, int32(1), v90+int32(32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v189 = v174
	goto L31
L45:
	;
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)) = uint8(v196)
	goto L4
L46:
	;
	m.G0 = v15 + int32(16)
	return v476
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[0])) = v18
	v476 = int32(0)
	goto L46
L48:
	;
	v437 = v404 + v420&int32(255)*int32(6) + int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[0])) = v18
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v440)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v442
	v444 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v444)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v444)
	v476 = int32(1)
	goto L46
L49:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v222 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v20)+40)) = v426
	goto L48
L51:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v397+v388<<(uint(int32(2))%32)-int32(4))))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+24))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+64)))
	if v405 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L52:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v223 != 0 {
		v386 = v222
		v388 = v223
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[2]))
	if v226 == int32(0) {
		goto L47
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v229 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L57:
	;
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v20)+32))
	v234 = int64(*(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[3])))
	if v232 < v234 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if v369 == int32(0) {
		goto L47
	} else {
		goto L92
	}
L59:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_LockPage(m, v279, int32(1), int32(5))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L10
	} else {
		goto L78
	}
L60:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236)+8))
	goto L64
L61:
	;
	v267 = v229
	goto L62
L62:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	if v268 == int32(0) {
		goto L47
	} else {
		goto L75
	}
L63:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if base.Ui32(v240) <= base.Ui32(v264) {
		goto L59
	} else {
		goto L74
	}
L64:
	;
	goto L63
L74:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v267 = v266
	goto L62
L75:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v272 = F_pairingheap_remove_first(m, v267)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	v276 = F_lappend(m, v271, v272-int32(12))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v276
	v369 = v276
	goto L58
L78:
	;
	v284 = int32(0)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	if v288 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v290 = v286 + int32(16)
	v291 = int32(0)
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[1]))
	if v293 <= v291 {
		v322 = v284
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v351 = v284
	v356 = v285
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v351
	F_UnlockPage(m, v356, int32(1), int32(5))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L10
	} else {
		goto L91
	}
L82:
	;
	v333 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v286)+24))
	v347 = F_HnswSearchLayer(m, v333, v286+int32(20), v322, v293, v333, v285, v286+int32(56), v339, v333, v333, v286+int32(12), v290, v333, v286+int32(32))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L10
	} else {
		goto L90
	}
L83:
	;
	v297 = v284
	v299 = v291
	goto L84
L84:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	if v309 == int32(0) {
		v322 = v297
		goto L82
	} else {
		goto L86
	}
L85:
	;
	v322 = v316
	goto L82
L86:
	;
	v312 = F_pairingheap_remove_first(m, v308)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	v316 = F_lappend(m, v297, v312-int32(12))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	v319 = v299 + int32(1)
	if v319 != v293 {
		v297 = v316
		v299 = v319
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v351 = v347
	v356 = v349
	goto L81
L91:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v369 = v367
	goto L58
L92:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v382 == int32(0) {
		goto L47
	} else {
		goto L93
	}
L93:
	;
	v386 = v369
	v388 = v382
	goto L51
L94:
	;
	v408 = F_list_delete_last(m, v386)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L10
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v420 = v405 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v404)+64)) = uint8(v420)
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[2]))
	if v423 != int32(2) {
		goto L48
	} else {
		goto L101
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v408
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_hnswgettuple[2]))
	if v412 == int32(0) {
		goto L49
	} else {
		goto L98
	}
L98:
	;
	F_pfree(m, v404)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	F_pfree(m, v403)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	goto L49
L101:
	;
	v426 = *(*float64)(unsafe.Add(mBase, uint32(v403)+32))
	v427 = *(*float64)(unsafe.Add(mBase, uint32(v20)+40))
	if base.F64_lt(v426, v427) != 0 {
		goto L49
	} else {
		goto L102
	}
L102:
	;
	goto L50
L103:
	;
	F_errmsg_internal(m, int32(_a_F_hnswgettuple_4), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_hnswgettuple_2), int32(214), int32(_a_F_hnswgettuple_3))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
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
	*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = int64(-4503599627370496)
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = v9
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
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
				v24 = v20 * int32(48)
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
				v36 = v32 * int32(48)
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v334 int32
	_ = v334
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v414 int32
	_ = v414
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v531 int32
	_ = v531
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L6
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v591
	v595 = v591 - int32(1)
	if v595 <= v8 {
		goto L136
	} else {
		goto L137
	}
L2:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v584))) = v580
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L85
L4:
	;
	if v130 != 0 {
		goto L3
	} else {
		goto L28
	}
L5:
	;
	v130 = v123
	goto L4
L6:
	;
	if v25 <= v12 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v123 = int32(0)
	goto L5
L8:
	;
	v130 = int32(-1)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v41 = int32(1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v26))))
	if base.Ui32(v43) < base.Ui32(int32(192)) {
		v100 = v43
		v101 = v41
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if int32(369) < v100 {
		v123 = v101
		goto L5
	} else {
		goto L24
	}
L12:
	;
	v47 = v12 + int32(1)
	if v47 == v25 {
		v100 = v43
		v101 = v41
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v26))))
	v52 = v50 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v43) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v26))))
	v68 = v66 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v43) {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v56 = v12 + int32(2)
	if v56 != v25 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v100 = v43<<(uint(int32(6))%32)&int32(1984) | v52
	v101 = int32(2)
	goto L11
L18:
	;
	goto L17
L19:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v72))))
	v100 = v85&int32(63) | (v43<<(uint(int32(18))%32)&int32(_a_F_hungarian_UTF_8_stem_0) | v52<<(uint(int32(12))%32) | v68<<(uint(int32(6))%32))
	v101 = int32(4)
	goto L11
L20:
	;
	v72 = v12 + int32(3)
	if v72 != v25 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v100 = v43<<(uint(int32(12))%32)&int32(_a_F_hungarian_UTF_8_stem_1) | v52<<(uint(int32(6))%32) | v68
	v101 = int32(3)
	goto L11
L23:
	;
	goto L22
L24:
	;
	v105 = v100 - int32(97)
	if v105 < int32(0) {
		v123 = v101
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_UTF_8_stem[0]))))
	if int32(base.Ui32(v111)>>(uint(v105&int32(7))%32))&int32(1) == int32(0) {
		v123 = v101
		goto L5
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v101 + v12
	goto L27
L27:
	;
	goto L7
L28:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v152 = v142
	goto L31
L29:
	;
	if v248 < int32(0) {
		goto L3
	} else {
		goto L53
	}
L30:
	;
	v248 = v219
	goto L29
L31:
	;
	if v143 <= v152 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v248 = int32(-1)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v159 = int32(1)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v144))))
	if base.Ui32(v161) < base.Ui32(int32(192)) {
		v218 = v161
		v219 = v159
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if int32(369) < v218 {
		goto L30
	} else {
		goto L49
	}
L37:
	;
	v165 = v152 + int32(1)
	if v165 == v143 {
		v218 = v161
		v219 = v159
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v144))))
	v170 = v168 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v161) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v144))))
	v186 = v184 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v161) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v174 = v152 + int32(2)
	if v174 != v143 {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v218 = v161<<(uint(int32(6))%32)&int32(1984) | v170
	v219 = int32(2)
	goto L36
L43:
	;
	goto L42
L44:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v190))))
	v218 = v203&int32(63) | (v161<<(uint(int32(18))%32)&int32(_a_F_hungarian_UTF_8_stem_0) | v170<<(uint(int32(12))%32) | v186<<(uint(int32(6))%32))
	v219 = int32(4)
	goto L36
L45:
	;
	v190 = v152 + int32(3)
	if v190 != v143 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v218 = v161<<(uint(int32(12))%32)&int32(_a_F_hungarian_UTF_8_stem_1) | v170<<(uint(int32(6))%32) | v186
	v219 = int32(3)
	goto L36
L48:
	;
	goto L47
L49:
	;
	v223 = v218 - int32(97)
	if v223 < int32(0) {
		goto L30
	} else {
		goto L50
	}
L50:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v223)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_UTF_8_stem[0]))))
	if int32(base.Ui32(v229)>>(uint(v223&int32(7))%32))&int32(1) == int32(0) {
		goto L30
	} else {
		goto L51
	}
L51:
	;
	v237 = v219 + v152
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v237
	v152 = v237
	goto L31
L53:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v254 = v252 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v255 <= v254 {
		v279 = v251
		v281 = v255
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
	goto L64
L55:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+v254))))
	if base.B2i32(v258&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v258)%32)&int32(101187584) == int32(0)) != 0 {
		v279 = v251
		v281 = v255
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v272 = F_find_among(m, l0, int32(_a_F_hungarian_UTF_8_stem_2), int32(8))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	if v272 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v580 = v276
	goto L2
L60:
	;
	goto L61
L61:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v279 = v278
	v281 = v277
	goto L54
L62:
	;
	if int32(0) <= v334 {
		v580 = v334
		goto L2
	} else {
		goto L82
	}
L64:
	;
	goto L65
L65:
	;
	goto L66
L66:
	;
	v289 = v252
	v291 = int32(1)
	goto L69
L68:
	;
	v334 = v319
	goto L62
L69:
	;
	if v281 <= v289 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L68
L71:
	;
	v334 = int32(-1)
	goto L62
L72:
	;
	goto L73
L73:
	;
	v296 = v289 + int32(1)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v289))))
	if base.Ui32(v298) < base.Ui32(int32(192)) {
		v319 = v296
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v320 = int32(1)
	if v320 < v291 {
		v289 = v319
		v291 = v291 - v320
		goto L69
	} else {
		goto L81
	}
L75:
	;
	if v281 <= v296 {
		v319 = v296
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v305 = v296
	goto L77
L77:
	;
	v308 = int32(*(*int8)(unsafe.Add(mBase, uint32(v279+v305))))
	if int32(-65) < v308 {
		v319 = v305
		goto L74
	} else {
		goto L79
	}
L78:
	;
	v319 = v281
	goto L74
L79:
	;
	v312 = v305 + int32(1)
	if v312 != v281 {
		v305 = v312
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	goto L70
L82:
	;
	goto L3
L83:
	;
	if v458 != 0 {
		goto L1
	} else {
		goto L108
	}
L84:
	;
	v458 = v451
	goto L83
L85:
	;
	if v354 <= v12 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v451 = int32(0)
	goto L84
L87:
	;
	v458 = int32(-1)
	goto L83
L88:
	;
	goto L89
L89:
	;
	v370 = int32(1)
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v355))))
	if base.Ui32(v372) < base.Ui32(int32(192)) {
		v429 = v372
		v430 = v370
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if int32(369) < v429 {
		goto L103
	} else {
		goto L104
	}
L91:
	;
	v376 = v12 + int32(1)
	if v376 == v354 {
		v429 = v372
		v430 = v370
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+v355))))
	v381 = v379 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v372) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+v355))))
	v397 = v395 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v372) {
		goto L99
	} else {
		goto L100
	}
L94:
	;
	v385 = v12 + int32(2)
	if v385 != v354 {
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v429 = v372<<(uint(int32(6))%32)&int32(1984) | v381
	v430 = int32(2)
	goto L90
L97:
	;
	goto L96
L98:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v401))))
	v429 = v414&int32(63) | (v372<<(uint(int32(18))%32)&int32(_a_F_hungarian_UTF_8_stem_0) | v381<<(uint(int32(12))%32) | v397<<(uint(int32(6))%32))
	v430 = int32(4)
	goto L90
L99:
	;
	v401 = v12 + int32(3)
	if v401 != v354 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v429 = v372<<(uint(int32(12))%32)&int32(_a_F_hungarian_UTF_8_stem_1) | v381<<(uint(int32(6))%32) | v397
	v430 = int32(3)
	goto L90
L102:
	;
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v430 + v12
	goto L107
L104:
	;
	v434 = v429 - int32(97)
	if v434 < int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v434)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_UTF_8_stem[0]))))
	if int32(base.Ui32(v440)>>(uint(v434&int32(7))%32))&int32(1) != 0 {
		v451 = v430
		goto L84
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	goto L86
L108:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v480 = v470
	goto L111
L109:
	;
	if v575 < int32(0) {
		goto L1
	} else {
		goto L134
	}
L110:
	;
	v575 = v547
	goto L109
L111:
	;
	if v471 <= v480 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v575 = int32(-1)
	goto L109
L114:
	;
	goto L115
L115:
	;
	v487 = int32(1)
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480+v472))))
	if base.Ui32(v489) < base.Ui32(int32(192)) {
		v546 = v489
		v547 = v487
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if int32(369) < v546 {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	v493 = v480 + int32(1)
	if v493 == v471 {
		v546 = v489
		v547 = v487
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493+v472))))
	v498 = v496 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v489) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502+v472))))
	v514 = v512 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v489) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v502 = v480 + int32(2)
	if v502 != v471 {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v546 = v489<<(uint(int32(6))%32)&int32(1984) | v498
	v547 = int32(2)
	goto L116
L123:
	;
	goto L122
L124:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v518))))
	v546 = v531&int32(63) | (v489<<(uint(int32(18))%32)&int32(_a_F_hungarian_UTF_8_stem_0) | v498<<(uint(int32(12))%32) | v514<<(uint(int32(6))%32))
	v547 = int32(4)
	goto L116
L125:
	;
	v518 = v480 + int32(3)
	if v518 != v471 {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v546 = v489<<(uint(int32(12))%32)&int32(_a_F_hungarian_UTF_8_stem_1) | v498<<(uint(int32(6))%32) | v514
	v547 = int32(3)
	goto L116
L128:
	;
	goto L127
L129:
	;
	v564 = v547 + v480
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v564
	v480 = v564
	goto L111
L130:
	;
	v551 = v546 - int32(97)
	if v551 < int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v551)>>(uint(int32(3))%32)))+uint32(_c_F_hungarian_UTF_8_stem[0]))))
	if int32(base.Ui32(v557)>>(uint(v551&int32(7))%32))&int32(1) != 0 {
		goto L110
	} else {
		goto L132
	}
L132:
	;
	goto L129
L134:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v580 = v578 + v575
	goto L2
L135:
	;
	return v1064
L136:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v652
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v652
	v657 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_3), int32(44))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L57
	} else {
		goto L151
	}
L137:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597+v595))))
	if v599 != int32(108) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v604 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_4), int32(2))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L57
	} else {
		goto L139
	}
L139:
	;
	if v604 == int32(0) {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v608
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	if v608 < v611 {
		goto L136
	} else {
		goto L141
	}
L141:
	;
	v614 = v608 - int32(1)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v614 <= v615 {
		goto L136
	} else {
		goto L142
	}
L142:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617+v614))))
	if base.B2i32(v619&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v619)%32)&int32(106790108) == int32(0)) != 0 {
		goto L136
	} else {
		goto L143
	}
L143:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v634 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_5), int32(23))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L57
	} else {
		goto L144
	}
L144:
	;
	if v634 == int32(0) {
		goto L136
	} else {
		goto L145
	}
L145:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v638 + (v608 - v631)
	v642 = F_slice_del(m, l0)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L57
	} else {
		goto L146
	}
L146:
	;
	if v642 < int32(0) {
		v1064 = v642
		goto L135
	} else {
		goto L147
	}
L147:
	;
	v646 = F_r_undouble_4(m, l0)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L57
	} else {
		goto L148
	}
L148:
	;
	if v646 < int32(0) {
		v1064 = v646
		goto L135
	} else {
		goto L149
	}
L149:
	;
	goto L136
L150:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v708
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v708-int32(2) <= v711 {
		goto L167
	} else {
		goto L168
	}
L151:
	;
	if v657 == int32(0) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v661
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	if v661 < v664 {
		goto L150
	} else {
		goto L153
	}
L153:
	;
	v666 = F_slice_del(m, l0)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L57
	} else {
		goto L154
	}
L154:
	;
	if v666 < int32(0) {
		v1064 = v666
		goto L135
	} else {
		goto L155
	}
L155:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v670
	v673 = v670 - int32(1)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v673 <= v674 {
		goto L150
	} else {
		goto L156
	}
L156:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676+v673))))
	switch v678 - int32(161) {
	case 0, 8:
		goto L157
	default:
		goto L150
	}
L157:
	;
	v683 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_6), int32(2))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L57
	} else {
		goto L158
	}
L158:
	;
	if v683 == int32(0) {
		goto L150
	} else {
		goto L159
	}
L159:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v687
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)))
	if v687 < v690 {
		goto L150
	} else {
		goto L160
	}
L160:
	;
	switch v683 - int32(1) {
	case 0:
		goto L162
	case 1:
		goto L161
	default:
		goto L150
	}
L161:
	;
	v702 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_7))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L57
	} else {
		goto L165
	}
L162:
	;
	v696 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_8))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L57
	} else {
		goto L163
	}
L163:
	;
	if int32(0) <= v696 {
		goto L150
	} else {
		goto L164
	}
L164:
	;
	v1064 = v696
	goto L135
L165:
	;
	if v702 < int32(0) {
		v1064 = v702
		goto L135
	} else {
		goto L166
	}
L166:
	;
	goto L150
L167:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v749
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v749
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v749-int32(3) <= v752 {
		goto L179
	} else {
		goto L180
	}
L168:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715+v708-int32(1)))))
	switch v719 - int32(110) {
	case 0, 6:
		goto L169
	default:
		goto L167
	}
L169:
	;
	v724 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_9), int32(3))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L57
	} else {
		goto L170
	}
L170:
	;
	if v724 == int32(0) {
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v728
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	if v728 < v731 {
		goto L167
	} else {
		goto L172
	}
L172:
	;
	switch v724 - int32(1) {
	case 0:
		goto L174
	case 1:
		goto L173
	default:
		goto L167
	}
L173:
	;
	v743 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_10))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L57
	} else {
		goto L177
	}
L174:
	;
	v737 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_11))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L57
	} else {
		goto L175
	}
L175:
	;
	if int32(0) <= v737 {
		goto L167
	} else {
		goto L176
	}
L176:
	;
	v1064 = v737
	goto L135
L177:
	;
	if v743 < int32(0) {
		v1064 = v743
		goto L135
	} else {
		goto L178
	}
L178:
	;
	goto L167
L179:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v794
	v796 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v794
	v800 = v794 - int32(1)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v800 <= v801 {
		v854 = v796
		goto L194
	} else {
		goto L195
	}
L180:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756+v749-int32(1)))))
	if v760 != int32(108) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v765 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_12), int32(6))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L57
	} else {
		goto L182
	}
L182:
	;
	if v765 == int32(0) {
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v769
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v771)))
	if v769 < v772 {
		goto L179
	} else {
		goto L184
	}
L184:
	;
	switch v765 - int32(1) {
	case 0:
		goto L187
	case 1:
		goto L186
	case 2:
		goto L185
	default:
		goto L179
	}
L185:
	;
	v788 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_13))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L57
	} else {
		goto L192
	}
L186:
	;
	v782 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_14))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L57
	} else {
		goto L190
	}
L187:
	;
	v776 = F_slice_del(m, l0)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L57
	} else {
		goto L188
	}
L188:
	;
	if int32(0) <= v776 {
		goto L179
	} else {
		goto L189
	}
L189:
	;
	v1064 = v776
	goto L135
L190:
	;
	if int32(0) <= v782 {
		goto L179
	} else {
		goto L191
	}
L191:
	;
	v1064 = v782
	goto L135
L192:
	;
	if v788 < int32(0) {
		v1064 = v788
		goto L135
	} else {
		goto L193
	}
L193:
	;
	goto L179
L194:
	;
	if v854 < int32(0) {
		v1064 = v854
		goto L135
	} else {
		goto L207
	}
L195:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803+v800))))
	switch v805 - int32(161) {
	case 0, 8:
		goto L196
	default:
		v854 = v796
		goto L194
	}
L196:
	;
	v810 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_15), int32(2))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L57
	} else {
		goto L197
	}
L197:
	;
	if v810 == int32(0) {
		v854 = v796
		goto L194
	} else {
		goto L198
	}
L198:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v814
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	if v814 < v817 {
		v854 = v796
		goto L194
	} else {
		goto L199
	}
L199:
	;
	v820 = v814 - int32(1)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v820 <= v821 {
		v854 = v796
		goto L194
	} else {
		goto L200
	}
L200:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823+v820))))
	if base.B2i32(v825&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v825)%32)&int32(106790108) == int32(0)) != 0 {
		v854 = v796
		goto L194
	} else {
		goto L201
	}
L201:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v840 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_5), int32(23))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L57
	} else {
		goto L202
	}
L202:
	;
	if v840 == int32(0) {
		v854 = v796
		goto L194
	} else {
		goto L203
	}
L203:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v844 + (v814 - v837)
	v848 = F_slice_del(m, l0)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L57
	} else {
		goto L204
	}
L204:
	;
	if v848 < int32(0) {
		v854 = v848
		goto L194
	} else {
		goto L205
	}
L205:
	;
	v852 = F_r_undouble_4(m, l0)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L57
	} else {
		goto L206
	}
L206:
	;
	v854 = v852
	goto L194
L207:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v859
	v861 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v859
	v865 = v859 - int32(1)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v865 <= v866 {
		v907 = v861
		goto L208
	} else {
		goto L209
	}
L208:
	;
	if v907 < int32(0) {
		v1064 = v907
		goto L135
	} else {
		goto L224
	}
L209:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868+v865))))
	if base.B2i32(v870 != int32(169))&base.B2i32(v870 != int32(105)) != 0 {
		v907 = v861
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v878 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_16), int32(12))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L57
	} else {
		goto L211
	}
L211:
	;
	if v878 == int32(0) {
		v907 = v861
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v882
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v884)))
	if v882 < v885 {
		v907 = v861
		goto L208
	} else {
		goto L213
	}
L213:
	;
	switch v878 - int32(1) {
	case 0:
		goto L217
	case 1:
		goto L216
	case 2:
		goto L215
	default:
		goto L214
	}
L214:
	;
	v907 = int32(1)
	goto L208
L215:
	;
	v901 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_17))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L57
	} else {
		goto L222
	}
L216:
	;
	v895 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_18))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L57
	} else {
		goto L220
	}
L217:
	;
	v889 = F_slice_del(m, l0)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L57
	} else {
		goto L218
	}
L218:
	;
	if int32(0) <= v889 {
		goto L214
	} else {
		goto L219
	}
L219:
	;
	v907 = v889
	goto L208
L220:
	;
	if int32(0) <= v895 {
		goto L214
	} else {
		goto L221
	}
L221:
	;
	v907 = v895
	goto L208
L222:
	;
	if v901 < int32(0) {
		v907 = v901
		goto L208
	} else {
		goto L223
	}
L223:
	;
	goto L214
L224:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v912
	v914 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v912
	v919 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_19), int32(31))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L57
	} else {
		goto L226
	}
L225:
	;
	if v948 < int32(0) {
		v1064 = v948
		goto L135
	} else {
		goto L239
	}
L226:
	;
	if v919 == int32(0) {
		v948 = v914
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v923
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v925)))
	if v923 < v926 {
		v948 = v914
		goto L225
	} else {
		goto L228
	}
L228:
	;
	switch v919 - int32(1) {
	case 0:
		goto L232
	case 1:
		goto L231
	case 2:
		goto L230
	default:
		goto L229
	}
L229:
	;
	v948 = int32(1)
	goto L225
L230:
	;
	v942 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_20))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L57
	} else {
		goto L237
	}
L231:
	;
	v936 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_21))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L57
	} else {
		goto L235
	}
L232:
	;
	v930 = F_slice_del(m, l0)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L57
	} else {
		goto L233
	}
L233:
	;
	if int32(0) <= v930 {
		goto L229
	} else {
		goto L234
	}
L234:
	;
	v948 = v930
	goto L225
L235:
	;
	if int32(0) <= v936 {
		goto L229
	} else {
		goto L236
	}
L236:
	;
	v948 = v936
	goto L225
L237:
	;
	if v942 < int32(0) {
		v948 = v942
		goto L225
	} else {
		goto L238
	}
L238:
	;
	goto L229
L239:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v952
	v954 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v952
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v952 <= v957 {
		v1006 = v954
		goto L240
	} else {
		goto L241
	}
L240:
	;
	if v1006 < int32(0) {
		v1064 = v1006
		goto L135
	} else {
		goto L256
	}
L241:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v961 = int32(1)
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959+v952-v961))))
	if base.B2i32(v963&int32(224) != int32(96))|base.B2i32(v961<<(uint(v963)%32)&int32(_a_F_hungarian_UTF_8_stem_22) == int32(0)) != 0 {
		v1006 = v954
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v977 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_23), int32(42))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L57
	} else {
		goto L243
	}
L243:
	;
	if v977 == int32(0) {
		v1006 = v954
		goto L240
	} else {
		goto L244
	}
L244:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v981
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v983)))
	if v981 < v984 {
		v1006 = v954
		goto L240
	} else {
		goto L245
	}
L245:
	;
	switch v977 - int32(1) {
	case 0:
		goto L249
	case 1:
		goto L248
	case 2:
		goto L247
	default:
		goto L246
	}
L246:
	;
	v1006 = int32(1)
	goto L240
L247:
	;
	v1000 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_24))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L57
	} else {
		goto L254
	}
L248:
	;
	v994 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_25))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L57
	} else {
		goto L252
	}
L249:
	;
	v988 = F_slice_del(m, l0)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L57
	} else {
		goto L250
	}
L250:
	;
	if int32(0) <= v988 {
		goto L246
	} else {
		goto L251
	}
L251:
	;
	v1006 = v988
	goto L240
L252:
	;
	if int32(0) <= v994 {
		goto L246
	} else {
		goto L253
	}
L253:
	;
	v1006 = v994
	goto L240
L254:
	;
	if v1000 < int32(0) {
		v1006 = v1000
		goto L240
	} else {
		goto L255
	}
L255:
	;
	goto L246
L256:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1011
	v1013 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1011
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1011 <= v1016 {
		v1056 = v1013
		goto L257
	} else {
		goto L258
	}
L257:
	;
	if v1056 < int32(0) {
		v1064 = v1056
		goto L135
	} else {
		goto L273
	}
L258:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1018+v1011-int32(1)))))
	if v1022 != int32(107) {
		v1056 = v1013
		goto L257
	} else {
		goto L259
	}
L259:
	;
	v1027 = F_find_among_b(m, l0, int32(_a_F_hungarian_UTF_8_stem_26), int32(7))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L57
	} else {
		goto L260
	}
L260:
	;
	if v1027 == int32(0) {
		v1056 = v1013
		goto L257
	} else {
		goto L261
	}
L261:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1031
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1033)))
	if v1031 < v1034 {
		v1056 = v1013
		goto L257
	} else {
		goto L262
	}
L262:
	;
	switch v1027 - int32(1) {
	case 0:
		goto L266
	case 1:
		goto L265
	case 2:
		goto L264
	default:
		goto L263
	}
L263:
	;
	v1056 = int32(1)
	goto L257
L264:
	;
	v1050 = F_slice_del(m, l0)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L57
	} else {
		goto L271
	}
L265:
	;
	v1046 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_27))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L57
	} else {
		goto L269
	}
L266:
	;
	v1040 = F_slice_from_s(m, l0, int32(1), int32(_a_F_hungarian_UTF_8_stem_28))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L57
	} else {
		goto L267
	}
L267:
	;
	if int32(0) <= v1040 {
		goto L263
	} else {
		goto L268
	}
L268:
	;
	v1056 = v1040
	goto L257
L269:
	;
	if int32(0) <= v1046 {
		goto L263
	} else {
		goto L270
	}
L270:
	;
	v1056 = v1046
	goto L257
L271:
	;
	if v1050 < int32(0) {
		v1056 = v1050
		goto L257
	} else {
		goto L272
	}
L272:
	;
	goto L263
L273:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1061
	v1064 = int32(1)
	goto L135
}
func F_hypothetical_dense_rank_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
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
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
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
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var __phi250 int32
	_ = __phi250
	var v251 int32
	_ = v251
	var __phi251 int32
	_ = __phi251
	var v259 int32
	_ = v259
	var __phi259 int32
	_ = __phi259
	var v261 int64
	_ = v261
	var __phi261 int64
	_ = __phi261
	var v262 int64
	_ = v262
	var __phi262 int64
	_ = __phi262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v282 int32
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
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v312 int64
	_ = v312
	var v313 int32
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
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	v2 = int32(0)
	v14 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v2
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v23 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L66
	}
L2:
	;
	m.G0 = v18 + int32(16)
	return v370
L3:
	;
	v27 = F_Int64GetDatum(m, int64(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v32 = l0 + int32(20)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v35 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	v370 = v27
	goto L2
L8:
	;
	v38 = int32(_a_F_hypothetical_dense_rank_final_0)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v41
	v43 = F_CreateStandaloneExprContext(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v53 = v35
	goto L10
L10:
	;
	v54 = int32(1)
	v55 = v20 - v54
	if v55&v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v43
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v39
	v53 = v48
	goto L10
L12:
	;
	v59 = v55 >> (uint(int32(1)) % 32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	F_hypothetical_check_argtypes(m, l0, v59, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	if v65 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v91 = v65
	v92 = v64
	goto L16
L15:
	;
	v66 = int32(_a_F_hypothetical_dense_rank_final_0)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+36))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+40))
	v80 = F_execTuplesMatchPrepare(m, v74, v69-int32(1), v68, v77, v78, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L17
	}
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	m.T0[v95].(func(*base.Module, int32))(m, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v67
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+48)) = v80
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v91 = v80
	v92 = v86
	goto L16
L18:
	;
	v98 = int32(0)
	if v59 <= v98 {
		v201 = v98
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v202+v201<<(uint(int32(2))%32)))) = int32(-1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v201))) = uint8(v210)
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
	v214 = v212 & int32(_a_F_hypothetical_dense_rank_final_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)) = uint16(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+6)) = uint16(v217)
	goto L28
L20:
	;
	v101 = int32(0)
	if v55 != int32(2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v110 = v101
	v119 = v2
	goto L24
L22:
	;
	v160 = v101
	goto L23
L23:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v179 = v32 + v160<<(uint(int32(3))%32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v173+v160<<(uint(int32(2))%32)))) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v182+v160))) = uint8(v184)
	v201 = v59
	goto L19
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v124 = int32(2)
	v128 = v110 | int32(1)
	v129 = int32(3)
	v131 = v32 + v128<<(uint(v129)%32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v123+v110<<(uint(v124)%32)))) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v134+v110))) = uint8(v136)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v143 = v110 + v124
	v146 = v32 + v143<<(uint(v129)%32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v128<<(uint(v124)%32)))) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v128+v149))) = uint8(v151)
	v154 = v119 + v124
	if v154 != v59&int32(2147483646) {
		v110 = v143
		v119 = v154
		goto L24
	} else {
		goto L26
	}
L25:
	;
	if v59&int32(1) == int32(0) {
		v201 = v59
		goto L19
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v160 = v143
	goto L23
L28:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	F_tuplesort_puttupleslot(m, v219, v93)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	F_tuplesort_performsort(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v225 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+24)) = uint8(v225)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	v230 = F_MakeTupleTableSlot(m, v228, int32(_a_F_hypothetical_dense_rank_final_2))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v233 = int32(1)
	v237 = F_tuplesort_gettupleslot(m, v232, v233, v233, v93, v18+int32(8))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	m.T0[v343].(func(*base.Module, int32))(m, v327)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L62
	}
L33:
	;
	if v237 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v327 = v93
	v328 = v230
	v340 = int64(1)
	v341 = v14
	goto L32
L35:
	;
	goto L36
L36:
	;
	__phi250 = v93
	__phi251 = v230
	__phi259 = int32(0)
	__phi261 = int64(1)
	__phi262 = v14
	v250 = __phi250
	v251 = __phi251
	v259 = __phi259
	v261 = __phi261
	v262 = __phi262
	goto L37
L37:
	;
	v263 = int32(*(*int16)(unsafe.Add(mBase, uint32(v250)+6)))
	if v263 <= v59 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v327 = v251
	v328 = v250
	v340 = v319
	v341 = v312
	goto L32
L39:
	;
	F_slot_getsomeattrs_int(m, v250, v59+int32(1))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v250)+20))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v59))))
	if v269 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v250
	if v251 == int32(0) {
		v312 = v262
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v270+v59<<(uint(int32(2))%32))))
	if v272 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v327 = v250
	v328 = v251
	v340 = v261
	v341 = v262
	goto L32
L46:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[1]))
	if v315 != 0 {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+4)))
	if v279&int32(2) != 0 {
		v312 = v262
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v282 != v259 {
		v312 = v262
		goto L46
	} else {
		goto L49
	}
L49:
	;
	if v91 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	F_MemoryContextReset(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L6
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v291 = int32(_a_F_hypothetical_dense_rank_final_0)
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v294
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v299 = m.T0[v298].(func(*base.Module, int32, int32, int32) int32)(m, v91, v53, v18+int32(15))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L54
	}
L53:
	;
	v312 = v262 + int64(1)
	goto L46
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hypothetical_dense_rank_final[0])) = v292
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	F_MemoryContextReset(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v312 = v262 + base.I64_extend_i32_u(base.B2i32(v299 != int32(0)))
	goto L46
L56:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v319 = v261 + int64(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v321 = int32(1)
	v325 = F_tuplesort_gettupleslot(m, v320, v321, v321, v251, v18+int32(8))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L6
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	if v325 != 0 {
		__phi250 = v251
		__phi251 = v250
		__phi259 = v313
		__phi261 = v319
		__phi262 = v312
		v250 = __phi250
		v251 = __phi251
		v259 = __phi259
		v261 = __phi261
		v262 = __phi262
		goto L37
	} else {
		goto L61
	}
L61:
	;
	goto L38
L62:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	m.T0[v347].(func(*base.Module, int32))(m, v328)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	F_ExecDropSingleTupleTableSlot(m, v230)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v353 = F_Int64GetDatum(m, v340-v341)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v370 = v353
	goto L2
L66:
	;
	F_errmsg_internal(m, int32(_a_F_hypothetical_dense_rank_final_3), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_hypothetical_dense_rank_final_4), int32(1332), int32(_a_F_hypothetical_dense_rank_final_5))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
