package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecJustApplyFuncToCase(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v13)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	if v17 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)) = uint8(v43)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
	v46 = m.T0[v45].(func(*base.Module, int32) int32)(m, v16)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v20 = v4
	goto L3
L3:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v20<<(uint(int32(3))%32))+24)))
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v34)
	return int32(0)
L5:
	;
	v32 = v20 + int32(1)
	if v17 != v32 {
		v20 = v32
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
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v50)
	return v46
}
func F_ExecJustAssignOuterVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v13 = v11 + int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+6)))
	if v15 < v13 {
		F_slot_getsomeattrs_int(m, v14, v13)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v11))))
			*(*uint8)(unsafe.Add(mBase, uint32(v10+v8))) = uint8(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v27 = int32(2)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v11<<(uint(v27)%32))))
			*(*int32)(unsafe.Add(mBase, uint32(v26+v10<<(uint(v27)%32)))) = v34
			return int32(0)
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v11))))
		*(*uint8)(unsafe.Add(mBase, uint32(v10+v8))) = uint8(v24)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
		v27 = int32(2)
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v11<<(uint(v27)%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v26+v10<<(uint(v27)%32)))) = v34
		return int32(0)
	}
}
