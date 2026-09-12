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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1238]))
	F_hash_seq_init(m, v5+int32(12), v10)
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
	v15 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L23
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v17 = v15
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+84)))
	if v20 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v34 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L17
	}
L12:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+85)))
	if v23 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v24 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
	if v25 != int32(2) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	F_HoldPortal(m, v19)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+85)) = uint8(v30)
	goto L11
L17:
	;
	if v34 != 0 {
		v17 = v34
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
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(19372), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(488237), int32(1232), int32(150836))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
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
	F_errmsg_internal(m, int32(424407), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(488237), int32(1236), int32(150836))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
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
	var v25 int32
	_ = v25
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v13 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v111
L2:
	;
	v111 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 <= int32(0) {
		v111 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v25 = v3
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v25<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v34 = int32(0)
	if v27 == v34 {
		v75 = v34
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v111 = v3
	goto L1
L8:
	;
	v103 = v25 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v103 < v104 {
		v25 = v103
		goto L6
	} else {
		goto L31
	}
L9:
	;
	if v75 != 0 {
		goto L8
	} else {
		goto L23
	}
L10:
	;
	goto L9
L11:
	;
	if v33 == int32(0) {
		v75 = v34
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v43 < v44 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v43
	goto L15
L14:
	;
	v46 = v44
	goto L15
L15:
	;
	if v46 <= int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = int32(1)
	goto L18
L17:
	;
	v49 = v46
	goto L18
L18:
	;
	v50 = int32(8)
	v55 = int32(0)
	goto L19
L19:
	;
	v62 = v55 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33+v50+v62)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+(v27+v50))))
	v67 = v64 & v66
	v69 = base.B2i32(v67 != int32(0))
	if v67 != 0 {
		v75 = v69
		goto L10
	} else {
		goto L21
	}
L20:
	;
	v75 = v69
	goto L10
L21:
	;
	v71 = v55 + int32(1)
	if v71 != v49 {
		v55 = v71
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v79 = F_have_relevant_joinclause(m, l0, l1, v32)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	if v79 == int32(0) {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v87 = F_bms_union(m, v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v93 = F_join_is_legal(m, l0, l1, v32, v87, v11+int32(12), v11+int32(11))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	F_bms_free(m, v87)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	if v93 == int32(0) {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	v111 = int32(1)
	goto L1
L31:
	;
	goto L7
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
			v13 = F_convert_any_priv_string(m, v10, int32(1632848))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[276]))
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v4 = int32(1)
	if l0 == l1 {
		v54 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v54
L2:
	;
	v6 = F_superuser_arg(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v6 != 0 {
		v54 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v11 = int32(0)
	v13 = F_roles_is_member_of(m, l0, int32(1), v11, v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v15 = int32(0)
	if v13 == v15 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v54 = v53
	goto L1
L8:
	;
	v53 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21 <= int32(0) {
		v46 = v15
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v46
	goto L7
L12:
	;
	v24 = int32(0)
	if v24 < v21 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v27 = v21
	goto L15
L14:
	;
	v27 = v24
	goto L15
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v30 = int32(0)
	goto L16
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30<<(uint(int32(2))%32))))
	v39 = base.B2i32(v38 == l1)
	if v38 == l1 {
		v46 = v39
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v46 = v39
	goto L11
L18:
	;
	v41 = v30 + int32(1)
	if v41 != v27 {
		v30 = v41
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
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
	var v43 float64
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 float64
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v89 float64
	_ = v89
	var v91 int32
	_ = v91
	var v98 float64
	_ = v98
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v201 float64
	_ = v201
	var v203 float64
	_ = v203
	var v205 float64
	_ = v205
	var v211 float64
	_ = v211
	var v228 float64
	_ = v228
	var v232 float64
	_ = v232
	var v251 float64
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
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
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_pfree(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L44
	} else {
		goto L50
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
	v43 = float64(0)
	v45 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v280 = v20
	goto L11
L11:
	;
	v285 = v21 + int32(1)
	if v285 < v280 {
		v20 = v280
		v21 = v285
		goto L7
	} else {
		goto L49
	}
L12:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v252+v40)+16))
	F_pfree(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L44
	} else {
		goto L45
	}
L13:
	;
	v251 = v232
	goto L12
L14:
	;
	if base.F64_gt(v211, float64(1.4316557653333333e+08)) == int32(0) {
		v232 = v211
		goto L13
	} else {
		goto L43
	}
L15:
	;
	v53 = int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v52 == v53 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v203 = *(*float64)(unsafe.Add(mBase, uint32(v42)+8))
	v205 = base.F64_div(v203, float64(0))
	if base.F64_le(v205, base.F64_mul(base.F64_convert_i32_u(v52), float64(2.5))) != 0 {
		v232 = v205
		goto L13
	} else {
		goto L42
	}
L18:
	;
	if v52&v53 != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v89 = v43
	v91 = v45
	goto L18
L20:
	;
	goto L21
L21:
	;
	v61 = v43
	v63 = v45
	v65 = v45
	goto L22
L22:
	;
	v70 = float64(1)
	v71 = v63 + v55
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v73 = F_ldexp(m, v70, v72)
	mBase = m.M
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v77 = F_ldexp(m, v70, v76)
	mBase = m.M
	v82 = base.F64_add(base.F64_add(v61, base.F64_div(v70, v77)), base.F64_div(v70, v73))
	v83 = int32(2)
	v84 = v63 + v83
	v86 = v65 + v83
	if v86 != v52&int32(-2) {
		v61 = v82
		v63 = v84
		v65 = v86
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v89 = v82
	v91 = v84
	goto L18
L24:
	;
	goto L23
L25:
	;
	v98 = float64(1)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v55))))
	v102 = F_ldexp(m, v98, v101)
	mBase = m.M
	v105 = base.F64_add(v89, base.F64_div(v98, v102))
	goto L27
L26:
	;
	v105 = v89
	goto L27
L27:
	;
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v42)+8))
	v107 = base.F64_div(v106, v105)
	v108 = base.F64_convert_i32_u(v52)
	if base.F64_le(v107, base.F64_mul(v108, float64(2.5))) == int32(0) {
		v211 = v107
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v115 = v52 & int32(3)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v117 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v52) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v128 = v117
	v129 = int32(0)
	v130 = v117
	goto L32
L30:
	;
	v160 = v117
	v162 = v117
	goto L31
L31:
	;
	if v115 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v135 = v128 + v116
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v137 = int32(0)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+3)))
	v151 = v130 + base.B2i32(v136 == v137) + base.B2i32(v140 == v137) + base.B2i32(v144 == v137) + base.B2i32(v148 == v137)
	v152 = int32(4)
	v153 = v128 + v152
	v155 = v129 + v152
	if v155 != v52&int32(-4) {
		v128 = v153
		v129 = v155
		v130 = v151
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v160 = v153
	v162 = v151
	goto L31
L34:
	;
	goto L33
L35:
	;
	v170 = v160
	v172 = v162
	v173 = v117
	goto L38
L36:
	;
	v192 = v162
	goto L37
L37:
	;
	if v192 == int32(0) {
		v232 = v107
		goto L13
	} else {
		goto L41
	}
L38:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v116))))
	v181 = v172 + base.B2i32(v178 == int32(0))
	v182 = int32(1)
	v185 = v173 + v182
	if v185 != v115 {
		v170 = v170 + v182
		v172 = v181
		v173 = v185
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v192 = v181
	goto L37
L40:
	;
	goto L39
L41:
	;
	v201 = F_log(m, base.F64_div(v108, base.F64_convert_i32_s(v192)))
	mBase = m.M
	v251 = base.F64_mul(v201, v108)
	goto L12
L42:
	;
	v211 = v205
	goto L14
L43:
	;
	v228 = F_log(m, base.F64_add(base.F64_mul(v211, float64(-2.3283064365386963e-10)), float64(1)))
	mBase = m.M
	v232 = base.F64_mul(v228, float64(-4.294967296e+09))
	goto L13
L44:
	;
	return
L45:
	;
	F_LogicalTapeRewindForRead(m, v38, int32(8192))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v260+v28)))
	v264 = F_palloc0(m, int32(32))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v264)+24)) = v251
	*(*int64)(unsafe.Add(mBase, uint32(v264)+16)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = int32(32) - v15
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = l2
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v272 = F_lappend(m, v271, v264)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v272
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+336)) = v275 + int32(1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v280 = v279
	goto L11
L49:
	;
	goto L8
L50:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_pfree(m, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_pfree(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L44
	} else {
		goto L52
	}
L52:
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
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
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
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
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
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
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
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
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
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
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
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
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
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
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L130
	}
L4:
	;
	v15 = int32(1)
	v16 = v10 + v15
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v21 = v19 & v15
	if v21 != 0 {
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
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L125
	}
L7:
	;
	v22 = v15
	goto L9
L8:
	;
	v22 = int32(4)
	goto L9
L9:
	;
	if v19 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v51 = v10 + v22
	if v21 != 0 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v25 = int32(4)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v27&int32(254) == int32(2) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v40 = int32(1)
	if v21 != 0 {
		v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v36 = v25
	goto L16
L15:
	;
	v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
	goto L16
L16:
	;
	if v27 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v39 = v25
	goto L19
L18:
	;
	v39 = v36
	goto L19
L19:
	;
	v50 = v39
	goto L10
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	v54 = v16
	goto L23
L22:
	;
	v54 = v10 + int32(4)
	goto L23
L23:
	;
	v59 = v50
	goto L24
L24:
	;
	if v59 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v76 = F_pg_newlocale_from_collation(m, v14)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v75 = v50 >> (uint(int32(31)) % 32) & v50
	goto L26
L28:
	;
	goto L29
L29:
	;
	v69 = v59 - int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v69))))
	if v71 == int32(32) {
		v59 = v69
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v75 = v59
	goto L26
L31:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v626 != v10 {
		goto L121
	} else {
		goto L122
	}
L32:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v78 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v86 = v75 - int32(1636608432)
	if v51&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	goto L35
L35:
	;
	v345 = int32(0)
	v347 = F_pg_strnxfrm(m, v345, v345, v51, v75, v76)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L76
	}
