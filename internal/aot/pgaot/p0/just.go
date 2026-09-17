package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecJustConst(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v5)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	return v7
}
func F_ExecJustHashOuterVarStrict(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+100))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	if v13 < v11 {
		F_slot_getsomeattrs_int(m, v12, v11)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v20 = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v9<<(uint(int32(2))%32))))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v25
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v9))))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v30)
			if v30 == int32(0) {
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v34)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
				v37 = m.T0[v36].(func(*base.Module, int32) int32)(m, v10)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					return v37
				}
			} else {
				v40 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v40)
				return int32(0)
			}
		}
	} else {
		v20 = v12
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v9<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v9))))
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v30)
		if v30 == int32(0) {
			v34 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v34)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
			v37 = m.T0[v36].(func(*base.Module, int32) int32)(m, v10)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				return v37
			}
		} else {
			v40 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v40)
			return int32(0)
		}
	}
}
func F_ExecJustInnerVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+56))
	v8 = v6 + int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+6)))
	if v10 < v8 {
		F_slot_getsomeattrs_int(m, v9, v8)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v6))))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v18)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v6<<(uint(int32(2))%32))))
			return v24
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v6))))
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v18)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v6<<(uint(int32(2))%32))))
		return v24
	}
}
func F_ExecJustScanVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v7))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v9)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v5<<(uint(int32(2))%32))))
	return v15
}
