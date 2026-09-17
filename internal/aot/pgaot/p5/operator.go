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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v106
L2:
	;
	return int32(0)
L3:
	;
	if v12 == int32(0) {
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
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v18)
	v106 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(_a_F_OperatorIsVisibleExt_0), v9)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_OperatorIsVisibleExt_1), int32(2076), int32(_a_F_OperatorIsVisibleExt_2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	v38 = v34 + v35
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v39 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v12)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L35
	}
L15:
	;
	v42 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_OperatorIsVisibleExt[0]))
	if v44 == v42 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v89 = F_makeString(m, v38+int32(4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v103 = v42
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v77 = v42
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v77
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v39)
	if v68 == v39 {
		v77 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v77 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v89
	v96 = F_list_make1_impl(m, int32(1), v9+int32(8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v38)+84))
	v100 = F_OpernameGetOprid(m, v96, v98, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v103 = base.B2i32(v100 == l0)
	goto L14
L35:
	;
	v106 = v103
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
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
	v26 = F_SearchSysCacheCopy(m, int32(40), l1, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v26 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v32 = v30 + v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+92))
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v77
	F_CatalogTupleUpdate(m, v20, v26+int32(4), v26)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L28
	}
L12:
	;
	if v33 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if l0 == v33 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	v77 = v5
	goto L11
L16:
	;
	if v33 == int32(0) {
		v77 = l0
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v39 = F_get_opname(m, v33)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v49 = v32 + int32(4)
	if v39 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v49
	F_errmsg(m, int32(_a_F_OperatorUpd_4), v10+int32(-32))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v49
	F_errmsg(m, int32(_a_F_OperatorUpd_5), v10+int32(-16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	F_errfinish(m, int32(_a_F_OperatorUpd_1), int32(745), int32(_a_F_OperatorUpd_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
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
	v76 = m.ExcPending
	if v76 != 0 {
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
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L7
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v122
	F_errmsg(m, int32(_a_F_OperatorUpd_0), v10+int32(-48))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L53
	}
L31:
	;
	F_relation_close(m, v20, int32(3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L52
	}
L32:
	;
	v93 = F_SearchSysCacheCopy(m, int32(40), l2, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	if v93 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
	v99 = v97 + v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+96))
	if l3 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v100 == int32(0) {
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if l0 == v100 {
		goto L31
	} else {
		goto L41
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+96)) = int32(0)
	F_CatalogTupleUpdate(m, v20, v93+int32(4), v93)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	goto L31
L41:
	;
	if v100 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v112 = F_get_opname(m, v100)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+96)) = l0
	F_CatalogTupleUpdate(m, v20, v93+int32(4), v93)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L51
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v122 = v99 + int32(4)
	if v112 != 0 {
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v99)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v122
	F_errmsg(m, int32(_a_F_OperatorUpd_3), v12)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_OperatorUpd_1), int32(813), int32(_a_F_OperatorUpd_2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v159 = m.ExcPending
	if v159 != 0 {
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
