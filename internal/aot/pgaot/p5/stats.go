package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StatsShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemInit[0]))
	v11 = v9 + int32(_a_F_StatsShmemInit_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
	v16 = F_dsa_create_in_place_ext(m, v11, int32(_a_F_StatsShmemInit_1), int32(83), int32(0))
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
	F_dsa_pin(m, v16)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_dsa_set_size_limit(m, v16, int32(_a_F_StatsShmemInit_1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = F_dshash_create(m, v16, int32(_a_F_StatsShmemInit_2), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v28
	F_dsa_set_size_limit(m, v16, int32(-1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_pfree(m, v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_dsa_detach(m, v16)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(1)
	v46 = v9 + int32(_a_F_StatsShmemInit_3)
	v49 = int32(1)
	goto L10
L10:
	;
	if base.Ui32(v49-int32(1)) <= base.Ui32(int32(12)) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	return
L12:
	;
	v121 = v49 + int32(1)
	if v121 != int32(33) {
		v46 = v118
		v49 = v121
		goto L10
	} else {
		goto L32
	}
L13:
	;
	if v81 == int32(0) {
		v118 = v46
		goto L12
	} else {
		goto L20
	}
L14:
	;
	v81 = v49*int32(84) + int32(_a_F_StatsShmemInit_4)
	goto L13
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(int32(8)) < base.Ui32(v49-int32(24)) {
		v79 = int32(0)
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v81 = v79
	goto L13
L18:
	;
	v67 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemInit[1]))
	if v69 == v67 {
		v79 = v67
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v69+v49<<(uint(int32(2))%32)-int32(96))))
	v79 = v77
	goto L17
L20:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v84&int32(8) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if base.Ui32(v49) <= base.Ui32(int32(13)) {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)+v49<<(uint(int32(3))%32)))) = int64(0)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v92&int32(1) != 0 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v84&int32(1) == int32(0) {
		v118 = v46
		goto L12
	} else {
		goto L26
	}
L25:
	;
	v118 = v46
	goto L12
L26:
	;
	goto L21
L27:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v81)+64))
	m.T0[v115].(func(*base.Module, int32))(m, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v113 = v46
	v114 = v9 + v101
	goto L27
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(_a_F_StatsShmemInit_5)+v49<<(uint(int32(2))%32)))) = v46
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v113 = v46 + (v107+int32(7))&int32(-8)
	v114 = v46
	goto L27
L31:
	;
	v118 = v113
	goto L12
L32:
	;
	goto L11
}
func F_get_stats_slot_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8))))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l6)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if l1 != v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_fmgr_info(m, l1, l2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v24 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	if v17&int32(1) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v60 = int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v60 < v61 {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	v55 = v50
	v56 = v53
	v58 = v52
	v59 = int32(1)
	goto L8
L10:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v33)
	v50 = v28
	v52 = v33
	v53 = v28
	goto L9
L11:
	;
	goto L12
L12:
	;
	v36 = F_FunctionCall2Coll(m, l2, l3, v28, v19)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
	v41 = base.B2i32(v36 != int64(0))
	if v36 != int64(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = v39
	goto L16
L15:
	;
	v42 = v19
	goto L16
L16:
	;
	v44 = F_FunctionCall2Coll(m, l2, l3, v18, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if v44 == int64(0) {
		v55 = v42
		v56 = v18
		v58 = v41
		v59 = int32(0)
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
	v50 = v42
	v52 = v41
	v53 = v49
	goto L9
L19:
	;
	v65 = v60
	v72 = v59
	v73 = v55
	v74 = v56
	v77 = v58
	goto L22
L20:
	;
	v116 = v59
	v117 = v55
	v118 = v56
	v121 = v58
	goto L21
L21:
	;
	if v121&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v81 = v65 << (uint(int32(3)) % 32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v81+v82)))
	v85 = F_FunctionCall2Coll(m, l2, l3, v84, v73)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	v116 = v100
	v117 = v102
	v118 = v101
	v121 = v103
	goto L21
L24:
	;
	v88 = base.B2i32(v85 != int64(0))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v89+v81)))
	v92 = F_FunctionCall2Coll(m, l2, l3, v74, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v92 != int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v96+v81)))
	v100 = int32(1)
	v101 = v98
	goto L28
L27:
	;
	v100 = v72
	v101 = v74
	goto L28
L28:
	;
	if v85 != int64(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v102 = v91
	goto L31
L30:
	;
	v102 = v73
	goto L31
L31:
	;
	v103 = v77 | v88
	v105 = v65 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v105 < v106 {
		v65 = v105
		v72 = v100
		v73 = v102
		v74 = v101
		v77 = v103
		goto L22
	} else {
		goto L32
	}
L32:
	;
	goto L23
L33:
	;
	v131 = F_datumCopy(m, v118, l5, l4)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L40
	}
L34:
	;
	v126 = F_datumCopy(m, v117, l5, l4)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v116 == int32(0) {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l6))) = v126
	if v116 != 0 {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L6
L39:
	;
	goto L33
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l7))) = v131
	goto L6
}
func F_transformStatsStmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v4 = int32(0)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v8 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L25
	}
L2:
	;
	v12 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	return l1
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
	v18 = F_relation_open(m, l0, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v20 = int32(1)
	v21 = int32(0)
	v24 = F_addRangeTableEntryForRelation(m, v12, v18, v20, v21, v21, v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v27 = int32(1)
	F_addNSItemToQuery(m, v12, v24, int32(0), v27, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v31 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v68 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v34 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v42 = v4
	goto L13
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v42<<(uint(int32(2))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L10
L15:
	;
	v51 = F_transformExpr(m, v12, v49, int32(34))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v58 = v42 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v58 < v59 {
		v42 = v58
		goto L13
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v51
	F_assign_expr_collations(m, v12, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L14
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 != int32(1) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_free_parsestate(m, v12)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_relation_close(m, v18, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v79)
	goto L4
L25:
	;
	F_errcode(m, int32(_a_F_transformStatsStmt_0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_transformStatsStmt_1), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_transformStatsStmt_2), int32(3226), int32(_a_F_transformStatsStmt_3))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
