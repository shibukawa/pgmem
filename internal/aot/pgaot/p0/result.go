package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecGetResultSlotOps(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)))
	if v4 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
		if v7 != 0 {
			if l1 == int32(0) {
				v42 = v7
				return v42
			} else {
				v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)))
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v10)
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				return v12
			}
		} else {
			if l1 == int32(0) {
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)))
				v30 = v16
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
			}
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if v34 == int32(0) {
				return int32(_a_F_ExecGetResultSlotOps_0)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
				v42 = v39
				return v42
			}
		}
	} else {
		if l1 == int32(0) {
		} else {
			v19 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if v20 == v19 {
				v30 = v19
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
				v30 = int32(base.Ui32(v23)>>(uint(int32(4))%32)) & int32(1)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
		}
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		if v34 == int32(0) {
			return int32(_a_F_ExecGetResultSlotOps_0)
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
			v42 = v39
			return v42
		}
	}
}
func F_ExecInitResultSlot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v10 == int32(0) {
		v28 = v9
		v29 = int32(2)
	} else {
		v14 = int32(7)
		v16 = int32(-8)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v28 = (v9+v14)&v16 + v18<<(uint(int32(3))%32) + (v18+v14)&v16
		v29 = int32(18)
	}
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = F_palloc0(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v10
		*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)) = uint16(v29)
		*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(449)
		*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = l1
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitResultSlot[0]))
		v40 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v31)+6)) = uint16(v40)
		*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v39
		if v10 != 0 {
			v47 = v31 + (v9+int32(7))&int32(-8)
			*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v47 + v49<<(uint(int32(3))%32)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			if int32(0) <= v54 {
				F_IncrTupleDescRefCount(m, v10)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v60 = v59
					*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = int32(0)
					v63 = v60
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
					m.T0[v64].(func(*base.Module, int32))(m, v31)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v30)+104))
						v68 = F_lappend(m, v67, v31)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v30)+104)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v31
							v72 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v72)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(base.B2i32(v75 != int32(0)))
							return
						}
					}
				}
			} else {
				v60 = l1
				*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = int32(0)
				v63 = v60
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
				m.T0[v64].(func(*base.Module, int32))(m, v31)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v30)+104))
					v68 = F_lappend(m, v67, v31)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v30)+104)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v31
						v72 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v72)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(base.B2i32(v75 != int32(0)))
						return
					}
				}
			}
		} else {
			v63 = l1
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			m.T0[v64].(func(*base.Module, int32))(m, v31)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v30)+104))
				v68 = F_lappend(m, v67, v31)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+104)) = v68
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v31
					v72 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v72)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(base.B2i32(v75 != int32(0)))
					return
				}
			}
		}
	}
}
func F_get_call_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = F_internal_get_result_type(m, v6, v7, v8, l1, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_wait_result_is_any_signal(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	v10 = int32(255)
	return base.B2i32(l0&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(l0)>>(uint(int32(8))%32))&v10)) | base.B2i32(base.Ui32(l0&int32(_a_F_wait_result_is_any_signal_0)-int32(1)) < base.Ui32(v10))
}
