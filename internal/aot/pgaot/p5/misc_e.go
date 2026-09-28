package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EOH_init_header(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = int32(513)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v4)
	v6 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v6)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l0)+14)) = l0
	return
}
func F_ExecASDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v5 == int32(0) {
		return
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+22)))
		if v8 != int32(1) {
			return
		} else {
			v11 = int32(0)
			F_AfterTriggerSaveEvent(m, l0, l1, v11, v11, int32(1), v11, v11, v11, v11, v11, l2, v11)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ExecBSTruncateTriggers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v19 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L20
	}
L2:
	;
	m.G0 = v9 + int32(48)
	return
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+23)))
	if v22 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(47244640704)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v29 <= int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v37 = v3
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v41 = v38 + v37*int32(60)
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	if v42&int32(99) != int32(34) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L2
L8:
	;
	v69 = v37 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v69 < v70 {
		v37 = v69
		goto L6
	} else {
		goto L19
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v48 = int32(0)
	v51 = F_TriggerEnabled(m, l0, l1, v41, v47, v48, v48, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v51 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v41
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v60 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v63 = v60
	goto L15
L14:
	;
	v61 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v65 = F_ExecCallTriggerFunc(m, v9+int32(4), v37, v58, v59, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v63 = v61
	goto L15
L17:
	;
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L8
L19:
	;
	goto L7
L20:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_ExecBSTruncateTriggers_0), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ExecBSTruncateTriggers_1), int32(3338), int32(_a_F_ExecBSTruncateTriggers_2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecCloseTrigTargetRelations(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v6 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = int32(0)
	goto L4
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v10<<(uint(int32(2))%32))))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_relation_close(m, v17, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v22 = v10 + int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v22 < v23 {
		v10 = v22
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_ExecFindJunkAttributeInTlist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = int32(0)
	if v14 < v11 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v78 = int32(0)
	goto L6
L6:
	;
	return base.I32_extend16_s(v78)
L7:
	;
	v17 = v11
	goto L9
L8:
	;
	v17 = v14
	goto L9
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = int32(0)
	goto L11
L10:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+8)))
	v78 = v70
	goto L6
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+26)))
	if v30 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	v66 = v20 + int32(1)
	if v66 != v17 {
		v20 = v66
		goto L11
	} else {
		goto L24
	}
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v33 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v38 == int32(0))|base.B2i32(v38 != v41) != 0 {
		v59 = v38
		v60 = v41
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v59-v60 == int32(0) {
		goto L10
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v44 = v33
	v45 = l1
	goto L19
L19:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v49 == int32(0) {
		v59 = v49
		v60 = v48
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v59 = v49
	v60 = v48
	goto L17
L21:
	;
	v52 = int32(1)
	if v49 == v48 {
		v44 = v44 + v52
		v45 = v45 + v52
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L13
L24:
	;
	goto L12
}
func F_ExecFindMatchingSubPlans(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v4
	v15 = int32(_a_F_ExecFindMatchingSubPlans_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindMatchingSubPlans[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindMatchingSubPlans[0])) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 < v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = v4
	goto L4
L2:
	;
	v66 = int32(0)
	goto L3
L3:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v68 = F_bms_add_members(m, v66, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L13
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(24)+v30<<(uint(int32(2))%32))))
	F_find_matching_subplans_recurse(m, v36, v36+int32(4), l1, v11+int32(12), l2)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v66 = v56
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = v30 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v53 < v54 {
		v30 = v53
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+116))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	F_MemoryContextReset(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L5
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindMatchingSubPlans[0])) = v16
	v72 = F_bms_copy(m, v68)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v75 = F_bms_copy(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_MemoryContextReset(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v75
	goto L17
L19:
	;
	m.G0 = v11 + int32(16)
	return v72
}
func F_ExecGetReturningSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v5 == int32(0) {
		v8 = int32(_a_F_ExecGetReturningSlot_0)
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGetReturningSlot[0]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecGetReturningSlot[0])) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
		v15 = F_table_slot_callbacks(m, v10)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_ExecInitExtraTupleSlot(m, l0, v14, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v19
				*(*int32)(unsafe.Add(mBase, _c_F_ExecGetReturningSlot[0])) = v9
				v24 = v19
				return v24
			}
		}
	} else {
		v24 = v5
		return v24
	}
}
func F_ExecInitExtraTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l1 == int32(0) {
		v26 = int32(2)
		v27 = v7
	} else {
		v12 = int32(7)
		v14 = int32(-8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v26 = int32(18)
		v27 = (v7+v12)&v14 + v16<<(uint(int32(3))%32) + (v16+v12)&v14
	}
	v28 = F_palloc0(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = l1
		*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)) = uint16(v26)
		*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(449)
		*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l2
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitExtraTupleSlot[0]))
		v39 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v28)+6)) = uint16(v39)
		*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v38
		if l1 != 0 {
			v46 = v28 + (v7+int32(7))&int32(-8)
			*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v46
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v46 + v48<<(uint(int32(3))%32)
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if int32(0) <= v53 {
				F_IncrTupleDescRefCount(m, l1)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v59 = v58
					*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = int32(0)
					v62 = v59
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					m.T0[v64].(func(*base.Module, int32))(m, v28)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v68 = F_lappend(m, v67, v28)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v68
							return v28
						}
					}
				}
			} else {
				v59 = l2
				*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = int32(0)
				v62 = v59
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
				m.T0[v64].(func(*base.Module, int32))(m, v28)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v68 = F_lappend(m, v67, v28)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v68
						return v28
					}
				}
			}
		} else {
			v62 = l2
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
			m.T0[v64].(func(*base.Module, int32))(m, v28)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v68 = F_lappend(m, v67, v28)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v68
					return v28
				}
			}
		}
	}
}
func F_ExecLockRows(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[0]))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
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
	v26 = l0 + int32(108)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L6
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v17 + int32(80)
	return v379
L8:
	;
	F_ExecReScan(m, v27)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v47 = m.T0[v46].(func(*base.Module, int32) int32)(m, v27)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	goto L7
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v57 == int32(0) {
		v379 = v47
		goto L12
	} else {
		goto L20
	}
L14:
	;
	if v47 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v49&int32(2) == int32(0) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_EvalPlanQualEnd(m, v26)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v379 = int32(0)
	goto L12
L20:
	;
	v60 = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v62 <= v60 {
		v379 = v47
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v73 = v60
	v77 = v60
	goto L22
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77<<(uint(int32(2))%32))))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v87 = F_EvalPlanQualSlot(m, v26, v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	if v345&int32(1) == int32(0) {
		v379 = v47
		goto L12
	} else {
		goto L95
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	m.T0[v90].(func(*base.Module, int32))(m, v87)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	if v93 == v94 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v348 = v77 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v348 < v349 {
		v73 = v345
		v77 = v348
		goto L22
	} else {
		goto L94
	}
L27:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L91
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L88
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L84
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L81
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L78
	}
L32:
	;
	v124 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+32)) = uint8(v124)
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+4)))
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+6)))
	if v127 < v126 {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+6)))
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+6)))
	if v97 < v96 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	m.T0[v100].(func(*base.Module, int32, int32))(m, v47, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v103 = int32(1)
	v104 = v96 - v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+v105))))
	if v107 == v103 {
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v104<<(uint(int32(3))%32))))
	if v110 == v115 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v117 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+38)) = uint16(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+34)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+32)) = uint8(v117)
	v345 = v73
	goto L26
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	m.T0[v130].(func(*base.Module, int32, int32))(m, v47, v126)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v133 = int32(1)
	v134 = v126 - v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v135))))
	if v137 == v133 {
		goto L30
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v140+v134<<(uint(int32(3))%32))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+48))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+119)))
	if v147 == int32(102) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+51)) = uint8(v150)
	v153 = F_GetFdwRoutineForRelation(m, v145, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v166 = base.I32_wrap_i64(v144)
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+76)) = uint16(v167)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if base.Ui32(int32(4)) <= base.Ui32(v171) {
		goto L28
	} else {
		goto L51
	}
L48:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+108))
	if v155 == int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	m.T0[v155].(func(*base.Module, int32, int32, int64, int32, int32))(m, v28, v84, v144, v87, v17+int32(51))
	mBase = m.M
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v161&int32(2) != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+51)))
	v345 = v164 | v73
	goto L26
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v179 = int32(3)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	v182 = int32(1)
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[1]))
	if v182 < v185 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L75
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L72
	}
L54:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[1]))
	if v224 < int32(2) {
		goto L6
	} else {
		goto L67
	}
L55:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[1]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L63
	}
L56:
	;
	v188 = v182
	goto L58
L57:
	;
	v188 = v179
	goto L58
L58:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v174)+188))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+104))
	v193 = m.T0[v192].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v174, v17+int32(72), v177, v87, v178, v179-v171, v181, v188, v17+int32(52))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	if v193 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	switch v193 - int32(1) {
	case 0:
		goto L53
	case 1, 5:
		goto L6
	case 2:
		goto L55
	case 3:
		goto L54
	default:
		goto L52
	}
L61:
	;
	goto L62
L62:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+68)))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+38)) = uint16(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+34)) = v200
	v345 = v197 | v73
	goto L26
L63:
	;
	if int32(2) <= v204 {
		goto L27
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(3)
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_0), v17+int32(32))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(230), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_ExecLockRows_3), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(237), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_4), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(242), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v193
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_5), v17+int32(16))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(247), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_6), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(102), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_7), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(122), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v305 + int32(4)
	F_errmsg(m, int32(_a_F_ExecLockRows_8), v17)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(136), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errmsg_internal(m, int32(_a_F_ExecLockRows_9), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(176), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errmsg(m, int32(_a_F_ExecLockRows_10), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_ExecLockRows_1), int32(228), int32(_a_F_ExecLockRows_2))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	goto L23
L95:
	;
	F_EvalPlanQualBegin(m, v26)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v47
	v358 = int32(_a_F_ExecLockRows_11)
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[2]))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[2])) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+52))
	if v365 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_ExecReScan(m, v364)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v369 = m.T0[v368].(func(*base.Module, int32) int32)(m, v364)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecLockRows[2])) = v359
	if v369 == int32(0) {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+4)))
	if v375&int32(2) != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	v379 = v369
	goto L12
}
func F_ExecModifyTable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v530 int32
	_ = v530
	var v560 int32
	_ = v560
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int64
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int64
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int64
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v828 int32
	_ = v828
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int64
	_ = v889
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v932 int64
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1011 float64
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1024 int32
	_ = v1024
	var v1025 float64
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 float64
	_ = v1111
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1133 int64
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1262 int32
	_ = v1262
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1416 int64
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1433 int32
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1439 float64
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1574 int32
	_ = v1574
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1597 int32
	_ = v1597
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1608 int32
	_ = v1608
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1680 int32
	_ = v1680
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1867 int32
	_ = v1867
	var v1900 int32
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1917 int32
	_ = v1917
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1943 int32
	_ = v1943
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1971 int32
	_ = v1971
	var v1978 int32
	_ = v1978
	var v1982 int32
	_ = v1982
	var v1987 int32
	_ = v1987
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v1998 int32
	_ = v1998
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2019 int32
	_ = v2019
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2052 int32
	_ = v2052
	var v2055 int32
	_ = v2055
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2071 int32
	_ = v2071
	var v2102 int32
	_ = v2102
	var v2106 int32
	_ = v2106
	var v2111 int32
	_ = v2111
	v28 = m.G0
	v30 = v28 - int32(176)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[0]))
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+156))
	if v40 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L4
	} else {
		goto L593
	}
L7:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L4
	} else {
		goto L589
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L4
	} else {
		goto L584
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L4
	} else {
		goto L581
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L4
	} else {
		goto L577
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L4
	} else {
		goto L573
	}
L12:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L4
	} else {
		goto L569
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L4
	} else {
		goto L564
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L4
	} else {
		goto L561
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L4
	} else {
		goto L558
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L4
	} else {
		goto L555
	}
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)))
	if v43 != 0 {
		v1867 = int32(0)
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L4
	} else {
		goto L552
	}
L20:
	;
	m.G0 = v30 + int32(176)
	return v1867
L21:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)))
	if v44 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	switch v48 - int32(2) {
	case 0:
		goto L26
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	default:
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+116)) = v33
	v111 = int32(124)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = l0 + v111
	*(*int32)(unsafe.Add(mBase, uint32(v30)+108)) = l0
	v119 = v30 + v111
	v127 = v108 + v109*int32(216)
	goto L49
L25:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)) = uint8(v103)
	goto L24
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v99, v47)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L48
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L45
	}
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v64&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSDeleteTriggers(m, v61, v47)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L34
	}
L30:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSInsertTriggers(m, v52, v47)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+128))
	if v55 != int32(2) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v58, v47)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	goto L25
L35:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSInsertTriggers(m, v67, v47)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v71 = v64
	goto L37
L37:
	;
	if v71&int32(2) != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v71 = v70
	goto L37
L39:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSUpdateTriggers(m, v74, v47)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	v78 = v71
	goto L41
L41:
	;
	if v78&int32(4) == int32(0) {
		goto L25
	} else {
		goto L43
	}
L42:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v78 = v77
	goto L41
L43:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecBSDeleteTriggers(m, v83, v47)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L25
L45:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_0), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_2), int32(_a_F_ExecModifyTable_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	goto L25
L49:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)+152))
	if v151 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v33)+188))
	if v1797 != 0 {
		goto L522
	} else {
		goto L523
	}
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	F_MemoryContextReset(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v155 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	F_MemoryContextReset(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v159 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+144)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+120)) = v159
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v167 = F_ExecMergeNotMatched(m, v30+int32(108), v165, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v107)+52))
	if v173 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v169
	if v167 == v169 {
		goto L49
	} else {
		goto L63
	}
L63:
	;
	v1867 = v167
	goto L20
L64:
	;
	F_ExecReScan(m, v107)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v177 = m.T0[v176].(func(*base.Module, int32) int32)(m, v107)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v179 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+144)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v30)+120)) = v177
	if v177 == v179 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	goto L50
L70:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+4)))
	if v184&int32(2) != 0 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v187 == int32(0) {
		v241 = v127
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+92)))
	if v243 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L73:
	;
	v190 = base.I32_extend16_s(v187)
	v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(v177)+6)))
	if v191 < v190 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	m.T0[v194].(func(*base.Module, int32, int32))(m, v177, v190)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v197 = int32(1)
	v198 = v190 - v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v199))))
	if v201 == v197 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	if v32 == int32(5) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229+v198<<(uint(int32(3))%32))))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v233 == v234 {
		v241 = v127
		goto L72
	} else {
		goto L89
	}
L81:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v206
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v212 = F_ExecMergeNotMatched(m, v30+int32(108), v210, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L86
	}
L84:
	;
	if v212 == int32(0) {
		goto L49
	} else {
		goto L85
	}
L85:
	;
	v1867 = v212
	goto L20
L86:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_4), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_5), int32(_a_F_ExecModifyTable_6))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v238 = F_ExecLookupResultRelByOid(m, l0, v233, int32(0), int32(1))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v241 = v238
	goto L72
L91:
	;
	v250 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	v253 = F_ExecProcessReturning(m, v30+int32(108), v241, base.B2i32(v32 == int32(4)), v250, v250, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v255
	v257 = int32(0)
	if base.Ui32(int32(5)) < base.Ui32(v32) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v1867 = v253
	goto L20
L95:
	;
	switch v32 - int32(2) {
	case 0:
		goto L144
	case 1:
		goto L145
	case 2:
		goto L141
	case 3:
		goto L143
	default:
		goto L142
	}
L96:
	;
	v400 = int32(0)
	v402 = v257
	goto L95
L97:
	;
	goto L98
L98:
	;
	if int32(1)<<(uint(v32)%32)&int32(52) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v400 = int32(0)
	v402 = v257
	goto L95
L100:
	;
	goto L101
L101:
	;
	v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241)+24)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+48))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+119)))
	v271 = v269 - int32(109)
	v278 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v271))|base.B2i32(int32(1)<<(uint(v271)%32)&int32(41) == v278) == v278 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+6)))
	if v283 < v266 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v332 = int32(0)
	if v266 == v332 {
		v400 = v332
		v402 = v257
		goto L95
	} else {
		goto L120
	}
L105:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
	m.T0[v286].(func(*base.Module, int32, int32))(m, v255, v266)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v289 = int32(1)
	v290 = v266 - v289
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v291))))
	if v293 == v289 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	if v32 == int32(5) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321+v290<<(uint(int32(3))%32))))
	v326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v325)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+104)) = uint16(v326)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+100)) = v328
	v400 = v30 + int32(100)
	v402 = v257
	goto L95
L112:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v298
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v304 = F_ExecMergeNotMatched(m, v30+int32(108), v302, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L117
	}
L115:
	;
	if v304 == int32(0) {
		v127 = v241
		goto L49
	} else {
		goto L116
	}
L116:
	;
	v1867 = v304
	goto L20
L117:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_7), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_8), int32(_a_F_ExecModifyTable_6))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v335 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+6)))
	if v335 < v266 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	m.T0[v338].(func(*base.Module, int32, int32))(m, v255, v266)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v341 = int32(1)
	v342 = v266 - v341
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342+v343))))
	if v345 == v341 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L123
L125:
	;
	if v32 == int32(5) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373+v342<<(uint(int32(3))%32))))
	v378 = F_pg_detoast_datum(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L136
	}
L128:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v350
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v356 = F_ExecMergeNotMatched(m, v30+int32(108), v354, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L133
	}
L131:
	;
	if v356 == int32(0) {
		v127 = v241
		goto L49
	} else {
		goto L132
	}
L132:
	;
	v1867 = v356
	goto L20
L133:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_9), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_10), int32(_a_F_ExecModifyTable_6))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v378
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v382 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+88)) = uint16(v382)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+84)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = int32(base.Ui32(v381) >> (uint(int32(2)) % 32))
	if v269 != int32(118) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+56))
	v394 = v392
	goto L139
L138:
	;
	v394 = int32(0)
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+92)) = v394
	v400 = v332
	v402 = v30 + int32(80)
	goto L95
L140:
	;
	if v1769 == int32(0) {
		v127 = v241
		goto L49
	} else {
		goto L521
	}
L141:
	;
	v1761 = int32(0)
	v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v1766 = F_ExecDelete(m, v30+int32(108), v241, v400, v402, int32(1), v1761, v1762, v1761, v1761, v1761)
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L4
	} else {
		goto L520
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L4
	} else {
		goto L517
	}
L143:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v400|v402 != 0 {
		goto L212
	} else {
		goto L213
	}
L144:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+48)))
	if v642 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L145:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+48)))
	if v403 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v406 = int32(0)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+52))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+44))
	if v412 != 0 {
		goto L152
	} else {
		goto L153
	}
L147:
	;
	goto L148
L148:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v241)+36))
	if v590 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L149:
	;
	v560 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+48)) = uint8(v560)
	goto L148
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504+v241))) = v530
	goto L149
L151:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	F_ExecCheckPlanOutput(m, v493, int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L174
	}
L152:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v413 <= int32(0) {
		goto L151
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	F_ExecCheckPlanOutput(m, v483, int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L172
	}
L155:
	;
	v417 = v406
	v422 = v406
	v432 = v406
	goto L156
L156:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v412)+12))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v443+v417<<(uint(int32(2))%32))))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+26)))
	if v448 != 0 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	F_ExecCheckPlanOutput(m, v458, v453)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L164
	}
L158:
	;
	v455 = v417 + int32(1)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v455 < v456 {
		v417 = v455
		v422 = v452
		v432 = v453
		goto L156
	} else {
		goto L163
	}
L159:
	;
	v452 = int32(1)
	v453 = v432
	goto L158
L160:
	;
	goto L161
L161:
	;
	v450 = F_lappend(m, v432, v447)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	v452 = v422
	v453 = v450
	goto L158
L163:
	;
	goto L157
L164:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v464 = F_table_slot_create(m, v461, v409+int32(104))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+40)) = v464
	if v452 == int32(0) {
		goto L149
	} else {
		goto L166
	}
L166:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+52))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v471 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	F_ExecAssignExprContext(m, v409, l0)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L170
	}
L168:
	;
	v478 = v464
	v479 = v471
	goto L169
L169:
	;
	v481 = F_ExecBuildProjectionInfo(m, v453, v479, v478, l0, v470)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L171
	}
L170:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v241)+40))
	v478 = v477
	v479 = v476
	goto L169
L171:
	;
	v504 = int32(36)
	v530 = v481
	goto L150
L172:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v491 = F_table_slot_create(m, v488, v409+int32(104))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v504 = int32(40)
	v530 = v491
	goto L150
L174:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v501 = F_table_slot_create(m, v498, v409+int32(104))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v504 = int32(40)
	v530 = v501
	goto L150
L176:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v636 = int32(0)
	v638 = F_ExecInsert(m, v30+int32(108), v241, v629, v635, v636, v636)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L4
	} else {
		goto L184
	}
L177:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v241)+40))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v589)+8))
	if v594 == v595 {
		v629 = v589
		goto L176
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v590)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+12)) = v589
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v590)+80))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v590)+24))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)+8))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+12))
	m.T0[v606].(func(*base.Module, int32))(m, v604)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L182
	}
L180:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v594)+32))
	m.T0[v597].(func(*base.Module, int32, int32))(m, v593, v589)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v241)+40))
	v629 = v600
	goto L176
L182:
	;
	v609 = int32(_a_F_ExecModifyTable_11)
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v612
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v590)+32))
	v618 = m.T0[v617].(func(*base.Module, int32, int32, int32) int64)(m, v590+int32(8), v603, int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v610
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+4)))
	v624 = v622 & int32(_a_F_ExecModifyTable_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v604)+4)) = uint16(v624)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v604)+12))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	*(*uint16)(unsafe.Add(mBase, uint32(v604)+6)) = uint16(v627)
	v629 = v604
	goto L176
L184:
	;
	if v638 == int32(0) {
		v127 = v241
		goto L49
	} else {
		goto L185
	}
L185:
	;
	v1867 = v638
	goto L20
L186:
	;
	F_ExecInitUpdateProjection(m, l0, v241)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L4
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	if v402 != 0 {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	goto L188
L190:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v241)+36))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v677)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v678)+4)) = v647
	*(*int32)(unsafe.Add(mBase, uint32(v678)+12)) = v676
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v677)+80))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v677)+24))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)+8))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)+12))
	m.T0[v684].(func(*base.Module, int32))(m, v682)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L205
	}
L191:
	;
	v648 = int32(0)
	F_ExecForceStoreHeapTuple(m, v402, v647, v648)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+49)))
	if v653 == int32(1) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v675 = v648
	goto L190
L195:
	;
	F_LockTuple(m, v652, v400, int32(7))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L4
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v660 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[2]))
	if v660 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L197
L199:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecModifyTable[3])))
	if v662&int32(1) == int32(0) {
		goto L6
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v652)+188))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)+60))
	v670 = m.T0[v669].(func(*base.Module, int32, int32, int32, int32) int32)(m, v652, v400, int32(_a_F_ExecModifyTable_13), v647)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L4
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	if v670 == int32(0) {
		goto L16
	} else {
		goto L204
	}
L204:
	;
	v675 = v653
	goto L190
L205:
	;
	v687 = int32(_a_F_ExecModifyTable_11)
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v681)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v690
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v677)+32))
	v696 = m.T0[v695].(func(*base.Module, int32, int32, int32) int64)(m, v677+int32(8), v681, int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v688
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v682)+4)))
	v702 = v700 & int32(_a_F_ExecModifyTable_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v682)+4)) = uint16(v702)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v682)+12))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	*(*uint16)(unsafe.Add(mBase, uint32(v682)+6)) = uint16(v705)
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	v710 = F_ExecUpdate(m, v30+int32(108), v241, v400, v402, v647, v682, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	if v675 == int32(0) {
		v1769 = v710
		goto L140
	} else {
		goto L208
	}
L208:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	F_UnlockTuple(m, v714, v400, int32(7))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L209
	}
L209:
	;
	if v710 == int32(0) {
		v127 = v241
		goto L49
	} else {
		goto L210
	}
L210:
	;
	v1867 = v710
	goto L20
L211:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v30)+108))
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1742)+220)) = v1743
	v1867 = v1674
	goto L20
L212:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v30)+108))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+64))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v241)+164))
	if v725 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v1738 = F_ExecMergeNotMatched(m, v30+int32(108), v241, v720&int32(1))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L4
	} else {
		goto L515
	}
L215:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v241)+168))
	if v728 == int32(0) {
		v127 = v241
		goto L49
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v723)+4)) = v731
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	v734 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v723)+12)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v723)+8)) = v733
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+172)) = uint16(v734)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+168)) = int32(-1)
	if v402 != 0 {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	goto L217
L219:
	;
	v775 = v241 + int32(164)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v241)+176))
	if v776 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L220:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	F_ExecForceStoreHeapTuple(m, v402, v741, int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L4
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+49)))
	if v745 == int32(1) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	goto L219
L224:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	F_LockTuple(m, v748, v400, int32(7))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L4
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[2]))
	if v757 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v752 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+172)) = uint16(v752)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+168)) = v754
	goto L226
L228:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecModifyTable[3])))
	if v759&int32(1) == int32(0) {
		goto L6
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v764)+188))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v767)+60))
	v769 = m.T0[v768].(func(*base.Module, int32, int32, int32, int32) int32)(m, v764, v400, int32(_a_F_ExecModifyTable_13), v766)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L4
	} else {
		goto L232
	}
L231:
	;
	goto L230
L232:
	;
	if v769 == int32(0) {
		goto L15
	} else {
		goto L233
	}
L233:
	;
	goto L219
L234:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v798)))
	if v802 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L235:
	;
	v798 = v775
	v800 = v241 + int32(168)
	goto L234
L236:
	;
	goto L237
L237:
	;
	v781 = int32(_a_F_ExecModifyTable_11)
	v782 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v723)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v784
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v776)+24))
	v789 = m.T0[v788].(func(*base.Module, int32, int32, int32) int64)(m, v776, v723, v30+int32(152))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v782
	v794 = v241 + int32(168)
	if v789 == int64(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v797 = v794
	goto L241
L240:
	;
	v797 = v775
	goto L241
L241:
	;
	v798 = v797
	v800 = v794
	goto L234
L242:
	;
	v1700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+172)))
	if v1700 != 0 {
		goto L509
	} else {
		goto L510
	}
L243:
	;
	v1674 = int32(0)
	v1680 = int32(1)
	goto L242
L244:
	;
	goto L245
L245:
	;
	v808 = v722 + int32(124)
	v817 = int32(0)
	v818 = int32(1)
	v828 = v802
	goto L247
L246:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	v1671 = F_ExecProcessReturning(m, v30+int32(108), v241, int32(0), v1669, v1634, v1670)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L4
	} else {
		goto L508
	}
L247:
	;
	v838 = int32(0)
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	if v839 <= v838 {
		v1674 = v838
		v1680 = v818
		goto L242
	} else {
		goto L249
	}
L248:
	;
	v1674 = int32(0)
	v1680 = v818
	goto L242
L249:
	;
	v843 = v838
	goto L250
L250:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v828)+12))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v869+v843<<(uint(int32(2))%32))))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+152)) = int64(0)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	if v880 != 0 {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L248
L252:
	;
	v1662 = v843 + int32(1)
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	if v1662 < v1663 {
		v843 = v1662
		goto L250
	} else {
		goto L507
	}
L253:
	;
	v881 = int32(_a_F_ExecModifyTable_11)
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v723)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v884
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v880)+24))
	v889 = m.T0[v888].(func(*base.Module, int32, int32, int32) int64)(m, v880, v723, v30+int32(175))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L4
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v241)+116))
	v897 = int32(0)
	if base.B2i32(v896 == v897)|base.B2i32(v875 == int32(7)) == v897 {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v882
	if v889 == int64(0) {
		goto L252
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	if v875 == int32(2) {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	goto L260
L260:
	;
	v915 = v875 - int32(2)
	switch v915 {
	case 0:
		goto L270
	default:
		goto L14
	case 2:
		goto L269
	case 5:
		goto L266
	}
L261:
	;
	v908 = int32(4)
	goto L263
L262:
	;
	v908 = int32(5)
	goto L263
L263:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v30)+108))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v910)+8))
	F_ExecWithCheckOptions(m, v908, v241, v909, v911)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L4
	} else {
		goto L264
	}
L264:
	;
	goto L260
L265:
	;
	v1636 = int32(0)
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v241)+152))
	if v1637 == v1636 {
		v1674 = v1636
		v1680 = v818
		goto L242
	} else {
		goto L500
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+164)) = int32(0)
	v1634 = v817
	goto L265
L267:
	;
	if v1120 != int32(3) {
		goto L337
	} else {
		goto L338
	}
L268:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v30)+164))
	v1120 = v1119
	v1122 = v1117
	goto L267
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v722)+216)) = v873
	v1030 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+164)) = v1030
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v241)+52))
	if v1032 == v1030 {
		goto L304
	} else {
		goto L305
	}
L270:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v873)+8))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)+80))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v916)+24))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v919)+12))
	m.T0[v920].(func(*base.Module, int32))(m, v918)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	v923 = int32(_a_F_ExecModifyTable_11)
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v917)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v926
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v916)+32))
	v932 = m.T0[v931].(func(*base.Module, int32, int32, int32) int64)(m, v916+int32(8), v917, int32(0))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v924
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v918)+4)))
	v938 = v936 & int32(_a_F_ExecModifyTable_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v918)+4)) = uint16(v938)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v918)+12))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)))
	*(*uint16)(unsafe.Add(mBase, uint32(v918)+6)) = uint16(v941)
	*(*int32)(unsafe.Add(mBase, uint32(v722)+216)) = v873
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+164)) = int32(0)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v947)+28))
	m.T0[v948].(func(*base.Module, int32))(m, v918)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v944)+48))
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951)+116)))
	if v952 != int32(1) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v241)+52))
	if v959 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L275:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	if v955 != 0 {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	F_ExecOpenIndices(m, v241, int32(0))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	goto L274
L278:
	;
	if v1017 != 0 {
		v1117 = v918
		goto L268
	} else {
		goto L300
	}
L279:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v30)+164))
	v1017 = v1016
	goto L278
L280:
	;
	v1003 = F_ExecUpdateAct(m, v30+int32(108), v241, v400, int32(0), v918, v720&int32(1), v30+int32(152))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L4
	} else {
		goto L298
	}
L281:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959)+13)))
	if v962 == int32(1) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v965)+188))
	if v966 != 0 {
		goto L285
	} else {
		goto L286
	}
L283:
	;
	v988 = v959
	goto L284
L284:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988)+15)))
	if v989 != int32(1) {
		goto L280
	} else {
		goto L295
	}
L285:
	;
	F_ExecPendingInserts(m, v965)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L4
	} else {
		goto L288
	}
L286:
	;
	v970 = v965
	goto L287
L287:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v30)+112))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v30)+108))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v975)+104))
	v979 = F_ExecBRUpdateTriggers(m, v970, v971, v241, v400, int32(0), v918, v30+int32(164), v119, base.B2i32(v976 == int32(5)))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L4
	} else {
		goto L289
	}
L288:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	v970 = v969
	goto L287
L289:
	;
	if v979 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v30)+164))
	if v983 != 0 {
		v1120 = v983
		v1122 = v918
		goto L267
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v241)+52))
	if v985 == int32(0) {
		goto L280
	} else {
		goto L294
	}
L293:
	;
	v1674 = int32(0)
	v1680 = v818
	goto L242
L294:
	;
	v988 = v985
	goto L284
L295:
	;
	v992 = F_ExecIRUpdateTriggers(m, v724, v241, v402, v918)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L4
	} else {
		goto L296
	}
L296:
	;
	if v992 != 0 {
		goto L279
	} else {
		goto L297
	}
L297:
	;
	v1674 = int32(0)
	v1680 = v818
	goto L242
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+164)) = v1003
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+152)))
	if v1006&int32(1) == int32(0) {
		v1017 = v1003
		goto L278
	} else {
		goto L299
	}
L299:
	;
	v1011 = *(*float64)(unsafe.Add(mBase, uint32(v722)+232))
	*(*float64)(unsafe.Add(mBase, uint32(v722)+232)) = base.F64_add(v1011, float64(1))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v30)+148))
	v1674 = v1015
	v1680 = v818
	goto L242
L300:
	;
	F_ExecUpdateEpilogue(m, v30+int32(108), v30+int32(152), v241, v400, int32(0), v918)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L4
	} else {
		goto L301
	}
L301:
	;
	v1025 = *(*float64)(unsafe.Add(mBase, uint32(v722)+232))
	*(*float64)(unsafe.Add(mBase, uint32(v722)+232)) = base.F64_add(v1025, float64(1))
	v1117 = v918
	goto L268
L302:
	;
	if v1083 != 0 {
		v1117 = v817
		goto L268
	} else {
		goto L323
	}
L303:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v30)+164))
	v1083 = v1082
	goto L302
L304:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+64))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+8))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+12))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+188))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+96))
	v1079 = m.T0[v1078].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1070, v400, v1072, int32(0), v1074, v1075, int32(1), v119)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L4
	} else {
		goto L322
	}
L305:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032)+18)))
	if v1035 == int32(1) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+188))
	if v1039 != 0 {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	v1062 = v1032
	goto L308
L308:
	;
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062)+20)))
	if v1063 != int32(1) {
		goto L304
	} else {
		goto L319
	}