L36:
	;
	v622 = v340 ^ v332 - base.I32_rotl(v340, int32(24))
	goto L31
L37:
	;
	v318 = int32(14)
	v320 = v314 ^ v315 - base.I32_rotl(v314, v318)
	v324 = v320 ^ v313 - base.I32_rotl(v320, int32(11))
	v328 = v324 ^ v314 - base.I32_rotl(v324, int32(25))
	v332 = v328 ^ v320 - base.I32_rotl(v328, int32(16))
	v336 = v332 ^ v324 - base.I32_rotl(v332, int32(4))
	v340 = v336 ^ v328 - base.I32_rotl(v336, v318)
	goto L36
L38:
	;
	switch v244 - int32(1) {
	case 0:
		v306 = v245
		v307 = v246
		v308 = v247
		goto L65
	case 1:
		v299 = v245
		v300 = v246
		v301 = v247
		goto L66
	case 2:
		v292 = v245
		v293 = v246
		v294 = v247
		goto L67
	case 3:
		v286 = v246
		v287 = v247
		goto L68
	case 4:
		v282 = v246
		v283 = v247
		goto L69
	case 5:
		v276 = v246
		v277 = v247
		goto L70
	case 6:
		v270 = v246
		v271 = v247
		goto L71
	case 7:
		v265 = v247
		goto L72
	case 8:
		v260 = v247
		goto L73
	case 9:
		v255 = v247
		goto L74
	case 10:
		goto L75
	default:
		v313 = v245
		v314 = v246
		v315 = v247
		goto L37
	}
L39:
	;
	v195 = v51
	v196 = v75
	v197 = v86
	v198 = v86
	v199 = v86
	goto L62
L40:
	;
	if base.Ui32(int32(11)) < base.Ui32(v75) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(v75) < base.Ui32(int32(12)) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v243 = v51
	v244 = v75
	v245 = v86
	v246 = v86
	v247 = v86
	goto L38
L44:
	;
	switch v142 - int32(1) {
	case 0:
		v192 = v143
		goto L51
	case 1:
		v187 = v143
		goto L52
	case 2:
		goto L53
	case 3:
		v180 = v144
		goto L54
	case 4:
		v177 = v144
		goto L55
	case 5:
		v172 = v144
		goto L56
	case 6:
		goto L57
	case 7:
		v163 = v145
		goto L58
	case 8:
		v158 = v145
		goto L59
	case 9:
		v153 = v145
		goto L60
	case 10:
		goto L61
	default:
		v313 = v143
		v314 = v144
		v315 = v145
		goto L37
	}
L45:
	;
	v141 = v51
	v142 = v75
	v143 = v86
	v144 = v86
	v145 = v86
	goto L44
L46:
	;
	goto L47
L47:
	;
	v93 = v51
	v94 = v75
	v95 = v86
	v96 = v86
	v97 = v86
	goto L48
L48:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v100 = v99 + v96
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v104 = v103 + v97
	v106 = int32(4)
	v108 = v101 + v95 - v104 ^ base.I32_rotl(v104, v106)
	v112 = v100 - v108 ^ base.I32_rotl(v108, int32(6))
	v113 = v104 + v100
	v114 = v108 + v113
	v115 = v112 + v114
	v119 = v113 - v112 ^ base.I32_rotl(v112, int32(8))
	v123 = v114 - v119 ^ base.I32_rotl(v119, int32(16))
	v127 = v115 - v123 ^ base.I32_rotl(v123, int32(19))
	v128 = v119 + v115
	v129 = v123 + v128
	v130 = v127 + v129
	v134 = v128 - v127 ^ base.I32_rotl(v127, v106)
	v135 = int32(12)
	v136 = v93 + v135
	v138 = v94 - v135
	if base.Ui32(int32(11)) < base.Ui32(v138) {
		v93 = v136
		v94 = v138
		v95 = v129
		v96 = v130
		v97 = v134
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v141 = v136
	v142 = v138
	v143 = v129
	v144 = v130
	v145 = v134
	goto L44
L50:
	;
	goto L49
L51:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v313 = v192 + v193
	v314 = v144
	v315 = v145
	goto L37
L52:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	v192 = v188<<(uint(int32(8))%32) + v187
	goto L51
L53:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+2)))
	v187 = v183<<(uint(int32(16))%32) + v143
	goto L52
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v313 = v181 + v143
	v314 = v180
	v315 = v145
	goto L37
L55:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+4)))
	v180 = v177 + v178
	goto L54
L56:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+5)))
	v177 = v173<<(uint(int32(8))%32) + v172
	goto L55
L57:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+6)))
	v172 = v168<<(uint(int32(16))%32) + v144
	goto L56
L58:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v313 = v164 + v143
	v314 = v166 + v144
	v315 = v163
	goto L37
L59:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+8)))
	v163 = v159<<(uint(int32(8))%32) + v158
	goto L58
L60:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+9)))
	v158 = v154<<(uint(int32(16))%32) + v153
	goto L59
L61:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+10)))
	v153 = v149<<(uint(int32(24))%32) + v145
	goto L60
L62:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v202 = v201 + v198
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	v206 = v205 + v199
	v208 = int32(4)
	v210 = v203 + v197 - v206 ^ base.I32_rotl(v206, v208)
	v214 = v202 - v210 ^ base.I32_rotl(v210, int32(6))
	v215 = v206 + v202
	v216 = v210 + v215
	v217 = v214 + v216
	v221 = v215 - v214 ^ base.I32_rotl(v214, int32(8))
	v225 = v216 - v221 ^ base.I32_rotl(v221, int32(16))
	v229 = v217 - v225 ^ base.I32_rotl(v225, int32(19))
	v230 = v221 + v217
	v231 = v225 + v230
	v232 = v229 + v231
	v236 = v230 - v229 ^ base.I32_rotl(v229, v208)
	v237 = int32(12)
	v238 = v195 + v237
	v240 = v196 - v237
	if base.Ui32(int32(11)) < base.Ui32(v240) {
		v195 = v238
		v196 = v240
		v197 = v231
		v198 = v232
		v199 = v236
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v243 = v238
	v244 = v240
	v245 = v231
	v246 = v232
	v247 = v236
	goto L38
L64:
	;
	goto L63
L65:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v313 = v306 + v309
	v314 = v307
	v315 = v308
	goto L37
L66:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	v306 = v302<<(uint(int32(8))%32) + v299
	v307 = v300
	v308 = v301
	goto L65
L67:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+2)))
	v299 = v295<<(uint(int32(16))%32) + v292
	v300 = v293
	v301 = v294
	goto L66
L68:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+3)))
	v292 = v288<<(uint(int32(24))%32) + v245
	v293 = v286
	v294 = v287
	goto L67
L69:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+4)))
	v286 = v282 + v284
	v287 = v283
	goto L68
L70:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+5)))
	v282 = v278<<(uint(int32(8))%32) + v276
	v283 = v277
	goto L69
L71:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+6)))
	v276 = v272<<(uint(int32(16))%32) + v270
	v277 = v271
	goto L70
L72:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+7)))
	v270 = v266<<(uint(int32(24))%32) + v246
	v271 = v265
	goto L71
L73:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+8)))
	v265 = v261<<(uint(int32(8))%32) + v260
	goto L72
L74:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+9)))
	v260 = v256<<(uint(int32(16))%32) + v255
	goto L73
L75:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+10)))
	v255 = v251<<(uint(int32(24))%32) + v247
	goto L74
