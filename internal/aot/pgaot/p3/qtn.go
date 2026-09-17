package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTNEq(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = v7 & v8
	if base.B2i32(v9 != v8)|base.B2i32(v7 != v9) != 0 {
		v18 = int32(1)
		return base.B2i32(v18 == int32(0))
	} else {
		v14 = F_QTNodeCompare(m, l0, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = v14
			return base.B2i32(v18 == int32(0))
		}
	}
}
func F_QTNFree(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	return
L5:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v46 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v22 = v6
	goto L9
L9:
	;
	if v22&int32(255) != int32(2) {
		goto L6
	} else {
		goto L13
	}
L10:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v12&int32(4) == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_pfree(m, v9)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v22 = v20
	goto L9
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v27 <= int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v32 = int32(0)
	goto L15
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v32<<(uint(int32(2))%32))))
	F_QTNFree(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L6
L17:
	;
	v41 = v32 + int32(1)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v41 < v42 {
		v32 = v41
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_pfree(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v49&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pfree(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_pfree(m, l0)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L3
}
func F_QTNSort(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	F_check_stack_depth(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v7 != int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v14 = int32(0)
	goto L6
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v14<<(uint(int32(2))%32))))
	F_QTNSort(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	if v25 < int32(2) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v24 = v14 + int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 < v25 {
		v14 = v24
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v30 == int32(4) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_qsort(m, v33, v25, int32(4), int32(1515))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L3
}
