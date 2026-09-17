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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn13838(m, l0, l1, int32(77), int32(_a_F_TSParserIsVisibleExt_0), int32(2804), int32(_a_F_TSParserIsVisibleExt_1), int32(78))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
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
	v5 = F_TS_execute_recurse(m, l0, l1, int32(2), int32(1502))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_get_ts_parser_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13903(m, l0, l1, int32(_a_F_get_ts_parser_oid_0), int32(2765), int32(_a_F_get_ts_parser_oid_1), int32(77))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
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
	var v39 int64
	_ = v39
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+36)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v12)+28)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v12)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(32)
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
	v136 = m.ExcPending
	if v136 != 0 {
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
	F_hlparsetext(m, v51, v12+int32(12), v20, v61, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L27
	}
L17:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v67 == int32(18) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v78 = int32(1)
	if v60 != 0 {
		v88 = int32(base.Ui32(v58)>>(uint(v78)%32)) - v78
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v70 = int32(16)
	goto L22
L21:
	;
	v70 = int32(0)
	goto L22
L22:
	;
	if base.Ui32((v67-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = int32(4)
	goto L25
L24:
	;
	v77 = v70
	goto L25
L25:
	;
	v88 = v77
	goto L16
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	if v29 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v96 = F_deserialize_deflist(m, v29)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v99 = int32(0)
	goto L30
L30:
	;
	v100 = F_FunctionCall3Coll(m, v34+int32(112), int32(0), v12+int32(12), v99, v20)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v99 = v96
	goto L30
L32:
	;
	v104 = F_generateHeadline(m, v12+int32(12))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v106 != v16 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_pfree(m, v16)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v110 != v20 {
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
	v113 = m.ExcPending
	if v113 != 0 {
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
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_pfree(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L46
	}
L43:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v29 == v116 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_pfree(m, v29)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_pfree(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	F_pfree(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	m.G0 = v12 + int32(48)
	return v104
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_ts_headline_byid_opt_0), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_ts_headline_byid_opt_1), int32(306), int32(_a_F_ts_headline_byid_opt_2))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
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
		v11 = F_DirectFunctionCall4Coll(m, int32(1174), int32(0), v4, v8, v9, v10)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1523), int32(0), v4, v5)
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 float32
	_ = v23
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_getWeights(m, v12, v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = F_calc_rank_cd(m, v9, v17, v19, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v17 {
								F_pfree(m, v17)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v33 != v19 {
										F_pfree(m, v19)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return base.I32_reinterpret_f32(v23)
										}
									} else {
										m.G0 = v9 + int32(16)
										return base.I32_reinterpret_f32(v23)
									}
								}
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v33 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return base.I32_reinterpret_f32(v23)
									}
								} else {
									m.G0 = v9 + int32(16)
									return base.I32_reinterpret_f32(v23)
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v33 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return base.I32_reinterpret_f32(v23)
									}
								} else {
									m.G0 = v9 + int32(16)
									return base.I32_reinterpret_f32(v23)
								}
							}
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v33 != v19 {
								F_pfree(m, v19)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return base.I32_reinterpret_f32(v23)
								}
							} else {
								m.G0 = v9 + int32(16)
								return base.I32_reinterpret_f32(v23)
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v140 int64
	_ = v140
	var v148 int32
	_ = v148
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v232 int32
	_ = v232
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
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
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
	v285 = m.ExcPending
	if v285 != 0 {
		goto L3
	} else {
		goto L81
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L78
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
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L75
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
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ts_stat_sql[0]))
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
	v140 = *(*int64)(unsafe.Add(mBase, _c_F_ts_stat_sql[1]))
	if v140 != int64(0) {
		goto L46
	} else {
		goto L47
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
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v98 = v62
	goto L36
L24:
	;
	if v91 == int32(0) {
		goto L18
	} else {
		goto L35
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
	goto L24
L26:
	;
	v96 = v62 + int32(4)
	goto L23
L27:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v61 == int32(0) {
		goto L25
	} else {
		goto L34
	}
L30:
	;
	if v65 == int32(18) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v76 = int32(16)
	goto L33
L32:
	;
	v76 = int32(0)
	goto L33
L33:
	;
	v91 = v76
	goto L24
L34:
	;
	v79 = int32(1)
	v91 = int32(base.Ui32(v59)>>(uint(v79)%32)) - v79
	goto L24
L35:
	;
	v96 = v62 + v91
	goto L23
L36:
	;
	v109 = F_pg_mblen_range(m, v98, v96)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L18
L38:
	;
	v125 = v98 + v109
	if base.Ui32(v125) < base.Ui32(v96) {
		v98 = v125
		goto L36
	} else {
		goto L45
	}
L39:
	;
	if v109 != int32(1) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	switch v114 - int32(65) {
	case 0, 32:
		v120 = int32(8)
		goto L41
	case 1, 33:
		goto L44
	case 2, 34:
		goto L43
	case 3, 35:
		goto L42
	default:
		goto L38
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v121 | v120
	goto L38
L42:
	;
	v120 = int32(1)
	goto L41
L43:
	;
	v120 = int32(2)
	goto L41
L44:
	;
	v120 = int32(4)
	goto L41
L45:
	;
	goto L37
L46:
	;
	v148 = v49
	v154 = int64(0)
	goto L49
L47:
	;
	v232 = v49
	goto L48
L48:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_ts_stat_sql[0]))
	F_SPI_freetuptable(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L3
	} else {
		goto L71
	}
L49:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_ts_stat_sql[0]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157+base.I32_wrap_i64(v154)<<(uint(int32(2))%32))))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v167 = F_SPI_getbinval(m, v162, v163, int32(1), v15+int32(31))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L51
	}