L76:
	;
	v350 = v347 + int32(1)
	v351 = F_palloc(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v353 = F_pg_strnxfrm(m, v351, v350, v51, v75, v76)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(v347) < base.Ui32(v353) {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v361 = v350 - int32(1636608432)
	if v351&int32(3) != 0 {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	F_pfree(m, v351)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L120
	}
L81:
	;
	v593 = int32(14)
	v595 = v589 ^ v590 - base.I32_rotl(v589, v593)
	v599 = v595 ^ v588 - base.I32_rotl(v595, int32(11))
	v603 = v599 ^ v589 - base.I32_rotl(v599, int32(25))
	v607 = v603 ^ v595 - base.I32_rotl(v603, int32(16))
	v611 = v607 ^ v599 - base.I32_rotl(v607, int32(4))
	v615 = v611 ^ v603 - base.I32_rotl(v611, v593)
	goto L80
L82:
	;
	switch v519 - int32(1) {
	case 0:
		v581 = v520
		v582 = v521
		v583 = v522
		goto L109
	case 1:
		v574 = v520
		v575 = v521
		v576 = v522
		goto L110
	case 2:
		v567 = v520
		v568 = v521
		v569 = v522
		goto L111
	case 3:
		v561 = v521
		v562 = v522
		goto L112
	case 4:
		v557 = v521
		v558 = v522
		goto L113
	case 5:
		v551 = v521
		v552 = v522
		goto L114
	case 6:
		v545 = v521
		v546 = v522
		goto L115
	case 7:
		v540 = v522
		goto L116
	case 8:
		v535 = v522
		goto L117
	case 9:
		v530 = v522
		goto L118
	case 10:
		goto L119
	default:
		v588 = v520
		v589 = v521
		v590 = v522
		goto L81
	}
L83:
	;
	v470 = v351
	v471 = v350
	v472 = v361
	v473 = v361
	v474 = v361
	goto L106
L84:
	;
	if base.Ui32(int32(11)) < base.Ui32(v350) {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(v350) < base.Ui32(int32(12)) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v518 = v351
	v519 = v350
	v520 = v361
	v521 = v361
	v522 = v361
	goto L82
L88:
	;
	switch v417 - int32(1) {
	case 0:
		v467 = v418
		goto L95
	case 1:
		v462 = v418
		goto L96
	case 2:
		goto L97
	case 3:
		v455 = v419
		goto L98
	case 4:
		v452 = v419
		goto L99
	case 5:
		v447 = v419
		goto L100
	case 6:
		goto L101
	case 7:
		v438 = v420
		goto L102
	case 8:
		v433 = v420
		goto L103
	case 9:
		v428 = v420
		goto L104
	case 10:
		goto L105
	default:
		v588 = v418
		v589 = v419
		v590 = v420
		goto L81
	}
L89:
	;
	v416 = v351
	v417 = v350
	v418 = v361
	v419 = v361
	v420 = v361
	goto L88
L90:
	;
	goto L91
L91:
	;
	v368 = v351
	v369 = v350
	v370 = v361
	v371 = v361
	v372 = v361
	goto L92
L92:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	v375 = v374 + v371
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v379 = v378 + v372
	v381 = int32(4)
	v383 = v376 + v370 - v379 ^ base.I32_rotl(v379, v381)
	v387 = v375 - v383 ^ base.I32_rotl(v383, int32(6))
	v388 = v379 + v375
	v389 = v383 + v388
	v390 = v387 + v389
	v394 = v388 - v387 ^ base.I32_rotl(v387, int32(8))
	v398 = v389 - v394 ^ base.I32_rotl(v394, int32(16))
	v402 = v390 - v398 ^ base.I32_rotl(v398, int32(19))
	v403 = v394 + v390
	v404 = v398 + v403
	v405 = v402 + v404
	v409 = v403 - v402 ^ base.I32_rotl(v402, v381)
	v410 = int32(12)
	v411 = v368 + v410
	v413 = v369 - v410
	if base.Ui32(int32(11)) < base.Ui32(v413) {
		v368 = v411
		v369 = v413
		v370 = v404
		v371 = v405
		v372 = v409
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v416 = v411
	v417 = v413
	v418 = v404
	v419 = v405
	v420 = v409
	goto L88
L94:
	;
	goto L93
L95:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	v588 = v467 + v468
	v589 = v419
	v590 = v420
	goto L81
L96:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+1)))
	v467 = v463<<(uint(int32(8))%32) + v462
	goto L95
L97:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+2)))
	v462 = v458<<(uint(int32(16))%32) + v418
	goto L96
L98:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v588 = v456 + v418
	v589 = v455
	v590 = v420
	goto L81
L99:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+4)))
	v455 = v452 + v453
	goto L98
L100:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+5)))
	v452 = v448<<(uint(int32(8))%32) + v447
	goto L99
L101:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+6)))
	v447 = v443<<(uint(int32(16))%32) + v419
	goto L100
L102:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v588 = v439 + v418
	v589 = v441 + v419
	v590 = v438
	goto L81
L103:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+8)))
	v438 = v434<<(uint(int32(8))%32) + v433
	goto L102
L104:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+9)))
	v433 = v429<<(uint(int32(16))%32) + v428
	goto L103
L105:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+10)))
	v428 = v424<<(uint(int32(24))%32) + v420
	goto L104
L106:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	v477 = v476 + v473
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v470)+8))
	v481 = v480 + v474
	v483 = int32(4)
	v485 = v478 + v472 - v481 ^ base.I32_rotl(v481, v483)
	v489 = v477 - v485 ^ base.I32_rotl(v485, int32(6))
	v490 = v481 + v477
	v491 = v485 + v490
	v492 = v489 + v491
	v496 = v490 - v489 ^ base.I32_rotl(v489, int32(8))
	v500 = v491 - v496 ^ base.I32_rotl(v496, int32(16))
	v504 = v492 - v500 ^ base.I32_rotl(v500, int32(19))
	v505 = v496 + v492
	v506 = v500 + v505
	v507 = v504 + v506
	v511 = v505 - v504 ^ base.I32_rotl(v504, v483)
	v512 = int32(12)
	v513 = v470 + v512
	v515 = v471 - v512
	if base.Ui32(int32(11)) < base.Ui32(v515) {
		v470 = v513
		v471 = v515
		v472 = v506
		v473 = v507
		v474 = v511
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v518 = v513
	v519 = v515
	v520 = v506
	v521 = v507
	v522 = v511
	goto L82
L108:
	;
	goto L107
L109:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	v588 = v581 + v584
	v589 = v582
	v590 = v583
	goto L81
L110:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+1)))
	v581 = v577<<(uint(int32(8))%32) + v574
	v582 = v575
	v583 = v576
	goto L109
L111:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+2)))
	v574 = v570<<(uint(int32(16))%32) + v567
	v575 = v568
	v576 = v569
	goto L110
L112:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+3)))
	v567 = v563<<(uint(int32(24))%32) + v520
	v568 = v561
	v569 = v562
	goto L111
L113:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+4)))
	v561 = v557 + v559
	v562 = v558
	goto L112
L114:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+5)))
	v557 = v553<<(uint(int32(8))%32) + v551
	v558 = v552
	goto L113
L115:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+6)))
	v551 = v547<<(uint(int32(16))%32) + v545
	v552 = v546
	goto L114
L116:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+7)))
	v545 = v541<<(uint(int32(24))%32) + v521
	v546 = v540
	goto L115
L117:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+8)))
	v540 = v536<<(uint(int32(8))%32) + v535
	goto L116
L118:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+9)))
	v535 = v531<<(uint(int32(16))%32) + v530
	goto L117
L119:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+10)))
	v530 = v526<<(uint(int32(24))%32) + v522
	goto L118
L120:
	;
	v622 = v615 ^ v607 - base.I32_rotl(v615, int32(24))
	goto L31
L121:
	;
	F_pfree(m, v10)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	return v622
L124:
	;
	goto L123