L309:
	;
	F_ExecPendingInserts(m, v1038)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L4
	} else {
		goto L312
	}
L310:
	;
	v1043 = v1038
	goto L311
L311:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v30)+112))
	v1045 = int32(0)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v30)+108))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1049)+104))
	v1053 = F_ExecBRDeleteTriggers(m, v1043, v1044, v241, v400, v1045, v1045, v30+int32(164), v119, base.B2i32(v1050 == int32(5)))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L4
	} else {
		goto L313
	}
L312:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	v1043 = v1042
	goto L311
L313:
	;
	if v1053 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v30)+164))
	if v1057 != 0 {
		v1120 = v1057
		v1122 = v817
		goto L267
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v241)+52))
	if v1059 == int32(0) {
		goto L304
	} else {
		goto L318
	}
L317:
	;
	v1674 = int32(0)
	v1680 = v818
	goto L242
L318:
	;
	v1062 = v1059
	goto L308
L319:
	;
	v1066 = F_ExecIRDeleteTriggers(m, v724, v241, v402)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L4
	} else {
		goto L320
	}
L320:
	;
	if v1066 != 0 {
		goto L303
	} else {
		goto L321
	}
L321:
	;
	v1674 = int32(0)
	v1680 = v818
	goto L242
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+164)) = v1079
	v1083 = v1079
	goto L302
L323:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v30)+108))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+204))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+104))
	if v1088 != int32(2) {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1107 = int32(0)
	F_ExecARDeleteTriggers(m, v1087, v241, v400, v1107, v1106, v1107)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L4
	} else {
		goto L335
	}
L325:
	;
	v1106 = v1086
	goto L324
L326:
	;
	goto L327
L327:
	;
	if v1086 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1106 = int32(0)
	goto L324
L329:
	;
	goto L330
L330:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086)+1)))
	if v1094 != int32(1) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1106 = v1086
	goto L324
L332:
	;
	goto L333
L333:
	;
	v1097 = int32(0)
	F_ExecARUpdateTriggers(m, v1087, v241, v1097, v1097, v400, v1097, v1097, v1097, v1086, v1097)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L4
	} else {
		goto L334
	}
L334:
	;
	v1106 = v1097
	goto L324
L335:
	;
	v1111 = *(*float64)(unsafe.Add(mBase, uint32(v722)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v722)+240)) = base.F64_add(v1111, float64(1))
	v1117 = v817
	goto L268
L336:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L4
	} else {
		goto L497
	}
L337:
	;
	switch v1120 {
	case 0:
		goto L342
	case 1, 5, 6:
		goto L336
	case 2:
		goto L341
	default:
		v1634 = v1122
		goto L265
	case 4:
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[4]))
	if int32(2) <= v1309 {
		goto L11
	} else {
		goto L394
	}
L340:
	;
	v1286 = int32(0)
	v1289 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[4]))
	if v1289 < int32(2) {
		v1674 = v1286
		v1680 = v1286
		goto L242
	} else {
		goto L389
	}
L341:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v30)+136))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v724)+64))
	if v1137 != v1138 {
		goto L13
	} else {
		goto L344
	}
L342:
	;
	if base.B2i32(v720&int32(1) == int32(0))|base.B2i32(v875 == int32(7)) != 0 {
		v1634 = v1122
		goto L265
	} else {
		goto L343
	}
L343:
	;
	v1133 = *(*int64)(unsafe.Add(mBase, uint32(v724)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v724)+112)) = v1133 + int64(1)
	v1634 = v1122
	goto L265
L344:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v30)+132))
	if base.Ui32(v1140) < base.Ui32(int32(3)) {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L4
	} else {
		goto L385
	}
L346:
	;
	v1272 = int32(0)
	goto L345
L347:
	;
	goto L348
L348:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[5]))
	if v1152 == v1140 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1272 = int32(1)
	goto L345
L350:
	;
	goto L351
L351:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[6]))
	if v1156 <= int32(0) {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v1272 = v1262
	goto L345
L353:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[7]))
	if v1160 == int32(0) {
		v1262 = int32(0)
		goto L352
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[8]))
	v1232 = int32(0)
	v1235 = v1156 - int32(1)
	goto L375
L356:
	;
	v1165 = v1160
	goto L357
L357:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+20))
	if v1171 == int32(4) {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	v1262 = int32(0)
	goto L352
L359:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+80))
	if v1225 != 0 {
		v1165 = v1225
		goto L357
	} else {
		goto L374
	}
L360:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1165)))
	if v1174 == int32(0) {
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v1177 = int32(1)
	if v1140 == v1174 {
		v1262 = v1177
		goto L352
	} else {
		goto L362
	}
L362:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+52))
	v1181 = v1179 - int32(1)
	if v1181 < int32(0) {
		goto L359
	} else {
		goto L363
	}
L363:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+48))
	v1187 = int32(0)
	v1190 = v1181
	goto L364
L364:
	;
	v1195 = int32(2)
	v1196 = base.I32_div_s(v1190-v1187, v1195)
	v1197 = v1196 + v1187
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1184+v1197<<(uint(v1195)%32))))
	if v1201 == v1140 {
		v1262 = v1177
		goto L352
	} else {
		goto L366
	}
L365:
	;
	goto L359
L366:
	;
	v1210 = base.B2i32(v1201-v1140 < int32(0)) | base.B2i32(base.Ui32(v1201) < base.Ui32(int32(3)))
	if v1210 != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1211 = v1197 + int32(1)
	goto L369
L368:
	;
	v1211 = v1187
	goto L369
L369:
	;
	if v1210 != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1214 = v1190
	goto L372
L371:
	;
	v1214 = v1197 - int32(1)
	goto L372
L372:
	;
	if v1211 <= v1214 {
		v1187 = v1211
		v1190 = v1214
		goto L364
	} else {
		goto L373
	}
L373:
	;
	goto L365
L374:
	;
	goto L358
L375:
	;
	v1240 = int32(2)
	v1241 = base.I32_div_s(v1235-v1232, v1240)
	v1242 = v1241 + v1232
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1230+v1242<<(uint(v1240)%32))))
	v1247 = base.B2i32(v1246 == v1140)
	if v1246 == v1140 {
		v1262 = v1247
		goto L352
	} else {
		goto L377
	}
L376:
	;
	v1262 = v1247
	goto L352
L377:
	;
	v1250 = base.B2i32(base.Ui32(v1246) < base.Ui32(v1140))
	if base.Ui32(v1246) < base.Ui32(v1140) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1251 = v1242 + int32(1)
	goto L380
L379:
	;
	v1251 = v1232
	goto L380
L380:
	;
	if base.Ui32(v1246) < base.Ui32(v1140) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1254 = v1235
	goto L383
L382:
	;
	v1254 = v1242 - int32(1)
	goto L383
L383:
	;
	if v1251 <= v1254 {
		v1232 = v1251
		v1235 = v1254
		goto L375
	} else {
		goto L384
	}
L384:
	;
	goto L376
L385:
	;
	if v1272 != 0 {
		goto L12
	} else {
		goto L386
	}
L386:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_14), int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3554), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L4
	} else {
		goto L390
	}
L390:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L4
	} else {
		goto L391
	}
L391:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_16), int32(0))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L4
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3561), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L4
	} else {
		goto L393
	}
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L394:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+4))
	v1315 = F_ExecUpdateLockMode(m, v724, v241)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L4
	} else {
		goto L395
	}
L395:
	;
	if v1314 == int32(0) {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	v1324 = int32(0)
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v724)+8))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v724)+64))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+188))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+104))
	v1331 = m.T0[v1330].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1312, v400, v1325, v1323, v1326, v1315, v1324, int32(2), v119)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L4
	} else {
		goto L401
	}
L397:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	v1320 = F_EvalPlanQualSlot(m, v808, v1312, v1319)
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L4
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	v1323 = v1322
	goto L396
L400:
	;
	v1323 = v1320
	goto L396
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+164)) = v1331
	if v1331 != 0 {
		goto L404
	} else {
		goto L405
	}
L402:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L4
	} else {
		goto L494
	}
L403:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v30)+136))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v724)+64))
	if v1449 != v1450 {
		goto L8
	} else {
		goto L449
	}
L404:
	;
	switch v1331 - int32(2) {
	case 0:
		goto L403
	default:
		goto L402
	case 2:
		v1674 = v1324
		v1680 = int32(0)
		goto L242
	}
L405:
	;
	goto L406
L406:
	;
	v1337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400)+4)))
	if v1337 == int32(_a_F_ExecModifyTable_12) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400))))
	v1341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400)+2)))
	if v1340&v1341 == int32(_a_F_ExecModifyTable_17) {
		goto L10
	} else {
		goto L410
	}
L408:
	;
	goto L409
L409:
	;
	if v1314 != 0 {
		v1445 = v818
		v1447 = v828
		goto L411
	} else {
		goto L412
	}
L410:
	;
	goto L409
L411:
	;
	if v1447 != 0 {
		v817 = v1122
		v818 = v1445
		v828 = v1447
		goto L247
	} else {
		goto L448
	}
L412:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	v1346 = F_EvalPlanQual(m, v808, v1312, v1345, v1323)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L4
	} else {
		goto L413
	}
L413:
	;
	if v1346 == int32(0) {
		v1674 = v1324
		v1680 = v818
		goto L242
	} else {
		goto L414
	}
L414:
	;
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1346)+4)))
	if v1350&int32(2) != 0 {
		v1674 = v1324
		v1680 = v818
		goto L242
	} else {
		goto L415
	}
L415:
	;
	v1353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241)+24)))
	v1354 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1346)+6)))
	if v1354 < v1353 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1346)+8))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+16))
	m.T0[v1357].(func(*base.Module, int32, int32))(m, v1346, v1353)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L4
	} else {
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1346)+20))
	v1362 = int32(1)
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1360+v1353-v1362))))
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+49)))
	if v1365 == v1362 {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	goto L418
L420:
	;
	v1368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+172)))
	if v1368 != 0 {
		goto L423
	} else {
		goto L424
	}
L421:
	;
	goto L422
L422:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[2]))
	if v1384 != 0 {
		goto L428
	} else {
		goto L429
	}
L423:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	F_UnlockTuple(m, v1369, v30+int32(168), int32(7))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L4
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	F_LockTuple(m, v1375, v400, int32(7))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L4
	} else {
		goto L427
	}
L426:
	;
	goto L425
L427:
	;
	v1379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+172)) = uint16(v1379)
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+168)) = v1381
	goto L422
L428:
	;
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecModifyTable[3])))
	if v1386&int32(1) == int32(0) {
		goto L6
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+188))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+60))
	v1395 = m.T0[v1394].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1312, v400, int32(_a_F_ExecModifyTable_13), v1392)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L4
	} else {
		goto L432
	}
L431:
	;
	goto L430
L432:
	;
	if v1395 == int32(0) {
		goto L9
	} else {
		goto L433
	}
L433:
	;
	if v818&(v1364^int32(-1)) == int32(0) {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v800)))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v722)+36))
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+20))
	if v1427 == int32(0) {
		goto L441
	} else {
		goto L442
	}
L435:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v241)+176))
	if v1404 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v1445 = int32(1)
	v1447 = v828
	goto L411
L437:
	;
	goto L438
L438:
	;
	v1408 = int32(_a_F_ExecModifyTable_11)
	v1409 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1]))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v723)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v1411
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+24))
	v1416 = m.T0[v1415].(func(*base.Module, int32, int32, int32) int64)(m, v1404, v723, v30+int32(175))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L4
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[1])) = v1409
	if v1416 == int64(0) {
		goto L434
	} else {
		goto L440
	}
L440:
	;
	v1445 = int32(1)
	v1447 = v828
	goto L411
L441:
	;
	v1445 = int32(0)
	v1447 = v1425
	goto L411
L442:
	;
	goto L443
L443:
	;
	if v1425 == int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1433 = int32(0)
	v1674 = v1433
	v1680 = v1433
	goto L242
L445:
	;
	goto L446
L446:
	;
	v1435 = int32(0)
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v241)+172))
	if v1436 == v1435 {
		v1445 = v1435
		v1447 = v1425
		goto L411
	} else {
		goto L447
	}
L447:
	;
	v1439 = *(*float64)(unsafe.Add(mBase, uint32(v1427)+384))
	*(*float64)(unsafe.Add(mBase, uint32(v1427)+384)) = base.F64_add(v1439, float64(1))
	v1445 = v1435
	v1447 = v1425
	goto L411
L448:
	;
	v1674 = int32(0)
	v1680 = v1445
	goto L242
L449:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v30)+132))
	if base.Ui32(v1452) < base.Ui32(int32(3)) {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L4
	} else {
		goto L490
	}
L451:
	;
	v1584 = int32(0)
	goto L450
L452:
	;
	goto L453
L453:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[5]))
	if v1464 == v1452 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v1584 = int32(1)
	goto L450
L455:
	;
	goto L456
L456:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[6]))
	if v1468 <= int32(0) {
		goto L458
	} else {
		goto L459
	}
L457:
	;
	v1584 = v1574
	goto L450
L458:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[7]))
	if v1472 == int32(0) {
		v1574 = int32(0)
		goto L457
	} else {
		goto L461
	}
L459:
	;
	goto L460
L460:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, _c_F_ExecModifyTable[8]))
	v1544 = int32(0)
	v1547 = v1468 - int32(1)
	goto L480
L461:
	;
	v1477 = v1472
	goto L462
L462:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+20))
	if v1483 == int32(4) {
		goto L464
	} else {
		goto L465
	}
L463:
	;
	v1574 = int32(0)
	goto L457
L464:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+80))
	if v1537 != 0 {
		v1477 = v1537
		goto L462
	} else {
		goto L479
	}
L465:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1477)))
	if v1486 == int32(0) {
		goto L464
	} else {
		goto L466
	}
L466:
	;
	v1489 = int32(1)
	if v1452 == v1486 {
		v1574 = v1489
		goto L457
	} else {
		goto L467
	}
L467:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+52))
	v1493 = v1491 - int32(1)
	if v1493 < int32(0) {
		goto L464
	} else {
		goto L468
	}
L468:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+48))
	v1499 = int32(0)
	v1502 = v1493
	goto L469
L469:
	;
	v1507 = int32(2)
	v1508 = base.I32_div_s(v1502-v1499, v1507)
	v1509 = v1508 + v1499
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1496+v1509<<(uint(v1507)%32))))
	if v1513 == v1452 {
		v1574 = v1489
		goto L457
	} else {
		goto L471
	}
L470:
	;
	goto L464
L471:
	;
	v1522 = base.B2i32(v1513-v1452 < int32(0)) | base.B2i32(base.Ui32(v1513) < base.Ui32(int32(3)))
	if v1522 != 0 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v1523 = v1509 + int32(1)
	goto L474
L473:
	;
	v1523 = v1499
	goto L474
L474:
	;
	if v1522 != 0 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v1526 = v1502
	goto L477
L476:
	;
	v1526 = v1509 - int32(1)
	goto L477
L477:
	;
	if v1523 <= v1526 {
		v1499 = v1523
		v1502 = v1526
		goto L469
	} else {
		goto L478
	}
L478:
	;
	goto L470
L479:
	;
	goto L463
L480:
	;
	v1552 = int32(2)
	v1553 = base.I32_div_s(v1547-v1544, v1552)
	v1554 = v1553 + v1544
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1542+v1554<<(uint(v1552)%32))))
	v1559 = base.B2i32(v1558 == v1452)
	if v1558 == v1452 {
		v1574 = v1559
		goto L457
	} else {
		goto L482
	}
L481:
	;
	v1574 = v1559
	goto L457
L482:
	;
	v1562 = base.B2i32(base.Ui32(v1558) < base.Ui32(v1452))
	if base.Ui32(v1558) < base.Ui32(v1452) {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v1563 = v1554 + int32(1)
	goto L485
L484:
	;
	v1563 = v1544
	goto L485
L485:
	;
	if base.Ui32(v1558) < base.Ui32(v1452) {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v1566 = v1547
	goto L488
L487:
	;
	v1566 = v1554 - int32(1)
	goto L488
L488:
	;
	if v1563 <= v1566 {
		v1544 = v1563
		v1547 = v1566
		goto L480
	} else {
		goto L489
	}
L489:
	;
	goto L481
L490:
	;
	if v1584 != 0 {
		goto L7
	} else {
		goto L491
	}
L491:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_14), int32(0))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L4
	} else {
		goto L492
	}
L492:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3751), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L4
	} else {
		goto L493
	}
L493:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L494:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v30)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v1602
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_18), v30+int32(32))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L4
	} else {
		goto L495
	}
L495:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3757), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L4
	} else {
		goto L496
	}
L496:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L497:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v30)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v1618
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_19), v30-int32(-64))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L4
	} else {
		goto L498
	}
L498:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3766), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L4
	} else {
		goto L499
	}
L499:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L500:
	;
	switch v915 {
	case 0:
		goto L246
	default:
		goto L501
	case 2:
		goto L502
	case 5:
		v1674 = v1636
		v1680 = v818
		goto L242
	}
L501:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L4
	} else {
		goto L504
	}
L502:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v241)+44))
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	v1646 = F_ExecProcessReturning(m, v30+int32(108), v241, int32(1), v1643, int32(0), v1645)
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L4
	} else {
		goto L503
	}
L503:
	;
	v1674 = v1646
	v1680 = v818
	goto L242
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v875
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_20), v30)
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L4
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3798), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L4
	} else {
		goto L506
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	goto L251
L508:
	;
	v1674 = v1671
	v1680 = v818
	goto L242
L509:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	F_UnlockTuple(m, v1701, v30+int32(168), int32(7))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L4
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	if v1680 != 0 {
		v1769 = v1674
		goto L140
	} else {
		goto L513
	}
L512:
	;
	goto L511
L513:
	;
	if v1674 != 0 {
		goto L211
	} else {
		goto L514
	}
L514:
	;
	goto L214
L515:
	;
	if v1738 == int32(0) {
		v127 = v241
		goto L49
	} else {
		goto L516
	}
L516:
	;
	v1867 = v1738
	goto L20
L517:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_0), int32(0))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L4
	} else {
		goto L518
	}
L518:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_21), int32(_a_F_ExecModifyTable_6))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L4
	} else {
		goto L519
	}
L519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L520:
	;
	v1769 = v1766
	goto L140
L521:
	;
	v1867 = v1769
	goto L20
L522:
	;
	F_ExecPendingInserts(m, v33)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L4
	} else {
		goto L525
	}
L523:
	;
	goto L524
L524:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	switch v1801 - int32(2) {
	case 0:
		goto L527
	case 1:
		goto L531
	case 2:
		goto L530
	case 3:
		goto L529
	default:
		goto L528
	}
L525:
	;
	goto L524
L526:
	;
	v1863 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v1863)
	v1867 = int32(0)
	goto L20
L527:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASUpdateTriggers(m, v1858, v1800, v1859)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L4
	} else {
		goto L551
	}
L528:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L4
	} else {
		goto L548
	}
L529:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v1820&int32(4) != 0 {
		goto L538
	} else {
		goto L539
	}
L530:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASDeleteTriggers(m, v1816, v1800, v1817)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L4
	} else {
		goto L537
	}
L531:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1804)+128))
	if v1805 == int32(2) {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	F_ExecASUpdateTriggers(m, v1808, v1800, v1809)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L4
	} else {
		goto L535
	}
L533:
	;
	goto L534
L534:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASInsertTriggers(m, v1812, v1800, v1813)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L4
	} else {
		goto L536
	}
L535:
	;
	goto L534
L536:
	;
	goto L526
L537:
	;
	goto L526
L538:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASDeleteTriggers(m, v1823, v1800, v1824)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L4
	} else {
		goto L541
	}
L539:
	;
	v1828 = v1820
	goto L540
L540:
	;
	if v1828&int32(2) != 0 {
		goto L542
	} else {
		goto L543
	}
L541:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v1828 = v1827
	goto L540
L542:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASUpdateTriggers(m, v1831, v1800, v1832)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L4
	} else {
		goto L545
	}
L543:
	;
	v1836 = v1828
	goto L544
L544:
	;
	if v1836&int32(1) == int32(0) {
		goto L526
	} else {
		goto L546
	}
L545:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v1836 = v1835
	goto L544
L546:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecASInsertTriggers(m, v1841, v1800, v1842)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L4
	} else {
		goto L547
	}
L547:
	;
	goto L526
L548:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_0), int32(0))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L4
	} else {
		goto L549
	}
L549:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_22), int32(_a_F_ExecModifyTable_23))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L4
	} else {
		goto L550
	}
L550:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L551:
	;
	goto L526
L552:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_24), int32(0))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L4
	} else {
		goto L553
	}
L553:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_25), int32(_a_F_ExecModifyTable_6))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L4
	} else {
		goto L554
	}
L554:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L555:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_26), int32(0))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L4
	} else {
		goto L556
	}
L556:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(_a_F_ExecModifyTable_27), int32(_a_F_ExecModifyTable_6))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L4
	} else {
		goto L557
	}
L557:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L558:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_28), int32(0))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L4
	} else {
		goto L559
	}
L559:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3345), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L4
	} else {
		goto L560
	}
L560:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L561:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_29), int32(0))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L4
	} else {
		goto L562
	}
L562:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3507), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L4
	} else {
		goto L563
	}
L563:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L564:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L4
	} else {
		goto L565
	}
L565:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_30), int32(0))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L4
	} else {
		goto L566
	}
L566:
	;
	F_errhint(m, int32(_a_F_ExecModifyTable_31), int32(0))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L4
	} else {
		goto L567
	}
L567:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3543), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L4
	} else {
		goto L568
	}
L568:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = int32(_a_F_ExecModifyTable_32)
	F_errmsg(m, int32(_a_F_ExecModifyTable_33), v30+int32(16))
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L4
	} else {
		goto L570
	}
L570:
	;
	F_errhint(m, int32(_a_F_ExecModifyTable_34), int32(0))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L4
	} else {
		goto L571
	}
L571:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3551), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L4
	} else {
		goto L572
	}
L572:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L573:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L4
	} else {
		goto L574
	}
L574:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_35), int32(0))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L4
	} else {
		goto L575
	}
L575:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3581), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L4
	} else {
		goto L576
	}
L576:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L577:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L4
	} else {
		goto L578
	}
L578:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_36), int32(0))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L4
	} else {
		goto L579
	}
L579:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3625), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L4
	} else {
		goto L580
	}
L580:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L581:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_28), int32(0))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L4
	} else {
		goto L582
	}
L582:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3679), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L4
	} else {
		goto L583
	}
L583:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L584:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L4
	} else {
		goto L585
	}
L585:
	;
	F_errmsg(m, int32(_a_F_ExecModifyTable_30), int32(0))
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L4
	} else {
		goto L586
	}
L586:
	;
	F_errhint(m, int32(_a_F_ExecModifyTable_31), int32(0))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L4
	} else {
		goto L587
	}
L587:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3740), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L4
	} else {
		goto L588
	}
L588:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = int32(_a_F_ExecModifyTable_32)
	F_errmsg(m, int32(_a_F_ExecModifyTable_33), v30+int32(48))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L4
	} else {
		goto L590
	}
L590:
	;
	F_errhint(m, int32(_a_F_ExecModifyTable_34), int32(0))
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L4
	} else {
		goto L591
	}
L591:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_1), int32(3748), int32(_a_F_ExecModifyTable_15))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L4
	} else {
		goto L592
	}
L592:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L593:
	;
	F_errmsg_internal(m, int32(_a_F_ExecModifyTable_37), int32(0))
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L4
	} else {
		goto L594
	}
L594:
	;
	F_errfinish(m, int32(_a_F_ExecModifyTable_38), int32(1355), int32(_a_F_ExecModifyTable_39))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L4
	} else {
		goto L595
	}
L595:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecProjectSRF(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v195 int32
	_ = v195
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v336 int64
	_ = v336
	var v338 int64
	_ = v338
	var v343 int64
	_ = v343
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int64
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v581 int64
	_ = v581
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int64
	_ = v596
	var v597 int32
	_ = v597
	var v614 int32
	_ = v614
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	v3 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	m.T0[v25].(func(*base.Module, int32))(m, v23)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = int32(_a_F_ExecProjectSRF_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v33
	v35 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v35)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v35 < v37 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return int32(0)
L4:
	;
	v40 = l0
	v41 = l1
	v49 = v22
	v50 = v23
	v51 = v3
	v52 = v31
	v53 = v3
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v31
	goto L3
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	v62 = v61 + v51
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v66 = v63 + v51<<(uint(int32(3))%32)
	v68 = v51 << (uint(int32(2)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v40)+108))
	v70 = v68 + v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71+v68)))
	if v41 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v52
	if v614&int32(1) == int32(0) {
		goto L3
	} else {
		goto L117
	}
L9:
	;
	v623 = v51 + int32(1)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v40)+112))
	if v623 < v624 {
		v51 = v623
		v53 = v614
		goto L7
	} else {
		goto L116
	}
L10:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v83 == int32(397) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v76 != int32(2) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = int64(0)
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v81)
	v614 = v53
	goto L9
L13:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+120))
	v87 = m.G0
	v89 = v87 - int32(80)
	m.G0 = v89
	F_check_stack_depth(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
	v596 = m.T0[v595].(func(*base.Module, int32, int32, int32) int64)(m, v73, v49, v62)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L115
	}
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v73)+44))
	if v93 != 0 {
		v477 = v93
		goto L19
	} else {
		goto L20
	}
L17:
	;
	m.G0 = v89 + int32(80)
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = v581
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v590 = base.B2i32(v587 != int32(2)) | v53
	if v587 != int32(1) {
		v614 = v590
		goto L9
	} else {
		goto L114
	}
L18:
	;
	v581 = int64(0)
	goto L17
L19:
	;
	v495 = int32(_a_F_ExecProjectSRF_0)
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0]))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v73)+48))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v499
	v503 = F_tuplestore_gettupleslot(m, v477, int32(1), int32(0), v498)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L101
	}
L20:
	;
	v94 = base.I64_extend_i32_u(v73)
	goto L24
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L97
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L93
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L89
	}
L24:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v73)+60))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+58)))
	if v117 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(2)
	v423 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v423)
	goto L18
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = v89 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = int32(389)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v73)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+28)) = int64(4294967299)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v224
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+26)))
	if v230 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v120 = int32(_a_F_ExecProjectSRF_0)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v86
	if v122 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+58)) = uint8(v195)
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v121
	goto L26
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v127 <= int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v136 = int32(0)
	goto L33
L33:
	;
	v156 = v116 + int32(24) + v136<<(uint(int32(4))%32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157+v136<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	v165 = m.T0[v164].(func(*base.Module, int32, int32, int32) int64)(m, v161, v49, v156+int32(8))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L30
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = v165
	v169 = v136 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v169 < v170 {
		v136 = v169
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v293 = v89 + int32(48)
	F_pgstat_init_function_usage(m, v116, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L46
	}
L38:
	;
	v233 = int32(0)
	v234 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+18)))
	if v234 <= v233 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v240 = v233
	goto L40
L40:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v240<<(uint(int32(4))%32))+32)))
	if v261 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v267 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v267)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(2)
	goto L18
L42:
	;
	v265 = v240 + int32(1)
	if v234 != v265 {
		v240 = v265
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	goto L37
L46:
	;
	v296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116)+16)) = uint8(v296)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+36)) = v296
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v302 = m.T0[v301].(func(*base.Module, int32) int64)(m, v116)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v304)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v306
	v316 = m.G0
	v318 = v316 - int32(16)
	m.G0 = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v320 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	if v355 != int32(2) {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	F___clock_gettime(m, int32(1), v318)
	mBase = m.M
	v323 = int32(_a_F_ExecProjectSRF_1)
	v324 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProjectSRF[1]))
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v293)+16))
	v327 = int64(*(*int32)(unsafe.Add(mBase, uint32(v318)+8)))
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v318)))
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v293)+24))
	v333 = v327 + v328*int64(1000000000) - v332
	*(*int64)(unsafe.Add(mBase, _c_F_ExecProjectSRF[1])) = v326 + v333
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v293)+8))
	if v306 != int32(1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	m.G0 = v318 + int32(16)
	goto L48
L52:
	;
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v338 + int64(1)
	goto L54
L53:
	;
	goto L54
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v336 + v333
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v320)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+16)) = v343 + (v333 - v324 + v326)
	goto L51
L55:
	;
	if v355 != int32(1) {
		goto L23
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	if v371 != 0 {
		goto L22
	} else {
		goto L62
	}
L58:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v360 != int32(1) {
		v581 = v302
		goto L17
	} else {
		goto L59
	}
L59:
	;
	v363 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+58)) = uint8(v363)
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+59)))
	if v365 != 0 {
		v581 = v302
		goto L17
	} else {
		goto L60
	}
L60:
	;
	F_RegisterExprContextCallback(m, v49, int32(681), v94)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v369 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+59)) = uint8(v369)
	v581 = v302
	goto L17
L62:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	if v372 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v89)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+44)) = v372
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v73)+48))
	if v375 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	goto L25
L66:
	;
	v378 = int32(_a_F_ExecProjectSRF_0)
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0]))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v73)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v73)+52))
	if v383 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	if v373 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	v388 = v383
	goto L71
L70:
	;
	if v373 == int32(0) {
		goto L21
	} else {
		goto L72
	}
L71:
	;
	v390 = F_MakeSingleTupleTableSlot(m, v388, int32(_a_F_ExecProjectSRF_2))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	v386 = F_CreateTupleDescCopy(m, v373)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v388 = v386
	goto L71
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+48)) = v390
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v379
	goto L68
L75:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+59)))
	if v408 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v73)+52))
	if v399 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_tupledesc_match(m, v399, v373)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	if v402 != int32(-1) {
		goto L75
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	F_FreeTupleDesc(m, v373)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L75
L83:
	;
	F_RegisterExprContextCallback(m, v49, int32(681), v94)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	v414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+59)) = uint8(v414)
	goto L85
L87:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v73)+44))
	if v418 == int32(0) {
		goto L24
	} else {
		goto L88
	}
L88:
	;
	v477 = v418
	goto L19
L89:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v432
	F_errmsg(m, int32(_a_F_ExecProjectSRF_3), v89)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ExecProjectSRF_4), int32(688), int32(_a_F_ExecProjectSRF_5))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_ExecProjectSRF_6), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ExecProjectSRF_4), int32(669), int32(_a_F_ExecProjectSRF_5))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_ExecProjectSRF_7), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_ExecProjectSRF_4), int32(898), int32(_a_F_ExecProjectSRF_8))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSRF[0])) = v496
	if v503 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v507 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v507
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+56)))
	if v509 == v507 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v73)+44))
	F_tuplestore_end(m, v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L113
	}
L105:
	;
	v512 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v512)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v73)+48))
	v515 = F_ExecFetchSlotHeapTupleDatum(m, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v73)+48))
	v518 = int32(*(*int16)(unsafe.Add(mBase, uint32(v517)+6)))
	if v518 <= int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v581 = v515
	goto L17
L109:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v517)+8))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+16))
	m.T0[v523].(func(*base.Module, int32, int32))(m, v517, int32(1))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v517)+20))
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v527)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v517)+16))
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v529)))
	v581 = v530
	goto L17
L112:
	;
	goto L111
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(2)
	v538 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v538)
	goto L18
L114:
	;
	v593 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+116)) = uint8(v593)
	v614 = v590
	goto L9
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = v596
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
	v614 = v53
	goto L9
L116:
	;
	goto L8
L117:
	;
	v632 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
	v634 = v632 & int32(_a_F_ExecProjectSRF_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v634)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	*(*uint16)(unsafe.Add(mBase, uint32(v50)+6)) = uint16(v637)
	goto L118
L118:
	;
	return v50
}
func F_ExecSort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v204 int64
	_ = v204
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v225 int64
	_ = v225
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSort[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v249 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v243 = v17
	goto L6
L8:
	;
	goto L9
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v29 = v23 | v24<<(uint(v19)%32)&int32(2)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v30 == v19 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	if v59 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22+v33<<(uint(int32(3))%32))+96))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSort[1]))
	v46 = F_tuplesort_begin_datum(m, v37, v39, v41, v43, v45, v29)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSort[1]))
	v56 = F_tuplesort_begin_heap(m, v22, v48, v49, v50, v51, v52, v54, int32(0), v29)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	v58 = v46
	goto L10
L15:
	;
	v58 = v56
	goto L10
L16:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+236))
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v58
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+149)))
	if v92 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L19:
	;
	goto L18
L20:
	;
	goto L19
L21:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v58)+72)) = uint32(v62)
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+68)) = uint8(v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
	if v80 != 0 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	if int64(1073741823) < v62 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if int64(1073741823) < v62 {
		goto L20
	} else {
		goto L27
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+232))
	if v68 != int32(-1) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	goto L21
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	v83 = v82
	goto L30
L29:
	;
	v83 = v79
	goto L30
L30:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v84
	goto L20
L31:
	;
	F_tuplesort_performsort(m, v58)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L59
	}
L32:
	;
	goto L35
L33:
	;
	goto L34
L34:
	;
	goto L49
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v102 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v106 = m.T0[v105].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	if v106 == int32(0) {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
	if v110&int32(2) != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106)+6)))
	if v113 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	m.T0[v118].(func(*base.Module, int32, int32))(m, v106, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	F_tuplesort_putdatum(m, v58, v122, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	goto L35
L49:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v134 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v138 = m.T0[v137].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	if v138 == int32(0) {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+4)))
	if v142&int32(2) != 0 {
		goto L31
	} else {
		goto L57
	}
