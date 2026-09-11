package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildOnConflictExcludedTargetlist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v57 int32
	_ = v57
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v8 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+120)))
	if v3 < v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v3
	v15 = v3
	goto L4
L2:
	;
	v66 = v7
	v68 = v3
	goto L3
L3:
	;
	v70 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	v75 = F_makeVar(m, l1, v70, v71, int32(-1), v70, v70)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L17
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v24 = v17 + v18<<(uint(int32(4))%32) + v14*int32(100)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+111)))
	if v25 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v66 = v61
	v68 = v59
	goto L3
L6:
	;
	v54 = v14 + int32(1)
	v57 = F_makeTargetEntry(m, v52, base.I32_extend16_s(v54), v51, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L14
	}
L7:
	;
	v28 = int32(0)
	v32 = F_makeNullConst(m, int32(23), int32(-1), v28)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v40 = v24 + int32(20)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+76))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+96))
	v45 = F_makeVar(m, l1, base.I32_extend16_s(v14+int32(1)), v41, v42, v43, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v51 = v28
	v52 = v32
	goto L6
L12:
	;
	v49 = F_pstrdup(m, v24+int32(24))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v51 = v49
	v52 = v45
	goto L6
L14:
	;
	v59 = F_lappend(m, v15, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+120)))
	if v54 < v62 {
		v14 = v54
		v15 = v59
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L5
L17:
	;
	v77 = int32(0)
	v80 = F_makeTargetEntry(m, v75, v77, v77, int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v82 = F_lappend(m, v68, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	return v82
}
func F_on_pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_line_contain_point(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