L125:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(329797), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errhint(m, int32(549953), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(486699), int32(1001), int32(226150))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
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
	F_errmsg_internal(m, int32(96536), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(486699), int32(1025), int32(226150))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		v12 = int32(-1636608428)
		v50 = v12
		v52 = v12
		v55 = v12
	} else {
		v15 = base.I32_wrap_i64(v4)
		v20 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v26 = v15 - v20 - int32(1636608428) ^ base.I32_rotl(v20, int32(6))
		v28 = v15 + int32(1021750440)
		v29 = v20 + v28
		v30 = v26 + v29
		v34 = v28 - v26 ^ base.I32_rotl(v26, int32(8))
		v38 = v29 - v34 ^ base.I32_rotl(v34, int32(16))
		v42 = v30 - v38 ^ base.I32_rotl(v38, int32(19))
		v43 = v34 + v30
		v44 = v38 + v43
		v50 = v42 + v44
		v52 = v44
		v55 = v43 - v42 ^ base.I32_rotl(v42, int32(4))
	}
	v57 = int32(14)
	v59 = v50 ^ v55 - base.I32_rotl(v50, v57)
	v64 = v59 ^ (v2 + v52) - base.I32_rotl(v59, int32(11))
	v68 = v64 ^ v50 - base.I32_rotl(v64, int32(25))
	v72 = v68 ^ v59 - base.I32_rotl(v68, int32(16))
	v76 = v72 ^ v64 - base.I32_rotl(v72, int32(4))
	v80 = v76 ^ v68 - base.I32_rotl(v76, v57)
	v90 = F_Int64GetDatum(m, base.I64_extend_i32_u(v80)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v80^v72-base.I32_rotl(v80, int32(24))))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		return int32(0)
	} else {
		return v90
	}
}
func F_hemdistcache_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int64
	_ = v145
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v161 int32
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
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
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v243 int64
	_ = v243
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int64
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int64
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int64
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int64
	_ = v328
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int64
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int64
	_ = v348
	var v349 int64
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v361 int64
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int64
	_ = v370
	var v371 int32
	_ = v371
	var v374 int64
	_ = v374
	var v375 int32
	_ = v375
	var v378 int64
	_ = v378
	var v379 int32
	_ = v379
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v386 int64
	_ = v386
	var v390 int64
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v401 int64
	_ = v401
	var v406 int64
	_ = v406
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int64
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int64
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int64
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int64
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int64
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int64
	_ = v491
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int64
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int64
	_ = v511
	var v512 int64
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int64
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int64
	_ = v533
	var v534 int32
	_ = v534
	var v537 int64
	_ = v537
	var v538 int32
	_ = v538
	var v541 int64
	_ = v541
	var v542 int32
	_ = v542
	var v545 int64
	_ = v545
	var v546 int32
	_ = v546
	var v549 int64
	_ = v549
	var v553 int64
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v564 int64
	_ = v564
	var v568 int64
	_ = v568
	var v579 int64
	_ = v579
	v4 = int64(0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(1) {
		if v9&int32(1) != 0 {
			return int32(0)
		} else {
			v17 = int32(3)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v17 < l2 {
				v243 = int64(0)
				if l2 < int32(4) {
					v322 = v19
					v323 = l2
					v328 = v243
				} else {
					if v19 != (v19+int32(3))&int32(-4) {
						v322 = v19
						v323 = l2
						v328 = v243
					} else {
						v252 = l2 - int32(4)
						v256 = int32(base.Ui32(v252)>>(uint(int32(2))%32)) + int32(1)
						v258 = v256 & int32(3)
						if base.Ui32(v252) < base.Ui32(int32(12)) {
							v294 = v19
							v295 = l2
							v300 = v243
						} else {
							v264 = v19
							v265 = l2
							v266 = int32(0)
							v270 = v243
							for {
								v271 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
								v274 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
								v277 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
								v280 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
								v286 = base.I64_extend_i32_u(base.I32_popcnt(v271)) + (base.I64_extend_i32_u(base.I32_popcnt(v274)) + (base.I64_extend_i32_u(base.I32_popcnt(v277)) + (v270 + base.I64_extend_i32_u(base.I32_popcnt(v280)))))
								v287 = int32(16)
								v288 = v265 - v287
								v290 = v264 + v287
								v292 = v266 + int32(4)
								if v292 != v256&int32(2147483644) {
									v264 = v290
									v265 = v288
									v266 = v292
									v270 = v286
									continue
								} else {
									break
								}
								break
							}
							v294 = v290
							v295 = v288
							v300 = v286
						}
						if v258 == int32(0) {
							v322 = v294
							v323 = v295
							v328 = v300
						} else {
							v305 = v295
							v306 = v294
							v307 = int32(0)
							v310 = v300
							for {
								v311 = int32(4)
								v312 = v305 - v311
								v313 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
								v316 = v310 + base.I64_extend_i32_u(base.I32_popcnt(v313))
								v318 = v306 + v311
								v320 = v307 + int32(1)
								if v320 != v258 {
									v305 = v312
									v306 = v318
									v307 = v320
									v310 = v316
									continue
								} else {
									break
								}
								break
							}
							v322 = v318
							v323 = v312
							v328 = v316
						}
					}
				}
				if v323 == int32(0) {
					v401 = v328
				} else {
					v332 = v323 & int32(3)
					if v332 == int32(0) {
						v355 = v322
						v357 = v323
						v361 = v328
					} else {
						v338 = v323
						v339 = v322
						v340 = int32(0)
						v342 = v328
						for {
							v343 = int32(1)
							v344 = v338 - v343
							v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
							v348 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v345)+uint32(_consts[117]))))
							v349 = v342 + v348
							v351 = v339 + v343
							v353 = v340 + v343
							if v353 != v332 {
								v338 = v344
								v339 = v351
								v340 = v353
								v342 = v349
								continue
							} else {
								break
							}
							break
						}
						v355 = v351
						v357 = v344
						v361 = v349
					}
					if base.Ui32(v323) < base.Ui32(int32(4)) {
						v401 = v361
					} else {
						v364 = v355
						v366 = v357
						v370 = v361
						for {
							v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+3)))
							v374 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v371)+uint32(_consts[117]))))
							v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+2)))
							v378 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v375)+uint32(_consts[117]))))
							v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)))
							v382 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v379)+uint32(_consts[117]))))
							v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
							v386 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v383)+uint32(_consts[117]))))
							v390 = v374 + (v378 + (v382 + (v370 + v386)))
							v391 = int32(4)
							v394 = v366 - v391
							if v394 != 0 {
								v364 = v364 + v391
								v366 = v394
								v370 = v390
								continue
							} else {
								break
							}
							break
						}
						v401 = v390
					}
				}
				v579 = v401
			} else {
				if l2 == int32(0) {
					v579 = v4
				} else {
					v25 = l2 & int32(3)
					if base.Ui32(l2) < base.Ui32(int32(4)) {
						v64 = v19
						v67 = v4
					} else {
						v31 = v19
						v33 = int32(0)
						v34 = v4
						for {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)))
							v42 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[117]))))
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+2)))
							v46 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[117]))))
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
							v50 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[117]))))
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
							v54 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v51)+uint32(_consts[117]))))
							v58 = v42 + (v46 + (v50 + (v34 + v54)))
							v59 = int32(4)
							v60 = v31 + v59
							v62 = v33 + v59
							if v62 != l2&int32(-4) {
								v31 = v60
								v33 = v62
								v34 = v58
								continue
							} else {
								break
							}
							break
						}
						v64 = v60
						v67 = v58
					}
					if v25 == int32(0) {
						v579 = v67
					} else {
						v75 = v64
						v77 = int32(0)
						v78 = v67
						for {
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
							v86 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[117]))))
							v87 = v78 + v86
							v88 = int32(1)
							v91 = v77 + v88
							if v91 != v25 {
								v75 = v75 + v88
								v77 = v91
								v78 = v87
								continue
							} else {
								break
							}
							break
						}
						v579 = v87
					}
				}
			}
			return l2<<(uint(v17)%32) - base.I32_wrap_i64(v579)
		}
	} else {
		if v9&int32(1) != 0 {
			v95 = int32(3)
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v95 < l2 {
				v406 = int64(0)
				if l2 < int32(4) {
					v485 = v97
					v486 = l2
					v491 = v406
				} else {
					if v97 != (v97+int32(3))&int32(-4) {
						v485 = v97
						v486 = l2
						v491 = v406
					} else {
						v415 = l2 - int32(4)
						v419 = int32(base.Ui32(v415)>>(uint(int32(2))%32)) + int32(1)
						v421 = v419 & int32(3)
						if base.Ui32(v415) < base.Ui32(int32(12)) {
							v457 = v97
							v458 = l2
							v463 = v406
						} else {
							v427 = v97
							v428 = l2
							v429 = int32(0)
							v433 = v406
							for {
								v434 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
								v437 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
								v440 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
								v443 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
								v449 = base.I64_extend_i32_u(base.I32_popcnt(v434)) + (base.I64_extend_i32_u(base.I32_popcnt(v437)) + (base.I64_extend_i32_u(base.I32_popcnt(v440)) + (v433 + base.I64_extend_i32_u(base.I32_popcnt(v443)))))
								v450 = int32(16)
								v451 = v428 - v450
								v453 = v427 + v450
								v455 = v429 + int32(4)
								if v455 != v419&int32(2147483644) {
									v427 = v453
									v428 = v451
									v429 = v455
									v433 = v449
									continue
								} else {
									break
								}
								break
							}
							v457 = v453
							v458 = v451
							v463 = v449
						}
						if v421 == int32(0) {
							v485 = v457
							v486 = v458
							v491 = v463
						} else {
							v468 = v458
							v469 = v457
							v470 = int32(0)
							v473 = v463
							for {
								v474 = int32(4)
								v475 = v468 - v474
								v476 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
								v479 = v473 + base.I64_extend_i32_u(base.I32_popcnt(v476))
								v481 = v469 + v474
								v483 = v470 + int32(1)
								if v483 != v421 {
									v468 = v475
									v469 = v481
									v470 = v483
									v473 = v479
									continue
								} else {
									break
								}
								break
							}
							v485 = v481
							v486 = v475
							v491 = v479
						}
					}
				}
				if v486 == int32(0) {
					v564 = v491
				} else {
					v495 = v486 & int32(3)
					if v495 == int32(0) {
						v518 = v485
						v520 = v486
						v524 = v491
					} else {
						v501 = v486
						v502 = v485
						v503 = int32(0)
						v505 = v491
						for {
							v506 = int32(1)
							v507 = v501 - v506
							v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
							v511 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[117]))))
							v512 = v505 + v511
							v514 = v502 + v506
							v516 = v503 + v506
							if v516 != v495 {
								v501 = v507
								v502 = v514
								v503 = v516
								v505 = v512
								continue
							} else {
								break
							}
							break
						}
						v518 = v514
						v520 = v507
						v524 = v512
					}
					if base.Ui32(v486) < base.Ui32(int32(4)) {
						v564 = v524
					} else {
						v527 = v518
						v529 = v520
						v533 = v524
						for {
							v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+3)))
							v537 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v534)+uint32(_consts[117]))))
							v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+2)))
							v541 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v538)+uint32(_consts[117]))))
							v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+1)))
							v545 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v542)+uint32(_consts[117]))))
							v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527))))
							v549 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v546)+uint32(_consts[117]))))
							v553 = v537 + (v541 + (v545 + (v533 + v549)))
							v554 = int32(4)
							v557 = v529 - v554
							if v557 != 0 {
								v527 = v527 + v554
								v529 = v557
								v533 = v553
								continue
							} else {
								break
							}
							break
						}
						v564 = v553
					}
				}
				v568 = v564
			} else {
				if l2 == int32(0) {
					v568 = v4
				} else {
					v103 = l2 & int32(3)
					if base.Ui32(l2) < base.Ui32(int32(4)) {
						v142 = v97
						v145 = v4
					} else {
						v109 = v97
						v111 = int32(0)
						v112 = v4
						for {
							v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+3)))
							v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[117]))))
							v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+2)))
							v124 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[117]))))
							v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
							v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v125)+uint32(_consts[117]))))
							v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
							v132 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[117]))))
							v136 = v120 + (v124 + (v128 + (v112 + v132)))
							v137 = int32(4)
							v138 = v109 + v137
							v140 = v111 + v137
							if v140 != l2&int32(-4) {
								v109 = v138
								v111 = v140
								v112 = v136
								continue
							} else {
								break
							}
							break
						}
						v142 = v138
						v145 = v136
					}
					if v103 == int32(0) {
						v568 = v145
					} else {
						v153 = v142
						v155 = int32(0)
						v156 = v145
						for {
							v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
							v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[117]))))
							v165 = v156 + v164
							v166 = int32(1)
							v169 = v155 + v166
							if v169 != v103 {
								v153 = v153 + v166
								v155 = v169
								v156 = v165
								continue
							} else {
								break
							}
							break
						}
						v568 = v165
					}
				}
			}
			return l2<<(uint(v95)%32) - base.I32_wrap_i64(v568)
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v177 = int32(1)
				if l2 == v177 {
					v181 = int32(0)
					v220 = v181
					v222 = v181
				} else {
					v185 = int32(0)
					v187 = v185
					v189 = v185
					v192 = int32(0)
					for {
						v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v175))))
						v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v176))))
						v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196^v198)+uint32(_consts[117]))))
						v205 = v187 | int32(1)
						v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v205))))
						v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v176))))
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207^v209)+uint32(_consts[117]))))
						v214 = v189 + v202 + v213
						v215 = int32(2)
						v216 = v187 + v215
						v218 = v192 + v215
						if v218 != l2&int32(2147483646) {
							v187 = v216
							v189 = v214
							v192 = v218
							continue
						} else {
							break
						}
						break
					}
					v220 = v216
					v222 = v214
				}
				if l2&v177 != 0 {
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220+v175))))
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220+v176))))
					v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229^v231)+uint32(_consts[117]))))
					v237 = v222 + v235
				} else {
					v237 = v222
				}
				return v237
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
	if v29 != int32(32767) {
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
	v96 = int32(32767)
	*(*uint16)(unsafe.Add(mBase, uint32(v28))) = uint16(v96)
	goto L9
L15:
	;
	v56 = int32(1)
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17+v46<<(uint(v56)%32)))))
	v62 = v32 + v59<<(uint(v56)%32)
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62))))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v63<<(uint(v56)%32)))))
	if v67 != (v29+int32(1))&int32(65535) {
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
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v333 int32
	_ = v333
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v413 int32
	_ = v413
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v530 int32
	_ = v530
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
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
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v590
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v590
	v594 = v590 - int32(1)
	if v594 <= v8 {
		goto L137
	} else {
		goto L138
	}
L2:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v579
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L86
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
	v100 = v85&int32(63) | (v43<<(uint(int32(18))%32)&int32(1835008) | v52<<(uint(int32(12))%32) | v68<<(uint(int32(6))%32))
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
	v100 = v43<<(uint(int32(12))%32)&int32(61440) | v52<<(uint(int32(6))%32) | v68
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
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_consts[1324]))))
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
	v218 = v203&int32(63) | (v161<<(uint(int32(18))%32)&int32(1835008) | v170<<(uint(int32(12))%32) | v186<<(uint(int32(6))%32))
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
	v218 = v161<<(uint(int32(12))%32)&int32(61440) | v170<<(uint(int32(6))%32) | v186
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
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v223)>>(uint(int32(3))%32)))+uint32(_consts[1324]))))
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
		v278 = v251
		v280 = v255
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
	goto L65