L57:
	;
	F_tuplesort_puttupleslot(m, v58, v138)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	goto L49
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v15
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v157)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)) = uint8(v159)
	v161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v163 == int32(0) {
		v243 = v58
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v166 != int32(1) {
		v243 = v58
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSort[2]))
	v175 = v163 + v170<<(uint(int32(4))%32) + int32(8)
	v176 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v58)+128))
	if v180 == v176 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v243 = v58
	goto L6
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v218
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v58)+112))
	v225 = base.I64_div_s(v221+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v175)+8)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v58)+124))
	switch v227 - int32(3) {
	case 0:
		goto L78
	case 1:
		v238 = v227
		goto L75
	case 2:
		goto L77
	default:
		goto L76
	}
L64:
	;
	if v197&int32(255) != base.B2i32(v180 != int32(0)) {
		goto L70
	} else {
		goto L71
	}
L65:
	;
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v58)+96))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v58)+88))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+120)))
	v197 = v186
	v198 = v183 - v184
	goto L64
L66:
	;
	goto L67
L67:
	;
	v187 = F_LogicalTapeSetBlocks(m, v180)
	mBase = m.M
	v189 = v187 << (uint(int64(13)) % 64)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+120)))
	if v191 != 0 {
		v197 = int32(1)
		v198 = v189
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+120)) = uint8(v192)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+112)) = v189
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v58)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+124)) = v195
	v218 = v176
	goto L63
L69:
	;
	v218 = int32(1)
	goto L63
L70:
	;
	if v197&int32(1) != 0 {
		v218 = v176
		goto L63
	} else {
		goto L74
	}
L71:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v58)+112))
	if v198 <= v204 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+120)) = uint8(v197)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+112)) = v198
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v58)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+124)) = v208
	if v197&int32(1) == int32(0) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v218 = v176
	goto L63
L74:
	;
	goto L69
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v238
	goto L62
L76:
	;
	v238 = int32(0)
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = int32(8)
	goto L62
L78:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+69)))
	if v232 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v233 = int32(1)
	goto L81
L80:
	;
	v233 = int32(2)
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v233
	goto L62
L82:
	;
	return v248
L83:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	m.T0[v253].(func(*base.Module, int32))(m, v248)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v276 = int32(0)
	v278 = F_tuplesort_gettupleslot(m, v243, base.B2i32(v15 == int32(1)), v276, v248, v276)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L90
	}
L86:
	;
	v258 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v248)+20))
	v262 = F_tuplesort_getdatum(m, v243, base.B2i32(v15 == int32(1)), v258, v259, v260, v258)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	if v262 == int32(0) {
		goto L82
	} else {
		goto L88
	}
L88:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248)+4)))
	v268 = v266 & int32(_a_F_ExecSort_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v248)+4)) = uint16(v268)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	*(*uint16)(unsafe.Add(mBase, uint32(v248)+6)) = uint16(v271)
	goto L89
L89:
	;
	return v248
L90:
	;
	goto L82
}
func F_ExecWithCheckOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v232 int64
	_ = v232
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(144)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	if v22 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = F_MakePerTupleExprContext(m, l3)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v27 = v22
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = l2
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v37 = v5
	goto L9
L4:
	;
	return
L5:
	;
	v27 = v25
	goto L3
L6:
	;
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = base.I64_rotl(v260, int64(32))
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_0), v18+int32(128))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L75
	}
L7:
	;
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = base.I64_rotl(v246, int64(32))
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_1), v18+int32(96))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L73
	}
L8:
	;
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = base.I64_rotl(v232, int64(32))
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_2), v18-int32(-64))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L71
	}
L9:
	;
	v46 = int32(0)
	if v30 == v46 {
		v56 = v46
		goto L11
	} else {
		goto L12
	}
L10:
	;
	m.G0 = v18 + int32(144)
	return
L11:
	;
	if v29 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v50 <= v37 {
		v56 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v56 = v52 + v37<<(uint(int32(2))%32)
	goto L11
L14:
	;
	goto L10
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if base.B2i32(v56 == int32(0))|base.B2i32(v61 <= v37) != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v64 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v68 != l0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v37 = v37 + int32(1)
	goto L9
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64+v37<<(uint(int32(2))%32))))
	if v73 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v76 = int32(_a_F_ExecWithCheckOptions_3)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWithCheckOptions[0]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWithCheckOptions[0])) = v79
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
	v84 = m.T0[v83].(func(*base.Module, int32, int32, int32) int64)(m, v73, v27, v18+int32(143))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWithCheckOptions[0])) = v77
	if v84 != int64(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	switch v90 {
	case 0:
		goto L27
	case 1, 2:
		goto L26
	case 3:
		goto L24
	case 4, 5:
		goto L25
	default:
		goto L23
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L68
	}
L24:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L63
	}
L25:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L58
	}
L26:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L53
	}
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v91 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+56))
	v123 = F_ExecBuildSlotValueDescription(m, v122, v119, v121, v118)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L44
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v96 = F_build_attrmap_by_name_if_req(m, v92, v94, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v112 = F_ExecGetInsertedCols(m, l1, l3)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L41
	}
L32:
	;
	if v96 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v100 = F_MakeTupleTableSlot(m, v94, int32(_a_F_ExecWithCheckOptions_4), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	v104 = l2
	goto L35
L35:
	;
	v105 = F_ExecGetInsertedCols(m, v91, l3)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v102 = F_execute_attr_map_slot(m, v96, l2, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v104 = v102
	goto L35
L38:
	;
	v107 = F_ExecGetUpdatedCols(m, v91, l3)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v109 = F_bms_union(m, v105, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v118 = v109
	v119 = v104
	v120 = v111
	v121 = v94
	goto L28
L41:
	;
	v114 = F_ExecGetUpdatedCols(m, l1, l3)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v116 = F_bms_union(m, v112, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v118 = v116
	v119 = l2
	v120 = v20
	v121 = v21
	goto L28
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(260))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v132
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_5), v18+int32(32))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	if v123 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v123
	v143 = F_errdetail(m, int32(_a_F_ExecWithCheckOptions_6), v18+int32(16))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2354), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	if v150 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v158
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_9), v18+int32(48))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2367), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	if v170 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v178
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_10), v18+int32(80))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2380), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	if v190 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v198
	F_errmsg(m, int32(_a_F_ExecWithCheckOptions_11), v18+int32(112))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2392), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v214
	F_errmsg_internal(m, int32(_a_F_ExecWithCheckOptions_12), v18)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2395), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2362), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2375), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errfinish(m, int32(_a_F_ExecWithCheckOptions_7), int32(2387), int32(_a_F_ExecWithCheckOptions_8))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__equalCaseWhen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = F_equal(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v11 = F_equal(m, v9, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v14 = v11
				return v14
			}
		} else {
			v14 = int32(0)
			return v14
		}
	}
}
func F_each_worker_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_MemoryContextDelete(m, v30)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L52
	}
L2:
	;
	v103 = v98
	v104 = v95
	v106 = v96
	goto L29
L3:
	;
	v95 = v4
	v96 = v4
	v98 = int32(2)
	goto L2
L4:
	;
	v95 = v91
	v96 = v45
	v98 = int32(1)
	goto L2
L5:
	;
	return
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if v17&int32(32) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L25
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0]))
	v30 = F_AllocSetContextCreateInternal(m, v25, int32(_a_F_each_worker_jsonb_0), int32(0), int32(_a_F_each_worker_jsonb_1), int32(_a_F_each_worker_jsonb_2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v34 = F_JsonbIteratorInit(m, v15+int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v34
	v42 = F_JsonbIteratorNext(m, v12+int32(76), v12+int32(40), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	v44 = int32(_a_F_each_worker_jsonb_3)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0])) = v30
	v48 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v52 = F_cstring_to_text_with_len(m, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	switch v42 {
	case 0:
		goto L1
	case 1:
		goto L13
	default:
		goto L3
	}
L15:
	;
	v57 = v12 + int32(40)
	v59 = F_JsonbIteratorNext(m, v12+int32(76), v57, int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = base.I64_extend_i32_u(v52)
	if l2 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = F_JsonbValueToJsonb(m, v57)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v91 = v65
	goto L4
L21:
	;
	v70 = F_JsonbValueAsText(m, v12+int32(40))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v72)
	v95 = v4
	v96 = v45
	v98 = int32(0)
	goto L2
L24:
	;
	v91 = v70
	goto L4
L25:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	F_errmsg(m, int32(_a_F_each_worker_jsonb_4), v12)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_each_worker_jsonb_5), int32(1989), int32(_a_F_each_worker_jsonb_6))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	switch v103 {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L33
	}
L31:
	;
	v103 = int32(0)
	v106 = v141
	goto L29
L32:
	;
	v103 = int32(1)
	v104 = v181
	v106 = v179
	goto L29
L33:
	;
	goto L38
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = base.I64_extend_i32_u(v104)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	F_tuplestore_putvalues(m, v111, v112, v12+int32(16), v12+int32(14))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	v179 = v106
	v181 = int32(0)
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0])) = v106
	F_MemoryContextReset(m, v30)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v103 = int32(2)
	goto L29
L38:
	;
	v138 = F_JsonbIteratorNext(m, v12+int32(76), v12+int32(40), int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	v140 = int32(_a_F_each_worker_jsonb_3)
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_each_worker_jsonb[0])) = v30
	v144 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v144)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v148 = F_cstring_to_text_with_len(m, v146, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	switch v138 {
	case 0:
		goto L1
	case 1:
		goto L40
	default:
		goto L38
	}
L42:
	;
	v155 = F_JsonbIteratorNext(m, v12+int32(76), v12+int32(40), int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = base.I64_extend_i32_u(v148)
	if l2 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v159 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v170 = F_JsonbValueToJsonb(m, v12+int32(40))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L51
	}
L47:
	;
	v162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v162)
	goto L31
L48:
	;
	goto L49
L49:
	;
	v166 = F_JsonbValueAsText(m, v12+int32(40))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v179 = v141
	v181 = v166
	goto L32
L51:
	;
	v179 = v141
	v181 = v170
	goto L32
L52:
	;
	v195 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v195)
	m.G0 = v12 + int32(80)
	return
}
func F_ec_search_derived_clause_for_ems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	var v39 float64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v54 int64
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v76 int64
	_ = v76
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v16 != 0 {
		v188 = v16
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L9
	} else {
		goto L114
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v513
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l4
	v195 = base.B2i32(base.Ui32(l2) < base.Ui32(l3))
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L53
	} else {
		goto L54
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v17 == int32(0) {
		v513 = v6
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if int32(32) <= v20 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	v25 = F_MemoryContextAllocZero(m, v23, int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v137 == int32(0) {
		v513 = v6
		goto L2
	} else {
		goto L36
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v23
	v33 = float64(4.294967296e+09)
	v36 = base.F64_div(base.F64_convert_i32_u(v20), float64(0.9))
	if base.F64_ge(v36, v33) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = v33
	goto L13
L12:
	;
	v39 = v36
	goto L13
L13:
	;
	v40 = base.I64_trunc_sat_f64_u(v39)
	if base.Ui64(v40) <= base.Ui64(int64(2)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v43 = int64(2)
	goto L16
L15:
	;
	v43 = v40
	goto L16
L16:
	;
	v44 = int64(1)
	if v43&(v43-v44) == int64(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v54 = v43
	goto L19
L18:
	;
	v54 = v44 << (uint(int64(64)-base.I64_clz(v43)) % 64)
	goto L19
L19:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v54*int64(20)) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v63 = F_MemoryContextAllocExtended(m, v23, base.I32_wrap_i64(v54)*int32(20), int32(5))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v63
	v66 = int64(1)
	if v54&(v54-v66) == int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = v54
	goto L24
L23:
	;
	v76 = v66 << (uint(int64(64)-base.I64_clz(v54)) % 64)
	goto L24
L24:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v76*int64(20)) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = base.I32_wrap_i64(v76) - int32(1)
	if v76 == int64(4294967296) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v93 = int32(-85899346)
	goto L28
L27:
	;
	v93 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v76), float64(0.9)))
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v25
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v96 == int32(0) {
		v188 = v25
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 <= int32(0) {
		v188 = v25
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v103 = int32(0)
	goto L31
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v103<<(uint(int32(2))%32))))
	F_ec_add_clause_to_derives_hash(m, l1, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L33
	}
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v125 != 0 {
		v188 = v125
		goto L3
	} else {
		goto L35
	}
L33:
	;
	v122 = v103 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v122 < v123 {
		v103 = v122
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L8
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v140 <= int32(0) {
		v513 = v6
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v143 = int32(0)
	if v143 < v140 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v146 = v140
	goto L40
L39:
	;
	v146 = v143
	goto L40
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v150 = int32(0)
	goto L41
L41:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v147+v150<<(uint(int32(2))%32))))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	if base.B2i32(l3 == int32(0))&base.B2i32(l2 == v166) != 0 {
		v513 = v165
		goto L2
	} else {
		goto L43
	}
L42:
	;
	v513 = int32(0)
	goto L2
L43:
	;
	if v166 != l2 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v166 != l3 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	if v170 != l3 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v165)+60))
	if v172 == l4 {
		v513 = v165
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v181 = v150 + int32(1)
	if v181 != v146 {
		v150 = v181
		goto L41
	} else {
		goto L52
	}
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	if v175 != l2 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v165)+60))
	if v177 == l4 {
		v513 = v165
		goto L2
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	goto L42
L53:
	;
	v196 = l3
	goto L55
L54:
	;
	v196 = l2
	goto L55
L55:
	;
	if l3 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v197 = v196
	goto L58
L57:
	;
	v197 = l2
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v197
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v199 = l2
	goto L61
L60:
	;
	v199 = l3
	goto L61
L61:
	;
	if l3 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v201 = v199
	goto L64
L63:
	;
	v201 = int32(0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v201
	v203 = int32(12)
	v209 = int32(-1636608420)
	if v14&int32(3) != 0 {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v470 = (v463 ^ v455 - base.I32_rotl(v463, int32(24))) & v469
	v473 = v468 + v470*int32(20)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	if v474 == int32(0) {
		v513 = v6
		goto L2
	} else {
		goto L105
	}
L66:
	;
	v441 = int32(14)
	v443 = v437 ^ v438 - base.I32_rotl(v437, v441)
	v447 = v443 ^ v436 - base.I32_rotl(v443, int32(11))
	v451 = v447 ^ v437 - base.I32_rotl(v447, int32(25))
	v455 = v451 ^ v443 - base.I32_rotl(v451, int32(16))
	v459 = v455 ^ v447 - base.I32_rotl(v455, int32(4))
	v463 = v459 ^ v451 - base.I32_rotl(v459, v441)
	goto L65
L67:
	;
	switch v363 - int32(1) {
	case 0:
		v429 = v354
		v430 = v355
		v431 = v359
		goto L94
	case 1:
		v422 = v354
		v423 = v355
		v424 = v359
		goto L95
	case 2:
		v415 = v354
		v416 = v355
		v417 = v359
		goto L96
	case 3:
		v409 = v355
		v410 = v359
		goto L97
	case 4:
		v405 = v355
		v406 = v359
		goto L98
	case 5:
		v399 = v355
		v400 = v359
		goto L99
	case 6:
		v393 = v355
		v394 = v359
		goto L100
	case 7:
		v388 = v359
		goto L101
	case 8:
		v383 = v359
		goto L102
	case 9:
		v378 = v359
		goto L103
	case 10:
		goto L104
	default:
		v436 = v354
		v437 = v355
		v438 = v359
		goto L66
	}
L68:
	;
	v318 = v14
	v319 = v203
	v320 = v209
	v321 = v209
	v322 = v209
	goto L91
L69:
	;
	goto L68
L70:
	;
	goto L71
L71:
	;
	goto L75
L73:
	;
	switch v261 - int32(1) {
	case 0:
		v315 = v252
		goto L80
	case 1:
		v310 = v252
		goto L81
	case 2:
		goto L82
	case 3:
		v303 = v253
		goto L83
	case 4:
		v300 = v253
		goto L84
	case 5:
		v295 = v253
		goto L85
	case 6:
		goto L86
	case 7:
		v286 = v257
		goto L87
	case 8:
		v281 = v257
		goto L88
	case 9:
		v276 = v257
		goto L89
	case 10:
		goto L90
	default:
		v436 = v252
		v437 = v253
		v438 = v257
		goto L66
	}
L75:
	;
	goto L76
L76:
	;
	v216 = v14
	v217 = v203
	v218 = v209
	v219 = v209
	v220 = v209
	goto L77
L77:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v223 = v222 + v219
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v227 = v226 + v220
	v229 = int32(4)
	v231 = v224 + v218 - v227 ^ base.I32_rotl(v227, v229)
	v235 = v223 - v231 ^ base.I32_rotl(v231, int32(6))
	v236 = v227 + v223
	v237 = v231 + v236
	v238 = v235 + v237
	v242 = v236 - v235 ^ base.I32_rotl(v235, int32(8))
	v246 = v237 - v242 ^ base.I32_rotl(v242, int32(16))
	v250 = v238 - v246 ^ base.I32_rotl(v246, int32(19))
	v251 = v242 + v238
	v252 = v246 + v251
	v253 = v250 + v252
	v257 = v251 - v250 ^ base.I32_rotl(v250, v229)
	v258 = int32(12)
	v259 = v216 + v258
	v261 = v217 - v258
	if base.Ui32(int32(11)) < base.Ui32(v261) {
		v216 = v259
		v217 = v261
		v218 = v252
		v219 = v253
		v220 = v257
		goto L77
	} else {
		goto L79
	}
L78:
	;
	goto L73
L79:
	;
	goto L78
L80:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	v436 = v315 + v316
	v437 = v253
	v438 = v257
	goto L66
L81:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v315 = v311<<(uint(int32(8))%32) + v310
	goto L80
L82:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+2)))
	v310 = v306<<(uint(int32(16))%32) + v252
	goto L81
L83:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v436 = v304 + v252
	v437 = v303
	v438 = v257
	goto L66
L84:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+4)))
	v303 = v300 + v301
	goto L83
L85:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+5)))
	v300 = v296<<(uint(int32(8))%32) + v295
	goto L84
L86:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+6)))
	v295 = v291<<(uint(int32(16))%32) + v253
	goto L85
L87:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	v436 = v287 + v252
	v437 = v289 + v253
	v438 = v286
	goto L66
L88:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+8)))
	v286 = v282<<(uint(int32(8))%32) + v281
	goto L87
L89:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+9)))
	v281 = v277<<(uint(int32(16))%32) + v276
	goto L88
L90:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+10)))
	v276 = v272<<(uint(int32(24))%32) + v257
	goto L89
L91:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v325 = v324 + v321
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v329 = v328 + v322
	v331 = int32(4)
	v333 = v326 + v320 - v329 ^ base.I32_rotl(v329, v331)
	v337 = v325 - v333 ^ base.I32_rotl(v333, int32(6))
	v338 = v329 + v325
	v339 = v333 + v338
	v340 = v337 + v339
	v344 = v338 - v337 ^ base.I32_rotl(v337, int32(8))
	v348 = v339 - v344 ^ base.I32_rotl(v344, int32(16))
	v352 = v340 - v348 ^ base.I32_rotl(v348, int32(19))
	v353 = v344 + v340
	v354 = v348 + v353
	v355 = v352 + v354
	v359 = v353 - v352 ^ base.I32_rotl(v352, v331)
	v360 = int32(12)
	v361 = v318 + v360
	v363 = v319 - v360
	if base.Ui32(int32(11)) < base.Ui32(v363) {
		v318 = v361
		v319 = v363
		v320 = v354
		v321 = v355
		v322 = v359
		goto L91
	} else {
		goto L93
	}
L92:
	;
	goto L67
L93:
	;
	goto L92
L94:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	v436 = v429 + v432
	v437 = v430
	v438 = v431
	goto L66
L95:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+1)))
	v429 = v425<<(uint(int32(8))%32) + v422
	v430 = v423
	v431 = v424
	goto L94
L96:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+2)))
	v422 = v418<<(uint(int32(16))%32) + v415
	v423 = v416
	v424 = v417
	goto L95
L97:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+3)))
	v415 = v411<<(uint(int32(24))%32) + v354
	v416 = v409
	v417 = v410
	goto L96
L98:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+4)))
	v409 = v405 + v407
	v410 = v406
	goto L97
L99:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+5)))
	v405 = v401<<(uint(int32(8))%32) + v399
	v406 = v400
	goto L98
L100:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+6)))
	v399 = v395<<(uint(int32(16))%32) + v393
	v400 = v394
	goto L99
L101:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+7)))
	v393 = v389<<(uint(int32(24))%32) + v355
	v394 = v388
	goto L100
L102:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+8)))
	v388 = v384<<(uint(int32(8))%32) + v383
	goto L101
L103:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+9)))
	v383 = v379<<(uint(int32(16))%32) + v378
	goto L102
L104:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+10)))
	v378 = v374<<(uint(int32(24))%32) + v359
	goto L103
L105:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v482 = v470
	v483 = v473
	goto L106
L106:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	if v491 != v479 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	v513 = v504
	goto L2
L108:
	;
	goto L107
L109:
	;
	v499 = (v482 + int32(1)) & v469
	v502 = v468 + v499*int32(20)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v503 != 0 {
		v482 = v499
		v483 = v502
		goto L106
	} else {
		goto L113
	}
L110:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	if v493 != v478 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v483)+12))
	if v495 == v477 {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	v513 = v6
	goto L2
L114:
	;
	F_errmsg_internal(m, int32(_a_F_ec_search_derived_clause_for_ems_0), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L9
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_ec_search_derived_clause_for_ems_1), int32(332), int32(_a_F_ec_search_derived_clause_for_ems_2))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L9
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_element_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_element_hash[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_FunctionCall1Coll(m, v5, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v8)
	}
}
func F_elements_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v4 == int32(1) {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v7 != int32(1) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v17
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
			if v10 != int32(1) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v17
				return int32(0)
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v13)
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_elements_worker_jsonb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_MemoryContextDelete(m, v35)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L54
	}
L2:
	;
	v110 = v106
	v112 = v104
	v114 = v105
	goto L33
L3:
	;
	v104 = v3
	v105 = v3
	v106 = int32(2)
	goto L2
L4:
	;
	v104 = v101
	v105 = v50
	v106 = int32(1)
	goto L2
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L29
	}
L6:
	;
	return
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v16&int32(268435456) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v16&int32(1073741824) == int32(0) {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L25
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0]))
	v35 = F_AllocSetContextCreateInternal(m, v30, int32(_a_F_elements_worker_jsonb_0), int32(0), int32(_a_F_elements_worker_jsonb_1), int32(_a_F_elements_worker_jsonb_2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v39 = F_JsonbIteratorInit(m, v14+int32(4))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v39
	v47 = F_JsonbIteratorNext(m, v9+int32(-4), v9+int32(-40), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L16
	}
L15:
	;
	v49 = int32(_a_F_elements_worker_jsonb_3)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0])) = v35
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v53)
	if l1 == v53 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	switch v47 {
	case 0:
		goto L1
	default:
		goto L3
	case 3:
		goto L15
	}
L17:
	;
	v59 = F_JsonbValueToJsonb(m, v9+int32(-40))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v61 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v101 = v59
	goto L4
L21:
	;
	v64 = F_JsonbValueAsText(m, v9+int32(-40))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v66)
	v104 = v3
	v105 = v50
	v106 = int32(0)
	goto L2
L24:
	;
	v101 = v64
	goto L4
L25:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_elements_worker_jsonb_4), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_elements_worker_jsonb_5), int32(2235), int32(_a_F_elements_worker_jsonb_6))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_elements_worker_jsonb_7), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_elements_worker_jsonb_5), int32(2239), int32(_a_F_elements_worker_jsonb_6))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	switch v110 {
	case 0:
		goto L39
	case 1:
		goto L38
	default:
		goto L37
	}
L35:
	;
	v110 = int32(0)
	v114 = v147
	goto L33
L36:
	;
	v110 = int32(1)
	v112 = v173
	v114 = v172
	goto L33
L37:
	;
	goto L42
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = base.I64_extend_i32_u(v112)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v118, v119, v9+int32(-48), v9+int32(-49))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L40
	}
L39:
	;
	v172 = v114
	v173 = int32(0)
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0])) = v114
	F_MemoryContextReset(m, v35)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v110 = int32(2)
	goto L33
L42:
	;
	v144 = F_JsonbIteratorNext(m, v9+int32(-4), v9+int32(-40), int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	v146 = int32(_a_F_elements_worker_jsonb_3)
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_elements_worker_jsonb[0])) = v35
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v150)
	if l1 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L43
L45:
	;
	switch v144 {
	case 0:
		goto L1
	default:
		goto L42
	case 3:
		goto L44
	}
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v152 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v163 = F_JsonbValueToJsonb(m, v9+int32(-40))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L53
	}
L49:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v155)
	goto L35
L50:
	;
	goto L51
L51:
	;
	v159 = F_JsonbValueAsText(m, v9+int32(-40))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	v172 = v147
	v173 = v159
	goto L36
L53:
	;
	v172 = v147
	v173 = v163
	goto L36
L54:
	;
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v186)
	m.G0 = v11 - int32(-64)
	return
}
func F_encrypt_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	F_init_work(m, v9+int32(-12), l1, l4, v9+int32(-56))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(0)
	if l1 == v21 {
		v81 = l2
		v82 = v21
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v85 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	if v25 == int32(0) {
		v81 = l2
		v82 = v21
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_encrypt_internal[1]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	goto L6
L6:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v31 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v61 = int32(1)
	if v31&v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v37 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v48 = int32(1)
	if v31&v48 != 0 {
		v60 = int32(base.Ui32(v31)>>(uint(v48)%32)) - v48
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v40 = int32(16)
	goto L13
L12:
	;
	v40 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = int32(4)
	goto L16
L15:
	;
	v47 = v40
	goto L16
L16:
	;
	v60 = v47
	goto L7
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v65 = v61
	goto L20
L19:
	;
	v65 = int32(4)
	goto L20
L20:
	;
	v66 = v65 + l2
	v68 = F_pg_do_encoding_conversion(m, v66, v60, v30, int32(6))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v68 != v66 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = F_cstring_to_text(m, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v75 = l2
	goto L24
L24:
	;
	v76 = base.B2i32(l2 == v75)
	if l2 == v75 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	F_pfree(m, v68)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v75 = v71
	goto L24
L27:
	;
	v77 = l2
	goto L29
L28:
	;
	v77 = v75
	goto L29
L29:
	;
	if l2 == v75 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v79 = int32(0)
	goto L32
L31:
	;
	v79 = v75
	goto L32
L32:
	;
	v81 = v77
	v82 = v79
	goto L3
L33:
	;
	v115 = int32(1)
	if v85&v115 != 0 {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v91 == int32(18) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v102 = int32(1)
	if v85&v102 != 0 {
		v114 = int32(base.Ui32(v85)>>(uint(v102)%32)) - v102
		goto L33
	} else {
		goto L43
	}
L37:
	;
	v94 = int32(16)
	goto L39
L38:
	;
	v94 = int32(0)
	goto L39
L39:
	;
	if base.Ui32((v91-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v101 = int32(4)
	goto L42
L41:
	;
	v101 = v94
	goto L42
L42:
	;
	v114 = v101
	goto L33
L43:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v114 = int32(base.Ui32(v108)>>(uint(int32(2))%32)) - int32(4)
	goto L33
L44:
	;
	v119 = v115
	goto L46
L45:
	;
	v119 = int32(4)
	goto L46
L46:
	;
	v121 = F_mbuf_create_from_data(m, v81+v119, v114)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v123 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = F_mbuf_create(m, v148+int32(128))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L59
	}
L49:
	;
	v127 = int32(18)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v129 == v127 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v140 = int32(1)
	if v123&v140 != 0 {
		v148 = int32(base.Ui32(v123) >> (uint(v140) % 32))
		goto L48
	} else {
		goto L58
	}
L52:
	;
	v132 = v127
	goto L54
L53:
	;
	v132 = int32(2)
	goto L54
L54:
	;
	if base.Ui32((v129-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v139 = int32(6)
	goto L57
L56:
	;
	v139 = v132
	goto L57
L57:
	;
	v148 = v139
	goto L48
L58:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v148 = int32(base.Ui32(v144) >> (uint(int32(2)) % 32))
	goto L48
L59:
	;
	v156 = F_mbuf_append(m, v151, v9+int32(-4), int32(4))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if l0 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if int32(0) <= v253 {
		goto L101
	} else {
		goto L102
	}
L62:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v158 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	goto L64
L64:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v205 = int32(1)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v209 = v207 & v205
	if v209 != 0 {
		goto L82
	} else {
		goto L83
	}
L65:
	;
	v188 = int32(1)
	if v158&v188 != 0 {
		goto L76
	} else {
		goto L77
	}
L66:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v164 == int32(18) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v175 = int32(1)
	if v158&v175 != 0 {
		v187 = int32(base.Ui32(v158)>>(uint(v175)%32)) - v175
		goto L65
	} else {
		goto L75
	}
L69:
	;
	v167 = int32(16)
	goto L71
L70:
	;
	v167 = int32(0)
	goto L71
L71:
	;
	if base.Ui32((v164-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v174 = int32(4)
	goto L74
L73:
	;
	v174 = v167
	goto L74
L74:
	;
	v187 = v174
	goto L65
L75:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v187 = int32(base.Ui32(v181)>>(uint(int32(2))%32)) - int32(4)
	goto L65
L76:
	;
	v192 = v188
	goto L78
L77:
	;
	v192 = int32(4)
	goto L78
L78:
	;
	v194 = F_mbuf_create_from_data(m, l3+v192, v187)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v197 = int32(0)
	v200 = F_pgp_set_pubkey(m, v196, v194, v197, v197, v197)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v202 = F_mbuf_free(m, v194)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v253 = v200
	goto L61
L82:
	;
	v210 = v205
	goto L84
L83:
	;
	v210 = int32(4)
	goto L84
L84:
	;
	v211 = l3 + v210
	if v207 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v239 = int32(0)
	if base.B2i32(v211 == v239)|base.B2i32(v238 <= v239) != 0 {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v217 == int32(18) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v228 = int32(1)
	if v209 != 0 {
		v238 = int32(base.Ui32(v207)>>(uint(v228)%32)) - v228
		goto L85
	} else {
		goto L95
	}
L89:
	;
	v220 = int32(16)
	goto L91
L90:
	;
	v220 = int32(0)
	goto L91
L91:
	;
	if base.Ui32((v217-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v227 = int32(4)
	goto L94
L93:
	;
	v227 = v220
	goto L94
L94:
	;
	v238 = v227
	goto L85
L95:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v238 = int32(base.Ui32(v232)>>(uint(int32(2))%32)) - int32(4)
	goto L85
L96:
	;
	v253 = v250
	goto L61
L97:
	;
	v250 = int32(-13)
	goto L99
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+132)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v204)+128)) = v211
	v250 = int32(0)
	goto L99
L99:
	;
	goto L96
L100:
	;
	v309 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+16)) = uint16(v309)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(-8)))) = v312
	v314 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v151)+8)) = v314
	*(*int64)(unsafe.Add(mBase, uint32(v151))) = v314
	goto L133
L101:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v257 = F_pgp_encrypt(m, v256, v121, v151)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	v261 = v253
	goto L103
L103:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v262 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	if v257 == int32(0) {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v261 = v257
	goto L103
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_encrypt_internal[0])) = int32(0)
	goto L109
L107:
	;
	goto L108
L108:
	;
	if v82 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v267 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	goto L112
L112:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v298 = F_pgp_free(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L129
	}
L113:
	;
	if v292 != 0 {
		goto L125
	} else {
		goto L126
	}
L114:
	;
	v271 = int32(18)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v273 == v271 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	v284 = int32(1)
	if v267&v284 != 0 {
		v292 = int32(base.Ui32(v267) >> (uint(v284) % 32))
		goto L113
	} else {
		goto L123
	}
L117:
	;
	v276 = v271
	goto L119
L118:
	;
	v276 = int32(2)
	goto L119
L119:
	;
	if base.Ui32((v273-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v283 = int32(6)
	goto L122
L121:
	;
	v283 = v276
	goto L122
L122:
	;
	v292 = v283
	goto L113
L123:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v292 = int32(base.Ui32(v288) >> (uint(int32(2)) % 32))
	goto L113
L124:
	;
	F_pfree(m, v82)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L128
	}
L125:
	;
	base.MemoryFill(m, v82, int32(0), v292)
	goto L127
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	goto L112
L129:
	;
	v300 = F_mbuf_free(m, v121)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v302 = F_mbuf_free(m, v151)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_px_THROW_ERROR(m, v261)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = (v311 - v312) << (uint(int32(2)) % 32)
	if v82 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v324 == int32(1) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	goto L136
L136:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v355 = F_pgp_free(m, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L153
	}
L137:
	;
	if v349 != 0 {
		goto L149
	} else {
		goto L150
	}
L138:
	;
	v328 = int32(18)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v330 == v328 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	v341 = int32(1)
	if v324&v341 != 0 {
		v349 = int32(base.Ui32(v324) >> (uint(v341) % 32))
		goto L137
	} else {
		goto L147
	}
L141:
	;
	v333 = v328
	goto L143
L142:
	;
	v333 = int32(2)
	goto L143
L143:
	;
	if base.Ui32((v330-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v340 = int32(6)
	goto L146
L145:
	;
	v340 = v333
	goto L146
L146:
	;
	v349 = v340
	goto L137
L147:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v349 = int32(base.Ui32(v345) >> (uint(int32(2)) % 32))
	goto L137
L148:
	;
	F_pfree(m, v82)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L152
	}
L149:
	;
	base.MemoryFill(m, v82, int32(0), v349)
	goto L151
L150:
	;
	goto L151
L151:
	;
	goto L148
L152:
	;
	goto L136
L153:
	;
	v357 = F_mbuf_free(m, v121)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v359 = F_mbuf_free(m, v151)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_encrypt_internal[0])) = int32(0)
	goto L156
