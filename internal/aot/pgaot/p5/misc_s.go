package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_SearchNamedReplicationSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v3 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_SearchNamedReplicationSlot[0]))
	v12 = F_LWLockAcquire(m, v8+int32(_a_F_SearchNamedReplicationSlot_0), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_SearchNamedReplicationSlot[1]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_SearchNamedReplicationSlot[2]))
	v20 = v17 + v19
	if v20 <= int32(0) {
		v73 = v3
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	if l1 != 0 {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_SearchNamedReplicationSlot[3]))
	v28 = v3
	goto L8
L8:
	;
	v33 = v24 + v28*int32(296)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
	if v34 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v73 = int32(0)
	goto L6
L10:
	;
	v38 = v33 + int32(24)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if base.B2i32(v41 == int32(0))|base.B2i32(v41 != v44) != 0 {
		v62 = v41
		v63 = v44
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v68 = v28 + int32(1)
	if v68 != v20 {
		v28 = v68
		goto L8
	} else {
		goto L21
	}
L13:
	;
	if v62-v63 == int32(0) {
		v73 = v33
		goto L6
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v47 = l0
	v48 = v38
	goto L16
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v52 == int32(0) {
		v62 = v52
		v63 = v51
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v62 = v52
	v63 = v51
	goto L14
L18:
	;
	v55 = int32(1)
	if v52 == v51 {
		v47 = v47 + v55
		v48 = v48 + v55
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L12
L21:
	;
	goto L9
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_SearchNamedReplicationSlot[0]))
	F_LWLockRelease(m, v78+int32(_a_F_SearchNamedReplicationSlot_0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	return v73
L25:
	;
	goto L24
}
func F_SetUserIdAndSecContext(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_SetUserIdAndSecContext[0])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_SetUserIdAndSecContext[1])) = l0
	return
}
func F_SimpleLruReadPage_ReadOnly(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v13 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
	v14 = base.I64_rem_s(l1, v13)
	v15 = base.I32_wrap_i64(v14)
	v18 = v12 + v15<<(uint(int32(7))%32)
	v20 = F_LWLockAcquire(m, v18, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = v15 << (uint(int32(4)) % 32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v32 = v25
	goto L4
L3:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v90 = int32(2)
	v92 = v87 + v85>>(uint(int32(4))%32)<<(uint(v90)%32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v95 = v85 << (uint(v90) % 32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95+v96)))
	if v93 != v98 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28+v32<<(uint(int32(2))%32))))
	if base.B2i32(v42 == int32(0))|base.B2i32(v42 == int32(1)) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_LWLockRelease(m, v18)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v54 = int32(1)
	v55 = v32 | v54
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v28+v55<<(uint(int32(2))%32))))
	v60 = int32(0)
	if base.B2i32(v59 == v60)|base.B2i32(v59 == v54) == v60 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v48+v32<<(uint(int32(3))%32))))
	if v52 != l1 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v85 = v32
	goto L3
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v67+v55<<(uint(int32(3))%32))))
	if v71 == l1 {
		v85 = v55
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v32 != v25|int32(14) {
		v32 = v32 + int32(2)
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	goto L5
L14:
	;
	v79 = F_LWLockAcquire(m, v18, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v82 = F_SimpleLruReadPage(m, l0, l1, int32(1), l2)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	return v82
L17:
	;
	v101 = v93 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v95))) = v101
	goto L19
L18:
	;
	goto L19
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage_ReadOnly[0])) = uint8(v109)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruReadPage_ReadOnly[1])) = uint8(v109)
	v115 = v107 << (uint(int32(6)) % 32)
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_SimpleLruReadPage_ReadOnly[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_SimpleLruReadPage_ReadOnly[2]))) = v118 + int64(1)
	goto L20
L20:
	;
	return v85
}
func F_StatisticsObjIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_SearchSysCache1(m, int32(64), base.I64_extend_i32_u(l0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v138
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v138 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	F_errmsg_internal(m, int32(_a_F_StatisticsObjIsVisibleExt_0), v10)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_StatisticsObjIsVisibleExt_1), int32(2728), int32(_a_F_StatisticsObjIsVisibleExt_2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v40 = v36 + v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+72))
	if v41 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L46
	}
L15:
	;
	v44 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_StatisticsObjIsVisibleExt[0]))
	if v46 == v44 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v89 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_StatisticsObjIsVisibleExt[0]))
	if v91 == v89 {
		v129 = v89
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v85 == int32(0) {
		v129 = v44
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v85 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v53 <= int32(0) {
		v79 = v44
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v85 = v79
	goto L18
L23:
	;
	v56 = int32(0)
	if v56 < v53 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v59 = v53
	goto L26
L25:
	;
	v59 = v56
	goto L26
L26:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v62 = int32(0)
	goto L27
L27:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60+v62<<(uint(int32(2))%32))))
	v71 = base.B2i32(v70 == v41)
	if v70 == v41 {
		v79 = v71
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v79 = v71
	goto L22
L29:
	;
	v73 = v62 + int32(1)
	if v73 != v59 {
		v62 = v73
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v94 <= int32(0) {
		v129 = v89
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v100 = v89
	goto L34
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v100<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_StatisticsObjIsVisibleExt[1]))
	if v111 != v113 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v129 = int32(0)
	goto L14
L36:
	;
	goto L35
L37:
	;
	if v111 == v41 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v124 = v100 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v124 < v125 {
		v100 = v124
		goto L34
	} else {
		goto L45
	}
L40:
	;
	v129 = int32(1)
	goto L14
L41:
	;
	goto L42
L42:
	;
	v119 = int64(0)
	v121 = F_SearchSysCacheExists(m, int32(63), base.I64_extend_i32_u(v40+int32(8)), base.I64_extend_i32_u(v111), v119, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	if v121 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	goto L39
L45:
	;
	goto L36
L46:
	;
	v138 = v129
	goto L1
}
func F_StoreConnectionWarning(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StoreConnectionWarning[0])))
	if v7 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_StoreConnectionWarning_0), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_StoreConnectionWarning_1), int32(1515), int32(_a_F_StoreConnectionWarning_2))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v21 = int32(_a_F_StoreConnectionWarning_3)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_StoreConnectionWarning[1]))
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_StoreConnectionWarning[2]))
		*(*int32)(unsafe.Add(mBase, _c_F_StoreConnectionWarning[1])) = v25
		v28 = F_palloc(m, int32(12))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = l0
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_StoreConnectionWarning[3]))
			v35 = F_lappend(m, v34, v28)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_StoreConnectionWarning[1])) = v22
				*(*int32)(unsafe.Add(mBase, _c_F_StoreConnectionWarning[3])) = v35
				return
			}
		}
	}
}
func F_StorePreparedStatement(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v3 = l2
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_StorePreparedStatement[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_StorePreparedStatement[1]))
	if v14 == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(343597383744)
		v25 = F_hash_create(m, int32(_a_F_StorePreparedStatement_0), int64(32), v7+int32(-48), int32(24))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_StorePreparedStatement[1])) = v25
			v28 = v25
			v32 = F_hash_search(m, v28, l0, int32(1), v7+int32(-48))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
				if v34 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errcode(m, int32(84017284))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg(m, int32(_a_F_StorePreparedStatement_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_StorePreparedStatement_2), int32(417), int32(_a_F_StorePreparedStatement_3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v12
					*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l1
					F_SaveCachedPlan(m, l1)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						m.G0 = v9 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		v28 = v14
		v32 = F_hash_search(m, v28, l0, int32(1), v7+int32(-48))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
			if v34 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errcode(m, int32(84017284))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(_a_F_StorePreparedStatement_1), v9)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_StorePreparedStatement_2), int32(417), int32(_a_F_StorePreparedStatement_3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v12
				*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l1
				F_SaveCachedPlan(m, l1)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					m.G0 = v9 - int32(-64)
					return
				}
			}
		}
	}
}
func F_StoreSingleInheritance(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_s(l2)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(l1)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(l0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
		v29 = F_heap_form_tuple(m, v24, v8+int32(16), v8+int32(12))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			F_CatalogTupleInsert(m, v12, v29)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_pfree(m, v29)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_relation_close(m, v12, int32(3))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v8 + int32(48)
						return
					}
				}
			}
		}
	}
}
func F_SwitchBackToLocalLatch(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	*(*int32)(unsafe.Add(mBase, _c_F_SwitchBackToLocalLatch[0])) = int32(_a_F_SwitchBackToLocalLatch_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchBackToLocalLatch[1]))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = int32(1)
	F_ModifyWaitEvent(m, v6, v7, v7, int32(_a_F_SwitchBackToLocalLatch_0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v15 = int32(_a_F_SwitchBackToLocalLatch_0)
	goto L3
L3:
	;
	v16 = int32(0)
	v19 = base.AtomicRmwOr32(m, v16, int32(_a_F_SwitchBackToLocalLatch_1), v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchBackToLocalLatch[0]))
	v15 = v13
	goto L3
L6:
	;
	return
L7:
	;
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
	v23 = int32(0)
	v26 = base.AtomicRmwOr32(m, v23, int32(_a_F_SwitchBackToLocalLatch_1), v23)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v27 == v23 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v30 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchBackToLocalLatch[2]))
	if v34 == v30 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v36 = m.G0
	v38 = v36 - int32(16)
	m.G0 = v38
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchBackToLocalLatch[3]))
	if v41 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v64 = F_pgmem_kill(m, v30, int32(23))
	mBase = m.M
	goto L7
L14:
	;
	m.G0 = v38 + int32(16)
	goto L6
L15:
	;
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+15)) = uint8(v44)
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchBackToLocalLatch[4]))
	v52 = F_write(m, v48, v38+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v52 {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	goto L14
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchBackToLocalLatch[5]))
	if v56 == int32(27) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
}
func F___strxfrm_l(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v5 = F_strlen(m, l1)
	mBase = m.M
	if base.Ui32(v5) < base.Ui32(l2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if (l1^l0)&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L3
L3:
	;
	return v5
L4:
	;
	goto L3
L5:
	;
	goto L4
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v60)
	if v60&int32(255) == int32(0) {
		goto L5
	} else {
		goto L21
	}
L7:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v59 = l1
	v60 = v12
	v61 = l0
	goto L6
L8:
	;
	goto L9
L9:
	;
	if l1&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v16 = l1
	v18 = l0
	goto L13
L11:
	;
	v30 = l1
	v32 = l0
	goto L12
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 != v37 {
		v59 = v30
		v60 = v34
		v61 = v32
		goto L6
	} else {
		goto L17
	}
L13:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v19)
	if v19 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v30 = v26
	v32 = v24
	goto L12
L15:
	;
	v23 = int32(1)
	v24 = v18 + v23
	v26 = v16 + v23
	if v26&int32(3) != 0 {
		v16 = v26
		v18 = v24
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v42 = v30
	v43 = v34
	v44 = v32
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v43
	v46 = int32(4)
	v47 = v44 + v46
	v49 = v42 + v46
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v54 = int32(-2139062144)
	if (int32(16843008)-v51|v51)&v54 == v54 {
		v42 = v49
		v43 = v51
		v44 = v47
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v59 = v49
	v60 = v51
	v61 = v47
	goto L6
L20:
	;
	goto L19
L21:
	;
	v68 = v59
	v70 = v61
	goto L22
L22:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)) = uint8(v71)
	v73 = int32(1)
	if v71 != 0 {
		v68 = v68 + v73
		v70 = v70 + v73
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L5
L24:
	;
	goto L23
}
func F___subtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	F___addtf3(m, v9, l1, l2, l3, l4^int64(-9223372036854775807-1))
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v14
	m.G0 = v9 + v8
	return
}
func F_sampler_random_fract(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v30 float64
	_ = v30
	for {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = v7 ^ v8
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = base.I64_rotl(v9, int64(37))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9<<(uint(int64(16))%64) ^ base.I64_rotl(v7, int64(24)) ^ v9
		v30 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v7*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
		mBase = m.M
		if base.F64_eq(v30, float64(0)) != 0 {
			continue
		} else {
			break
		}
		break
	}
	return v30
}
func F_scalarineqsel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v82 float64
	_ = v82
	var v86 float64
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v111 float64
	_ = v111
	var v114 float32
	_ = v114
	var v117 float64
	_ = v117
	var v120 float64
	_ = v120
	var v140 float64
	_ = v140
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v21 == int32(0) {
		v24 = float64(0.3333333333333333)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
		if v25 == int32(0) {
			v140 = v24
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			if base.B2i32(v28 != int32(6))|base.B2i32(l7 != int32(27)) != 0 {
				v140 = v24
			} else {
				v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)))
				if v34 != int32(_a_F_scalarineqsel_0) {
					v140 = v24
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+124))
					if v38 == int32(0) {
						v140 = float64(1)
					} else {
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v37)+128))
						v44 = base.F64_add(base.F64_convert_i32_u(v38), float64(-0.5))
						v45 = base.F64_div(v41, v44)
						v48 = base.I32_wrap_i64(l6)
						v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48))))
						v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)))
						v54 = base.F64_convert_i32_u(v49<<(uint(int32(16))%32) | v52)
						if base.F64_ge(v54, base.F64_convert_i32_u(v38-int32(1))) != 0 {
							v59 = base.F64_mul(v45, float64(0.5))
						} else {
							v59 = v45
						}
						if base.F64_gt(v59, float64(0)) != 0 {
							v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)))
							v64 = base.F64_div(base.F64_convert_i32_u(v62), v59)
							v65 = float64(1)
							if base.F64_lt(v64, v65) != 0 {
								v68 = v64
							} else {
								v68 = v65
							}
							v71 = base.F64_add(v68, v54)
						} else {
							v71 = v54
						}
						v72 = base.F64_div(v71, v44)
						if base.B2i32(base.F64_ge(v41, float64(1)) == int32(0))|base.B2i32(l2 != l3) != 0 {
							v82 = v72
						} else {
							v82 = base.F64_add(v72, base.F64_div(float64(-1), v41))
						}
						if l2 != 0 {
							v86 = base.F64_sub(float64(1), v82)
						} else {
							v86 = v82
						}
						if base.F64_lt(v86, float64(0)) != 0 {
							v140 = float64(0)
						} else {
							if base.F64_gt(v86, float64(1)) != 0 {
								v140 = float64(1)
							} else {
								v140 = v86
							}
						}
					}
				}
			}
		}
		m.G0 = v19 + int32(48)
		return v140
	} else {
		v91 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
		v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
		v93 = F_get_opcode(m, l1)
		mBase = m.M
		v96 = m.ExcPending
		if v96 != 0 {
			return float64(0)
		} else {
			v98 = v19 + int32(20)
			F_fmgr_info(m, v93, v98)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return float64(0)
			} else {
				v104 = F_mcv_selectivity(m, l5, v98, l4, l6, int32(1), v19+int32(8))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return float64(0)
				} else {
					v106 = F_ineq_histogram_selectivity(m, l0, l5, l1, v98, l2, l3, l4, l6, l7)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return float64(0)
					} else {
						if base.F64_ge(v106, float64(0)) != 0 {
							v111 = v106
						} else {
							v111 = float64(0.5)
						}
						v114 = *(*float32)(unsafe.Add(mBase, uint32(v91+v92)+8))
						v117 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
						v120 = base.F64_add(v104, base.F64_mul(v111, base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v114)), v117)))
						if base.F64_lt(v120, float64(0)) != 0 {
							v140 = float64(0)
						} else {
							if base.F64_gt(v120, float64(1)) == int32(0) {
								v140 = v120
							} else {
								v140 = float64(1)
							}
						}
						m.G0 = v19 + int32(48)
						return v140
					}
				}
			}
		}
	}
}
func F_scalbn(m *base.Module, l0 float64, l1 int32) float64 {
	var v6 float64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 float64
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	if int32(1024) <= l1 {
		v6 = base.F64_mul(l0, float64(8.98846567431158e+307))
		if base.Ui32(l1) < base.Ui32(int32(2047)) {
			v35 = v6
			v36 = l1 - int32(1023)
		} else {
			v13 = int32(3069)
			if base.Ui32(v13) <= base.Ui32(l1) {
				v16 = v13
			} else {
				v16 = l1
			}
			v35 = base.F64_mul(v6, float64(8.98846567431158e+307))
			v36 = v16 - int32(2046)
		}
	} else {
		if int32(-1023) < l1 {
			v35 = l0
			v36 = l1
		} else {
			v22 = base.F64_mul(l0, float64(2.004168360008973e-292))
			if base.Ui32(int32(-1992)) < base.Ui32(l1) {
				v35 = v22
				v36 = l1 + int32(969)
			} else {
				v29 = int32(-2960)
				if base.Ui32(l1) <= base.Ui32(v29) {
					v32 = v29
				} else {
					v32 = l1
				}
				v35 = base.F64_mul(v22, float64(2.004168360008973e-292))
				v36 = v32 + int32(1938)
			}
		}
	}
	return base.F64_mul(v35, base.F64_reinterpret_i64(base.I64_extend_i32_u(v36+int32(1023))<<(uint(int64(52))%64)))
}
func F_scan_identifier(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
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
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v10 != int32(34) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v220
L2:
	;
	v96 = v46 - v9
	v98 = F_downcase_truncate_identifier(m, v9, v96, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L37
	} else {
		goto L38
	}
L3:
	;
	if v10 == int32(0) {
		v220 = v4
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v50 = v9 + int32(1)
	v51 = int32(34)
	v52 = F___strchrnul(m, v50, v51)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v54 == v51 {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v16 = l2 & int32(255)
	if v10 == v16 {
		v220 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v21 = v10
	goto L9
L8:
	;
	if v9 == v46 {
		v220 = v4
		goto L1
	} else {
		goto L17
	}
L9:
	;
	v26 = base.I32_extend8_s(v21)
	goto L11
L10:
	;
	v46 = v38
	goto L8
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(v26 == int32(32))|base.B2i32(base.Ui32((v26-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v46 = v36
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v38 = v36 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v40 != v16 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v43 = v40
	goto L15
L14:
	;
	v43 = int32(0)
	goto L15
L15:
	;
	if v43 != 0 {
		v21 = v40
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v46
	goto L2
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v58
	if v58 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v58 = v52
	goto L21
L20:
	;
	v58 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	return int32(0)
L23:
	;
	goto L24
L24:
	;
	v67 = v58
	goto L26
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v73
	return v50
L26:
	;
	v73 = v67 + int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v74 != int32(34) {
		goto L25
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	v77 = F_strlen(m, v67)
	mBase = m.M
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	base.MemoryCopy(m, v67, v73, v77)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v79
	v83 = int32(34)
	v84 = F___strchrnul(m, v79+int32(1), v83)
	mBase = m.M
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v86 == v83 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v90
	if v90 != 0 {
		v67 = v90
		goto L26
	} else {
		goto L36
	}
L33:
	;
	v90 = v84
	goto L35
L34:
	;
	v90 = int32(0)
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L27
L37:
	;
	return int32(0)
L38:
	;
	if (v98^v9)&int32(3) != 0 {
		v172 = v98
		v173 = v96
		v174 = v9
		goto L43
	} else {
		goto L44
	}
L39:
	;
	F_pfree(m, v98)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L37
	} else {
		goto L64
	}
L40:
	;
	F___memset(m, v207, int32(0), v206)
	mBase = m.M
	goto L39
L41:
	;
	v206 = int32(0)
	v207 = v201
	goto L40
L42:
	;
	v184 = v179
	v185 = v180
	v186 = v181
	goto L60
L43:
	;
	if v173 == int32(0) {
		v201 = v174
		goto L41
	} else {
		goto L59
	}
L44:
	;
	v107 = int32(0)
	if base.B2i32(v98&int32(3) == v107)|base.B2i32(v96 == v107) != 0 {
		v138 = v98
		v139 = v96
		v140 = v9
		v141 = base.B2i32(v96 != v107)
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v141 == int32(0) {
		v201 = v140
		goto L41
	} else {
		goto L52
	}
L46:
	;
	v117 = v98
	v118 = v96
	v119 = v9
	goto L47
L47:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v121)
	if v121 == int32(0) {
		v206 = v118
		v207 = v119
		goto L40
	} else {
		goto L49
	}
L48:
	;
	v138 = v132
	v139 = v128
	v140 = v126
	v141 = v130
	goto L45
L49:
	;
	v125 = int32(1)
	v126 = v119 + v125
	v128 = v118 - v125
	v129 = int32(0)
	v130 = base.B2i32(v128 != v129)
	v132 = v117 + v125
	if v132&int32(3) == v129 {
		v138 = v132
		v139 = v128
		v140 = v126
		v141 = v130
		goto L45
	} else {
		goto L50
	}
L50:
	;
	if v128 != 0 {
		v117 = v132
		v118 = v128
		v119 = v126
		goto L47
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v144 == int32(0) {
		v206 = v139
		v207 = v140
		goto L40
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32(v139) < base.Ui32(int32(4)) {
		v172 = v138
		v173 = v139
		v174 = v140
		goto L43
	} else {
		goto L54
	}
L54:
	;
	v150 = v138
	v151 = v139
	v152 = v140
	goto L55
L55:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v158 = int32(-2139062144)
	if (int32(16843008)-v155|v155)&v158 != v158 {
		v179 = v150
		v180 = v151
		v181 = v152
		goto L42
	} else {
		goto L57
	}
L56:
	;
	v172 = v166
	v173 = v168
	v174 = v164
	goto L43
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v155
	v163 = int32(4)
	v164 = v152 + v163
	v166 = v150 + v163
	v168 = v151 - v163
	if base.Ui32(int32(3)) < base.Ui32(v168) {
		v150 = v166
		v151 = v168
		v152 = v164
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v179 = v172
	v180 = v173
	v181 = v174
	goto L42
L60:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v188)
	if v188 == int32(0) {
		v206 = v185
		v207 = v186
		goto L40
	} else {
		goto L62
	}
L61:
	;
	v201 = v193
	goto L41
L62:
	;
	v192 = int32(1)
	v193 = v186 + v192
	v197 = v185 - v192
	if v197 != 0 {
		v184 = v184 + v192
		v185 = v197
		v186 = v193
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v220 = v9
	goto L1
}
func F_scanner_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v7 = F_strlen(m, l0)
	mBase = m.M
	v9 = F_palloc(m, int32(100))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			base.MemoryFill(m, v9+int32(4), int32(0), int32(96))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_scanner_init[0]))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v22
			v25 = v7 + int32(2)
			v26 = F_palloc(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
				if v7 != 0 {
					base.MemoryCopy(m, v26, l0, v7)
				} else {
				}
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v33 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v31+v7)+1)) = uint8(v33)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*uint8)(unsafe.Add(mBase, uint32(v35+v7))) = uint8(v33)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui32(v25) < base.Ui32(int32(2)) {
					v127 = int32(1024)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v127
					v130 = F_palloc(m, v127)
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v130
						return v9
					}
				} else {
					v43 = v25 - int32(2)
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v43))))
					if v45 != 0 {
						v127 = int32(1024)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v127
						v130 = F_palloc(m, v127)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v130
							return v9
						}
					} else {
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v25-int32(1)))))
						if v49 != 0 {
							v127 = int32(1024)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v127
							v130 = F_palloc(m, v127)
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v130
								return v9
							}
						} else {
							v51 = F_palloc(m, int32(48))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v51 != 0 {
									v53 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = v53
									*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v43
									*(*int64)(unsafe.Add(mBase, uint32(v51)+40)) = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = int64(4294967296)
									*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v43
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = v53
									F_core_yyensure_buffer_stack(m, v9)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v67+v68<<(uint(int32(2))%32))))
										if v72 == v51 {
										} else {
											if v72 != 0 {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
												*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v75)
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v79 = int32(2)
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v77+v78<<(uint(v79)%32))))
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
												*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v83
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v85+v86<<(uint(v79)%32))))
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
												*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v91
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v95 = v93
												v96 = v94
											} else {
												v95 = v67
												v96 = v68
											}
											v97 = int32(2)
											*(*int32)(unsafe.Add(mBase, uint32(v96<<(uint(v97)%32)+v95))) = v51
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											v105 = v101 + v102<<(uint(v97)%32)
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v107
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v110
											*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v110
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v114
											v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v116)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(1)
										}
										v127 = int32(1024)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v127
										v130 = F_palloc(m, v127)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v130
											return v9
										}
									}
								} else {
									F_yy_fatal_error_2(m, int32(_a_F_scanner_init_0))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
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
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_scanner_init[1])) = int32(48)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v142 = m.ExcPending
			if v142 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_scanner_init_1), int32(0))
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_scanner_init_2), int32(1233), int32(_a_F_scanner_init_3))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
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
func F_schedule_alarm(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v96 int32
	_ = v96
	var v97 float64
	_ = v97
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v106 float64
	_ = v106
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[0]))
	if v14 <= int32(0) {
		m.G0 = v11 + int32(48)
		return
	} else {
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v17
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[1]))
		if v26 == int32(0) {
		} else {
			v30 = *(*int64)(unsafe.Add(mBase, _c_F_schedule_alarm[2]))
			if l0 <= v30+int64(10000) {
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[1])) = int32(0)
			}
		}
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[3]))
		v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
		if v39 < l0 {
			v41 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v41
			*(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[1])) = v41
			v73 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v73
			v77 = int32(0)
			v78 = v73
		} else {
			v52 = v39 - l0
			if v52 <= int64(0) {
				v64 = int32(0)
				v65 = int32(0)
			} else {
				v56 = int64(1000000)
				v57 = base.I64_div_u_s(v52, v56)
				v64 = base.I32_wrap_i64(v57)
				v65 = base.I32_wrap_i64(v52 - v57*v56)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11+int32(12)))) = v64
			*(*int32)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v65
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			if v68|v69 != 0 {
				v77 = v68
				v78 = v69
			} else {
				v73 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v73
				v77 = int32(0)
				v78 = v73
			}
		}
		*(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[4])) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v78
		*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = base.I64_extend_i32_s(v77)
		v86 = *(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[1]))
		if v86 != 0 {
			v88 = *(*int64)(unsafe.Add(mBase, _c_F_schedule_alarm[2]))
			if v88 <= v39 {
				m.G0 = v11 + int32(48)
				return
			} else {
				*(*int64)(unsafe.Add(mBase, _c_F_schedule_alarm[2])) = v39
				*(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[1])) = int32(1)
				v96 = v11 + int32(16)
				v97 = m.Env.Emscripten_get_now(m)
				mBase = m.M
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
				v102 = *(*int64)(unsafe.Add(mBase, uint32(v96)+16))
				v106 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v98), float64(1000)), base.F64_convert_i64_s(v102*int64(1000)))
				if v102 == int64(0) {
					if v98 == int32(0) {
						v123 = float64(0)
						v124 = float64(0)
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
						v117 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
						v123 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v113), float64(1000)), base.F64_convert_i64_s(v117*int64(1000)))
						v124 = base.F64_add(v97, v106)
					}
				} else {
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
					v123 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v113), float64(1000)), base.F64_convert_i64_s(v117*int64(1000)))
					v124 = base.F64_add(v97, v106)
				}
				*(*float64)(unsafe.Add(mBase, _c_F_schedule_alarm[5])) = v124
				*(*float64)(unsafe.Add(mBase, _c_F_schedule_alarm[6])) = v123
				v129 = m.Env.X_setitimer_js(m, int32(0), v106)
				mBase = m.M
				if v129 != 0 {
					v141 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[1])) = v141
					F_errstart_cold(m, int32(22), v141)
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_schedule_alarm_0), int32(0))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_schedule_alarm_1), int32(347), int32(_a_F_schedule_alarm_2))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					m.G0 = v11 + int32(48)
					return
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, _c_F_schedule_alarm[2])) = v39
			*(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[1])) = int32(1)
			v96 = v11 + int32(16)
			v97 = m.Env.Emscripten_get_now(m)
			mBase = m.M
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
			v102 = *(*int64)(unsafe.Add(mBase, uint32(v96)+16))
			v106 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v98), float64(1000)), base.F64_convert_i64_s(v102*int64(1000)))
			if v102 == int64(0) {
				if v98 == int32(0) {
					v123 = float64(0)
					v124 = float64(0)
				} else {
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
					v123 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v113), float64(1000)), base.F64_convert_i64_s(v117*int64(1000)))
					v124 = base.F64_add(v97, v106)
				}
			} else {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
				v117 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
				v123 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v113), float64(1000)), base.F64_convert_i64_s(v117*int64(1000)))
				v124 = base.F64_add(v97, v106)
			}
			*(*float64)(unsafe.Add(mBase, _c_F_schedule_alarm[5])) = v124
			*(*float64)(unsafe.Add(mBase, _c_F_schedule_alarm[6])) = v123
			v129 = m.Env.X_setitimer_js(m, int32(0), v106)
			mBase = m.M
			if v129 != 0 {
				v141 = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_schedule_alarm[1])) = v141
				F_errstart_cold(m, int32(22), v141)
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_schedule_alarm_0), int32(0))
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_schedule_alarm_1), int32(347), int32(_a_F_schedule_alarm_2))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v11 + int32(48)
				return
			}
		}
	}
}
func F_search_indexed_tlist_for_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 <= v6 {
		v249 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L9
	} else {
		goto L60
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return v249
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v25 = l1
	v31 = v17
	goto L4
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v36 != v20 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v249 = v6
	goto L2
L6:
	;
	v240 = int32(1)
	if base.Ui32(v240) < base.Ui32(v31) {
		v25 = v25 + int32(12)
		v31 = v31 - v240
		goto L4
	} else {
		goto L59
	}
L7:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+16)))
	if v38 != v21&int32(_a_F_search_indexed_tlist_for_var_0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v41 = F_palloc(m, int32(48))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+40)) = v45
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+32)) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+24)) = v49
	v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+16)) = v51
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = v53
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v55
	if v21 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = l2
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+8)) = uint16(v231)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	if v233 == int32(0) {
		v249 = v41
		goto L2
	} else {
		goto L58
	}
