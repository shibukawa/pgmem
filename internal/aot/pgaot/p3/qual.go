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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v15 = F__emscripten_memset_bulkmem(m, v8, base.I32_extend8_s(int32(0)), v11<<(uint(int32(2))%32))
	mBase = m.M
	goto L4
L2:
	;
	goto L3
L3:
	;
	if v6 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ExecResetTupleTable(m, v16, int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	goto L3
L7:
	;
	v22 = int32(4489440)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_ExecEndNode(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	return
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+144))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v6)+104))
	F_ExecResetTupleTable(m, v57, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L18
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v38 = int32(0)
	goto L14
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v38<<(uint(int32(2))%32))))
	F_ExecEndNode(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	v49 = v38 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v49 < v50 {
		v38 = v49
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	F_ExecCloseResultRelations(m, v6)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = int32(0)
	F_FreeExecutorState(m, v6)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(l0)+28)) = v69
	goto L9
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
	var v25 int32
	_ = v25
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
	var v66 int32
	_ = v66
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
	v67 = l1 + v66
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
		v66 = v16
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		v66 = v16
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v25 = v3
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v25<<(uint(int32(2))%32))))
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
	v66 = v16
	goto L3
L11:
	;
	return
L12:
	;
	v37 = v25 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v37 < v38 {
		v25 = v37
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
	v66 = int32(28)
	goto L3
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v51
	F_errmsg_internal(m, int32(485403), v9)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(494727), int32(1386), int32(138977))
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