L156:
	;
	m.G0 = v11 - int32(-64)
	return v319
}
func F_encrypt_password(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int64
	_ = v916
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v4
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v19 != int32(109) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v15 + int32(112)
	return v881
L2:
	;
	v1054 = F_parse_scram_secret(m, v881, v15+int32(104), v15+int32(96), v15+int32(100), v15+int32(108), v15-int32(-64), v15+int32(32))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L33
	} else {
		goto L261
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L33
	} else {
		goto L256
	}
L4:
	;
	v890 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[0])))
	if v890 != int32(1) {
		goto L1
	} else {
		goto L224
	}
L5:
	;
	if v864 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(0)
	switch l0 {
	case 0:
		goto L38
	case 1:
		goto L39
	case 2:
		goto L37
	default:
		v881 = v4
		goto L4
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(0)
	v139 = F_pstrdup(m, l2)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L33
	} else {
		goto L36
	}
L8:
	;
	v131 = F_parse_scram_secret(m, l2, v15+int32(104), v15+int32(96), v15+int32(100), v15+int32(108), v15-int32(-64), v15+int32(32))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v22 != int32(100) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
	if v25 != int32(53) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v28 = F_strlen(m, l2)
	mBase = m.M
	if v28 != int32(35) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v32 = l2 + int32(3)
	v33 = int32(_a_F_encrypt_password_0)
	v37 = m.G0
	v39 = v37 - int32(32)
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v40
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[1])))
	if v48 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v116 == int32(32) {
		goto L7
	} else {
		goto L32
	}
L14:
	;
	v116 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[2])))
	if v52 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v56 = v32
	goto L20
L18:
	;
	goto L19
L19:
	;
	v66 = v33
	v67 = v48
	goto L23
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v62 == v48 {
		v56 = v56 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v116 = v56 - v32
	goto L13
L22:
	;
	goto L21
L23:
	;
	v74 = v39 + int32(base.Ui32(v67)>>(uint(int32(3))%32))&int32(28)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v76 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75 | v76<<(uint(v67)%32)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v80 != 0 {
		v66 = v66 + v76
		v67 = v80
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v83 == int32(0) {
		v106 = v32
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v116 = v106 - v32
	goto L13
L27:
	;
	v87 = v32
	v88 = v83
	goto L28
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(base.Ui32(v88)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v96)>>(uint(v88)%32))&int32(1) == int32(0) {
		v106 = v87
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v106 = v104
	goto L26
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v104 = v87 + int32(1)
	if v102 != 0 {
		v87 = v104
		v88 = v102
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L8
L33:
	;
	return int32(0)
L34:
	;
	if v131 == int32(0) {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L7
L36:
	;
	v864 = v139
	goto L5
L37:
	;
	v180 = m.G0
	v182 = v180 - int32(48)
	m.G0 = v182
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = int32(0)
	v188 = F_pg_saslprep(m, l2, v182+int32(44))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L33
	} else {
		goto L50
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L33
	} else {
		goto L46
	}
L39:
	;
	v144 = F_palloc(m, int32(36))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v146 = F_strlen(m, l1)
	mBase = m.M
	v149 = F_pg_md5_encrypt(m, l2, l1, v146, v144, v15+int32(28))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	if v149 != 0 {
		v864 = v144
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v155
	F_errmsg_internal(m, int32(_a_F_encrypt_password_1), v15+int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_2), int32(204), int32(_a_F_encrypt_password_3))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L33
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errmsg_internal(m, int32(_a_F_encrypt_password_4), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L33
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_2), int32(212), int32(_a_F_encrypt_password_3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L33
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v864 = v397
	goto L5
L50:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)+44))
	v191 = int32(16)
	v192 = v182 + v191
	v194 = int32(0)
	v198 = m.G0
	v200 = v198 - v191
	m.G0 = v200
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v194
	v206 = F_open(m, int32(_a_F_encrypt_password_5), v194, v200)
	mBase = m.M
	if v206 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v239 != 0 {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	goto L56
L53:
	;
	v239 = v194
	goto L54
L54:
	;
	m.G0 = v200 + int32(16)
	goto L51
L55:
	;
	v234 = F_close(m, v206)
	mBase = m.M
	v239 = v232
	goto L54
L56:
	;
	v212 = v192
	v213 = v191
	goto L57
L57:
	;
	v218 = F_read(m, v206, v212, v213)
	mBase = m.M
	if v218 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v232 = int32(1)
	goto L55
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_encrypt_password[3]))
	if v222 == int32(27) {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v227 = v213 - v218
	if v227 != 0 {
		v212 = v212 + v218
		v213 = v227
		goto L57
	} else {
		goto L63
	}
L62:
	;
	v232 = int32(0)
	goto L55
L63:
	;
	goto L58
L64:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_encrypt_password[4]))
	v246 = m.G0
	v248 = v246 - int32(176)
	m.G0 = v248
	if v188 != 0 {
		goto L72
	} else {
		goto L73
	}
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L33
	} else {
		goto L216
	}
L67:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v182)+44))
	if v839 != 0 {
		goto L212
	} else {
		goto L213
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_encrypt_password_6)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L33
	} else {
		goto L209
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_encrypt_password_7)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L33
	} else {
		goto L206
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_encrypt_password_8)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L33
	} else {
		goto L203
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L33
	} else {
		goto L200
	}
L72:
	;
	v250 = l2
	goto L74
L73:
	;
	v250 = v190
	goto L74
L74:
	;
	v257 = v182 + int32(12)
	v258 = F_scram_SaltedPassword(m, v250, int32(3), int32(32), v192, int32(16), v245, v248+int32(144), v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L33
	} else {
		goto L75
	}
L75:
	;
	if v258 < int32(0) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v263 = F_pg_hmac_create(m, int32(3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L33
	} else {
		goto L78
	}
L77:
	;
	if v351 < int32(0) {
		goto L71
	} else {
		goto L124
	}
L78:
	;
	if v263 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	goto L83
L80:
	;
	goto L81
L81:
	;
	v295 = F_pg_hmac_init(m, v263, v248+int32(144), int32(32))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L33
	} else {
		goto L97
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(_a_F_encrypt_password_9)
	v351 = int32(-1)
	goto L77
L83:
	;
	goto L82
L95:
	;
	F_pg_hmac_free(m, v263)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L33
	} else {
		goto L123
	}
L96:
	;
	if v263 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L97:
	;
	if v295 < int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	if v263 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v315 < int32(0) {
		goto L96
	} else {
		goto L106
	}
L100:
	;
	v315 = int32(-1)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v305 = F_pg_cryptohash_update(m, v304, int32(_a_F_encrypt_password_10), int32(10))
	mBase = m.M
	if int32(0) <= v305 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v315 = int32(0)
	goto L99
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+8)) = int32(2)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v312 = F_pg_cryptohash_error(m, v311)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v263)+12)) = v312
	v315 = int32(-1)
	goto L99
L106:
	;
	v319 = F_pg_hmac_final(m, v263, v248+int32(112), int32(32))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L33
	} else {
		goto L107
	}
L107:
	;
	if int32(0) <= v319 {
		goto L95
	} else {
		goto L108
	}
L108:
	;
	goto L96
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v342
	F_pg_hmac_free(m, v263)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L33
	} else {
		goto L122
	}
L110:
	;
	v342 = int32(_a_F_encrypt_password_9)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	if v327 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v339 = v327
	goto L115
L114:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	if v331 == int32(2) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v342 = v339
	goto L109
L116:
	;
	v334 = int32(_a_F_encrypt_password_11)
	goto L118
L117:
	;
	v334 = int32(_a_F_encrypt_password_12)
	goto L118
L118:
	;
	if v331 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v337 = int32(_a_F_encrypt_password_9)
	goto L121
L120:
	;
	v337 = v334
	goto L121
L121:
	;
	v339 = v337
	goto L115
L122:
	;
	v351 = int32(-1)
	goto L77
L123:
	;
	v351 = int32(0)
	goto L77
L124:
	;
	v355 = v248 + int32(112)
	v358 = F_scram_H(m, v355, int32(3), int32(32), v355, v257)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L33
	} else {
		goto L125
	}
L125:
	;
	if v358 < int32(0) {
		goto L71
	} else {
		goto L126
	}
L126:
	;
	v367 = v248 + int32(80)
	v368 = F_scram_ServerKey(m, v248+int32(144), int32(3), int32(32), v367, v257)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L33
	} else {
		goto L127
	}
L127:
	;
	if v368 < int32(0) {
		goto L71
	} else {
		goto L128
	}
L128:
	;
	v376 = base.I32_div_s(int32(18), int32(3))
	v378 = v376 << (uint(int32(2)) % 32)
	goto L129
L129:
	;
	v383 = base.I32_div_s(int32(34), int32(3))
	v385 = v383 << (uint(int32(2)) % 32)
	goto L130
L130:
	;
	v391 = base.I32_div_s(int32(34), int32(3))
	v393 = v391 << (uint(int32(2)) % 32)
	goto L131
L131:
	;
	v397 = F_palloc(m, v378+v385+v393+int32(28))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L33
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+64)) = v245
	v404 = F_pg_sprintf(m, v397, int32(_a_F_encrypt_password_13), v248-int32(-64))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L33
	} else {
		goto L133
	}
L133:
	;
	v406 = v404 + v397
	goto L137
L134:
	;
	if v518 < int32(0) {
		goto L70
	} else {
		goto L155
	}
L135:
	;
	if v378 != 0 {
		goto L152
	} else {
		goto L153
	}
L136:
	;
	if v378 < v460-v406+int32(4) {
		goto L135
	} else {
		goto L148
	}
L137:
	;
	v415 = v192
	v416 = int32(0)
	v419 = v406
	v420 = int32(2)
	goto L140
L139:
	;
	v518 = v460 - v406
	goto L134
L140:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	v426 = v422<<(uint(v420<<(uint(int32(3))%32))%32) | v416
	if int32(0) < v420 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	if v461 != int32(2) {
		goto L136
	} else {
		goto L147
	}
L142:
	;
	v459 = v426
	v460 = v419
	v461 = v420 - int32(1)
	goto L144
L143:
	;
	if v378 < v419-v406+int32(4) {
		goto L135
	} else {
		goto L145
	}
L144:
	;
	v463 = v415 + int32(1)
	if base.Ui32(v463) < base.Ui32(v182+int32(32)) {
		v415 = v463
		v416 = v459
		v419 = v460
		v420 = v461
		goto L140
	} else {
		goto L146
	}
L145:
	;
	v435 = int32(63)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426&v435)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+3)) = uint8(v437)
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v426)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v419))) = uint8(v441)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v426)>>(uint(int32(6))%32))&v435)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+2)) = uint8(v447)
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v426)>>(uint(int32(12))%32))&v435)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)) = uint8(v453)
	v459 = int32(0)
	v460 = v419 + int32(4)
	v461 = int32(2)
	goto L144
L146:
	;
	goto L141
L147:
	;
	goto L139
L148:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v459)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v481)
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v459)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)) = uint8(v487)
	if v461 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v459)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	v497 = v496
	goto L151
L150:
	;
	v497 = int32(61)
	goto L151
L151:
	;
	v498 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v460)+3)) = uint8(v498)
	*(*uint8)(unsafe.Add(mBase, uint32(v460)+2)) = uint8(v497)
	v518 = v460 + int32(4) - v406
	goto L134
L152:
	;
	base.MemoryFill(m, v406, int32(0), v378)
	goto L154
L153:
	;
	goto L154
L154:
	;
	v518 = int32(-1)
	goto L134
L155:
	;
	v521 = v406 + v518
	v522 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v521))) = uint8(v522)
	v526 = v521 + int32(1)
	goto L159
L156:
	;
	if v638 < int32(0) {
		goto L69
	} else {
		goto L177
	}
L157:
	;
	if v385 != 0 {
		goto L174
	} else {
		goto L175
	}
L158:
	;
	if v385 < v580-v526+int32(4) {
		goto L157
	} else {
		goto L170
	}
L159:
	;
	v535 = v355
	v536 = int32(0)
	v539 = v526
	v540 = int32(2)
	goto L162
L161:
	;
	v638 = v580 - v526
	goto L156
L162:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	v546 = v542<<(uint(v540<<(uint(int32(3))%32))%32) | v536
	if int32(0) < v540 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v581 != int32(2) {
		goto L158
	} else {
		goto L169
	}
L164:
	;
	v579 = v546
	v580 = v539
	v581 = v540 - int32(1)
	goto L166
L165:
	;
	if v385 < v539-v526+int32(4) {
		goto L157
	} else {
		goto L167
	}
L166:
	;
	v583 = v535 + int32(1)
	if base.Ui32(v583) < base.Ui32(v248+int32(144)) {
		v535 = v583
		v536 = v579
		v539 = v580
		v540 = v581
		goto L162
	} else {
		goto L168
	}
L167:
	;
	v555 = int32(63)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546&v555)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+3)) = uint8(v557)
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v546)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v539))) = uint8(v561)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v546)>>(uint(int32(6))%32))&v555)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+2)) = uint8(v567)
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v546)>>(uint(int32(12))%32))&v555)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+1)) = uint8(v573)
	v579 = int32(0)
	v580 = v539 + int32(4)
	v581 = int32(2)
	goto L166
L168:
	;
	goto L163
L169:
	;
	goto L161
L170:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v579)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v580))) = uint8(v601)
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v579)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+1)) = uint8(v607)
	if v581 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v579)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	v617 = v616
	goto L173
L172:
	;
	v617 = int32(61)
	goto L173
L173:
	;
	v618 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+3)) = uint8(v618)
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+2)) = uint8(v617)
	v638 = v580 + int32(4) - v526
	goto L156
L174:
	;
	base.MemoryFill(m, v526, int32(0), v385)
	goto L176
L175:
	;
	goto L176
L176:
	;
	v638 = int32(-1)
	goto L156
L177:
	;
	v641 = v526 + v638
	v642 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v641))) = uint8(v642)
	v646 = v641 + int32(1)
	goto L181
L178:
	;
	if v758 < int32(0) {
		goto L68
	} else {
		goto L199
	}
L179:
	;
	if v393 != 0 {
		goto L196
	} else {
		goto L197
	}
L180:
	;
	if v393 < v700-v646+int32(4) {
		goto L179
	} else {
		goto L192
	}
L181:
	;
	v655 = v367
	v656 = int32(0)
	v659 = v646
	v660 = int32(2)
	goto L184
L183:
	;
	v758 = v700 - v646
	goto L178
L184:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655))))
	v666 = v662<<(uint(v660<<(uint(int32(3))%32))%32) | v656
	if int32(0) < v660 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	if v701 != int32(2) {
		goto L180
	} else {
		goto L191
	}
L186:
	;
	v699 = v666
	v700 = v659
	v701 = v660 - int32(1)
	goto L188
L187:
	;
	if v393 < v659-v646+int32(4) {
		goto L179
	} else {
		goto L189
	}
L188:
	;
	v703 = v655 + int32(1)
	if base.Ui32(v703) < base.Ui32(v248+int32(112)) {
		v655 = v703
		v656 = v699
		v659 = v700
		v660 = v701
		goto L184
	} else {
		goto L190
	}
L189:
	;
	v675 = int32(63)
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666&v675)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+3)) = uint8(v677)
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v666)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v659))) = uint8(v681)
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v666)>>(uint(int32(6))%32))&v675)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+2)) = uint8(v687)
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v666)>>(uint(int32(12))%32))&v675)+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+1)) = uint8(v693)
	v699 = int32(0)
	v700 = v659 + int32(4)
	v701 = int32(2)
	goto L188
L190:
	;
	goto L185
L191:
	;
	goto L183
L192:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v699)>>(uint(int32(18))%32)))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v700))) = uint8(v721)
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v699)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v700)+1)) = uint8(v727)
	if v701 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v699)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_encrypt_password[5]))))
	v737 = v736
	goto L195
L194:
	;
	v737 = int32(61)
	goto L195
L195:
	;
	v738 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v700)+3)) = uint8(v738)
	*(*uint8)(unsafe.Add(mBase, uint32(v700)+2)) = uint8(v737)
	v758 = v700 + int32(4) - v646
	goto L178
L196:
	;
	base.MemoryFill(m, v646, int32(0), v393)
	goto L198
L197:
	;
	goto L198
L198:
	;
	v758 = int32(-1)
	goto L178
L199:
	;
	v762 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v646+v758))) = uint8(v762)
	m.G0 = v248 + int32(176)
	goto L67
L200:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v775
	F_errmsg_internal(m, int32(_a_F_encrypt_password_14), v248)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L33
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_15), int32(245), int32(_a_F_encrypt_password_16))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L33
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = v791
	F_errmsg_internal(m, int32(_a_F_encrypt_password_17), v248+int32(16))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L33
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_15), int32(286), int32(_a_F_encrypt_password_16))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L33
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+32)) = v809
	F_errmsg_internal(m, int32(_a_F_encrypt_password_17), v248+int32(32))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L33
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_15), int32(302), int32(_a_F_encrypt_password_16))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L33
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+48)) = v827
	F_errmsg_internal(m, int32(_a_F_encrypt_password_17), v248+int32(48))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L33
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_15), int32(319), int32(_a_F_encrypt_password_16))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L33
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	F_pfree(m, v839)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L33
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	m.G0 = v182 + int32(48)
	goto L49
L215:
	;
	goto L214
L216:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L33
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(_a_F_encrypt_password_18), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L33
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_19), int32(502), int32(_a_F_encrypt_password_20))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L33
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	v881 = int32(0)
	goto L4
L221:
	;
	goto L222
L222:
	;
	v875 = F_strlen(m, v864)
	mBase = m.M
	if base.Ui32(int32(513)) <= base.Ui32(v875) {
		goto L3
	} else {
		goto L223
	}
L223:
	;
	v881 = v864
	goto L4
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = int32(0)
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881))))
	if v895 != int32(109) {
		goto L2
	} else {
		goto L225
	}
L225:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881)+1)))
	if v898 != int32(100) {
		goto L2
	} else {
		goto L226
	}
L226:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881)+2)))
	if v901 != int32(53) {
		goto L2
	} else {
		goto L227
	}
L227:
	;
	v904 = F_strlen(m, v881)
	mBase = m.M
	if v904 != int32(35) {
		goto L2
	} else {
		goto L228
	}
L228:
	;
	v908 = v881 + int32(3)
	v909 = int32(_a_F_encrypt_password_0)
	v913 = m.G0
	v915 = v913 - int32(32)
	v916 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v915)+24)) = v916
	*(*int64)(unsafe.Add(mBase, uint32(v915)+16)) = v916
	*(*int64)(unsafe.Add(mBase, uint32(v915)+8)) = v916
	*(*int64)(unsafe.Add(mBase, uint32(v915))) = v916
	v924 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[1])))
	if v924 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	if v992 != int32(32) {
		goto L2
	} else {
		goto L248
	}
L230:
	;
	v992 = int32(0)
	goto L229
L231:
	;
	goto L232
L232:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_encrypt_password[2])))
	if v928 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v932 = v908
	goto L236
L234:
	;
	goto L235
L235:
	;
	v942 = v909
	v943 = v924
	goto L239
L236:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932))))
	if v938 == v924 {
		v932 = v932 + int32(1)
		goto L236
	} else {
		goto L238
	}
L237:
	;
	v992 = v932 - v908
	goto L229
L238:
	;
	goto L237
L239:
	;
	v950 = v915 + int32(base.Ui32(v943)>>(uint(int32(3))%32))&int32(28)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)))
	v952 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v950))) = v951 | v952<<(uint(v943)%32)
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942)+1)))
	if v956 != 0 {
		v942 = v942 + v952
		v943 = v956
		goto L239
	} else {
		goto L241
	}
L240:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	if v959 == int32(0) {
		v982 = v908
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L240
L242:
	;
	v992 = v982 - v908
	goto L229
L243:
	;
	v963 = v908
	v964 = v959
	goto L244
L244:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v915+int32(base.Ui32(v964)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v972)>>(uint(v964)%32))&int32(1) == int32(0) {
		v982 = v963
		goto L242
	} else {
		goto L246
	}
L245:
	;
	v982 = v980
	goto L242
L246:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963)+1)))
	v980 = v963 + int32(1)
	if v978 != 0 {
		v963 = v980
		v964 = v978
		goto L244
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	v997 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L33
	} else {
		goto L249
	}
L249:
	;
	if v997 == int32(0) {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_errcode(m, int32(16908352))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L33
	} else {
		goto L251
	}
L251:
	;
	F_errmsg(m, int32(_a_F_encrypt_password_21), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L33
	} else {
		goto L252
	}
L252:
	;
	v1010 = F_errdetail(m, int32(_a_F_encrypt_password_22), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L33
	} else {
		goto L253
	}
L253:
	;
	F_errhint(m, int32(_a_F_encrypt_password_23), int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L33
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_2), int32(248), int32(_a_F_encrypt_password_3))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L33
	} else {
		goto L255
	}
L255:
	;
	goto L1
L256:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L33
	} else {
		goto L257
	}
L257:
	;
	F_errmsg(m, int32(_a_F_encrypt_password_24), int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L33
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(512)
	v1035 = F_errdetail(m, int32(_a_F_encrypt_password_25), v15)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L33
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_encrypt_password_2), int32(239), int32(_a_F_encrypt_password_3))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L33
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	goto L1
}
func F_encrypt_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	if l3 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v13 = l1 + int32(4)
	v16 = l2
	v17 = l3
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = int32(_a_F_encrypt_process_0)
	if base.Ui32(v22) <= base.Ui32(v17) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return v43
L6:
	;
	goto L5
L7:
	;
	v25 = v22
	goto L9
L8:
	;
	v25 = v17
	goto L9
L9:
	;
	v26 = F_pgp_cfb_encrypt(m, v21, v16, v25, v13)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v26 < int32(0) {
		v43 = v26
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v32 = F_pushf_write(m, l0, v13, v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v32 < int32(0) {
		v43 = v32
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(0)
	v38 = v17 - v25
	if v37 < v38 {
		v16 = v16 + v25
		v17 = v38
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v43 = v37
	goto L6
}
func F_end_MultiFuncCall(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	F_UnregisterExprContextCallback(m, v4, int32(1828), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
		F_MemoryContextDelete(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	}
}
func F_end_tup_output(m *base.Module, l0 int32) {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	m.T0[v4].(func(*base.Module, int32))(m, v3)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_ExecDropSingleTupleTableSlot(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_english_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v610 int32
	_ = v610
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v725 int32
	_ = v725
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1048 int32
	_ = v1048
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1495 int32
	_ = v1495
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1537 int32
	_ = v1537
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1598 int32
	_ = v1598
	var v1610 int32
	_ = v1610
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	v10 = v7 + int32(2)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 <= v10 {
		v91 = v11
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v1610
L2:
	;
	v1610 = int32(1)
	goto L1
L3:
	;
	v94 = v7 + int32(3)
	v95 = base.B2i32(v91 < v94)
	if v91 < v94 {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v10))))
	if base.B2i32(v15&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v15)%32)&int32(42750482) == int32(0)) != 0 {
		v91 = v11
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = F_find_among(m, l0, int32(_a_F_english_ISO_8859_1_stem_0), int32(15), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v30 == int32(0) {
		v91 = v34
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
	if v37 < v34 {
		v91 = v34
		goto L3
	} else {
		goto L9
	}
L9:
	;
	switch v30 - int32(1) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	case 5:
		goto L12
	case 6:
		goto L11
	case 7:
		goto L10
	default:
		goto L2
	}
L10:
	;
	v86 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_ISO_8859_1_stem_1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L32
	}
L11:
	;
	v80 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_2))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L30
	}
L12:
	;
	v74 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_ISO_8859_1_stem_3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L28
	}
L13:
	;
	v68 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L26
	}
L14:
	;
	v62 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_ISO_8859_1_stem_5))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L24
	}
L15:
	;
	v56 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_6))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L22
	}
L16:
	;
	v50 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_7))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L20
	}
L17:
	;
	v44 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if int32(0) <= v44 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v1610 = v44
	goto L1
L20:
	;
	if int32(0) <= v50 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v1610 = v50
	goto L1
L22:
	;
	if int32(0) <= v56 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v1610 = v56
	goto L1
L24:
	;
	if int32(0) <= v62 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v1610 = v62
	goto L1
L26:
	;
	if int32(0) <= v68 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v1610 = v68
	goto L1
L28:
	;
	if int32(0) <= v74 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v1610 = v74
	goto L1
L30:
	;
	if int32(0) <= v80 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v1610 = v80
	goto L1
L32:
	;
	if int32(0) <= v86 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v1610 = v86
	goto L1
L34:
	;
	v96 = v7
	goto L36
L35:
	;
	v96 = v94
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
	if v91 < v94 {
		v1610 = int32(1)
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v99)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	if v91 == v7 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v142 = v7
	goto L48
L39:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+v7))))
	if v106 == int32(39) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v110 = v7 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110
	v113 = F_slice_del(m, l0)
	mBase = m.M
	if v113 < int32(0) {
		v1610 = v113
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v122 = v104
	goto L42
L42:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v7))))
	if v124 != int32(121) {
		goto L38
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 == v118 {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v122 = v120
	goto L42
L45:
	;
	v127 = int32(1)
	v128 = v7 + v127
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v128
	v133 = F_slice_from_s(m, l0, v127, int32(_a_F_english_ISO_8859_1_stem_9))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	if v133 < int32(0) {
		v1610 = v133
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v137 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v137)
	goto L38
L48:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v156 < v155 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v200
	if v200 <= v94 {
		goto L75
	} else {
		goto L76
	}
L50:
	;
	goto L49
L51:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v199 != 0 {
		goto L66
	} else {
		goto L67
	}
L52:
	;
	v158 = v155
	goto L54
L53:
	;
	v158 = v156
	goto L54
L54:
	;
	goto L56
L55:
	;
	v199 = v194
	goto L51
L56:
	;
	if v155 == v158 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v194 = int32(0)
	goto L55
L58:
	;
	v199 = int32(-1)
	goto L51
L59:
	;
	goto L60
L60:
	;
	v170 = int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v155))))
	if int32(121) < v173 {
		v194 = v170
		goto L55
	} else {
		goto L61
	}
L61:
	;
	v175 = v173 - int32(97)
	if v175 < int32(0) {
		v194 = v170
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v175)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v181)>>(uint(v175&int32(7))%32))&int32(1) == int32(0) {
		v194 = v170
		goto L55
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155 + int32(1)
	goto L64
L64:
	;
	goto L57
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v142
	v215 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v201 + v215
	v220 = F_slice_from_s(m, l0, v215, int32(_a_F_english_ISO_8859_1_stem_10))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L71
	}
L66:
	;
	if v200 <= v142 {
		goto L50
	} else {
		goto L70
	}
L67:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v201
	if v200 == v201 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v201))))
	if v206 == int32(121) {
		goto L65
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	v212 = v142 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v212
	v142 = v212
	goto L48
L71:
	;
	if v220 < int32(0) {
		v1610 = v220
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v224)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v142 = v226
	goto L48
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v483
	if v483 <= v7 {
		v512 = v483
		goto L142
	} else {
		goto L143
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v367
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v377 < v376 {
		goto L112
	} else {
		goto L113
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v263 < v7 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v94))))
	if base.B2i32(v233&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v233)%32)&int32(_a_F_english_ISO_8859_1_stem_11) == int32(0)) != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v248 = F_find_among(m, l0, int32(_a_F_english_ISO_8859_1_stem_12), int32(9), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	if v248 == int32(0) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v367 = v252
	goto L74
L80:
	;
	if v303 < int32(0) {
		goto L73
	} else {
		goto L95
	}
L81:
	;
	v265 = v7
	goto L83
L82:
	;
	v265 = v263
	goto L83
L83:
	;
	v272 = v7
	goto L85
L84:
	;
	v303 = v283
	goto L80
L85:
	;
	if v272 == v265 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v303 = int32(-1)
	goto L80
L88:
	;
	goto L89
L89:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276+v272))))
	if int32(121) < v278 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v295 = v272 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v295
	v272 = v295
	goto L85
L91:
	;
	v280 = v278 - int32(97)
	if v280 < int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v283 = int32(1)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v280)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v287)>>(uint(v280&int32(7))%32))&v283 != 0 {
		goto L84
	} else {
		goto L93
	}
L93:
	;
	goto L90
L95:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v307 = v306 + v303
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v307
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v318 < v307 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v361 < int32(0) {
		goto L73
	} else {
		goto L110
	}
L97:
	;
	v320 = v307
	goto L99
L98:
	;
	v320 = v318
	goto L99
L99:
	;
	v326 = v307
	goto L101
L100:
	;
	v361 = int32(1)
	goto L96
L101:
	;
	if v326 == v320 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v361 = int32(-1)
	goto L96
L104:
	;
	goto L105
L105:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333+v326))))
	if int32(121) < v335 {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v337 = v335 - int32(97)
	if v337 < int32(0) {
		goto L100
	} else {
		goto L107
	}
L107:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v337)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v343)>>(uint(v337&int32(7))%32))&int32(1) == int32(0) {
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v352 = v326 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v352
	v326 = v352
	goto L101
L110:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v365 = v364 + v361
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365
	v367 = v365
	goto L74
L111:
	;
	if v417 < int32(0) {
		goto L73
	} else {
		goto L126
	}
L112:
	;
	v379 = v376
	goto L114
L113:
	;
	v379 = v377
	goto L114
L114:
	;
	v386 = v376
	goto L116
L115:
	;
	v417 = v397
	goto L111
L116:
	;
	if v386 == v379 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v417 = int32(-1)
	goto L111
L119:
	;
	goto L120
L120:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+v386))))
	if int32(121) < v392 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v409 = v386 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v409
	v386 = v409
	goto L116
L122:
	;
	v394 = v392 - int32(97)
	if v394 < int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v397 = int32(1)
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v394)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v401)>>(uint(v394&int32(7))%32))&v397 != 0 {
		goto L115
	} else {
		goto L124
	}
L124:
	;
	goto L121
L126:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v421 = v420 + v417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v421
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v432 < v421 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v475 < int32(0) {
		goto L73
	} else {
		goto L141
	}
L128:
	;
	v434 = v421
	goto L130
L129:
	;
	v434 = v432
	goto L130
L130:
	;
	v440 = v421
	goto L132
L131:
	;
	v475 = int32(1)
	goto L127
L132:
	;
	if v440 == v434 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v475 = int32(-1)
	goto L127
L135:
	;
	goto L136
L136:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v440))))
	if int32(121) < v449 {
		goto L131
	} else {
		goto L137
	}
L137:
	;
	v451 = v449 - int32(97)
	if v451 < int32(0) {
		goto L131
	} else {
		goto L138
	}
L138:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v451)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v457)>>(uint(v451&int32(7))%32))&int32(1) == int32(0) {
		goto L131
	} else {
		goto L139
	}
L139:
	;
	v466 = v440 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v466
	v440 = v466
	goto L132
L141:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v478 + v475
	goto L73
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v512
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v512 <= v515 {
		goto L150
	} else {
		goto L151
	}
L143:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487+v483-int32(1)))))
	if base.B2i32(v491 != int32(115))&base.B2i32(v491 != int32(39)) != 0 {
		v512 = v483
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v500 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_13), int32(3), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	if v500 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v504
	v512 = v504
	goto L142
L147:
	;
	goto L148
L148:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v506
	v508 = F_slice_del(m, l0)
	mBase = m.M
	if v508 < int32(0) {
		v1610 = v508
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v512 = v511
	goto L142
L150:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v629
	v633 = v629 - int32(1)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v633 <= v634 {
		goto L184
	} else {
		goto L185
	}
L151:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517+v512-int32(1)))))
	v523 = v521 - int32(100)
	v524 = int32(0)
	if base.B2i32(v523 == v524)|base.B2i32(v523 == int32(15)) == v524 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v534 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_14), int32(6), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	if v534 == int32(0) {
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v538
	switch v534 - int32(1) {
	case 0:
		goto L157
	case 1:
		goto L156
	case 2:
		goto L155
	default:
		goto L150
	}
L155:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v538 <= v565 {
		goto L150
	} else {
		goto L167
	}
L156:
	;
	v549 = v538 - int32(2)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v550 <= v549 {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	v544 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_15))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	if int32(0) <= v544 {
		goto L150
	} else {
		goto L159
	}
L159:
	;
	v1610 = v544
	goto L1
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v549
	v555 = F_slice_from_s(m, l0, int32(1), int32(_a_F_english_ISO_8859_1_stem_16))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L6
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v561 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_17))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L6
	} else {
		goto L165
	}
L163:
	;
	if int32(0) <= v555 {
		goto L150
	} else {
		goto L164
	}
L164:
	;
	v1610 = v555
	goto L1
L165:
	;
	if int32(0) <= v561 {
		goto L150
	} else {
		goto L166
	}