L12:
	;
	switch l4 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L15
	default:
		goto L14
	}
L13:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v174 = int32(0)
	if v172 == v174 {
		goto L44
	} else {
		goto L45
	}
L14:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v121 = int32(0)
	if base.B2i32(v119 == v121)|base.B2i32(v120 == v121) != 0 {
		v167 = base.B2i32(v119|v120 == v121)
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v63 = int32(0)
	if v61 == v63 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v116 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L17:
	;
	v116 = int32(1)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v62 == int32(0) {
		v109 = v63
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v116 = v109
	goto L16
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v73 < v72 {
		v109 = v63
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v75 = int32(1)
	if v72 <= v75 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = v75
	goto L25
L24:
	;
	v78 = v72
	goto L25
L25:
	;
	v79 = int32(8)
	v84 = int32(0)
	goto L26
L26:
	;
	v91 = v84 << (uint(int32(2)) % 32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v61+v79+v91)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v62+v79+v91)))
	v98 = v93 & (v95 ^ int32(-1))
	v100 = base.B2i32(v98 == int32(0))
	if v98 != 0 {
		v109 = v100
		goto L20
	} else {
		goto L28
	}
L27:
	;
	v109 = v100
	goto L20
L28:
	;
	v102 = v84 + int32(1)
	if v102 != v78 {
		v84 = v102
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L11
L31:
	;
	if v167 != 0 {
		goto L11
	} else {
		goto L42
	}
L32:
	;
	goto L31
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v135 != v136 {
		v167 = int32(0)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v138 = int32(1)
	if v135 <= v138 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v141 = v138
	goto L37
L36:
	;
	v141 = v135
	goto L37
L37:
	;
	v142 = int32(8)
	v147 = int32(0)
	goto L38
L38:
	;
	v155 = v147 << (uint(int32(2)) % 32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v119+v142+v155)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v120+v142+v155)))
	v160 = base.B2i32(v157 == v159)
	if v157 != v159 {
		v167 = v160
		goto L32
	} else {
		goto L40
	}
L39:
	;
	v167 = v160
	goto L32
L40:
	;
	v163 = v147 + int32(1)
	if v163 != v141 {
		v147 = v163
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L1
L43:
	;
	if v227 == int32(0) {
		goto L1
	} else {
		goto L57
	}
L44:
	;
	v227 = int32(1)
	goto L43
L45:
	;
	goto L46
L46:
	;
	if v173 == int32(0) {
		v220 = v174
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v227 = v220
	goto L43
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v184 < v183 {
		v220 = v174
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v186 = int32(1)
	if v183 <= v186 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v189 = v186
	goto L52
L51:
	;
	v189 = v183
	goto L52
L52:
	;
	v190 = int32(8)
	v195 = int32(0)
	goto L53
L53:
	;
	v202 = v195 << (uint(int32(2)) % 32)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v172+v190+v202)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v173+v190+v202)))
	v209 = v204 & (v206 ^ int32(-1))
	v211 = base.B2i32(v209 == int32(0))
	if v209 != 0 {
		v220 = v211
		goto L47
	} else {
		goto L55
	}
L54:
	;
	v220 = v211
	goto L47
L55:
	;
	v213 = v195 + int32(1)
	if v213 != v189 {
		v195 = v213
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	goto L11
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v233 + l3
	v249 = v41
	goto L2
L59:
	;
	goto L5
L60:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v265 = F_bmsToString(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v268 = F_bmsToString(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v265
	F_errmsg_internal(m, int32(_a_F_search_indexed_tlist_for_var_1), v15)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_search_indexed_tlist_for_var_2), int32(2987), int32(_a_F_search_indexed_tlist_for_var_3))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_search_plan_tree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	v4 = int32(0)
	if l0 == v4 {
		v249 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v249
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v11 - int32(400) {
	case 0, 43:
		v131 = int32(36)
		goto L6
	default:
		v249 = v4
		goto L1
	case 3:
		goto L8
	case 9, 10, 11, 12, 14, 15, 16, 24, 25:
		goto L4
	case 17:
		goto L7
	}
L3:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v240 != 0 {
		goto L81
	} else {
		goto L82
	}
L4:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v228 == int32(0) {
		v249 = v4
		goto L1
	} else {
		goto L79
	}
L5:
	;
	if v222 == int32(0) {
		v249 = v4
		goto L1
	} else {
		goto L78
	}
L6:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0+v131)))
	v134 = int32(0)
	if v133 == v134 {
		v210 = v134
		goto L51
	} else {
		goto L52
	}
L7:
	;
	v131 = int32(116)
	goto L6
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v14 <= int32(0) {
		v249 = v4
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v21 = int32(0)
	v22 = v4
	goto L10
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v22<<(uint(int32(2))%32))))
	v30 = int32(0)
	if v29 == v30 {
		v106 = v30
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v222 = v125
	goto L5
L12:
	;
	v115 = int32(0)
	if base.B2i32(v114 == v115)|base.B2i32(v21 == v115) == v115 {
		goto L40
	} else {
		goto L41
	}
L13:
	;
	v114 = v106
	goto L12
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	switch v37 - int32(400) {
	case 0, 43:
		v72 = int32(36)
		goto L18
	default:
		v106 = v30
		goto L13
	case 3:
		goto L20
	case 9, 10, 11, 12, 14, 15, 16, 24, 25:
		goto L16
	case 17:
		goto L19
	}
L15:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v97 != 0 {
		goto L37
	} else {
		goto L38
	}
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v29)+104))
	if v85 == int32(0) {
		v106 = v30
		goto L13
	} else {
		goto L35
	}
L17:
	;
	if v79 == int32(0) {
		v106 = v30
		goto L13
	} else {
		goto L34
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v29+v72)))
	v75 = F_search_plan_tree(m, v74, l1, l2)
	mBase = m.M
	v79 = v75
	goto L17
L19:
	;
	v72 = int32(116)
	goto L18
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	if v40 <= int32(0) {
		v106 = v30
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v47 = int32(0)
	v48 = v30
	goto L22
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v29)+104))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v48<<(uint(int32(2))%32))))
	v56 = F_search_plan_tree(m, v55, l1, l2)
	mBase = m.M
	v57 = int32(0)
	if base.B2i32(v56 == v57)|base.B2i32(v47 == v57) == v57 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v79 = v66
	goto L17
L24:
	;
	v114 = int32(0)
	goto L12
L25:
	;
	goto L26
L26:
	;
	if v47 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v65 = v47
	goto L29
L28:
	;
	v65 = v56
	goto L29
L29:
	;
	if v56 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v66 = v65
	goto L32
L31:
	;
	v66 = v47
	goto L32
L32:
	;
	v68 = v48 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	if v68 < v69 {
		v47 = v66
		v48 = v68
		goto L22
	} else {
		goto L33
	}
L33:
	;
	goto L23
L34:
	;
	v93 = v79
	goto L15
L35:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+56))
	if v88 != l1 {
		v106 = v30
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v93 = v29
	goto L15
L37:
	;
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v98)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v106 = v93
	goto L13
L40:
	;
	return int32(0)
L41:
	;
	goto L42
L42:
	;
	if v21 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v124 = v21
	goto L45
L44:
	;
	v124 = v114
	goto L45
L45:
	;
	if v114 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v125 = v124
	goto L48
L47:
	;
	v125 = v21
	goto L48
L48:
	;
	v127 = v22 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v127 < v128 {
		v21 = v125
		v22 = v127
		goto L10
	} else {
		goto L49
	}
L49:
	;
	goto L11
L50:
	;
	v222 = v218
	goto L5
L51:
	;
	v218 = v210
	goto L50
L52:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	switch v141 - int32(400) {
	case 0, 43:
		v176 = int32(36)
		goto L56
	default:
		v210 = v134
		goto L51
	case 3:
		goto L58
	case 9, 10, 11, 12, 14, 15, 16, 24, 25:
		goto L54
	case 17:
		goto L57
	}
L53:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	if v201 != 0 {
		goto L75
	} else {
		goto L76
	}
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v133)+104))
	if v189 == int32(0) {
		v210 = v134
		goto L51
	} else {
		goto L73
	}
L55:
	;
	if v183 == int32(0) {
		v210 = v134
		goto L51
	} else {
		goto L72
	}
L56:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v133+v176)))
	v179 = F_search_plan_tree(m, v178, l1, l2)
	mBase = m.M
	v183 = v179
	goto L55
L57:
	;
	v176 = int32(116)
	goto L56
L58:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v133)+108))
	if v144 <= int32(0) {
		v210 = v134
		goto L51
	} else {
		goto L59
	}
L59:
	;
	v151 = int32(0)
	v152 = v134
	goto L60
L60:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v133)+104))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155+v152<<(uint(int32(2))%32))))
	v160 = F_search_plan_tree(m, v159, l1, l2)
	mBase = m.M
	v161 = int32(0)
	if base.B2i32(v160 == v161)|base.B2i32(v151 == v161) == v161 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v183 = v170
	goto L55
L62:
	;
	v218 = int32(0)
	goto L50
L63:
	;
	goto L64
L64:
	;
	if v151 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v169 = v151
	goto L67
L66:
	;
	v169 = v160
	goto L67
L67:
	;
	if v160 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v170 = v169
	goto L70
L69:
	;
	v170 = v151
	goto L70
L70:
	;
	v172 = v152 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v133)+108))
	if v172 < v173 {
		v151 = v170
		v152 = v172
		goto L60
	} else {
		goto L71
	}
L71:
	;
	goto L61
L72:
	;
	v197 = v183
	goto L53
L73:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+56))
	if v192 != l1 {
		v210 = v134
		goto L51
	} else {
		goto L74
	}
L74:
	;
	v197 = v133
	goto L53
L75:
	;
	v202 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v202)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v210 = v197
	goto L51
L78:
	;
	v236 = v222
	goto L3
L79:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)+56))
	if v231 != l1 {
		v249 = v4
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v236 = l0
	goto L3
L81:
	;
	v241 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v241)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v249 = v236
	goto L1
}
func F_self_join_candidates_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return base.B2i32(base.Ui32(v4) < base.Ui32(v3)) - base.B2i32(base.Ui32(v3) < base.Ui32(v4))
}
func F_sendAuthRequest(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
	if v11 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_pq_beginmessage(m, v8, int32(82))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_enlargeStringInfo(m, v8, int32(4))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v25 = int32(16711935)
					*(*int32)(unsafe.Add(mBase, uint32(v20+v21))) = base.I32_rotr(l0, int32(24))&v25 | base.I32_rotr(l0&v25, int32(8))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20 + int32(4)
					if int32(0) < l2 {
						F_appendBinaryStringInfo(m, v8, l1, l2)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							F_pq_endmessage(m, v8)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = int32(0)
								if base.B2i32(l0 == v42)|base.B2i32(l0 == int32(12)) == v42 {
									v50 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[1]))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
									v52 = m.T0[v51].(func(*base.Module) int32)(m)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
										if v55 != 0 {
											F_ProcessInterrupts(m)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
												return
											}
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
									if v55 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						F_pq_endmessage(m, v8)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v42 = int32(0)
							if base.B2i32(l0 == v42)|base.B2i32(l0 == int32(12)) == v42 {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[1]))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
								v52 = m.T0[v51].(func(*base.Module) int32)(m)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
									if v55 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
								if v55 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_pq_beginmessage(m, v8, int32(82))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_enlargeStringInfo(m, v8, int32(4))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v25 = int32(16711935)
				*(*int32)(unsafe.Add(mBase, uint32(v20+v21))) = base.I32_rotr(l0, int32(24))&v25 | base.I32_rotr(l0&v25, int32(8))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20 + int32(4)
				if int32(0) < l2 {
					F_appendBinaryStringInfo(m, v8, l1, l2)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_pq_endmessage(m, v8)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v42 = int32(0)
							if base.B2i32(l0 == v42)|base.B2i32(l0 == int32(12)) == v42 {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[1]))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
								v52 = m.T0[v51].(func(*base.Module) int32)(m)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
									if v55 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
								if v55 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_pq_endmessage(m, v8)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v42 = int32(0)
						if base.B2i32(l0 == v42)|base.B2i32(l0 == int32(12)) == v42 {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[1]))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							v52 = m.T0[v51].(func(*base.Module) int32)(m)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
								if v55 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_sendAuthRequest[0]))
							if v55 != 0 {
								F_ProcessInterrupts(m)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_setenv(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v952 int32
	_ = v952
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v105) {
		goto L52
	} else {
		goto L53
	}
L2:
	;
	return int32(-1)
L3:
	;
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_setenv[0])) = int32(28)
	goto L2
