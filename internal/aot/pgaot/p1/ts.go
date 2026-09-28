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
	v8 = Fn14223(m, l0, l1, int32(77), int32(_a_F_TSParserIsVisibleExt_0), int32(2873), int32(_a_F_TSParserIsVisibleExt_1), int32(78))
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
	v5 = F_TS_execute_recurse(m, l0, l1, int32(2), int32(1715))
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
	v7 = Fn14294(m, l0, l1, int32(_a_F_get_ts_parser_oid_0), int32(2834), int32(_a_F_get_ts_parser_oid_1), int32(77))
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
func F_ts_ckpt_progress_comparator(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v10 = *(*float64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+8))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))+8))
	if base.F64_ne(v10, v12) != 0 {
		v14 = int32(-1)
	} else {
		v14 = int32(0)
	}
	if base.F64_lt(v10, v12) != 0 {
		v16 = int32(1)
	} else {
		v16 = v14
	}
	return v16
}
func F_ts_headline_byid_opt(m *base.Module, l0 int32) int64 {
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
	var v20 int64
	_ = v20
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
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
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
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v21 < int32(4) {
		v30 = v2
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v24 == int32(0) {
		v30 = v2
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
	v30 = v27
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
	v37 = base.I32_wrap_i64(v20)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+36)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v12)+28)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v12)+20)) = v40
	v46 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v46
	v50 = F_palloc_mul(m, int32(16), v46)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
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
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L49
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v50
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v58 = int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v62 = v60 & v58
	if v62 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v63 = v58
	goto L15
L14:
	;
	v63 = int32(4)
	goto L15
L15:
	;
	if v60 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_hlparsetext(m, v55, v12+int32(12), v37, v16+v63, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L27
	}