L166:
	;
	v1610 = v561
	goto L1
L167:
	;
	v568 = v538 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v568
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v584 = v568
	goto L170
L168:
	;
	if v618 < int32(0) {
		goto L150
	} else {
		goto L180
	}
L169:
	;
	v618 = v598
	goto L168
L170:
	;
	if v584 <= v578 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v618 = int32(-1)
	goto L168
L173:
	;
	goto L174
L174:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589+v584-int32(1)))))
	if int32(121) < v593 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v610 = v584 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v610
	v584 = v610
	goto L170
L176:
	;
	v595 = v593 - int32(97)
	if v595 < int32(0) {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v598 = int32(1)
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v595)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v602)>>(uint(v595&int32(7))%32))&v598 != 0 {
		goto L169
	} else {
		goto L178
	}
L178:
	;
	goto L175
L180:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v621 - v618
	v624 = F_slice_del(m, l0)
	mBase = m.M
	if v624 < int32(0) {
		v1610 = v624
		goto L1
	} else {
		goto L181
	}
L181:
	;
	goto L150
L182:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1022
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1022 <= v1025 {
		v1100 = v1025
		goto L276
	} else {
		goto L277
	}
L183:
	;
	v652 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_18), int32(7), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L6
	} else {
		goto L188
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v629
	goto L182
L185:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636+v633))))
	if v638&int32(224) != int32(96) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	if int32(1)<<(uint(v638)%32)&int32(33554576) != 0 {
		goto L183
	} else {
		goto L187
	}
L187:
	;
	goto L184
L188:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v654
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v652 - int32(1) {
	case 0:
		goto L191
	case 1:
		goto L189
	case 2:
		goto L190
	default:
		goto L182
	}
L189:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v786 = v656 - v654
	v787 = v785 - v786
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v787
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v803 = v787
	goto L225
L190:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v654 <= v692 {
		goto L189
	} else {
		goto L201
	}
L191:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v654 < v659 {
		goto L182
	} else {
		goto L192
	}
L192:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v654-int32(2) <= v661 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v682 + (v654 - v656)
	v688 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_19))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L6
	} else {
		goto L199
	}
L194:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665+v654-int32(1)))))
	if v669 != int32(99) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v675 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_20), int32(3), int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	if v675 == int32(0) {
		goto L193
	} else {
		goto L197
	}
L197:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v679 <= v680 {
		goto L182
	} else {
		goto L198
	}
L198:
	;
	goto L193
L199:
	;
	if int32(0) <= v688 {
		goto L182
	} else {
		goto L200
	}
L200:
	;
	v1610 = v688
	goto L1
L201:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v696 = int32(1)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694+v654-v696))))
	if base.B2i32(v698&int32(224) != int32(96))|base.B2i32(v696<<(uint(v698)%32)&int32(34881536) == int32(0)) != 0 {
		goto L189
	} else {
		goto L202
	}
L202:
	;
	v713 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_21), int32(7), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L6
	} else {
		goto L205
	}
L203:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v780 <= v781 {
		goto L182
	} else {
		goto L222
	}
L204:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L208
L205:
	;
	switch v713 {
	case 0:
		goto L189
	case 1:
		goto L204
	case 2:
		goto L203
	default:
		goto L182
	}
L206:
	;
	if v765 != 0 {
		goto L189
	} else {
		goto L218
	}
L207:
	;
	v765 = v762
	goto L206
L208:
	;
	if v716 <= v725 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v762 = int32(0)
	goto L207
L210:
	;
	v765 = int32(-1)
	goto L206
L211:
	;
	goto L212
L212:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736+v716-int32(1)))))
	if int32(121) < v740 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v716 - int32(1)
	goto L217
L214:
	;
	v742 = v740 - int32(97)
	if v742 < int32(0) {
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v745 = int32(1)
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v742)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v749)>>(uint(v742&int32(7))%32))&v745 != 0 {
		v762 = v745
		goto L207
	} else {
		goto L216
	}
L216:
	;
	goto L213
L217:
	;
	goto L209
L218:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v767 < v766 {
		goto L189
	} else {
		goto L219
	}
L219:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v771 = v769 + (v716 - v715)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v771
	v776 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_22))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L6
	} else {
		goto L220
	}
L220:
	;
	if int32(0) <= v776 {
		goto L182
	} else {
		goto L221
	}
L221:
	;
	v1610 = v776
	goto L1
L222:
	;
	goto L189
L223:
	;
	if v837 < int32(0) {
		goto L182
	} else {
		goto L235
	}
L224:
	;
	v837 = v817
	goto L223
L225:
	;
	if v803 <= v797 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v837 = int32(-1)
	goto L223
L228:
	;
	goto L229
L229:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v808+v803-int32(1)))))
	if int32(121) < v812 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v829 = v803 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v829
	v803 = v829
	goto L225
L231:
	;
	v814 = v812 - int32(97)
	if v814 < int32(0) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v817 = int32(1)
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v814)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v821)>>(uint(v814&int32(7))%32))&v817 != 0 {
		goto L224
	} else {
		goto L233
	}
L233:
	;
	goto L230
L235:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v840 - v786
	v843 = F_slice_del(m, l0)
	mBase = m.M
	if v843 < int32(0) {
		v1610 = v843
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v846
	v850 = v846 - int32(1)
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v850 <= v851 {
		v933 = v846
		goto L240
	} else {
		goto L241
	}
L237:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1007 = v1005 + (v846 - v867)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1007
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1007
	if v1007 <= v1004 {
		goto L182
	} else {
		goto L274
	}
L238:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1004 = v1003
	goto L237
L239:
	;
	v999 = F_slice_from_s(m, l0, int32(1), v996)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L6
	} else {
		goto L272
	}
L240:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v933 != v936 {
		goto L182
	} else {
		goto L259
	}
L241:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853+v850))))
	if base.B2i32(v855&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v855)%32)&int32(68514004) == int32(0)) != 0 {
		v933 = v846
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v872 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_23), int32(13), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L6
	} else {
		goto L245
	}
L243:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v933 = v932
	goto L240
L244:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L248
L245:
	;
	switch v872 - int32(1) {
	case 0:
		v996 = int32(_a_F_english_ISO_8859_1_stem_24)
		goto L239
	case 1:
		goto L244
	case 2:
		goto L243
	default:
		goto L238
	}
L246:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v928 != 0 {
		v1004 = v929
		goto L237
	} else {
		goto L257
	}
L247:
	;
	v928 = v924
	goto L246
L248:
	;
	if v884 <= v885 {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v924 = int32(0)
	goto L247
L250:
	;
	v928 = int32(-1)
	goto L246
L251:
	;
	goto L252
L252:
	;
	v897 = int32(1)
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898+v884-v897))))
	if int32(111) < v902 {
		v924 = v897
		goto L247
	} else {
		goto L253
	}
L253:
	;
	v904 = v902 - int32(97)
	if v904 < int32(0) {
		v924 = v897
		goto L247
	} else {
		goto L254
	}
L254:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v904)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v910)>>(uint(v904&int32(7))%32))&int32(1) == int32(0) {
		v924 = v897
		goto L247
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v884 - int32(1)
	goto L256
L256:
	;
	goto L249
L257:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v929 < v930 {
		v1004 = v929
		goto L237
	} else {
		goto L258
	}
L258:
	;
	goto L182
L259:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v947 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_25), int32(89), int32(121), int32(0))
	mBase = m.M
	if v947 != 0 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	if v987 == int32(0) {
		goto L182
	} else {
		goto L271
	}
L261:
	;
	v987 = int32(1)
	goto L260
L262:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v961 = v938 - v941
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v960 - v961
	v968 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(97), int32(121), int32(0))
	mBase = m.M
	if v968 != 0 {
		goto L266
	} else {
		goto L267
	}
L263:
	;
	v952 = F_in_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(97), int32(121), int32(0))
	mBase = m.M
	if v952 != 0 {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v956 = int32(0)
	v957 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(97), int32(121), v956)
	mBase = m.M
	if v957 == v956 {
		goto L261
	} else {
		goto L265
	}
L265:
	;
	goto L262
L266:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v977 - v961
	v982 = F_eq_s_b(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_27))
	mBase = m.M
	if v982 != 0 {
		goto L261
	} else {
		goto L270
	}
L267:
	;
	v973 = F_in_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(97), int32(121), int32(0))
	mBase = m.M
	if v973 != 0 {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v974 <= v975 {
		goto L261
	} else {
		goto L269
	}
L269:
	;
	goto L266
L270:
	;
	v987 = int32(0)
	goto L260
L271:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v990 + (v933 - v938)
	v996 = int32(_a_F_english_ISO_8859_1_stem_28)
	goto L239
L272:
	;
	if int32(0) <= v999 {
		goto L182
	} else {
		goto L273
	}
L273:
	;
	v1610 = v999
	goto L1
L274:
	;
	v1012 = v1007 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1012
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1012
	v1015 = F_slice_del(m, l0)
	mBase = m.M
	if v1015 < int32(0) {
		v1610 = v1015
		goto L1
	} else {
		goto L275
	}
L275:
	;
	goto L182
L276:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1101
	v1105 = v1101 - int32(1)
	if v1105 <= v1100 {
		goto L295
	} else {
		goto L296
	}
L277:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027+v1022-int32(1)))))
	if v1031|int32(32) != int32(121) {
		v1100 = v1025
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1037 = v1022 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1037
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1037
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L281
L279:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1088 != 0 {
		v1100 = v1089
		goto L276
	} else {
		goto L291
	}
L280:
	;
	v1088 = v1085
	goto L279
L281:
	;
	if v1037 <= v1048 {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	v1085 = int32(0)
	goto L280
L283:
	;
	v1088 = int32(-1)
	goto L279
L284:
	;
	goto L285
L285:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059+v1037-int32(1)))))
	if int32(121) < v1063 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1037 - int32(1)
	goto L290
L287:
	;
	v1065 = v1063 - int32(97)
	if v1065 < int32(0) {
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v1068 = int32(1)
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1065)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v1072)>>(uint(v1065&int32(7))%32))&v1068 != 0 {
		v1085 = v1068
		goto L280
	} else {
		goto L289
	}
L289:
	;
	goto L286
L290:
	;
	goto L282
L291:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1090 <= v1089 {
		v1100 = v1089
		goto L276
	} else {
		goto L292
	}
L292:
	;
	v1094 = F_slice_from_s(m, l0, int32(1), int32(_a_F_english_ISO_8859_1_stem_29))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L6
	} else {
		goto L293
	}
L293:
	;
	if v1094 < int32(0) {
		v1610 = v1094
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1100 = v1098
	goto L276
L295:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1294
	v1296 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1294
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1294-int32(2) <= v1299 {
		v1366 = v1296
		goto L362
	} else {
		goto L363
	}
L296:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107+v1105))))
	if base.B2i32(v1109&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1109)%32)&int32(_a_F_english_ISO_8859_1_stem_30) == int32(0)) != 0 {
		goto L295
	} else {
		goto L297
	}
L297:
	;
	v1124 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_31), int32(25), int32(0))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L6
	} else {
		goto L298
	}
L298:
	;
	if v1124 == int32(0) {
		goto L295
	} else {
		goto L299
	}
L299:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1128
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1128 < v1130 {
		goto L295
	} else {
		goto L300
	}
L300:
	;
	switch v1124 - int32(1) {
	case 0:
		goto L316
	case 1:
		goto L315
	case 2:
		goto L314
	case 3:
		goto L313
	case 4:
		goto L312
	case 5:
		goto L311
	case 6:
		goto L310
	case 7:
		goto L309
	case 8:
		goto L308
	case 9:
		goto L307
	case 10:
		goto L306
	case 11:
		goto L305
	case 12:
		goto L304
	case 13:
		goto L303
	case 14:
		goto L302
	case 15:
		goto L301
	default:
		goto L295
	}
L301:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L351
L302:
	;
	v1232 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_32))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L6
	} else {
		goto L347
	}
L303:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1128 <= v1212 {
		goto L295
	} else {
		goto L343
	}
L304:
	;
	v1208 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_33))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L6
	} else {
		goto L341
	}
L305:
	;
	v1202 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_34))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L6
	} else {
		goto L339
	}
L306:
	;
	v1196 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_35))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L6
	} else {
		goto L337
	}
L307:
	;
	v1190 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_36))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L6
	} else {
		goto L335
	}
L308:
	;
	v1184 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_37))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L6
	} else {
		goto L333
	}
L309:
	;
	v1178 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_38))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L6
	} else {
		goto L331
	}
L310:
	;
	v1172 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_39))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L6
	} else {
		goto L329
	}
L311:
	;
	v1166 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_40))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L6
	} else {
		goto L327
	}
L312:
	;
	v1160 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_41))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L6
	} else {
		goto L325
	}
L313:
	;
	v1154 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_42))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L6
	} else {
		goto L323
	}
L314:
	;
	v1148 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_43))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L6
	} else {
		goto L321
	}
L315:
	;
	v1142 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_44))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L6
	} else {
		goto L319
	}
L316:
	;
	v1136 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_45))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L6
	} else {
		goto L317
	}
L317:
	;
	if int32(0) <= v1136 {
		goto L295
	} else {
		goto L318
	}
L318:
	;
	v1610 = v1136
	goto L1
L319:
	;
	if int32(0) <= v1142 {
		goto L295
	} else {
		goto L320
	}
L320:
	;
	v1610 = v1142
	goto L1
L321:
	;
	if int32(0) <= v1148 {
		goto L295
	} else {
		goto L322
	}
L322:
	;
	v1610 = v1148
	goto L1
L323:
	;
	if int32(0) <= v1154 {
		goto L295
	} else {
		goto L324
	}
L324:
	;
	v1610 = v1154
	goto L1
L325:
	;
	if int32(0) <= v1160 {
		goto L295
	} else {
		goto L326
	}
L326:
	;
	v1610 = v1160
	goto L1
L327:
	;
	if int32(0) <= v1166 {
		goto L295
	} else {
		goto L328
	}
L328:
	;
	v1610 = v1166
	goto L1
L329:
	;
	if int32(0) <= v1172 {
		goto L295
	} else {
		goto L330
	}
L330:
	;
	v1610 = v1172
	goto L1
L331:
	;
	if int32(0) <= v1178 {
		goto L295
	} else {
		goto L332
	}
L332:
	;
	v1610 = v1178
	goto L1
L333:
	;
	if int32(0) <= v1184 {
		goto L295
	} else {
		goto L334
	}
L334:
	;
	v1610 = v1184
	goto L1
L335:
	;
	if int32(0) <= v1190 {
		goto L295
	} else {
		goto L336
	}
L336:
	;
	v1610 = v1190
	goto L1
L337:
	;
	if int32(0) <= v1196 {
		goto L295
	} else {
		goto L338
	}
L338:
	;
	v1610 = v1196
	goto L1
L339:
	;
	if int32(0) <= v1202 {
		goto L295
	} else {
		goto L340
	}
L340:
	;
	v1610 = v1202
	goto L1
L341:
	;
	if int32(0) <= v1208 {
		goto L295
	} else {
		goto L342
	}
L342:
	;
	v1610 = v1208
	goto L1
L343:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214+v1128-int32(1)))))
	if v1218 != int32(108) {
		goto L295
	} else {
		goto L344
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1128 - int32(1)
	v1226 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_46))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L6
	} else {
		goto L345
	}
L345:
	;
	if int32(0) <= v1226 {
		goto L295
	} else {
		goto L346
	}
L346:
	;
	v1610 = v1226
	goto L1
L347:
	;
	if int32(0) <= v1232 {
		goto L295
	} else {
		goto L348
	}
L348:
	;
	v1610 = v1232
	goto L1
L349:
	;
	if v1288 != 0 {
		goto L295
	} else {
		goto L360
	}
L350:
	;
	v1288 = v1284
	goto L349
L351:
	;
	if v1244 <= v1245 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v1284 = int32(0)
	goto L350
L353:
	;
	v1288 = int32(-1)
	goto L349
L354:
	;
	goto L355
L355:
	;
	v1257 = int32(1)
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1258+v1244-v1257))))
	if int32(116) < v1262 {
		v1284 = v1257
		goto L350
	} else {
		goto L356
	}
L356:
	;
	v1264 = v1262 - int32(99)
	if v1264 < int32(0) {
		v1284 = v1257
		goto L350
	} else {
		goto L357
	}
L357:
	;
	v1270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1264)>>(uint(int32(3))%32)))+uint32(_c_F_english_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v1270)>>(uint(v1264&int32(7))%32))&int32(1) == int32(0) {
		v1284 = v1257
		goto L350
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1244 - int32(1)
	goto L359
L359:
	;
	goto L352
L360:
	;
	v1289 = F_slice_del(m, l0)
	mBase = m.M
	if v1289 < int32(0) {
		v1610 = v1289
		goto L1
	} else {
		goto L361
	}
L361:
	;
	goto L295
L362:
	;
	if v1366 < int32(0) {
		v1610 = v1366
		goto L1
	} else {
		goto L386
	}
L363:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1305 = int32(1)
	v1307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303+v1294-v1305))))
	if base.B2i32(v1307&int32(224) != int32(96))|base.B2i32(v1305<<(uint(v1307)%32)&int32(_a_F_english_ISO_8859_1_stem_47) == int32(0)) != 0 {
		v1366 = v1296
		goto L362
	} else {
		goto L364
	}
L364:
	;
	v1322 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_48), int32(9), int32(0))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L6
	} else {
		goto L365
	}
L365:
	;
	if v1322 == int32(0) {
		v1366 = v1296
		goto L362
	} else {
		goto L366
	}
L366:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1326
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1326 < v1328 {
		v1366 = v1296
		goto L362
	} else {
		goto L367
	}
L367:
	;
	switch v1322 - int32(1) {
	case 0:
		goto L374
	case 1:
		goto L373
	case 2:
		goto L372
	case 3:
		goto L371
	case 4:
		goto L370
	case 5:
		goto L369
	default:
		goto L368
	}
L368:
	;
	v1366 = int32(1)
	goto L362
L369:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1326 < v1359 {
		v1366 = v1296
		goto L362
	} else {
		goto L384
	}
L370:
	;
	v1356 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1356 {
		goto L368
	} else {
		goto L383
	}
L371:
	;
	v1352 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_49))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L6
	} else {
		goto L381
	}
L372:
	;
	v1346 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_ISO_8859_1_stem_50))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L6
	} else {
		goto L379
	}
L373:
	;
	v1340 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_ISO_8859_1_stem_51))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L6
	} else {
		goto L377
	}
L374:
	;
	v1334 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_52))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L6
	} else {
		goto L375
	}
L375:
	;
	if int32(0) <= v1334 {
		goto L368
	} else {
		goto L376
	}
L376:
	;
	v1366 = v1334
	goto L362
L377:
	;
	if int32(0) <= v1340 {
		goto L368
	} else {
		goto L378
	}
L378:
	;
	v1366 = v1340
	goto L362
L379:
	;
	if int32(0) <= v1346 {
		goto L368
	} else {
		goto L380
	}
L380:
	;
	v1366 = v1346
	goto L362
L381:
	;
	if int32(0) <= v1352 {
		goto L368
	} else {
		goto L382
	}
L382:
	;
	v1366 = v1352
	goto L362
L383:
	;
	v1366 = v1356
	goto L362
L384:
	;
	v1361 = F_slice_del(m, l0)
	mBase = m.M
	if v1361 < int32(0) {
		v1366 = v1361
		goto L362
	} else {
		goto L385
	}
L385:
	;
	goto L368
L386:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1371
	v1373 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1371
	v1377 = v1371 - int32(1)
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1377 <= v1378 {
		v1431 = v1373
		goto L387
	} else {
		goto L388
	}
L387:
	;
	if v1431 < int32(0) {
		v1610 = v1431
		goto L1
	} else {
		goto L400
	}
L388:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1380+v1377))))
	if base.B2i32(v1382&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1382)%32)&int32(_a_F_english_ISO_8859_1_stem_53) == int32(0)) != 0 {
		v1431 = v1373
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1397 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_54), int32(18), int32(0))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L6
	} else {
		goto L390
	}
L390:
	;
	if v1397 == int32(0) {
		v1431 = v1373
		goto L387
	} else {
		goto L391
	}
L391:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1401
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1401 < v1403 {
		v1431 = v1373
		goto L387
	} else {
		goto L392
	}
L392:
	;
	switch v1397 - int32(1) {
	case 0:
		goto L395
	case 1:
		goto L394
	default:
		goto L393
	}
L393:
	;
	v1431 = int32(1)
	goto L387
L394:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1401 <= v1410 {
		v1431 = v1373
		goto L387
	} else {
		goto L397
	}
L395:
	;
	v1407 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1407 {
		goto L393
	} else {
		goto L396
	}
L396:
	;
	v1431 = v1407
	goto L387
L397:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1414 = int32(1)
	v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1412+v1401-v1414))))
	if base.Ui32(v1414) < base.Ui32((v1416-int32(115))&int32(255)) {
		v1431 = v1373
		goto L387
	} else {
		goto L398
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1401 - int32(1)
	v1426 = F_slice_del(m, l0)
	mBase = m.M
	if v1426 < int32(0) {
		v1431 = v1426
		goto L387
	} else {
		goto L399
	}
L399:
	;
	goto L393
L400:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1436
	v1438 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1436
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1436 <= v1441 {
		v1544 = v1438
		goto L401
	} else {
		goto L402
	}
L401:
	;
	if v1544 < int32(0) {
		v1610 = v1544
		goto L1
	} else {
		goto L430
	}
L402:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443+v1436-int32(1)))))
	switch v1447 - int32(101) {
	case 0, 7:
		goto L403
	default:
		v1544 = v1438
		goto L401
	}
L403:
	;
	v1453 = F_find_among_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_55), int32(2), int32(0))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L6
	} else {
		goto L404
	}
L404:
	;
	if v1453 == int32(0) {
		v1544 = v1438
		goto L401
	} else {
		goto L405
	}
L405:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1457
	switch v1453 - int32(1) {
	case 0:
		goto L408
	case 1:
		goto L407
	default:
		goto L406
	}
L406:
	;
	v1544 = int32(1)
	goto L401
L407:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1457 < v1523 {
		v1544 = v1438
		goto L401
	} else {
		goto L426
	}
L408:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1457 < v1461 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1457 < v1463 {
		v1544 = v1438
		goto L401
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	v1520 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1520 {
		goto L406
	} else {
		goto L425
	}
L412:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1474 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_25), int32(89), int32(121), int32(0))
	mBase = m.M
	if v1474 != 0 {
		goto L415
	} else {
		goto L416
	}
L413:
	;
	if v1514 != 0 {
		v1544 = v1438
		goto L401
	} else {
		goto L424
	}
L414:
	;
	v1514 = int32(1)
	goto L413
L415:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1488 = v1465 - v1468
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1487 - v1488
	v1495 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1495 != 0 {
		goto L419
	} else {
		goto L420
	}
L416:
	;
	v1479 = F_in_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1479 != 0 {
		goto L415
	} else {
		goto L417
	}
L417:
	;
	v1483 = int32(0)
	v1484 = F_out_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(97), int32(121), v1483)
	mBase = m.M
	if v1484 == v1483 {
		goto L414
	} else {
		goto L418
	}
L418:
	;
	goto L415
L419:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1504 - v1488
	v1509 = F_eq_s_b(m, l0, int32(4), int32(_a_F_english_ISO_8859_1_stem_27))
	mBase = m.M
	if v1509 != 0 {
		goto L414
	} else {
		goto L423
	}
L420:
	;
	v1500 = F_in_grouping_b(m, l0, int32(_a_F_english_ISO_8859_1_stem_26), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1500 != 0 {
		goto L419
	} else {
		goto L421
	}
L421:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1501 <= v1502 {
		goto L414
	} else {
		goto L422
	}
L422:
	;
	goto L419
L423:
	;
	v1514 = int32(0)
	goto L413
L424:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1515 + (v1457 - v1465)
	goto L411
L425:
	;
	v1544 = v1520
	goto L401
L426:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1457 <= v1525 {
		v1544 = v1438
		goto L401
	} else {
		goto L427
	}
L427:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527+v1457-int32(1)))))
	if v1531 != int32(108) {
		v1544 = v1438
		goto L401
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1457 - int32(1)
	v1537 = F_slice_del(m, l0)
	mBase = m.M
	if v1537 < int32(0) {
		v1544 = v1537
		goto L401
	} else {
		goto L429
	}
L429:
	;
	goto L406
L430:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1548
	v1550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v1550 != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	goto L435
L432:
	;
	v1598 = int32(0)
	goto L433
L433:
	;
	if v1598 < int32(0) {
		v1610 = v1598
		goto L1
	} else {
		goto L451
	}
L434:
	;
	v1598 = v1590
	goto L433
L435:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1558 < v1557 {
		goto L437
	} else {
		goto L438
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1557
	v1590 = int32(1)
	goto L434
L437:
	;
	v1560 = v1557
	goto L439
L438:
	;
	v1560 = v1558
	goto L439
L439:
	;
	v1562 = v1557
	goto L441
L440:
	;
	goto L436
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1562
	if v1562 != v1558 {
		goto L444
	} else {
		goto L445
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1562
	v1579 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1562 + v1579
	v1584 = F_slice_from_s(m, l0, v1579, int32(_a_F_english_ISO_8859_1_stem_56))
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L6
	} else {
		goto L449
	}
L443:
	;
	goto L442
L444:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1569+v1562))))
	if v1571 == int32(89) {
		goto L443
	} else {
		goto L447
	}
L445:
	;
	goto L446
L446:
	;
	if v1562 == v1560 {
		goto L440
	} else {
		goto L448
	}
L447:
	;
	goto L446
L448:
	;
	v1576 = v1562 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1576
	v1562 = v1576
	goto L441
L449:
	;
	if int32(0) <= v1584 {
		goto L435
	} else {
		goto L450
	}
L450:
	;
	v1590 = v1584
	goto L434
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1548
	goto L2
}
func F_equality_ops_are_compatible(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	if l0 != l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = int64(0)
	v13 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(l0), v11, v11)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v72 = int32(1)
	goto L3
L3:
	;
	return v72
L4:
	;
	F_ReleaseCatCacheList(m, v13)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L18
	}
L5:
	;
	return int32(0)
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if int32(0) < v17 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v24 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	v67 = int32(0)
	goto L4
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(-64)+v24<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	v37 = v35 + v36
	v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+4)))
	v40 = F_SearchSysCacheExists(m, int32(3), base.I64_extend_i32_u(l1), int64(115), v38, int64(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v53 = v24 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if v53 < v54 {
		v24 = v53
		goto L10
	} else {
		goto L17
	}
L13:
	;
	if v40 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v46 = F_GetIndexAmRoutineByAmId(m, v44, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+13)))
	if v48 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v67 = int32(1)
	goto L4
L17:
	;
	goto L11
L18:
	;
	v72 = v67
	goto L3
}
func F_errcode_for_socket_access(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_errcode_for_socket_access[0]))
	if int32(0) <= v5 {
		v9 = v5 * int32(100)
		v12 = int32(100663808)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_errcode_for_socket_access[1])))
		switch v13 - int32(13) {
		case 0, 2, 10, 25, 26, 27, 51, 60:
			v19 = v12
		case 1, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 52, 53, 54, 55, 56, 57, 58, 59:
			v19 = int32(2600)
		default:
			if v13 == int32(142) {
				v19 = v12
			} else {
				v19 = int32(2600)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_errcode_for_socket_access[2]))) = v19
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_errcode_for_socket_access[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_errcode_for_socket_access_0), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_errcode_for_socket_access_1), int32(982), int32(_a_F_errcode_for_socket_access_2))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_errdatatype(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_errdatatype_0), v6)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_errdatatype_1), int32(414), int32(_a_F_errdatatype_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v30 = v28 + v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
			v32 = F_get_namespace_name(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_err_generic_string(m, int32(115), v32)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_err_generic_string(m, int32(100), v30+int32(4))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_ReleaseCatCache(m, v10)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_errorMissingRTE(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v113 int32
	_ = v113
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v520 int32
	_ = v520
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l0 != 0 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	v154 = int32(0)
	v158 = F_RangeVarGetRelidExtended(m, l1, v154, int32(1), v154, v154)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L28
	} else {
		goto L29
	}
L3:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = l0
	v31 = v3
	goto L7
L5:
	;
	goto L6
L6:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v130 = F_get_visible_ENR_metadata(m, v129, v20)
	mBase = m.M
	goto L26
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v113 != 0 {
		v24 = v113
		v31 = v31 + int32(1)
		goto L7
	} else {
		goto L25
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v48 = int32(0)
	goto L12
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v43+v48<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v67 == int32(0))|base.B2i32(v67 != v70) != 0 {
		v88 = v67
		v89 = v70
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v164 = int32(0)
	v168 = v3
	v171 = v31
	v177 = int32(1)
	goto L1
L14:
	;
	if v88-v89 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	v73 = v64
	v74 = v20
	goto L17
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v78 == int32(0) {
		v88 = v78
		v89 = v77
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v88 = v78
	v89 = v77
	goto L15
L19:
	;
	v81 = int32(1)
	if v78 == v77 {
		v73 = v73 + v81
		v74 = v74 + v81
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v92 = v48 + int32(1)
	if v92 != v40 {
		v48 = v92
		goto L12
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L13
L24:
	;
	goto L9
L25:
	;
	goto L8
L26:
	;
	if base.B2i32(v130 != int32(0)) == int32(0) {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v135 = int32(0)
	v164 = v135
	v168 = int32(1)
	v171 = v135
	v177 = v135
	goto L1
L28:
	;
	return
L29:
	;
	v160 = int32(0)
	v164 = v158
	v168 = v3
	v171 = v160
	v177 = v160
	goto L1
L30:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L28
	} else {
		goto L124
	}
L31:
	;
	v178 = int32(1)
	v188 = l0
	v192 = v3
	goto L34
L32:
	;
	goto L33
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L28
	} else {
		goto L119
	}
L34:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	if v197 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if v520 != 0 {
		v188 = v520
		v192 = v192 + int32(1)
		goto L34
	} else {
		goto L118
	}
L37:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v200 <= int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v208 = int32(0)
	goto L39
L39:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v203+v208<<(uint(int32(2))%32))))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v225 = int32(0)
	if v224|base.B2i32(v164 == v225) == v225 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v332 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L41:
	;
	goto L40
L42:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v303 == int32(0))|base.B2i32(v303 != v306) != 0 {
		v324 = v303
		v325 = v306
		goto L69
	} else {
		goto L70
	}
L43:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	if v230 != v164 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v177^v178|base.B2i32(v224 != int32(6)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L41
L47:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v223)+88))
	if v237+v192 != v171 {
		goto L42
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v168^v178|base.B2i32(v224 != int32(7)) != 0 {
		goto L42
	} else {
		goto L59
	}
L50:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v223)+84))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v243 == int32(0))|base.B2i32(v243 != v246) != 0 {
		v264 = v243
		v265 = v246
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v264-v265 != 0 {
		goto L42
	} else {
		goto L58
	}
L52:
	;
	goto L51
L53:
	;
	v249 = v240
	v250 = v20
	goto L54
L54:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	if v254 == int32(0) {
		v264 = v254
		v265 = v253
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v264 = v254
	v265 = v253
	goto L52
L56:
	;
	v257 = int32(1)
	if v254 == v253 {
		v249 = v249 + v257
		v250 = v250 + v257
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L41
L59:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v223)+108))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v273 == int32(0))|base.B2i32(v273 != v276) != 0 {
		v294 = v273
		v295 = v276
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v294-v295 == int32(0) {
		goto L41
	} else {
		goto L67
	}
L61:
	;
	goto L60
L62:
	;
	v279 = v270
	v280 = v20
	goto L63
L63:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if v284 == int32(0) {
		v294 = v284
		v295 = v283
		goto L61
	} else {
		goto L65
	}
L64:
	;
	v294 = v284
	v295 = v283
	goto L61
L65:
	;
	v287 = int32(1)
	if v284 == v283 {
		v279 = v279 + v287
		v280 = v280 + v287
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L42
L68:
	;
	if v324-v325 == int32(0) {
		goto L41
	} else {
		goto L75
	}
L69:
	;
	goto L68
L70:
	;
	v309 = v300
	v310 = v20
	goto L71
L71:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	if v314 == int32(0) {
		v324 = v314
		v325 = v313
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v324 = v314
	v325 = v313
	goto L69
L73:
	;
	v317 = int32(1)
	if v314 == v313 {
		v309 = v309 + v317
		v310 = v310 + v317
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v330 = v208 + int32(1)
	if v330 != v200 {
		v208 = v330
		goto L39
	} else {
		goto L76
	}
L76:
	;
	goto L36
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L28
	} else {
		goto L112
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L28
	} else {
		goto L92
	}
L79:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	if base.B2i32(v340 == int32(0))|base.B2i32(v340 != v343) != 0 {
		v361 = v340
		v362 = v343
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v361-v362 == int32(0) {
		goto L78
	} else {
		goto L87
	}
L81:
	;
	goto L80
L82:
	;
	v346 = v336
	v347 = v337
	goto L83
L83:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+1)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+1)))
	if v351 == int32(0) {
		v361 = v351
		v362 = v350
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v361 = v351
	v362 = v350
	goto L81
L85:
	;
	v354 = int32(1)
	if v351 == v350 {
		v346 = v346 + v354
		v347 = v347 + v354
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v370 = F_refnameNamespaceItem(m, l0, int32(0), v336, v367, v18+int32(76))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L28
	} else {
		goto L88
	}
L88:
	;
	if v370 == int32(0) {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v374 != v223 {
		goto L78
	} else {
		goto L90
	}
L90:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v377 != 0 {
		goto L77
	} else {
		goto L91
	}
L91:
	;
	goto L78
L92:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L28
	} else {
		goto L93
	}
L93:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v386
	F_errmsg(m, int32(_a_F_errorMissingRTE_0), v18+int32(32))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L28
	} else {
		goto L94
	}