L5:
	;
	goto L10
L6:
	;
	if v93 == l0 {
		goto L4
	} else {
		goto L29
	}
L7:
	;
	goto L6
L8:
	;
	v83 = v78
	goto L25
L9:
	;
	v78 = v70
	goto L8
L10:
	;
	if l0&int32(3) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v16 = l0
	goto L16
L14:
	;
	v30 = l0
	goto L15
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 != v39 {
		v70 = v30
		goto L9
	} else {
		goto L20
	}
L16:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if base.B2i32(v21 == int32(0))|base.B2i32(int32(61) == v21) != 0 {
		v93 = v16
		goto L7
	} else {
		goto L18
	}
L17:
	;
	v30 = v27
	goto L15
L18:
	;
	v27 = v16 + int32(1)
	if v27&int32(3) != 0 {
		v16 = v27
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v45 = v30
	v47 = v36
	goto L21
L21:
	;
	v51 = v47 ^ int32(1027423549)
	v54 = int32(-2139062144)
	if (int32(16843008)-v51|v51)&v54 != v54 {
		v70 = v45
		goto L9
	} else {
		goto L23
	}
L22:
	;
	v78 = v60
	goto L8
L23:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v60 = v45 + int32(4)
	v64 = int32(-2139062144)
	if (v58|(int32(16843008)-v58))&v64 == v64 {
		v45 = v60
		v47 = v58
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v85 == int32(0) {
		v93 = v83
		goto L7
	} else {
		goto L27
	}
L26:
	;
	v93 = v83
	goto L7
L27:
	;
	if v85 != int32(61) {
		v83 = v83 + int32(1)
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v105 = v93 - l0
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v105))))
	if v107 == int32(0) {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	goto L4
L31:
	;
	v166 = F_strlen(m, l1)
	mBase = m.M
	v170 = F_emscripten_builtin_malloc(m, v105+v166+int32(2))
	mBase = m.M
	if v170 != 0 {
		goto L1
	} else {
		goto L50
	}
L32:
	;
	v114 = int32(0)
	v119 = F___strchrnul(m, l0, int32(61))
	mBase = m.M
	if l0 == v119 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v161 == int32(0) {
		goto L31
	} else {
		goto L49
	}
L34:
	;
	v161 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v122 = v119 - l0
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v122))))
	if v124 != 0 {
		v155 = v114
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v161 = v155
	goto L33
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[1]))
	if v126 == int32(0) {
		v155 = v114
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v129 == int32(0) {
		v155 = v114
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v133 = v126
	v134 = v129
	goto L41
L41:
	;
	v137 = F_strncmp(m, l0, v134, v122)
	mBase = m.M
	if v137 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v155 = v141 + int32(1)
	goto L37
L43:
	;
	goto L42
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v141 = v140 + v122
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v142 == int32(61) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v146 != 0 {
		v133 = v133 + int32(4)
		v134 = v146
		goto L41
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v155 = v114
	goto L37
L49:
	;
	return int32(0)
L50:
	;
	goto L2
L51:
	;
	v345 = v170 + v105
	v346 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v346)
	v348 = int32(1)
	v349 = v345 + v348
	v351 = v166 + v348
	if base.Ui32(int32(512)) <= base.Ui32(v351) {
		goto L99
	} else {
		goto L100
	}
L52:
	;
	if v105 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v182 = v170 + v105
	if (v170^l0)&int32(3) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	base.MemoryCopy(m, v170, l0, v105)
	goto L57
L56:
	;
	goto L57
L57:
	;
	goto L51
L58:
	;
	if base.Ui32(v314) < base.Ui32(v182) {
		goto L92
	} else {
		goto L93
	}
L59:
	;
	if v170&int32(3) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	goto L61
L61:
	;
	if base.Ui32(v182) < base.Ui32(int32(4)) {
		goto L83
	} else {
		goto L84
	}
L62:
	;
	v218 = v182 & int32(-4)
	if base.Ui32(v182) < base.Ui32(int32(64)) {
		v268 = v212
		v269 = v213
		goto L73
	} else {
		goto L74
	}
L63:
	;
	v212 = l0
	v213 = v170
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v105 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v212 = l0
	v213 = v170
	goto L62
L67:
	;
	goto L68
L68:
	;
	v195 = l0
	v196 = v170
	goto L69
L69:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v200)
	v202 = int32(1)
	v203 = v195 + v202
	v205 = v196 + v202
	if v205&int32(3) == int32(0) {
		v212 = v203
		v213 = v205
		goto L62
	} else {
		goto L71
	}
L70:
	;
	v212 = v203
	v213 = v205
	goto L62
L71:
	;
	if base.Ui32(v205) < base.Ui32(v182) {
		v195 = v203
		v196 = v205
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	if base.Ui32(v218) <= base.Ui32(v269) {
		v313 = v268
		v314 = v269
		goto L58
	} else {
		goto L79
	}
L74:
	;
	v222 = v218 + int32(-64)
	if base.Ui32(v222) < base.Ui32(v213) {
		v268 = v212
		v269 = v213
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v225 = v212
	v226 = v213
	goto L76
L76:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+4)) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+8)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+12)) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+16)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v225)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+20)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v225)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+24)) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v225)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+28)) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v225)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+32)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v225)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+36)) = v248
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v225)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+40)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v225)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+44)) = v252
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v225)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+48)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v225)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+52)) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v225)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+56)) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v225)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+60)) = v260
	v262 = int32(-64)
	v263 = v225 - v262
	v265 = v226 - v262
	if base.Ui32(v265) <= base.Ui32(v222) {
		v225 = v263
		v226 = v265
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v268 = v263
	v269 = v265
	goto L73
L78:
	;
	goto L77
L79:
	;
	v275 = v268
	v276 = v269
	goto L80
L80:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v280
	v282 = int32(4)
	v283 = v275 + v282
	v285 = v276 + v282
	if base.Ui32(v285) < base.Ui32(v218) {
		v275 = v283
		v276 = v285
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v313 = v283
	v314 = v285
	goto L58
L82:
	;
	goto L81
L83:
	;
	v313 = l0
	v314 = v170
	goto L58
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(v105) < base.Ui32(int32(4)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v313 = l0
	v314 = v170
	goto L58
L87:
	;
	goto L88
L88:
	;
	v294 = l0
	v295 = v170
	goto L89
L89:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	*(*uint8)(unsafe.Add(mBase, uint32(v295))) = uint8(v299)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)) = uint8(v301)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+2)) = uint8(v303)
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+3)) = uint8(v305)
	v307 = int32(4)
	v308 = v294 + v307
	v310 = v295 + v307
	if base.Ui32(v310) <= base.Ui32(v182-int32(4)) {
		v294 = v308
		v295 = v310
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v313 = v308
	v314 = v310
	goto L58
L91:
	;
	goto L90
L92:
	;
	v320 = v313
	v321 = v314
	goto L95
L93:
	;
	goto L94
L94:
	;
	goto L51
L95:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	*(*uint8)(unsafe.Add(mBase, uint32(v321))) = uint8(v325)
	v327 = int32(1)
	v330 = v321 + v327
	if v330 != v182 {
		v320 = v320 + v327
		v321 = v330
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L94
L97:
	;
	goto L96
L98:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[1]))
	if v522 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L99:
	;
	if v351 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v358 = v349 + v351
	if (v349^l1)&int32(3) == int32(0) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	base.MemoryCopy(m, v349, l1, v351)
	goto L104
L103:
	;
	goto L104
L104:
	;
	goto L98
L105:
	;
	if base.Ui32(v490) < base.Ui32(v358) {
		goto L139
	} else {
		goto L140
	}
L106:
	;
	if v349&int32(3) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	if base.Ui32(v358) < base.Ui32(int32(4)) {
		goto L130
	} else {
		goto L131
	}
L109:
	;
	v394 = v358 & int32(-4)
	if base.Ui32(v358) < base.Ui32(int32(64)) {
		v444 = v388
		v445 = v389
		goto L120
	} else {
		goto L121
	}
L110:
	;
	v388 = l1
	v389 = v349
	goto L109
L111:
	;
	goto L112
L112:
	;
	if v351 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v388 = l1
	v389 = v349
	goto L109
L114:
	;
	goto L115
L115:
	;
	v371 = l1
	v372 = v349
	goto L116
L116:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	*(*uint8)(unsafe.Add(mBase, uint32(v372))) = uint8(v376)
	v378 = int32(1)
	v379 = v371 + v378
	v381 = v372 + v378
	if v381&int32(3) == int32(0) {
		v388 = v379
		v389 = v381
		goto L109
	} else {
		goto L118
	}
L117:
	;
	v388 = v379
	v389 = v381
	goto L109
L118:
	;
	if base.Ui32(v381) < base.Ui32(v358) {
		v371 = v379
		v372 = v381
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	if base.Ui32(v394) <= base.Ui32(v445) {
		v489 = v444
		v490 = v445
		goto L105
	} else {
		goto L126
	}
L121:
	;
	v398 = v394 + int32(-64)
	if base.Ui32(v398) < base.Ui32(v389) {
		v444 = v388
		v445 = v389
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v401 = v388
	v402 = v389
	goto L123
L123:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+4)) = v408
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v401)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+8)) = v410
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+12)) = v412
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v401)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+16)) = v414
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v401)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+20)) = v416
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v401)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+24)) = v418
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v401)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+28)) = v420
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v401)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+32)) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v401)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+36)) = v424
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v401)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+40)) = v426
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v401)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+44)) = v428
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v401)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+48)) = v430
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v401)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+52)) = v432
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v401)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+56)) = v434
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v401)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+60)) = v436
	v438 = int32(-64)
	v439 = v401 - v438
	v441 = v402 - v438
	if base.Ui32(v441) <= base.Ui32(v398) {
		v401 = v439
		v402 = v441
		goto L123
	} else {
		goto L125
	}
L124:
	;
	v444 = v439
	v445 = v441
	goto L120
L125:
	;
	goto L124
L126:
	;
	v451 = v444
	v452 = v445
	goto L127
L127:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = v456
	v458 = int32(4)
	v459 = v451 + v458
	v461 = v452 + v458
	if base.Ui32(v461) < base.Ui32(v394) {
		v451 = v459
		v452 = v461
		goto L127
	} else {
		goto L129
	}
L128:
	;
	v489 = v459
	v490 = v461
	goto L105
L129:
	;
	goto L128
L130:
	;
	v489 = l1
	v490 = v349
	goto L105
L131:
	;
	goto L132
L132:
	;
	if base.Ui32(v351) < base.Ui32(int32(4)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v489 = l1
	v490 = v349
	goto L105
L134:
	;
	goto L135
L135:
	;
	v470 = l1
	v471 = v349
	goto L136
L136:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	*(*uint8)(unsafe.Add(mBase, uint32(v471))) = uint8(v475)
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v471)+1)) = uint8(v477)
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v471)+2)) = uint8(v479)
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v471)+3)) = uint8(v481)
	v483 = int32(4)
	v484 = v470 + v483
	v486 = v471 + v483
	if base.Ui32(v486) <= base.Ui32(v358-int32(4)) {
		v470 = v484
		v471 = v486
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v489 = v484
	v490 = v486
	goto L105
L138:
	;
	goto L137
L139:
	;
	v496 = v489
	v497 = v490
	goto L142
L140:
	;
	goto L141
L141:
	;
	goto L98
L142:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	*(*uint8)(unsafe.Add(mBase, uint32(v497))) = uint8(v501)
	v503 = int32(1)
	v506 = v497 + v503
	if v506 != v358 {
		v496 = v496 + v503
		v497 = v506
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	goto L143
L145:
	;
	return v952
L146:
	;
	v674 = v669 << (uint(int32(2)) % 32)
	v676 = v674 + int32(8)
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[2]))
	if v668 == v678 {
		goto L190
	} else {
		goto L191
	}
L147:
	;
	v531 = v105 + int32(1)
	v533 = v522
	v534 = int32(0)
	v537 = v526
	goto L153
L148:
	;
	v668 = v527
	v669 = int32(0)
	goto L146
L149:
	;
	v527 = int32(0)
	goto L148
L150:
	;
	goto L151
L151:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	if v526 != 0 {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	v527 = v522
	goto L148
L153:
	;
	if v531 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[1]))
	v668 = v667
	v669 = v662
	goto L146
L155:
	;
	if v582 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L156:
	;
	v582 = int32(0)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v543 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v544 = v170
	v545 = v537
	v546 = v531
	v547 = v543
	goto L163
L160:
	;
	v570 = v537
	v574 = int32(0)
	goto L161
L161:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	v582 = v574 - v575
	goto L155
L162:
	;
	v570 = v565
	v574 = v567
	goto L161
L163:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	if base.B2i32(v547 != v549)|base.B2i32(v549 == int32(0)) != 0 {
		v565 = v545
		v567 = v547
		goto L162
	} else {
		goto L165
	}
L164:
	;
	v565 = v559
	v567 = int32(0)
	goto L162
L165:
	;
	v555 = v546 - int32(1)
	if v555 == int32(0) {
		v565 = v545
		v567 = v547
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v558 = int32(1)
	v559 = v545 + v558
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+1)))
	if v560 != 0 {
		v544 = v544 + v558
		v545 = v559
		v546 = v555
		v547 = v560
		goto L163
	} else {
		goto L167
	}
L167:
	;
	goto L164
L168:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v170
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[3]))
	if v593 != 0 {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	goto L170
L170:
	;
	v662 = v534 + int32(1)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v663 != 0 {
		v533 = v533 + int32(4)
		v534 = v662
		v537 = v663
		goto L153
	} else {
		goto L187
	}
L171:
	;
	v952 = int32(0)
	goto L145
L172:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[4]))
	v597 = v170
	v599 = int32(0)
	goto L175
L173:
	;
	v622 = v170
	goto L174
L174:
	;
	if v622 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L175:
	;
	v605 = v595 + v599<<(uint(int32(2))%32)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	if v585 == v606 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v622 = v617
	goto L174
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v605))) = v597
	F_emscripten_builtin_free(m, v585)
	mBase = m.M
	goto L171
L178:
	;
	goto L179
L179:
	;
	v610 = int32(0)
	if v606|base.B2i32(v597 == v610) == v610 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v605))) = v597
	v617 = int32(0)
	goto L182
L181:
	;
	v617 = v597
	goto L182
L182:
	;
	v619 = v599 + int32(1)
	if v619 != v593 {
		v597 = v617
		v599 = v619
		goto L175
	} else {
		goto L183
	}
L183:
	;
	goto L176
L184:
	;
	goto L171
L185:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[4]))
	v636 = F_emscripten_builtin_realloc(m, v631, v593<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	if v636 == int32(0) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_setenv[4])) = v636
	v641 = int32(_a_F_setenv_0)
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_setenv[3])) = v643 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v636+v643<<(uint(int32(2))%32)))) = v622
	goto L184
L187:
	;
	goto L154
L188:
	;
	F_emscripten_builtin_free(m, v170)
	mBase = m.M
	v952 = int32(-1)
	goto L145
L189:
	;
	v861 = v858 + v669<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v861))) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v861)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_setenv[1])) = v858
	*(*int32)(unsafe.Add(mBase, _c_F_setenv[2])) = v858
	if v170 != 0 {
		goto L245
	} else {
		goto L246
	}
L190:
	;
	v680 = F_emscripten_builtin_realloc(m, v678, v676)
	mBase = m.M
	if v680 != 0 {
		v858 = v680
		goto L189
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v681 = F_emscripten_builtin_malloc(m, v676)
	mBase = m.M
	if v681 == int32(0) {
		goto L188
	} else {
		goto L194
	}
L193:
	;
	goto L188
L194:
	;
	if v669 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[1]))
	if base.Ui32(int32(512)) <= base.Ui32(v674) {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	goto L197
L197:
	;
	v856 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[2]))
	F_emscripten_builtin_free(m, v856)
	mBase = m.M
	v858 = v681
	goto L189
L198:
	;
	goto L197
L199:
	;
	if v674 != 0 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	v692 = v681 + v674
	if (v681^v685)&int32(3) == int32(0) {
		goto L206
	} else {
		goto L207
	}
L202:
	;
	base.MemoryCopy(m, v681, v685, v674)
	goto L204
L203:
	;
	goto L204
L204:
	;
	goto L198
L205:
	;
	if base.Ui32(v824) < base.Ui32(v692) {
		goto L239
	} else {
		goto L240
	}
L206:
	;
	if v681&int32(3) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L207:
	;
	goto L208
L208:
	;
	if base.Ui32(v692) < base.Ui32(int32(4)) {
		goto L230
	} else {
		goto L231
	}
L209:
	;
	v728 = v692 & int32(-4)
	if base.Ui32(v692) < base.Ui32(int32(64)) {
		v778 = v722
		v779 = v723
		goto L220
	} else {
		goto L221
	}
L210:
	;
	v722 = v685
	v723 = v681
	goto L209
L211:
	;
	goto L212
L212:
	;
	if v674 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v722 = v685
	v723 = v681
	goto L209
L214:
	;
	goto L215
L215:
	;
	v705 = v685
	v706 = v681
	goto L216
L216:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	*(*uint8)(unsafe.Add(mBase, uint32(v706))) = uint8(v710)
	v712 = int32(1)
	v713 = v705 + v712
	v715 = v706 + v712
	if v715&int32(3) == int32(0) {
		v722 = v713
		v723 = v715
		goto L209
	} else {
		goto L218
	}
L217:
	;
	v722 = v713
	v723 = v715
	goto L209
L218:
	;
	if base.Ui32(v715) < base.Ui32(v692) {
		v705 = v713
		v706 = v715
		goto L216
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	if base.Ui32(v728) <= base.Ui32(v779) {
		v823 = v778
		v824 = v779
		goto L205
	} else {
		goto L226
	}
L221:
	;
	v732 = v728 + int32(-64)
	if base.Ui32(v732) < base.Ui32(v723) {
		v778 = v722
		v779 = v723
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v735 = v722
	v736 = v723
	goto L223
L223:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v735)))
	*(*int32)(unsafe.Add(mBase, uint32(v736))) = v740
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v735)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+4)) = v742
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v735)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+8)) = v744
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v735)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+12)) = v746
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v735)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+16)) = v748
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v735)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+20)) = v750
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v735)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+24)) = v752
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v735)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+28)) = v754
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v735)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+32)) = v756
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v735)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+36)) = v758
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v735)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+40)) = v760
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v735)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+44)) = v762
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v735)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+48)) = v764
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v735)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+52)) = v766
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v735)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+56)) = v768
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v735)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+60)) = v770
	v772 = int32(-64)
	v773 = v735 - v772
	v775 = v736 - v772
	if base.Ui32(v775) <= base.Ui32(v732) {
		v735 = v773
		v736 = v775
		goto L223
	} else {
		goto L225
	}
L224:
	;
	v778 = v773
	v779 = v775
	goto L220
L225:
	;
	goto L224
L226:
	;
	v785 = v778
	v786 = v779
	goto L227
L227:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	*(*int32)(unsafe.Add(mBase, uint32(v786))) = v790
	v792 = int32(4)
	v793 = v785 + v792
	v795 = v786 + v792
	if base.Ui32(v795) < base.Ui32(v728) {
		v785 = v793
		v786 = v795
		goto L227
	} else {
		goto L229
	}
L228:
	;
	v823 = v793
	v824 = v795
	goto L205
L229:
	;
	goto L228
L230:
	;
	v823 = v685
	v824 = v681
	goto L205
L231:
	;
	goto L232
L232:
	;
	if base.Ui32(v674) < base.Ui32(int32(4)) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v823 = v685
	v824 = v681
	goto L205
L234:
	;
	goto L235
L235:
	;
	v804 = v685
	v805 = v681
	goto L236
L236:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	*(*uint8)(unsafe.Add(mBase, uint32(v805))) = uint8(v809)
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v805)+1)) = uint8(v811)
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v805)+2)) = uint8(v813)
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v805)+3)) = uint8(v815)
	v817 = int32(4)
	v818 = v804 + v817
	v820 = v805 + v817
	if base.Ui32(v820) <= base.Ui32(v692-int32(4)) {
		v804 = v818
		v805 = v820
		goto L236
	} else {
		goto L238
	}
L237:
	;
	v823 = v818
	v824 = v820
	goto L205
L238:
	;
	goto L237
L239:
	;
	v830 = v823
	v831 = v824
	goto L242
L240:
	;
	goto L241
L241:
	;
	goto L198