L55:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+v254))))
	if v258&int32(224) != int32(96) {
		v278 = v251
		v280 = v255
		goto L54
	} else {
		goto L56
	}
L56:
	;
	if int32(1)<<(uint(v258)%32)&int32(101187584) == int32(0) {
		v278 = v251
		v280 = v255
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v271 = F_find_among(m, l0, int32(4263264), int32(8))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	return int32(0)
L59:
	;
	if v271 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v579 = v275
	goto L2
L61:
	;
	goto L62
L62:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v278 = v277
	v280 = v276
	goto L54
L63:
	;
	if int32(0) <= v333 {
		v579 = v333
		goto L2
	} else {
		goto L83
	}
L65:
	;
	goto L66
L66:
	;
	goto L67
L67:
	;
	v288 = v252
	v290 = int32(1)
	goto L70
L69:
	;
	v333 = v318
	goto L63
L70:
	;
	if v280 <= v288 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L69
L72:
	;
	v333 = int32(-1)
	goto L63
L73:
	;
	goto L74
L74:
	;
	v295 = v288 + int32(1)
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278+v288))))
	if base.Ui32(v297) < base.Ui32(int32(192)) {
		v318 = v295
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v319 = int32(1)
	if v319 < v290 {
		v288 = v318
		v290 = v290 - v319
		goto L70
	} else {
		goto L82
	}
L76:
	;
	if v280 <= v295 {
		v318 = v295
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v304 = v295
	goto L78
L78:
	;
	v307 = int32(*(*int8)(unsafe.Add(mBase, uint32(v278+v304))))
	if int32(-65) < v307 {
		v318 = v304
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v318 = v280
	goto L75
L80:
	;
	v311 = v304 + int32(1)
	if v311 != v280 {
		v304 = v311
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L71
L83:
	;
	goto L3
L84:
	;
	if v457 != 0 {
		goto L1
	} else {
		goto L109
	}
L85:
	;
	v457 = v450
	goto L84
L86:
	;
	if v353 <= v12 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v450 = int32(0)
	goto L85
L88:
	;
	v457 = int32(-1)
	goto L84
L89:
	;
	goto L90
L90:
	;
	v369 = int32(1)
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v354))))
	if base.Ui32(v371) < base.Ui32(int32(192)) {
		v428 = v371
		v429 = v369
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if int32(369) < v428 {
		goto L104
	} else {
		goto L105
	}
L92:
	;
	v375 = v12 + int32(1)
	if v375 == v353 {
		v428 = v371
		v429 = v369
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375+v354))))
	v380 = v378 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v371) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384+v354))))
	v396 = v394 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v371) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v384 = v12 + int32(2)
	if v384 != v353 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v428 = v371<<(uint(int32(6))%32)&int32(1984) | v380
	v429 = int32(2)
	goto L91
L98:
	;
	goto L97
L99:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+v400))))
	v428 = v413&int32(63) | (v371<<(uint(int32(18))%32)&int32(1835008) | v380<<(uint(int32(12))%32) | v396<<(uint(int32(6))%32))
	v429 = int32(4)
	goto L91
L100:
	;
	v400 = v12 + int32(3)
	if v400 != v353 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v428 = v371<<(uint(int32(12))%32)&int32(61440) | v380<<(uint(int32(6))%32) | v396
	v429 = int32(3)
	goto L91
L103:
	;
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 + v12
	goto L108
L105:
	;
	v433 = v428 - int32(97)
	if v433 < int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v433)>>(uint(int32(3))%32)))+uint32(_consts[1324]))))
	if int32(base.Ui32(v439)>>(uint(v433&int32(7))%32))&int32(1) != 0 {
		v450 = v429
		goto L85
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	goto L87
L109:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v479 = v469
	goto L112
L110:
	;
	if v574 < int32(0) {
		goto L1
	} else {
		goto L135
	}
L111:
	;
	v574 = v546
	goto L110
L112:
	;
	if v470 <= v479 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v574 = int32(-1)
	goto L110
L115:
	;
	goto L116
L116:
	;
	v486 = int32(1)
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479+v471))))
	if base.Ui32(v488) < base.Ui32(int32(192)) {
		v545 = v488
		v546 = v486
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if int32(369) < v545 {
		goto L130
	} else {
		goto L131
	}
L118:
	;
	v492 = v479 + int32(1)
	if v492 == v470 {
		v545 = v488
		v546 = v486
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492+v471))))
	v497 = v495 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v488) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501+v471))))
	v513 = v511 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v488) {
		goto L126
	} else {
		goto L127
	}