L94:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v394
	v399 = F_errdetail(m, int32(_a_F_errorMissingRTE_1), v18+int32(16))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L28
	} else {
		goto L95
	}
L95:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v401 != 0 {
		goto L30
	} else {
		goto L96
	}
L96:
	;
	v408 = l0
	goto L97
L97:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v408)+28))
	if v417 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L30
L99:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	if v474 != 0 {
		v408 = v474
		goto L97
	} else {
		goto L111
	}
L100:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	if v420 <= int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v428 = int32(0)
	goto L102
L102:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v423+v428<<(uint(int32(2))%32))))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v223 != v444 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+22)))
	if v449 != int32(1) {
		goto L30
	} else {
		goto L108
	}
L104:
	;
	v447 = v428 + int32(1)
	if v447 != v420 {
		v428 = v447
		goto L102
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	goto L99
L108:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+23)))
	if v452 != int32(1) {
		goto L30
	} else {
		goto L109
	}
L109:
	;
	F_errhint(m, int32(_a_F_errorMissingRTE_2), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L28
	} else {
		goto L110
	}
L110:
	;
	goto L30
L111:
	;
	goto L98
L112:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L28
	} else {
		goto L113
	}
L113:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v482
	F_errmsg(m, int32(_a_F_errorMissingRTE_0), v18-int32(-64))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L28
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v377
	F_errhint(m, int32(_a_F_errorMissingRTE_3), v18+int32(48))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L28
	} else {
		goto L115
	}
L115:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L28
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_errorMissingRTE_4), int32(3766), int32(_a_F_errorMissingRTE_5))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L28
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	goto L35
L119:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L28
	} else {
		goto L120
	}
L120:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v543
	F_errmsg(m, int32(_a_F_errorMissingRTE_6), v18)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L28
	} else {
		goto L121
	}
L121:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L28
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_errorMissingRTE_4), int32(3784), int32(_a_F_errorMissingRTE_5))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L28
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errfinish(m, int32(_a_F_errorMissingRTE_4), int32(3777), int32(_a_F_errorMissingRTE_5))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L28
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_escape_json_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int64
	_ = v71
	var v77 int64
	_ = v77
	var v98 int64
	_ = v98
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v167 int64
	_ = v167
	var v170 int64
	_ = v170
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	F_enlargeStringInfo(m, l0, l2+int32(2))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 <= v22+int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v43 = l2 & int32(-8)
	v48 = int32(0)
	goto L8
L4:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v29+v22))) = uint8(v31)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = v33 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37+v35))) = uint8(v39)
	goto L3
L7:
	;
	goto L3
L8:
	;
	if v48 < v43 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v61 = v48
	v62 = v48
	goto L13
L11:
	;
	v227 = v48
	goto L12
L12:
	;
	v237 = v227 + int32(8)
	v241 = v227
	goto L34
L13:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(l1+v62)))
	if int64(0) <= v71 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v214 < v215 {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	goto L14
L16:
	;
	v204 = v62 - v61
	if int32(512) <= v204 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v77 = v71&int64(36170086419038336) ^ int64(-9187201950435737472)
	if v77&(v71-int64(2314885530818453536))|v77&(v71^int64(2459565876494606882)-int64(72340172838076673)) != int64(0) {
		v214 = v61
		v215 = v62
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v98 = int64(0)
	if base.B2i32(v71&int64(63050394783186944) == v98)|base.B2i32(v71&int64(246290604621824) == v98)|(base.B2i32(v71&int64(962072674304) == v98)|base.B2i32(v71&int64(3758096384) == v98))|(base.B2i32(v71&int64(57344) == v98)|(base.B2i32(v71&int64(14680064) == v98)|base.B2i32(v71&int64(224) == v98))) != 0 {
		v214 = v61
		v215 = v62
		goto L15
	} else {
		goto L22
	}
L20:
	;
	if v77&(v71^int64(6655295901103053916)-int64(72340172838076673)) == int64(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v214 = v61
	v215 = v62
	goto L15
L22:
	;
	v131 = v71 ^ int64(2459565876494606882)
	v134 = int64(0)
	if base.B2i32(v131&int64(71776119061217280) == v134)|base.B2i32(v131&int64(280375465082880) == v134)|(base.B2i32(v131&int64(1095216660480) == v134)|base.B2i32(v131&int64(4278190080) == v134))|(base.B2i32(v131&int64(65280) == v134)|(base.B2i32(v131&int64(16711680) == v134)|base.B2i32(v131&int64(255) == v134))) != 0 {
		v214 = v61
		v215 = v62
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v167 = v71 ^ int64(6655295901103053916)
	v170 = int64(0)
	if base.B2i32(v167&int64(71776119061217280) == v170)|base.B2i32(v167&int64(280375465082880) == v170)|(base.B2i32(v167&int64(1095216660480) == v170)|base.B2i32(v167&int64(4278190080) == v170))|(base.B2i32(v167&int64(65280) == v170)|(base.B2i32(v167&int64(16711680) == v170)|base.B2i32(v167&int64(255) == v170))) != 0 {
		v214 = v61
		v215 = v62
		goto L15
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	F_appendBinaryStringInfo(m, l0, l1+v61, v204)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v210 = v61
	goto L27
L27:
	;
	v212 = v62 + int32(8)
	if v212 < v43 {
		v61 = v210
		v62 = v212
		goto L13
	} else {
		goto L29
	}
L28:
	;
	v210 = v62
	goto L27
L29:
	;
	v214 = v210
	v215 = v212
	goto L15
L30:
	;
	F_appendBinaryStringInfo(m, l0, l1+v214, v215-v214)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v227 = v215
	goto L12
L33:
	;
	goto L32
L34:
	;
	if l2 != v241 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v48 = v237
	goto L8
L36:
	;
	v331 = v241 + int32(1)
	if v331 != v237 {
		v241 = v331
		goto L34
	} else {
		goto L70
	}
L37:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L69
	}
L38:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v241))))
	switch v252 - int32(8) {
	case 0:
		goto L48
	case 1:
		goto L44
	case 2:
		goto L46
	case 3, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		goto L41
	case 4:
		goto L47
	case 5:
		goto L45
	case 26:
		goto L43
	default:
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v300 <= v301+int32(1) {
		goto L65
	} else {
		goto L66
	}
L41:
	;
	v275 = base.I32_extend8_s(v252)
	if base.Ui32(v252) <= base.Ui32(int32(31)) {
		goto L56
	} else {
		goto L57
	}
L42:
	;
	if v252 == int32(92) {
		goto L37
	} else {
		goto L55
	}
L43:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L54
	}
L44:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_2))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L53
	}
L45:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_3))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L52
	}
L46:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_4))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L51
	}
L47:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_5))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_with_len_6))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L36
L50:
	;
	goto L36
L51:
	;
	goto L36
L52:
	;
	goto L36
L53:
	;
	goto L36
L54:
	;
	goto L36
L55:
	;
	goto L41
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v275
	F_appendStringInfo(m, l0, int32(_a_F_escape_json_with_len_7), v15)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v282 <= v283+int32(1) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L36
L60:
	;
	F_appendStringInfoChar(m, l0, v275)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v289+v283))) = uint8(v252)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v294 = v292 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v298 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v296+v294))) = uint8(v298)
	goto L36
L63:
	;
	goto L36
L64:
	;
	m.G0 = v15 + int32(16)
	return
L65:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v310 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v308+v301))) = uint8(v310)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v314 = v312 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v318 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v316+v314))) = uint8(v318)
	goto L64
L68:
	;
	goto L64
L69:
	;
	goto L36
L70:
	;
	goto L35
}
func F_eval_const_expressions_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int64
	_ = v265
	var v274 int32
	_ = v274
	var v277 int64
	_ = v277
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int64
	_ = v429
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int64
	_ = v666
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v790 int64
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v807 int32
	_ = v807
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int64
	_ = v905
	var v907 int64
	_ = v907
	var v909 int64
	_ = v909
	var v911 int64
	_ = v911
	var v913 int64
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v984 int64
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int64
	_ = v1047
	var v1049 int64
	_ = v1049
	var v1051 int64
	_ = v1051
	var v1053 int64
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1183 int32
	_ = v1183
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1295 int32
	_ = v1295
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1548 int32
	_ = v1548
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1668 int64
	_ = v1668
	var v1672 int32
	_ = v1672
	var v1673 int64
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1678 int64
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1701 int64
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1912 int32
	_ = v1912
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(160)
	m.G0 = v16
	F_check_stack_depth(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		v1912 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v16 + int32(160)
	return v1912
L4:
	;
	v24 = l0
	goto L7
L5:
	;
	v1908 = F_copyObjectImpl(m, v24)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L601
	}
L6:
	;
	v1885 = F_palloc0(m, int32(36))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L1
	} else {
		goto L600
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	switch v37 - int32(8) {
	case 0:
		goto L39
	case 1:
		goto L36
	default:
		goto L12
	case 3:
		goto L38
	case 6, 27, 28, 31:
		goto L22
	case 7:
		goto L37
	case 9:
		goto L35
	case 10:
		goto L34
	case 11:
		goto L33
	case 12:
		goto L32
	case 13:
		goto L31
	case 15, 16:
		v1912 = v24
		goto L3
	case 17:
		goto L18
	case 19:
		goto L28
	case 20:
		goto L27
	case 21:
		goto L26
	case 22:
		goto L13
	case 23:
		goto L25
	case 24:
		goto L24
	case 26:
		goto L23
	case 30:
		goto L21
	case 32:
		goto L20
	case 33:
		goto L19
	case 36:
		goto L30
	case 37:
		goto L29
	case 44:
		goto L17
	case 45:
		goto L16
	case 47:
		goto L15
	case 313:
		goto L14
	}
L8:
	;
	v1912 = int32(0)
	goto L3
L9:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v24+v1876)))
	F_check_stack_depth(m)
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L598
	}
L10:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+4))
	v1865 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+8)))
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+28))
	v1870 = F_makeVar(m, v1864, v1865, v1866, v1867, v1868, v1869)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L1
	} else {
		goto L597
	}
L11:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v1912 = v1863
	goto L3
L12:
	;
	v1861 = F_expression_tree_mutator_impl(m, v24, int32(919), l1)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L596
	}
L13:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v1819 = F_eval_const_expressions_mutator(m, v1818, l1)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L1
	} else {
		goto L578
	}
L14:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v1813 != 0 {
		goto L5
	} else {
		goto L576
	}
L15:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v1754 = F_eval_const_expressions_mutator(m, v1753, l1)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L1
	} else {
		goto L558
	}
L16:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v1654 = F_eval_const_expressions_mutator(m, v1653, l1)
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L1
	} else {
		goto L523
	}
L17:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v1459 = F_eval_const_expressions_mutator(m, v1458, l1)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L460
	}
L18:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v1369 = F_eval_const_expressions_mutator(m, v1368, l1)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L1
	} else {
		goto L429
	}
L19:
	;
	v1347 = F_expression_tree_mutator_impl(m, v24, int32(919), l1)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L1
	} else {
		goto L417
	}
L20:
	;
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1338 != int32(1) {
		goto L5
	} else {
		goto L415
	}
L21:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v1245 != 0 {
		goto L388
	} else {
		goto L389
	}
L22:
	;
	v1231 = F_expression_tree_mutator_impl(m, v24, int32(919), l1)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L1
	} else {
		goto L380
	}
L23:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1225 == int32(0) {
		goto L5
	} else {
		goto L378
	}
L24:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1104 = F_eval_const_expressions_mutator(m, v1103, l1)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L353
	}
L25:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v1091 = F_eval_const_expressions_mutator(m, v1090, l1)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L347
	}
L26:
	;
	v1043 = F_palloc0(m, int32(32))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L1
	} else {
		goto L334
	}
L27:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v942
	*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v942
	v948 = F_list_make1_impl(m, int32(1), v16+int32(44))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L320
	}
L28:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v932 = F_eval_const_expressions_mutator(m, v931, l1)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L318
	}
L29:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v896 == int32(3) {
		goto L311
	} else {
		goto L312
	}
L30:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v884 = F_eval_const_expressions_mutator(m, v883, l1)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L303
	}
L31:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	switch v606 {
	case 0:
		goto L204
	case 1:
		goto L205
	case 2:
		goto L203
	default:
		goto L202
	}
L32:
	;
	v578 = F_expression_tree_mutator_impl(m, v24, int32(919), l1)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L188
	}
L33:
	;
	v497 = F_expression_tree_mutator_impl(m, v24, int32(919), l1)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L162
	}
L34:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v319 = F_expression_tree_mutator_impl(m, v317, int32(919), l1)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L111
	}
L35:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v237
	F_set_opfuncid(m, v24)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L86
	}
L36:
	;
	v208 = F_expression_tree_mutator_impl(m, v24, int32(919), l1)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L77
	}
L37:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v173 = F_exprTypmod(m, v24)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L73
	}
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v112 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v110))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L61
	}
L39:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v40 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v41 == int32(0) {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v44 <= int32(0) {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	if v47 < v44 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v49 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v61 == int32(0) {
		goto L5
	} else {
		goto L49
	}
L45:
	;
	v53 = m.T0[v49].(func(*base.Module, int32, int32, int32, int32) int32)(m, v41, v44, int32(1), v16+int32(144))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v60 = v41 + v44<<(uint(int32(4))%32) + int32(16)
	goto L44
L48:
	;
	v60 = v53
	goto L44
L49:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v61 != v64 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v66 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)))
	if v69&int32(1) == int32(0) {
		goto L5
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_get_typlenbyval(m, v61, v16+int32(140), v16+int32(132))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+132)))
	if v81|v82&int32(1) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+140)))
	v90 = F_datumCopy(m, v80, int32(0), v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	v94 = v82
	v95 = v81
	v96 = v80
	goto L58
L58:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+140)))
	v101 = int32(1)
	v105 = F_makeConst(m, v97, v98, v99, v100, v96, v95&v101, v94&v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+132)))
	v94 = v93
	v95 = v92
	v96 = v90
	goto L58
L60:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+36)) = v107
	v1912 = v105
	goto L3
L61:
	;
	if v112 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v132 = F_expand_function_arguments(m, v129, int32(0), v131, v112)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v110
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_0), v16)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(2889), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_ReleaseCatCache(m, v112)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v137 = F_expression_tree_mutator_impl(m, v132, int32(919), l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v140 = F_eval_const_expressions_mutator(m, v139, l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v143 = F_palloc0(m, int32(48))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = int32(11)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+12)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+24)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v143)+16)) = v153
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+28)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+32)) = v159
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+36)) = uint8(v161)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+37)) = uint8(v163)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+40)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+44)) = v167
	v1912 = v143
	goto L3
L73:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+13)))
	v180 = int32(1)
	v182 = F_simplify_function(m, v171, v172, v173, v175, v176, v16+int32(144), v179, v180, v180, l1)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v182 != 0 {
		v1912 = v182
		goto L3
	} else {
		goto L75
	}
L75:
	;
	v185 = F_palloc0(m, int32(36))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(15)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+4)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+8)) = v191
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+12)) = uint8(v193)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+13)) = uint8(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+16)) = v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+20)) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+24)) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+28)) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+32)) = v205
	v1912 = v185
	goto L3
L77:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v210 == int32(0) {
		v1912 = v208
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v213 = m.G0
	v215 = v213 - int32(16)
	m.G0 = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v218 = F_get_func_support(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	m.G0 = v215 + int32(16)
	v1912 = v233
	goto L3
L80:
	;
	if v218 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+4)) = int32(464)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v215)+8)) = v222
	v229 = F_OidFunctionCall1Coll(m, v218, int32(0), base.I64_extend_i32_u(v215+int32(4)))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v233 = v208
	goto L79
L84:
	;
	v231 = base.I32_wrap_i64(v229)
	if v231 != 0 {
		v233 = v231
		goto L79
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v249 = int32(1)
	v251 = F_simplify_function(m, v241, v242, int32(-1), v244, v245, v16+int32(144), int32(0), v249, v249, l1)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v251 != 0 {
		v1912 = v251
		goto L3
	} else {
		goto L88
	}
L88:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	switch v254 - int32(85) {
	case 0, 6:
		goto L90
	default:
		goto L89
	}
L89:
	;
	v298 = F_palloc0(m, int32(36))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L110
	}
L90:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v259 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	if v290 != 0 {
		v1912 = v290
		goto L3
	} else {
		goto L109
	}
L92:
	;
	v287 = F_negate_clause(m, v284)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L108
	}
L93:
	;
	if v258 == int32(0) {
		v290 = v3
		goto L91
	} else {
		goto L101
	}
L94:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v262 != int32(7) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v259)+24))
	if v254 == int32(91) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if v265 == int64(0) {
		v284 = v258
		goto L92
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v265 != int64(0) {
		v284 = v258
		goto L92
	} else {
		goto L100
	}
L99:
	;
	v290 = v258
	goto L91
L100:
	;
	v290 = v258
	goto L91
L101:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v274 != int32(7) {
		v290 = v3
		goto L91
	} else {
		goto L102
	}
L102:
	;
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v258)+24))
	if v254 == int32(91) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v277 == int64(0) {
		v284 = v259
		goto L92
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v277 == int64(0) {
		v290 = v259
		goto L91
	} else {
		goto L107
	}
L106:
	;
	v290 = v259
	goto L91
L107:
	;
	v284 = v259
	goto L92
L108:
	;
	v290 = v287
	goto L91
L109:
	;
	goto L89
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = int32(17)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+4)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+8)) = v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+12)) = v306
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v298)+16)) = uint8(v308)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+20)) = v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+28)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v298)+24)) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+32)) = v315
	v1912 = v298
	goto L3
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v319
	if v319 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	if v381&int32(1) == int32(0) {
		goto L144
	} else {
		goto L145
	}
L113:
	;
	if v380&int32(1) != 0 {
		goto L135
	} else {
		goto L136
	}
L114:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v322 <= int32(0) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	goto L116
L116:
	;
	v402 = int32(0)
	v404 = F_makeBoolConst(m, v402, v402)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L134
	}
L117:
	;
	if v383 != 0 {
		goto L112
	} else {
		goto L132
	}
L118:
	;
	v379 = int32(1)
	v380 = v3
	v381 = v3
	v383 = v3
	goto L117
L119:
	;
	goto L120
L120:
	;
	v330 = int32(0)
	v333 = int32(1)
	v334 = v3
	v335 = v3
	v337 = v3
	goto L121
L121:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341+v330<<(uint(int32(2))%32))))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if v346 == int32(7) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v379 = v365
	v380 = v366
	v381 = v367
	v383 = v369
	goto L117
L123:
	;
	v371 = v330 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v371 < v372 {
		v330 = v371
		v333 = v365
		v334 = v366
		v335 = v367
		v337 = v369
		goto L121
	} else {
		goto L131
	}
L124:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+32)))
	v365 = v333 & v349
	v366 = (v334 | v349) & int32(1)
	v367 = v335
	v369 = v337
	goto L123
L125:
	;
	goto L126
L126:
	;
	v354 = int32(1)
	v355 = int32(0)
	if v335&v354 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v365 = v355
	v366 = v334
	v367 = int32(1)
	v369 = v354
	goto L123
L128:
	;
	goto L129
L129:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v361 = F_expr_is_nonnullable(m, v359, v345, int32(1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v365 = v355
	v366 = v334
	v367 = v361 ^ int32(1)
	v369 = v354
	goto L123
L131:
	;
	goto L122
L132:
	;
	if v379 == int32(0) {
		goto L113
	} else {
		goto L133
	}
L133:
	;
	goto L116
L134:
	;
	v1912 = v404
	goto L3
L135:
	;
	v410 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	F_set_opfuncid(m, v24)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L139
	}
L138:
	;
	v1912 = v410
	goto L3
L139:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v421 = int32(0)
	v424 = F_simplify_function(m, v414, v415, int32(-1), v417, v418, v16+int32(144), v421, v421, v421, l1)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	if v424 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	v1883 = v428
	goto L6
L142:
	;
	goto L143
L143:
	;
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v424)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v424)+24)) = base.I64_extend_i32_u(base.B2i32(v429 == int64(0)))
	v1912 = v424
	goto L3
L144:
	;
	if v380&int32(1) != 0 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	if v380&int32(1) == int32(0) {
		v1883 = v319
		goto L6
	} else {
		goto L154
	}
L147:
	;
	v442 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v445 = F_palloc0(m, int32(36))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L151
	}
L150:
	;
	v1912 = v442
	goto L3
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445))) = int32(17)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+4)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+12)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v451
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v445)+16)) = uint8(v455)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+20)) = v457
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+28)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v445)+24)) = v459
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+32)) = v462
	v464 = F_negate_clause(m, v445)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v466 = F_eval_const_expressions_mutator(m, v464, l1)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v1912 = v466
	goto L3
L154:
	;
	v473 = F_palloc0(m, int32(20))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = int32(52)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	if v481 == int32(7) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v484 = int32(4)
	goto L158
L157:
	;
	v484 = int32(0)
	goto L158
L158:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v477+v484)))
	v487 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v473)+12)) = uint8(v487)
	*(*int32)(unsafe.Add(mBase, uint32(v473)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v473)+4)) = v486
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v473)+16)) = v492
	v494 = F_eval_const_expressions_mutator(m, v473, l1)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v1912 = v494
	goto L3
L160:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v497)+8))
	v558 = F_func_volatile(m, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L178
	}
L161:
	;
	F_set_opfuncid(m, v497)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L177
	}
L162:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v497)+28))
	if v499 == int32(0) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v502 <= int32(0) {
		goto L161
	} else {
		goto L164
	}
L164:
	;
	v505 = int32(0)
	if v505 < v502 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v508 = v502
	goto L167
L166:
	;
	v508 = v505
	goto L167
L167:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
	v511 = int32(0)
	v516 = v3
	goto L168
L168:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v509+v511<<(uint(int32(2))%32))))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	if v529 == int32(7) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	F_set_opfuncid(m, v497)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L175
	}
L170:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+32)))
	if v532 != 0 {
		goto L11
	} else {
		goto L173
	}
L171:
	;
	v533 = int32(1)
	goto L172
L172:
	;
	v535 = v511 + int32(1)
	if v535 != v508 {
		v511 = v535
		v516 = v533
		goto L168
	} else {
		goto L174
	}
L173:
	;
	v533 = v516
	goto L172
L174:
	;
	goto L169
L175:
	;
	if v533&int32(1) != 0 {
		v1912 = v497
		goto L3
	} else {
		goto L176
	}
L176:
	;
	goto L160
L177:
	;
	goto L160
L178:
	;
	if v558 != int32(105) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	if v558 != int32(115) {
		v1912 = v497
		goto L3
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v569 = F_exprType(m, v497)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v564&int32(1) == int32(0) {
		v1912 = v497
		goto L3
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v571 = F_exprTypmod(m, v497)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v573 = F_exprCollation(m, v497)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v575 = F_evaluate_expr(m, v497, v569, v571, v573)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v1912 = v575
	goto L3
L188:
	;
	F_set_opfuncid(m, v578)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v584 = F_expression_tree_walker_impl(m, v578, int32(920), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	if v584 != 0 {
		v1912 = v578
		goto L3
	} else {
		goto L191
	}
L191:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v578)+8))
	v587 = F_func_volatile(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	if v587 != int32(105) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if v587 != int32(115) {
		v1912 = v578
		goto L3
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v598 = F_exprType(m, v578)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v593&int32(1) == int32(0) {
		v1912 = v578
		goto L3
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v600 = F_exprTypmod(m, v578)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	v602 = F_exprCollation(m, v578)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v604 = F_evaluate_expr(m, v578, v598, v600, v602)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v1912 = v604
	goto L3
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L300
	}
L203:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)+12))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v860)))
	v862 = F_eval_const_expressions_mutator(m, v861, l1)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L298
	}
L204:
	;
	v735 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+144)) = uint8(v735)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+140)) = uint8(v735)
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v740 = F_list_copy(m, v739)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L253
	}
L205:
	;
	v607 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+144)) = uint8(v607)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+140)) = uint8(v607)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v612 = F_list_copy(m, v611)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L207
	}
L206:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+140)))
	if v705 == int32(1) {
		goto L235
	} else {
		goto L236
	}
L207:
	;
	if v612 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v618 = v612
	v623 = v3
	goto L211
L209:
	;
	v683 = v3
	goto L210
L210:
	;
	v704 = v683
	goto L206
L211:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v618)+12))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	v633 = F_list_delete_first(m, v618)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L213
	}
L212:
	;
	v683 = v676
	goto L210
L213:
	;
	if v632 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	if v674 != 0 {
		v618 = v674
		v623 = v676
		goto L211
	} else {
		goto L234
	}
L215:
	;
	v648 = F_eval_const_expressions_mutator(m, v632, l1)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L222
	}
L216:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v632)))
	if v637 != int32(21) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	if v640 != int32(1) {
		goto L215
	} else {
		goto L218
	}
L218:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v632)+8))
	v644 = F_list_concat_copy(m, v643, v633)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	F_list_free(m, v633)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v674 = v644
	v676 = v623
	goto L214
L221:
	;
	v672 = F_lappend(m, v623, v648)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L233
	}
L222:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	v652 = v650 - int32(7)
	if v652 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	if v652 != int32(14) {
		goto L221
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648)+32)))
	if v661 == int32(1) {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v648)+4))
	if v655 != int32(1) {
		goto L221
	} else {
		goto L227
	}
L227:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v648)+8))
	v659 = F_list_concat_copy(m, v658, v633)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v674 = v659
	v676 = v623
	goto L214
L229:
	;
	v664 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(144)))) = uint8(v664)
	v674 = v633
	v676 = v623
	goto L214
L230:
	;
	goto L231
L231:
	;
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v648)+24))
	if v666 == int64(0) {
		v674 = v633
		v676 = v623
		goto L214
	} else {
		goto L232
	}
L232:
	;
	v669 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(140)))) = uint8(v669)
	v704 = int32(0)
	goto L206
L233:
	;
	v674 = v633
	v676 = v672
	goto L214
L234:
	;
	goto L212
L235:
	;
	v710 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+144)))
	if v712 == int32(1) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1912 = v710
	goto L3
L239:
	;
	v717 = F_makeBoolConst(m, int32(0), int32(1))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	v721 = v704
	goto L241
L241:
	;
	if v721 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	v719 = F_lappend(m, v704, v717)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v721 = v719
	goto L241
L244:
	;
	v724 = int32(0)
	v726 = F_makeBoolConst(m, v724, v724)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v721)+4))
	if v728 == int32(1) {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v1912 = v726
	goto L3
L248:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v721)+12))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v1912 = v732
	goto L3
L249:
	;
	goto L250
L250:
	;
	v733 = F_make_orclause(m, v721)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1912 = v733
	goto L3
L252:
	;
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+140)))
	if v829 == int32(1) {
		goto L281
	} else {
		goto L282
	}
L253:
	;
	if v740 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v746 = v740
	v751 = v3
	goto L257
L255:
	;
	v807 = v3
	goto L256
L256:
	;
	v828 = v807
	goto L252
L257:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v746)+12))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	v761 = F_list_delete_first(m, v746)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L259
	}
L258:
	;
	v807 = v800
	goto L256
L259:
	;
	if v760 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	if v798 != 0 {
		v746 = v798
		v751 = v800
		goto L257
	} else {
		goto L280
	}
L261:
	;
	v774 = F_eval_const_expressions_mutator(m, v760, l1)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L268
	}
L262:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	if v765 != int32(21) {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v768 != 0 {
		goto L261
	} else {
		goto L264
	}
L264:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v760)+8))
	v770 = F_list_concat_copy(m, v769, v761)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	F_list_free(m, v761)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v798 = v770
	v800 = v751
	goto L260
L267:
	;
	v796 = F_lappend(m, v751, v774)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L279
	}
L268:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	v778 = v776 - int32(7)
	if v778 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	if v778 != int32(14) {
		goto L267
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v774)+32)))
	if v785 == int32(1) {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v774)+4))
	if v781 != 0 {
		goto L267
	} else {
		goto L273
	}
L273:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v774)+8))
	v783 = F_list_concat_copy(m, v782, v761)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	v798 = v783
	v800 = v751
	goto L260
L275:
	;
	v788 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(144)))) = uint8(v788)
	v798 = v761
	v800 = v751
	goto L260
L276:
	;
	goto L277
L277:
	;
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v774)+24))
	if v790 != int64(0) {
		v798 = v761
		v800 = v751
		goto L260
	} else {
		goto L278
	}
L278:
	;
	v793 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(140)))) = uint8(v793)
	v828 = int32(0)
	goto L252
L279:
	;
	v798 = v761
	v800 = v796
	goto L260
L280:
	;
	goto L258
L281:
	;
	v832 = int32(0)
	v834 = F_makeBoolConst(m, v832, v832)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+144)))
	if v836 == int32(1) {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	v1912 = v834
	goto L3
L285:
	;
	v841 = F_makeBoolConst(m, int32(0), int32(1))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L288
	}
L286:
	;
	v845 = v828
	goto L287
L287:
	;
	if v845 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	v843 = F_lappend(m, v828, v841)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v845 = v843
	goto L287
L290:
	;
	v850 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	if v852 == int32(1) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v1912 = v850
	goto L3
L294:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v845)+12))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v855)))
	v1912 = v856
	goto L3
L295:
	;
	goto L296
L296:
	;
	v857 = F_make_andclause(m, v845)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1912 = v857
	goto L3
L298:
	;
	v864 = F_negate_clause(m, v862)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	v1912 = v864
	goto L3
L300:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v870
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_3), v16+int32(16))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(3342), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	if v884 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v884)))
	if v886 == int32(7) {
		v1912 = v884
		goto L3
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v889 = F_eval_const_expressions_mutator(m, v882, l1)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L308
	}
L307:
	;
	goto L306
L308:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v892 = F_copyObjectImpl(m, v891)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v894 = F_makeJsonValueExpr(m, v889, v884, v892)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v1912 = v894
	goto L3
L311:
	;
	v1876 = int32(12)
	goto L9
L312:
	;
	goto L313
L313:
	;
	v901 = F_palloc0(m, int32(40))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v901))) = int32(45)
	v905 = *(*int64)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v901)+8)) = v905
	v907 = *(*int64)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v901)+16)) = v907
	v909 = *(*int64)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v901)+24)) = v909
	v911 = *(*int64)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v901)+32)) = v911
	v913 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	*(*int64)(unsafe.Add(mBase, uint32(v901))) = v913
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v916 = F_eval_const_expressions_mutator(m, v915, l1)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v901)+8)) = v916
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v920 = F_eval_const_expressions_mutator(m, v919, l1)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v901)+12)) = v920
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v927 = F_eval_const_expressions_mutator(m, v926, l1)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v901)+16)) = v927
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v923
	v1912 = v901
	goto L3
L318:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v940 = F_applyRelabelType(m, v932, v934, v935, v936, v937, v938, int32(1))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	v1912 = v940
	goto L3
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v948
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v952 = F_exprType(m, v951)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_getTypeOutputInfo(m, v952, v16+int32(140), v16+int32(139))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_getTypeInputInfo(m, v960, v16+int32(132), v16+int32(128))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v16)+140))
	v970 = int32(0)
	v973 = v16 + int32(144)
	v975 = int32(1)
	v977 = F_simplify_function(m, v967, int32(2275), int32(-1), v970, v970, v973, v970, v975, v975, l1)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	if v977 != 0 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v977
	v982 = int32(0)
	v984 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+128)))
	v987 = F_makeConst(m, int32(26), int32(-1), v982, int32(4), v984, v982, int32(1))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L1
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1026 = F_palloc0(m, int32(24))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L333
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v987
	v992 = int32(0)
	v997 = F_makeConst(m, int32(23), int32(-1), v992, int32(4), int64(-1), v992, int32(1))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v997
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v16)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v1001
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v1003
	v1011 = F_list_make3_impl(m, v16+int32(40), v16+int32(36), v16+int32(32))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v1011
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v16)+132))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1018 = int32(0)
	v1022 = F_simplify_function(m, v1014, v1015, int32(-1), v1017, v1018, v973, v1018, v1018, int32(1), l1)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	if v1022 != 0 {
		v1912 = v1022
		goto L3
	} else {
		goto L332
	}
L332:
	;
	goto L327
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1026))) = int32(28)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+12))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1031)))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+4)) = v1032
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+8)) = v1034
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+12)) = v1036
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+16)) = v1038
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+20)) = v1040
	v1912 = v1026
	goto L3
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1043))) = int32(29)
	v1047 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	*(*int64)(unsafe.Add(mBase, uint32(v1043))) = v1047
	v1049 = *(*int64)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1043)+8)) = v1049
	v1051 = *(*int64)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1043)+16)) = v1051
	v1053 = *(*int64)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1043)+24)) = v1053
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	v1056 = F_eval_const_expressions_mutator(m, v1055, l1)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+4)) = v1056
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+8))
	v1063 = F_eval_const_expressions_mutator(m, v1062, l1)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+8)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v1059
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	if v1067 == int32(0) {
		v1912 = v1043
		goto L3
	} else {
		goto L337
	}