L242:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830))))
	*(*uint8)(unsafe.Add(mBase, uint32(v831))) = uint8(v835)
	v837 = int32(1)
	v840 = v831 + v837
	if v840 != v692 {
		v830 = v830 + v837
		v831 = v840
		goto L242
	} else {
		goto L244
	}
L243:
	;
	goto L241
L244:
	;
	goto L243
L245:
	;
	v869 = int32(0)
	v876 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[3]))
	if v876 != 0 {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	goto L247
L247:
	;
	v952 = int32(0)
	goto L145
L248:
	;
	goto L247
L249:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[4]))
	v880 = v170
	v882 = v869
	goto L252
L250:
	;
	v905 = v170
	goto L251
L251:
	;
	if v905 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L252:
	;
	v888 = v878 + v882<<(uint(int32(2))%32)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	if v869 == v889 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v905 = v900
	goto L251
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v888))) = v880
	F_emscripten_builtin_free(m, v869)
	mBase = m.M
	goto L248
L255:
	;
	goto L256
L256:
	;
	v893 = int32(0)
	if v889|base.B2i32(v880 == v893) == v893 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v888))) = v880
	v900 = int32(0)
	goto L259
L258:
	;
	v900 = v880
	goto L259
L259:
	;
	v902 = v882 + int32(1)
	if v902 != v876 {
		v880 = v900
		v882 = v902
		goto L252
	} else {
		goto L260
	}
L260:
	;
	goto L253
L261:
	;
	goto L248
L262:
	;
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[4]))
	v919 = F_emscripten_builtin_realloc(m, v914, v876<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	if v919 == int32(0) {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_setenv[4])) = v919
	v924 = int32(_a_F_setenv_0)
	v926 = *(*int32)(unsafe.Add(mBase, _c_F_setenv[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_setenv[3])) = v926 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v919+v926<<(uint(int32(2))%32)))) = v905
	goto L261
}
func F_setval3_oid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	F_SetSequence(m, v3, v4, base.B2i32(v5 != int64(0)))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_sha256_bytea(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_cryptohash_internal(m, int32(3), v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_shim_system(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = int32(_a_F_shim_system_0)
	v4 = m.Env.Pgmem_run(m, l0, v2, v2)
	return v4 << (uint(int32(8)) % 32) & int32(_a_F_shim_system_1)
}
func F_show_effective_wal_level(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_show_effective_wal_level[0]))
	if v4 == int32(0) {
		return int32(_a_F_show_effective_wal_level_0)
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_show_effective_wal_level[1])))
		if v10 != int32(1) {
			v45 = int32(_a_F_show_effective_wal_level_1)
			v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_show_effective_wal_level[2])))
			if v49&int32(1) != 0 {
				v52 = v45
			} else {
				v52 = int32(_a_F_show_effective_wal_level_2)
			}
			if int32(1) < v4 {
				v55 = v45
			} else {
				v55 = v52
			}
			return v55
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_show_effective_wal_level[3]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+308))
			v17 = int32(2)
			*(*uint8)(unsafe.Add(mBase, _c_F_show_effective_wal_level[1])) = uint8(base.B2i32(v16 != v17))
			if v16 == v17 {
				v45 = int32(_a_F_show_effective_wal_level_1)
				v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_show_effective_wal_level[2])))
				if v49&int32(1) != 0 {
					v52 = v45
				} else {
					v52 = int32(_a_F_show_effective_wal_level_2)
				}
				if int32(1) < v4 {
					v55 = v45
				} else {
					v55 = v52
				}
				return v55
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_show_effective_wal_level[4]))
				v27 = F_LWLockAcquire(m, v23+int32(_a_F_show_effective_wal_level_3), int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_show_effective_wal_level[5]))
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_show_effective_wal_level[4]))
					F_LWLockRelease(m, v35+int32(_a_F_show_effective_wal_level_3))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if v33 != 0 {
							v42 = int32(_a_F_show_effective_wal_level_1)
						} else {
							v42 = int32(_a_F_show_effective_wal_level_2)
						}
						return v42
					}
				}
			}
		}
	}
}
func F_show_random_seed(m *base.Module) int32 {
	return int32(_a_F_show_random_seed_0)
}
func F_show_timing_clock_source(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_show_timing_clock_source[0]))
	if v5 == int32(1) {
		v8 = int32(_a_F_show_timing_clock_source_0)
	} else {
		v8 = int32(_a_F_show_timing_clock_source_1)
	}
	if v5 != 0 {
		v10 = v8
	} else {
		v10 = int32(_a_F_show_timing_clock_source_2)
	}
	return v10
}
func F_sigismember(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v4 = l1 - int32(1)
	if base.Ui32(v4) <= base.Ui32(int32(63)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v4)>>(uint(int32(3))%32))&int32(536870908))))
		v17 = int32(base.Ui32(v12)>>(uint(v4)%32)) & int32(1)
	} else {
		v17 = int32(0)
	}
	return v17
}
func F_signal_child(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v15) <= base.Ui32(int32(17)) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(2))%32))+uint32(_c_F_signal_child[0])))
				v22 = v20
			} else {
				v22 = int32(_a_F_signal_child_0)
			}
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_c_F_signal_child[1])))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v31
			F_errmsg_internal(m, int32(_a_F_signal_child_1), v8+int32(32))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_signal_child_2), int32(3505), int32(_a_F_signal_child_3))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v44 = F_pgmem_kill(m, v10, l1)
					mBase = m.M
					if int32(0) <= v44 {
						if base.B2i32(int32(1)<<(uint(l1)%32)&int32(_a_F_signal_child_4) == int32(0))|base.B2i32(base.Ui32(int32(15)) < base.Ui32(l1)) != 0 {
							m.G0 = v8 + int32(48)
							return
						} else {
							v74 = int32(0)
							v75 = v74 - v10
							v76 = F_pgmem_kill(m, v75, l1)
							mBase = m.M
							if v74 <= v76 {
								m.G0 = v8 + int32(48)
								return
							} else {
								v81 = F_errstart(m, int32(12), int32(0))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									if v81 == int32(0) {
										m.G0 = v8 + int32(48)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
										F_errmsg_internal(m, int32(_a_F_signal_child_5), v8)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_signal_child_2), int32(3518), int32(_a_F_signal_child_3))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												m.G0 = v8 + int32(48)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v49 = F_errstart(m, int32(12), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							if v49 == int32(0) {
								if base.B2i32(int32(1)<<(uint(l1)%32)&int32(_a_F_signal_child_4) == int32(0))|base.B2i32(base.Ui32(int32(15)) < base.Ui32(l1)) != 0 {
									m.G0 = v8 + int32(48)
									return
								} else {
									v74 = int32(0)
									v75 = v74 - v10
									v76 = F_pgmem_kill(m, v75, l1)
									mBase = m.M
									if v74 <= v76 {
										m.G0 = v8 + int32(48)
										return
									} else {
										v81 = F_errstart(m, int32(12), int32(0))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											if v81 == int32(0) {
												m.G0 = v8 + int32(48)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
												F_errmsg_internal(m, int32(_a_F_signal_child_5), v8)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_signal_child_2), int32(3518), int32(_a_F_signal_child_3))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														m.G0 = v8 + int32(48)
														return
													}
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
								F_errmsg_internal(m, int32(_a_F_signal_child_5), v8+int32(16))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_signal_child_2), int32(3508), int32(_a_F_signal_child_3))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										if base.B2i32(int32(1)<<(uint(l1)%32)&int32(_a_F_signal_child_4) == int32(0))|base.B2i32(base.Ui32(int32(15)) < base.Ui32(l1)) != 0 {
											m.G0 = v8 + int32(48)
											return
										} else {
											v74 = int32(0)
											v75 = v74 - v10
											v76 = F_pgmem_kill(m, v75, l1)
											mBase = m.M
											if v74 <= v76 {
												m.G0 = v8 + int32(48)
												return
											} else {
												v81 = F_errstart(m, int32(12), int32(0))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													if v81 == int32(0) {
														m.G0 = v8 + int32(48)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
														F_errmsg_internal(m, int32(_a_F_signal_child_5), v8)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_signal_child_2), int32(3518), int32(_a_F_signal_child_3))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return
															} else {
																m.G0 = v8 + int32(48)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v44 = F_pgmem_kill(m, v10, l1)
			mBase = m.M
			if int32(0) <= v44 {
				if base.B2i32(int32(1)<<(uint(l1)%32)&int32(_a_F_signal_child_4) == int32(0))|base.B2i32(base.Ui32(int32(15)) < base.Ui32(l1)) != 0 {
					m.G0 = v8 + int32(48)
					return
				} else {
					v74 = int32(0)
					v75 = v74 - v10
					v76 = F_pgmem_kill(m, v75, l1)
					mBase = m.M
					if v74 <= v76 {
						m.G0 = v8 + int32(48)
						return
					} else {
						v81 = F_errstart(m, int32(12), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							if v81 == int32(0) {
								m.G0 = v8 + int32(48)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
								F_errmsg_internal(m, int32(_a_F_signal_child_5), v8)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_signal_child_2), int32(3518), int32(_a_F_signal_child_3))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										m.G0 = v8 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v49 = F_errstart(m, int32(12), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					if v49 == int32(0) {
						if base.B2i32(int32(1)<<(uint(l1)%32)&int32(_a_F_signal_child_4) == int32(0))|base.B2i32(base.Ui32(int32(15)) < base.Ui32(l1)) != 0 {
							m.G0 = v8 + int32(48)
							return
						} else {
							v74 = int32(0)
							v75 = v74 - v10
							v76 = F_pgmem_kill(m, v75, l1)
							mBase = m.M
							if v74 <= v76 {
								m.G0 = v8 + int32(48)
								return
							} else {
								v81 = F_errstart(m, int32(12), int32(0))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									if v81 == int32(0) {
										m.G0 = v8 + int32(48)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
										F_errmsg_internal(m, int32(_a_F_signal_child_5), v8)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_signal_child_2), int32(3518), int32(_a_F_signal_child_3))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												m.G0 = v8 + int32(48)
												return
											}
										}
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
						F_errmsg_internal(m, int32(_a_F_signal_child_5), v8+int32(16))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_signal_child_2), int32(3508), int32(_a_F_signal_child_3))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								if base.B2i32(int32(1)<<(uint(l1)%32)&int32(_a_F_signal_child_4) == int32(0))|base.B2i32(base.Ui32(int32(15)) < base.Ui32(l1)) != 0 {
									m.G0 = v8 + int32(48)
									return
								} else {
									v74 = int32(0)
									v75 = v74 - v10
									v76 = F_pgmem_kill(m, v75, l1)
									mBase = m.M
									if v74 <= v76 {
										m.G0 = v8 + int32(48)
										return
									} else {
										v81 = F_errstart(m, int32(12), int32(0))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											if v81 == int32(0) {
												m.G0 = v8 + int32(48)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
												F_errmsg_internal(m, int32(_a_F_signal_child_5), v8)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_signal_child_2), int32(3518), int32(_a_F_signal_child_3))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														m.G0 = v8 + int32(48)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_similar_to_escape_1(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_similar_escape_internal(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_similarity(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v185 int64
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v27 = v25 & v23
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v23
	goto L6
L5:
	;
	v28 = int32(4)
	goto L6
L6:
	;
	if v25 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = F_generate_trgm(m, v16+v28, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v35 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v46 = int32(1)
	if v27 != 0 {
		v56 = int32(base.Ui32(v25)>>(uint(v46)%32)) - v46
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v38 = int32(16)
	goto L13
L12:
	;
	v38 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = int32(4)
	goto L16
L15:
	;
	v45 = v38
	goto L16
L16:
	;
	v56 = v45
	goto L7
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v59 = int32(1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v63 = v61 & v59
	if v63 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v64 = v59
	goto L21
L20:
	;
	v64 = int32(4)
	goto L21
L21:
	;
	if v61 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v93 = F_generate_trgm(m, v21+v64, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L33
	}
L23:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v71 == int32(18) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v82 = int32(1)
	if v63 != 0 {
		v92 = int32(base.Ui32(v61)>>(uint(v82)%32)) - v82
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v74 = int32(16)
	goto L28
L27:
	;
	v74 = int32(0)
	goto L28
L28:
	;
	if base.Ui32((v71-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = int32(4)
	goto L31
L30:
	;
	v81 = v74
	goto L31
L31:
	;
	v92 = v81
	goto L22
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v92 = int32(base.Ui32(v86)>>(uint(int32(2))%32)) - int32(4)
	goto L22
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v96 = int32(2)
	v98 = int32(5)
	v99 = int32(base.Ui32(v95)>>(uint(v96)%32)) - v98
	v100 = int32(3)
	v101 = base.I32_div_u_s(v99, v100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v106 = int32(base.Ui32(v102)>>(uint(v96)%32)) - v98
	v108 = base.I32_div_u_s(v106, v100)
	if base.B2i32(base.Ui32(v106) < base.Ui32(v100))|base.B2i32(base.Ui32(v99) < base.Ui32(v100)) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v116 = int32(5)
	v117 = v57 + v116
	v119 = v93 + v116
	v121 = v119
	v122 = v117
	v126 = int32(0)
	goto L37
L35:
	;
	v185 = int64(0)
	goto L36
L36:
	;
	F_pfree(m, v57)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L51
	}
L37:
	;
	v136 = base.I32_div_s(v121-v119, int32(3))
	if v136 < v101 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v185 = base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_div(base.F32_convert_i32_s(v163), base.F32_convert_i32_s(v101+v108-v163))))
	goto L36
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_similarity[0]))
	v140 = m.T0[v139].(func(*base.Module, int32, int32) int32)(m, v122, v121)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v163 = v126
	goto L41
L41:
	;
	goto L38
L42:
	;
	v159 = base.I32_div_s(v155-v117, int32(3))
	if v159 < v108 {
		v121 = v154
		v122 = v155
		v126 = v156
		goto L37
	} else {
		goto L50
	}
L43:
	;
	if int32(0) <= v140 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v140 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v150 = v121
	v151 = v126
	goto L46
L46:
	;
	v154 = v150
	v155 = v122 + int32(3)
	v156 = v151
	goto L42
L47:
	;
	v154 = v121 + int32(3)
	v155 = v122
	v156 = v126
	goto L42
L48:
	;
	goto L49
L49:
	;
	v150 = v121 + int32(3)
	v151 = v126 + int32(1)
	goto L46
L50:
	;
	v163 = v156
	goto L41
L51:
	;
	F_pfree(m, v93)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v190 != v16 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v16)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v194 != v21 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v21)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return v185
L60:
	;
	goto L59
}
func F_similarity_op(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 float64
	_ = v12
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_DirectFunctionCall2Coll(m, int32(_a_F_similarity_op_0), int32(0), v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, _c_F_similarity_op[0]))
		return base.I64_extend_i32_u(base.F64_le(v12, base.F64_promote_f32(base.F32_reinterpret_i32(base.I32_wrap_i64(v7)))))
	}
}
func F_size_box(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v9 float64
	_ = v9
	var v15 float64
	_ = v15
	var v18 int64
	_ = v18
	var v23 float64
	_ = v23
	var v29 float64
	_ = v29
	var v32 int64
	_ = v32
	var v38 int64
	_ = v38
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v58 float64
	_ = v58
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v100 float64
	_ = v100
	var v101 int32
	_ = v101
	var v103 float64
	_ = v103
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v117 float64
	_ = v117
	v2 = float64(0)
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v9)&int64(9223372036854775807)) {
		v117 = v2
		return v117
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
		v18 = base.I64_reinterpret_f64(v15) & int64(9223372036854775807)
		if base.B2i32(base.Ui64(v18) < base.Ui64(int64(9218868437227405313)))&base.F64_le(v15, v9) != 0 {
			v117 = v2
			return v117
		} else {
			v23 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v23)&int64(9223372036854775807)) {
				v117 = v2
				return v117
			} else {
				v29 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
				v32 = base.I64_reinterpret_f64(v29) & int64(9223372036854775807)
				if base.B2i32(base.Ui64(v32) < base.Ui64(int64(9218868437227405313)))&base.F64_le(v29, v23) != 0 {
					v117 = v2
					return v117
				} else {
					v38 = int64(9218868437227405312)
					if base.B2i32(base.Ui64(v38) < base.Ui64(v18))|base.B2i32(base.Ui64(v38) < base.Ui64(v32)) != 0 {
						v117 = math.Float64frombits(uint64(0x7ff0000000000000))
						return v117
					} else {
						v44 = math.Float64frombits(uint64(0x7ff0000000000000))
						v46 = base.F64_sub(v15, v9)
						if base.F64_eq(base.F64_abs(v15), v44)|base.F64_ne(base.F64_abs(v46), v44)|base.F64_eq(base.F64_abs(v9), v44) == int32(0) {
							v58 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return float64(0)
							} else {
								v62 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
								v63 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
								v64 = v62
								v65 = v58
								v66 = v63
								v68 = math.Float64frombits(uint64(0x7ff0000000000000))
								v70 = base.F64_sub(v64, v66)
								if base.F64_eq(base.F64_abs(v64), v68)|base.F64_ne(base.F64_abs(v70), v68)|base.F64_eq(base.F64_abs(v66), v68) == int32(0) {
									v82 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return float64(0)
									} else {
										v84 = v82
										v86 = math.Float64frombits(uint64(0x7ff0000000000000))
										v88 = base.F64_mul(v65, v84)
										if base.F64_eq(base.F64_abs(v65), v86)|base.F64_ne(base.F64_abs(v88), v86)|base.F64_eq(base.F64_abs(v84), v86) == int32(0) {
											v100 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return float64(0)
											} else {
												return v100
											}
										} else {
											v103 = float64(0)
											if base.F64_eq(v65, v103)|base.F64_ne(v88, v103)|base.F64_eq(v84, v103) != 0 {
												v117 = v88
												return v117
											} else {
												v112 = F_float_underflow_error_ext(m, int32(0))
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return float64(0)
												} else {
													v117 = v112
													return v117
												}
											}
										}
									}
								} else {
									v84 = v70
									v86 = math.Float64frombits(uint64(0x7ff0000000000000))
									v88 = base.F64_mul(v65, v84)
									if base.F64_eq(base.F64_abs(v65), v86)|base.F64_ne(base.F64_abs(v88), v86)|base.F64_eq(base.F64_abs(v84), v86) == int32(0) {
										v100 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return float64(0)
										} else {
											return v100
										}
									} else {
										v103 = float64(0)
										if base.F64_eq(v65, v103)|base.F64_ne(v88, v103)|base.F64_eq(v84, v103) != 0 {
											v117 = v88
											return v117
										} else {
											v112 = F_float_underflow_error_ext(m, int32(0))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return float64(0)
											} else {
												v117 = v112
												return v117
											}
										}
									}
								}
							}
						} else {
							v64 = v29
							v65 = v46
							v66 = v23
							v68 = math.Float64frombits(uint64(0x7ff0000000000000))
							v70 = base.F64_sub(v64, v66)
							if base.F64_eq(base.F64_abs(v64), v68)|base.F64_ne(base.F64_abs(v70), v68)|base.F64_eq(base.F64_abs(v66), v68) == int32(0) {
								v82 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return float64(0)
								} else {
									v84 = v82
									v86 = math.Float64frombits(uint64(0x7ff0000000000000))
									v88 = base.F64_mul(v65, v84)
									if base.F64_eq(base.F64_abs(v65), v86)|base.F64_ne(base.F64_abs(v88), v86)|base.F64_eq(base.F64_abs(v84), v86) == int32(0) {
										v100 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return float64(0)
										} else {
											return v100
										}
									} else {
										v103 = float64(0)
										if base.F64_eq(v65, v103)|base.F64_ne(v88, v103)|base.F64_eq(v84, v103) != 0 {
											v117 = v88
											return v117
										} else {
											v112 = F_float_underflow_error_ext(m, int32(0))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return float64(0)
											} else {
												v117 = v112
												return v117
											}
										}
									}
								}
							} else {
								v84 = v70
								v86 = math.Float64frombits(uint64(0x7ff0000000000000))
								v88 = base.F64_mul(v65, v84)
								if base.F64_eq(base.F64_abs(v65), v86)|base.F64_ne(base.F64_abs(v88), v86)|base.F64_eq(base.F64_abs(v84), v86) == int32(0) {
									v100 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return float64(0)
									} else {
										return v100
									}
								} else {
									v103 = float64(0)
									if base.F64_eq(v65, v103)|base.F64_ne(v88, v103)|base.F64_eq(v84, v103) != 0 {
										v117 = v88
										return v117
									} else {
										v112 = F_float_underflow_error_ext(m, int32(0))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return float64(0)
										} else {
											v117 = v112
											return v117
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_skeys(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_skeys(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_smgrcreate(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
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
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	v9 = int32(_a_F_smgrcreate_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[0])) = v11 + int32(1)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v335 = int32(_a_F_smgrcreate_0)
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[0])) = v337 - int32(1)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[1])) = v256
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L16
	} else {
		goto L100
	}
