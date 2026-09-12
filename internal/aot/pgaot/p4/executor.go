package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecutorEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
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
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v10 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.T0[v10].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+164))
	if int32(0) < v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+168)))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v39 = int32(4515600)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_ExecEndNode(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L13
	}
L9:
	;
	v24 = F_pgstat_prep_pending_entry(m, int32(1), v19, int64(0), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v27 + base.I64_extend_i32_u(v14)
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v26)+248))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v30 + v17
	goto L11
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+144))
	if v47 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	F_ExecResetTupleTable(m, v80, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L21
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v50 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v56 = int32(0)
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v56<<(uint(int32(2))%32))))
	F_ExecEndNode(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L14
L19:
	;
	v69 = v56 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v69 < v70 {
		v56 = v69
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	F_ExecCloseResultRelations(m, v13)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v86 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v90 = v86
	v91 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_UnregisterSnapshot(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L33
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v91<<(uint(int32(2))%32))))
	if v100 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	F_sequence_close(m, v100, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	v105 = v90
	goto L30
L30:
	;
	v107 = v91 + int32(1)
	if base.Ui32(v107) < base.Ui32(v105) {
		v90 = v105
		v91 = v107
		goto L26
	} else {
		goto L32
	}
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v105 = v104
	goto L30
L32:
	;
	goto L27
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_UnregisterSnapshot(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v40
	F_FreeExecutorState(m, v13)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v127 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
	return
}
func F_ExecutorRewind(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = int32(4515600)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+100))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_ExecReScan(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[28])) = v4
		return
	}
}