L337:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	if v1070 != int32(7) {
		v1912 = v1043
		goto L3
	} else {
		goto L338
	}
L338:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+8))
	if v1073 == int32(0) {
		v1912 = v1043
		goto L3
	} else {
		goto L339
	}
L339:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1073)))
	if v1076 == int32(55) {
		v1912 = v1043
		goto L3
	} else {
		goto L340
	}
L340:
	;
	v1080 = F_contain_mutable_functions_walker(m, v1073, int32(0))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	if v1080 != 0 {
		v1912 = v1043
		goto L3
	} else {
		goto L342
	}
L342:
	;
	v1082 = F_exprType(m, v1043)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	v1084 = F_exprTypmod(m, v1043)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	v1086 = F_exprCollation(m, v1043)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	v1088 = F_evaluate_expr(m, v1043, v1082, v1084, v1086)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1912 = v1088
	goto L3
L347:
	;
	v1093 = F_exprType(m, v1091)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	v1095 = F_exprTypmod(m, v1091)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1101 = F_applyRelabelType(m, v1091, v1093, v1095, v1097, int32(2), v1099, int32(1))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	v1912 = v1101
	goto L3
L351:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v1113
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v1117 == int32(0) {
		v1183 = v3
		goto L357
	} else {
		goto L358
	}
L352:
	;
	v1113 = int32(0)
	v1114 = v1104
	goto L351
L353:
	;
	if v1104 == int32(0) {
		goto L352
	} else {
		goto L354
	}
L354:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1104)))
	if v1108 != int32(7) {
		goto L352
	} else {
		goto L355
	}
L355:
	;
	v1113 = v1104
	v1114 = int32(0)
	goto L351
L356:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1204)))
	v1206 = F_eval_const_expressions_mutator(m, v1205, l1)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L373
	}
L357:
	;
	v1198 = v1183
	v1204 = v24 + int32(20)
	goto L356
L358:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+4))
	if v1120 <= int32(0) {
		v1183 = v3
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1126 = v3
	v1130 = v3
	goto L360
L360:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+12))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1136+v1126<<(uint(int32(2))%32))))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+4))
	v1142 = F_eval_const_expressions_mutator(m, v1141, l1)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L364
	}
L361:
	;
	v1183 = v1170
	goto L357
L362:
	;
	v1173 = v1126 + int32(1)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+4))
	if v1173 < v1174 {
		v1126 = v1173
		v1130 = v1170
		goto L360
	} else {
		goto L372
	}
L363:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+8))
	v1156 = F_eval_const_expressions_mutator(m, v1155, l1)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L369
	}
L364:
	;
	if v1142 == int32(0) {
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1142)))
	if v1146 != int32(7) {
		goto L363
	} else {
		goto L366
	}
L366:
	;
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1142)+32)))
	if v1149 != 0 {
		v1170 = v1130
		goto L362
	} else {
		goto L367
	}
L367:
	;
	v1150 = *(*int64)(unsafe.Add(mBase, uint32(v1142)+24))
	if v1150 == int64(0) {
		v1170 = v1130
		goto L362
	} else {
		goto L368
	}
L368:
	;
	v1198 = v1130
	v1204 = v1140 + int32(8)
	goto L356
L369:
	;
	v1159 = F_palloc0(m, int32(16))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1159)+8)) = v1156
	*(*int32)(unsafe.Add(mBase, uint32(v1159)+4)) = v1142
	*(*int32)(unsafe.Add(mBase, uint32(v1159))) = int32(33)
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1159)+12)) = v1165
	v1167 = F_lappend(m, v1130, v1159)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	v1170 = v1167
	goto L362
L372:
	;
	goto L361
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v1115
	if v1198 == int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1912 = v1206
	goto L3
L375:
	;
	goto L376
L376:
	;
	v1212 = F_palloc0(m, int32(28))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1212))) = int32(32)
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+4)) = v1216
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+20)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+16)) = v1198
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+12)) = v1114
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+8)) = v1218
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+24)) = v1223
	v1912 = v1212
	goto L3
L378:
	;
	v1228 = F_copyObjectImpl(m, v1225)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1912 = v1228
	goto L3
L380:
	;
	v1235 = F_expression_tree_walker_impl(m, v1231, int32(920), int32(0))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	if v1235 != 0 {
		v1912 = v1231
		goto L3
	} else {
		goto L382
	}
L382:
	;
	v1237 = F_exprType(m, v1231)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v1239 = F_exprTypmod(m, v1231)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	v1241 = F_exprCollation(m, v1231)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	v1243 = F_evaluate_expr(m, v1231, v1237, v1239, v1241)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v1912 = v1243
	goto L3
L387:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+4))
	if v1321 == int32(1) {
		goto L411
	} else {
		goto L412
	}
L388:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	if v1246 <= int32(0) {
		v1295 = v3
		goto L391
	} else {
		goto L392
	}
L389:
	;
	goto L390
L390:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v1319 = F_makeNullConst(m, v1316, int32(-1), v1318)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L410
	}
L391:
	;
	if v1295 != 0 {
		goto L387
	} else {
		goto L409
	}
L392:
	;
	v1252 = v3
	v1254 = v3
	goto L393
L393:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+12))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1262+v1252<<(uint(int32(2))%32))))
	v1267 = F_eval_const_expressions_mutator(m, v1266, l1)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L1
	} else {
		goto L398
	}
L394:
	;
	v1295 = v1285
	goto L391
L395:
	;
	v1287 = v1252 + int32(1)
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	if v1287 < v1288 {
		v1252 = v1287
		v1254 = v1285
		goto L393
	} else {
		goto L408
	}
L396:
	;
	v1283 = F_lappend(m, v1254, v1267)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L407
	}
L397:
	;
	if v1254 == int32(0) {
		v1912 = v1267
		goto L3
	} else {
		goto L405
	}
L398:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1267)))
	if v1269 == int32(7) {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1267)+32)))
	if v1272 != 0 {
		v1285 = v1254
		goto L395
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1275 = F_expr_is_nonnullable(m, v1273, v1267, int32(1))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L1
	} else {
		goto L403
	}
L402:
	;
	goto L397
L403:
	;
	if v1275 == int32(0) {
		goto L396
	} else {
		goto L404
	}
L404:
	;
	goto L397
L405:
	;
	v1281 = F_lappend(m, v1254, v1267)
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v1295 = v1281
	goto L391
L407:
	;
	v1285 = v1283
	goto L395
L408:
	;
	goto L394
L409:
	;
	goto L390
L410:
	;
	v1912 = v1319
	goto L3
L411:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+12))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1324)))
	v1912 = v1325
	goto L3
L412:
	;
	goto L413
L413:
	;
	v1327 = F_palloc0(m, int32(20))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1327))) = int32(38)
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+4)) = v1331
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+12)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+8)) = v1333
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+16)) = v1336
	v1912 = v1327
	goto L3
L415:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1344 = F_evaluate_expr(m, v24, v1341, v1342, int32(0))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v1912 = v1344
	goto L3
L417:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1349 != 0 {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1358 = F_expression_tree_walker_impl(m, v1347, int32(920), int32(0))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L1
	} else {
		goto L422
	}
L419:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+4))
	if base.Ui32(v1350-int32(3)) < base.Ui32(int32(5)) {
		goto L418
	} else {
		goto L420
	}
L420:
	;
	if v1350 != 0 {
		v1912 = v1347
		goto L3
	} else {
		goto L421
	}
L421:
	;
	goto L418
L422:
	;
	if v1358 != 0 {
		v1912 = v1347
		goto L3
	} else {
		goto L423
	}
L423:
	;
	v1360 = F_exprType(m, v1347)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v1362 = F_exprTypmod(m, v1347)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	v1364 = F_exprCollation(m, v1347)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v1366 = F_evaluate_expr(m, v1347, v1360, v1362, v1364)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	v1912 = v1366
	goto L3
L428:
	;
	v1427 = F_palloc0(m, int32(24))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L1
	} else {
		goto L450
	}
L429:
	;
	if v1369 == int32(0) {
		goto L428
	} else {
		goto L430
	}
L430:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	if v1373 == int32(6) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1369)+8)))
	if v1376 != 0 {
		goto L428
	} else {
		goto L434
	}
L432:
	;
	v1386 = v1373
	goto L433
L433:
	;
	if v1386 != int32(36) {
		goto L428
	} else {
		goto L438
	}
L434:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+28))
	if v1377 != 0 {
		goto L428
	} else {
		goto L435
	}
L435:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+12))
	v1379 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+8)))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v1383 = F_rowtype_field_matches(m, v1378, v1379, v1380, v1381, v1382)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	if v1383 != 0 {
		goto L10
	} else {
		goto L437
	}
L437:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	v1386 = v1385
	goto L433
L438:
	;
	v1389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+8)))
	if v1389 <= int32(0) {
		goto L428
	} else {
		goto L439
	}
L439:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+4))
	if v1392 == int32(0) {
		goto L428
	} else {
		goto L440
	}
L440:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+4))
	if v1395 < v1389 {
		goto L428
	} else {
		goto L441
	}
L441:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+12))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1397+v1389<<(uint(int32(2))%32)-int32(4))))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+8))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v1408 = F_rowtype_field_matches(m, v1404, v1389, v1405, v1406, v1407)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	if v1408 == int32(0) {
		goto L428
	} else {
		goto L443
	}
L443:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1413 = F_exprType(m, v1403)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	if v1412 != v1413 {
		goto L428
	} else {
		goto L445
	}
L445:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v1417 = F_exprTypmod(m, v1403)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L1
	} else {
		goto L446
	}
L446:
	;
	if v1416 != v1417 {
		goto L428
	} else {
		goto L447
	}
L447:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v1421 = F_exprCollation(m, v1403)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	if v1420 == v1421 {
		v1912 = v1403
		goto L3
	} else {
		goto L449
	}
L449:
	;
	goto L428
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1427)+4)) = v1369
	*(*int32)(unsafe.Add(mBase, uint32(v1427))) = int32(25)
	v1432 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1427)+8)) = uint16(v1432)
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1427)+12)) = v1434
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1427)+16)) = v1436
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1427)+20)) = v1438
	if v1369 == int32(0) {
		v1912 = v1427
		goto L3
	} else {
		goto L451
	}
L451:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	if v1442 != int32(7) {
		v1912 = v1427
		goto L3
	} else {
		goto L452
	}
L452:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+4))
	v1446 = F_rowtype_field_matches(m, v1445, v1432, v1434, v1436, v1438)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	if v1446 == int32(0) {
		v1912 = v1427
		goto L3
	} else {
		goto L454
	}
L454:
	;
	v1450 = F_exprType(m, v1427)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v1452 = F_exprTypmod(m, v1427)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	v1454 = F_exprCollation(m, v1427)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v1456 = F_evaluate_expr(m, v1427, v1450, v1452, v1454)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v1912 = v1456
	goto L3
L459:
	;
	v1642 = F_palloc0(m, int32(20))
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L1
	} else {
		goto L521
	}
L460:
	;
	v1461 = int32(0)
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+12)))
	if base.B2i32(v1459 == v1461)|base.B2i32(v1463 != int32(1)) == v1461 {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1459)))
	if v1469 != int32(36) {
		goto L459
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	if v1463|base.B2i32(v1459 == int32(0)) != 0 {
		goto L459
	} else {
		goto L500
	}
L464:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+4))
	if v1472 != 0 {
		goto L466
	} else {
		goto L467
	}
L465:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1548)+4))
	if v1572 == int32(1) {
		goto L496
	} else {
		goto L497
	}
L466:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+4))
	if int32(0) < v1473 {
		goto L469
	} else {
		goto L470
	}
L467:
	;
	goto L468
L468:
	;
	v1570 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L1
	} else {
		goto L495
	}
L469:
	;
	v1480 = int32(0)
	v1483 = v3
	goto L472
L470:
	;
	v1548 = v3
	goto L471
L471:
	;
	if v1548 != 0 {
		goto L465
	} else {
		goto L494
	}
L472:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+12))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1490+v1480<<(uint(int32(2))%32))))
	if v1494 == int32(0) {
		goto L475
	} else {
		goto L476
	}
L473:
	;
	v1548 = v1537
	goto L471
L474:
	;
	v1539 = v1480 + int32(1)
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+4))
	if v1539 < v1540 {
		v1480 = v1539
		v1483 = v1537
		goto L472
	} else {
		goto L493
	}
L475:
	;
	v1522 = F_palloc0(m, int32(20))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L1
	} else {
		goto L491
	}
L476:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1494)))
	if v1497 == int32(7) {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1494)+32)))
	if v1501 == int32(1) {
		goto L481
	} else {
		goto L482
	}
L478:
	;
	goto L479
L479:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1512 = F_expr_is_nonnullable(m, v1510, v1494, int32(1))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L487
	}
L480:
	;
	v1506 = int32(0)
	v1508 = F_makeBoolConst(m, v1506, v1506)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L1
	} else {
		goto L486
	}
L481:
	;
	if v1500 == int32(1) {
		goto L480
	} else {
		goto L484
	}
L482:
	;
	goto L483
L483:
	;
	if v1500 != 0 {
		v1537 = v1483
		goto L474
	} else {
		goto L485
	}
L484:
	;
	v1537 = v1483
	goto L474
L485:
	;
	goto L480
L486:
	;
	v1912 = v1508
	goto L3
L487:
	;
	if v1512 == int32(0) {
		goto L475
	} else {
		goto L488
	}
L488:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v1516 != 0 {
		v1537 = v1483
		goto L474
	} else {
		goto L489
	}
L489:
	;
	v1517 = int32(0)
	v1519 = F_makeBoolConst(m, v1517, v1517)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v1912 = v1519
	goto L3
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1522)+4)) = v1494
	*(*int32)(unsafe.Add(mBase, uint32(v1522))) = int32(52)
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v1528 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1522)+12)) = uint8(v1528)
	*(*int32)(unsafe.Add(mBase, uint32(v1522)+8)) = v1527
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1522)+16)) = v1531
	v1533 = F_lappend(m, v1483, v1522)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	v1537 = v1533
	goto L474
L493:
	;
	goto L473
L494:
	;
	goto L468
L495:
	;
	v1912 = v1570
	goto L3
L496:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1548)+12))
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1575)))
	v1912 = v1576
	goto L3
L497:
	;
	goto L498
L498:
	;
	v1577 = F_make_andclause(m, v1548)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	v1912 = v1577
	goto L3
L500:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1459)))
	if v1582 == int32(7) {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	switch v1585 {
	case 0:
		goto L505
	case 1:
		goto L507
	default:
		goto L506
	}
L502:
	;
	goto L503
L503:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1614 = F_expr_is_nonnullable(m, v1612, v1459, int32(1))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L1
	} else {
		goto L512
	}
L504:
	;
	v1610 = F_makeBoolConst(m, v1606&int32(1), int32(0))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L1
	} else {
		goto L511
	}
L505:
	;
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1459)+32)))
	v1606 = v1605
	goto L504
L506:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L1
	} else {
		goto L508
	}
L507:
	;
	v1586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1459)+32)))
	v1606 = v1586 ^ int32(1)
	goto L504
L508:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v1593
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_4), v16+int32(48))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(4050), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L511:
	;
	v1912 = v1610
	goto L3
L512:
	;
	if v1614 == int32(0) {
		goto L459
	} else {
		goto L513
	}
L513:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	switch v1619 {
	case 0:
		v1637 = int32(0)
		goto L514
	case 1:
		goto L515
	default:
		goto L516
	}
L514:
	;
	v1639 = F_makeBoolConst(m, v1637, int32(0))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L1
	} else {
		goto L520
	}
L515:
	;
	v1637 = int32(1)
	goto L514
L516:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v1624
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_4), v16-int32(-64))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(4073), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L520:
	;
	v1912 = v1639
	goto L3
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1642)+4)) = v1459
	*(*int32)(unsafe.Add(mBase, uint32(v1642))) = int32(52)
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1642)+8)) = v1647
	v1649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1642)+12)) = uint8(v1649)
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1642)+16)) = v1651
	v1912 = v1642
	goto L3
L522:
	;
	v1744 = F_palloc0(m, int32(16))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L557
	}
L523:
	;
	if v1654 == int32(0) {
		goto L522
	} else {
		goto L524
	}
L524:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1654)))
	if v1658 == int32(7) {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	switch v1661 {
	case 0:
		goto L536
	case 1:
		goto L535
	case 2:
		goto L534
	case 3:
		goto L533
	case 4:
		goto L532
	case 5:
		goto L531
	default:
		goto L530
	}
L526:
	;
	goto L527
L527:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1712 = F_expr_is_nonnullable(m, v1710, v1654, int32(1))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L1
	} else {
		goto L545
	}
L528:
	;
	v1708 = F_makeBoolConst(m, v1704&int32(1), int32(0))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L1
	} else {
		goto L544
	}
L529:
	;
	v1701 = *(*int64)(unsafe.Add(mBase, uint32(v1654)+24))
	v1704 = base.B2i32(v1701 != int64(0))
	goto L528
L530:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L541
	}
L531:
	;
	v1682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654)+32)))
	v1704 = v1682 ^ int32(1)
	goto L528
L532:
	;
	v1681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654)+32)))
	v1704 = v1681
	goto L528
L533:
	;
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654)+32)))
	if v1677 != 0 {
		v1704 = int32(1)
		goto L528
	} else {
		goto L540
	}
L534:
	;
	v1672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654)+32)))
	if v1672 != 0 {
		v1704 = int32(0)
		goto L528
	} else {
		goto L539
	}
L535:
	;
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654)+32)))
	if v1667 != 0 {
		v1704 = int32(1)
		goto L528
	} else {
		goto L538
	}
L536:
	;
	v1662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654)+32)))
	if v1662 == int32(0) {
		goto L529
	} else {
		goto L537
	}
L537:
	;
	v1704 = int32(0)
	goto L528
L538:
	;
	v1668 = *(*int64)(unsafe.Add(mBase, uint32(v1654)+24))
	v1704 = base.B2i32(v1668 == int64(0))
	goto L528
L539:
	;
	v1673 = *(*int64)(unsafe.Add(mBase, uint32(v1654)+24))
	v1704 = base.B2i32(v1673 == int64(0))
	goto L528
L540:
	;
	v1678 = *(*int64)(unsafe.Add(mBase, uint32(v1654)+24))
	v1704 = base.B2i32(v1678 != int64(0))
	goto L528
L541:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v1689
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_5), v16+int32(80))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(_a_F_eval_const_expressions_mutator_6), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L544:
	;
	v1912 = v1708
	goto L3
L545:
	;
	if v1712 == int32(0) {
		goto L522
	} else {
		goto L546
	}
L546:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	switch v1716 {
	case 0, 3:
		v1912 = v1654
		goto L3
	case 1, 2:
		goto L550
	case 4:
		goto L549
	case 5:
		goto L548
	default:
		goto L547
	}
L547:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L1
	} else {
		goto L554
	}
L548:
	;
	v1725 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L1
	} else {
		goto L553
	}
L549:
	;
	v1719 = int32(0)
	v1721 = F_makeBoolConst(m, v1719, v1719)
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L1
	} else {
		goto L552
	}
L550:
	;
	v1717 = F_make_notclause(m, v1654)
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	v1912 = v1717
	goto L3
L552:
	;
	v1912 = v1721
	goto L3
L553:
	;
	v1912 = v1725
	goto L3
L554:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v1731
	F_errmsg_internal(m, int32(_a_F_eval_const_expressions_mutator_5), v16+int32(96))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	F_errfinish(m, int32(_a_F_eval_const_expressions_mutator_1), int32(_a_F_eval_const_expressions_mutator_7), int32(_a_F_eval_const_expressions_mutator_2))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1744)+4)) = v1654
	*(*int32)(unsafe.Add(mBase, uint32(v1744))) = int32(53)
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1744)+8)) = v1749
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1744)+12)) = v1751
	v1912 = v1744
	goto L3
L558:
	;
	v1756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1756 == int32(0) {
		goto L560
	} else {
		goto L561
	}
L559:
	;
	v1798 = F_palloc0(m, int32(28))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L575
	}
L560:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v1760 = F_DomainHasConstraints(m, v1759)
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L1
	} else {
		goto L563
	}
L561:
	;
	goto L562
L562:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1762 == int32(0) {
		goto L565
	} else {
		goto L566
	}
L563:
	;
	if v1760 != 0 {
		goto L559
	} else {
		goto L564
	}
L564:
	;
	goto L562
L565:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v1795 = F_applyRelabelType(m, v1754, v1789, v1790, v1791, v1792, v1793, int32(1))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L1
	} else {
		goto L574
	}
L566:
	;
	v1765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1765 != 0 {
		goto L565
	} else {
		goto L567
	}
L567:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if base.Ui32(int32(_a_F_eval_const_expressions_mutator_8)) <= base.Ui32(v1766) {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v1770 = F_palloc0(m, int32(12))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	goto L570
L570:
	;
	goto L565
L571:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1770))) = int64(352187318655)
	v1777 = F_GetSysCacheHashValue(m, int32(82), base.I64_extend_i32_u(v1766), int64(0))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1770)+8)) = v1777
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1762)+8))
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+68))
	v1782 = F_lappend(m, v1781, v1770)
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1762)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1784)+68)) = v1782
	goto L570
L574:
	;
	v1912 = v1795
	goto L3
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+4)) = v1754
	*(*int32)(unsafe.Add(mBase, uint32(v1798))) = int32(55)
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+8)) = v1803
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+12)) = v1805
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+16)) = v1807
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+20)) = v1809
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+24)) = v1811
	v1912 = v1798
	goto L3
L576:
	;
	v1814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1814 == int32(0) {
		goto L12
	} else {
		goto L577
	}
L577:
	;
	v1876 = int32(4)
	goto L9
L578:
	;
	v1822 = F_palloc0(m, int32(20))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822))) = int32(30)
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+8)) = v1826
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+12)) = v1828
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+16)) = v1830
	if v1819 == int32(0) {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+4)) = int32(0)
	v1912 = v1822
	goto L3
L581:
	;
	goto L582
L582:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1819)))
	if v1836 != int32(30) {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1848)))
	if v1849 != int32(7) {
		v1912 = v1822
		goto L3
	} else {
		goto L591
	}
L584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+4)) = v1819
	v1848 = v1819
	goto L583
L585:
	;
	goto L586
L586:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+4))
	if v1828 == int32(2) {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+12)) = v1843
	goto L589
L588:
	;
	goto L589
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+4)) = v1840
	if v1840 == int32(0) {
		v1912 = v1822
		goto L3
	} else {
		goto L590
	}
L590:
	;
	v1848 = v1840
	goto L583
L591:
	;
	v1852 = F_exprType(m, v1822)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	v1854 = F_exprTypmod(m, v1822)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	v1856 = F_exprCollation(m, v1822)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	v1858 = F_evaluate_expr(m, v1822, v1852, v1854, v1856)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	v1912 = v1858
	goto L3
L596:
	;
	v1912 = v1861
	goto L3
L597:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1870)+32)) = v1872
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1870)+24)) = v1874
	v1912 = v1870
	goto L3
L598:
	;
	if v1878 != 0 {
		v24 = v1878
		goto L7
	} else {
		goto L599
	}
L599:
	;
	goto L8
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1885))) = int32(18)
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+4)) = v1889
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+8)) = v1891
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+12)) = v1893
	v1895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1885)+16)) = uint8(v1895)
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+20)) = v1897
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+28)) = v1883
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+24)) = v1899
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+32)) = v1902
	v1912 = v1885
	goto L3
L601:
	;
	v1912 = v1908
	goto L3
}
func F_exec_dynquery_with_params(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v14 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
		v22 = F_AllocSetContextCreateInternal(m, v17, int32(_a_F_exec_dynquery_with_params_0), int32(0), int32(_a_F_exec_dynquery_with_params_1), int32(_a_F_exec_dynquery_with_params_2))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v22
			v27 = v22
			v34 = F_exec_eval_expr(m, l0, l1, v12+int32(30), v12+int32(24), v12+int32(20))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
				if v36 != int32(1) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					v40 = int32(_a_F_exec_dynquery_with_params_3)
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v44
					F_getTypeOutputInfo(m, v39, v12+int32(8), v12+int32(31))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						v53 = F_OidOutputFunctionCall(m, v52, v34)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v41
							v57 = F_MemoryContextStrdup(m, v27, v53)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
								if v59 != 0 {
									F_SPI_freetuptable(m, v59)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										if v64 != 0 {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
											F_MemoryContextReset(m, v65)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
												v70 = F_exec_eval_using_params(m, l0, l2)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
													v76 = m.G0
													v78 = v76 - int32(48)
													m.G0 = v78
													v80 = int32(0)
													v83 = v12 + int32(8)
													if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
														v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														if v90 == int32(0) {
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1546), int32(_a_F_exec_dynquery_with_params_6))
																	mBase = m.M
																	v173 = m.ExcPending
																	if v173 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
															v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
															*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
															v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
															v103 = int32(0)
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
															v105 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
															*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
															*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															if v119 != 0 {
																v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
																v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
																*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
															} else {
															}
															v125 = v78 + int32(8)
															F__SPI_prepare_plan(m, v57, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
																v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
																v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
																	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
																	*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
																	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
																	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
																	F_MemoryContextReset(m, v139)
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v78 + int32(48)
																		if v130 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																			mBase = m.M
																			v201 = m.ExcPending
																			if v201 != 0 {
																				return int32(0)
																			} else {
																				v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																				v204 = F_SPI_result_code_string(m, v203)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																					F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																					mBase = m.M
																					v210 = m.ExcPending
																					if v210 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																						mBase = m.M
																						v215 = m.ExcPending
																						if v215 != 0 {
																							return int32(0)
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			}
																		} else {
																			F_MemoryContextReset(m, v27)
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v12 + int32(32)
																				return v130
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1542), int32(_a_F_exec_dynquery_with_params_6))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
											v70 = F_exec_eval_using_params(m, l0, l2)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
												v76 = m.G0
												v78 = v76 - int32(48)
												m.G0 = v78
												v80 = int32(0)
												v83 = v12 + int32(8)
												if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													if v90 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1546), int32(_a_F_exec_dynquery_with_params_6))
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
														v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
														v103 = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
														v105 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
														*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														if v119 != 0 {
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
														} else {
														}
														v125 = v78 + int32(8)
														F__SPI_prepare_plan(m, v57, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
															v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
																*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
																*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
																F_MemoryContextReset(m, v139)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v78 + int32(48)
																	if v130 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																			v204 = F_SPI_result_code_string(m, v203)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																				mBase = m.M
																				v210 = m.ExcPending
																				if v210 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
																						return int32(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				}
																			}
																		}
																	} else {
																		F_MemoryContextReset(m, v27)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v130
																		}
																	}
																}
															}
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1542), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									if v64 != 0 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
										F_MemoryContextReset(m, v65)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
											v70 = F_exec_eval_using_params(m, l0, l2)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
												v76 = m.G0
												v78 = v76 - int32(48)
												m.G0 = v78
												v80 = int32(0)
												v83 = v12 + int32(8)
												if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													if v90 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1546), int32(_a_F_exec_dynquery_with_params_6))
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
														v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
														v103 = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
														v105 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
														*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														if v119 != 0 {
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
														} else {
														}
														v125 = v78 + int32(8)
														F__SPI_prepare_plan(m, v57, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
															v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
																*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
																*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
																F_MemoryContextReset(m, v139)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v78 + int32(48)
																	if v130 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																			v204 = F_SPI_result_code_string(m, v203)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																				mBase = m.M
																				v210 = m.ExcPending
																				if v210 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
																						return int32(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				}
																			}
																		}
																	} else {
																		F_MemoryContextReset(m, v27)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v130
																		}
																	}
																}
															}
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1542), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
										v70 = F_exec_eval_using_params(m, l0, l2)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
											v76 = m.G0
											v78 = v76 - int32(48)
											m.G0 = v78
											v80 = int32(0)
											v83 = v12 + int32(8)
											if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
												if v90 == int32(0) {
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1546), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
													v103 = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
													v105 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
													*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
													*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													if v119 != 0 {
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
													} else {
													}
													v125 = v78 + int32(8)
													F__SPI_prepare_plan(m, v57, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
														v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
															*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
															F_MemoryContextReset(m, v139)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																m.G0 = v78 + int32(48)
																if v130 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																		v204 = F_SPI_result_code_string(m, v203)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																			F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	}
																} else {
																	F_MemoryContextReset(m, v27)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v12 + int32(32)
																		return v130
																	}
																}
															}
														}
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1542), int32(_a_F_exec_dynquery_with_params_6))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
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
					}
				} else {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v188 = m.ExcPending
						if v188 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_exec_dynquery_with_params_13), int32(0))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_14), int32(_a_F_exec_dynquery_with_params_11))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v27 = v14
		v34 = F_exec_eval_expr(m, l0, l1, v12+int32(30), v12+int32(24), v12+int32(20))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
			if v36 != int32(1) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
				v40 = int32(_a_F_exec_dynquery_with_params_3)
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0]))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
				*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v44
				F_getTypeOutputInfo(m, v39, v12+int32(8), v12+int32(31))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v53 = F_OidOutputFunctionCall(m, v52, v34)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v41
						v57 = F_MemoryContextStrdup(m, v27, v53)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							if v59 != 0 {
								F_SPI_freetuptable(m, v59)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									if v64 != 0 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
										F_MemoryContextReset(m, v65)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
											v70 = F_exec_eval_using_params(m, l0, l2)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
												v76 = m.G0
												v78 = v76 - int32(48)
												m.G0 = v78
												v80 = int32(0)
												v83 = v12 + int32(8)
												if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													if v90 == int32(0) {
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1546), int32(_a_F_exec_dynquery_with_params_6))
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
														v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
														v103 = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
														v105 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
														*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
														*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														if v119 != 0 {
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
														} else {
														}
														v125 = v78 + int32(8)
														F__SPI_prepare_plan(m, v57, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
															v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
															v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
																*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
																*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
																F_MemoryContextReset(m, v139)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v78 + int32(48)
																	if v130 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																			v204 = F_SPI_result_code_string(m, v203)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																				F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																				mBase = m.M
																				v210 = m.ExcPending
																				if v210 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
																						return int32(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				}
																			}
																		}
																	} else {
																		F_MemoryContextReset(m, v27)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v12 + int32(32)
																			return v130
																		}
																	}
																}
															}
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1542), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
										v70 = F_exec_eval_using_params(m, l0, l2)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
											v76 = m.G0
											v78 = v76 - int32(48)
											m.G0 = v78
											v80 = int32(0)
											v83 = v12 + int32(8)
											if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
												if v90 == int32(0) {
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1546), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
													v103 = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
													v105 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
													*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
													*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													if v119 != 0 {
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
													} else {
													}
													v125 = v78 + int32(8)
													F__SPI_prepare_plan(m, v57, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
														v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
															*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
															F_MemoryContextReset(m, v139)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																m.G0 = v78 + int32(48)
																if v130 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																		v204 = F_SPI_result_code_string(m, v203)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																			F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	}
																} else {
																	F_MemoryContextReset(m, v27)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v12 + int32(32)
																		return v130
																	}
																}
															}
														}
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1542), int32(_a_F_exec_dynquery_with_params_6))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								if v64 != 0 {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
									F_MemoryContextReset(m, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
										v70 = F_exec_eval_using_params(m, l0, l2)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
											v76 = m.G0
											v78 = v76 - int32(48)
											m.G0 = v78
											v80 = int32(0)
											v83 = v12 + int32(8)
											if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
												if v90 == int32(0) {
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1546), int32(_a_F_exec_dynquery_with_params_6))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
													*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
													v103 = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
													v105 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
													*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
													*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
													*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													if v119 != 0 {
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
													} else {
													}
													v125 = v78 + int32(8)
													F__SPI_prepare_plan(m, v57, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
														v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
														v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
															*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
															*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
															F_MemoryContextReset(m, v139)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																m.G0 = v78 + int32(48)
																if v130 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																		v204 = F_SPI_result_code_string(m, v203)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																			F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	}
																} else {
																	F_MemoryContextReset(m, v27)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v12 + int32(32)
																		return v130
																	}
																}
															}
														}
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1542), int32(_a_F_exec_dynquery_with_params_6))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
									v70 = F_exec_eval_using_params(m, l0, l2)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
										v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v74)
										v76 = m.G0
										v78 = v76 - int32(48)
										m.G0 = v78
										v80 = int32(0)
										v83 = v12 + int32(8)
										if base.B2i32(v57 == v80)|base.B2i32(v83 == v80) == v80 {
											v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
											if v90 == int32(0) {
												*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = int32(-4)
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_4), int32(0))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1546), int32(_a_F_exec_dynquery_with_params_6))
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[3]))
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
												*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v95
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
												*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v100
												v103 = int32(0)
												*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2])) = v103
												v105 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v105
												*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v105
												*(*int64)(unsafe.Add(mBase, uint32(v78)+28)) = v105
												*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v105
												*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v103
												*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(569278163)
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v117
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
												if v119 != 0 {
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v120
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v122
												} else {
												}
												v125 = v78 + int32(8)
												F__SPI_prepare_plan(m, v57, v125)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
													v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
													v130 = F_SPI_cursor_open_internal(m, l3, v125, v128, v129)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int32(0)
													} else {
														v134 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[1]))
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
														*(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[0])) = v135
														*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = int32(0)
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
														F_MemoryContextReset(m, v139)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int32(0)
														} else {
															m.G0 = v78 + int32(48)
															if v130 == int32(0) {
																F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	v203 = *(*int32)(unsafe.Add(mBase, _c_F_exec_dynquery_with_params[2]))
																	v204 = F_SPI_result_code_string(m, v203)
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v204
																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
																		F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_8), v12)
																		mBase = m.M
																		v210 = m.ExcPending
																		if v210 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_10), int32(_a_F_exec_dynquery_with_params_11))
																			mBase = m.M
																			v215 = m.ExcPending
																			if v215 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																}
															} else {
																F_MemoryContextReset(m, v27)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v12 + int32(32)
																	return v130
																}
															}
														}
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(_a_F_exec_dynquery_with_params_12), int32(0))
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_exec_dynquery_with_params_5), int32(1542), int32(_a_F_exec_dynquery_with_params_6))
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
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
				}
			} else {
				F_errstart_cold(m, int32(21), int32(_a_F_exec_dynquery_with_params_7))
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v188 = m.ExcPending
					if v188 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_exec_dynquery_with_params_13), int32(0))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_exec_dynquery_with_params_9), int32(_a_F_exec_dynquery_with_params_14), int32(_a_F_exec_dynquery_with_params_11))
							mBase = m.M
							v197 = m.ExcPending
							if v197 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v6 = int32(0)
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v13 == int32(2) {
		v76 = l0
		v82 = v6
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return (v91 ^ v92) & int32(1)
L4:
	;
	v84 = m.T0[l4].(func(*base.Module, int32, int32, int32) int32)(m, l1, v76, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L5:
	;
	v16 = l0
	v19 = l3
	v22 = v6
	goto L6
L6:
	;
	v26 = v16
	goto L8
L7:
	;
	v76 = v72
	v82 = v22
	goto L4
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	switch v34 - int32(33) {
	case 0:
		goto L13
	default:
		goto L11
	case 5:
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v59 = int32(1)
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+2)))
	v66 = F_execute(m, v26+v60<<(uint(int32(3))%32), l1, l2, v19&v59, l4)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L19
	}