L3:
	;
	m.G0 = v17 + int32(80)
	goto L1
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+40))
	if int32(0) < v22 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = m.G0
	v29 = v27 - int32(160)
	m.G0 = v29
	if v25 != int32(1664) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v237 = v17 + int32(8)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v237, v238, v239, v240, v241, l1)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L77
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L16
	} else {
		goto L73
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L16
	} else {
		goto L69
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L16
	} else {
		goto L65
	}
L12:
	;
	v33 = F_GetDatabasePath(m, v26, v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	m.G0 = v29 + int32(160)
	goto L8
L15:
	;
	F_pfree(m, v33)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L16
	} else {
		goto L64
	}
L16:
	;
	return
L17:
	;
	v36 = v29 - int32(-64)
	v39 = F___fstatat(m, int32(-100), v33, v36, int32(0))
	mBase = m.M
	goto L18
L18:
	;
	if v39 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[1]))
	if v43 == int32(44) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	if v174&int32(_a_F_smgrcreate_1) != int32(_a_F_smgrcreate_2) {
		goto L9
	} else {
		goto L63
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[2]))
	v51 = F_LWLockAcquire(m, v47+int32(2432), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L16
	} else {
		goto L59
	}
L25:
	;
	v55 = F___fstatat(m, int32(-100), v33, v36, int32(0))
	mBase = m.M
	goto L27
L26:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[2]))
	F_LWLockRelease(m, v152+int32(2432))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L16
	} else {
		goto L58
	}
L27:
	;
	if v55 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	if v58&int32(_a_F_smgrcreate_1) == int32(_a_F_smgrcreate_2) {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[3]))
	v65 = F_mkdir(m, v33, v64)
	mBase = m.M
	goto L32
L31:
	;
	goto L30
L32:
	;
	if int32(0) <= v65 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	if l2 == int32(0) {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[1]))
	if v71 != int32(44) {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[3]))
	v81 = m.G0
	v83 = v81 - int32(96)
	m.G0 = v83
	v86 = F_umask(m, int32(0))
	mBase = m.M
	v89 = F_umask(m, v86&int32(-193))
	mBase = m.M
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v96 = v33 + base.B2i32(v90 == int32(47))
	goto L37
L36:
	;
	if v143 < int32(0) {
		goto L10
	} else {
		goto L57
	}
L37:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v101 != int32(47) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	v145 = F_umask(m, v86)
	mBase = m.M
	m.G0 = v83 + int32(96)
	goto L36
L39:
	;
	goto L38
L40:
	;
	v96 = v96 + int32(1)
	goto L37
L41:
	;
	if v115 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L42:
	;
	v117 = F_mkdir(m, v33, v116)
	mBase = m.M
	if int32(0) <= v117 {
		goto L41
	} else {
		goto L49
	}
L43:
	;
	v113 = F_umask(m, v86)
	mBase = m.M
	v115 = int32(0)
	v116 = v75
	goto L42
L44:
	;
	if v101 != 0 {
		goto L40
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v106)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v108 == v106 {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v104)
	goto L43
L48:
	;
	v115 = int32(1)
	v116 = int32(511)
	goto L42
L49:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[1]))
	if v121 != int32(20) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[1])) = v121
	v143 = int32(-1)
	goto L39
L51:
	;
	v124 = F_stat(m, v33, v83)
	mBase = m.M
	if v124 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v125&int32(_a_F_smgrcreate_1) == int32(_a_F_smgrcreate_2) {
		goto L41
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	v143 = int32(0)
	goto L39
L55:
	;
	goto L56
L56:
	;
	v137 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v137)
	goto L40
L57:
	;
	goto L26
L58:
	;
	goto L15
L59:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v33
	F_errmsg(m, int32(_a_F_smgrcreate_3), v29+int32(32))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L16
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_smgrcreate_4), int32(187), int32(_a_F_smgrcreate_5))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	goto L15
L64:
	;
	goto L14
L65:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L16
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v33
	F_errmsg(m, int32(_a_F_smgrcreate_6), v29+int32(16))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_smgrcreate_4), int32(161), int32(_a_F_smgrcreate_5))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v33
	F_errmsg(m, int32(_a_F_smgrcreate_6), v29)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_smgrcreate_4), int32(177), int32(_a_F_smgrcreate_5))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v33
	F_errmsg(m, int32(_a_F_smgrcreate_7), v29+int32(48))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_smgrcreate_4), int32(197), int32(_a_F_smgrcreate_5))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[4]))
	if v247&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v250 = int32(_a_F_smgrcreate_8)
	goto L80
L79:
	;
	v250 = int32(194)
	goto L80
L80:
	;
	v251 = F_PathNameOpenFile(m, v237, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	if v251 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[1]))
	if l2 == int32(0) {
		goto L2
	} else {
		goto L85
	}
L83:
	;
	v271 = v251
	goto L84
L84:
	;
	v274 = l0 + l1<<(uint(int32(2))%32)
	v276 = v274 + int32(40)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v277 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[4]))
	if v262&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v265 = int32(_a_F_smgrcreate_9)
	goto L88
L87:
	;
	v265 = int32(2)
	goto L88
L88:
	;
	v266 = F_PathNameOpenFile(m, v237, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L16
	} else {
		goto L89
	}
L89:
	;
	if v266 < int32(0) {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	v271 = v266
	goto L84
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v271
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v302 != int32(-1) {
		goto L3
	} else {
		goto L98
	}
L92:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_smgrcreate[5]))
	v283 = F_MemoryContextAlloc(m, v281, int32(8))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L16
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v287 = v274 + int32(56)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	if int32(0) < v277 {
		v295 = v288
		goto L91
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+56)) = v283
	v295 = v283
	goto L91
L96:
	;
	v292 = F_repalloc(m, v288, int32(8))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L16
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v292
	v295 = v292
	goto L91
L98:
	;
	F_register_dirty_segment(m, l0, l1, v295)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	goto L3
L100:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v17 + int32(8)
	F_errmsg(m, int32(_a_F_smgrcreate_10), v17)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_smgrcreate_11), int32(263), int32(_a_F_smgrcreate_12))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L16
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_smgropen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int64
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = int32(_a_F_smgropen_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_smgropen[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgropen[0])) = v11 + int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_smgropen[1]))
	if v16 == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(360777252880)
		v26 = F_hash_create(m, int32(_a_F_smgropen_1), int64(400), v5+int32(-56), int32(40))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v31 = int32(_a_F_smgropen_2)
			*(*int32)(unsafe.Add(mBase, _c_F_smgropen[2])) = v31
			*(*int32)(unsafe.Add(mBase, _c_F_smgropen[3])) = v31
			*(*int32)(unsafe.Add(mBase, _c_F_smgropen[1])) = v26
			v38 = v26
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v39
			v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v41
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
			v49 = F_hash_search(m, v38, v5+int32(-56), int32(1), v5+int32(-1))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+63)))
				if v51 == int32(0) {
					v54 = int64(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = v54
					*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(v49)+72)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = int64(4294967295)
					v63 = v49 + int32(76)
					v65 = *(*int32)(unsafe.Add(mBase, _c_F_smgropen[2]))
					if v65 != 0 {
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_smgropen[3]))
						v72 = v67
					} else {
						v69 = int32(_a_F_smgropen_2)
						*(*int32)(unsafe.Add(mBase, _c_F_smgropen[2])) = v69
						v72 = v69
					}
					*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = v72
					v74 = int32(_a_F_smgropen_2)
					*(*int32)(unsafe.Add(mBase, uint32(v49)+80)) = v74
					*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v63
					*(*int32)(unsafe.Add(mBase, _c_F_smgropen[3])) = v63
					v79 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v49)+48)) = v79
					*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = v79
				} else {
				}
				v85 = int32(_a_F_smgropen_0)
				v87 = *(*int32)(unsafe.Add(mBase, _c_F_smgropen[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_smgropen[0])) = v87 - int32(1)
				m.G0 = v7 - int32(-64)
				return v49
			}
		}
	} else {
		v38 = v16
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v39
		v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v41
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
		v49 = F_hash_search(m, v38, v5+int32(-56), int32(1), v5+int32(-1))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+63)))
			if v51 == int32(0) {
				v54 = int64(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = v54
				*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v49)+72)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = int64(4294967295)
				v63 = v49 + int32(76)
				v65 = *(*int32)(unsafe.Add(mBase, _c_F_smgropen[2]))
				if v65 != 0 {
					v67 = *(*int32)(unsafe.Add(mBase, _c_F_smgropen[3]))
					v72 = v67
				} else {
					v69 = int32(_a_F_smgropen_2)
					*(*int32)(unsafe.Add(mBase, _c_F_smgropen[2])) = v69
					v72 = v69
				}
				*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = v72
				v74 = int32(_a_F_smgropen_2)
				*(*int32)(unsafe.Add(mBase, uint32(v49)+80)) = v74
				*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v63
				*(*int32)(unsafe.Add(mBase, _c_F_smgropen[3])) = v63
				v79 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v49)+48)) = v79
				*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = v79
			} else {
			}
			v85 = int32(_a_F_smgropen_0)
			v87 = *(*int32)(unsafe.Add(mBase, _c_F_smgropen[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_smgropen[0])) = v87 - int32(1)
			m.G0 = v7 - int32(-64)
			return v49
		}
	}
}
func F_smgrregistersync(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	v10 = int32(_a_F_smgrregistersync_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_smgrregistersync[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrregistersync[0])) = v12 + int32(1)
	v16 = F_mdnblocks(m, l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = l0 + l1<<(uint(int32(2))%32) + int32(40)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = v23
	goto L3
L3:
	;
	v36 = F__mdfd_openseg(m, l0, l1, v24, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	if int32(0) < v24 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v36 != 0 {
		v24 = v24 + int32(1)
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v44 = l0 + l1<<(uint(int32(2))%32) + int32(56)
	v46 = v24
	goto L10
L8:
	;
	goto L9
L9:
	;
	v100 = int32(_a_F_smgrregistersync_0)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_smgrregistersync[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrregistersync[0])) = v102 - int32(1)
	return
L10:
	;
	v55 = v46 - int32(1)
	v57 = v55 << (uint(int32(3)) % 32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v59 = v57 + v58
	F_register_dirty_segment(m, l0, l1, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	if v23 < v46 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	F_FileClose(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(int32(1)) < base.Ui32(v46) {
		v46 = v55
		goto L10
	} else {
		goto L30
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v55 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v55
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v85
	goto L17
L19:
	;
	if v66 <= int32(0) {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v66 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	F_pfree(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v85 = int32(0)
	goto L18
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_smgrregistersync[1]))
	v79 = F_MemoryContextAlloc(m, v78, v57)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v55 <= v66 {
		goto L17
	} else {
		goto L28
	}
L27:
	;
	v85 = v79
	goto L18
L28:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v83 = F_repalloc(m, v82, v57)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v85 = v83
	goto L18
L30:
	;
	goto L11
}
func F_smgrzeroextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int64
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v266 int32
	_ = v266
	var v282 int32
	_ = v282
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int64
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int64
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v520 int32
	_ = v520
	var v536 int32
	_ = v536
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	v15 = int32(_a_F_smgrzeroextend_0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[0])) = v17 + int32(1)
	v21 = m.G0
	v23 = v21 - int32(128)
	m.G0 = v23
	if base.Ui64(base.I64_extend_i32_s(l3)+base.I64_extend_i32_u(l2)) <= base.Ui64(int64(4294967294)) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v683 = l0 + l1<<(uint(int32(2))%32) + int32(20)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	if v686 != l2 {
		goto L148
	} else {
		goto L149
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L12
	} else {
		goto L142
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L12
	} else {
		goto L139
	}
L4:
	;
	if int32(0) < l3 {
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
	v611 = m.ExcPending
	if v611 != 0 {
		goto L12
	} else {
		goto L134
	}
L7:
	;
	v39 = l3
	v41 = l2
	goto L10
L8:
	;
	goto L9
L9:
	;
	m.G0 = v23 + int32(128)
	goto L1
L10:
	;
	v47 = v41 & int32(_a_F_smgrzeroextend_1)
	v50 = base.I64_extend_i32_u(v47 << (uint(int32(13)) % 32))
	v53 = F__mdfd_getseg(m, l0, l1, v41, int32(0), int32(4))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	return
L13:
	;
	v55 = int32(_a_F_smgrzeroextend_2)
	if base.Ui32(v55) < base.Ui32(v47+v39) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v582 == int32(-1) {
		goto L129
	} else {
		goto L130
	}
L15:
	;
	v370 = base.I64_extend_i32_u(v60 << (uint(int32(13)) % 32))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v372 = F_FileAccess(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L12
	} else {
		goto L84
	}
L16:
	;
	v60 = v55 - v47
	goto L18
L17:
	;
	v60 = v39
	goto L18
L18:
	;
	if base.Ui32(v60) < base.Ui32(int32(9)) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[1]))
	if v64 == int32(1) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	if v64 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v69 = base.I64_extend_i32_u(v60 << (uint(int32(13)) % 32))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v71 = F_FileAccess(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L24
	}
L22:
	;
	if v327 == int32(0) {
		goto L14
	} else {
		goto L75
	}
L23:
	;
	v327 = int32(-1)
	goto L22
L24:
	;
	if v71 < int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	goto L26
L26:
	;
	v91 = int32(_a_F_smgrzeroextend_3)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(167772179)
	v95 = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[3]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97+v70*int32(48))))
	v101 = m.Env.X__syscall_fallocate(m, v99, v95, v50, v69)
	mBase = m.M
	v102 = v95 - v101
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v95
	if v102 == int32(27) {
		goto L26
	} else {
		goto L28
	}
L27:
	;
	if v102 == int32(0) {
		v327 = v102
		goto L22
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[4])) = v102
	if base.B2i32(v102 != int32(138))&base.B2i32(v102 != int32(28)) != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v118 = F_FileAccess(m, v70)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	if v118 < int32(0) {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = int32(167772179)
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[3]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v70*int32(48))))
	v142 = m.G0
	v144 = v142 - int32(1024)
	m.G0 = v144
	v147 = base.I32_wrap_i64(v69)
	v148 = v50
	v156 = int32(0)
	goto L34
L33:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[2]))
	v296 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v296
	if v282 < v296 {
		goto L23
	} else {
		goto L72
	}
L34:
	;
	v158 = int32(0)
	if v147 == v158 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	m.G0 = v144 + int32(1024)
	goto L33
L36:
	;
	goto L35
L37:
	;
	v282 = v156
	goto L36
L38:
	;
	goto L39
L39:
	;
	v162 = v147
	v164 = v158
	goto L40
L40:
	;
	v175 = v144 + v164<<(uint(int32(3))%32)
	v176 = int32(_a_F_smgrzeroextend_4)
	if base.Ui32(v176) <= base.Ui32(v162) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v190 = m.G0
	v192 = v190 - int32(1024)
	m.G0 = v192
	if v184 <= int32(128) {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	goto L41
L43:
	;
	v179 = v176
	goto L45
L44:
	;
	v179 = v162
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = int32(_a_F_smgrzeroextend_5)
	v184 = v164 + int32(1)
	v185 = v162 - v179
	if base.Ui32(int32(126)) < base.Ui32(v164) {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	if v185 != 0 {
		v162 = v185
		v164 = v184
		goto L40
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	m.G0 = v192 + int32(1024)
	if int32(0) <= v266 {
		v147 = v185
		v148 = v148 + base.I64_extend_i32_u(v266)
		v156 = v156 + v266
		goto L34
	} else {
		goto L71
	}
L49:
	;
	v199 = v144
	v200 = v184
	v203 = int32(0)
	v207 = v148
	goto L52
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[4])) = int32(28)
	v266 = int32(-1)
	goto L48
L52:
	;
	if v200 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v266 = v218
	goto L48
L54:
	;
	if v214 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v212 = F_pwrite(m, v131, v210, v211, v207)
	mBase = m.M
	v214 = v212
	goto L54
L56:
	;
	goto L57
L57:
	;
	v213 = F_pwritev(m, v131, v199, v200, v207)
	mBase = m.M
	v214 = v213
	goto L54
L58:
	;
	v266 = int32(-1)
	goto L48
L59:
	;
	goto L60
L60:
	;
	v218 = v214 + v203
	v224 = v199
	v225 = v200
	v227 = v214
	goto L61
L61:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if base.Ui32(v233) <= base.Ui32(v227) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v224 == v192 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v239 = v225 - int32(1)
	if v239 != 0 {
		v224 = v224 + int32(8)
		v225 = v239
		v227 = v227 - v233
		goto L61
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	v266 = v218
	goto L48
L67:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v247 + v227
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v250 - v227
	if int32(0) < v225 {
		v199 = v192
		v200 = v225
		v203 = v218
		v207 = v207 + base.I64_extend_i32_u(v214)
		goto L52
	} else {
		goto L70
	}
L68:
	;
	v242 = v225 << (uint(int32(3)) % 32)
	if v242 == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	base.MemoryCopy(m, v192, v224, v242)
	goto L67
L70:
	;
	goto L53
L71:
	;
	v282 = v266
	goto L36
L72:
	;
	if v69 == base.I64_extend_i32_u(v282) {
		v327 = int32(0)
		goto L22
	} else {
		goto L73
	}
L73:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[4]))
	if v304 != 0 {
		goto L23
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[4])) = int32(51)
	goto L23
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[3]))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v347+v345*int32(48))+32))
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v351
	F_errmsg(m, int32(_a_F_smgrzeroextend_6), v23+int32(16))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	F_errhint(m, int32(_a_F_smgrzeroextend_7), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L12
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_smgrzeroextend_8), int32(630), int32(_a_F_smgrzeroextend_9))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L12
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	if v565 < int32(0) {
		goto L2
	} else {
		goto L128
	}
L83:
	;
	v565 = int32(-1)
	goto L82
L84:
	;
	if v372 < int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = int32(167772179)
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[3]))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v381+v371*int32(48))))
	v396 = m.G0
	v398 = v396 - int32(1024)
	m.G0 = v398
	v401 = base.I32_wrap_i64(v370)
	v402 = v50
	v410 = int32(0)
	goto L87
L86:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[2]))
	v550 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = v550
	if v536 < v550 {
		goto L83
	} else {
		goto L125
	}
L87:
	;
	v412 = int32(0)
	if v401 == v412 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	m.G0 = v398 + int32(1024)
	goto L86
L89:
	;
	goto L88
L90:
	;
	v536 = v410
	goto L89
L91:
	;
	goto L92
L92:
	;
	v416 = v401
	v418 = v412
	goto L93
L93:
	;
	v429 = v398 + v418<<(uint(int32(3))%32)
	v430 = int32(_a_F_smgrzeroextend_4)
	if base.Ui32(v430) <= base.Ui32(v416) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v444 = m.G0
	v446 = v444 - int32(1024)
	m.G0 = v446
	if v438 <= int32(128) {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	goto L94
L96:
	;
	v433 = v430
	goto L98
L97:
	;
	v433 = v416
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+4)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = int32(_a_F_smgrzeroextend_5)
	v438 = v418 + int32(1)
	v439 = v416 - v433
	if base.Ui32(int32(126)) < base.Ui32(v418) {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	if v439 != 0 {
		v416 = v439
		v418 = v438
		goto L93
	} else {
		goto L100
	}
L100:
	;
	goto L95
L101:
	;
	m.G0 = v446 + int32(1024)
	if int32(0) <= v520 {
		v401 = v439
		v402 = v402 + base.I64_extend_i32_u(v520)
		v410 = v410 + v520
		goto L87
	} else {
		goto L124
	}
L102:
	;
	v453 = v398
	v454 = v438
	v457 = int32(0)
	v461 = v402
	goto L105
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[4])) = int32(28)
	v520 = int32(-1)
	goto L101
L105:
	;
	if v454 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v520 = v472
	goto L101
L107:
	;
	if v468 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v466 = F_pwrite(m, v385, v464, v465, v461)
	mBase = m.M
	v468 = v466
	goto L107
L109:
	;
	goto L110
L110:
	;
	v467 = F_pwritev(m, v385, v453, v454, v461)
	mBase = m.M
	v468 = v467
	goto L107
L111:
	;
	v520 = int32(-1)
	goto L101
L112:
	;
	goto L113
L113:
	;
	v472 = v468 + v457
	v478 = v453
	v479 = v454
	v481 = v468
	goto L114
L114:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	if base.Ui32(v487) <= base.Ui32(v481) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v478 == v446 {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	v493 = v479 - int32(1)
	if v493 != 0 {
		v478 = v478 + int32(8)
		v479 = v493
		v481 = v481 - v487
		goto L114
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	goto L115
L119:
	;
	v520 = v472
	goto L101
L120:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	*(*int32)(unsafe.Add(mBase, uint32(v446))) = v501 + v481
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v446)+4)) = v504 - v481
	if int32(0) < v479 {
		v453 = v446
		v454 = v479
		v457 = v472
		v461 = v461 + base.I64_extend_i32_u(v468)
		goto L105
	} else {
		goto L123
	}
