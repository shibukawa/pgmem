package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecJustAssignInnerVar(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
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
func F_ExecJustHashInnerVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v9+v10<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v18)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v4)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)))
	if v22 != 0 {
		v29 = int64(0)
		return v29
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
		v25 = m.T0[v24].(func(*base.Module, int32) int64)(m, v7)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v29 = v25
			return v29
		}
	}
}
func F_ExecJustHashInnerVarWithIV(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+140))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
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
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v22 = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v23+v9<<(uint(int32(3))%32))))
			*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v27
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v9))))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(v32)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
			v36 = base.I32_rotl(v34, int32(1))
			if v32 == int32(0) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
				v40 = m.T0[v39].(func(*base.Module, int32) int64)(m, v10)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int64(0)
				} else {
					v44 = v36 ^ base.I32_wrap_i64(v40)
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
					return base.I64_extend_i32_u(v44)
				}
			} else {
				v44 = v36
				v45 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
				return base.I64_extend_i32_u(v44)
			}
		}
	} else {
		v22 = v12
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v23+v9<<(uint(int32(3))%32))))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v27
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v9))))
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(v32)
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
		v36 = base.I32_rotl(v34, int32(1))
		if v32 == int32(0) {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
			v40 = m.T0[v39].(func(*base.Module, int32) int64)(m, v10)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				v44 = v36 ^ base.I32_wrap_i64(v40)
				v45 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
				return base.I64_extend_i32_u(v44)
			}
		} else {
			v44 = v36
			v45 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
			return base.I64_extend_i32_u(v44)
		}
	}
}
func F_ExecJustHashOuterVar(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+100))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+6)))
	if v12 < v10 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
		m.T0[v15].(func(*base.Module, int32, int32))(m, v11, v10)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v20+v8<<(uint(int32(3))%32))))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v8))))
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+32)) = uint8(v28)
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v30)
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+32)))
			if v32 != 0 {
				v37 = int64(0)
				return v37
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+104))
				v35 = m.T0[v34].(func(*base.Module, int32) int64)(m, v9)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					v37 = v35
					return v37
				}
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v20+v8<<(uint(int32(3))%32))))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v8))))
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+32)) = uint8(v28)
		v30 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v30)
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+32)))
		if v32 != 0 {
			v37 = int64(0)
			return v37
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+104))
			v35 = m.T0[v34].(func(*base.Module, int32) int64)(m, v9)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				v37 = v35
				return v37
			}
		}
	}
}
func F_ExecJustInnerVarVirt(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v7))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v9)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v11+v5<<(uint(int32(3))%32))))
	return v15
}
func F_ExecJustScanVar(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
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