L12:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+2)))
	v56 = F_execute(m, v26+v52<<(uint(int32(3))%32), l1, l2, v19&int32(1), l4)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	v37 = int32(1)
	if v19&v37 == int32(0) {
		v91 = v37
		v92 = v22
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v42 = int32(1)
	v44 = v22 ^ v42
	F_check_stack_depth(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v48 = v26 - int32(8)
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48))))
	if v49 != int32(2) {
		v16 = v48
		v19 = v42
		v22 = v44
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v76 = v48
	v82 = v44
	goto L4
L17:
	;
	if v56 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v91 = int32(0)
	v92 = v22
	goto L3
L19:
	;
	if v66 != 0 {
		v91 = v59
		v92 = v22
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	v72 = v26 - int32(8)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72))))
	if v73 != int32(2) {
		v26 = v72
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L9
L23:
	;
	v91 = v84
	v92 = v82
	goto L3
}
func F_existsTimeLineHistory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v3 = m.G0
	v5 = v3 - int32(1136)
	m.G0 = v5
	if l0 != int32(1) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_existsTimeLineHistory[0])))
		if v10 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
			v15 = v5 + int32(48)
			v20 = F_pg_snprintf(m, v15, int32(64), int32(_a_F_existsTimeLineHistory_0), v5+int32(16))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v29 = F_RestoreArchivedFile(m, v5+int32(112), v15, int32(_a_F_existsTimeLineHistory_1), int64(0), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v44 = F_AllocateFile(m, v5+int32(112), int32(_a_F_existsTimeLineHistory_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						if v44 != 0 {
							v46 = F_FreeFile(m, v44)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v56 = int32(1)
								m.G0 = v5 + int32(1136)
								return v56
							}
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_existsTimeLineHistory[1]))
							if v50 != int32(44) {
								F_errstart_cold(m, int32(22), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(112)
										F_errmsg(m, int32(_a_F_existsTimeLineHistory_3), v5)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_existsTimeLineHistory_4), int32(252), int32(_a_F_existsTimeLineHistory_5))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v56 = int32(0)
								m.G0 = v5 + int32(1136)
								return v56
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = l0
			v38 = F_pg_snprintf(m, v5+int32(112), int32(1024), int32(_a_F_existsTimeLineHistory_6), v5+int32(32))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v44 = F_AllocateFile(m, v5+int32(112), int32(_a_F_existsTimeLineHistory_2))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					if v44 != 0 {
						v46 = F_FreeFile(m, v44)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v56 = int32(1)
							m.G0 = v5 + int32(1136)
							return v56
						}
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_existsTimeLineHistory[1]))
						if v50 != int32(44) {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(112)
									F_errmsg(m, int32(_a_F_existsTimeLineHistory_3), v5)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_existsTimeLineHistory_4), int32(252), int32(_a_F_existsTimeLineHistory_5))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v56 = int32(0)
							m.G0 = v5 + int32(1136)
							return v56
						}
					}
				}
			}
		}
	} else {
		v56 = int32(0)
		m.G0 = v5 + int32(1136)
		return v56
	}
}
func F_expand_partitioned_rtentry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int64
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v455 int32
	_ = v455
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v520 int32
	_ = v520
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	F_check_stack_depth(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+112))
	v28 = F_PartitionDirectoryLookup(m, v27, l4)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L138
	}
L4:
	;
	m.G0 = v22 + int32(16)
	return
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v30 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = m.G0
	v35 = v33 + int32(-64)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	if v37 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v35 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+280)) = v115
	v123 = int64(0)
	if v115 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L8:
	;
	v115 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_expand_partitioned_rtentry[0])))
	if v42 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+60)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v35)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = l1
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
	if v62 == int32(-1) {
		v70 = v45
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v45 != 0 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v47 = int32(0)
	v51 = F_bms_add_range(m, v47, v47, v37-int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v115 = v51
	goto L7
L17:
	;
	v74 = F_gen_partprune_steps_internal(m, v33+int32(-20), v70)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+272))
	if v65 == int32(0) {
		v70 = v45
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v68 = F_list_concat_copy(m, v45, v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v70 = v68
	goto L17
L21:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+59)))
	if v76 != 0 {
		v115 = v53
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	if v77 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v80 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	v85 = F_bms_add_range(m, v80, v80, v82-int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v88)
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v98
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v103 = F_palloc0_mul(m, int32(28), v90*v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	v115 = v85
	goto L7
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
	*(*int64)(unsafe.Add(mBase, uint32(v35)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+40)) = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_expand_partitioned_rtentry[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v111
	v113 = F_get_matching_partitions(m, v35, v77)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v115 = v113
	goto L7
L29:
	;
	if int32(0) < v167 {
		goto L44
	} else {
		goto L45
	}
L30:
	;
	v167 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v128 = v115 + int32(8)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v129 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v167 = base.I32_popcnt(v132)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v135 = v129 << (uint(int32(2)) % 32)
	if v135 <= int32(7) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v167 = base.I32_wrap_i64(v162)
	goto L29
L37:
	;
	if v135 == int32(0) {
		v162 = v123
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v159 = F_pg_popcount_optimized(m, v128, v135)
	mBase = m.M
	v162 = v159
	goto L36
L40:
	;
	v140 = v135
	v141 = v128
	v142 = v123
	goto L41
L41:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+3)))
	v144 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_expand_partitioned_rtentry[2]))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+2)))
	v146 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v145)+uint32(_c_F_expand_partitioned_rtentry[2]))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	v148 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_expand_partitioned_rtentry[2]))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v150 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_expand_partitioned_rtentry[2]))))
	v154 = v144 + (v146 + (v148 + (v142 + v150)))
	v155 = int32(4)
	v158 = v140 - v155
	if v158 != 0 {
		v140 = v158
		v141 = v141 + v155
		v142 = v154
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v162 = v154
	goto L36
L43:
	;
	goto L42
L44:
	;
	F_expand_planner_arrays(m, l0, v167)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	v175 = F_palloc0(m, v172<<(uint(int32(2))%32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+276)) = v175
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+280))
	if v178 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v235 < int32(0) {
		goto L4
	} else {
		goto L60
	}
L50:
	;
	v235 = base.I32_ctz(v221) | v222<<(uint(int32(5))%32)
	goto L49
L51:
	;
	v235 = int32(-2)
	goto L49
L52:
	;
	v186 = int32(0)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if v189 <= v186 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v192 = v178 + int32(8)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v199 = v196 & int32(-1)
	if v199 != 0 {
		v221 = v199
		v222 = v186
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v200 = int32(1)
	if v200 == v189 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v204 = v200
	goto L56
L56:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v192+v204<<(uint(int32(2))%32))))
	if v211 != 0 {
		v221 = v211
		v222 = v204
		goto L50
	} else {
		goto L58
	}
L57:
	;
	goto L51
L58:
	;
	v213 = v204 + int32(1)
	if v213 != v189 {
		v204 = v213
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v248 = v235
	goto L61
L61:
	;
	v258 = v248 << (uint(int32(2)) % 32)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258+v259)))
	v262 = F_try_table_open(m, v261, l7)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L4
L63:
	;
	if v455 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L64:
	;
	if v262 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+280))
	v267 = F_bms_del_member(m, v266, v248)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v262)+48))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+118)))
	if v271 == int32(116) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+280)) = v267
	v455 = v267
	goto L63
L69:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+24)))
	if v274 == int32(0) {
		goto L3
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_expand_single_inheritance_child(m, l0, l2, l3, l4, l6, v262, v22+int32(12), v22+int32(8))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v284 = F_build_simple_rel(m, l0, v283, l1)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	*(*int32)(unsafe.Add(mBase, uint32(v286+v258))) = v284
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+284))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	v291 = F_bms_add_members(m, v289, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+284)) = v291
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v262)+48))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+119)))
	if v295 == int32(112) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298+v283<<(uint(int32(2))%32))))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+20))
	v304 = int32(0)
	v307 = F_bms_is_member(m, int32(1), l5)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_relation_close(m, v262, int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L125
	}
L79:
	;
	if v307 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v311 = F_bms_add_member(m, int32(0), int32(1))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	v313 = v304
	goto L82
L82:
	;
	v315 = F_bms_is_member(m, int32(2), l5)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	v313 = v311
	goto L82
L84:
	;
	if v315 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v318 = F_bms_add_member(m, v313, int32(2))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	v320 = v313
	goto L87
L87:
	;
	v322 = F_bms_is_member(m, int32(3), l5)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	v320 = v318
	goto L87
L89:
	;
	if v322 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v325 = F_bms_add_member(m, v320, int32(3))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	v327 = v320
	goto L92
L92:
	;
	v329 = F_bms_is_member(m, int32(4), l5)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v327 = v325
	goto L92
L94:
	;
	if v329 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v332 = F_bms_add_member(m, v327, int32(4))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	v334 = v327
	goto L97
L97:
	;
	v336 = F_bms_is_member(m, int32(5), l5)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L99
	}
L98:
	;
	v334 = v332
	goto L97
L99:
	;
	if v336 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v339 = F_bms_add_member(m, v334, int32(5))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	v341 = v334
	goto L102
L102:
	;
	v343 = F_bms_is_member(m, int32(6), l5)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	v341 = v339
	goto L102
L104:
	;
	if v343 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v346 = F_bms_add_member(m, v341, int32(6))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v348 = v341
	goto L107
L107:
	;
	v350 = F_bms_is_member(m, int32(7), l5)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	v348 = v346
	goto L107
L109:
	;
	if v303 == int32(0) {
		v409 = v348
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	F_expand_partitioned_rtentry(m, l0, v284, v420, v283, v262, v409, l6, l7)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L124
	}
L111:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v354 <= int32(0) {
		v409 = v348
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v365 = v348
	v370 = v304
	goto L113
L113:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v303)+12))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376+v370<<(uint(int32(2))%32))))
	if v380 == int32(0) {
		v396 = v365
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v409 = v396
	goto L110
L115:
	;
	v398 = v370 + int32(1)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v398 < v399 {
		v365 = v396
		v370 = v398
		goto L113
	} else {
		goto L123
	}
L116:
	;
	if v350 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v387 = F_bms_is_member(m, v370+int32(8), l5)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v380)+8)))
	v394 = F_bms_add_member(m, v365, v391+int32(7))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	if v387 == int32(0) {
		v396 = v365
		goto L115
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v396 = v394
	goto L115
L123:
	;
	goto L114
L124:
	;
	goto L78
L125:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+280))
	v455 = v445
	goto L63
L126:
	;
	if int32(0) <= v520 {
		v248 = v520
		goto L61
	} else {
		goto L137
	}
L127:
	;
	v520 = base.I32_ctz(v506) | v507<<(uint(int32(5))%32)
	goto L126
L128:
	;
	v520 = int32(-2)
	goto L126
L129:
	;
	v471 = v248 + int32(1)
	v473 = int32(base.Ui32(v471) >> (uint(int32(5)) % 32))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	if v474 <= v473 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v477 = v455 + int32(8)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v477+v473<<(uint(int32(2))%32))))
	v484 = v481 & (int32(-1) << (uint(v471) % 32))
	if v484 != 0 {
		v506 = v484
		v507 = v473
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v486 = v473 + int32(1)
	if v486 == v474 {
		goto L128
	} else {
		goto L132
	}
L132:
	;
	v489 = v486
	goto L133
L133:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v477+v489<<(uint(int32(2))%32))))
	if v496 != 0 {
		v506 = v496
		v507 = v489
		goto L127
	} else {
		goto L135
	}
L134:
	;
	goto L128
L135:
	;
	v498 = v489 + int32(1)
	if v498 != v474 {
		v489 = v498
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	goto L62
L138:
	;
	F_errmsg_internal(m, int32(_a_F_expand_partitioned_rtentry_0), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_expand_partitioned_rtentry_1), int32(398), int32(_a_F_expand_partitioned_rtentry_2))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_expm1(m *base.Module, l0 float64) float64 {
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v99 float64
	_ = v99
	var v109 float64
	_ = v109
	var v126 float64
	_ = v126
	var v136 float64
	_ = v136
	var v141 float64
	_ = v141
	var v148 float64
	_ = v148
	var v152 float64
	_ = v152
	var v158 float64
	_ = v158
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	v8 = base.I64_reinterpret_f64(l0)
	v13 = base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1078159482)) <= base.Ui32(v13) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
			v170 = l0
			return v170
		} else {
			if v8 < int64(0) {
				return float64(-1)
			} else {
				if base.F64_gt(l0, float64(709.782712893384)) == int32(0) {
					v51 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(l0, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), l0)))
					v52 = base.F64_convert_i32_s(v51)
					v58 = v51
					v59 = base.F64_mul(v52, float64(1.9082149292705877e-10))
					v61 = base.F64_add(l0, base.F64_mul(v52, float64(-0.6931471803691238)))
					v62 = base.F64_sub(v61, v59)
					v68 = v62
					v69 = v58
					v70 = base.F64_sub(base.F64_sub(v61, v62), v59)
					v73 = base.F64_mul(v68, float64(0.5))
					v74 = base.F64_mul(v68, v73)
					v90 = base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v93 = base.F64_sub(float64(3), base.F64_mul(v90, v73))
					v99 = base.F64_mul(v74, base.F64_div(base.F64_sub(v90, v93), base.F64_sub(float64(6), base.F64_mul(v68, v93))))
					if v69 == int32(0) {
						return base.F64_sub(v68, base.F64_sub(base.F64_mul(v68, v99), v74))
					} else {
						v109 = base.F64_sub(base.F64_sub(base.F64_mul(v68, base.F64_sub(v99, v70)), v70), v74)
						switch v69 + int32(1) {
						case 0:
							return base.F64_add(base.F64_mul(base.F64_sub(v68, v109), float64(0.5)), float64(-0.5))
						default:
							v136 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v69+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v69) {
								v141 = base.F64_add(base.F64_sub(v68, v109), float64(1))
								if v69 == int32(1024) {
									v148 = base.F64_mul(base.F64_add(v141, v141), float64(8.98846567431158e+307))
								} else {
									v148 = base.F64_mul(v141, v136)
								}
								return base.F64_add(v148, float64(-1))
							} else {
								v152 = float64(1)
								v158 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v69) << (uint(int64(52)) % 64))
								if base.Ui32(v69) <= base.Ui32(int32(19)) {
									v168 = base.F64_add(base.F64_sub(v152, v158), base.F64_sub(v68, v109))
								} else {
									v168 = base.F64_add(base.F64_sub(v68, base.F64_add(v109, v158)), v152)
								}
								v170 = base.F64_mul(v168, v136)
								return v170
							}
						case 2:
							if base.F64_lt(v68, float64(-0.25)) != 0 {
								return base.F64_mul(base.F64_sub(v109, base.F64_add(v68, float64(0.5))), float64(-2))
							} else {
								v126 = base.F64_sub(v68, v109)
								return base.F64_add(base.F64_add(v126, v126), float64(1))
							}
						}
					}
				} else {
					return base.F64_mul(l0, float64(8.98846567431158e+307))
				}
			}
		}
	} else {
		if base.Ui32(v13) < base.Ui32(int32(1071001155)) {
			if base.Ui32(v13) < base.Ui32(int32(1016070144)) {
				v170 = l0
				return v170
			} else {
				v68 = l0
				v69 = int32(0)
				v70 = float64(0)
				v73 = base.F64_mul(v68, float64(0.5))
				v74 = base.F64_mul(v68, v73)
				v90 = base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
				v93 = base.F64_sub(float64(3), base.F64_mul(v90, v73))
				v99 = base.F64_mul(v74, base.F64_div(base.F64_sub(v90, v93), base.F64_sub(float64(6), base.F64_mul(v68, v93))))
				if v69 == int32(0) {
					return base.F64_sub(v68, base.F64_sub(base.F64_mul(v68, v99), v74))
				} else {
					v109 = base.F64_sub(base.F64_sub(base.F64_mul(v68, base.F64_sub(v99, v70)), v70), v74)
					switch v69 + int32(1) {
					case 0:
						return base.F64_add(base.F64_mul(base.F64_sub(v68, v109), float64(0.5)), float64(-0.5))
					default:
						v136 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v69+int32(1023)) << (uint(int64(52)) % 64))
						if base.Ui32(int32(57)) <= base.Ui32(v69) {
							v141 = base.F64_add(base.F64_sub(v68, v109), float64(1))
							if v69 == int32(1024) {
								v148 = base.F64_mul(base.F64_add(v141, v141), float64(8.98846567431158e+307))
							} else {
								v148 = base.F64_mul(v141, v136)
							}
							return base.F64_add(v148, float64(-1))
						} else {
							v152 = float64(1)
							v158 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v69) << (uint(int64(52)) % 64))
							if base.Ui32(v69) <= base.Ui32(int32(19)) {
								v168 = base.F64_add(base.F64_sub(v152, v158), base.F64_sub(v68, v109))
							} else {
								v168 = base.F64_add(base.F64_sub(v68, base.F64_add(v109, v158)), v152)
							}
							v170 = base.F64_mul(v168, v136)
							return v170
						}
					case 2:
						if base.F64_lt(v68, float64(-0.25)) != 0 {
							return base.F64_mul(base.F64_sub(v109, base.F64_add(v68, float64(0.5))), float64(-2))
						} else {
							v126 = base.F64_sub(v68, v109)
							return base.F64_add(base.F64_add(v126, v126), float64(1))
						}
					}
				}
			}
		} else {
			if base.Ui32(int32(1072734897)) < base.Ui32(v13) {
				v51 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(l0, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), l0)))
				v52 = base.F64_convert_i32_s(v51)
				v58 = v51
				v59 = base.F64_mul(v52, float64(1.9082149292705877e-10))
				v61 = base.F64_add(l0, base.F64_mul(v52, float64(-0.6931471803691238)))
			} else {
				if int64(0) <= v8 {
					v58 = int32(1)
					v59 = float64(1.9082149292705877e-10)
					v61 = base.F64_add(l0, float64(-0.6931471803691238))
				} else {
					v58 = int32(-1)
					v59 = float64(-1.9082149292705877e-10)
					v61 = base.F64_add(l0, float64(0.6931471803691238))
				}
			}
			v62 = base.F64_sub(v61, v59)
			v68 = v62
			v69 = v58
			v70 = base.F64_sub(base.F64_sub(v61, v62), v59)
			v73 = base.F64_mul(v68, float64(0.5))
			v74 = base.F64_mul(v68, v73)
			v90 = base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, base.F64_add(base.F64_mul(v74, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
			v93 = base.F64_sub(float64(3), base.F64_mul(v90, v73))
			v99 = base.F64_mul(v74, base.F64_div(base.F64_sub(v90, v93), base.F64_sub(float64(6), base.F64_mul(v68, v93))))
			if v69 == int32(0) {
				return base.F64_sub(v68, base.F64_sub(base.F64_mul(v68, v99), v74))
			} else {
				v109 = base.F64_sub(base.F64_sub(base.F64_mul(v68, base.F64_sub(v99, v70)), v70), v74)
				switch v69 + int32(1) {
				case 0:
					return base.F64_add(base.F64_mul(base.F64_sub(v68, v109), float64(0.5)), float64(-0.5))
				default:
					v136 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v69+int32(1023)) << (uint(int64(52)) % 64))
					if base.Ui32(int32(57)) <= base.Ui32(v69) {
						v141 = base.F64_add(base.F64_sub(v68, v109), float64(1))
						if v69 == int32(1024) {
							v148 = base.F64_mul(base.F64_add(v141, v141), float64(8.98846567431158e+307))
						} else {
							v148 = base.F64_mul(v141, v136)
						}
						return base.F64_add(v148, float64(-1))
					} else {
						v152 = float64(1)
						v158 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v69) << (uint(int64(52)) % 64))
						if base.Ui32(v69) <= base.Ui32(int32(19)) {
							v168 = base.F64_add(base.F64_sub(v152, v158), base.F64_sub(v68, v109))
						} else {
							v168 = base.F64_add(base.F64_sub(v68, base.F64_add(v109, v158)), v152)
						}
						v170 = base.F64_mul(v168, v136)
						return v170
					}
				case 2:
					if base.F64_lt(v68, float64(-0.25)) != 0 {
						return base.F64_mul(base.F64_sub(v109, base.F64_add(v68, float64(0.5))), float64(-2))
					} else {
						v126 = base.F64_sub(v68, v109)
						return base.F64_add(base.F64_add(v126, v126), float64(1))
					}
				}
			}
		}
	}
}
func F_extractModify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_defGetString(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L29
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return v100
L3:
	;
	return int32(0)
L4:
	;
	v13 = int32(_a_F_extractModify_0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_extractModify[0])))
	if base.B2i32(v16 == int32(0))|base.B2i32(v16 != v19) != 0 {
		v37 = v16
		v38 = v19
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v37-v38 == int32(0) {
		v100 = int32(114)
		goto L2
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v22 = v9
	v23 = v13
	goto L8
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v27
		v38 = v26
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v37 = v27
	v38 = v26
	goto L6
L10:
	;
	v30 = int32(1)
	if v27 == v26 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v43 = int32(_a_F_extractModify_1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_extractModify[1])))
	if base.B2i32(v46 == int32(0))|base.B2i32(v46 != v49) != 0 {
		v67 = v46
		v68 = v49
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v67-v68 == int32(0) {
		v100 = int32(115)
		goto L2
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v52 = v9
	v53 = v43
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v57
		v68 = v56
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v67 = v57
	v68 = v56
	goto L14
L18:
	;
	v60 = int32(1)
	if v57 == v56 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v72 = int32(_a_F_extractModify_2)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_extractModify[2])))
	if base.B2i32(v75 == int32(0))|base.B2i32(v75 != v78) != 0 {
		v96 = v75
		v97 = v78
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v96-v97 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v81 = v9
	v82 = v72
	goto L24
L24:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v86 == int32(0) {
		v96 = v86
		v97 = v85
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v86
	v97 = v85
	goto L22
L26:
	;
	v89 = int32(1)
	if v86 == v85 {
		v81 = v81 + v89
		v82 = v82 + v89
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v100 = int32(119)
	goto L2
L29:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v112
	F_errmsg(m, int32(_a_F_extractModify_3), v6)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_extractModify_4), int32(495), int32(_a_F_extractModify_5))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_extract_actual_clauses(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v3
	v17 = v3
	goto L7
L5:
	;
	v47 = v3
	goto L6
L6:
	;
	return v47
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v16<<(uint(int32(2))%32))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
	if v23 != l1 {
		v38 = v17
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v47 = v38
	goto L6
L9:
	;
	v40 = v16 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v40 < v41 {
		v16 = v40
		v17 = v38
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 != int32(7) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = F_lappend(m, v17, v25)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)))
	if v29 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	if v30 != int64(0) {
		v38 = v17
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	return int32(0)
L16:
	;
	v38 = v33
	goto L9
L17:
	;
	goto L8
}
func F_extract_or_clause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v10 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if int32(0) < v15 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v247
L5:
	;
	v22 = v3
	v25 = v3
	goto L8
L6:
	;
	v236 = v3
	goto L7
L7:
	;
	v240 = int32(0)
	if v236 == v240 {
		v247 = v240
		goto L4
	} else {
		goto L75
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v25<<(uint(int32(2))%32))))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v236 = v227
	goto L7
L10:
	;
	if v203 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L11:
	;
	v136 = int32(0)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+10)))
	if v137 != 0 {
		v247 = v136
		goto L4
	} else {
		goto L47
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v33 != int32(21) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v36 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v37 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L17
L17:
	;
	v42 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v44 <= v42 {
		v203 = v42
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v47 = v42
	v50 = v42
	goto L19
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v47<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	goto L23
L20:
	;
	v203 = v131
	goto L10
L21:
	;
	v133 = v47 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v133 < v134 {
		v47 = v133
		v50 = v131
		goto L19
	} else {
		goto L46
	}
L22:
	;
	v128 = F_lappend(m, v50, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L27
	} else {
		goto L45
	}
L23:
	;
	if v60 != int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v63 = F_extract_or_clause(m, v59, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+10)))
	if v67 != 0 {
		v131 = v50
		goto L21
	} else {
		goto L30
	}
L27:
	;
	return int32(0)
L28:
	;
	if v63 != 0 {
		v127 = v63
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v131 = v50
	goto L21
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v70 = int32(0)
	if base.B2i32(v68 == v70)|base.B2i32(v69 == v70) != 0 {
		v116 = base.B2i32(v68|v69 == v70)
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v116 == int32(0) {
		v131 = v50
		goto L21
	} else {
		goto L42
	}
L32:
	;
	goto L31
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v84 != v85 {
		v116 = int32(0)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v87 = int32(1)
	if v84 <= v87 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v90 = v87
	goto L37
L36:
	;
	v90 = v84
	goto L37
L37:
	;
	v91 = int32(8)
	v96 = int32(0)
	goto L38
L38:
	;
	v104 = v96 << (uint(int32(2)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v68+v91+v104)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v69+v91+v104)))
	v109 = base.B2i32(v106 == v108)
	if v106 != v108 {
		v116 = v109
		goto L32
	} else {
		goto L40
	}
L39:
	;
	v116 = v109
	goto L32
L40:
	;
	v112 = v96 + int32(1)
	if v112 != v90 {
		v96 = v112
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v124 = F_contain_volatile_functions(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	if v124 != 0 {
		v131 = v50
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v127 = v126
	goto L22
L45:
	;
	v131 = v128
	goto L21
L46:
	;
	goto L20
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v140 = int32(0)
	if base.B2i32(v138 == v140)|base.B2i32(v139 == v140) != 0 {
		v186 = base.B2i32(v138|v139 == v140)
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v186 == int32(0) {
		v247 = v136
		goto L4
	} else {
		goto L59
	}
L49:
	;
	goto L48
L50:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v154 != v155 {
		v186 = int32(0)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v157 = int32(1)
	if v154 <= v157 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v160 = v157
	goto L54
L53:
	;
	v160 = v154
	goto L54
L54:
	;
	v161 = int32(8)
	v166 = int32(0)
	goto L55
L55:
	;
	v174 = v166 << (uint(int32(2)) % 32)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v138+v161+v174)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v139+v161+v174)))
	v179 = base.B2i32(v176 == v178)
	if v176 != v178 {
		v186 = v179
		goto L49
	} else {
		goto L57
	}
L56:
	;
	v186 = v179
	goto L49
L57:
	;
	v182 = v166 + int32(1)
	if v182 != v160 {
		v166 = v182
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v194 = F_contain_volatile_functions(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L27
	} else {
		goto L60
	}
L60:
	;
	if v194 != 0 {
		v247 = v136
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v198 = F_lappend(m, int32(0), v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L27
	} else {
		goto L62
	}
L62:
	;
	v203 = v198
	goto L10
L63:
	;
	return int32(0)
L64:
	;
	goto L65
L65:
	;
	v212 = F_make_ands_explicit(m, v203)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L27
	} else {
		goto L68
	}
L66:
	;
	v229 = v25 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v229 < v230 {
		v22 = v227
		v25 = v229
		goto L8
	} else {
		goto L74
	}
L67:
	;
	v225 = F_lappend(m, v22, v212)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L27
	} else {
		goto L73
	}
L68:
	;
	if v212 == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v216 != int32(21) {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v219 != int32(1) {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v223 = F_list_concat(m, v22, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L27
	} else {
		goto L72
	}
L72:
	;
	v227 = v223
	goto L66
L73:
	;
	v227 = v225
	goto L66
L74:
	;
	goto L9
L75:
	;
	v243 = F_make_orclause(m, v236)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L27
	} else {
		goto L76
	}
L76:
	;
	v247 = v243
	goto L4
}
func F_extract_variadic_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == v5 {
		v31 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v34
	if v31&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	goto L1
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v23 == int32(0) {
		v31 = v5
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v26 != int32(15) {
		v31 = v5
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+13)))
	v31 = v29
	goto L2
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L14
	} else {
		goto L50
	}
L7:
	;
	m.G0 = v17 + int32(32)
	return v255
L8:
	;
	v255 = int32(-1)
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v241
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v238
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v240
	v255 = v237
	goto L7
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v40 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v114 = F_palloc0(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L25
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v42 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	F_get_typlenbyvalalign(m, v46, v17+int32(16), v17+int32(19), v17+int32(18))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+16)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+19)))
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+18)))
	F_deconstruct_array(m, v42, v55, v56, v57, v17+int32(28), v17+int32(24), v17+int32(20))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v69 = F_palloc0(m, v66<<(uint(int32(2))%32))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if int32(0) < v71 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v80 = int32(0)
	goto L22
L20:
	;
	v103 = v71
	goto L21
L21:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v237 = v103
	v238 = v111
	v240 = v69
	v241 = v112
	goto L9
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69+v80<<(uint(int32(2))%32)))) = v46
	v94 = v80 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v94 < v95 {
		v80 = v94
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v103 = v95
	goto L21
L24:
	;
	goto L23
L25:
	;
	v118 = F_palloc0(m, v113<<(uint(int32(3))%32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v122 = F_palloc0(m, v113<<(uint(int32(2))%32))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	if v113 <= int32(0) {
		v237 = v113
		v238 = v114
		v240 = v122
		v241 = v118
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v134 = int32(0)
	goto L29
L29:
	;
	v146 = l0 + int32(24) + v134<<(uint(int32(4))%32)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v134+v114))) = uint8(v147)
	v151 = v122 + v134<<(uint(int32(2))%32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v153 = F_get_fn_expr_argtype(m, v152, v134)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L14
	} else {
		goto L31
	}
L30:
	;
	v237 = v113
	v238 = v114
	v240 = v122
	v241 = v118
	goto L9
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v153
	if v153 != int32(705) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118+v134<<(uint(int32(3))%32)))) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if base.B2i32(v222 == int32(0))|base.B2i32(v222 == int32(705)) != 0 {
		goto L6
	} else {
		goto L48
	}
L33:
	;
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
	v220 = v219
	goto L32
L34:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v162 = int32(0)
	if v161 == v162 {
		v208 = v162
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v208 == int32(0) {
		goto L33
	} else {
		goto L45
	}
L36:
	;
	goto L35
L37:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	if v166 == int32(0) {
		v208 = v162
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v171 = v169 - int32(11)
	v178 = int32(0)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v171))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v171)%32))&int32(1) == v178)|base.B2i32(v134 < v178) != 0 {
		v208 = v162
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v171<<(uint(int32(2))%32))+uint32(_c_F_extract_variadic_args[0])))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v166+v186)))
	if v188 == int32(0) {
		v208 = v162
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v191 <= v134 {
		v208 = v162
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v193 = int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194+v134<<(uint(int32(2))%32))))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	switch v199 - int32(7) {
	case 0:
		v208 = v193
		goto L36
	case 1:
		goto L43
	default:
		goto L42
	}
L42:
	;
	v208 = int32(0)
	goto L36
L43:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v202 == int32(0) {
		v208 = v193
		goto L36
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(25)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+8)))
	if v214 != 0 {
		v220 = int64(0)
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v216 = F_cstring_to_text(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	v220 = base.I64_extend_i32_u(v216)
	goto L32
L48:
	;
	v229 = v134 + int32(1)
	if v229 < v113 {
		v134 = v229
		goto L29
	} else {
		goto L49
	}
L49:
	;
	goto L30
L50:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v134 + int32(1)
	F_errmsg(m, int32(_a_F_extract_variadic_args_0), v17)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_extract_variadic_args_1), int32(2097), int32(_a_F_extract_variadic_args_2))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