L121:
	;
	v496 = v479 << (uint(int32(3)) % 32)
	if v496 == int32(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	base.MemoryCopy(m, v446, v478, v496)
	goto L120
L123:
	;
	goto L106
L124:
	;
	v536 = v520
	goto L89
L125:
	;
	if v370 == base.I64_extend_i32_u(v536) {
		v565 = int32(0)
		goto L82
	} else {
		goto L126
	}
L126:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[4]))
	if v558 != 0 {
		goto L83
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[4])) = int32(51)
	goto L83
L128:
	;
	goto L14
L129:
	;
	F_register_dirty_segment(m, l0, l1, v53)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L12
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v588 = v39 - v60
	if int32(0) < v588 {
		v39 = v588
		v41 = v41 + v60
		goto L10
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	goto L11
L134:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	v616 = v23 + int32(56)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v616, v617, v618, v619, v620, l1)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L12
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v616
	F_errmsg(m, int32(_a_F_smgrzeroextend_10), v23)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L12
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_smgrzeroextend_8), int32(577), int32(_a_F_smgrzeroextend_9))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L12
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v639
	F_errmsg_internal(m, int32(_a_F_smgrzeroextend_11), v23+int32(32))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L12
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_smgrzeroextend_8), int32(622), int32(_a_F_smgrzeroextend_9))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v659 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[3]))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v659+v657*int32(48))+32))
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v663
	F_errmsg(m, int32(_a_F_smgrzeroextend_12), v23+int32(48))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L12
	} else {
		goto L145
	}
L145:
	;
	F_errhint(m, int32(_a_F_smgrzeroextend_7), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_smgrzeroextend_8), int32(652), int32(_a_F_smgrzeroextend_9))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v688 = int32(-1)
	goto L150
L149:
	;
	v688 = l2 + l3
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = v688
	v690 = int32(_a_F_smgrzeroextend_0)
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrzeroextend[0])) = v692 - int32(1)
	return
}
func F_sortins_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v8 < v11 {
		return int32(-1)
	} else {
		v15 = int32(1)
		if v11 < v8 {
			v28 = v15
			return v28
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6)+4)))
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			if v17 < v18 {
				return int32(-1)
			} else {
				if v18 < v17 {
					v28 = v15
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					if v24 < v25 {
						v28 = int32(-1)
					} else {
						v28 = base.B2i32(v25 < v24)
					}
				}
				return v28
			}
		}
	}
}
func F_spanish_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 < v10 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v509 < v10 {
		goto L149
	} else {
		goto L150
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v496
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v290 < v10 {
		goto L83
	} else {
		goto L84
	}
L4:
	;
	if v63 != 0 {
		goto L3
	} else {
		goto L18
	}
L5:
	;
	v22 = v10
	goto L7
L6:
	;
	v22 = v6
	goto L7
L7:
	;
	goto L9
L8:
	;
	v63 = v58
	goto L4
L9:
	;
	if v10 == v22 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v58 = int32(0)
	goto L8
L11:
	;
	v63 = int32(-1)
	goto L4
L12:
	;
	goto L13
L13:
	;
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v10))))
	if int32(252) < v37 {
		v58 = v34
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v39 = v37 - int32(97)
	if v39 < int32(0) {
		v58 = v34
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v39)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v45)>>(uint(v39&int32(7))%32))&int32(1) == int32(0) {
		v58 = v34
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
	goto L17
L17:
	;
	goto L10
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v73 < v64 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v178 < v64 {
		goto L53
	} else {
		goto L54
	}
L20:
	;
	if v113 != 0 {
		goto L19
	} else {
		goto L35
	}
L21:
	;
	v75 = v64
	goto L23
L22:
	;
	v75 = v73
	goto L23
L23:
	;
	goto L25
L24:
	;
	v113 = v110
	goto L20
L25:
	;
	if v64 == v75 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v110 = int32(0)
	goto L24
L27:
	;
	v113 = int32(-1)
	goto L20
L28:
	;
	goto L29
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v64))))
	if int32(252) < v88 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64 + int32(1)
	goto L34
L31:
	;
	v90 = v88 - int32(97)
	if v90 < int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v93 = int32(1)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v90)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v97)>>(uint(v90&int32(7))%32))&v93 != 0 {
		v110 = v93
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	goto L26
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v122 < v121 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v162 < int32(0) {
		goto L19
	} else {
		goto L51
	}
L37:
	;
	v124 = v121
	goto L39
L38:
	;
	v124 = v122
	goto L39
L39:
	;
	v131 = v121
	goto L41
L40:
	;
	v162 = v142
	goto L36
L41:
	;
	if v131 == v124 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v162 = int32(-1)
	goto L36
L44:
	;
	goto L45
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v131))))
	if int32(252) < v137 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v154 = v131 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	v131 = v154
	goto L41
L47:
	;
	v139 = v137 - int32(97)
	if v139 < int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v142 = int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v139)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v146)>>(uint(v139&int32(7))%32))&v142 != 0 {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L46
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v496 = v165 + v162
	goto L2
L52:
	;
	if v221 != 0 {
		goto L3
	} else {
		goto L66
	}
L53:
	;
	v180 = v64
	goto L55
L54:
	;
	v180 = v178
	goto L55
L55:
	;
	goto L57
L56:
	;
	v221 = v216
	goto L52
L57:
	;
	if v64 == v180 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v216 = int32(0)
	goto L56
L59:
	;
	v221 = int32(-1)
	goto L52
L60:
	;
	goto L61
L61:
	;
	v192 = int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v64))))
	if int32(252) < v195 {
		v216 = v192
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v197 = v195 - int32(97)
	if v197 < int32(0) {
		v216 = v192
		goto L56
	} else {
		goto L63
	}
L63:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v197)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v203)>>(uint(v197&int32(7))%32))&int32(1) == int32(0) {
		v216 = v192
		goto L56
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64 + int32(1)
	goto L65
L65:
	;
	goto L58
L66:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v231 < v230 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v274 < int32(0) {
		goto L3
	} else {
		goto L81
	}
L68:
	;
	v233 = v230
	goto L70
L69:
	;
	v233 = v231
	goto L70
L70:
	;
	v239 = v230
	goto L72
L71:
	;
	v274 = int32(1)
	goto L67
L72:
	;
	if v239 == v233 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v274 = int32(-1)
	goto L67
L75:
	;
	goto L76
L76:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v239))))
	if int32(252) < v248 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v250 = v248 - int32(97)
	if v250 < int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v250)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v256)>>(uint(v250&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v265 = v239 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v265
	v239 = v265
	goto L72
L81:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v496 = v277 + v274
	goto L2
L82:
	;
	if v330 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v292 = v10
	goto L85
L84:
	;
	v292 = v290
	goto L85
L85:
	;
	goto L87
L86:
	;
	v330 = v327
	goto L82
L87:
	;
	if v10 == v292 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v327 = int32(0)
	goto L86
L89:
	;
	v330 = int32(-1)
	goto L82
L90:
	;
	goto L91
L91:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v10))))
	if int32(252) < v305 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
	goto L96
L93:
	;
	v307 = v305 - int32(97)
	if v307 < int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v310 = int32(1)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v307)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v314)>>(uint(v307&int32(7))%32))&v310 != 0 {
		v327 = v310
		goto L86
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	goto L88
L97:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v340 < v331 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v445 < v331 {
		goto L132
	} else {
		goto L133
	}
L99:
	;
	if v380 != 0 {
		goto L98
	} else {
		goto L114
	}
L100:
	;
	v342 = v331
	goto L102
L101:
	;
	v342 = v340
	goto L102
L102:
	;
	goto L104
L103:
	;
	v380 = v377
	goto L99
L104:
	;
	if v331 == v342 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v377 = int32(0)
	goto L103
L106:
	;
	v380 = int32(-1)
	goto L99
L107:
	;
	goto L108
L108:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v331))))
	if int32(252) < v355 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331 + int32(1)
	goto L113
L110:
	;
	v357 = v355 - int32(97)
	if v357 < int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v360 = int32(1)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v357)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v364)>>(uint(v357&int32(7))%32))&v360 != 0 {
		v377 = v360
		goto L103
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	goto L105
L114:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v389 < v388 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v429 < int32(0) {
		goto L98
	} else {
		goto L130
	}
L116:
	;
	v391 = v388
	goto L118
L117:
	;
	v391 = v389
	goto L118
L118:
	;
	v398 = v388
	goto L120
L119:
	;
	v429 = v409
	goto L115
L120:
	;
	if v398 == v391 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v429 = int32(-1)
	goto L115
L123:
	;
	goto L124
L124:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v398))))
	if int32(252) < v404 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v421 = v398 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v421
	v398 = v421
	goto L120
L126:
	;
	v406 = v404 - int32(97)
	if v406 < int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v409 = int32(1)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v406)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v413)>>(uint(v406&int32(7))%32))&v409 != 0 {
		goto L119
	} else {
		goto L128
	}
L128:
	;
	goto L125
L130:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v496 = v432 + v429
	goto L2
L131:
	;
	if v488 != 0 {
		goto L1
	} else {
		goto L145
	}
L132:
	;
	v447 = v331
	goto L134
L133:
	;
	v447 = v445
	goto L134
L134:
	;
	goto L136
L135:
	;
	v488 = v483
	goto L131
L136:
	;
	if v331 == v447 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v483 = int32(0)
	goto L135
L138:
	;
	v488 = int32(-1)
	goto L131
L139:
	;
	goto L140
L140:
	;
	v459 = int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460+v331))))
	if int32(252) < v462 {
		v483 = v459
		goto L135
	} else {
		goto L141
	}
L141:
	;
	v464 = v462 - int32(97)
	if v464 < int32(0) {
		v483 = v459
		goto L135
	} else {
		goto L142
	}
L142:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v464)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v470)>>(uint(v464&int32(7))%32))&int32(1) == int32(0) {
		v483 = v459
		goto L135
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331 + int32(1)
	goto L144
L144:
	;
	goto L137
L145:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v490 <= v489 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v496 = v489 + int32(1)
	goto L2
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v728
	v732 = v728 - int32(1)
	if v732 <= v10 {
		goto L211
	} else {
		goto L212
	}
L148:
	;
	if v549 < int32(0) {
		goto L147
	} else {
		goto L163
	}
L149:
	;
	v511 = v10
	goto L151
L150:
	;
	v511 = v509
	goto L151
L151:
	;
	v518 = v10
	goto L153
L152:
	;
	v549 = v529
	goto L148
L153:
	;
	if v518 == v511 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v549 = int32(-1)
	goto L148
L156:
	;
	goto L157
L157:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522+v518))))
	if int32(252) < v524 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v541 = v518 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v541
	v518 = v541
	goto L153
L159:
	;
	v526 = v524 - int32(97)
	if v526 < int32(0) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v529 = int32(1)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v526)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v533)>>(uint(v526&int32(7))%32))&v529 != 0 {
		goto L152
	} else {
		goto L161
	}
L161:
	;
	goto L158
L163:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v553 = v552 + v549
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v553
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v564 < v553 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v607 < int32(0) {
		goto L147
	} else {
		goto L178
	}
L165:
	;
	v566 = v553
	goto L167
L166:
	;
	v566 = v564
	goto L167
L167:
	;
	v572 = v553
	goto L169
L168:
	;
	v607 = int32(1)
	goto L164
L169:
	;
	if v572 == v566 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v607 = int32(-1)
	goto L164
L172:
	;
	goto L173
L173:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579+v572))))
	if int32(252) < v581 {
		goto L168
	} else {
		goto L174
	}
L174:
	;
	v583 = v581 - int32(97)
	if v583 < int32(0) {
		goto L168
	} else {
		goto L175
	}
L175:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v583)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v589)>>(uint(v583&int32(7))%32))&int32(1) == int32(0) {
		goto L168
	} else {
		goto L176
	}
L176:
	;
	v598 = v572 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v598
	v572 = v598
	goto L169
L178:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v611 = v610 + v607
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v611
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v622 < v611 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v662 < int32(0) {
		goto L147
	} else {
		goto L194
	}
L180:
	;
	v624 = v611
	goto L182
L181:
	;
	v624 = v622
	goto L182
L182:
	;
	v631 = v611
	goto L184
L183:
	;
	v662 = v642
	goto L179
L184:
	;
	if v631 == v624 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v662 = int32(-1)
	goto L179
L187:
	;
	goto L188
L188:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635+v631))))
	if int32(252) < v637 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v654 = v631 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v654
	v631 = v654
	goto L184
L190:
	;
	v639 = v637 - int32(97)
	if v639 < int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v642 = int32(1)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v639)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v646)>>(uint(v639&int32(7))%32))&v642 != 0 {
		goto L183
	} else {
		goto L192
	}
L192:
	;
	goto L189
L194:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v666 = v665 + v662
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v666
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v677 < v666 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v720 < int32(0) {
		goto L147
	} else {
		goto L209
	}
L196:
	;
	v679 = v666
	goto L198
L197:
	;
	v679 = v677
	goto L198
L198:
	;
	v685 = v666
	goto L200
L199:
	;
	v720 = int32(1)
	goto L195
L200:
	;
	if v685 == v679 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v720 = int32(-1)
	goto L195
L203:
	;
	goto L204
L204:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692+v685))))
	if int32(252) < v694 {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	v696 = v694 - int32(97)
	if v696 < int32(0) {
		goto L199
	} else {
		goto L206
	}
L206:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v696)>>(uint(int32(3))%32)))+uint32(_c_F_spanish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v702)>>(uint(v696&int32(7))%32))&int32(1) == int32(0) {
		goto L199
	} else {
		goto L207
	}
L207:
	;
	v711 = v685 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v711
	v685 = v711
	goto L200
L209:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v723 + v720
	goto L147
L210:
	;
	return v1339
L211:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v835
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v835-int32(2) <= v838 {
		goto L244
	} else {
		goto L245
	}
L212:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734+v732))))
	if base.B2i32(v736&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v736)%32)&int32(_a_F_spanish_ISO_8859_1_stem_0) == int32(0)) != 0 {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v751 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_1), int32(13), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	return int32(0)
L215:
	;
	if v751 == int32(0) {
		goto L211
	} else {
		goto L216
	}
L216:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v757
	v760 = v757 - int32(1)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v760 <= v761 {
		goto L211
	} else {
		goto L217
	}
L217:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763+v760))))
	switch v765 - int32(111) {
	case 0, 3:
		goto L218
	default:
		goto L211
	}
L218:
	;
	v771 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_2), int32(11), int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L214
	} else {
		goto L219
	}
L219:
	;
	if v771 == int32(0) {
		goto L211
	} else {
		goto L220
	}
L220:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v775 < v776 {
		goto L211
	} else {
		goto L221
	}
L221:
	;
	switch v771 - int32(1) {
	case 0:
		goto L228
	case 1:
		goto L227
	case 2:
		goto L226
	case 3:
		goto L225
	case 4:
		goto L224
	case 5:
		goto L223
	case 6:
		goto L222
	default:
		goto L211
	}
L222:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v775 <= v818 {
		goto L211
	} else {
		goto L240
	}
L223:
	;
	v815 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v815 {
		goto L211
	} else {
		goto L239
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v775
	v811 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_ISO_8859_1_stem_3))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L214
	} else {
		goto L237
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v775
	v804 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_ISO_8859_1_stem_4))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L214
	} else {
		goto L235
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v775
	v797 = F_slice_from_s(m, l0, int32(2), int32(_a_F_spanish_ISO_8859_1_stem_5))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L214
	} else {
		goto L233
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v775
	v790 = F_slice_from_s(m, l0, int32(4), int32(_a_F_spanish_ISO_8859_1_stem_6))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L214
	} else {
		goto L231
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v775
	v783 = F_slice_from_s(m, l0, int32(5), int32(_a_F_spanish_ISO_8859_1_stem_7))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L214
	} else {
		goto L229
	}
L229:
	;
	if int32(0) <= v783 {
		goto L211
	} else {
		goto L230
	}
L230:
	;
	v1339 = v783
	goto L210
L231:
	;
	if int32(0) <= v790 {
		goto L211
	} else {
		goto L232
	}
L232:
	;
	v1339 = v790
	goto L210
L233:
	;
	if int32(0) <= v797 {
		goto L211
	} else {
		goto L234
	}
L234:
	;
	v1339 = v797
	goto L210
L235:
	;
	if int32(0) <= v804 {
		goto L211
	} else {
		goto L236
	}
L236:
	;
	v1339 = v804
	goto L210
L237:
	;
	if int32(0) <= v811 {
		goto L211
	} else {
		goto L238
	}
L238:
	;
	v1339 = v811
	goto L210
L239:
	;
	v1339 = v815
	goto L210
L240:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820+v775-int32(1)))))
	if v824 != int32(117) {
		goto L211
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v775 - int32(1)
	v830 = F_slice_del(m, l0)
	mBase = m.M
	if v830 < int32(0) {
		v1339 = v830
		goto L210
	} else {
		goto L242
	}
L242:
	;
	goto L211
L243:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1206
	v1212 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_8), int32(8), int32(0))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L214
	} else {
		goto L355
	}
L244:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1103
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1103 < v1105 {
		goto L319
	} else {
		goto L320
	}
L245:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v844 = int32(1)
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842+v835-v844))))
	if base.B2i32(v846&int32(224) != int32(96))|base.B2i32(v844<<(uint(v846)%32)&int32(_a_F_spanish_ISO_8859_1_stem_9) == int32(0)) != 0 {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v861 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_10), int32(48), int32(0))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L214
	} else {
		goto L247
	}
L247:
	;
	if v861 == int32(0) {
		goto L244
	} else {
		goto L248
	}
L248:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v865
	switch v861 - int32(1) {
	case 0:
		goto L257
	case 1:
		goto L256
	case 2:
		goto L255
	case 3:
		goto L254
	case 4:
		goto L253
	case 5:
		goto L252
	case 6:
		goto L251
	case 7:
		goto L250
	case 8:
		goto L249
	default:
		goto L243
	}
L249:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 < v1069 {
		goto L244
	} else {
		goto L310
	}
L250:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 < v1030 {
		goto L244
	} else {
		goto L302
	}
L251:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 < v998 {
		goto L244
	} else {
		goto L294
	}
L252:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v865 < v930 {
		goto L244
	} else {
		goto L278
	}
L253:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 < v922 {
		goto L244
	} else {
		goto L275
	}
L254:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 < v914 {
		goto L244
	} else {
		goto L272
	}
L255:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 < v906 {
		goto L244
	} else {
		goto L269
	}
L256:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 < v874 {
		goto L244
	} else {
		goto L260
	}
L257:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 < v869 {
		goto L244
	} else {
		goto L258
	}
L258:
	;
	v871 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v871 {
		goto L243
	} else {
		goto L259
	}
L259:
	;
	v1339 = v871
	goto L210
L260:
	;
	v876 = F_slice_del(m, l0)
	mBase = m.M
	if v876 < int32(0) {
		v1339 = v876
		goto L210
	} else {
		goto L261
	}
L261:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v879
	v881 = int32(2)
	v883 = int32(0)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v879-v886 < v881 {
		v896 = v883
		goto L263
	} else {
		goto L264
	}
L262:
	;
	if v896 == int32(0) {
		goto L243
	} else {
		goto L266
	}
L263:
	;
	goto L262
L264:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v892 = F_memcmp(m, v889+v879-v881, int32(_a_F_spanish_ISO_8859_1_stem_11), v881)
	mBase = m.M
	if v892 != 0 {
		v896 = v883
		goto L263
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v879 - v881
	v896 = int32(1)
	goto L263
L266:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v899
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v899 < v901 {
		goto L243
	} else {
		goto L267
	}
L267:
	;
	v903 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v903 {
		goto L243
	} else {
		goto L268
	}
L268:
	;
	v1339 = v903
	goto L210
L269:
	;
	v910 = F_slice_from_s(m, l0, int32(3), int32(_a_F_spanish_ISO_8859_1_stem_12))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L214
	} else {
		goto L270
	}
L270:
	;
	if int32(0) <= v910 {
		goto L243
	} else {
		goto L271
	}
L271:
	;
	v1339 = v910
	goto L210
L272:
	;
	v918 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_ISO_8859_1_stem_13))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L214
	} else {
		goto L273
	}
L273:
	;
	if int32(0) <= v918 {
		goto L243
	} else {
		goto L274
	}
L274:
	;
	v1339 = v918
	goto L210
L275:
	;
	v926 = F_slice_from_s(m, l0, int32(4), int32(_a_F_spanish_ISO_8859_1_stem_14))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L214
	} else {
		goto L276
	}
L276:
	;
	if int32(0) <= v926 {
		goto L243
	} else {
		goto L277
	}
L277:
	;
	v1339 = v926
	goto L210
L278:
	;
	v932 = F_slice_del(m, l0)
	mBase = m.M
	if v932 < int32(0) {
		v1339 = v932
		goto L210
	} else {
		goto L279
	}
L279:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v935
	v938 = v935 - int32(1)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v938 <= v939 {
		goto L243
	} else {
		goto L280
	}
L280:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v941+v938))))
	if base.B2i32(v943&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v943)%32)&int32(_a_F_spanish_ISO_8859_1_stem_15) == int32(0)) != 0 {
		goto L243
	} else {
		goto L281
	}
