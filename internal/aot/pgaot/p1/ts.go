package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TSDictionaryIsVisible(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_TSDictionaryIsVisibleExt(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_TSParserIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(78), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v148
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
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
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v148 = v3
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
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(41240), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(476407), int32(2804), int32(61485))
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
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L47
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	if v87 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v138 = v3
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
		v76 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v76
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
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v76 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v76 = v69
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
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 < v91 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v90
	goto L36
L34:
	;
	goto L35
L35:
	;
	v138 = v3
	goto L14
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v96<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if v108 != v110 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v108 == v40 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v120 = v96 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v120 < v121 {
		v96 = v120
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v138 = int32(1)
	goto L14
L42:
	;
	goto L43
L43:
	;
	v115 = int32(0)
	v117 = F_SearchSysCacheExists(m, int32(77), v39+int32(4), v108, v115, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v117 != 0 {
		v138 = v3
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	goto L37
L47:
	;
	v148 = v138
	goto L1
}
func F_TS_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_TS_execute_recurse(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 != int32(0))
	}
}
func F_TS_execute_ternary(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_TS_execute_recurse(m, l0, l1, int32(2), int32(1517))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_get_ts_parser_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_DeconstructQualifiedName(m, l0, v10+int32(12), v10+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v106
L4:
	;
	if l1 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L26
	}
L5:
	;
	v83 = int32(0)
	goto L4
L6:
	;
	v22 = F_LookupExplicitNamespace(m, v20, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = int32(0)
	goto L12
L11:
	;
	v24 = l1
	goto L12
L12:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v27 = int32(0)
	v29 = F_GetSysCacheOid(m, int32(77), v26, v22, v27, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v83 = v29
	goto L4
L15:
	;
	v33 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	if v35 == v33 {
		v83 = v33
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v39 <= v38 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v46 = v38
	goto L18
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v46<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if v54 != v56 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L5
L20:
	;
	v59 = int32(0)
	v61 = F_GetSysCacheOid(m, int32(77), v42, v54, v59, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v65 = v46 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v65 < v66 {
		v46 = v65
		goto L18
	} else {
		goto L25
	}
L23:
	;
	if v61 != 0 {
		v106 = v61
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L19
L26:
	;
	if v83 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v91 = F_NameListToString(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v91
	F_errmsg(m, int32(67677), v10)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(476407), int32(2765), int32(413390))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsCompareString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l4 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l3 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v11 = int32(0)
	if v11 < l3 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v14 = int32(-1)
	goto L9
L8:
	;
	v14 = v11
	goto L9
L9:
	;
	return v14
L10:
	;
	return base.B2i32(int32(0) < l1)
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(l1) < base.Ui32(l3) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v22 = l1
	goto L15
L14:
	;
	v22 = l3
	goto L15
L15:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v22) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if l4 != 0 {
		goto L35
	} else {
		goto L36
	}
L17:
	;
	v84 = int32(0)
	goto L16
L18:
	;
	v58 = v53
	v59 = v54
	v60 = v55
	goto L28
L19:
	;
	if (l0|l2)&int32(3) != 0 {
		v53 = l0
		v54 = l2
		v55 = v22
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v46 = l0
	v47 = l2
	v48 = v22
	goto L21
L21:
	;
	if v48 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v30 = l0
	v31 = l2
	v32 = v22
	goto L23
L23:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v35 != v36 {
		v53 = v30
		v54 = v31
		v55 = v32
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v46 = v41
	v47 = v39
	v48 = v43
	goto L21
L25:
	;
	v38 = int32(4)
	v39 = v31 + v38
	v41 = v30 + v38
	v43 = v32 - v38
	if base.Ui32(int32(3)) < base.Ui32(v43) {
		v30 = v41
		v31 = v39
		v32 = v43
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v53 = v46
	v54 = v47
	v55 = v48
	goto L18
L28:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v63 == v64 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v84 = v63 - v64
	goto L16
L30:
	;
	v66 = int32(1)
	v71 = v60 - v66
	if v71 != 0 {
		v58 = v58 + v66
		v59 = v59 + v66
		v60 = v71
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	return v94
L35:
	;
	if v84 != 0 {
		v94 = v84
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v84 != 0 {
		v94 = v84
		goto L34
	} else {
		goto L39
	}
L38:
	;
	return base.B2i32(l3 < l1)
L39:
	;
	if l1 == l3 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	goto L42
L42:
	;
	if l1 < l3 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v93 = int32(-1)
	goto L45
L44:
	;
	v93 = int32(1)
	goto L45
L45:
	;
	v94 = v93
	goto L34
}
func F_ts_ckpt_progress_comparator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_ne(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_lt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_ts_headline_byid_opt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v21 < int32(4) {
		v29 = v2
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = F_lookup_ts_config_cache(m, v14)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v24 == int32(0) {
		v29 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = F_pg_detoast_datum_packed(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v29 = v27
	goto L3
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v34 = F_lookup_ts_parser_cache(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+36)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v12)+28)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+20)) = v37
	v48 = F_palloc(m, int32(512))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L49
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v54 = int32(1)
	v55 = v16 + v54
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v60 = v58 & v54
	if v60 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v61 = v55
	goto L15
L14:
	;
	v61 = v16 + int32(4)
	goto L15
L15:
	;
	if v58 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_hlparsetext(m, v51, v12+int32(12), v20, v61, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L27
	}
L17:
	;
	v64 = int32(4)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v66&int32(254) == int32(2) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v79 = int32(1)
	if v60 != 0 {
		v89 = int32(base.Ui32(v58)>>(uint(v79)%32)) - v79
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v75 = v64
	goto L22
L21:
	;
	v75 = base.B2i32(v66 == int32(18)) << (uint(v64) % 32)
	goto L22
L22:
	;
	if v66 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = v64
	goto L25
L24:
	;
	v78 = v75
	goto L25
L25:
	;
	v89 = v78
	goto L16
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	v92 = int32(0)
	if v29 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = F_deserialize_deflist(m, v29)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v100 = v92
	goto L30
L30:
	;
	v101 = F_FunctionCall3Coll(m, v34+int32(112), v92, v12+int32(12), v100, v20)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v100 = v98
	goto L30
L32:
	;
	v105 = F_generateHeadline(m, v12+int32(12))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v107 != v16 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_pfree(m, v16)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v111 != v20 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	F_pfree(m, v20)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v29 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_pfree(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L46
	}
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v29 == v117 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_pfree(m, v29)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_pfree(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	F_pfree(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	m.G0 = v12 + int32(48)
	return v105
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(252091), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(472389), int32(306), int32(79511))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ts_headline_json_opt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v12 int32
	_ = v12
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_DirectFunctionCall4Coll(m, int32(1189), int32(0), v4, v8, v9, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_ts_match_qv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(1538), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_ts_rankcd_wtt(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 float32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = l0 + int32(28)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = F_pg_detoast_datum(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_getWeights(m, v13, v10)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v26 = F_calc_rank_cd(m, v10, v20, v22, int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v28 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
							if v32 != v20 {
								F_pfree(m, v20)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v36 != v22 {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I32_reinterpret_f32(v26)
										}
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v26)
									}
								}
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v36 != v22 {
									F_pfree(m, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v26)
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v26)
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						if v32 != v20 {
							F_pfree(m, v20)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v36 != v22 {
									F_pfree(m, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v26)
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v26)
								}
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v36 != v22 {
								F_pfree(m, v22)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v26)
								}
							} else {
								m.G0 = v10 + int32(16)
								return base.I32_reinterpret_f32(v26)
							}
						}
					}
				}
			}
		}
	}
}
func F_ts_stat1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_SPI_connect_ext(m, int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
					v21 = F_ts_stat_sql(m, v19, v10, int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v23 != v10 {
							F_pfree(m, v10)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								F_ts_setup_firstcall(m, l0, v14, v21)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return int32(0)
								} else {
									v29 = F_SPI_finish(m)
									mBase = m.M
									v30 = m.ExcPending
									if v30 != 0 {
										return int32(0)
									} else {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
										v36 = F_ts_process_call(m, v35)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											if v36 != 0 {
												v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
												*(*int64)(unsafe.Add(mBase, uint32(v35))) = v38 + int64(1)
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = int32(1)
												return v36
											} else {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return int32(0)
												} else {
													v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(2)
													v51 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
													return v36
												}
											}
										}
									}
								}
							}
						} else {
							F_ts_setup_firstcall(m, l0, v14, v21)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = F_SPI_finish(m)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
									v36 = F_ts_process_call(m, v35)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										if v36 != 0 {
											v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
											*(*int64)(unsafe.Add(mBase, uint32(v35))) = v38 + int64(1)
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = int32(1)
											return v36
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return int32(0)
											} else {
												v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(2)
												v51 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
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
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
		v36 = F_ts_process_call(m, v35)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			if v36 != 0 {
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
				*(*int64)(unsafe.Add(mBase, uint32(v35))) = v38 + int64(1)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = int32(1)
				return v36
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(2)
					v51 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
					return v36
				}
			}
		}
	}
}
func F_ts_stat_sql(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v136 int64
	_ = v136
	var v144 int32
	_ = v144
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v230 int32
	_ = v230
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
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = F_text_to_cstring(m, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L3
	} else {
		goto L83
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L3
	} else {
		goto L80
	}
L3:
	;
	return int32(0)
L4:
	;
	v21 = int32(0)
	v23 = F_SPI_prepare(m, v17, v21, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = F_SPI_cursor_open(m, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L3
	} else {
		goto L77
	}
L9:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_SPI_cursor_fetch(m, v25, int32(100))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[388]))
	if v33 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != int32(1) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v41 = F_SPI_gettypeid(m, v36, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v44 = F_IsBinaryCoercible(m, v41, int32(3614))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if v44 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v49 = F_MemoryContextAllocZero(m, l0, int32(20))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = int32(1)
	if l2 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v136 = *(*int64)(unsafe.Add(mBase, _consts[387]))
	if v136 != int64(0) {
		goto L45
	} else {
		goto L46
	}
L19:
	;
	v55 = int32(1)
	v56 = l2 + v55
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v61 = v59 & v55
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v62 = v56
	goto L22
L21:
	;
	v62 = l2 + int32(4)
	goto L22
L22:
	;
	if v59 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v91 = v62 + v90
	if base.Ui32(v91) <= base.Ui32(v62) {
		goto L18
	} else {
		goto L34
	}
L24:
	;
	v65 = int32(4)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v67&int32(254) == int32(2) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v80 = int32(1)
	if v61 != 0 {
		v90 = int32(base.Ui32(v59)>>(uint(v80)%32)) - v80
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v76 = v65
	goto L29
L28:
	;
	v76 = base.B2i32(v67 == int32(18)) << (uint(v65) % 32)
	goto L29
L29:
	;
	if v67 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v79 = v65
	goto L32
L31:
	;
	v79 = v76
	goto L32
L32:
	;
	v90 = v79
	goto L23
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
	goto L23
L34:
	;
	v94 = v62
	goto L35
L35:
	;
	v105 = F_pg_mblen_range(m, v94, v91)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L38
	}
L36:
	;
	goto L18
L37:
	;
	v121 = v94 + v105
	if base.Ui32(v121) < base.Ui32(v91) {
		v94 = v121
		goto L35
	} else {
		goto L44
	}
L38:
	;
	if v105 != int32(1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	switch v110 - int32(65) {
	case 0, 32:
		v116 = int32(8)
		goto L40
	case 1, 33:
		goto L43
	case 2, 34:
		goto L42
	case 3, 35:
		goto L41
	default:
		goto L37
	}
L40:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v117 | v116
	goto L37
L41:
	;
	v116 = int32(1)
	goto L40
L42:
	;
	v116 = int32(2)
	goto L40
L43:
	;
	v116 = int32(4)
	goto L40
L44:
	;
	goto L36
L45:
	;
	v144 = v49
	v150 = int64(0)
	goto L48
L46:
	;
	v230 = v49
	goto L47
L47:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[388]))
	F_SPI_freetuptable(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L3
	} else {
		goto L73
	}
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[388]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153+base.I32_wrap_i64(v150)<<(uint(int32(2))%32))))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v163 = F_SPI_getbinval(m, v158, v159, int32(1), v15+int32(31))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L50
	}
