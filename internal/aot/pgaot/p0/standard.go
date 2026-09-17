package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_adjust_standard_join_alias_expression(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v4 = l0
	goto L2
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	switch v7 - int32(6) {
	case 0:
		goto L11
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 24, 25, 26, 27, 28, 29, 30, 31:
		goto L1
	case 9:
		goto L8
	case 21:
		goto L7
	case 22:
		goto L6
	case 23:
		goto L5
	case 32:
		goto L4
	default:
		goto L10
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v33 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L4:
	;
	goto L3
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v4 = v32
	goto L2
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v4 = v31
	goto L2
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v4 = v30
	goto L2
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v4 = v29
	goto L2
L9:
	;
	v21 = v20 + v4
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v24 = F_bms_add_members(m, v22, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if v7 != int32(319) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v10 != v11 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v20 = int32(24)
	goto L9
L13:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v16 != v17 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v20 = int32(12)
	goto L9
L15:
	;
	return
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v24
	return
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v40 = int32(0)
	goto L19
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v40<<(uint(int32(2))%32))))
	F_adjust_standard_join_alias_expression(m, v47, l1)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L21
	}
L20:
	;
	goto L1
L21:
	;
	v51 = v40 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v51 < v52 {
		v40 = v51
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
func F_standard_ExecutorEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+164))
	if int32(0) < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+168)))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorEnd[0]))
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
	v36 = int32(_a_F_standard_ExecutorEnd_0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorEnd[1]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorEnd[1])) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_ExecEndNode(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L9
	}
L4:
	;
	v20 = F_pgstat_prep_pending_entry(m, int32(1), v15, int64(0), int32(0))
	mBase = m.M
	v21 = m.ExcPending
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
	goto L3
L7:
	;
	return
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+240)) = v23 + base.I64_extend_i32_u(v10)
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v22)+248))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+248)) = v26 + v13
	goto L6
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+144))
	if v44 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+104))
	F_ExecResetTupleTable(m, v77, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v47 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v51 = int32(0)
	goto L13
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v51<<(uint(int32(2))%32))))
	F_ExecEndNode(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L10
L15:
	;
	v66 = v51 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v66 < v67 {
		v51 = v66
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_ExecCloseResultRelations(m, v9)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v83 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v86 = int32(0)
	v88 = v83
	goto L22
L20:
	;
	goto L21
L21:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	F_UnregisterSnapshot(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L29
	}
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+v86<<(uint(int32(2))%32))))
	if v97 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	F_relation_close(m, v97, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L27
	}
L25:
	;
	v102 = v88
	goto L26
L26:
	;
	v104 = v86 + int32(1)
	if base.Ui32(v104) < base.Ui32(v102) {
		v86 = v104
		v88 = v102
		goto L22
	} else {
		goto L28
	}
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v102 = v101
	goto L26
L28:
	;
	goto L23
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	F_UnregisterSnapshot(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorEnd[1])) = v37
	F_FreeExecutorState(m, v9)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v124 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
	return
}