L281:
	;
	v958 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_16), int32(4), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L214
	} else {
		goto L282
	}
L282:
	;
	if v958 == int32(0) {
		goto L243
	} else {
		goto L283
	}
L283:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v962
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v962 < v964 {
		goto L243
	} else {
		goto L284
	}
L284:
	;
	v966 = F_slice_del(m, l0)
	mBase = m.M
	if v966 < int32(0) {
		v1339 = v966
		goto L210
	} else {
		goto L285
	}
L285:
	;
	if v958 != int32(1) {
		goto L243
	} else {
		goto L286
	}
L286:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v971
	v973 = int32(2)
	v975 = int32(0)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v971-v978 < v973 {
		v988 = v975
		goto L288
	} else {
		goto L289
	}
L287:
	;
	if v988 == int32(0) {
		goto L243
	} else {
		goto L291
	}
L288:
	;
	goto L287
L289:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v984 = F_memcmp(m, v981+v971-v973, int32(_a_F_spanish_ISO_8859_1_stem_17), v973)
	mBase = m.M
	if v984 != 0 {
		v988 = v975
		goto L288
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v971 - v973
	v988 = int32(1)
	goto L288
L291:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v991
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v991 < v993 {
		goto L243
	} else {
		goto L292
	}
L292:
	;
	v995 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v995 {
		goto L243
	} else {
		goto L293
	}
L293:
	;
	v1339 = v995
	goto L210
L294:
	;
	v1000 = F_slice_del(m, l0)
	mBase = m.M
	if v1000 < int32(0) {
		v1339 = v1000
		goto L210
	} else {
		goto L295
	}
L295:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1003
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1003-int32(3) <= v1005 {
		goto L243
	} else {
		goto L296
	}
L296:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009+v1003-int32(1)))))
	if v1013 != int32(101) {
		goto L243
	} else {
		goto L297
	}
L297:
	;
	v1019 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_18), int32(3), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L214
	} else {
		goto L298
	}
L298:
	;
	if v1019 == int32(0) {
		goto L243
	} else {
		goto L299
	}
L299:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1023
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1023 < v1025 {
		goto L243
	} else {
		goto L300
	}
L300:
	;
	v1027 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1027 {
		goto L243
	} else {
		goto L301
	}
L301:
	;
	v1339 = v1027
	goto L210
L302:
	;
	v1032 = F_slice_del(m, l0)
	mBase = m.M
	if v1032 < int32(0) {
		v1339 = v1032
		goto L210
	} else {
		goto L303
	}
L303:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1035
	v1038 = v1035 - int32(1)
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1038 <= v1039 {
		goto L243
	} else {
		goto L304
	}
L304:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1041+v1038))))
	if base.B2i32(v1043&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1043)%32)&int32(_a_F_spanish_ISO_8859_1_stem_19) == int32(0)) != 0 {
		goto L243
	} else {
		goto L305
	}
L305:
	;
	v1058 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_20), int32(3), int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L214
	} else {
		goto L306
	}
L306:
	;
	if v1058 == int32(0) {
		goto L243
	} else {
		goto L307
	}
L307:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1062
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1062 < v1064 {
		goto L243
	} else {
		goto L308
	}
L308:
	;
	v1066 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1066 {
		goto L243
	} else {
		goto L309
	}
L309:
	;
	v1339 = v1066
	goto L210
L310:
	;
	v1071 = F_slice_del(m, l0)
	mBase = m.M
	if v1071 < int32(0) {
		v1339 = v1071
		goto L210
	} else {
		goto L311
	}
L311:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1074
	v1076 = int32(2)
	v1078 = int32(0)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1074-v1081 < v1076 {
		v1091 = v1078
		goto L313
	} else {
		goto L314
	}
L312:
	;
	if v1091 == int32(0) {
		goto L243
	} else {
		goto L316
	}
L313:
	;
	goto L312
L314:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1087 = F_memcmp(m, v1084+v1074-v1076, int32(_a_F_spanish_ISO_8859_1_stem_21), v1076)
	mBase = m.M
	if v1087 != 0 {
		v1091 = v1078
		goto L313
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1074 - v1076
	v1091 = int32(1)
	goto L313
L316:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1094
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1094 < v1096 {
		goto L243
	} else {
		goto L317
	}
L317:
	;
	v1098 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1098 {
		goto L243
	} else {
		goto L318
	}
L318:
	;
	v1339 = v1098
	goto L210
L319:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1153
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1153 < v1155 {
		goto L243
	} else {
		goto L338
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1103
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1105
	v1113 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_22), int32(12), int32(0))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L214
	} else {
		goto L321
	}
L321:
	;
	if v1113 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1108
	goto L319
L323:
	;
	goto L324
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1108
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1119
	if v1119 <= v1108 {
		goto L319
	} else {
		goto L325
	}
L325:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122+v1119-int32(1)))))
	if v1126 != int32(117) {
		goto L319
	} else {
		goto L326
	}
L326:
	;
	v1129 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1119 - v1129
	v1133 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1133 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1144 = int32(0)
	v1145 = base.B2i32(v1139 < v1144)
	if v1145 == v1144 {
		goto L243
	} else {
		goto L334
	}
L328:
	;
	v1139 = v1129
	goto L330
L329:
	;
	v1139 = v1133 >> (uint(int32(31)) % 32) & v1133
	goto L330
L330:
	;
	if v1139 != 0 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1143 = int32(base.Ui32(v1139) >> (uint(int32(31)) % 32))
	goto L333
L332:
	;
	v1143 = int32(5)
	goto L333
L333:
	;
	switch v1143 {
	case 0:
		goto L243
	default:
		goto L327
	case 5:
		goto L319
	}
L334:
	;
	if v1139 < v1144 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1149 = v1139
	goto L337
L336:
	;
	v1149 = int32(1)
	goto L337
L337:
	;
	return v1149
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1153
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1155
	v1163 = F_find_among_b(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_23), int32(96), int32(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L214
	} else {
		goto L339
	}
L339:
	;
	if v1163 == int32(0) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1158
	goto L243
L341:
	;
	goto L342
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1158
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1169
	switch v1163 - int32(1) {
	case 0:
		goto L344
	case 1:
		goto L343
	default:
		goto L243
	}
L343:
	;
	v1199 = F_slice_del(m, l0)
	mBase = m.M
	if v1199 < int32(0) {
		v1339 = v1199
		goto L210
	} else {
		goto L353
	}
L344:
	;
	if v1169 <= v1158 {
		v1193 = v1169
		goto L345
	} else {
		goto L346
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1193
	v1196 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1196 {
		goto L243
	} else {
		goto L352
	}
L346:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1175 = v1174 + v1169
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175-int32(1)))))
	if v1178 != int32(117) {
		v1193 = v1169
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1182 = v1169 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1182
	if v1182 <= v1158 {
		v1193 = v1169
		goto L345
	} else {
		goto L348
	}
L348:
	;
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175-int32(2)))))
	if v1187 == int32(103) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1190 = v1182
	goto L351
L350:
	;
	v1190 = v1169
	goto L351
L351:
	;
	v1193 = v1190
	goto L345
L352:
	;
	v1339 = v1196
	goto L210
L353:
	;
	goto L243
L354:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1260
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1265 = v1260
	v1267 = v1262
	goto L369
L355:
	;
	if v1212 == int32(0) {
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1216
	switch v1212 - int32(1) {
	case 0:
		goto L358
	case 1:
		goto L357
	default:
		goto L354
	}
L357:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1216 < v1225 {
		goto L354
	} else {
		goto L361
	}
L358:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1216 < v1220 {
		goto L354
	} else {
		goto L359
	}
L359:
	;
	v1222 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1222 {
		goto L354
	} else {
		goto L360
	}
L360:
	;
	v1339 = v1222
	goto L210
L361:
	;
	v1227 = F_slice_del(m, l0)
	mBase = m.M
	if v1227 < int32(0) {
		v1339 = v1227
		goto L210
	} else {
		goto L362
	}
L362:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1230
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1230 <= v1232 {
		goto L354
	} else {
		goto L363
	}
L363:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1235 = v1234 + v1230
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1235-int32(1)))))
	if v1238 != int32(117) {
		goto L354
	} else {
		goto L364
	}
L364:
	;
	v1242 = v1230 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1242
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1242
	if v1242 <= v1232 {
		goto L354
	} else {
		goto L365
	}
L365:
	;
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1235-int32(2)))))
	if v1248 != int32(103) {
		goto L354
	} else {
		goto L366
	}
L366:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1230 <= v1251 {
		goto L354
	} else {
		goto L367
	}
L367:
	;
	v1253 = F_slice_del(m, l0)
	mBase = m.M
	if v1253 < int32(0) {
		v1339 = v1253
		goto L210
	} else {
		goto L368
	}
L368:
	;
	goto L354
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1265
	if v1267 <= v1265 {
		goto L375
	} else {
		goto L376
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1260
	v1339 = int32(1)
	goto L210
L371:
	;
	goto L370
L372:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1265 = v1334
	v1267 = v1333
	goto L369
L373:
	;
	if v1325 <= v1323 {
		goto L371
	} else {
		goto L396
	}
L374:
	;
	v1286 = F_find_among(m, l0, int32(_a_F_spanish_ISO_8859_1_stem_24), int32(6), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L214
	} else {
		goto L379
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1265
	v1323 = v1265
	v1325 = v1267
	goto L373
L376:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1270+v1265))))
	v1273 = int32(224)
	if v1272&v1273 != v1273 {
		goto L375
	} else {
		goto L377
	}
L377:
	;
	if int32(1)<<(uint(v1272)%32)&int32(67641858) != 0 {
		goto L374
	} else {
		goto L378
	}
L378:
	;
	goto L375
L379:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1288
	switch v1286 - int32(1) {
	case 0:
		goto L385
	case 1:
		goto L384
	case 2:
		goto L383
	case 3:
		goto L382
	case 4:
		goto L381
	case 5:
		goto L380
	default:
		goto L372
	}
L380:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1323 = v1288
	v1325 = v1322
	goto L373
L381:
	;
	v1318 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_ISO_8859_1_stem_25))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L214
	} else {
		goto L394
	}
L382:
	;
	v1312 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_ISO_8859_1_stem_26))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L214
	} else {
		goto L392
	}
L383:
	;
	v1306 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_ISO_8859_1_stem_27))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L214
	} else {
		goto L390
	}
L384:
	;
	v1300 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_ISO_8859_1_stem_28))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L214
	} else {
		goto L388
	}
L385:
	;
	v1294 = F_slice_from_s(m, l0, int32(1), int32(_a_F_spanish_ISO_8859_1_stem_29))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L214
	} else {
		goto L386
	}
L386:
	;
	if int32(0) <= v1294 {
		goto L372
	} else {
		goto L387
	}
L387:
	;
	v1339 = v1294
	goto L210
L388:
	;
	if int32(0) <= v1300 {
		goto L372
	} else {
		goto L389
	}
L389:
	;
	v1339 = v1300
	goto L210
L390:
	;
	if int32(0) <= v1306 {
		goto L372
	} else {
		goto L391
	}
L391:
	;
	v1339 = v1306
	goto L210
L392:
	;
	if int32(0) <= v1312 {
		goto L372
	} else {
		goto L393
	}
L393:
	;
	v1339 = v1312
	goto L210
L394:
	;
	if int32(0) <= v1318 {
		goto L372
	} else {
		goto L395
	}
L395:
	;
	v1339 = v1318
	goto L210
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1323 + int32(1)
	goto L372
}
func F_spgbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
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
	v4 = F_smgr_bulk_start_rel(m, l0, int32(3))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = F_smgr_bulk_get_buf(m, v4)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_PageInit(m, v6, int32(_a_F_spgbuildempty_0), int32(8))
			mBase = m.M
			v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+16)))
			v13 = v6 + v12
			v14 = int32(_a_F_spgbuildempty_1)
			*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)) = uint16(v14)
			v16 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v13))) = uint16(v16)
			v18 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+80)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+72)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+64)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+48)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v6)+88)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = int64(-1173640210)
			v36 = int32(92)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)) = uint16(v36)
			v38 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v38
			F_smgr_bulk_write(m, v4, int32(0), v6, int32(1))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				v56 = F_smgr_bulk_get_buf(m, v4)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = int32(4)
					F_PageInit(m, v56, int32(_a_F_spgbuildempty_0), int32(8))
					mBase = m.M
					v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+16)))
					v63 = v56 + v62
					v64 = int32(_a_F_spgbuildempty_1)
					*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v64)
					*(*uint16)(unsafe.Add(mBase, uint32(v63))) = uint16(v58)
					v67 = int32(1)
					F_smgr_bulk_write(m, v4, v67, v56, v67)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v71 = F_smgr_bulk_get_buf(m, v4)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v73 = int32(12)
							F_PageInit(m, v71, int32(_a_F_spgbuildempty_0), int32(8))
							mBase = m.M
							v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+16)))
							v78 = v71 + v77
							v79 = int32(_a_F_spgbuildempty_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v78)+6)) = uint16(v79)
							*(*uint16)(unsafe.Add(mBase, uint32(v78))) = uint16(v73)
							F_smgr_bulk_write(m, v4, int32(2), v71, int32(1))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_smgr_bulk_finish(m, v4)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_spgcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 float64
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v72 float64
	_ = v72
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v21 = v17 + int32(8)
	base.MemoryFill(m, v21, int32(0), int32(72))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	F_genericcostestimate(m, l0, l1, l2, v21)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
		if v29 < int32(0) {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			if base.Ui32(int32(2)) <= base.Ui32(v33) {
				v37 = F_log(m, base.F64_convert_i32_u(v33))
				mBase = m.M
				v41 = base.I32_trunc_sat_f64_s(base.F64_div(v37, float64(4.605170185988092)))
			} else {
				v41 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v41
			v44 = v41
		} else {
			v44 = v29
		}
		v46 = *(*float64)(unsafe.Add(mBase, _c_F_spgcostestimate[0]))
		v47 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
		v48 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
		if base.F64_gt(v48, float64(1)) == int32(0) {
			v53 = *(*float64)(unsafe.Add(mBase, uint32(v17)+64))
			v54 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
			v63 = v47
			v64 = v53
			v66 = v54
		} else {
			v55 = F_log(m, v48)
			mBase = m.M
			v57 = base.F64_mul(base.F64_ceil(v55), v46)
			v59 = *(*float64)(unsafe.Add(mBase, uint32(v17)+64))
			v61 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
			v63 = base.F64_add(v47, v57)
			v64 = v59
			v66 = base.F64_add(base.F64_mul(v59, v57), v61)
		}
		v72 = base.F64_mul(v46, base.F64_mul(base.F64_convert_i32_s(v44+int32(1)), float64(50)))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v63, v72)
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(v64, v72), v66)
		v78 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v78
		v80 = *(*float64)(unsafe.Add(mBase, uint32(v17)+32))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v80
		v82 = *(*float64)(unsafe.Add(mBase, uint32(v17)+40))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v82
		m.G0 = v17 + int32(80)
		return
	}
}
func F_spginsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_spginsert[0]))
	v19 = F_AllocSetContextCreateInternal(m, v14, int32(_a_F_spginsert_0), int32(0), int32(_a_F_spginsert_1), int32(_a_F_spginsert_2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(_a_F_spginsert_3)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_spginsert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_spginsert[0])) = v19
	v28 = v11 + int32(12)
	F_initSpGistState(m, v28, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = F_spgdoinsert(m, l0, v28, l3, l1, l2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v31 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_SpGistUpdateMetaPage(m, l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	F_MemoryContextReset(m, v19)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v46 = v11 + int32(12)
	F_initSpGistState(m, v46, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v49 = F_spgdoinsert(m, l0, v46, l3, l1, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v49 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spginsert[0])) = v24
	F_MemoryContextDelete(m, v19)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	m.G0 = v11 + int32(96)
	return int32(0)
}
func F_spgproperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(1)
	v21 = base.B2i32(l2 == int32(6)) & base.B2i32(l1 != int32(0))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = F_get_index_column_opclass(m, l0, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v14 + int32(16)
	return v21
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v112)
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	if v22 == int32(0) {
		v112 = v16
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v32 = F_get_opclass_opfamily_and_input_type(m, v22, v14+int32(12), v14+int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v32 == int32(0) {
		v112 = v16
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v14)+12)))
	v40 = int64(0)
	v42 = F_SearchSysCacheList(m, int32(4), int32(1), v39, v40, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v44)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+56))
	if v46 <= v44 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_ReleaseCatCacheList(m, v42)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L26
	}
L12:
	;
	v52 = v46
	v53 = int32(0)
	goto L13
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v42-int32(-64)+v53<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
	v69 = v67 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+18)))
	if v70 != int32(111) {
		v85 = v52
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v90)
	goto L11
L15:
	;
	goto L14