L121:
	;
	v501 = v479 + int32(2)
	if v501 != v470 {
		goto L120
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v545 = v488<<(uint(int32(6))%32)&int32(1984) | v497
	v546 = int32(2)
	goto L117
L124:
	;
	goto L123
L125:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471+v517))))
	v545 = v530&int32(63) | (v488<<(uint(int32(18))%32)&int32(1835008) | v497<<(uint(int32(12))%32) | v513<<(uint(int32(6))%32))
	v546 = int32(4)
	goto L117
L126:
	;
	v517 = v479 + int32(3)
	if v517 != v470 {
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v545 = v488<<(uint(int32(12))%32)&int32(61440) | v497<<(uint(int32(6))%32) | v513
	v546 = int32(3)
	goto L117
L129:
	;
	goto L128
L130:
	;
	v563 = v546 + v479
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v563
	v479 = v563
	goto L112
L131:
	;
	v550 = v545 - int32(97)
	if v550 < int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v550)>>(uint(int32(3))%32)))+uint32(_consts[1324]))))
	if int32(base.Ui32(v556)>>(uint(v550&int32(7))%32))&int32(1) != 0 {
		goto L111
	} else {
		goto L133
	}
L133:
	;
	goto L130
L135:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v579 = v577 + v574
	goto L2
L136:
	;
	return v1290
L137:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v775
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v775
	v780 = F_find_among_b(m, l0, int32(4263936), int32(44))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L58
	} else {
		goto L199
	}
L138:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596+v594))))
	if v598 != int32(108) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v603 = F_find_among_b(m, l0, int32(4263424), int32(2))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L58
	} else {
		goto L140
	}
L140:
	;
	if v603 == int32(0) {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v607
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	if v607 < v610 {
		goto L137
	} else {
		goto L142
	}
L142:
	;
	v613 = v607 - int32(1)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v613 <= v614 {
		goto L137
	} else {
		goto L143
	}
L143:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616+v613))))
	if v618&int32(224) != int32(96) {
		goto L137
	} else {
		goto L144
	}
L144:
	;
	if int32(1)<<(uint(v618)%32)&int32(106790108) == int32(0) {
		goto L137
	} else {
		goto L145
	}
L145:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v632 = F_find_among_b(m, l0, int32(4263472), int32(23))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L58
	} else {
		goto L146
	}
L146:
	;
	if v632 == int32(0) {
		goto L137
	} else {
		goto L147
	}
L147:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v636 + (v607 - v629)
	v640 = F_slice_del(m, l0)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L58
	} else {
		goto L148
	}
L148:
	;
	if v640 < int32(0) {
		v1290 = v640
		goto L136
	} else {
		goto L149
	}
L149:
	;
	v644 = int32(0)
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L153
L150:
	;
	if v770 < int32(0) {
		v1290 = v770
		goto L136
	} else {
		goto L197
	}
L151:
	;
	if v698 < int32(0) {
		v770 = v644
		goto L150
	} else {
		goto L171
	}
L153:
	;
	goto L154
L154:
	;
	goto L155
L155:
	;
	v654 = v646
	v656 = int32(1)
	goto L158
L157:
	;
	v698 = v680
	goto L151
L158:
	;
	if v654 <= v647 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L157
L160:
	;
	v698 = int32(-1)
	goto L151
L161:
	;
	goto L162
L162:
	;
	v661 = v654 - int32(1)
	v663 = int32(*(*int8)(unsafe.Add(mBase, uint32(v645+v661))))
	if int32(0) <= v663 {
		v680 = v661
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v684 = int32(1)
	if v684 < v656 {
		v654 = v680
		v656 = v656 - v684
		goto L158
	} else {
		goto L170
	}
L164:
	;
	if v661 <= v647 {
		v680 = v661
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v668 = v661
	goto L166
L166:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645+v668))))
	if base.Ui32(int32(191)) < base.Ui32(v673) {
		v680 = v668
		goto L163
	} else {
		goto L168
	}
L167:
	;
	v680 = v647
	goto L163
L168:
	;
	v677 = v668 - int32(1)
	if v647 < v677 {
		v668 = v677
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	goto L159
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v698
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v698
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L174
L172:
	;
	if v755 < int32(0) {
		v770 = v644
		goto L150
	} else {
		goto L192
	}
L174:
	;
	goto L175
L175:
	;
	goto L176
L176:
	;
	v711 = v698
	v713 = int32(1)
	goto L179
L178:
	;
	v755 = v737
	goto L172
L179:
	;
	if v711 <= v704 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	goto L178
L181:
	;
	v755 = int32(-1)
	goto L172
L182:
	;
	goto L183
L183:
	;
	v718 = v711 - int32(1)
	v720 = int32(*(*int8)(unsafe.Add(mBase, uint32(v703+v718))))
	if int32(0) <= v720 {
		v737 = v718
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v741 = int32(1)
	if v741 < v713 {
		v711 = v737
		v713 = v713 - v741
		goto L179
	} else {
		goto L191
	}
L185:
	;
	if v718 <= v704 {
		v737 = v718
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v725 = v718
	goto L187
L187:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703+v725))))
	if base.Ui32(int32(191)) < base.Ui32(v730) {
		v737 = v725
		goto L184
	} else {
		goto L189
	}
L188:
	;
	v737 = v704
	goto L184
L189:
	;
	v734 = v725 - int32(1)
	if v704 < v734 {
		v725 = v734
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	goto L180
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v755
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v755
	v761 = F_slice_del(m, l0)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L58
	} else {
		goto L193
	}
L193:
	;
	if int32(0) <= v761 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v768 = int32(1)
	goto L196
L195:
	;
	v768 = v761 >> (uint(int32(31)) % 32) & v761
	goto L196
L196:
	;
	v770 = v768
	goto L150
L197:
	;
	goto L137
L198:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v831
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v831
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v831-int32(2) <= v834 {
		goto L215
	} else {
		goto L216
	}
L199:
	;
	if v780 == int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v784
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	if v784 < v787 {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	v789 = F_slice_del(m, l0)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L58
	} else {
		goto L202
	}
L202:
	;
	if v789 < int32(0) {
		v1290 = v789
		goto L136
	} else {
		goto L203
	}
L203:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v793
	v796 = v793 - int32(1)
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v796 <= v797 {
		goto L198
	} else {
		goto L204
	}
L204:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799+v796))))
	switch v801 - int32(161) {
	case 0, 8:
		goto L205
	default:
		goto L198
	}
L205:
	;
	v806 = F_find_among_b(m, l0, int32(4264816), int32(2))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L58
	} else {
		goto L206
	}
L206:
	;
	if v806 == int32(0) {
		goto L198
	} else {
		goto L207
	}
L207:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v810
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	if v810 < v813 {
		goto L198
	} else {
		goto L208
	}
L208:
	;
	switch v806 - int32(1) {
	case 0:
		goto L210
	case 1:
		goto L209
	default:
		goto L198
	}
L209:
	;
	v825 = F_slice_from_s(m, l0, int32(1), int32(2175664))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L58
	} else {
		goto L213
	}
L210:
	;
	v819 = F_slice_from_s(m, l0, int32(1), int32(2175663))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L58
	} else {
		goto L211
	}
L211:
	;
	if int32(0) <= v819 {
		goto L198
	} else {
		goto L212
	}
L212:
	;
	v1290 = v819
	goto L136
L213:
	;
	if v825 < int32(0) {
		v1290 = v825
		goto L136
	} else {
		goto L214
	}
L214:
	;
	goto L198
L215:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v872
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v872
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v872-int32(3) <= v875 {
		goto L227
	} else {
		goto L228
	}
L216:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838+v831-int32(1)))))
	switch v842 - int32(110) {
	case 0, 6:
		goto L217
	default:
		goto L215
	}
L217:
	;
	v847 = F_find_among_b(m, l0, int32(4264864), int32(3))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L58
	} else {
		goto L218
	}
L218:
	;
	if v847 == int32(0) {
		goto L215
	} else {
		goto L219
	}
L219:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v851
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v853)))
	if v851 < v854 {
		goto L215
	} else {
		goto L220
	}
L220:
	;
	switch v847 - int32(1) {
	case 0:
		goto L222
	case 1:
		goto L221
	default:
		goto L215
	}
L221:
	;
	v866 = F_slice_from_s(m, l0, int32(1), int32(2175670))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L58
	} else {
		goto L225
	}
L222:
	;
	v860 = F_slice_from_s(m, l0, int32(1), int32(2175669))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L58
	} else {
		goto L223
	}
L223:
	;
	if int32(0) <= v860 {
		goto L215
	} else {
		goto L224
	}
L224:
	;
	v1290 = v860
	goto L136
L225:
	;
	if v866 < int32(0) {
		v1290 = v866
		goto L136
	} else {
		goto L226
	}
L226:
	;
	goto L215
L227:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v917
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v917
	v921 = v917 - int32(1)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v921 <= v922 {
		goto L242
	} else {
		goto L243
	}
L228:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879+v872-int32(1)))))
	if v883 != int32(108) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v888 = F_find_among_b(m, l0, int32(4264928), int32(6))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L58
	} else {
		goto L230
	}
L230:
	;
	if v888 == int32(0) {
		goto L227
	} else {
		goto L231
	}
L231:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v892
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v894)))
	if v892 < v895 {
		goto L227
	} else {
		goto L232
	}
