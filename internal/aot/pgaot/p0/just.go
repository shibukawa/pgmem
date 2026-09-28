package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecJustConst(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v5)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v4)+16))
	return v7
}
func F_ExecJustHashOuterVarStrict(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+100))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	if v13 < v11 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		m.T0[v16].(func(*base.Module, int32, int32))(m, v12, v11)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v22 = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v23+v9<<(uint(int32(3))%32))))
			*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v27
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v9))))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(v32)
			if v32 == int32(0) {
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
				v39 = m.T0[v38].(func(*base.Module, int32) int64)(m, v10)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					return v39
				}
			} else {
				v42 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v42)
				return int64(0)
			}
		}
	} else {
		v22 = v12
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v23+v9<<(uint(int32(3))%32))))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v27
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v9))))
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(v32)
		if v32 == int32(0) {
			v36 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
			v39 = m.T0[v38].(func(*base.Module, int32) int64)(m, v10)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				return v39
			}
		} else {
			v42 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v42)
			return int64(0)
		}
	}
}
func F_ExecJustInnerVar(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+56))
	v8 = v6 + int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+6)))
	if v10 < v8 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		m.T0[v13].(func(*base.Module, int32, int32))(m, v9, v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v6))))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v20)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v22+v6<<(uint(int32(3))%32))))
			return v26
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v6))))
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v20)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v22+v6<<(uint(int32(3))%32))))
		return v26
	}
}
func F_ExecJustScanVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v15 int64
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v7))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v9)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v11+v5<<(uint(int32(3))%32))))
	return v15
}