L16:
	;
	v88 = v53 + int32(1)
	if v88 < v85 {
		v52 = v85
		v53 = v88
		goto L13
	} else {
		goto L25
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v73 != v74 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v76 != v73 {
		v85 = v52
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v80 = F_get_op_rettype(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v82 = F_opfamily_can_sort_type(m, v78, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v82 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v42)+56))
	v85 = v84
	goto L16
L25:
	;
	goto L11
L26:
	;
	v112 = int32(0)
	goto L4
}
func F_ssup_datum_signed_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	return base.B2i32(l1 < l0) - base.B2i32(l0 < l1)
}
func F_stack_is_too_deep(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_stack_is_too_deep[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_stack_is_too_deep[1]))
	v7 = m.G0
	v8 = v6 - v7
	v10 = v8 >> (uint(int32(31)) % 32)
	return base.B2i32(v4 < v8^v10-v10) & base.B2i32(v6 != int32(0))
}
func F_statatt_init_empty_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v21 int32
	_ = v21
	var v31 int64
	_ = v31
	v7 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+29)) = uint16(v7)
	v9 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(l4)+21)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l5)+23)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l5)+16)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = base.I64_extend_i32_u(l0)
	v21 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = base.I64_extend_i32_s(l1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = base.I64_extend_i32_u(l2)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+2)) = uint8(v21)
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+32)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+40)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+48)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+6)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+88)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+11)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+128)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+56)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+7)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+96)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+12)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+136)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+17)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+64)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+104)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+13)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+144)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+18)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+72)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+9)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+112)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+14)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+152)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+19)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+80)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+10)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+120)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)) = uint8(v21)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+160)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+20)) = uint8(v21)
	return
}
func F_statistic_proc_security_check(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v10 != 0 {
		v36 = v9
		m.G0 = v7 + int32(16)
		return v36
	} else {
		if l1 == int32(0) {
			v36 = int32(0)
			m.G0 = v7 + int32(16)
			return v36
		} else {
			v14 = F_get_func_leakproof(m, l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v36 = v9
					m.G0 = v7 + int32(16)
					return v36
				} else {
					v18 = int32(0)
					v21 = F_errstart(m, int32(13), v18)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						if v21 == int32(0) {
							v36 = v18
							m.G0 = v7 + int32(16)
							return v36
						} else {
							v25 = F_get_func_name(m, l1)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v25
								F_errmsg_internal(m, int32(_a_F_statistic_proc_security_check_0), v7)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_statistic_proc_security_check_1), int32(_a_F_statistic_proc_security_check_2), int32(_a_F_statistic_proc_security_check_3))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										v36 = v18
										m.G0 = v7 + int32(16)
										return v36
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_storeGettuple(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	v4 = l3
	v6 = l5
	v7 = l6
	v12 = m.G0
	v14 = v12 - int32(288)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v19 = l0 + v16*int32(6)
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+228)) = uint16(v20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v24+int32(2672)))) = uint8(v6)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v29+int32(3080)))) = uint8(v7)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if int32(0) < v34 {
		v37 = int32(0)
		if v4 != 0 {
			v89 = v37
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
			*(*int32)(unsafe.Add(mBase, uint32(l0+v95<<(uint(int32(2))%32))+uint32(_c_F_storeGettuple[0]))) = v89
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
			if v113 == int32(1) {
				v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
				if int32(2) <= v117 {
					v121 = v14 + int32(32)
					v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
					if v4 == int32(0) {
						F_index_deform_tuple_internal(m, v116, v121, v14, l4+int32(16), l4+int32(12), int32(base.Ui32(v123&int32(_a_F_storeGettuple_0))>>(uint(int32(15))%32)))
						mBase = m.M
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
						if v126 != int32(1) {
							F_index_deform_tuple_internal(m, v116, v121, v14, l4+int32(16), l4+int32(12), int32(base.Ui32(v123&int32(_a_F_storeGettuple_0))>>(uint(int32(15))%32)))
							mBase = m.M
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v121))) = int64(0)
							v131 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v131)
						}
					}
				} else {
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
				*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = l2
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
				v147 = F_heap_form_tuple(m, v144, v14+int32(32), v14)
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return
				} else {
					v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0+v149<<(uint(int32(2))%32)+int32(3488)))) = v147
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v157 + int32(1)
					m.G0 = v14 + int32(288)
					return
				}
			} else {
				v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v157 + int32(1)
				m.G0 = v14 + int32(288)
				return
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			if v38 <= int32(0) {
				v89 = v37
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				*(*int32)(unsafe.Add(mBase, uint32(l0+v95<<(uint(int32(2))%32))+uint32(_c_F_storeGettuple[0]))) = v89
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
				if v113 == int32(1) {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
					if int32(2) <= v117 {
						v121 = v14 + int32(32)
						v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
						if v4 == int32(0) {
							F_index_deform_tuple_internal(m, v116, v121, v14, l4+int32(16), l4+int32(12), int32(base.Ui32(v123&int32(_a_F_storeGettuple_0))>>(uint(int32(15))%32)))
							mBase = m.M
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
							if v126 != int32(1) {
								F_index_deform_tuple_internal(m, v116, v121, v14, l4+int32(16), l4+int32(12), int32(base.Ui32(v123&int32(_a_F_storeGettuple_0))>>(uint(int32(15))%32)))
								mBase = m.M
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v121))) = int64(0)
								v131 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v131)
							}
						}
					} else {
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
					*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = l2
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
					v147 = F_heap_form_tuple(m, v144, v14+int32(32), v14)
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return
					} else {
						v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0+v149<<(uint(int32(2))%32)+int32(3488)))) = v147
						v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v157 + int32(1)
						m.G0 = v14 + int32(288)
						return
					}
				} else {
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v157 + int32(1)
					m.G0 = v14 + int32(288)
					return
				}
			} else {
				v42 = F_palloc_mul(m, int32(16), v34)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
					if v44 <= int32(0) {
						v89 = v42
					} else {
						v49 = int32(0)
						for {
							v61 = v42 + v49<<(uint(int32(4))%32)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v49<<(uint(int32(2))%32))))
							if v66 < int32(0) {
								v76 = float64(0)
								v77 = int32(1)
							} else {
								v74 = *(*float64)(unsafe.Add(mBase, uint32(l7+v66<<(uint(int32(3))%32))))
								v76 = v74
								v77 = int32(0)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v61)+8)) = uint8(v77)
							*(*float64)(unsafe.Add(mBase, uint32(v61))) = v76
							v81 = v49 + int32(1)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
							if v81 < v82 {
								v49 = v81
								continue
							} else {
								break
							}
							break
						}
						v89 = v42
					}
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0+v95<<(uint(int32(2))%32))+uint32(_c_F_storeGettuple[0]))) = v89
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
					if v113 == int32(1) {
						v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
						if int32(2) <= v117 {
							v121 = v14 + int32(32)
							v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
							if v4 == int32(0) {
								F_index_deform_tuple_internal(m, v116, v121, v14, l4+int32(16), l4+int32(12), int32(base.Ui32(v123&int32(_a_F_storeGettuple_0))>>(uint(int32(15))%32)))
								mBase = m.M
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
								if v126 != int32(1) {
									F_index_deform_tuple_internal(m, v116, v121, v14, l4+int32(16), l4+int32(12), int32(base.Ui32(v123&int32(_a_F_storeGettuple_0))>>(uint(int32(15))%32)))
									mBase = m.M
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v121))) = int64(0)
									v131 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v131)
								}
							}
						} else {
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
						*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = l2
						v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
						v147 = F_heap_form_tuple(m, v144, v14+int32(32), v14)
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return
						} else {
							v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0+v149<<(uint(int32(2))%32)+int32(3488)))) = v147
							v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v157 + int32(1)
							m.G0 = v14 + int32(288)
							return
						}
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v157 + int32(1)
						m.G0 = v14 + int32(288)
						return
					}
				}
			}
		}
	} else {
		v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
		if v113 == int32(1) {
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
			if int32(2) <= v117 {
				v121 = v14 + int32(32)
				v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
				if v4 == int32(0) {
					F_index_deform_tuple_internal(m, v116, v121, v14, l4+int32(16), l4+int32(12), int32(base.Ui32(v123&int32(_a_F_storeGettuple_0))>>(uint(int32(15))%32)))
					mBase = m.M
				} else {
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
					if v126 != int32(1) {
						F_index_deform_tuple_internal(m, v116, v121, v14, l4+int32(16), l4+int32(12), int32(base.Ui32(v123&int32(_a_F_storeGettuple_0))>>(uint(int32(15))%32)))
						mBase = m.M
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v121))) = int64(0)
						v131 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v131)
					}
				}
			} else {
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = l2
			v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
			v147 = F_heap_form_tuple(m, v144, v14+int32(32), v14)
			mBase = m.M
			v148 = m.ExcPending
			if v148 != 0 {
				return
			} else {
				v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				*(*int32)(unsafe.Add(mBase, uint32(l0+v149<<(uint(int32(2))%32)+int32(3488)))) = v147
				v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v157 + int32(1)
				m.G0 = v14 + int32(288)
				return
			}
		} else {
			v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v157 + int32(1)
			m.G0 = v14 + int32(288)
			return
		}
	}
}
func F_strict_word_similarity_dist_op(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14409(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_strip_noop_phvs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v2 = int32(0)
	if l0 == v2 {
		v30 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v30
L2:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 == int32(321) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v18 = l0
	goto L11
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v13 = F_expression_tree_walker_impl(m, l0, int32(931), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	return int32(0)
L9:
	;
	if v13 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	return l0
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 != int32(321) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v27 = F_expression_tree_mutator_impl(m, v18, int32(932), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L17
	}
L13:
	;
	goto L12
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v23 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v24 != 0 {
		v18 = v24
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v30 = v2
	goto L1
L17:
	;
	v30 = v27
	goto L1
}
func F_strtoint(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, int32(10), int64(2147483648))
	return base.I32_wrap_i64(v5)
}
func F_strtoull(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(-1))
	return v5
}
func F_strupper_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
	v8 = int32(0)
	v10 = F_convert_case(m, l0, l1, l2, l3, int32(2), v7, v8, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_subcolor(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9+l1<<(uint(int32(1))%32)))))
	v16 = v8 + v13*int32(24)
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+8)))
	if v17 != int32(_a_F_subcolor_0) {
		v42 = v17
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
		if v44 != 0 {
			v87 = int32(_a_F_subcolor_0)
			return base.I32_extend16_s(v87)
		} else {
			v46 = int32(_a_F_subcolor_0)
			if v13&v46 == v42&v46 {
				return v13
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v53 = int32(24)
				v55 = v52 + v13*v53
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
				*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56 - int32(1)
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v63 = base.I32_extend16_s(v42) * v53
				v64 = v60 + v63
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
				if v65 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = l1
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v63)))
					v72 = v71
					v73 = v69
				} else {
					v72 = v65
					v73 = v60
				}
				v75 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v73+v63))) = v72 + v75
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*uint16)(unsafe.Add(mBase, uint32(v78+l1<<(uint(v75)%32)))) = uint16(v42)
				v87 = v42
				return base.I32_extend16_s(v87)
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v20+v21 == int32(1) {
			v42 = v13
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
			if v44 != 0 {
				v87 = int32(_a_F_subcolor_0)
				return base.I32_extend16_s(v87)
			} else {
				v46 = int32(_a_F_subcolor_0)
				if v13&v46 == v42&v46 {
					return v13
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v53 = int32(24)
					v55 = v52 + v13*v53
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56 - int32(1)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v63 = base.I32_extend16_s(v42) * v53
					v64 = v60 + v63
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
					if v65 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = l1
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v63)))
						v72 = v71
						v73 = v69
					} else {
						v72 = v65
						v73 = v60
					}
					v75 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v73+v63))) = v72 + v75
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*uint16)(unsafe.Add(mBase, uint32(v78+l1<<(uint(v75)%32)))) = uint16(v42)
					v87 = v42
					return base.I32_extend16_s(v87)
				}
			}
		} else {
			v25 = F_newcolor(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v25 == int32(-1) {
					v42 = int32(_a_F_subcolor_0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v33 = int32(24)
					*(*uint16)(unsafe.Add(mBase, uint32(v32+v13*v33)+8)) = uint16(v25)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*uint16)(unsafe.Add(mBase, uint32(v37+v25*v33)+8)) = uint16(v25)
					v42 = v25
				}
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
				if v44 != 0 {
					v87 = int32(_a_F_subcolor_0)
					return base.I32_extend16_s(v87)
				} else {
					v46 = int32(_a_F_subcolor_0)
					if v13&v46 == v42&v46 {
						return v13
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v53 = int32(24)
						v55 = v52 + v13*v53
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56 - int32(1)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v63 = base.I32_extend16_s(v42) * v53
						v64 = v60 + v63
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
						if v65 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = l1
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v63)))
							v72 = v71
							v73 = v69
						} else {
							v72 = v65
							v73 = v60
						}
						v75 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v73+v63))) = v72 + v75
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*uint16)(unsafe.Add(mBase, uint32(v78+l1<<(uint(v75)%32)))) = uint16(v42)
						v87 = v42
						return base.I32_extend16_s(v87)
					}
				}
			}
		}
	}
}
func F_subcoloronerow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	v6 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	if v13 <= v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v29 = v16 + l1*v13<<(uint(int32(1))%32)
	v31 = v6
	goto L3
L3:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29))))
	v34 = v32 * int32(24)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v36 = v34 + v35
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+8)))
	if v37 != int32(_a_F_subcoloronerow_0) {
		v58 = v37
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v61 != 0 {
		v84 = int32(_a_F_subcoloronerow_0)
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v40+v41 == int32(1) {
		v58 = v32
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v45 = F_newcolor(m, v12)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v45 == int32(-1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v58 = int32(_a_F_subcoloronerow_0)
	goto L5
L11:
	;
	goto L12
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v50+v34)+8)) = uint16(v45)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v53+v45*int32(24))+8)) = uint16(v45)
	v58 = v45
	goto L5
L13:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v85 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v62 = int32(_a_F_subcoloronerow_0)
	if v32&v62 == v58&v62 {
		v84 = v32
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v68 = v67 + v34
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v69 - v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v77 = v73 + base.I32_extend16_s(v58)*int32(24)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v78 + v70
	*(*uint16)(unsafe.Add(mBase, uint32(v29))) = uint16(v58)
	v84 = v58
	goto L13
L16:
	;
	v87 = v84 & int32(_a_F_subcoloronerow_0)
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	if v87 != v88 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_subcoloronerow[0]))
	if v92 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v184 = v31 + int32(1)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	if v184 < v185 {
		v29 = v29 + int32(2)
		v31 = v184
		goto L3
	} else {
		goto L47
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v95 <= v96 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L22
L24:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v168 != 0 {
		goto L1
	} else {
		goto L46
	}
L25:
	;
	F_createarc(m, v90, int32(112), base.I32_extend16_s(v84), l2, l3)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L8
	} else {
		goto L45
	}
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v98 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v120 == int32(0) {
		goto L25
	} else {
		goto L37
	}
L29:
	;
	v102 = v98
	goto L30
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	if v112 != l3 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L25
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v119 != 0 {
		v102 = v119
		goto L30
	} else {
		goto L36
	}
L33:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+4)))
	if v114 != v87 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v116 == int32(112) {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	goto L31
L37:
	;
	v124 = v120
	goto L38
L38:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	if v134 != l2 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L25
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v124)+24))
	if v141 != 0 {
		v124 = v141
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	if v136 != v87 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v138 == int32(112) {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	goto L39
L45:
	;
	goto L24
L46:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v84)
	goto L19
L47:
	;
	goto L4
}
func F_subvector(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 float32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 float32
	_ = v85
	var v88 int32
	_ = v88
	var v91 float32
	_ = v91
	var v94 int32
	_ = v94
	var v97 float32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 float32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if int32(0) < v16 {
			v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if int32(0) < v21 {
				if v19 < v21 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_subvector_0), int32(0))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_subvector_1), int32(1015), int32(_a_F_subvector_2))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
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
					v25 = v21
					if v19-v16 < v21 {
						v31 = v19 + int32(1)
					} else {
						v31 = v21 + v16
					}
					v32 = v31 - v25
					F_CheckDim_3(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v37 = F_mul_size(m, int32(4), v32)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v39 = F_add_size(m, int32(8), v37)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v41 = F_palloc0(m, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)) = uint16(v32)
									*(*int32)(unsafe.Add(mBase, uint32(v41))) = v39 << (uint(int32(2)) % 32)
									if v32 <= int32(0) {
									} else {
										v50 = v32 & int32(3)
										v52 = v41 + int32(8)
										v57 = v12 + int32(4) + v25<<(uint(int32(2))%32)
										v58 = int32(0)
										if base.Ui32(v25-v31) <= base.Ui32(int32(-4)) {
											v65 = v58
											v67 = int32(0)
											for {
												v76 = v65 << (uint(int32(2)) % 32)
												v79 = *(*float32)(unsafe.Add(mBase, uint32(v76+v57)))
												*(*float32)(unsafe.Add(mBase, uint32(v52+v76))) = v79
												v81 = int32(4)
												v82 = v76 | v81
												v85 = *(*float32)(unsafe.Add(mBase, uint32(v57+v82)))
												*(*float32)(unsafe.Add(mBase, uint32(v52+v82))) = v85
												v88 = v76 | int32(8)
												v91 = *(*float32)(unsafe.Add(mBase, uint32(v57+v88)))
												*(*float32)(unsafe.Add(mBase, uint32(v52+v88))) = v91
												v94 = v76 | int32(12)
												v97 = *(*float32)(unsafe.Add(mBase, uint32(v94+v57)))
												*(*float32)(unsafe.Add(mBase, uint32(v52+v94))) = v97
												v100 = v65 + v81
												v102 = v67 + v81
												if v102 != v32&int32(2147483644) {
													v65 = v100
													v67 = v102
													continue
												} else {
													break
												}
												break
											}
											if v50 == int32(0) {
											} else {
												v106 = v100
												v116 = v106
												v125 = int32(0)
												for {
													v127 = v116 << (uint(int32(2)) % 32)
													v130 = *(*float32)(unsafe.Add(mBase, uint32(v127+v57)))
													*(*float32)(unsafe.Add(mBase, uint32(v52+v127))) = v130
													v132 = int32(1)
													v135 = v125 + v132
													if v135 != v50 {
														v116 = v116 + v132
														v125 = v135
														continue
													} else {
														break
													}
													break
												}
											}
										} else {
											v106 = v58
											v116 = v106
											v125 = int32(0)
											for {
												v127 = v116 << (uint(int32(2)) % 32)
												v130 = *(*float32)(unsafe.Add(mBase, uint32(v127+v57)))
												*(*float32)(unsafe.Add(mBase, uint32(v52+v127))) = v130
												v132 = int32(1)
												v135 = v125 + v132
												if v135 != v50 {
													v116 = v116 + v132
													v125 = v135
													continue
												} else {
													break
												}
												break
											}
										}
									}
									return base.I64_extend_i32_u(v41)
								}
							}
						}
					}
				}
			} else {
				v25 = int32(1)
				if v19-v16 < v21 {
					v31 = v19 + int32(1)
				} else {
					v31 = v21 + v16
				}
				v32 = v31 - v25
				F_CheckDim_3(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v37 = F_mul_size(m, int32(4), v32)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v39 = F_add_size(m, int32(8), v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							v41 = F_palloc0(m, v39)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)) = uint16(v32)
								*(*int32)(unsafe.Add(mBase, uint32(v41))) = v39 << (uint(int32(2)) % 32)
								if v32 <= int32(0) {
								} else {
									v50 = v32 & int32(3)
									v52 = v41 + int32(8)
									v57 = v12 + int32(4) + v25<<(uint(int32(2))%32)
									v58 = int32(0)
									if base.Ui32(v25-v31) <= base.Ui32(int32(-4)) {
										v65 = v58
										v67 = int32(0)
										for {
											v76 = v65 << (uint(int32(2)) % 32)
											v79 = *(*float32)(unsafe.Add(mBase, uint32(v76+v57)))
											*(*float32)(unsafe.Add(mBase, uint32(v52+v76))) = v79
											v81 = int32(4)
											v82 = v76 | v81
											v85 = *(*float32)(unsafe.Add(mBase, uint32(v57+v82)))
											*(*float32)(unsafe.Add(mBase, uint32(v52+v82))) = v85
											v88 = v76 | int32(8)
											v91 = *(*float32)(unsafe.Add(mBase, uint32(v57+v88)))
											*(*float32)(unsafe.Add(mBase, uint32(v52+v88))) = v91
											v94 = v76 | int32(12)
											v97 = *(*float32)(unsafe.Add(mBase, uint32(v94+v57)))
											*(*float32)(unsafe.Add(mBase, uint32(v52+v94))) = v97
											v100 = v65 + v81
											v102 = v67 + v81
											if v102 != v32&int32(2147483644) {
												v65 = v100
												v67 = v102
												continue
											} else {
												break
											}
											break
										}
										if v50 == int32(0) {
										} else {
											v106 = v100
											v116 = v106
											v125 = int32(0)
											for {
												v127 = v116 << (uint(int32(2)) % 32)
												v130 = *(*float32)(unsafe.Add(mBase, uint32(v127+v57)))
												*(*float32)(unsafe.Add(mBase, uint32(v52+v127))) = v130
												v132 = int32(1)
												v135 = v125 + v132
												if v135 != v50 {
													v116 = v116 + v132
													v125 = v135
													continue
												} else {
													break
												}
												break
											}
										}
									} else {
										v106 = v58
										v116 = v106
										v125 = int32(0)
										for {
											v127 = v116 << (uint(int32(2)) % 32)
											v130 = *(*float32)(unsafe.Add(mBase, uint32(v127+v57)))
											*(*float32)(unsafe.Add(mBase, uint32(v52+v127))) = v130
											v132 = int32(1)
											v135 = v125 + v132
											if v135 != v50 {
												v116 = v116 + v132
												v125 = v135
												continue
											} else {
												break
											}
											break
										}
									}
								}
								return base.I64_extend_i32_u(v41)
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v152 = m.ExcPending
			if v152 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_subvector_0), int32(0))
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_subvector_1), int32(998), int32(_a_F_subvector_2))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
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
func F_summarizer_wal_segment_open(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v11 = m.G0
	v13 = v11 - int32(1072)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = v15
	v23 = int32(0)
	goto L3
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L18
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
	m.G0 = v13 + int32(1072)
	return
L3:
	;
	v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v21
	v30 = base.I64_div_u_s(int64(4294967296), v27)
	v31 = base.I64_div_u_s(l1, v30)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v31)
	v34 = l1 - v30*v31
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v34)
	v37 = v13 + int32(48)
	v42 = F_pg_snprintf(m, v37, int32(1024), int32(_a_F_summarizer_wal_segment_open_0), v13+int32(32))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v65 = m.G0
	v67 = v65 - int32(16)
	m.G0 = v67
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v63
	v72 = base.I64_div_u_s(int64(4294967296), base.I64_extend_i32_s(v64))
	v73 = base.I64_div_u_s(l1, v72)
	*(*uint32)(unsafe.Add(mBase, uint32(v67)+4)) = uint32(v73)
	v76 = l1 - v72*v73
	*(*uint32)(unsafe.Add(mBase, uint32(v67)+8)) = uint32(v76)
	v79 = v13 + int32(48)
	v82 = F_pg_snprintf(m, v79, int32(1024), int32(_a_F_summarizer_wal_segment_open_0), v67)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L13
	}
L5:
	;
	return
L6:
	;
	v45 = F_BasicOpenFile(m, v37, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1168)) = v45
	if int32(0) <= v45 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_summarizer_wal_segment_open[0]))
	if v51 != int32(44) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v23 < v54 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v23<<(uint(int32(2))%32))))
	v21 = v60
	v23 = v23 + int32(1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	goto L4
L13:
	;
	m.G0 = v67 + int32(16)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v79
	F_errmsg(m, int32(_a_F_summarizer_wal_segment_open_1), v13)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_summarizer_wal_segment_open_2), int32(1637), int32(_a_F_summarizer_wal_segment_open_3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(48)
	F_errmsg(m, int32(_a_F_summarizer_wal_segment_open_4), v13+int32(16))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_summarizer_wal_segment_open_2), int32(1623), int32(_a_F_summarizer_wal_segment_open_3))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_svals(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_svals(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_syncrep_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(_a_F_syncrep_yy_create_buffer_0)
			v15 = F_palloc(m, int32(_a_F_syncrep_yy_create_buffer_1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_4(m, int32(_a_F_syncrep_yy_create_buffer_2))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_syncrep_yy_create_buffer[0]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v37 == v24 {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v43 = v37 + v40<<(uint(int32(2))%32)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v8 != v44 {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v55)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v62 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32))))
						if v8 == v67 {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_syncrep_yy_create_buffer[0])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_4(m, int32(_a_F_syncrep_yy_create_buffer_2))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_syncrep_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_4(m, int32(_a_F_syncrep_yyensure_buffer_stack_0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_4(m, int32(_a_F_syncrep_yyensure_buffer_stack_0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
