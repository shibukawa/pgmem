package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EvalPlanQualEnd(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v11 = v9 << (uint(int32(2)) % 32)
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if v6 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	base.MemoryFill(m, v12, int32(0), v11)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ExecResetTupleTable(m, v15, int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	goto L3
L9:
	;
	v23 = int32(_a_F_EvalPlanQualEnd_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualEnd[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualEnd[0])) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_ExecEndNode(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	return
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+144))
	if v31 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v6)+104))
	F_ExecResetTupleTable(m, v58, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L20
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v34 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v39 = int32(0)
	goto L16
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v39<<(uint(int32(2))%32))))
	F_ExecEndNode(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	v50 = v39 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v50 < v51 {
		v39 = v50
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	F_ExecCloseResultRelations(m, v6)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualEnd[0])) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = int32(0)
	F_FreeExecutorState(m, v6)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v70 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(l0)+28)) = v70
	goto L11
}
func F_preprocess_qual_conditions(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v13 - int32(63) {
	case 0:
		goto L1
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L4
	}
L3:
	;
	v67 = l1 + v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v70 = F_preprocess_expression(m, l0, v68, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
	} else {
		goto L19
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L16
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_preprocess_qual_conditions(m, l0, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L14
	}
L6:
	;
	v16 = int32(8)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 == int32(0) {
		v65 = v16
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		v65 = v16
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v28 = v3
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v28<<(uint(int32(2))%32))))
	F_preprocess_qual_conditions(m, l0, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v65 = v16
	goto L3
L11:
	;
	return
L12:
	;
	v37 = v28 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v37 < v38 {
		v28 = v37
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_preprocess_qual_conditions(m, l0, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v65 = int32(28)
	goto L3
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v51
	F_errmsg_internal(m, int32(_a_F_preprocess_qual_conditions_0), v9)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_preprocess_qual_conditions_1), int32(1386), int32(_a_F_preprocess_qual_conditions_2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v70
	goto L1
}