L232:
	;
	switch v888 - int32(1) {
	case 0:
		goto L235
	case 1:
		goto L234
	case 2:
		goto L233
	default:
		goto L227
	}
L233:
	;
	v911 = F_slice_from_s(m, l0, int32(1), int32(2175686))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L58
	} else {
		goto L240
	}
L234:
	;
	v905 = F_slice_from_s(m, l0, int32(1), int32(2175685))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L58
	} else {
		goto L238
	}
L235:
	;
	v899 = F_slice_del(m, l0)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L58
	} else {
		goto L236
	}
L236:
	;
	if int32(0) <= v899 {
		goto L227
	} else {
		goto L237
	}
L237:
	;
	v1290 = v899
	goto L136
L238:
	;
	if int32(0) <= v905 {
		goto L227
	} else {
		goto L239
	}
L239:
	;
	v1290 = v905
	goto L136
L240:
	;
	if v911 < int32(0) {
		v1290 = v911
		goto L136
	} else {
		goto L241
	}
L241:
	;
	goto L227
L242:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1093
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1093
	v1097 = v1093 - int32(1)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1097 <= v1098 {
		goto L299
	} else {
		goto L300
	}
L243:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924+v921))))
	switch v926 - int32(161) {
	case 0, 8:
		goto L244
	default:
		goto L242
	}
L244:
	;
	v931 = F_find_among_b(m, l0, int32(4265056), int32(2))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L58
	} else {
		goto L245
	}
L245:
	;
	if v931 == int32(0) {
		goto L242
	} else {
		goto L246
	}
L246:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v935
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	if v935 < v938 {
		goto L242
	} else {
		goto L247
	}
L247:
	;
	v941 = v935 - int32(1)
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v941 <= v942 {
		goto L242
	} else {
		goto L248
	}
L248:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v944+v941))))
	if v946&int32(224) != int32(96) {
		goto L242
	} else {
		goto L249
	}
L249:
	;
	if int32(1)<<(uint(v946)%32)&int32(106790108) == int32(0) {
		goto L242
	} else {
		goto L250
	}
L250:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v960 = F_find_among_b(m, l0, int32(4263472), int32(23))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L58
	} else {
		goto L251
	}
L251:
	;
	if v960 == int32(0) {
		goto L242
	} else {
		goto L252
	}
L252:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v964 + (v935 - v957)
	v968 = F_slice_del(m, l0)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L58
	} else {
		goto L253
	}
L253:
	;
	if v968 < int32(0) {
		v1290 = v968
		goto L136
	} else {
		goto L254
	}
L254:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L257
L255:
	;
	if v1025 < int32(0) {
		goto L242
	} else {
		goto L275
	}
L257:
	;
	goto L258
L258:
	;
	goto L259
L259:
	;
	v981 = v973
	v983 = int32(1)
	goto L262
L261:
	;
	v1025 = v1007
	goto L255
L262:
	;
	if v981 <= v974 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	goto L261
L264:
	;
	v1025 = int32(-1)
	goto L255
L265:
	;
	goto L266
L266:
	;
	v988 = v981 - int32(1)
	v990 = int32(*(*int8)(unsafe.Add(mBase, uint32(v972+v988))))
	if int32(0) <= v990 {
		v1007 = v988
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1011 = int32(1)
	if v1011 < v983 {
		v981 = v1007
		v983 = v983 - v1011
		goto L262
	} else {
		goto L274
	}
L268:
	;
	if v988 <= v974 {
		v1007 = v988
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v995 = v988
	goto L270
L270:
	;
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972+v995))))
	if base.Ui32(int32(191)) < base.Ui32(v1000) {
		v1007 = v995
		goto L267
	} else {
		goto L272
	}
L271:
	;
	v1007 = v974
	goto L267
L272:
	;
	v1004 = v995 - int32(1)
	if v974 < v1004 {
		v995 = v1004
		goto L270
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	goto L263
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1025
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L278
L276:
	;
	if v1082 < int32(0) {
		goto L242
	} else {
		goto L296
	}
L278:
	;
	goto L279
L279:
	;
	goto L280
L280:
	;
	v1038 = v1025
	v1040 = int32(1)
	goto L283
L282:
	;
	v1082 = v1064
	goto L276
L283:
	;
	if v1038 <= v1031 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	goto L282
L285:
	;
	v1082 = int32(-1)
	goto L276
L286:
	;
	goto L287
L287:
	;
	v1045 = v1038 - int32(1)
	v1047 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1030+v1045))))
	if int32(0) <= v1047 {
		v1064 = v1045
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1068 = int32(1)
	if v1068 < v1040 {
		v1038 = v1064
		v1040 = v1040 - v1068
		goto L283
	} else {
		goto L295
	}
L289:
	;
	if v1045 <= v1031 {
		v1064 = v1045
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1052 = v1045
	goto L291
L291:
	;
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1030+v1052))))
	if base.Ui32(int32(191)) < base.Ui32(v1057) {
		v1064 = v1052
		goto L288
	} else {
		goto L293
	}
L292:
	;
	v1064 = v1031
	goto L288
L293:
	;
	v1061 = v1052 - int32(1)
	if v1031 < v1061 {
		v1052 = v1061
		goto L291
	} else {
		goto L294
	}
L294:
	;
	goto L292
L295:
	;
	goto L284
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1082
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1082
	v1087 = F_slice_del(m, l0)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L58
	} else {
		goto L297
	}
L297:
	;
	if v1087 < int32(0) {
		v1290 = v1087
		goto L136
	} else {
		goto L298
	}
L298:
	;
	goto L242
L299:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1139
	v1141 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1139
	v1146 = F_find_among_b(m, l0, int32(4265344), int32(31))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L58
	} else {
		goto L315
	}
L300:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1100+v1097))))
	if base.B2i32(v1102 != int32(169))&base.B2i32(v1102 != int32(105)) != 0 {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v1110 = F_find_among_b(m, l0, int32(4265104), int32(12))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L58
	} else {
		goto L302
	}
L302:
	;
	if v1110 == int32(0) {
		goto L299
	} else {
		goto L303
	}
L303:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1114
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)))
	if v1114 < v1117 {
		goto L299
	} else {
		goto L304
	}
L304:
	;
	switch v1110 - int32(1) {
	case 0:
		goto L307
	case 1:
		goto L306
	case 2:
		goto L305
	default:
		goto L299
	}
L305:
	;
	v1133 = F_slice_from_s(m, l0, int32(1), int32(2175725))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L58
	} else {
		goto L312
	}
L306:
	;
	v1127 = F_slice_from_s(m, l0, int32(1), int32(2175724))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L58
	} else {
		goto L310
	}
L307:
	;
	v1121 = F_slice_del(m, l0)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L58
	} else {
		goto L308
	}
L308:
	;
	if int32(0) <= v1121 {
		goto L299
	} else {
		goto L309
	}
L309:
	;
	v1290 = v1121
	goto L136
L310:
	;
	if int32(0) <= v1127 {
		goto L299
	} else {
		goto L311
	}
L311:
	;
	v1290 = v1127
	goto L136
L312:
	;
	if v1133 < int32(0) {
		v1290 = v1133
		goto L136
	} else {
		goto L313
	}
L313:
	;
	goto L299
L314:
	;
	if v1175 < int32(0) {
		v1290 = v1175
		goto L136
	} else {
		goto L328
	}
L315:
	;
	if v1146 == int32(0) {
		v1175 = v1141
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1150
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	if v1150 < v1153 {
		v1175 = v1141
		goto L314
	} else {
		goto L317
	}
L317:
	;
	switch v1146 - int32(1) {
	case 0:
		goto L321
	case 1:
		goto L320
	case 2:
		goto L319
	default:
		goto L318
	}
L318:
	;
	v1175 = int32(1)
	goto L314
L319:
	;
	v1169 = F_slice_from_s(m, l0, int32(1), int32(2175776))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L58
	} else {
		goto L326
	}
L320:
	;
	v1163 = F_slice_from_s(m, l0, int32(1), int32(2175775))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L58
	} else {
		goto L324
	}
L321:
	;
	v1157 = F_slice_del(m, l0)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L58
	} else {
		goto L322
	}
L322:
	;
	if int32(0) <= v1157 {
		goto L318
	} else {
		goto L323
	}
L323:
	;
	v1175 = v1157
	goto L314
L324:
	;
	if int32(0) <= v1163 {
		goto L318
	} else {
		goto L325
	}
L325:
	;
	v1175 = v1163
	goto L314
L326:
	;
	if v1169 < int32(0) {
		v1175 = v1169
		goto L314
	} else {
		goto L327
	}
L327:
	;
	goto L318
L328:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1179
	v1181 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1179
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1179 <= v1184 {
		v1233 = v1181
		goto L329
	} else {
		goto L330
	}
L329:
	;
	if v1233 < int32(0) {
		v1290 = v1233
		goto L136
	} else {
		goto L346
	}
L330:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186+v1179-int32(1)))))
	if v1190&int32(224) != int32(96) {
		v1233 = v1181
		goto L329
	} else {
		goto L331
	}
L331:
	;
	if int32(1)<<(uint(v1190)%32)&int32(10768) == int32(0) {
		v1233 = v1181
		goto L329
	} else {
		goto L332
	}
