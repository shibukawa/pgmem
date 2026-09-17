package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecJustAssignInnerVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = int32(2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14<<(uint(v9)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v6+v8<<(uint(v9)%32)))) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v22))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8+v20))) = uint8(v24)
	return int32(0)
}
func F_ExecJustAssignScanVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
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
func F_ExecJustHashInnerVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v12 int32
	_ = v12
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+100))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+6)))
	if v12 < v10 {
		F_slot_getsomeattrs_int(m, v11, v10)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v8<<(uint(int32(2))%32))))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v8))))
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v26)
			v28 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v28)
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
			if v30 != 0 {
				v35 = int32(0)
				return v35
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+104))
				v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, v9)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = v33
					return v35
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v8<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v8))))
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v26)
		v28 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v28)
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
		if v30 != 0 {
			v35 = int32(0)
			return v35
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+104))
			v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, v9)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = v33
				return v35
			}
		}
	}
}
func F_ExecJustHashOuterVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)) = uint8(v18)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v4)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)))
	if v22 != 0 {
		v29 = int32(0)
		return v29
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
		v25 = m.T0[v24].(func(*base.Module, int32) int32)(m, v7)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = v25
			return v29
		}
	}
}
func F_ExecJustOuterVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v7))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v9)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v5<<(uint(int32(2))%32))))
	return v15
}
