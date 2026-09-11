package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecJustApplyFuncToCase(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+68))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)) = uint8(v48)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v7-int32(-64))))
	v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v16)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v23 = int32(0)
	goto L3
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(24)+v23<<(uint(int32(3))%32)))))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v38)
	return int32(0)
L5:
	;
	v36 = v23 + int32(1)
	if v17 != v36 {
		v23 = v36
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v57)
	return v53
}
func F_ExecJustAssignOuterVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	v14 = v12 + int32(1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+6)))
	if v16 < v14 {
		F_slot_getsomeattrs_int(m, v15, v14)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v12))))
			*(*uint8)(unsafe.Add(mBase, uint32(v9+v11))) = uint8(v25)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v28 = int32(2)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v12<<(uint(v28)%32))))
			*(*int32)(unsafe.Add(mBase, uint32(v27+v9<<(uint(v28)%32)))) = v35
			return int32(0)
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v12))))
		*(*uint8)(unsafe.Add(mBase, uint32(v9+v11))) = uint8(v25)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
		v28 = int32(2)
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v12<<(uint(v28)%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v27+v9<<(uint(v28)%32)))) = v35
		return int32(0)
	}
}
