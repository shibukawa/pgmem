package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OperatorIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v107
L2:
	;
	return int32(0)
L3:
	;
	if v13 == int32(0) {
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v19)
	v107 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(_a_F_OperatorIsVisibleExt_0), v9)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_OperatorIsVisibleExt_1), int32(2145), int32(_a_F_OperatorIsVisibleExt_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v13)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L35
	}
L15:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_OperatorIsVisibleExt[0]))
	if v45 == v43 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v90 = F_makeString(m, v39+int32(4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v84 == int32(0) {
		v104 = v43
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v84 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v52 <= int32(0) {
		v78 = v43
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v78
	goto L18
L23:
	;
	v55 = int32(0)
	if v55 < v52 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v58 = v52
	goto L26
L25:
	;
	v58 = v55
	goto L26
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v61 = int32(0)
	goto L27
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59+v61<<(uint(int32(2))%32))))
	v70 = base.B2i32(v69 == v40)
	if v69 == v40 {
		v78 = v70
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v78 = v70
	goto L22
L29:
	;
	v72 = v61 + int32(1)
	if v72 != v58 {
		v61 = v72
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v90
	v97 = F_list_make1_impl(m, int32(1), v9+int32(8))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	v101 = F_OpernameGetOprid(m, v97, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v104 = base.B2i32(v101 == l0)
	goto L14
L35:
	;
	v107 = v104
	goto L1
}
func F_OperatorUpd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l3 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = F_table_open(m, int32(2617), int32(3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l2 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L8:
	;
	v27 = F_SearchSysCacheCopy(m, int32(40), base.I64_extend_i32_u(l1), int64(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v27 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+92))
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v78
	F_CatalogTupleUpdate(m, v20, v27+int32(4), v27)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L28
	}
L12:
	;
	if v34 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if l0 == v34 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	v78 = v5
	goto L11
L16:
	;
	if v34 == int32(0) {
		v78 = l0
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v40 = F_get_opname(m, v34)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v50 = v33 + int32(4)
	if v40 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v50
	F_errmsg(m, int32(_a_F_OperatorUpd_4), v10+int32(-32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v50
	F_errmsg(m, int32(_a_F_OperatorUpd_5), v10+int32(-16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	F_errfinish(m, int32(_a_F_OperatorUpd_1), int32(745), int32(_a_F_OperatorUpd_2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F_errfinish(m, int32(_a_F_OperatorUpd_1), int32(740), int32(_a_F_OperatorUpd_2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L7
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v124
	F_errmsg(m, int32(_a_F_OperatorUpd_0), v10+int32(-48))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L53
	}
L31:
	;
	F_relation_close(m, v20, int32(3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L52
	}
L32:
	;
	v95 = F_SearchSysCacheCopy(m, int32(40), base.I64_extend_i32_u(l2), int64(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	if v95 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+22)))
	v101 = v99 + v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+96))
	if l3 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v102 == int32(0) {
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if l0 == v102 {
		goto L31
	} else {
		goto L41
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+96)) = int32(0)
	F_CatalogTupleUpdate(m, v20, v95+int32(4), v95)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	goto L31
L41:
	;
	if v102 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = F_get_opname(m, v102)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+96)) = l0
	F_CatalogTupleUpdate(m, v20, v95+int32(4), v95)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L51
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v124 = v101 + int32(4)
	if v114 != 0 {
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v101)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v124
	F_errmsg(m, int32(_a_F_OperatorUpd_3), v12)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_OperatorUpd_1), int32(813), int32(_a_F_OperatorUpd_2))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	goto L31
L52:
	;
	m.G0 = v12 - int32(-64)
	return
L53:
	;
	F_errfinish(m, int32(_a_F_OperatorUpd_1), int32(808), int32(_a_F_OperatorUpd_2))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makeOperatorDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2617)
	v16 = v12 + v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = F_deleteDependencyRecordsFor(m, int32(2617), v17, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v29 = F_new_object_addresses(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	F_deleteSharedDependencyRecordsFor(m, int32(2617), v17, int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(2615)
	F_add_exact_object_address(m, v10+int32(4), v29)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1247)
	F_add_exact_object_address(m, v10+int32(4), v29)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	if v51 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1247)
	F_add_exact_object_address(m, v10+int32(4), v29)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1247)
	F_add_exact_object_address(m, v10+int32(4), v29)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	if v71 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1255)
	F_add_exact_object_address(m, v10+int32(4), v29)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	if v81 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1255)
	F_add_exact_object_address(m, v10+int32(4), v29)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	if v91 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1255)
	F_add_exact_object_address(m, v10+int32(4), v29)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_record_object_address_dependencies(m, l0, v29, int32(110))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	F_free_object_addresses(m, v29)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	F_recordDependencyOnOwner(m, int32(2617), v107, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_recordDependencyOnCurrentExtension(m, l0, l3)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	m.G0 = v10 + int32(16)
	return
L42:
	;
	goto L41
}
