package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecJustApplyFuncToCase(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v14)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+68))
	if v18 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)) = uint8(v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v7)+64))
	v49 = m.T0[v48].(func(*base.Module, int32) int64)(m, v17)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v21 = v4
	goto L3
L3:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v21<<(uint(int32(4))%32))+32)))
	if v30 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
	return int64(0)
L5:
	;
	v34 = v21 + int32(1)
	if v18 != v34 {
		v21 = v34
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
	return int64(0)
L10:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v53)
	return v49
}
func F_ExecJustAssignOuterVar(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v13 = v11 + int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+6)))
	if v15 < v13 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
		m.T0[v18].(func(*base.Module, int32, int32))(m, v14, v13)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v11))))
			*(*uint8)(unsafe.Add(mBase, uint32(v10+v8))) = uint8(v26)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v29 = int32(3)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v36 = *(*int64)(unsafe.Add(mBase, uint32(v32+v11<<(uint(v29)%32))))
			*(*int64)(unsafe.Add(mBase, uint32(v28+v10<<(uint(v29)%32)))) = v36
			return int64(0)
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v11))))
		*(*uint8)(unsafe.Add(mBase, uint32(v10+v8))) = uint8(v26)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
		v29 = int32(3)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v32+v11<<(uint(v29)%32))))
		*(*int64)(unsafe.Add(mBase, uint32(v28+v10<<(uint(v29)%32)))) = v36
		return int64(0)
	}
}
