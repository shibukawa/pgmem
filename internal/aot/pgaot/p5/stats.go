package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_stats_slot_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8))))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v17&int32(1) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v61 = int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v61 < v62 {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	v55 = v49
	v57 = v51
	v59 = v53
	v60 = int32(1)
	goto L8
L10:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v33)
	v49 = v28
	v51 = v33
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = v39
	goto L16
L15:
	;
	v40 = v19
	goto L16
L16:
	;
	v42 = base.B2i32(v36 != int32(0))
	v43 = F_FunctionCall2Coll(m, l2, l3, v18, v39)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if v43 == int32(0) {
		v55 = v40
		v57 = v42
		v59 = v18
		v60 = int32(0)
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = v40
	v51 = v42
	v53 = v48
	goto L9
L19:
	;
	v66 = v55
	v73 = v61
	v74 = v57
	v76 = v59
	v77 = v60
	goto L22
L20:
	;
	v108 = v55
	v116 = v57
	v118 = v59
	v119 = v60
	goto L21
L21:
	;
	if v116&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v82 = v73 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)))
	v86 = F_FunctionCall2Coll(m, l2, l3, v85, v66)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	v108 = v101
	v116 = v102
	v118 = v99
	v119 = v100
	goto L21
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90+v82)))
	v93 = F_FunctionCall2Coll(m, l2, l3, v76, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v93 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96+v82)))
	v99 = v98
	v100 = int32(1)
	goto L28
L27:
	;
	v99 = v76
	v100 = v77
	goto L28
L28:
	;
	if v86 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v101 = v92
	goto L31
L30:
	;
	v101 = v66
	goto L31
L31:
	;
	v102 = v74 | base.B2i32(v86 != int32(0))
	v104 = v73 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v104 < v105 {
		v66 = v101
		v73 = v104
		v74 = v102
		v76 = v99
		v77 = v100
		goto L22
	} else {
		goto L32
	}
L32:
	;
	goto L23
L33:
	;
	v130 = F_datumCopy(m, v118, l5, l4)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L40
	}
L34:
	;
	v125 = F_datumCopy(m, v108, l5, l4)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v119 == int32(0) {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v125
	if v119 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v130
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
	var v41 int32
	_ = v41
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
	v41 = v4
	goto L13
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v41<<(uint(int32(2))%32))))
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
	v58 = v41 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v58 < v59 {
		v41 = v58
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
	F_sequence_close(m, v18, int32(0))
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
	F_errcode(m, int32(393348))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(435440), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(468988), int32(3195), int32(90322))
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