L50:
	;
	v232 = v209
	goto L48
L51:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v169 != 0 {
		v209 = v148
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v211 = v154 + int64(1)
	v213 = *(*int64)(unsafe.Add(mBase, _c_F_ts_stat_sql[1]))
	if base.Ui64(v211) < base.Ui64(v213) {
		v148 = v209
		v154 = v211
		goto L49
	} else {
		goto L67
	}
L53:
	;
	v170 = F_pg_detoast_datum(m, v167)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	if v148 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v175 = F_MemoryContextAllocZero(m, l0, int32(20))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	v179 = v148
	goto L57
L57:
	;
	if v170 == int32(0) {
		v209 = v179
		goto L52
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = int32(1)
	v179 = v175
	goto L57
L59:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v182 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v170 == v167 {
		v209 = v179
		goto L52
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v188 = int32(1)
	v194 = v188 << (uint(int32(0)-base.I32_clz(v182-v188)) % 32)
	v199 = int32(base.Ui32(v194-v182) >> (uint(v188) % 32))
	F_insertStatEntry(m, l0, v179, v170, int32(base.Ui32(v194)>>(uint(v188)%32))-v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L3
	} else {
		goto L65
	}
L63:
	;
	F_pfree(m, v170)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v209 = v179
	goto L52
L65:
	;
	F_chooseNextStatEntry(m, l0, v179, v170, int32(0), v194, v199)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v209 = v179
	goto L52
L67:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_ts_stat_sql[0]))
	F_SPI_freetuptable(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	F_SPI_cursor_fetch(m, v25, int32(100))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v222 = int64(0)
	v224 = *(*int64)(unsafe.Add(mBase, _c_F_ts_stat_sql[1]))
	if v224 != v222 {
		v148 = v209
		v154 = v222
		goto L49
	} else {
		goto L70
	}
L70:
	;
	goto L50
L71:
	;
	F_SPI_cursor_close(m, v25)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	F_SPI_freeplan(m, v23)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	F_pfree(m, v17)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	m.G0 = v15 + int32(32)
	return v232
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	F_errmsg_internal(m, int32(_a_F_ts_stat_sql_0), v15)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_ts_stat_sql_1), int32(2585), int32(_a_F_ts_stat_sql_2))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v17
	F_errmsg_internal(m, int32(_a_F_ts_stat_sql_3), v15+int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_ts_stat_sql_1), int32(2589), int32(_a_F_ts_stat_sql_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_ts_stat_sql_4), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_ts_stat_sql_1), int32(2599), int32(_a_F_ts_stat_sql_2))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