L17:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v70 == int32(18) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v81 = int32(1)
	if v62 != 0 {
		v91 = int32(base.Ui32(v60)>>(uint(v81)%32)) - v81
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v73 = int32(16)
	goto L22
L21:
	;
	v73 = int32(0)
	goto L22
L22:
	;
	if base.Ui32((v70-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v80 = int32(4)
	goto L25
L24:
	;
	v80 = v73
	goto L25
L25:
	;
	v91 = v80
	goto L16
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	if v30 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v101 = F_deserialize_deflist(m, base.I64_extend_i32_u(v30))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v105 = int64(0)
	goto L30
L30:
	;
	v106 = F_FunctionCall3Coll(m, v34+int32(112), int32(0), base.I64_extend_i32_u(v12+int32(12)), v105, v20&int64(4294967295))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v105 = base.I64_extend_i32_u(v101)
	goto L30
L32:
	;
	v110 = F_generateHeadline(m, v12+int32(12))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v112 != v16 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_pfree(m, v16)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v116 != v37 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	F_pfree(m, v37)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v30 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_pfree(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L46
	}
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v30 == v122 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_pfree(m, v30)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_pfree(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	F_pfree(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	m.G0 = v12 + int32(48)
	return base.I64_extend_i32_u(v110)
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_ts_headline_byid_opt_0), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_ts_headline_byid_opt_1), int32(306), int32(_a_F_ts_headline_byid_opt_2))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
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
func F_ts_headline_json_opt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v12 = F_DirectFunctionCall4Coll(m, int32(1289), int32(0), base.I64_extend_i32_u(v4), v9, v10, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			return v12
		}
	}
}
func F_ts_match_qv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(1736), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_ts_rankcd_wtt(m *base.Module, l0 int32) int64 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			F_getWeights(m, v12, v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v23 = F_calc_rank_cd(m, v9, v17, v19, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v25 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int64(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v29 != v17 {
								F_pfree(m, v17)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int64(0)
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v33 != v19 {
										F_pfree(m, v19)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int64(0)
										} else {
											m.G0 = v9 + int32(16)
											return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
										}
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
									}
								}
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v33 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int64(0)
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
									}
								} else {
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v29 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v33 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int64(0)
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
									}
								} else {
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
								}
							}
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							if v33 != v19 {
								F_pfree(m, v19)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int64(0)
								} else {
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
								}
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
							}
						}
					}
				}
			}
		}
	}
}
func F_ts_stat1(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_SPI_connect_ext(m, int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
					v22 = F_ts_stat_sql(m, v20, v11, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v24 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int64(0)
							} else {
								F_ts_setup_firstcall(m, l0, v15, v22)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int64(0)
								} else {
									v30 = F_SPI_finish(m)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return int64(0)
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
										v37 = F_ts_process_call(m, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return int64(0)
										} else {
											if v37 != int64(0) {
												v41 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
												*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41 + int64(1)
												v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(1)
												return v37
											} else {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return int64(0)
												} else {
													v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = int32(2)
													v54 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
													return v37
												}
											}
										}
									}
								}
							}
						} else {
							F_ts_setup_firstcall(m, l0, v15, v22)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int64(0)
							} else {
								v30 = F_SPI_finish(m)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int64(0)
								} else {
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
									v37 = F_ts_process_call(m, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int64(0)
									} else {
										if v37 != int64(0) {
											v41 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
											*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41 + int64(1)
											v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(1)
											return v37
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int64(0)
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = int32(2)
												v54 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
												return v37
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
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
		v37 = F_ts_process_call(m, v36)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int64(0)
		} else {
			if v37 != int64(0) {
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
				*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41 + int64(1)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(1)
				return v37
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = int32(2)
					v54 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
					return v37
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v141 int64
	_ = v141
	var v149 int32
	_ = v149
	var v154 int64
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v234 int32
	_ = v234
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
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
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
	v287 = m.ExcPending
	if v287 != 0 {
		goto L3
	} else {
		goto L81
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
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
	v258 = m.ExcPending
	if v258 != 0 {
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
	v141 = *(*int64)(unsafe.Add(mBase, _c_F_ts_stat_sql[1]))
	if v141 != int64(0) {
		goto L46
	} else {
		goto L47
	}
L19:
	;
	v55 = int32(1)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v57&v55 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v60 = v55
	goto L22
L21:
	;
	v60 = int32(4)
	goto L22
L22:
	;
	v61 = l2 + v60
	if v57 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v99 = v61
	goto L36
L24:
	;
	if v92 == int32(0) {
		goto L18
	} else {
		goto L35
	}
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v92 = int32(base.Ui32(v86)>>(uint(int32(2))%32)) - int32(4)
	goto L24
L26:
	;
	v97 = v61 + int32(4)
	goto L23
L27:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v57&int32(1) == int32(0) {
		goto L25
	} else {
		goto L34
	}
L30:
	;
	if v64 == int32(18) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v75 = int32(16)
	goto L33
L32:
	;
	v75 = int32(0)
	goto L33
L33:
	;
	v92 = v75
	goto L24
L34:
	;
	v80 = int32(1)
	v92 = int32(base.Ui32(v57)>>(uint(v80)%32)) - v80
	goto L24
L35:
	;
	v97 = v61 + v92
	goto L23
L36:
	;
	v110 = F_pg_mblen_range(m, v99, v97)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L18
L38:
	;
	v126 = v99 + v110
	if base.Ui32(v126) < base.Ui32(v97) {
		v99 = v126
		goto L36
	} else {
		goto L45
	}
L39:
	;
	if v110 != int32(1) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	switch v115 - int32(65) {
	case 0, 32:
		v121 = int32(8)
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
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v122 | v121
	goto L38
L42:
	;
	v121 = int32(1)
	goto L41
L43:
	;
	v121 = int32(2)
	goto L41
L44:
	;
	v121 = int32(4)
	goto L41
L45:
	;
	goto L37
L46:
	;
	v149 = v49
	v154 = int64(0)
	goto L49
L47:
	;
	v234 = v49
	goto L48
L48:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_ts_stat_sql[0]))
	F_SPI_freetuptable(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L3
	} else {
		goto L71
	}
L49:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_ts_stat_sql[0]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158+base.I32_wrap_i64(v154)<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v168 = F_SPI_getbinval(m, v163, v164, int32(1), v15+int32(31))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L51
	}
L50:
	;
	v234 = v211
	goto L48
L51:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	if v170 != 0 {
		v211 = v149
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v213 = v154 + int64(1)
	v215 = *(*int64)(unsafe.Add(mBase, _c_F_ts_stat_sql[1]))
	if base.Ui64(v213) < base.Ui64(v215) {
		v149 = v211
		v154 = v213
		goto L49
	} else {
		goto L67
	}
L53:
	;
	v171 = base.I32_wrap_i64(v168)
	v172 = F_pg_detoast_datum(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	if v149 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v177 = F_MemoryContextAllocZero(m, l0, int32(20))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	v181 = v149
	goto L57
L57:
	;
	if v172 == int32(0) {
		v211 = v181
		goto L52
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = int32(1)
	v181 = v177
	goto L57
L59:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v184 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v172 == v171 {
		v211 = v181
		goto L52
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v190 = int32(1)
	v196 = v190 << (uint(int32(0)-base.I32_clz(v184-v190)) % 32)
	v201 = int32(base.Ui32(v196-v184) >> (uint(v190) % 32))
	F_insertStatEntry(m, l0, v181, v172, int32(base.Ui32(v196)>>(uint(v190)%32))-v201)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L65
	}
L63:
	;
	F_pfree(m, v172)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v211 = v181
	goto L52
L65:
	;
	F_chooseNextStatEntry(m, l0, v181, v172, int32(0), v196, v201)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v211 = v181
	goto L52
L67:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_ts_stat_sql[0]))
	F_SPI_freetuptable(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	F_SPI_cursor_fetch(m, v25, int32(100))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v224 = int64(0)
	v226 = *(*int64)(unsafe.Add(mBase, _c_F_ts_stat_sql[1]))
	if v226 != v224 {
		v149 = v211
		v154 = v224
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
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	F_SPI_freeplan(m, v23)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	F_pfree(m, v17)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	m.G0 = v15 + int32(32)
	return v234
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	F_errmsg_internal(m, int32(_a_F_ts_stat_sql_0), v15)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_ts_stat_sql_1), int32(2589), int32(_a_F_ts_stat_sql_2))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	v277 = m.ExcPending
	if v277 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_ts_stat_sql_1), int32(2593), int32(_a_F_ts_stat_sql_2))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
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
	v290 = m.ExcPending
	if v290 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_ts_stat_sql_4), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_ts_stat_sql_1), int32(2603), int32(_a_F_ts_stat_sql_2))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
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