L332:
	;
	v1203 = F_find_among_b(m, l0, int32(4265968), int32(42))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L58
	} else {
		goto L333
	}
L333:
	;
	if v1203 == int32(0) {
		v1233 = v1181
		goto L329
	} else {
		goto L334
	}
L334:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1207
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1209)))
	if v1207 < v1210 {
		v1233 = v1181
		goto L329
	} else {
		goto L335
	}
L335:
	;
	switch v1203 - int32(1) {
	case 0:
		goto L339
	case 1:
		goto L338
	case 2:
		goto L337
	default:
		goto L336
	}
L336:
	;
	v1233 = int32(1)
	goto L329
L337:
	;
	v1226 = F_slice_from_s(m, l0, int32(1), int32(2175858))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L58
	} else {
		goto L344
	}
L338:
	;
	v1220 = F_slice_from_s(m, l0, int32(1), int32(2175857))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L58
	} else {
		goto L342
	}
L339:
	;
	v1214 = F_slice_del(m, l0)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L58
	} else {
		goto L340
	}
L340:
	;
	if int32(0) <= v1214 {
		goto L336
	} else {
		goto L341
	}
L341:
	;
	v1233 = v1214
	goto L329
L342:
	;
	if int32(0) <= v1220 {
		goto L336
	} else {
		goto L343
	}
L343:
	;
	v1233 = v1220
	goto L329
L344:
	;
	if v1226 < int32(0) {
		v1233 = v1226
		goto L329
	} else {
		goto L345
	}
L345:
	;
	goto L336
L346:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1237
	v1239 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1237
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1237 <= v1242 {
		v1282 = v1239
		goto L347
	} else {
		goto L348
	}
L347:
	;
	if v1282 < int32(0) {
		v1290 = v1282
		goto L136
	} else {
		goto L363
	}
L348:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1244+v1237-int32(1)))))
	if v1248 != int32(107) {
		v1282 = v1239
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1253 = F_find_among_b(m, l0, int32(4266816), int32(7))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L58
	} else {
		goto L350
	}
L350:
	;
	if v1253 == int32(0) {
		v1282 = v1239
		goto L347
	} else {
		goto L351
	}
L351:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1257
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1259)))
	if v1257 < v1260 {
		v1282 = v1239
		goto L347
	} else {
		goto L352
	}
L352:
	;
	switch v1253 - int32(1) {
	case 0:
		goto L356
	case 1:
		goto L355
	case 2:
		goto L354
	default:
		goto L353
	}
L353:
	;
	v1282 = int32(1)
	goto L347
L354:
	;
	v1276 = F_slice_del(m, l0)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L58
	} else {
		goto L361
	}
L355:
	;
	v1272 = F_slice_from_s(m, l0, int32(1), int32(2176018))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L58
	} else {
		goto L359
	}
L356:
	;
	v1266 = F_slice_from_s(m, l0, int32(1), int32(2176017))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L58
	} else {
		goto L357
	}
L357:
	;
	if int32(0) <= v1266 {
		goto L353
	} else {
		goto L358
	}
L358:
	;
	v1282 = v1266
	goto L347
L359:
	;
	if int32(0) <= v1272 {
		goto L353
	} else {
		goto L360
	}
L360:
	;
	v1282 = v1272
	goto L347
L361:
	;
	if v1276 < int32(0) {
		v1282 = v1276
		goto L347
	} else {
		goto L362
	}
L362:
	;
	goto L353
L363:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1287
	v1290 = int32(1)
	goto L136
}
func F_hypothetical_dense_rank_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var __phi250 int32
	_ = __phi250
	var v254 int32
	_ = v254
	var __phi254 int32
	_ = __phi254
	var v258 int32
	_ = v258
	var __phi258 int32
	_ = __phi258
	var v262 int64
	_ = v262
	var __phi262 int64
	_ = __phi262
	var v263 int64
	_ = v263
	var __phi263 int64
	_ = __phi263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int64
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v341 int64
	_ = v341
	var v342 int64
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	v14 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
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
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L66
	}
L2:
	;
	m.G0 = v18 + int32(16)
	return v371
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
	v32 = v20 - int32(1)
	v34 = l0 + int32(20)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v37 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	v371 = v27
	goto L2
L8:
	;
	v40 = int32(4470400)
	v41 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v43
	v45 = F_CreateStandaloneExprContext(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v55 = v37
	goto L10
L10:
	;
	if v32&int32(1) != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v41
	v55 = v50
	goto L10
L12:
	;
	v59 = v32 >> (uint(int32(1)) % 32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
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
	v66 = int32(4470400)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v67
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+48)) = v80
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v91 = v80
	v92 = v86
	goto L16
L18:
	;
	v98 = int32(0)
	if v59 <= v98 {
		v202 = v98
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v203+v202<<(uint(int32(2))%32)))) = int32(-1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v209+v202))) = uint8(v211)
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
	v215 = v213 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)) = uint16(v215)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+6)) = uint16(v218)
	goto L28
L20:
	;
	v101 = int32(0)
	if v32 != int32(2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v108 = v101
	v116 = int32(0)
	goto L24
L22:
	;
	v156 = v101
	goto L23
L23:
	;
	if v32&int32(2) == int32(0) {
		v202 = v59
		goto L19
	} else {
		goto L27
	}
L24:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v123 = int32(2)
	v127 = v108 | int32(1)
	v128 = int32(3)
	v130 = v34 + v127<<(uint(v128)%32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	*(*int32)(unsafe.Add(mBase, uint32(v122+v108<<(uint(v123)%32)))) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133+v108))) = uint8(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v142 = v108 + v123
	v145 = v34 + v142<<(uint(v128)%32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	*(*int32)(unsafe.Add(mBase, uint32(v137+v127<<(uint(v123)%32)))) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v127+v148))) = uint8(v150)
	v153 = v116 + v123
	if v153 != v59&int32(2147483646) {
		v108 = v142
		v116 = v153
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v156 = v142
	goto L23
L26:
	;
	goto L25
L27:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v180 = v156<<(uint(int32(3))%32) + v34
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v174+v156<<(uint(int32(2))%32)))) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v183+v156))) = uint8(v185)
	v202 = v59
	goto L19
L28:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	F_tuplesort_puttupleslot(m, v220, v93)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	F_tuplesort_performsort(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v226)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	v231 = F_MakeSingleTupleTableSlot(m, v229, int32(1591740))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v234 = int32(1)
	v238 = F_tuplesort_gettupleslot(m, v233, v234, v234, v93, v18+int32(8))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	m.T0[v344].(func(*base.Module, int32))(m, v328)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L62
	}
L33:
	;
	if v238 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v328 = v93
	v330 = v231
	v341 = v14
	v342 = int64(1)
	goto L32
L35:
	;
	goto L36
L36:
	;
	__phi250 = v93
	__phi254 = v231
	__phi258 = int32(0)
	__phi262 = v14
	__phi263 = int64(1)
	v250 = __phi250
	v254 = __phi254
	v258 = __phi258
	v262 = __phi262
	v263 = __phi263
	goto L37
L37:
	;
	v264 = int32(*(*int16)(unsafe.Add(mBase, uint32(v250)+6)))
	if v264 <= v59 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v328 = v254
	v330 = v250
	v341 = v313
	v342 = v320
	goto L32
L39:
	;
	F_slot_getsomeattrs_int(m, v250, v59+int32(1))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v250)+20))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v59))))
	if v270 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v250
	if v254 == int32(0) {
		v313 = v262
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v271+v59<<(uint(int32(2))%32))))
	if v273 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v328 = v250
	v330 = v254
	v341 = v262
	v342 = v263
	goto L32
L46:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v316 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v316 != 0 {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+4)))
	if v280&int32(2) != 0 {
		v313 = v262
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v283 != v258 {
		v313 = v262
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
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	F_MemoryContextReset(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v292 = int32(4470400)
	v293 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v295
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v300 = m.T0[v299].(func(*base.Module, int32, int32, int32) int32)(m, v91, v55, v18+int32(15))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L54
	}
L53:
	;
	v313 = v262 + int64(1)
	goto L46
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v293
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	F_MemoryContextReset(m, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v313 = v262 + base.I64_extend_i32_u(base.B2i32(v300 != int32(0)))
	goto L46
L56:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v320 = v263 + int64(1)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v322 = int32(1)
	v326 = F_tuplesort_gettupleslot(m, v321, v322, v322, v254, v18+int32(8))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	if v326 != 0 {
		__phi250 = v254
		__phi254 = v250
		__phi258 = v314
		__phi262 = v313
		__phi263 = v320
		v250 = __phi250
		v254 = __phi254
		v258 = __phi258
		v262 = __phi262
		v263 = __phi263
		goto L37
	} else {
		goto L61
	}
L61:
	;
	goto L38
L62:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	m.T0[v348].(func(*base.Module, int32))(m, v330)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	F_ExecDropSingleTupleTableSlot(m, v231)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v354 = F_Int64GetDatum(m, v342-v341)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v371 = v354
	goto L2
L66:
	;
	F_errmsg_internal(m, int32(247740), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(485213), int32(1332), int32(307885))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
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