L49:
	;
	v230 = v207
	goto L47
L50:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v165 != 0 {
		v207 = v144
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v209 = v150 + int64(1)
	v211 = *(*int64)(unsafe.Add(mBase, _consts[387]))
	if base.Ui64(v209) < base.Ui64(v211) {
		v144 = v207
		v150 = v209
		goto L48
	} else {
		goto L69
	}
L52:
	;
	v166 = F_pg_detoast_datum(m, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	if v144 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v171 = F_MemoryContextAllocZero(m, l0, int32(20))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L57
	}
L55:
	;
	v175 = v144
	goto L56
L56:
	;
	if v166 == int32(0) {
		v207 = v175
		goto L51
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = int32(1)
	v175 = v171
	goto L56
L58:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v178 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if v166 == v163 {
		v207 = v175
		goto L51
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v184 = int32(1)
	v187 = v178 - v184
	if v187 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	F_pfree(m, v166)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	v207 = v175
	goto L51
L64:
	;
	v192 = v184 << (uint(int32(32)-base.I32_clz(v187)) % 32)
	goto L66
L65:
	;
	v192 = v184
	goto L66
L66:
	;
	v193 = int32(1)
	v197 = int32(base.Ui32(v192-v178) >> (uint(v193) % 32))
	F_insertStatEntry(m, l0, v175, v166, int32(base.Ui32(v192)>>(uint(v193)%32))-v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	F_chooseNextStatEntry(m, l0, v175, v166, int32(0), v192, v197)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v207 = v175
	goto L51
L69:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[388]))
	F_SPI_freetuptable(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	F_SPI_cursor_fetch(m, v25, int32(100))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v220 = int64(0)
	v222 = *(*int64)(unsafe.Add(mBase, _consts[387]))
	if v222 != v220 {
		v144 = v207
		v150 = v220
		goto L48
	} else {
		goto L72
	}
L72:
	;
	goto L49
L73:
	;
	F_SPI_cursor_close(m, v25)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	F_SPI_freeplan(m, v23)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	F_pfree(m, v17)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	m.G0 = v15 + int32(32)
	return v230
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	F_errmsg_internal(m, int32(433596), v15)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(472897), int32(2585), int32(286317))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v17
	F_errmsg_internal(m, int32(433567), v15+int32(16))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L3
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(472897), int32(2589), int32(286317))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(260144), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(472897), int32(2599), int32(286317))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
